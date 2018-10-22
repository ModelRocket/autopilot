package prometheus

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// Cluster stats
	Cluster = "Cluster"
	// Node stats
	Node = "Node"
	// Disk stats
	Disk = "Disk"
	// Volume stats
	Volume = "Volume"
	// Pool stats
	Pool = "Pool"
	// Proc stats - on the PX processes
	Proc = "Proc"
	// Alerts metrics are a special case for outputs
	Alerts = "ALERTS"
	// Instance field in the CSV
	Instance = "instance"
	// Timestamp field in the CSV
	Timestamp = "timestamp"
)

type (
	// Metric is the metric that comes from Prometheus apiproxy
	Metric struct {
		Name          string  `json:"__name__"`
		Cluster       string  `json:"cluster"`
		Instance      string  `json:"instance"`
		Node          string  `json:"node_id"`
		NodeName      string  `json:"node"`
		Job           *string `json:"job,omitempty"`
		Volume        *string `json:"volumeid,omitempty"`
		VolumeName    *string `json:"volumename,omitempty"`
		VolumePVC     *string `json:"volumepvc,omitempty"`
		Disk          *string `json:"disk,omitempty"`
		Pool          *string `json:"pool,omitempty"`
		AlertName     *string `json:"alertname,omitempty"`
		AlertState    *string `json:"alertstate,omitempty"`
		AlertSeverity *string `json:"severity,omitempty"`
		AlertIssue    *string `json:"issue,omitempty"`
		Proc          *string `json:"proc,omitempty"`
	}

	// Result for a single metric and value
	Result struct {
		Metric Metric `json:"metric"`
		// Value is [timestamp, value] such as "value": [1231231233.232, "0"], its mixed type
		Value []interface{} `json:"value,omitempty"`
		// Values is for range queries its an array of Value above
		Values [][]interface{} `json:"values,omitempty"`
	}

	// ClusterResults is the complete set of metrics for the cluster
	ClusterResults struct {
		Status string `json:"status"`
		Data   struct {
			ResultType string   `json:"resultType"`
			Results    []Result `json:"result"`
		} `json:"data"`
	}

	// Structures for Underlying Representation for ML

	// CSVRow is the common characteristics of every metric - this defines a row in the CSV
	CSVRow struct {
		Timestamp uint32
		Cluster   string
		Instance  string
		Node      string
	}

	// CSVMetrics contain ALL the BaseAttributes for a row in the CSV Each map is keyed by the Name of the field (VolumeID, Disk, Pool, Proc)
	// This probably could contain a map of maps of maps, but its becoming unreadable with that level of nesting.
	// Cluster and Proc metrics each are node specific, so they will be located in the Node map
	CSVMetrics struct {
		// Node are the node based metrics keyed on node field
		Node map[string]map[string]string
		// Volume are the volume based metrics keyed on Volume field
		Volume map[string]map[string]string
		// Disk are the disk based metrics keyed on Disk field
		Disk map[string]map[string]string
		// Pool are the pool based metrics keyed on Pool field
		Pool map[string]map[string]string
	}

	// AlertRow is a special type of metric that we use for the Output of the ML.  This is one of the predictors
	AlertRow struct {
		csvRow        CSVRow
		AlertName     string
		AlertState    string
		AlertSeverity string
		AlertIssue    string
		AlertValue    string
	}

	// Prometheus defines the prometheus provider
	Prometheus struct {
		URL string
		Log *logrus.Logger
	}
)

// NewCSVMetrics constructor to initialize internal maps
func NewCSVMetrics() *CSVMetrics {
	return &CSVMetrics{
		Node:   make(map[string]map[string]string),
		Volume: make(map[string]map[string]string),
		Disk:   make(map[string]map[string]string),
		Pool:   make(map[string]map[string]string),
	}
}

/* JNT - Im not sure how to do this, what I want to do is handle the goofy structure here that has Value or Values,
 * if its Value I want to make it Values: [Value]
func (cr *ClusterResults) UnmarsalJSON(b []byte) error {
	var f interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}
	m := f.(map[string]interface{})

	cr.Status =

}
*/

// NewAlertRow Create a new AlertRow
func NewAlertRow(csvRow CSVRow, result *Result, value string) *AlertRow {
	return &AlertRow{
		csvRow:        csvRow,
		AlertName:     *result.Metric.AlertName,
		AlertState:    *result.Metric.AlertState,
		AlertSeverity: *result.Metric.AlertSeverity,
		AlertIssue:    *result.Metric.AlertIssue,
		AlertValue:    value,
	}
}

// Keys Return keys of the given map - used to pull the values from a particular VolumeID, Disk, Pool, Proc of a BaseMetric
func Keys(m map[string]string) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Collect implements the provider interface
func (p *Prometheus) Collect(params map[string]string, stagingPath string) error {
	endDate := time.Now()
	startDate := endDate.Add(time.Hour * -24)
	step := "15m"
	query := ""

	if endParam, ok := params["end"]; ok {
		tmp, err := time.Parse(time.RFC3339, endParam)
		if err != nil {
			return err
		}
		endDate = tmp
	}
	if startParam, ok := params["start"]; ok {
		tmp, err := time.Parse(time.RFC3339, startParam)
		if err != nil {
			return err
		}
		startDate = tmp
	}
	if stepParam, ok := params["step"]; ok {
		step = stepParam
	}
	if queryParam, ok := params["query"]; ok {
		query = queryParam
	}

	client := &http.Client{}

	req, _ := http.NewRequest("GET", p.URL+"/query_range", nil)

	q := req.URL.Query()

	q.Add("query", query)
	q.Add("start", startDate.Format(time.RFC3339))
	q.Add("end", endDate.Format(time.RFC3339))
	q.Add("step", step)

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	results := ClusterResults{}
	err = json.Unmarshal(data, &results)
	if err != nil {
		return err
	}

	timeseries, alerts := p.TransformToRows(&results)
	//fmt.Printf("TimeSeries Length %# v", pretty.Formatter(timeseries))

	base := fmt.Sprintf("%s/prometheus", stagingPath)

	p.WriteCSV(timeseries, base, Volume)
	p.WriteCSV(timeseries, base, Disk)
	p.WriteCSV(timeseries, base, Pool)
	p.WriteCSV(timeseries, base, Cluster)
	p.WriteAlertCSV(base, alerts)

	return nil
}

// CreateCSVRows creates the rows for a particular MAP in the CSVMetrics
func (p *Prometheus) CreateCSVRows(br CSVRow, m map[string]map[string]string, label string) (rows [][]string, rowHeaders []string) {
	for rowKey, rowValue := range m {
		var row = []string{strconv.FormatUint(uint64(br.Timestamp), 10), br.Cluster, br.Instance, br.Node}
		var rowHeader = []string{Timestamp, Cluster, Instance, Node}
		if label != "Cluster" && label != "Node" {
			row = append(row, rowKey)
			rowHeader = append(rowHeader, label)
		}
		for _, k := range Keys(rowValue) {
			v := rowValue[k]
			row = append(row, v)
			rowHeader = append(rowHeader, k)
		}
		rows = append(rows, row)
		rowHeaders = rowHeader
	}
	return rows, rowHeaders
}

// WriteCSV for a particular metric category, create the CSV file
func (p *Prometheus) WriteCSV(timeSeries map[CSVRow]*CSVMetrics, base, name string) {
	f, err := os.Create(base + "-" + name + ".csv")
	if err != nil {
		p.Log.Fatalln("Error Creating File", name, err)
	}

	defer f.Close()
	w := csv.NewWriter(f)
	var wroteHeader = false

	for br, bm := range timeSeries {
		// This is a bit of a hack to avoid reflection. If the CSVMetrics had a map instead of fields, we could avoid this.
		var m map[string]map[string]string
		if name == Volume {
			m = bm.Volume
		} else if name == Disk {
			m = bm.Disk
		} else if name == Pool {
			m = bm.Pool
		} else if name == Node {
			m = bm.Node
		}

		csvRows, csvHeader := p.CreateCSVRows(br, m, name)
		if wroteHeader == false {
			w.Write(csvHeader)
			wroteHeader = true
		}
		for _, csvRow := range csvRows {
			w.Write(csvRow)
		}
	}
	w.Flush()
}

// WriteAlertCSV for the alerts csv
func (p *Prometheus) WriteAlertCSV(base string, alerts []*AlertRow) {
	f, err := os.Create(base + "-alerts.csv")
	if err != nil {
		p.Log.Fatalln("Error Creating Alerts File", err)
	}

	defer f.Close()
	w := csv.NewWriter(f)
	w.Write([]string{Timestamp, Cluster, Instance, Node, "alert_name", "alert_state", "alert_severity", "alert_value"})
	for _, alert := range alerts {
		var row []string
		row = append(row, strconv.FormatUint(uint64(alert.csvRow.Timestamp), 10), alert.csvRow.Cluster, alert.csvRow.Instance, alert.csvRow.Node, alert.AlertName, alert.AlertState, alert.AlertSeverity, alert.AlertValue)
		w.Write(row)
	}
	w.Flush()
}

// TransformToRows takes the Prometheus API Calls ClusterResults and flattens it to the structure that can exported as CSV
func (p *Prometheus) TransformToRows(results *ClusterResults) (timeseries map[CSVRow]*CSVMetrics, alerts []*AlertRow) {
	// Go through the results for the cluster and generate the appropriate CSV files
	timeseries = make(map[CSVRow]*CSVMetrics)

	for _, result := range results.Data.Results {
		var values [][]interface{}
		values = result.Values
		if result.Values == nil {
			values = append(values, result.Value)
		}
		for _, value := range values {
			csvRow := CSVRow{
				Timestamp: uint32(value[0].(float64)),
				Cluster:   result.Metric.Cluster,
				Instance:  result.Metric.Instance,
				Node:      result.Metric.Node,
			}
			// fmt.Println(csvRow)
			csvMetrics, ok := timeseries[csvRow]
			if !ok {
				csvMetrics = new(CSVMetrics)
				timeseries[csvRow] = csvMetrics
			}
			if result.Metric.Name == Alerts {
				//	fmt.Printf(" RESULTS %# v\n", pretty.Formatter(result))

				alert := NewAlertRow(csvRow, &result, value[1].(string))
				alerts = append(alerts, alert)
			} else if result.Metric.Volume != nil {
				if csvMetrics.Volume == nil {
					csvMetrics.Volume = make(map[string]map[string]string)
				}
				if csvMetrics.Volume[*result.Metric.Volume] == nil {
					csvMetrics.Volume[*result.Metric.Volume] = make(map[string]string)
				}
				csvMetrics.Volume[*result.Metric.Volume][result.Metric.Name] = value[1].(string)
			} else if result.Metric.Disk != nil {
				if csvMetrics.Disk == nil {
					csvMetrics.Disk = make(map[string]map[string]string)
				}
				if csvMetrics.Disk[*result.Metric.Disk] == nil {
					csvMetrics.Disk[*result.Metric.Disk] = make(map[string]string)
				}
				csvMetrics.Disk[*result.Metric.Disk][result.Metric.Name] = value[1].(string)
			} else if result.Metric.Pool != nil {
				if csvMetrics.Pool == nil {
					csvMetrics.Pool = make(map[string]map[string]string)
				}
				if csvMetrics.Pool[*result.Metric.Pool] == nil {
					csvMetrics.Pool[*result.Metric.Pool] = make(map[string]string)
				}
				csvMetrics.Pool[*result.Metric.Pool][result.Metric.Name] = value[1].(string)
			} else if result.Metric.Proc != nil {
				if csvMetrics.Node == nil {
					csvMetrics.Node = make(map[string]map[string]string)
				}
				if csvMetrics.Node[result.Metric.Node] == nil {
					csvMetrics.Node[result.Metric.Node] = make(map[string]string)
				}
				csvMetrics.Node[result.Metric.Node][*result.Metric.Proc+"_"+result.Metric.Name] = value[1].(string)
			} else if strings.HasPrefix(result.Metric.Name, "px_node_stats") == true || strings.HasPrefix(result.Metric.Name, "px_network_") == true ||
				strings.HasPrefix(result.Metric.Name, "px_cluster_") == true {
				if csvMetrics.Node == nil {
					csvMetrics.Node = make(map[string]map[string]string)
				}
				if csvMetrics.Node[result.Metric.Node] == nil {
					csvMetrics.Node[result.Metric.Node] = make(map[string]string)
				}
				csvMetrics.Node[result.Metric.Node][result.Metric.Name] = value[1].(string)
			}
		}
	}
	return timeseries, alerts
}

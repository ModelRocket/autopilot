// Code generated by hiro; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	context "golang.org/x/net/context"

	"github.com/libopenstorage/autopilot/api/autopilot/rest/operations/collector"
	"github.com/libopenstorage/autopilot/api/autopilot/rest/operations/emitter"
	"github.com/libopenstorage/autopilot/api/autopilot/rest/operations/rule"
	"github.com/libopenstorage/autopilot/api/autopilot/rest/operations/sample"
	"github.com/libopenstorage/autopilot/api/autopilot/rest/operations/task"
)

// NewAutopilotAPI creates a new Autopilot instance
func NewAutopilotAPI(spec *loads.Document) *AutopilotAPI {
	return &AutopilotAPI{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		customConsumers:       make(map[string]runtime.Consumer),
		customProducers:       make(map[string]runtime.Producer),
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		BasicAuthenticator:    security.BasicAuthCtx,
		APIKeyAuthenticator:   security.APIKeyAuthCtx,
		BearerAuthenticator:   security.BearerAuthCtx,
		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		JSONProducer:          runtime.JSONProducer(),
		CollectorCollectorListHandler: collector.CollectorListHandlerFunc(func(params collector.CollectorListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CollectorCollectorList has not yet been implemented")
		}),
		CollectorCollectorPollHandler: collector.CollectorPollHandlerFunc(func(params collector.CollectorPollParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CollectorCollectorPoll has not yet been implemented")
		}),
		EmitterEmitterListHandler: emitter.EmitterListHandlerFunc(func(params emitter.EmitterListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation EmitterEmitterList has not yet been implemented")
		}),
		RecommendationsGetHandler: RecommendationsGetHandlerFunc(func(params RecommendationsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation RecommendationsGet has not yet been implemented")
		}),
		RuleRuleListHandler: rule.RuleListHandlerFunc(func(params rule.RuleListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation RuleRuleList has not yet been implemented")
		}),
		SampleSampleDeleteHandler: sample.SampleDeleteHandlerFunc(func(params sample.SampleDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation SampleSampleDelete has not yet been implemented")
		}),
		SampleSampleListHandler: sample.SampleListHandlerFunc(func(params sample.SampleListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation SampleSampleList has not yet been implemented")
		}),
		TaskTaskListHandler: task.TaskListHandlerFunc(func(params task.TaskListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation TaskTaskList has not yet been implemented")
		}),

		// Applies when the Authorization header is set with the Basic scheme
		BasicAuthAuth: func(ctx context.Context, user string, pass string) (context.Context, interface{}, error) {
			return ctx, nil, errors.NotImplemented("basic auth  (basicAuth) has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*AutopilotAPI libopenstorage autopilot API */
type AutopilotAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthenticationCtx) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthenticationCtx) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthenticationCtx) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// BasicAuthAuth registers a function that takes username and password and returns a principal
	// it performs authentication with basic auth
	BasicAuthAuth func(context.Context, string, string) (context.Context, interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// CollectorCollectorListHandler sets the operation handler for the collector list operation
	CollectorCollectorListHandler collector.CollectorListHandler
	// CollectorCollectorPollHandler sets the operation handler for the collector poll operation
	CollectorCollectorPollHandler collector.CollectorPollHandler
	// EmitterEmitterListHandler sets the operation handler for the emitter list operation
	EmitterEmitterListHandler emitter.EmitterListHandler
	// RecommendationsGetHandler sets the operation handler for the recommendations get operation
	RecommendationsGetHandler RecommendationsGetHandler
	// RuleRuleListHandler sets the operation handler for the rule list operation
	RuleRuleListHandler rule.RuleListHandler
	// SampleSampleDeleteHandler sets the operation handler for the sample delete operation
	SampleSampleDeleteHandler sample.SampleDeleteHandler
	// SampleSampleListHandler sets the operation handler for the sample list operation
	SampleSampleListHandler sample.SampleListHandler
	// TaskTaskListHandler sets the operation handler for the task list operation
	TaskTaskListHandler task.TaskListHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *AutopilotAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *AutopilotAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *AutopilotAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *AutopilotAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *AutopilotAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *AutopilotAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *AutopilotAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the AutopilotAPI
func (o *AutopilotAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BasicAuthAuth == nil {
		unregistered = append(unregistered, "BasicAuthAuth")
	}

	if o.CollectorCollectorListHandler == nil {
		unregistered = append(unregistered, "collector.CollectorListHandler")
	}

	if o.CollectorCollectorPollHandler == nil {
		unregistered = append(unregistered, "collector.CollectorPollHandler")
	}

	if o.EmitterEmitterListHandler == nil {
		unregistered = append(unregistered, "emitter.EmitterListHandler")
	}

	if o.RecommendationsGetHandler == nil {
		unregistered = append(unregistered, "RecommendationsGetHandler")
	}

	if o.RuleRuleListHandler == nil {
		unregistered = append(unregistered, "rule.RuleListHandler")
	}

	if o.SampleSampleDeleteHandler == nil {
		unregistered = append(unregistered, "sample.SampleDeleteHandler")
	}

	if o.SampleSampleListHandler == nil {
		unregistered = append(unregistered, "sample.SampleListHandler")
	}

	if o.TaskTaskListHandler == nil {
		unregistered = append(unregistered, "task.TaskListHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *AutopilotAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *AutopilotAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "basicAuth":
			_ = scheme
			result[name] = o.BasicAuthenticator(o.BasicAuthAuth)

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *AutopilotAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *AutopilotAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *AutopilotAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *AutopilotAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the autopilot API
func (o *AutopilotAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *AutopilotAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/collectors"] = collector.NewCollectorList(o.context, o.CollectorCollectorListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/collectors/{collector}/poll"] = collector.NewCollectorPoll(o.context, o.CollectorCollectorPollHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/emitters"] = emitter.NewEmitterList(o.context, o.EmitterEmitterListHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/recommend"] = NewRecommendationsGet(o.context, o.RecommendationsGetHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rules"] = rule.NewRuleList(o.context, o.RuleRuleListHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/samples/{sample_id}"] = sample.NewSampleDelete(o.context, o.SampleSampleDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/samples"] = sample.NewSampleList(o.context, o.SampleSampleListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tasks"] = task.NewTaskList(o.context, o.TaskTaskListHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *AutopilotAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *AutopilotAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *AutopilotAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *AutopilotAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

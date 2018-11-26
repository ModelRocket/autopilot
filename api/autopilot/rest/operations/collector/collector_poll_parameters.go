// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package collector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCollectorPollParams creates a new CollectorPollParams object
// no default values defined in spec.
func NewCollectorPollParams() CollectorPollParams {

	return CollectorPollParams{}
}

// CollectorPollParams contains all the bound params for the collector poll operation
// typically these are obtained from a http.Request
//
// swagger:parameters collectorPoll
type CollectorPollParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*the collector name
	  Required: true
	  In: path
	*/
	Collector string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCollectorPollParams() beforehand.
func (o *CollectorPollParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rCollector, rhkCollector, _ := route.Params.GetOK("collector")
	if err := o.bindCollector(rCollector, rhkCollector, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCollector binds and validates parameter Collector from path.
func (o *CollectorPollParams) bindCollector(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Collector = raw

	return nil
}

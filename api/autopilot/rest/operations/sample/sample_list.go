// Code generated by hiro; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package sample

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"gitlab.com/ModelRocket/sparks/cloud/provider"
)

// SampleListHandlerFunc turns a function with the right signature into a sample list handler
type SampleListHandlerFunc func(SampleListParams, provider.AuthToken) middleware.Responder

// Handle executing the request and returning a response
func (fn SampleListHandlerFunc) Handle(params SampleListParams, principal provider.AuthToken) middleware.Responder {
	return fn(params, principal)
}

// SampleListHandler interface for that can handle valid sample list params
type SampleListHandler interface {
	Handle(SampleListParams, provider.AuthToken) middleware.Responder
}

// NewSampleList creates a new http.Handler for the sample list operation
func NewSampleList(ctx *middleware.Context, handler SampleListHandler) *SampleList {
	return &SampleList{Context: ctx, Handler: handler}
}

/*SampleList swagger:route GET /samples sample sampleList

Get a list of telemetry samples

Returns an array of telemetry samples defined in the system

*/
type SampleList struct {
	Context *middleware.Context
	Handler SampleListHandler
}

func (o *SampleList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSampleListParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal provider.AuthToken
	if uprinc != nil {
		principal = uprinc.(provider.AuthToken) // this is really a provider.AuthToken, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

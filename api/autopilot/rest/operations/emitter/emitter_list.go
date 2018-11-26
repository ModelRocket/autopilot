// Code generated by hiro; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package emitter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// EmitterListHandlerFunc turns a function with the right signature into a emitter list handler
type EmitterListHandlerFunc func(EmitterListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn EmitterListHandlerFunc) Handle(params EmitterListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// EmitterListHandler interface for that can handle valid emitter list params
type EmitterListHandler interface {
	Handle(EmitterListParams, interface{}) middleware.Responder
}

// NewEmitterList creates a new http.Handler for the emitter list operation
func NewEmitterList(ctx *middleware.Context, handler EmitterListHandler) *EmitterList {
	return &EmitterList{Context: ctx, Handler: handler}
}

/*EmitterList swagger:route GET /emitters emitter emitterList

Get a list of telemetry emitters

Returns an array of telemetry emitters defined in the system

*/
type EmitterList struct {
	Context *middleware.Context
	Handler EmitterListHandler
}

func (o *EmitterList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewEmitterListParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

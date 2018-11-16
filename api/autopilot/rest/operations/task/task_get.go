// Code generated by hiro; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"gitlab.com/ModelRocket/sparks/cloud/provider"
)

// TaskGetHandlerFunc turns a function with the right signature into a task get handler
type TaskGetHandlerFunc func(TaskGetParams, provider.AuthToken) middleware.Responder

// Handle executing the request and returning a response
func (fn TaskGetHandlerFunc) Handle(params TaskGetParams, principal provider.AuthToken) middleware.Responder {
	return fn(params, principal)
}

// TaskGetHandler interface for that can handle valid task get params
type TaskGetHandler interface {
	Handle(TaskGetParams, provider.AuthToken) middleware.Responder
}

// NewTaskGet creates a new http.Handler for the task get operation
func NewTaskGet(ctx *middleware.Context, handler TaskGetHandler) *TaskGet {
	return &TaskGet{Context: ctx, Handler: handler}
}

/*TaskGet swagger:route GET /tasks/{task_id} task taskGet

Get a task

Returns the request task object

*/
type TaskGet struct {
	Context *middleware.Context
	Handler TaskGetHandler
}

func (o *TaskGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTaskGetParams()

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

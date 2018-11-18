// Code generated by hiro; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RuleGetHandlerFunc turns a function with the right signature into a rule get handler
type RuleGetHandlerFunc func(RuleGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn RuleGetHandlerFunc) Handle(params RuleGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// RuleGetHandler interface for that can handle valid rule get params
type RuleGetHandler interface {
	Handle(RuleGetParams, interface{}) middleware.Responder
}

// NewRuleGet creates a new http.Handler for the rule get operation
func NewRuleGet(ctx *middleware.Context, handler RuleGetHandler) *RuleGet {
	return &RuleGet{Context: ctx, Handler: handler}
}

/*RuleGet swagger:route GET /rules/{rule_id} rule ruleGet

Get a list of telemetry rules

Returns the request collected object

*/
type RuleGet struct {
	Context *middleware.Context
	Handler RuleGetHandler
}

func (o *RuleGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRuleGetParams()

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

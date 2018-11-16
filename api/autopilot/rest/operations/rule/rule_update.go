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
	"gitlab.com/ModelRocket/sparks/cloud/provider"
)

// RuleUpdateHandlerFunc turns a function with the right signature into a rule update handler
type RuleUpdateHandlerFunc func(RuleUpdateParams, provider.AuthToken) middleware.Responder

// Handle executing the request and returning a response
func (fn RuleUpdateHandlerFunc) Handle(params RuleUpdateParams, principal provider.AuthToken) middleware.Responder {
	return fn(params, principal)
}

// RuleUpdateHandler interface for that can handle valid rule update params
type RuleUpdateHandler interface {
	Handle(RuleUpdateParams, provider.AuthToken) middleware.Responder
}

// NewRuleUpdate creates a new http.Handler for the rule update operation
func NewRuleUpdate(ctx *middleware.Context, handler RuleUpdateHandler) *RuleUpdate {
	return &RuleUpdate{Context: ctx, Handler: handler}
}

/*RuleUpdate swagger:route PUT /rules/{rule_id} rule ruleUpdate

Update a rule object

Update the properties of the specified rule

*/
type RuleUpdate struct {
	Context *middleware.Context
	Handler RuleUpdateHandler
}

func (o *RuleUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRuleUpdateParams()

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

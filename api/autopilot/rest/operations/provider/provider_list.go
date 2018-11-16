// Code generated by hiro; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"gitlab.com/ModelRocket/sparks/cloud/provider"
)

// ProviderListHandlerFunc turns a function with the right signature into a provider list handler
type ProviderListHandlerFunc func(ProviderListParams, provider.AuthToken) middleware.Responder

// Handle executing the request and returning a response
func (fn ProviderListHandlerFunc) Handle(params ProviderListParams, principal provider.AuthToken) middleware.Responder {
	return fn(params, principal)
}

// ProviderListHandler interface for that can handle valid provider list params
type ProviderListHandler interface {
	Handle(ProviderListParams, provider.AuthToken) middleware.Responder
}

// NewProviderList creates a new http.Handler for the provider list operation
func NewProviderList(ctx *middleware.Context, handler ProviderListHandler) *ProviderList {
	return &ProviderList{Context: ctx, Handler: handler}
}

/*ProviderList swagger:route GET /providers provider providerList

Get a list of telemetry providers

Returns an array of telemetry providers defined in the system

*/
type ProviderList struct {
	Context *middleware.Context
	Handler ProviderListHandler
}

func (o *ProviderList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewProviderListParams()

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

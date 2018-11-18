// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	types "github.com/libopenstorage/autopilot/api/autopilot/types"
)

// NewRuleCreateParams creates a new RuleCreateParams object
// no default values defined in spec.
func NewRuleCreateParams() RuleCreateParams {

	return RuleCreateParams{}
}

// RuleCreateParams contains all the bound params for the rule create operation
// typically these are obtained from a http.Request
//
// swagger:parameters ruleCreate
type RuleCreateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The rule to create
	  Required: true
	  In: body
	*/
	Rule *types.RuleSet
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRuleCreateParams() beforehand.
func (o *RuleCreateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body types.RuleSet
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("rule", "body"))
			} else {
				res = append(res, errors.NewParseError("rule", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Rule = &body
			}
		}
	} else {
		res = append(res, errors.Required("rule", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

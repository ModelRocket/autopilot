// Code generated by go-swagger; DO NOT EDIT.

//
//  MODEL ROCKET LLC CONFIDENTIAL
//  _________________
//   Copyright (c) 2018 - 2019 MODEL ROCKET LLC
//   All Rights Reserved.
//
//   NOTICE:  All information contained herein is, and remains
//   the property of MODEL ROCKET LLC and its suppliers,
//   if any.  The intellectual and technical concepts contained
//   herein are proprietary to MODEL ROCKET LLC
//   and its suppliers and may be covered by U.S. and Foreign Patents,
//   patents in process, and are protected by trade secret or copyright law.
//   Dissemination of this information or reproduction of this material
//   is strictly forbidden unless prior written permission is obtained
//   from MODEL ROCKET LLC.
//

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewSourceListParams creates a new SourceListParams object
// no default values defined in spec.
func NewSourceListParams() SourceListParams {

	return SourceListParams{}
}

// SourceListParams contains all the bound params for the source list operation
// typically these are obtained from a http.Request
//
// swagger:parameters sourceList
type SourceListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSourceListParams() beforehand.
func (o *SourceListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package sample

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSampleGetParams creates a new SampleGetParams object
// no default values defined in spec.
func NewSampleGetParams() SampleGetParams {

	return SampleGetParams{}
}

// SampleGetParams contains all the bound params for the sample get operation
// typically these are obtained from a http.Request
//
// swagger:parameters sampleGet
type SampleGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The id of the sample
	  Required: true
	  In: path
	*/
	SampleID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSampleGetParams() beforehand.
func (o *SampleGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSampleID, rhkSampleID, _ := route.Params.GetOK("sample_id")
	if err := o.bindSampleID(rSampleID, rhkSampleID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSampleID binds and validates parameter SampleID from path.
func (o *SampleGetParams) bindSampleID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("sample_id", "path", "strfmt.UUID", raw)
	}
	o.SampleID = *(value.(*strfmt.UUID))

	if err := o.validateSampleID(formats); err != nil {
		return err
	}

	return nil
}

// validateSampleID carries on validations for parameter SampleID
func (o *SampleGetParams) validateSampleID(formats strfmt.Registry) error {

	if err := validate.FormatOf("sample_id", "path", "uuid", o.SampleID.String(), formats); err != nil {
		return err
	}
	return nil
}

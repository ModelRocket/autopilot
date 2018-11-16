// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	types "github.com/libopenstorage/autopilot/api/autopilot/types"
)

// SourceListOKCode is the HTTP code returned for type SourceListOK
const SourceListOKCode int = 200

/*SourceListOK OK

swagger:response sourceListOK
*/
type SourceListOK struct {

	/*
	  In: Body
	*/
	Payload []*types.Source `json:"body,omitempty"`
}

// NewSourceListOK creates SourceListOK with default headers values
func NewSourceListOK() *SourceListOK {

	return &SourceListOK{}
}

// WithPayload adds the payload to the source list o k response
func (o *SourceListOK) WithPayload(payload []*types.Source) *SourceListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the source list o k response
func (o *SourceListOK) SetPayload(payload []*types.Source) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SourceListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*types.Source, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// SourceListInternalServerErrorCode is the HTTP code returned for type SourceListInternalServerError
const SourceListInternalServerErrorCode int = 500

/*SourceListInternalServerError ServerError

swagger:response sourceListInternalServerError
*/
type SourceListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *types.Error `json:"body,omitempty"`
}

// NewSourceListInternalServerError creates SourceListInternalServerError with default headers values
func NewSourceListInternalServerError() *SourceListInternalServerError {

	return &SourceListInternalServerError{}
}

// WithPayload adds the payload to the source list internal server error response
func (o *SourceListInternalServerError) WithPayload(payload *types.Error) *SourceListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the source list internal server error response
func (o *SourceListInternalServerError) SetPayload(payload *types.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SourceListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

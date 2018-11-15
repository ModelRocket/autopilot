// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	types "github.com/libopenstorage/autopilot/api/autopilot/types"
)

// CollectorCreateCreatedCode is the HTTP code returned for type CollectorCreateCreated
const CollectorCreateCreatedCode int = 201

/*CollectorCreateCreated Created

swagger:response collectorCreateCreated
*/
type CollectorCreateCreated struct {

	/*
	  In: Body
	*/
	Payload *types.Collector `json:"body,omitempty"`
}

// NewCollectorCreateCreated creates CollectorCreateCreated with default headers values
func NewCollectorCreateCreated() *CollectorCreateCreated {

	return &CollectorCreateCreated{}
}

// WithPayload adds the payload to the collector create created response
func (o *CollectorCreateCreated) WithPayload(payload *types.Collector) *CollectorCreateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the collector create created response
func (o *CollectorCreateCreated) SetPayload(payload *types.Collector) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CollectorCreateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CollectorCreateBadRequestCode is the HTTP code returned for type CollectorCreateBadRequest
const CollectorCreateBadRequestCode int = 400

/*CollectorCreateBadRequest Bad Request

swagger:response collectorCreateBadRequest
*/
type CollectorCreateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *types.Error `json:"body,omitempty"`
}

// NewCollectorCreateBadRequest creates CollectorCreateBadRequest with default headers values
func NewCollectorCreateBadRequest() *CollectorCreateBadRequest {

	return &CollectorCreateBadRequest{}
}

// WithPayload adds the payload to the collector create bad request response
func (o *CollectorCreateBadRequest) WithPayload(payload *types.Error) *CollectorCreateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the collector create bad request response
func (o *CollectorCreateBadRequest) SetPayload(payload *types.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CollectorCreateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CollectorCreateInternalServerErrorCode is the HTTP code returned for type CollectorCreateInternalServerError
const CollectorCreateInternalServerErrorCode int = 500

/*CollectorCreateInternalServerError Server Error

swagger:response collectorCreateInternalServerError
*/
type CollectorCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *types.Error `json:"body,omitempty"`
}

// NewCollectorCreateInternalServerError creates CollectorCreateInternalServerError with default headers values
func NewCollectorCreateInternalServerError() *CollectorCreateInternalServerError {

	return &CollectorCreateInternalServerError{}
}

// WithPayload adds the payload to the collector create internal server error response
func (o *CollectorCreateInternalServerError) WithPayload(payload *types.Error) *CollectorCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the collector create internal server error response
func (o *CollectorCreateInternalServerError) SetPayload(payload *types.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CollectorCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

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

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Source A definition for a telemetry source
// swagger:model Source
type Source struct {

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this source
func (m *Source) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Source) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Source) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Source) UnmarshalBinary(b []byte) error {
	var res Source
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/strfmt"
)

// ProviderUpdateURL generates an URL for the provider update operation
type ProviderUpdateURL struct {
	ProviderID strfmt.UUID

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ProviderUpdateURL) WithBasePath(bp string) *ProviderUpdateURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ProviderUpdateURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ProviderUpdateURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/providers/{provider_id}"

	providerID := o.ProviderID.String()
	if providerID != "" {
		_path = strings.Replace(_path, "{provider_id}", providerID, -1)
	} else {
		return nil, errors.New("ProviderID is required on ProviderUpdateURL")
	}

	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ProviderUpdateURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ProviderUpdateURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ProviderUpdateURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ProviderUpdateURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ProviderUpdateURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ProviderUpdateURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

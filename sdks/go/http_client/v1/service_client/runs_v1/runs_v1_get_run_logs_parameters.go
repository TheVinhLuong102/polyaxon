// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewRunsV1GetRunLogsParams creates a new RunsV1GetRunLogsParams object
// with the default values initialized.
func NewRunsV1GetRunLogsParams() *RunsV1GetRunLogsParams {
	var ()
	return &RunsV1GetRunLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunsV1GetRunLogsParamsWithTimeout creates a new RunsV1GetRunLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunsV1GetRunLogsParamsWithTimeout(timeout time.Duration) *RunsV1GetRunLogsParams {
	var ()
	return &RunsV1GetRunLogsParams{

		timeout: timeout,
	}
}

// NewRunsV1GetRunLogsParamsWithContext creates a new RunsV1GetRunLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunsV1GetRunLogsParamsWithContext(ctx context.Context) *RunsV1GetRunLogsParams {
	var ()
	return &RunsV1GetRunLogsParams{

		Context: ctx,
	}
}

// NewRunsV1GetRunLogsParamsWithHTTPClient creates a new RunsV1GetRunLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunsV1GetRunLogsParamsWithHTTPClient(client *http.Client) *RunsV1GetRunLogsParams {
	var ()
	return &RunsV1GetRunLogsParams{
		HTTPClient: client,
	}
}

/*RunsV1GetRunLogsParams contains all the parameters to send to the API endpoint
for the runs v1 get run logs operation typically these are written to a http.Request
*/
type RunsV1GetRunLogsParams struct {

	/*LastFile
	  last file.

	*/
	LastFile *string
	/*LastTime
	  last time.

	*/
	LastTime *strfmt.DateTime
	/*Namespace*/
	Namespace string
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project where the run will be assigned

	*/
	Project string
	/*UUID
	  Uuid identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) WithTimeout(timeout time.Duration) *RunsV1GetRunLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) WithContext(ctx context.Context) *RunsV1GetRunLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) WithHTTPClient(client *http.Client) *RunsV1GetRunLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLastFile adds the lastFile to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) WithLastFile(lastFile *string) *RunsV1GetRunLogsParams {
	o.SetLastFile(lastFile)
	return o
}

// SetLastFile adds the lastFile to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) SetLastFile(lastFile *string) {
	o.LastFile = lastFile
}

// WithLastTime adds the lastTime to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) WithLastTime(lastTime *strfmt.DateTime) *RunsV1GetRunLogsParams {
	o.SetLastTime(lastTime)
	return o
}

// SetLastTime adds the lastTime to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) SetLastTime(lastTime *strfmt.DateTime) {
	o.LastTime = lastTime
}

// WithNamespace adds the namespace to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) WithNamespace(namespace string) *RunsV1GetRunLogsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOwner adds the owner to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) WithOwner(owner string) *RunsV1GetRunLogsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) WithProject(project string) *RunsV1GetRunLogsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) SetProject(project string) {
	o.Project = project
}

// WithUUID adds the uuid to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) WithUUID(uuid string) *RunsV1GetRunLogsParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the runs v1 get run logs params
func (o *RunsV1GetRunLogsParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *RunsV1GetRunLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LastFile != nil {

		// query param last_file
		var qrLastFile string
		if o.LastFile != nil {
			qrLastFile = *o.LastFile
		}
		qLastFile := qrLastFile
		if qLastFile != "" {
			if err := r.SetQueryParam("last_file", qLastFile); err != nil {
				return err
			}
		}

	}

	if o.LastTime != nil {

		// query param last_time
		var qrLastTime strfmt.DateTime
		if o.LastTime != nil {
			qrLastTime = *o.LastTime
		}
		qLastTime := qrLastTime.String()
		if qLastTime != "" {
			if err := r.SetQueryParam("last_time", qLastTime); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
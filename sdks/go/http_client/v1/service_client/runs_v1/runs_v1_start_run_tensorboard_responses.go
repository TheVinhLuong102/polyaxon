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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// RunsV1StartRunTensorboardReader is a Reader for the RunsV1StartRunTensorboard structure.
type RunsV1StartRunTensorboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunsV1StartRunTensorboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunsV1StartRunTensorboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewRunsV1StartRunTensorboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRunsV1StartRunTensorboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunsV1StartRunTensorboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRunsV1StartRunTensorboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunsV1StartRunTensorboardOK creates a RunsV1StartRunTensorboardOK with default headers values
func NewRunsV1StartRunTensorboardOK() *RunsV1StartRunTensorboardOK {
	return &RunsV1StartRunTensorboardOK{}
}

/*RunsV1StartRunTensorboardOK handles this case with default header values.

A successful response.
*/
type RunsV1StartRunTensorboardOK struct {
}

func (o *RunsV1StartRunTensorboardOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/tensorboard/start][%d] runsV1StartRunTensorboardOK ", 200)
}

func (o *RunsV1StartRunTensorboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRunsV1StartRunTensorboardNoContent creates a RunsV1StartRunTensorboardNoContent with default headers values
func NewRunsV1StartRunTensorboardNoContent() *RunsV1StartRunTensorboardNoContent {
	return &RunsV1StartRunTensorboardNoContent{}
}

/*RunsV1StartRunTensorboardNoContent handles this case with default header values.

No content.
*/
type RunsV1StartRunTensorboardNoContent struct {
	Payload interface{}
}

func (o *RunsV1StartRunTensorboardNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/tensorboard/start][%d] runsV1StartRunTensorboardNoContent  %+v", 204, o.Payload)
}

func (o *RunsV1StartRunTensorboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1StartRunTensorboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1StartRunTensorboardForbidden creates a RunsV1StartRunTensorboardForbidden with default headers values
func NewRunsV1StartRunTensorboardForbidden() *RunsV1StartRunTensorboardForbidden {
	return &RunsV1StartRunTensorboardForbidden{}
}

/*RunsV1StartRunTensorboardForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type RunsV1StartRunTensorboardForbidden struct {
	Payload interface{}
}

func (o *RunsV1StartRunTensorboardForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/tensorboard/start][%d] runsV1StartRunTensorboardForbidden  %+v", 403, o.Payload)
}

func (o *RunsV1StartRunTensorboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1StartRunTensorboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1StartRunTensorboardNotFound creates a RunsV1StartRunTensorboardNotFound with default headers values
func NewRunsV1StartRunTensorboardNotFound() *RunsV1StartRunTensorboardNotFound {
	return &RunsV1StartRunTensorboardNotFound{}
}

/*RunsV1StartRunTensorboardNotFound handles this case with default header values.

Resource does not exist.
*/
type RunsV1StartRunTensorboardNotFound struct {
	Payload interface{}
}

func (o *RunsV1StartRunTensorboardNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/tensorboard/start][%d] runsV1StartRunTensorboardNotFound  %+v", 404, o.Payload)
}

func (o *RunsV1StartRunTensorboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *RunsV1StartRunTensorboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsV1StartRunTensorboardDefault creates a RunsV1StartRunTensorboardDefault with default headers values
func NewRunsV1StartRunTensorboardDefault(code int) *RunsV1StartRunTensorboardDefault {
	return &RunsV1StartRunTensorboardDefault{
		_statusCode: code,
	}
}

/*RunsV1StartRunTensorboardDefault handles this case with default header values.

An unexpected error response
*/
type RunsV1StartRunTensorboardDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the runs v1 start run tensorboard default response
func (o *RunsV1StartRunTensorboardDefault) Code() int {
	return o._statusCode
}

func (o *RunsV1StartRunTensorboardDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/tensorboard/start][%d] RunsV1_StartRunTensorboard default  %+v", o._statusCode, o.Payload)
}

func (o *RunsV1StartRunTensorboardDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *RunsV1StartRunTensorboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
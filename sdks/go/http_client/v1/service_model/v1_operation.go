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

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Operation Operation specification
//
// swagger:model v1Operation
type V1Operation struct {

	// Optional flag to disable cache validation and force run this component
	Cache *V1Cache `json:"cache,omitempty"`

	// component
	Component *V1Component `json:"component,omitempty"`

	// An optional list of condition to check before starting the run, entities should be a valid Condition
	Conditions []interface{} `json:"conditions"`

	// dag ref
	DagRef string `json:"dag_ref,omitempty"`

	// Optional graph dependencies of this op
	Dependencies []string `json:"dependencies"`

	// Optional component description
	Description string `json:"description,omitempty"`

	// hub ref
	HubRef string `json:"hub_ref,omitempty"`

	// Optional component kind, should be equal to "component"
	Kind string `json:"kind,omitempty"`

	// Optional component name, should a valid slug
	Name string `json:"name,omitempty"`

	// Optional parallel section, must be a valid Parallel option (Random/Grid/BO/Hyperband/Hyperopt/Mapping/Iterative)
	Parallel interface{} `json:"parallel,omitempty"`

	// Optional dict of params
	Params map[string]V1Param `json:"params,omitempty"`

	// path ref
	PathRef string `json:"path_ref,omitempty"`

	// Optional plugins to enable
	Plugins *V1Plugins `json:"plugins,omitempty"`

	// Optional profile to use for running this component
	Profile string `json:"profile,omitempty"`

	// Optional queue to use for running this component
	Queue string `json:"queue,omitempty"`

	// Optional a run section to override the the content of the run in the template
	// should be one of: Job/Service/Spark/Flink/Kubeflow/Dask/Dag
	RunPatch interface{} `json:"run_patch,omitempty"`

	// Optional schedule section, must be a valid Schedule option (Cron/Interval/Repeatable/ExactTime)
	Schedule interface{} `json:"schedule,omitempty"`

	// Optional flag to skip this run if upstream was skipped
	SkipOnUpstreamSkip bool `json:"skip_on_upstream_skip,omitempty"`

	// Optional component tag version
	Tag string `json:"tag,omitempty"`

	// Optional component tags
	Tags []string `json:"tags"`

	// optional termination section
	Termination *V1Termination `json:"termination,omitempty"`

	// Optional trigger policy
	Trigger V1TriggerPolicy `json:"trigger,omitempty"`

	// url ref
	URLRef string `json:"url_ref,omitempty"`

	// Spec version
	Version float32 `json:"version,omitempty"`
}

// Validate validates this v1 operation
func (m *V1Operation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCache(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlugins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTermination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Operation) validateCache(formats strfmt.Registry) error {

	if swag.IsZero(m.Cache) { // not required
		return nil
	}

	if m.Cache != nil {
		if err := m.Cache.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cache")
			}
			return err
		}
	}

	return nil
}

func (m *V1Operation) validateComponent(formats strfmt.Registry) error {

	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *V1Operation) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(m.Params) { // not required
		return nil
	}

	for k := range m.Params {

		if err := validate.Required("params"+"."+k, "body", m.Params[k]); err != nil {
			return err
		}
		if val, ok := m.Params[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V1Operation) validatePlugins(formats strfmt.Registry) error {

	if swag.IsZero(m.Plugins) { // not required
		return nil
	}

	if m.Plugins != nil {
		if err := m.Plugins.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plugins")
			}
			return err
		}
	}

	return nil
}

func (m *V1Operation) validateTermination(formats strfmt.Registry) error {

	if swag.IsZero(m.Termination) { // not required
		return nil
	}

	if m.Termination != nil {
		if err := m.Termination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("termination")
			}
			return err
		}
	}

	return nil
}

func (m *V1Operation) validateTrigger(formats strfmt.Registry) error {

	if swag.IsZero(m.Trigger) { // not required
		return nil
	}

	if err := m.Trigger.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trigger")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Operation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Operation) UnmarshalBinary(b []byte) error {
	var res V1Operation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
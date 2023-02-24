// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Aggregate aggregate
//
// swagger:model Aggregate
type Aggregate struct {

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Date added
	// Format: date
	DateAdded *strfmt.Date `json:"date_added,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// family
	Family *AggregateFamily `json:"family,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// Prefix
	// Required: true
	Prefix *string `json:"prefix"`

	// rir
	// Required: true
	Rir *NestedRIR `json:"rir"`

	// tags
	Tags []*NestedTag `json:"tags,omitempty"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this aggregate
func (m *Aggregate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateAdded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Aggregate) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Aggregate) validateDateAdded(formats strfmt.Registry) error {
	if swag.IsZero(m.DateAdded) { // not required
		return nil
	}

	if err := validate.FormatOf("date_added", "body", "date", m.DateAdded.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Aggregate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *Aggregate) validateFamily(formats strfmt.Registry) error {
	if swag.IsZero(m.Family) { // not required
		return nil
	}

	if m.Family != nil {
		if err := m.Family.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("family")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("family")
			}
			return err
		}
	}

	return nil
}

func (m *Aggregate) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Aggregate) validatePrefix(formats strfmt.Registry) error {

	if err := validate.Required("prefix", "body", m.Prefix); err != nil {
		return err
	}

	return nil
}

func (m *Aggregate) validateRir(formats strfmt.Registry) error {

	if err := validate.Required("rir", "body", m.Rir); err != nil {
		return err
	}

	if m.Rir != nil {
		if err := m.Rir.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rir")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rir")
			}
			return err
		}
	}

	return nil
}

func (m *Aggregate) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Aggregate) validateTenant(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *Aggregate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this aggregate based on the context it is used
func (m *Aggregate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFamily(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRir(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Aggregate) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *Aggregate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *Aggregate) contextValidateFamily(ctx context.Context, formats strfmt.Registry) error {

	if m.Family != nil {
		if err := m.Family.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("family")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("family")
			}
			return err
		}
	}

	return nil
}

func (m *Aggregate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Aggregate) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *Aggregate) contextValidateRir(ctx context.Context, formats strfmt.Registry) error {

	if m.Rir != nil {
		if err := m.Rir.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rir")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rir")
			}
			return err
		}
	}

	return nil
}

func (m *Aggregate) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Aggregate) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

	if m.Tenant != nil {
		if err := m.Tenant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *Aggregate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Aggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Aggregate) UnmarshalBinary(b []byte) error {
	var res Aggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AggregateFamily Family
//
// swagger:model AggregateFamily
type AggregateFamily struct {

	// label
	// Required: true
	// Enum: [IPv4 IPv6]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [4 6]
	Value *int64 `json:"value"`
}

// Validate validates this aggregate family
func (m *AggregateFamily) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var aggregateFamilyTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPv4","IPv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aggregateFamilyTypeLabelPropEnum = append(aggregateFamilyTypeLabelPropEnum, v)
	}
}

const (

	// AggregateFamilyLabelIPV4 captures enum value "IPv4"
	AggregateFamilyLabelIPV4 string = "IPv4"

	// AggregateFamilyLabelIPV6 captures enum value "IPv6"
	AggregateFamilyLabelIPV6 string = "IPv6"
)

// prop value enum
func (m *AggregateFamily) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, aggregateFamilyTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AggregateFamily) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("family"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("family"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var aggregateFamilyTypeValuePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[4,6]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aggregateFamilyTypeValuePropEnum = append(aggregateFamilyTypeValuePropEnum, v)
	}
}

// prop value enum
func (m *AggregateFamily) validateValueEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, aggregateFamilyTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AggregateFamily) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("family"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("family"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this aggregate family based on the context it is used
func (m *AggregateFamily) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AggregateFamily) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregateFamily) UnmarshalBinary(b []byte) error {
	var res AggregateFamily
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

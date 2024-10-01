/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.2 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the IKEProposalAuthenticationAlgorithm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IKEProposalAuthenticationAlgorithm{}

// IKEProposalAuthenticationAlgorithm struct for IKEProposalAuthenticationAlgorithm
type IKEProposalAuthenticationAlgorithm struct {
	Value *IKEProposalAuthenticationAlgorithmValue `json:"value,omitempty"`
	Label *IKEProposalAuthenticationAlgorithmLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IKEProposalAuthenticationAlgorithm IKEProposalAuthenticationAlgorithm

// NewIKEProposalAuthenticationAlgorithm instantiates a new IKEProposalAuthenticationAlgorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIKEProposalAuthenticationAlgorithm() *IKEProposalAuthenticationAlgorithm {
	this := IKEProposalAuthenticationAlgorithm{}
	return &this
}

// NewIKEProposalAuthenticationAlgorithmWithDefaults instantiates a new IKEProposalAuthenticationAlgorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIKEProposalAuthenticationAlgorithmWithDefaults() *IKEProposalAuthenticationAlgorithm {
	this := IKEProposalAuthenticationAlgorithm{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IKEProposalAuthenticationAlgorithm) GetValue() IKEProposalAuthenticationAlgorithmValue {
	if o == nil || IsNil(o.Value) {
		var ret IKEProposalAuthenticationAlgorithmValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposalAuthenticationAlgorithm) GetValueOk() (*IKEProposalAuthenticationAlgorithmValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IKEProposalAuthenticationAlgorithm) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given IKEProposalAuthenticationAlgorithmValue and assigns it to the Value field.
func (o *IKEProposalAuthenticationAlgorithm) SetValue(v IKEProposalAuthenticationAlgorithmValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IKEProposalAuthenticationAlgorithm) GetLabel() IKEProposalAuthenticationAlgorithmLabel {
	if o == nil || IsNil(o.Label) {
		var ret IKEProposalAuthenticationAlgorithmLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposalAuthenticationAlgorithm) GetLabelOk() (*IKEProposalAuthenticationAlgorithmLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IKEProposalAuthenticationAlgorithm) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given IKEProposalAuthenticationAlgorithmLabel and assigns it to the Label field.
func (o *IKEProposalAuthenticationAlgorithm) SetLabel(v IKEProposalAuthenticationAlgorithmLabel) {
	o.Label = &v
}

func (o IKEProposalAuthenticationAlgorithm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IKEProposalAuthenticationAlgorithm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IKEProposalAuthenticationAlgorithm) UnmarshalJSON(data []byte) (err error) {
	varIKEProposalAuthenticationAlgorithm := _IKEProposalAuthenticationAlgorithm{}

	err = json.Unmarshal(data, &varIKEProposalAuthenticationAlgorithm)

	if err != nil {
		return err
	}

	*o = IKEProposalAuthenticationAlgorithm(varIKEProposalAuthenticationAlgorithm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIKEProposalAuthenticationAlgorithm struct {
	value *IKEProposalAuthenticationAlgorithm
	isSet bool
}

func (v NullableIKEProposalAuthenticationAlgorithm) Get() *IKEProposalAuthenticationAlgorithm {
	return v.value
}

func (v *NullableIKEProposalAuthenticationAlgorithm) Set(val *IKEProposalAuthenticationAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEProposalAuthenticationAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEProposalAuthenticationAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEProposalAuthenticationAlgorithm(val *IKEProposalAuthenticationAlgorithm) *NullableIKEProposalAuthenticationAlgorithm {
	return &NullableIKEProposalAuthenticationAlgorithm{value: val, isSet: true}
}

func (v NullableIKEProposalAuthenticationAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEProposalAuthenticationAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



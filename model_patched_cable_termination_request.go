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

// checks if the PatchedCableTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedCableTerminationRequest{}

// PatchedCableTerminationRequest Adds support for custom fields and tags.
type PatchedCableTerminationRequest struct {
	Cable *int32 `json:"cable,omitempty"`
	CableEnd *End1 `json:"cable_end,omitempty"`
	TerminationType *string `json:"termination_type,omitempty"`
	TerminationId *int64 `json:"termination_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedCableTerminationRequest PatchedCableTerminationRequest

// NewPatchedCableTerminationRequest instantiates a new PatchedCableTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCableTerminationRequest() *PatchedCableTerminationRequest {
	this := PatchedCableTerminationRequest{}
	return &this
}

// NewPatchedCableTerminationRequestWithDefaults instantiates a new PatchedCableTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCableTerminationRequestWithDefaults() *PatchedCableTerminationRequest {
	this := PatchedCableTerminationRequest{}
	return &this
}

// GetCable returns the Cable field value if set, zero value otherwise.
func (o *PatchedCableTerminationRequest) GetCable() int32 {
	if o == nil || IsNil(o.Cable) {
		var ret int32
		return ret
	}
	return *o.Cable
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCableTerminationRequest) GetCableOk() (*int32, bool) {
	if o == nil || IsNil(o.Cable) {
		return nil, false
	}
	return o.Cable, true
}

// HasCable returns a boolean if a field has been set.
func (o *PatchedCableTerminationRequest) HasCable() bool {
	if o != nil && !IsNil(o.Cable) {
		return true
	}

	return false
}

// SetCable gets a reference to the given int32 and assigns it to the Cable field.
func (o *PatchedCableTerminationRequest) SetCable(v int32) {
	o.Cable = &v
}

// GetCableEnd returns the CableEnd field value if set, zero value otherwise.
func (o *PatchedCableTerminationRequest) GetCableEnd() End1 {
	if o == nil || IsNil(o.CableEnd) {
		var ret End1
		return ret
	}
	return *o.CableEnd
}

// GetCableEndOk returns a tuple with the CableEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCableTerminationRequest) GetCableEndOk() (*End1, bool) {
	if o == nil || IsNil(o.CableEnd) {
		return nil, false
	}
	return o.CableEnd, true
}

// HasCableEnd returns a boolean if a field has been set.
func (o *PatchedCableTerminationRequest) HasCableEnd() bool {
	if o != nil && !IsNil(o.CableEnd) {
		return true
	}

	return false
}

// SetCableEnd gets a reference to the given End1 and assigns it to the CableEnd field.
func (o *PatchedCableTerminationRequest) SetCableEnd(v End1) {
	o.CableEnd = &v
}

// GetTerminationType returns the TerminationType field value if set, zero value otherwise.
func (o *PatchedCableTerminationRequest) GetTerminationType() string {
	if o == nil || IsNil(o.TerminationType) {
		var ret string
		return ret
	}
	return *o.TerminationType
}

// GetTerminationTypeOk returns a tuple with the TerminationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCableTerminationRequest) GetTerminationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TerminationType) {
		return nil, false
	}
	return o.TerminationType, true
}

// HasTerminationType returns a boolean if a field has been set.
func (o *PatchedCableTerminationRequest) HasTerminationType() bool {
	if o != nil && !IsNil(o.TerminationType) {
		return true
	}

	return false
}

// SetTerminationType gets a reference to the given string and assigns it to the TerminationType field.
func (o *PatchedCableTerminationRequest) SetTerminationType(v string) {
	o.TerminationType = &v
}

// GetTerminationId returns the TerminationId field value if set, zero value otherwise.
func (o *PatchedCableTerminationRequest) GetTerminationId() int64 {
	if o == nil || IsNil(o.TerminationId) {
		var ret int64
		return ret
	}
	return *o.TerminationId
}

// GetTerminationIdOk returns a tuple with the TerminationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCableTerminationRequest) GetTerminationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TerminationId) {
		return nil, false
	}
	return o.TerminationId, true
}

// HasTerminationId returns a boolean if a field has been set.
func (o *PatchedCableTerminationRequest) HasTerminationId() bool {
	if o != nil && !IsNil(o.TerminationId) {
		return true
	}

	return false
}

// SetTerminationId gets a reference to the given int64 and assigns it to the TerminationId field.
func (o *PatchedCableTerminationRequest) SetTerminationId(v int64) {
	o.TerminationId = &v
}

func (o PatchedCableTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedCableTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cable) {
		toSerialize["cable"] = o.Cable
	}
	if !IsNil(o.CableEnd) {
		toSerialize["cable_end"] = o.CableEnd
	}
	if !IsNil(o.TerminationType) {
		toSerialize["termination_type"] = o.TerminationType
	}
	if !IsNil(o.TerminationId) {
		toSerialize["termination_id"] = o.TerminationId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedCableTerminationRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedCableTerminationRequest := _PatchedCableTerminationRequest{}

	err = json.Unmarshal(data, &varPatchedCableTerminationRequest)

	if err != nil {
		return err
	}

	*o = PatchedCableTerminationRequest(varPatchedCableTerminationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cable")
		delete(additionalProperties, "cable_end")
		delete(additionalProperties, "termination_type")
		delete(additionalProperties, "termination_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedCableTerminationRequest struct {
	value *PatchedCableTerminationRequest
	isSet bool
}

func (v NullablePatchedCableTerminationRequest) Get() *PatchedCableTerminationRequest {
	return v.value
}

func (v *NullablePatchedCableTerminationRequest) Set(val *PatchedCableTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCableTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCableTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCableTerminationRequest(val *PatchedCableTerminationRequest) *NullablePatchedCableTerminationRequest {
	return &NullablePatchedCableTerminationRequest{value: val, isSet: true}
}

func (v NullablePatchedCableTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCableTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



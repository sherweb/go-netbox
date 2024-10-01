/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.2 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the CableTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CableTerminationRequest{}

// CableTerminationRequest Adds support for custom fields and tags.
type CableTerminationRequest struct {
	Cable int32 `json:"cable"`
	CableEnd End1 `json:"cable_end"`
	TerminationType string `json:"termination_type"`
	TerminationId int64 `json:"termination_id"`
	AdditionalProperties map[string]interface{}
}

type _CableTerminationRequest CableTerminationRequest

// NewCableTerminationRequest instantiates a new CableTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCableTerminationRequest(cable int32, cableEnd End1, terminationType string, terminationId int64) *CableTerminationRequest {
	this := CableTerminationRequest{}
	this.Cable = cable
	this.CableEnd = cableEnd
	this.TerminationType = terminationType
	this.TerminationId = terminationId
	return &this
}

// NewCableTerminationRequestWithDefaults instantiates a new CableTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCableTerminationRequestWithDefaults() *CableTerminationRequest {
	this := CableTerminationRequest{}
	return &this
}

// GetCable returns the Cable field value
func (o *CableTerminationRequest) GetCable() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cable
}

// GetCableOk returns a tuple with the Cable field value
// and a boolean to check if the value has been set.
func (o *CableTerminationRequest) GetCableOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cable, true
}

// SetCable sets field value
func (o *CableTerminationRequest) SetCable(v int32) {
	o.Cable = v
}

// GetCableEnd returns the CableEnd field value
func (o *CableTerminationRequest) GetCableEnd() End1 {
	if o == nil {
		var ret End1
		return ret
	}

	return o.CableEnd
}

// GetCableEndOk returns a tuple with the CableEnd field value
// and a boolean to check if the value has been set.
func (o *CableTerminationRequest) GetCableEndOk() (*End1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CableEnd, true
}

// SetCableEnd sets field value
func (o *CableTerminationRequest) SetCableEnd(v End1) {
	o.CableEnd = v
}

// GetTerminationType returns the TerminationType field value
func (o *CableTerminationRequest) GetTerminationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerminationType
}

// GetTerminationTypeOk returns a tuple with the TerminationType field value
// and a boolean to check if the value has been set.
func (o *CableTerminationRequest) GetTerminationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminationType, true
}

// SetTerminationType sets field value
func (o *CableTerminationRequest) SetTerminationType(v string) {
	o.TerminationType = v
}

// GetTerminationId returns the TerminationId field value
func (o *CableTerminationRequest) GetTerminationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TerminationId
}

// GetTerminationIdOk returns a tuple with the TerminationId field value
// and a boolean to check if the value has been set.
func (o *CableTerminationRequest) GetTerminationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminationId, true
}

// SetTerminationId sets field value
func (o *CableTerminationRequest) SetTerminationId(v int64) {
	o.TerminationId = v
}

func (o CableTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CableTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cable"] = o.Cable
	toSerialize["cable_end"] = o.CableEnd
	toSerialize["termination_type"] = o.TerminationType
	toSerialize["termination_id"] = o.TerminationId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CableTerminationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cable",
		"cable_end",
		"termination_type",
		"termination_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCableTerminationRequest := _CableTerminationRequest{}

	err = json.Unmarshal(data, &varCableTerminationRequest)

	if err != nil {
		return err
	}

	*o = CableTerminationRequest(varCableTerminationRequest)

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

type NullableCableTerminationRequest struct {
	value *CableTerminationRequest
	isSet bool
}

func (v NullableCableTerminationRequest) Get() *CableTerminationRequest {
	return v.value
}

func (v *NullableCableTerminationRequest) Set(val *CableTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCableTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCableTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableTerminationRequest(val *CableTerminationRequest) *NullableCableTerminationRequest {
	return &NullableCableTerminationRequest{value: val, isSet: true}
}

func (v NullableCableTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



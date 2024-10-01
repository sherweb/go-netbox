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

// checks if the BriefInterfaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefInterfaceRequest{}

// BriefInterfaceRequest Adds support for custom fields and tags.
type BriefInterfaceRequest struct {
	Device BriefDeviceRequest `json:"device"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefInterfaceRequest BriefInterfaceRequest

// NewBriefInterfaceRequest instantiates a new BriefInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefInterfaceRequest(device BriefDeviceRequest, name string) *BriefInterfaceRequest {
	this := BriefInterfaceRequest{}
	this.Device = device
	this.Name = name
	return &this
}

// NewBriefInterfaceRequestWithDefaults instantiates a new BriefInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefInterfaceRequestWithDefaults() *BriefInterfaceRequest {
	this := BriefInterfaceRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *BriefInterfaceRequest) GetDevice() BriefDeviceRequest {
	if o == nil {
		var ret BriefDeviceRequest
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *BriefInterfaceRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *BriefInterfaceRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = v
}

// GetName returns the Name field value
func (o *BriefInterfaceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefInterfaceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefInterfaceRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefInterfaceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefInterfaceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefInterfaceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefInterfaceRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefInterfaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefInterfaceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
		"name",
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

	varBriefInterfaceRequest := _BriefInterfaceRequest{}

	err = json.Unmarshal(data, &varBriefInterfaceRequest)

	if err != nil {
		return err
	}

	*o = BriefInterfaceRequest(varBriefInterfaceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefInterfaceRequest struct {
	value *BriefInterfaceRequest
	isSet bool
}

func (v NullableBriefInterfaceRequest) Get() *BriefInterfaceRequest {
	return v.value
}

func (v *NullableBriefInterfaceRequest) Set(val *BriefInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefInterfaceRequest(val *BriefInterfaceRequest) *NullableBriefInterfaceRequest {
	return &NullableBriefInterfaceRequest{value: val, isSet: true}
}

func (v NullableBriefInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



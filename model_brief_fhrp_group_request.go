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

// checks if the BriefFHRPGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefFHRPGroupRequest{}

// BriefFHRPGroupRequest Adds support for custom fields and tags.
type BriefFHRPGroupRequest struct {
	Protocol BriefFHRPGroupProtocol `json:"protocol"`
	GroupId int32 `json:"group_id"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefFHRPGroupRequest BriefFHRPGroupRequest

// NewBriefFHRPGroupRequest instantiates a new BriefFHRPGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefFHRPGroupRequest(protocol BriefFHRPGroupProtocol, groupId int32) *BriefFHRPGroupRequest {
	this := BriefFHRPGroupRequest{}
	this.Protocol = protocol
	this.GroupId = groupId
	return &this
}

// NewBriefFHRPGroupRequestWithDefaults instantiates a new BriefFHRPGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefFHRPGroupRequestWithDefaults() *BriefFHRPGroupRequest {
	this := BriefFHRPGroupRequest{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *BriefFHRPGroupRequest) GetProtocol() BriefFHRPGroupProtocol {
	if o == nil {
		var ret BriefFHRPGroupProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *BriefFHRPGroupRequest) GetProtocolOk() (*BriefFHRPGroupProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *BriefFHRPGroupRequest) SetProtocol(v BriefFHRPGroupProtocol) {
	o.Protocol = v
}

// GetGroupId returns the GroupId field value
func (o *BriefFHRPGroupRequest) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *BriefFHRPGroupRequest) GetGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *BriefFHRPGroupRequest) SetGroupId(v int32) {
	o.GroupId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefFHRPGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefFHRPGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefFHRPGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefFHRPGroupRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefFHRPGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefFHRPGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	toSerialize["group_id"] = o.GroupId
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefFHRPGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"protocol",
		"group_id",
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

	varBriefFHRPGroupRequest := _BriefFHRPGroupRequest{}

	err = json.Unmarshal(data, &varBriefFHRPGroupRequest)

	if err != nil {
		return err
	}

	*o = BriefFHRPGroupRequest(varBriefFHRPGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "group_id")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefFHRPGroupRequest struct {
	value *BriefFHRPGroupRequest
	isSet bool
}

func (v NullableBriefFHRPGroupRequest) Get() *BriefFHRPGroupRequest {
	return v.value
}

func (v *NullableBriefFHRPGroupRequest) Set(val *BriefFHRPGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefFHRPGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefFHRPGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefFHRPGroupRequest(val *BriefFHRPGroupRequest) *NullableBriefFHRPGroupRequest {
	return &NullableBriefFHRPGroupRequest{value: val, isSet: true}
}

func (v NullableBriefFHRPGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefFHRPGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



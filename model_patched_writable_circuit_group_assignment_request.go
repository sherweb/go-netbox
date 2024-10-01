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

// checks if the PatchedWritableCircuitGroupAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableCircuitGroupAssignmentRequest{}

// PatchedWritableCircuitGroupAssignmentRequest Base serializer for group assignments under CircuitSerializer.
type PatchedWritableCircuitGroupAssignmentRequest struct {
	Group *BriefCircuitGroupRequest `json:"group,omitempty"`
	Circuit *BriefCircuitRequest `json:"circuit,omitempty"`
	Priority *BriefCircuitGroupAssignmentSerializerPriorityValue `json:"priority,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableCircuitGroupAssignmentRequest PatchedWritableCircuitGroupAssignmentRequest

// NewPatchedWritableCircuitGroupAssignmentRequest instantiates a new PatchedWritableCircuitGroupAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableCircuitGroupAssignmentRequest() *PatchedWritableCircuitGroupAssignmentRequest {
	this := PatchedWritableCircuitGroupAssignmentRequest{}
	return &this
}

// NewPatchedWritableCircuitGroupAssignmentRequestWithDefaults instantiates a new PatchedWritableCircuitGroupAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableCircuitGroupAssignmentRequestWithDefaults() *PatchedWritableCircuitGroupAssignmentRequest {
	this := PatchedWritableCircuitGroupAssignmentRequest{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PatchedWritableCircuitGroupAssignmentRequest) GetGroup() BriefCircuitGroupRequest {
	if o == nil || IsNil(o.Group) {
		var ret BriefCircuitGroupRequest
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitGroupAssignmentRequest) GetGroupOk() (*BriefCircuitGroupRequest, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedWritableCircuitGroupAssignmentRequest) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given BriefCircuitGroupRequest and assigns it to the Group field.
func (o *PatchedWritableCircuitGroupAssignmentRequest) SetGroup(v BriefCircuitGroupRequest) {
	o.Group = &v
}

// GetCircuit returns the Circuit field value if set, zero value otherwise.
func (o *PatchedWritableCircuitGroupAssignmentRequest) GetCircuit() BriefCircuitRequest {
	if o == nil || IsNil(o.Circuit) {
		var ret BriefCircuitRequest
		return ret
	}
	return *o.Circuit
}

// GetCircuitOk returns a tuple with the Circuit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitGroupAssignmentRequest) GetCircuitOk() (*BriefCircuitRequest, bool) {
	if o == nil || IsNil(o.Circuit) {
		return nil, false
	}
	return o.Circuit, true
}

// HasCircuit returns a boolean if a field has been set.
func (o *PatchedWritableCircuitGroupAssignmentRequest) HasCircuit() bool {
	if o != nil && !IsNil(o.Circuit) {
		return true
	}

	return false
}

// SetCircuit gets a reference to the given BriefCircuitRequest and assigns it to the Circuit field.
func (o *PatchedWritableCircuitGroupAssignmentRequest) SetCircuit(v BriefCircuitRequest) {
	o.Circuit = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *PatchedWritableCircuitGroupAssignmentRequest) GetPriority() BriefCircuitGroupAssignmentSerializerPriorityValue {
	if o == nil || IsNil(o.Priority) {
		var ret BriefCircuitGroupAssignmentSerializerPriorityValue
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitGroupAssignmentRequest) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *PatchedWritableCircuitGroupAssignmentRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given BriefCircuitGroupAssignmentSerializerPriorityValue and assigns it to the Priority field.
func (o *PatchedWritableCircuitGroupAssignmentRequest) SetPriority(v BriefCircuitGroupAssignmentSerializerPriorityValue) {
	o.Priority = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableCircuitGroupAssignmentRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitGroupAssignmentRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableCircuitGroupAssignmentRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableCircuitGroupAssignmentRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

func (o PatchedWritableCircuitGroupAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableCircuitGroupAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Circuit) {
		toSerialize["circuit"] = o.Circuit
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableCircuitGroupAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableCircuitGroupAssignmentRequest := _PatchedWritableCircuitGroupAssignmentRequest{}

	err = json.Unmarshal(data, &varPatchedWritableCircuitGroupAssignmentRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableCircuitGroupAssignmentRequest(varPatchedWritableCircuitGroupAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "circuit")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableCircuitGroupAssignmentRequest struct {
	value *PatchedWritableCircuitGroupAssignmentRequest
	isSet bool
}

func (v NullablePatchedWritableCircuitGroupAssignmentRequest) Get() *PatchedWritableCircuitGroupAssignmentRequest {
	return v.value
}

func (v *NullablePatchedWritableCircuitGroupAssignmentRequest) Set(val *PatchedWritableCircuitGroupAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCircuitGroupAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCircuitGroupAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCircuitGroupAssignmentRequest(val *PatchedWritableCircuitGroupAssignmentRequest) *NullablePatchedWritableCircuitGroupAssignmentRequest {
	return &NullablePatchedWritableCircuitGroupAssignmentRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableCircuitGroupAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCircuitGroupAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



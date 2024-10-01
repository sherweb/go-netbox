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

// checks if the PatchedWritableEventRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableEventRuleRequest{}

// PatchedWritableEventRuleRequest Adds support for custom fields and tags.
type PatchedWritableEventRuleRequest struct {
	ObjectTypes []string `json:"object_types,omitempty"`
	Name *string `json:"name,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// The types of event which will trigger this rule.
	EventTypes []EventRuleEventTypesInner `json:"event_types,omitempty"`
	// A set of conditions which determine whether the event will be generated.
	Conditions interface{} `json:"conditions,omitempty"`
	ActionType *EventRuleActionTypeValue `json:"action_type,omitempty"`
	ActionObjectType *string `json:"action_object_type,omitempty"`
	ActionObjectId NullableInt64 `json:"action_object_id,omitempty"`
	Description *string `json:"description,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableEventRuleRequest PatchedWritableEventRuleRequest

// NewPatchedWritableEventRuleRequest instantiates a new PatchedWritableEventRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableEventRuleRequest() *PatchedWritableEventRuleRequest {
	this := PatchedWritableEventRuleRequest{}
	return &this
}

// NewPatchedWritableEventRuleRequestWithDefaults instantiates a new PatchedWritableEventRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableEventRuleRequestWithDefaults() *PatchedWritableEventRuleRequest {
	this := PatchedWritableEventRuleRequest{}
	return &this
}

// GetObjectTypes returns the ObjectTypes field value if set, zero value otherwise.
func (o *PatchedWritableEventRuleRequest) GetObjectTypes() []string {
	if o == nil || IsNil(o.ObjectTypes) {
		var ret []string
		return ret
	}
	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableEventRuleRequest) GetObjectTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ObjectTypes) {
		return nil, false
	}
	return o.ObjectTypes, true
}

// HasObjectTypes returns a boolean if a field has been set.
func (o *PatchedWritableEventRuleRequest) HasObjectTypes() bool {
	if o != nil && !IsNil(o.ObjectTypes) {
		return true
	}

	return false
}

// SetObjectTypes gets a reference to the given []string and assigns it to the ObjectTypes field.
func (o *PatchedWritableEventRuleRequest) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableEventRuleRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableEventRuleRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableEventRuleRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableEventRuleRequest) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedWritableEventRuleRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableEventRuleRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedWritableEventRuleRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedWritableEventRuleRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *PatchedWritableEventRuleRequest) GetEventTypes() []EventRuleEventTypesInner {
	if o == nil || IsNil(o.EventTypes) {
		var ret []EventRuleEventTypesInner
		return ret
	}
	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableEventRuleRequest) GetEventTypesOk() ([]EventRuleEventTypesInner, bool) {
	if o == nil || IsNil(o.EventTypes) {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *PatchedWritableEventRuleRequest) HasEventTypes() bool {
	if o != nil && !IsNil(o.EventTypes) {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []EventRuleEventTypesInner and assigns it to the EventTypes field.
func (o *PatchedWritableEventRuleRequest) SetEventTypes(v []EventRuleEventTypesInner) {
	o.EventTypes = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableEventRuleRequest) GetConditions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableEventRuleRequest) GetConditionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return &o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *PatchedWritableEventRuleRequest) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given interface{} and assigns it to the Conditions field.
func (o *PatchedWritableEventRuleRequest) SetConditions(v interface{}) {
	o.Conditions = v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *PatchedWritableEventRuleRequest) GetActionType() EventRuleActionTypeValue {
	if o == nil || IsNil(o.ActionType) {
		var ret EventRuleActionTypeValue
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableEventRuleRequest) GetActionTypeOk() (*EventRuleActionTypeValue, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *PatchedWritableEventRuleRequest) HasActionType() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given EventRuleActionTypeValue and assigns it to the ActionType field.
func (o *PatchedWritableEventRuleRequest) SetActionType(v EventRuleActionTypeValue) {
	o.ActionType = &v
}

// GetActionObjectType returns the ActionObjectType field value if set, zero value otherwise.
func (o *PatchedWritableEventRuleRequest) GetActionObjectType() string {
	if o == nil || IsNil(o.ActionObjectType) {
		var ret string
		return ret
	}
	return *o.ActionObjectType
}

// GetActionObjectTypeOk returns a tuple with the ActionObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableEventRuleRequest) GetActionObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActionObjectType) {
		return nil, false
	}
	return o.ActionObjectType, true
}

// HasActionObjectType returns a boolean if a field has been set.
func (o *PatchedWritableEventRuleRequest) HasActionObjectType() bool {
	if o != nil && !IsNil(o.ActionObjectType) {
		return true
	}

	return false
}

// SetActionObjectType gets a reference to the given string and assigns it to the ActionObjectType field.
func (o *PatchedWritableEventRuleRequest) SetActionObjectType(v string) {
	o.ActionObjectType = &v
}

// GetActionObjectId returns the ActionObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableEventRuleRequest) GetActionObjectId() int64 {
	if o == nil || IsNil(o.ActionObjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.ActionObjectId.Get()
}

// GetActionObjectIdOk returns a tuple with the ActionObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableEventRuleRequest) GetActionObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionObjectId.Get(), o.ActionObjectId.IsSet()
}

// HasActionObjectId returns a boolean if a field has been set.
func (o *PatchedWritableEventRuleRequest) HasActionObjectId() bool {
	if o != nil && o.ActionObjectId.IsSet() {
		return true
	}

	return false
}

// SetActionObjectId gets a reference to the given NullableInt64 and assigns it to the ActionObjectId field.
func (o *PatchedWritableEventRuleRequest) SetActionObjectId(v int64) {
	o.ActionObjectId.Set(&v)
}
// SetActionObjectIdNil sets the value for ActionObjectId to be an explicit nil
func (o *PatchedWritableEventRuleRequest) SetActionObjectIdNil() {
	o.ActionObjectId.Set(nil)
}

// UnsetActionObjectId ensures that no value is present for ActionObjectId, not even an explicit nil
func (o *PatchedWritableEventRuleRequest) UnsetActionObjectId() {
	o.ActionObjectId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableEventRuleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableEventRuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableEventRuleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableEventRuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableEventRuleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableEventRuleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableEventRuleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableEventRuleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableEventRuleRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableEventRuleRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableEventRuleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableEventRuleRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

func (o PatchedWritableEventRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableEventRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectTypes) {
		toSerialize["object_types"] = o.ObjectTypes
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.EventTypes) {
		toSerialize["event_types"] = o.EventTypes
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.ActionType) {
		toSerialize["action_type"] = o.ActionType
	}
	if !IsNil(o.ActionObjectType) {
		toSerialize["action_object_type"] = o.ActionObjectType
	}
	if o.ActionObjectId.IsSet() {
		toSerialize["action_object_id"] = o.ActionObjectId.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableEventRuleRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableEventRuleRequest := _PatchedWritableEventRuleRequest{}

	err = json.Unmarshal(data, &varPatchedWritableEventRuleRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableEventRuleRequest(varPatchedWritableEventRuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "event_types")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "action_type")
		delete(additionalProperties, "action_object_type")
		delete(additionalProperties, "action_object_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableEventRuleRequest struct {
	value *PatchedWritableEventRuleRequest
	isSet bool
}

func (v NullablePatchedWritableEventRuleRequest) Get() *PatchedWritableEventRuleRequest {
	return v.value
}

func (v *NullablePatchedWritableEventRuleRequest) Set(val *PatchedWritableEventRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableEventRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableEventRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableEventRuleRequest(val *PatchedWritableEventRuleRequest) *NullablePatchedWritableEventRuleRequest {
	return &NullablePatchedWritableEventRuleRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableEventRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableEventRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



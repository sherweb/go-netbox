/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.2 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the L2VPNTermination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &L2VPNTermination{}

// L2VPNTermination Adds support for custom fields and tags.
type L2VPNTermination struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl *string `json:"display_url,omitempty"`
	Display string `json:"display"`
	L2vpn BriefL2VPN `json:"l2vpn"`
	AssignedObjectType string `json:"assigned_object_type"`
	AssignedObjectId int64 `json:"assigned_object_id"`
	AssignedObject interface{} `json:"assigned_object"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _L2VPNTermination L2VPNTermination

// NewL2VPNTermination instantiates a new L2VPNTermination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewL2VPNTermination(id int32, url string, display string, l2vpn BriefL2VPN, assignedObjectType string, assignedObjectId int64, assignedObject interface{}, created NullableTime, lastUpdated NullableTime) *L2VPNTermination {
	this := L2VPNTermination{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.L2vpn = l2vpn
	this.AssignedObjectType = assignedObjectType
	this.AssignedObjectId = assignedObjectId
	this.AssignedObject = assignedObject
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewL2VPNTerminationWithDefaults instantiates a new L2VPNTermination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewL2VPNTerminationWithDefaults() *L2VPNTermination {
	this := L2VPNTermination{}
	return &this
}

// GetId returns the Id field value
func (o *L2VPNTermination) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *L2VPNTermination) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *L2VPNTermination) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *L2VPNTermination) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *L2VPNTermination) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *L2VPNTermination) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value if set, zero value otherwise.
func (o *L2VPNTermination) GetDisplayUrl() string {
	if o == nil || IsNil(o.DisplayUrl) {
		var ret string
		return ret
	}
	return *o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNTermination) GetDisplayUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayUrl) {
		return nil, false
	}
	return o.DisplayUrl, true
}

// HasDisplayUrl returns a boolean if a field has been set.
func (o *L2VPNTermination) HasDisplayUrl() bool {
	if o != nil && !IsNil(o.DisplayUrl) {
		return true
	}

	return false
}

// SetDisplayUrl gets a reference to the given string and assigns it to the DisplayUrl field.
func (o *L2VPNTermination) SetDisplayUrl(v string) {
	o.DisplayUrl = &v
}

// GetDisplay returns the Display field value
func (o *L2VPNTermination) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *L2VPNTermination) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *L2VPNTermination) SetDisplay(v string) {
	o.Display = v
}

// GetL2vpn returns the L2vpn field value
func (o *L2VPNTermination) GetL2vpn() BriefL2VPN {
	if o == nil {
		var ret BriefL2VPN
		return ret
	}

	return o.L2vpn
}

// GetL2vpnOk returns a tuple with the L2vpn field value
// and a boolean to check if the value has been set.
func (o *L2VPNTermination) GetL2vpnOk() (*BriefL2VPN, bool) {
	if o == nil {
		return nil, false
	}
	return &o.L2vpn, true
}

// SetL2vpn sets field value
func (o *L2VPNTermination) SetL2vpn(v BriefL2VPN) {
	o.L2vpn = v
}

// GetAssignedObjectType returns the AssignedObjectType field value
func (o *L2VPNTermination) GetAssignedObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedObjectType
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value
// and a boolean to check if the value has been set.
func (o *L2VPNTermination) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectType, true
}

// SetAssignedObjectType sets field value
func (o *L2VPNTermination) SetAssignedObjectType(v string) {
	o.AssignedObjectType = v
}

// GetAssignedObjectId returns the AssignedObjectId field value
func (o *L2VPNTermination) GetAssignedObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AssignedObjectId
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value
// and a boolean to check if the value has been set.
func (o *L2VPNTermination) GetAssignedObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectId, true
}

// SetAssignedObjectId sets field value
func (o *L2VPNTermination) SetAssignedObjectId(v int64) {
	o.AssignedObjectId = v
}

// GetAssignedObject returns the AssignedObject field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *L2VPNTermination) GetAssignedObject() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.AssignedObject
}

// GetAssignedObjectOk returns a tuple with the AssignedObject field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *L2VPNTermination) GetAssignedObjectOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AssignedObject) {
		return nil, false
	}
	return &o.AssignedObject, true
}

// SetAssignedObject sets field value
func (o *L2VPNTermination) SetAssignedObject(v interface{}) {
	o.AssignedObject = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *L2VPNTermination) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNTermination) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *L2VPNTermination) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *L2VPNTermination) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *L2VPNTermination) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2VPNTermination) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *L2VPNTermination) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *L2VPNTermination) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *L2VPNTermination) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *L2VPNTermination) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *L2VPNTermination) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *L2VPNTermination) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *L2VPNTermination) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *L2VPNTermination) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o L2VPNTermination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o L2VPNTermination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	if !IsNil(o.DisplayUrl) {
		toSerialize["display_url"] = o.DisplayUrl
	}
	toSerialize["display"] = o.Display
	toSerialize["l2vpn"] = o.L2vpn
	toSerialize["assigned_object_type"] = o.AssignedObjectType
	toSerialize["assigned_object_id"] = o.AssignedObjectId
	if o.AssignedObject != nil {
		toSerialize["assigned_object"] = o.AssignedObject
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *L2VPNTermination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"l2vpn",
		"assigned_object_type",
		"assigned_object_id",
		"assigned_object",
		"created",
		"last_updated",
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

	varL2VPNTermination := _L2VPNTermination{}

	err = json.Unmarshal(data, &varL2VPNTermination)

	if err != nil {
		return err
	}

	*o = L2VPNTermination(varL2VPNTermination)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "l2vpn")
		delete(additionalProperties, "assigned_object_type")
		delete(additionalProperties, "assigned_object_id")
		delete(additionalProperties, "assigned_object")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableL2VPNTermination struct {
	value *L2VPNTermination
	isSet bool
}

func (v NullableL2VPNTermination) Get() *L2VPNTermination {
	return v.value
}

func (v *NullableL2VPNTermination) Set(val *L2VPNTermination) {
	v.value = val
	v.isSet = true
}

func (v NullableL2VPNTermination) IsSet() bool {
	return v.isSet
}

func (v *NullableL2VPNTermination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL2VPNTermination(val *L2VPNTermination) *NullableL2VPNTermination {
	return &NullableL2VPNTermination{value: val, isSet: true}
}

func (v NullableL2VPNTermination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL2VPNTermination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



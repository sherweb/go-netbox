/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.3 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the DeviceRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceRole{}

// DeviceRole Adds support for custom fields and tags.
type DeviceRole struct {
	Id                   int32   `json:"id"`
	Url                  string  `json:"url"`
	Display              string  `json:"display"`
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug"`
	Description          *string `json:"description,omitempty"`
	DeviceCount          *int64  `json:"device_count,omitempty"`
	VirtualmachineCount  *int64  `json:"virtualmachine_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceRole DeviceRole

// NewDeviceRole instantiates a new DeviceRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceRole(id int32, url string, display string, name string, slug string) *DeviceRole {
	this := DeviceRole{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	return &this
}

// NewDeviceRoleWithDefaults instantiates a new DeviceRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceRoleWithDefaults() *DeviceRole {
	this := DeviceRole{}
	return &this
}

// GetId returns the Id field value
func (o *DeviceRole) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceRole) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *DeviceRole) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DeviceRole) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *DeviceRole) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *DeviceRole) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *DeviceRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceRole) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *DeviceRole) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *DeviceRole) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceRole) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceRole) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceRole) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *DeviceRole) GetDeviceCount() int64 {
	if o == nil || IsNil(o.DeviceCount) {
		var ret int64
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetDeviceCountOk() (*int64, bool) {
	if o == nil || IsNil(o.DeviceCount) {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *DeviceRole) HasDeviceCount() bool {
	if o != nil && !IsNil(o.DeviceCount) {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int64 and assigns it to the DeviceCount field.
func (o *DeviceRole) SetDeviceCount(v int64) {
	o.DeviceCount = &v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value if set, zero value otherwise.
func (o *DeviceRole) GetVirtualmachineCount() int64 {
	if o == nil || IsNil(o.VirtualmachineCount) {
		var ret int64
		return ret
	}
	return *o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetVirtualmachineCountOk() (*int64, bool) {
	if o == nil || IsNil(o.VirtualmachineCount) {
		return nil, false
	}
	return o.VirtualmachineCount, true
}

// HasVirtualmachineCount returns a boolean if a field has been set.
func (o *DeviceRole) HasVirtualmachineCount() bool {
	if o != nil && !IsNil(o.VirtualmachineCount) {
		return true
	}

	return false
}

// SetVirtualmachineCount gets a reference to the given int64 and assigns it to the VirtualmachineCount field.
func (o *DeviceRole) SetVirtualmachineCount(v int64) {
	o.VirtualmachineCount = &v
}

func (o DeviceRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DeviceCount) {
		toSerialize["device_count"] = o.DeviceCount
	}
	if !IsNil(o.VirtualmachineCount) {
		toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"slug",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeviceRole := _DeviceRole{}

	err = json.Unmarshal(data, &varDeviceRole)

	if err != nil {
		return err
	}

	*o = DeviceRole(varDeviceRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "device_count")
		delete(additionalProperties, "virtualmachine_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceRole struct {
	value *DeviceRole
	isSet bool
}

func (v NullableDeviceRole) Get() *DeviceRole {
	return v.value
}

func (v *NullableDeviceRole) Set(val *DeviceRole) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceRole) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceRole(val *DeviceRole) *NullableDeviceRole {
	return &NullableDeviceRole{value: val, isSet: true}
}

func (v NullableDeviceRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

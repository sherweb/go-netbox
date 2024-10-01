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

// checks if the BriefModuleType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefModuleType{}

// BriefModuleType Adds support for custom fields and tags.
type BriefModuleType struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Manufacturer BriefManufacturer `json:"manufacturer"`
	Model string `json:"model"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefModuleType BriefModuleType

// NewBriefModuleType instantiates a new BriefModuleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefModuleType(id int32, url string, display string, manufacturer BriefManufacturer, model string) *BriefModuleType {
	this := BriefModuleType{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Manufacturer = manufacturer
	this.Model = model
	return &this
}

// NewBriefModuleTypeWithDefaults instantiates a new BriefModuleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefModuleTypeWithDefaults() *BriefModuleType {
	this := BriefModuleType{}
	return &this
}

// GetId returns the Id field value
func (o *BriefModuleType) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefModuleType) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefModuleType) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefModuleType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefModuleType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefModuleType) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefModuleType) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefModuleType) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefModuleType) SetDisplay(v string) {
	o.Display = v
}

// GetManufacturer returns the Manufacturer field value
func (o *BriefModuleType) GetManufacturer() BriefManufacturer {
	if o == nil {
		var ret BriefManufacturer
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *BriefModuleType) GetManufacturerOk() (*BriefManufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *BriefModuleType) SetManufacturer(v BriefManufacturer) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *BriefModuleType) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *BriefModuleType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *BriefModuleType) SetModel(v string) {
	o.Model = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefModuleType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefModuleType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefModuleType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefModuleType) SetDescription(v string) {
	o.Description = &v
}

func (o BriefModuleType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefModuleType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["manufacturer"] = o.Manufacturer
	toSerialize["model"] = o.Model
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefModuleType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"manufacturer",
		"model",
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

	varBriefModuleType := _BriefModuleType{}

	err = json.Unmarshal(data, &varBriefModuleType)

	if err != nil {
		return err
	}

	*o = BriefModuleType(varBriefModuleType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "model")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefModuleType struct {
	value *BriefModuleType
	isSet bool
}

func (v NullableBriefModuleType) Get() *BriefModuleType {
	return v.value
}

func (v *NullableBriefModuleType) Set(val *BriefModuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefModuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefModuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefModuleType(val *BriefModuleType) *NullableBriefModuleType {
	return &NullableBriefModuleType{value: val, isSet: true}
}

func (v NullableBriefModuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefModuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



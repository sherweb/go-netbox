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

// checks if the BriefCustomFieldChoiceSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefCustomFieldChoiceSet{}

// BriefCustomFieldChoiceSet Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type BriefCustomFieldChoiceSet struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	ChoicesCount string `json:"choices_count"`
	AdditionalProperties map[string]interface{}
}

type _BriefCustomFieldChoiceSet BriefCustomFieldChoiceSet

// NewBriefCustomFieldChoiceSet instantiates a new BriefCustomFieldChoiceSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefCustomFieldChoiceSet(id int32, url string, display string, name string, choicesCount string) *BriefCustomFieldChoiceSet {
	this := BriefCustomFieldChoiceSet{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.ChoicesCount = choicesCount
	return &this
}

// NewBriefCustomFieldChoiceSetWithDefaults instantiates a new BriefCustomFieldChoiceSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefCustomFieldChoiceSetWithDefaults() *BriefCustomFieldChoiceSet {
	this := BriefCustomFieldChoiceSet{}
	return &this
}

// GetId returns the Id field value
func (o *BriefCustomFieldChoiceSet) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefCustomFieldChoiceSet) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefCustomFieldChoiceSet) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefCustomFieldChoiceSet) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefCustomFieldChoiceSet) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefCustomFieldChoiceSet) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefCustomFieldChoiceSet) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefCustomFieldChoiceSet) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefCustomFieldChoiceSet) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *BriefCustomFieldChoiceSet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefCustomFieldChoiceSet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefCustomFieldChoiceSet) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefCustomFieldChoiceSet) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefCustomFieldChoiceSet) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefCustomFieldChoiceSet) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefCustomFieldChoiceSet) SetDescription(v string) {
	o.Description = &v
}

// GetChoicesCount returns the ChoicesCount field value
func (o *BriefCustomFieldChoiceSet) GetChoicesCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChoicesCount
}

// GetChoicesCountOk returns a tuple with the ChoicesCount field value
// and a boolean to check if the value has been set.
func (o *BriefCustomFieldChoiceSet) GetChoicesCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChoicesCount, true
}

// SetChoicesCount sets field value
func (o *BriefCustomFieldChoiceSet) SetChoicesCount(v string) {
	o.ChoicesCount = v
}

func (o BriefCustomFieldChoiceSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefCustomFieldChoiceSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["choices_count"] = o.ChoicesCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefCustomFieldChoiceSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"choices_count",
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

	varBriefCustomFieldChoiceSet := _BriefCustomFieldChoiceSet{}

	err = json.Unmarshal(data, &varBriefCustomFieldChoiceSet)

	if err != nil {
		return err
	}

	*o = BriefCustomFieldChoiceSet(varBriefCustomFieldChoiceSet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "choices_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefCustomFieldChoiceSet struct {
	value *BriefCustomFieldChoiceSet
	isSet bool
}

func (v NullableBriefCustomFieldChoiceSet) Get() *BriefCustomFieldChoiceSet {
	return v.value
}

func (v *NullableBriefCustomFieldChoiceSet) Set(val *BriefCustomFieldChoiceSet) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefCustomFieldChoiceSet) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefCustomFieldChoiceSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefCustomFieldChoiceSet(val *BriefCustomFieldChoiceSet) *NullableBriefCustomFieldChoiceSet {
	return &NullableBriefCustomFieldChoiceSet{value: val, isSet: true}
}

func (v NullableBriefCustomFieldChoiceSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefCustomFieldChoiceSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



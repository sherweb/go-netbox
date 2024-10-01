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

// checks if the AvailablePrefix type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailablePrefix{}

// AvailablePrefix Representation of a prefix which does not exist in the database.
type AvailablePrefix struct {
	Family int32 `json:"family"`
	Prefix string `json:"prefix"`
	Vrf NullableBriefVRF `json:"vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AvailablePrefix AvailablePrefix

// NewAvailablePrefix instantiates a new AvailablePrefix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailablePrefix(family int32, prefix string) *AvailablePrefix {
	this := AvailablePrefix{}
	this.Family = family
	this.Prefix = prefix
	return &this
}

// NewAvailablePrefixWithDefaults instantiates a new AvailablePrefix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailablePrefixWithDefaults() *AvailablePrefix {
	this := AvailablePrefix{}
	return &this
}

// GetFamily returns the Family field value
func (o *AvailablePrefix) GetFamily() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *AvailablePrefix) GetFamilyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *AvailablePrefix) SetFamily(v int32) {
	o.Family = v
}


// GetPrefix returns the Prefix field value
func (o *AvailablePrefix) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *AvailablePrefix) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *AvailablePrefix) SetPrefix(v string) {
	o.Prefix = v
}


// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailablePrefix) GetVrf() BriefVRF {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret BriefVRF
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailablePrefix) GetVrfOk() (*BriefVRF, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *AvailablePrefix) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableBriefVRF and assigns it to the Vrf field.
func (o *AvailablePrefix) SetVrf(v BriefVRF) {
	o.Vrf.Set(&v)
}
// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *AvailablePrefix) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *AvailablePrefix) UnsetVrf() {
	o.Vrf.Unset()
}

func (o AvailablePrefix) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailablePrefix) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["family"] = o.Family
	toSerialize["prefix"] = o.Prefix
	if o.Vrf.IsSet() {
		toSerialize["vrf"] = o.Vrf.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AvailablePrefix) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"family",
		"prefix",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{} {
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == ""{
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil{
			return err
		}
	}
	varAvailablePrefix := _AvailablePrefix{}

	err = json.Unmarshal(data, &varAvailablePrefix)

	if err != nil {
		return err
	}

	*o = AvailablePrefix(varAvailablePrefix)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "family")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "vrf")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAvailablePrefix struct {
	value *AvailablePrefix
	isSet bool
}

func (v NullableAvailablePrefix) Get() *AvailablePrefix {
	return v.value
}

func (v *NullableAvailablePrefix) Set(val *AvailablePrefix) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailablePrefix) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailablePrefix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailablePrefix(val *AvailablePrefix) *NullableAvailablePrefix {
	return &NullableAvailablePrefix{value: val, isSet: true}
}

func (v NullableAvailablePrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailablePrefix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



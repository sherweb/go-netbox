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

// WirelessRole1 * `ap` - Access point * `station` - Station
type WirelessRole1 string

// List of Wireless_role_1
const (
	WIRELESSROLE1_AP WirelessRole1 = "ap"
	WIRELESSROLE1_STATION WirelessRole1 = "station"
	WIRELESSROLE1_EMPTY WirelessRole1 = ""
)

// All allowed values of WirelessRole1 enum
var AllowedWirelessRole1EnumValues = []WirelessRole1{
	"ap",
	"station",
	"",
}

func (v *WirelessRole1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WirelessRole1(value)
	for _, existing := range AllowedWirelessRole1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WirelessRole1", value)
}

// NewWirelessRole1FromValue returns a pointer to a valid WirelessRole1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWirelessRole1FromValue(v string) (*WirelessRole1, error) {
	ev := WirelessRole1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WirelessRole1: valid values are %v", v, AllowedWirelessRole1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WirelessRole1) IsValid() bool {
	for _, existing := range AllowedWirelessRole1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Wireless_role_1 value
func (v WirelessRole1) Ptr() *WirelessRole1 {
	return &v
}

type NullableWirelessRole1 struct {
	value *WirelessRole1
	isSet bool
}

func (v NullableWirelessRole1) Get() *WirelessRole1 {
	return v.value
}

func (v *NullableWirelessRole1) Set(val *WirelessRole1) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessRole1) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessRole1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessRole1(val *WirelessRole1) *NullableWirelessRole1 {
	return &NullableWirelessRole1{value: val, isSet: true}
}

func (v NullableWirelessRole1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessRole1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


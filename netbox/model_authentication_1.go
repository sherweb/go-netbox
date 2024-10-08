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

// Authentication1 * `hmac-sha1` - SHA-1 HMAC * `hmac-sha256` - SHA-256 HMAC * `hmac-sha384` - SHA-384 HMAC * `hmac-sha512` - SHA-512 HMAC * `hmac-md5` - MD5 HMAC
type Authentication1 string

// List of Authentication_1
const (
	AUTHENTICATION1_HMAC_SHA1 Authentication1 = "hmac-sha1"
	AUTHENTICATION1_HMAC_SHA256 Authentication1 = "hmac-sha256"
	AUTHENTICATION1_HMAC_SHA384 Authentication1 = "hmac-sha384"
	AUTHENTICATION1_HMAC_SHA512 Authentication1 = "hmac-sha512"
	AUTHENTICATION1_HMAC_MD5 Authentication1 = "hmac-md5"
	AUTHENTICATION1_EMPTY Authentication1 = ""
)

// All allowed values of Authentication1 enum
var AllowedAuthentication1EnumValues = []Authentication1{
	"hmac-sha1",
	"hmac-sha256",
	"hmac-sha384",
	"hmac-sha512",
	"hmac-md5",
	"",
}

func (v *Authentication1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Authentication1(value)
	for _, existing := range AllowedAuthentication1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Authentication1", value)
}

// NewAuthentication1FromValue returns a pointer to a valid Authentication1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthentication1FromValue(v string) (*Authentication1, error) {
	ev := Authentication1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Authentication1: valid values are %v", v, AllowedAuthentication1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Authentication1) IsValid() bool {
	for _, existing := range AllowedAuthentication1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Authentication_1 value
func (v Authentication1) Ptr() *Authentication1 {
	return &v
}

type NullableAuthentication1 struct {
	value *Authentication1
	isSet bool
}

func (v NullableAuthentication1) Get() *Authentication1 {
	return v.value
}

func (v *NullableAuthentication1) Set(val *Authentication1) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthentication1) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthentication1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthentication1(val *Authentication1) *NullableAuthentication1 {
	return &NullableAuthentication1{value: val, isSet: true}
}

func (v NullableAuthentication1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthentication1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


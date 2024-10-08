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

// AuthenticationType2 * `plaintext` - Plaintext * `md5` - MD5
type AuthenticationType2 string

// List of Authentication_type_2
const (
	AUTHENTICATIONTYPE2_PLAINTEXT AuthenticationType2 = "plaintext"
	AUTHENTICATIONTYPE2_MD5 AuthenticationType2 = "md5"
	AUTHENTICATIONTYPE2_EMPTY AuthenticationType2 = ""
)

// All allowed values of AuthenticationType2 enum
var AllowedAuthenticationType2EnumValues = []AuthenticationType2{
	"plaintext",
	"md5",
	"",
}

func (v *AuthenticationType2) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthenticationType2(value)
	for _, existing := range AllowedAuthenticationType2EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthenticationType2", value)
}

// NewAuthenticationType2FromValue returns a pointer to a valid AuthenticationType2
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthenticationType2FromValue(v string) (*AuthenticationType2, error) {
	ev := AuthenticationType2(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthenticationType2: valid values are %v", v, AllowedAuthenticationType2EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthenticationType2) IsValid() bool {
	for _, existing := range AllowedAuthenticationType2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Authentication_type_2 value
func (v AuthenticationType2) Ptr() *AuthenticationType2 {
	return &v
}

type NullableAuthenticationType2 struct {
	value *AuthenticationType2
	isSet bool
}

func (v NullableAuthenticationType2) Get() *AuthenticationType2 {
	return v.value
}

func (v *NullableAuthenticationType2) Set(val *AuthenticationType2) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationType2) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationType2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationType2(val *AuthenticationType2) *NullableAuthenticationType2 {
	return &NullableAuthenticationType2{value: val, isSet: true}
}

func (v NullableAuthenticationType2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationType2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// VpnIkeProposalsListEncryptionAlgorithmIcParameterInner the model 'VpnIkeProposalsListEncryptionAlgorithmIcParameterInner'
type VpnIkeProposalsListEncryptionAlgorithmIcParameterInner string

// List of vpn_ike_proposals_list_encryption_algorithm__ic_parameter_inner
const (
	VPNIKEPROPOSALSLISTENCRYPTIONALGORITHMICPARAMETERINNER__3DES_CBC VpnIkeProposalsListEncryptionAlgorithmIcParameterInner = "3des-cbc"
	VPNIKEPROPOSALSLISTENCRYPTIONALGORITHMICPARAMETERINNER_AES_128_CBC VpnIkeProposalsListEncryptionAlgorithmIcParameterInner = "aes-128-cbc"
	VPNIKEPROPOSALSLISTENCRYPTIONALGORITHMICPARAMETERINNER_AES_128_GCM VpnIkeProposalsListEncryptionAlgorithmIcParameterInner = "aes-128-gcm"
	VPNIKEPROPOSALSLISTENCRYPTIONALGORITHMICPARAMETERINNER_AES_192_CBC VpnIkeProposalsListEncryptionAlgorithmIcParameterInner = "aes-192-cbc"
	VPNIKEPROPOSALSLISTENCRYPTIONALGORITHMICPARAMETERINNER_AES_192_GCM VpnIkeProposalsListEncryptionAlgorithmIcParameterInner = "aes-192-gcm"
	VPNIKEPROPOSALSLISTENCRYPTIONALGORITHMICPARAMETERINNER_AES_256_CBC VpnIkeProposalsListEncryptionAlgorithmIcParameterInner = "aes-256-cbc"
	VPNIKEPROPOSALSLISTENCRYPTIONALGORITHMICPARAMETERINNER_AES_256_GCM VpnIkeProposalsListEncryptionAlgorithmIcParameterInner = "aes-256-gcm"
	VPNIKEPROPOSALSLISTENCRYPTIONALGORITHMICPARAMETERINNER_DES_CBC VpnIkeProposalsListEncryptionAlgorithmIcParameterInner = "des-cbc"
)

// All allowed values of VpnIkeProposalsListEncryptionAlgorithmIcParameterInner enum
var AllowedVpnIkeProposalsListEncryptionAlgorithmIcParameterInnerEnumValues = []VpnIkeProposalsListEncryptionAlgorithmIcParameterInner{
	"3des-cbc",
	"aes-128-cbc",
	"aes-128-gcm",
	"aes-192-cbc",
	"aes-192-gcm",
	"aes-256-cbc",
	"aes-256-gcm",
	"des-cbc",
}

func (v *VpnIkeProposalsListEncryptionAlgorithmIcParameterInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VpnIkeProposalsListEncryptionAlgorithmIcParameterInner(value)
	for _, existing := range AllowedVpnIkeProposalsListEncryptionAlgorithmIcParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VpnIkeProposalsListEncryptionAlgorithmIcParameterInner", value)
}

// NewVpnIkeProposalsListEncryptionAlgorithmIcParameterInnerFromValue returns a pointer to a valid VpnIkeProposalsListEncryptionAlgorithmIcParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVpnIkeProposalsListEncryptionAlgorithmIcParameterInnerFromValue(v string) (*VpnIkeProposalsListEncryptionAlgorithmIcParameterInner, error) {
	ev := VpnIkeProposalsListEncryptionAlgorithmIcParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VpnIkeProposalsListEncryptionAlgorithmIcParameterInner: valid values are %v", v, AllowedVpnIkeProposalsListEncryptionAlgorithmIcParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VpnIkeProposalsListEncryptionAlgorithmIcParameterInner) IsValid() bool {
	for _, existing := range AllowedVpnIkeProposalsListEncryptionAlgorithmIcParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vpn_ike_proposals_list_encryption_algorithm__ic_parameter_inner value
func (v VpnIkeProposalsListEncryptionAlgorithmIcParameterInner) Ptr() *VpnIkeProposalsListEncryptionAlgorithmIcParameterInner {
	return &v
}

type NullableVpnIkeProposalsListEncryptionAlgorithmIcParameterInner struct {
	value *VpnIkeProposalsListEncryptionAlgorithmIcParameterInner
	isSet bool
}

func (v NullableVpnIkeProposalsListEncryptionAlgorithmIcParameterInner) Get() *VpnIkeProposalsListEncryptionAlgorithmIcParameterInner {
	return v.value
}

func (v *NullableVpnIkeProposalsListEncryptionAlgorithmIcParameterInner) Set(val *VpnIkeProposalsListEncryptionAlgorithmIcParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVpnIkeProposalsListEncryptionAlgorithmIcParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVpnIkeProposalsListEncryptionAlgorithmIcParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpnIkeProposalsListEncryptionAlgorithmIcParameterInner(val *VpnIkeProposalsListEncryptionAlgorithmIcParameterInner) *NullableVpnIkeProposalsListEncryptionAlgorithmIcParameterInner {
	return &NullableVpnIkeProposalsListEncryptionAlgorithmIcParameterInner{value: val, isSet: true}
}

func (v NullableVpnIkeProposalsListEncryptionAlgorithmIcParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpnIkeProposalsListEncryptionAlgorithmIcParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


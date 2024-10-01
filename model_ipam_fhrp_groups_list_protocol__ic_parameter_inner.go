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

// IpamFhrpGroupsListProtocolIcParameterInner the model 'IpamFhrpGroupsListProtocolIcParameterInner'
type IpamFhrpGroupsListProtocolIcParameterInner string

// List of ipam_fhrp_groups_list_protocol__ic_parameter_inner
const (
	IPAMFHRPGROUPSLISTPROTOCOLICPARAMETERINNER_CARP IpamFhrpGroupsListProtocolIcParameterInner = "carp"
	IPAMFHRPGROUPSLISTPROTOCOLICPARAMETERINNER_CLUSTERXL IpamFhrpGroupsListProtocolIcParameterInner = "clusterxl"
	IPAMFHRPGROUPSLISTPROTOCOLICPARAMETERINNER_GLBP IpamFhrpGroupsListProtocolIcParameterInner = "glbp"
	IPAMFHRPGROUPSLISTPROTOCOLICPARAMETERINNER_HSRP IpamFhrpGroupsListProtocolIcParameterInner = "hsrp"
	IPAMFHRPGROUPSLISTPROTOCOLICPARAMETERINNER_OTHER IpamFhrpGroupsListProtocolIcParameterInner = "other"
	IPAMFHRPGROUPSLISTPROTOCOLICPARAMETERINNER_VRRP2 IpamFhrpGroupsListProtocolIcParameterInner = "vrrp2"
	IPAMFHRPGROUPSLISTPROTOCOLICPARAMETERINNER_VRRP3 IpamFhrpGroupsListProtocolIcParameterInner = "vrrp3"
)

// All allowed values of IpamFhrpGroupsListProtocolIcParameterInner enum
var AllowedIpamFhrpGroupsListProtocolIcParameterInnerEnumValues = []IpamFhrpGroupsListProtocolIcParameterInner{
	"carp",
	"clusterxl",
	"glbp",
	"hsrp",
	"other",
	"vrrp2",
	"vrrp3",
}

func (v *IpamFhrpGroupsListProtocolIcParameterInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IpamFhrpGroupsListProtocolIcParameterInner(value)
	for _, existing := range AllowedIpamFhrpGroupsListProtocolIcParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IpamFhrpGroupsListProtocolIcParameterInner", value)
}

// NewIpamFhrpGroupsListProtocolIcParameterInnerFromValue returns a pointer to a valid IpamFhrpGroupsListProtocolIcParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIpamFhrpGroupsListProtocolIcParameterInnerFromValue(v string) (*IpamFhrpGroupsListProtocolIcParameterInner, error) {
	ev := IpamFhrpGroupsListProtocolIcParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IpamFhrpGroupsListProtocolIcParameterInner: valid values are %v", v, AllowedIpamFhrpGroupsListProtocolIcParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IpamFhrpGroupsListProtocolIcParameterInner) IsValid() bool {
	for _, existing := range AllowedIpamFhrpGroupsListProtocolIcParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ipam_fhrp_groups_list_protocol__ic_parameter_inner value
func (v IpamFhrpGroupsListProtocolIcParameterInner) Ptr() *IpamFhrpGroupsListProtocolIcParameterInner {
	return &v
}

type NullableIpamFhrpGroupsListProtocolIcParameterInner struct {
	value *IpamFhrpGroupsListProtocolIcParameterInner
	isSet bool
}

func (v NullableIpamFhrpGroupsListProtocolIcParameterInner) Get() *IpamFhrpGroupsListProtocolIcParameterInner {
	return v.value
}

func (v *NullableIpamFhrpGroupsListProtocolIcParameterInner) Set(val *IpamFhrpGroupsListProtocolIcParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamFhrpGroupsListProtocolIcParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamFhrpGroupsListProtocolIcParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamFhrpGroupsListProtocolIcParameterInner(val *IpamFhrpGroupsListProtocolIcParameterInner) *NullableIpamFhrpGroupsListProtocolIcParameterInner {
	return &NullableIpamFhrpGroupsListProtocolIcParameterInner{value: val, isSet: true}
}

func (v NullableIpamFhrpGroupsListProtocolIcParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamFhrpGroupsListProtocolIcParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


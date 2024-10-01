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

// PatchedWritableVirtualDeviceContextRequestStatus * `active` - Active * `planned` - Planned * `offline` - Offline
type PatchedWritableVirtualDeviceContextRequestStatus string

// List of PatchedWritableVirtualDeviceContextRequest_status
const (
	PATCHEDWRITABLEVIRTUALDEVICECONTEXTREQUESTSTATUS_ACTIVE PatchedWritableVirtualDeviceContextRequestStatus = "active"
	PATCHEDWRITABLEVIRTUALDEVICECONTEXTREQUESTSTATUS_PLANNED PatchedWritableVirtualDeviceContextRequestStatus = "planned"
	PATCHEDWRITABLEVIRTUALDEVICECONTEXTREQUESTSTATUS_OFFLINE PatchedWritableVirtualDeviceContextRequestStatus = "offline"
)

// All allowed values of PatchedWritableVirtualDeviceContextRequestStatus enum
var AllowedPatchedWritableVirtualDeviceContextRequestStatusEnumValues = []PatchedWritableVirtualDeviceContextRequestStatus{
	"active",
	"planned",
	"offline",
}

func (v *PatchedWritableVirtualDeviceContextRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableVirtualDeviceContextRequestStatus(value)
	for _, existing := range AllowedPatchedWritableVirtualDeviceContextRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableVirtualDeviceContextRequestStatus", value)
}

// NewPatchedWritableVirtualDeviceContextRequestStatusFromValue returns a pointer to a valid PatchedWritableVirtualDeviceContextRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableVirtualDeviceContextRequestStatusFromValue(v string) (*PatchedWritableVirtualDeviceContextRequestStatus, error) {
	ev := PatchedWritableVirtualDeviceContextRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableVirtualDeviceContextRequestStatus: valid values are %v", v, AllowedPatchedWritableVirtualDeviceContextRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableVirtualDeviceContextRequestStatus) IsValid() bool {
	for _, existing := range AllowedPatchedWritableVirtualDeviceContextRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableVirtualDeviceContextRequest_status value
func (v PatchedWritableVirtualDeviceContextRequestStatus) Ptr() *PatchedWritableVirtualDeviceContextRequestStatus {
	return &v
}

type NullablePatchedWritableVirtualDeviceContextRequestStatus struct {
	value *PatchedWritableVirtualDeviceContextRequestStatus
	isSet bool
}

func (v NullablePatchedWritableVirtualDeviceContextRequestStatus) Get() *PatchedWritableVirtualDeviceContextRequestStatus {
	return v.value
}

func (v *NullablePatchedWritableVirtualDeviceContextRequestStatus) Set(val *PatchedWritableVirtualDeviceContextRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableVirtualDeviceContextRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableVirtualDeviceContextRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableVirtualDeviceContextRequestStatus(val *PatchedWritableVirtualDeviceContextRequestStatus) *NullablePatchedWritableVirtualDeviceContextRequestStatus {
	return &NullablePatchedWritableVirtualDeviceContextRequestStatus{value: val, isSet: true}
}

func (v NullablePatchedWritableVirtualDeviceContextRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableVirtualDeviceContextRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


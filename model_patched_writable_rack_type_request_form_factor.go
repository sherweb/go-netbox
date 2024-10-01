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

// PatchedWritableRackTypeRequestFormFactor * `2-post-frame` - 2-post frame * `4-post-frame` - 4-post frame * `4-post-cabinet` - 4-post cabinet * `wall-frame` - Wall-mounted frame * `wall-frame-vertical` - Wall-mounted frame (vertical) * `wall-cabinet` - Wall-mounted cabinet * `wall-cabinet-vertical` - Wall-mounted cabinet (vertical)
type PatchedWritableRackTypeRequestFormFactor string

// List of PatchedWritableRackTypeRequest_form_factor
const (
	PATCHEDWRITABLERACKTYPEREQUESTFORMFACTOR__2_POST_FRAME PatchedWritableRackTypeRequestFormFactor = "2-post-frame"
	PATCHEDWRITABLERACKTYPEREQUESTFORMFACTOR__4_POST_FRAME PatchedWritableRackTypeRequestFormFactor = "4-post-frame"
	PATCHEDWRITABLERACKTYPEREQUESTFORMFACTOR__4_POST_CABINET PatchedWritableRackTypeRequestFormFactor = "4-post-cabinet"
	PATCHEDWRITABLERACKTYPEREQUESTFORMFACTOR_WALL_FRAME PatchedWritableRackTypeRequestFormFactor = "wall-frame"
	PATCHEDWRITABLERACKTYPEREQUESTFORMFACTOR_WALL_FRAME_VERTICAL PatchedWritableRackTypeRequestFormFactor = "wall-frame-vertical"
	PATCHEDWRITABLERACKTYPEREQUESTFORMFACTOR_WALL_CABINET PatchedWritableRackTypeRequestFormFactor = "wall-cabinet"
	PATCHEDWRITABLERACKTYPEREQUESTFORMFACTOR_WALL_CABINET_VERTICAL PatchedWritableRackTypeRequestFormFactor = "wall-cabinet-vertical"
)

// All allowed values of PatchedWritableRackTypeRequestFormFactor enum
var AllowedPatchedWritableRackTypeRequestFormFactorEnumValues = []PatchedWritableRackTypeRequestFormFactor{
	"2-post-frame",
	"4-post-frame",
	"4-post-cabinet",
	"wall-frame",
	"wall-frame-vertical",
	"wall-cabinet",
	"wall-cabinet-vertical",
}

func (v *PatchedWritableRackTypeRequestFormFactor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableRackTypeRequestFormFactor(value)
	for _, existing := range AllowedPatchedWritableRackTypeRequestFormFactorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableRackTypeRequestFormFactor", value)
}

// NewPatchedWritableRackTypeRequestFormFactorFromValue returns a pointer to a valid PatchedWritableRackTypeRequestFormFactor
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableRackTypeRequestFormFactorFromValue(v string) (*PatchedWritableRackTypeRequestFormFactor, error) {
	ev := PatchedWritableRackTypeRequestFormFactor(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableRackTypeRequestFormFactor: valid values are %v", v, AllowedPatchedWritableRackTypeRequestFormFactorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableRackTypeRequestFormFactor) IsValid() bool {
	for _, existing := range AllowedPatchedWritableRackTypeRequestFormFactorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableRackTypeRequest_form_factor value
func (v PatchedWritableRackTypeRequestFormFactor) Ptr() *PatchedWritableRackTypeRequestFormFactor {
	return &v
}

type NullablePatchedWritableRackTypeRequestFormFactor struct {
	value *PatchedWritableRackTypeRequestFormFactor
	isSet bool
}

func (v NullablePatchedWritableRackTypeRequestFormFactor) Get() *PatchedWritableRackTypeRequestFormFactor {
	return v.value
}

func (v *NullablePatchedWritableRackTypeRequestFormFactor) Set(val *PatchedWritableRackTypeRequestFormFactor) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableRackTypeRequestFormFactor) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableRackTypeRequestFormFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableRackTypeRequestFormFactor(val *PatchedWritableRackTypeRequestFormFactor) *NullablePatchedWritableRackTypeRequestFormFactor {
	return &NullablePatchedWritableRackTypeRequestFormFactor{value: val, isSet: true}
}

func (v NullablePatchedWritableRackTypeRequestFormFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableRackTypeRequestFormFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


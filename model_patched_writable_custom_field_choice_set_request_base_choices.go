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

// PatchedWritableCustomFieldChoiceSetRequestBaseChoices Base set of predefined choices (optional)  * `IATA` - IATA (Airport codes) * `ISO_3166` - ISO 3166 (Country codes) * `UN_LOCODE` - UN/LOCODE (Location codes)
type PatchedWritableCustomFieldChoiceSetRequestBaseChoices string

// List of PatchedWritableCustomFieldChoiceSetRequest_base_choices
const (
	PATCHEDWRITABLECUSTOMFIELDCHOICESETREQUESTBASECHOICES_IATA PatchedWritableCustomFieldChoiceSetRequestBaseChoices = "IATA"
	PATCHEDWRITABLECUSTOMFIELDCHOICESETREQUESTBASECHOICES_ISO_3166 PatchedWritableCustomFieldChoiceSetRequestBaseChoices = "ISO_3166"
	PATCHEDWRITABLECUSTOMFIELDCHOICESETREQUESTBASECHOICES_UN_LOCODE PatchedWritableCustomFieldChoiceSetRequestBaseChoices = "UN_LOCODE"
	PATCHEDWRITABLECUSTOMFIELDCHOICESETREQUESTBASECHOICES_EMPTY PatchedWritableCustomFieldChoiceSetRequestBaseChoices = ""
)

// All allowed values of PatchedWritableCustomFieldChoiceSetRequestBaseChoices enum
var AllowedPatchedWritableCustomFieldChoiceSetRequestBaseChoicesEnumValues = []PatchedWritableCustomFieldChoiceSetRequestBaseChoices{
	"IATA",
	"ISO_3166",
	"UN_LOCODE",
	"",
}

func (v *PatchedWritableCustomFieldChoiceSetRequestBaseChoices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableCustomFieldChoiceSetRequestBaseChoices(value)
	for _, existing := range AllowedPatchedWritableCustomFieldChoiceSetRequestBaseChoicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableCustomFieldChoiceSetRequestBaseChoices", value)
}

// NewPatchedWritableCustomFieldChoiceSetRequestBaseChoicesFromValue returns a pointer to a valid PatchedWritableCustomFieldChoiceSetRequestBaseChoices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableCustomFieldChoiceSetRequestBaseChoicesFromValue(v string) (*PatchedWritableCustomFieldChoiceSetRequestBaseChoices, error) {
	ev := PatchedWritableCustomFieldChoiceSetRequestBaseChoices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableCustomFieldChoiceSetRequestBaseChoices: valid values are %v", v, AllowedPatchedWritableCustomFieldChoiceSetRequestBaseChoicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableCustomFieldChoiceSetRequestBaseChoices) IsValid() bool {
	for _, existing := range AllowedPatchedWritableCustomFieldChoiceSetRequestBaseChoicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableCustomFieldChoiceSetRequest_base_choices value
func (v PatchedWritableCustomFieldChoiceSetRequestBaseChoices) Ptr() *PatchedWritableCustomFieldChoiceSetRequestBaseChoices {
	return &v
}

type NullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices struct {
	value *PatchedWritableCustomFieldChoiceSetRequestBaseChoices
	isSet bool
}

func (v NullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices) Get() *PatchedWritableCustomFieldChoiceSetRequestBaseChoices {
	return v.value
}

func (v *NullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices) Set(val *PatchedWritableCustomFieldChoiceSetRequestBaseChoices) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices(val *PatchedWritableCustomFieldChoiceSetRequestBaseChoices) *NullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices {
	return &NullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices{value: val, isSet: true}
}

func (v NullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


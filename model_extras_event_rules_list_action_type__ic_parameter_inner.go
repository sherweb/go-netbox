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

// ExtrasEventRulesListActionTypeIcParameterInner the model 'ExtrasEventRulesListActionTypeIcParameterInner'
type ExtrasEventRulesListActionTypeIcParameterInner string

// List of extras_event_rules_list_action_type__ic_parameter_inner
const (
	EXTRASEVENTRULESLISTACTIONTYPEICPARAMETERINNER_NOTIFICATION ExtrasEventRulesListActionTypeIcParameterInner = "notification"
	EXTRASEVENTRULESLISTACTIONTYPEICPARAMETERINNER_SCRIPT ExtrasEventRulesListActionTypeIcParameterInner = "script"
	EXTRASEVENTRULESLISTACTIONTYPEICPARAMETERINNER_WEBHOOK ExtrasEventRulesListActionTypeIcParameterInner = "webhook"
)

// All allowed values of ExtrasEventRulesListActionTypeIcParameterInner enum
var AllowedExtrasEventRulesListActionTypeIcParameterInnerEnumValues = []ExtrasEventRulesListActionTypeIcParameterInner{
	"notification",
	"script",
	"webhook",
}

func (v *ExtrasEventRulesListActionTypeIcParameterInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExtrasEventRulesListActionTypeIcParameterInner(value)
	for _, existing := range AllowedExtrasEventRulesListActionTypeIcParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExtrasEventRulesListActionTypeIcParameterInner", value)
}

// NewExtrasEventRulesListActionTypeIcParameterInnerFromValue returns a pointer to a valid ExtrasEventRulesListActionTypeIcParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExtrasEventRulesListActionTypeIcParameterInnerFromValue(v string) (*ExtrasEventRulesListActionTypeIcParameterInner, error) {
	ev := ExtrasEventRulesListActionTypeIcParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExtrasEventRulesListActionTypeIcParameterInner: valid values are %v", v, AllowedExtrasEventRulesListActionTypeIcParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExtrasEventRulesListActionTypeIcParameterInner) IsValid() bool {
	for _, existing := range AllowedExtrasEventRulesListActionTypeIcParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to extras_event_rules_list_action_type__ic_parameter_inner value
func (v ExtrasEventRulesListActionTypeIcParameterInner) Ptr() *ExtrasEventRulesListActionTypeIcParameterInner {
	return &v
}

type NullableExtrasEventRulesListActionTypeIcParameterInner struct {
	value *ExtrasEventRulesListActionTypeIcParameterInner
	isSet bool
}

func (v NullableExtrasEventRulesListActionTypeIcParameterInner) Get() *ExtrasEventRulesListActionTypeIcParameterInner {
	return v.value
}

func (v *NullableExtrasEventRulesListActionTypeIcParameterInner) Set(val *ExtrasEventRulesListActionTypeIcParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExtrasEventRulesListActionTypeIcParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExtrasEventRulesListActionTypeIcParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtrasEventRulesListActionTypeIcParameterInner(val *ExtrasEventRulesListActionTypeIcParameterInner) *NullableExtrasEventRulesListActionTypeIcParameterInner {
	return &NullableExtrasEventRulesListActionTypeIcParameterInner{value: val, isSet: true}
}

func (v NullableExtrasEventRulesListActionTypeIcParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtrasEventRulesListActionTypeIcParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


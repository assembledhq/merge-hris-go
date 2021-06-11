/*
 * Merge HRIS API
 *
 * The unified API for building rich integrations with multiple HR Information System platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_hris_client

import (
	"encoding/json"
	"fmt"
)

// UnitsEnum the model 'UnitsEnum'
type UnitsEnum string

// List of UnitsEnum
const (
	UNITSENUM_HOURS UnitsEnum = "HOURS"
	UNITSENUM_DAYS UnitsEnum = "DAYS"
)

var allowedUnitsEnumEnumValues = []UnitsEnum{
	"HOURS",
	"DAYS",
}

func (v *UnitsEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UnitsEnum(value)
	for _, existing := range allowedUnitsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UnitsEnum", value)
}

// NewUnitsEnumFromValue returns a pointer to a valid UnitsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUnitsEnumFromValue(v string) (*UnitsEnum, error) {
	ev := UnitsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UnitsEnum: valid values are %v", v, allowedUnitsEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UnitsEnum) IsValid() bool {
	for _, existing := range allowedUnitsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnitsEnum value
func (v UnitsEnum) Ptr() *UnitsEnum {
	return &v
}

type NullableUnitsEnum struct {
	value *UnitsEnum
	isSet bool
}

func (v NullableUnitsEnum) Get() *UnitsEnum {
	return v.value
}

func (v *NullableUnitsEnum) Set(val *UnitsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableUnitsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableUnitsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnitsEnum(val *UnitsEnum) *NullableUnitsEnum {
	return &NullableUnitsEnum{value: val, isSet: true}
}

func (v NullableUnitsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnitsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


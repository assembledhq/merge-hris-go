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

// LocationTypeEnum the model 'LocationTypeEnum'
type LocationTypeEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of LocationTypeEnum
const (
    LOCATIONTYPEENUM_MERGE_NONSTANDARD_VALUE LocationTypeEnum = "MERGE_NONSTANDARD_VALUE"
    
	LOCATIONTYPEENUM_HOME LocationTypeEnum = "HOME"
	LOCATIONTYPEENUM_WORK LocationTypeEnum = "WORK"
)

var allowedLocationTypeEnumEnumValues = []LocationTypeEnum{
	"HOME",
	"WORK",
}

func (v *LocationTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocationTypeEnum(value)
	for _, existing := range allowedLocationTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = LOCATIONTYPEENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewLocationTypeEnumFromValue returns a pointer to a valid LocationTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocationTypeEnumFromValue(v string) (*LocationTypeEnum, error) {
	ev := LocationTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := LOCATIONTYPEENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocationTypeEnum) IsValid() bool {
	for _, existing := range allowedLocationTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocationTypeEnum value
func (v LocationTypeEnum) Ptr() *LocationTypeEnum {
	return &v
}

type NullableLocationTypeEnum struct {
	value *LocationTypeEnum
	isSet bool
}

func (v NullableLocationTypeEnum) Get() *LocationTypeEnum {
	return v.value
}

func (v *NullableLocationTypeEnum) Set(val *LocationTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationTypeEnum(val *LocationTypeEnum) *NullableLocationTypeEnum {
	return &NullableLocationTypeEnum{value: val, isSet: true}
}

func (v NullableLocationTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


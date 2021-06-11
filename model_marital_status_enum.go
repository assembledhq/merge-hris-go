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

// MaritalStatusEnum the model 'MaritalStatusEnum'
type MaritalStatusEnum string

// List of MaritalStatusEnum
const (
	MARITALSTATUSENUM_SINGLE MaritalStatusEnum = "SINGLE"
	MARITALSTATUSENUM_MARRIED_FILING_JOINTLY MaritalStatusEnum = "MARRIED_FILING_JOINTLY"
	MARITALSTATUSENUM_MARRIED_FILING_SEPARATELY MaritalStatusEnum = "MARRIED_FILING_SEPARATELY"
	MARITALSTATUSENUM_HEAD_OF_HOUSEHOLD MaritalStatusEnum = "HEAD_OF_HOUSEHOLD"
	MARITALSTATUSENUM_QUALIFYING_WIDOW_OR_WIDOWER_WITH_DEPENDENT_CHILD MaritalStatusEnum = "QUALIFYING_WIDOW_OR_WIDOWER_WITH_DEPENDENT_CHILD"
)

var allowedMaritalStatusEnumEnumValues = []MaritalStatusEnum{
	"SINGLE",
	"MARRIED_FILING_JOINTLY",
	"MARRIED_FILING_SEPARATELY",
	"HEAD_OF_HOUSEHOLD",
	"QUALIFYING_WIDOW_OR_WIDOWER_WITH_DEPENDENT_CHILD",
}

func (v *MaritalStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MaritalStatusEnum(value)
	for _, existing := range allowedMaritalStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MaritalStatusEnum", value)
}

// NewMaritalStatusEnumFromValue returns a pointer to a valid MaritalStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMaritalStatusEnumFromValue(v string) (*MaritalStatusEnum, error) {
	ev := MaritalStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MaritalStatusEnum: valid values are %v", v, allowedMaritalStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MaritalStatusEnum) IsValid() bool {
	for _, existing := range allowedMaritalStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MaritalStatusEnum value
func (v MaritalStatusEnum) Ptr() *MaritalStatusEnum {
	return &v
}

type NullableMaritalStatusEnum struct {
	value *MaritalStatusEnum
	isSet bool
}

func (v NullableMaritalStatusEnum) Get() *MaritalStatusEnum {
	return v.value
}

func (v *NullableMaritalStatusEnum) Set(val *MaritalStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMaritalStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMaritalStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaritalStatusEnum(val *MaritalStatusEnum) *NullableMaritalStatusEnum {
	return &NullableMaritalStatusEnum{value: val, isSet: true}
}

func (v NullableMaritalStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaritalStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


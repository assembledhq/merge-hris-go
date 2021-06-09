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

// EmploymentStatusEnum the model 'EmploymentStatusEnum'
type EmploymentStatusEnum string

// List of EmploymentStatusEnum
const (
	EMPLOYMENTSTATUSENUM_ACTIVE EmploymentStatusEnum = "ACTIVE"
	EMPLOYMENTSTATUSENUM_PENDING EmploymentStatusEnum = "PENDING"
	EMPLOYMENTSTATUSENUM_INACTIVE EmploymentStatusEnum = "INACTIVE"
)

var allowedEmploymentStatusEnumEnumValues = []EmploymentStatusEnum{
	"ACTIVE",
	"PENDING",
	"INACTIVE",
}

func (v *EmploymentStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmploymentStatusEnum(value)
	for _, existing := range allowedEmploymentStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmploymentStatusEnum", value)
}

// NewEmploymentStatusEnumFromValue returns a pointer to a valid EmploymentStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmploymentStatusEnumFromValue(v string) (*EmploymentStatusEnum, error) {
	ev := EmploymentStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmploymentStatusEnum: valid values are %v", v, allowedEmploymentStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmploymentStatusEnum) IsValid() bool {
	for _, existing := range allowedEmploymentStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmploymentStatusEnum value
func (v EmploymentStatusEnum) Ptr() *EmploymentStatusEnum {
	return &v
}

type NullableEmploymentStatusEnum struct {
	value *EmploymentStatusEnum
	isSet bool
}

func (v NullableEmploymentStatusEnum) Get() *EmploymentStatusEnum {
	return v.value
}

func (v *NullableEmploymentStatusEnum) Set(val *EmploymentStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableEmploymentStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableEmploymentStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmploymentStatusEnum(val *EmploymentStatusEnum) *NullableEmploymentStatusEnum {
	return &NullableEmploymentStatusEnum{value: val, isSet: true}
}

func (v NullableEmploymentStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmploymentStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


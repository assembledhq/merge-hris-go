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

// FlsaStatusEnum the model 'FlsaStatusEnum'
type FlsaStatusEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of FlsaStatusEnum
const (
    FLSASTATUSENUM_MERGE_NONSTANDARD_VALUE FlsaStatusEnum = "MERGE_NONSTANDARD_VALUE"
    
	FLSASTATUSENUM_EXEMPT FlsaStatusEnum = "EXEMPT"
	FLSASTATUSENUM_SALARIED_NONEXEMPT FlsaStatusEnum = "SALARIED_NONEXEMPT"
	FLSASTATUSENUM_NONEXEMPT FlsaStatusEnum = "NONEXEMPT"
	FLSASTATUSENUM_OWNER FlsaStatusEnum = "OWNER"
)

var allowedFlsaStatusEnumEnumValues = []FlsaStatusEnum{
	"EXEMPT",
	"SALARIED_NONEXEMPT",
	"NONEXEMPT",
	"OWNER",
}

func (v *FlsaStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FlsaStatusEnum(value)
	for _, existing := range allowedFlsaStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = FLSASTATUSENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewFlsaStatusEnumFromValue returns a pointer to a valid FlsaStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlsaStatusEnumFromValue(v string) (*FlsaStatusEnum, error) {
	ev := FlsaStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := FLSASTATUSENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlsaStatusEnum) IsValid() bool {
	for _, existing := range allowedFlsaStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlsaStatusEnum value
func (v FlsaStatusEnum) Ptr() *FlsaStatusEnum {
	return &v
}

type NullableFlsaStatusEnum struct {
	value *FlsaStatusEnum
	isSet bool
}

func (v NullableFlsaStatusEnum) Get() *FlsaStatusEnum {
	return v.value
}

func (v *NullableFlsaStatusEnum) Set(val *FlsaStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableFlsaStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableFlsaStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlsaStatusEnum(val *FlsaStatusEnum) *NullableFlsaStatusEnum {
	return &NullableFlsaStatusEnum{value: val, isSet: true}
}

func (v NullableFlsaStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlsaStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


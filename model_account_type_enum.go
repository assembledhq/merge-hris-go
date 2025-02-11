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

// AccountTypeEnum the model 'AccountTypeEnum'
type AccountTypeEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of AccountTypeEnum
const (
    ACCOUNTTYPEENUM_MERGE_NONSTANDARD_VALUE AccountTypeEnum = "MERGE_NONSTANDARD_VALUE"
    
	ACCOUNTTYPEENUM_SAVINGS AccountTypeEnum = "SAVINGS"
	ACCOUNTTYPEENUM_CHECKING AccountTypeEnum = "CHECKING"
)

var allowedAccountTypeEnumEnumValues = []AccountTypeEnum{
	"SAVINGS",
	"CHECKING",
}

func (v *AccountTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountTypeEnum(value)
	for _, existing := range allowedAccountTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = ACCOUNTTYPEENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewAccountTypeEnumFromValue returns a pointer to a valid AccountTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountTypeEnumFromValue(v string) (*AccountTypeEnum, error) {
	ev := AccountTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := ACCOUNTTYPEENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountTypeEnum) IsValid() bool {
	for _, existing := range allowedAccountTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountTypeEnum value
func (v AccountTypeEnum) Ptr() *AccountTypeEnum {
	return &v
}

type NullableAccountTypeEnum struct {
	value *AccountTypeEnum
	isSet bool
}

func (v NullableAccountTypeEnum) Get() *AccountTypeEnum {
	return v.value
}

func (v *NullableAccountTypeEnum) Set(val *AccountTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountTypeEnum(val *AccountTypeEnum) *NullableAccountTypeEnum {
	return &NullableAccountTypeEnum{value: val, isSet: true}
}

func (v NullableAccountTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


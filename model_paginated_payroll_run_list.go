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
)

// PaginatedPayrollRunList struct for PaginatedPayrollRunList
type PaginatedPayrollRunList struct {
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results *[]PayrollRun `json:"results,omitempty"`
    // raw json response by property name
    responseRaw map[string]json.RawMessage `json:"-"`
}

// NewPaginatedPayrollRunList instantiates a new PaginatedPayrollRunList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPayrollRunList() *PaginatedPayrollRunList {
	this := PaginatedPayrollRunList{}
	return &this
}

// NewPaginatedPayrollRunListWithDefaults instantiates a new PaginatedPayrollRunList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPayrollRunListWithDefaults() *PaginatedPayrollRunList {
	this := PaginatedPayrollRunList{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedPayrollRunList) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedPayrollRunList) GetNextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedPayrollRunList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedPayrollRunList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedPayrollRunList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedPayrollRunList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedPayrollRunList) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedPayrollRunList) GetPreviousOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedPayrollRunList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedPayrollRunList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedPayrollRunList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedPayrollRunList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedPayrollRunList) GetResults() []PayrollRun {
	if o == nil || o.Results == nil {
		var ret []PayrollRun
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedPayrollRunList) GetResultsOk() (*[]PayrollRun, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedPayrollRunList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []PayrollRun and assigns it to the Results field.
func (o *PaginatedPayrollRunList) SetResults(v []PayrollRun) {
	o.Results = &v
}

func (o PaginatedPayrollRunList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

func (v *PaginatedPayrollRunList) UnmarshalJSON(src []byte) error {
    type PaginatedPayrollRunListUnmarshalTarget PaginatedPayrollRunList

	var intermediateResult PaginatedPayrollRunListUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.responseRaw)
	if err2 != nil {
		return err2
	}

	*v = PaginatedPayrollRunList(intermediateResult)
	return nil
}
type NullablePaginatedPayrollRunList struct {
	value *PaginatedPayrollRunList
	isSet bool
}

func (v NullablePaginatedPayrollRunList) Get() *PaginatedPayrollRunList {
	return v.value
}

func (v *NullablePaginatedPayrollRunList) Set(val *PaginatedPayrollRunList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPayrollRunList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPayrollRunList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPayrollRunList(val *PaginatedPayrollRunList) *NullablePaginatedPayrollRunList {
	return &NullablePaginatedPayrollRunList{value: val, isSet: true}
}

func (v NullablePaginatedPayrollRunList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPayrollRunList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.responseRaw)
}



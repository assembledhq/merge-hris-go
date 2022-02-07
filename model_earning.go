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

// Earning # The Earning Object ### Description The `Earning` object is used to represent an earning for a given employee's payroll run. One run could include several earnings.  ### Usage Example Fetch from the `LIST Earnings` endpoint and filter by `ID` to show all earnings.
type Earning struct {
	Id *string `json:"id,omitempty"`
	EmployeePayrollRun NullableString `json:"employee_payroll_run,omitempty"`
	// The amount earned.
	Amount NullableFloat32 `json:"amount,omitempty"`
	// The type of earning.
	Type NullableTypeEnum `json:"type,omitempty"`
	RemoteData []map[string]interface{} `json:"remote_data,omitempty"`
}

// NewEarning instantiates a new Earning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEarning() *Earning {
	this := Earning{}
	return &this
}

// NewEarningWithDefaults instantiates a new Earning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEarningWithDefaults() *Earning {
	this := Earning{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Earning) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Earning) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Earning) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Earning) SetId(v string) {
	o.Id = &v
}

// GetEmployeePayrollRun returns the EmployeePayrollRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Earning) GetEmployeePayrollRun() string {
	if o == nil || o.EmployeePayrollRun.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmployeePayrollRun.Get()
}

// GetEmployeePayrollRunOk returns a tuple with the EmployeePayrollRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Earning) GetEmployeePayrollRunOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployeePayrollRun.Get(), o.EmployeePayrollRun.IsSet()
}

// HasEmployeePayrollRun returns a boolean if a field has been set.
func (o *Earning) HasEmployeePayrollRun() bool {
	if o != nil && o.EmployeePayrollRun.IsSet() {
		return true
	}

	return false
}

// SetEmployeePayrollRun gets a reference to the given NullableString and assigns it to the EmployeePayrollRun field.
func (o *Earning) SetEmployeePayrollRun(v string) {
	o.EmployeePayrollRun.Set(&v)
}
// SetEmployeePayrollRunNil sets the value for EmployeePayrollRun to be an explicit nil
func (o *Earning) SetEmployeePayrollRunNil() {
	o.EmployeePayrollRun.Set(nil)
}

// UnsetEmployeePayrollRun ensures that no value is present for EmployeePayrollRun, not even an explicit nil
func (o *Earning) UnsetEmployeePayrollRun() {
	o.EmployeePayrollRun.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Earning) GetAmount() float32 {
	if o == nil || o.Amount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Earning) GetAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *Earning) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat32 and assigns it to the Amount field.
func (o *Earning) SetAmount(v float32) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *Earning) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *Earning) UnsetAmount() {
	o.Amount.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Earning) GetType() TypeEnum {
	if o == nil || o.Type.Get() == nil {
		var ret TypeEnum
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Earning) GetTypeOk() (*TypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Earning) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableTypeEnum and assigns it to the Type field.
func (o *Earning) SetType(v TypeEnum) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Earning) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Earning) UnsetType() {
	o.Type.Unset()
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Earning) GetRemoteData() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Earning) GetRemoteDataOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *Earning) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []map[string]interface{} and assigns it to the RemoteData field.
func (o *Earning) SetRemoteData(v []map[string]interface{}) {
	o.RemoteData = v
}

func (o Earning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.EmployeePayrollRun.IsSet() {
		toSerialize["employee_payroll_run"] = o.EmployeePayrollRun.Get()
	}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	return json.Marshal(toSerialize)
}

type NullableEarning struct {
	value *Earning
	isSet bool
}

func (v NullableEarning) Get() *Earning {
	return v.value
}

func (v *NullableEarning) Set(val *Earning) {
	v.value = val
	v.isSet = true
}

func (v NullableEarning) IsSet() bool {
	return v.isSet
}

func (v *NullableEarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarning(val *Earning) *NullableEarning {
	return &NullableEarning{value: val, isSet: true}
}

func (v NullableEarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// TimeOffBalance # The TimeOffBalance Object ### Description The `TimeOffBalance` object is used to represent a Time Off Balance for an employee.  ### Usage Example Fetch from the `LIST TimeOffBalances` endpoint and filter by `ID` to show all time off balances.
type TimeOffBalance struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	Employee NullableString `json:"employee,omitempty"`
	// The current PTO balance in terms of hours.
	Balance NullableFloat32 `json:"balance,omitempty"`
	// The amount of PTO used in terms of hours.
	Used NullableFloat32 `json:"used,omitempty"`
	// The policy type of this time off balance.
	PolicyType NullablePolicyTypeEnum `json:"policy_type,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewTimeOffBalance instantiates a new TimeOffBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeOffBalance() *TimeOffBalance {
	this := TimeOffBalance{}
	return &this
}

// NewTimeOffBalanceWithDefaults instantiates a new TimeOffBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeOffBalanceWithDefaults() *TimeOffBalance {
	this := TimeOffBalance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TimeOffBalance) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeOffBalance) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TimeOffBalance) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TimeOffBalance) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffBalance) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffBalance) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *TimeOffBalance) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *TimeOffBalance) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *TimeOffBalance) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *TimeOffBalance) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetEmployee returns the Employee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffBalance) GetEmployee() string {
	if o == nil || o.Employee.Get() == nil {
		var ret string
		return ret
	}
	return *o.Employee.Get()
}

// GetEmployeeOk returns a tuple with the Employee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffBalance) GetEmployeeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Employee.Get(), o.Employee.IsSet()
}

// HasEmployee returns a boolean if a field has been set.
func (o *TimeOffBalance) HasEmployee() bool {
	if o != nil && o.Employee.IsSet() {
		return true
	}

	return false
}

// SetEmployee gets a reference to the given NullableString and assigns it to the Employee field.
func (o *TimeOffBalance) SetEmployee(v string) {
	o.Employee.Set(&v)
}
// SetEmployeeNil sets the value for Employee to be an explicit nil
func (o *TimeOffBalance) SetEmployeeNil() {
	o.Employee.Set(nil)
}

// UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
func (o *TimeOffBalance) UnsetEmployee() {
	o.Employee.Unset()
}

// GetBalance returns the Balance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffBalance) GetBalance() float32 {
	if o == nil || o.Balance.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffBalance) GetBalanceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// HasBalance returns a boolean if a field has been set.
func (o *TimeOffBalance) HasBalance() bool {
	if o != nil && o.Balance.IsSet() {
		return true
	}

	return false
}

// SetBalance gets a reference to the given NullableFloat32 and assigns it to the Balance field.
func (o *TimeOffBalance) SetBalance(v float32) {
	o.Balance.Set(&v)
}
// SetBalanceNil sets the value for Balance to be an explicit nil
func (o *TimeOffBalance) SetBalanceNil() {
	o.Balance.Set(nil)
}

// UnsetBalance ensures that no value is present for Balance, not even an explicit nil
func (o *TimeOffBalance) UnsetBalance() {
	o.Balance.Unset()
}

// GetUsed returns the Used field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffBalance) GetUsed() float32 {
	if o == nil || o.Used.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Used.Get()
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffBalance) GetUsedOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Used.Get(), o.Used.IsSet()
}

// HasUsed returns a boolean if a field has been set.
func (o *TimeOffBalance) HasUsed() bool {
	if o != nil && o.Used.IsSet() {
		return true
	}

	return false
}

// SetUsed gets a reference to the given NullableFloat32 and assigns it to the Used field.
func (o *TimeOffBalance) SetUsed(v float32) {
	o.Used.Set(&v)
}
// SetUsedNil sets the value for Used to be an explicit nil
func (o *TimeOffBalance) SetUsedNil() {
	o.Used.Set(nil)
}

// UnsetUsed ensures that no value is present for Used, not even an explicit nil
func (o *TimeOffBalance) UnsetUsed() {
	o.Used.Unset()
}

// GetPolicyType returns the PolicyType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffBalance) GetPolicyType() PolicyTypeEnum {
	if o == nil || o.PolicyType.Get() == nil {
		var ret PolicyTypeEnum
		return ret
	}
	return *o.PolicyType.Get()
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffBalance) GetPolicyTypeOk() (*PolicyTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyType.Get(), o.PolicyType.IsSet()
}

// HasPolicyType returns a boolean if a field has been set.
func (o *TimeOffBalance) HasPolicyType() bool {
	if o != nil && o.PolicyType.IsSet() {
		return true
	}

	return false
}

// SetPolicyType gets a reference to the given NullablePolicyTypeEnum and assigns it to the PolicyType field.
func (o *TimeOffBalance) SetPolicyType(v PolicyTypeEnum) {
	o.PolicyType.Set(&v)
}
// SetPolicyTypeNil sets the value for PolicyType to be an explicit nil
func (o *TimeOffBalance) SetPolicyTypeNil() {
	o.PolicyType.Set(nil)
}

// UnsetPolicyType ensures that no value is present for PolicyType, not even an explicit nil
func (o *TimeOffBalance) UnsetPolicyType() {
	o.PolicyType.Unset()
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeOffBalance) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffBalance) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *TimeOffBalance) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *TimeOffBalance) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

func (o TimeOffBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.Employee.IsSet() {
		toSerialize["employee"] = o.Employee.Get()
	}
	if o.Balance.IsSet() {
		toSerialize["balance"] = o.Balance.Get()
	}
	if o.Used.IsSet() {
		toSerialize["used"] = o.Used.Get()
	}
	if o.PolicyType.IsSet() {
		toSerialize["policy_type"] = o.PolicyType.Get()
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	return json.Marshal(toSerialize)
}

func (v *TimeOffBalance) UnmarshalJSON(src []byte) error {
    type TimeOffBalanceUnmarshalTarget TimeOffBalance

	var intermediateResult TimeOffBalanceUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = TimeOffBalance(intermediateResult)
	return nil
}
type NullableTimeOffBalance struct {
	value *TimeOffBalance
	isSet bool
}

func (v NullableTimeOffBalance) Get() *TimeOffBalance {
	return v.value
}

func (v *NullableTimeOffBalance) Set(val *TimeOffBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeOffBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeOffBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeOffBalance(val *TimeOffBalance) *NullableTimeOffBalance {
	return &NullableTimeOffBalance{value: val, isSet: true}
}

func (v NullableTimeOffBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeOffBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



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
	"time"
)

// BankInfo # The BankInfo Object ### Description The `BankInfo` object is used to represent the Bank Account information for an Employee. This is often referenced with an Employee object.  ### Usage Example Fetch from the `LIST BankInfo` endpoint and filter by `ID` to show all bank information.
type BankInfo struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	Employee NullableString `json:"employee,omitempty"`
	// The account number.
	AccountNumber NullableString `json:"account_number,omitempty"`
	// The routing number.
	RoutingNumber NullableString `json:"routing_number,omitempty"`
	// The bank name.
	BankName NullableString `json:"bank_name,omitempty"`
	// The bank account type
	AccountType NullableAccountTypeEnum `json:"account_type,omitempty"`
	// When the matching bank object was created in the third party system.
	RemoteCreatedAt NullableTime `json:"remote_created_at,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewBankInfo instantiates a new BankInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankInfo() *BankInfo {
	this := BankInfo{}
	return &this
}

// NewBankInfoWithDefaults instantiates a new BankInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankInfoWithDefaults() *BankInfo {
	this := BankInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BankInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BankInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BankInfo) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankInfo) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankInfo) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *BankInfo) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *BankInfo) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *BankInfo) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *BankInfo) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetEmployee returns the Employee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankInfo) GetEmployee() string {
	if o == nil || o.Employee.Get() == nil {
		var ret string
		return ret
	}
	return *o.Employee.Get()
}

// GetEmployeeOk returns a tuple with the Employee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankInfo) GetEmployeeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Employee.Get(), o.Employee.IsSet()
}

// HasEmployee returns a boolean if a field has been set.
func (o *BankInfo) HasEmployee() bool {
	if o != nil && o.Employee.IsSet() {
		return true
	}

	return false
}

// SetEmployee gets a reference to the given NullableString and assigns it to the Employee field.
func (o *BankInfo) SetEmployee(v string) {
	o.Employee.Set(&v)
}
// SetEmployeeNil sets the value for Employee to be an explicit nil
func (o *BankInfo) SetEmployeeNil() {
	o.Employee.Set(nil)
}

// UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
func (o *BankInfo) UnsetEmployee() {
	o.Employee.Unset()
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankInfo) GetAccountNumber() string {
	if o == nil || o.AccountNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber.Get()
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankInfo) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountNumber.Get(), o.AccountNumber.IsSet()
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *BankInfo) HasAccountNumber() bool {
	if o != nil && o.AccountNumber.IsSet() {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given NullableString and assigns it to the AccountNumber field.
func (o *BankInfo) SetAccountNumber(v string) {
	o.AccountNumber.Set(&v)
}
// SetAccountNumberNil sets the value for AccountNumber to be an explicit nil
func (o *BankInfo) SetAccountNumberNil() {
	o.AccountNumber.Set(nil)
}

// UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
func (o *BankInfo) UnsetAccountNumber() {
	o.AccountNumber.Unset()
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankInfo) GetRoutingNumber() string {
	if o == nil || o.RoutingNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.RoutingNumber.Get()
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankInfo) GetRoutingNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RoutingNumber.Get(), o.RoutingNumber.IsSet()
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *BankInfo) HasRoutingNumber() bool {
	if o != nil && o.RoutingNumber.IsSet() {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given NullableString and assigns it to the RoutingNumber field.
func (o *BankInfo) SetRoutingNumber(v string) {
	o.RoutingNumber.Set(&v)
}
// SetRoutingNumberNil sets the value for RoutingNumber to be an explicit nil
func (o *BankInfo) SetRoutingNumberNil() {
	o.RoutingNumber.Set(nil)
}

// UnsetRoutingNumber ensures that no value is present for RoutingNumber, not even an explicit nil
func (o *BankInfo) UnsetRoutingNumber() {
	o.RoutingNumber.Unset()
}

// GetBankName returns the BankName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankInfo) GetBankName() string {
	if o == nil || o.BankName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BankName.Get()
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankInfo) GetBankNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankName.Get(), o.BankName.IsSet()
}

// HasBankName returns a boolean if a field has been set.
func (o *BankInfo) HasBankName() bool {
	if o != nil && o.BankName.IsSet() {
		return true
	}

	return false
}

// SetBankName gets a reference to the given NullableString and assigns it to the BankName field.
func (o *BankInfo) SetBankName(v string) {
	o.BankName.Set(&v)
}
// SetBankNameNil sets the value for BankName to be an explicit nil
func (o *BankInfo) SetBankNameNil() {
	o.BankName.Set(nil)
}

// UnsetBankName ensures that no value is present for BankName, not even an explicit nil
func (o *BankInfo) UnsetBankName() {
	o.BankName.Unset()
}

// GetAccountType returns the AccountType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankInfo) GetAccountType() AccountTypeEnum {
	if o == nil || o.AccountType.Get() == nil {
		var ret AccountTypeEnum
		return ret
	}
	return *o.AccountType.Get()
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankInfo) GetAccountTypeOk() (*AccountTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountType.Get(), o.AccountType.IsSet()
}

// HasAccountType returns a boolean if a field has been set.
func (o *BankInfo) HasAccountType() bool {
	if o != nil && o.AccountType.IsSet() {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given NullableAccountTypeEnum and assigns it to the AccountType field.
func (o *BankInfo) SetAccountType(v AccountTypeEnum) {
	o.AccountType.Set(&v)
}
// SetAccountTypeNil sets the value for AccountType to be an explicit nil
func (o *BankInfo) SetAccountTypeNil() {
	o.AccountType.Set(nil)
}

// UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
func (o *BankInfo) UnsetAccountType() {
	o.AccountType.Unset()
}

// GetRemoteCreatedAt returns the RemoteCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankInfo) GetRemoteCreatedAt() time.Time {
	if o == nil || o.RemoteCreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteCreatedAt.Get()
}

// GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankInfo) GetRemoteCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteCreatedAt.Get(), o.RemoteCreatedAt.IsSet()
}

// HasRemoteCreatedAt returns a boolean if a field has been set.
func (o *BankInfo) HasRemoteCreatedAt() bool {
	if o != nil && o.RemoteCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteCreatedAt gets a reference to the given NullableTime and assigns it to the RemoteCreatedAt field.
func (o *BankInfo) SetRemoteCreatedAt(v time.Time) {
	o.RemoteCreatedAt.Set(&v)
}
// SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil
func (o *BankInfo) SetRemoteCreatedAtNil() {
	o.RemoteCreatedAt.Set(nil)
}

// UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
func (o *BankInfo) UnsetRemoteCreatedAt() {
	o.RemoteCreatedAt.Unset()
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankInfo) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankInfo) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *BankInfo) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *BankInfo) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

func (o BankInfo) MarshalJSON() ([]byte, error) {
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
	if o.AccountNumber.IsSet() {
		toSerialize["account_number"] = o.AccountNumber.Get()
	}
	if o.RoutingNumber.IsSet() {
		toSerialize["routing_number"] = o.RoutingNumber.Get()
	}
	if o.BankName.IsSet() {
		toSerialize["bank_name"] = o.BankName.Get()
	}
	if o.AccountType.IsSet() {
		toSerialize["account_type"] = o.AccountType.Get()
	}
	if o.RemoteCreatedAt.IsSet() {
		toSerialize["remote_created_at"] = o.RemoteCreatedAt.Get()
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	return json.Marshal(toSerialize)
}

func (v *BankInfo) UnmarshalJSON(src []byte) error {
    type BankInfoUnmarshalTarget BankInfo

	var intermediateResult BankInfoUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = BankInfo(intermediateResult)
	return nil
}
type NullableBankInfo struct {
	value *BankInfo
	isSet bool
}

func (v NullableBankInfo) Get() *BankInfo {
	return v.value
}

func (v *NullableBankInfo) Set(val *BankInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBankInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBankInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankInfo(val *BankInfo) *NullableBankInfo {
	return &NullableBankInfo{value: val, isSet: true}
}

func (v NullableBankInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


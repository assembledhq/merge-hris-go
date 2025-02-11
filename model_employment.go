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

// Employment # The Employment Object ### Description The `Employment` object is used to represent an employment position at a company. These are associated with the employee filling the role.  ### Usage Example Fetch from the `LIST Employments` endpoint and filter by `ID` to show all employees.
type Employment struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	Employee NullableString `json:"employee,omitempty"`
	// The position's title.
	JobTitle NullableString `json:"job_title,omitempty"`
	// The position's pay rate in dollars.
	PayRate NullableFloat32 `json:"pay_rate,omitempty"`
	// The time period this pay rate encompasses.
	PayPeriod NullablePayPeriodEnum `json:"pay_period,omitempty"`
	// The position's pay frequency.
	PayFrequency NullablePayFrequencyEnum `json:"pay_frequency,omitempty"`
	// The position's currency code.
	PayCurrency NullablePayCurrencyEnum `json:"pay_currency,omitempty"`
	// The position's FLSA status.
	FlsaStatus NullableFlsaStatusEnum `json:"flsa_status,omitempty"`
	// The position's effective date.
	EffectiveDate NullableTime `json:"effective_date,omitempty"`
	// The position's type of employment.
	EmploymentType NullableEmploymentTypeEnum `json:"employment_type,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewEmployment instantiates a new Employment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployment() *Employment {
	this := Employment{}
	return &this
}

// NewEmploymentWithDefaults instantiates a new Employment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmploymentWithDefaults() *Employment {
	this := Employment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Employment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Employment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Employment) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employment) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employment) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *Employment) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *Employment) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *Employment) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *Employment) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetEmployee returns the Employee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employment) GetEmployee() string {
	if o == nil || o.Employee.Get() == nil {
		var ret string
		return ret
	}
	return *o.Employee.Get()
}

// GetEmployeeOk returns a tuple with the Employee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employment) GetEmployeeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Employee.Get(), o.Employee.IsSet()
}

// HasEmployee returns a boolean if a field has been set.
func (o *Employment) HasEmployee() bool {
	if o != nil && o.Employee.IsSet() {
		return true
	}

	return false
}

// SetEmployee gets a reference to the given NullableString and assigns it to the Employee field.
func (o *Employment) SetEmployee(v string) {
	o.Employee.Set(&v)
}
// SetEmployeeNil sets the value for Employee to be an explicit nil
func (o *Employment) SetEmployeeNil() {
	o.Employee.Set(nil)
}

// UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
func (o *Employment) UnsetEmployee() {
	o.Employee.Unset()
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employment) GetJobTitle() string {
	if o == nil || o.JobTitle.Get() == nil {
		var ret string
		return ret
	}
	return *o.JobTitle.Get()
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employment) GetJobTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JobTitle.Get(), o.JobTitle.IsSet()
}

// HasJobTitle returns a boolean if a field has been set.
func (o *Employment) HasJobTitle() bool {
	if o != nil && o.JobTitle.IsSet() {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given NullableString and assigns it to the JobTitle field.
func (o *Employment) SetJobTitle(v string) {
	o.JobTitle.Set(&v)
}
// SetJobTitleNil sets the value for JobTitle to be an explicit nil
func (o *Employment) SetJobTitleNil() {
	o.JobTitle.Set(nil)
}

// UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
func (o *Employment) UnsetJobTitle() {
	o.JobTitle.Unset()
}

// GetPayRate returns the PayRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employment) GetPayRate() float32 {
	if o == nil || o.PayRate.Get() == nil {
		var ret float32
		return ret
	}
	return *o.PayRate.Get()
}

// GetPayRateOk returns a tuple with the PayRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employment) GetPayRateOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayRate.Get(), o.PayRate.IsSet()
}

// HasPayRate returns a boolean if a field has been set.
func (o *Employment) HasPayRate() bool {
	if o != nil && o.PayRate.IsSet() {
		return true
	}

	return false
}

// SetPayRate gets a reference to the given NullableFloat32 and assigns it to the PayRate field.
func (o *Employment) SetPayRate(v float32) {
	o.PayRate.Set(&v)
}
// SetPayRateNil sets the value for PayRate to be an explicit nil
func (o *Employment) SetPayRateNil() {
	o.PayRate.Set(nil)
}

// UnsetPayRate ensures that no value is present for PayRate, not even an explicit nil
func (o *Employment) UnsetPayRate() {
	o.PayRate.Unset()
}

// GetPayPeriod returns the PayPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employment) GetPayPeriod() PayPeriodEnum {
	if o == nil || o.PayPeriod.Get() == nil {
		var ret PayPeriodEnum
		return ret
	}
	return *o.PayPeriod.Get()
}

// GetPayPeriodOk returns a tuple with the PayPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employment) GetPayPeriodOk() (*PayPeriodEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayPeriod.Get(), o.PayPeriod.IsSet()
}

// HasPayPeriod returns a boolean if a field has been set.
func (o *Employment) HasPayPeriod() bool {
	if o != nil && o.PayPeriod.IsSet() {
		return true
	}

	return false
}

// SetPayPeriod gets a reference to the given NullablePayPeriodEnum and assigns it to the PayPeriod field.
func (o *Employment) SetPayPeriod(v PayPeriodEnum) {
	o.PayPeriod.Set(&v)
}
// SetPayPeriodNil sets the value for PayPeriod to be an explicit nil
func (o *Employment) SetPayPeriodNil() {
	o.PayPeriod.Set(nil)
}

// UnsetPayPeriod ensures that no value is present for PayPeriod, not even an explicit nil
func (o *Employment) UnsetPayPeriod() {
	o.PayPeriod.Unset()
}

// GetPayFrequency returns the PayFrequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employment) GetPayFrequency() PayFrequencyEnum {
	if o == nil || o.PayFrequency.Get() == nil {
		var ret PayFrequencyEnum
		return ret
	}
	return *o.PayFrequency.Get()
}

// GetPayFrequencyOk returns a tuple with the PayFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employment) GetPayFrequencyOk() (*PayFrequencyEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayFrequency.Get(), o.PayFrequency.IsSet()
}

// HasPayFrequency returns a boolean if a field has been set.
func (o *Employment) HasPayFrequency() bool {
	if o != nil && o.PayFrequency.IsSet() {
		return true
	}

	return false
}

// SetPayFrequency gets a reference to the given NullablePayFrequencyEnum and assigns it to the PayFrequency field.
func (o *Employment) SetPayFrequency(v PayFrequencyEnum) {
	o.PayFrequency.Set(&v)
}
// SetPayFrequencyNil sets the value for PayFrequency to be an explicit nil
func (o *Employment) SetPayFrequencyNil() {
	o.PayFrequency.Set(nil)
}

// UnsetPayFrequency ensures that no value is present for PayFrequency, not even an explicit nil
func (o *Employment) UnsetPayFrequency() {
	o.PayFrequency.Unset()
}

// GetPayCurrency returns the PayCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employment) GetPayCurrency() PayCurrencyEnum {
	if o == nil || o.PayCurrency.Get() == nil {
		var ret PayCurrencyEnum
		return ret
	}
	return *o.PayCurrency.Get()
}

// GetPayCurrencyOk returns a tuple with the PayCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employment) GetPayCurrencyOk() (*PayCurrencyEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayCurrency.Get(), o.PayCurrency.IsSet()
}

// HasPayCurrency returns a boolean if a field has been set.
func (o *Employment) HasPayCurrency() bool {
	if o != nil && o.PayCurrency.IsSet() {
		return true
	}

	return false
}

// SetPayCurrency gets a reference to the given NullablePayCurrencyEnum and assigns it to the PayCurrency field.
func (o *Employment) SetPayCurrency(v PayCurrencyEnum) {
	o.PayCurrency.Set(&v)
}
// SetPayCurrencyNil sets the value for PayCurrency to be an explicit nil
func (o *Employment) SetPayCurrencyNil() {
	o.PayCurrency.Set(nil)
}

// UnsetPayCurrency ensures that no value is present for PayCurrency, not even an explicit nil
func (o *Employment) UnsetPayCurrency() {
	o.PayCurrency.Unset()
}

// GetFlsaStatus returns the FlsaStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employment) GetFlsaStatus() FlsaStatusEnum {
	if o == nil || o.FlsaStatus.Get() == nil {
		var ret FlsaStatusEnum
		return ret
	}
	return *o.FlsaStatus.Get()
}

// GetFlsaStatusOk returns a tuple with the FlsaStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employment) GetFlsaStatusOk() (*FlsaStatusEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FlsaStatus.Get(), o.FlsaStatus.IsSet()
}

// HasFlsaStatus returns a boolean if a field has been set.
func (o *Employment) HasFlsaStatus() bool {
	if o != nil && o.FlsaStatus.IsSet() {
		return true
	}

	return false
}

// SetFlsaStatus gets a reference to the given NullableFlsaStatusEnum and assigns it to the FlsaStatus field.
func (o *Employment) SetFlsaStatus(v FlsaStatusEnum) {
	o.FlsaStatus.Set(&v)
}
// SetFlsaStatusNil sets the value for FlsaStatus to be an explicit nil
func (o *Employment) SetFlsaStatusNil() {
	o.FlsaStatus.Set(nil)
}

// UnsetFlsaStatus ensures that no value is present for FlsaStatus, not even an explicit nil
func (o *Employment) UnsetFlsaStatus() {
	o.FlsaStatus.Unset()
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employment) GetEffectiveDate() time.Time {
	if o == nil || o.EffectiveDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EffectiveDate.Get()
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employment) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EffectiveDate.Get(), o.EffectiveDate.IsSet()
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *Employment) HasEffectiveDate() bool {
	if o != nil && o.EffectiveDate.IsSet() {
		return true
	}

	return false
}

// SetEffectiveDate gets a reference to the given NullableTime and assigns it to the EffectiveDate field.
func (o *Employment) SetEffectiveDate(v time.Time) {
	o.EffectiveDate.Set(&v)
}
// SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil
func (o *Employment) SetEffectiveDateNil() {
	o.EffectiveDate.Set(nil)
}

// UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil
func (o *Employment) UnsetEffectiveDate() {
	o.EffectiveDate.Unset()
}

// GetEmploymentType returns the EmploymentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employment) GetEmploymentType() EmploymentTypeEnum {
	if o == nil || o.EmploymentType.Get() == nil {
		var ret EmploymentTypeEnum
		return ret
	}
	return *o.EmploymentType.Get()
}

// GetEmploymentTypeOk returns a tuple with the EmploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employment) GetEmploymentTypeOk() (*EmploymentTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmploymentType.Get(), o.EmploymentType.IsSet()
}

// HasEmploymentType returns a boolean if a field has been set.
func (o *Employment) HasEmploymentType() bool {
	if o != nil && o.EmploymentType.IsSet() {
		return true
	}

	return false
}

// SetEmploymentType gets a reference to the given NullableEmploymentTypeEnum and assigns it to the EmploymentType field.
func (o *Employment) SetEmploymentType(v EmploymentTypeEnum) {
	o.EmploymentType.Set(&v)
}
// SetEmploymentTypeNil sets the value for EmploymentType to be an explicit nil
func (o *Employment) SetEmploymentTypeNil() {
	o.EmploymentType.Set(nil)
}

// UnsetEmploymentType ensures that no value is present for EmploymentType, not even an explicit nil
func (o *Employment) UnsetEmploymentType() {
	o.EmploymentType.Unset()
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employment) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employment) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *Employment) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *Employment) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

func (o Employment) MarshalJSON() ([]byte, error) {
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
	if o.JobTitle.IsSet() {
		toSerialize["job_title"] = o.JobTitle.Get()
	}
	if o.PayRate.IsSet() {
		toSerialize["pay_rate"] = o.PayRate.Get()
	}
	if o.PayPeriod.IsSet() {
		toSerialize["pay_period"] = o.PayPeriod.Get()
	}
	if o.PayFrequency.IsSet() {
		toSerialize["pay_frequency"] = o.PayFrequency.Get()
	}
	if o.PayCurrency.IsSet() {
		toSerialize["pay_currency"] = o.PayCurrency.Get()
	}
	if o.FlsaStatus.IsSet() {
		toSerialize["flsa_status"] = o.FlsaStatus.Get()
	}
	if o.EffectiveDate.IsSet() {
		toSerialize["effective_date"] = o.EffectiveDate.Get()
	}
	if o.EmploymentType.IsSet() {
		toSerialize["employment_type"] = o.EmploymentType.Get()
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	return json.Marshal(toSerialize)
}

func (v *Employment) UnmarshalJSON(src []byte) error {
    type EmploymentUnmarshalTarget Employment

	var intermediateResult EmploymentUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = Employment(intermediateResult)
	return nil
}
type NullableEmployment struct {
	value *Employment
	isSet bool
}

func (v NullableEmployment) Get() *Employment {
	return v.value
}

func (v *NullableEmployment) Set(val *Employment) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployment) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployment(val *Employment) *NullableEmployment {
	return &NullableEmployment{value: val, isSet: true}
}

func (v NullableEmployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



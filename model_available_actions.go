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

// AvailableActions # The AvailableActions Object ### Description The `Activity` object is used to see all available model/operation combinations for an integration.  ### Usage Example Fetch all the actions available for the `Zenefits` integration.
type AvailableActions struct {
	Integration AccountIntegration `json:"integration"`
	PassthroughAvailable bool `json:"passthrough_available"`
	AvailableModelOperations *[]ModelOperation `json:"available_model_operations,omitempty"`
}

// NewAvailableActions instantiates a new AvailableActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableActions(integration AccountIntegration, passthroughAvailable bool) *AvailableActions {
	this := AvailableActions{}
	this.Integration = integration
	this.PassthroughAvailable = passthroughAvailable
	return &this
}

// NewAvailableActionsWithDefaults instantiates a new AvailableActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableActionsWithDefaults() *AvailableActions {
	this := AvailableActions{}
	return &this
}

// GetIntegration returns the Integration field value
func (o *AvailableActions) GetIntegration() AccountIntegration {
	if o == nil {
		var ret AccountIntegration
		return ret
	}

	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *AvailableActions) GetIntegrationOk() (*AccountIntegration, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value
func (o *AvailableActions) SetIntegration(v AccountIntegration) {
	o.Integration = v
}

// GetPassthroughAvailable returns the PassthroughAvailable field value
func (o *AvailableActions) GetPassthroughAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PassthroughAvailable
}

// GetPassthroughAvailableOk returns a tuple with the PassthroughAvailable field value
// and a boolean to check if the value has been set.
func (o *AvailableActions) GetPassthroughAvailableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PassthroughAvailable, true
}

// SetPassthroughAvailable sets field value
func (o *AvailableActions) SetPassthroughAvailable(v bool) {
	o.PassthroughAvailable = v
}

// GetAvailableModelOperations returns the AvailableModelOperations field value if set, zero value otherwise.
func (o *AvailableActions) GetAvailableModelOperations() []ModelOperation {
	if o == nil || o.AvailableModelOperations == nil {
		var ret []ModelOperation
		return ret
	}
	return *o.AvailableModelOperations
}

// GetAvailableModelOperationsOk returns a tuple with the AvailableModelOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableActions) GetAvailableModelOperationsOk() (*[]ModelOperation, bool) {
	if o == nil || o.AvailableModelOperations == nil {
		return nil, false
	}
	return o.AvailableModelOperations, true
}

// HasAvailableModelOperations returns a boolean if a field has been set.
func (o *AvailableActions) HasAvailableModelOperations() bool {
	if o != nil && o.AvailableModelOperations != nil {
		return true
	}

	return false
}

// SetAvailableModelOperations gets a reference to the given []ModelOperation and assigns it to the AvailableModelOperations field.
func (o *AvailableActions) SetAvailableModelOperations(v []ModelOperation) {
	o.AvailableModelOperations = &v
}

func (o AvailableActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["integration"] = o.Integration
	}
	if true {
		toSerialize["passthrough_available"] = o.PassthroughAvailable
	}
	if o.AvailableModelOperations != nil {
		toSerialize["available_model_operations"] = o.AvailableModelOperations
	}
	return json.Marshal(toSerialize)
}

type NullableAvailableActions struct {
	value *AvailableActions
	isSet bool
}

func (v NullableAvailableActions) Get() *AvailableActions {
	return v.value
}

func (v *NullableAvailableActions) Set(val *AvailableActions) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableActions) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableActions(val *AvailableActions) *NullableAvailableActions {
	return &NullableAvailableActions{value: val, isSet: true}
}

func (v NullableAvailableActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// GenerateRemoteKeyRequest # The GenerateRemoteKey Object ### Description The `GenerateRemoteKey` object is used to represent a request for a new remote key.  ### Usage Example Post a `GenerateRemoteKey` to create a new remote key.
type GenerateRemoteKeyRequest struct {
	Name string `json:"name"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewGenerateRemoteKeyRequest instantiates a new GenerateRemoteKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateRemoteKeyRequest(name string) *GenerateRemoteKeyRequest {
	this := GenerateRemoteKeyRequest{}
	this.Name = name
	return &this
}

// NewGenerateRemoteKeyRequestWithDefaults instantiates a new GenerateRemoteKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateRemoteKeyRequestWithDefaults() *GenerateRemoteKeyRequest {
	this := GenerateRemoteKeyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *GenerateRemoteKeyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GenerateRemoteKeyRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GenerateRemoteKeyRequest) SetName(v string) {
	o.Name = v
}

func (o GenerateRemoteKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

func (v *GenerateRemoteKeyRequest) UnmarshalJSON(src []byte) error {
    type GenerateRemoteKeyRequestUnmarshalTarget GenerateRemoteKeyRequest

	var intermediateResult GenerateRemoteKeyRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = GenerateRemoteKeyRequest(intermediateResult)
	return nil
}
type NullableGenerateRemoteKeyRequest struct {
	value *GenerateRemoteKeyRequest
	isSet bool
}

func (v NullableGenerateRemoteKeyRequest) Get() *GenerateRemoteKeyRequest {
	return v.value
}

func (v *NullableGenerateRemoteKeyRequest) Set(val *GenerateRemoteKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateRemoteKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateRemoteKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateRemoteKeyRequest(val *GenerateRemoteKeyRequest) *NullableGenerateRemoteKeyRequest {
	return &NullableGenerateRemoteKeyRequest{value: val, isSet: true}
}

func (v NullableGenerateRemoteKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateRemoteKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



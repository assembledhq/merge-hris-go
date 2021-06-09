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

// RemoteData struct for RemoteData
type RemoteData struct {
	Path string `json:"path"`
	Data *map[string]interface{} `json:"data,omitempty"`
}

// NewRemoteData instantiates a new RemoteData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteData(path string) *RemoteData {
	this := RemoteData{}
	this.Path = path
	return &this
}

// NewRemoteDataWithDefaults instantiates a new RemoteData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteDataWithDefaults() *RemoteData {
	this := RemoteData{}
	return &this
}

// GetPath returns the Path field value
func (o *RemoteData) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *RemoteData) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *RemoteData) SetPath(v string) {
	o.Path = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RemoteData) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteData) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RemoteData) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *RemoteData) SetData(v map[string]interface{}) {
	o.Data = &v
}

func (o RemoteData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteData struct {
	value *RemoteData
	isSet bool
}

func (v NullableRemoteData) Get() *RemoteData {
	return v.value
}

func (v *NullableRemoteData) Set(val *RemoteData) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteData) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteData(val *RemoteData) *NullableRemoteData {
	return &NullableRemoteData{value: val, isSet: true}
}

func (v NullableRemoteData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



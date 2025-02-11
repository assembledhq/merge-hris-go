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

// DebugModelLogSummary struct for DebugModelLogSummary
type DebugModelLogSummary struct {
	Url string `json:"url"`
	Method string `json:"method"`
	StatusCode int32 `json:"status_code"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewDebugModelLogSummary instantiates a new DebugModelLogSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebugModelLogSummary(url string, method string, statusCode int32) *DebugModelLogSummary {
	this := DebugModelLogSummary{}
	this.Url = url
	this.Method = method
	this.StatusCode = statusCode
	return &this
}

// NewDebugModelLogSummaryWithDefaults instantiates a new DebugModelLogSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebugModelLogSummaryWithDefaults() *DebugModelLogSummary {
	this := DebugModelLogSummary{}
	return &this
}

// GetUrl returns the Url field value
func (o *DebugModelLogSummary) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DebugModelLogSummary) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DebugModelLogSummary) SetUrl(v string) {
	o.Url = v
}

// GetMethod returns the Method field value
func (o *DebugModelLogSummary) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *DebugModelLogSummary) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *DebugModelLogSummary) SetMethod(v string) {
	o.Method = v
}

// GetStatusCode returns the StatusCode field value
func (o *DebugModelLogSummary) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *DebugModelLogSummary) GetStatusCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *DebugModelLogSummary) SetStatusCode(v int32) {
	o.StatusCode = v
}

func (o DebugModelLogSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["status_code"] = o.StatusCode
	}
	return json.Marshal(toSerialize)
}

func (v *DebugModelLogSummary) UnmarshalJSON(src []byte) error {
    type DebugModelLogSummaryUnmarshalTarget DebugModelLogSummary

	var intermediateResult DebugModelLogSummaryUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = DebugModelLogSummary(intermediateResult)
	return nil
}
type NullableDebugModelLogSummary struct {
	value *DebugModelLogSummary
	isSet bool
}

func (v NullableDebugModelLogSummary) Get() *DebugModelLogSummary {
	return v.value
}

func (v *NullableDebugModelLogSummary) Set(val *DebugModelLogSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDebugModelLogSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDebugModelLogSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebugModelLogSummary(val *DebugModelLogSummary) *NullableDebugModelLogSummary {
	return &NullableDebugModelLogSummary{value: val, isSet: true}
}

func (v NullableDebugModelLogSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebugModelLogSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



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

// MultipartFormFieldRequest # The MultipartFormField Object ### Description The `MultipartFormField` object is used to represent fields in an HTTP request using `multipart/form-data`.  ### Usage Example Create a `MultipartFormField` to define a multipart form entry.
type MultipartFormFieldRequest struct {
	// The name of the form field
	Name string `json:"name"`
	// The data for the form field.
	Data string `json:"data"`
	// The encoding of the value of `data`. Defaults to `RAW` if not defined.
	Encoding NullableEncodingEnum `json:"encoding,omitempty"`
	// The file name of the form field, if the field is for a file.
	FileName NullableString `json:"file_name,omitempty"`
	// The MIME type of the file, if the field is for a file.
	ContentType NullableString `json:"content_type,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewMultipartFormFieldRequest instantiates a new MultipartFormFieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipartFormFieldRequest(name string, data string) *MultipartFormFieldRequest {
	this := MultipartFormFieldRequest{}
	this.Name = name
	this.Data = data
	return &this
}

// NewMultipartFormFieldRequestWithDefaults instantiates a new MultipartFormFieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipartFormFieldRequestWithDefaults() *MultipartFormFieldRequest {
	this := MultipartFormFieldRequest{}
	return &this
}

// GetName returns the Name field value
func (o *MultipartFormFieldRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MultipartFormFieldRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MultipartFormFieldRequest) SetName(v string) {
	o.Name = v
}

// GetData returns the Data field value
func (o *MultipartFormFieldRequest) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MultipartFormFieldRequest) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MultipartFormFieldRequest) SetData(v string) {
	o.Data = v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MultipartFormFieldRequest) GetEncoding() EncodingEnum {
	if o == nil || o.Encoding.Get() == nil {
		var ret EncodingEnum
		return ret
	}
	return *o.Encoding.Get()
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultipartFormFieldRequest) GetEncodingOk() (*EncodingEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Encoding.Get(), o.Encoding.IsSet()
}

// HasEncoding returns a boolean if a field has been set.
func (o *MultipartFormFieldRequest) HasEncoding() bool {
	if o != nil && o.Encoding.IsSet() {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given NullableEncodingEnum and assigns it to the Encoding field.
func (o *MultipartFormFieldRequest) SetEncoding(v EncodingEnum) {
	o.Encoding.Set(&v)
}
// SetEncodingNil sets the value for Encoding to be an explicit nil
func (o *MultipartFormFieldRequest) SetEncodingNil() {
	o.Encoding.Set(nil)
}

// UnsetEncoding ensures that no value is present for Encoding, not even an explicit nil
func (o *MultipartFormFieldRequest) UnsetEncoding() {
	o.Encoding.Unset()
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MultipartFormFieldRequest) GetFileName() string {
	if o == nil || o.FileName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultipartFormFieldRequest) GetFileNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *MultipartFormFieldRequest) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *MultipartFormFieldRequest) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *MultipartFormFieldRequest) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *MultipartFormFieldRequest) UnsetFileName() {
	o.FileName.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MultipartFormFieldRequest) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultipartFormFieldRequest) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *MultipartFormFieldRequest) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *MultipartFormFieldRequest) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *MultipartFormFieldRequest) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *MultipartFormFieldRequest) UnsetContentType() {
	o.ContentType.Unset()
}

func (o MultipartFormFieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Encoding.IsSet() {
		toSerialize["encoding"] = o.Encoding.Get()
	}
	if o.FileName.IsSet() {
		toSerialize["file_name"] = o.FileName.Get()
	}
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	return json.Marshal(toSerialize)
}

func (v *MultipartFormFieldRequest) UnmarshalJSON(src []byte) error {
    type MultipartFormFieldRequestUnmarshalTarget MultipartFormFieldRequest

	var intermediateResult MultipartFormFieldRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = MultipartFormFieldRequest(intermediateResult)
	return nil
}
type NullableMultipartFormFieldRequest struct {
	value *MultipartFormFieldRequest
	isSet bool
}

func (v NullableMultipartFormFieldRequest) Get() *MultipartFormFieldRequest {
	return v.value
}

func (v *NullableMultipartFormFieldRequest) Set(val *MultipartFormFieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipartFormFieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipartFormFieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipartFormFieldRequest(val *MultipartFormFieldRequest) *NullableMultipartFormFieldRequest {
	return &NullableMultipartFormFieldRequest{value: val, isSet: true}
}

func (v NullableMultipartFormFieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipartFormFieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



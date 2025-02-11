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

// Team # The Team Object ### Description The `Team` object is used to represent a Team within a company. `Employee` objects are often grouped this way. Note that in the Merge HRIS API, company subdivisions are all represented with `Teams`, rather than `Teams` and `Departments`.  ### Usage Example If you're building a way to filter by `Team`, you'd hit the `GET Teams` endpoint to fetch the `Teams`, and then use the `ID` of the team your user selects to filter the `GET Employees` endpoint.
type Team struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The team's name.
	Name NullableString `json:"name,omitempty"`
	ParentTeam NullableString `json:"parent_team,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewTeam instantiates a new Team object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeam() *Team {
	this := Team{}
	return &this
}

// NewTeamWithDefaults instantiates a new Team object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamWithDefaults() *Team {
	this := Team{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Team) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Team) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Team) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Team) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Team) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *Team) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *Team) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *Team) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *Team) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Team) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Team) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Team) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Team) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Team) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Team) UnsetName() {
	o.Name.Unset()
}

// GetParentTeam returns the ParentTeam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Team) GetParentTeam() string {
	if o == nil || o.ParentTeam.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentTeam.Get()
}

// GetParentTeamOk returns a tuple with the ParentTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Team) GetParentTeamOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentTeam.Get(), o.ParentTeam.IsSet()
}

// HasParentTeam returns a boolean if a field has been set.
func (o *Team) HasParentTeam() bool {
	if o != nil && o.ParentTeam.IsSet() {
		return true
	}

	return false
}

// SetParentTeam gets a reference to the given NullableString and assigns it to the ParentTeam field.
func (o *Team) SetParentTeam(v string) {
	o.ParentTeam.Set(&v)
}
// SetParentTeamNil sets the value for ParentTeam to be an explicit nil
func (o *Team) SetParentTeamNil() {
	o.ParentTeam.Set(nil)
}

// UnsetParentTeam ensures that no value is present for ParentTeam, not even an explicit nil
func (o *Team) UnsetParentTeam() {
	o.ParentTeam.Unset()
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Team) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Team) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *Team) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *Team) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

func (o Team) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ParentTeam.IsSet() {
		toSerialize["parent_team"] = o.ParentTeam.Get()
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	return json.Marshal(toSerialize)
}

func (v *Team) UnmarshalJSON(src []byte) error {
    type TeamUnmarshalTarget Team

	var intermediateResult TeamUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = Team(intermediateResult)
	return nil
}
type NullableTeam struct {
	value *Team
	isSet bool
}

func (v NullableTeam) Get() *Team {
	return v.value
}

func (v *NullableTeam) Set(val *Team) {
	v.value = val
	v.isSet = true
}

func (v NullableTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeam(val *Team) *NullableTeam {
	return &NullableTeam{value: val, isSet: true}
}

func (v NullableTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



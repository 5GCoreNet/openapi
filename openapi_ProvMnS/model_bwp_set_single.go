/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the BWPSetSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BWPSetSingle{}

// BWPSetSingle struct for BWPSetSingle
type BWPSetSingle struct {
	Top
	BWPlist []string `json:"bWPlist,omitempty"`
}

// NewBWPSetSingle instantiates a new BWPSetSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBWPSetSingle(id NullableString) *BWPSetSingle {
	this := BWPSetSingle{}
	this.Id = id
	return &this
}

// NewBWPSetSingleWithDefaults instantiates a new BWPSetSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBWPSetSingleWithDefaults() *BWPSetSingle {
	this := BWPSetSingle{}
	return &this
}

// GetBWPlist returns the BWPlist field value if set, zero value otherwise.
func (o *BWPSetSingle) GetBWPlist() []string {
	if o == nil || IsNil(o.BWPlist) {
		var ret []string
		return ret
	}
	return o.BWPlist
}

// GetBWPlistOk returns a tuple with the BWPlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BWPSetSingle) GetBWPlistOk() ([]string, bool) {
	if o == nil || IsNil(o.BWPlist) {
		return nil, false
	}
	return o.BWPlist, true
}

// HasBWPlist returns a boolean if a field has been set.
func (o *BWPSetSingle) HasBWPlist() bool {
	if o != nil && !IsNil(o.BWPlist) {
		return true
	}

	return false
}

// SetBWPlist gets a reference to the given []string and assigns it to the BWPlist field.
func (o *BWPSetSingle) SetBWPlist(v []string) {
	o.BWPlist = v
}

func (o BWPSetSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BWPSetSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop, errTop := json.Marshal(o.Top)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	errTop = json.Unmarshal([]byte(serializedTop), &toSerialize)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	if !IsNil(o.BWPlist) {
		toSerialize["bWPlist"] = o.BWPlist
	}
	return toSerialize, nil
}

type NullableBWPSetSingle struct {
	value *BWPSetSingle
	isSet bool
}

func (v NullableBWPSetSingle) Get() *BWPSetSingle {
	return v.value
}

func (v *NullableBWPSetSingle) Set(val *BWPSetSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableBWPSetSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableBWPSetSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBWPSetSingle(val *BWPSetSingle) *NullableBWPSetSingle {
	return &NullableBWPSetSingle{value: val, isSet: true}
}

func (v NullableBWPSetSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBWPSetSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

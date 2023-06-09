/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
)

// checks if the AlarmsGet200ResponseValueAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlarmsGet200ResponseValueAllOf1{}

// AlarmsGet200ResponseValueAllOf1 struct for AlarmsGet200ResponseValueAllOf1
type AlarmsGet200ResponseValueAllOf1 struct {
	// Collection of comments. The comment identifiers are allocated by the MnS producer and used as key in the map.
	Comments *map[string]Comment `json:"comments,omitempty"`
}

// NewAlarmsGet200ResponseValueAllOf1 instantiates a new AlarmsGet200ResponseValueAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmsGet200ResponseValueAllOf1() *AlarmsGet200ResponseValueAllOf1 {
	this := AlarmsGet200ResponseValueAllOf1{}
	return &this
}

// NewAlarmsGet200ResponseValueAllOf1WithDefaults instantiates a new AlarmsGet200ResponseValueAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmsGet200ResponseValueAllOf1WithDefaults() *AlarmsGet200ResponseValueAllOf1 {
	this := AlarmsGet200ResponseValueAllOf1{}
	return &this
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *AlarmsGet200ResponseValueAllOf1) GetComments() map[string]Comment {
	if o == nil || IsNil(o.Comments) {
		var ret map[string]Comment
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmsGet200ResponseValueAllOf1) GetCommentsOk() (*map[string]Comment, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *AlarmsGet200ResponseValueAllOf1) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given map[string]Comment and assigns it to the Comments field.
func (o *AlarmsGet200ResponseValueAllOf1) SetComments(v map[string]Comment) {
	o.Comments = &v
}

func (o AlarmsGet200ResponseValueAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlarmsGet200ResponseValueAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	return toSerialize, nil
}

type NullableAlarmsGet200ResponseValueAllOf1 struct {
	value *AlarmsGet200ResponseValueAllOf1
	isSet bool
}

func (v NullableAlarmsGet200ResponseValueAllOf1) Get() *AlarmsGet200ResponseValueAllOf1 {
	return v.value
}

func (v *NullableAlarmsGet200ResponseValueAllOf1) Set(val *AlarmsGet200ResponseValueAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmsGet200ResponseValueAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmsGet200ResponseValueAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmsGet200ResponseValueAllOf1(val *AlarmsGet200ResponseValueAllOf1) *NullableAlarmsGet200ResponseValueAllOf1 {
	return &NullableAlarmsGet200ResponseValueAllOf1{value: val, isSet: true}
}

func (v NullableAlarmsGet200ResponseValueAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmsGet200ResponseValueAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

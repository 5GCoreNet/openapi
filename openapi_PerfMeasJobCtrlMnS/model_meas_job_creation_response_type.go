/*
TS 28.550 Performance Measurement Job Control Service

OAS 3.0.1 specification of the Performance Measurement Job Control Service @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 16.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PerfMeasJobCtrlMnS

import (
	"encoding/json"
)

// checks if the MeasJobCreationResponseType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeasJobCreationResponseType{}

// MeasJobCreationResponseType struct for MeasJobCreationResponseType
type MeasJobCreationResponseType struct {
	UnsupportedList []UnsupportedMeasType `json:"unsupportedList,omitempty"`
}

// NewMeasJobCreationResponseType instantiates a new MeasJobCreationResponseType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasJobCreationResponseType() *MeasJobCreationResponseType {
	this := MeasJobCreationResponseType{}
	return &this
}

// NewMeasJobCreationResponseTypeWithDefaults instantiates a new MeasJobCreationResponseType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasJobCreationResponseTypeWithDefaults() *MeasJobCreationResponseType {
	this := MeasJobCreationResponseType{}
	return &this
}

// GetUnsupportedList returns the UnsupportedList field value if set, zero value otherwise.
func (o *MeasJobCreationResponseType) GetUnsupportedList() []UnsupportedMeasType {
	if o == nil || IsNil(o.UnsupportedList) {
		var ret []UnsupportedMeasType
		return ret
	}
	return o.UnsupportedList
}

// GetUnsupportedListOk returns a tuple with the UnsupportedList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationResponseType) GetUnsupportedListOk() ([]UnsupportedMeasType, bool) {
	if o == nil || IsNil(o.UnsupportedList) {
		return nil, false
	}
	return o.UnsupportedList, true
}

// HasUnsupportedList returns a boolean if a field has been set.
func (o *MeasJobCreationResponseType) HasUnsupportedList() bool {
	if o != nil && !IsNil(o.UnsupportedList) {
		return true
	}

	return false
}

// SetUnsupportedList gets a reference to the given []UnsupportedMeasType and assigns it to the UnsupportedList field.
func (o *MeasJobCreationResponseType) SetUnsupportedList(v []UnsupportedMeasType) {
	o.UnsupportedList = v
}

func (o MeasJobCreationResponseType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeasJobCreationResponseType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UnsupportedList) {
		toSerialize["unsupportedList"] = o.UnsupportedList
	}
	return toSerialize, nil
}

type NullableMeasJobCreationResponseType struct {
	value *MeasJobCreationResponseType
	isSet bool
}

func (v NullableMeasJobCreationResponseType) Get() *MeasJobCreationResponseType {
	return v.value
}

func (v *NullableMeasJobCreationResponseType) Set(val *MeasJobCreationResponseType) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasJobCreationResponseType) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasJobCreationResponseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasJobCreationResponseType(val *MeasJobCreationResponseType) *NullableMeasJobCreationResponseType {
	return &NullableMeasJobCreationResponseType{value: val, isSet: true}
}

func (v NullableMeasJobCreationResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasJobCreationResponseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

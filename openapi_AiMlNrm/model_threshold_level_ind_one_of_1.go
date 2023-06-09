/*
AI/ML NRM

OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AiMlNrm

import (
	"encoding/json"
)

// checks if the ThresholdLevelIndOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThresholdLevelIndOneOf1{}

// ThresholdLevelIndOneOf1 struct for ThresholdLevelIndOneOf1
type ThresholdLevelIndOneOf1 struct {
	Down *ThresholdHysteresis `json:"down,omitempty"`
}

// NewThresholdLevelIndOneOf1 instantiates a new ThresholdLevelIndOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdLevelIndOneOf1() *ThresholdLevelIndOneOf1 {
	this := ThresholdLevelIndOneOf1{}
	return &this
}

// NewThresholdLevelIndOneOf1WithDefaults instantiates a new ThresholdLevelIndOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdLevelIndOneOf1WithDefaults() *ThresholdLevelIndOneOf1 {
	this := ThresholdLevelIndOneOf1{}
	return &this
}

// GetDown returns the Down field value if set, zero value otherwise.
func (o *ThresholdLevelIndOneOf1) GetDown() ThresholdHysteresis {
	if o == nil || IsNil(o.Down) {
		var ret ThresholdHysteresis
		return ret
	}
	return *o.Down
}

// GetDownOk returns a tuple with the Down field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevelIndOneOf1) GetDownOk() (*ThresholdHysteresis, bool) {
	if o == nil || IsNil(o.Down) {
		return nil, false
	}
	return o.Down, true
}

// HasDown returns a boolean if a field has been set.
func (o *ThresholdLevelIndOneOf1) HasDown() bool {
	if o != nil && !IsNil(o.Down) {
		return true
	}

	return false
}

// SetDown gets a reference to the given ThresholdHysteresis and assigns it to the Down field.
func (o *ThresholdLevelIndOneOf1) SetDown(v ThresholdHysteresis) {
	o.Down = &v
}

func (o ThresholdLevelIndOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThresholdLevelIndOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Down) {
		toSerialize["down"] = o.Down
	}
	return toSerialize, nil
}

type NullableThresholdLevelIndOneOf1 struct {
	value *ThresholdLevelIndOneOf1
	isSet bool
}

func (v NullableThresholdLevelIndOneOf1) Get() *ThresholdLevelIndOneOf1 {
	return v.value
}

func (v *NullableThresholdLevelIndOneOf1) Set(val *ThresholdLevelIndOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdLevelIndOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdLevelIndOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdLevelIndOneOf1(val *ThresholdLevelIndOneOf1) *NullableThresholdLevelIndOneOf1 {
	return &NullableThresholdLevelIndOneOf1{value: val, isSet: true}
}

func (v NullableThresholdLevelIndOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdLevelIndOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

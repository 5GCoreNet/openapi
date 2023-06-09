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

// checks if the ThresholdLevelIndOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThresholdLevelIndOneOf{}

// ThresholdLevelIndOneOf struct for ThresholdLevelIndOneOf
type ThresholdLevelIndOneOf struct {
	Up *ThresholdHysteresis `json:"up,omitempty"`
}

// NewThresholdLevelIndOneOf instantiates a new ThresholdLevelIndOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdLevelIndOneOf() *ThresholdLevelIndOneOf {
	this := ThresholdLevelIndOneOf{}
	return &this
}

// NewThresholdLevelIndOneOfWithDefaults instantiates a new ThresholdLevelIndOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdLevelIndOneOfWithDefaults() *ThresholdLevelIndOneOf {
	this := ThresholdLevelIndOneOf{}
	return &this
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *ThresholdLevelIndOneOf) GetUp() ThresholdHysteresis {
	if o == nil || IsNil(o.Up) {
		var ret ThresholdHysteresis
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevelIndOneOf) GetUpOk() (*ThresholdHysteresis, bool) {
	if o == nil || IsNil(o.Up) {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *ThresholdLevelIndOneOf) HasUp() bool {
	if o != nil && !IsNil(o.Up) {
		return true
	}

	return false
}

// SetUp gets a reference to the given ThresholdHysteresis and assigns it to the Up field.
func (o *ThresholdLevelIndOneOf) SetUp(v ThresholdHysteresis) {
	o.Up = &v
}

func (o ThresholdLevelIndOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThresholdLevelIndOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Up) {
		toSerialize["up"] = o.Up
	}
	return toSerialize, nil
}

type NullableThresholdLevelIndOneOf struct {
	value *ThresholdLevelIndOneOf
	isSet bool
}

func (v NullableThresholdLevelIndOneOf) Get() *ThresholdLevelIndOneOf {
	return v.value
}

func (v *NullableThresholdLevelIndOneOf) Set(val *ThresholdLevelIndOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdLevelIndOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdLevelIndOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdLevelIndOneOf(val *ThresholdLevelIndOneOf) *NullableThresholdLevelIndOneOf {
	return &NullableThresholdLevelIndOneOf{value: val, isSet: true}
}

func (v NullableThresholdLevelIndOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdLevelIndOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

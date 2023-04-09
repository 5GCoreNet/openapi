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

// checks if the ThresholdInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThresholdInfo{}

// ThresholdInfo struct for ThresholdInfo
type ThresholdInfo struct {
	ThresholdDirection *string                      `json:"thresholdDirection,omitempty"`
	ThresholdValue     *ThresholdInfoThresholdValue `json:"thresholdValue,omitempty"`
	Hysteresis         *ThresholdInfoHysteresis     `json:"hysteresis,omitempty"`
}

// NewThresholdInfo instantiates a new ThresholdInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdInfo() *ThresholdInfo {
	this := ThresholdInfo{}
	return &this
}

// NewThresholdInfoWithDefaults instantiates a new ThresholdInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdInfoWithDefaults() *ThresholdInfo {
	this := ThresholdInfo{}
	return &this
}

// GetThresholdDirection returns the ThresholdDirection field value if set, zero value otherwise.
func (o *ThresholdInfo) GetThresholdDirection() string {
	if o == nil || IsNil(o.ThresholdDirection) {
		var ret string
		return ret
	}
	return *o.ThresholdDirection
}

// GetThresholdDirectionOk returns a tuple with the ThresholdDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdInfo) GetThresholdDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.ThresholdDirection) {
		return nil, false
	}
	return o.ThresholdDirection, true
}

// HasThresholdDirection returns a boolean if a field has been set.
func (o *ThresholdInfo) HasThresholdDirection() bool {
	if o != nil && !IsNil(o.ThresholdDirection) {
		return true
	}

	return false
}

// SetThresholdDirection gets a reference to the given string and assigns it to the ThresholdDirection field.
func (o *ThresholdInfo) SetThresholdDirection(v string) {
	o.ThresholdDirection = &v
}

// GetThresholdValue returns the ThresholdValue field value if set, zero value otherwise.
func (o *ThresholdInfo) GetThresholdValue() ThresholdInfoThresholdValue {
	if o == nil || IsNil(o.ThresholdValue) {
		var ret ThresholdInfoThresholdValue
		return ret
	}
	return *o.ThresholdValue
}

// GetThresholdValueOk returns a tuple with the ThresholdValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdInfo) GetThresholdValueOk() (*ThresholdInfoThresholdValue, bool) {
	if o == nil || IsNil(o.ThresholdValue) {
		return nil, false
	}
	return o.ThresholdValue, true
}

// HasThresholdValue returns a boolean if a field has been set.
func (o *ThresholdInfo) HasThresholdValue() bool {
	if o != nil && !IsNil(o.ThresholdValue) {
		return true
	}

	return false
}

// SetThresholdValue gets a reference to the given ThresholdInfoThresholdValue and assigns it to the ThresholdValue field.
func (o *ThresholdInfo) SetThresholdValue(v ThresholdInfoThresholdValue) {
	o.ThresholdValue = &v
}

// GetHysteresis returns the Hysteresis field value if set, zero value otherwise.
func (o *ThresholdInfo) GetHysteresis() ThresholdInfoHysteresis {
	if o == nil || IsNil(o.Hysteresis) {
		var ret ThresholdInfoHysteresis
		return ret
	}
	return *o.Hysteresis
}

// GetHysteresisOk returns a tuple with the Hysteresis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdInfo) GetHysteresisOk() (*ThresholdInfoHysteresis, bool) {
	if o == nil || IsNil(o.Hysteresis) {
		return nil, false
	}
	return o.Hysteresis, true
}

// HasHysteresis returns a boolean if a field has been set.
func (o *ThresholdInfo) HasHysteresis() bool {
	if o != nil && !IsNil(o.Hysteresis) {
		return true
	}

	return false
}

// SetHysteresis gets a reference to the given ThresholdInfoHysteresis and assigns it to the Hysteresis field.
func (o *ThresholdInfo) SetHysteresis(v ThresholdInfoHysteresis) {
	o.Hysteresis = &v
}

func (o ThresholdInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThresholdInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThresholdDirection) {
		toSerialize["thresholdDirection"] = o.ThresholdDirection
	}
	if !IsNil(o.ThresholdValue) {
		toSerialize["thresholdValue"] = o.ThresholdValue
	}
	if !IsNil(o.Hysteresis) {
		toSerialize["hysteresis"] = o.Hysteresis
	}
	return toSerialize, nil
}

type NullableThresholdInfo struct {
	value *ThresholdInfo
	isSet bool
}

func (v NullableThresholdInfo) Get() *ThresholdInfo {
	return v.value
}

func (v *NullableThresholdInfo) Set(val *ThresholdInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdInfo(val *ThresholdInfo) *NullableThresholdInfo {
	return &NullableThresholdInfo{value: val, isSet: true}
}

func (v NullableThresholdInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
	"time"
)

// checks if the ThresholdInfo1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThresholdInfo1{}

// ThresholdInfo1 struct for ThresholdInfo1
type ThresholdInfo1 struct {
	ObservedMeasurement string             `json:"observedMeasurement"`
	ObservedValue       float32            `json:"observedValue"`
	ThresholdLevel      *ThresholdLevelInd `json:"thresholdLevel,omitempty"`
	ArmTime             *time.Time         `json:"armTime,omitempty"`
}

// NewThresholdInfo1 instantiates a new ThresholdInfo1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdInfo1(observedMeasurement string, observedValue float32) *ThresholdInfo1 {
	this := ThresholdInfo1{}
	this.ObservedMeasurement = observedMeasurement
	this.ObservedValue = observedValue
	return &this
}

// NewThresholdInfo1WithDefaults instantiates a new ThresholdInfo1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdInfo1WithDefaults() *ThresholdInfo1 {
	this := ThresholdInfo1{}
	return &this
}

// GetObservedMeasurement returns the ObservedMeasurement field value
func (o *ThresholdInfo1) GetObservedMeasurement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObservedMeasurement
}

// GetObservedMeasurementOk returns a tuple with the ObservedMeasurement field value
// and a boolean to check if the value has been set.
func (o *ThresholdInfo1) GetObservedMeasurementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObservedMeasurement, true
}

// SetObservedMeasurement sets field value
func (o *ThresholdInfo1) SetObservedMeasurement(v string) {
	o.ObservedMeasurement = v
}

// GetObservedValue returns the ObservedValue field value
func (o *ThresholdInfo1) GetObservedValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ObservedValue
}

// GetObservedValueOk returns a tuple with the ObservedValue field value
// and a boolean to check if the value has been set.
func (o *ThresholdInfo1) GetObservedValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObservedValue, true
}

// SetObservedValue sets field value
func (o *ThresholdInfo1) SetObservedValue(v float32) {
	o.ObservedValue = v
}

// GetThresholdLevel returns the ThresholdLevel field value if set, zero value otherwise.
func (o *ThresholdInfo1) GetThresholdLevel() ThresholdLevelInd {
	if o == nil || IsNil(o.ThresholdLevel) {
		var ret ThresholdLevelInd
		return ret
	}
	return *o.ThresholdLevel
}

// GetThresholdLevelOk returns a tuple with the ThresholdLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdInfo1) GetThresholdLevelOk() (*ThresholdLevelInd, bool) {
	if o == nil || IsNil(o.ThresholdLevel) {
		return nil, false
	}
	return o.ThresholdLevel, true
}

// HasThresholdLevel returns a boolean if a field has been set.
func (o *ThresholdInfo1) HasThresholdLevel() bool {
	if o != nil && !IsNil(o.ThresholdLevel) {
		return true
	}

	return false
}

// SetThresholdLevel gets a reference to the given ThresholdLevelInd and assigns it to the ThresholdLevel field.
func (o *ThresholdInfo1) SetThresholdLevel(v ThresholdLevelInd) {
	o.ThresholdLevel = &v
}

// GetArmTime returns the ArmTime field value if set, zero value otherwise.
func (o *ThresholdInfo1) GetArmTime() time.Time {
	if o == nil || IsNil(o.ArmTime) {
		var ret time.Time
		return ret
	}
	return *o.ArmTime
}

// GetArmTimeOk returns a tuple with the ArmTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdInfo1) GetArmTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ArmTime) {
		return nil, false
	}
	return o.ArmTime, true
}

// HasArmTime returns a boolean if a field has been set.
func (o *ThresholdInfo1) HasArmTime() bool {
	if o != nil && !IsNil(o.ArmTime) {
		return true
	}

	return false
}

// SetArmTime gets a reference to the given time.Time and assigns it to the ArmTime field.
func (o *ThresholdInfo1) SetArmTime(v time.Time) {
	o.ArmTime = &v
}

func (o ThresholdInfo1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThresholdInfo1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["observedMeasurement"] = o.ObservedMeasurement
	toSerialize["observedValue"] = o.ObservedValue
	if !IsNil(o.ThresholdLevel) {
		toSerialize["thresholdLevel"] = o.ThresholdLevel
	}
	if !IsNil(o.ArmTime) {
		toSerialize["armTime"] = o.ArmTime
	}
	return toSerialize, nil
}

type NullableThresholdInfo1 struct {
	value *ThresholdInfo1
	isSet bool
}

func (v NullableThresholdInfo1) Get() *ThresholdInfo1 {
	return v.value
}

func (v *NullableThresholdInfo1) Set(val *ThresholdInfo1) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdInfo1) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdInfo1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdInfo1(val *ThresholdInfo1) *NullableThresholdInfo1 {
	return &NullableThresholdInfo1{value: val, isSet: true}
}

func (v NullableThresholdInfo1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdInfo1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

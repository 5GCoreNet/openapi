/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
	"time"
)

// checks if the ThresholdInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThresholdInfo{}

// ThresholdInfo struct for ThresholdInfo
type ThresholdInfo struct {
	ObservedMeasurement string `json:"observedMeasurement"`
	ObservedValue float32 `json:"observedValue"`
	ThresholdLevel *ThresholdLevelInd `json:"thresholdLevel,omitempty"`
	ArmTime *time.Time `json:"armTime,omitempty"`
}

// NewThresholdInfo instantiates a new ThresholdInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdInfo(observedMeasurement string, observedValue float32) *ThresholdInfo {
	this := ThresholdInfo{}
	this.ObservedMeasurement = observedMeasurement
	this.ObservedValue = observedValue
	return &this
}

// NewThresholdInfoWithDefaults instantiates a new ThresholdInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdInfoWithDefaults() *ThresholdInfo {
	this := ThresholdInfo{}
	return &this
}

// GetObservedMeasurement returns the ObservedMeasurement field value
func (o *ThresholdInfo) GetObservedMeasurement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObservedMeasurement
}

// GetObservedMeasurementOk returns a tuple with the ObservedMeasurement field value
// and a boolean to check if the value has been set.
func (o *ThresholdInfo) GetObservedMeasurementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObservedMeasurement, true
}

// SetObservedMeasurement sets field value
func (o *ThresholdInfo) SetObservedMeasurement(v string) {
	o.ObservedMeasurement = v
}

// GetObservedValue returns the ObservedValue field value
func (o *ThresholdInfo) GetObservedValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ObservedValue
}

// GetObservedValueOk returns a tuple with the ObservedValue field value
// and a boolean to check if the value has been set.
func (o *ThresholdInfo) GetObservedValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObservedValue, true
}

// SetObservedValue sets field value
func (o *ThresholdInfo) SetObservedValue(v float32) {
	o.ObservedValue = v
}

// GetThresholdLevel returns the ThresholdLevel field value if set, zero value otherwise.
func (o *ThresholdInfo) GetThresholdLevel() ThresholdLevelInd {
	if o == nil || isNil(o.ThresholdLevel) {
		var ret ThresholdLevelInd
		return ret
	}
	return *o.ThresholdLevel
}

// GetThresholdLevelOk returns a tuple with the ThresholdLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdInfo) GetThresholdLevelOk() (*ThresholdLevelInd, bool) {
	if o == nil || isNil(o.ThresholdLevel) {
		return nil, false
	}
	return o.ThresholdLevel, true
}

// HasThresholdLevel returns a boolean if a field has been set.
func (o *ThresholdInfo) HasThresholdLevel() bool {
	if o != nil && !isNil(o.ThresholdLevel) {
		return true
	}

	return false
}

// SetThresholdLevel gets a reference to the given ThresholdLevelInd and assigns it to the ThresholdLevel field.
func (o *ThresholdInfo) SetThresholdLevel(v ThresholdLevelInd) {
	o.ThresholdLevel = &v
}

// GetArmTime returns the ArmTime field value if set, zero value otherwise.
func (o *ThresholdInfo) GetArmTime() time.Time {
	if o == nil || isNil(o.ArmTime) {
		var ret time.Time
		return ret
	}
	return *o.ArmTime
}

// GetArmTimeOk returns a tuple with the ArmTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdInfo) GetArmTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ArmTime) {
		return nil, false
	}
	return o.ArmTime, true
}

// HasArmTime returns a boolean if a field has been set.
func (o *ThresholdInfo) HasArmTime() bool {
	if o != nil && !isNil(o.ArmTime) {
		return true
	}

	return false
}

// SetArmTime gets a reference to the given time.Time and assigns it to the ArmTime field.
func (o *ThresholdInfo) SetArmTime(v time.Time) {
	o.ArmTime = &v
}

func (o ThresholdInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThresholdInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["observedMeasurement"] = o.ObservedMeasurement
	toSerialize["observedValue"] = o.ObservedValue
	if !isNil(o.ThresholdLevel) {
		toSerialize["thresholdLevel"] = o.ThresholdLevel
	}
	if !isNil(o.ArmTime) {
		toSerialize["armTime"] = o.ArmTime
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



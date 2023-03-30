/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Location

import (
	"encoding/json"
)

// checks if the LocationQoS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationQoS{}

// LocationQoS QoS of Location request.
type LocationQoS struct {
	// Indicates value of accuracy.
	HAccuracy *float32 `json:"hAccuracy,omitempty"`
	// Indicates value of accuracy.
	VAccuracy *float32 `json:"vAccuracy,omitempty"`
	VerticalRequested *bool `json:"verticalRequested,omitempty"`
	ResponseTime *ResponseTime `json:"responseTime,omitempty"`
	MinorLocQoses []MinorLocationQoS `json:"minorLocQoses,omitempty"`
	LcsQosClass *LcsQosClass `json:"lcsQosClass,omitempty"`
}

// NewLocationQoS instantiates a new LocationQoS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationQoS() *LocationQoS {
	this := LocationQoS{}
	return &this
}

// NewLocationQoSWithDefaults instantiates a new LocationQoS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationQoSWithDefaults() *LocationQoS {
	this := LocationQoS{}
	return &this
}

// GetHAccuracy returns the HAccuracy field value if set, zero value otherwise.
func (o *LocationQoS) GetHAccuracy() float32 {
	if o == nil || IsNil(o.HAccuracy) {
		var ret float32
		return ret
	}
	return *o.HAccuracy
}

// GetHAccuracyOk returns a tuple with the HAccuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationQoS) GetHAccuracyOk() (*float32, bool) {
	if o == nil || IsNil(o.HAccuracy) {
		return nil, false
	}
	return o.HAccuracy, true
}

// HasHAccuracy returns a boolean if a field has been set.
func (o *LocationQoS) HasHAccuracy() bool {
	if o != nil && !IsNil(o.HAccuracy) {
		return true
	}

	return false
}

// SetHAccuracy gets a reference to the given float32 and assigns it to the HAccuracy field.
func (o *LocationQoS) SetHAccuracy(v float32) {
	o.HAccuracy = &v
}

// GetVAccuracy returns the VAccuracy field value if set, zero value otherwise.
func (o *LocationQoS) GetVAccuracy() float32 {
	if o == nil || IsNil(o.VAccuracy) {
		var ret float32
		return ret
	}
	return *o.VAccuracy
}

// GetVAccuracyOk returns a tuple with the VAccuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationQoS) GetVAccuracyOk() (*float32, bool) {
	if o == nil || IsNil(o.VAccuracy) {
		return nil, false
	}
	return o.VAccuracy, true
}

// HasVAccuracy returns a boolean if a field has been set.
func (o *LocationQoS) HasVAccuracy() bool {
	if o != nil && !IsNil(o.VAccuracy) {
		return true
	}

	return false
}

// SetVAccuracy gets a reference to the given float32 and assigns it to the VAccuracy field.
func (o *LocationQoS) SetVAccuracy(v float32) {
	o.VAccuracy = &v
}

// GetVerticalRequested returns the VerticalRequested field value if set, zero value otherwise.
func (o *LocationQoS) GetVerticalRequested() bool {
	if o == nil || IsNil(o.VerticalRequested) {
		var ret bool
		return ret
	}
	return *o.VerticalRequested
}

// GetVerticalRequestedOk returns a tuple with the VerticalRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationQoS) GetVerticalRequestedOk() (*bool, bool) {
	if o == nil || IsNil(o.VerticalRequested) {
		return nil, false
	}
	return o.VerticalRequested, true
}

// HasVerticalRequested returns a boolean if a field has been set.
func (o *LocationQoS) HasVerticalRequested() bool {
	if o != nil && !IsNil(o.VerticalRequested) {
		return true
	}

	return false
}

// SetVerticalRequested gets a reference to the given bool and assigns it to the VerticalRequested field.
func (o *LocationQoS) SetVerticalRequested(v bool) {
	o.VerticalRequested = &v
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise.
func (o *LocationQoS) GetResponseTime() ResponseTime {
	if o == nil || IsNil(o.ResponseTime) {
		var ret ResponseTime
		return ret
	}
	return *o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationQoS) GetResponseTimeOk() (*ResponseTime, bool) {
	if o == nil || IsNil(o.ResponseTime) {
		return nil, false
	}
	return o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *LocationQoS) HasResponseTime() bool {
	if o != nil && !IsNil(o.ResponseTime) {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given ResponseTime and assigns it to the ResponseTime field.
func (o *LocationQoS) SetResponseTime(v ResponseTime) {
	o.ResponseTime = &v
}

// GetMinorLocQoses returns the MinorLocQoses field value if set, zero value otherwise.
func (o *LocationQoS) GetMinorLocQoses() []MinorLocationQoS {
	if o == nil || IsNil(o.MinorLocQoses) {
		var ret []MinorLocationQoS
		return ret
	}
	return o.MinorLocQoses
}

// GetMinorLocQosesOk returns a tuple with the MinorLocQoses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationQoS) GetMinorLocQosesOk() ([]MinorLocationQoS, bool) {
	if o == nil || IsNil(o.MinorLocQoses) {
		return nil, false
	}
	return o.MinorLocQoses, true
}

// HasMinorLocQoses returns a boolean if a field has been set.
func (o *LocationQoS) HasMinorLocQoses() bool {
	if o != nil && !IsNil(o.MinorLocQoses) {
		return true
	}

	return false
}

// SetMinorLocQoses gets a reference to the given []MinorLocationQoS and assigns it to the MinorLocQoses field.
func (o *LocationQoS) SetMinorLocQoses(v []MinorLocationQoS) {
	o.MinorLocQoses = v
}

// GetLcsQosClass returns the LcsQosClass field value if set, zero value otherwise.
func (o *LocationQoS) GetLcsQosClass() LcsQosClass {
	if o == nil || IsNil(o.LcsQosClass) {
		var ret LcsQosClass
		return ret
	}
	return *o.LcsQosClass
}

// GetLcsQosClassOk returns a tuple with the LcsQosClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationQoS) GetLcsQosClassOk() (*LcsQosClass, bool) {
	if o == nil || IsNil(o.LcsQosClass) {
		return nil, false
	}
	return o.LcsQosClass, true
}

// HasLcsQosClass returns a boolean if a field has been set.
func (o *LocationQoS) HasLcsQosClass() bool {
	if o != nil && !IsNil(o.LcsQosClass) {
		return true
	}

	return false
}

// SetLcsQosClass gets a reference to the given LcsQosClass and assigns it to the LcsQosClass field.
func (o *LocationQoS) SetLcsQosClass(v LcsQosClass) {
	o.LcsQosClass = &v
}

func (o LocationQoS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationQoS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HAccuracy) {
		toSerialize["hAccuracy"] = o.HAccuracy
	}
	if !IsNil(o.VAccuracy) {
		toSerialize["vAccuracy"] = o.VAccuracy
	}
	if !IsNil(o.VerticalRequested) {
		toSerialize["verticalRequested"] = o.VerticalRequested
	}
	if !IsNil(o.ResponseTime) {
		toSerialize["responseTime"] = o.ResponseTime
	}
	if !IsNil(o.MinorLocQoses) {
		toSerialize["minorLocQoses"] = o.MinorLocQoses
	}
	if !IsNil(o.LcsQosClass) {
		toSerialize["lcsQosClass"] = o.LcsQosClass
	}
	return toSerialize, nil
}

type NullableLocationQoS struct {
	value *LocationQoS
	isSet bool
}

func (v NullableLocationQoS) Get() *LocationQoS {
	return v.value
}

func (v *NullableLocationQoS) Set(val *LocationQoS) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationQoS) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationQoS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationQoS(val *LocationQoS) *NullableLocationQoS {
	return &NullableLocationQoS{value: val, isSet: true}
}

func (v NullableLocationQoS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationQoS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


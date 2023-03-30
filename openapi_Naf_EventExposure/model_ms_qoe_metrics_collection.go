/*
Naf_EventExposure

AF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_EventExposure

import (
	"encoding/json"
)

// checks if the MsQoeMetricsCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MsQoeMetricsCollection{}

// MsQoeMetricsCollection Contains the Media Streaming QoE metrics information collected for an UE Application via AF. 
type MsQoeMetricsCollection struct {
	MsQoeMetrics []string `json:"msQoeMetrics"`
}

// NewMsQoeMetricsCollection instantiates a new MsQoeMetricsCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsQoeMetricsCollection(msQoeMetrics []string) *MsQoeMetricsCollection {
	this := MsQoeMetricsCollection{}
	this.MsQoeMetrics = msQoeMetrics
	return &this
}

// NewMsQoeMetricsCollectionWithDefaults instantiates a new MsQoeMetricsCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsQoeMetricsCollectionWithDefaults() *MsQoeMetricsCollection {
	this := MsQoeMetricsCollection{}
	return &this
}

// GetMsQoeMetrics returns the MsQoeMetrics field value
func (o *MsQoeMetricsCollection) GetMsQoeMetrics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MsQoeMetrics
}

// GetMsQoeMetricsOk returns a tuple with the MsQoeMetrics field value
// and a boolean to check if the value has been set.
func (o *MsQoeMetricsCollection) GetMsQoeMetricsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MsQoeMetrics, true
}

// SetMsQoeMetrics sets field value
func (o *MsQoeMetricsCollection) SetMsQoeMetrics(v []string) {
	o.MsQoeMetrics = v
}

func (o MsQoeMetricsCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MsQoeMetricsCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["msQoeMetrics"] = o.MsQoeMetrics
	return toSerialize, nil
}

type NullableMsQoeMetricsCollection struct {
	value *MsQoeMetricsCollection
	isSet bool
}

func (v NullableMsQoeMetricsCollection) Get() *MsQoeMetricsCollection {
	return v.value
}

func (v *NullableMsQoeMetricsCollection) Set(val *MsQoeMetricsCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableMsQoeMetricsCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableMsQoeMetricsCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsQoeMetricsCollection(val *MsQoeMetricsCollection) *NullableMsQoeMetricsCollection {
	return &NullableMsQoeMetricsCollection{value: val, isSet: true}
}

func (v NullableMsQoeMetricsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsQoeMetricsCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



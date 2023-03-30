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

// checks if the QFPacketDelayThresholdsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QFPacketDelayThresholdsType{}

// QFPacketDelayThresholdsType struct for QFPacketDelayThresholdsType
type QFPacketDelayThresholdsType struct {
	ThresholdDl *int32 `json:"thresholdDl,omitempty"`
	ThresholdUl *int32 `json:"thresholdUl,omitempty"`
	ThresholdRtt *int32 `json:"thresholdRtt,omitempty"`
}

// NewQFPacketDelayThresholdsType instantiates a new QFPacketDelayThresholdsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQFPacketDelayThresholdsType() *QFPacketDelayThresholdsType {
	this := QFPacketDelayThresholdsType{}
	return &this
}

// NewQFPacketDelayThresholdsTypeWithDefaults instantiates a new QFPacketDelayThresholdsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQFPacketDelayThresholdsTypeWithDefaults() *QFPacketDelayThresholdsType {
	this := QFPacketDelayThresholdsType{}
	return &this
}

// GetThresholdDl returns the ThresholdDl field value if set, zero value otherwise.
func (o *QFPacketDelayThresholdsType) GetThresholdDl() int32 {
	if o == nil || IsNil(o.ThresholdDl) {
		var ret int32
		return ret
	}
	return *o.ThresholdDl
}

// GetThresholdDlOk returns a tuple with the ThresholdDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFPacketDelayThresholdsType) GetThresholdDlOk() (*int32, bool) {
	if o == nil || IsNil(o.ThresholdDl) {
		return nil, false
	}
	return o.ThresholdDl, true
}

// HasThresholdDl returns a boolean if a field has been set.
func (o *QFPacketDelayThresholdsType) HasThresholdDl() bool {
	if o != nil && !IsNil(o.ThresholdDl) {
		return true
	}

	return false
}

// SetThresholdDl gets a reference to the given int32 and assigns it to the ThresholdDl field.
func (o *QFPacketDelayThresholdsType) SetThresholdDl(v int32) {
	o.ThresholdDl = &v
}

// GetThresholdUl returns the ThresholdUl field value if set, zero value otherwise.
func (o *QFPacketDelayThresholdsType) GetThresholdUl() int32 {
	if o == nil || IsNil(o.ThresholdUl) {
		var ret int32
		return ret
	}
	return *o.ThresholdUl
}

// GetThresholdUlOk returns a tuple with the ThresholdUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFPacketDelayThresholdsType) GetThresholdUlOk() (*int32, bool) {
	if o == nil || IsNil(o.ThresholdUl) {
		return nil, false
	}
	return o.ThresholdUl, true
}

// HasThresholdUl returns a boolean if a field has been set.
func (o *QFPacketDelayThresholdsType) HasThresholdUl() bool {
	if o != nil && !IsNil(o.ThresholdUl) {
		return true
	}

	return false
}

// SetThresholdUl gets a reference to the given int32 and assigns it to the ThresholdUl field.
func (o *QFPacketDelayThresholdsType) SetThresholdUl(v int32) {
	o.ThresholdUl = &v
}

// GetThresholdRtt returns the ThresholdRtt field value if set, zero value otherwise.
func (o *QFPacketDelayThresholdsType) GetThresholdRtt() int32 {
	if o == nil || IsNil(o.ThresholdRtt) {
		var ret int32
		return ret
	}
	return *o.ThresholdRtt
}

// GetThresholdRttOk returns a tuple with the ThresholdRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFPacketDelayThresholdsType) GetThresholdRttOk() (*int32, bool) {
	if o == nil || IsNil(o.ThresholdRtt) {
		return nil, false
	}
	return o.ThresholdRtt, true
}

// HasThresholdRtt returns a boolean if a field has been set.
func (o *QFPacketDelayThresholdsType) HasThresholdRtt() bool {
	if o != nil && !IsNil(o.ThresholdRtt) {
		return true
	}

	return false
}

// SetThresholdRtt gets a reference to the given int32 and assigns it to the ThresholdRtt field.
func (o *QFPacketDelayThresholdsType) SetThresholdRtt(v int32) {
	o.ThresholdRtt = &v
}

func (o QFPacketDelayThresholdsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QFPacketDelayThresholdsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThresholdDl) {
		toSerialize["thresholdDl"] = o.ThresholdDl
	}
	if !IsNil(o.ThresholdUl) {
		toSerialize["thresholdUl"] = o.ThresholdUl
	}
	if !IsNil(o.ThresholdRtt) {
		toSerialize["thresholdRtt"] = o.ThresholdRtt
	}
	return toSerialize, nil
}

type NullableQFPacketDelayThresholdsType struct {
	value *QFPacketDelayThresholdsType
	isSet bool
}

func (v NullableQFPacketDelayThresholdsType) Get() *QFPacketDelayThresholdsType {
	return v.value
}

func (v *NullableQFPacketDelayThresholdsType) Set(val *QFPacketDelayThresholdsType) {
	v.value = val
	v.isSet = true
}

func (v NullableQFPacketDelayThresholdsType) IsSet() bool {
	return v.isSet
}

func (v *NullableQFPacketDelayThresholdsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQFPacketDelayThresholdsType(val *QFPacketDelayThresholdsType) *NullableQFPacketDelayThresholdsType {
	return &NullableQFPacketDelayThresholdsType{value: val, isSet: true}
}

func (v NullableQFPacketDelayThresholdsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQFPacketDelayThresholdsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_StreamingDataMnS

import (
	"encoding/json"
)

// checks if the EventThresholdL1Type type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventThresholdL1Type{}

// EventThresholdL1Type See details in 3GPP TS 32.422 clause 5.10.X.
type EventThresholdL1Type struct {
	RSRP *int32 `json:"RSRP,omitempty"`
	RSRQ *int32 `json:"RSRQ,omitempty"`
}

// NewEventThresholdL1Type instantiates a new EventThresholdL1Type object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventThresholdL1Type() *EventThresholdL1Type {
	this := EventThresholdL1Type{}
	return &this
}

// NewEventThresholdL1TypeWithDefaults instantiates a new EventThresholdL1Type object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventThresholdL1TypeWithDefaults() *EventThresholdL1Type {
	this := EventThresholdL1Type{}
	return &this
}

// GetRSRP returns the RSRP field value if set, zero value otherwise.
func (o *EventThresholdL1Type) GetRSRP() int32 {
	if o == nil || isNil(o.RSRP) {
		var ret int32
		return ret
	}
	return *o.RSRP
}

// GetRSRPOk returns a tuple with the RSRP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventThresholdL1Type) GetRSRPOk() (*int32, bool) {
	if o == nil || isNil(o.RSRP) {
		return nil, false
	}
	return o.RSRP, true
}

// HasRSRP returns a boolean if a field has been set.
func (o *EventThresholdL1Type) HasRSRP() bool {
	if o != nil && !isNil(o.RSRP) {
		return true
	}

	return false
}

// SetRSRP gets a reference to the given int32 and assigns it to the RSRP field.
func (o *EventThresholdL1Type) SetRSRP(v int32) {
	o.RSRP = &v
}

// GetRSRQ returns the RSRQ field value if set, zero value otherwise.
func (o *EventThresholdL1Type) GetRSRQ() int32 {
	if o == nil || isNil(o.RSRQ) {
		var ret int32
		return ret
	}
	return *o.RSRQ
}

// GetRSRQOk returns a tuple with the RSRQ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventThresholdL1Type) GetRSRQOk() (*int32, bool) {
	if o == nil || isNil(o.RSRQ) {
		return nil, false
	}
	return o.RSRQ, true
}

// HasRSRQ returns a boolean if a field has been set.
func (o *EventThresholdL1Type) HasRSRQ() bool {
	if o != nil && !isNil(o.RSRQ) {
		return true
	}

	return false
}

// SetRSRQ gets a reference to the given int32 and assigns it to the RSRQ field.
func (o *EventThresholdL1Type) SetRSRQ(v int32) {
	o.RSRQ = &v
}

func (o EventThresholdL1Type) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventThresholdL1Type) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RSRP) {
		toSerialize["RSRP"] = o.RSRP
	}
	if !isNil(o.RSRQ) {
		toSerialize["RSRQ"] = o.RSRQ
	}
	return toSerialize, nil
}

type NullableEventThresholdL1Type struct {
	value *EventThresholdL1Type
	isSet bool
}

func (v NullableEventThresholdL1Type) Get() *EventThresholdL1Type {
	return v.value
}

func (v *NullableEventThresholdL1Type) Set(val *EventThresholdL1Type) {
	v.value = val
	v.isSet = true
}

func (v NullableEventThresholdL1Type) IsSet() bool {
	return v.isSet
}

func (v *NullableEventThresholdL1Type) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventThresholdL1Type(val *EventThresholdL1Type) *NullableEventThresholdL1Type {
	return &NullableEventThresholdL1Type{value: val, isSet: true}
}

func (v NullableEventThresholdL1Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventThresholdL1Type) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



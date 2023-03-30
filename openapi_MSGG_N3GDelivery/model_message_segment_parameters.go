/*
MSGG_N3GDelivery

API for MSGG N3G Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSGG_N3GDelivery

import (
	"encoding/json"
)

// checks if the MessageSegmentParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageSegmentParameters{}

// MessageSegmentParameters Contains the message segment parameters data
type MessageSegmentParameters struct {
	SegId *string `json:"segId,omitempty"`
	TotalSegCount *int32 `json:"totalSegCount,omitempty"`
	SegNumb *int32 `json:"segNumb,omitempty"`
	LastSegFlag *bool `json:"lastSegFlag,omitempty"`
}

// NewMessageSegmentParameters instantiates a new MessageSegmentParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageSegmentParameters() *MessageSegmentParameters {
	this := MessageSegmentParameters{}
	return &this
}

// NewMessageSegmentParametersWithDefaults instantiates a new MessageSegmentParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageSegmentParametersWithDefaults() *MessageSegmentParameters {
	this := MessageSegmentParameters{}
	return &this
}

// GetSegId returns the SegId field value if set, zero value otherwise.
func (o *MessageSegmentParameters) GetSegId() string {
	if o == nil || IsNil(o.SegId) {
		var ret string
		return ret
	}
	return *o.SegId
}

// GetSegIdOk returns a tuple with the SegId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSegmentParameters) GetSegIdOk() (*string, bool) {
	if o == nil || IsNil(o.SegId) {
		return nil, false
	}
	return o.SegId, true
}

// HasSegId returns a boolean if a field has been set.
func (o *MessageSegmentParameters) HasSegId() bool {
	if o != nil && !IsNil(o.SegId) {
		return true
	}

	return false
}

// SetSegId gets a reference to the given string and assigns it to the SegId field.
func (o *MessageSegmentParameters) SetSegId(v string) {
	o.SegId = &v
}

// GetTotalSegCount returns the TotalSegCount field value if set, zero value otherwise.
func (o *MessageSegmentParameters) GetTotalSegCount() int32 {
	if o == nil || IsNil(o.TotalSegCount) {
		var ret int32
		return ret
	}
	return *o.TotalSegCount
}

// GetTotalSegCountOk returns a tuple with the TotalSegCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSegmentParameters) GetTotalSegCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSegCount) {
		return nil, false
	}
	return o.TotalSegCount, true
}

// HasTotalSegCount returns a boolean if a field has been set.
func (o *MessageSegmentParameters) HasTotalSegCount() bool {
	if o != nil && !IsNil(o.TotalSegCount) {
		return true
	}

	return false
}

// SetTotalSegCount gets a reference to the given int32 and assigns it to the TotalSegCount field.
func (o *MessageSegmentParameters) SetTotalSegCount(v int32) {
	o.TotalSegCount = &v
}

// GetSegNumb returns the SegNumb field value if set, zero value otherwise.
func (o *MessageSegmentParameters) GetSegNumb() int32 {
	if o == nil || IsNil(o.SegNumb) {
		var ret int32
		return ret
	}
	return *o.SegNumb
}

// GetSegNumbOk returns a tuple with the SegNumb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSegmentParameters) GetSegNumbOk() (*int32, bool) {
	if o == nil || IsNil(o.SegNumb) {
		return nil, false
	}
	return o.SegNumb, true
}

// HasSegNumb returns a boolean if a field has been set.
func (o *MessageSegmentParameters) HasSegNumb() bool {
	if o != nil && !IsNil(o.SegNumb) {
		return true
	}

	return false
}

// SetSegNumb gets a reference to the given int32 and assigns it to the SegNumb field.
func (o *MessageSegmentParameters) SetSegNumb(v int32) {
	o.SegNumb = &v
}

// GetLastSegFlag returns the LastSegFlag field value if set, zero value otherwise.
func (o *MessageSegmentParameters) GetLastSegFlag() bool {
	if o == nil || IsNil(o.LastSegFlag) {
		var ret bool
		return ret
	}
	return *o.LastSegFlag
}

// GetLastSegFlagOk returns a tuple with the LastSegFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSegmentParameters) GetLastSegFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.LastSegFlag) {
		return nil, false
	}
	return o.LastSegFlag, true
}

// HasLastSegFlag returns a boolean if a field has been set.
func (o *MessageSegmentParameters) HasLastSegFlag() bool {
	if o != nil && !IsNil(o.LastSegFlag) {
		return true
	}

	return false
}

// SetLastSegFlag gets a reference to the given bool and assigns it to the LastSegFlag field.
func (o *MessageSegmentParameters) SetLastSegFlag(v bool) {
	o.LastSegFlag = &v
}

func (o MessageSegmentParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageSegmentParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SegId) {
		toSerialize["segId"] = o.SegId
	}
	if !IsNil(o.TotalSegCount) {
		toSerialize["totalSegCount"] = o.TotalSegCount
	}
	if !IsNil(o.SegNumb) {
		toSerialize["segNumb"] = o.SegNumb
	}
	if !IsNil(o.LastSegFlag) {
		toSerialize["lastSegFlag"] = o.LastSegFlag
	}
	return toSerialize, nil
}

type NullableMessageSegmentParameters struct {
	value *MessageSegmentParameters
	isSet bool
}

func (v NullableMessageSegmentParameters) Get() *MessageSegmentParameters {
	return v.value
}

func (v *NullableMessageSegmentParameters) Set(val *MessageSegmentParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSegmentParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSegmentParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSegmentParameters(val *MessageSegmentParameters) *NullableMessageSegmentParameters {
	return &NullableMessageSegmentParameters{value: val, isSet: true}
}

func (v NullableMessageSegmentParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSegmentParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



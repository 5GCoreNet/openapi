/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
)

// checks if the TraceTargetType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceTargetType{}

// TraceTargetType Trace target conveying both the type and value of the target ID. For additional details see 3GPP TS 32.422
type TraceTargetType struct {
	TargetIdType string `json:"TargetIdType"`
	TargetIdValue string `json:"TargetIdValue"`
}

// NewTraceTargetType instantiates a new TraceTargetType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceTargetType(targetIdType string, targetIdValue string) *TraceTargetType {
	this := TraceTargetType{}
	this.TargetIdType = targetIdType
	this.TargetIdValue = targetIdValue
	return &this
}

// NewTraceTargetTypeWithDefaults instantiates a new TraceTargetType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceTargetTypeWithDefaults() *TraceTargetType {
	this := TraceTargetType{}
	return &this
}

// GetTargetIdType returns the TargetIdType field value
func (o *TraceTargetType) GetTargetIdType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetIdType
}

// GetTargetIdTypeOk returns a tuple with the TargetIdType field value
// and a boolean to check if the value has been set.
func (o *TraceTargetType) GetTargetIdTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetIdType, true
}

// SetTargetIdType sets field value
func (o *TraceTargetType) SetTargetIdType(v string) {
	o.TargetIdType = v
}

// GetTargetIdValue returns the TargetIdValue field value
func (o *TraceTargetType) GetTargetIdValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetIdValue
}

// GetTargetIdValueOk returns a tuple with the TargetIdValue field value
// and a boolean to check if the value has been set.
func (o *TraceTargetType) GetTargetIdValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetIdValue, true
}

// SetTargetIdValue sets field value
func (o *TraceTargetType) SetTargetIdValue(v string) {
	o.TargetIdValue = v
}

func (o TraceTargetType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceTargetType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["TargetIdType"] = o.TargetIdType
	toSerialize["TargetIdValue"] = o.TargetIdValue
	return toSerialize, nil
}

type NullableTraceTargetType struct {
	value *TraceTargetType
	isSet bool
}

func (v NullableTraceTargetType) Get() *TraceTargetType {
	return v.value
}

func (v *NullableTraceTargetType) Set(val *TraceTargetType) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceTargetType) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceTargetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceTargetType(val *TraceTargetType) *NullableTraceTargetType {
	return &NullableTraceTargetType{value: val, isSet: true}
}

func (v NullableTraceTargetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceTargetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



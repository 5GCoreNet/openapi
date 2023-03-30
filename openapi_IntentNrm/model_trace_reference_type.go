/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
)

// checks if the TraceReferenceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceReferenceType{}

// TraceReferenceType The Trace Reference parameter shall be globally unique, therefore the Trace Reference shall compose as follows - MCC+MNC+Trace ID, where the MCC and MNC are coming with the Trace activation request from the management system to identify one PLMN containing the management system, and Trace ID is a 3 byte Octet String. See 3GPP TS 32.422 clause 5.6 for additional details.
type TraceReferenceType struct {
	Mcc string `json:"mcc"`
	Mnc string `json:"mnc"`
	TraceId string `json:"traceId"`
}

// NewTraceReferenceType instantiates a new TraceReferenceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceReferenceType(mcc string, mnc string, traceId string) *TraceReferenceType {
	this := TraceReferenceType{}
	this.Mcc = mcc
	this.Mnc = mnc
	this.TraceId = traceId
	return &this
}

// NewTraceReferenceTypeWithDefaults instantiates a new TraceReferenceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceReferenceTypeWithDefaults() *TraceReferenceType {
	this := TraceReferenceType{}
	return &this
}

// GetMcc returns the Mcc field value
func (o *TraceReferenceType) GetMcc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value
// and a boolean to check if the value has been set.
func (o *TraceReferenceType) GetMccOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mcc, true
}

// SetMcc sets field value
func (o *TraceReferenceType) SetMcc(v string) {
	o.Mcc = v
}

// GetMnc returns the Mnc field value
func (o *TraceReferenceType) GetMnc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value
// and a boolean to check if the value has been set.
func (o *TraceReferenceType) GetMncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mnc, true
}

// SetMnc sets field value
func (o *TraceReferenceType) SetMnc(v string) {
	o.Mnc = v
}

// GetTraceId returns the TraceId field value
func (o *TraceReferenceType) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *TraceReferenceType) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value
func (o *TraceReferenceType) SetTraceId(v string) {
	o.TraceId = v
}

func (o TraceReferenceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceReferenceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mcc"] = o.Mcc
	toSerialize["mnc"] = o.Mnc
	toSerialize["traceId"] = o.TraceId
	return toSerialize, nil
}

type NullableTraceReferenceType struct {
	value *TraceReferenceType
	isSet bool
}

func (v NullableTraceReferenceType) Get() *TraceReferenceType {
	return v.value
}

func (v *NullableTraceReferenceType) Set(val *TraceReferenceType) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceReferenceType) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceReferenceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceReferenceType(val *TraceReferenceType) *NullableTraceReferenceType {
	return &NullableTraceReferenceType{value: val, isSet: true}
}

func (v NullableTraceReferenceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceReferenceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



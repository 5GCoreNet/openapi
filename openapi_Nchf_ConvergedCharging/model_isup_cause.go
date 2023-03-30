/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the ISUPCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ISUPCause{}

// ISUPCause struct for ISUPCause
type ISUPCause struct {
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	ISUPCauseLocation *int32 `json:"iSUPCauseLocation,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	ISUPCauseValue *int32 `json:"iSUPCauseValue,omitempty"`
	ISUPCauseDiagnostics *string `json:"iSUPCauseDiagnostics,omitempty"`
}

// NewISUPCause instantiates a new ISUPCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewISUPCause() *ISUPCause {
	this := ISUPCause{}
	return &this
}

// NewISUPCauseWithDefaults instantiates a new ISUPCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewISUPCauseWithDefaults() *ISUPCause {
	this := ISUPCause{}
	return &this
}

// GetISUPCauseLocation returns the ISUPCauseLocation field value if set, zero value otherwise.
func (o *ISUPCause) GetISUPCauseLocation() int32 {
	if o == nil || IsNil(o.ISUPCauseLocation) {
		var ret int32
		return ret
	}
	return *o.ISUPCauseLocation
}

// GetISUPCauseLocationOk returns a tuple with the ISUPCauseLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISUPCause) GetISUPCauseLocationOk() (*int32, bool) {
	if o == nil || IsNil(o.ISUPCauseLocation) {
		return nil, false
	}
	return o.ISUPCauseLocation, true
}

// HasISUPCauseLocation returns a boolean if a field has been set.
func (o *ISUPCause) HasISUPCauseLocation() bool {
	if o != nil && !IsNil(o.ISUPCauseLocation) {
		return true
	}

	return false
}

// SetISUPCauseLocation gets a reference to the given int32 and assigns it to the ISUPCauseLocation field.
func (o *ISUPCause) SetISUPCauseLocation(v int32) {
	o.ISUPCauseLocation = &v
}

// GetISUPCauseValue returns the ISUPCauseValue field value if set, zero value otherwise.
func (o *ISUPCause) GetISUPCauseValue() int32 {
	if o == nil || IsNil(o.ISUPCauseValue) {
		var ret int32
		return ret
	}
	return *o.ISUPCauseValue
}

// GetISUPCauseValueOk returns a tuple with the ISUPCauseValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISUPCause) GetISUPCauseValueOk() (*int32, bool) {
	if o == nil || IsNil(o.ISUPCauseValue) {
		return nil, false
	}
	return o.ISUPCauseValue, true
}

// HasISUPCauseValue returns a boolean if a field has been set.
func (o *ISUPCause) HasISUPCauseValue() bool {
	if o != nil && !IsNil(o.ISUPCauseValue) {
		return true
	}

	return false
}

// SetISUPCauseValue gets a reference to the given int32 and assigns it to the ISUPCauseValue field.
func (o *ISUPCause) SetISUPCauseValue(v int32) {
	o.ISUPCauseValue = &v
}

// GetISUPCauseDiagnostics returns the ISUPCauseDiagnostics field value if set, zero value otherwise.
func (o *ISUPCause) GetISUPCauseDiagnostics() string {
	if o == nil || IsNil(o.ISUPCauseDiagnostics) {
		var ret string
		return ret
	}
	return *o.ISUPCauseDiagnostics
}

// GetISUPCauseDiagnosticsOk returns a tuple with the ISUPCauseDiagnostics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISUPCause) GetISUPCauseDiagnosticsOk() (*string, bool) {
	if o == nil || IsNil(o.ISUPCauseDiagnostics) {
		return nil, false
	}
	return o.ISUPCauseDiagnostics, true
}

// HasISUPCauseDiagnostics returns a boolean if a field has been set.
func (o *ISUPCause) HasISUPCauseDiagnostics() bool {
	if o != nil && !IsNil(o.ISUPCauseDiagnostics) {
		return true
	}

	return false
}

// SetISUPCauseDiagnostics gets a reference to the given string and assigns it to the ISUPCauseDiagnostics field.
func (o *ISUPCause) SetISUPCauseDiagnostics(v string) {
	o.ISUPCauseDiagnostics = &v
}

func (o ISUPCause) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ISUPCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ISUPCauseLocation) {
		toSerialize["iSUPCauseLocation"] = o.ISUPCauseLocation
	}
	if !IsNil(o.ISUPCauseValue) {
		toSerialize["iSUPCauseValue"] = o.ISUPCauseValue
	}
	if !IsNil(o.ISUPCauseDiagnostics) {
		toSerialize["iSUPCauseDiagnostics"] = o.ISUPCauseDiagnostics
	}
	return toSerialize, nil
}

type NullableISUPCause struct {
	value *ISUPCause
	isSet bool
}

func (v NullableISUPCause) Get() *ISUPCause {
	return v.value
}

func (v *NullableISUPCause) Set(val *ISUPCause) {
	v.value = val
	v.isSet = true
}

func (v NullableISUPCause) IsSet() bool {
	return v.isSet
}

func (v *NullableISUPCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableISUPCause(val *ISUPCause) *NullableISUPCause {
	return &NullableISUPCause{value: val, isSet: true}
}

func (v NullableISUPCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableISUPCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



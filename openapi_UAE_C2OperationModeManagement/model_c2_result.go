/*
UAE Server C2 Operation Mode Management Service

UAE Server C2 Operation Mode Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_UAE_C2OperationModeManagement

import (
	"encoding/json"
)

// checks if the C2Result type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C2Result{}

// C2Result Represents the result of an action related to C2 of a UAS.
type C2Result struct {
	C2OpConfirmed bool `json:"c2OpConfirmed"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewC2Result instantiates a new C2Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC2Result(c2OpConfirmed bool) *C2Result {
	this := C2Result{}
	this.C2OpConfirmed = c2OpConfirmed
	return &this
}

// NewC2ResultWithDefaults instantiates a new C2Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC2ResultWithDefaults() *C2Result {
	this := C2Result{}
	return &this
}

// GetC2OpConfirmed returns the C2OpConfirmed field value
func (o *C2Result) GetC2OpConfirmed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.C2OpConfirmed
}

// GetC2OpConfirmedOk returns a tuple with the C2OpConfirmed field value
// and a boolean to check if the value has been set.
func (o *C2Result) GetC2OpConfirmedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.C2OpConfirmed, true
}

// SetC2OpConfirmed sets field value
func (o *C2Result) SetC2OpConfirmed(v bool) {
	o.C2OpConfirmed = v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *C2Result) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C2Result) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *C2Result) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *C2Result) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o C2Result) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C2Result) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["c2OpConfirmed"] = o.C2OpConfirmed
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableC2Result struct {
	value *C2Result
	isSet bool
}

func (v NullableC2Result) Get() *C2Result {
	return v.value
}

func (v *NullableC2Result) Set(val *C2Result) {
	v.value = val
	v.isSet = true
}

func (v NullableC2Result) IsSet() bool {
	return v.isSet
}

func (v *NullableC2Result) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC2Result(val *C2Result) *NullableC2Result {
	return &NullableC2Result{value: val, isSet: true}
}

func (v NullableC2Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC2Result) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



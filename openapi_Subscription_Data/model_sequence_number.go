/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// checks if the SequenceNumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SequenceNumber{}

// SequenceNumber Contains the SQN.
type SequenceNumber struct {
	SqnScheme *SqnScheme `json:"sqnScheme,omitempty"`
	Sqn *string `json:"sqn,omitempty"`
	LastIndexes *map[string]int32 `json:"lastIndexes,omitempty"`
	IndLength *int32 `json:"indLength,omitempty"`
	DifSign *Sign `json:"difSign,omitempty"`
}

// NewSequenceNumber instantiates a new SequenceNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSequenceNumber() *SequenceNumber {
	this := SequenceNumber{}
	return &this
}

// NewSequenceNumberWithDefaults instantiates a new SequenceNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSequenceNumberWithDefaults() *SequenceNumber {
	this := SequenceNumber{}
	return &this
}

// GetSqnScheme returns the SqnScheme field value if set, zero value otherwise.
func (o *SequenceNumber) GetSqnScheme() SqnScheme {
	if o == nil || isNil(o.SqnScheme) {
		var ret SqnScheme
		return ret
	}
	return *o.SqnScheme
}

// GetSqnSchemeOk returns a tuple with the SqnScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SequenceNumber) GetSqnSchemeOk() (*SqnScheme, bool) {
	if o == nil || isNil(o.SqnScheme) {
		return nil, false
	}
	return o.SqnScheme, true
}

// HasSqnScheme returns a boolean if a field has been set.
func (o *SequenceNumber) HasSqnScheme() bool {
	if o != nil && !isNil(o.SqnScheme) {
		return true
	}

	return false
}

// SetSqnScheme gets a reference to the given SqnScheme and assigns it to the SqnScheme field.
func (o *SequenceNumber) SetSqnScheme(v SqnScheme) {
	o.SqnScheme = &v
}

// GetSqn returns the Sqn field value if set, zero value otherwise.
func (o *SequenceNumber) GetSqn() string {
	if o == nil || isNil(o.Sqn) {
		var ret string
		return ret
	}
	return *o.Sqn
}

// GetSqnOk returns a tuple with the Sqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SequenceNumber) GetSqnOk() (*string, bool) {
	if o == nil || isNil(o.Sqn) {
		return nil, false
	}
	return o.Sqn, true
}

// HasSqn returns a boolean if a field has been set.
func (o *SequenceNumber) HasSqn() bool {
	if o != nil && !isNil(o.Sqn) {
		return true
	}

	return false
}

// SetSqn gets a reference to the given string and assigns it to the Sqn field.
func (o *SequenceNumber) SetSqn(v string) {
	o.Sqn = &v
}

// GetLastIndexes returns the LastIndexes field value if set, zero value otherwise.
func (o *SequenceNumber) GetLastIndexes() map[string]int32 {
	if o == nil || isNil(o.LastIndexes) {
		var ret map[string]int32
		return ret
	}
	return *o.LastIndexes
}

// GetLastIndexesOk returns a tuple with the LastIndexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SequenceNumber) GetLastIndexesOk() (*map[string]int32, bool) {
	if o == nil || isNil(o.LastIndexes) {
		return nil, false
	}
	return o.LastIndexes, true
}

// HasLastIndexes returns a boolean if a field has been set.
func (o *SequenceNumber) HasLastIndexes() bool {
	if o != nil && !isNil(o.LastIndexes) {
		return true
	}

	return false
}

// SetLastIndexes gets a reference to the given map[string]int32 and assigns it to the LastIndexes field.
func (o *SequenceNumber) SetLastIndexes(v map[string]int32) {
	o.LastIndexes = &v
}

// GetIndLength returns the IndLength field value if set, zero value otherwise.
func (o *SequenceNumber) GetIndLength() int32 {
	if o == nil || isNil(o.IndLength) {
		var ret int32
		return ret
	}
	return *o.IndLength
}

// GetIndLengthOk returns a tuple with the IndLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SequenceNumber) GetIndLengthOk() (*int32, bool) {
	if o == nil || isNil(o.IndLength) {
		return nil, false
	}
	return o.IndLength, true
}

// HasIndLength returns a boolean if a field has been set.
func (o *SequenceNumber) HasIndLength() bool {
	if o != nil && !isNil(o.IndLength) {
		return true
	}

	return false
}

// SetIndLength gets a reference to the given int32 and assigns it to the IndLength field.
func (o *SequenceNumber) SetIndLength(v int32) {
	o.IndLength = &v
}

// GetDifSign returns the DifSign field value if set, zero value otherwise.
func (o *SequenceNumber) GetDifSign() Sign {
	if o == nil || isNil(o.DifSign) {
		var ret Sign
		return ret
	}
	return *o.DifSign
}

// GetDifSignOk returns a tuple with the DifSign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SequenceNumber) GetDifSignOk() (*Sign, bool) {
	if o == nil || isNil(o.DifSign) {
		return nil, false
	}
	return o.DifSign, true
}

// HasDifSign returns a boolean if a field has been set.
func (o *SequenceNumber) HasDifSign() bool {
	if o != nil && !isNil(o.DifSign) {
		return true
	}

	return false
}

// SetDifSign gets a reference to the given Sign and assigns it to the DifSign field.
func (o *SequenceNumber) SetDifSign(v Sign) {
	o.DifSign = &v
}

func (o SequenceNumber) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SequenceNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SqnScheme) {
		toSerialize["sqnScheme"] = o.SqnScheme
	}
	if !isNil(o.Sqn) {
		toSerialize["sqn"] = o.Sqn
	}
	if !isNil(o.LastIndexes) {
		toSerialize["lastIndexes"] = o.LastIndexes
	}
	if !isNil(o.IndLength) {
		toSerialize["indLength"] = o.IndLength
	}
	if !isNil(o.DifSign) {
		toSerialize["difSign"] = o.DifSign
	}
	return toSerialize, nil
}

type NullableSequenceNumber struct {
	value *SequenceNumber
	isSet bool
}

func (v NullableSequenceNumber) Get() *SequenceNumber {
	return v.value
}

func (v *NullableSequenceNumber) Set(val *SequenceNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableSequenceNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableSequenceNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSequenceNumber(val *SequenceNumber) *NullableSequenceNumber {
	return &NullableSequenceNumber{value: val, isSet: true}
}

func (v NullableSequenceNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSequenceNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



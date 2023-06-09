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

// checks if the ExternalGnbDuFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalGnbDuFunctionSingleAllOfAttributesAllOf{}

// ExternalGnbDuFunctionSingleAllOfAttributesAllOf struct for ExternalGnbDuFunctionSingleAllOfAttributesAllOf
type ExternalGnbDuFunctionSingleAllOfAttributesAllOf struct {
	GnbId       *string `json:"gnbId,omitempty"`
	GnbIdLength *int32  `json:"gnbIdLength,omitempty"`
}

// NewExternalGnbDuFunctionSingleAllOfAttributesAllOf instantiates a new ExternalGnbDuFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGnbDuFunctionSingleAllOfAttributesAllOf() *ExternalGnbDuFunctionSingleAllOfAttributesAllOf {
	this := ExternalGnbDuFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewExternalGnbDuFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new ExternalGnbDuFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGnbDuFunctionSingleAllOfAttributesAllOfWithDefaults() *ExternalGnbDuFunctionSingleAllOfAttributesAllOf {
	this := ExternalGnbDuFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetGnbId returns the GnbId field value if set, zero value otherwise.
func (o *ExternalGnbDuFunctionSingleAllOfAttributesAllOf) GetGnbId() string {
	if o == nil || IsNil(o.GnbId) {
		var ret string
		return ret
	}
	return *o.GnbId
}

// GetGnbIdOk returns a tuple with the GnbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbDuFunctionSingleAllOfAttributesAllOf) GetGnbIdOk() (*string, bool) {
	if o == nil || IsNil(o.GnbId) {
		return nil, false
	}
	return o.GnbId, true
}

// HasGnbId returns a boolean if a field has been set.
func (o *ExternalGnbDuFunctionSingleAllOfAttributesAllOf) HasGnbId() bool {
	if o != nil && !IsNil(o.GnbId) {
		return true
	}

	return false
}

// SetGnbId gets a reference to the given string and assigns it to the GnbId field.
func (o *ExternalGnbDuFunctionSingleAllOfAttributesAllOf) SetGnbId(v string) {
	o.GnbId = &v
}

// GetGnbIdLength returns the GnbIdLength field value if set, zero value otherwise.
func (o *ExternalGnbDuFunctionSingleAllOfAttributesAllOf) GetGnbIdLength() int32 {
	if o == nil || IsNil(o.GnbIdLength) {
		var ret int32
		return ret
	}
	return *o.GnbIdLength
}

// GetGnbIdLengthOk returns a tuple with the GnbIdLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbDuFunctionSingleAllOfAttributesAllOf) GetGnbIdLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.GnbIdLength) {
		return nil, false
	}
	return o.GnbIdLength, true
}

// HasGnbIdLength returns a boolean if a field has been set.
func (o *ExternalGnbDuFunctionSingleAllOfAttributesAllOf) HasGnbIdLength() bool {
	if o != nil && !IsNil(o.GnbIdLength) {
		return true
	}

	return false
}

// SetGnbIdLength gets a reference to the given int32 and assigns it to the GnbIdLength field.
func (o *ExternalGnbDuFunctionSingleAllOfAttributesAllOf) SetGnbIdLength(v int32) {
	o.GnbIdLength = &v
}

func (o ExternalGnbDuFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalGnbDuFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GnbId) {
		toSerialize["gnbId"] = o.GnbId
	}
	if !IsNil(o.GnbIdLength) {
		toSerialize["gnbIdLength"] = o.GnbIdLength
	}
	return toSerialize, nil
}

type NullableExternalGnbDuFunctionSingleAllOfAttributesAllOf struct {
	value *ExternalGnbDuFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableExternalGnbDuFunctionSingleAllOfAttributesAllOf) Get() *ExternalGnbDuFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableExternalGnbDuFunctionSingleAllOfAttributesAllOf) Set(val *ExternalGnbDuFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGnbDuFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGnbDuFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGnbDuFunctionSingleAllOfAttributesAllOf(val *ExternalGnbDuFunctionSingleAllOfAttributesAllOf) *NullableExternalGnbDuFunctionSingleAllOfAttributesAllOf {
	return &NullableExternalGnbDuFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableExternalGnbDuFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGnbDuFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

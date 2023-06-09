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

// checks if the ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf{}

// ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf struct for ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf
type ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf struct {
	GnbId       *string `json:"gnbId,omitempty"`
	GnbIdLength *int32  `json:"gnbIdLength,omitempty"`
	PlmnId      *PlmnId `json:"plmnId,omitempty"`
}

// NewExternalGnbCuCpFunctionSingleAllOfAttributesAllOf instantiates a new ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGnbCuCpFunctionSingleAllOfAttributesAllOf() *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf {
	this := ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewExternalGnbCuCpFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGnbCuCpFunctionSingleAllOfAttributesAllOfWithDefaults() *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf {
	this := ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetGnbId returns the GnbId field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) GetGnbId() string {
	if o == nil || IsNil(o.GnbId) {
		var ret string
		return ret
	}
	return *o.GnbId
}

// GetGnbIdOk returns a tuple with the GnbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) GetGnbIdOk() (*string, bool) {
	if o == nil || IsNil(o.GnbId) {
		return nil, false
	}
	return o.GnbId, true
}

// HasGnbId returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) HasGnbId() bool {
	if o != nil && !IsNil(o.GnbId) {
		return true
	}

	return false
}

// SetGnbId gets a reference to the given string and assigns it to the GnbId field.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) SetGnbId(v string) {
	o.GnbId = &v
}

// GetGnbIdLength returns the GnbIdLength field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) GetGnbIdLength() int32 {
	if o == nil || IsNil(o.GnbIdLength) {
		var ret int32
		return ret
	}
	return *o.GnbIdLength
}

// GetGnbIdLengthOk returns a tuple with the GnbIdLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) GetGnbIdLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.GnbIdLength) {
		return nil, false
	}
	return o.GnbIdLength, true
}

// HasGnbIdLength returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) HasGnbIdLength() bool {
	if o != nil && !IsNil(o.GnbIdLength) {
		return true
	}

	return false
}

// SetGnbIdLength gets a reference to the given int32 and assigns it to the GnbIdLength field.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) SetGnbIdLength(v int32) {
	o.GnbIdLength = &v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) GetPlmnId() PlmnId {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

func (o ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GnbId) {
		toSerialize["gnbId"] = o.GnbId
	}
	if !IsNil(o.GnbIdLength) {
		toSerialize["gnbIdLength"] = o.GnbIdLength
	}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	return toSerialize, nil
}

type NullableExternalGnbCuCpFunctionSingleAllOfAttributesAllOf struct {
	value *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) Get() *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) Set(val *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGnbCuCpFunctionSingleAllOfAttributesAllOf(val *ExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) *NullableExternalGnbCuCpFunctionSingleAllOfAttributesAllOf {
	return &NullableExternalGnbCuCpFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGnbCuCpFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

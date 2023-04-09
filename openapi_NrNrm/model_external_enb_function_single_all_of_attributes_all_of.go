/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the ExternalENBFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalENBFunctionSingleAllOfAttributesAllOf{}

// ExternalENBFunctionSingleAllOfAttributesAllOf struct for ExternalENBFunctionSingleAllOfAttributesAllOf
type ExternalENBFunctionSingleAllOfAttributesAllOf struct {
	ENBId *int32 `json:"eNBId,omitempty"`
}

// NewExternalENBFunctionSingleAllOfAttributesAllOf instantiates a new ExternalENBFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalENBFunctionSingleAllOfAttributesAllOf() *ExternalENBFunctionSingleAllOfAttributesAllOf {
	this := ExternalENBFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewExternalENBFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new ExternalENBFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalENBFunctionSingleAllOfAttributesAllOfWithDefaults() *ExternalENBFunctionSingleAllOfAttributesAllOf {
	this := ExternalENBFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetENBId returns the ENBId field value if set, zero value otherwise.
func (o *ExternalENBFunctionSingleAllOfAttributesAllOf) GetENBId() int32 {
	if o == nil || IsNil(o.ENBId) {
		var ret int32
		return ret
	}
	return *o.ENBId
}

// GetENBIdOk returns a tuple with the ENBId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalENBFunctionSingleAllOfAttributesAllOf) GetENBIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ENBId) {
		return nil, false
	}
	return o.ENBId, true
}

// HasENBId returns a boolean if a field has been set.
func (o *ExternalENBFunctionSingleAllOfAttributesAllOf) HasENBId() bool {
	if o != nil && !IsNil(o.ENBId) {
		return true
	}

	return false
}

// SetENBId gets a reference to the given int32 and assigns it to the ENBId field.
func (o *ExternalENBFunctionSingleAllOfAttributesAllOf) SetENBId(v int32) {
	o.ENBId = &v
}

func (o ExternalENBFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalENBFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ENBId) {
		toSerialize["eNBId"] = o.ENBId
	}
	return toSerialize, nil
}

type NullableExternalENBFunctionSingleAllOfAttributesAllOf struct {
	value *ExternalENBFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableExternalENBFunctionSingleAllOfAttributesAllOf) Get() *ExternalENBFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableExternalENBFunctionSingleAllOfAttributesAllOf) Set(val *ExternalENBFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalENBFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalENBFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalENBFunctionSingleAllOfAttributesAllOf(val *ExternalENBFunctionSingleAllOfAttributesAllOf) *NullableExternalENBFunctionSingleAllOfAttributesAllOf {
	return &NullableExternalENBFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableExternalENBFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalENBFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

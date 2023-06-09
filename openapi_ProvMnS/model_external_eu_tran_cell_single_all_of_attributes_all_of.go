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

// checks if the ExternalEUTranCellSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalEUTranCellSingleAllOfAttributesAllOf{}

// ExternalEUTranCellSingleAllOfAttributesAllOf struct for ExternalEUTranCellSingleAllOfAttributesAllOf
type ExternalEUTranCellSingleAllOfAttributesAllOf struct {
	EUtranFrequencyRef *string `json:"EUtranFrequencyRef,omitempty"`
}

// NewExternalEUTranCellSingleAllOfAttributesAllOf instantiates a new ExternalEUTranCellSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalEUTranCellSingleAllOfAttributesAllOf() *ExternalEUTranCellSingleAllOfAttributesAllOf {
	this := ExternalEUTranCellSingleAllOfAttributesAllOf{}
	return &this
}

// NewExternalEUTranCellSingleAllOfAttributesAllOfWithDefaults instantiates a new ExternalEUTranCellSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalEUTranCellSingleAllOfAttributesAllOfWithDefaults() *ExternalEUTranCellSingleAllOfAttributesAllOf {
	this := ExternalEUTranCellSingleAllOfAttributesAllOf{}
	return &this
}

// GetEUtranFrequencyRef returns the EUtranFrequencyRef field value if set, zero value otherwise.
func (o *ExternalEUTranCellSingleAllOfAttributesAllOf) GetEUtranFrequencyRef() string {
	if o == nil || IsNil(o.EUtranFrequencyRef) {
		var ret string
		return ret
	}
	return *o.EUtranFrequencyRef
}

// GetEUtranFrequencyRefOk returns a tuple with the EUtranFrequencyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEUTranCellSingleAllOfAttributesAllOf) GetEUtranFrequencyRefOk() (*string, bool) {
	if o == nil || IsNil(o.EUtranFrequencyRef) {
		return nil, false
	}
	return o.EUtranFrequencyRef, true
}

// HasEUtranFrequencyRef returns a boolean if a field has been set.
func (o *ExternalEUTranCellSingleAllOfAttributesAllOf) HasEUtranFrequencyRef() bool {
	if o != nil && !IsNil(o.EUtranFrequencyRef) {
		return true
	}

	return false
}

// SetEUtranFrequencyRef gets a reference to the given string and assigns it to the EUtranFrequencyRef field.
func (o *ExternalEUTranCellSingleAllOfAttributesAllOf) SetEUtranFrequencyRef(v string) {
	o.EUtranFrequencyRef = &v
}

func (o ExternalEUTranCellSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalEUTranCellSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EUtranFrequencyRef) {
		toSerialize["EUtranFrequencyRef"] = o.EUtranFrequencyRef
	}
	return toSerialize, nil
}

type NullableExternalEUTranCellSingleAllOfAttributesAllOf struct {
	value *ExternalEUTranCellSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableExternalEUTranCellSingleAllOfAttributesAllOf) Get() *ExternalEUTranCellSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableExternalEUTranCellSingleAllOfAttributesAllOf) Set(val *ExternalEUTranCellSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalEUTranCellSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalEUTranCellSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalEUTranCellSingleAllOfAttributesAllOf(val *ExternalEUTranCellSingleAllOfAttributesAllOf) *NullableExternalEUTranCellSingleAllOfAttributesAllOf {
	return &NullableExternalEUTranCellSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableExternalEUTranCellSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalEUTranCellSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

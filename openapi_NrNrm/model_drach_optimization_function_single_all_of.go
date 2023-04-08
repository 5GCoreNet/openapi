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

// checks if the DRACHOptimizationFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DRACHOptimizationFunctionSingleAllOf{}

// DRACHOptimizationFunctionSingleAllOf struct for DRACHOptimizationFunctionSingleAllOf
type DRACHOptimizationFunctionSingleAllOf struct {
	Attributes *DRACHOptimizationFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewDRACHOptimizationFunctionSingleAllOf instantiates a new DRACHOptimizationFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDRACHOptimizationFunctionSingleAllOf() *DRACHOptimizationFunctionSingleAllOf {
	this := DRACHOptimizationFunctionSingleAllOf{}
	return &this
}

// NewDRACHOptimizationFunctionSingleAllOfWithDefaults instantiates a new DRACHOptimizationFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDRACHOptimizationFunctionSingleAllOfWithDefaults() *DRACHOptimizationFunctionSingleAllOf {
	this := DRACHOptimizationFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DRACHOptimizationFunctionSingleAllOf) GetAttributes() DRACHOptimizationFunctionSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret DRACHOptimizationFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DRACHOptimizationFunctionSingleAllOf) GetAttributesOk() (*DRACHOptimizationFunctionSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DRACHOptimizationFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given DRACHOptimizationFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *DRACHOptimizationFunctionSingleAllOf) SetAttributes(v DRACHOptimizationFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o DRACHOptimizationFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DRACHOptimizationFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableDRACHOptimizationFunctionSingleAllOf struct {
	value *DRACHOptimizationFunctionSingleAllOf
	isSet bool
}

func (v NullableDRACHOptimizationFunctionSingleAllOf) Get() *DRACHOptimizationFunctionSingleAllOf {
	return v.value
}

func (v *NullableDRACHOptimizationFunctionSingleAllOf) Set(val *DRACHOptimizationFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDRACHOptimizationFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDRACHOptimizationFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDRACHOptimizationFunctionSingleAllOf(val *DRACHOptimizationFunctionSingleAllOf) *NullableDRACHOptimizationFunctionSingleAllOf {
	return &NullableDRACHOptimizationFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableDRACHOptimizationFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDRACHOptimizationFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



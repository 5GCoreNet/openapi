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

// checks if the AmfRegionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfRegionSingleAllOf{}

// AmfRegionSingleAllOf struct for AmfRegionSingleAllOf
type AmfRegionSingleAllOf struct {
	Attributes *AmfRegionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewAmfRegionSingleAllOf instantiates a new AmfRegionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfRegionSingleAllOf() *AmfRegionSingleAllOf {
	this := AmfRegionSingleAllOf{}
	return &this
}

// NewAmfRegionSingleAllOfWithDefaults instantiates a new AmfRegionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfRegionSingleAllOfWithDefaults() *AmfRegionSingleAllOf {
	this := AmfRegionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOf) GetAttributes() AmfRegionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AmfRegionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOf) GetAttributesOk() (*AmfRegionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AmfRegionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *AmfRegionSingleAllOf) SetAttributes(v AmfRegionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o AmfRegionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfRegionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableAmfRegionSingleAllOf struct {
	value *AmfRegionSingleAllOf
	isSet bool
}

func (v NullableAmfRegionSingleAllOf) Get() *AmfRegionSingleAllOf {
	return v.value
}

func (v *NullableAmfRegionSingleAllOf) Set(val *AmfRegionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfRegionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfRegionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfRegionSingleAllOf(val *AmfRegionSingleAllOf) *NullableAmfRegionSingleAllOf {
	return &NullableAmfRegionSingleAllOf{value: val, isSet: true}
}

func (v NullableAmfRegionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfRegionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the BeamSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BeamSingleAllOf{}

// BeamSingleAllOf struct for BeamSingleAllOf
type BeamSingleAllOf struct {
	Attributes *BeamSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewBeamSingleAllOf instantiates a new BeamSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeamSingleAllOf() *BeamSingleAllOf {
	this := BeamSingleAllOf{}
	return &this
}

// NewBeamSingleAllOfWithDefaults instantiates a new BeamSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeamSingleAllOfWithDefaults() *BeamSingleAllOf {
	this := BeamSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BeamSingleAllOf) GetAttributes() BeamSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BeamSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOf) GetAttributesOk() (*BeamSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BeamSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BeamSingleAllOfAttributes and assigns it to the Attributes field.
func (o *BeamSingleAllOf) SetAttributes(v BeamSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o BeamSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BeamSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableBeamSingleAllOf struct {
	value *BeamSingleAllOf
	isSet bool
}

func (v NullableBeamSingleAllOf) Get() *BeamSingleAllOf {
	return v.value
}

func (v *NullableBeamSingleAllOf) Set(val *BeamSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBeamSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBeamSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeamSingleAllOf(val *BeamSingleAllOf) *NullableBeamSingleAllOf {
	return &NullableBeamSingleAllOf{value: val, isSet: true}
}

func (v NullableBeamSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeamSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



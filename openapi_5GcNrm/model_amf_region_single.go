/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the AmfRegionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfRegionSingle{}

// AmfRegionSingle struct for AmfRegionSingle
type AmfRegionSingle struct {
	Top
	Attributes *AmfRegionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewAmfRegionSingle instantiates a new AmfRegionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfRegionSingle(id NullableString) *AmfRegionSingle {
	this := AmfRegionSingle{}
	this.Id = id
	return &this
}

// NewAmfRegionSingleWithDefaults instantiates a new AmfRegionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfRegionSingleWithDefaults() *AmfRegionSingle {
	this := AmfRegionSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AmfRegionSingle) GetAttributes() AmfRegionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AmfRegionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingle) GetAttributesOk() (*AmfRegionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AmfRegionSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AmfRegionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *AmfRegionSingle) SetAttributes(v AmfRegionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o AmfRegionSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfRegionSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop, errTop := json.Marshal(o.Top)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	errTop = json.Unmarshal([]byte(serializedTop), &toSerialize)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableAmfRegionSingle struct {
	value *AmfRegionSingle
	isSet bool
}

func (v NullableAmfRegionSingle) Get() *AmfRegionSingle {
	return v.value
}

func (v *NullableAmfRegionSingle) Set(val *AmfRegionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfRegionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfRegionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfRegionSingle(val *AmfRegionSingle) *NullableAmfRegionSingle {
	return &NullableAmfRegionSingle{value: val, isSet: true}
}

func (v NullableAmfRegionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfRegionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GenericNrm

import (
	"encoding/json"
)

// checks if the ManagementDataCollectionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementDataCollectionSingleAllOf{}

// ManagementDataCollectionSingleAllOf struct for ManagementDataCollectionSingleAllOf
type ManagementDataCollectionSingleAllOf struct {
	Attributes *ManagementDataCollectionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewManagementDataCollectionSingleAllOf instantiates a new ManagementDataCollectionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementDataCollectionSingleAllOf() *ManagementDataCollectionSingleAllOf {
	this := ManagementDataCollectionSingleAllOf{}
	return &this
}

// NewManagementDataCollectionSingleAllOfWithDefaults instantiates a new ManagementDataCollectionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementDataCollectionSingleAllOfWithDefaults() *ManagementDataCollectionSingleAllOf {
	this := ManagementDataCollectionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagementDataCollectionSingleAllOf) GetAttributes() ManagementDataCollectionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ManagementDataCollectionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementDataCollectionSingleAllOf) GetAttributesOk() (*ManagementDataCollectionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagementDataCollectionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagementDataCollectionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ManagementDataCollectionSingleAllOf) SetAttributes(v ManagementDataCollectionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o ManagementDataCollectionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementDataCollectionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableManagementDataCollectionSingleAllOf struct {
	value *ManagementDataCollectionSingleAllOf
	isSet bool
}

func (v NullableManagementDataCollectionSingleAllOf) Get() *ManagementDataCollectionSingleAllOf {
	return v.value
}

func (v *NullableManagementDataCollectionSingleAllOf) Set(val *ManagementDataCollectionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementDataCollectionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementDataCollectionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementDataCollectionSingleAllOf(val *ManagementDataCollectionSingleAllOf) *NullableManagementDataCollectionSingleAllOf {
	return &NullableManagementDataCollectionSingleAllOf{value: val, isSet: true}
}

func (v NullableManagementDataCollectionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementDataCollectionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

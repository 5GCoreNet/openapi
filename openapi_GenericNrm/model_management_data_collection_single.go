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

// checks if the ManagementDataCollectionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementDataCollectionSingle{}

// ManagementDataCollectionSingle struct for ManagementDataCollectionSingle
type ManagementDataCollectionSingle struct {
	Top
	Attributes *ManagementDataCollectionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewManagementDataCollectionSingle instantiates a new ManagementDataCollectionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementDataCollectionSingle(id NullableString) *ManagementDataCollectionSingle {
	this := ManagementDataCollectionSingle{}
	this.Id = id
	return &this
}

// NewManagementDataCollectionSingleWithDefaults instantiates a new ManagementDataCollectionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementDataCollectionSingleWithDefaults() *ManagementDataCollectionSingle {
	this := ManagementDataCollectionSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagementDataCollectionSingle) GetAttributes() ManagementDataCollectionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ManagementDataCollectionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementDataCollectionSingle) GetAttributesOk() (*ManagementDataCollectionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagementDataCollectionSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagementDataCollectionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ManagementDataCollectionSingle) SetAttributes(v ManagementDataCollectionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o ManagementDataCollectionSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementDataCollectionSingle) ToMap() (map[string]interface{}, error) {
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

type NullableManagementDataCollectionSingle struct {
	value *ManagementDataCollectionSingle
	isSet bool
}

func (v NullableManagementDataCollectionSingle) Get() *ManagementDataCollectionSingle {
	return v.value
}

func (v *NullableManagementDataCollectionSingle) Set(val *ManagementDataCollectionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementDataCollectionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementDataCollectionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementDataCollectionSingle(val *ManagementDataCollectionSingle) *NullableManagementDataCollectionSingle {
	return &NullableManagementDataCollectionSingle{value: val, isSet: true}
}

func (v NullableManagementDataCollectionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementDataCollectionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

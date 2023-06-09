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

// checks if the Dynamic5QISetSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dynamic5QISetSingle{}

// Dynamic5QISetSingle struct for Dynamic5QISetSingle
type Dynamic5QISetSingle struct {
	Top
	Attributes *Dynamic5QISetSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewDynamic5QISetSingle instantiates a new Dynamic5QISetSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamic5QISetSingle(id NullableString) *Dynamic5QISetSingle {
	this := Dynamic5QISetSingle{}
	this.Id = id
	return &this
}

// NewDynamic5QISetSingleWithDefaults instantiates a new Dynamic5QISetSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamic5QISetSingleWithDefaults() *Dynamic5QISetSingle {
	this := Dynamic5QISetSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Dynamic5QISetSingle) GetAttributes() Dynamic5QISetSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret Dynamic5QISetSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dynamic5QISetSingle) GetAttributesOk() (*Dynamic5QISetSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Dynamic5QISetSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Dynamic5QISetSingleAllOfAttributes and assigns it to the Attributes field.
func (o *Dynamic5QISetSingle) SetAttributes(v Dynamic5QISetSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o Dynamic5QISetSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dynamic5QISetSingle) ToMap() (map[string]interface{}, error) {
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

type NullableDynamic5QISetSingle struct {
	value *Dynamic5QISetSingle
	isSet bool
}

func (v NullableDynamic5QISetSingle) Get() *Dynamic5QISetSingle {
	return v.value
}

func (v *NullableDynamic5QISetSingle) Set(val *Dynamic5QISetSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamic5QISetSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamic5QISetSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamic5QISetSingle(val *Dynamic5QISetSingle) *NullableDynamic5QISetSingle {
	return &NullableDynamic5QISetSingle{value: val, isSet: true}
}

func (v NullableDynamic5QISetSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamic5QISetSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

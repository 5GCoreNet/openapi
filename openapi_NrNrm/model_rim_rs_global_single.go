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

// checks if the RimRSGlobalSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RimRSGlobalSingle{}

// RimRSGlobalSingle struct for RimRSGlobalSingle
type RimRSGlobalSingle struct {
	Top
	Attributes *RimRSGlobalSingleAllOfAttributes `json:"attributes,omitempty"`
	RimRSSet   []RimRSSetSingle                  `json:"RimRSSet,omitempty"`
}

// NewRimRSGlobalSingle instantiates a new RimRSGlobalSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRimRSGlobalSingle(id NullableString) *RimRSGlobalSingle {
	this := RimRSGlobalSingle{}
	this.Id = id
	return &this
}

// NewRimRSGlobalSingleWithDefaults instantiates a new RimRSGlobalSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRimRSGlobalSingleWithDefaults() *RimRSGlobalSingle {
	this := RimRSGlobalSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RimRSGlobalSingle) GetAttributes() RimRSGlobalSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret RimRSGlobalSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSGlobalSingle) GetAttributesOk() (*RimRSGlobalSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RimRSGlobalSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given RimRSGlobalSingleAllOfAttributes and assigns it to the Attributes field.
func (o *RimRSGlobalSingle) SetAttributes(v RimRSGlobalSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetRimRSSet returns the RimRSSet field value if set, zero value otherwise.
func (o *RimRSGlobalSingle) GetRimRSSet() []RimRSSetSingle {
	if o == nil || IsNil(o.RimRSSet) {
		var ret []RimRSSetSingle
		return ret
	}
	return o.RimRSSet
}

// GetRimRSSetOk returns a tuple with the RimRSSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSGlobalSingle) GetRimRSSetOk() ([]RimRSSetSingle, bool) {
	if o == nil || IsNil(o.RimRSSet) {
		return nil, false
	}
	return o.RimRSSet, true
}

// HasRimRSSet returns a boolean if a field has been set.
func (o *RimRSGlobalSingle) HasRimRSSet() bool {
	if o != nil && !IsNil(o.RimRSSet) {
		return true
	}

	return false
}

// SetRimRSSet gets a reference to the given []RimRSSetSingle and assigns it to the RimRSSet field.
func (o *RimRSGlobalSingle) SetRimRSSet(v []RimRSSetSingle) {
	o.RimRSSet = v
}

func (o RimRSGlobalSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RimRSGlobalSingle) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RimRSSet) {
		toSerialize["RimRSSet"] = o.RimRSSet
	}
	return toSerialize, nil
}

type NullableRimRSGlobalSingle struct {
	value *RimRSGlobalSingle
	isSet bool
}

func (v NullableRimRSGlobalSingle) Get() *RimRSGlobalSingle {
	return v.value
}

func (v *NullableRimRSGlobalSingle) Set(val *RimRSGlobalSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableRimRSGlobalSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableRimRSGlobalSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRimRSGlobalSingle(val *RimRSGlobalSingle) *NullableRimRSGlobalSingle {
	return &NullableRimRSGlobalSingle{value: val, isSet: true}
}

func (v NullableRimRSGlobalSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRimRSGlobalSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

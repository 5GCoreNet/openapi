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

// checks if the RRMPolicyRatioSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RRMPolicyRatioSingle{}

// RRMPolicyRatioSingle struct for RRMPolicyRatioSingle
type RRMPolicyRatioSingle struct {
	Top
	Attributes *RRMPolicyRatioSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewRRMPolicyRatioSingle instantiates a new RRMPolicyRatioSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRRMPolicyRatioSingle(id NullableString) *RRMPolicyRatioSingle {
	this := RRMPolicyRatioSingle{}
	this.Id = id
	return &this
}

// NewRRMPolicyRatioSingleWithDefaults instantiates a new RRMPolicyRatioSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRRMPolicyRatioSingleWithDefaults() *RRMPolicyRatioSingle {
	this := RRMPolicyRatioSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RRMPolicyRatioSingle) GetAttributes() RRMPolicyRatioSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret RRMPolicyRatioSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RRMPolicyRatioSingle) GetAttributesOk() (*RRMPolicyRatioSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RRMPolicyRatioSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given RRMPolicyRatioSingleAllOfAttributes and assigns it to the Attributes field.
func (o *RRMPolicyRatioSingle) SetAttributes(v RRMPolicyRatioSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o RRMPolicyRatioSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RRMPolicyRatioSingle) ToMap() (map[string]interface{}, error) {
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

type NullableRRMPolicyRatioSingle struct {
	value *RRMPolicyRatioSingle
	isSet bool
}

func (v NullableRRMPolicyRatioSingle) Get() *RRMPolicyRatioSingle {
	return v.value
}

func (v *NullableRRMPolicyRatioSingle) Set(val *RRMPolicyRatioSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableRRMPolicyRatioSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableRRMPolicyRatioSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRRMPolicyRatioSingle(val *RRMPolicyRatioSingle) *NullableRRMPolicyRatioSingle {
	return &NullableRRMPolicyRatioSingle{value: val, isSet: true}
}

func (v NullableRRMPolicyRatioSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRRMPolicyRatioSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

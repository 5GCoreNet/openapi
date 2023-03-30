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

// checks if the RRMPolicyRatioSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RRMPolicyRatioSingleAllOf{}

// RRMPolicyRatioSingleAllOf struct for RRMPolicyRatioSingleAllOf
type RRMPolicyRatioSingleAllOf struct {
	Attributes *RRMPolicyRatioSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewRRMPolicyRatioSingleAllOf instantiates a new RRMPolicyRatioSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRRMPolicyRatioSingleAllOf() *RRMPolicyRatioSingleAllOf {
	this := RRMPolicyRatioSingleAllOf{}
	return &this
}

// NewRRMPolicyRatioSingleAllOfWithDefaults instantiates a new RRMPolicyRatioSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRRMPolicyRatioSingleAllOfWithDefaults() *RRMPolicyRatioSingleAllOf {
	this := RRMPolicyRatioSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RRMPolicyRatioSingleAllOf) GetAttributes() RRMPolicyRatioSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret RRMPolicyRatioSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RRMPolicyRatioSingleAllOf) GetAttributesOk() (*RRMPolicyRatioSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RRMPolicyRatioSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given RRMPolicyRatioSingleAllOfAttributes and assigns it to the Attributes field.
func (o *RRMPolicyRatioSingleAllOf) SetAttributes(v RRMPolicyRatioSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o RRMPolicyRatioSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RRMPolicyRatioSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableRRMPolicyRatioSingleAllOf struct {
	value *RRMPolicyRatioSingleAllOf
	isSet bool
}

func (v NullableRRMPolicyRatioSingleAllOf) Get() *RRMPolicyRatioSingleAllOf {
	return v.value
}

func (v *NullableRRMPolicyRatioSingleAllOf) Set(val *RRMPolicyRatioSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRRMPolicyRatioSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRRMPolicyRatioSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRRMPolicyRatioSingleAllOf(val *RRMPolicyRatioSingleAllOf) *NullableRRMPolicyRatioSingleAllOf {
	return &NullableRRMPolicyRatioSingleAllOf{value: val, isSet: true}
}

func (v NullableRRMPolicyRatioSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRRMPolicyRatioSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// checks if the AssuranceReportSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssuranceReportSingleAllOf{}

// AssuranceReportSingleAllOf struct for AssuranceReportSingleAllOf
type AssuranceReportSingleAllOf struct {
	Attributes *AssuranceReportSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewAssuranceReportSingleAllOf instantiates a new AssuranceReportSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssuranceReportSingleAllOf() *AssuranceReportSingleAllOf {
	this := AssuranceReportSingleAllOf{}
	return &this
}

// NewAssuranceReportSingleAllOfWithDefaults instantiates a new AssuranceReportSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssuranceReportSingleAllOfWithDefaults() *AssuranceReportSingleAllOf {
	this := AssuranceReportSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AssuranceReportSingleAllOf) GetAttributes() AssuranceReportSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AssuranceReportSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceReportSingleAllOf) GetAttributesOk() (*AssuranceReportSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AssuranceReportSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AssuranceReportSingleAllOfAttributes and assigns it to the Attributes field.
func (o *AssuranceReportSingleAllOf) SetAttributes(v AssuranceReportSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o AssuranceReportSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssuranceReportSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableAssuranceReportSingleAllOf struct {
	value *AssuranceReportSingleAllOf
	isSet bool
}

func (v NullableAssuranceReportSingleAllOf) Get() *AssuranceReportSingleAllOf {
	return v.value
}

func (v *NullableAssuranceReportSingleAllOf) Set(val *AssuranceReportSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssuranceReportSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssuranceReportSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssuranceReportSingleAllOf(val *AssuranceReportSingleAllOf) *NullableAssuranceReportSingleAllOf {
	return &NullableAssuranceReportSingleAllOf{value: val, isSet: true}
}

func (v NullableAssuranceReportSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssuranceReportSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

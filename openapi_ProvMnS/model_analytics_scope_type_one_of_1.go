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

// checks if the AnalyticsScopeTypeOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsScopeTypeOneOf1{}

// AnalyticsScopeTypeOneOf1 struct for AnalyticsScopeTypeOneOf1
type AnalyticsScopeTypeOneOf1 struct {
	AreaScope *GeoArea `json:"areaScope,omitempty"`
}

// NewAnalyticsScopeTypeOneOf1 instantiates a new AnalyticsScopeTypeOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsScopeTypeOneOf1() *AnalyticsScopeTypeOneOf1 {
	this := AnalyticsScopeTypeOneOf1{}
	return &this
}

// NewAnalyticsScopeTypeOneOf1WithDefaults instantiates a new AnalyticsScopeTypeOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsScopeTypeOneOf1WithDefaults() *AnalyticsScopeTypeOneOf1 {
	this := AnalyticsScopeTypeOneOf1{}
	return &this
}

// GetAreaScope returns the AreaScope field value if set, zero value otherwise.
func (o *AnalyticsScopeTypeOneOf1) GetAreaScope() GeoArea {
	if o == nil || IsNil(o.AreaScope) {
		var ret GeoArea
		return ret
	}
	return *o.AreaScope
}

// GetAreaScopeOk returns a tuple with the AreaScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsScopeTypeOneOf1) GetAreaScopeOk() (*GeoArea, bool) {
	if o == nil || IsNil(o.AreaScope) {
		return nil, false
	}
	return o.AreaScope, true
}

// HasAreaScope returns a boolean if a field has been set.
func (o *AnalyticsScopeTypeOneOf1) HasAreaScope() bool {
	if o != nil && !IsNil(o.AreaScope) {
		return true
	}

	return false
}

// SetAreaScope gets a reference to the given GeoArea and assigns it to the AreaScope field.
func (o *AnalyticsScopeTypeOneOf1) SetAreaScope(v GeoArea) {
	o.AreaScope = &v
}

func (o AnalyticsScopeTypeOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsScopeTypeOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AreaScope) {
		toSerialize["areaScope"] = o.AreaScope
	}
	return toSerialize, nil
}

type NullableAnalyticsScopeTypeOneOf1 struct {
	value *AnalyticsScopeTypeOneOf1
	isSet bool
}

func (v NullableAnalyticsScopeTypeOneOf1) Get() *AnalyticsScopeTypeOneOf1 {
	return v.value
}

func (v *NullableAnalyticsScopeTypeOneOf1) Set(val *AnalyticsScopeTypeOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsScopeTypeOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsScopeTypeOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsScopeTypeOneOf1(val *AnalyticsScopeTypeOneOf1) *NullableAnalyticsScopeTypeOneOf1 {
	return &NullableAnalyticsScopeTypeOneOf1{value: val, isSet: true}
}

func (v NullableAnalyticsScopeTypeOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsScopeTypeOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



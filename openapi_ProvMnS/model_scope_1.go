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

// checks if the Scope1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Scope1{}

// Scope1 struct for Scope1
type Scope1 struct {
	ScopeType  *string `json:"scopeType,omitempty"`
	ScopeLevel *int32  `json:"scopeLevel,omitempty"`
}

// NewScope1 instantiates a new Scope1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScope1() *Scope1 {
	this := Scope1{}
	return &this
}

// NewScope1WithDefaults instantiates a new Scope1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScope1WithDefaults() *Scope1 {
	this := Scope1{}
	return &this
}

// GetScopeType returns the ScopeType field value if set, zero value otherwise.
func (o *Scope1) GetScopeType() string {
	if o == nil || IsNil(o.ScopeType) {
		var ret string
		return ret
	}
	return *o.ScopeType
}

// GetScopeTypeOk returns a tuple with the ScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope1) GetScopeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeType) {
		return nil, false
	}
	return o.ScopeType, true
}

// HasScopeType returns a boolean if a field has been set.
func (o *Scope1) HasScopeType() bool {
	if o != nil && !IsNil(o.ScopeType) {
		return true
	}

	return false
}

// SetScopeType gets a reference to the given string and assigns it to the ScopeType field.
func (o *Scope1) SetScopeType(v string) {
	o.ScopeType = &v
}

// GetScopeLevel returns the ScopeLevel field value if set, zero value otherwise.
func (o *Scope1) GetScopeLevel() int32 {
	if o == nil || IsNil(o.ScopeLevel) {
		var ret int32
		return ret
	}
	return *o.ScopeLevel
}

// GetScopeLevelOk returns a tuple with the ScopeLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope1) GetScopeLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.ScopeLevel) {
		return nil, false
	}
	return o.ScopeLevel, true
}

// HasScopeLevel returns a boolean if a field has been set.
func (o *Scope1) HasScopeLevel() bool {
	if o != nil && !IsNil(o.ScopeLevel) {
		return true
	}

	return false
}

// SetScopeLevel gets a reference to the given int32 and assigns it to the ScopeLevel field.
func (o *Scope1) SetScopeLevel(v int32) {
	o.ScopeLevel = &v
}

func (o Scope1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Scope1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScopeType) {
		toSerialize["scopeType"] = o.ScopeType
	}
	if !IsNil(o.ScopeLevel) {
		toSerialize["scopeLevel"] = o.ScopeLevel
	}
	return toSerialize, nil
}

type NullableScope1 struct {
	value *Scope1
	isSet bool
}

func (v NullableScope1) Get() *Scope1 {
	return v.value
}

func (v *NullableScope1) Set(val *Scope1) {
	v.value = val
	v.isSet = true
}

func (v NullableScope1) IsSet() bool {
	return v.isSet
}

func (v *NullableScope1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScope1(val *Scope1) *NullableScope1 {
	return &NullableScope1{value: val, isSet: true}
}

func (v NullableScope1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScope1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

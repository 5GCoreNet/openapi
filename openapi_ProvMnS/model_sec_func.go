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

// checks if the SecFunc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecFunc{}

// SecFunc struct for SecFunc
type SecFunc struct {
	SecFunId   *string  `json:"secFunId,omitempty"`
	SecFunType *string  `json:"secFunType,omitempty"`
	SecRules   []string `json:"secRules,omitempty"`
}

// NewSecFunc instantiates a new SecFunc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecFunc() *SecFunc {
	this := SecFunc{}
	return &this
}

// NewSecFuncWithDefaults instantiates a new SecFunc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecFuncWithDefaults() *SecFunc {
	this := SecFunc{}
	return &this
}

// GetSecFunId returns the SecFunId field value if set, zero value otherwise.
func (o *SecFunc) GetSecFunId() string {
	if o == nil || IsNil(o.SecFunId) {
		var ret string
		return ret
	}
	return *o.SecFunId
}

// GetSecFunIdOk returns a tuple with the SecFunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecFunc) GetSecFunIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecFunId) {
		return nil, false
	}
	return o.SecFunId, true
}

// HasSecFunId returns a boolean if a field has been set.
func (o *SecFunc) HasSecFunId() bool {
	if o != nil && !IsNil(o.SecFunId) {
		return true
	}

	return false
}

// SetSecFunId gets a reference to the given string and assigns it to the SecFunId field.
func (o *SecFunc) SetSecFunId(v string) {
	o.SecFunId = &v
}

// GetSecFunType returns the SecFunType field value if set, zero value otherwise.
func (o *SecFunc) GetSecFunType() string {
	if o == nil || IsNil(o.SecFunType) {
		var ret string
		return ret
	}
	return *o.SecFunType
}

// GetSecFunTypeOk returns a tuple with the SecFunType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecFunc) GetSecFunTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SecFunType) {
		return nil, false
	}
	return o.SecFunType, true
}

// HasSecFunType returns a boolean if a field has been set.
func (o *SecFunc) HasSecFunType() bool {
	if o != nil && !IsNil(o.SecFunType) {
		return true
	}

	return false
}

// SetSecFunType gets a reference to the given string and assigns it to the SecFunType field.
func (o *SecFunc) SetSecFunType(v string) {
	o.SecFunType = &v
}

// GetSecRules returns the SecRules field value if set, zero value otherwise.
func (o *SecFunc) GetSecRules() []string {
	if o == nil || IsNil(o.SecRules) {
		var ret []string
		return ret
	}
	return o.SecRules
}

// GetSecRulesOk returns a tuple with the SecRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecFunc) GetSecRulesOk() ([]string, bool) {
	if o == nil || IsNil(o.SecRules) {
		return nil, false
	}
	return o.SecRules, true
}

// HasSecRules returns a boolean if a field has been set.
func (o *SecFunc) HasSecRules() bool {
	if o != nil && !IsNil(o.SecRules) {
		return true
	}

	return false
}

// SetSecRules gets a reference to the given []string and assigns it to the SecRules field.
func (o *SecFunc) SetSecRules(v []string) {
	o.SecRules = v
}

func (o SecFunc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecFunc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecFunId) {
		toSerialize["secFunId"] = o.SecFunId
	}
	if !IsNil(o.SecFunType) {
		toSerialize["secFunType"] = o.SecFunType
	}
	if !IsNil(o.SecRules) {
		toSerialize["secRules"] = o.SecRules
	}
	return toSerialize, nil
}

type NullableSecFunc struct {
	value *SecFunc
	isSet bool
}

func (v NullableSecFunc) Get() *SecFunc {
	return v.value
}

func (v *NullableSecFunc) Set(val *SecFunc) {
	v.value = val
	v.isSet = true
}

func (v NullableSecFunc) IsSet() bool {
	return v.isSet
}

func (v *NullableSecFunc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecFunc(val *SecFunc) *NullableSecFunc {
	return &NullableSecFunc{value: val, isSet: true}
}

func (v NullableSecFunc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecFunc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

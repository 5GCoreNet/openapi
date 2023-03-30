/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the KeyAmf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyAmf{}

// KeyAmf Represents the Kamf or K'amf
type KeyAmf struct {
	KeyType KeyAmfType `json:"keyType"`
	KeyVal string `json:"keyVal"`
}

// NewKeyAmf instantiates a new KeyAmf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyAmf(keyType KeyAmfType, keyVal string) *KeyAmf {
	this := KeyAmf{}
	this.KeyType = keyType
	this.KeyVal = keyVal
	return &this
}

// NewKeyAmfWithDefaults instantiates a new KeyAmf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyAmfWithDefaults() *KeyAmf {
	this := KeyAmf{}
	return &this
}

// GetKeyType returns the KeyType field value
func (o *KeyAmf) GetKeyType() KeyAmfType {
	if o == nil {
		var ret KeyAmfType
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *KeyAmf) GetKeyTypeOk() (*KeyAmfType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyType, true
}

// SetKeyType sets field value
func (o *KeyAmf) SetKeyType(v KeyAmfType) {
	o.KeyType = v
}

// GetKeyVal returns the KeyVal field value
func (o *KeyAmf) GetKeyVal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyVal
}

// GetKeyValOk returns a tuple with the KeyVal field value
// and a boolean to check if the value has been set.
func (o *KeyAmf) GetKeyValOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyVal, true
}

// SetKeyVal sets field value
func (o *KeyAmf) SetKeyVal(v string) {
	o.KeyVal = v
}

func (o KeyAmf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyAmf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keyType"] = o.KeyType
	toSerialize["keyVal"] = o.KeyVal
	return toSerialize, nil
}

type NullableKeyAmf struct {
	value *KeyAmf
	isSet bool
}

func (v NullableKeyAmf) Get() *KeyAmf {
	return v.value
}

func (v *NullableKeyAmf) Set(val *KeyAmf) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyAmf) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyAmf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyAmf(val *KeyAmf) *NullableKeyAmf {
	return &NullableKeyAmf{value: val, isSet: true}
}

func (v NullableKeyAmf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyAmf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



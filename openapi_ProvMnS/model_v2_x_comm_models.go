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

// checks if the V2XCommModels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2XCommModels{}

// V2XCommModels struct for V2XCommModels
type V2XCommModels struct {
	ServAttrCom *ServAttrCom `json:"servAttrCom,omitempty"`
	V2XMode     *Support     `json:"v2XMode,omitempty"`
}

// NewV2XCommModels instantiates a new V2XCommModels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2XCommModels() *V2XCommModels {
	this := V2XCommModels{}
	return &this
}

// NewV2XCommModelsWithDefaults instantiates a new V2XCommModels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2XCommModelsWithDefaults() *V2XCommModels {
	this := V2XCommModels{}
	return &this
}

// GetServAttrCom returns the ServAttrCom field value if set, zero value otherwise.
func (o *V2XCommModels) GetServAttrCom() ServAttrCom {
	if o == nil || IsNil(o.ServAttrCom) {
		var ret ServAttrCom
		return ret
	}
	return *o.ServAttrCom
}

// GetServAttrComOk returns a tuple with the ServAttrCom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2XCommModels) GetServAttrComOk() (*ServAttrCom, bool) {
	if o == nil || IsNil(o.ServAttrCom) {
		return nil, false
	}
	return o.ServAttrCom, true
}

// HasServAttrCom returns a boolean if a field has been set.
func (o *V2XCommModels) HasServAttrCom() bool {
	if o != nil && !IsNil(o.ServAttrCom) {
		return true
	}

	return false
}

// SetServAttrCom gets a reference to the given ServAttrCom and assigns it to the ServAttrCom field.
func (o *V2XCommModels) SetServAttrCom(v ServAttrCom) {
	o.ServAttrCom = &v
}

// GetV2XMode returns the V2XMode field value if set, zero value otherwise.
func (o *V2XCommModels) GetV2XMode() Support {
	if o == nil || IsNil(o.V2XMode) {
		var ret Support
		return ret
	}
	return *o.V2XMode
}

// GetV2XModeOk returns a tuple with the V2XMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2XCommModels) GetV2XModeOk() (*Support, bool) {
	if o == nil || IsNil(o.V2XMode) {
		return nil, false
	}
	return o.V2XMode, true
}

// HasV2XMode returns a boolean if a field has been set.
func (o *V2XCommModels) HasV2XMode() bool {
	if o != nil && !IsNil(o.V2XMode) {
		return true
	}

	return false
}

// SetV2XMode gets a reference to the given Support and assigns it to the V2XMode field.
func (o *V2XCommModels) SetV2XMode(v Support) {
	o.V2XMode = &v
}

func (o V2XCommModels) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2XCommModels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServAttrCom) {
		toSerialize["servAttrCom"] = o.ServAttrCom
	}
	if !IsNil(o.V2XMode) {
		toSerialize["v2XMode"] = o.V2XMode
	}
	return toSerialize, nil
}

type NullableV2XCommModels struct {
	value *V2XCommModels
	isSet bool
}

func (v NullableV2XCommModels) Get() *V2XCommModels {
	return v.value
}

func (v *NullableV2XCommModels) Set(val *V2XCommModels) {
	v.value = val
	v.isSet = true
}

func (v NullableV2XCommModels) IsSet() bool {
	return v.isSet
}

func (v *NullableV2XCommModels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2XCommModels(val *V2XCommModels) *NullableV2XCommModels {
	return &NullableV2XCommModels{value: val, isSet: true}
}

func (v NullableV2XCommModels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2XCommModels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

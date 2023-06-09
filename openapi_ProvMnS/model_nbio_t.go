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

// checks if the NBIoT type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NBIoT{}

// NBIoT struct for NBIoT
type NBIoT struct {
	ServAttrCom *ServAttrCom `json:"servAttrCom,omitempty"`
	Support     *Support     `json:"support,omitempty"`
}

// NewNBIoT instantiates a new NBIoT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNBIoT() *NBIoT {
	this := NBIoT{}
	return &this
}

// NewNBIoTWithDefaults instantiates a new NBIoT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNBIoTWithDefaults() *NBIoT {
	this := NBIoT{}
	return &this
}

// GetServAttrCom returns the ServAttrCom field value if set, zero value otherwise.
func (o *NBIoT) GetServAttrCom() ServAttrCom {
	if o == nil || IsNil(o.ServAttrCom) {
		var ret ServAttrCom
		return ret
	}
	return *o.ServAttrCom
}

// GetServAttrComOk returns a tuple with the ServAttrCom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NBIoT) GetServAttrComOk() (*ServAttrCom, bool) {
	if o == nil || IsNil(o.ServAttrCom) {
		return nil, false
	}
	return o.ServAttrCom, true
}

// HasServAttrCom returns a boolean if a field has been set.
func (o *NBIoT) HasServAttrCom() bool {
	if o != nil && !IsNil(o.ServAttrCom) {
		return true
	}

	return false
}

// SetServAttrCom gets a reference to the given ServAttrCom and assigns it to the ServAttrCom field.
func (o *NBIoT) SetServAttrCom(v ServAttrCom) {
	o.ServAttrCom = &v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *NBIoT) GetSupport() Support {
	if o == nil || IsNil(o.Support) {
		var ret Support
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NBIoT) GetSupportOk() (*Support, bool) {
	if o == nil || IsNil(o.Support) {
		return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *NBIoT) HasSupport() bool {
	if o != nil && !IsNil(o.Support) {
		return true
	}

	return false
}

// SetSupport gets a reference to the given Support and assigns it to the Support field.
func (o *NBIoT) SetSupport(v Support) {
	o.Support = &v
}

func (o NBIoT) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NBIoT) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServAttrCom) {
		toSerialize["servAttrCom"] = o.ServAttrCom
	}
	if !IsNil(o.Support) {
		toSerialize["support"] = o.Support
	}
	return toSerialize, nil
}

type NullableNBIoT struct {
	value *NBIoT
	isSet bool
}

func (v NullableNBIoT) Get() *NBIoT {
	return v.value
}

func (v *NullableNBIoT) Set(val *NBIoT) {
	v.value = val
	v.isSet = true
}

func (v NullableNBIoT) IsSet() bool {
	return v.isSet
}

func (v *NullableNBIoT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNBIoT(val *NBIoT) *NullableNBIoT {
	return &NullableNBIoT{value: val, isSet: true}
}

func (v NullableNBIoT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNBIoT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

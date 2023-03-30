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

// checks if the MappingSetIDBackhaulAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MappingSetIDBackhaulAddress{}

// MappingSetIDBackhaulAddress struct for MappingSetIDBackhaulAddress
type MappingSetIDBackhaulAddress struct {
	SetID *int32 `json:"setID,omitempty"`
	BackhaulAddress *BackhaulAddress `json:"backhaulAddress,omitempty"`
}

// NewMappingSetIDBackhaulAddress instantiates a new MappingSetIDBackhaulAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMappingSetIDBackhaulAddress() *MappingSetIDBackhaulAddress {
	this := MappingSetIDBackhaulAddress{}
	return &this
}

// NewMappingSetIDBackhaulAddressWithDefaults instantiates a new MappingSetIDBackhaulAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMappingSetIDBackhaulAddressWithDefaults() *MappingSetIDBackhaulAddress {
	this := MappingSetIDBackhaulAddress{}
	return &this
}

// GetSetID returns the SetID field value if set, zero value otherwise.
func (o *MappingSetIDBackhaulAddress) GetSetID() int32 {
	if o == nil || IsNil(o.SetID) {
		var ret int32
		return ret
	}
	return *o.SetID
}

// GetSetIDOk returns a tuple with the SetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappingSetIDBackhaulAddress) GetSetIDOk() (*int32, bool) {
	if o == nil || IsNil(o.SetID) {
		return nil, false
	}
	return o.SetID, true
}

// HasSetID returns a boolean if a field has been set.
func (o *MappingSetIDBackhaulAddress) HasSetID() bool {
	if o != nil && !IsNil(o.SetID) {
		return true
	}

	return false
}

// SetSetID gets a reference to the given int32 and assigns it to the SetID field.
func (o *MappingSetIDBackhaulAddress) SetSetID(v int32) {
	o.SetID = &v
}

// GetBackhaulAddress returns the BackhaulAddress field value if set, zero value otherwise.
func (o *MappingSetIDBackhaulAddress) GetBackhaulAddress() BackhaulAddress {
	if o == nil || IsNil(o.BackhaulAddress) {
		var ret BackhaulAddress
		return ret
	}
	return *o.BackhaulAddress
}

// GetBackhaulAddressOk returns a tuple with the BackhaulAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappingSetIDBackhaulAddress) GetBackhaulAddressOk() (*BackhaulAddress, bool) {
	if o == nil || IsNil(o.BackhaulAddress) {
		return nil, false
	}
	return o.BackhaulAddress, true
}

// HasBackhaulAddress returns a boolean if a field has been set.
func (o *MappingSetIDBackhaulAddress) HasBackhaulAddress() bool {
	if o != nil && !IsNil(o.BackhaulAddress) {
		return true
	}

	return false
}

// SetBackhaulAddress gets a reference to the given BackhaulAddress and assigns it to the BackhaulAddress field.
func (o *MappingSetIDBackhaulAddress) SetBackhaulAddress(v BackhaulAddress) {
	o.BackhaulAddress = &v
}

func (o MappingSetIDBackhaulAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MappingSetIDBackhaulAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SetID) {
		toSerialize["setID"] = o.SetID
	}
	if !IsNil(o.BackhaulAddress) {
		toSerialize["backhaulAddress"] = o.BackhaulAddress
	}
	return toSerialize, nil
}

type NullableMappingSetIDBackhaulAddress struct {
	value *MappingSetIDBackhaulAddress
	isSet bool
}

func (v NullableMappingSetIDBackhaulAddress) Get() *MappingSetIDBackhaulAddress {
	return v.value
}

func (v *NullableMappingSetIDBackhaulAddress) Set(val *MappingSetIDBackhaulAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableMappingSetIDBackhaulAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableMappingSetIDBackhaulAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMappingSetIDBackhaulAddress(val *MappingSetIDBackhaulAddress) *NullableMappingSetIDBackhaulAddress {
	return &NullableMappingSetIDBackhaulAddress{value: val, isSet: true}
}

func (v NullableMappingSetIDBackhaulAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMappingSetIDBackhaulAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the AppPortId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPortId{}

// AppPortId struct for AppPortId
type AppPortId struct {
	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	DestinationPort *int32 `json:"destinationPort,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	OriginatorPort *int32 `json:"originatorPort,omitempty"`
}

// NewAppPortId instantiates a new AppPortId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPortId() *AppPortId {
	this := AppPortId{}
	return &this
}

// NewAppPortIdWithDefaults instantiates a new AppPortId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPortIdWithDefaults() *AppPortId {
	this := AppPortId{}
	return &this
}

// GetDestinationPort returns the DestinationPort field value if set, zero value otherwise.
func (o *AppPortId) GetDestinationPort() int32 {
	if o == nil || IsNil(o.DestinationPort) {
		var ret int32
		return ret
	}
	return *o.DestinationPort
}

// GetDestinationPortOk returns a tuple with the DestinationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPortId) GetDestinationPortOk() (*int32, bool) {
	if o == nil || IsNil(o.DestinationPort) {
		return nil, false
	}
	return o.DestinationPort, true
}

// HasDestinationPort returns a boolean if a field has been set.
func (o *AppPortId) HasDestinationPort() bool {
	if o != nil && !IsNil(o.DestinationPort) {
		return true
	}

	return false
}

// SetDestinationPort gets a reference to the given int32 and assigns it to the DestinationPort field.
func (o *AppPortId) SetDestinationPort(v int32) {
	o.DestinationPort = &v
}

// GetOriginatorPort returns the OriginatorPort field value if set, zero value otherwise.
func (o *AppPortId) GetOriginatorPort() int32 {
	if o == nil || IsNil(o.OriginatorPort) {
		var ret int32
		return ret
	}
	return *o.OriginatorPort
}

// GetOriginatorPortOk returns a tuple with the OriginatorPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPortId) GetOriginatorPortOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginatorPort) {
		return nil, false
	}
	return o.OriginatorPort, true
}

// HasOriginatorPort returns a boolean if a field has been set.
func (o *AppPortId) HasOriginatorPort() bool {
	if o != nil && !IsNil(o.OriginatorPort) {
		return true
	}

	return false
}

// SetOriginatorPort gets a reference to the given int32 and assigns it to the OriginatorPort field.
func (o *AppPortId) SetOriginatorPort(v int32) {
	o.OriginatorPort = &v
}

func (o AppPortId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPortId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DestinationPort) {
		toSerialize["destinationPort"] = o.DestinationPort
	}
	if !IsNil(o.OriginatorPort) {
		toSerialize["originatorPort"] = o.OriginatorPort
	}
	return toSerialize, nil
}

type NullableAppPortId struct {
	value *AppPortId
	isSet bool
}

func (v NullableAppPortId) Get() *AppPortId {
	return v.value
}

func (v *NullableAppPortId) Set(val *AppPortId) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPortId) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPortId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPortId(val *AppPortId) *NullableAppPortId {
	return &NullableAppPortId{value: val, isSet: true}
}

func (v NullableAppPortId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPortId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

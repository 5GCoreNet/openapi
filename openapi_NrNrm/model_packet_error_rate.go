/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the PacketErrorRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PacketErrorRate{}

// PacketErrorRate struct for PacketErrorRate
type PacketErrorRate struct {
	Scalar   *int32 `json:"scalar,omitempty"`
	Exponent *int32 `json:"exponent,omitempty"`
}

// NewPacketErrorRate instantiates a new PacketErrorRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPacketErrorRate() *PacketErrorRate {
	this := PacketErrorRate{}
	return &this
}

// NewPacketErrorRateWithDefaults instantiates a new PacketErrorRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPacketErrorRateWithDefaults() *PacketErrorRate {
	this := PacketErrorRate{}
	return &this
}

// GetScalar returns the Scalar field value if set, zero value otherwise.
func (o *PacketErrorRate) GetScalar() int32 {
	if o == nil || IsNil(o.Scalar) {
		var ret int32
		return ret
	}
	return *o.Scalar
}

// GetScalarOk returns a tuple with the Scalar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketErrorRate) GetScalarOk() (*int32, bool) {
	if o == nil || IsNil(o.Scalar) {
		return nil, false
	}
	return o.Scalar, true
}

// HasScalar returns a boolean if a field has been set.
func (o *PacketErrorRate) HasScalar() bool {
	if o != nil && !IsNil(o.Scalar) {
		return true
	}

	return false
}

// SetScalar gets a reference to the given int32 and assigns it to the Scalar field.
func (o *PacketErrorRate) SetScalar(v int32) {
	o.Scalar = &v
}

// GetExponent returns the Exponent field value if set, zero value otherwise.
func (o *PacketErrorRate) GetExponent() int32 {
	if o == nil || IsNil(o.Exponent) {
		var ret int32
		return ret
	}
	return *o.Exponent
}

// GetExponentOk returns a tuple with the Exponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PacketErrorRate) GetExponentOk() (*int32, bool) {
	if o == nil || IsNil(o.Exponent) {
		return nil, false
	}
	return o.Exponent, true
}

// HasExponent returns a boolean if a field has been set.
func (o *PacketErrorRate) HasExponent() bool {
	if o != nil && !IsNil(o.Exponent) {
		return true
	}

	return false
}

// SetExponent gets a reference to the given int32 and assigns it to the Exponent field.
func (o *PacketErrorRate) SetExponent(v int32) {
	o.Exponent = &v
}

func (o PacketErrorRate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PacketErrorRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scalar) {
		toSerialize["scalar"] = o.Scalar
	}
	if !IsNil(o.Exponent) {
		toSerialize["exponent"] = o.Exponent
	}
	return toSerialize, nil
}

type NullablePacketErrorRate struct {
	value *PacketErrorRate
	isSet bool
}

func (v NullablePacketErrorRate) Get() *PacketErrorRate {
	return v.value
}

func (v *NullablePacketErrorRate) Set(val *PacketErrorRate) {
	v.value = val
	v.isSet = true
}

func (v NullablePacketErrorRate) IsSet() bool {
	return v.isSet
}

func (v *NullablePacketErrorRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePacketErrorRate(val *PacketErrorRate) *NullablePacketErrorRate {
	return &NullablePacketErrorRate{value: val, isSet: true}
}

func (v NullablePacketErrorRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePacketErrorRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

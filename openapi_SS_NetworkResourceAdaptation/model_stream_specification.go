/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceAdaptation

import (
	"encoding/json"
)

// checks if the StreamSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamSpecification{}

// StreamSpecification Stream specification includes MAC addresses of the source and destination DS-TT ports.
type StreamSpecification struct {
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	SrcMacAddr string `json:"srcMacAddr"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	DstMacAddr string `json:"dstMacAddr"`
}

// NewStreamSpecification instantiates a new StreamSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamSpecification(srcMacAddr string, dstMacAddr string) *StreamSpecification {
	this := StreamSpecification{}
	this.SrcMacAddr = srcMacAddr
	this.DstMacAddr = dstMacAddr
	return &this
}

// NewStreamSpecificationWithDefaults instantiates a new StreamSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamSpecificationWithDefaults() *StreamSpecification {
	this := StreamSpecification{}
	return &this
}

// GetSrcMacAddr returns the SrcMacAddr field value
func (o *StreamSpecification) GetSrcMacAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SrcMacAddr
}

// GetSrcMacAddrOk returns a tuple with the SrcMacAddr field value
// and a boolean to check if the value has been set.
func (o *StreamSpecification) GetSrcMacAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SrcMacAddr, true
}

// SetSrcMacAddr sets field value
func (o *StreamSpecification) SetSrcMacAddr(v string) {
	o.SrcMacAddr = v
}

// GetDstMacAddr returns the DstMacAddr field value
func (o *StreamSpecification) GetDstMacAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DstMacAddr
}

// GetDstMacAddrOk returns a tuple with the DstMacAddr field value
// and a boolean to check if the value has been set.
func (o *StreamSpecification) GetDstMacAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DstMacAddr, true
}

// SetDstMacAddr sets field value
func (o *StreamSpecification) SetDstMacAddr(v string) {
	o.DstMacAddr = v
}

func (o StreamSpecification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["srcMacAddr"] = o.SrcMacAddr
	toSerialize["dstMacAddr"] = o.DstMacAddr
	return toSerialize, nil
}

type NullableStreamSpecification struct {
	value *StreamSpecification
	isSet bool
}

func (v NullableStreamSpecification) Get() *StreamSpecification {
	return v.value
}

func (v *NullableStreamSpecification) Set(val *StreamSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamSpecification(val *StreamSpecification) *NullableStreamSpecification {
	return &NullableStreamSpecification{value: val, isSet: true}
}

func (v NullableStreamSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

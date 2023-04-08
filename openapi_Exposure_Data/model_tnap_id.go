/*
Unified Data Repository Service API file for structured data for exposure

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Exposure_Data

import (
	"encoding/json"
)

// checks if the TnapId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TnapId{}

// TnapId Contain the TNAP Identifier see clause5.6.2 of 3GPP TS 23.501.
type TnapId struct {
	// This IE shall be present if the UE is accessing the 5GC via a trusted WLAN access network.When present, it shall contain the SSID of the access point to which the UE is attached, that is received over NGAP,  see IEEE Std 802.11-2012.  
	SsId *string `json:"ssId,omitempty"`
	// When present, it shall contain the BSSID of the access point to which the UE is attached, that is received over NGAP, see IEEE Std 802.11-2012.  
	BssId *string `json:"bssId,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	CivicAddress *string `json:"civicAddress,omitempty"`
}

// NewTnapId instantiates a new TnapId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTnapId() *TnapId {
	this := TnapId{}
	return &this
}

// NewTnapIdWithDefaults instantiates a new TnapId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTnapIdWithDefaults() *TnapId {
	this := TnapId{}
	return &this
}

// GetSsId returns the SsId field value if set, zero value otherwise.
func (o *TnapId) GetSsId() string {
	if o == nil || isNil(o.SsId) {
		var ret string
		return ret
	}
	return *o.SsId
}

// GetSsIdOk returns a tuple with the SsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TnapId) GetSsIdOk() (*string, bool) {
	if o == nil || isNil(o.SsId) {
		return nil, false
	}
	return o.SsId, true
}

// HasSsId returns a boolean if a field has been set.
func (o *TnapId) HasSsId() bool {
	if o != nil && !isNil(o.SsId) {
		return true
	}

	return false
}

// SetSsId gets a reference to the given string and assigns it to the SsId field.
func (o *TnapId) SetSsId(v string) {
	o.SsId = &v
}

// GetBssId returns the BssId field value if set, zero value otherwise.
func (o *TnapId) GetBssId() string {
	if o == nil || isNil(o.BssId) {
		var ret string
		return ret
	}
	return *o.BssId
}

// GetBssIdOk returns a tuple with the BssId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TnapId) GetBssIdOk() (*string, bool) {
	if o == nil || isNil(o.BssId) {
		return nil, false
	}
	return o.BssId, true
}

// HasBssId returns a boolean if a field has been set.
func (o *TnapId) HasBssId() bool {
	if o != nil && !isNil(o.BssId) {
		return true
	}

	return false
}

// SetBssId gets a reference to the given string and assigns it to the BssId field.
func (o *TnapId) SetBssId(v string) {
	o.BssId = &v
}

// GetCivicAddress returns the CivicAddress field value if set, zero value otherwise.
func (o *TnapId) GetCivicAddress() string {
	if o == nil || isNil(o.CivicAddress) {
		var ret string
		return ret
	}
	return *o.CivicAddress
}

// GetCivicAddressOk returns a tuple with the CivicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TnapId) GetCivicAddressOk() (*string, bool) {
	if o == nil || isNil(o.CivicAddress) {
		return nil, false
	}
	return o.CivicAddress, true
}

// HasCivicAddress returns a boolean if a field has been set.
func (o *TnapId) HasCivicAddress() bool {
	if o != nil && !isNil(o.CivicAddress) {
		return true
	}

	return false
}

// SetCivicAddress gets a reference to the given string and assigns it to the CivicAddress field.
func (o *TnapId) SetCivicAddress(v string) {
	o.CivicAddress = &v
}

func (o TnapId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TnapId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SsId) {
		toSerialize["ssId"] = o.SsId
	}
	if !isNil(o.BssId) {
		toSerialize["bssId"] = o.BssId
	}
	if !isNil(o.CivicAddress) {
		toSerialize["civicAddress"] = o.CivicAddress
	}
	return toSerialize, nil
}

type NullableTnapId struct {
	value *TnapId
	isSet bool
}

func (v NullableTnapId) Get() *TnapId {
	return v.value
}

func (v *NullableTnapId) Set(val *TnapId) {
	v.value = val
	v.isSet = true
}

func (v NullableTnapId) IsSet() bool {
	return v.isSet
}

func (v *NullableTnapId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTnapId(val *TnapId) *NullableTnapId {
	return &NullableTnapId{value: val, isSet: true}
}

func (v NullableTnapId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTnapId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



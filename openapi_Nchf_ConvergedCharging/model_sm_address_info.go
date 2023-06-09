/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the SMAddressInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SMAddressInfo{}

// SMAddressInfo struct for SMAddressInfo
type SMAddressInfo struct {
	SMaddressType   *SMAddressType   `json:"sMaddressType,omitempty"`
	SMaddressData   *string          `json:"sMaddressData,omitempty"`
	SMaddressDomain *SMAddressDomain `json:"sMaddressDomain,omitempty"`
}

// NewSMAddressInfo instantiates a new SMAddressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMAddressInfo() *SMAddressInfo {
	this := SMAddressInfo{}
	return &this
}

// NewSMAddressInfoWithDefaults instantiates a new SMAddressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMAddressInfoWithDefaults() *SMAddressInfo {
	this := SMAddressInfo{}
	return &this
}

// GetSMaddressType returns the SMaddressType field value if set, zero value otherwise.
func (o *SMAddressInfo) GetSMaddressType() SMAddressType {
	if o == nil || IsNil(o.SMaddressType) {
		var ret SMAddressType
		return ret
	}
	return *o.SMaddressType
}

// GetSMaddressTypeOk returns a tuple with the SMaddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMAddressInfo) GetSMaddressTypeOk() (*SMAddressType, bool) {
	if o == nil || IsNil(o.SMaddressType) {
		return nil, false
	}
	return o.SMaddressType, true
}

// HasSMaddressType returns a boolean if a field has been set.
func (o *SMAddressInfo) HasSMaddressType() bool {
	if o != nil && !IsNil(o.SMaddressType) {
		return true
	}

	return false
}

// SetSMaddressType gets a reference to the given SMAddressType and assigns it to the SMaddressType field.
func (o *SMAddressInfo) SetSMaddressType(v SMAddressType) {
	o.SMaddressType = &v
}

// GetSMaddressData returns the SMaddressData field value if set, zero value otherwise.
func (o *SMAddressInfo) GetSMaddressData() string {
	if o == nil || IsNil(o.SMaddressData) {
		var ret string
		return ret
	}
	return *o.SMaddressData
}

// GetSMaddressDataOk returns a tuple with the SMaddressData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMAddressInfo) GetSMaddressDataOk() (*string, bool) {
	if o == nil || IsNil(o.SMaddressData) {
		return nil, false
	}
	return o.SMaddressData, true
}

// HasSMaddressData returns a boolean if a field has been set.
func (o *SMAddressInfo) HasSMaddressData() bool {
	if o != nil && !IsNil(o.SMaddressData) {
		return true
	}

	return false
}

// SetSMaddressData gets a reference to the given string and assigns it to the SMaddressData field.
func (o *SMAddressInfo) SetSMaddressData(v string) {
	o.SMaddressData = &v
}

// GetSMaddressDomain returns the SMaddressDomain field value if set, zero value otherwise.
func (o *SMAddressInfo) GetSMaddressDomain() SMAddressDomain {
	if o == nil || IsNil(o.SMaddressDomain) {
		var ret SMAddressDomain
		return ret
	}
	return *o.SMaddressDomain
}

// GetSMaddressDomainOk returns a tuple with the SMaddressDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMAddressInfo) GetSMaddressDomainOk() (*SMAddressDomain, bool) {
	if o == nil || IsNil(o.SMaddressDomain) {
		return nil, false
	}
	return o.SMaddressDomain, true
}

// HasSMaddressDomain returns a boolean if a field has been set.
func (o *SMAddressInfo) HasSMaddressDomain() bool {
	if o != nil && !IsNil(o.SMaddressDomain) {
		return true
	}

	return false
}

// SetSMaddressDomain gets a reference to the given SMAddressDomain and assigns it to the SMaddressDomain field.
func (o *SMAddressInfo) SetSMaddressDomain(v SMAddressDomain) {
	o.SMaddressDomain = &v
}

func (o SMAddressInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SMAddressInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SMaddressType) {
		toSerialize["sMaddressType"] = o.SMaddressType
	}
	if !IsNil(o.SMaddressData) {
		toSerialize["sMaddressData"] = o.SMaddressData
	}
	if !IsNil(o.SMaddressDomain) {
		toSerialize["sMaddressDomain"] = o.SMaddressDomain
	}
	return toSerialize, nil
}

type NullableSMAddressInfo struct {
	value *SMAddressInfo
	isSet bool
}

func (v NullableSMAddressInfo) Get() *SMAddressInfo {
	return v.value
}

func (v *NullableSMAddressInfo) Set(val *SMAddressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSMAddressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSMAddressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMAddressInfo(val *SMAddressInfo) *NullableSMAddressInfo {
	return &NullableSMAddressInfo{value: val, isSet: true}
}

func (v NullableSMAddressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMAddressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

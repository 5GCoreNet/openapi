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

// checks if the RecipientAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientAddress{}

// RecipientAddress struct for RecipientAddress
type RecipientAddress struct {
	RecipientAddressInfo *SMAddressInfo `json:"recipientAddressInfo,omitempty"`
	SMaddresseeType *SMAddresseeType `json:"sMaddresseeType,omitempty"`
}

// NewRecipientAddress instantiates a new RecipientAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientAddress() *RecipientAddress {
	this := RecipientAddress{}
	return &this
}

// NewRecipientAddressWithDefaults instantiates a new RecipientAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientAddressWithDefaults() *RecipientAddress {
	this := RecipientAddress{}
	return &this
}

// GetRecipientAddressInfo returns the RecipientAddressInfo field value if set, zero value otherwise.
func (o *RecipientAddress) GetRecipientAddressInfo() SMAddressInfo {
	if o == nil || IsNil(o.RecipientAddressInfo) {
		var ret SMAddressInfo
		return ret
	}
	return *o.RecipientAddressInfo
}

// GetRecipientAddressInfoOk returns a tuple with the RecipientAddressInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientAddress) GetRecipientAddressInfoOk() (*SMAddressInfo, bool) {
	if o == nil || IsNil(o.RecipientAddressInfo) {
		return nil, false
	}
	return o.RecipientAddressInfo, true
}

// HasRecipientAddressInfo returns a boolean if a field has been set.
func (o *RecipientAddress) HasRecipientAddressInfo() bool {
	if o != nil && !IsNil(o.RecipientAddressInfo) {
		return true
	}

	return false
}

// SetRecipientAddressInfo gets a reference to the given SMAddressInfo and assigns it to the RecipientAddressInfo field.
func (o *RecipientAddress) SetRecipientAddressInfo(v SMAddressInfo) {
	o.RecipientAddressInfo = &v
}

// GetSMaddresseeType returns the SMaddresseeType field value if set, zero value otherwise.
func (o *RecipientAddress) GetSMaddresseeType() SMAddresseeType {
	if o == nil || IsNil(o.SMaddresseeType) {
		var ret SMAddresseeType
		return ret
	}
	return *o.SMaddresseeType
}

// GetSMaddresseeTypeOk returns a tuple with the SMaddresseeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientAddress) GetSMaddresseeTypeOk() (*SMAddresseeType, bool) {
	if o == nil || IsNil(o.SMaddresseeType) {
		return nil, false
	}
	return o.SMaddresseeType, true
}

// HasSMaddresseeType returns a boolean if a field has been set.
func (o *RecipientAddress) HasSMaddresseeType() bool {
	if o != nil && !IsNil(o.SMaddresseeType) {
		return true
	}

	return false
}

// SetSMaddresseeType gets a reference to the given SMAddresseeType and assigns it to the SMaddresseeType field.
func (o *RecipientAddress) SetSMaddresseeType(v SMAddresseeType) {
	o.SMaddresseeType = &v
}

func (o RecipientAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecipientAddressInfo) {
		toSerialize["recipientAddressInfo"] = o.RecipientAddressInfo
	}
	if !IsNil(o.SMaddresseeType) {
		toSerialize["sMaddresseeType"] = o.SMaddresseeType
	}
	return toSerialize, nil
}

type NullableRecipientAddress struct {
	value *RecipientAddress
	isSet bool
}

func (v NullableRecipientAddress) Get() *RecipientAddress {
	return v.value
}

func (v *NullableRecipientAddress) Set(val *RecipientAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientAddress(val *RecipientAddress) *NullableRecipientAddress {
	return &NullableRecipientAddress{value: val, isSet: true}
}

func (v NullableRecipientAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// checks if the MMRecipientInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MMRecipientInfo{}

// MMRecipientInfo struct for MMRecipientInfo
type MMRecipientInfo struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	RecipientSUPI *string `json:"recipientSUPI,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	RecipientGPSI *string `json:"recipientGPSI,omitempty"`
	RecipientOtherAddress []SMAddressInfo `json:"recipientOtherAddress,omitempty"`
}

// NewMMRecipientInfo instantiates a new MMRecipientInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMMRecipientInfo() *MMRecipientInfo {
	this := MMRecipientInfo{}
	return &this
}

// NewMMRecipientInfoWithDefaults instantiates a new MMRecipientInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMMRecipientInfoWithDefaults() *MMRecipientInfo {
	this := MMRecipientInfo{}
	return &this
}

// GetRecipientSUPI returns the RecipientSUPI field value if set, zero value otherwise.
func (o *MMRecipientInfo) GetRecipientSUPI() string {
	if o == nil || IsNil(o.RecipientSUPI) {
		var ret string
		return ret
	}
	return *o.RecipientSUPI
}

// GetRecipientSUPIOk returns a tuple with the RecipientSUPI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MMRecipientInfo) GetRecipientSUPIOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientSUPI) {
		return nil, false
	}
	return o.RecipientSUPI, true
}

// HasRecipientSUPI returns a boolean if a field has been set.
func (o *MMRecipientInfo) HasRecipientSUPI() bool {
	if o != nil && !IsNil(o.RecipientSUPI) {
		return true
	}

	return false
}

// SetRecipientSUPI gets a reference to the given string and assigns it to the RecipientSUPI field.
func (o *MMRecipientInfo) SetRecipientSUPI(v string) {
	o.RecipientSUPI = &v
}

// GetRecipientGPSI returns the RecipientGPSI field value if set, zero value otherwise.
func (o *MMRecipientInfo) GetRecipientGPSI() string {
	if o == nil || IsNil(o.RecipientGPSI) {
		var ret string
		return ret
	}
	return *o.RecipientGPSI
}

// GetRecipientGPSIOk returns a tuple with the RecipientGPSI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MMRecipientInfo) GetRecipientGPSIOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientGPSI) {
		return nil, false
	}
	return o.RecipientGPSI, true
}

// HasRecipientGPSI returns a boolean if a field has been set.
func (o *MMRecipientInfo) HasRecipientGPSI() bool {
	if o != nil && !IsNil(o.RecipientGPSI) {
		return true
	}

	return false
}

// SetRecipientGPSI gets a reference to the given string and assigns it to the RecipientGPSI field.
func (o *MMRecipientInfo) SetRecipientGPSI(v string) {
	o.RecipientGPSI = &v
}

// GetRecipientOtherAddress returns the RecipientOtherAddress field value if set, zero value otherwise.
func (o *MMRecipientInfo) GetRecipientOtherAddress() []SMAddressInfo {
	if o == nil || IsNil(o.RecipientOtherAddress) {
		var ret []SMAddressInfo
		return ret
	}
	return o.RecipientOtherAddress
}

// GetRecipientOtherAddressOk returns a tuple with the RecipientOtherAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MMRecipientInfo) GetRecipientOtherAddressOk() ([]SMAddressInfo, bool) {
	if o == nil || IsNil(o.RecipientOtherAddress) {
		return nil, false
	}
	return o.RecipientOtherAddress, true
}

// HasRecipientOtherAddress returns a boolean if a field has been set.
func (o *MMRecipientInfo) HasRecipientOtherAddress() bool {
	if o != nil && !IsNil(o.RecipientOtherAddress) {
		return true
	}

	return false
}

// SetRecipientOtherAddress gets a reference to the given []SMAddressInfo and assigns it to the RecipientOtherAddress field.
func (o *MMRecipientInfo) SetRecipientOtherAddress(v []SMAddressInfo) {
	o.RecipientOtherAddress = v
}

func (o MMRecipientInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MMRecipientInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecipientSUPI) {
		toSerialize["recipientSUPI"] = o.RecipientSUPI
	}
	if !IsNil(o.RecipientGPSI) {
		toSerialize["recipientGPSI"] = o.RecipientGPSI
	}
	if !IsNil(o.RecipientOtherAddress) {
		toSerialize["recipientOtherAddress"] = o.RecipientOtherAddress
	}
	return toSerialize, nil
}

type NullableMMRecipientInfo struct {
	value *MMRecipientInfo
	isSet bool
}

func (v NullableMMRecipientInfo) Get() *MMRecipientInfo {
	return v.value
}

func (v *NullableMMRecipientInfo) Set(val *MMRecipientInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMMRecipientInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMMRecipientInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMMRecipientInfo(val *MMRecipientInfo) *NullableMMRecipientInfo {
	return &NullableMMRecipientInfo{value: val, isSet: true}
}

func (v NullableMMRecipientInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMMRecipientInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


/*
EES EEL Managed ACR Service

EES EEL Managed ACR Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EELManagedACR

import (
	"encoding/json"
)

// checks if the EELACRResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EELACRResp{}

// EELACRResp Represents the feedback of the EES on EEL Managed ACR request.
type EELACRResp struct {
	// string providing an URI formatted according to IETF RFC 3986.
	AppCtxtStoreAddr *string `json:"appCtxtStoreAddr,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewEELACRResp instantiates a new EELACRResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEELACRResp() *EELACRResp {
	this := EELACRResp{}
	return &this
}

// NewEELACRRespWithDefaults instantiates a new EELACRResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEELACRRespWithDefaults() *EELACRResp {
	this := EELACRResp{}
	return &this
}

// GetAppCtxtStoreAddr returns the AppCtxtStoreAddr field value if set, zero value otherwise.
func (o *EELACRResp) GetAppCtxtStoreAddr() string {
	if o == nil || IsNil(o.AppCtxtStoreAddr) {
		var ret string
		return ret
	}
	return *o.AppCtxtStoreAddr
}

// GetAppCtxtStoreAddrOk returns a tuple with the AppCtxtStoreAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EELACRResp) GetAppCtxtStoreAddrOk() (*string, bool) {
	if o == nil || IsNil(o.AppCtxtStoreAddr) {
		return nil, false
	}
	return o.AppCtxtStoreAddr, true
}

// HasAppCtxtStoreAddr returns a boolean if a field has been set.
func (o *EELACRResp) HasAppCtxtStoreAddr() bool {
	if o != nil && !IsNil(o.AppCtxtStoreAddr) {
		return true
	}

	return false
}

// SetAppCtxtStoreAddr gets a reference to the given string and assigns it to the AppCtxtStoreAddr field.
func (o *EELACRResp) SetAppCtxtStoreAddr(v string) {
	o.AppCtxtStoreAddr = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *EELACRResp) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EELACRResp) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *EELACRResp) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *EELACRResp) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o EELACRResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EELACRResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppCtxtStoreAddr) {
		toSerialize["appCtxtStoreAddr"] = o.AppCtxtStoreAddr
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableEELACRResp struct {
	value *EELACRResp
	isSet bool
}

func (v NullableEELACRResp) Get() *EELACRResp {
	return v.value
}

func (v *NullableEELACRResp) Set(val *EELACRResp) {
	v.value = val
	v.isSet = true
}

func (v NullableEELACRResp) IsSet() bool {
	return v.isSet
}

func (v *NullableEELACRResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEELACRResp(val *EELACRResp) *NullableEELACRResp {
	return &NullableEELACRResp{value: val, isSet: true}
}

func (v NullableEELACRResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEELACRResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



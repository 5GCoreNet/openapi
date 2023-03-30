/*
AEF_Security_API

API for AEF security management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AEF_Security_API

import (
	"encoding/json"
)

// checks if the CheckAuthenticationReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckAuthenticationReq{}

// CheckAuthenticationReq Represents authentication check request data.
type CheckAuthenticationReq struct {
	// API invoker ID assigned by the CAPIF core function to the API invoker while on-boarding the API invoker. 
	ApiInvokerId string `json:"apiInvokerId"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures"`
}

// NewCheckAuthenticationReq instantiates a new CheckAuthenticationReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckAuthenticationReq(apiInvokerId string, supportedFeatures string) *CheckAuthenticationReq {
	this := CheckAuthenticationReq{}
	this.ApiInvokerId = apiInvokerId
	this.SupportedFeatures = supportedFeatures
	return &this
}

// NewCheckAuthenticationReqWithDefaults instantiates a new CheckAuthenticationReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckAuthenticationReqWithDefaults() *CheckAuthenticationReq {
	this := CheckAuthenticationReq{}
	return &this
}

// GetApiInvokerId returns the ApiInvokerId field value
func (o *CheckAuthenticationReq) GetApiInvokerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiInvokerId
}

// GetApiInvokerIdOk returns a tuple with the ApiInvokerId field value
// and a boolean to check if the value has been set.
func (o *CheckAuthenticationReq) GetApiInvokerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiInvokerId, true
}

// SetApiInvokerId sets field value
func (o *CheckAuthenticationReq) SetApiInvokerId(v string) {
	o.ApiInvokerId = v
}

// GetSupportedFeatures returns the SupportedFeatures field value
func (o *CheckAuthenticationReq) GetSupportedFeatures() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value
// and a boolean to check if the value has been set.
func (o *CheckAuthenticationReq) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedFeatures, true
}

// SetSupportedFeatures sets field value
func (o *CheckAuthenticationReq) SetSupportedFeatures(v string) {
	o.SupportedFeatures = v
}

func (o CheckAuthenticationReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckAuthenticationReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiInvokerId"] = o.ApiInvokerId
	toSerialize["supportedFeatures"] = o.SupportedFeatures
	return toSerialize, nil
}

type NullableCheckAuthenticationReq struct {
	value *CheckAuthenticationReq
	isSet bool
}

func (v NullableCheckAuthenticationReq) Get() *CheckAuthenticationReq {
	return v.value
}

func (v *NullableCheckAuthenticationReq) Set(val *CheckAuthenticationReq) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckAuthenticationReq) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckAuthenticationReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckAuthenticationReq(val *CheckAuthenticationReq) *NullableCheckAuthenticationReq {
	return &NullableCheckAuthenticationReq{value: val, isSet: true}
}

func (v NullableCheckAuthenticationReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckAuthenticationReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



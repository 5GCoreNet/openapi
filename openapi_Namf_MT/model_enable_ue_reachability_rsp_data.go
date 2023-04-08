/*
Namf_MT

AMF Mobile Terminated Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MT

import (
	"encoding/json"
)

// checks if the EnableUeReachabilityRspData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableUeReachabilityRspData{}

// EnableUeReachabilityRspData Data within the Enable UE Reachability Response
type EnableUeReachabilityRspData struct {
	Reachability UeReachability `json:"reachability"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewEnableUeReachabilityRspData instantiates a new EnableUeReachabilityRspData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableUeReachabilityRspData(reachability UeReachability) *EnableUeReachabilityRspData {
	this := EnableUeReachabilityRspData{}
	this.Reachability = reachability
	return &this
}

// NewEnableUeReachabilityRspDataWithDefaults instantiates a new EnableUeReachabilityRspData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableUeReachabilityRspDataWithDefaults() *EnableUeReachabilityRspData {
	this := EnableUeReachabilityRspData{}
	return &this
}

// GetReachability returns the Reachability field value
func (o *EnableUeReachabilityRspData) GetReachability() UeReachability {
	if o == nil {
		var ret UeReachability
		return ret
	}

	return o.Reachability
}

// GetReachabilityOk returns a tuple with the Reachability field value
// and a boolean to check if the value has been set.
func (o *EnableUeReachabilityRspData) GetReachabilityOk() (*UeReachability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reachability, true
}

// SetReachability sets field value
func (o *EnableUeReachabilityRspData) SetReachability(v UeReachability) {
	o.Reachability = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *EnableUeReachabilityRspData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableUeReachabilityRspData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *EnableUeReachabilityRspData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *EnableUeReachabilityRspData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o EnableUeReachabilityRspData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableUeReachabilityRspData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reachability"] = o.Reachability
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableEnableUeReachabilityRspData struct {
	value *EnableUeReachabilityRspData
	isSet bool
}

func (v NullableEnableUeReachabilityRspData) Get() *EnableUeReachabilityRspData {
	return v.value
}

func (v *NullableEnableUeReachabilityRspData) Set(val *EnableUeReachabilityRspData) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableUeReachabilityRspData) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableUeReachabilityRspData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableUeReachabilityRspData(val *EnableUeReachabilityRspData) *NullableEnableUeReachabilityRspData {
	return &NullableEnableUeReachabilityRspData{value: val, isSet: true}
}

func (v NullableEnableUeReachabilityRspData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableUeReachabilityRspData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



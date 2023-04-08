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

// checks if the EnableUeReachabilityReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableUeReachabilityReqData{}

// EnableUeReachabilityReqData Data within the Enable UE Reachability Request
type EnableUeReachabilityReqData struct {
	Reachability UeReachability `json:"reachability"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	OldGuami *Guami `json:"oldGuami,omitempty"`
	ExtBufSupport *bool `json:"extBufSupport,omitempty"`
}

// NewEnableUeReachabilityReqData instantiates a new EnableUeReachabilityReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableUeReachabilityReqData(reachability UeReachability) *EnableUeReachabilityReqData {
	this := EnableUeReachabilityReqData{}
	this.Reachability = reachability
	var extBufSupport bool = false
	this.ExtBufSupport = &extBufSupport
	return &this
}

// NewEnableUeReachabilityReqDataWithDefaults instantiates a new EnableUeReachabilityReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableUeReachabilityReqDataWithDefaults() *EnableUeReachabilityReqData {
	this := EnableUeReachabilityReqData{}
	var extBufSupport bool = false
	this.ExtBufSupport = &extBufSupport
	return &this
}

// GetReachability returns the Reachability field value
func (o *EnableUeReachabilityReqData) GetReachability() UeReachability {
	if o == nil {
		var ret UeReachability
		return ret
	}

	return o.Reachability
}

// GetReachabilityOk returns a tuple with the Reachability field value
// and a boolean to check if the value has been set.
func (o *EnableUeReachabilityReqData) GetReachabilityOk() (*UeReachability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reachability, true
}

// SetReachability sets field value
func (o *EnableUeReachabilityReqData) SetReachability(v UeReachability) {
	o.Reachability = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *EnableUeReachabilityReqData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableUeReachabilityReqData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *EnableUeReachabilityReqData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *EnableUeReachabilityReqData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetOldGuami returns the OldGuami field value if set, zero value otherwise.
func (o *EnableUeReachabilityReqData) GetOldGuami() Guami {
	if o == nil || isNil(o.OldGuami) {
		var ret Guami
		return ret
	}
	return *o.OldGuami
}

// GetOldGuamiOk returns a tuple with the OldGuami field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableUeReachabilityReqData) GetOldGuamiOk() (*Guami, bool) {
	if o == nil || isNil(o.OldGuami) {
		return nil, false
	}
	return o.OldGuami, true
}

// HasOldGuami returns a boolean if a field has been set.
func (o *EnableUeReachabilityReqData) HasOldGuami() bool {
	if o != nil && !isNil(o.OldGuami) {
		return true
	}

	return false
}

// SetOldGuami gets a reference to the given Guami and assigns it to the OldGuami field.
func (o *EnableUeReachabilityReqData) SetOldGuami(v Guami) {
	o.OldGuami = &v
}

// GetExtBufSupport returns the ExtBufSupport field value if set, zero value otherwise.
func (o *EnableUeReachabilityReqData) GetExtBufSupport() bool {
	if o == nil || isNil(o.ExtBufSupport) {
		var ret bool
		return ret
	}
	return *o.ExtBufSupport
}

// GetExtBufSupportOk returns a tuple with the ExtBufSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableUeReachabilityReqData) GetExtBufSupportOk() (*bool, bool) {
	if o == nil || isNil(o.ExtBufSupport) {
		return nil, false
	}
	return o.ExtBufSupport, true
}

// HasExtBufSupport returns a boolean if a field has been set.
func (o *EnableUeReachabilityReqData) HasExtBufSupport() bool {
	if o != nil && !isNil(o.ExtBufSupport) {
		return true
	}

	return false
}

// SetExtBufSupport gets a reference to the given bool and assigns it to the ExtBufSupport field.
func (o *EnableUeReachabilityReqData) SetExtBufSupport(v bool) {
	o.ExtBufSupport = &v
}

func (o EnableUeReachabilityReqData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableUeReachabilityReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reachability"] = o.Reachability
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.OldGuami) {
		toSerialize["oldGuami"] = o.OldGuami
	}
	if !isNil(o.ExtBufSupport) {
		toSerialize["extBufSupport"] = o.ExtBufSupport
	}
	return toSerialize, nil
}

type NullableEnableUeReachabilityReqData struct {
	value *EnableUeReachabilityReqData
	isSet bool
}

func (v NullableEnableUeReachabilityReqData) Get() *EnableUeReachabilityReqData {
	return v.value
}

func (v *NullableEnableUeReachabilityReqData) Set(val *EnableUeReachabilityReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableUeReachabilityReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableUeReachabilityReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableUeReachabilityReqData(val *EnableUeReachabilityReqData) *NullableEnableUeReachabilityReqData {
	return &NullableEnableUeReachabilityReqData{value: val, isSet: true}
}

func (v NullableEnableUeReachabilityReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableUeReachabilityReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Neasdf_BaselineDNSPattern

EASDF Baseline DNS Pattern Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Neasdf_BaselineDNSPattern

import (
	"encoding/json"
)

// checks if the BaseDnsPatternCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseDnsPatternCreateData{}

// BaseDnsPatternCreateData Data in Baseline DNS Pattern Create request
type BaseDnsPatternCreateData struct {
	Label *string `json:"label,omitempty"`
	// map of baseline DNS message detection templates where a valid JSON string serves as key
	BaseDnsMdtList *map[string]BaselineDnsMdt `json:"baseDnsMdtList,omitempty"`
	// map of Baseline DNS action information Template where a valid JSON string serves as key
	BaseDnsAitList *map[string]BaselineDnsAit `json:"baseDnsAitList,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewBaseDnsPatternCreateData instantiates a new BaseDnsPatternCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseDnsPatternCreateData() *BaseDnsPatternCreateData {
	this := BaseDnsPatternCreateData{}
	return &this
}

// NewBaseDnsPatternCreateDataWithDefaults instantiates a new BaseDnsPatternCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseDnsPatternCreateDataWithDefaults() *BaseDnsPatternCreateData {
	this := BaseDnsPatternCreateData{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BaseDnsPatternCreateData) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsPatternCreateData) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BaseDnsPatternCreateData) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BaseDnsPatternCreateData) SetLabel(v string) {
	o.Label = &v
}

// GetBaseDnsMdtList returns the BaseDnsMdtList field value if set, zero value otherwise.
func (o *BaseDnsPatternCreateData) GetBaseDnsMdtList() map[string]BaselineDnsMdt {
	if o == nil || isNil(o.BaseDnsMdtList) {
		var ret map[string]BaselineDnsMdt
		return ret
	}
	return *o.BaseDnsMdtList
}

// GetBaseDnsMdtListOk returns a tuple with the BaseDnsMdtList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsPatternCreateData) GetBaseDnsMdtListOk() (*map[string]BaselineDnsMdt, bool) {
	if o == nil || isNil(o.BaseDnsMdtList) {
		return nil, false
	}
	return o.BaseDnsMdtList, true
}

// HasBaseDnsMdtList returns a boolean if a field has been set.
func (o *BaseDnsPatternCreateData) HasBaseDnsMdtList() bool {
	if o != nil && !isNil(o.BaseDnsMdtList) {
		return true
	}

	return false
}

// SetBaseDnsMdtList gets a reference to the given map[string]BaselineDnsMdt and assigns it to the BaseDnsMdtList field.
func (o *BaseDnsPatternCreateData) SetBaseDnsMdtList(v map[string]BaselineDnsMdt) {
	o.BaseDnsMdtList = &v
}

// GetBaseDnsAitList returns the BaseDnsAitList field value if set, zero value otherwise.
func (o *BaseDnsPatternCreateData) GetBaseDnsAitList() map[string]BaselineDnsAit {
	if o == nil || isNil(o.BaseDnsAitList) {
		var ret map[string]BaselineDnsAit
		return ret
	}
	return *o.BaseDnsAitList
}

// GetBaseDnsAitListOk returns a tuple with the BaseDnsAitList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsPatternCreateData) GetBaseDnsAitListOk() (*map[string]BaselineDnsAit, bool) {
	if o == nil || isNil(o.BaseDnsAitList) {
		return nil, false
	}
	return o.BaseDnsAitList, true
}

// HasBaseDnsAitList returns a boolean if a field has been set.
func (o *BaseDnsPatternCreateData) HasBaseDnsAitList() bool {
	if o != nil && !isNil(o.BaseDnsAitList) {
		return true
	}

	return false
}

// SetBaseDnsAitList gets a reference to the given map[string]BaselineDnsAit and assigns it to the BaseDnsAitList field.
func (o *BaseDnsPatternCreateData) SetBaseDnsAitList(v map[string]BaselineDnsAit) {
	o.BaseDnsAitList = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *BaseDnsPatternCreateData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsPatternCreateData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *BaseDnsPatternCreateData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *BaseDnsPatternCreateData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o BaseDnsPatternCreateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseDnsPatternCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.BaseDnsMdtList) {
		toSerialize["baseDnsMdtList"] = o.BaseDnsMdtList
	}
	if !isNil(o.BaseDnsAitList) {
		toSerialize["baseDnsAitList"] = o.BaseDnsAitList
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableBaseDnsPatternCreateData struct {
	value *BaseDnsPatternCreateData
	isSet bool
}

func (v NullableBaseDnsPatternCreateData) Get() *BaseDnsPatternCreateData {
	return v.value
}

func (v *NullableBaseDnsPatternCreateData) Set(val *BaseDnsPatternCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseDnsPatternCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseDnsPatternCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseDnsPatternCreateData(val *BaseDnsPatternCreateData) *NullableBaseDnsPatternCreateData {
	return &NullableBaseDnsPatternCreateData{value: val, isSet: true}
}

func (v NullableBaseDnsPatternCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseDnsPatternCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



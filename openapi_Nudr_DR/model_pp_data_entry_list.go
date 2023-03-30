/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the PpDataEntryList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PpDataEntryList{}

// PpDataEntryList Contains a list of the Provisioned Parameters entries
type PpDataEntryList struct {
	PpDataEntryList []PpDataEntry `json:"ppDataEntryList,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewPpDataEntryList instantiates a new PpDataEntryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPpDataEntryList() *PpDataEntryList {
	this := PpDataEntryList{}
	return &this
}

// NewPpDataEntryListWithDefaults instantiates a new PpDataEntryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPpDataEntryListWithDefaults() *PpDataEntryList {
	this := PpDataEntryList{}
	return &this
}

// GetPpDataEntryList returns the PpDataEntryList field value if set, zero value otherwise.
func (o *PpDataEntryList) GetPpDataEntryList() []PpDataEntry {
	if o == nil || IsNil(o.PpDataEntryList) {
		var ret []PpDataEntry
		return ret
	}
	return o.PpDataEntryList
}

// GetPpDataEntryListOk returns a tuple with the PpDataEntryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpDataEntryList) GetPpDataEntryListOk() ([]PpDataEntry, bool) {
	if o == nil || IsNil(o.PpDataEntryList) {
		return nil, false
	}
	return o.PpDataEntryList, true
}

// HasPpDataEntryList returns a boolean if a field has been set.
func (o *PpDataEntryList) HasPpDataEntryList() bool {
	if o != nil && !IsNil(o.PpDataEntryList) {
		return true
	}

	return false
}

// SetPpDataEntryList gets a reference to the given []PpDataEntry and assigns it to the PpDataEntryList field.
func (o *PpDataEntryList) SetPpDataEntryList(v []PpDataEntry) {
	o.PpDataEntryList = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *PpDataEntryList) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpDataEntryList) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *PpDataEntryList) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *PpDataEntryList) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o PpDataEntryList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PpDataEntryList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PpDataEntryList) {
		toSerialize["ppDataEntryList"] = o.PpDataEntryList
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullablePpDataEntryList struct {
	value *PpDataEntryList
	isSet bool
}

func (v NullablePpDataEntryList) Get() *PpDataEntryList {
	return v.value
}

func (v *NullablePpDataEntryList) Set(val *PpDataEntryList) {
	v.value = val
	v.isSet = true
}

func (v NullablePpDataEntryList) IsSet() bool {
	return v.isSet
}

func (v *NullablePpDataEntryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePpDataEntryList(val *PpDataEntryList) *NullablePpDataEntryList {
	return &NullablePpDataEntryList{value: val, isSet: true}
}

func (v NullablePpDataEntryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePpDataEntryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


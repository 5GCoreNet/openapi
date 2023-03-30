/*
Nucmf_UECapabilityManagement

Nucmf_UECapabilityManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nucmf_UERCM

import (
	"encoding/json"
	"os"
)

// checks if the RetrieveDictionaryEntry200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetrieveDictionaryEntry200Response{}

// RetrieveDictionaryEntry200Response struct for RetrieveDictionaryEntry200Response
type RetrieveDictionaryEntry200Response struct {
	JsonData *DicEntryData `json:"jsonData,omitempty"`
	BinaryDataUeRadioCapability5GS **os.File `json:"binaryDataUeRadioCapability5GS,omitempty"`
	BinaryDataUeRadioCapabilityEPS **os.File `json:"binaryDataUeRadioCapabilityEPS,omitempty"`
	BinaryDataUeRadioCap5GSForPaging **os.File `json:"binaryDataUeRadioCap5GSForPaging,omitempty"`
	BinaryDataUeRadioCapEPSForPaging **os.File `json:"binaryDataUeRadioCapEPSForPaging,omitempty"`
}

// NewRetrieveDictionaryEntry200Response instantiates a new RetrieveDictionaryEntry200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetrieveDictionaryEntry200Response() *RetrieveDictionaryEntry200Response {
	this := RetrieveDictionaryEntry200Response{}
	return &this
}

// NewRetrieveDictionaryEntry200ResponseWithDefaults instantiates a new RetrieveDictionaryEntry200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetrieveDictionaryEntry200ResponseWithDefaults() *RetrieveDictionaryEntry200Response {
	this := RetrieveDictionaryEntry200Response{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *RetrieveDictionaryEntry200Response) GetJsonData() DicEntryData {
	if o == nil || IsNil(o.JsonData) {
		var ret DicEntryData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveDictionaryEntry200Response) GetJsonDataOk() (*DicEntryData, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *RetrieveDictionaryEntry200Response) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given DicEntryData and assigns it to the JsonData field.
func (o *RetrieveDictionaryEntry200Response) SetJsonData(v DicEntryData) {
	o.JsonData = &v
}

// GetBinaryDataUeRadioCapability5GS returns the BinaryDataUeRadioCapability5GS field value if set, zero value otherwise.
func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCapability5GS() *os.File {
	if o == nil || IsNil(o.BinaryDataUeRadioCapability5GS) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataUeRadioCapability5GS
}

// GetBinaryDataUeRadioCapability5GSOk returns a tuple with the BinaryDataUeRadioCapability5GS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCapability5GSOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataUeRadioCapability5GS) {
		return nil, false
	}
	return o.BinaryDataUeRadioCapability5GS, true
}

// HasBinaryDataUeRadioCapability5GS returns a boolean if a field has been set.
func (o *RetrieveDictionaryEntry200Response) HasBinaryDataUeRadioCapability5GS() bool {
	if o != nil && !IsNil(o.BinaryDataUeRadioCapability5GS) {
		return true
	}

	return false
}

// SetBinaryDataUeRadioCapability5GS gets a reference to the given *os.File and assigns it to the BinaryDataUeRadioCapability5GS field.
func (o *RetrieveDictionaryEntry200Response) SetBinaryDataUeRadioCapability5GS(v *os.File) {
	o.BinaryDataUeRadioCapability5GS = &v
}

// GetBinaryDataUeRadioCapabilityEPS returns the BinaryDataUeRadioCapabilityEPS field value if set, zero value otherwise.
func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCapabilityEPS() *os.File {
	if o == nil || IsNil(o.BinaryDataUeRadioCapabilityEPS) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataUeRadioCapabilityEPS
}

// GetBinaryDataUeRadioCapabilityEPSOk returns a tuple with the BinaryDataUeRadioCapabilityEPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCapabilityEPSOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataUeRadioCapabilityEPS) {
		return nil, false
	}
	return o.BinaryDataUeRadioCapabilityEPS, true
}

// HasBinaryDataUeRadioCapabilityEPS returns a boolean if a field has been set.
func (o *RetrieveDictionaryEntry200Response) HasBinaryDataUeRadioCapabilityEPS() bool {
	if o != nil && !IsNil(o.BinaryDataUeRadioCapabilityEPS) {
		return true
	}

	return false
}

// SetBinaryDataUeRadioCapabilityEPS gets a reference to the given *os.File and assigns it to the BinaryDataUeRadioCapabilityEPS field.
func (o *RetrieveDictionaryEntry200Response) SetBinaryDataUeRadioCapabilityEPS(v *os.File) {
	o.BinaryDataUeRadioCapabilityEPS = &v
}

// GetBinaryDataUeRadioCap5GSForPaging returns the BinaryDataUeRadioCap5GSForPaging field value if set, zero value otherwise.
func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCap5GSForPaging() *os.File {
	if o == nil || IsNil(o.BinaryDataUeRadioCap5GSForPaging) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataUeRadioCap5GSForPaging
}

// GetBinaryDataUeRadioCap5GSForPagingOk returns a tuple with the BinaryDataUeRadioCap5GSForPaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCap5GSForPagingOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataUeRadioCap5GSForPaging) {
		return nil, false
	}
	return o.BinaryDataUeRadioCap5GSForPaging, true
}

// HasBinaryDataUeRadioCap5GSForPaging returns a boolean if a field has been set.
func (o *RetrieveDictionaryEntry200Response) HasBinaryDataUeRadioCap5GSForPaging() bool {
	if o != nil && !IsNil(o.BinaryDataUeRadioCap5GSForPaging) {
		return true
	}

	return false
}

// SetBinaryDataUeRadioCap5GSForPaging gets a reference to the given *os.File and assigns it to the BinaryDataUeRadioCap5GSForPaging field.
func (o *RetrieveDictionaryEntry200Response) SetBinaryDataUeRadioCap5GSForPaging(v *os.File) {
	o.BinaryDataUeRadioCap5GSForPaging = &v
}

// GetBinaryDataUeRadioCapEPSForPaging returns the BinaryDataUeRadioCapEPSForPaging field value if set, zero value otherwise.
func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCapEPSForPaging() *os.File {
	if o == nil || IsNil(o.BinaryDataUeRadioCapEPSForPaging) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataUeRadioCapEPSForPaging
}

// GetBinaryDataUeRadioCapEPSForPagingOk returns a tuple with the BinaryDataUeRadioCapEPSForPaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCapEPSForPagingOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataUeRadioCapEPSForPaging) {
		return nil, false
	}
	return o.BinaryDataUeRadioCapEPSForPaging, true
}

// HasBinaryDataUeRadioCapEPSForPaging returns a boolean if a field has been set.
func (o *RetrieveDictionaryEntry200Response) HasBinaryDataUeRadioCapEPSForPaging() bool {
	if o != nil && !IsNil(o.BinaryDataUeRadioCapEPSForPaging) {
		return true
	}

	return false
}

// SetBinaryDataUeRadioCapEPSForPaging gets a reference to the given *os.File and assigns it to the BinaryDataUeRadioCapEPSForPaging field.
func (o *RetrieveDictionaryEntry200Response) SetBinaryDataUeRadioCapEPSForPaging(v *os.File) {
	o.BinaryDataUeRadioCapEPSForPaging = &v
}

func (o RetrieveDictionaryEntry200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetrieveDictionaryEntry200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !IsNil(o.BinaryDataUeRadioCapability5GS) {
		toSerialize["binaryDataUeRadioCapability5GS"] = o.BinaryDataUeRadioCapability5GS
	}
	if !IsNil(o.BinaryDataUeRadioCapabilityEPS) {
		toSerialize["binaryDataUeRadioCapabilityEPS"] = o.BinaryDataUeRadioCapabilityEPS
	}
	if !IsNil(o.BinaryDataUeRadioCap5GSForPaging) {
		toSerialize["binaryDataUeRadioCap5GSForPaging"] = o.BinaryDataUeRadioCap5GSForPaging
	}
	if !IsNil(o.BinaryDataUeRadioCapEPSForPaging) {
		toSerialize["binaryDataUeRadioCapEPSForPaging"] = o.BinaryDataUeRadioCapEPSForPaging
	}
	return toSerialize, nil
}

type NullableRetrieveDictionaryEntry200Response struct {
	value *RetrieveDictionaryEntry200Response
	isSet bool
}

func (v NullableRetrieveDictionaryEntry200Response) Get() *RetrieveDictionaryEntry200Response {
	return v.value
}

func (v *NullableRetrieveDictionaryEntry200Response) Set(val *RetrieveDictionaryEntry200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrieveDictionaryEntry200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrieveDictionaryEntry200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrieveDictionaryEntry200Response(val *RetrieveDictionaryEntry200Response) *NullableRetrieveDictionaryEntry200Response {
	return &NullableRetrieveDictionaryEntry200Response{value: val, isSet: true}
}

func (v NullableRetrieveDictionaryEntry200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetrieveDictionaryEntry200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



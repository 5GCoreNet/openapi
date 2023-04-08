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

// checks if the CreateDictionaryEntryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDictionaryEntryRequest{}

// CreateDictionaryEntryRequest struct for CreateDictionaryEntryRequest
type CreateDictionaryEntryRequest struct {
	JsonData *DicEntryCreateData `json:"jsonData,omitempty"`
	BinaryDataUeRadioCapability5GS *os.File `json:"binaryDataUeRadioCapability5GS,omitempty"`
	BinaryDataUeRadioCapabilityEPS *os.File `json:"binaryDataUeRadioCapabilityEPS,omitempty"`
	BinaryDataUeRadioCap5GSForPaging *os.File `json:"binaryDataUeRadioCap5GSForPaging,omitempty"`
	BinaryDataUeRadioCapEPSForPaging *os.File `json:"binaryDataUeRadioCapEPSForPaging,omitempty"`
}

// NewCreateDictionaryEntryRequest instantiates a new CreateDictionaryEntryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDictionaryEntryRequest() *CreateDictionaryEntryRequest {
	this := CreateDictionaryEntryRequest{}
	return &this
}

// NewCreateDictionaryEntryRequestWithDefaults instantiates a new CreateDictionaryEntryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDictionaryEntryRequestWithDefaults() *CreateDictionaryEntryRequest {
	this := CreateDictionaryEntryRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *CreateDictionaryEntryRequest) GetJsonData() DicEntryCreateData {
	if o == nil || isNil(o.JsonData) {
		var ret DicEntryCreateData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDictionaryEntryRequest) GetJsonDataOk() (*DicEntryCreateData, bool) {
	if o == nil || isNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *CreateDictionaryEntryRequest) HasJsonData() bool {
	if o != nil && !isNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given DicEntryCreateData and assigns it to the JsonData field.
func (o *CreateDictionaryEntryRequest) SetJsonData(v DicEntryCreateData) {
	o.JsonData = &v
}

// GetBinaryDataUeRadioCapability5GS returns the BinaryDataUeRadioCapability5GS field value if set, zero value otherwise.
func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCapability5GS() os.File {
	if o == nil || isNil(o.BinaryDataUeRadioCapability5GS) {
		var ret os.File
		return ret
	}
	return *o.BinaryDataUeRadioCapability5GS
}

// GetBinaryDataUeRadioCapability5GSOk returns a tuple with the BinaryDataUeRadioCapability5GS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCapability5GSOk() (*os.File, bool) {
	if o == nil || isNil(o.BinaryDataUeRadioCapability5GS) {
		return nil, false
	}
	return o.BinaryDataUeRadioCapability5GS, true
}

// HasBinaryDataUeRadioCapability5GS returns a boolean if a field has been set.
func (o *CreateDictionaryEntryRequest) HasBinaryDataUeRadioCapability5GS() bool {
	if o != nil && !isNil(o.BinaryDataUeRadioCapability5GS) {
		return true
	}

	return false
}

// SetBinaryDataUeRadioCapability5GS gets a reference to the given os.File and assigns it to the BinaryDataUeRadioCapability5GS field.
func (o *CreateDictionaryEntryRequest) SetBinaryDataUeRadioCapability5GS(v os.File) {
	o.BinaryDataUeRadioCapability5GS = &v
}

// GetBinaryDataUeRadioCapabilityEPS returns the BinaryDataUeRadioCapabilityEPS field value if set, zero value otherwise.
func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCapabilityEPS() os.File {
	if o == nil || isNil(o.BinaryDataUeRadioCapabilityEPS) {
		var ret os.File
		return ret
	}
	return *o.BinaryDataUeRadioCapabilityEPS
}

// GetBinaryDataUeRadioCapabilityEPSOk returns a tuple with the BinaryDataUeRadioCapabilityEPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCapabilityEPSOk() (*os.File, bool) {
	if o == nil || isNil(o.BinaryDataUeRadioCapabilityEPS) {
		return nil, false
	}
	return o.BinaryDataUeRadioCapabilityEPS, true
}

// HasBinaryDataUeRadioCapabilityEPS returns a boolean if a field has been set.
func (o *CreateDictionaryEntryRequest) HasBinaryDataUeRadioCapabilityEPS() bool {
	if o != nil && !isNil(o.BinaryDataUeRadioCapabilityEPS) {
		return true
	}

	return false
}

// SetBinaryDataUeRadioCapabilityEPS gets a reference to the given os.File and assigns it to the BinaryDataUeRadioCapabilityEPS field.
func (o *CreateDictionaryEntryRequest) SetBinaryDataUeRadioCapabilityEPS(v os.File) {
	o.BinaryDataUeRadioCapabilityEPS = &v
}

// GetBinaryDataUeRadioCap5GSForPaging returns the BinaryDataUeRadioCap5GSForPaging field value if set, zero value otherwise.
func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCap5GSForPaging() os.File {
	if o == nil || isNil(o.BinaryDataUeRadioCap5GSForPaging) {
		var ret os.File
		return ret
	}
	return *o.BinaryDataUeRadioCap5GSForPaging
}

// GetBinaryDataUeRadioCap5GSForPagingOk returns a tuple with the BinaryDataUeRadioCap5GSForPaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCap5GSForPagingOk() (*os.File, bool) {
	if o == nil || isNil(o.BinaryDataUeRadioCap5GSForPaging) {
		return nil, false
	}
	return o.BinaryDataUeRadioCap5GSForPaging, true
}

// HasBinaryDataUeRadioCap5GSForPaging returns a boolean if a field has been set.
func (o *CreateDictionaryEntryRequest) HasBinaryDataUeRadioCap5GSForPaging() bool {
	if o != nil && !isNil(o.BinaryDataUeRadioCap5GSForPaging) {
		return true
	}

	return false
}

// SetBinaryDataUeRadioCap5GSForPaging gets a reference to the given os.File and assigns it to the BinaryDataUeRadioCap5GSForPaging field.
func (o *CreateDictionaryEntryRequest) SetBinaryDataUeRadioCap5GSForPaging(v os.File) {
	o.BinaryDataUeRadioCap5GSForPaging = &v
}

// GetBinaryDataUeRadioCapEPSForPaging returns the BinaryDataUeRadioCapEPSForPaging field value if set, zero value otherwise.
func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCapEPSForPaging() os.File {
	if o == nil || isNil(o.BinaryDataUeRadioCapEPSForPaging) {
		var ret os.File
		return ret
	}
	return *o.BinaryDataUeRadioCapEPSForPaging
}

// GetBinaryDataUeRadioCapEPSForPagingOk returns a tuple with the BinaryDataUeRadioCapEPSForPaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCapEPSForPagingOk() (*os.File, bool) {
	if o == nil || isNil(o.BinaryDataUeRadioCapEPSForPaging) {
		return nil, false
	}
	return o.BinaryDataUeRadioCapEPSForPaging, true
}

// HasBinaryDataUeRadioCapEPSForPaging returns a boolean if a field has been set.
func (o *CreateDictionaryEntryRequest) HasBinaryDataUeRadioCapEPSForPaging() bool {
	if o != nil && !isNil(o.BinaryDataUeRadioCapEPSForPaging) {
		return true
	}

	return false
}

// SetBinaryDataUeRadioCapEPSForPaging gets a reference to the given os.File and assigns it to the BinaryDataUeRadioCapEPSForPaging field.
func (o *CreateDictionaryEntryRequest) SetBinaryDataUeRadioCapEPSForPaging(v os.File) {
	o.BinaryDataUeRadioCapEPSForPaging = &v
}

func (o CreateDictionaryEntryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDictionaryEntryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !isNil(o.BinaryDataUeRadioCapability5GS) {
		toSerialize["binaryDataUeRadioCapability5GS"] = o.BinaryDataUeRadioCapability5GS
	}
	if !isNil(o.BinaryDataUeRadioCapabilityEPS) {
		toSerialize["binaryDataUeRadioCapabilityEPS"] = o.BinaryDataUeRadioCapabilityEPS
	}
	if !isNil(o.BinaryDataUeRadioCap5GSForPaging) {
		toSerialize["binaryDataUeRadioCap5GSForPaging"] = o.BinaryDataUeRadioCap5GSForPaging
	}
	if !isNil(o.BinaryDataUeRadioCapEPSForPaging) {
		toSerialize["binaryDataUeRadioCapEPSForPaging"] = o.BinaryDataUeRadioCapEPSForPaging
	}
	return toSerialize, nil
}

type NullableCreateDictionaryEntryRequest struct {
	value *CreateDictionaryEntryRequest
	isSet bool
}

func (v NullableCreateDictionaryEntryRequest) Get() *CreateDictionaryEntryRequest {
	return v.value
}

func (v *NullableCreateDictionaryEntryRequest) Set(val *CreateDictionaryEntryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDictionaryEntryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDictionaryEntryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDictionaryEntryRequest(val *CreateDictionaryEntryRequest) *NullableCreateDictionaryEntryRequest {
	return &NullableCreateDictionaryEntryRequest{value: val, isSet: true}
}

func (v NullableCreateDictionaryEntryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDictionaryEntryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



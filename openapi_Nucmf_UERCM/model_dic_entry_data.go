/*
Nucmf_UECapabilityManagement

Nucmf_UECapabilityManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nucmf_UERCM

import (
	"encoding/json"
)

// checks if the DicEntryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DicEntryData{}

// DicEntryData A dictionary entry for a UE radio capability ID in the UCMF
type DicEntryData struct {
	DicEntryId *int32 `json:"dicEntryId,omitempty"`
	// Type Allocation Code (TAC) of the UE, comprising the initial eight-digit portion of the 15-digit IMEI and 16-digit IMEISV codes. See clause 6.2 of 3GPP TS 23.003. 
	TypeAllocationCode string `json:"typeAllocationCode"`
	// string with format 'bytes' as defined in OpenAPI
	PlmnAssiUeRadioCapId *string `json:"plmnAssiUeRadioCapId,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	ManAssiUeRadioCapId *string `json:"manAssiUeRadioCapId,omitempty"`
	UeRadioCapability5GS *RefToBinaryData `json:"ueRadioCapability5GS,omitempty"`
	UeRadioCapabilityEPS *RefToBinaryData `json:"ueRadioCapabilityEPS,omitempty"`
	UeRadioCap5GSForPaging *RefToBinaryData `json:"ueRadioCap5GSForPaging,omitempty"`
	UeRadioCapEPSForPaging *RefToBinaryData `json:"ueRadioCapEPSForPaging,omitempty"`
}

// NewDicEntryData instantiates a new DicEntryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDicEntryData(typeAllocationCode string) *DicEntryData {
	this := DicEntryData{}
	this.TypeAllocationCode = typeAllocationCode
	return &this
}

// NewDicEntryDataWithDefaults instantiates a new DicEntryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDicEntryDataWithDefaults() *DicEntryData {
	this := DicEntryData{}
	return &this
}

// GetDicEntryId returns the DicEntryId field value if set, zero value otherwise.
func (o *DicEntryData) GetDicEntryId() int32 {
	if o == nil || IsNil(o.DicEntryId) {
		var ret int32
		return ret
	}
	return *o.DicEntryId
}

// GetDicEntryIdOk returns a tuple with the DicEntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DicEntryData) GetDicEntryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DicEntryId) {
		return nil, false
	}
	return o.DicEntryId, true
}

// HasDicEntryId returns a boolean if a field has been set.
func (o *DicEntryData) HasDicEntryId() bool {
	if o != nil && !IsNil(o.DicEntryId) {
		return true
	}

	return false
}

// SetDicEntryId gets a reference to the given int32 and assigns it to the DicEntryId field.
func (o *DicEntryData) SetDicEntryId(v int32) {
	o.DicEntryId = &v
}

// GetTypeAllocationCode returns the TypeAllocationCode field value
func (o *DicEntryData) GetTypeAllocationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeAllocationCode
}

// GetTypeAllocationCodeOk returns a tuple with the TypeAllocationCode field value
// and a boolean to check if the value has been set.
func (o *DicEntryData) GetTypeAllocationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeAllocationCode, true
}

// SetTypeAllocationCode sets field value
func (o *DicEntryData) SetTypeAllocationCode(v string) {
	o.TypeAllocationCode = v
}

// GetPlmnAssiUeRadioCapId returns the PlmnAssiUeRadioCapId field value if set, zero value otherwise.
func (o *DicEntryData) GetPlmnAssiUeRadioCapId() string {
	if o == nil || IsNil(o.PlmnAssiUeRadioCapId) {
		var ret string
		return ret
	}
	return *o.PlmnAssiUeRadioCapId
}

// GetPlmnAssiUeRadioCapIdOk returns a tuple with the PlmnAssiUeRadioCapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DicEntryData) GetPlmnAssiUeRadioCapIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlmnAssiUeRadioCapId) {
		return nil, false
	}
	return o.PlmnAssiUeRadioCapId, true
}

// HasPlmnAssiUeRadioCapId returns a boolean if a field has been set.
func (o *DicEntryData) HasPlmnAssiUeRadioCapId() bool {
	if o != nil && !IsNil(o.PlmnAssiUeRadioCapId) {
		return true
	}

	return false
}

// SetPlmnAssiUeRadioCapId gets a reference to the given string and assigns it to the PlmnAssiUeRadioCapId field.
func (o *DicEntryData) SetPlmnAssiUeRadioCapId(v string) {
	o.PlmnAssiUeRadioCapId = &v
}

// GetManAssiUeRadioCapId returns the ManAssiUeRadioCapId field value if set, zero value otherwise.
func (o *DicEntryData) GetManAssiUeRadioCapId() string {
	if o == nil || IsNil(o.ManAssiUeRadioCapId) {
		var ret string
		return ret
	}
	return *o.ManAssiUeRadioCapId
}

// GetManAssiUeRadioCapIdOk returns a tuple with the ManAssiUeRadioCapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DicEntryData) GetManAssiUeRadioCapIdOk() (*string, bool) {
	if o == nil || IsNil(o.ManAssiUeRadioCapId) {
		return nil, false
	}
	return o.ManAssiUeRadioCapId, true
}

// HasManAssiUeRadioCapId returns a boolean if a field has been set.
func (o *DicEntryData) HasManAssiUeRadioCapId() bool {
	if o != nil && !IsNil(o.ManAssiUeRadioCapId) {
		return true
	}

	return false
}

// SetManAssiUeRadioCapId gets a reference to the given string and assigns it to the ManAssiUeRadioCapId field.
func (o *DicEntryData) SetManAssiUeRadioCapId(v string) {
	o.ManAssiUeRadioCapId = &v
}

// GetUeRadioCapability5GS returns the UeRadioCapability5GS field value if set, zero value otherwise.
func (o *DicEntryData) GetUeRadioCapability5GS() RefToBinaryData {
	if o == nil || IsNil(o.UeRadioCapability5GS) {
		var ret RefToBinaryData
		return ret
	}
	return *o.UeRadioCapability5GS
}

// GetUeRadioCapability5GSOk returns a tuple with the UeRadioCapability5GS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DicEntryData) GetUeRadioCapability5GSOk() (*RefToBinaryData, bool) {
	if o == nil || IsNil(o.UeRadioCapability5GS) {
		return nil, false
	}
	return o.UeRadioCapability5GS, true
}

// HasUeRadioCapability5GS returns a boolean if a field has been set.
func (o *DicEntryData) HasUeRadioCapability5GS() bool {
	if o != nil && !IsNil(o.UeRadioCapability5GS) {
		return true
	}

	return false
}

// SetUeRadioCapability5GS gets a reference to the given RefToBinaryData and assigns it to the UeRadioCapability5GS field.
func (o *DicEntryData) SetUeRadioCapability5GS(v RefToBinaryData) {
	o.UeRadioCapability5GS = &v
}

// GetUeRadioCapabilityEPS returns the UeRadioCapabilityEPS field value if set, zero value otherwise.
func (o *DicEntryData) GetUeRadioCapabilityEPS() RefToBinaryData {
	if o == nil || IsNil(o.UeRadioCapabilityEPS) {
		var ret RefToBinaryData
		return ret
	}
	return *o.UeRadioCapabilityEPS
}

// GetUeRadioCapabilityEPSOk returns a tuple with the UeRadioCapabilityEPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DicEntryData) GetUeRadioCapabilityEPSOk() (*RefToBinaryData, bool) {
	if o == nil || IsNil(o.UeRadioCapabilityEPS) {
		return nil, false
	}
	return o.UeRadioCapabilityEPS, true
}

// HasUeRadioCapabilityEPS returns a boolean if a field has been set.
func (o *DicEntryData) HasUeRadioCapabilityEPS() bool {
	if o != nil && !IsNil(o.UeRadioCapabilityEPS) {
		return true
	}

	return false
}

// SetUeRadioCapabilityEPS gets a reference to the given RefToBinaryData and assigns it to the UeRadioCapabilityEPS field.
func (o *DicEntryData) SetUeRadioCapabilityEPS(v RefToBinaryData) {
	o.UeRadioCapabilityEPS = &v
}

// GetUeRadioCap5GSForPaging returns the UeRadioCap5GSForPaging field value if set, zero value otherwise.
func (o *DicEntryData) GetUeRadioCap5GSForPaging() RefToBinaryData {
	if o == nil || IsNil(o.UeRadioCap5GSForPaging) {
		var ret RefToBinaryData
		return ret
	}
	return *o.UeRadioCap5GSForPaging
}

// GetUeRadioCap5GSForPagingOk returns a tuple with the UeRadioCap5GSForPaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DicEntryData) GetUeRadioCap5GSForPagingOk() (*RefToBinaryData, bool) {
	if o == nil || IsNil(o.UeRadioCap5GSForPaging) {
		return nil, false
	}
	return o.UeRadioCap5GSForPaging, true
}

// HasUeRadioCap5GSForPaging returns a boolean if a field has been set.
func (o *DicEntryData) HasUeRadioCap5GSForPaging() bool {
	if o != nil && !IsNil(o.UeRadioCap5GSForPaging) {
		return true
	}

	return false
}

// SetUeRadioCap5GSForPaging gets a reference to the given RefToBinaryData and assigns it to the UeRadioCap5GSForPaging field.
func (o *DicEntryData) SetUeRadioCap5GSForPaging(v RefToBinaryData) {
	o.UeRadioCap5GSForPaging = &v
}

// GetUeRadioCapEPSForPaging returns the UeRadioCapEPSForPaging field value if set, zero value otherwise.
func (o *DicEntryData) GetUeRadioCapEPSForPaging() RefToBinaryData {
	if o == nil || IsNil(o.UeRadioCapEPSForPaging) {
		var ret RefToBinaryData
		return ret
	}
	return *o.UeRadioCapEPSForPaging
}

// GetUeRadioCapEPSForPagingOk returns a tuple with the UeRadioCapEPSForPaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DicEntryData) GetUeRadioCapEPSForPagingOk() (*RefToBinaryData, bool) {
	if o == nil || IsNil(o.UeRadioCapEPSForPaging) {
		return nil, false
	}
	return o.UeRadioCapEPSForPaging, true
}

// HasUeRadioCapEPSForPaging returns a boolean if a field has been set.
func (o *DicEntryData) HasUeRadioCapEPSForPaging() bool {
	if o != nil && !IsNil(o.UeRadioCapEPSForPaging) {
		return true
	}

	return false
}

// SetUeRadioCapEPSForPaging gets a reference to the given RefToBinaryData and assigns it to the UeRadioCapEPSForPaging field.
func (o *DicEntryData) SetUeRadioCapEPSForPaging(v RefToBinaryData) {
	o.UeRadioCapEPSForPaging = &v
}

func (o DicEntryData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DicEntryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DicEntryId) {
		toSerialize["dicEntryId"] = o.DicEntryId
	}
	toSerialize["typeAllocationCode"] = o.TypeAllocationCode
	if !IsNil(o.PlmnAssiUeRadioCapId) {
		toSerialize["plmnAssiUeRadioCapId"] = o.PlmnAssiUeRadioCapId
	}
	if !IsNil(o.ManAssiUeRadioCapId) {
		toSerialize["manAssiUeRadioCapId"] = o.ManAssiUeRadioCapId
	}
	if !IsNil(o.UeRadioCapability5GS) {
		toSerialize["ueRadioCapability5GS"] = o.UeRadioCapability5GS
	}
	if !IsNil(o.UeRadioCapabilityEPS) {
		toSerialize["ueRadioCapabilityEPS"] = o.UeRadioCapabilityEPS
	}
	if !IsNil(o.UeRadioCap5GSForPaging) {
		toSerialize["ueRadioCap5GSForPaging"] = o.UeRadioCap5GSForPaging
	}
	if !IsNil(o.UeRadioCapEPSForPaging) {
		toSerialize["ueRadioCapEPSForPaging"] = o.UeRadioCapEPSForPaging
	}
	return toSerialize, nil
}

type NullableDicEntryData struct {
	value *DicEntryData
	isSet bool
}

func (v NullableDicEntryData) Get() *DicEntryData {
	return v.value
}

func (v *NullableDicEntryData) Set(val *DicEntryData) {
	v.value = val
	v.isSet = true
}

func (v NullableDicEntryData) IsSet() bool {
	return v.isSet
}

func (v *NullableDicEntryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDicEntryData(val *DicEntryData) *NullableDicEntryData {
	return &NullableDicEntryData{value: val, isSet: true}
}

func (v NullableDicEntryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDicEntryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



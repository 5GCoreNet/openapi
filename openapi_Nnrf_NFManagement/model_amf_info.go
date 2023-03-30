/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
)

// checks if the AmfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfInfo{}

// AmfInfo Information of an AMF NF Instance
type AmfInfo struct {
	// String identifying the AMF Set ID (10 bits) as specified in clause 2.10.1 of 3GPP TS 23.003.  It is encoded as a string of 3 hexadecimal characters where the first character is limited to  values 0 to 3 (i.e. 10 bits). 
	AmfSetId string `json:"amfSetId"`
	// String identifying the AMF Set ID (10 bits) as specified in clause 2.10.1 of 3GPP TS 23.003.  It is encoded as a string of 3 hexadecimal characters where the first character is limited to  values 0 to 3 (i.e. 10 bits) 
	AmfRegionId string `json:"amfRegionId"`
	GuamiList []Guami `json:"guamiList"`
	TaiList []Tai `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	BackupInfoAmfFailure []Guami `json:"backupInfoAmfFailure,omitempty"`
	BackupInfoAmfRemoval []Guami `json:"backupInfoAmfRemoval,omitempty"`
	N2InterfaceAmfInfo *N2InterfaceAmfInfo `json:"n2InterfaceAmfInfo,omitempty"`
	AmfOnboardingCapability *bool `json:"amfOnboardingCapability,omitempty"`
}

// NewAmfInfo instantiates a new AmfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfInfo(amfSetId string, amfRegionId string, guamiList []Guami) *AmfInfo {
	this := AmfInfo{}
	this.AmfSetId = amfSetId
	this.AmfRegionId = amfRegionId
	this.GuamiList = guamiList
	var amfOnboardingCapability bool = false
	this.AmfOnboardingCapability = &amfOnboardingCapability
	return &this
}

// NewAmfInfoWithDefaults instantiates a new AmfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfInfoWithDefaults() *AmfInfo {
	this := AmfInfo{}
	var amfOnboardingCapability bool = false
	this.AmfOnboardingCapability = &amfOnboardingCapability
	return &this
}

// GetAmfSetId returns the AmfSetId field value
func (o *AmfInfo) GetAmfSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmfSetId
}

// GetAmfSetIdOk returns a tuple with the AmfSetId field value
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetAmfSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmfSetId, true
}

// SetAmfSetId sets field value
func (o *AmfInfo) SetAmfSetId(v string) {
	o.AmfSetId = v
}

// GetAmfRegionId returns the AmfRegionId field value
func (o *AmfInfo) GetAmfRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmfRegionId
}

// GetAmfRegionIdOk returns a tuple with the AmfRegionId field value
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetAmfRegionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmfRegionId, true
}

// SetAmfRegionId sets field value
func (o *AmfInfo) SetAmfRegionId(v string) {
	o.AmfRegionId = v
}

// GetGuamiList returns the GuamiList field value
func (o *AmfInfo) GetGuamiList() []Guami {
	if o == nil {
		var ret []Guami
		return ret
	}

	return o.GuamiList
}

// GetGuamiListOk returns a tuple with the GuamiList field value
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetGuamiListOk() ([]Guami, bool) {
	if o == nil {
		return nil, false
	}
	return o.GuamiList, true
}

// SetGuamiList sets field value
func (o *AmfInfo) SetGuamiList(v []Guami) {
	o.GuamiList = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *AmfInfo) GetTaiList() []Tai {
	if o == nil || IsNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TaiList) {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *AmfInfo) HasTaiList() bool {
	if o != nil && !IsNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *AmfInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *AmfInfo) GetTaiRangeList() []TaiRange {
	if o == nil || IsNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || IsNil(o.TaiRangeList) {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *AmfInfo) HasTaiRangeList() bool {
	if o != nil && !IsNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *AmfInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetBackupInfoAmfFailure returns the BackupInfoAmfFailure field value if set, zero value otherwise.
func (o *AmfInfo) GetBackupInfoAmfFailure() []Guami {
	if o == nil || IsNil(o.BackupInfoAmfFailure) {
		var ret []Guami
		return ret
	}
	return o.BackupInfoAmfFailure
}

// GetBackupInfoAmfFailureOk returns a tuple with the BackupInfoAmfFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetBackupInfoAmfFailureOk() ([]Guami, bool) {
	if o == nil || IsNil(o.BackupInfoAmfFailure) {
		return nil, false
	}
	return o.BackupInfoAmfFailure, true
}

// HasBackupInfoAmfFailure returns a boolean if a field has been set.
func (o *AmfInfo) HasBackupInfoAmfFailure() bool {
	if o != nil && !IsNil(o.BackupInfoAmfFailure) {
		return true
	}

	return false
}

// SetBackupInfoAmfFailure gets a reference to the given []Guami and assigns it to the BackupInfoAmfFailure field.
func (o *AmfInfo) SetBackupInfoAmfFailure(v []Guami) {
	o.BackupInfoAmfFailure = v
}

// GetBackupInfoAmfRemoval returns the BackupInfoAmfRemoval field value if set, zero value otherwise.
func (o *AmfInfo) GetBackupInfoAmfRemoval() []Guami {
	if o == nil || IsNil(o.BackupInfoAmfRemoval) {
		var ret []Guami
		return ret
	}
	return o.BackupInfoAmfRemoval
}

// GetBackupInfoAmfRemovalOk returns a tuple with the BackupInfoAmfRemoval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetBackupInfoAmfRemovalOk() ([]Guami, bool) {
	if o == nil || IsNil(o.BackupInfoAmfRemoval) {
		return nil, false
	}
	return o.BackupInfoAmfRemoval, true
}

// HasBackupInfoAmfRemoval returns a boolean if a field has been set.
func (o *AmfInfo) HasBackupInfoAmfRemoval() bool {
	if o != nil && !IsNil(o.BackupInfoAmfRemoval) {
		return true
	}

	return false
}

// SetBackupInfoAmfRemoval gets a reference to the given []Guami and assigns it to the BackupInfoAmfRemoval field.
func (o *AmfInfo) SetBackupInfoAmfRemoval(v []Guami) {
	o.BackupInfoAmfRemoval = v
}

// GetN2InterfaceAmfInfo returns the N2InterfaceAmfInfo field value if set, zero value otherwise.
func (o *AmfInfo) GetN2InterfaceAmfInfo() N2InterfaceAmfInfo {
	if o == nil || IsNil(o.N2InterfaceAmfInfo) {
		var ret N2InterfaceAmfInfo
		return ret
	}
	return *o.N2InterfaceAmfInfo
}

// GetN2InterfaceAmfInfoOk returns a tuple with the N2InterfaceAmfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetN2InterfaceAmfInfoOk() (*N2InterfaceAmfInfo, bool) {
	if o == nil || IsNil(o.N2InterfaceAmfInfo) {
		return nil, false
	}
	return o.N2InterfaceAmfInfo, true
}

// HasN2InterfaceAmfInfo returns a boolean if a field has been set.
func (o *AmfInfo) HasN2InterfaceAmfInfo() bool {
	if o != nil && !IsNil(o.N2InterfaceAmfInfo) {
		return true
	}

	return false
}

// SetN2InterfaceAmfInfo gets a reference to the given N2InterfaceAmfInfo and assigns it to the N2InterfaceAmfInfo field.
func (o *AmfInfo) SetN2InterfaceAmfInfo(v N2InterfaceAmfInfo) {
	o.N2InterfaceAmfInfo = &v
}

// GetAmfOnboardingCapability returns the AmfOnboardingCapability field value if set, zero value otherwise.
func (o *AmfInfo) GetAmfOnboardingCapability() bool {
	if o == nil || IsNil(o.AmfOnboardingCapability) {
		var ret bool
		return ret
	}
	return *o.AmfOnboardingCapability
}

// GetAmfOnboardingCapabilityOk returns a tuple with the AmfOnboardingCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfInfo) GetAmfOnboardingCapabilityOk() (*bool, bool) {
	if o == nil || IsNil(o.AmfOnboardingCapability) {
		return nil, false
	}
	return o.AmfOnboardingCapability, true
}

// HasAmfOnboardingCapability returns a boolean if a field has been set.
func (o *AmfInfo) HasAmfOnboardingCapability() bool {
	if o != nil && !IsNil(o.AmfOnboardingCapability) {
		return true
	}

	return false
}

// SetAmfOnboardingCapability gets a reference to the given bool and assigns it to the AmfOnboardingCapability field.
func (o *AmfInfo) SetAmfOnboardingCapability(v bool) {
	o.AmfOnboardingCapability = &v
}

func (o AmfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amfSetId"] = o.AmfSetId
	toSerialize["amfRegionId"] = o.AmfRegionId
	toSerialize["guamiList"] = o.GuamiList
	if !IsNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !IsNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !IsNil(o.BackupInfoAmfFailure) {
		toSerialize["backupInfoAmfFailure"] = o.BackupInfoAmfFailure
	}
	if !IsNil(o.BackupInfoAmfRemoval) {
		toSerialize["backupInfoAmfRemoval"] = o.BackupInfoAmfRemoval
	}
	if !IsNil(o.N2InterfaceAmfInfo) {
		toSerialize["n2InterfaceAmfInfo"] = o.N2InterfaceAmfInfo
	}
	if !IsNil(o.AmfOnboardingCapability) {
		toSerialize["amfOnboardingCapability"] = o.AmfOnboardingCapability
	}
	return toSerialize, nil
}

type NullableAmfInfo struct {
	value *AmfInfo
	isSet bool
}

func (v NullableAmfInfo) Get() *AmfInfo {
	return v.value
}

func (v *NullableAmfInfo) Set(val *AmfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfInfo(val *AmfInfo) *NullableAmfInfo {
	return &NullableAmfInfo{value: val, isSet: true}
}

func (v NullableAmfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



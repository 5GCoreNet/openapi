/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the TrustAfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrustAfInfo{}

// TrustAfInfo Information of a trusted AF Instance
type TrustAfInfo struct {
	SNssaiInfoList []SnssaiInfoItem `json:"sNssaiInfoList,omitempty"`
	AfEvents []AfEvent `json:"afEvents,omitempty"`
	AppIds []string `json:"appIds,omitempty"`
	InternalGroupId []string `json:"internalGroupId,omitempty"`
	MappingInd *bool `json:"mappingInd,omitempty"`
}

// NewTrustAfInfo instantiates a new TrustAfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustAfInfo() *TrustAfInfo {
	this := TrustAfInfo{}
	var mappingInd bool = false
	this.MappingInd = &mappingInd
	return &this
}

// NewTrustAfInfoWithDefaults instantiates a new TrustAfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustAfInfoWithDefaults() *TrustAfInfo {
	this := TrustAfInfo{}
	var mappingInd bool = false
	this.MappingInd = &mappingInd
	return &this
}

// GetSNssaiInfoList returns the SNssaiInfoList field value if set, zero value otherwise.
func (o *TrustAfInfo) GetSNssaiInfoList() []SnssaiInfoItem {
	if o == nil || isNil(o.SNssaiInfoList) {
		var ret []SnssaiInfoItem
		return ret
	}
	return o.SNssaiInfoList
}

// GetSNssaiInfoListOk returns a tuple with the SNssaiInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustAfInfo) GetSNssaiInfoListOk() ([]SnssaiInfoItem, bool) {
	if o == nil || isNil(o.SNssaiInfoList) {
		return nil, false
	}
	return o.SNssaiInfoList, true
}

// HasSNssaiInfoList returns a boolean if a field has been set.
func (o *TrustAfInfo) HasSNssaiInfoList() bool {
	if o != nil && !isNil(o.SNssaiInfoList) {
		return true
	}

	return false
}

// SetSNssaiInfoList gets a reference to the given []SnssaiInfoItem and assigns it to the SNssaiInfoList field.
func (o *TrustAfInfo) SetSNssaiInfoList(v []SnssaiInfoItem) {
	o.SNssaiInfoList = v
}

// GetAfEvents returns the AfEvents field value if set, zero value otherwise.
func (o *TrustAfInfo) GetAfEvents() []AfEvent {
	if o == nil || isNil(o.AfEvents) {
		var ret []AfEvent
		return ret
	}
	return o.AfEvents
}

// GetAfEventsOk returns a tuple with the AfEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustAfInfo) GetAfEventsOk() ([]AfEvent, bool) {
	if o == nil || isNil(o.AfEvents) {
		return nil, false
	}
	return o.AfEvents, true
}

// HasAfEvents returns a boolean if a field has been set.
func (o *TrustAfInfo) HasAfEvents() bool {
	if o != nil && !isNil(o.AfEvents) {
		return true
	}

	return false
}

// SetAfEvents gets a reference to the given []AfEvent and assigns it to the AfEvents field.
func (o *TrustAfInfo) SetAfEvents(v []AfEvent) {
	o.AfEvents = v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *TrustAfInfo) GetAppIds() []string {
	if o == nil || isNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustAfInfo) GetAppIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AppIds) {
		return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *TrustAfInfo) HasAppIds() bool {
	if o != nil && !isNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *TrustAfInfo) SetAppIds(v []string) {
	o.AppIds = v
}

// GetInternalGroupId returns the InternalGroupId field value if set, zero value otherwise.
func (o *TrustAfInfo) GetInternalGroupId() []string {
	if o == nil || isNil(o.InternalGroupId) {
		var ret []string
		return ret
	}
	return o.InternalGroupId
}

// GetInternalGroupIdOk returns a tuple with the InternalGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustAfInfo) GetInternalGroupIdOk() ([]string, bool) {
	if o == nil || isNil(o.InternalGroupId) {
		return nil, false
	}
	return o.InternalGroupId, true
}

// HasInternalGroupId returns a boolean if a field has been set.
func (o *TrustAfInfo) HasInternalGroupId() bool {
	if o != nil && !isNil(o.InternalGroupId) {
		return true
	}

	return false
}

// SetInternalGroupId gets a reference to the given []string and assigns it to the InternalGroupId field.
func (o *TrustAfInfo) SetInternalGroupId(v []string) {
	o.InternalGroupId = v
}

// GetMappingInd returns the MappingInd field value if set, zero value otherwise.
func (o *TrustAfInfo) GetMappingInd() bool {
	if o == nil || isNil(o.MappingInd) {
		var ret bool
		return ret
	}
	return *o.MappingInd
}

// GetMappingIndOk returns a tuple with the MappingInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustAfInfo) GetMappingIndOk() (*bool, bool) {
	if o == nil || isNil(o.MappingInd) {
		return nil, false
	}
	return o.MappingInd, true
}

// HasMappingInd returns a boolean if a field has been set.
func (o *TrustAfInfo) HasMappingInd() bool {
	if o != nil && !isNil(o.MappingInd) {
		return true
	}

	return false
}

// SetMappingInd gets a reference to the given bool and assigns it to the MappingInd field.
func (o *TrustAfInfo) SetMappingInd(v bool) {
	o.MappingInd = &v
}

func (o TrustAfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrustAfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SNssaiInfoList) {
		toSerialize["sNssaiInfoList"] = o.SNssaiInfoList
	}
	if !isNil(o.AfEvents) {
		toSerialize["afEvents"] = o.AfEvents
	}
	if !isNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	if !isNil(o.InternalGroupId) {
		toSerialize["internalGroupId"] = o.InternalGroupId
	}
	if !isNil(o.MappingInd) {
		toSerialize["mappingInd"] = o.MappingInd
	}
	return toSerialize, nil
}

type NullableTrustAfInfo struct {
	value *TrustAfInfo
	isSet bool
}

func (v NullableTrustAfInfo) Get() *TrustAfInfo {
	return v.value
}

func (v *NullableTrustAfInfo) Set(val *TrustAfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustAfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustAfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustAfInfo(val *TrustAfInfo) *NullableTrustAfInfo {
	return &NullableTrustAfInfo{value: val, isSet: true}
}

func (v NullableTrustAfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustAfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the EeSubscriptionExt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EeSubscriptionExt{}

// EeSubscriptionExt struct for EeSubscriptionExt
type EeSubscriptionExt struct {
	EeSubscription1
	AmfSubscriptionInfoList []AmfSubscriptionInfo `json:"amfSubscriptionInfoList,omitempty"`
	SmfSubscriptionInfo     *SmfSubscriptionInfo  `json:"smfSubscriptionInfo,omitempty"`
	HssSubscriptionInfo     *HssSubscriptionInfo  `json:"hssSubscriptionInfo,omitempty"`
}

// NewEeSubscriptionExt instantiates a new EeSubscriptionExt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEeSubscriptionExt(callbackReference string, monitoringConfigurations map[string]MonitoringConfiguration1) *EeSubscriptionExt {
	this := EeSubscriptionExt{}
	this.CallbackReference = callbackReference
	this.MonitoringConfigurations = monitoringConfigurations
	var epcAppliedInd bool = false
	this.EpcAppliedInd = &epcAppliedInd
	return &this
}

// NewEeSubscriptionExtWithDefaults instantiates a new EeSubscriptionExt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEeSubscriptionExtWithDefaults() *EeSubscriptionExt {
	this := EeSubscriptionExt{}
	return &this
}

// GetAmfSubscriptionInfoList returns the AmfSubscriptionInfoList field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetAmfSubscriptionInfoList() []AmfSubscriptionInfo {
	if o == nil || IsNil(o.AmfSubscriptionInfoList) {
		var ret []AmfSubscriptionInfo
		return ret
	}
	return o.AmfSubscriptionInfoList
}

// GetAmfSubscriptionInfoListOk returns a tuple with the AmfSubscriptionInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetAmfSubscriptionInfoListOk() ([]AmfSubscriptionInfo, bool) {
	if o == nil || IsNil(o.AmfSubscriptionInfoList) {
		return nil, false
	}
	return o.AmfSubscriptionInfoList, true
}

// HasAmfSubscriptionInfoList returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasAmfSubscriptionInfoList() bool {
	if o != nil && !IsNil(o.AmfSubscriptionInfoList) {
		return true
	}

	return false
}

// SetAmfSubscriptionInfoList gets a reference to the given []AmfSubscriptionInfo and assigns it to the AmfSubscriptionInfoList field.
func (o *EeSubscriptionExt) SetAmfSubscriptionInfoList(v []AmfSubscriptionInfo) {
	o.AmfSubscriptionInfoList = v
}

// GetSmfSubscriptionInfo returns the SmfSubscriptionInfo field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetSmfSubscriptionInfo() SmfSubscriptionInfo {
	if o == nil || IsNil(o.SmfSubscriptionInfo) {
		var ret SmfSubscriptionInfo
		return ret
	}
	return *o.SmfSubscriptionInfo
}

// GetSmfSubscriptionInfoOk returns a tuple with the SmfSubscriptionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetSmfSubscriptionInfoOk() (*SmfSubscriptionInfo, bool) {
	if o == nil || IsNil(o.SmfSubscriptionInfo) {
		return nil, false
	}
	return o.SmfSubscriptionInfo, true
}

// HasSmfSubscriptionInfo returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasSmfSubscriptionInfo() bool {
	if o != nil && !IsNil(o.SmfSubscriptionInfo) {
		return true
	}

	return false
}

// SetSmfSubscriptionInfo gets a reference to the given SmfSubscriptionInfo and assigns it to the SmfSubscriptionInfo field.
func (o *EeSubscriptionExt) SetSmfSubscriptionInfo(v SmfSubscriptionInfo) {
	o.SmfSubscriptionInfo = &v
}

// GetHssSubscriptionInfo returns the HssSubscriptionInfo field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetHssSubscriptionInfo() HssSubscriptionInfo {
	if o == nil || IsNil(o.HssSubscriptionInfo) {
		var ret HssSubscriptionInfo
		return ret
	}
	return *o.HssSubscriptionInfo
}

// GetHssSubscriptionInfoOk returns a tuple with the HssSubscriptionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetHssSubscriptionInfoOk() (*HssSubscriptionInfo, bool) {
	if o == nil || IsNil(o.HssSubscriptionInfo) {
		return nil, false
	}
	return o.HssSubscriptionInfo, true
}

// HasHssSubscriptionInfo returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasHssSubscriptionInfo() bool {
	if o != nil && !IsNil(o.HssSubscriptionInfo) {
		return true
	}

	return false
}

// SetHssSubscriptionInfo gets a reference to the given HssSubscriptionInfo and assigns it to the HssSubscriptionInfo field.
func (o *EeSubscriptionExt) SetHssSubscriptionInfo(v HssSubscriptionInfo) {
	o.HssSubscriptionInfo = &v
}

func (o EeSubscriptionExt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EeSubscriptionExt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEeSubscription1, errEeSubscription1 := json.Marshal(o.EeSubscription1)
	if errEeSubscription1 != nil {
		return map[string]interface{}{}, errEeSubscription1
	}
	errEeSubscription1 = json.Unmarshal([]byte(serializedEeSubscription1), &toSerialize)
	if errEeSubscription1 != nil {
		return map[string]interface{}{}, errEeSubscription1
	}
	if !IsNil(o.AmfSubscriptionInfoList) {
		toSerialize["amfSubscriptionInfoList"] = o.AmfSubscriptionInfoList
	}
	if !IsNil(o.SmfSubscriptionInfo) {
		toSerialize["smfSubscriptionInfo"] = o.SmfSubscriptionInfo
	}
	if !IsNil(o.HssSubscriptionInfo) {
		toSerialize["hssSubscriptionInfo"] = o.HssSubscriptionInfo
	}
	return toSerialize, nil
}

type NullableEeSubscriptionExt struct {
	value *EeSubscriptionExt
	isSet bool
}

func (v NullableEeSubscriptionExt) Get() *EeSubscriptionExt {
	return v.value
}

func (v *NullableEeSubscriptionExt) Set(val *EeSubscriptionExt) {
	v.value = val
	v.isSet = true
}

func (v NullableEeSubscriptionExt) IsSet() bool {
	return v.isSet
}

func (v *NullableEeSubscriptionExt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEeSubscriptionExt(val *EeSubscriptionExt) *NullableEeSubscriptionExt {
	return &NullableEeSubscriptionExt{value: val, isSet: true}
}

func (v NullableEeSubscriptionExt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEeSubscriptionExt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

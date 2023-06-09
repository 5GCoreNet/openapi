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

// checks if the WlanPerformanceReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WlanPerformanceReq{}

// WlanPerformanceReq Represents other WLAN performance analytics requirements.
type WlanPerformanceReq struct {
	SsIds           []string               `json:"ssIds,omitempty"`
	BssIds          []string               `json:"bssIds,omitempty"`
	WlanOrderCriter *WlanOrderingCriterion `json:"wlanOrderCriter,omitempty"`
	Order           *MatchingDirection     `json:"order,omitempty"`
}

// NewWlanPerformanceReq instantiates a new WlanPerformanceReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWlanPerformanceReq() *WlanPerformanceReq {
	this := WlanPerformanceReq{}
	return &this
}

// NewWlanPerformanceReqWithDefaults instantiates a new WlanPerformanceReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWlanPerformanceReqWithDefaults() *WlanPerformanceReq {
	this := WlanPerformanceReq{}
	return &this
}

// GetSsIds returns the SsIds field value if set, zero value otherwise.
func (o *WlanPerformanceReq) GetSsIds() []string {
	if o == nil || IsNil(o.SsIds) {
		var ret []string
		return ret
	}
	return o.SsIds
}

// GetSsIdsOk returns a tuple with the SsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanPerformanceReq) GetSsIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SsIds) {
		return nil, false
	}
	return o.SsIds, true
}

// HasSsIds returns a boolean if a field has been set.
func (o *WlanPerformanceReq) HasSsIds() bool {
	if o != nil && !IsNil(o.SsIds) {
		return true
	}

	return false
}

// SetSsIds gets a reference to the given []string and assigns it to the SsIds field.
func (o *WlanPerformanceReq) SetSsIds(v []string) {
	o.SsIds = v
}

// GetBssIds returns the BssIds field value if set, zero value otherwise.
func (o *WlanPerformanceReq) GetBssIds() []string {
	if o == nil || IsNil(o.BssIds) {
		var ret []string
		return ret
	}
	return o.BssIds
}

// GetBssIdsOk returns a tuple with the BssIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanPerformanceReq) GetBssIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.BssIds) {
		return nil, false
	}
	return o.BssIds, true
}

// HasBssIds returns a boolean if a field has been set.
func (o *WlanPerformanceReq) HasBssIds() bool {
	if o != nil && !IsNil(o.BssIds) {
		return true
	}

	return false
}

// SetBssIds gets a reference to the given []string and assigns it to the BssIds field.
func (o *WlanPerformanceReq) SetBssIds(v []string) {
	o.BssIds = v
}

// GetWlanOrderCriter returns the WlanOrderCriter field value if set, zero value otherwise.
func (o *WlanPerformanceReq) GetWlanOrderCriter() WlanOrderingCriterion {
	if o == nil || IsNil(o.WlanOrderCriter) {
		var ret WlanOrderingCriterion
		return ret
	}
	return *o.WlanOrderCriter
}

// GetWlanOrderCriterOk returns a tuple with the WlanOrderCriter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanPerformanceReq) GetWlanOrderCriterOk() (*WlanOrderingCriterion, bool) {
	if o == nil || IsNil(o.WlanOrderCriter) {
		return nil, false
	}
	return o.WlanOrderCriter, true
}

// HasWlanOrderCriter returns a boolean if a field has been set.
func (o *WlanPerformanceReq) HasWlanOrderCriter() bool {
	if o != nil && !IsNil(o.WlanOrderCriter) {
		return true
	}

	return false
}

// SetWlanOrderCriter gets a reference to the given WlanOrderingCriterion and assigns it to the WlanOrderCriter field.
func (o *WlanPerformanceReq) SetWlanOrderCriter(v WlanOrderingCriterion) {
	o.WlanOrderCriter = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *WlanPerformanceReq) GetOrder() MatchingDirection {
	if o == nil || IsNil(o.Order) {
		var ret MatchingDirection
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanPerformanceReq) GetOrderOk() (*MatchingDirection, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *WlanPerformanceReq) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given MatchingDirection and assigns it to the Order field.
func (o *WlanPerformanceReq) SetOrder(v MatchingDirection) {
	o.Order = &v
}

func (o WlanPerformanceReq) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WlanPerformanceReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SsIds) {
		toSerialize["ssIds"] = o.SsIds
	}
	if !IsNil(o.BssIds) {
		toSerialize["bssIds"] = o.BssIds
	}
	if !IsNil(o.WlanOrderCriter) {
		toSerialize["wlanOrderCriter"] = o.WlanOrderCriter
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableWlanPerformanceReq struct {
	value *WlanPerformanceReq
	isSet bool
}

func (v NullableWlanPerformanceReq) Get() *WlanPerformanceReq {
	return v.value
}

func (v *NullableWlanPerformanceReq) Set(val *WlanPerformanceReq) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanPerformanceReq) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanPerformanceReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanPerformanceReq(val *WlanPerformanceReq) *NullableWlanPerformanceReq {
	return &NullableWlanPerformanceReq{value: val, isSet: true}
}

func (v NullableWlanPerformanceReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanPerformanceReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

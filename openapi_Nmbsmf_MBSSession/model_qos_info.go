/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
)

// checks if the QosInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosInfo{}

// QosInfo QoS Information
type QosInfo struct {
	QosFlowsAddModRequestList []QosFlowAddModifyRequestItem `json:"qosFlowsAddModRequestList,omitempty"`
	QosFlowsRelRequestList    []int32                       `json:"qosFlowsRelRequestList,omitempty"`
}

// NewQosInfo instantiates a new QosInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosInfo() *QosInfo {
	this := QosInfo{}
	return &this
}

// NewQosInfoWithDefaults instantiates a new QosInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosInfoWithDefaults() *QosInfo {
	this := QosInfo{}
	return &this
}

// GetQosFlowsAddModRequestList returns the QosFlowsAddModRequestList field value if set, zero value otherwise.
func (o *QosInfo) GetQosFlowsAddModRequestList() []QosFlowAddModifyRequestItem {
	if o == nil || IsNil(o.QosFlowsAddModRequestList) {
		var ret []QosFlowAddModifyRequestItem
		return ret
	}
	return o.QosFlowsAddModRequestList
}

// GetQosFlowsAddModRequestListOk returns a tuple with the QosFlowsAddModRequestList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosInfo) GetQosFlowsAddModRequestListOk() ([]QosFlowAddModifyRequestItem, bool) {
	if o == nil || IsNil(o.QosFlowsAddModRequestList) {
		return nil, false
	}
	return o.QosFlowsAddModRequestList, true
}

// HasQosFlowsAddModRequestList returns a boolean if a field has been set.
func (o *QosInfo) HasQosFlowsAddModRequestList() bool {
	if o != nil && !IsNil(o.QosFlowsAddModRequestList) {
		return true
	}

	return false
}

// SetQosFlowsAddModRequestList gets a reference to the given []QosFlowAddModifyRequestItem and assigns it to the QosFlowsAddModRequestList field.
func (o *QosInfo) SetQosFlowsAddModRequestList(v []QosFlowAddModifyRequestItem) {
	o.QosFlowsAddModRequestList = v
}

// GetQosFlowsRelRequestList returns the QosFlowsRelRequestList field value if set, zero value otherwise.
func (o *QosInfo) GetQosFlowsRelRequestList() []int32 {
	if o == nil || IsNil(o.QosFlowsRelRequestList) {
		var ret []int32
		return ret
	}
	return o.QosFlowsRelRequestList
}

// GetQosFlowsRelRequestListOk returns a tuple with the QosFlowsRelRequestList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosInfo) GetQosFlowsRelRequestListOk() ([]int32, bool) {
	if o == nil || IsNil(o.QosFlowsRelRequestList) {
		return nil, false
	}
	return o.QosFlowsRelRequestList, true
}

// HasQosFlowsRelRequestList returns a boolean if a field has been set.
func (o *QosInfo) HasQosFlowsRelRequestList() bool {
	if o != nil && !IsNil(o.QosFlowsRelRequestList) {
		return true
	}

	return false
}

// SetQosFlowsRelRequestList gets a reference to the given []int32 and assigns it to the QosFlowsRelRequestList field.
func (o *QosInfo) SetQosFlowsRelRequestList(v []int32) {
	o.QosFlowsRelRequestList = v
}

func (o QosInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QosFlowsAddModRequestList) {
		toSerialize["qosFlowsAddModRequestList"] = o.QosFlowsAddModRequestList
	}
	if !IsNil(o.QosFlowsRelRequestList) {
		toSerialize["qosFlowsRelRequestList"] = o.QosFlowsRelRequestList
	}
	return toSerialize, nil
}

type NullableQosInfo struct {
	value *QosInfo
	isSet bool
}

func (v NullableQosInfo) Get() *QosInfo {
	return v.value
}

func (v *NullableQosInfo) Set(val *QosInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableQosInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableQosInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosInfo(val *QosInfo) *NullableQosInfo {
	return &NullableQosInfo{value: val, isSet: true}
}

func (v NullableQosInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
)

// checks if the QosNotificationControlInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosNotificationControlInfo{}

// QosNotificationControlInfo Contains the QoS Notification Control Information.
type QosNotificationControlInfo struct {
	// An array of PCC rule id references to the PCC rules associated with the QoS notification  control info.
	RefPccRuleIds []string     `json:"refPccRuleIds"`
	NotifType     QosNotifType `json:"notifType"`
	// Represents the content version of some content.
	ContVer *int32 `json:"contVer,omitempty"`
	// Indicates the alternative QoS parameter set the NG-RAN can guarantee. When it is omitted and the notifType attribute is set to NOT_GUAARANTEED it indicates that the lowest priority alternative QoS profile could not be fulfilled.
	AltQosParamId *string `json:"altQosParamId,omitempty"`
	// When present and set to true it indicates that the Alternative QoS profiles are not  supported by NG-RAN.
	AltQosNotSuppInd *bool `json:"altQosNotSuppInd,omitempty"`
}

// NewQosNotificationControlInfo instantiates a new QosNotificationControlInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosNotificationControlInfo(refPccRuleIds []string, notifType QosNotifType) *QosNotificationControlInfo {
	this := QosNotificationControlInfo{}
	this.RefPccRuleIds = refPccRuleIds
	this.NotifType = notifType
	return &this
}

// NewQosNotificationControlInfoWithDefaults instantiates a new QosNotificationControlInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosNotificationControlInfoWithDefaults() *QosNotificationControlInfo {
	this := QosNotificationControlInfo{}
	return &this
}

// GetRefPccRuleIds returns the RefPccRuleIds field value
func (o *QosNotificationControlInfo) GetRefPccRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RefPccRuleIds
}

// GetRefPccRuleIdsOk returns a tuple with the RefPccRuleIds field value
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetRefPccRuleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefPccRuleIds, true
}

// SetRefPccRuleIds sets field value
func (o *QosNotificationControlInfo) SetRefPccRuleIds(v []string) {
	o.RefPccRuleIds = v
}

// GetNotifType returns the NotifType field value
func (o *QosNotificationControlInfo) GetNotifType() QosNotifType {
	if o == nil {
		var ret QosNotifType
		return ret
	}

	return o.NotifType
}

// GetNotifTypeOk returns a tuple with the NotifType field value
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetNotifTypeOk() (*QosNotifType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifType, true
}

// SetNotifType sets field value
func (o *QosNotificationControlInfo) SetNotifType(v QosNotifType) {
	o.NotifType = v
}

// GetContVer returns the ContVer field value if set, zero value otherwise.
func (o *QosNotificationControlInfo) GetContVer() int32 {
	if o == nil || IsNil(o.ContVer) {
		var ret int32
		return ret
	}
	return *o.ContVer
}

// GetContVerOk returns a tuple with the ContVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetContVerOk() (*int32, bool) {
	if o == nil || IsNil(o.ContVer) {
		return nil, false
	}
	return o.ContVer, true
}

// HasContVer returns a boolean if a field has been set.
func (o *QosNotificationControlInfo) HasContVer() bool {
	if o != nil && !IsNil(o.ContVer) {
		return true
	}

	return false
}

// SetContVer gets a reference to the given int32 and assigns it to the ContVer field.
func (o *QosNotificationControlInfo) SetContVer(v int32) {
	o.ContVer = &v
}

// GetAltQosParamId returns the AltQosParamId field value if set, zero value otherwise.
func (o *QosNotificationControlInfo) GetAltQosParamId() string {
	if o == nil || IsNil(o.AltQosParamId) {
		var ret string
		return ret
	}
	return *o.AltQosParamId
}

// GetAltQosParamIdOk returns a tuple with the AltQosParamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetAltQosParamIdOk() (*string, bool) {
	if o == nil || IsNil(o.AltQosParamId) {
		return nil, false
	}
	return o.AltQosParamId, true
}

// HasAltQosParamId returns a boolean if a field has been set.
func (o *QosNotificationControlInfo) HasAltQosParamId() bool {
	if o != nil && !IsNil(o.AltQosParamId) {
		return true
	}

	return false
}

// SetAltQosParamId gets a reference to the given string and assigns it to the AltQosParamId field.
func (o *QosNotificationControlInfo) SetAltQosParamId(v string) {
	o.AltQosParamId = &v
}

// GetAltQosNotSuppInd returns the AltQosNotSuppInd field value if set, zero value otherwise.
func (o *QosNotificationControlInfo) GetAltQosNotSuppInd() bool {
	if o == nil || IsNil(o.AltQosNotSuppInd) {
		var ret bool
		return ret
	}
	return *o.AltQosNotSuppInd
}

// GetAltQosNotSuppIndOk returns a tuple with the AltQosNotSuppInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetAltQosNotSuppIndOk() (*bool, bool) {
	if o == nil || IsNil(o.AltQosNotSuppInd) {
		return nil, false
	}
	return o.AltQosNotSuppInd, true
}

// HasAltQosNotSuppInd returns a boolean if a field has been set.
func (o *QosNotificationControlInfo) HasAltQosNotSuppInd() bool {
	if o != nil && !IsNil(o.AltQosNotSuppInd) {
		return true
	}

	return false
}

// SetAltQosNotSuppInd gets a reference to the given bool and assigns it to the AltQosNotSuppInd field.
func (o *QosNotificationControlInfo) SetAltQosNotSuppInd(v bool) {
	o.AltQosNotSuppInd = &v
}

func (o QosNotificationControlInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosNotificationControlInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["refPccRuleIds"] = o.RefPccRuleIds
	toSerialize["notifType"] = o.NotifType
	if !IsNil(o.ContVer) {
		toSerialize["contVer"] = o.ContVer
	}
	if !IsNil(o.AltQosParamId) {
		toSerialize["altQosParamId"] = o.AltQosParamId
	}
	if !IsNil(o.AltQosNotSuppInd) {
		toSerialize["altQosNotSuppInd"] = o.AltQosNotSuppInd
	}
	return toSerialize, nil
}

type NullableQosNotificationControlInfo struct {
	value *QosNotificationControlInfo
	isSet bool
}

func (v NullableQosNotificationControlInfo) Get() *QosNotificationControlInfo {
	return v.value
}

func (v *NullableQosNotificationControlInfo) Set(val *QosNotificationControlInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableQosNotificationControlInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableQosNotificationControlInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosNotificationControlInfo(val *QosNotificationControlInfo) *NullableQosNotificationControlInfo {
	return &NullableQosNotificationControlInfo{value: val, isSet: true}
}

func (v NullableQosNotificationControlInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosNotificationControlInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

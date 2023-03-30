/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// checks if the QosNotificationControlInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosNotificationControlInfo{}

// QosNotificationControlInfo Indicates whether the QoS targets for a GRB flow are not guaranteed or guaranteed again. 
type QosNotificationControlInfo struct {
	NotifType QosNotifType `json:"notifType"`
	Flows []Flows `json:"flows,omitempty"`
	// Indicates the alternative service requirement NG-RAN can guarantee. When it is omitted and the notifType attribute is set to NOT_GUAARANTEED it indicates that the lowest priority alternative alternative service requirement could not be fulfilled by NG-RAN. 
	AltSerReq *string `json:"altSerReq,omitempty"`
	// When present and set to true it indicates that Alternative Service Requirements are not  supported by NG-RAN. 
	AltSerReqNotSuppInd *bool `json:"altSerReqNotSuppInd,omitempty"`
}

// NewQosNotificationControlInfo instantiates a new QosNotificationControlInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosNotificationControlInfo(notifType QosNotifType) *QosNotificationControlInfo {
	this := QosNotificationControlInfo{}
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

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *QosNotificationControlInfo) GetFlows() []Flows {
	if o == nil || IsNil(o.Flows) {
		var ret []Flows
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetFlowsOk() ([]Flows, bool) {
	if o == nil || IsNil(o.Flows) {
		return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *QosNotificationControlInfo) HasFlows() bool {
	if o != nil && !IsNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []Flows and assigns it to the Flows field.
func (o *QosNotificationControlInfo) SetFlows(v []Flows) {
	o.Flows = v
}

// GetAltSerReq returns the AltSerReq field value if set, zero value otherwise.
func (o *QosNotificationControlInfo) GetAltSerReq() string {
	if o == nil || IsNil(o.AltSerReq) {
		var ret string
		return ret
	}
	return *o.AltSerReq
}

// GetAltSerReqOk returns a tuple with the AltSerReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetAltSerReqOk() (*string, bool) {
	if o == nil || IsNil(o.AltSerReq) {
		return nil, false
	}
	return o.AltSerReq, true
}

// HasAltSerReq returns a boolean if a field has been set.
func (o *QosNotificationControlInfo) HasAltSerReq() bool {
	if o != nil && !IsNil(o.AltSerReq) {
		return true
	}

	return false
}

// SetAltSerReq gets a reference to the given string and assigns it to the AltSerReq field.
func (o *QosNotificationControlInfo) SetAltSerReq(v string) {
	o.AltSerReq = &v
}

// GetAltSerReqNotSuppInd returns the AltSerReqNotSuppInd field value if set, zero value otherwise.
func (o *QosNotificationControlInfo) GetAltSerReqNotSuppInd() bool {
	if o == nil || IsNil(o.AltSerReqNotSuppInd) {
		var ret bool
		return ret
	}
	return *o.AltSerReqNotSuppInd
}

// GetAltSerReqNotSuppIndOk returns a tuple with the AltSerReqNotSuppInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetAltSerReqNotSuppIndOk() (*bool, bool) {
	if o == nil || IsNil(o.AltSerReqNotSuppInd) {
		return nil, false
	}
	return o.AltSerReqNotSuppInd, true
}

// HasAltSerReqNotSuppInd returns a boolean if a field has been set.
func (o *QosNotificationControlInfo) HasAltSerReqNotSuppInd() bool {
	if o != nil && !IsNil(o.AltSerReqNotSuppInd) {
		return true
	}

	return false
}

// SetAltSerReqNotSuppInd gets a reference to the given bool and assigns it to the AltSerReqNotSuppInd field.
func (o *QosNotificationControlInfo) SetAltSerReqNotSuppInd(v bool) {
	o.AltSerReqNotSuppInd = &v
}

func (o QosNotificationControlInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosNotificationControlInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifType"] = o.NotifType
	if !IsNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	if !IsNil(o.AltSerReq) {
		toSerialize["altSerReq"] = o.AltSerReq
	}
	if !IsNil(o.AltSerReqNotSuppInd) {
		toSerialize["altSerReqNotSuppInd"] = o.AltSerReqNotSuppInd
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



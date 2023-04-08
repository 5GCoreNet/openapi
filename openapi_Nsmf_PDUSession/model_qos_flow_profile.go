/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// checks if the QosFlowProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosFlowProfile{}

// QosFlowProfile QoS flow profile
type QosFlowProfile struct {
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255. 
	Var5qi int32 `json:"5qi"`
	NonDynamic5Qi *NonDynamic5Qi `json:"nonDynamic5Qi,omitempty"`
	Dynamic5Qi *Dynamic5Qi `json:"dynamic5Qi,omitempty"`
	Arp *Arp `json:"arp,omitempty"`
	GbrQosFlowInfo *GbrQosFlowInformation `json:"gbrQosFlowInfo,omitempty"`
	Rqa *ReflectiveQoSAttribute `json:"rqa,omitempty"`
	AdditionalQosFlowInfo *AdditionalQosFlowInfo `json:"additionalQosFlowInfo,omitempty"`
	QosMonitoringReq *QosMonitoringReq `json:"qosMonitoringReq,omitempty"`
	// indicating a time in seconds.
	QosRepPeriod *int32 `json:"qosRepPeriod,omitempty"`
}

// NewQosFlowProfile instantiates a new QosFlowProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosFlowProfile(var5qi int32) *QosFlowProfile {
	this := QosFlowProfile{}
	this.Var5qi = var5qi
	return &this
}

// NewQosFlowProfileWithDefaults instantiates a new QosFlowProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosFlowProfileWithDefaults() *QosFlowProfile {
	this := QosFlowProfile{}
	return &this
}

// GetVar5qi returns the Var5qi field value
func (o *QosFlowProfile) GetVar5qi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var5qi
}

// GetVar5qiOk returns a tuple with the Var5qi field value
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetVar5qiOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var5qi, true
}

// SetVar5qi sets field value
func (o *QosFlowProfile) SetVar5qi(v int32) {
	o.Var5qi = v
}

// GetNonDynamic5Qi returns the NonDynamic5Qi field value if set, zero value otherwise.
func (o *QosFlowProfile) GetNonDynamic5Qi() NonDynamic5Qi {
	if o == nil || isNil(o.NonDynamic5Qi) {
		var ret NonDynamic5Qi
		return ret
	}
	return *o.NonDynamic5Qi
}

// GetNonDynamic5QiOk returns a tuple with the NonDynamic5Qi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetNonDynamic5QiOk() (*NonDynamic5Qi, bool) {
	if o == nil || isNil(o.NonDynamic5Qi) {
		return nil, false
	}
	return o.NonDynamic5Qi, true
}

// HasNonDynamic5Qi returns a boolean if a field has been set.
func (o *QosFlowProfile) HasNonDynamic5Qi() bool {
	if o != nil && !isNil(o.NonDynamic5Qi) {
		return true
	}

	return false
}

// SetNonDynamic5Qi gets a reference to the given NonDynamic5Qi and assigns it to the NonDynamic5Qi field.
func (o *QosFlowProfile) SetNonDynamic5Qi(v NonDynamic5Qi) {
	o.NonDynamic5Qi = &v
}

// GetDynamic5Qi returns the Dynamic5Qi field value if set, zero value otherwise.
func (o *QosFlowProfile) GetDynamic5Qi() Dynamic5Qi {
	if o == nil || isNil(o.Dynamic5Qi) {
		var ret Dynamic5Qi
		return ret
	}
	return *o.Dynamic5Qi
}

// GetDynamic5QiOk returns a tuple with the Dynamic5Qi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetDynamic5QiOk() (*Dynamic5Qi, bool) {
	if o == nil || isNil(o.Dynamic5Qi) {
		return nil, false
	}
	return o.Dynamic5Qi, true
}

// HasDynamic5Qi returns a boolean if a field has been set.
func (o *QosFlowProfile) HasDynamic5Qi() bool {
	if o != nil && !isNil(o.Dynamic5Qi) {
		return true
	}

	return false
}

// SetDynamic5Qi gets a reference to the given Dynamic5Qi and assigns it to the Dynamic5Qi field.
func (o *QosFlowProfile) SetDynamic5Qi(v Dynamic5Qi) {
	o.Dynamic5Qi = &v
}

// GetArp returns the Arp field value if set, zero value otherwise.
func (o *QosFlowProfile) GetArp() Arp {
	if o == nil || isNil(o.Arp) {
		var ret Arp
		return ret
	}
	return *o.Arp
}

// GetArpOk returns a tuple with the Arp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetArpOk() (*Arp, bool) {
	if o == nil || isNil(o.Arp) {
		return nil, false
	}
	return o.Arp, true
}

// HasArp returns a boolean if a field has been set.
func (o *QosFlowProfile) HasArp() bool {
	if o != nil && !isNil(o.Arp) {
		return true
	}

	return false
}

// SetArp gets a reference to the given Arp and assigns it to the Arp field.
func (o *QosFlowProfile) SetArp(v Arp) {
	o.Arp = &v
}

// GetGbrQosFlowInfo returns the GbrQosFlowInfo field value if set, zero value otherwise.
func (o *QosFlowProfile) GetGbrQosFlowInfo() GbrQosFlowInformation {
	if o == nil || isNil(o.GbrQosFlowInfo) {
		var ret GbrQosFlowInformation
		return ret
	}
	return *o.GbrQosFlowInfo
}

// GetGbrQosFlowInfoOk returns a tuple with the GbrQosFlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetGbrQosFlowInfoOk() (*GbrQosFlowInformation, bool) {
	if o == nil || isNil(o.GbrQosFlowInfo) {
		return nil, false
	}
	return o.GbrQosFlowInfo, true
}

// HasGbrQosFlowInfo returns a boolean if a field has been set.
func (o *QosFlowProfile) HasGbrQosFlowInfo() bool {
	if o != nil && !isNil(o.GbrQosFlowInfo) {
		return true
	}

	return false
}

// SetGbrQosFlowInfo gets a reference to the given GbrQosFlowInformation and assigns it to the GbrQosFlowInfo field.
func (o *QosFlowProfile) SetGbrQosFlowInfo(v GbrQosFlowInformation) {
	o.GbrQosFlowInfo = &v
}

// GetRqa returns the Rqa field value if set, zero value otherwise.
func (o *QosFlowProfile) GetRqa() ReflectiveQoSAttribute {
	if o == nil || isNil(o.Rqa) {
		var ret ReflectiveQoSAttribute
		return ret
	}
	return *o.Rqa
}

// GetRqaOk returns a tuple with the Rqa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetRqaOk() (*ReflectiveQoSAttribute, bool) {
	if o == nil || isNil(o.Rqa) {
		return nil, false
	}
	return o.Rqa, true
}

// HasRqa returns a boolean if a field has been set.
func (o *QosFlowProfile) HasRqa() bool {
	if o != nil && !isNil(o.Rqa) {
		return true
	}

	return false
}

// SetRqa gets a reference to the given ReflectiveQoSAttribute and assigns it to the Rqa field.
func (o *QosFlowProfile) SetRqa(v ReflectiveQoSAttribute) {
	o.Rqa = &v
}

// GetAdditionalQosFlowInfo returns the AdditionalQosFlowInfo field value if set, zero value otherwise.
func (o *QosFlowProfile) GetAdditionalQosFlowInfo() AdditionalQosFlowInfo {
	if o == nil || isNil(o.AdditionalQosFlowInfo) {
		var ret AdditionalQosFlowInfo
		return ret
	}
	return *o.AdditionalQosFlowInfo
}

// GetAdditionalQosFlowInfoOk returns a tuple with the AdditionalQosFlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetAdditionalQosFlowInfoOk() (*AdditionalQosFlowInfo, bool) {
	if o == nil || isNil(o.AdditionalQosFlowInfo) {
		return nil, false
	}
	return o.AdditionalQosFlowInfo, true
}

// HasAdditionalQosFlowInfo returns a boolean if a field has been set.
func (o *QosFlowProfile) HasAdditionalQosFlowInfo() bool {
	if o != nil && !isNil(o.AdditionalQosFlowInfo) {
		return true
	}

	return false
}

// SetAdditionalQosFlowInfo gets a reference to the given AdditionalQosFlowInfo and assigns it to the AdditionalQosFlowInfo field.
func (o *QosFlowProfile) SetAdditionalQosFlowInfo(v AdditionalQosFlowInfo) {
	o.AdditionalQosFlowInfo = &v
}

// GetQosMonitoringReq returns the QosMonitoringReq field value if set, zero value otherwise.
func (o *QosFlowProfile) GetQosMonitoringReq() QosMonitoringReq {
	if o == nil || isNil(o.QosMonitoringReq) {
		var ret QosMonitoringReq
		return ret
	}
	return *o.QosMonitoringReq
}

// GetQosMonitoringReqOk returns a tuple with the QosMonitoringReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetQosMonitoringReqOk() (*QosMonitoringReq, bool) {
	if o == nil || isNil(o.QosMonitoringReq) {
		return nil, false
	}
	return o.QosMonitoringReq, true
}

// HasQosMonitoringReq returns a boolean if a field has been set.
func (o *QosFlowProfile) HasQosMonitoringReq() bool {
	if o != nil && !isNil(o.QosMonitoringReq) {
		return true
	}

	return false
}

// SetQosMonitoringReq gets a reference to the given QosMonitoringReq and assigns it to the QosMonitoringReq field.
func (o *QosFlowProfile) SetQosMonitoringReq(v QosMonitoringReq) {
	o.QosMonitoringReq = &v
}

// GetQosRepPeriod returns the QosRepPeriod field value if set, zero value otherwise.
func (o *QosFlowProfile) GetQosRepPeriod() int32 {
	if o == nil || isNil(o.QosRepPeriod) {
		var ret int32
		return ret
	}
	return *o.QosRepPeriod
}

// GetQosRepPeriodOk returns a tuple with the QosRepPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowProfile) GetQosRepPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.QosRepPeriod) {
		return nil, false
	}
	return o.QosRepPeriod, true
}

// HasQosRepPeriod returns a boolean if a field has been set.
func (o *QosFlowProfile) HasQosRepPeriod() bool {
	if o != nil && !isNil(o.QosRepPeriod) {
		return true
	}

	return false
}

// SetQosRepPeriod gets a reference to the given int32 and assigns it to the QosRepPeriod field.
func (o *QosFlowProfile) SetQosRepPeriod(v int32) {
	o.QosRepPeriod = &v
}

func (o QosFlowProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosFlowProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["5qi"] = o.Var5qi
	if !isNil(o.NonDynamic5Qi) {
		toSerialize["nonDynamic5Qi"] = o.NonDynamic5Qi
	}
	if !isNil(o.Dynamic5Qi) {
		toSerialize["dynamic5Qi"] = o.Dynamic5Qi
	}
	if !isNil(o.Arp) {
		toSerialize["arp"] = o.Arp
	}
	if !isNil(o.GbrQosFlowInfo) {
		toSerialize["gbrQosFlowInfo"] = o.GbrQosFlowInfo
	}
	if !isNil(o.Rqa) {
		toSerialize["rqa"] = o.Rqa
	}
	if !isNil(o.AdditionalQosFlowInfo) {
		toSerialize["additionalQosFlowInfo"] = o.AdditionalQosFlowInfo
	}
	if !isNil(o.QosMonitoringReq) {
		toSerialize["qosMonitoringReq"] = o.QosMonitoringReq
	}
	if !isNil(o.QosRepPeriod) {
		toSerialize["qosRepPeriod"] = o.QosRepPeriod
	}
	return toSerialize, nil
}

type NullableQosFlowProfile struct {
	value *QosFlowProfile
	isSet bool
}

func (v NullableQosFlowProfile) Get() *QosFlowProfile {
	return v.value
}

func (v *NullableQosFlowProfile) Set(val *QosFlowProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableQosFlowProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableQosFlowProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosFlowProfile(val *QosFlowProfile) *NullableQosFlowProfile {
	return &NullableQosFlowProfile{value: val, isSet: true}
}

func (v NullableQosFlowProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosFlowProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



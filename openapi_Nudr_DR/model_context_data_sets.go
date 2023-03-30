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

// checks if the ContextDataSets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextDataSets{}

// ContextDataSets Contains the context data sets.
type ContextDataSets struct {
	Amf3Gpp *Amf3GppAccessRegistration `json:"amf3Gpp,omitempty"`
	AmfNon3Gpp *AmfNon3GppAccessRegistration `json:"amfNon3Gpp,omitempty"`
	SdmSubscriptions []SdmSubscription `json:"sdmSubscriptions,omitempty"`
	EeSubscriptions []EeSubscription `json:"eeSubscriptions,omitempty"`
	Smsf3GppAccess *SmsfRegistration `json:"smsf3GppAccess,omitempty"`
	SmsfNon3GppAccess *SmsfRegistration `json:"smsfNon3GppAccess,omitempty"`
	SubscriptionDataSubscriptions []SubscriptionDataSubscriptions `json:"subscriptionDataSubscriptions,omitempty"`
	// The list of all the SMF registrations of a UE.
	SmfRegistrations []SmfRegistration `json:"smfRegistrations,omitempty"`
	IpSmGw *IpSmGwRegistration `json:"ipSmGw,omitempty"`
	RoamingInfo *RoamingInfoUpdate `json:"roamingInfo,omitempty"`
	PeiInfo *PeiUpdateInfo `json:"peiInfo,omitempty"`
}

// NewContextDataSets instantiates a new ContextDataSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextDataSets() *ContextDataSets {
	this := ContextDataSets{}
	return &this
}

// NewContextDataSetsWithDefaults instantiates a new ContextDataSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextDataSetsWithDefaults() *ContextDataSets {
	this := ContextDataSets{}
	return &this
}

// GetAmf3Gpp returns the Amf3Gpp field value if set, zero value otherwise.
func (o *ContextDataSets) GetAmf3Gpp() Amf3GppAccessRegistration {
	if o == nil || IsNil(o.Amf3Gpp) {
		var ret Amf3GppAccessRegistration
		return ret
	}
	return *o.Amf3Gpp
}

// GetAmf3GppOk returns a tuple with the Amf3Gpp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextDataSets) GetAmf3GppOk() (*Amf3GppAccessRegistration, bool) {
	if o == nil || IsNil(o.Amf3Gpp) {
		return nil, false
	}
	return o.Amf3Gpp, true
}

// HasAmf3Gpp returns a boolean if a field has been set.
func (o *ContextDataSets) HasAmf3Gpp() bool {
	if o != nil && !IsNil(o.Amf3Gpp) {
		return true
	}

	return false
}

// SetAmf3Gpp gets a reference to the given Amf3GppAccessRegistration and assigns it to the Amf3Gpp field.
func (o *ContextDataSets) SetAmf3Gpp(v Amf3GppAccessRegistration) {
	o.Amf3Gpp = &v
}

// GetAmfNon3Gpp returns the AmfNon3Gpp field value if set, zero value otherwise.
func (o *ContextDataSets) GetAmfNon3Gpp() AmfNon3GppAccessRegistration {
	if o == nil || IsNil(o.AmfNon3Gpp) {
		var ret AmfNon3GppAccessRegistration
		return ret
	}
	return *o.AmfNon3Gpp
}

// GetAmfNon3GppOk returns a tuple with the AmfNon3Gpp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextDataSets) GetAmfNon3GppOk() (*AmfNon3GppAccessRegistration, bool) {
	if o == nil || IsNil(o.AmfNon3Gpp) {
		return nil, false
	}
	return o.AmfNon3Gpp, true
}

// HasAmfNon3Gpp returns a boolean if a field has been set.
func (o *ContextDataSets) HasAmfNon3Gpp() bool {
	if o != nil && !IsNil(o.AmfNon3Gpp) {
		return true
	}

	return false
}

// SetAmfNon3Gpp gets a reference to the given AmfNon3GppAccessRegistration and assigns it to the AmfNon3Gpp field.
func (o *ContextDataSets) SetAmfNon3Gpp(v AmfNon3GppAccessRegistration) {
	o.AmfNon3Gpp = &v
}

// GetSdmSubscriptions returns the SdmSubscriptions field value if set, zero value otherwise.
func (o *ContextDataSets) GetSdmSubscriptions() []SdmSubscription {
	if o == nil || IsNil(o.SdmSubscriptions) {
		var ret []SdmSubscription
		return ret
	}
	return o.SdmSubscriptions
}

// GetSdmSubscriptionsOk returns a tuple with the SdmSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextDataSets) GetSdmSubscriptionsOk() ([]SdmSubscription, bool) {
	if o == nil || IsNil(o.SdmSubscriptions) {
		return nil, false
	}
	return o.SdmSubscriptions, true
}

// HasSdmSubscriptions returns a boolean if a field has been set.
func (o *ContextDataSets) HasSdmSubscriptions() bool {
	if o != nil && !IsNil(o.SdmSubscriptions) {
		return true
	}

	return false
}

// SetSdmSubscriptions gets a reference to the given []SdmSubscription and assigns it to the SdmSubscriptions field.
func (o *ContextDataSets) SetSdmSubscriptions(v []SdmSubscription) {
	o.SdmSubscriptions = v
}

// GetEeSubscriptions returns the EeSubscriptions field value if set, zero value otherwise.
func (o *ContextDataSets) GetEeSubscriptions() []EeSubscription {
	if o == nil || IsNil(o.EeSubscriptions) {
		var ret []EeSubscription
		return ret
	}
	return o.EeSubscriptions
}

// GetEeSubscriptionsOk returns a tuple with the EeSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextDataSets) GetEeSubscriptionsOk() ([]EeSubscription, bool) {
	if o == nil || IsNil(o.EeSubscriptions) {
		return nil, false
	}
	return o.EeSubscriptions, true
}

// HasEeSubscriptions returns a boolean if a field has been set.
func (o *ContextDataSets) HasEeSubscriptions() bool {
	if o != nil && !IsNil(o.EeSubscriptions) {
		return true
	}

	return false
}

// SetEeSubscriptions gets a reference to the given []EeSubscription and assigns it to the EeSubscriptions field.
func (o *ContextDataSets) SetEeSubscriptions(v []EeSubscription) {
	o.EeSubscriptions = v
}

// GetSmsf3GppAccess returns the Smsf3GppAccess field value if set, zero value otherwise.
func (o *ContextDataSets) GetSmsf3GppAccess() SmsfRegistration {
	if o == nil || IsNil(o.Smsf3GppAccess) {
		var ret SmsfRegistration
		return ret
	}
	return *o.Smsf3GppAccess
}

// GetSmsf3GppAccessOk returns a tuple with the Smsf3GppAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextDataSets) GetSmsf3GppAccessOk() (*SmsfRegistration, bool) {
	if o == nil || IsNil(o.Smsf3GppAccess) {
		return nil, false
	}
	return o.Smsf3GppAccess, true
}

// HasSmsf3GppAccess returns a boolean if a field has been set.
func (o *ContextDataSets) HasSmsf3GppAccess() bool {
	if o != nil && !IsNil(o.Smsf3GppAccess) {
		return true
	}

	return false
}

// SetSmsf3GppAccess gets a reference to the given SmsfRegistration and assigns it to the Smsf3GppAccess field.
func (o *ContextDataSets) SetSmsf3GppAccess(v SmsfRegistration) {
	o.Smsf3GppAccess = &v
}

// GetSmsfNon3GppAccess returns the SmsfNon3GppAccess field value if set, zero value otherwise.
func (o *ContextDataSets) GetSmsfNon3GppAccess() SmsfRegistration {
	if o == nil || IsNil(o.SmsfNon3GppAccess) {
		var ret SmsfRegistration
		return ret
	}
	return *o.SmsfNon3GppAccess
}

// GetSmsfNon3GppAccessOk returns a tuple with the SmsfNon3GppAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextDataSets) GetSmsfNon3GppAccessOk() (*SmsfRegistration, bool) {
	if o == nil || IsNil(o.SmsfNon3GppAccess) {
		return nil, false
	}
	return o.SmsfNon3GppAccess, true
}

// HasSmsfNon3GppAccess returns a boolean if a field has been set.
func (o *ContextDataSets) HasSmsfNon3GppAccess() bool {
	if o != nil && !IsNil(o.SmsfNon3GppAccess) {
		return true
	}

	return false
}

// SetSmsfNon3GppAccess gets a reference to the given SmsfRegistration and assigns it to the SmsfNon3GppAccess field.
func (o *ContextDataSets) SetSmsfNon3GppAccess(v SmsfRegistration) {
	o.SmsfNon3GppAccess = &v
}

// GetSubscriptionDataSubscriptions returns the SubscriptionDataSubscriptions field value if set, zero value otherwise.
func (o *ContextDataSets) GetSubscriptionDataSubscriptions() []SubscriptionDataSubscriptions {
	if o == nil || IsNil(o.SubscriptionDataSubscriptions) {
		var ret []SubscriptionDataSubscriptions
		return ret
	}
	return o.SubscriptionDataSubscriptions
}

// GetSubscriptionDataSubscriptionsOk returns a tuple with the SubscriptionDataSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextDataSets) GetSubscriptionDataSubscriptionsOk() ([]SubscriptionDataSubscriptions, bool) {
	if o == nil || IsNil(o.SubscriptionDataSubscriptions) {
		return nil, false
	}
	return o.SubscriptionDataSubscriptions, true
}

// HasSubscriptionDataSubscriptions returns a boolean if a field has been set.
func (o *ContextDataSets) HasSubscriptionDataSubscriptions() bool {
	if o != nil && !IsNil(o.SubscriptionDataSubscriptions) {
		return true
	}

	return false
}

// SetSubscriptionDataSubscriptions gets a reference to the given []SubscriptionDataSubscriptions and assigns it to the SubscriptionDataSubscriptions field.
func (o *ContextDataSets) SetSubscriptionDataSubscriptions(v []SubscriptionDataSubscriptions) {
	o.SubscriptionDataSubscriptions = v
}

// GetSmfRegistrations returns the SmfRegistrations field value if set, zero value otherwise.
func (o *ContextDataSets) GetSmfRegistrations() []SmfRegistration {
	if o == nil || IsNil(o.SmfRegistrations) {
		var ret []SmfRegistration
		return ret
	}
	return o.SmfRegistrations
}

// GetSmfRegistrationsOk returns a tuple with the SmfRegistrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextDataSets) GetSmfRegistrationsOk() ([]SmfRegistration, bool) {
	if o == nil || IsNil(o.SmfRegistrations) {
		return nil, false
	}
	return o.SmfRegistrations, true
}

// HasSmfRegistrations returns a boolean if a field has been set.
func (o *ContextDataSets) HasSmfRegistrations() bool {
	if o != nil && !IsNil(o.SmfRegistrations) {
		return true
	}

	return false
}

// SetSmfRegistrations gets a reference to the given []SmfRegistration and assigns it to the SmfRegistrations field.
func (o *ContextDataSets) SetSmfRegistrations(v []SmfRegistration) {
	o.SmfRegistrations = v
}

// GetIpSmGw returns the IpSmGw field value if set, zero value otherwise.
func (o *ContextDataSets) GetIpSmGw() IpSmGwRegistration {
	if o == nil || IsNil(o.IpSmGw) {
		var ret IpSmGwRegistration
		return ret
	}
	return *o.IpSmGw
}

// GetIpSmGwOk returns a tuple with the IpSmGw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextDataSets) GetIpSmGwOk() (*IpSmGwRegistration, bool) {
	if o == nil || IsNil(o.IpSmGw) {
		return nil, false
	}
	return o.IpSmGw, true
}

// HasIpSmGw returns a boolean if a field has been set.
func (o *ContextDataSets) HasIpSmGw() bool {
	if o != nil && !IsNil(o.IpSmGw) {
		return true
	}

	return false
}

// SetIpSmGw gets a reference to the given IpSmGwRegistration and assigns it to the IpSmGw field.
func (o *ContextDataSets) SetIpSmGw(v IpSmGwRegistration) {
	o.IpSmGw = &v
}

// GetRoamingInfo returns the RoamingInfo field value if set, zero value otherwise.
func (o *ContextDataSets) GetRoamingInfo() RoamingInfoUpdate {
	if o == nil || IsNil(o.RoamingInfo) {
		var ret RoamingInfoUpdate
		return ret
	}
	return *o.RoamingInfo
}

// GetRoamingInfoOk returns a tuple with the RoamingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextDataSets) GetRoamingInfoOk() (*RoamingInfoUpdate, bool) {
	if o == nil || IsNil(o.RoamingInfo) {
		return nil, false
	}
	return o.RoamingInfo, true
}

// HasRoamingInfo returns a boolean if a field has been set.
func (o *ContextDataSets) HasRoamingInfo() bool {
	if o != nil && !IsNil(o.RoamingInfo) {
		return true
	}

	return false
}

// SetRoamingInfo gets a reference to the given RoamingInfoUpdate and assigns it to the RoamingInfo field.
func (o *ContextDataSets) SetRoamingInfo(v RoamingInfoUpdate) {
	o.RoamingInfo = &v
}

// GetPeiInfo returns the PeiInfo field value if set, zero value otherwise.
func (o *ContextDataSets) GetPeiInfo() PeiUpdateInfo {
	if o == nil || IsNil(o.PeiInfo) {
		var ret PeiUpdateInfo
		return ret
	}
	return *o.PeiInfo
}

// GetPeiInfoOk returns a tuple with the PeiInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextDataSets) GetPeiInfoOk() (*PeiUpdateInfo, bool) {
	if o == nil || IsNil(o.PeiInfo) {
		return nil, false
	}
	return o.PeiInfo, true
}

// HasPeiInfo returns a boolean if a field has been set.
func (o *ContextDataSets) HasPeiInfo() bool {
	if o != nil && !IsNil(o.PeiInfo) {
		return true
	}

	return false
}

// SetPeiInfo gets a reference to the given PeiUpdateInfo and assigns it to the PeiInfo field.
func (o *ContextDataSets) SetPeiInfo(v PeiUpdateInfo) {
	o.PeiInfo = &v
}

func (o ContextDataSets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextDataSets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amf3Gpp) {
		toSerialize["amf3Gpp"] = o.Amf3Gpp
	}
	if !IsNil(o.AmfNon3Gpp) {
		toSerialize["amfNon3Gpp"] = o.AmfNon3Gpp
	}
	if !IsNil(o.SdmSubscriptions) {
		toSerialize["sdmSubscriptions"] = o.SdmSubscriptions
	}
	if !IsNil(o.EeSubscriptions) {
		toSerialize["eeSubscriptions"] = o.EeSubscriptions
	}
	if !IsNil(o.Smsf3GppAccess) {
		toSerialize["smsf3GppAccess"] = o.Smsf3GppAccess
	}
	if !IsNil(o.SmsfNon3GppAccess) {
		toSerialize["smsfNon3GppAccess"] = o.SmsfNon3GppAccess
	}
	if !IsNil(o.SubscriptionDataSubscriptions) {
		toSerialize["subscriptionDataSubscriptions"] = o.SubscriptionDataSubscriptions
	}
	if !IsNil(o.SmfRegistrations) {
		toSerialize["smfRegistrations"] = o.SmfRegistrations
	}
	if !IsNil(o.IpSmGw) {
		toSerialize["ipSmGw"] = o.IpSmGw
	}
	if !IsNil(o.RoamingInfo) {
		toSerialize["roamingInfo"] = o.RoamingInfo
	}
	if !IsNil(o.PeiInfo) {
		toSerialize["peiInfo"] = o.PeiInfo
	}
	return toSerialize, nil
}

type NullableContextDataSets struct {
	value *ContextDataSets
	isSet bool
}

func (v NullableContextDataSets) Get() *ContextDataSets {
	return v.value
}

func (v *NullableContextDataSets) Set(val *ContextDataSets) {
	v.value = val
	v.isSet = true
}

func (v NullableContextDataSets) IsSet() bool {
	return v.isSet
}

func (v *NullableContextDataSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextDataSets(val *ContextDataSets) *NullableContextDataSets {
	return &NullableContextDataSets{value: val, isSet: true}
}

func (v NullableContextDataSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextDataSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



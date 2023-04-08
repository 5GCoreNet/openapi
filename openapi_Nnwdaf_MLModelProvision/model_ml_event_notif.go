/*
Nnwdaf_MLModelProvision

Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_MLModelProvision

import (
	"encoding/json"
)

// checks if the MLEventNotif type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MLEventNotif{}

// MLEventNotif Represents a notification related to a single event that occurred.
type MLEventNotif struct {
	Event NwdafEvent `json:"event"`
	NotifCorreId *string `json:"notifCorreId,omitempty"`
	MLFileAddr MLModelAddr `json:"mLFileAddr"`
	ValidityPeriod *TimeWindow `json:"validityPeriod,omitempty"`
	SpatialValidity *NetworkAreaInfo `json:"spatialValidity,omitempty"`
}

// NewMLEventNotif instantiates a new MLEventNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMLEventNotif(event NwdafEvent, mLFileAddr MLModelAddr) *MLEventNotif {
	this := MLEventNotif{}
	this.Event = event
	this.MLFileAddr = mLFileAddr
	return &this
}

// NewMLEventNotifWithDefaults instantiates a new MLEventNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMLEventNotifWithDefaults() *MLEventNotif {
	this := MLEventNotif{}
	return &this
}

// GetEvent returns the Event field value
func (o *MLEventNotif) GetEvent() NwdafEvent {
	if o == nil {
		var ret NwdafEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *MLEventNotif) GetEventOk() (*NwdafEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *MLEventNotif) SetEvent(v NwdafEvent) {
	o.Event = v
}

// GetNotifCorreId returns the NotifCorreId field value if set, zero value otherwise.
func (o *MLEventNotif) GetNotifCorreId() string {
	if o == nil || isNil(o.NotifCorreId) {
		var ret string
		return ret
	}
	return *o.NotifCorreId
}

// GetNotifCorreIdOk returns a tuple with the NotifCorreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLEventNotif) GetNotifCorreIdOk() (*string, bool) {
	if o == nil || isNil(o.NotifCorreId) {
		return nil, false
	}
	return o.NotifCorreId, true
}

// HasNotifCorreId returns a boolean if a field has been set.
func (o *MLEventNotif) HasNotifCorreId() bool {
	if o != nil && !isNil(o.NotifCorreId) {
		return true
	}

	return false
}

// SetNotifCorreId gets a reference to the given string and assigns it to the NotifCorreId field.
func (o *MLEventNotif) SetNotifCorreId(v string) {
	o.NotifCorreId = &v
}

// GetMLFileAddr returns the MLFileAddr field value
func (o *MLEventNotif) GetMLFileAddr() MLModelAddr {
	if o == nil {
		var ret MLModelAddr
		return ret
	}

	return o.MLFileAddr
}

// GetMLFileAddrOk returns a tuple with the MLFileAddr field value
// and a boolean to check if the value has been set.
func (o *MLEventNotif) GetMLFileAddrOk() (*MLModelAddr, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MLFileAddr, true
}

// SetMLFileAddr sets field value
func (o *MLEventNotif) SetMLFileAddr(v MLModelAddr) {
	o.MLFileAddr = v
}

// GetValidityPeriod returns the ValidityPeriod field value if set, zero value otherwise.
func (o *MLEventNotif) GetValidityPeriod() TimeWindow {
	if o == nil || isNil(o.ValidityPeriod) {
		var ret TimeWindow
		return ret
	}
	return *o.ValidityPeriod
}

// GetValidityPeriodOk returns a tuple with the ValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLEventNotif) GetValidityPeriodOk() (*TimeWindow, bool) {
	if o == nil || isNil(o.ValidityPeriod) {
		return nil, false
	}
	return o.ValidityPeriod, true
}

// HasValidityPeriod returns a boolean if a field has been set.
func (o *MLEventNotif) HasValidityPeriod() bool {
	if o != nil && !isNil(o.ValidityPeriod) {
		return true
	}

	return false
}

// SetValidityPeriod gets a reference to the given TimeWindow and assigns it to the ValidityPeriod field.
func (o *MLEventNotif) SetValidityPeriod(v TimeWindow) {
	o.ValidityPeriod = &v
}

// GetSpatialValidity returns the SpatialValidity field value if set, zero value otherwise.
func (o *MLEventNotif) GetSpatialValidity() NetworkAreaInfo {
	if o == nil || isNil(o.SpatialValidity) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.SpatialValidity
}

// GetSpatialValidityOk returns a tuple with the SpatialValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLEventNotif) GetSpatialValidityOk() (*NetworkAreaInfo, bool) {
	if o == nil || isNil(o.SpatialValidity) {
		return nil, false
	}
	return o.SpatialValidity, true
}

// HasSpatialValidity returns a boolean if a field has been set.
func (o *MLEventNotif) HasSpatialValidity() bool {
	if o != nil && !isNil(o.SpatialValidity) {
		return true
	}

	return false
}

// SetSpatialValidity gets a reference to the given NetworkAreaInfo and assigns it to the SpatialValidity field.
func (o *MLEventNotif) SetSpatialValidity(v NetworkAreaInfo) {
	o.SpatialValidity = &v
}

func (o MLEventNotif) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MLEventNotif) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if !isNil(o.NotifCorreId) {
		toSerialize["notifCorreId"] = o.NotifCorreId
	}
	toSerialize["mLFileAddr"] = o.MLFileAddr
	if !isNil(o.ValidityPeriod) {
		toSerialize["validityPeriod"] = o.ValidityPeriod
	}
	if !isNil(o.SpatialValidity) {
		toSerialize["spatialValidity"] = o.SpatialValidity
	}
	return toSerialize, nil
}

type NullableMLEventNotif struct {
	value *MLEventNotif
	isSet bool
}

func (v NullableMLEventNotif) Get() *MLEventNotif {
	return v.value
}

func (v *NullableMLEventNotif) Set(val *MLEventNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableMLEventNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableMLEventNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMLEventNotif(val *MLEventNotif) *NullableMLEventNotif {
	return &NullableMLEventNotif{value: val, isSet: true}
}

func (v NullableMLEventNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMLEventNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



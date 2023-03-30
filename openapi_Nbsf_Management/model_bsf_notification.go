/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.4.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsf_Management

import (
	"encoding/json"
)

// checks if the BsfNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BsfNotification{}

// BsfNotification Contains the event notifications.
type BsfNotification struct {
	// Notification Correlation ID assigned by the NF service consumer.
	NotifCorreId string `json:"notifCorreId"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	PcfId *string `json:"pcfId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	PcfSetId *string `json:"pcfSetId,omitempty"`
	BindLevel *BindingLevel `json:"bindLevel,omitempty"`
	// Notifications about Individual Events.
	EventNotifs []BsfEventNotification `json:"eventNotifs"`
}

// NewBsfNotification instantiates a new BsfNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBsfNotification(notifCorreId string, eventNotifs []BsfEventNotification) *BsfNotification {
	this := BsfNotification{}
	this.NotifCorreId = notifCorreId
	this.EventNotifs = eventNotifs
	return &this
}

// NewBsfNotificationWithDefaults instantiates a new BsfNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBsfNotificationWithDefaults() *BsfNotification {
	this := BsfNotification{}
	return &this
}

// GetNotifCorreId returns the NotifCorreId field value
func (o *BsfNotification) GetNotifCorreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifCorreId
}

// GetNotifCorreIdOk returns a tuple with the NotifCorreId field value
// and a boolean to check if the value has been set.
func (o *BsfNotification) GetNotifCorreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifCorreId, true
}

// SetNotifCorreId sets field value
func (o *BsfNotification) SetNotifCorreId(v string) {
	o.NotifCorreId = v
}

// GetPcfId returns the PcfId field value if set, zero value otherwise.
func (o *BsfNotification) GetPcfId() string {
	if o == nil || IsNil(o.PcfId) {
		var ret string
		return ret
	}
	return *o.PcfId
}

// GetPcfIdOk returns a tuple with the PcfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfNotification) GetPcfIdOk() (*string, bool) {
	if o == nil || IsNil(o.PcfId) {
		return nil, false
	}
	return o.PcfId, true
}

// HasPcfId returns a boolean if a field has been set.
func (o *BsfNotification) HasPcfId() bool {
	if o != nil && !IsNil(o.PcfId) {
		return true
	}

	return false
}

// SetPcfId gets a reference to the given string and assigns it to the PcfId field.
func (o *BsfNotification) SetPcfId(v string) {
	o.PcfId = &v
}

// GetPcfSetId returns the PcfSetId field value if set, zero value otherwise.
func (o *BsfNotification) GetPcfSetId() string {
	if o == nil || IsNil(o.PcfSetId) {
		var ret string
		return ret
	}
	return *o.PcfSetId
}

// GetPcfSetIdOk returns a tuple with the PcfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfNotification) GetPcfSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.PcfSetId) {
		return nil, false
	}
	return o.PcfSetId, true
}

// HasPcfSetId returns a boolean if a field has been set.
func (o *BsfNotification) HasPcfSetId() bool {
	if o != nil && !IsNil(o.PcfSetId) {
		return true
	}

	return false
}

// SetPcfSetId gets a reference to the given string and assigns it to the PcfSetId field.
func (o *BsfNotification) SetPcfSetId(v string) {
	o.PcfSetId = &v
}

// GetBindLevel returns the BindLevel field value if set, zero value otherwise.
func (o *BsfNotification) GetBindLevel() BindingLevel {
	if o == nil || IsNil(o.BindLevel) {
		var ret BindingLevel
		return ret
	}
	return *o.BindLevel
}

// GetBindLevelOk returns a tuple with the BindLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfNotification) GetBindLevelOk() (*BindingLevel, bool) {
	if o == nil || IsNil(o.BindLevel) {
		return nil, false
	}
	return o.BindLevel, true
}

// HasBindLevel returns a boolean if a field has been set.
func (o *BsfNotification) HasBindLevel() bool {
	if o != nil && !IsNil(o.BindLevel) {
		return true
	}

	return false
}

// SetBindLevel gets a reference to the given BindingLevel and assigns it to the BindLevel field.
func (o *BsfNotification) SetBindLevel(v BindingLevel) {
	o.BindLevel = &v
}

// GetEventNotifs returns the EventNotifs field value
func (o *BsfNotification) GetEventNotifs() []BsfEventNotification {
	if o == nil {
		var ret []BsfEventNotification
		return ret
	}

	return o.EventNotifs
}

// GetEventNotifsOk returns a tuple with the EventNotifs field value
// and a boolean to check if the value has been set.
func (o *BsfNotification) GetEventNotifsOk() ([]BsfEventNotification, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventNotifs, true
}

// SetEventNotifs sets field value
func (o *BsfNotification) SetEventNotifs(v []BsfEventNotification) {
	o.EventNotifs = v
}

func (o BsfNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BsfNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifCorreId"] = o.NotifCorreId
	if !IsNil(o.PcfId) {
		toSerialize["pcfId"] = o.PcfId
	}
	if !IsNil(o.PcfSetId) {
		toSerialize["pcfSetId"] = o.PcfSetId
	}
	if !IsNil(o.BindLevel) {
		toSerialize["bindLevel"] = o.BindLevel
	}
	toSerialize["eventNotifs"] = o.EventNotifs
	return toSerialize, nil
}

type NullableBsfNotification struct {
	value *BsfNotification
	isSet bool
}

func (v NullableBsfNotification) Get() *BsfNotification {
	return v.value
}

func (v *NullableBsfNotification) Set(val *BsfNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableBsfNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableBsfNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBsfNotification(val *BsfNotification) *NullableBsfNotification {
	return &NullableBsfNotification{value: val, isSet: true}
}

func (v NullableBsfNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBsfNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the ExtAmfEventSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtAmfEventSubscription{}

// ExtAmfEventSubscription AMF event subscription extended with additional information received for the subscription
type ExtAmfEventSubscription struct {
	AmfEventSubscription
	BindingInfo       []string `json:"bindingInfo,omitempty"`
	SubscribingNfType *NFType  `json:"subscribingNfType,omitempty"`
	EventSyncInd      *bool    `json:"eventSyncInd,omitempty"`
	NfConsumerInfo    []string `json:"nfConsumerInfo,omitempty"`
	// Map of subscribed Area of Interest (AoI) Event State in the old AMF. The JSON pointer to an AmfEventArea element in the areaList IE (or a PresenceInfo element in  presenceInfoList IE) of the AmfEvent data type shall be the key of the map.
	AoiStateList *map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
}

// NewExtAmfEventSubscription instantiates a new ExtAmfEventSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtAmfEventSubscription(eventList []AmfEvent, eventNotifyUri string, notifyCorrelationId string, nfId string) *ExtAmfEventSubscription {
	this := ExtAmfEventSubscription{}
	this.EventList = eventList
	this.EventNotifyUri = eventNotifyUri
	this.NotifyCorrelationId = notifyCorrelationId
	this.NfId = nfId
	return &this
}

// NewExtAmfEventSubscriptionWithDefaults instantiates a new ExtAmfEventSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtAmfEventSubscriptionWithDefaults() *ExtAmfEventSubscription {
	this := ExtAmfEventSubscription{}
	return &this
}

// GetBindingInfo returns the BindingInfo field value if set, zero value otherwise.
func (o *ExtAmfEventSubscription) GetBindingInfo() []string {
	if o == nil || IsNil(o.BindingInfo) {
		var ret []string
		return ret
	}
	return o.BindingInfo
}

// GetBindingInfoOk returns a tuple with the BindingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtAmfEventSubscription) GetBindingInfoOk() ([]string, bool) {
	if o == nil || IsNil(o.BindingInfo) {
		return nil, false
	}
	return o.BindingInfo, true
}

// HasBindingInfo returns a boolean if a field has been set.
func (o *ExtAmfEventSubscription) HasBindingInfo() bool {
	if o != nil && !IsNil(o.BindingInfo) {
		return true
	}

	return false
}

// SetBindingInfo gets a reference to the given []string and assigns it to the BindingInfo field.
func (o *ExtAmfEventSubscription) SetBindingInfo(v []string) {
	o.BindingInfo = v
}

// GetSubscribingNfType returns the SubscribingNfType field value if set, zero value otherwise.
func (o *ExtAmfEventSubscription) GetSubscribingNfType() NFType {
	if o == nil || IsNil(o.SubscribingNfType) {
		var ret NFType
		return ret
	}
	return *o.SubscribingNfType
}

// GetSubscribingNfTypeOk returns a tuple with the SubscribingNfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtAmfEventSubscription) GetSubscribingNfTypeOk() (*NFType, bool) {
	if o == nil || IsNil(o.SubscribingNfType) {
		return nil, false
	}
	return o.SubscribingNfType, true
}

// HasSubscribingNfType returns a boolean if a field has been set.
func (o *ExtAmfEventSubscription) HasSubscribingNfType() bool {
	if o != nil && !IsNil(o.SubscribingNfType) {
		return true
	}

	return false
}

// SetSubscribingNfType gets a reference to the given NFType and assigns it to the SubscribingNfType field.
func (o *ExtAmfEventSubscription) SetSubscribingNfType(v NFType) {
	o.SubscribingNfType = &v
}

// GetEventSyncInd returns the EventSyncInd field value if set, zero value otherwise.
func (o *ExtAmfEventSubscription) GetEventSyncInd() bool {
	if o == nil || IsNil(o.EventSyncInd) {
		var ret bool
		return ret
	}
	return *o.EventSyncInd
}

// GetEventSyncIndOk returns a tuple with the EventSyncInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtAmfEventSubscription) GetEventSyncIndOk() (*bool, bool) {
	if o == nil || IsNil(o.EventSyncInd) {
		return nil, false
	}
	return o.EventSyncInd, true
}

// HasEventSyncInd returns a boolean if a field has been set.
func (o *ExtAmfEventSubscription) HasEventSyncInd() bool {
	if o != nil && !IsNil(o.EventSyncInd) {
		return true
	}

	return false
}

// SetEventSyncInd gets a reference to the given bool and assigns it to the EventSyncInd field.
func (o *ExtAmfEventSubscription) SetEventSyncInd(v bool) {
	o.EventSyncInd = &v
}

// GetNfConsumerInfo returns the NfConsumerInfo field value if set, zero value otherwise.
func (o *ExtAmfEventSubscription) GetNfConsumerInfo() []string {
	if o == nil || IsNil(o.NfConsumerInfo) {
		var ret []string
		return ret
	}
	return o.NfConsumerInfo
}

// GetNfConsumerInfoOk returns a tuple with the NfConsumerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtAmfEventSubscription) GetNfConsumerInfoOk() ([]string, bool) {
	if o == nil || IsNil(o.NfConsumerInfo) {
		return nil, false
	}
	return o.NfConsumerInfo, true
}

// HasNfConsumerInfo returns a boolean if a field has been set.
func (o *ExtAmfEventSubscription) HasNfConsumerInfo() bool {
	if o != nil && !IsNil(o.NfConsumerInfo) {
		return true
	}

	return false
}

// SetNfConsumerInfo gets a reference to the given []string and assigns it to the NfConsumerInfo field.
func (o *ExtAmfEventSubscription) SetNfConsumerInfo(v []string) {
	o.NfConsumerInfo = v
}

// GetAoiStateList returns the AoiStateList field value if set, zero value otherwise.
func (o *ExtAmfEventSubscription) GetAoiStateList() map[string]AreaOfInterestEventState {
	if o == nil || IsNil(o.AoiStateList) {
		var ret map[string]AreaOfInterestEventState
		return ret
	}
	return *o.AoiStateList
}

// GetAoiStateListOk returns a tuple with the AoiStateList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtAmfEventSubscription) GetAoiStateListOk() (*map[string]AreaOfInterestEventState, bool) {
	if o == nil || IsNil(o.AoiStateList) {
		return nil, false
	}
	return o.AoiStateList, true
}

// HasAoiStateList returns a boolean if a field has been set.
func (o *ExtAmfEventSubscription) HasAoiStateList() bool {
	if o != nil && !IsNil(o.AoiStateList) {
		return true
	}

	return false
}

// SetAoiStateList gets a reference to the given map[string]AreaOfInterestEventState and assigns it to the AoiStateList field.
func (o *ExtAmfEventSubscription) SetAoiStateList(v map[string]AreaOfInterestEventState) {
	o.AoiStateList = &v
}

func (o ExtAmfEventSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtAmfEventSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAmfEventSubscription, errAmfEventSubscription := json.Marshal(o.AmfEventSubscription)
	if errAmfEventSubscription != nil {
		return map[string]interface{}{}, errAmfEventSubscription
	}
	errAmfEventSubscription = json.Unmarshal([]byte(serializedAmfEventSubscription), &toSerialize)
	if errAmfEventSubscription != nil {
		return map[string]interface{}{}, errAmfEventSubscription
	}
	if !IsNil(o.BindingInfo) {
		toSerialize["bindingInfo"] = o.BindingInfo
	}
	if !IsNil(o.SubscribingNfType) {
		toSerialize["subscribingNfType"] = o.SubscribingNfType
	}
	if !IsNil(o.EventSyncInd) {
		toSerialize["eventSyncInd"] = o.EventSyncInd
	}
	if !IsNil(o.NfConsumerInfo) {
		toSerialize["nfConsumerInfo"] = o.NfConsumerInfo
	}
	if !IsNil(o.AoiStateList) {
		toSerialize["aoiStateList"] = o.AoiStateList
	}
	return toSerialize, nil
}

type NullableExtAmfEventSubscription struct {
	value *ExtAmfEventSubscription
	isSet bool
}

func (v NullableExtAmfEventSubscription) Get() *ExtAmfEventSubscription {
	return v.value
}

func (v *NullableExtAmfEventSubscription) Set(val *ExtAmfEventSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableExtAmfEventSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableExtAmfEventSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtAmfEventSubscription(val *ExtAmfEventSubscription) *NullableExtAmfEventSubscription {
	return &NullableExtAmfEventSubscription{value: val, isSet: true}
}

func (v NullableExtAmfEventSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtAmfEventSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the DataChangeNotify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataChangeNotify{}

// DataChangeNotify Container for data which have changed and notification was requested when changed.
type DataChangeNotify struct {
	OriginalCallbackReference []string `json:"originalCallbackReference,omitempty"`
	// String represents the SUPI or GPSI
	UeId *string `json:"ueId,omitempty"`
	NotifyItems []NotifyItem `json:"notifyItems,omitempty"`
	SdmSubscription *SdmSubscription1 `json:"sdmSubscription,omitempty"`
	AdditionalSdmSubscriptions []SdmSubscription1 `json:"additionalSdmSubscriptions,omitempty"`
	SubscriptionDataSubscriptions []SubscriptionDataSubscriptions `json:"subscriptionDataSubscriptions,omitempty"`
}

// NewDataChangeNotify instantiates a new DataChangeNotify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataChangeNotify() *DataChangeNotify {
	this := DataChangeNotify{}
	return &this
}

// NewDataChangeNotifyWithDefaults instantiates a new DataChangeNotify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataChangeNotifyWithDefaults() *DataChangeNotify {
	this := DataChangeNotify{}
	return &this
}

// GetOriginalCallbackReference returns the OriginalCallbackReference field value if set, zero value otherwise.
func (o *DataChangeNotify) GetOriginalCallbackReference() []string {
	if o == nil || isNil(o.OriginalCallbackReference) {
		var ret []string
		return ret
	}
	return o.OriginalCallbackReference
}

// GetOriginalCallbackReferenceOk returns a tuple with the OriginalCallbackReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChangeNotify) GetOriginalCallbackReferenceOk() ([]string, bool) {
	if o == nil || isNil(o.OriginalCallbackReference) {
		return nil, false
	}
	return o.OriginalCallbackReference, true
}

// HasOriginalCallbackReference returns a boolean if a field has been set.
func (o *DataChangeNotify) HasOriginalCallbackReference() bool {
	if o != nil && !isNil(o.OriginalCallbackReference) {
		return true
	}

	return false
}

// SetOriginalCallbackReference gets a reference to the given []string and assigns it to the OriginalCallbackReference field.
func (o *DataChangeNotify) SetOriginalCallbackReference(v []string) {
	o.OriginalCallbackReference = v
}

// GetUeId returns the UeId field value if set, zero value otherwise.
func (o *DataChangeNotify) GetUeId() string {
	if o == nil || isNil(o.UeId) {
		var ret string
		return ret
	}
	return *o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChangeNotify) GetUeIdOk() (*string, bool) {
	if o == nil || isNil(o.UeId) {
		return nil, false
	}
	return o.UeId, true
}

// HasUeId returns a boolean if a field has been set.
func (o *DataChangeNotify) HasUeId() bool {
	if o != nil && !isNil(o.UeId) {
		return true
	}

	return false
}

// SetUeId gets a reference to the given string and assigns it to the UeId field.
func (o *DataChangeNotify) SetUeId(v string) {
	o.UeId = &v
}

// GetNotifyItems returns the NotifyItems field value if set, zero value otherwise.
func (o *DataChangeNotify) GetNotifyItems() []NotifyItem {
	if o == nil || isNil(o.NotifyItems) {
		var ret []NotifyItem
		return ret
	}
	return o.NotifyItems
}

// GetNotifyItemsOk returns a tuple with the NotifyItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChangeNotify) GetNotifyItemsOk() ([]NotifyItem, bool) {
	if o == nil || isNil(o.NotifyItems) {
		return nil, false
	}
	return o.NotifyItems, true
}

// HasNotifyItems returns a boolean if a field has been set.
func (o *DataChangeNotify) HasNotifyItems() bool {
	if o != nil && !isNil(o.NotifyItems) {
		return true
	}

	return false
}

// SetNotifyItems gets a reference to the given []NotifyItem and assigns it to the NotifyItems field.
func (o *DataChangeNotify) SetNotifyItems(v []NotifyItem) {
	o.NotifyItems = v
}

// GetSdmSubscription returns the SdmSubscription field value if set, zero value otherwise.
func (o *DataChangeNotify) GetSdmSubscription() SdmSubscription1 {
	if o == nil || isNil(o.SdmSubscription) {
		var ret SdmSubscription1
		return ret
	}
	return *o.SdmSubscription
}

// GetSdmSubscriptionOk returns a tuple with the SdmSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChangeNotify) GetSdmSubscriptionOk() (*SdmSubscription1, bool) {
	if o == nil || isNil(o.SdmSubscription) {
		return nil, false
	}
	return o.SdmSubscription, true
}

// HasSdmSubscription returns a boolean if a field has been set.
func (o *DataChangeNotify) HasSdmSubscription() bool {
	if o != nil && !isNil(o.SdmSubscription) {
		return true
	}

	return false
}

// SetSdmSubscription gets a reference to the given SdmSubscription1 and assigns it to the SdmSubscription field.
func (o *DataChangeNotify) SetSdmSubscription(v SdmSubscription1) {
	o.SdmSubscription = &v
}

// GetAdditionalSdmSubscriptions returns the AdditionalSdmSubscriptions field value if set, zero value otherwise.
func (o *DataChangeNotify) GetAdditionalSdmSubscriptions() []SdmSubscription1 {
	if o == nil || isNil(o.AdditionalSdmSubscriptions) {
		var ret []SdmSubscription1
		return ret
	}
	return o.AdditionalSdmSubscriptions
}

// GetAdditionalSdmSubscriptionsOk returns a tuple with the AdditionalSdmSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChangeNotify) GetAdditionalSdmSubscriptionsOk() ([]SdmSubscription1, bool) {
	if o == nil || isNil(o.AdditionalSdmSubscriptions) {
		return nil, false
	}
	return o.AdditionalSdmSubscriptions, true
}

// HasAdditionalSdmSubscriptions returns a boolean if a field has been set.
func (o *DataChangeNotify) HasAdditionalSdmSubscriptions() bool {
	if o != nil && !isNil(o.AdditionalSdmSubscriptions) {
		return true
	}

	return false
}

// SetAdditionalSdmSubscriptions gets a reference to the given []SdmSubscription1 and assigns it to the AdditionalSdmSubscriptions field.
func (o *DataChangeNotify) SetAdditionalSdmSubscriptions(v []SdmSubscription1) {
	o.AdditionalSdmSubscriptions = v
}

// GetSubscriptionDataSubscriptions returns the SubscriptionDataSubscriptions field value if set, zero value otherwise.
func (o *DataChangeNotify) GetSubscriptionDataSubscriptions() []SubscriptionDataSubscriptions {
	if o == nil || isNil(o.SubscriptionDataSubscriptions) {
		var ret []SubscriptionDataSubscriptions
		return ret
	}
	return o.SubscriptionDataSubscriptions
}

// GetSubscriptionDataSubscriptionsOk returns a tuple with the SubscriptionDataSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChangeNotify) GetSubscriptionDataSubscriptionsOk() ([]SubscriptionDataSubscriptions, bool) {
	if o == nil || isNil(o.SubscriptionDataSubscriptions) {
		return nil, false
	}
	return o.SubscriptionDataSubscriptions, true
}

// HasSubscriptionDataSubscriptions returns a boolean if a field has been set.
func (o *DataChangeNotify) HasSubscriptionDataSubscriptions() bool {
	if o != nil && !isNil(o.SubscriptionDataSubscriptions) {
		return true
	}

	return false
}

// SetSubscriptionDataSubscriptions gets a reference to the given []SubscriptionDataSubscriptions and assigns it to the SubscriptionDataSubscriptions field.
func (o *DataChangeNotify) SetSubscriptionDataSubscriptions(v []SubscriptionDataSubscriptions) {
	o.SubscriptionDataSubscriptions = v
}

func (o DataChangeNotify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataChangeNotify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OriginalCallbackReference) {
		toSerialize["originalCallbackReference"] = o.OriginalCallbackReference
	}
	if !isNil(o.UeId) {
		toSerialize["ueId"] = o.UeId
	}
	if !isNil(o.NotifyItems) {
		toSerialize["notifyItems"] = o.NotifyItems
	}
	if !isNil(o.SdmSubscription) {
		toSerialize["sdmSubscription"] = o.SdmSubscription
	}
	if !isNil(o.AdditionalSdmSubscriptions) {
		toSerialize["additionalSdmSubscriptions"] = o.AdditionalSdmSubscriptions
	}
	if !isNil(o.SubscriptionDataSubscriptions) {
		toSerialize["subscriptionDataSubscriptions"] = o.SubscriptionDataSubscriptions
	}
	return toSerialize, nil
}

type NullableDataChangeNotify struct {
	value *DataChangeNotify
	isSet bool
}

func (v NullableDataChangeNotify) Get() *DataChangeNotify {
	return v.value
}

func (v *NullableDataChangeNotify) Set(val *DataChangeNotify) {
	v.value = val
	v.isSet = true
}

func (v NullableDataChangeNotify) IsSet() bool {
	return v.isSet
}

func (v *NullableDataChangeNotify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataChangeNotify(val *DataChangeNotify) *NullableDataChangeNotify {
	return &NullableDataChangeNotify{value: val, isSet: true}
}

func (v NullableDataChangeNotify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataChangeNotify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



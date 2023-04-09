/*
3gpp-nidd

API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NIDD

import (
	"encoding/json"
)

// checks if the GmdNiddDownlinkDataDeliveryNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GmdNiddDownlinkDataDeliveryNotification{}

// GmdNiddDownlinkDataDeliveryNotification Represents the delivery status of a specific group NIDD downlink data delivery.
type GmdNiddDownlinkDataDeliveryNotification struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NiddDownlinkDataTransfer string `json:"niddDownlinkDataTransfer"`
	// Indicates the group message delivery result.
	GmdResults []GmdResult `json:"gmdResults"`
}

// NewGmdNiddDownlinkDataDeliveryNotification instantiates a new GmdNiddDownlinkDataDeliveryNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGmdNiddDownlinkDataDeliveryNotification(niddDownlinkDataTransfer string, gmdResults []GmdResult) *GmdNiddDownlinkDataDeliveryNotification {
	this := GmdNiddDownlinkDataDeliveryNotification{}
	this.NiddDownlinkDataTransfer = niddDownlinkDataTransfer
	this.GmdResults = gmdResults
	return &this
}

// NewGmdNiddDownlinkDataDeliveryNotificationWithDefaults instantiates a new GmdNiddDownlinkDataDeliveryNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGmdNiddDownlinkDataDeliveryNotificationWithDefaults() *GmdNiddDownlinkDataDeliveryNotification {
	this := GmdNiddDownlinkDataDeliveryNotification{}
	return &this
}

// GetNiddDownlinkDataTransfer returns the NiddDownlinkDataTransfer field value
func (o *GmdNiddDownlinkDataDeliveryNotification) GetNiddDownlinkDataTransfer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NiddDownlinkDataTransfer
}

// GetNiddDownlinkDataTransferOk returns a tuple with the NiddDownlinkDataTransfer field value
// and a boolean to check if the value has been set.
func (o *GmdNiddDownlinkDataDeliveryNotification) GetNiddDownlinkDataTransferOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NiddDownlinkDataTransfer, true
}

// SetNiddDownlinkDataTransfer sets field value
func (o *GmdNiddDownlinkDataDeliveryNotification) SetNiddDownlinkDataTransfer(v string) {
	o.NiddDownlinkDataTransfer = v
}

// GetGmdResults returns the GmdResults field value
func (o *GmdNiddDownlinkDataDeliveryNotification) GetGmdResults() []GmdResult {
	if o == nil {
		var ret []GmdResult
		return ret
	}

	return o.GmdResults
}

// GetGmdResultsOk returns a tuple with the GmdResults field value
// and a boolean to check if the value has been set.
func (o *GmdNiddDownlinkDataDeliveryNotification) GetGmdResultsOk() ([]GmdResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.GmdResults, true
}

// SetGmdResults sets field value
func (o *GmdNiddDownlinkDataDeliveryNotification) SetGmdResults(v []GmdResult) {
	o.GmdResults = v
}

func (o GmdNiddDownlinkDataDeliveryNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GmdNiddDownlinkDataDeliveryNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["niddDownlinkDataTransfer"] = o.NiddDownlinkDataTransfer
	toSerialize["gmdResults"] = o.GmdResults
	return toSerialize, nil
}

type NullableGmdNiddDownlinkDataDeliveryNotification struct {
	value *GmdNiddDownlinkDataDeliveryNotification
	isSet bool
}

func (v NullableGmdNiddDownlinkDataDeliveryNotification) Get() *GmdNiddDownlinkDataDeliveryNotification {
	return v.value
}

func (v *NullableGmdNiddDownlinkDataDeliveryNotification) Set(val *GmdNiddDownlinkDataDeliveryNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableGmdNiddDownlinkDataDeliveryNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableGmdNiddDownlinkDataDeliveryNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGmdNiddDownlinkDataDeliveryNotification(val *GmdNiddDownlinkDataDeliveryNotification) *NullableGmdNiddDownlinkDataDeliveryNotification {
	return &NullableGmdNiddDownlinkDataDeliveryNotification{value: val, isSet: true}
}

func (v NullableGmdNiddDownlinkDataDeliveryNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGmdNiddDownlinkDataDeliveryNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

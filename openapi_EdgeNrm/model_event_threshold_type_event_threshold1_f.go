/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
)

// checks if the EventThresholdTypeEventThreshold1F type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventThresholdTypeEventThreshold1F{}

// EventThresholdTypeEventThreshold1F struct for EventThresholdTypeEventThreshold1F
type EventThresholdTypeEventThreshold1F struct {
	CPICH_RSCP *int32 `json:"CPICH_RSCP,omitempty"`
	CPICHEcNo  *int32 `json:"CPICH_EcNo,omitempty"`
	PathLoss   *int32 `json:"PathLoss,omitempty"`
}

// NewEventThresholdTypeEventThreshold1F instantiates a new EventThresholdTypeEventThreshold1F object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventThresholdTypeEventThreshold1F() *EventThresholdTypeEventThreshold1F {
	this := EventThresholdTypeEventThreshold1F{}
	return &this
}

// NewEventThresholdTypeEventThreshold1FWithDefaults instantiates a new EventThresholdTypeEventThreshold1F object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventThresholdTypeEventThreshold1FWithDefaults() *EventThresholdTypeEventThreshold1F {
	this := EventThresholdTypeEventThreshold1F{}
	return &this
}

// GetCPICH_RSCP returns the CPICH_RSCP field value if set, zero value otherwise.
func (o *EventThresholdTypeEventThreshold1F) GetCPICH_RSCP() int32 {
	if o == nil || IsNil(o.CPICH_RSCP) {
		var ret int32
		return ret
	}
	return *o.CPICH_RSCP
}

// GetCPICH_RSCPOk returns a tuple with the CPICH_RSCP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventThresholdTypeEventThreshold1F) GetCPICH_RSCPOk() (*int32, bool) {
	if o == nil || IsNil(o.CPICH_RSCP) {
		return nil, false
	}
	return o.CPICH_RSCP, true
}

// HasCPICH_RSCP returns a boolean if a field has been set.
func (o *EventThresholdTypeEventThreshold1F) HasCPICH_RSCP() bool {
	if o != nil && !IsNil(o.CPICH_RSCP) {
		return true
	}

	return false
}

// SetCPICH_RSCP gets a reference to the given int32 and assigns it to the CPICH_RSCP field.
func (o *EventThresholdTypeEventThreshold1F) SetCPICH_RSCP(v int32) {
	o.CPICH_RSCP = &v
}

// GetCPICHEcNo returns the CPICHEcNo field value if set, zero value otherwise.
func (o *EventThresholdTypeEventThreshold1F) GetCPICHEcNo() int32 {
	if o == nil || IsNil(o.CPICHEcNo) {
		var ret int32
		return ret
	}
	return *o.CPICHEcNo
}

// GetCPICHEcNoOk returns a tuple with the CPICHEcNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventThresholdTypeEventThreshold1F) GetCPICHEcNoOk() (*int32, bool) {
	if o == nil || IsNil(o.CPICHEcNo) {
		return nil, false
	}
	return o.CPICHEcNo, true
}

// HasCPICHEcNo returns a boolean if a field has been set.
func (o *EventThresholdTypeEventThreshold1F) HasCPICHEcNo() bool {
	if o != nil && !IsNil(o.CPICHEcNo) {
		return true
	}

	return false
}

// SetCPICHEcNo gets a reference to the given int32 and assigns it to the CPICHEcNo field.
func (o *EventThresholdTypeEventThreshold1F) SetCPICHEcNo(v int32) {
	o.CPICHEcNo = &v
}

// GetPathLoss returns the PathLoss field value if set, zero value otherwise.
func (o *EventThresholdTypeEventThreshold1F) GetPathLoss() int32 {
	if o == nil || IsNil(o.PathLoss) {
		var ret int32
		return ret
	}
	return *o.PathLoss
}

// GetPathLossOk returns a tuple with the PathLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventThresholdTypeEventThreshold1F) GetPathLossOk() (*int32, bool) {
	if o == nil || IsNil(o.PathLoss) {
		return nil, false
	}
	return o.PathLoss, true
}

// HasPathLoss returns a boolean if a field has been set.
func (o *EventThresholdTypeEventThreshold1F) HasPathLoss() bool {
	if o != nil && !IsNil(o.PathLoss) {
		return true
	}

	return false
}

// SetPathLoss gets a reference to the given int32 and assigns it to the PathLoss field.
func (o *EventThresholdTypeEventThreshold1F) SetPathLoss(v int32) {
	o.PathLoss = &v
}

func (o EventThresholdTypeEventThreshold1F) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventThresholdTypeEventThreshold1F) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CPICH_RSCP) {
		toSerialize["CPICH_RSCP"] = o.CPICH_RSCP
	}
	if !IsNil(o.CPICHEcNo) {
		toSerialize["CPICH_EcNo"] = o.CPICHEcNo
	}
	if !IsNil(o.PathLoss) {
		toSerialize["PathLoss"] = o.PathLoss
	}
	return toSerialize, nil
}

type NullableEventThresholdTypeEventThreshold1F struct {
	value *EventThresholdTypeEventThreshold1F
	isSet bool
}

func (v NullableEventThresholdTypeEventThreshold1F) Get() *EventThresholdTypeEventThreshold1F {
	return v.value
}

func (v *NullableEventThresholdTypeEventThreshold1F) Set(val *EventThresholdTypeEventThreshold1F) {
	v.value = val
	v.isSet = true
}

func (v NullableEventThresholdTypeEventThreshold1F) IsSet() bool {
	return v.isSet
}

func (v *NullableEventThresholdTypeEventThreshold1F) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventThresholdTypeEventThreshold1F(val *EventThresholdTypeEventThreshold1F) *NullableEventThresholdTypeEventThreshold1F {
	return &NullableEventThresholdTypeEventThreshold1F{value: val, isSet: true}
}

func (v NullableEventThresholdTypeEventThreshold1F) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventThresholdTypeEventThreshold1F) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

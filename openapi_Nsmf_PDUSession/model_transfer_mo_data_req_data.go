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

// checks if the TransferMoDataReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferMoDataReqData{}

// TransferMoDataReqData Data within Transfer MO Data Request
type TransferMoDataReqData struct {
	MoData RefToBinaryData `json:"moData"`
	MoExpDataCounter *MoExpDataCounter `json:"moExpDataCounter,omitempty"`
	UeLocation *UserLocation `json:"ueLocation,omitempty"`
}

// NewTransferMoDataReqData instantiates a new TransferMoDataReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferMoDataReqData(moData RefToBinaryData) *TransferMoDataReqData {
	this := TransferMoDataReqData{}
	this.MoData = moData
	return &this
}

// NewTransferMoDataReqDataWithDefaults instantiates a new TransferMoDataReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferMoDataReqDataWithDefaults() *TransferMoDataReqData {
	this := TransferMoDataReqData{}
	return &this
}

// GetMoData returns the MoData field value
func (o *TransferMoDataReqData) GetMoData() RefToBinaryData {
	if o == nil {
		var ret RefToBinaryData
		return ret
	}

	return o.MoData
}

// GetMoDataOk returns a tuple with the MoData field value
// and a boolean to check if the value has been set.
func (o *TransferMoDataReqData) GetMoDataOk() (*RefToBinaryData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MoData, true
}

// SetMoData sets field value
func (o *TransferMoDataReqData) SetMoData(v RefToBinaryData) {
	o.MoData = v
}

// GetMoExpDataCounter returns the MoExpDataCounter field value if set, zero value otherwise.
func (o *TransferMoDataReqData) GetMoExpDataCounter() MoExpDataCounter {
	if o == nil || isNil(o.MoExpDataCounter) {
		var ret MoExpDataCounter
		return ret
	}
	return *o.MoExpDataCounter
}

// GetMoExpDataCounterOk returns a tuple with the MoExpDataCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMoDataReqData) GetMoExpDataCounterOk() (*MoExpDataCounter, bool) {
	if o == nil || isNil(o.MoExpDataCounter) {
		return nil, false
	}
	return o.MoExpDataCounter, true
}

// HasMoExpDataCounter returns a boolean if a field has been set.
func (o *TransferMoDataReqData) HasMoExpDataCounter() bool {
	if o != nil && !isNil(o.MoExpDataCounter) {
		return true
	}

	return false
}

// SetMoExpDataCounter gets a reference to the given MoExpDataCounter and assigns it to the MoExpDataCounter field.
func (o *TransferMoDataReqData) SetMoExpDataCounter(v MoExpDataCounter) {
	o.MoExpDataCounter = &v
}

// GetUeLocation returns the UeLocation field value if set, zero value otherwise.
func (o *TransferMoDataReqData) GetUeLocation() UserLocation {
	if o == nil || isNil(o.UeLocation) {
		var ret UserLocation
		return ret
	}
	return *o.UeLocation
}

// GetUeLocationOk returns a tuple with the UeLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMoDataReqData) GetUeLocationOk() (*UserLocation, bool) {
	if o == nil || isNil(o.UeLocation) {
		return nil, false
	}
	return o.UeLocation, true
}

// HasUeLocation returns a boolean if a field has been set.
func (o *TransferMoDataReqData) HasUeLocation() bool {
	if o != nil && !isNil(o.UeLocation) {
		return true
	}

	return false
}

// SetUeLocation gets a reference to the given UserLocation and assigns it to the UeLocation field.
func (o *TransferMoDataReqData) SetUeLocation(v UserLocation) {
	o.UeLocation = &v
}

func (o TransferMoDataReqData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferMoDataReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["moData"] = o.MoData
	if !isNil(o.MoExpDataCounter) {
		toSerialize["moExpDataCounter"] = o.MoExpDataCounter
	}
	if !isNil(o.UeLocation) {
		toSerialize["ueLocation"] = o.UeLocation
	}
	return toSerialize, nil
}

type NullableTransferMoDataReqData struct {
	value *TransferMoDataReqData
	isSet bool
}

func (v NullableTransferMoDataReqData) Get() *TransferMoDataReqData {
	return v.value
}

func (v *NullableTransferMoDataReqData) Set(val *TransferMoDataReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferMoDataReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferMoDataReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferMoDataReqData(val *TransferMoDataReqData) *NullableTransferMoDataReqData {
	return &NullableTransferMoDataReqData{value: val, isSet: true}
}

func (v NullableTransferMoDataReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferMoDataReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



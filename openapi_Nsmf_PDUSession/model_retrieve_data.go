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

// checks if the RetrieveData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetrieveData{}

// RetrieveData Data within Retrieve Request
type RetrieveData struct {
	SmallDataRateStatusReq *bool `json:"smallDataRateStatusReq,omitempty"`
	PduSessionContextType *PduSessionContextType `json:"pduSessionContextType,omitempty"`
}

// NewRetrieveData instantiates a new RetrieveData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetrieveData() *RetrieveData {
	this := RetrieveData{}
	var smallDataRateStatusReq bool = false
	this.SmallDataRateStatusReq = &smallDataRateStatusReq
	return &this
}

// NewRetrieveDataWithDefaults instantiates a new RetrieveData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetrieveDataWithDefaults() *RetrieveData {
	this := RetrieveData{}
	var smallDataRateStatusReq bool = false
	this.SmallDataRateStatusReq = &smallDataRateStatusReq
	return &this
}

// GetSmallDataRateStatusReq returns the SmallDataRateStatusReq field value if set, zero value otherwise.
func (o *RetrieveData) GetSmallDataRateStatusReq() bool {
	if o == nil || IsNil(o.SmallDataRateStatusReq) {
		var ret bool
		return ret
	}
	return *o.SmallDataRateStatusReq
}

// GetSmallDataRateStatusReqOk returns a tuple with the SmallDataRateStatusReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveData) GetSmallDataRateStatusReqOk() (*bool, bool) {
	if o == nil || IsNil(o.SmallDataRateStatusReq) {
		return nil, false
	}
	return o.SmallDataRateStatusReq, true
}

// HasSmallDataRateStatusReq returns a boolean if a field has been set.
func (o *RetrieveData) HasSmallDataRateStatusReq() bool {
	if o != nil && !IsNil(o.SmallDataRateStatusReq) {
		return true
	}

	return false
}

// SetSmallDataRateStatusReq gets a reference to the given bool and assigns it to the SmallDataRateStatusReq field.
func (o *RetrieveData) SetSmallDataRateStatusReq(v bool) {
	o.SmallDataRateStatusReq = &v
}

// GetPduSessionContextType returns the PduSessionContextType field value if set, zero value otherwise.
func (o *RetrieveData) GetPduSessionContextType() PduSessionContextType {
	if o == nil || IsNil(o.PduSessionContextType) {
		var ret PduSessionContextType
		return ret
	}
	return *o.PduSessionContextType
}

// GetPduSessionContextTypeOk returns a tuple with the PduSessionContextType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveData) GetPduSessionContextTypeOk() (*PduSessionContextType, bool) {
	if o == nil || IsNil(o.PduSessionContextType) {
		return nil, false
	}
	return o.PduSessionContextType, true
}

// HasPduSessionContextType returns a boolean if a field has been set.
func (o *RetrieveData) HasPduSessionContextType() bool {
	if o != nil && !IsNil(o.PduSessionContextType) {
		return true
	}

	return false
}

// SetPduSessionContextType gets a reference to the given PduSessionContextType and assigns it to the PduSessionContextType field.
func (o *RetrieveData) SetPduSessionContextType(v PduSessionContextType) {
	o.PduSessionContextType = &v
}

func (o RetrieveData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetrieveData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SmallDataRateStatusReq) {
		toSerialize["smallDataRateStatusReq"] = o.SmallDataRateStatusReq
	}
	if !IsNil(o.PduSessionContextType) {
		toSerialize["pduSessionContextType"] = o.PduSessionContextType
	}
	return toSerialize, nil
}

type NullableRetrieveData struct {
	value *RetrieveData
	isSet bool
}

func (v NullableRetrieveData) Get() *RetrieveData {
	return v.value
}

func (v *NullableRetrieveData) Set(val *RetrieveData) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrieveData) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrieveData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrieveData(val *RetrieveData) *NullableRetrieveData {
	return &NullableRetrieveData{value: val, isSet: true}
}

func (v NullableRetrieveData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetrieveData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



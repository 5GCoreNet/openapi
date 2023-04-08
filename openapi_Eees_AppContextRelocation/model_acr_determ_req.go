/*
Eees Application Context Relocation Service

Eees Application Context Relocation Service.   © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_AppContextRelocation

import (
	"encoding/json"
)

// checks if the AcrDetermReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcrDetermReq{}

// AcrDetermReq Represents the parameters to request ACR with action determination.
type AcrDetermReq struct {
	RequestorId string `json:"requestorId"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	UeId *string `json:"ueId,omitempty"`
	AcId *string `json:"acId,omitempty"`
	EasId *string `json:"easId,omitempty"`
	SEasEndpoint EndPoint `json:"sEasEndpoint"`
}

// NewAcrDetermReq instantiates a new AcrDetermReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcrDetermReq(requestorId string, sEasEndpoint EndPoint) *AcrDetermReq {
	this := AcrDetermReq{}
	this.RequestorId = requestorId
	this.SEasEndpoint = sEasEndpoint
	return &this
}

// NewAcrDetermReqWithDefaults instantiates a new AcrDetermReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcrDetermReqWithDefaults() *AcrDetermReq {
	this := AcrDetermReq{}
	return &this
}

// GetRequestorId returns the RequestorId field value
func (o *AcrDetermReq) GetRequestorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestorId
}

// GetRequestorIdOk returns a tuple with the RequestorId field value
// and a boolean to check if the value has been set.
func (o *AcrDetermReq) GetRequestorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestorId, true
}

// SetRequestorId sets field value
func (o *AcrDetermReq) SetRequestorId(v string) {
	o.RequestorId = v
}

// GetUeId returns the UeId field value if set, zero value otherwise.
func (o *AcrDetermReq) GetUeId() string {
	if o == nil || isNil(o.UeId) {
		var ret string
		return ret
	}
	return *o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrDetermReq) GetUeIdOk() (*string, bool) {
	if o == nil || isNil(o.UeId) {
		return nil, false
	}
	return o.UeId, true
}

// HasUeId returns a boolean if a field has been set.
func (o *AcrDetermReq) HasUeId() bool {
	if o != nil && !isNil(o.UeId) {
		return true
	}

	return false
}

// SetUeId gets a reference to the given string and assigns it to the UeId field.
func (o *AcrDetermReq) SetUeId(v string) {
	o.UeId = &v
}

// GetAcId returns the AcId field value if set, zero value otherwise.
func (o *AcrDetermReq) GetAcId() string {
	if o == nil || isNil(o.AcId) {
		var ret string
		return ret
	}
	return *o.AcId
}

// GetAcIdOk returns a tuple with the AcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrDetermReq) GetAcIdOk() (*string, bool) {
	if o == nil || isNil(o.AcId) {
		return nil, false
	}
	return o.AcId, true
}

// HasAcId returns a boolean if a field has been set.
func (o *AcrDetermReq) HasAcId() bool {
	if o != nil && !isNil(o.AcId) {
		return true
	}

	return false
}

// SetAcId gets a reference to the given string and assigns it to the AcId field.
func (o *AcrDetermReq) SetAcId(v string) {
	o.AcId = &v
}

// GetEasId returns the EasId field value if set, zero value otherwise.
func (o *AcrDetermReq) GetEasId() string {
	if o == nil || isNil(o.EasId) {
		var ret string
		return ret
	}
	return *o.EasId
}

// GetEasIdOk returns a tuple with the EasId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrDetermReq) GetEasIdOk() (*string, bool) {
	if o == nil || isNil(o.EasId) {
		return nil, false
	}
	return o.EasId, true
}

// HasEasId returns a boolean if a field has been set.
func (o *AcrDetermReq) HasEasId() bool {
	if o != nil && !isNil(o.EasId) {
		return true
	}

	return false
}

// SetEasId gets a reference to the given string and assigns it to the EasId field.
func (o *AcrDetermReq) SetEasId(v string) {
	o.EasId = &v
}

// GetSEasEndpoint returns the SEasEndpoint field value
func (o *AcrDetermReq) GetSEasEndpoint() EndPoint {
	if o == nil {
		var ret EndPoint
		return ret
	}

	return o.SEasEndpoint
}

// GetSEasEndpointOk returns a tuple with the SEasEndpoint field value
// and a boolean to check if the value has been set.
func (o *AcrDetermReq) GetSEasEndpointOk() (*EndPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEasEndpoint, true
}

// SetSEasEndpoint sets field value
func (o *AcrDetermReq) SetSEasEndpoint(v EndPoint) {
	o.SEasEndpoint = v
}

func (o AcrDetermReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcrDetermReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestorId"] = o.RequestorId
	if !isNil(o.UeId) {
		toSerialize["ueId"] = o.UeId
	}
	if !isNil(o.AcId) {
		toSerialize["acId"] = o.AcId
	}
	if !isNil(o.EasId) {
		toSerialize["easId"] = o.EasId
	}
	toSerialize["sEasEndpoint"] = o.SEasEndpoint
	return toSerialize, nil
}

type NullableAcrDetermReq struct {
	value *AcrDetermReq
	isSet bool
}

func (v NullableAcrDetermReq) Get() *AcrDetermReq {
	return v.value
}

func (v *NullableAcrDetermReq) Set(val *AcrDetermReq) {
	v.value = val
	v.isSet = true
}

func (v NullableAcrDetermReq) IsSet() bool {
	return v.isSet
}

func (v *NullableAcrDetermReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcrDetermReq(val *AcrDetermReq) *NullableAcrDetermReq {
	return &NullableAcrDetermReq{value: val, isSet: true}
}

func (v NullableAcrDetermReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcrDetermReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



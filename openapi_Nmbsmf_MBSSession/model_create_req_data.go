/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
)

// checks if the CreateReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReqData{}

// CreateReqData Data within Create Request
type CreateReqData struct {
	MbsSession ExtMbsSession `json:"mbsSession"`
}

// NewCreateReqData instantiates a new CreateReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReqData(mbsSession ExtMbsSession) *CreateReqData {
	this := CreateReqData{}
	this.MbsSession = mbsSession
	return &this
}

// NewCreateReqDataWithDefaults instantiates a new CreateReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReqDataWithDefaults() *CreateReqData {
	this := CreateReqData{}
	return &this
}

// GetMbsSession returns the MbsSession field value
func (o *CreateReqData) GetMbsSession() ExtMbsSession {
	if o == nil {
		var ret ExtMbsSession
		return ret
	}

	return o.MbsSession
}

// GetMbsSessionOk returns a tuple with the MbsSession field value
// and a boolean to check if the value has been set.
func (o *CreateReqData) GetMbsSessionOk() (*ExtMbsSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MbsSession, true
}

// SetMbsSession sets field value
func (o *CreateReqData) SetMbsSession(v ExtMbsSession) {
	o.MbsSession = v
}

func (o CreateReqData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mbsSession"] = o.MbsSession
	return toSerialize, nil
}

type NullableCreateReqData struct {
	value *CreateReqData
	isSet bool
}

func (v NullableCreateReqData) Get() *CreateReqData {
	return v.value
}

func (v *NullableCreateReqData) Set(val *CreateReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReqData(val *CreateReqData) *NullableCreateReqData {
	return &NullableCreateReqData{value: val, isSet: true}
}

func (v NullableCreateReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



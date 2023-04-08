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

// checks if the CreateRspData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRspData{}

// CreateRspData Data within Create Response
type CreateRspData struct {
	MbsSession ExtMbsSession `json:"mbsSession"`
	EventList *MbsSessionEventReportList `json:"eventList,omitempty"`
}

// NewCreateRspData instantiates a new CreateRspData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRspData(mbsSession ExtMbsSession) *CreateRspData {
	this := CreateRspData{}
	this.MbsSession = mbsSession
	return &this
}

// NewCreateRspDataWithDefaults instantiates a new CreateRspData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRspDataWithDefaults() *CreateRspData {
	this := CreateRspData{}
	return &this
}

// GetMbsSession returns the MbsSession field value
func (o *CreateRspData) GetMbsSession() ExtMbsSession {
	if o == nil {
		var ret ExtMbsSession
		return ret
	}

	return o.MbsSession
}

// GetMbsSessionOk returns a tuple with the MbsSession field value
// and a boolean to check if the value has been set.
func (o *CreateRspData) GetMbsSessionOk() (*ExtMbsSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MbsSession, true
}

// SetMbsSession sets field value
func (o *CreateRspData) SetMbsSession(v ExtMbsSession) {
	o.MbsSession = v
}

// GetEventList returns the EventList field value if set, zero value otherwise.
func (o *CreateRspData) GetEventList() MbsSessionEventReportList {
	if o == nil || isNil(o.EventList) {
		var ret MbsSessionEventReportList
		return ret
	}
	return *o.EventList
}

// GetEventListOk returns a tuple with the EventList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRspData) GetEventListOk() (*MbsSessionEventReportList, bool) {
	if o == nil || isNil(o.EventList) {
		return nil, false
	}
	return o.EventList, true
}

// HasEventList returns a boolean if a field has been set.
func (o *CreateRspData) HasEventList() bool {
	if o != nil && !isNil(o.EventList) {
		return true
	}

	return false
}

// SetEventList gets a reference to the given MbsSessionEventReportList and assigns it to the EventList field.
func (o *CreateRspData) SetEventList(v MbsSessionEventReportList) {
	o.EventList = &v
}

func (o CreateRspData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRspData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mbsSession"] = o.MbsSession
	if !isNil(o.EventList) {
		toSerialize["eventList"] = o.EventList
	}
	return toSerialize, nil
}

type NullableCreateRspData struct {
	value *CreateRspData
	isSet bool
}

func (v NullableCreateRspData) Get() *CreateRspData {
	return v.value
}

func (v *NullableCreateRspData) Set(val *CreateRspData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRspData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRspData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRspData(val *CreateRspData) *NullableCreateRspData {
	return &NullableCreateRspData{value: val, isSet: true}
}

func (v NullableCreateRspData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRspData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



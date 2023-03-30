/*
3gpp-mbs-session

API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSSession

import (
	"encoding/json"
)

// checks if the MbsSessionCreateRsp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsSessionCreateRsp{}

// MbsSessionCreateRsp Represents the parameters to be returned in an MBS Session creation response..
type MbsSessionCreateRsp struct {
	MbsSession MbsSession `json:"mbsSession"`
	EventList *MbsSessionEventReportList `json:"eventList,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewMbsSessionCreateRsp instantiates a new MbsSessionCreateRsp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSessionCreateRsp(mbsSession MbsSession) *MbsSessionCreateRsp {
	this := MbsSessionCreateRsp{}
	this.MbsSession = mbsSession
	return &this
}

// NewMbsSessionCreateRspWithDefaults instantiates a new MbsSessionCreateRsp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSessionCreateRspWithDefaults() *MbsSessionCreateRsp {
	this := MbsSessionCreateRsp{}
	return &this
}

// GetMbsSession returns the MbsSession field value
func (o *MbsSessionCreateRsp) GetMbsSession() MbsSession {
	if o == nil {
		var ret MbsSession
		return ret
	}

	return o.MbsSession
}

// GetMbsSessionOk returns a tuple with the MbsSession field value
// and a boolean to check if the value has been set.
func (o *MbsSessionCreateRsp) GetMbsSessionOk() (*MbsSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MbsSession, true
}

// SetMbsSession sets field value
func (o *MbsSessionCreateRsp) SetMbsSession(v MbsSession) {
	o.MbsSession = v
}

// GetEventList returns the EventList field value if set, zero value otherwise.
func (o *MbsSessionCreateRsp) GetEventList() MbsSessionEventReportList {
	if o == nil || IsNil(o.EventList) {
		var ret MbsSessionEventReportList
		return ret
	}
	return *o.EventList
}

// GetEventListOk returns a tuple with the EventList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionCreateRsp) GetEventListOk() (*MbsSessionEventReportList, bool) {
	if o == nil || IsNil(o.EventList) {
		return nil, false
	}
	return o.EventList, true
}

// HasEventList returns a boolean if a field has been set.
func (o *MbsSessionCreateRsp) HasEventList() bool {
	if o != nil && !IsNil(o.EventList) {
		return true
	}

	return false
}

// SetEventList gets a reference to the given MbsSessionEventReportList and assigns it to the EventList field.
func (o *MbsSessionCreateRsp) SetEventList(v MbsSessionEventReportList) {
	o.EventList = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *MbsSessionCreateRsp) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionCreateRsp) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *MbsSessionCreateRsp) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *MbsSessionCreateRsp) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o MbsSessionCreateRsp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsSessionCreateRsp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mbsSession"] = o.MbsSession
	if !IsNil(o.EventList) {
		toSerialize["eventList"] = o.EventList
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableMbsSessionCreateRsp struct {
	value *MbsSessionCreateRsp
	isSet bool
}

func (v NullableMbsSessionCreateRsp) Get() *MbsSessionCreateRsp {
	return v.value
}

func (v *NullableMbsSessionCreateRsp) Set(val *MbsSessionCreateRsp) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionCreateRsp) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionCreateRsp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionCreateRsp(val *MbsSessionCreateRsp) *NullableMbsSessionCreateRsp {
	return &NullableMbsSessionCreateRsp{value: val, isSet: true}
}

func (v NullableMbsSessionCreateRsp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionCreateRsp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSUserDataIngestSession

import (
	"encoding/json"
)

// checks if the MBSUserDataIngSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MBSUserDataIngSession{}

// MBSUserDataIngSession Represents MBS User Data Ingest Session information.
type MBSUserDataIngSession struct {
	MbsUserServId string `json:"mbsUserServId"`
	// Represents one or more MBS Distribution Session(s) composing the MBS User Data Ingest  Session. The key of the map shall be set to the value ofthe \"mbsDistSessionId\" attribute  of the MBSDistributionSessionInfo data structure encoding the corresponding map entry.
	MbsDisSessInfos map[string]MBSDistributionSessionInfo `json:"mbsDisSessInfos"`
	ActPeriods      []TimeWindow                          `json:"actPeriods,omitempty"`
	MbsUserServAnmt *MBSUserServAnmt                      `json:"mbsUserServAnmt,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewMBSUserDataIngSession instantiates a new MBSUserDataIngSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMBSUserDataIngSession(mbsUserServId string, mbsDisSessInfos map[string]MBSDistributionSessionInfo) *MBSUserDataIngSession {
	this := MBSUserDataIngSession{}
	this.MbsUserServId = mbsUserServId
	this.MbsDisSessInfos = mbsDisSessInfos
	return &this
}

// NewMBSUserDataIngSessionWithDefaults instantiates a new MBSUserDataIngSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMBSUserDataIngSessionWithDefaults() *MBSUserDataIngSession {
	this := MBSUserDataIngSession{}
	return &this
}

// GetMbsUserServId returns the MbsUserServId field value
func (o *MBSUserDataIngSession) GetMbsUserServId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MbsUserServId
}

// GetMbsUserServIdOk returns a tuple with the MbsUserServId field value
// and a boolean to check if the value has been set.
func (o *MBSUserDataIngSession) GetMbsUserServIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MbsUserServId, true
}

// SetMbsUserServId sets field value
func (o *MBSUserDataIngSession) SetMbsUserServId(v string) {
	o.MbsUserServId = v
}

// GetMbsDisSessInfos returns the MbsDisSessInfos field value
func (o *MBSUserDataIngSession) GetMbsDisSessInfos() map[string]MBSDistributionSessionInfo {
	if o == nil {
		var ret map[string]MBSDistributionSessionInfo
		return ret
	}

	return o.MbsDisSessInfos
}

// GetMbsDisSessInfosOk returns a tuple with the MbsDisSessInfos field value
// and a boolean to check if the value has been set.
func (o *MBSUserDataIngSession) GetMbsDisSessInfosOk() (*map[string]MBSDistributionSessionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MbsDisSessInfos, true
}

// SetMbsDisSessInfos sets field value
func (o *MBSUserDataIngSession) SetMbsDisSessInfos(v map[string]MBSDistributionSessionInfo) {
	o.MbsDisSessInfos = v
}

// GetActPeriods returns the ActPeriods field value if set, zero value otherwise.
func (o *MBSUserDataIngSession) GetActPeriods() []TimeWindow {
	if o == nil || IsNil(o.ActPeriods) {
		var ret []TimeWindow
		return ret
	}
	return o.ActPeriods
}

// GetActPeriodsOk returns a tuple with the ActPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSUserDataIngSession) GetActPeriodsOk() ([]TimeWindow, bool) {
	if o == nil || IsNil(o.ActPeriods) {
		return nil, false
	}
	return o.ActPeriods, true
}

// HasActPeriods returns a boolean if a field has been set.
func (o *MBSUserDataIngSession) HasActPeriods() bool {
	if o != nil && !IsNil(o.ActPeriods) {
		return true
	}

	return false
}

// SetActPeriods gets a reference to the given []TimeWindow and assigns it to the ActPeriods field.
func (o *MBSUserDataIngSession) SetActPeriods(v []TimeWindow) {
	o.ActPeriods = v
}

// GetMbsUserServAnmt returns the MbsUserServAnmt field value if set, zero value otherwise.
func (o *MBSUserDataIngSession) GetMbsUserServAnmt() MBSUserServAnmt {
	if o == nil || IsNil(o.MbsUserServAnmt) {
		var ret MBSUserServAnmt
		return ret
	}
	return *o.MbsUserServAnmt
}

// GetMbsUserServAnmtOk returns a tuple with the MbsUserServAnmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSUserDataIngSession) GetMbsUserServAnmtOk() (*MBSUserServAnmt, bool) {
	if o == nil || IsNil(o.MbsUserServAnmt) {
		return nil, false
	}
	return o.MbsUserServAnmt, true
}

// HasMbsUserServAnmt returns a boolean if a field has been set.
func (o *MBSUserDataIngSession) HasMbsUserServAnmt() bool {
	if o != nil && !IsNil(o.MbsUserServAnmt) {
		return true
	}

	return false
}

// SetMbsUserServAnmt gets a reference to the given MBSUserServAnmt and assigns it to the MbsUserServAnmt field.
func (o *MBSUserDataIngSession) SetMbsUserServAnmt(v MBSUserServAnmt) {
	o.MbsUserServAnmt = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *MBSUserDataIngSession) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSUserDataIngSession) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *MBSUserDataIngSession) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *MBSUserDataIngSession) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o MBSUserDataIngSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MBSUserDataIngSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mbsUserServId"] = o.MbsUserServId
	toSerialize["mbsDisSessInfos"] = o.MbsDisSessInfos
	if !IsNil(o.ActPeriods) {
		toSerialize["actPeriods"] = o.ActPeriods
	}
	if !IsNil(o.MbsUserServAnmt) {
		toSerialize["mbsUserServAnmt"] = o.MbsUserServAnmt
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableMBSUserDataIngSession struct {
	value *MBSUserDataIngSession
	isSet bool
}

func (v NullableMBSUserDataIngSession) Get() *MBSUserDataIngSession {
	return v.value
}

func (v *NullableMBSUserDataIngSession) Set(val *MBSUserDataIngSession) {
	v.value = val
	v.isSet = true
}

func (v NullableMBSUserDataIngSession) IsSet() bool {
	return v.isSet
}

func (v *NullableMBSUserDataIngSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMBSUserDataIngSession(val *MBSUserDataIngSession) *NullableMBSUserDataIngSession {
	return &NullableMBSUserDataIngSession{value: val, isSet: true}
}

func (v NullableMBSUserDataIngSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMBSUserDataIngSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSUserDataIngestSession

import (
	"encoding/json"
	"time"
)

// checks if the MBSUserServAnmt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MBSUserServAnmt{}

// MBSUserServAnmt Represents the MBS User Service Announcement currently associated with the MBS User Data  Ingest Session.
type MBSUserServAnmt struct {
	ExtServiceId []string `json:"extServiceId"`
	ServClass    string   `json:"servClass"`
	// string with format \"date-time\" as defined in OpenAPI.
	StartTime *time.Time `json:"startTime,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	EndTime       *time.Time               `json:"endTime,omitempty"`
	ServNameDescs []ServiceNameDescription `json:"servNameDescs"`
	MainServLang  *string                  `json:"mainServLang,omitempty"`
	// Represents the set of MBS Distribution Session Announcements currently associated with  this MBS User Service Announcement.
	MbsDistSessAnmt *map[string]MBSDistSessionAnmt `json:"mbsDistSessAnmt,omitempty"`
}

// NewMBSUserServAnmt instantiates a new MBSUserServAnmt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMBSUserServAnmt(extServiceId []string, servClass string, servNameDescs []ServiceNameDescription) *MBSUserServAnmt {
	this := MBSUserServAnmt{}
	this.ExtServiceId = extServiceId
	this.ServClass = servClass
	this.ServNameDescs = servNameDescs
	return &this
}

// NewMBSUserServAnmtWithDefaults instantiates a new MBSUserServAnmt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMBSUserServAnmtWithDefaults() *MBSUserServAnmt {
	this := MBSUserServAnmt{}
	return &this
}

// GetExtServiceId returns the ExtServiceId field value
func (o *MBSUserServAnmt) GetExtServiceId() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ExtServiceId
}

// GetExtServiceIdOk returns a tuple with the ExtServiceId field value
// and a boolean to check if the value has been set.
func (o *MBSUserServAnmt) GetExtServiceIdOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtServiceId, true
}

// SetExtServiceId sets field value
func (o *MBSUserServAnmt) SetExtServiceId(v []string) {
	o.ExtServiceId = v
}

// GetServClass returns the ServClass field value
func (o *MBSUserServAnmt) GetServClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServClass
}

// GetServClassOk returns a tuple with the ServClass field value
// and a boolean to check if the value has been set.
func (o *MBSUserServAnmt) GetServClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServClass, true
}

// SetServClass sets field value
func (o *MBSUserServAnmt) SetServClass(v string) {
	o.ServClass = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *MBSUserServAnmt) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSUserServAnmt) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *MBSUserServAnmt) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *MBSUserServAnmt) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *MBSUserServAnmt) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSUserServAnmt) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *MBSUserServAnmt) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *MBSUserServAnmt) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetServNameDescs returns the ServNameDescs field value
func (o *MBSUserServAnmt) GetServNameDescs() []ServiceNameDescription {
	if o == nil {
		var ret []ServiceNameDescription
		return ret
	}

	return o.ServNameDescs
}

// GetServNameDescsOk returns a tuple with the ServNameDescs field value
// and a boolean to check if the value has been set.
func (o *MBSUserServAnmt) GetServNameDescsOk() ([]ServiceNameDescription, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServNameDescs, true
}

// SetServNameDescs sets field value
func (o *MBSUserServAnmt) SetServNameDescs(v []ServiceNameDescription) {
	o.ServNameDescs = v
}

// GetMainServLang returns the MainServLang field value if set, zero value otherwise.
func (o *MBSUserServAnmt) GetMainServLang() string {
	if o == nil || IsNil(o.MainServLang) {
		var ret string
		return ret
	}
	return *o.MainServLang
}

// GetMainServLangOk returns a tuple with the MainServLang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSUserServAnmt) GetMainServLangOk() (*string, bool) {
	if o == nil || IsNil(o.MainServLang) {
		return nil, false
	}
	return o.MainServLang, true
}

// HasMainServLang returns a boolean if a field has been set.
func (o *MBSUserServAnmt) HasMainServLang() bool {
	if o != nil && !IsNil(o.MainServLang) {
		return true
	}

	return false
}

// SetMainServLang gets a reference to the given string and assigns it to the MainServLang field.
func (o *MBSUserServAnmt) SetMainServLang(v string) {
	o.MainServLang = &v
}

// GetMbsDistSessAnmt returns the MbsDistSessAnmt field value if set, zero value otherwise.
func (o *MBSUserServAnmt) GetMbsDistSessAnmt() map[string]MBSDistSessionAnmt {
	if o == nil || IsNil(o.MbsDistSessAnmt) {
		var ret map[string]MBSDistSessionAnmt
		return ret
	}
	return *o.MbsDistSessAnmt
}

// GetMbsDistSessAnmtOk returns a tuple with the MbsDistSessAnmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSUserServAnmt) GetMbsDistSessAnmtOk() (*map[string]MBSDistSessionAnmt, bool) {
	if o == nil || IsNil(o.MbsDistSessAnmt) {
		return nil, false
	}
	return o.MbsDistSessAnmt, true
}

// HasMbsDistSessAnmt returns a boolean if a field has been set.
func (o *MBSUserServAnmt) HasMbsDistSessAnmt() bool {
	if o != nil && !IsNil(o.MbsDistSessAnmt) {
		return true
	}

	return false
}

// SetMbsDistSessAnmt gets a reference to the given map[string]MBSDistSessionAnmt and assigns it to the MbsDistSessAnmt field.
func (o *MBSUserServAnmt) SetMbsDistSessAnmt(v map[string]MBSDistSessionAnmt) {
	o.MbsDistSessAnmt = &v
}

func (o MBSUserServAnmt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MBSUserServAnmt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["extServiceId"] = o.ExtServiceId
	toSerialize["servClass"] = o.ServClass
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	toSerialize["servNameDescs"] = o.ServNameDescs
	if !IsNil(o.MainServLang) {
		toSerialize["mainServLang"] = o.MainServLang
	}
	if !IsNil(o.MbsDistSessAnmt) {
		toSerialize["mbsDistSessAnmt"] = o.MbsDistSessAnmt
	}
	return toSerialize, nil
}

type NullableMBSUserServAnmt struct {
	value *MBSUserServAnmt
	isSet bool
}

func (v NullableMBSUserServAnmt) Get() *MBSUserServAnmt {
	return v.value
}

func (v *NullableMBSUserServAnmt) Set(val *MBSUserServAnmt) {
	v.value = val
	v.isSet = true
}

func (v NullableMBSUserServAnmt) IsSet() bool {
	return v.isSet
}

func (v *NullableMBSUserServAnmt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMBSUserServAnmt(val *MBSUserServAnmt) *NullableMBSUserServAnmt {
	return &NullableMBSUserServAnmt{value: val, isSet: true}
}

func (v NullableMBSUserServAnmt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMBSUserServAnmt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
)

// checks if the NotifyCommentsAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyCommentsAllOf{}

// NotifyCommentsAllOf struct for NotifyCommentsAllOf
type NotifyCommentsAllOf struct {
	AlarmId string `json:"alarmId"`
	AlarmType AlarmType `json:"alarmType"`
	ProbableCause ProbableCause `json:"probableCause"`
	PerceivedSeverity PerceivedSeverity `json:"perceivedSeverity"`
	// Collection of comments. The comment identifiers are allocated by the MnS producer and used as key in the map.
	Comments map[string]Comment `json:"comments"`
}

// NewNotifyCommentsAllOf instantiates a new NotifyCommentsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyCommentsAllOf(alarmId string, alarmType AlarmType, probableCause ProbableCause, perceivedSeverity PerceivedSeverity, comments map[string]Comment) *NotifyCommentsAllOf {
	this := NotifyCommentsAllOf{}
	this.AlarmId = alarmId
	this.AlarmType = alarmType
	this.ProbableCause = probableCause
	this.PerceivedSeverity = perceivedSeverity
	this.Comments = comments
	return &this
}

// NewNotifyCommentsAllOfWithDefaults instantiates a new NotifyCommentsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyCommentsAllOfWithDefaults() *NotifyCommentsAllOf {
	this := NotifyCommentsAllOf{}
	return &this
}

// GetAlarmId returns the AlarmId field value
func (o *NotifyCommentsAllOf) GetAlarmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlarmId
}

// GetAlarmIdOk returns a tuple with the AlarmId field value
// and a boolean to check if the value has been set.
func (o *NotifyCommentsAllOf) GetAlarmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmId, true
}

// SetAlarmId sets field value
func (o *NotifyCommentsAllOf) SetAlarmId(v string) {
	o.AlarmId = v
}

// GetAlarmType returns the AlarmType field value
func (o *NotifyCommentsAllOf) GetAlarmType() AlarmType {
	if o == nil {
		var ret AlarmType
		return ret
	}

	return o.AlarmType
}

// GetAlarmTypeOk returns a tuple with the AlarmType field value
// and a boolean to check if the value has been set.
func (o *NotifyCommentsAllOf) GetAlarmTypeOk() (*AlarmType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmType, true
}

// SetAlarmType sets field value
func (o *NotifyCommentsAllOf) SetAlarmType(v AlarmType) {
	o.AlarmType = v
}

// GetProbableCause returns the ProbableCause field value
func (o *NotifyCommentsAllOf) GetProbableCause() ProbableCause {
	if o == nil {
		var ret ProbableCause
		return ret
	}

	return o.ProbableCause
}

// GetProbableCauseOk returns a tuple with the ProbableCause field value
// and a boolean to check if the value has been set.
func (o *NotifyCommentsAllOf) GetProbableCauseOk() (*ProbableCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProbableCause, true
}

// SetProbableCause sets field value
func (o *NotifyCommentsAllOf) SetProbableCause(v ProbableCause) {
	o.ProbableCause = v
}

// GetPerceivedSeverity returns the PerceivedSeverity field value
func (o *NotifyCommentsAllOf) GetPerceivedSeverity() PerceivedSeverity {
	if o == nil {
		var ret PerceivedSeverity
		return ret
	}

	return o.PerceivedSeverity
}

// GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field value
// and a boolean to check if the value has been set.
func (o *NotifyCommentsAllOf) GetPerceivedSeverityOk() (*PerceivedSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerceivedSeverity, true
}

// SetPerceivedSeverity sets field value
func (o *NotifyCommentsAllOf) SetPerceivedSeverity(v PerceivedSeverity) {
	o.PerceivedSeverity = v
}

// GetComments returns the Comments field value
func (o *NotifyCommentsAllOf) GetComments() map[string]Comment {
	if o == nil {
		var ret map[string]Comment
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *NotifyCommentsAllOf) GetCommentsOk() (*map[string]Comment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *NotifyCommentsAllOf) SetComments(v map[string]Comment) {
	o.Comments = v
}

func (o NotifyCommentsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyCommentsAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alarmId"] = o.AlarmId
	toSerialize["alarmType"] = o.AlarmType
	toSerialize["probableCause"] = o.ProbableCause
	toSerialize["perceivedSeverity"] = o.PerceivedSeverity
	toSerialize["comments"] = o.Comments
	return toSerialize, nil
}

type NullableNotifyCommentsAllOf struct {
	value *NotifyCommentsAllOf
	isSet bool
}

func (v NullableNotifyCommentsAllOf) Get() *NotifyCommentsAllOf {
	return v.value
}

func (v *NullableNotifyCommentsAllOf) Set(val *NotifyCommentsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyCommentsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyCommentsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyCommentsAllOf(val *NotifyCommentsAllOf) *NullableNotifyCommentsAllOf {
	return &NullableNotifyCommentsAllOf{value: val, isSet: true}
}

func (v NullableNotifyCommentsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyCommentsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



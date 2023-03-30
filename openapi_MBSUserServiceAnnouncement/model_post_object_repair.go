/*
MBS User Service Announcement Element units’ definition

MBS User Service Announcement Element units. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSUserServiceAnnouncement

import (
	"encoding/json"
)

// checks if the PostObjectRepair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostObjectRepair{}

// PostObjectRepair struct for PostObjectRepair
type PostObjectRepair struct {
	ServiceURIs []string `json:"serviceURIs,omitempty"`
	// indicating a time in seconds.
	OffsetTime *int32 `json:"offsetTime,omitempty"`
	// indicating a time in seconds.
	RandomTimePeriod *int32 `json:"randomTimePeriod,omitempty"`
}

// NewPostObjectRepair instantiates a new PostObjectRepair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostObjectRepair() *PostObjectRepair {
	this := PostObjectRepair{}
	return &this
}

// NewPostObjectRepairWithDefaults instantiates a new PostObjectRepair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostObjectRepairWithDefaults() *PostObjectRepair {
	this := PostObjectRepair{}
	return &this
}

// GetServiceURIs returns the ServiceURIs field value if set, zero value otherwise.
func (o *PostObjectRepair) GetServiceURIs() []string {
	if o == nil || IsNil(o.ServiceURIs) {
		var ret []string
		return ret
	}
	return o.ServiceURIs
}

// GetServiceURIsOk returns a tuple with the ServiceURIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostObjectRepair) GetServiceURIsOk() ([]string, bool) {
	if o == nil || IsNil(o.ServiceURIs) {
		return nil, false
	}
	return o.ServiceURIs, true
}

// HasServiceURIs returns a boolean if a field has been set.
func (o *PostObjectRepair) HasServiceURIs() bool {
	if o != nil && !IsNil(o.ServiceURIs) {
		return true
	}

	return false
}

// SetServiceURIs gets a reference to the given []string and assigns it to the ServiceURIs field.
func (o *PostObjectRepair) SetServiceURIs(v []string) {
	o.ServiceURIs = v
}

// GetOffsetTime returns the OffsetTime field value if set, zero value otherwise.
func (o *PostObjectRepair) GetOffsetTime() int32 {
	if o == nil || IsNil(o.OffsetTime) {
		var ret int32
		return ret
	}
	return *o.OffsetTime
}

// GetOffsetTimeOk returns a tuple with the OffsetTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostObjectRepair) GetOffsetTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.OffsetTime) {
		return nil, false
	}
	return o.OffsetTime, true
}

// HasOffsetTime returns a boolean if a field has been set.
func (o *PostObjectRepair) HasOffsetTime() bool {
	if o != nil && !IsNil(o.OffsetTime) {
		return true
	}

	return false
}

// SetOffsetTime gets a reference to the given int32 and assigns it to the OffsetTime field.
func (o *PostObjectRepair) SetOffsetTime(v int32) {
	o.OffsetTime = &v
}

// GetRandomTimePeriod returns the RandomTimePeriod field value if set, zero value otherwise.
func (o *PostObjectRepair) GetRandomTimePeriod() int32 {
	if o == nil || IsNil(o.RandomTimePeriod) {
		var ret int32
		return ret
	}
	return *o.RandomTimePeriod
}

// GetRandomTimePeriodOk returns a tuple with the RandomTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostObjectRepair) GetRandomTimePeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RandomTimePeriod) {
		return nil, false
	}
	return o.RandomTimePeriod, true
}

// HasRandomTimePeriod returns a boolean if a field has been set.
func (o *PostObjectRepair) HasRandomTimePeriod() bool {
	if o != nil && !IsNil(o.RandomTimePeriod) {
		return true
	}

	return false
}

// SetRandomTimePeriod gets a reference to the given int32 and assigns it to the RandomTimePeriod field.
func (o *PostObjectRepair) SetRandomTimePeriod(v int32) {
	o.RandomTimePeriod = &v
}

func (o PostObjectRepair) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostObjectRepair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceURIs) {
		toSerialize["serviceURIs"] = o.ServiceURIs
	}
	if !IsNil(o.OffsetTime) {
		toSerialize["offsetTime"] = o.OffsetTime
	}
	if !IsNil(o.RandomTimePeriod) {
		toSerialize["randomTimePeriod"] = o.RandomTimePeriod
	}
	return toSerialize, nil
}

type NullablePostObjectRepair struct {
	value *PostObjectRepair
	isSet bool
}

func (v NullablePostObjectRepair) Get() *PostObjectRepair {
	return v.value
}

func (v *NullablePostObjectRepair) Set(val *PostObjectRepair) {
	v.value = val
	v.isSet = true
}

func (v NullablePostObjectRepair) IsSet() bool {
	return v.isSet
}

func (v *NullablePostObjectRepair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostObjectRepair(val *PostObjectRepair) *NullablePostObjectRepair {
	return &NullablePostObjectRepair{value: val, isSet: true}
}

func (v NullablePostObjectRepair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostObjectRepair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



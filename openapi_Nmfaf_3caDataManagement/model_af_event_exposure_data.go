/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// checks if the AfEventExposureData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AfEventExposureData{}

// AfEventExposureData AF Event Exposure data managed by a given NEF Instance
type AfEventExposureData struct {
	AfEvents []AfEvent `json:"afEvents"`
	AfIds []string `json:"afIds,omitempty"`
	AppIds []string `json:"appIds,omitempty"`
}

// NewAfEventExposureData instantiates a new AfEventExposureData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfEventExposureData(afEvents []AfEvent) *AfEventExposureData {
	this := AfEventExposureData{}
	this.AfEvents = afEvents
	return &this
}

// NewAfEventExposureDataWithDefaults instantiates a new AfEventExposureData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfEventExposureDataWithDefaults() *AfEventExposureData {
	this := AfEventExposureData{}
	return &this
}

// GetAfEvents returns the AfEvents field value
func (o *AfEventExposureData) GetAfEvents() []AfEvent {
	if o == nil {
		var ret []AfEvent
		return ret
	}

	return o.AfEvents
}

// GetAfEventsOk returns a tuple with the AfEvents field value
// and a boolean to check if the value has been set.
func (o *AfEventExposureData) GetAfEventsOk() ([]AfEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.AfEvents, true
}

// SetAfEvents sets field value
func (o *AfEventExposureData) SetAfEvents(v []AfEvent) {
	o.AfEvents = v
}

// GetAfIds returns the AfIds field value if set, zero value otherwise.
func (o *AfEventExposureData) GetAfIds() []string {
	if o == nil || isNil(o.AfIds) {
		var ret []string
		return ret
	}
	return o.AfIds
}

// GetAfIdsOk returns a tuple with the AfIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventExposureData) GetAfIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AfIds) {
		return nil, false
	}
	return o.AfIds, true
}

// HasAfIds returns a boolean if a field has been set.
func (o *AfEventExposureData) HasAfIds() bool {
	if o != nil && !isNil(o.AfIds) {
		return true
	}

	return false
}

// SetAfIds gets a reference to the given []string and assigns it to the AfIds field.
func (o *AfEventExposureData) SetAfIds(v []string) {
	o.AfIds = v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *AfEventExposureData) GetAppIds() []string {
	if o == nil || isNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventExposureData) GetAppIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AppIds) {
		return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *AfEventExposureData) HasAppIds() bool {
	if o != nil && !isNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *AfEventExposureData) SetAppIds(v []string) {
	o.AppIds = v
}

func (o AfEventExposureData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AfEventExposureData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["afEvents"] = o.AfEvents
	if !isNil(o.AfIds) {
		toSerialize["afIds"] = o.AfIds
	}
	if !isNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	return toSerialize, nil
}

type NullableAfEventExposureData struct {
	value *AfEventExposureData
	isSet bool
}

func (v NullableAfEventExposureData) Get() *AfEventExposureData {
	return v.value
}

func (v *NullableAfEventExposureData) Set(val *AfEventExposureData) {
	v.value = val
	v.isSet = true
}

func (v NullableAfEventExposureData) IsSet() bool {
	return v.isSet
}

func (v *NullableAfEventExposureData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfEventExposureData(val *AfEventExposureData) *NullableAfEventExposureData {
	return &NullableAfEventExposureData{value: val, isSet: true}
}

func (v NullableAfEventExposureData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfEventExposureData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



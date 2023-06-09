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

// checks if the AvailabilityInformationBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailabilityInformationBinding{}

// AvailabilityInformationBinding struct for AvailabilityInformationBinding
type AvailabilityInformationBinding struct {
	MbsServiceArea []MbsServiceArea `json:"mbsServiceArea,omitempty"`
	// MBS Frequency Selection Area Identifier
	MbsFSAId       *string `json:"mbsFSAId,omitempty"`
	RadioFrequency []int32 `json:"radioFrequency,omitempty"`
}

// NewAvailabilityInformationBinding instantiates a new AvailabilityInformationBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityInformationBinding() *AvailabilityInformationBinding {
	this := AvailabilityInformationBinding{}
	return &this
}

// NewAvailabilityInformationBindingWithDefaults instantiates a new AvailabilityInformationBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityInformationBindingWithDefaults() *AvailabilityInformationBinding {
	this := AvailabilityInformationBinding{}
	return &this
}

// GetMbsServiceArea returns the MbsServiceArea field value if set, zero value otherwise.
func (o *AvailabilityInformationBinding) GetMbsServiceArea() []MbsServiceArea {
	if o == nil || IsNil(o.MbsServiceArea) {
		var ret []MbsServiceArea
		return ret
	}
	return o.MbsServiceArea
}

// GetMbsServiceAreaOk returns a tuple with the MbsServiceArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityInformationBinding) GetMbsServiceAreaOk() ([]MbsServiceArea, bool) {
	if o == nil || IsNil(o.MbsServiceArea) {
		return nil, false
	}
	return o.MbsServiceArea, true
}

// HasMbsServiceArea returns a boolean if a field has been set.
func (o *AvailabilityInformationBinding) HasMbsServiceArea() bool {
	if o != nil && !IsNil(o.MbsServiceArea) {
		return true
	}

	return false
}

// SetMbsServiceArea gets a reference to the given []MbsServiceArea and assigns it to the MbsServiceArea field.
func (o *AvailabilityInformationBinding) SetMbsServiceArea(v []MbsServiceArea) {
	o.MbsServiceArea = v
}

// GetMbsFSAId returns the MbsFSAId field value if set, zero value otherwise.
func (o *AvailabilityInformationBinding) GetMbsFSAId() string {
	if o == nil || IsNil(o.MbsFSAId) {
		var ret string
		return ret
	}
	return *o.MbsFSAId
}

// GetMbsFSAIdOk returns a tuple with the MbsFSAId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityInformationBinding) GetMbsFSAIdOk() (*string, bool) {
	if o == nil || IsNil(o.MbsFSAId) {
		return nil, false
	}
	return o.MbsFSAId, true
}

// HasMbsFSAId returns a boolean if a field has been set.
func (o *AvailabilityInformationBinding) HasMbsFSAId() bool {
	if o != nil && !IsNil(o.MbsFSAId) {
		return true
	}

	return false
}

// SetMbsFSAId gets a reference to the given string and assigns it to the MbsFSAId field.
func (o *AvailabilityInformationBinding) SetMbsFSAId(v string) {
	o.MbsFSAId = &v
}

// GetRadioFrequency returns the RadioFrequency field value if set, zero value otherwise.
func (o *AvailabilityInformationBinding) GetRadioFrequency() []int32 {
	if o == nil || IsNil(o.RadioFrequency) {
		var ret []int32
		return ret
	}
	return o.RadioFrequency
}

// GetRadioFrequencyOk returns a tuple with the RadioFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityInformationBinding) GetRadioFrequencyOk() ([]int32, bool) {
	if o == nil || IsNil(o.RadioFrequency) {
		return nil, false
	}
	return o.RadioFrequency, true
}

// HasRadioFrequency returns a boolean if a field has been set.
func (o *AvailabilityInformationBinding) HasRadioFrequency() bool {
	if o != nil && !IsNil(o.RadioFrequency) {
		return true
	}

	return false
}

// SetRadioFrequency gets a reference to the given []int32 and assigns it to the RadioFrequency field.
func (o *AvailabilityInformationBinding) SetRadioFrequency(v []int32) {
	o.RadioFrequency = v
}

func (o AvailabilityInformationBinding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailabilityInformationBinding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MbsServiceArea) {
		toSerialize["mbsServiceArea"] = o.MbsServiceArea
	}
	if !IsNil(o.MbsFSAId) {
		toSerialize["mbsFSAId"] = o.MbsFSAId
	}
	if !IsNil(o.RadioFrequency) {
		toSerialize["radioFrequency"] = o.RadioFrequency
	}
	return toSerialize, nil
}

type NullableAvailabilityInformationBinding struct {
	value *AvailabilityInformationBinding
	isSet bool
}

func (v NullableAvailabilityInformationBinding) Get() *AvailabilityInformationBinding {
	return v.value
}

func (v *NullableAvailabilityInformationBinding) Set(val *AvailabilityInformationBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityInformationBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityInformationBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityInformationBinding(val *AvailabilityInformationBinding) *NullableAvailabilityInformationBinding {
	return &NullableAvailabilityInformationBinding{value: val, isSet: true}
}

func (v NullableAvailabilityInformationBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityInformationBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

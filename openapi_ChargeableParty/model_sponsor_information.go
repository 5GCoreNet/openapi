/*
3gpp-chargeable-party

API for Chargeable Party management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ChargeableParty

import (
	"encoding/json"
)

// checks if the SponsorInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsorInformation{}

// SponsorInformation Represents a sponsor information.
type SponsorInformation struct {
	// It indicates Sponsor ID.
	SponsorId string `json:"sponsorId"`
	// It indicates Application Service Provider ID.
	AspId string `json:"aspId"`
}

// NewSponsorInformation instantiates a new SponsorInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsorInformation(sponsorId string, aspId string) *SponsorInformation {
	this := SponsorInformation{}
	this.SponsorId = sponsorId
	this.AspId = aspId
	return &this
}

// NewSponsorInformationWithDefaults instantiates a new SponsorInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsorInformationWithDefaults() *SponsorInformation {
	this := SponsorInformation{}
	return &this
}

// GetSponsorId returns the SponsorId field value
func (o *SponsorInformation) GetSponsorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SponsorId
}

// GetSponsorIdOk returns a tuple with the SponsorId field value
// and a boolean to check if the value has been set.
func (o *SponsorInformation) GetSponsorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SponsorId, true
}

// SetSponsorId sets field value
func (o *SponsorInformation) SetSponsorId(v string) {
	o.SponsorId = v
}

// GetAspId returns the AspId field value
func (o *SponsorInformation) GetAspId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AspId
}

// GetAspIdOk returns a tuple with the AspId field value
// and a boolean to check if the value has been set.
func (o *SponsorInformation) GetAspIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AspId, true
}

// SetAspId sets field value
func (o *SponsorInformation) SetAspId(v string) {
	o.AspId = v
}

func (o SponsorInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SponsorInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sponsorId"] = o.SponsorId
	toSerialize["aspId"] = o.AspId
	return toSerialize, nil
}

type NullableSponsorInformation struct {
	value *SponsorInformation
	isSet bool
}

func (v NullableSponsorInformation) Get() *SponsorInformation {
	return v.value
}

func (v *NullableSponsorInformation) Set(val *SponsorInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsorInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsorInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsorInformation(val *SponsorInformation) *NullableSponsorInformation {
	return &NullableSponsorInformation{value: val, isSet: true}
}

func (v NullableSponsorInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSponsorInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



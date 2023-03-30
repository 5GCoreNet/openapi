/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the SgsnLocationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SgsnLocationData{}

// SgsnLocationData Location information as retrieved from the SGSN serving node
type SgsnLocationData struct {
	SgsnNumber string `json:"sgsnNumber"`
	PlmnId PlmnId `json:"plmnId"`
	SgsnLocation *UtraLocation `json:"sgsnLocation,omitempty"`
	CsgInformation *CsgInformation `json:"csgInformation,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	TimeZone *string `json:"timeZone,omitempty"`
	RatType *RatType `json:"ratType,omitempty"`
}

// NewSgsnLocationData instantiates a new SgsnLocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSgsnLocationData(sgsnNumber string, plmnId PlmnId) *SgsnLocationData {
	this := SgsnLocationData{}
	this.SgsnNumber = sgsnNumber
	this.PlmnId = plmnId
	return &this
}

// NewSgsnLocationDataWithDefaults instantiates a new SgsnLocationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSgsnLocationDataWithDefaults() *SgsnLocationData {
	this := SgsnLocationData{}
	return &this
}

// GetSgsnNumber returns the SgsnNumber field value
func (o *SgsnLocationData) GetSgsnNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SgsnNumber
}

// GetSgsnNumberOk returns a tuple with the SgsnNumber field value
// and a boolean to check if the value has been set.
func (o *SgsnLocationData) GetSgsnNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SgsnNumber, true
}

// SetSgsnNumber sets field value
func (o *SgsnLocationData) SetSgsnNumber(v string) {
	o.SgsnNumber = v
}

// GetPlmnId returns the PlmnId field value
func (o *SgsnLocationData) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *SgsnLocationData) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *SgsnLocationData) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetSgsnLocation returns the SgsnLocation field value if set, zero value otherwise.
func (o *SgsnLocationData) GetSgsnLocation() UtraLocation {
	if o == nil || IsNil(o.SgsnLocation) {
		var ret UtraLocation
		return ret
	}
	return *o.SgsnLocation
}

// GetSgsnLocationOk returns a tuple with the SgsnLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SgsnLocationData) GetSgsnLocationOk() (*UtraLocation, bool) {
	if o == nil || IsNil(o.SgsnLocation) {
		return nil, false
	}
	return o.SgsnLocation, true
}

// HasSgsnLocation returns a boolean if a field has been set.
func (o *SgsnLocationData) HasSgsnLocation() bool {
	if o != nil && !IsNil(o.SgsnLocation) {
		return true
	}

	return false
}

// SetSgsnLocation gets a reference to the given UtraLocation and assigns it to the SgsnLocation field.
func (o *SgsnLocationData) SetSgsnLocation(v UtraLocation) {
	o.SgsnLocation = &v
}

// GetCsgInformation returns the CsgInformation field value if set, zero value otherwise.
func (o *SgsnLocationData) GetCsgInformation() CsgInformation {
	if o == nil || IsNil(o.CsgInformation) {
		var ret CsgInformation
		return ret
	}
	return *o.CsgInformation
}

// GetCsgInformationOk returns a tuple with the CsgInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SgsnLocationData) GetCsgInformationOk() (*CsgInformation, bool) {
	if o == nil || IsNil(o.CsgInformation) {
		return nil, false
	}
	return o.CsgInformation, true
}

// HasCsgInformation returns a boolean if a field has been set.
func (o *SgsnLocationData) HasCsgInformation() bool {
	if o != nil && !IsNil(o.CsgInformation) {
		return true
	}

	return false
}

// SetCsgInformation gets a reference to the given CsgInformation and assigns it to the CsgInformation field.
func (o *SgsnLocationData) SetCsgInformation(v CsgInformation) {
	o.CsgInformation = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *SgsnLocationData) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SgsnLocationData) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *SgsnLocationData) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *SgsnLocationData) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *SgsnLocationData) GetRatType() RatType {
	if o == nil || IsNil(o.RatType) {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SgsnLocationData) GetRatTypeOk() (*RatType, bool) {
	if o == nil || IsNil(o.RatType) {
		return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *SgsnLocationData) HasRatType() bool {
	if o != nil && !IsNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *SgsnLocationData) SetRatType(v RatType) {
	o.RatType = &v
}

func (o SgsnLocationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SgsnLocationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sgsnNumber"] = o.SgsnNumber
	toSerialize["plmnId"] = o.PlmnId
	if !IsNil(o.SgsnLocation) {
		toSerialize["sgsnLocation"] = o.SgsnLocation
	}
	if !IsNil(o.CsgInformation) {
		toSerialize["csgInformation"] = o.CsgInformation
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	return toSerialize, nil
}

type NullableSgsnLocationData struct {
	value *SgsnLocationData
	isSet bool
}

func (v NullableSgsnLocationData) Get() *SgsnLocationData {
	return v.value
}

func (v *NullableSgsnLocationData) Set(val *SgsnLocationData) {
	v.value = val
	v.isSet = true
}

func (v NullableSgsnLocationData) IsSet() bool {
	return v.isSet
}

func (v *NullableSgsnLocationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSgsnLocationData(val *SgsnLocationData) *NullableSgsnLocationData {
	return &NullableSgsnLocationData{value: val, isSet: true}
}

func (v NullableSgsnLocationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSgsnLocationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



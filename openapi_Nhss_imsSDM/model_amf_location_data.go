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

// checks if the AmfLocationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfLocationData{}

// AmfLocationData Location information as retrieved from the AMF serving node
type AmfLocationData struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	AmfAddress  string      `json:"amfAddress"`
	PlmnId      PlmnId      `json:"plmnId"`
	AmfLocation *NrLocation `json:"amfLocation,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SmsfAddress *string `json:"smsfAddress,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	TimeZone *string  `json:"timeZone,omitempty"`
	RatType  *RatType `json:"ratType,omitempty"`
}

// NewAmfLocationData instantiates a new AmfLocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfLocationData(amfAddress string, plmnId PlmnId) *AmfLocationData {
	this := AmfLocationData{}
	this.AmfAddress = amfAddress
	this.PlmnId = plmnId
	return &this
}

// NewAmfLocationDataWithDefaults instantiates a new AmfLocationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfLocationDataWithDefaults() *AmfLocationData {
	this := AmfLocationData{}
	return &this
}

// GetAmfAddress returns the AmfAddress field value
func (o *AmfLocationData) GetAmfAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmfAddress
}

// GetAmfAddressOk returns a tuple with the AmfAddress field value
// and a boolean to check if the value has been set.
func (o *AmfLocationData) GetAmfAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmfAddress, true
}

// SetAmfAddress sets field value
func (o *AmfLocationData) SetAmfAddress(v string) {
	o.AmfAddress = v
}

// GetPlmnId returns the PlmnId field value
func (o *AmfLocationData) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *AmfLocationData) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *AmfLocationData) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetAmfLocation returns the AmfLocation field value if set, zero value otherwise.
func (o *AmfLocationData) GetAmfLocation() NrLocation {
	if o == nil || IsNil(o.AmfLocation) {
		var ret NrLocation
		return ret
	}
	return *o.AmfLocation
}

// GetAmfLocationOk returns a tuple with the AmfLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfLocationData) GetAmfLocationOk() (*NrLocation, bool) {
	if o == nil || IsNil(o.AmfLocation) {
		return nil, false
	}
	return o.AmfLocation, true
}

// HasAmfLocation returns a boolean if a field has been set.
func (o *AmfLocationData) HasAmfLocation() bool {
	if o != nil && !IsNil(o.AmfLocation) {
		return true
	}

	return false
}

// SetAmfLocation gets a reference to the given NrLocation and assigns it to the AmfLocation field.
func (o *AmfLocationData) SetAmfLocation(v NrLocation) {
	o.AmfLocation = &v
}

// GetSmsfAddress returns the SmsfAddress field value if set, zero value otherwise.
func (o *AmfLocationData) GetSmsfAddress() string {
	if o == nil || IsNil(o.SmsfAddress) {
		var ret string
		return ret
	}
	return *o.SmsfAddress
}

// GetSmsfAddressOk returns a tuple with the SmsfAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfLocationData) GetSmsfAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SmsfAddress) {
		return nil, false
	}
	return o.SmsfAddress, true
}

// HasSmsfAddress returns a boolean if a field has been set.
func (o *AmfLocationData) HasSmsfAddress() bool {
	if o != nil && !IsNil(o.SmsfAddress) {
		return true
	}

	return false
}

// SetSmsfAddress gets a reference to the given string and assigns it to the SmsfAddress field.
func (o *AmfLocationData) SetSmsfAddress(v string) {
	o.SmsfAddress = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *AmfLocationData) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfLocationData) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *AmfLocationData) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *AmfLocationData) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *AmfLocationData) GetRatType() RatType {
	if o == nil || IsNil(o.RatType) {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfLocationData) GetRatTypeOk() (*RatType, bool) {
	if o == nil || IsNil(o.RatType) {
		return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *AmfLocationData) HasRatType() bool {
	if o != nil && !IsNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *AmfLocationData) SetRatType(v RatType) {
	o.RatType = &v
}

func (o AmfLocationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfLocationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amfAddress"] = o.AmfAddress
	toSerialize["plmnId"] = o.PlmnId
	if !IsNil(o.AmfLocation) {
		toSerialize["amfLocation"] = o.AmfLocation
	}
	if !IsNil(o.SmsfAddress) {
		toSerialize["smsfAddress"] = o.SmsfAddress
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	return toSerialize, nil
}

type NullableAmfLocationData struct {
	value *AmfLocationData
	isSet bool
}

func (v NullableAmfLocationData) Get() *AmfLocationData {
	return v.value
}

func (v *NullableAmfLocationData) Set(val *AmfLocationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfLocationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfLocationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfLocationData(val *AmfLocationData) *NullableAmfLocationData {
	return &NullableAmfLocationData{value: val, isSet: true}
}

func (v NullableAmfLocationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfLocationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

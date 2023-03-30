/*
EES UE Location Information_API

API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_UELocation

import (
	"encoding/json"
)

// checks if the LocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationRequest{}

// LocationRequest To request location information request.
type LocationRequest struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	UeId string `json:"ueId"`
	Gran *Accuracy `json:"gran,omitempty"`
	LocQos *LocationQoS `json:"locQos,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewLocationRequest instantiates a new LocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationRequest(ueId string) *LocationRequest {
	this := LocationRequest{}
	this.UeId = ueId
	return &this
}

// NewLocationRequestWithDefaults instantiates a new LocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationRequestWithDefaults() *LocationRequest {
	this := LocationRequest{}
	return &this
}

// GetUeId returns the UeId field value
func (o *LocationRequest) GetUeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetUeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UeId, true
}

// SetUeId sets field value
func (o *LocationRequest) SetUeId(v string) {
	o.UeId = v
}

// GetGran returns the Gran field value if set, zero value otherwise.
func (o *LocationRequest) GetGran() Accuracy {
	if o == nil || IsNil(o.Gran) {
		var ret Accuracy
		return ret
	}
	return *o.Gran
}

// GetGranOk returns a tuple with the Gran field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetGranOk() (*Accuracy, bool) {
	if o == nil || IsNil(o.Gran) {
		return nil, false
	}
	return o.Gran, true
}

// HasGran returns a boolean if a field has been set.
func (o *LocationRequest) HasGran() bool {
	if o != nil && !IsNil(o.Gran) {
		return true
	}

	return false
}

// SetGran gets a reference to the given Accuracy and assigns it to the Gran field.
func (o *LocationRequest) SetGran(v Accuracy) {
	o.Gran = &v
}

// GetLocQos returns the LocQos field value if set, zero value otherwise.
func (o *LocationRequest) GetLocQos() LocationQoS {
	if o == nil || IsNil(o.LocQos) {
		var ret LocationQoS
		return ret
	}
	return *o.LocQos
}

// GetLocQosOk returns a tuple with the LocQos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetLocQosOk() (*LocationQoS, bool) {
	if o == nil || IsNil(o.LocQos) {
		return nil, false
	}
	return o.LocQos, true
}

// HasLocQos returns a boolean if a field has been set.
func (o *LocationRequest) HasLocQos() bool {
	if o != nil && !IsNil(o.LocQos) {
		return true
	}

	return false
}

// SetLocQos gets a reference to the given LocationQoS and assigns it to the LocQos field.
func (o *LocationRequest) SetLocQos(v LocationQoS) {
	o.LocQos = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *LocationRequest) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *LocationRequest) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *LocationRequest) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o LocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ueId"] = o.UeId
	if !IsNil(o.Gran) {
		toSerialize["gran"] = o.Gran
	}
	if !IsNil(o.LocQos) {
		toSerialize["locQos"] = o.LocQos
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableLocationRequest struct {
	value *LocationRequest
	isSet bool
}

func (v NullableLocationRequest) Get() *LocationRequest {
	return v.value
}

func (v *NullableLocationRequest) Set(val *LocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationRequest(val *LocationRequest) *NullableLocationRequest {
	return &NullableLocationRequest{value: val, isSet: true}
}

func (v NullableLocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


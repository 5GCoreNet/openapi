/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the LocationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationInfo{}

// LocationInfo Represents UE location information.
type LocationInfo struct {
	Loc UserLocation `json:"loc"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	Ratio *int32 `json:"ratio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
}

// NewLocationInfo instantiates a new LocationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationInfo(loc UserLocation) *LocationInfo {
	this := LocationInfo{}
	this.Loc = loc
	return &this
}

// NewLocationInfoWithDefaults instantiates a new LocationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationInfoWithDefaults() *LocationInfo {
	this := LocationInfo{}
	return &this
}

// GetLoc returns the Loc field value
func (o *LocationInfo) GetLoc() UserLocation {
	if o == nil {
		var ret UserLocation
		return ret
	}

	return o.Loc
}

// GetLocOk returns a tuple with the Loc field value
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetLocOk() (*UserLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Loc, true
}

// SetLoc sets field value
func (o *LocationInfo) SetLoc(v UserLocation) {
	o.Loc = v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *LocationInfo) GetRatio() int32 {
	if o == nil || isNil(o.Ratio) {
		var ret int32
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetRatioOk() (*int32, bool) {
	if o == nil || isNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *LocationInfo) HasRatio() bool {
	if o != nil && !isNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given int32 and assigns it to the Ratio field.
func (o *LocationInfo) SetRatio(v int32) {
	o.Ratio = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *LocationInfo) GetConfidence() int32 {
	if o == nil || isNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetConfidenceOk() (*int32, bool) {
	if o == nil || isNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *LocationInfo) HasConfidence() bool {
	if o != nil && !isNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *LocationInfo) SetConfidence(v int32) {
	o.Confidence = &v
}

func (o LocationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["loc"] = o.Loc
	if !isNil(o.Ratio) {
		toSerialize["ratio"] = o.Ratio
	}
	if !isNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	return toSerialize, nil
}

type NullableLocationInfo struct {
	value *LocationInfo
	isSet bool
}

func (v NullableLocationInfo) Get() *LocationInfo {
	return v.value
}

func (v *NullableLocationInfo) Set(val *LocationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationInfo(val *LocationInfo) *NullableLocationInfo {
	return &NullableLocationInfo{value: val, isSet: true}
}

func (v NullableLocationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



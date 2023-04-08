/*
Nnwdaf_MLModelProvision

Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_MLModelProvision

import (
	"encoding/json"
)

// checks if the RatFreqInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatFreqInformation{}

// RatFreqInformation Represents the RAT type and/or Frequency information.
type RatFreqInformation struct {
	// Set to \"true\" to indicate to handle all the frequencies the NWDAF received, otherwise  set to \"false\" or omit. The \"allFreq\" attribute and the \"freq\" attribute are mutually  exclusive. 
	AllFreq *bool `json:"allFreq,omitempty"`
	// Set to \"true\" to indicate to handle all the RAT Types the NWDAF received, otherwise  set to \"false\" or omit. The \"allRat\" attribute and the \"ratType\" attribute are mutually  exclusive. 
	AllRat *bool `json:"allRat,omitempty"`
	// Integer value indicating the ARFCN applicable for a downlink, uplink or bi-directional (TDD) NR global frequency raster, as definition of \"ARFCN-ValueNR\" IE in clause 6.3.2 of 3GPP TS 38.331. 
	Freq *int32 `json:"freq,omitempty"`
	RatType *RatType `json:"ratType,omitempty"`
	SvcExpThreshold *ThresholdLevel `json:"svcExpThreshold,omitempty"`
	MatchingDir *MatchingDirection `json:"matchingDir,omitempty"`
}

// NewRatFreqInformation instantiates a new RatFreqInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatFreqInformation() *RatFreqInformation {
	this := RatFreqInformation{}
	return &this
}

// NewRatFreqInformationWithDefaults instantiates a new RatFreqInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatFreqInformationWithDefaults() *RatFreqInformation {
	this := RatFreqInformation{}
	return &this
}

// GetAllFreq returns the AllFreq field value if set, zero value otherwise.
func (o *RatFreqInformation) GetAllFreq() bool {
	if o == nil || isNil(o.AllFreq) {
		var ret bool
		return ret
	}
	return *o.AllFreq
}

// GetAllFreqOk returns a tuple with the AllFreq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatFreqInformation) GetAllFreqOk() (*bool, bool) {
	if o == nil || isNil(o.AllFreq) {
		return nil, false
	}
	return o.AllFreq, true
}

// HasAllFreq returns a boolean if a field has been set.
func (o *RatFreqInformation) HasAllFreq() bool {
	if o != nil && !isNil(o.AllFreq) {
		return true
	}

	return false
}

// SetAllFreq gets a reference to the given bool and assigns it to the AllFreq field.
func (o *RatFreqInformation) SetAllFreq(v bool) {
	o.AllFreq = &v
}

// GetAllRat returns the AllRat field value if set, zero value otherwise.
func (o *RatFreqInformation) GetAllRat() bool {
	if o == nil || isNil(o.AllRat) {
		var ret bool
		return ret
	}
	return *o.AllRat
}

// GetAllRatOk returns a tuple with the AllRat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatFreqInformation) GetAllRatOk() (*bool, bool) {
	if o == nil || isNil(o.AllRat) {
		return nil, false
	}
	return o.AllRat, true
}

// HasAllRat returns a boolean if a field has been set.
func (o *RatFreqInformation) HasAllRat() bool {
	if o != nil && !isNil(o.AllRat) {
		return true
	}

	return false
}

// SetAllRat gets a reference to the given bool and assigns it to the AllRat field.
func (o *RatFreqInformation) SetAllRat(v bool) {
	o.AllRat = &v
}

// GetFreq returns the Freq field value if set, zero value otherwise.
func (o *RatFreqInformation) GetFreq() int32 {
	if o == nil || isNil(o.Freq) {
		var ret int32
		return ret
	}
	return *o.Freq
}

// GetFreqOk returns a tuple with the Freq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatFreqInformation) GetFreqOk() (*int32, bool) {
	if o == nil || isNil(o.Freq) {
		return nil, false
	}
	return o.Freq, true
}

// HasFreq returns a boolean if a field has been set.
func (o *RatFreqInformation) HasFreq() bool {
	if o != nil && !isNil(o.Freq) {
		return true
	}

	return false
}

// SetFreq gets a reference to the given int32 and assigns it to the Freq field.
func (o *RatFreqInformation) SetFreq(v int32) {
	o.Freq = &v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *RatFreqInformation) GetRatType() RatType {
	if o == nil || isNil(o.RatType) {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatFreqInformation) GetRatTypeOk() (*RatType, bool) {
	if o == nil || isNil(o.RatType) {
		return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *RatFreqInformation) HasRatType() bool {
	if o != nil && !isNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *RatFreqInformation) SetRatType(v RatType) {
	o.RatType = &v
}

// GetSvcExpThreshold returns the SvcExpThreshold field value if set, zero value otherwise.
func (o *RatFreqInformation) GetSvcExpThreshold() ThresholdLevel {
	if o == nil || isNil(o.SvcExpThreshold) {
		var ret ThresholdLevel
		return ret
	}
	return *o.SvcExpThreshold
}

// GetSvcExpThresholdOk returns a tuple with the SvcExpThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatFreqInformation) GetSvcExpThresholdOk() (*ThresholdLevel, bool) {
	if o == nil || isNil(o.SvcExpThreshold) {
		return nil, false
	}
	return o.SvcExpThreshold, true
}

// HasSvcExpThreshold returns a boolean if a field has been set.
func (o *RatFreqInformation) HasSvcExpThreshold() bool {
	if o != nil && !isNil(o.SvcExpThreshold) {
		return true
	}

	return false
}

// SetSvcExpThreshold gets a reference to the given ThresholdLevel and assigns it to the SvcExpThreshold field.
func (o *RatFreqInformation) SetSvcExpThreshold(v ThresholdLevel) {
	o.SvcExpThreshold = &v
}

// GetMatchingDir returns the MatchingDir field value if set, zero value otherwise.
func (o *RatFreqInformation) GetMatchingDir() MatchingDirection {
	if o == nil || isNil(o.MatchingDir) {
		var ret MatchingDirection
		return ret
	}
	return *o.MatchingDir
}

// GetMatchingDirOk returns a tuple with the MatchingDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatFreqInformation) GetMatchingDirOk() (*MatchingDirection, bool) {
	if o == nil || isNil(o.MatchingDir) {
		return nil, false
	}
	return o.MatchingDir, true
}

// HasMatchingDir returns a boolean if a field has been set.
func (o *RatFreqInformation) HasMatchingDir() bool {
	if o != nil && !isNil(o.MatchingDir) {
		return true
	}

	return false
}

// SetMatchingDir gets a reference to the given MatchingDirection and assigns it to the MatchingDir field.
func (o *RatFreqInformation) SetMatchingDir(v MatchingDirection) {
	o.MatchingDir = &v
}

func (o RatFreqInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatFreqInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllFreq) {
		toSerialize["allFreq"] = o.AllFreq
	}
	if !isNil(o.AllRat) {
		toSerialize["allRat"] = o.AllRat
	}
	if !isNil(o.Freq) {
		toSerialize["freq"] = o.Freq
	}
	if !isNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	if !isNil(o.SvcExpThreshold) {
		toSerialize["svcExpThreshold"] = o.SvcExpThreshold
	}
	if !isNil(o.MatchingDir) {
		toSerialize["matchingDir"] = o.MatchingDir
	}
	return toSerialize, nil
}

type NullableRatFreqInformation struct {
	value *RatFreqInformation
	isSet bool
}

func (v NullableRatFreqInformation) Get() *RatFreqInformation {
	return v.value
}

func (v *NullableRatFreqInformation) Set(val *RatFreqInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableRatFreqInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableRatFreqInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatFreqInformation(val *RatFreqInformation) *NullableRatFreqInformation {
	return &NullableRatFreqInformation{value: val, isSet: true}
}

func (v NullableRatFreqInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatFreqInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



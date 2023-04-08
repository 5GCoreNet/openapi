/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReporting

import (
	"encoding/json"
)

// checks if the SvcExperience type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SvcExperience{}

// SvcExperience Contains a mean opinion score with the customized range.
type SvcExperience struct {
	// string with format 'float' as defined in OpenAPI.
	Mos *float32 `json:"mos,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	UpperRange *float32 `json:"upperRange,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	LowerRange *float32 `json:"lowerRange,omitempty"`
}

// NewSvcExperience instantiates a new SvcExperience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSvcExperience() *SvcExperience {
	this := SvcExperience{}
	return &this
}

// NewSvcExperienceWithDefaults instantiates a new SvcExperience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSvcExperienceWithDefaults() *SvcExperience {
	this := SvcExperience{}
	return &this
}

// GetMos returns the Mos field value if set, zero value otherwise.
func (o *SvcExperience) GetMos() float32 {
	if o == nil || isNil(o.Mos) {
		var ret float32
		return ret
	}
	return *o.Mos
}

// GetMosOk returns a tuple with the Mos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SvcExperience) GetMosOk() (*float32, bool) {
	if o == nil || isNil(o.Mos) {
		return nil, false
	}
	return o.Mos, true
}

// HasMos returns a boolean if a field has been set.
func (o *SvcExperience) HasMos() bool {
	if o != nil && !isNil(o.Mos) {
		return true
	}

	return false
}

// SetMos gets a reference to the given float32 and assigns it to the Mos field.
func (o *SvcExperience) SetMos(v float32) {
	o.Mos = &v
}

// GetUpperRange returns the UpperRange field value if set, zero value otherwise.
func (o *SvcExperience) GetUpperRange() float32 {
	if o == nil || isNil(o.UpperRange) {
		var ret float32
		return ret
	}
	return *o.UpperRange
}

// GetUpperRangeOk returns a tuple with the UpperRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SvcExperience) GetUpperRangeOk() (*float32, bool) {
	if o == nil || isNil(o.UpperRange) {
		return nil, false
	}
	return o.UpperRange, true
}

// HasUpperRange returns a boolean if a field has been set.
func (o *SvcExperience) HasUpperRange() bool {
	if o != nil && !isNil(o.UpperRange) {
		return true
	}

	return false
}

// SetUpperRange gets a reference to the given float32 and assigns it to the UpperRange field.
func (o *SvcExperience) SetUpperRange(v float32) {
	o.UpperRange = &v
}

// GetLowerRange returns the LowerRange field value if set, zero value otherwise.
func (o *SvcExperience) GetLowerRange() float32 {
	if o == nil || isNil(o.LowerRange) {
		var ret float32
		return ret
	}
	return *o.LowerRange
}

// GetLowerRangeOk returns a tuple with the LowerRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SvcExperience) GetLowerRangeOk() (*float32, bool) {
	if o == nil || isNil(o.LowerRange) {
		return nil, false
	}
	return o.LowerRange, true
}

// HasLowerRange returns a boolean if a field has been set.
func (o *SvcExperience) HasLowerRange() bool {
	if o != nil && !isNil(o.LowerRange) {
		return true
	}

	return false
}

// SetLowerRange gets a reference to the given float32 and assigns it to the LowerRange field.
func (o *SvcExperience) SetLowerRange(v float32) {
	o.LowerRange = &v
}

func (o SvcExperience) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SvcExperience) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mos) {
		toSerialize["mos"] = o.Mos
	}
	if !isNil(o.UpperRange) {
		toSerialize["upperRange"] = o.UpperRange
	}
	if !isNil(o.LowerRange) {
		toSerialize["lowerRange"] = o.LowerRange
	}
	return toSerialize, nil
}

type NullableSvcExperience struct {
	value *SvcExperience
	isSet bool
}

func (v NullableSvcExperience) Get() *SvcExperience {
	return v.value
}

func (v *NullableSvcExperience) Set(val *SvcExperience) {
	v.value = val
	v.isSet = true
}

func (v NullableSvcExperience) IsSet() bool {
	return v.isSet
}

func (v *NullableSvcExperience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSvcExperience(val *SvcExperience) *NullableSvcExperience {
	return &NullableSvcExperience{value: val, isSet: true}
}

func (v NullableSvcExperience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSvcExperience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



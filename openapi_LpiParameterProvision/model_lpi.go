/*
3gpp-lpi-pp

API for Location Privacy Indication Parameters Provisioning.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_LpiParameterProvision

import (
	"encoding/json"
)

// checks if the Lpi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Lpi{}

// Lpi struct for Lpi
type Lpi struct {
	LocationPrivacyInd LocationPrivacyInd `json:"locationPrivacyInd"`
	ValidTimePeriod *ValidTimePeriod `json:"validTimePeriod,omitempty"`
}

// NewLpi instantiates a new Lpi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLpi(locationPrivacyInd LocationPrivacyInd) *Lpi {
	this := Lpi{}
	this.LocationPrivacyInd = locationPrivacyInd
	return &this
}

// NewLpiWithDefaults instantiates a new Lpi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLpiWithDefaults() *Lpi {
	this := Lpi{}
	return &this
}

// GetLocationPrivacyInd returns the LocationPrivacyInd field value
func (o *Lpi) GetLocationPrivacyInd() LocationPrivacyInd {
	if o == nil {
		var ret LocationPrivacyInd
		return ret
	}

	return o.LocationPrivacyInd
}

// GetLocationPrivacyIndOk returns a tuple with the LocationPrivacyInd field value
// and a boolean to check if the value has been set.
func (o *Lpi) GetLocationPrivacyIndOk() (*LocationPrivacyInd, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationPrivacyInd, true
}

// SetLocationPrivacyInd sets field value
func (o *Lpi) SetLocationPrivacyInd(v LocationPrivacyInd) {
	o.LocationPrivacyInd = v
}

// GetValidTimePeriod returns the ValidTimePeriod field value if set, zero value otherwise.
func (o *Lpi) GetValidTimePeriod() ValidTimePeriod {
	if o == nil || IsNil(o.ValidTimePeriod) {
		var ret ValidTimePeriod
		return ret
	}
	return *o.ValidTimePeriod
}

// GetValidTimePeriodOk returns a tuple with the ValidTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lpi) GetValidTimePeriodOk() (*ValidTimePeriod, bool) {
	if o == nil || IsNil(o.ValidTimePeriod) {
		return nil, false
	}
	return o.ValidTimePeriod, true
}

// HasValidTimePeriod returns a boolean if a field has been set.
func (o *Lpi) HasValidTimePeriod() bool {
	if o != nil && !IsNil(o.ValidTimePeriod) {
		return true
	}

	return false
}

// SetValidTimePeriod gets a reference to the given ValidTimePeriod and assigns it to the ValidTimePeriod field.
func (o *Lpi) SetValidTimePeriod(v ValidTimePeriod) {
	o.ValidTimePeriod = &v
}

func (o Lpi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Lpi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locationPrivacyInd"] = o.LocationPrivacyInd
	if !IsNil(o.ValidTimePeriod) {
		toSerialize["validTimePeriod"] = o.ValidTimePeriod
	}
	return toSerialize, nil
}

type NullableLpi struct {
	value *Lpi
	isSet bool
}

func (v NullableLpi) Get() *Lpi {
	return v.value
}

func (v *NullableLpi) Set(val *Lpi) {
	v.value = val
	v.isSet = true
}

func (v NullableLpi) IsSet() bool {
	return v.isSet
}

func (v *NullableLpi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLpi(val *Lpi) *NullableLpi {
	return &NullableLpi{value: val, isSet: true}
}

func (v NullableLpi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLpi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


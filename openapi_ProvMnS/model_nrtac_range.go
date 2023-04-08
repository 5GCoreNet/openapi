/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the NRTACRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NRTACRange{}

// NRTACRange struct for NRTACRange
type NRTACRange struct {
	NRTACstart *string `json:"nRTACstart,omitempty"`
	NRTACend *string `json:"nRTACend,omitempty"`
	NRTACpattern *string `json:"nRTACpattern,omitempty"`
}

// NewNRTACRange instantiates a new NRTACRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNRTACRange() *NRTACRange {
	this := NRTACRange{}
	return &this
}

// NewNRTACRangeWithDefaults instantiates a new NRTACRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNRTACRangeWithDefaults() *NRTACRange {
	this := NRTACRange{}
	return &this
}

// GetNRTACstart returns the NRTACstart field value if set, zero value otherwise.
func (o *NRTACRange) GetNRTACstart() string {
	if o == nil || isNil(o.NRTACstart) {
		var ret string
		return ret
	}
	return *o.NRTACstart
}

// GetNRTACstartOk returns a tuple with the NRTACstart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRTACRange) GetNRTACstartOk() (*string, bool) {
	if o == nil || isNil(o.NRTACstart) {
		return nil, false
	}
	return o.NRTACstart, true
}

// HasNRTACstart returns a boolean if a field has been set.
func (o *NRTACRange) HasNRTACstart() bool {
	if o != nil && !isNil(o.NRTACstart) {
		return true
	}

	return false
}

// SetNRTACstart gets a reference to the given string and assigns it to the NRTACstart field.
func (o *NRTACRange) SetNRTACstart(v string) {
	o.NRTACstart = &v
}

// GetNRTACend returns the NRTACend field value if set, zero value otherwise.
func (o *NRTACRange) GetNRTACend() string {
	if o == nil || isNil(o.NRTACend) {
		var ret string
		return ret
	}
	return *o.NRTACend
}

// GetNRTACendOk returns a tuple with the NRTACend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRTACRange) GetNRTACendOk() (*string, bool) {
	if o == nil || isNil(o.NRTACend) {
		return nil, false
	}
	return o.NRTACend, true
}

// HasNRTACend returns a boolean if a field has been set.
func (o *NRTACRange) HasNRTACend() bool {
	if o != nil && !isNil(o.NRTACend) {
		return true
	}

	return false
}

// SetNRTACend gets a reference to the given string and assigns it to the NRTACend field.
func (o *NRTACRange) SetNRTACend(v string) {
	o.NRTACend = &v
}

// GetNRTACpattern returns the NRTACpattern field value if set, zero value otherwise.
func (o *NRTACRange) GetNRTACpattern() string {
	if o == nil || isNil(o.NRTACpattern) {
		var ret string
		return ret
	}
	return *o.NRTACpattern
}

// GetNRTACpatternOk returns a tuple with the NRTACpattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRTACRange) GetNRTACpatternOk() (*string, bool) {
	if o == nil || isNil(o.NRTACpattern) {
		return nil, false
	}
	return o.NRTACpattern, true
}

// HasNRTACpattern returns a boolean if a field has been set.
func (o *NRTACRange) HasNRTACpattern() bool {
	if o != nil && !isNil(o.NRTACpattern) {
		return true
	}

	return false
}

// SetNRTACpattern gets a reference to the given string and assigns it to the NRTACpattern field.
func (o *NRTACRange) SetNRTACpattern(v string) {
	o.NRTACpattern = &v
}

func (o NRTACRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NRTACRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NRTACstart) {
		toSerialize["nRTACstart"] = o.NRTACstart
	}
	if !isNil(o.NRTACend) {
		toSerialize["nRTACend"] = o.NRTACend
	}
	if !isNil(o.NRTACpattern) {
		toSerialize["nRTACpattern"] = o.NRTACpattern
	}
	return toSerialize, nil
}

type NullableNRTACRange struct {
	value *NRTACRange
	isSet bool
}

func (v NullableNRTACRange) Get() *NRTACRange {
	return v.value
}

func (v *NullableNRTACRange) Set(val *NRTACRange) {
	v.value = val
	v.isSet = true
}

func (v NullableNRTACRange) IsSet() bool {
	return v.isSet
}

func (v *NullableNRTACRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNRTACRange(val *NRTACRange) *NullableNRTACRange {
	return &NullableNRTACRange{value: val, isSet: true}
}

func (v NullableNRTACRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNRTACRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



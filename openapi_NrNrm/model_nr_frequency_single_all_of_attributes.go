/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the NRFrequencySingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NRFrequencySingleAllOfAttributes{}

// NRFrequencySingleAllOfAttributes struct for NRFrequencySingleAllOfAttributes
type NRFrequencySingleAllOfAttributes struct {
	AbsoluteFrequencySSB *int32 `json:"absoluteFrequencySSB,omitempty"`
	SsbSubCarrierSpacing *SsbSubCarrierSpacing `json:"ssbSubCarrierSpacing,omitempty"`
	MultiFrequencyBandListNR *int32 `json:"multiFrequencyBandListNR,omitempty"`
}

// NewNRFrequencySingleAllOfAttributes instantiates a new NRFrequencySingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNRFrequencySingleAllOfAttributes() *NRFrequencySingleAllOfAttributes {
	this := NRFrequencySingleAllOfAttributes{}
	return &this
}

// NewNRFrequencySingleAllOfAttributesWithDefaults instantiates a new NRFrequencySingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNRFrequencySingleAllOfAttributesWithDefaults() *NRFrequencySingleAllOfAttributes {
	this := NRFrequencySingleAllOfAttributes{}
	return &this
}

// GetAbsoluteFrequencySSB returns the AbsoluteFrequencySSB field value if set, zero value otherwise.
func (o *NRFrequencySingleAllOfAttributes) GetAbsoluteFrequencySSB() int32 {
	if o == nil || IsNil(o.AbsoluteFrequencySSB) {
		var ret int32
		return ret
	}
	return *o.AbsoluteFrequencySSB
}

// GetAbsoluteFrequencySSBOk returns a tuple with the AbsoluteFrequencySSB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRFrequencySingleAllOfAttributes) GetAbsoluteFrequencySSBOk() (*int32, bool) {
	if o == nil || IsNil(o.AbsoluteFrequencySSB) {
		return nil, false
	}
	return o.AbsoluteFrequencySSB, true
}

// HasAbsoluteFrequencySSB returns a boolean if a field has been set.
func (o *NRFrequencySingleAllOfAttributes) HasAbsoluteFrequencySSB() bool {
	if o != nil && !IsNil(o.AbsoluteFrequencySSB) {
		return true
	}

	return false
}

// SetAbsoluteFrequencySSB gets a reference to the given int32 and assigns it to the AbsoluteFrequencySSB field.
func (o *NRFrequencySingleAllOfAttributes) SetAbsoluteFrequencySSB(v int32) {
	o.AbsoluteFrequencySSB = &v
}

// GetSsbSubCarrierSpacing returns the SsbSubCarrierSpacing field value if set, zero value otherwise.
func (o *NRFrequencySingleAllOfAttributes) GetSsbSubCarrierSpacing() SsbSubCarrierSpacing {
	if o == nil || IsNil(o.SsbSubCarrierSpacing) {
		var ret SsbSubCarrierSpacing
		return ret
	}
	return *o.SsbSubCarrierSpacing
}

// GetSsbSubCarrierSpacingOk returns a tuple with the SsbSubCarrierSpacing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRFrequencySingleAllOfAttributes) GetSsbSubCarrierSpacingOk() (*SsbSubCarrierSpacing, bool) {
	if o == nil || IsNil(o.SsbSubCarrierSpacing) {
		return nil, false
	}
	return o.SsbSubCarrierSpacing, true
}

// HasSsbSubCarrierSpacing returns a boolean if a field has been set.
func (o *NRFrequencySingleAllOfAttributes) HasSsbSubCarrierSpacing() bool {
	if o != nil && !IsNil(o.SsbSubCarrierSpacing) {
		return true
	}

	return false
}

// SetSsbSubCarrierSpacing gets a reference to the given SsbSubCarrierSpacing and assigns it to the SsbSubCarrierSpacing field.
func (o *NRFrequencySingleAllOfAttributes) SetSsbSubCarrierSpacing(v SsbSubCarrierSpacing) {
	o.SsbSubCarrierSpacing = &v
}

// GetMultiFrequencyBandListNR returns the MultiFrequencyBandListNR field value if set, zero value otherwise.
func (o *NRFrequencySingleAllOfAttributes) GetMultiFrequencyBandListNR() int32 {
	if o == nil || IsNil(o.MultiFrequencyBandListNR) {
		var ret int32
		return ret
	}
	return *o.MultiFrequencyBandListNR
}

// GetMultiFrequencyBandListNROk returns a tuple with the MultiFrequencyBandListNR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRFrequencySingleAllOfAttributes) GetMultiFrequencyBandListNROk() (*int32, bool) {
	if o == nil || IsNil(o.MultiFrequencyBandListNR) {
		return nil, false
	}
	return o.MultiFrequencyBandListNR, true
}

// HasMultiFrequencyBandListNR returns a boolean if a field has been set.
func (o *NRFrequencySingleAllOfAttributes) HasMultiFrequencyBandListNR() bool {
	if o != nil && !IsNil(o.MultiFrequencyBandListNR) {
		return true
	}

	return false
}

// SetMultiFrequencyBandListNR gets a reference to the given int32 and assigns it to the MultiFrequencyBandListNR field.
func (o *NRFrequencySingleAllOfAttributes) SetMultiFrequencyBandListNR(v int32) {
	o.MultiFrequencyBandListNR = &v
}

func (o NRFrequencySingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NRFrequencySingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbsoluteFrequencySSB) {
		toSerialize["absoluteFrequencySSB"] = o.AbsoluteFrequencySSB
	}
	if !IsNil(o.SsbSubCarrierSpacing) {
		toSerialize["ssbSubCarrierSpacing"] = o.SsbSubCarrierSpacing
	}
	if !IsNil(o.MultiFrequencyBandListNR) {
		toSerialize["multiFrequencyBandListNR"] = o.MultiFrequencyBandListNR
	}
	return toSerialize, nil
}

type NullableNRFrequencySingleAllOfAttributes struct {
	value *NRFrequencySingleAllOfAttributes
	isSet bool
}

func (v NullableNRFrequencySingleAllOfAttributes) Get() *NRFrequencySingleAllOfAttributes {
	return v.value
}

func (v *NullableNRFrequencySingleAllOfAttributes) Set(val *NRFrequencySingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNRFrequencySingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNRFrequencySingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNRFrequencySingleAllOfAttributes(val *NRFrequencySingleAllOfAttributes) *NullableNRFrequencySingleAllOfAttributes {
	return &NullableNRFrequencySingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableNRFrequencySingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNRFrequencySingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



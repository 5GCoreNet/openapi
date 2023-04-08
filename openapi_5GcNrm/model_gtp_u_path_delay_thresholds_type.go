/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the GtpUPathDelayThresholdsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GtpUPathDelayThresholdsType{}

// GtpUPathDelayThresholdsType struct for GtpUPathDelayThresholdsType
type GtpUPathDelayThresholdsType struct {
	N3AveragePacketDelayThreshold *int32 `json:"n3AveragePacketDelayThreshold,omitempty"`
	N3MinPacketDelayThreshold *int32 `json:"n3MinPacketDelayThreshold,omitempty"`
	N3MaxPacketDelayThreshold *int32 `json:"n3MaxPacketDelayThreshold,omitempty"`
	N9AveragePacketDelayThreshold *int32 `json:"n9AveragePacketDelayThreshold,omitempty"`
	N9MinPacketDelayThreshold *int32 `json:"n9MinPacketDelayThreshold,omitempty"`
	N9MaxPacketDelayThreshold *int32 `json:"n9MaxPacketDelayThreshold,omitempty"`
}

// NewGtpUPathDelayThresholdsType instantiates a new GtpUPathDelayThresholdsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGtpUPathDelayThresholdsType() *GtpUPathDelayThresholdsType {
	this := GtpUPathDelayThresholdsType{}
	return &this
}

// NewGtpUPathDelayThresholdsTypeWithDefaults instantiates a new GtpUPathDelayThresholdsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGtpUPathDelayThresholdsTypeWithDefaults() *GtpUPathDelayThresholdsType {
	this := GtpUPathDelayThresholdsType{}
	return &this
}

// GetN3AveragePacketDelayThreshold returns the N3AveragePacketDelayThreshold field value if set, zero value otherwise.
func (o *GtpUPathDelayThresholdsType) GetN3AveragePacketDelayThreshold() int32 {
	if o == nil || isNil(o.N3AveragePacketDelayThreshold) {
		var ret int32
		return ret
	}
	return *o.N3AveragePacketDelayThreshold
}

// GetN3AveragePacketDelayThresholdOk returns a tuple with the N3AveragePacketDelayThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathDelayThresholdsType) GetN3AveragePacketDelayThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.N3AveragePacketDelayThreshold) {
		return nil, false
	}
	return o.N3AveragePacketDelayThreshold, true
}

// HasN3AveragePacketDelayThreshold returns a boolean if a field has been set.
func (o *GtpUPathDelayThresholdsType) HasN3AveragePacketDelayThreshold() bool {
	if o != nil && !isNil(o.N3AveragePacketDelayThreshold) {
		return true
	}

	return false
}

// SetN3AveragePacketDelayThreshold gets a reference to the given int32 and assigns it to the N3AveragePacketDelayThreshold field.
func (o *GtpUPathDelayThresholdsType) SetN3AveragePacketDelayThreshold(v int32) {
	o.N3AveragePacketDelayThreshold = &v
}

// GetN3MinPacketDelayThreshold returns the N3MinPacketDelayThreshold field value if set, zero value otherwise.
func (o *GtpUPathDelayThresholdsType) GetN3MinPacketDelayThreshold() int32 {
	if o == nil || isNil(o.N3MinPacketDelayThreshold) {
		var ret int32
		return ret
	}
	return *o.N3MinPacketDelayThreshold
}

// GetN3MinPacketDelayThresholdOk returns a tuple with the N3MinPacketDelayThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathDelayThresholdsType) GetN3MinPacketDelayThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.N3MinPacketDelayThreshold) {
		return nil, false
	}
	return o.N3MinPacketDelayThreshold, true
}

// HasN3MinPacketDelayThreshold returns a boolean if a field has been set.
func (o *GtpUPathDelayThresholdsType) HasN3MinPacketDelayThreshold() bool {
	if o != nil && !isNil(o.N3MinPacketDelayThreshold) {
		return true
	}

	return false
}

// SetN3MinPacketDelayThreshold gets a reference to the given int32 and assigns it to the N3MinPacketDelayThreshold field.
func (o *GtpUPathDelayThresholdsType) SetN3MinPacketDelayThreshold(v int32) {
	o.N3MinPacketDelayThreshold = &v
}

// GetN3MaxPacketDelayThreshold returns the N3MaxPacketDelayThreshold field value if set, zero value otherwise.
func (o *GtpUPathDelayThresholdsType) GetN3MaxPacketDelayThreshold() int32 {
	if o == nil || isNil(o.N3MaxPacketDelayThreshold) {
		var ret int32
		return ret
	}
	return *o.N3MaxPacketDelayThreshold
}

// GetN3MaxPacketDelayThresholdOk returns a tuple with the N3MaxPacketDelayThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathDelayThresholdsType) GetN3MaxPacketDelayThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.N3MaxPacketDelayThreshold) {
		return nil, false
	}
	return o.N3MaxPacketDelayThreshold, true
}

// HasN3MaxPacketDelayThreshold returns a boolean if a field has been set.
func (o *GtpUPathDelayThresholdsType) HasN3MaxPacketDelayThreshold() bool {
	if o != nil && !isNil(o.N3MaxPacketDelayThreshold) {
		return true
	}

	return false
}

// SetN3MaxPacketDelayThreshold gets a reference to the given int32 and assigns it to the N3MaxPacketDelayThreshold field.
func (o *GtpUPathDelayThresholdsType) SetN3MaxPacketDelayThreshold(v int32) {
	o.N3MaxPacketDelayThreshold = &v
}

// GetN9AveragePacketDelayThreshold returns the N9AveragePacketDelayThreshold field value if set, zero value otherwise.
func (o *GtpUPathDelayThresholdsType) GetN9AveragePacketDelayThreshold() int32 {
	if o == nil || isNil(o.N9AveragePacketDelayThreshold) {
		var ret int32
		return ret
	}
	return *o.N9AveragePacketDelayThreshold
}

// GetN9AveragePacketDelayThresholdOk returns a tuple with the N9AveragePacketDelayThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathDelayThresholdsType) GetN9AveragePacketDelayThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.N9AveragePacketDelayThreshold) {
		return nil, false
	}
	return o.N9AveragePacketDelayThreshold, true
}

// HasN9AveragePacketDelayThreshold returns a boolean if a field has been set.
func (o *GtpUPathDelayThresholdsType) HasN9AveragePacketDelayThreshold() bool {
	if o != nil && !isNil(o.N9AveragePacketDelayThreshold) {
		return true
	}

	return false
}

// SetN9AveragePacketDelayThreshold gets a reference to the given int32 and assigns it to the N9AveragePacketDelayThreshold field.
func (o *GtpUPathDelayThresholdsType) SetN9AveragePacketDelayThreshold(v int32) {
	o.N9AveragePacketDelayThreshold = &v
}

// GetN9MinPacketDelayThreshold returns the N9MinPacketDelayThreshold field value if set, zero value otherwise.
func (o *GtpUPathDelayThresholdsType) GetN9MinPacketDelayThreshold() int32 {
	if o == nil || isNil(o.N9MinPacketDelayThreshold) {
		var ret int32
		return ret
	}
	return *o.N9MinPacketDelayThreshold
}

// GetN9MinPacketDelayThresholdOk returns a tuple with the N9MinPacketDelayThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathDelayThresholdsType) GetN9MinPacketDelayThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.N9MinPacketDelayThreshold) {
		return nil, false
	}
	return o.N9MinPacketDelayThreshold, true
}

// HasN9MinPacketDelayThreshold returns a boolean if a field has been set.
func (o *GtpUPathDelayThresholdsType) HasN9MinPacketDelayThreshold() bool {
	if o != nil && !isNil(o.N9MinPacketDelayThreshold) {
		return true
	}

	return false
}

// SetN9MinPacketDelayThreshold gets a reference to the given int32 and assigns it to the N9MinPacketDelayThreshold field.
func (o *GtpUPathDelayThresholdsType) SetN9MinPacketDelayThreshold(v int32) {
	o.N9MinPacketDelayThreshold = &v
}

// GetN9MaxPacketDelayThreshold returns the N9MaxPacketDelayThreshold field value if set, zero value otherwise.
func (o *GtpUPathDelayThresholdsType) GetN9MaxPacketDelayThreshold() int32 {
	if o == nil || isNil(o.N9MaxPacketDelayThreshold) {
		var ret int32
		return ret
	}
	return *o.N9MaxPacketDelayThreshold
}

// GetN9MaxPacketDelayThresholdOk returns a tuple with the N9MaxPacketDelayThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathDelayThresholdsType) GetN9MaxPacketDelayThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.N9MaxPacketDelayThreshold) {
		return nil, false
	}
	return o.N9MaxPacketDelayThreshold, true
}

// HasN9MaxPacketDelayThreshold returns a boolean if a field has been set.
func (o *GtpUPathDelayThresholdsType) HasN9MaxPacketDelayThreshold() bool {
	if o != nil && !isNil(o.N9MaxPacketDelayThreshold) {
		return true
	}

	return false
}

// SetN9MaxPacketDelayThreshold gets a reference to the given int32 and assigns it to the N9MaxPacketDelayThreshold field.
func (o *GtpUPathDelayThresholdsType) SetN9MaxPacketDelayThreshold(v int32) {
	o.N9MaxPacketDelayThreshold = &v
}

func (o GtpUPathDelayThresholdsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GtpUPathDelayThresholdsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.N3AveragePacketDelayThreshold) {
		toSerialize["n3AveragePacketDelayThreshold"] = o.N3AveragePacketDelayThreshold
	}
	if !isNil(o.N3MinPacketDelayThreshold) {
		toSerialize["n3MinPacketDelayThreshold"] = o.N3MinPacketDelayThreshold
	}
	if !isNil(o.N3MaxPacketDelayThreshold) {
		toSerialize["n3MaxPacketDelayThreshold"] = o.N3MaxPacketDelayThreshold
	}
	if !isNil(o.N9AveragePacketDelayThreshold) {
		toSerialize["n9AveragePacketDelayThreshold"] = o.N9AveragePacketDelayThreshold
	}
	if !isNil(o.N9MinPacketDelayThreshold) {
		toSerialize["n9MinPacketDelayThreshold"] = o.N9MinPacketDelayThreshold
	}
	if !isNil(o.N9MaxPacketDelayThreshold) {
		toSerialize["n9MaxPacketDelayThreshold"] = o.N9MaxPacketDelayThreshold
	}
	return toSerialize, nil
}

type NullableGtpUPathDelayThresholdsType struct {
	value *GtpUPathDelayThresholdsType
	isSet bool
}

func (v NullableGtpUPathDelayThresholdsType) Get() *GtpUPathDelayThresholdsType {
	return v.value
}

func (v *NullableGtpUPathDelayThresholdsType) Set(val *GtpUPathDelayThresholdsType) {
	v.value = val
	v.isSet = true
}

func (v NullableGtpUPathDelayThresholdsType) IsSet() bool {
	return v.isSet
}

func (v *NullableGtpUPathDelayThresholdsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGtpUPathDelayThresholdsType(val *GtpUPathDelayThresholdsType) *NullableGtpUPathDelayThresholdsType {
	return &NullableGtpUPathDelayThresholdsType{value: val, isSet: true}
}

func (v NullableGtpUPathDelayThresholdsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGtpUPathDelayThresholdsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



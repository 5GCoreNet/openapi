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

// checks if the InterRatEsActivationOriginalCellParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterRatEsActivationOriginalCellParameters{}

// InterRatEsActivationOriginalCellParameters struct for InterRatEsActivationOriginalCellParameters
type InterRatEsActivationOriginalCellParameters struct {
	LoadThreshold *int32 `json:"loadThreshold,omitempty"`
	TimeDuration *int32 `json:"timeDuration,omitempty"`
}

// NewInterRatEsActivationOriginalCellParameters instantiates a new InterRatEsActivationOriginalCellParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterRatEsActivationOriginalCellParameters() *InterRatEsActivationOriginalCellParameters {
	this := InterRatEsActivationOriginalCellParameters{}
	return &this
}

// NewInterRatEsActivationOriginalCellParametersWithDefaults instantiates a new InterRatEsActivationOriginalCellParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterRatEsActivationOriginalCellParametersWithDefaults() *InterRatEsActivationOriginalCellParameters {
	this := InterRatEsActivationOriginalCellParameters{}
	return &this
}

// GetLoadThreshold returns the LoadThreshold field value if set, zero value otherwise.
func (o *InterRatEsActivationOriginalCellParameters) GetLoadThreshold() int32 {
	if o == nil || isNil(o.LoadThreshold) {
		var ret int32
		return ret
	}
	return *o.LoadThreshold
}

// GetLoadThresholdOk returns a tuple with the LoadThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterRatEsActivationOriginalCellParameters) GetLoadThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.LoadThreshold) {
		return nil, false
	}
	return o.LoadThreshold, true
}

// HasLoadThreshold returns a boolean if a field has been set.
func (o *InterRatEsActivationOriginalCellParameters) HasLoadThreshold() bool {
	if o != nil && !isNil(o.LoadThreshold) {
		return true
	}

	return false
}

// SetLoadThreshold gets a reference to the given int32 and assigns it to the LoadThreshold field.
func (o *InterRatEsActivationOriginalCellParameters) SetLoadThreshold(v int32) {
	o.LoadThreshold = &v
}

// GetTimeDuration returns the TimeDuration field value if set, zero value otherwise.
func (o *InterRatEsActivationOriginalCellParameters) GetTimeDuration() int32 {
	if o == nil || isNil(o.TimeDuration) {
		var ret int32
		return ret
	}
	return *o.TimeDuration
}

// GetTimeDurationOk returns a tuple with the TimeDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterRatEsActivationOriginalCellParameters) GetTimeDurationOk() (*int32, bool) {
	if o == nil || isNil(o.TimeDuration) {
		return nil, false
	}
	return o.TimeDuration, true
}

// HasTimeDuration returns a boolean if a field has been set.
func (o *InterRatEsActivationOriginalCellParameters) HasTimeDuration() bool {
	if o != nil && !isNil(o.TimeDuration) {
		return true
	}

	return false
}

// SetTimeDuration gets a reference to the given int32 and assigns it to the TimeDuration field.
func (o *InterRatEsActivationOriginalCellParameters) SetTimeDuration(v int32) {
	o.TimeDuration = &v
}

func (o InterRatEsActivationOriginalCellParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterRatEsActivationOriginalCellParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LoadThreshold) {
		toSerialize["loadThreshold"] = o.LoadThreshold
	}
	if !isNil(o.TimeDuration) {
		toSerialize["timeDuration"] = o.TimeDuration
	}
	return toSerialize, nil
}

type NullableInterRatEsActivationOriginalCellParameters struct {
	value *InterRatEsActivationOriginalCellParameters
	isSet bool
}

func (v NullableInterRatEsActivationOriginalCellParameters) Get() *InterRatEsActivationOriginalCellParameters {
	return v.value
}

func (v *NullableInterRatEsActivationOriginalCellParameters) Set(val *InterRatEsActivationOriginalCellParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableInterRatEsActivationOriginalCellParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableInterRatEsActivationOriginalCellParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterRatEsActivationOriginalCellParameters(val *InterRatEsActivationOriginalCellParameters) *NullableInterRatEsActivationOriginalCellParameters {
	return &NullableInterRatEsActivationOriginalCellParameters{value: val, isSet: true}
}

func (v NullableInterRatEsActivationOriginalCellParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterRatEsActivationOriginalCellParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the IntraRatEsActivationCandidateCellsLoadParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntraRatEsActivationCandidateCellsLoadParameters{}

// IntraRatEsActivationCandidateCellsLoadParameters struct for IntraRatEsActivationCandidateCellsLoadParameters
type IntraRatEsActivationCandidateCellsLoadParameters struct {
	LoadThreshold *int32 `json:"loadThreshold,omitempty"`
	TimeDuration  *int32 `json:"timeDuration,omitempty"`
}

// NewIntraRatEsActivationCandidateCellsLoadParameters instantiates a new IntraRatEsActivationCandidateCellsLoadParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntraRatEsActivationCandidateCellsLoadParameters() *IntraRatEsActivationCandidateCellsLoadParameters {
	this := IntraRatEsActivationCandidateCellsLoadParameters{}
	return &this
}

// NewIntraRatEsActivationCandidateCellsLoadParametersWithDefaults instantiates a new IntraRatEsActivationCandidateCellsLoadParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntraRatEsActivationCandidateCellsLoadParametersWithDefaults() *IntraRatEsActivationCandidateCellsLoadParameters {
	this := IntraRatEsActivationCandidateCellsLoadParameters{}
	return &this
}

// GetLoadThreshold returns the LoadThreshold field value if set, zero value otherwise.
func (o *IntraRatEsActivationCandidateCellsLoadParameters) GetLoadThreshold() int32 {
	if o == nil || IsNil(o.LoadThreshold) {
		var ret int32
		return ret
	}
	return *o.LoadThreshold
}

// GetLoadThresholdOk returns a tuple with the LoadThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntraRatEsActivationCandidateCellsLoadParameters) GetLoadThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.LoadThreshold) {
		return nil, false
	}
	return o.LoadThreshold, true
}

// HasLoadThreshold returns a boolean if a field has been set.
func (o *IntraRatEsActivationCandidateCellsLoadParameters) HasLoadThreshold() bool {
	if o != nil && !IsNil(o.LoadThreshold) {
		return true
	}

	return false
}

// SetLoadThreshold gets a reference to the given int32 and assigns it to the LoadThreshold field.
func (o *IntraRatEsActivationCandidateCellsLoadParameters) SetLoadThreshold(v int32) {
	o.LoadThreshold = &v
}

// GetTimeDuration returns the TimeDuration field value if set, zero value otherwise.
func (o *IntraRatEsActivationCandidateCellsLoadParameters) GetTimeDuration() int32 {
	if o == nil || IsNil(o.TimeDuration) {
		var ret int32
		return ret
	}
	return *o.TimeDuration
}

// GetTimeDurationOk returns a tuple with the TimeDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntraRatEsActivationCandidateCellsLoadParameters) GetTimeDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeDuration) {
		return nil, false
	}
	return o.TimeDuration, true
}

// HasTimeDuration returns a boolean if a field has been set.
func (o *IntraRatEsActivationCandidateCellsLoadParameters) HasTimeDuration() bool {
	if o != nil && !IsNil(o.TimeDuration) {
		return true
	}

	return false
}

// SetTimeDuration gets a reference to the given int32 and assigns it to the TimeDuration field.
func (o *IntraRatEsActivationCandidateCellsLoadParameters) SetTimeDuration(v int32) {
	o.TimeDuration = &v
}

func (o IntraRatEsActivationCandidateCellsLoadParameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntraRatEsActivationCandidateCellsLoadParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LoadThreshold) {
		toSerialize["loadThreshold"] = o.LoadThreshold
	}
	if !IsNil(o.TimeDuration) {
		toSerialize["timeDuration"] = o.TimeDuration
	}
	return toSerialize, nil
}

type NullableIntraRatEsActivationCandidateCellsLoadParameters struct {
	value *IntraRatEsActivationCandidateCellsLoadParameters
	isSet bool
}

func (v NullableIntraRatEsActivationCandidateCellsLoadParameters) Get() *IntraRatEsActivationCandidateCellsLoadParameters {
	return v.value
}

func (v *NullableIntraRatEsActivationCandidateCellsLoadParameters) Set(val *IntraRatEsActivationCandidateCellsLoadParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableIntraRatEsActivationCandidateCellsLoadParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableIntraRatEsActivationCandidateCellsLoadParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntraRatEsActivationCandidateCellsLoadParameters(val *IntraRatEsActivationCandidateCellsLoadParameters) *NullableIntraRatEsActivationCandidateCellsLoadParameters {
	return &NullableIntraRatEsActivationCandidateCellsLoadParameters{value: val, isSet: true}
}

func (v NullableIntraRatEsActivationCandidateCellsLoadParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntraRatEsActivationCandidateCellsLoadParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

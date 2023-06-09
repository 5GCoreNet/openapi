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

// checks if the ModelPerformance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPerformance{}

// ModelPerformance struct for ModelPerformance
type ModelPerformance struct {
	InferenceOutputName     *string  `json:"inferenceOutputName,omitempty"`
	PerformanceMetric       *string  `json:"performanceMetric,omitempty"`
	PerformanceScore        *float32 `json:"performanceScore,omitempty"`
	DecisionConfidenceScore *float32 `json:"decisionConfidenceScore,omitempty"`
}

// NewModelPerformance instantiates a new ModelPerformance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPerformance() *ModelPerformance {
	this := ModelPerformance{}
	return &this
}

// NewModelPerformanceWithDefaults instantiates a new ModelPerformance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPerformanceWithDefaults() *ModelPerformance {
	this := ModelPerformance{}
	return &this
}

// GetInferenceOutputName returns the InferenceOutputName field value if set, zero value otherwise.
func (o *ModelPerformance) GetInferenceOutputName() string {
	if o == nil || IsNil(o.InferenceOutputName) {
		var ret string
		return ret
	}
	return *o.InferenceOutputName
}

// GetInferenceOutputNameOk returns a tuple with the InferenceOutputName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPerformance) GetInferenceOutputNameOk() (*string, bool) {
	if o == nil || IsNil(o.InferenceOutputName) {
		return nil, false
	}
	return o.InferenceOutputName, true
}

// HasInferenceOutputName returns a boolean if a field has been set.
func (o *ModelPerformance) HasInferenceOutputName() bool {
	if o != nil && !IsNil(o.InferenceOutputName) {
		return true
	}

	return false
}

// SetInferenceOutputName gets a reference to the given string and assigns it to the InferenceOutputName field.
func (o *ModelPerformance) SetInferenceOutputName(v string) {
	o.InferenceOutputName = &v
}

// GetPerformanceMetric returns the PerformanceMetric field value if set, zero value otherwise.
func (o *ModelPerformance) GetPerformanceMetric() string {
	if o == nil || IsNil(o.PerformanceMetric) {
		var ret string
		return ret
	}
	return *o.PerformanceMetric
}

// GetPerformanceMetricOk returns a tuple with the PerformanceMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPerformance) GetPerformanceMetricOk() (*string, bool) {
	if o == nil || IsNil(o.PerformanceMetric) {
		return nil, false
	}
	return o.PerformanceMetric, true
}

// HasPerformanceMetric returns a boolean if a field has been set.
func (o *ModelPerformance) HasPerformanceMetric() bool {
	if o != nil && !IsNil(o.PerformanceMetric) {
		return true
	}

	return false
}

// SetPerformanceMetric gets a reference to the given string and assigns it to the PerformanceMetric field.
func (o *ModelPerformance) SetPerformanceMetric(v string) {
	o.PerformanceMetric = &v
}

// GetPerformanceScore returns the PerformanceScore field value if set, zero value otherwise.
func (o *ModelPerformance) GetPerformanceScore() float32 {
	if o == nil || IsNil(o.PerformanceScore) {
		var ret float32
		return ret
	}
	return *o.PerformanceScore
}

// GetPerformanceScoreOk returns a tuple with the PerformanceScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPerformance) GetPerformanceScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.PerformanceScore) {
		return nil, false
	}
	return o.PerformanceScore, true
}

// HasPerformanceScore returns a boolean if a field has been set.
func (o *ModelPerformance) HasPerformanceScore() bool {
	if o != nil && !IsNil(o.PerformanceScore) {
		return true
	}

	return false
}

// SetPerformanceScore gets a reference to the given float32 and assigns it to the PerformanceScore field.
func (o *ModelPerformance) SetPerformanceScore(v float32) {
	o.PerformanceScore = &v
}

// GetDecisionConfidenceScore returns the DecisionConfidenceScore field value if set, zero value otherwise.
func (o *ModelPerformance) GetDecisionConfidenceScore() float32 {
	if o == nil || IsNil(o.DecisionConfidenceScore) {
		var ret float32
		return ret
	}
	return *o.DecisionConfidenceScore
}

// GetDecisionConfidenceScoreOk returns a tuple with the DecisionConfidenceScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPerformance) GetDecisionConfidenceScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.DecisionConfidenceScore) {
		return nil, false
	}
	return o.DecisionConfidenceScore, true
}

// HasDecisionConfidenceScore returns a boolean if a field has been set.
func (o *ModelPerformance) HasDecisionConfidenceScore() bool {
	if o != nil && !IsNil(o.DecisionConfidenceScore) {
		return true
	}

	return false
}

// SetDecisionConfidenceScore gets a reference to the given float32 and assigns it to the DecisionConfidenceScore field.
func (o *ModelPerformance) SetDecisionConfidenceScore(v float32) {
	o.DecisionConfidenceScore = &v
}

func (o ModelPerformance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPerformance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InferenceOutputName) {
		toSerialize["inferenceOutputName"] = o.InferenceOutputName
	}
	if !IsNil(o.PerformanceMetric) {
		toSerialize["performanceMetric"] = o.PerformanceMetric
	}
	if !IsNil(o.PerformanceScore) {
		toSerialize["performanceScore"] = o.PerformanceScore
	}
	if !IsNil(o.DecisionConfidenceScore) {
		toSerialize["decisionConfidenceScore"] = o.DecisionConfidenceScore
	}
	return toSerialize, nil
}

type NullableModelPerformance struct {
	value *ModelPerformance
	isSet bool
}

func (v NullableModelPerformance) Get() *ModelPerformance {
	return v.value
}

func (v *NullableModelPerformance) Set(val *ModelPerformance) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPerformance) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPerformance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPerformance(val *ModelPerformance) *NullableModelPerformance {
	return &NullableModelPerformance{value: val, isSet: true}
}

func (v NullableModelPerformance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPerformance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

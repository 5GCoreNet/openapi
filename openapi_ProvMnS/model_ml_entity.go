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

// checks if the MLEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MLEntity{}

// MLEntity struct for MLEntity
type MLEntity struct {
	MLEntityId *string `json:"mLEntityId,omitempty"`
	InferenceType *string `json:"inferenceType,omitempty"`
	MLEntityVersion *string `json:"mLEntityVersion,omitempty"`
	ExpectedRunTimeContext *MLContext `json:"expectedRunTimeContext,omitempty"`
	TrainingContext *MLContext `json:"trainingContext,omitempty"`
	RunTimeContext *MLContext `json:"runTimeContext,omitempty"`
}

// NewMLEntity instantiates a new MLEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMLEntity() *MLEntity {
	this := MLEntity{}
	return &this
}

// NewMLEntityWithDefaults instantiates a new MLEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMLEntityWithDefaults() *MLEntity {
	this := MLEntity{}
	return &this
}

// GetMLEntityId returns the MLEntityId field value if set, zero value otherwise.
func (o *MLEntity) GetMLEntityId() string {
	if o == nil || isNil(o.MLEntityId) {
		var ret string
		return ret
	}
	return *o.MLEntityId
}

// GetMLEntityIdOk returns a tuple with the MLEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLEntity) GetMLEntityIdOk() (*string, bool) {
	if o == nil || isNil(o.MLEntityId) {
		return nil, false
	}
	return o.MLEntityId, true
}

// HasMLEntityId returns a boolean if a field has been set.
func (o *MLEntity) HasMLEntityId() bool {
	if o != nil && !isNil(o.MLEntityId) {
		return true
	}

	return false
}

// SetMLEntityId gets a reference to the given string and assigns it to the MLEntityId field.
func (o *MLEntity) SetMLEntityId(v string) {
	o.MLEntityId = &v
}

// GetInferenceType returns the InferenceType field value if set, zero value otherwise.
func (o *MLEntity) GetInferenceType() string {
	if o == nil || isNil(o.InferenceType) {
		var ret string
		return ret
	}
	return *o.InferenceType
}

// GetInferenceTypeOk returns a tuple with the InferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLEntity) GetInferenceTypeOk() (*string, bool) {
	if o == nil || isNil(o.InferenceType) {
		return nil, false
	}
	return o.InferenceType, true
}

// HasInferenceType returns a boolean if a field has been set.
func (o *MLEntity) HasInferenceType() bool {
	if o != nil && !isNil(o.InferenceType) {
		return true
	}

	return false
}

// SetInferenceType gets a reference to the given string and assigns it to the InferenceType field.
func (o *MLEntity) SetInferenceType(v string) {
	o.InferenceType = &v
}

// GetMLEntityVersion returns the MLEntityVersion field value if set, zero value otherwise.
func (o *MLEntity) GetMLEntityVersion() string {
	if o == nil || isNil(o.MLEntityVersion) {
		var ret string
		return ret
	}
	return *o.MLEntityVersion
}

// GetMLEntityVersionOk returns a tuple with the MLEntityVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLEntity) GetMLEntityVersionOk() (*string, bool) {
	if o == nil || isNil(o.MLEntityVersion) {
		return nil, false
	}
	return o.MLEntityVersion, true
}

// HasMLEntityVersion returns a boolean if a field has been set.
func (o *MLEntity) HasMLEntityVersion() bool {
	if o != nil && !isNil(o.MLEntityVersion) {
		return true
	}

	return false
}

// SetMLEntityVersion gets a reference to the given string and assigns it to the MLEntityVersion field.
func (o *MLEntity) SetMLEntityVersion(v string) {
	o.MLEntityVersion = &v
}

// GetExpectedRunTimeContext returns the ExpectedRunTimeContext field value if set, zero value otherwise.
func (o *MLEntity) GetExpectedRunTimeContext() MLContext {
	if o == nil || isNil(o.ExpectedRunTimeContext) {
		var ret MLContext
		return ret
	}
	return *o.ExpectedRunTimeContext
}

// GetExpectedRunTimeContextOk returns a tuple with the ExpectedRunTimeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLEntity) GetExpectedRunTimeContextOk() (*MLContext, bool) {
	if o == nil || isNil(o.ExpectedRunTimeContext) {
		return nil, false
	}
	return o.ExpectedRunTimeContext, true
}

// HasExpectedRunTimeContext returns a boolean if a field has been set.
func (o *MLEntity) HasExpectedRunTimeContext() bool {
	if o != nil && !isNil(o.ExpectedRunTimeContext) {
		return true
	}

	return false
}

// SetExpectedRunTimeContext gets a reference to the given MLContext and assigns it to the ExpectedRunTimeContext field.
func (o *MLEntity) SetExpectedRunTimeContext(v MLContext) {
	o.ExpectedRunTimeContext = &v
}

// GetTrainingContext returns the TrainingContext field value if set, zero value otherwise.
func (o *MLEntity) GetTrainingContext() MLContext {
	if o == nil || isNil(o.TrainingContext) {
		var ret MLContext
		return ret
	}
	return *o.TrainingContext
}

// GetTrainingContextOk returns a tuple with the TrainingContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLEntity) GetTrainingContextOk() (*MLContext, bool) {
	if o == nil || isNil(o.TrainingContext) {
		return nil, false
	}
	return o.TrainingContext, true
}

// HasTrainingContext returns a boolean if a field has been set.
func (o *MLEntity) HasTrainingContext() bool {
	if o != nil && !isNil(o.TrainingContext) {
		return true
	}

	return false
}

// SetTrainingContext gets a reference to the given MLContext and assigns it to the TrainingContext field.
func (o *MLEntity) SetTrainingContext(v MLContext) {
	o.TrainingContext = &v
}

// GetRunTimeContext returns the RunTimeContext field value if set, zero value otherwise.
func (o *MLEntity) GetRunTimeContext() MLContext {
	if o == nil || isNil(o.RunTimeContext) {
		var ret MLContext
		return ret
	}
	return *o.RunTimeContext
}

// GetRunTimeContextOk returns a tuple with the RunTimeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLEntity) GetRunTimeContextOk() (*MLContext, bool) {
	if o == nil || isNil(o.RunTimeContext) {
		return nil, false
	}
	return o.RunTimeContext, true
}

// HasRunTimeContext returns a boolean if a field has been set.
func (o *MLEntity) HasRunTimeContext() bool {
	if o != nil && !isNil(o.RunTimeContext) {
		return true
	}

	return false
}

// SetRunTimeContext gets a reference to the given MLContext and assigns it to the RunTimeContext field.
func (o *MLEntity) SetRunTimeContext(v MLContext) {
	o.RunTimeContext = &v
}

func (o MLEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MLEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MLEntityId) {
		toSerialize["mLEntityId"] = o.MLEntityId
	}
	if !isNil(o.InferenceType) {
		toSerialize["inferenceType"] = o.InferenceType
	}
	if !isNil(o.MLEntityVersion) {
		toSerialize["mLEntityVersion"] = o.MLEntityVersion
	}
	if !isNil(o.ExpectedRunTimeContext) {
		toSerialize["expectedRunTimeContext"] = o.ExpectedRunTimeContext
	}
	if !isNil(o.TrainingContext) {
		toSerialize["trainingContext"] = o.TrainingContext
	}
	if !isNil(o.RunTimeContext) {
		toSerialize["runTimeContext"] = o.RunTimeContext
	}
	return toSerialize, nil
}

type NullableMLEntity struct {
	value *MLEntity
	isSet bool
}

func (v NullableMLEntity) Get() *MLEntity {
	return v.value
}

func (v *NullableMLEntity) Set(val *MLEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableMLEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableMLEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMLEntity(val *MLEntity) *NullableMLEntity {
	return &NullableMLEntity{value: val, isSet: true}
}

func (v NullableMLEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMLEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
AI/ML NRM

OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AiMlNrm

import (
	"encoding/json"
	"time"
)

// checks if the MLTrainingRequestSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MLTrainingRequestSingleAllOfAttributesAllOf{}

// MLTrainingRequestSingleAllOfAttributesAllOf struct for MLTrainingRequestSingleAllOfAttributesAllOf
type MLTrainingRequestSingleAllOfAttributesAllOf struct {
	MLEntityId *string `json:"mLEntityId,omitempty"`
	CandidateTraingDataSource []string `json:"candidateTraingDataSource,omitempty"`
	TraingDataQualityScore *float32 `json:"traingDataQualityScore,omitempty"`
	TrainingRequestSource *string `json:"trainingRequestSource,omitempty"`
	RequestStatus *RequestStatus `json:"requestStatus,omitempty"`
	ExpectedRuntimeContext *time.Time `json:"expectedRuntimeContext,omitempty"`
	PerformanceRequirements []ModelPerformance `json:"performanceRequirements,omitempty"`
	CancelRequest *bool `json:"cancelRequest,omitempty"`
	SuspendRequest *bool `json:"suspendRequest,omitempty"`
}

// NewMLTrainingRequestSingleAllOfAttributesAllOf instantiates a new MLTrainingRequestSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMLTrainingRequestSingleAllOfAttributesAllOf() *MLTrainingRequestSingleAllOfAttributesAllOf {
	this := MLTrainingRequestSingleAllOfAttributesAllOf{}
	return &this
}

// NewMLTrainingRequestSingleAllOfAttributesAllOfWithDefaults instantiates a new MLTrainingRequestSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMLTrainingRequestSingleAllOfAttributesAllOfWithDefaults() *MLTrainingRequestSingleAllOfAttributesAllOf {
	this := MLTrainingRequestSingleAllOfAttributesAllOf{}
	return &this
}

// GetMLEntityId returns the MLEntityId field value if set, zero value otherwise.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetMLEntityId() string {
	if o == nil || IsNil(o.MLEntityId) {
		var ret string
		return ret
	}
	return *o.MLEntityId
}

// GetMLEntityIdOk returns a tuple with the MLEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetMLEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.MLEntityId) {
		return nil, false
	}
	return o.MLEntityId, true
}

// HasMLEntityId returns a boolean if a field has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) HasMLEntityId() bool {
	if o != nil && !IsNil(o.MLEntityId) {
		return true
	}

	return false
}

// SetMLEntityId gets a reference to the given string and assigns it to the MLEntityId field.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) SetMLEntityId(v string) {
	o.MLEntityId = &v
}

// GetCandidateTraingDataSource returns the CandidateTraingDataSource field value if set, zero value otherwise.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetCandidateTraingDataSource() []string {
	if o == nil || IsNil(o.CandidateTraingDataSource) {
		var ret []string
		return ret
	}
	return o.CandidateTraingDataSource
}

// GetCandidateTraingDataSourceOk returns a tuple with the CandidateTraingDataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetCandidateTraingDataSourceOk() ([]string, bool) {
	if o == nil || IsNil(o.CandidateTraingDataSource) {
		return nil, false
	}
	return o.CandidateTraingDataSource, true
}

// HasCandidateTraingDataSource returns a boolean if a field has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) HasCandidateTraingDataSource() bool {
	if o != nil && !IsNil(o.CandidateTraingDataSource) {
		return true
	}

	return false
}

// SetCandidateTraingDataSource gets a reference to the given []string and assigns it to the CandidateTraingDataSource field.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) SetCandidateTraingDataSource(v []string) {
	o.CandidateTraingDataSource = v
}

// GetTraingDataQualityScore returns the TraingDataQualityScore field value if set, zero value otherwise.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetTraingDataQualityScore() float32 {
	if o == nil || IsNil(o.TraingDataQualityScore) {
		var ret float32
		return ret
	}
	return *o.TraingDataQualityScore
}

// GetTraingDataQualityScoreOk returns a tuple with the TraingDataQualityScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetTraingDataQualityScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.TraingDataQualityScore) {
		return nil, false
	}
	return o.TraingDataQualityScore, true
}

// HasTraingDataQualityScore returns a boolean if a field has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) HasTraingDataQualityScore() bool {
	if o != nil && !IsNil(o.TraingDataQualityScore) {
		return true
	}

	return false
}

// SetTraingDataQualityScore gets a reference to the given float32 and assigns it to the TraingDataQualityScore field.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) SetTraingDataQualityScore(v float32) {
	o.TraingDataQualityScore = &v
}

// GetTrainingRequestSource returns the TrainingRequestSource field value if set, zero value otherwise.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetTrainingRequestSource() string {
	if o == nil || IsNil(o.TrainingRequestSource) {
		var ret string
		return ret
	}
	return *o.TrainingRequestSource
}

// GetTrainingRequestSourceOk returns a tuple with the TrainingRequestSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetTrainingRequestSourceOk() (*string, bool) {
	if o == nil || IsNil(o.TrainingRequestSource) {
		return nil, false
	}
	return o.TrainingRequestSource, true
}

// HasTrainingRequestSource returns a boolean if a field has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) HasTrainingRequestSource() bool {
	if o != nil && !IsNil(o.TrainingRequestSource) {
		return true
	}

	return false
}

// SetTrainingRequestSource gets a reference to the given string and assigns it to the TrainingRequestSource field.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) SetTrainingRequestSource(v string) {
	o.TrainingRequestSource = &v
}

// GetRequestStatus returns the RequestStatus field value if set, zero value otherwise.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetRequestStatus() RequestStatus {
	if o == nil || IsNil(o.RequestStatus) {
		var ret RequestStatus
		return ret
	}
	return *o.RequestStatus
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetRequestStatusOk() (*RequestStatus, bool) {
	if o == nil || IsNil(o.RequestStatus) {
		return nil, false
	}
	return o.RequestStatus, true
}

// HasRequestStatus returns a boolean if a field has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) HasRequestStatus() bool {
	if o != nil && !IsNil(o.RequestStatus) {
		return true
	}

	return false
}

// SetRequestStatus gets a reference to the given RequestStatus and assigns it to the RequestStatus field.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) SetRequestStatus(v RequestStatus) {
	o.RequestStatus = &v
}

// GetExpectedRuntimeContext returns the ExpectedRuntimeContext field value if set, zero value otherwise.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetExpectedRuntimeContext() time.Time {
	if o == nil || IsNil(o.ExpectedRuntimeContext) {
		var ret time.Time
		return ret
	}
	return *o.ExpectedRuntimeContext
}

// GetExpectedRuntimeContextOk returns a tuple with the ExpectedRuntimeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetExpectedRuntimeContextOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpectedRuntimeContext) {
		return nil, false
	}
	return o.ExpectedRuntimeContext, true
}

// HasExpectedRuntimeContext returns a boolean if a field has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) HasExpectedRuntimeContext() bool {
	if o != nil && !IsNil(o.ExpectedRuntimeContext) {
		return true
	}

	return false
}

// SetExpectedRuntimeContext gets a reference to the given time.Time and assigns it to the ExpectedRuntimeContext field.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) SetExpectedRuntimeContext(v time.Time) {
	o.ExpectedRuntimeContext = &v
}

// GetPerformanceRequirements returns the PerformanceRequirements field value if set, zero value otherwise.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetPerformanceRequirements() []ModelPerformance {
	if o == nil || IsNil(o.PerformanceRequirements) {
		var ret []ModelPerformance
		return ret
	}
	return o.PerformanceRequirements
}

// GetPerformanceRequirementsOk returns a tuple with the PerformanceRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetPerformanceRequirementsOk() ([]ModelPerformance, bool) {
	if o == nil || IsNil(o.PerformanceRequirements) {
		return nil, false
	}
	return o.PerformanceRequirements, true
}

// HasPerformanceRequirements returns a boolean if a field has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) HasPerformanceRequirements() bool {
	if o != nil && !IsNil(o.PerformanceRequirements) {
		return true
	}

	return false
}

// SetPerformanceRequirements gets a reference to the given []ModelPerformance and assigns it to the PerformanceRequirements field.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) SetPerformanceRequirements(v []ModelPerformance) {
	o.PerformanceRequirements = v
}

// GetCancelRequest returns the CancelRequest field value if set, zero value otherwise.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetCancelRequest() bool {
	if o == nil || IsNil(o.CancelRequest) {
		var ret bool
		return ret
	}
	return *o.CancelRequest
}

// GetCancelRequestOk returns a tuple with the CancelRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetCancelRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelRequest) {
		return nil, false
	}
	return o.CancelRequest, true
}

// HasCancelRequest returns a boolean if a field has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) HasCancelRequest() bool {
	if o != nil && !IsNil(o.CancelRequest) {
		return true
	}

	return false
}

// SetCancelRequest gets a reference to the given bool and assigns it to the CancelRequest field.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) SetCancelRequest(v bool) {
	o.CancelRequest = &v
}

// GetSuspendRequest returns the SuspendRequest field value if set, zero value otherwise.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetSuspendRequest() bool {
	if o == nil || IsNil(o.SuspendRequest) {
		var ret bool
		return ret
	}
	return *o.SuspendRequest
}

// GetSuspendRequestOk returns a tuple with the SuspendRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) GetSuspendRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.SuspendRequest) {
		return nil, false
	}
	return o.SuspendRequest, true
}

// HasSuspendRequest returns a boolean if a field has been set.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) HasSuspendRequest() bool {
	if o != nil && !IsNil(o.SuspendRequest) {
		return true
	}

	return false
}

// SetSuspendRequest gets a reference to the given bool and assigns it to the SuspendRequest field.
func (o *MLTrainingRequestSingleAllOfAttributesAllOf) SetSuspendRequest(v bool) {
	o.SuspendRequest = &v
}

func (o MLTrainingRequestSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MLTrainingRequestSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MLEntityId) {
		toSerialize["mLEntityId"] = o.MLEntityId
	}
	if !IsNil(o.CandidateTraingDataSource) {
		toSerialize["candidateTraingDataSource"] = o.CandidateTraingDataSource
	}
	if !IsNil(o.TraingDataQualityScore) {
		toSerialize["traingDataQualityScore"] = o.TraingDataQualityScore
	}
	if !IsNil(o.TrainingRequestSource) {
		toSerialize["trainingRequestSource"] = o.TrainingRequestSource
	}
	if !IsNil(o.RequestStatus) {
		toSerialize["requestStatus"] = o.RequestStatus
	}
	if !IsNil(o.ExpectedRuntimeContext) {
		toSerialize["expectedRuntimeContext"] = o.ExpectedRuntimeContext
	}
	if !IsNil(o.PerformanceRequirements) {
		toSerialize["performanceRequirements"] = o.PerformanceRequirements
	}
	if !IsNil(o.CancelRequest) {
		toSerialize["cancelRequest"] = o.CancelRequest
	}
	if !IsNil(o.SuspendRequest) {
		toSerialize["suspendRequest"] = o.SuspendRequest
	}
	return toSerialize, nil
}

type NullableMLTrainingRequestSingleAllOfAttributesAllOf struct {
	value *MLTrainingRequestSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableMLTrainingRequestSingleAllOfAttributesAllOf) Get() *MLTrainingRequestSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableMLTrainingRequestSingleAllOfAttributesAllOf) Set(val *MLTrainingRequestSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMLTrainingRequestSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMLTrainingRequestSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMLTrainingRequestSingleAllOfAttributesAllOf(val *MLTrainingRequestSingleAllOfAttributesAllOf) *NullableMLTrainingRequestSingleAllOfAttributesAllOf {
	return &NullableMLTrainingRequestSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableMLTrainingRequestSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMLTrainingRequestSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the MLTrainingReportSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MLTrainingReportSingleAllOfAttributes{}

// MLTrainingReportSingleAllOfAttributes struct for MLTrainingReportSingleAllOfAttributes
type MLTrainingReportSingleAllOfAttributes struct {
	MLEntityId *string `json:"mLEntityId,omitempty"`
	AreConsumerTrainingDataUsed *bool `json:"areConsumerTrainingDataUsed,omitempty"`
	UsedConsumerTrainingData []string `json:"usedConsumerTrainingData,omitempty"`
	ConfidenceIndication *int32 `json:"confidenceIndication,omitempty"`
	ModelPerformanceTraining []ModelPerformance `json:"modelPerformanceTraining,omitempty"`
	AreNewTrainingDataUsed *bool `json:"areNewTrainingDataUsed,omitempty"`
}

// NewMLTrainingReportSingleAllOfAttributes instantiates a new MLTrainingReportSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMLTrainingReportSingleAllOfAttributes() *MLTrainingReportSingleAllOfAttributes {
	this := MLTrainingReportSingleAllOfAttributes{}
	return &this
}

// NewMLTrainingReportSingleAllOfAttributesWithDefaults instantiates a new MLTrainingReportSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMLTrainingReportSingleAllOfAttributesWithDefaults() *MLTrainingReportSingleAllOfAttributes {
	this := MLTrainingReportSingleAllOfAttributes{}
	return &this
}

// GetMLEntityId returns the MLEntityId field value if set, zero value otherwise.
func (o *MLTrainingReportSingleAllOfAttributes) GetMLEntityId() string {
	if o == nil || IsNil(o.MLEntityId) {
		var ret string
		return ret
	}
	return *o.MLEntityId
}

// GetMLEntityIdOk returns a tuple with the MLEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingReportSingleAllOfAttributes) GetMLEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.MLEntityId) {
		return nil, false
	}
	return o.MLEntityId, true
}

// HasMLEntityId returns a boolean if a field has been set.
func (o *MLTrainingReportSingleAllOfAttributes) HasMLEntityId() bool {
	if o != nil && !IsNil(o.MLEntityId) {
		return true
	}

	return false
}

// SetMLEntityId gets a reference to the given string and assigns it to the MLEntityId field.
func (o *MLTrainingReportSingleAllOfAttributes) SetMLEntityId(v string) {
	o.MLEntityId = &v
}

// GetAreConsumerTrainingDataUsed returns the AreConsumerTrainingDataUsed field value if set, zero value otherwise.
func (o *MLTrainingReportSingleAllOfAttributes) GetAreConsumerTrainingDataUsed() bool {
	if o == nil || IsNil(o.AreConsumerTrainingDataUsed) {
		var ret bool
		return ret
	}
	return *o.AreConsumerTrainingDataUsed
}

// GetAreConsumerTrainingDataUsedOk returns a tuple with the AreConsumerTrainingDataUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingReportSingleAllOfAttributes) GetAreConsumerTrainingDataUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.AreConsumerTrainingDataUsed) {
		return nil, false
	}
	return o.AreConsumerTrainingDataUsed, true
}

// HasAreConsumerTrainingDataUsed returns a boolean if a field has been set.
func (o *MLTrainingReportSingleAllOfAttributes) HasAreConsumerTrainingDataUsed() bool {
	if o != nil && !IsNil(o.AreConsumerTrainingDataUsed) {
		return true
	}

	return false
}

// SetAreConsumerTrainingDataUsed gets a reference to the given bool and assigns it to the AreConsumerTrainingDataUsed field.
func (o *MLTrainingReportSingleAllOfAttributes) SetAreConsumerTrainingDataUsed(v bool) {
	o.AreConsumerTrainingDataUsed = &v
}

// GetUsedConsumerTrainingData returns the UsedConsumerTrainingData field value if set, zero value otherwise.
func (o *MLTrainingReportSingleAllOfAttributes) GetUsedConsumerTrainingData() []string {
	if o == nil || IsNil(o.UsedConsumerTrainingData) {
		var ret []string
		return ret
	}
	return o.UsedConsumerTrainingData
}

// GetUsedConsumerTrainingDataOk returns a tuple with the UsedConsumerTrainingData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingReportSingleAllOfAttributes) GetUsedConsumerTrainingDataOk() ([]string, bool) {
	if o == nil || IsNil(o.UsedConsumerTrainingData) {
		return nil, false
	}
	return o.UsedConsumerTrainingData, true
}

// HasUsedConsumerTrainingData returns a boolean if a field has been set.
func (o *MLTrainingReportSingleAllOfAttributes) HasUsedConsumerTrainingData() bool {
	if o != nil && !IsNil(o.UsedConsumerTrainingData) {
		return true
	}

	return false
}

// SetUsedConsumerTrainingData gets a reference to the given []string and assigns it to the UsedConsumerTrainingData field.
func (o *MLTrainingReportSingleAllOfAttributes) SetUsedConsumerTrainingData(v []string) {
	o.UsedConsumerTrainingData = v
}

// GetConfidenceIndication returns the ConfidenceIndication field value if set, zero value otherwise.
func (o *MLTrainingReportSingleAllOfAttributes) GetConfidenceIndication() int32 {
	if o == nil || IsNil(o.ConfidenceIndication) {
		var ret int32
		return ret
	}
	return *o.ConfidenceIndication
}

// GetConfidenceIndicationOk returns a tuple with the ConfidenceIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingReportSingleAllOfAttributes) GetConfidenceIndicationOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfidenceIndication) {
		return nil, false
	}
	return o.ConfidenceIndication, true
}

// HasConfidenceIndication returns a boolean if a field has been set.
func (o *MLTrainingReportSingleAllOfAttributes) HasConfidenceIndication() bool {
	if o != nil && !IsNil(o.ConfidenceIndication) {
		return true
	}

	return false
}

// SetConfidenceIndication gets a reference to the given int32 and assigns it to the ConfidenceIndication field.
func (o *MLTrainingReportSingleAllOfAttributes) SetConfidenceIndication(v int32) {
	o.ConfidenceIndication = &v
}

// GetModelPerformanceTraining returns the ModelPerformanceTraining field value if set, zero value otherwise.
func (o *MLTrainingReportSingleAllOfAttributes) GetModelPerformanceTraining() []ModelPerformance {
	if o == nil || IsNil(o.ModelPerformanceTraining) {
		var ret []ModelPerformance
		return ret
	}
	return o.ModelPerformanceTraining
}

// GetModelPerformanceTrainingOk returns a tuple with the ModelPerformanceTraining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingReportSingleAllOfAttributes) GetModelPerformanceTrainingOk() ([]ModelPerformance, bool) {
	if o == nil || IsNil(o.ModelPerformanceTraining) {
		return nil, false
	}
	return o.ModelPerformanceTraining, true
}

// HasModelPerformanceTraining returns a boolean if a field has been set.
func (o *MLTrainingReportSingleAllOfAttributes) HasModelPerformanceTraining() bool {
	if o != nil && !IsNil(o.ModelPerformanceTraining) {
		return true
	}

	return false
}

// SetModelPerformanceTraining gets a reference to the given []ModelPerformance and assigns it to the ModelPerformanceTraining field.
func (o *MLTrainingReportSingleAllOfAttributes) SetModelPerformanceTraining(v []ModelPerformance) {
	o.ModelPerformanceTraining = v
}

// GetAreNewTrainingDataUsed returns the AreNewTrainingDataUsed field value if set, zero value otherwise.
func (o *MLTrainingReportSingleAllOfAttributes) GetAreNewTrainingDataUsed() bool {
	if o == nil || IsNil(o.AreNewTrainingDataUsed) {
		var ret bool
		return ret
	}
	return *o.AreNewTrainingDataUsed
}

// GetAreNewTrainingDataUsedOk returns a tuple with the AreNewTrainingDataUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingReportSingleAllOfAttributes) GetAreNewTrainingDataUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.AreNewTrainingDataUsed) {
		return nil, false
	}
	return o.AreNewTrainingDataUsed, true
}

// HasAreNewTrainingDataUsed returns a boolean if a field has been set.
func (o *MLTrainingReportSingleAllOfAttributes) HasAreNewTrainingDataUsed() bool {
	if o != nil && !IsNil(o.AreNewTrainingDataUsed) {
		return true
	}

	return false
}

// SetAreNewTrainingDataUsed gets a reference to the given bool and assigns it to the AreNewTrainingDataUsed field.
func (o *MLTrainingReportSingleAllOfAttributes) SetAreNewTrainingDataUsed(v bool) {
	o.AreNewTrainingDataUsed = &v
}

func (o MLTrainingReportSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MLTrainingReportSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MLEntityId) {
		toSerialize["mLEntityId"] = o.MLEntityId
	}
	if !IsNil(o.AreConsumerTrainingDataUsed) {
		toSerialize["areConsumerTrainingDataUsed"] = o.AreConsumerTrainingDataUsed
	}
	if !IsNil(o.UsedConsumerTrainingData) {
		toSerialize["usedConsumerTrainingData"] = o.UsedConsumerTrainingData
	}
	if !IsNil(o.ConfidenceIndication) {
		toSerialize["confidenceIndication"] = o.ConfidenceIndication
	}
	if !IsNil(o.ModelPerformanceTraining) {
		toSerialize["modelPerformanceTraining"] = o.ModelPerformanceTraining
	}
	if !IsNil(o.AreNewTrainingDataUsed) {
		toSerialize["areNewTrainingDataUsed"] = o.AreNewTrainingDataUsed
	}
	return toSerialize, nil
}

type NullableMLTrainingReportSingleAllOfAttributes struct {
	value *MLTrainingReportSingleAllOfAttributes
	isSet bool
}

func (v NullableMLTrainingReportSingleAllOfAttributes) Get() *MLTrainingReportSingleAllOfAttributes {
	return v.value
}

func (v *NullableMLTrainingReportSingleAllOfAttributes) Set(val *MLTrainingReportSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMLTrainingReportSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMLTrainingReportSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMLTrainingReportSingleAllOfAttributes(val *MLTrainingReportSingleAllOfAttributes) *NullableMLTrainingReportSingleAllOfAttributes {
	return &NullableMLTrainingReportSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableMLTrainingReportSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMLTrainingReportSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



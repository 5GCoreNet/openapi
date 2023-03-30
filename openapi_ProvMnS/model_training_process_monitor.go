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

// checks if the TrainingProcessMonitor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrainingProcessMonitor{}

// TrainingProcessMonitor This data type is the \"ProcessMonitor\" data type defined in “genericNrm.yaml” with specialisations for usage in the \"MLTrainingProcess\".
type TrainingProcessMonitor struct {
	MLTrainingProcessId *string `json:"mLTrainingProcessId,omitempty"`
	Status *string `json:"status,omitempty"`
	ProgressPercentage *int32 `json:"progressPercentage,omitempty"`
	ProgressStateInfo *string `json:"progressStateInfo,omitempty"`
	ResultStateInfo *string `json:"resultStateInfo,omitempty"`
}

// NewTrainingProcessMonitor instantiates a new TrainingProcessMonitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrainingProcessMonitor() *TrainingProcessMonitor {
	this := TrainingProcessMonitor{}
	return &this
}

// NewTrainingProcessMonitorWithDefaults instantiates a new TrainingProcessMonitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrainingProcessMonitorWithDefaults() *TrainingProcessMonitor {
	this := TrainingProcessMonitor{}
	return &this
}

// GetMLTrainingProcessId returns the MLTrainingProcessId field value if set, zero value otherwise.
func (o *TrainingProcessMonitor) GetMLTrainingProcessId() string {
	if o == nil || IsNil(o.MLTrainingProcessId) {
		var ret string
		return ret
	}
	return *o.MLTrainingProcessId
}

// GetMLTrainingProcessIdOk returns a tuple with the MLTrainingProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainingProcessMonitor) GetMLTrainingProcessIdOk() (*string, bool) {
	if o == nil || IsNil(o.MLTrainingProcessId) {
		return nil, false
	}
	return o.MLTrainingProcessId, true
}

// HasMLTrainingProcessId returns a boolean if a field has been set.
func (o *TrainingProcessMonitor) HasMLTrainingProcessId() bool {
	if o != nil && !IsNil(o.MLTrainingProcessId) {
		return true
	}

	return false
}

// SetMLTrainingProcessId gets a reference to the given string and assigns it to the MLTrainingProcessId field.
func (o *TrainingProcessMonitor) SetMLTrainingProcessId(v string) {
	o.MLTrainingProcessId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TrainingProcessMonitor) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainingProcessMonitor) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TrainingProcessMonitor) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TrainingProcessMonitor) SetStatus(v string) {
	o.Status = &v
}

// GetProgressPercentage returns the ProgressPercentage field value if set, zero value otherwise.
func (o *TrainingProcessMonitor) GetProgressPercentage() int32 {
	if o == nil || IsNil(o.ProgressPercentage) {
		var ret int32
		return ret
	}
	return *o.ProgressPercentage
}

// GetProgressPercentageOk returns a tuple with the ProgressPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainingProcessMonitor) GetProgressPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.ProgressPercentage) {
		return nil, false
	}
	return o.ProgressPercentage, true
}

// HasProgressPercentage returns a boolean if a field has been set.
func (o *TrainingProcessMonitor) HasProgressPercentage() bool {
	if o != nil && !IsNil(o.ProgressPercentage) {
		return true
	}

	return false
}

// SetProgressPercentage gets a reference to the given int32 and assigns it to the ProgressPercentage field.
func (o *TrainingProcessMonitor) SetProgressPercentage(v int32) {
	o.ProgressPercentage = &v
}

// GetProgressStateInfo returns the ProgressStateInfo field value if set, zero value otherwise.
func (o *TrainingProcessMonitor) GetProgressStateInfo() string {
	if o == nil || IsNil(o.ProgressStateInfo) {
		var ret string
		return ret
	}
	return *o.ProgressStateInfo
}

// GetProgressStateInfoOk returns a tuple with the ProgressStateInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainingProcessMonitor) GetProgressStateInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ProgressStateInfo) {
		return nil, false
	}
	return o.ProgressStateInfo, true
}

// HasProgressStateInfo returns a boolean if a field has been set.
func (o *TrainingProcessMonitor) HasProgressStateInfo() bool {
	if o != nil && !IsNil(o.ProgressStateInfo) {
		return true
	}

	return false
}

// SetProgressStateInfo gets a reference to the given string and assigns it to the ProgressStateInfo field.
func (o *TrainingProcessMonitor) SetProgressStateInfo(v string) {
	o.ProgressStateInfo = &v
}

// GetResultStateInfo returns the ResultStateInfo field value if set, zero value otherwise.
func (o *TrainingProcessMonitor) GetResultStateInfo() string {
	if o == nil || IsNil(o.ResultStateInfo) {
		var ret string
		return ret
	}
	return *o.ResultStateInfo
}

// GetResultStateInfoOk returns a tuple with the ResultStateInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainingProcessMonitor) GetResultStateInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ResultStateInfo) {
		return nil, false
	}
	return o.ResultStateInfo, true
}

// HasResultStateInfo returns a boolean if a field has been set.
func (o *TrainingProcessMonitor) HasResultStateInfo() bool {
	if o != nil && !IsNil(o.ResultStateInfo) {
		return true
	}

	return false
}

// SetResultStateInfo gets a reference to the given string and assigns it to the ResultStateInfo field.
func (o *TrainingProcessMonitor) SetResultStateInfo(v string) {
	o.ResultStateInfo = &v
}

func (o TrainingProcessMonitor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrainingProcessMonitor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MLTrainingProcessId) {
		toSerialize["mLTrainingProcessId"] = o.MLTrainingProcessId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ProgressPercentage) {
		toSerialize["progressPercentage"] = o.ProgressPercentage
	}
	if !IsNil(o.ProgressStateInfo) {
		toSerialize["progressStateInfo"] = o.ProgressStateInfo
	}
	if !IsNil(o.ResultStateInfo) {
		toSerialize["resultStateInfo"] = o.ResultStateInfo
	}
	return toSerialize, nil
}

type NullableTrainingProcessMonitor struct {
	value *TrainingProcessMonitor
	isSet bool
}

func (v NullableTrainingProcessMonitor) Get() *TrainingProcessMonitor {
	return v.value
}

func (v *NullableTrainingProcessMonitor) Set(val *TrainingProcessMonitor) {
	v.value = val
	v.isSet = true
}

func (v NullableTrainingProcessMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullableTrainingProcessMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrainingProcessMonitor(val *TrainingProcessMonitor) *NullableTrainingProcessMonitor {
	return &NullableTrainingProcessMonitor{value: val, isSet: true}
}

func (v NullableTrainingProcessMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrainingProcessMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


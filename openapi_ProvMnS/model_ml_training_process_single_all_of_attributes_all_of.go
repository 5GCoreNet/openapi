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

// checks if the MLTrainingProcessSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MLTrainingProcessSingleAllOfAttributesAllOf{}

// MLTrainingProcessSingleAllOfAttributesAllOf struct for MLTrainingProcessSingleAllOfAttributesAllOf
type MLTrainingProcessSingleAllOfAttributesAllOf struct {
	MLTrainingProcessId   *string                 `json:"mLTrainingProcessId,omitempty"`
	Priority              *int32                  `json:"priority,omitempty"`
	TerminationConditions *string                 `json:"terminationConditions,omitempty"`
	ProgressStatus        *TrainingProcessMonitor `json:"progressStatus,omitempty"`
	CancelProcess         *bool                   `json:"cancelProcess,omitempty"`
	SuspendProcess        *bool                   `json:"suspendProcess,omitempty"`
	TrainingRequestRef    []string                `json:"trainingRequestRef,omitempty"`
	TrainingReportRef     *string                 `json:"trainingReportRef,omitempty"`
}

// NewMLTrainingProcessSingleAllOfAttributesAllOf instantiates a new MLTrainingProcessSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMLTrainingProcessSingleAllOfAttributesAllOf() *MLTrainingProcessSingleAllOfAttributesAllOf {
	this := MLTrainingProcessSingleAllOfAttributesAllOf{}
	return &this
}

// NewMLTrainingProcessSingleAllOfAttributesAllOfWithDefaults instantiates a new MLTrainingProcessSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMLTrainingProcessSingleAllOfAttributesAllOfWithDefaults() *MLTrainingProcessSingleAllOfAttributesAllOf {
	this := MLTrainingProcessSingleAllOfAttributesAllOf{}
	return &this
}

// GetMLTrainingProcessId returns the MLTrainingProcessId field value if set, zero value otherwise.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetMLTrainingProcessId() string {
	if o == nil || IsNil(o.MLTrainingProcessId) {
		var ret string
		return ret
	}
	return *o.MLTrainingProcessId
}

// GetMLTrainingProcessIdOk returns a tuple with the MLTrainingProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetMLTrainingProcessIdOk() (*string, bool) {
	if o == nil || IsNil(o.MLTrainingProcessId) {
		return nil, false
	}
	return o.MLTrainingProcessId, true
}

// HasMLTrainingProcessId returns a boolean if a field has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) HasMLTrainingProcessId() bool {
	if o != nil && !IsNil(o.MLTrainingProcessId) {
		return true
	}

	return false
}

// SetMLTrainingProcessId gets a reference to the given string and assigns it to the MLTrainingProcessId field.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) SetMLTrainingProcessId(v string) {
	o.MLTrainingProcessId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) SetPriority(v int32) {
	o.Priority = &v
}

// GetTerminationConditions returns the TerminationConditions field value if set, zero value otherwise.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetTerminationConditions() string {
	if o == nil || IsNil(o.TerminationConditions) {
		var ret string
		return ret
	}
	return *o.TerminationConditions
}

// GetTerminationConditionsOk returns a tuple with the TerminationConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetTerminationConditionsOk() (*string, bool) {
	if o == nil || IsNil(o.TerminationConditions) {
		return nil, false
	}
	return o.TerminationConditions, true
}

// HasTerminationConditions returns a boolean if a field has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) HasTerminationConditions() bool {
	if o != nil && !IsNil(o.TerminationConditions) {
		return true
	}

	return false
}

// SetTerminationConditions gets a reference to the given string and assigns it to the TerminationConditions field.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) SetTerminationConditions(v string) {
	o.TerminationConditions = &v
}

// GetProgressStatus returns the ProgressStatus field value if set, zero value otherwise.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetProgressStatus() TrainingProcessMonitor {
	if o == nil || IsNil(o.ProgressStatus) {
		var ret TrainingProcessMonitor
		return ret
	}
	return *o.ProgressStatus
}

// GetProgressStatusOk returns a tuple with the ProgressStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetProgressStatusOk() (*TrainingProcessMonitor, bool) {
	if o == nil || IsNil(o.ProgressStatus) {
		return nil, false
	}
	return o.ProgressStatus, true
}

// HasProgressStatus returns a boolean if a field has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) HasProgressStatus() bool {
	if o != nil && !IsNil(o.ProgressStatus) {
		return true
	}

	return false
}

// SetProgressStatus gets a reference to the given TrainingProcessMonitor and assigns it to the ProgressStatus field.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) SetProgressStatus(v TrainingProcessMonitor) {
	o.ProgressStatus = &v
}

// GetCancelProcess returns the CancelProcess field value if set, zero value otherwise.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetCancelProcess() bool {
	if o == nil || IsNil(o.CancelProcess) {
		var ret bool
		return ret
	}
	return *o.CancelProcess
}

// GetCancelProcessOk returns a tuple with the CancelProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetCancelProcessOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelProcess) {
		return nil, false
	}
	return o.CancelProcess, true
}

// HasCancelProcess returns a boolean if a field has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) HasCancelProcess() bool {
	if o != nil && !IsNil(o.CancelProcess) {
		return true
	}

	return false
}

// SetCancelProcess gets a reference to the given bool and assigns it to the CancelProcess field.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) SetCancelProcess(v bool) {
	o.CancelProcess = &v
}

// GetSuspendProcess returns the SuspendProcess field value if set, zero value otherwise.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetSuspendProcess() bool {
	if o == nil || IsNil(o.SuspendProcess) {
		var ret bool
		return ret
	}
	return *o.SuspendProcess
}

// GetSuspendProcessOk returns a tuple with the SuspendProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetSuspendProcessOk() (*bool, bool) {
	if o == nil || IsNil(o.SuspendProcess) {
		return nil, false
	}
	return o.SuspendProcess, true
}

// HasSuspendProcess returns a boolean if a field has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) HasSuspendProcess() bool {
	if o != nil && !IsNil(o.SuspendProcess) {
		return true
	}

	return false
}

// SetSuspendProcess gets a reference to the given bool and assigns it to the SuspendProcess field.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) SetSuspendProcess(v bool) {
	o.SuspendProcess = &v
}

// GetTrainingRequestRef returns the TrainingRequestRef field value if set, zero value otherwise.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetTrainingRequestRef() []string {
	if o == nil || IsNil(o.TrainingRequestRef) {
		var ret []string
		return ret
	}
	return o.TrainingRequestRef
}

// GetTrainingRequestRefOk returns a tuple with the TrainingRequestRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetTrainingRequestRefOk() ([]string, bool) {
	if o == nil || IsNil(o.TrainingRequestRef) {
		return nil, false
	}
	return o.TrainingRequestRef, true
}

// HasTrainingRequestRef returns a boolean if a field has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) HasTrainingRequestRef() bool {
	if o != nil && !IsNil(o.TrainingRequestRef) {
		return true
	}

	return false
}

// SetTrainingRequestRef gets a reference to the given []string and assigns it to the TrainingRequestRef field.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) SetTrainingRequestRef(v []string) {
	o.TrainingRequestRef = v
}

// GetTrainingReportRef returns the TrainingReportRef field value if set, zero value otherwise.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetTrainingReportRef() string {
	if o == nil || IsNil(o.TrainingReportRef) {
		var ret string
		return ret
	}
	return *o.TrainingReportRef
}

// GetTrainingReportRefOk returns a tuple with the TrainingReportRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) GetTrainingReportRefOk() (*string, bool) {
	if o == nil || IsNil(o.TrainingReportRef) {
		return nil, false
	}
	return o.TrainingReportRef, true
}

// HasTrainingReportRef returns a boolean if a field has been set.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) HasTrainingReportRef() bool {
	if o != nil && !IsNil(o.TrainingReportRef) {
		return true
	}

	return false
}

// SetTrainingReportRef gets a reference to the given string and assigns it to the TrainingReportRef field.
func (o *MLTrainingProcessSingleAllOfAttributesAllOf) SetTrainingReportRef(v string) {
	o.TrainingReportRef = &v
}

func (o MLTrainingProcessSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MLTrainingProcessSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MLTrainingProcessId) {
		toSerialize["mLTrainingProcessId"] = o.MLTrainingProcessId
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.TerminationConditions) {
		toSerialize["terminationConditions"] = o.TerminationConditions
	}
	if !IsNil(o.ProgressStatus) {
		toSerialize["progressStatus"] = o.ProgressStatus
	}
	if !IsNil(o.CancelProcess) {
		toSerialize["cancelProcess"] = o.CancelProcess
	}
	if !IsNil(o.SuspendProcess) {
		toSerialize["suspendProcess"] = o.SuspendProcess
	}
	if !IsNil(o.TrainingRequestRef) {
		toSerialize["trainingRequestRef"] = o.TrainingRequestRef
	}
	if !IsNil(o.TrainingReportRef) {
		toSerialize["trainingReportRef"] = o.TrainingReportRef
	}
	return toSerialize, nil
}

type NullableMLTrainingProcessSingleAllOfAttributesAllOf struct {
	value *MLTrainingProcessSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableMLTrainingProcessSingleAllOfAttributesAllOf) Get() *MLTrainingProcessSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableMLTrainingProcessSingleAllOfAttributesAllOf) Set(val *MLTrainingProcessSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMLTrainingProcessSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMLTrainingProcessSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMLTrainingProcessSingleAllOfAttributesAllOf(val *MLTrainingProcessSingleAllOfAttributesAllOf) *NullableMLTrainingProcessSingleAllOfAttributesAllOf {
	return &NullableMLTrainingProcessSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableMLTrainingProcessSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMLTrainingProcessSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

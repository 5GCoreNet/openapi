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

// checks if the ManagedElementSingle4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedElementSingle4{}

// ManagedElementSingle4 struct for ManagedElementSingle4
type ManagedElementSingle4 struct {
	Top
	Attributes             *ManagedElementAttr            `json:"attributes,omitempty"`
	MnsAgent               []MnsAgentSingle               `json:"MnsAgent,omitempty"`
	PerfMetricJob          []PerfMetricJobSingle          `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor       []ThresholdMonitorSingle       `json:"ThresholdMonitor,omitempty"`
	TraceJob               []TraceJobSingle               `json:"TraceJob,omitempty"`
	NtfSubscriptionControl []NtfSubscriptionControlSingle `json:"NtfSubscriptionControl,omitempty"`
	AlarmList              *AlarmListSingle               `json:"AlarmList,omitempty"`
	FileDownloadJob        []FileDownloadJobSingle        `json:"FileDownloadJob,omitempty"`
	Files                  []FilesSingle                  `json:"Files,omitempty"`
	MLTrainingFunction     []MLTrainingFunctionSingle     `json:"MLTrainingFunction,omitempty"`
}

// NewManagedElementSingle4 instantiates a new ManagedElementSingle4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedElementSingle4(id NullableString) *ManagedElementSingle4 {
	this := ManagedElementSingle4{}
	this.Id = id
	return &this
}

// NewManagedElementSingle4WithDefaults instantiates a new ManagedElementSingle4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedElementSingle4WithDefaults() *ManagedElementSingle4 {
	this := ManagedElementSingle4{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagedElementSingle4) GetAttributes() ManagedElementAttr {
	if o == nil || IsNil(o.Attributes) {
		var ret ManagedElementAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle4) GetAttributesOk() (*ManagedElementAttr, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagedElementSingle4) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedElementAttr and assigns it to the Attributes field.
func (o *ManagedElementSingle4) SetAttributes(v ManagedElementAttr) {
	o.Attributes = &v
}

// GetMnsAgent returns the MnsAgent field value if set, zero value otherwise.
func (o *ManagedElementSingle4) GetMnsAgent() []MnsAgentSingle {
	if o == nil || IsNil(o.MnsAgent) {
		var ret []MnsAgentSingle
		return ret
	}
	return o.MnsAgent
}

// GetMnsAgentOk returns a tuple with the MnsAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle4) GetMnsAgentOk() ([]MnsAgentSingle, bool) {
	if o == nil || IsNil(o.MnsAgent) {
		return nil, false
	}
	return o.MnsAgent, true
}

// HasMnsAgent returns a boolean if a field has been set.
func (o *ManagedElementSingle4) HasMnsAgent() bool {
	if o != nil && !IsNil(o.MnsAgent) {
		return true
	}

	return false
}

// SetMnsAgent gets a reference to the given []MnsAgentSingle and assigns it to the MnsAgent field.
func (o *ManagedElementSingle4) SetMnsAgent(v []MnsAgentSingle) {
	o.MnsAgent = v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *ManagedElementSingle4) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || IsNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle4) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || IsNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *ManagedElementSingle4) HasPerfMetricJob() bool {
	if o != nil && !IsNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *ManagedElementSingle4) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *ManagedElementSingle4) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || IsNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle4) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || IsNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *ManagedElementSingle4) HasThresholdMonitor() bool {
	if o != nil && !IsNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *ManagedElementSingle4) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *ManagedElementSingle4) GetTraceJob() []TraceJobSingle {
	if o == nil || IsNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle4) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || IsNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *ManagedElementSingle4) HasTraceJob() bool {
	if o != nil && !IsNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *ManagedElementSingle4) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetNtfSubscriptionControl returns the NtfSubscriptionControl field value if set, zero value otherwise.
func (o *ManagedElementSingle4) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle {
	if o == nil || IsNil(o.NtfSubscriptionControl) {
		var ret []NtfSubscriptionControlSingle
		return ret
	}
	return o.NtfSubscriptionControl
}

// GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle4) GetNtfSubscriptionControlOk() ([]NtfSubscriptionControlSingle, bool) {
	if o == nil || IsNil(o.NtfSubscriptionControl) {
		return nil, false
	}
	return o.NtfSubscriptionControl, true
}

// HasNtfSubscriptionControl returns a boolean if a field has been set.
func (o *ManagedElementSingle4) HasNtfSubscriptionControl() bool {
	if o != nil && !IsNil(o.NtfSubscriptionControl) {
		return true
	}

	return false
}

// SetNtfSubscriptionControl gets a reference to the given []NtfSubscriptionControlSingle and assigns it to the NtfSubscriptionControl field.
func (o *ManagedElementSingle4) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle) {
	o.NtfSubscriptionControl = v
}

// GetAlarmList returns the AlarmList field value if set, zero value otherwise.
func (o *ManagedElementSingle4) GetAlarmList() AlarmListSingle {
	if o == nil || IsNil(o.AlarmList) {
		var ret AlarmListSingle
		return ret
	}
	return *o.AlarmList
}

// GetAlarmListOk returns a tuple with the AlarmList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle4) GetAlarmListOk() (*AlarmListSingle, bool) {
	if o == nil || IsNil(o.AlarmList) {
		return nil, false
	}
	return o.AlarmList, true
}

// HasAlarmList returns a boolean if a field has been set.
func (o *ManagedElementSingle4) HasAlarmList() bool {
	if o != nil && !IsNil(o.AlarmList) {
		return true
	}

	return false
}

// SetAlarmList gets a reference to the given AlarmListSingle and assigns it to the AlarmList field.
func (o *ManagedElementSingle4) SetAlarmList(v AlarmListSingle) {
	o.AlarmList = &v
}

// GetFileDownloadJob returns the FileDownloadJob field value if set, zero value otherwise.
func (o *ManagedElementSingle4) GetFileDownloadJob() []FileDownloadJobSingle {
	if o == nil || IsNil(o.FileDownloadJob) {
		var ret []FileDownloadJobSingle
		return ret
	}
	return o.FileDownloadJob
}

// GetFileDownloadJobOk returns a tuple with the FileDownloadJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle4) GetFileDownloadJobOk() ([]FileDownloadJobSingle, bool) {
	if o == nil || IsNil(o.FileDownloadJob) {
		return nil, false
	}
	return o.FileDownloadJob, true
}

// HasFileDownloadJob returns a boolean if a field has been set.
func (o *ManagedElementSingle4) HasFileDownloadJob() bool {
	if o != nil && !IsNil(o.FileDownloadJob) {
		return true
	}

	return false
}

// SetFileDownloadJob gets a reference to the given []FileDownloadJobSingle and assigns it to the FileDownloadJob field.
func (o *ManagedElementSingle4) SetFileDownloadJob(v []FileDownloadJobSingle) {
	o.FileDownloadJob = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *ManagedElementSingle4) GetFiles() []FilesSingle {
	if o == nil || IsNil(o.Files) {
		var ret []FilesSingle
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle4) GetFilesOk() ([]FilesSingle, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *ManagedElementSingle4) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FilesSingle and assigns it to the Files field.
func (o *ManagedElementSingle4) SetFiles(v []FilesSingle) {
	o.Files = v
}

// GetMLTrainingFunction returns the MLTrainingFunction field value if set, zero value otherwise.
func (o *ManagedElementSingle4) GetMLTrainingFunction() []MLTrainingFunctionSingle {
	if o == nil || IsNil(o.MLTrainingFunction) {
		var ret []MLTrainingFunctionSingle
		return ret
	}
	return o.MLTrainingFunction
}

// GetMLTrainingFunctionOk returns a tuple with the MLTrainingFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle4) GetMLTrainingFunctionOk() ([]MLTrainingFunctionSingle, bool) {
	if o == nil || IsNil(o.MLTrainingFunction) {
		return nil, false
	}
	return o.MLTrainingFunction, true
}

// HasMLTrainingFunction returns a boolean if a field has been set.
func (o *ManagedElementSingle4) HasMLTrainingFunction() bool {
	if o != nil && !IsNil(o.MLTrainingFunction) {
		return true
	}

	return false
}

// SetMLTrainingFunction gets a reference to the given []MLTrainingFunctionSingle and assigns it to the MLTrainingFunction field.
func (o *ManagedElementSingle4) SetMLTrainingFunction(v []MLTrainingFunctionSingle) {
	o.MLTrainingFunction = v
}

func (o ManagedElementSingle4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedElementSingle4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop, errTop := json.Marshal(o.Top)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	errTop = json.Unmarshal([]byte(serializedTop), &toSerialize)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.MnsAgent) {
		toSerialize["MnsAgent"] = o.MnsAgent
	}
	if !IsNil(o.PerfMetricJob) {
		toSerialize["PerfMetricJob"] = o.PerfMetricJob
	}
	if !IsNil(o.ThresholdMonitor) {
		toSerialize["ThresholdMonitor"] = o.ThresholdMonitor
	}
	if !IsNil(o.TraceJob) {
		toSerialize["TraceJob"] = o.TraceJob
	}
	if !IsNil(o.NtfSubscriptionControl) {
		toSerialize["NtfSubscriptionControl"] = o.NtfSubscriptionControl
	}
	if !IsNil(o.AlarmList) {
		toSerialize["AlarmList"] = o.AlarmList
	}
	if !IsNil(o.FileDownloadJob) {
		toSerialize["FileDownloadJob"] = o.FileDownloadJob
	}
	if !IsNil(o.Files) {
		toSerialize["Files"] = o.Files
	}
	if !IsNil(o.MLTrainingFunction) {
		toSerialize["MLTrainingFunction"] = o.MLTrainingFunction
	}
	return toSerialize, nil
}

type NullableManagedElementSingle4 struct {
	value *ManagedElementSingle4
	isSet bool
}

func (v NullableManagedElementSingle4) Get() *ManagedElementSingle4 {
	return v.value
}

func (v *NullableManagedElementSingle4) Set(val *ManagedElementSingle4) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedElementSingle4) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedElementSingle4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedElementSingle4(val *ManagedElementSingle4) *NullableManagedElementSingle4 {
	return &NullableManagedElementSingle4{value: val, isSet: true}
}

func (v NullableManagedElementSingle4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedElementSingle4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

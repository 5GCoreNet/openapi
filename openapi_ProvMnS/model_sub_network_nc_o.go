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

// checks if the SubNetworkNcO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubNetworkNcO{}

// SubNetworkNcO struct for SubNetworkNcO
type SubNetworkNcO struct {
	ManagementNode []ManagementNodeSingle `json:"ManagementNode,omitempty"`
	MnsAgent []MnsAgentSingle `json:"MnsAgent,omitempty"`
	MeContext []MeContextSingle `json:"MeContext,omitempty"`
	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`
	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`
	ManagementDataCollection []ManagementDataCollectionSingle `json:"ManagementDataCollection,omitempty"`
	NtfSubscriptionControl []NtfSubscriptionControlSingle `json:"NtfSubscriptionControl,omitempty"`
	AlarmList *AlarmListSingle `json:"AlarmList,omitempty"`
	FileDownloadJob []FileDownloadJobSingle `json:"FileDownloadJob,omitempty"`
	Files []FilesSingle `json:"Files,omitempty"`
	MnsRegistry *MnsRegistrySingle `json:"MnsRegistry,omitempty"`
}

// NewSubNetworkNcO instantiates a new SubNetworkNcO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkNcO() *SubNetworkNcO {
	this := SubNetworkNcO{}
	return &this
}

// NewSubNetworkNcOWithDefaults instantiates a new SubNetworkNcO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkNcOWithDefaults() *SubNetworkNcO {
	this := SubNetworkNcO{}
	return &this
}

// GetManagementNode returns the ManagementNode field value if set, zero value otherwise.
func (o *SubNetworkNcO) GetManagementNode() []ManagementNodeSingle {
	if o == nil || IsNil(o.ManagementNode) {
		var ret []ManagementNodeSingle
		return ret
	}
	return o.ManagementNode
}

// GetManagementNodeOk returns a tuple with the ManagementNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkNcO) GetManagementNodeOk() ([]ManagementNodeSingle, bool) {
	if o == nil || IsNil(o.ManagementNode) {
		return nil, false
	}
	return o.ManagementNode, true
}

// HasManagementNode returns a boolean if a field has been set.
func (o *SubNetworkNcO) HasManagementNode() bool {
	if o != nil && !IsNil(o.ManagementNode) {
		return true
	}

	return false
}

// SetManagementNode gets a reference to the given []ManagementNodeSingle and assigns it to the ManagementNode field.
func (o *SubNetworkNcO) SetManagementNode(v []ManagementNodeSingle) {
	o.ManagementNode = v
}

// GetMnsAgent returns the MnsAgent field value if set, zero value otherwise.
func (o *SubNetworkNcO) GetMnsAgent() []MnsAgentSingle {
	if o == nil || IsNil(o.MnsAgent) {
		var ret []MnsAgentSingle
		return ret
	}
	return o.MnsAgent
}

// GetMnsAgentOk returns a tuple with the MnsAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkNcO) GetMnsAgentOk() ([]MnsAgentSingle, bool) {
	if o == nil || IsNil(o.MnsAgent) {
		return nil, false
	}
	return o.MnsAgent, true
}

// HasMnsAgent returns a boolean if a field has been set.
func (o *SubNetworkNcO) HasMnsAgent() bool {
	if o != nil && !IsNil(o.MnsAgent) {
		return true
	}

	return false
}

// SetMnsAgent gets a reference to the given []MnsAgentSingle and assigns it to the MnsAgent field.
func (o *SubNetworkNcO) SetMnsAgent(v []MnsAgentSingle) {
	o.MnsAgent = v
}

// GetMeContext returns the MeContext field value if set, zero value otherwise.
func (o *SubNetworkNcO) GetMeContext() []MeContextSingle {
	if o == nil || IsNil(o.MeContext) {
		var ret []MeContextSingle
		return ret
	}
	return o.MeContext
}

// GetMeContextOk returns a tuple with the MeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkNcO) GetMeContextOk() ([]MeContextSingle, bool) {
	if o == nil || IsNil(o.MeContext) {
		return nil, false
	}
	return o.MeContext, true
}

// HasMeContext returns a boolean if a field has been set.
func (o *SubNetworkNcO) HasMeContext() bool {
	if o != nil && !IsNil(o.MeContext) {
		return true
	}

	return false
}

// SetMeContext gets a reference to the given []MeContextSingle and assigns it to the MeContext field.
func (o *SubNetworkNcO) SetMeContext(v []MeContextSingle) {
	o.MeContext = v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *SubNetworkNcO) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || IsNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkNcO) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || IsNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *SubNetworkNcO) HasPerfMetricJob() bool {
	if o != nil && !IsNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *SubNetworkNcO) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *SubNetworkNcO) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || IsNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkNcO) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || IsNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *SubNetworkNcO) HasThresholdMonitor() bool {
	if o != nil && !IsNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *SubNetworkNcO) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *SubNetworkNcO) GetTraceJob() []TraceJobSingle {
	if o == nil || IsNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkNcO) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || IsNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *SubNetworkNcO) HasTraceJob() bool {
	if o != nil && !IsNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *SubNetworkNcO) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetManagementDataCollection returns the ManagementDataCollection field value if set, zero value otherwise.
func (o *SubNetworkNcO) GetManagementDataCollection() []ManagementDataCollectionSingle {
	if o == nil || IsNil(o.ManagementDataCollection) {
		var ret []ManagementDataCollectionSingle
		return ret
	}
	return o.ManagementDataCollection
}

// GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkNcO) GetManagementDataCollectionOk() ([]ManagementDataCollectionSingle, bool) {
	if o == nil || IsNil(o.ManagementDataCollection) {
		return nil, false
	}
	return o.ManagementDataCollection, true
}

// HasManagementDataCollection returns a boolean if a field has been set.
func (o *SubNetworkNcO) HasManagementDataCollection() bool {
	if o != nil && !IsNil(o.ManagementDataCollection) {
		return true
	}

	return false
}

// SetManagementDataCollection gets a reference to the given []ManagementDataCollectionSingle and assigns it to the ManagementDataCollection field.
func (o *SubNetworkNcO) SetManagementDataCollection(v []ManagementDataCollectionSingle) {
	o.ManagementDataCollection = v
}

// GetNtfSubscriptionControl returns the NtfSubscriptionControl field value if set, zero value otherwise.
func (o *SubNetworkNcO) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle {
	if o == nil || IsNil(o.NtfSubscriptionControl) {
		var ret []NtfSubscriptionControlSingle
		return ret
	}
	return o.NtfSubscriptionControl
}

// GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkNcO) GetNtfSubscriptionControlOk() ([]NtfSubscriptionControlSingle, bool) {
	if o == nil || IsNil(o.NtfSubscriptionControl) {
		return nil, false
	}
	return o.NtfSubscriptionControl, true
}

// HasNtfSubscriptionControl returns a boolean if a field has been set.
func (o *SubNetworkNcO) HasNtfSubscriptionControl() bool {
	if o != nil && !IsNil(o.NtfSubscriptionControl) {
		return true
	}

	return false
}

// SetNtfSubscriptionControl gets a reference to the given []NtfSubscriptionControlSingle and assigns it to the NtfSubscriptionControl field.
func (o *SubNetworkNcO) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle) {
	o.NtfSubscriptionControl = v
}

// GetAlarmList returns the AlarmList field value if set, zero value otherwise.
func (o *SubNetworkNcO) GetAlarmList() AlarmListSingle {
	if o == nil || IsNil(o.AlarmList) {
		var ret AlarmListSingle
		return ret
	}
	return *o.AlarmList
}

// GetAlarmListOk returns a tuple with the AlarmList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkNcO) GetAlarmListOk() (*AlarmListSingle, bool) {
	if o == nil || IsNil(o.AlarmList) {
		return nil, false
	}
	return o.AlarmList, true
}

// HasAlarmList returns a boolean if a field has been set.
func (o *SubNetworkNcO) HasAlarmList() bool {
	if o != nil && !IsNil(o.AlarmList) {
		return true
	}

	return false
}

// SetAlarmList gets a reference to the given AlarmListSingle and assigns it to the AlarmList field.
func (o *SubNetworkNcO) SetAlarmList(v AlarmListSingle) {
	o.AlarmList = &v
}

// GetFileDownloadJob returns the FileDownloadJob field value if set, zero value otherwise.
func (o *SubNetworkNcO) GetFileDownloadJob() []FileDownloadJobSingle {
	if o == nil || IsNil(o.FileDownloadJob) {
		var ret []FileDownloadJobSingle
		return ret
	}
	return o.FileDownloadJob
}

// GetFileDownloadJobOk returns a tuple with the FileDownloadJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkNcO) GetFileDownloadJobOk() ([]FileDownloadJobSingle, bool) {
	if o == nil || IsNil(o.FileDownloadJob) {
		return nil, false
	}
	return o.FileDownloadJob, true
}

// HasFileDownloadJob returns a boolean if a field has been set.
func (o *SubNetworkNcO) HasFileDownloadJob() bool {
	if o != nil && !IsNil(o.FileDownloadJob) {
		return true
	}

	return false
}

// SetFileDownloadJob gets a reference to the given []FileDownloadJobSingle and assigns it to the FileDownloadJob field.
func (o *SubNetworkNcO) SetFileDownloadJob(v []FileDownloadJobSingle) {
	o.FileDownloadJob = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *SubNetworkNcO) GetFiles() []FilesSingle {
	if o == nil || IsNil(o.Files) {
		var ret []FilesSingle
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkNcO) GetFilesOk() ([]FilesSingle, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *SubNetworkNcO) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FilesSingle and assigns it to the Files field.
func (o *SubNetworkNcO) SetFiles(v []FilesSingle) {
	o.Files = v
}

// GetMnsRegistry returns the MnsRegistry field value if set, zero value otherwise.
func (o *SubNetworkNcO) GetMnsRegistry() MnsRegistrySingle {
	if o == nil || IsNil(o.MnsRegistry) {
		var ret MnsRegistrySingle
		return ret
	}
	return *o.MnsRegistry
}

// GetMnsRegistryOk returns a tuple with the MnsRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkNcO) GetMnsRegistryOk() (*MnsRegistrySingle, bool) {
	if o == nil || IsNil(o.MnsRegistry) {
		return nil, false
	}
	return o.MnsRegistry, true
}

// HasMnsRegistry returns a boolean if a field has been set.
func (o *SubNetworkNcO) HasMnsRegistry() bool {
	if o != nil && !IsNil(o.MnsRegistry) {
		return true
	}

	return false
}

// SetMnsRegistry gets a reference to the given MnsRegistrySingle and assigns it to the MnsRegistry field.
func (o *SubNetworkNcO) SetMnsRegistry(v MnsRegistrySingle) {
	o.MnsRegistry = &v
}

func (o SubNetworkNcO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubNetworkNcO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManagementNode) {
		toSerialize["ManagementNode"] = o.ManagementNode
	}
	if !IsNil(o.MnsAgent) {
		toSerialize["MnsAgent"] = o.MnsAgent
	}
	if !IsNil(o.MeContext) {
		toSerialize["MeContext"] = o.MeContext
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
	if !IsNil(o.ManagementDataCollection) {
		toSerialize["ManagementDataCollection"] = o.ManagementDataCollection
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
	if !IsNil(o.MnsRegistry) {
		toSerialize["MnsRegistry"] = o.MnsRegistry
	}
	return toSerialize, nil
}

type NullableSubNetworkNcO struct {
	value *SubNetworkNcO
	isSet bool
}

func (v NullableSubNetworkNcO) Get() *SubNetworkNcO {
	return v.value
}

func (v *NullableSubNetworkNcO) Set(val *SubNetworkNcO) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkNcO) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkNcO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkNcO(val *SubNetworkNcO) *NullableSubNetworkNcO {
	return &NullableSubNetworkNcO{value: val, isSet: true}
}

func (v NullableSubNetworkNcO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkNcO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



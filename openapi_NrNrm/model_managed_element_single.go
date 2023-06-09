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

// checks if the ManagedElementSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedElementSingle{}

// ManagedElementSingle struct for ManagedElementSingle
type ManagedElementSingle struct {
	Top
	Attributes                *ManagedElementAttr              `json:"attributes,omitempty"`
	MnsAgent                  []MnsAgentSingle                 `json:"MnsAgent,omitempty"`
	PerfMetricJob             []PerfMetricJobSingle            `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor          []ThresholdMonitorSingle         `json:"ThresholdMonitor,omitempty"`
	TraceJob                  []TraceJobSingle                 `json:"TraceJob,omitempty"`
	NtfSubscriptionControl    []NtfSubscriptionControlSingle   `json:"NtfSubscriptionControl,omitempty"`
	AlarmList                 *AlarmListSingle                 `json:"AlarmList,omitempty"`
	FileDownloadJob           []FileDownloadJobSingle          `json:"FileDownloadJob,omitempty"`
	Files                     []FilesSingle                    `json:"Files,omitempty"`
	GnbDuFunction             []GnbDuFunctionSingle            `json:"GnbDuFunction,omitempty"`
	GnbCuUpFunction           []GnbCuUpFunctionSingle          `json:"GnbCuUpFunction,omitempty"`
	GnbCuCpFunction           []GnbCuCpFunctionSingle          `json:"GnbCuCpFunction,omitempty"`
	DESManagementFunction     *DESManagementFunctionSingle     `json:"DESManagementFunction,omitempty"`
	DRACHOptimizationFunction *DRACHOptimizationFunctionSingle `json:"DRACHOptimizationFunction,omitempty"`
	DMROFunction              *DMROFunctionSingle              `json:"DMROFunction,omitempty"`
	DLBOFunction              *DLBOFunctionSingle              `json:"DLBOFunction,omitempty"`
	DPCIConfigurationFunction *DPCIConfigurationFunctionSingle `json:"DPCIConfigurationFunction,omitempty"`
	CPCIConfigurationFunction *CPCIConfigurationFunctionSingle `json:"CPCIConfigurationFunction,omitempty"`
	CESManagementFunction     *CESManagementFunctionSingle     `json:"CESManagementFunction,omitempty"`
	Configurable5QISet        []Configurable5QISetSingle       `json:"Configurable5QISet,omitempty"`
	Dynamic5QISet             []Dynamic5QISetSingle            `json:"Dynamic5QISet,omitempty"`
}

// NewManagedElementSingle instantiates a new ManagedElementSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedElementSingle(id NullableString) *ManagedElementSingle {
	this := ManagedElementSingle{}
	this.Id = id
	return &this
}

// NewManagedElementSingleWithDefaults instantiates a new ManagedElementSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedElementSingleWithDefaults() *ManagedElementSingle {
	this := ManagedElementSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetAttributes() ManagedElementAttr {
	if o == nil || IsNil(o.Attributes) {
		var ret ManagedElementAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetAttributesOk() (*ManagedElementAttr, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedElementAttr and assigns it to the Attributes field.
func (o *ManagedElementSingle) SetAttributes(v ManagedElementAttr) {
	o.Attributes = &v
}

// GetMnsAgent returns the MnsAgent field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetMnsAgent() []MnsAgentSingle {
	if o == nil || IsNil(o.MnsAgent) {
		var ret []MnsAgentSingle
		return ret
	}
	return o.MnsAgent
}

// GetMnsAgentOk returns a tuple with the MnsAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetMnsAgentOk() ([]MnsAgentSingle, bool) {
	if o == nil || IsNil(o.MnsAgent) {
		return nil, false
	}
	return o.MnsAgent, true
}

// HasMnsAgent returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasMnsAgent() bool {
	if o != nil && !IsNil(o.MnsAgent) {
		return true
	}

	return false
}

// SetMnsAgent gets a reference to the given []MnsAgentSingle and assigns it to the MnsAgent field.
func (o *ManagedElementSingle) SetMnsAgent(v []MnsAgentSingle) {
	o.MnsAgent = v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || IsNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || IsNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasPerfMetricJob() bool {
	if o != nil && !IsNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *ManagedElementSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || IsNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || IsNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasThresholdMonitor() bool {
	if o != nil && !IsNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *ManagedElementSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || IsNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || IsNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasTraceJob() bool {
	if o != nil && !IsNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *ManagedElementSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetNtfSubscriptionControl returns the NtfSubscriptionControl field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle {
	if o == nil || IsNil(o.NtfSubscriptionControl) {
		var ret []NtfSubscriptionControlSingle
		return ret
	}
	return o.NtfSubscriptionControl
}

// GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetNtfSubscriptionControlOk() ([]NtfSubscriptionControlSingle, bool) {
	if o == nil || IsNil(o.NtfSubscriptionControl) {
		return nil, false
	}
	return o.NtfSubscriptionControl, true
}

// HasNtfSubscriptionControl returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasNtfSubscriptionControl() bool {
	if o != nil && !IsNil(o.NtfSubscriptionControl) {
		return true
	}

	return false
}

// SetNtfSubscriptionControl gets a reference to the given []NtfSubscriptionControlSingle and assigns it to the NtfSubscriptionControl field.
func (o *ManagedElementSingle) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle) {
	o.NtfSubscriptionControl = v
}

// GetAlarmList returns the AlarmList field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetAlarmList() AlarmListSingle {
	if o == nil || IsNil(o.AlarmList) {
		var ret AlarmListSingle
		return ret
	}
	return *o.AlarmList
}

// GetAlarmListOk returns a tuple with the AlarmList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetAlarmListOk() (*AlarmListSingle, bool) {
	if o == nil || IsNil(o.AlarmList) {
		return nil, false
	}
	return o.AlarmList, true
}

// HasAlarmList returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasAlarmList() bool {
	if o != nil && !IsNil(o.AlarmList) {
		return true
	}

	return false
}

// SetAlarmList gets a reference to the given AlarmListSingle and assigns it to the AlarmList field.
func (o *ManagedElementSingle) SetAlarmList(v AlarmListSingle) {
	o.AlarmList = &v
}

// GetFileDownloadJob returns the FileDownloadJob field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetFileDownloadJob() []FileDownloadJobSingle {
	if o == nil || IsNil(o.FileDownloadJob) {
		var ret []FileDownloadJobSingle
		return ret
	}
	return o.FileDownloadJob
}

// GetFileDownloadJobOk returns a tuple with the FileDownloadJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetFileDownloadJobOk() ([]FileDownloadJobSingle, bool) {
	if o == nil || IsNil(o.FileDownloadJob) {
		return nil, false
	}
	return o.FileDownloadJob, true
}

// HasFileDownloadJob returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasFileDownloadJob() bool {
	if o != nil && !IsNil(o.FileDownloadJob) {
		return true
	}

	return false
}

// SetFileDownloadJob gets a reference to the given []FileDownloadJobSingle and assigns it to the FileDownloadJob field.
func (o *ManagedElementSingle) SetFileDownloadJob(v []FileDownloadJobSingle) {
	o.FileDownloadJob = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetFiles() []FilesSingle {
	if o == nil || IsNil(o.Files) {
		var ret []FilesSingle
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetFilesOk() ([]FilesSingle, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FilesSingle and assigns it to the Files field.
func (o *ManagedElementSingle) SetFiles(v []FilesSingle) {
	o.Files = v
}

// GetGnbDuFunction returns the GnbDuFunction field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetGnbDuFunction() []GnbDuFunctionSingle {
	if o == nil || IsNil(o.GnbDuFunction) {
		var ret []GnbDuFunctionSingle
		return ret
	}
	return o.GnbDuFunction
}

// GetGnbDuFunctionOk returns a tuple with the GnbDuFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetGnbDuFunctionOk() ([]GnbDuFunctionSingle, bool) {
	if o == nil || IsNil(o.GnbDuFunction) {
		return nil, false
	}
	return o.GnbDuFunction, true
}

// HasGnbDuFunction returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasGnbDuFunction() bool {
	if o != nil && !IsNil(o.GnbDuFunction) {
		return true
	}

	return false
}

// SetGnbDuFunction gets a reference to the given []GnbDuFunctionSingle and assigns it to the GnbDuFunction field.
func (o *ManagedElementSingle) SetGnbDuFunction(v []GnbDuFunctionSingle) {
	o.GnbDuFunction = v
}

// GetGnbCuUpFunction returns the GnbCuUpFunction field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetGnbCuUpFunction() []GnbCuUpFunctionSingle {
	if o == nil || IsNil(o.GnbCuUpFunction) {
		var ret []GnbCuUpFunctionSingle
		return ret
	}
	return o.GnbCuUpFunction
}

// GetGnbCuUpFunctionOk returns a tuple with the GnbCuUpFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetGnbCuUpFunctionOk() ([]GnbCuUpFunctionSingle, bool) {
	if o == nil || IsNil(o.GnbCuUpFunction) {
		return nil, false
	}
	return o.GnbCuUpFunction, true
}

// HasGnbCuUpFunction returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasGnbCuUpFunction() bool {
	if o != nil && !IsNil(o.GnbCuUpFunction) {
		return true
	}

	return false
}

// SetGnbCuUpFunction gets a reference to the given []GnbCuUpFunctionSingle and assigns it to the GnbCuUpFunction field.
func (o *ManagedElementSingle) SetGnbCuUpFunction(v []GnbCuUpFunctionSingle) {
	o.GnbCuUpFunction = v
}

// GetGnbCuCpFunction returns the GnbCuCpFunction field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetGnbCuCpFunction() []GnbCuCpFunctionSingle {
	if o == nil || IsNil(o.GnbCuCpFunction) {
		var ret []GnbCuCpFunctionSingle
		return ret
	}
	return o.GnbCuCpFunction
}

// GetGnbCuCpFunctionOk returns a tuple with the GnbCuCpFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetGnbCuCpFunctionOk() ([]GnbCuCpFunctionSingle, bool) {
	if o == nil || IsNil(o.GnbCuCpFunction) {
		return nil, false
	}
	return o.GnbCuCpFunction, true
}

// HasGnbCuCpFunction returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasGnbCuCpFunction() bool {
	if o != nil && !IsNil(o.GnbCuCpFunction) {
		return true
	}

	return false
}

// SetGnbCuCpFunction gets a reference to the given []GnbCuCpFunctionSingle and assigns it to the GnbCuCpFunction field.
func (o *ManagedElementSingle) SetGnbCuCpFunction(v []GnbCuCpFunctionSingle) {
	o.GnbCuCpFunction = v
}

// GetDESManagementFunction returns the DESManagementFunction field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetDESManagementFunction() DESManagementFunctionSingle {
	if o == nil || IsNil(o.DESManagementFunction) {
		var ret DESManagementFunctionSingle
		return ret
	}
	return *o.DESManagementFunction
}

// GetDESManagementFunctionOk returns a tuple with the DESManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetDESManagementFunctionOk() (*DESManagementFunctionSingle, bool) {
	if o == nil || IsNil(o.DESManagementFunction) {
		return nil, false
	}
	return o.DESManagementFunction, true
}

// HasDESManagementFunction returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasDESManagementFunction() bool {
	if o != nil && !IsNil(o.DESManagementFunction) {
		return true
	}

	return false
}

// SetDESManagementFunction gets a reference to the given DESManagementFunctionSingle and assigns it to the DESManagementFunction field.
func (o *ManagedElementSingle) SetDESManagementFunction(v DESManagementFunctionSingle) {
	o.DESManagementFunction = &v
}

// GetDRACHOptimizationFunction returns the DRACHOptimizationFunction field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetDRACHOptimizationFunction() DRACHOptimizationFunctionSingle {
	if o == nil || IsNil(o.DRACHOptimizationFunction) {
		var ret DRACHOptimizationFunctionSingle
		return ret
	}
	return *o.DRACHOptimizationFunction
}

// GetDRACHOptimizationFunctionOk returns a tuple with the DRACHOptimizationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetDRACHOptimizationFunctionOk() (*DRACHOptimizationFunctionSingle, bool) {
	if o == nil || IsNil(o.DRACHOptimizationFunction) {
		return nil, false
	}
	return o.DRACHOptimizationFunction, true
}

// HasDRACHOptimizationFunction returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasDRACHOptimizationFunction() bool {
	if o != nil && !IsNil(o.DRACHOptimizationFunction) {
		return true
	}

	return false
}

// SetDRACHOptimizationFunction gets a reference to the given DRACHOptimizationFunctionSingle and assigns it to the DRACHOptimizationFunction field.
func (o *ManagedElementSingle) SetDRACHOptimizationFunction(v DRACHOptimizationFunctionSingle) {
	o.DRACHOptimizationFunction = &v
}

// GetDMROFunction returns the DMROFunction field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetDMROFunction() DMROFunctionSingle {
	if o == nil || IsNil(o.DMROFunction) {
		var ret DMROFunctionSingle
		return ret
	}
	return *o.DMROFunction
}

// GetDMROFunctionOk returns a tuple with the DMROFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetDMROFunctionOk() (*DMROFunctionSingle, bool) {
	if o == nil || IsNil(o.DMROFunction) {
		return nil, false
	}
	return o.DMROFunction, true
}

// HasDMROFunction returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasDMROFunction() bool {
	if o != nil && !IsNil(o.DMROFunction) {
		return true
	}

	return false
}

// SetDMROFunction gets a reference to the given DMROFunctionSingle and assigns it to the DMROFunction field.
func (o *ManagedElementSingle) SetDMROFunction(v DMROFunctionSingle) {
	o.DMROFunction = &v
}

// GetDLBOFunction returns the DLBOFunction field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetDLBOFunction() DLBOFunctionSingle {
	if o == nil || IsNil(o.DLBOFunction) {
		var ret DLBOFunctionSingle
		return ret
	}
	return *o.DLBOFunction
}

// GetDLBOFunctionOk returns a tuple with the DLBOFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetDLBOFunctionOk() (*DLBOFunctionSingle, bool) {
	if o == nil || IsNil(o.DLBOFunction) {
		return nil, false
	}
	return o.DLBOFunction, true
}

// HasDLBOFunction returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasDLBOFunction() bool {
	if o != nil && !IsNil(o.DLBOFunction) {
		return true
	}

	return false
}

// SetDLBOFunction gets a reference to the given DLBOFunctionSingle and assigns it to the DLBOFunction field.
func (o *ManagedElementSingle) SetDLBOFunction(v DLBOFunctionSingle) {
	o.DLBOFunction = &v
}

// GetDPCIConfigurationFunction returns the DPCIConfigurationFunction field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetDPCIConfigurationFunction() DPCIConfigurationFunctionSingle {
	if o == nil || IsNil(o.DPCIConfigurationFunction) {
		var ret DPCIConfigurationFunctionSingle
		return ret
	}
	return *o.DPCIConfigurationFunction
}

// GetDPCIConfigurationFunctionOk returns a tuple with the DPCIConfigurationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetDPCIConfigurationFunctionOk() (*DPCIConfigurationFunctionSingle, bool) {
	if o == nil || IsNil(o.DPCIConfigurationFunction) {
		return nil, false
	}
	return o.DPCIConfigurationFunction, true
}

// HasDPCIConfigurationFunction returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasDPCIConfigurationFunction() bool {
	if o != nil && !IsNil(o.DPCIConfigurationFunction) {
		return true
	}

	return false
}

// SetDPCIConfigurationFunction gets a reference to the given DPCIConfigurationFunctionSingle and assigns it to the DPCIConfigurationFunction field.
func (o *ManagedElementSingle) SetDPCIConfigurationFunction(v DPCIConfigurationFunctionSingle) {
	o.DPCIConfigurationFunction = &v
}

// GetCPCIConfigurationFunction returns the CPCIConfigurationFunction field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetCPCIConfigurationFunction() CPCIConfigurationFunctionSingle {
	if o == nil || IsNil(o.CPCIConfigurationFunction) {
		var ret CPCIConfigurationFunctionSingle
		return ret
	}
	return *o.CPCIConfigurationFunction
}

// GetCPCIConfigurationFunctionOk returns a tuple with the CPCIConfigurationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetCPCIConfigurationFunctionOk() (*CPCIConfigurationFunctionSingle, bool) {
	if o == nil || IsNil(o.CPCIConfigurationFunction) {
		return nil, false
	}
	return o.CPCIConfigurationFunction, true
}

// HasCPCIConfigurationFunction returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasCPCIConfigurationFunction() bool {
	if o != nil && !IsNil(o.CPCIConfigurationFunction) {
		return true
	}

	return false
}

// SetCPCIConfigurationFunction gets a reference to the given CPCIConfigurationFunctionSingle and assigns it to the CPCIConfigurationFunction field.
func (o *ManagedElementSingle) SetCPCIConfigurationFunction(v CPCIConfigurationFunctionSingle) {
	o.CPCIConfigurationFunction = &v
}

// GetCESManagementFunction returns the CESManagementFunction field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetCESManagementFunction() CESManagementFunctionSingle {
	if o == nil || IsNil(o.CESManagementFunction) {
		var ret CESManagementFunctionSingle
		return ret
	}
	return *o.CESManagementFunction
}

// GetCESManagementFunctionOk returns a tuple with the CESManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetCESManagementFunctionOk() (*CESManagementFunctionSingle, bool) {
	if o == nil || IsNil(o.CESManagementFunction) {
		return nil, false
	}
	return o.CESManagementFunction, true
}

// HasCESManagementFunction returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasCESManagementFunction() bool {
	if o != nil && !IsNil(o.CESManagementFunction) {
		return true
	}

	return false
}

// SetCESManagementFunction gets a reference to the given CESManagementFunctionSingle and assigns it to the CESManagementFunction field.
func (o *ManagedElementSingle) SetCESManagementFunction(v CESManagementFunctionSingle) {
	o.CESManagementFunction = &v
}

// GetConfigurable5QISet returns the Configurable5QISet field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetConfigurable5QISet() []Configurable5QISetSingle {
	if o == nil || IsNil(o.Configurable5QISet) {
		var ret []Configurable5QISetSingle
		return ret
	}
	return o.Configurable5QISet
}

// GetConfigurable5QISetOk returns a tuple with the Configurable5QISet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetConfigurable5QISetOk() ([]Configurable5QISetSingle, bool) {
	if o == nil || IsNil(o.Configurable5QISet) {
		return nil, false
	}
	return o.Configurable5QISet, true
}

// HasConfigurable5QISet returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasConfigurable5QISet() bool {
	if o != nil && !IsNil(o.Configurable5QISet) {
		return true
	}

	return false
}

// SetConfigurable5QISet gets a reference to the given []Configurable5QISetSingle and assigns it to the Configurable5QISet field.
func (o *ManagedElementSingle) SetConfigurable5QISet(v []Configurable5QISetSingle) {
	o.Configurable5QISet = v
}

// GetDynamic5QISet returns the Dynamic5QISet field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetDynamic5QISet() []Dynamic5QISetSingle {
	if o == nil || IsNil(o.Dynamic5QISet) {
		var ret []Dynamic5QISetSingle
		return ret
	}
	return o.Dynamic5QISet
}

// GetDynamic5QISetOk returns a tuple with the Dynamic5QISet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetDynamic5QISetOk() ([]Dynamic5QISetSingle, bool) {
	if o == nil || IsNil(o.Dynamic5QISet) {
		return nil, false
	}
	return o.Dynamic5QISet, true
}

// HasDynamic5QISet returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasDynamic5QISet() bool {
	if o != nil && !IsNil(o.Dynamic5QISet) {
		return true
	}

	return false
}

// SetDynamic5QISet gets a reference to the given []Dynamic5QISetSingle and assigns it to the Dynamic5QISet field.
func (o *ManagedElementSingle) SetDynamic5QISet(v []Dynamic5QISetSingle) {
	o.Dynamic5QISet = v
}

func (o ManagedElementSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedElementSingle) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.GnbDuFunction) {
		toSerialize["GnbDuFunction"] = o.GnbDuFunction
	}
	if !IsNil(o.GnbCuUpFunction) {
		toSerialize["GnbCuUpFunction"] = o.GnbCuUpFunction
	}
	if !IsNil(o.GnbCuCpFunction) {
		toSerialize["GnbCuCpFunction"] = o.GnbCuCpFunction
	}
	if !IsNil(o.DESManagementFunction) {
		toSerialize["DESManagementFunction"] = o.DESManagementFunction
	}
	if !IsNil(o.DRACHOptimizationFunction) {
		toSerialize["DRACHOptimizationFunction"] = o.DRACHOptimizationFunction
	}
	if !IsNil(o.DMROFunction) {
		toSerialize["DMROFunction"] = o.DMROFunction
	}
	if !IsNil(o.DLBOFunction) {
		toSerialize["DLBOFunction"] = o.DLBOFunction
	}
	if !IsNil(o.DPCIConfigurationFunction) {
		toSerialize["DPCIConfigurationFunction"] = o.DPCIConfigurationFunction
	}
	if !IsNil(o.CPCIConfigurationFunction) {
		toSerialize["CPCIConfigurationFunction"] = o.CPCIConfigurationFunction
	}
	if !IsNil(o.CESManagementFunction) {
		toSerialize["CESManagementFunction"] = o.CESManagementFunction
	}
	if !IsNil(o.Configurable5QISet) {
		toSerialize["Configurable5QISet"] = o.Configurable5QISet
	}
	if !IsNil(o.Dynamic5QISet) {
		toSerialize["Dynamic5QISet"] = o.Dynamic5QISet
	}
	return toSerialize, nil
}

type NullableManagedElementSingle struct {
	value *ManagedElementSingle
	isSet bool
}

func (v NullableManagedElementSingle) Get() *ManagedElementSingle {
	return v.value
}

func (v *NullableManagedElementSingle) Set(val *ManagedElementSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedElementSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedElementSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedElementSingle(val *ManagedElementSingle) *NullableManagedElementSingle {
	return &NullableManagedElementSingle{value: val, isSet: true}
}

func (v NullableManagedElementSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedElementSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// checks if the SubNetworkSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubNetworkSingle{}

// SubNetworkSingle struct for SubNetworkSingle
type SubNetworkSingle struct {
	Top
	Attributes                             *SubNetworkSingleAllOfAttributes               `json:"attributes,omitempty"`
	ManagementNode                         []ManagementNodeSingle                         `json:"ManagementNode,omitempty"`
	MnsAgent                               []MnsAgentSingle                               `json:"MnsAgent,omitempty"`
	MeContext                              []MeContextSingle                              `json:"MeContext,omitempty"`
	PerfMetricJob                          []PerfMetricJobSingle                          `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor                       []ThresholdMonitorSingle                       `json:"ThresholdMonitor,omitempty"`
	TraceJob                               []TraceJobSingle                               `json:"TraceJob,omitempty"`
	ManagementDataCollection               []ManagementDataCollectionSingle               `json:"ManagementDataCollection,omitempty"`
	NtfSubscriptionControl                 []NtfSubscriptionControlSingle                 `json:"NtfSubscriptionControl,omitempty"`
	AlarmList                              *AlarmListSingle                               `json:"AlarmList,omitempty"`
	FileDownloadJob                        []FileDownloadJobSingle                        `json:"FileDownloadJob,omitempty"`
	Files                                  []FilesSingle                                  `json:"Files,omitempty"`
	MnsRegistry                            *MnsRegistrySingle                             `json:"MnsRegistry,omitempty"`
	SubNetwork                             []SubNetworkSingle                             `json:"SubNetwork,omitempty"`
	NetworkSlice                           []NetworkSliceSingle                           `json:"NetworkSlice,omitempty"`
	NetworkSliceSubnet                     []NetworkSliceSubnetSingle                     `json:"NetworkSliceSubnet,omitempty"`
	EPTransport                            []EPTransportSingle                            `json:"EP_Transport,omitempty"`
	NetworkSliceSubnetProviderCapabilities []NetworkSliceSubnetProviderCapabilitiesSingle `json:"NetworkSliceSubnetProviderCapabilities,omitempty"`
	FeasibilityCheckAndReservationJob      []FeasibilityCheckAndReservationJobSingle      `json:"FeasibilityCheckAndReservationJob,omitempty"`
}

// NewSubNetworkSingle instantiates a new SubNetworkSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingle(id NullableString) *SubNetworkSingle {
	this := SubNetworkSingle{}
	this.Id = id
	return &this
}

// NewSubNetworkSingleWithDefaults instantiates a new SubNetworkSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingleWithDefaults() *SubNetworkSingle {
	this := SubNetworkSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetAttributes() SubNetworkSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubNetworkSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetAttributesOk() (*SubNetworkSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubNetworkSingleAllOfAttributes and assigns it to the Attributes field.
func (o *SubNetworkSingle) SetAttributes(v SubNetworkSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetManagementNode returns the ManagementNode field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetManagementNode() []ManagementNodeSingle {
	if o == nil || IsNil(o.ManagementNode) {
		var ret []ManagementNodeSingle
		return ret
	}
	return o.ManagementNode
}

// GetManagementNodeOk returns a tuple with the ManagementNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetManagementNodeOk() ([]ManagementNodeSingle, bool) {
	if o == nil || IsNil(o.ManagementNode) {
		return nil, false
	}
	return o.ManagementNode, true
}

// HasManagementNode returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasManagementNode() bool {
	if o != nil && !IsNil(o.ManagementNode) {
		return true
	}

	return false
}

// SetManagementNode gets a reference to the given []ManagementNodeSingle and assigns it to the ManagementNode field.
func (o *SubNetworkSingle) SetManagementNode(v []ManagementNodeSingle) {
	o.ManagementNode = v
}

// GetMnsAgent returns the MnsAgent field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetMnsAgent() []MnsAgentSingle {
	if o == nil || IsNil(o.MnsAgent) {
		var ret []MnsAgentSingle
		return ret
	}
	return o.MnsAgent
}

// GetMnsAgentOk returns a tuple with the MnsAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetMnsAgentOk() ([]MnsAgentSingle, bool) {
	if o == nil || IsNil(o.MnsAgent) {
		return nil, false
	}
	return o.MnsAgent, true
}

// HasMnsAgent returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasMnsAgent() bool {
	if o != nil && !IsNil(o.MnsAgent) {
		return true
	}

	return false
}

// SetMnsAgent gets a reference to the given []MnsAgentSingle and assigns it to the MnsAgent field.
func (o *SubNetworkSingle) SetMnsAgent(v []MnsAgentSingle) {
	o.MnsAgent = v
}

// GetMeContext returns the MeContext field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetMeContext() []MeContextSingle {
	if o == nil || IsNil(o.MeContext) {
		var ret []MeContextSingle
		return ret
	}
	return o.MeContext
}

// GetMeContextOk returns a tuple with the MeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetMeContextOk() ([]MeContextSingle, bool) {
	if o == nil || IsNil(o.MeContext) {
		return nil, false
	}
	return o.MeContext, true
}

// HasMeContext returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasMeContext() bool {
	if o != nil && !IsNil(o.MeContext) {
		return true
	}

	return false
}

// SetMeContext gets a reference to the given []MeContextSingle and assigns it to the MeContext field.
func (o *SubNetworkSingle) SetMeContext(v []MeContextSingle) {
	o.MeContext = v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || IsNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || IsNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasPerfMetricJob() bool {
	if o != nil && !IsNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *SubNetworkSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || IsNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || IsNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasThresholdMonitor() bool {
	if o != nil && !IsNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *SubNetworkSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || IsNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || IsNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasTraceJob() bool {
	if o != nil && !IsNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *SubNetworkSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetManagementDataCollection returns the ManagementDataCollection field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetManagementDataCollection() []ManagementDataCollectionSingle {
	if o == nil || IsNil(o.ManagementDataCollection) {
		var ret []ManagementDataCollectionSingle
		return ret
	}
	return o.ManagementDataCollection
}

// GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetManagementDataCollectionOk() ([]ManagementDataCollectionSingle, bool) {
	if o == nil || IsNil(o.ManagementDataCollection) {
		return nil, false
	}
	return o.ManagementDataCollection, true
}

// HasManagementDataCollection returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasManagementDataCollection() bool {
	if o != nil && !IsNil(o.ManagementDataCollection) {
		return true
	}

	return false
}

// SetManagementDataCollection gets a reference to the given []ManagementDataCollectionSingle and assigns it to the ManagementDataCollection field.
func (o *SubNetworkSingle) SetManagementDataCollection(v []ManagementDataCollectionSingle) {
	o.ManagementDataCollection = v
}

// GetNtfSubscriptionControl returns the NtfSubscriptionControl field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle {
	if o == nil || IsNil(o.NtfSubscriptionControl) {
		var ret []NtfSubscriptionControlSingle
		return ret
	}
	return o.NtfSubscriptionControl
}

// GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetNtfSubscriptionControlOk() ([]NtfSubscriptionControlSingle, bool) {
	if o == nil || IsNil(o.NtfSubscriptionControl) {
		return nil, false
	}
	return o.NtfSubscriptionControl, true
}

// HasNtfSubscriptionControl returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasNtfSubscriptionControl() bool {
	if o != nil && !IsNil(o.NtfSubscriptionControl) {
		return true
	}

	return false
}

// SetNtfSubscriptionControl gets a reference to the given []NtfSubscriptionControlSingle and assigns it to the NtfSubscriptionControl field.
func (o *SubNetworkSingle) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle) {
	o.NtfSubscriptionControl = v
}

// GetAlarmList returns the AlarmList field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetAlarmList() AlarmListSingle {
	if o == nil || IsNil(o.AlarmList) {
		var ret AlarmListSingle
		return ret
	}
	return *o.AlarmList
}

// GetAlarmListOk returns a tuple with the AlarmList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetAlarmListOk() (*AlarmListSingle, bool) {
	if o == nil || IsNil(o.AlarmList) {
		return nil, false
	}
	return o.AlarmList, true
}

// HasAlarmList returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasAlarmList() bool {
	if o != nil && !IsNil(o.AlarmList) {
		return true
	}

	return false
}

// SetAlarmList gets a reference to the given AlarmListSingle and assigns it to the AlarmList field.
func (o *SubNetworkSingle) SetAlarmList(v AlarmListSingle) {
	o.AlarmList = &v
}

// GetFileDownloadJob returns the FileDownloadJob field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetFileDownloadJob() []FileDownloadJobSingle {
	if o == nil || IsNil(o.FileDownloadJob) {
		var ret []FileDownloadJobSingle
		return ret
	}
	return o.FileDownloadJob
}

// GetFileDownloadJobOk returns a tuple with the FileDownloadJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetFileDownloadJobOk() ([]FileDownloadJobSingle, bool) {
	if o == nil || IsNil(o.FileDownloadJob) {
		return nil, false
	}
	return o.FileDownloadJob, true
}

// HasFileDownloadJob returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasFileDownloadJob() bool {
	if o != nil && !IsNil(o.FileDownloadJob) {
		return true
	}

	return false
}

// SetFileDownloadJob gets a reference to the given []FileDownloadJobSingle and assigns it to the FileDownloadJob field.
func (o *SubNetworkSingle) SetFileDownloadJob(v []FileDownloadJobSingle) {
	o.FileDownloadJob = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetFiles() []FilesSingle {
	if o == nil || IsNil(o.Files) {
		var ret []FilesSingle
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetFilesOk() ([]FilesSingle, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FilesSingle and assigns it to the Files field.
func (o *SubNetworkSingle) SetFiles(v []FilesSingle) {
	o.Files = v
}

// GetMnsRegistry returns the MnsRegistry field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetMnsRegistry() MnsRegistrySingle {
	if o == nil || IsNil(o.MnsRegistry) {
		var ret MnsRegistrySingle
		return ret
	}
	return *o.MnsRegistry
}

// GetMnsRegistryOk returns a tuple with the MnsRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetMnsRegistryOk() (*MnsRegistrySingle, bool) {
	if o == nil || IsNil(o.MnsRegistry) {
		return nil, false
	}
	return o.MnsRegistry, true
}

// HasMnsRegistry returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasMnsRegistry() bool {
	if o != nil && !IsNil(o.MnsRegistry) {
		return true
	}

	return false
}

// SetMnsRegistry gets a reference to the given MnsRegistrySingle and assigns it to the MnsRegistry field.
func (o *SubNetworkSingle) SetMnsRegistry(v MnsRegistrySingle) {
	o.MnsRegistry = &v
}

// GetSubNetwork returns the SubNetwork field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetSubNetwork() []SubNetworkSingle {
	if o == nil || IsNil(o.SubNetwork) {
		var ret []SubNetworkSingle
		return ret
	}
	return o.SubNetwork
}

// GetSubNetworkOk returns a tuple with the SubNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetSubNetworkOk() ([]SubNetworkSingle, bool) {
	if o == nil || IsNil(o.SubNetwork) {
		return nil, false
	}
	return o.SubNetwork, true
}

// HasSubNetwork returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasSubNetwork() bool {
	if o != nil && !IsNil(o.SubNetwork) {
		return true
	}

	return false
}

// SetSubNetwork gets a reference to the given []SubNetworkSingle and assigns it to the SubNetwork field.
func (o *SubNetworkSingle) SetSubNetwork(v []SubNetworkSingle) {
	o.SubNetwork = v
}

// GetNetworkSlice returns the NetworkSlice field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetNetworkSlice() []NetworkSliceSingle {
	if o == nil || IsNil(o.NetworkSlice) {
		var ret []NetworkSliceSingle
		return ret
	}
	return o.NetworkSlice
}

// GetNetworkSliceOk returns a tuple with the NetworkSlice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetNetworkSliceOk() ([]NetworkSliceSingle, bool) {
	if o == nil || IsNil(o.NetworkSlice) {
		return nil, false
	}
	return o.NetworkSlice, true
}

// HasNetworkSlice returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasNetworkSlice() bool {
	if o != nil && !IsNil(o.NetworkSlice) {
		return true
	}

	return false
}

// SetNetworkSlice gets a reference to the given []NetworkSliceSingle and assigns it to the NetworkSlice field.
func (o *SubNetworkSingle) SetNetworkSlice(v []NetworkSliceSingle) {
	o.NetworkSlice = v
}

// GetNetworkSliceSubnet returns the NetworkSliceSubnet field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetNetworkSliceSubnet() []NetworkSliceSubnetSingle {
	if o == nil || IsNil(o.NetworkSliceSubnet) {
		var ret []NetworkSliceSubnetSingle
		return ret
	}
	return o.NetworkSliceSubnet
}

// GetNetworkSliceSubnetOk returns a tuple with the NetworkSliceSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetNetworkSliceSubnetOk() ([]NetworkSliceSubnetSingle, bool) {
	if o == nil || IsNil(o.NetworkSliceSubnet) {
		return nil, false
	}
	return o.NetworkSliceSubnet, true
}

// HasNetworkSliceSubnet returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasNetworkSliceSubnet() bool {
	if o != nil && !IsNil(o.NetworkSliceSubnet) {
		return true
	}

	return false
}

// SetNetworkSliceSubnet gets a reference to the given []NetworkSliceSubnetSingle and assigns it to the NetworkSliceSubnet field.
func (o *SubNetworkSingle) SetNetworkSliceSubnet(v []NetworkSliceSubnetSingle) {
	o.NetworkSliceSubnet = v
}

// GetEPTransport returns the EPTransport field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetEPTransport() []EPTransportSingle {
	if o == nil || IsNil(o.EPTransport) {
		var ret []EPTransportSingle
		return ret
	}
	return o.EPTransport
}

// GetEPTransportOk returns a tuple with the EPTransport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetEPTransportOk() ([]EPTransportSingle, bool) {
	if o == nil || IsNil(o.EPTransport) {
		return nil, false
	}
	return o.EPTransport, true
}

// HasEPTransport returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasEPTransport() bool {
	if o != nil && !IsNil(o.EPTransport) {
		return true
	}

	return false
}

// SetEPTransport gets a reference to the given []EPTransportSingle and assigns it to the EPTransport field.
func (o *SubNetworkSingle) SetEPTransport(v []EPTransportSingle) {
	o.EPTransport = v
}

// GetNetworkSliceSubnetProviderCapabilities returns the NetworkSliceSubnetProviderCapabilities field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetNetworkSliceSubnetProviderCapabilities() []NetworkSliceSubnetProviderCapabilitiesSingle {
	if o == nil || IsNil(o.NetworkSliceSubnetProviderCapabilities) {
		var ret []NetworkSliceSubnetProviderCapabilitiesSingle
		return ret
	}
	return o.NetworkSliceSubnetProviderCapabilities
}

// GetNetworkSliceSubnetProviderCapabilitiesOk returns a tuple with the NetworkSliceSubnetProviderCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetNetworkSliceSubnetProviderCapabilitiesOk() ([]NetworkSliceSubnetProviderCapabilitiesSingle, bool) {
	if o == nil || IsNil(o.NetworkSliceSubnetProviderCapabilities) {
		return nil, false
	}
	return o.NetworkSliceSubnetProviderCapabilities, true
}

// HasNetworkSliceSubnetProviderCapabilities returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasNetworkSliceSubnetProviderCapabilities() bool {
	if o != nil && !IsNil(o.NetworkSliceSubnetProviderCapabilities) {
		return true
	}

	return false
}

// SetNetworkSliceSubnetProviderCapabilities gets a reference to the given []NetworkSliceSubnetProviderCapabilitiesSingle and assigns it to the NetworkSliceSubnetProviderCapabilities field.
func (o *SubNetworkSingle) SetNetworkSliceSubnetProviderCapabilities(v []NetworkSliceSubnetProviderCapabilitiesSingle) {
	o.NetworkSliceSubnetProviderCapabilities = v
}

// GetFeasibilityCheckAndReservationJob returns the FeasibilityCheckAndReservationJob field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetFeasibilityCheckAndReservationJob() []FeasibilityCheckAndReservationJobSingle {
	if o == nil || IsNil(o.FeasibilityCheckAndReservationJob) {
		var ret []FeasibilityCheckAndReservationJobSingle
		return ret
	}
	return o.FeasibilityCheckAndReservationJob
}

// GetFeasibilityCheckAndReservationJobOk returns a tuple with the FeasibilityCheckAndReservationJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetFeasibilityCheckAndReservationJobOk() ([]FeasibilityCheckAndReservationJobSingle, bool) {
	if o == nil || IsNil(o.FeasibilityCheckAndReservationJob) {
		return nil, false
	}
	return o.FeasibilityCheckAndReservationJob, true
}

// HasFeasibilityCheckAndReservationJob returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasFeasibilityCheckAndReservationJob() bool {
	if o != nil && !IsNil(o.FeasibilityCheckAndReservationJob) {
		return true
	}

	return false
}

// SetFeasibilityCheckAndReservationJob gets a reference to the given []FeasibilityCheckAndReservationJobSingle and assigns it to the FeasibilityCheckAndReservationJob field.
func (o *SubNetworkSingle) SetFeasibilityCheckAndReservationJob(v []FeasibilityCheckAndReservationJobSingle) {
	o.FeasibilityCheckAndReservationJob = v
}

func (o SubNetworkSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubNetworkSingle) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SubNetwork) {
		toSerialize["SubNetwork"] = o.SubNetwork
	}
	if !IsNil(o.NetworkSlice) {
		toSerialize["NetworkSlice"] = o.NetworkSlice
	}
	if !IsNil(o.NetworkSliceSubnet) {
		toSerialize["NetworkSliceSubnet"] = o.NetworkSliceSubnet
	}
	if !IsNil(o.EPTransport) {
		toSerialize["EP_Transport"] = o.EPTransport
	}
	if !IsNil(o.NetworkSliceSubnetProviderCapabilities) {
		toSerialize["NetworkSliceSubnetProviderCapabilities"] = o.NetworkSliceSubnetProviderCapabilities
	}
	if !IsNil(o.FeasibilityCheckAndReservationJob) {
		toSerialize["FeasibilityCheckAndReservationJob"] = o.FeasibilityCheckAndReservationJob
	}
	return toSerialize, nil
}

type NullableSubNetworkSingle struct {
	value *SubNetworkSingle
	isSet bool
}

func (v NullableSubNetworkSingle) Get() *SubNetworkSingle {
	return v.value
}

func (v *NullableSubNetworkSingle) Set(val *SubNetworkSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingle(val *SubNetworkSingle) *NullableSubNetworkSingle {
	return &NullableSubNetworkSingle{value: val, isSet: true}
}

func (v NullableSubNetworkSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

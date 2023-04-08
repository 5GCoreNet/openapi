/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the PcfFunctionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcfFunctionSingle{}

// PcfFunctionSingle struct for PcfFunctionSingle
type PcfFunctionSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`
	ManagedNFService []ManagedNFServiceSingle `json:"ManagedNFService,omitempty"`
	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`
	EPN5 []EPN5Single `json:"EP_N5,omitempty"`
	EPN7 []EPN7Single `json:"EP_N7,omitempty"`
	EPN15 []EPN15Single `json:"EP_N15,omitempty"`
	EPN16 []EPN16Single `json:"EP_N16,omitempty"`
	EPRx []EPRxSingle `json:"EP_Rx,omitempty"`
	PredefinedPccRuleSet *PredefinedPccRuleSetSingle `json:"PredefinedPccRuleSet,omitempty"`
}

// NewPcfFunctionSingle instantiates a new PcfFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfFunctionSingle(id NullableString) *PcfFunctionSingle {
	this := PcfFunctionSingle{}
	this.Id = id
	return &this
}

// NewPcfFunctionSingleWithDefaults instantiates a new PcfFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfFunctionSingleWithDefaults() *PcfFunctionSingle {
	this := PcfFunctionSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PcfFunctionSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PcfFunctionSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *PcfFunctionSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *PcfFunctionSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *PcfFunctionSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *PcfFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *PcfFunctionSingle) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || isNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || isNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasPerfMetricJob() bool {
	if o != nil && !isNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *PcfFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || isNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || isNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasThresholdMonitor() bool {
	if o != nil && !isNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *PcfFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetManagedNFService returns the ManagedNFService field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle {
	if o == nil || isNil(o.ManagedNFService) {
		var ret []ManagedNFServiceSingle
		return ret
	}
	return o.ManagedNFService
}

// GetManagedNFServiceOk returns a tuple with the ManagedNFService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetManagedNFServiceOk() ([]ManagedNFServiceSingle, bool) {
	if o == nil || isNil(o.ManagedNFService) {
		return nil, false
	}
	return o.ManagedNFService, true
}

// HasManagedNFService returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasManagedNFService() bool {
	if o != nil && !isNil(o.ManagedNFService) {
		return true
	}

	return false
}

// SetManagedNFService gets a reference to the given []ManagedNFServiceSingle and assigns it to the ManagedNFService field.
func (o *PcfFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle) {
	o.ManagedNFService = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || isNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || isNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasTraceJob() bool {
	if o != nil && !isNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *PcfFunctionSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetEPN5 returns the EPN5 field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetEPN5() []EPN5Single {
	if o == nil || isNil(o.EPN5) {
		var ret []EPN5Single
		return ret
	}
	return o.EPN5
}

// GetEPN5Ok returns a tuple with the EPN5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetEPN5Ok() ([]EPN5Single, bool) {
	if o == nil || isNil(o.EPN5) {
		return nil, false
	}
	return o.EPN5, true
}

// HasEPN5 returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasEPN5() bool {
	if o != nil && !isNil(o.EPN5) {
		return true
	}

	return false
}

// SetEPN5 gets a reference to the given []EPN5Single and assigns it to the EPN5 field.
func (o *PcfFunctionSingle) SetEPN5(v []EPN5Single) {
	o.EPN5 = v
}

// GetEPN7 returns the EPN7 field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetEPN7() []EPN7Single {
	if o == nil || isNil(o.EPN7) {
		var ret []EPN7Single
		return ret
	}
	return o.EPN7
}

// GetEPN7Ok returns a tuple with the EPN7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetEPN7Ok() ([]EPN7Single, bool) {
	if o == nil || isNil(o.EPN7) {
		return nil, false
	}
	return o.EPN7, true
}

// HasEPN7 returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasEPN7() bool {
	if o != nil && !isNil(o.EPN7) {
		return true
	}

	return false
}

// SetEPN7 gets a reference to the given []EPN7Single and assigns it to the EPN7 field.
func (o *PcfFunctionSingle) SetEPN7(v []EPN7Single) {
	o.EPN7 = v
}

// GetEPN15 returns the EPN15 field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetEPN15() []EPN15Single {
	if o == nil || isNil(o.EPN15) {
		var ret []EPN15Single
		return ret
	}
	return o.EPN15
}

// GetEPN15Ok returns a tuple with the EPN15 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetEPN15Ok() ([]EPN15Single, bool) {
	if o == nil || isNil(o.EPN15) {
		return nil, false
	}
	return o.EPN15, true
}

// HasEPN15 returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasEPN15() bool {
	if o != nil && !isNil(o.EPN15) {
		return true
	}

	return false
}

// SetEPN15 gets a reference to the given []EPN15Single and assigns it to the EPN15 field.
func (o *PcfFunctionSingle) SetEPN15(v []EPN15Single) {
	o.EPN15 = v
}

// GetEPN16 returns the EPN16 field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetEPN16() []EPN16Single {
	if o == nil || isNil(o.EPN16) {
		var ret []EPN16Single
		return ret
	}
	return o.EPN16
}

// GetEPN16Ok returns a tuple with the EPN16 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetEPN16Ok() ([]EPN16Single, bool) {
	if o == nil || isNil(o.EPN16) {
		return nil, false
	}
	return o.EPN16, true
}

// HasEPN16 returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasEPN16() bool {
	if o != nil && !isNil(o.EPN16) {
		return true
	}

	return false
}

// SetEPN16 gets a reference to the given []EPN16Single and assigns it to the EPN16 field.
func (o *PcfFunctionSingle) SetEPN16(v []EPN16Single) {
	o.EPN16 = v
}

// GetEPRx returns the EPRx field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetEPRx() []EPRxSingle {
	if o == nil || isNil(o.EPRx) {
		var ret []EPRxSingle
		return ret
	}
	return o.EPRx
}

// GetEPRxOk returns a tuple with the EPRx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetEPRxOk() ([]EPRxSingle, bool) {
	if o == nil || isNil(o.EPRx) {
		return nil, false
	}
	return o.EPRx, true
}

// HasEPRx returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasEPRx() bool {
	if o != nil && !isNil(o.EPRx) {
		return true
	}

	return false
}

// SetEPRx gets a reference to the given []EPRxSingle and assigns it to the EPRx field.
func (o *PcfFunctionSingle) SetEPRx(v []EPRxSingle) {
	o.EPRx = v
}

// GetPredefinedPccRuleSet returns the PredefinedPccRuleSet field value if set, zero value otherwise.
func (o *PcfFunctionSingle) GetPredefinedPccRuleSet() PredefinedPccRuleSetSingle {
	if o == nil || isNil(o.PredefinedPccRuleSet) {
		var ret PredefinedPccRuleSetSingle
		return ret
	}
	return *o.PredefinedPccRuleSet
}

// GetPredefinedPccRuleSetOk returns a tuple with the PredefinedPccRuleSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingle) GetPredefinedPccRuleSetOk() (*PredefinedPccRuleSetSingle, bool) {
	if o == nil || isNil(o.PredefinedPccRuleSet) {
		return nil, false
	}
	return o.PredefinedPccRuleSet, true
}

// HasPredefinedPccRuleSet returns a boolean if a field has been set.
func (o *PcfFunctionSingle) HasPredefinedPccRuleSet() bool {
	if o != nil && !isNil(o.PredefinedPccRuleSet) {
		return true
	}

	return false
}

// SetPredefinedPccRuleSet gets a reference to the given PredefinedPccRuleSetSingle and assigns it to the PredefinedPccRuleSet field.
func (o *PcfFunctionSingle) SetPredefinedPccRuleSet(v PredefinedPccRuleSetSingle) {
	o.PredefinedPccRuleSet = &v
}

func (o PcfFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcfFunctionSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if !isNil(o.ObjectClass) {
		toSerialize["objectClass"] = o.ObjectClass
	}
	if !isNil(o.ObjectInstance) {
		toSerialize["objectInstance"] = o.ObjectInstance
	}
	if !isNil(o.VsDataContainer) {
		toSerialize["VsDataContainer"] = o.VsDataContainer
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.PerfMetricJob) {
		toSerialize["PerfMetricJob"] = o.PerfMetricJob
	}
	if !isNil(o.ThresholdMonitor) {
		toSerialize["ThresholdMonitor"] = o.ThresholdMonitor
	}
	if !isNil(o.ManagedNFService) {
		toSerialize["ManagedNFService"] = o.ManagedNFService
	}
	if !isNil(o.TraceJob) {
		toSerialize["TraceJob"] = o.TraceJob
	}
	if !isNil(o.EPN5) {
		toSerialize["EP_N5"] = o.EPN5
	}
	if !isNil(o.EPN7) {
		toSerialize["EP_N7"] = o.EPN7
	}
	if !isNil(o.EPN15) {
		toSerialize["EP_N15"] = o.EPN15
	}
	if !isNil(o.EPN16) {
		toSerialize["EP_N16"] = o.EPN16
	}
	if !isNil(o.EPRx) {
		toSerialize["EP_Rx"] = o.EPRx
	}
	if !isNil(o.PredefinedPccRuleSet) {
		toSerialize["PredefinedPccRuleSet"] = o.PredefinedPccRuleSet
	}
	return toSerialize, nil
}

type NullablePcfFunctionSingle struct {
	value *PcfFunctionSingle
	isSet bool
}

func (v NullablePcfFunctionSingle) Get() *PcfFunctionSingle {
	return v.value
}

func (v *NullablePcfFunctionSingle) Set(val *PcfFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfFunctionSingle(val *PcfFunctionSingle) *NullablePcfFunctionSingle {
	return &NullablePcfFunctionSingle{value: val, isSet: true}
}

func (v NullablePcfFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



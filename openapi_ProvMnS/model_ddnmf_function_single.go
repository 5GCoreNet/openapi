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

// checks if the DDNMFFunctionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DDNMFFunctionSingle{}

// DDNMFFunctionSingle struct for DDNMFFunctionSingle
type DDNMFFunctionSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`
	ManagedNFService []ManagedNFServiceSingle `json:"ManagedNFService,omitempty"`
	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`
	EPNpc4 []EPNpc4Single `json:"EP_Npc4,omitempty"`
	EPNpc6 []EPNpc6Single `json:"EP_Npc6,omitempty"`
	EPNpc7 []EPNpc7Single `json:"EP_Npc7,omitempty"`
	EPNpc8 []EPNpc8Single `json:"EP_Npc8,omitempty"`
}

// NewDDNMFFunctionSingle instantiates a new DDNMFFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDDNMFFunctionSingle(id NullableString) *DDNMFFunctionSingle {
	this := DDNMFFunctionSingle{}
	this.Id = id
	return &this
}

// NewDDNMFFunctionSingleWithDefaults instantiates a new DDNMFFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDDNMFFunctionSingleWithDefaults() *DDNMFFunctionSingle {
	this := DDNMFFunctionSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DDNMFFunctionSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DDNMFFunctionSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *DDNMFFunctionSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *DDNMFFunctionSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *DDNMFFunctionSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *DDNMFFunctionSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *DDNMFFunctionSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *DDNMFFunctionSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *DDNMFFunctionSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *DDNMFFunctionSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *DDNMFFunctionSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *DDNMFFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DDNMFFunctionSingle) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DDNMFFunctionSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *DDNMFFunctionSingle) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *DDNMFFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || isNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || isNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *DDNMFFunctionSingle) HasPerfMetricJob() bool {
	if o != nil && !isNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *DDNMFFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *DDNMFFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || isNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || isNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *DDNMFFunctionSingle) HasThresholdMonitor() bool {
	if o != nil && !isNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *DDNMFFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetManagedNFService returns the ManagedNFService field value if set, zero value otherwise.
func (o *DDNMFFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle {
	if o == nil || isNil(o.ManagedNFService) {
		var ret []ManagedNFServiceSingle
		return ret
	}
	return o.ManagedNFService
}

// GetManagedNFServiceOk returns a tuple with the ManagedNFService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingle) GetManagedNFServiceOk() ([]ManagedNFServiceSingle, bool) {
	if o == nil || isNil(o.ManagedNFService) {
		return nil, false
	}
	return o.ManagedNFService, true
}

// HasManagedNFService returns a boolean if a field has been set.
func (o *DDNMFFunctionSingle) HasManagedNFService() bool {
	if o != nil && !isNil(o.ManagedNFService) {
		return true
	}

	return false
}

// SetManagedNFService gets a reference to the given []ManagedNFServiceSingle and assigns it to the ManagedNFService field.
func (o *DDNMFFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle) {
	o.ManagedNFService = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *DDNMFFunctionSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || isNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || isNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *DDNMFFunctionSingle) HasTraceJob() bool {
	if o != nil && !isNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *DDNMFFunctionSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetEPNpc4 returns the EPNpc4 field value if set, zero value otherwise.
func (o *DDNMFFunctionSingle) GetEPNpc4() []EPNpc4Single {
	if o == nil || isNil(o.EPNpc4) {
		var ret []EPNpc4Single
		return ret
	}
	return o.EPNpc4
}

// GetEPNpc4Ok returns a tuple with the EPNpc4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingle) GetEPNpc4Ok() ([]EPNpc4Single, bool) {
	if o == nil || isNil(o.EPNpc4) {
		return nil, false
	}
	return o.EPNpc4, true
}

// HasEPNpc4 returns a boolean if a field has been set.
func (o *DDNMFFunctionSingle) HasEPNpc4() bool {
	if o != nil && !isNil(o.EPNpc4) {
		return true
	}

	return false
}

// SetEPNpc4 gets a reference to the given []EPNpc4Single and assigns it to the EPNpc4 field.
func (o *DDNMFFunctionSingle) SetEPNpc4(v []EPNpc4Single) {
	o.EPNpc4 = v
}

// GetEPNpc6 returns the EPNpc6 field value if set, zero value otherwise.
func (o *DDNMFFunctionSingle) GetEPNpc6() []EPNpc6Single {
	if o == nil || isNil(o.EPNpc6) {
		var ret []EPNpc6Single
		return ret
	}
	return o.EPNpc6
}

// GetEPNpc6Ok returns a tuple with the EPNpc6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingle) GetEPNpc6Ok() ([]EPNpc6Single, bool) {
	if o == nil || isNil(o.EPNpc6) {
		return nil, false
	}
	return o.EPNpc6, true
}

// HasEPNpc6 returns a boolean if a field has been set.
func (o *DDNMFFunctionSingle) HasEPNpc6() bool {
	if o != nil && !isNil(o.EPNpc6) {
		return true
	}

	return false
}

// SetEPNpc6 gets a reference to the given []EPNpc6Single and assigns it to the EPNpc6 field.
func (o *DDNMFFunctionSingle) SetEPNpc6(v []EPNpc6Single) {
	o.EPNpc6 = v
}

// GetEPNpc7 returns the EPNpc7 field value if set, zero value otherwise.
func (o *DDNMFFunctionSingle) GetEPNpc7() []EPNpc7Single {
	if o == nil || isNil(o.EPNpc7) {
		var ret []EPNpc7Single
		return ret
	}
	return o.EPNpc7
}

// GetEPNpc7Ok returns a tuple with the EPNpc7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingle) GetEPNpc7Ok() ([]EPNpc7Single, bool) {
	if o == nil || isNil(o.EPNpc7) {
		return nil, false
	}
	return o.EPNpc7, true
}

// HasEPNpc7 returns a boolean if a field has been set.
func (o *DDNMFFunctionSingle) HasEPNpc7() bool {
	if o != nil && !isNil(o.EPNpc7) {
		return true
	}

	return false
}

// SetEPNpc7 gets a reference to the given []EPNpc7Single and assigns it to the EPNpc7 field.
func (o *DDNMFFunctionSingle) SetEPNpc7(v []EPNpc7Single) {
	o.EPNpc7 = v
}

// GetEPNpc8 returns the EPNpc8 field value if set, zero value otherwise.
func (o *DDNMFFunctionSingle) GetEPNpc8() []EPNpc8Single {
	if o == nil || isNil(o.EPNpc8) {
		var ret []EPNpc8Single
		return ret
	}
	return o.EPNpc8
}

// GetEPNpc8Ok returns a tuple with the EPNpc8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingle) GetEPNpc8Ok() ([]EPNpc8Single, bool) {
	if o == nil || isNil(o.EPNpc8) {
		return nil, false
	}
	return o.EPNpc8, true
}

// HasEPNpc8 returns a boolean if a field has been set.
func (o *DDNMFFunctionSingle) HasEPNpc8() bool {
	if o != nil && !isNil(o.EPNpc8) {
		return true
	}

	return false
}

// SetEPNpc8 gets a reference to the given []EPNpc8Single and assigns it to the EPNpc8 field.
func (o *DDNMFFunctionSingle) SetEPNpc8(v []EPNpc8Single) {
	o.EPNpc8 = v
}

func (o DDNMFFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DDNMFFunctionSingle) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.EPNpc4) {
		toSerialize["EP_Npc4"] = o.EPNpc4
	}
	if !isNil(o.EPNpc6) {
		toSerialize["EP_Npc6"] = o.EPNpc6
	}
	if !isNil(o.EPNpc7) {
		toSerialize["EP_Npc7"] = o.EPNpc7
	}
	if !isNil(o.EPNpc8) {
		toSerialize["EP_Npc8"] = o.EPNpc8
	}
	return toSerialize, nil
}

type NullableDDNMFFunctionSingle struct {
	value *DDNMFFunctionSingle
	isSet bool
}

func (v NullableDDNMFFunctionSingle) Get() *DDNMFFunctionSingle {
	return v.value
}

func (v *NullableDDNMFFunctionSingle) Set(val *DDNMFFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableDDNMFFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableDDNMFFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDDNMFFunctionSingle(val *DDNMFFunctionSingle) *NullableDDNMFFunctionSingle {
	return &NullableDDNMFFunctionSingle{value: val, isSet: true}
}

func (v NullableDDNMFFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDDNMFFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



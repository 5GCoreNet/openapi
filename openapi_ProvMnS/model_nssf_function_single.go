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

// checks if the NssfFunctionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NssfFunctionSingle{}

// NssfFunctionSingle struct for NssfFunctionSingle
type NssfFunctionSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *NssfFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`
	ManagedNFService []ManagedNFServiceSingle `json:"ManagedNFService,omitempty"`
	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`
	EPN22 []EPN22Single `json:"EP_N22,omitempty"`
	EPN31 []EPN31Single `json:"EP_N31,omitempty"`
}

// NewNssfFunctionSingle instantiates a new NssfFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNssfFunctionSingle(id NullableString) *NssfFunctionSingle {
	this := NssfFunctionSingle{}
	this.Id = id
	return &this
}

// NewNssfFunctionSingleWithDefaults instantiates a new NssfFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNssfFunctionSingleWithDefaults() *NssfFunctionSingle {
	this := NssfFunctionSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NssfFunctionSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NssfFunctionSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *NssfFunctionSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *NssfFunctionSingle) GetObjectClass() string {
	if o == nil || IsNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfFunctionSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *NssfFunctionSingle) HasObjectClass() bool {
	if o != nil && !IsNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *NssfFunctionSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *NssfFunctionSingle) GetObjectInstance() string {
	if o == nil || IsNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfFunctionSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *NssfFunctionSingle) HasObjectInstance() bool {
	if o != nil && !IsNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *NssfFunctionSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *NssfFunctionSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || IsNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfFunctionSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || IsNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *NssfFunctionSingle) HasVsDataContainer() bool {
	if o != nil && !IsNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *NssfFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NssfFunctionSingle) GetAttributes() NssfFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret NssfFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfFunctionSingle) GetAttributesOk() (*NssfFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NssfFunctionSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given NssfFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *NssfFunctionSingle) SetAttributes(v NssfFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *NssfFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || IsNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfFunctionSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || IsNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *NssfFunctionSingle) HasPerfMetricJob() bool {
	if o != nil && !IsNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *NssfFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *NssfFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || IsNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfFunctionSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || IsNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *NssfFunctionSingle) HasThresholdMonitor() bool {
	if o != nil && !IsNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *NssfFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetManagedNFService returns the ManagedNFService field value if set, zero value otherwise.
func (o *NssfFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle {
	if o == nil || IsNil(o.ManagedNFService) {
		var ret []ManagedNFServiceSingle
		return ret
	}
	return o.ManagedNFService
}

// GetManagedNFServiceOk returns a tuple with the ManagedNFService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfFunctionSingle) GetManagedNFServiceOk() ([]ManagedNFServiceSingle, bool) {
	if o == nil || IsNil(o.ManagedNFService) {
		return nil, false
	}
	return o.ManagedNFService, true
}

// HasManagedNFService returns a boolean if a field has been set.
func (o *NssfFunctionSingle) HasManagedNFService() bool {
	if o != nil && !IsNil(o.ManagedNFService) {
		return true
	}

	return false
}

// SetManagedNFService gets a reference to the given []ManagedNFServiceSingle and assigns it to the ManagedNFService field.
func (o *NssfFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle) {
	o.ManagedNFService = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *NssfFunctionSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || IsNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfFunctionSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || IsNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *NssfFunctionSingle) HasTraceJob() bool {
	if o != nil && !IsNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *NssfFunctionSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetEPN22 returns the EPN22 field value if set, zero value otherwise.
func (o *NssfFunctionSingle) GetEPN22() []EPN22Single {
	if o == nil || IsNil(o.EPN22) {
		var ret []EPN22Single
		return ret
	}
	return o.EPN22
}

// GetEPN22Ok returns a tuple with the EPN22 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfFunctionSingle) GetEPN22Ok() ([]EPN22Single, bool) {
	if o == nil || IsNil(o.EPN22) {
		return nil, false
	}
	return o.EPN22, true
}

// HasEPN22 returns a boolean if a field has been set.
func (o *NssfFunctionSingle) HasEPN22() bool {
	if o != nil && !IsNil(o.EPN22) {
		return true
	}

	return false
}

// SetEPN22 gets a reference to the given []EPN22Single and assigns it to the EPN22 field.
func (o *NssfFunctionSingle) SetEPN22(v []EPN22Single) {
	o.EPN22 = v
}

// GetEPN31 returns the EPN31 field value if set, zero value otherwise.
func (o *NssfFunctionSingle) GetEPN31() []EPN31Single {
	if o == nil || IsNil(o.EPN31) {
		var ret []EPN31Single
		return ret
	}
	return o.EPN31
}

// GetEPN31Ok returns a tuple with the EPN31 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfFunctionSingle) GetEPN31Ok() ([]EPN31Single, bool) {
	if o == nil || IsNil(o.EPN31) {
		return nil, false
	}
	return o.EPN31, true
}

// HasEPN31 returns a boolean if a field has been set.
func (o *NssfFunctionSingle) HasEPN31() bool {
	if o != nil && !IsNil(o.EPN31) {
		return true
	}

	return false
}

// SetEPN31 gets a reference to the given []EPN31Single and assigns it to the EPN31 field.
func (o *NssfFunctionSingle) SetEPN31(v []EPN31Single) {
	o.EPN31 = v
}

func (o NssfFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NssfFunctionSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if !IsNil(o.ObjectClass) {
		toSerialize["objectClass"] = o.ObjectClass
	}
	if !IsNil(o.ObjectInstance) {
		toSerialize["objectInstance"] = o.ObjectInstance
	}
	if !IsNil(o.VsDataContainer) {
		toSerialize["VsDataContainer"] = o.VsDataContainer
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.PerfMetricJob) {
		toSerialize["PerfMetricJob"] = o.PerfMetricJob
	}
	if !IsNil(o.ThresholdMonitor) {
		toSerialize["ThresholdMonitor"] = o.ThresholdMonitor
	}
	if !IsNil(o.ManagedNFService) {
		toSerialize["ManagedNFService"] = o.ManagedNFService
	}
	if !IsNil(o.TraceJob) {
		toSerialize["TraceJob"] = o.TraceJob
	}
	if !IsNil(o.EPN22) {
		toSerialize["EP_N22"] = o.EPN22
	}
	if !IsNil(o.EPN31) {
		toSerialize["EP_N31"] = o.EPN31
	}
	return toSerialize, nil
}

type NullableNssfFunctionSingle struct {
	value *NssfFunctionSingle
	isSet bool
}

func (v NullableNssfFunctionSingle) Get() *NssfFunctionSingle {
	return v.value
}

func (v *NullableNssfFunctionSingle) Set(val *NssfFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableNssfFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableNssfFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNssfFunctionSingle(val *NssfFunctionSingle) *NullableNssfFunctionSingle {
	return &NullableNssfFunctionSingle{value: val, isSet: true}
}

func (v NullableNssfFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNssfFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



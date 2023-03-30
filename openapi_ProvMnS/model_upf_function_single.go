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

// checks if the UpfFunctionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpfFunctionSingle{}

// UpfFunctionSingle struct for UpfFunctionSingle
type UpfFunctionSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *UpfFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`
	ManagedNFService []ManagedNFServiceSingle `json:"ManagedNFService,omitempty"`
	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`
	EPN3 []EPN3Single `json:"EP_N3,omitempty"`
	EPN4 []EPN4Single `json:"EP_N4,omitempty"`
	EPN6 []EPN6Single `json:"EP_N6,omitempty"`
	EPN9 []EPN9Single `json:"EP_N9,omitempty"`
	EPS5U []EPS5USingle `json:"EP_S5U,omitempty"`
}

// NewUpfFunctionSingle instantiates a new UpfFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpfFunctionSingle(id NullableString) *UpfFunctionSingle {
	this := UpfFunctionSingle{}
	this.Id = id
	return &this
}

// NewUpfFunctionSingleWithDefaults instantiates a new UpfFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpfFunctionSingleWithDefaults() *UpfFunctionSingle {
	this := UpfFunctionSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpfFunctionSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpfFunctionSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *UpfFunctionSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetObjectClass() string {
	if o == nil || IsNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasObjectClass() bool {
	if o != nil && !IsNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *UpfFunctionSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetObjectInstance() string {
	if o == nil || IsNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasObjectInstance() bool {
	if o != nil && !IsNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *UpfFunctionSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || IsNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || IsNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasVsDataContainer() bool {
	if o != nil && !IsNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *UpfFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetAttributes() UpfFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret UpfFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetAttributesOk() (*UpfFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given UpfFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *UpfFunctionSingle) SetAttributes(v UpfFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || IsNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || IsNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasPerfMetricJob() bool {
	if o != nil && !IsNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *UpfFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || IsNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || IsNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasThresholdMonitor() bool {
	if o != nil && !IsNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *UpfFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetManagedNFService returns the ManagedNFService field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle {
	if o == nil || IsNil(o.ManagedNFService) {
		var ret []ManagedNFServiceSingle
		return ret
	}
	return o.ManagedNFService
}

// GetManagedNFServiceOk returns a tuple with the ManagedNFService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetManagedNFServiceOk() ([]ManagedNFServiceSingle, bool) {
	if o == nil || IsNil(o.ManagedNFService) {
		return nil, false
	}
	return o.ManagedNFService, true
}

// HasManagedNFService returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasManagedNFService() bool {
	if o != nil && !IsNil(o.ManagedNFService) {
		return true
	}

	return false
}

// SetManagedNFService gets a reference to the given []ManagedNFServiceSingle and assigns it to the ManagedNFService field.
func (o *UpfFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle) {
	o.ManagedNFService = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || IsNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || IsNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasTraceJob() bool {
	if o != nil && !IsNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *UpfFunctionSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetEPN3 returns the EPN3 field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetEPN3() []EPN3Single {
	if o == nil || IsNil(o.EPN3) {
		var ret []EPN3Single
		return ret
	}
	return o.EPN3
}

// GetEPN3Ok returns a tuple with the EPN3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetEPN3Ok() ([]EPN3Single, bool) {
	if o == nil || IsNil(o.EPN3) {
		return nil, false
	}
	return o.EPN3, true
}

// HasEPN3 returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasEPN3() bool {
	if o != nil && !IsNil(o.EPN3) {
		return true
	}

	return false
}

// SetEPN3 gets a reference to the given []EPN3Single and assigns it to the EPN3 field.
func (o *UpfFunctionSingle) SetEPN3(v []EPN3Single) {
	o.EPN3 = v
}

// GetEPN4 returns the EPN4 field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetEPN4() []EPN4Single {
	if o == nil || IsNil(o.EPN4) {
		var ret []EPN4Single
		return ret
	}
	return o.EPN4
}

// GetEPN4Ok returns a tuple with the EPN4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetEPN4Ok() ([]EPN4Single, bool) {
	if o == nil || IsNil(o.EPN4) {
		return nil, false
	}
	return o.EPN4, true
}

// HasEPN4 returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasEPN4() bool {
	if o != nil && !IsNil(o.EPN4) {
		return true
	}

	return false
}

// SetEPN4 gets a reference to the given []EPN4Single and assigns it to the EPN4 field.
func (o *UpfFunctionSingle) SetEPN4(v []EPN4Single) {
	o.EPN4 = v
}

// GetEPN6 returns the EPN6 field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetEPN6() []EPN6Single {
	if o == nil || IsNil(o.EPN6) {
		var ret []EPN6Single
		return ret
	}
	return o.EPN6
}

// GetEPN6Ok returns a tuple with the EPN6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetEPN6Ok() ([]EPN6Single, bool) {
	if o == nil || IsNil(o.EPN6) {
		return nil, false
	}
	return o.EPN6, true
}

// HasEPN6 returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasEPN6() bool {
	if o != nil && !IsNil(o.EPN6) {
		return true
	}

	return false
}

// SetEPN6 gets a reference to the given []EPN6Single and assigns it to the EPN6 field.
func (o *UpfFunctionSingle) SetEPN6(v []EPN6Single) {
	o.EPN6 = v
}

// GetEPN9 returns the EPN9 field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetEPN9() []EPN9Single {
	if o == nil || IsNil(o.EPN9) {
		var ret []EPN9Single
		return ret
	}
	return o.EPN9
}

// GetEPN9Ok returns a tuple with the EPN9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetEPN9Ok() ([]EPN9Single, bool) {
	if o == nil || IsNil(o.EPN9) {
		return nil, false
	}
	return o.EPN9, true
}

// HasEPN9 returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasEPN9() bool {
	if o != nil && !IsNil(o.EPN9) {
		return true
	}

	return false
}

// SetEPN9 gets a reference to the given []EPN9Single and assigns it to the EPN9 field.
func (o *UpfFunctionSingle) SetEPN9(v []EPN9Single) {
	o.EPN9 = v
}

// GetEPS5U returns the EPS5U field value if set, zero value otherwise.
func (o *UpfFunctionSingle) GetEPS5U() []EPS5USingle {
	if o == nil || IsNil(o.EPS5U) {
		var ret []EPS5USingle
		return ret
	}
	return o.EPS5U
}

// GetEPS5UOk returns a tuple with the EPS5U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingle) GetEPS5UOk() ([]EPS5USingle, bool) {
	if o == nil || IsNil(o.EPS5U) {
		return nil, false
	}
	return o.EPS5U, true
}

// HasEPS5U returns a boolean if a field has been set.
func (o *UpfFunctionSingle) HasEPS5U() bool {
	if o != nil && !IsNil(o.EPS5U) {
		return true
	}

	return false
}

// SetEPS5U gets a reference to the given []EPS5USingle and assigns it to the EPS5U field.
func (o *UpfFunctionSingle) SetEPS5U(v []EPS5USingle) {
	o.EPS5U = v
}

func (o UpfFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpfFunctionSingle) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EPN3) {
		toSerialize["EP_N3"] = o.EPN3
	}
	if !IsNil(o.EPN4) {
		toSerialize["EP_N4"] = o.EPN4
	}
	if !IsNil(o.EPN6) {
		toSerialize["EP_N6"] = o.EPN6
	}
	if !IsNil(o.EPN9) {
		toSerialize["EP_N9"] = o.EPN9
	}
	if !IsNil(o.EPS5U) {
		toSerialize["EP_S5U"] = o.EPS5U
	}
	return toSerialize, nil
}

type NullableUpfFunctionSingle struct {
	value *UpfFunctionSingle
	isSet bool
}

func (v NullableUpfFunctionSingle) Get() *UpfFunctionSingle {
	return v.value
}

func (v *NullableUpfFunctionSingle) Set(val *UpfFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableUpfFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableUpfFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpfFunctionSingle(val *UpfFunctionSingle) *NullableUpfFunctionSingle {
	return &NullableUpfFunctionSingle{value: val, isSet: true}
}

func (v NullableUpfFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpfFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



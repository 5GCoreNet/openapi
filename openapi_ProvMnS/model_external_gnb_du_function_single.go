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

// checks if the ExternalGnbDuFunctionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalGnbDuFunctionSingle{}

// ExternalGnbDuFunctionSingle struct for ExternalGnbDuFunctionSingle
type ExternalGnbDuFunctionSingle struct {
	Top
	Attributes       *ExternalGnbDuFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
	PerfMetricJob    []PerfMetricJobSingle                       `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle                    `json:"ThresholdMonitor,omitempty"`
	ManagedNFService []ManagedNFServiceSingle                    `json:"ManagedNFService,omitempty"`
	TraceJob         []TraceJobSingle                            `json:"TraceJob,omitempty"`
	EPF1C            []EPF1CSingle                               `json:"EP_F1C,omitempty"`
	EPF1U            []EPF1USingle                               `json:"EP_F1U,omitempty"`
}

// NewExternalGnbDuFunctionSingle instantiates a new ExternalGnbDuFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGnbDuFunctionSingle(id NullableString) *ExternalGnbDuFunctionSingle {
	this := ExternalGnbDuFunctionSingle{}
	this.Id = id
	return &this
}

// NewExternalGnbDuFunctionSingleWithDefaults instantiates a new ExternalGnbDuFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGnbDuFunctionSingleWithDefaults() *ExternalGnbDuFunctionSingle {
	this := ExternalGnbDuFunctionSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalGnbDuFunctionSingle) GetAttributes() ExternalGnbDuFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ExternalGnbDuFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbDuFunctionSingle) GetAttributesOk() (*ExternalGnbDuFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalGnbDuFunctionSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ExternalGnbDuFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ExternalGnbDuFunctionSingle) SetAttributes(v ExternalGnbDuFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *ExternalGnbDuFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || IsNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbDuFunctionSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || IsNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *ExternalGnbDuFunctionSingle) HasPerfMetricJob() bool {
	if o != nil && !IsNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *ExternalGnbDuFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *ExternalGnbDuFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || IsNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbDuFunctionSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || IsNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *ExternalGnbDuFunctionSingle) HasThresholdMonitor() bool {
	if o != nil && !IsNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *ExternalGnbDuFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetManagedNFService returns the ManagedNFService field value if set, zero value otherwise.
func (o *ExternalGnbDuFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle {
	if o == nil || IsNil(o.ManagedNFService) {
		var ret []ManagedNFServiceSingle
		return ret
	}
	return o.ManagedNFService
}

// GetManagedNFServiceOk returns a tuple with the ManagedNFService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbDuFunctionSingle) GetManagedNFServiceOk() ([]ManagedNFServiceSingle, bool) {
	if o == nil || IsNil(o.ManagedNFService) {
		return nil, false
	}
	return o.ManagedNFService, true
}

// HasManagedNFService returns a boolean if a field has been set.
func (o *ExternalGnbDuFunctionSingle) HasManagedNFService() bool {
	if o != nil && !IsNil(o.ManagedNFService) {
		return true
	}

	return false
}

// SetManagedNFService gets a reference to the given []ManagedNFServiceSingle and assigns it to the ManagedNFService field.
func (o *ExternalGnbDuFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle) {
	o.ManagedNFService = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *ExternalGnbDuFunctionSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || IsNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbDuFunctionSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || IsNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *ExternalGnbDuFunctionSingle) HasTraceJob() bool {
	if o != nil && !IsNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *ExternalGnbDuFunctionSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetEPF1C returns the EPF1C field value if set, zero value otherwise.
func (o *ExternalGnbDuFunctionSingle) GetEPF1C() []EPF1CSingle {
	if o == nil || IsNil(o.EPF1C) {
		var ret []EPF1CSingle
		return ret
	}
	return o.EPF1C
}

// GetEPF1COk returns a tuple with the EPF1C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbDuFunctionSingle) GetEPF1COk() ([]EPF1CSingle, bool) {
	if o == nil || IsNil(o.EPF1C) {
		return nil, false
	}
	return o.EPF1C, true
}

// HasEPF1C returns a boolean if a field has been set.
func (o *ExternalGnbDuFunctionSingle) HasEPF1C() bool {
	if o != nil && !IsNil(o.EPF1C) {
		return true
	}

	return false
}

// SetEPF1C gets a reference to the given []EPF1CSingle and assigns it to the EPF1C field.
func (o *ExternalGnbDuFunctionSingle) SetEPF1C(v []EPF1CSingle) {
	o.EPF1C = v
}

// GetEPF1U returns the EPF1U field value if set, zero value otherwise.
func (o *ExternalGnbDuFunctionSingle) GetEPF1U() []EPF1USingle {
	if o == nil || IsNil(o.EPF1U) {
		var ret []EPF1USingle
		return ret
	}
	return o.EPF1U
}

// GetEPF1UOk returns a tuple with the EPF1U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbDuFunctionSingle) GetEPF1UOk() ([]EPF1USingle, bool) {
	if o == nil || IsNil(o.EPF1U) {
		return nil, false
	}
	return o.EPF1U, true
}

// HasEPF1U returns a boolean if a field has been set.
func (o *ExternalGnbDuFunctionSingle) HasEPF1U() bool {
	if o != nil && !IsNil(o.EPF1U) {
		return true
	}

	return false
}

// SetEPF1U gets a reference to the given []EPF1USingle and assigns it to the EPF1U field.
func (o *ExternalGnbDuFunctionSingle) SetEPF1U(v []EPF1USingle) {
	o.EPF1U = v
}

func (o ExternalGnbDuFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalGnbDuFunctionSingle) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EPF1C) {
		toSerialize["EP_F1C"] = o.EPF1C
	}
	if !IsNil(o.EPF1U) {
		toSerialize["EP_F1U"] = o.EPF1U
	}
	return toSerialize, nil
}

type NullableExternalGnbDuFunctionSingle struct {
	value *ExternalGnbDuFunctionSingle
	isSet bool
}

func (v NullableExternalGnbDuFunctionSingle) Get() *ExternalGnbDuFunctionSingle {
	return v.value
}

func (v *NullableExternalGnbDuFunctionSingle) Set(val *ExternalGnbDuFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGnbDuFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGnbDuFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGnbDuFunctionSingle(val *ExternalGnbDuFunctionSingle) *NullableExternalGnbDuFunctionSingle {
	return &NullableExternalGnbDuFunctionSingle{value: val, isSet: true}
}

func (v NullableExternalGnbDuFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGnbDuFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

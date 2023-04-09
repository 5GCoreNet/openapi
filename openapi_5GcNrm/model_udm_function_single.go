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

// checks if the UdmFunctionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UdmFunctionSingle{}

// UdmFunctionSingle struct for UdmFunctionSingle
type UdmFunctionSingle struct {
	Top
	Attributes       *UdmFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
	PerfMetricJob    []PerfMetricJobSingle             `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle          `json:"ThresholdMonitor,omitempty"`
	ManagedNFService []ManagedNFServiceSingle          `json:"ManagedNFService,omitempty"`
	TraceJob         []TraceJobSingle                  `json:"TraceJob,omitempty"`
	EPN8             []EPN8Single                      `json:"EP_N8,omitempty"`
	EPN10            []EPN10Single                     `json:"EP_N10,omitempty"`
	EPN13            []EPN13Single                     `json:"EP_N13,omitempty"`
}

// NewUdmFunctionSingle instantiates a new UdmFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdmFunctionSingle(id NullableString) *UdmFunctionSingle {
	this := UdmFunctionSingle{}
	this.Id = id
	return &this
}

// NewUdmFunctionSingleWithDefaults instantiates a new UdmFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdmFunctionSingleWithDefaults() *UdmFunctionSingle {
	this := UdmFunctionSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UdmFunctionSingle) GetAttributes() UdmFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret UdmFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingle) GetAttributesOk() (*UdmFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UdmFunctionSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given UdmFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *UdmFunctionSingle) SetAttributes(v UdmFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *UdmFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || IsNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || IsNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *UdmFunctionSingle) HasPerfMetricJob() bool {
	if o != nil && !IsNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *UdmFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *UdmFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || IsNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || IsNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *UdmFunctionSingle) HasThresholdMonitor() bool {
	if o != nil && !IsNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *UdmFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetManagedNFService returns the ManagedNFService field value if set, zero value otherwise.
func (o *UdmFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle {
	if o == nil || IsNil(o.ManagedNFService) {
		var ret []ManagedNFServiceSingle
		return ret
	}
	return o.ManagedNFService
}

// GetManagedNFServiceOk returns a tuple with the ManagedNFService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingle) GetManagedNFServiceOk() ([]ManagedNFServiceSingle, bool) {
	if o == nil || IsNil(o.ManagedNFService) {
		return nil, false
	}
	return o.ManagedNFService, true
}

// HasManagedNFService returns a boolean if a field has been set.
func (o *UdmFunctionSingle) HasManagedNFService() bool {
	if o != nil && !IsNil(o.ManagedNFService) {
		return true
	}

	return false
}

// SetManagedNFService gets a reference to the given []ManagedNFServiceSingle and assigns it to the ManagedNFService field.
func (o *UdmFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle) {
	o.ManagedNFService = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *UdmFunctionSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || IsNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || IsNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *UdmFunctionSingle) HasTraceJob() bool {
	if o != nil && !IsNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *UdmFunctionSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetEPN8 returns the EPN8 field value if set, zero value otherwise.
func (o *UdmFunctionSingle) GetEPN8() []EPN8Single {
	if o == nil || IsNil(o.EPN8) {
		var ret []EPN8Single
		return ret
	}
	return o.EPN8
}

// GetEPN8Ok returns a tuple with the EPN8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingle) GetEPN8Ok() ([]EPN8Single, bool) {
	if o == nil || IsNil(o.EPN8) {
		return nil, false
	}
	return o.EPN8, true
}

// HasEPN8 returns a boolean if a field has been set.
func (o *UdmFunctionSingle) HasEPN8() bool {
	if o != nil && !IsNil(o.EPN8) {
		return true
	}

	return false
}

// SetEPN8 gets a reference to the given []EPN8Single and assigns it to the EPN8 field.
func (o *UdmFunctionSingle) SetEPN8(v []EPN8Single) {
	o.EPN8 = v
}

// GetEPN10 returns the EPN10 field value if set, zero value otherwise.
func (o *UdmFunctionSingle) GetEPN10() []EPN10Single {
	if o == nil || IsNil(o.EPN10) {
		var ret []EPN10Single
		return ret
	}
	return o.EPN10
}

// GetEPN10Ok returns a tuple with the EPN10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingle) GetEPN10Ok() ([]EPN10Single, bool) {
	if o == nil || IsNil(o.EPN10) {
		return nil, false
	}
	return o.EPN10, true
}

// HasEPN10 returns a boolean if a field has been set.
func (o *UdmFunctionSingle) HasEPN10() bool {
	if o != nil && !IsNil(o.EPN10) {
		return true
	}

	return false
}

// SetEPN10 gets a reference to the given []EPN10Single and assigns it to the EPN10 field.
func (o *UdmFunctionSingle) SetEPN10(v []EPN10Single) {
	o.EPN10 = v
}

// GetEPN13 returns the EPN13 field value if set, zero value otherwise.
func (o *UdmFunctionSingle) GetEPN13() []EPN13Single {
	if o == nil || IsNil(o.EPN13) {
		var ret []EPN13Single
		return ret
	}
	return o.EPN13
}

// GetEPN13Ok returns a tuple with the EPN13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingle) GetEPN13Ok() ([]EPN13Single, bool) {
	if o == nil || IsNil(o.EPN13) {
		return nil, false
	}
	return o.EPN13, true
}

// HasEPN13 returns a boolean if a field has been set.
func (o *UdmFunctionSingle) HasEPN13() bool {
	if o != nil && !IsNil(o.EPN13) {
		return true
	}

	return false
}

// SetEPN13 gets a reference to the given []EPN13Single and assigns it to the EPN13 field.
func (o *UdmFunctionSingle) SetEPN13(v []EPN13Single) {
	o.EPN13 = v
}

func (o UdmFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UdmFunctionSingle) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EPN8) {
		toSerialize["EP_N8"] = o.EPN8
	}
	if !IsNil(o.EPN10) {
		toSerialize["EP_N10"] = o.EPN10
	}
	if !IsNil(o.EPN13) {
		toSerialize["EP_N13"] = o.EPN13
	}
	return toSerialize, nil
}

type NullableUdmFunctionSingle struct {
	value *UdmFunctionSingle
	isSet bool
}

func (v NullableUdmFunctionSingle) Get() *UdmFunctionSingle {
	return v.value
}

func (v *NullableUdmFunctionSingle) Set(val *UdmFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableUdmFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableUdmFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdmFunctionSingle(val *UdmFunctionSingle) *NullableUdmFunctionSingle {
	return &NullableUdmFunctionSingle{value: val, isSet: true}
}

func (v NullableUdmFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdmFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

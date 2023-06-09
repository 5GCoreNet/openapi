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

// checks if the AusfFunctionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AusfFunctionSingle{}

// AusfFunctionSingle struct for AusfFunctionSingle
type AusfFunctionSingle struct {
	Top
	Attributes       *AusfFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
	PerfMetricJob    []PerfMetricJobSingle              `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle           `json:"ThresholdMonitor,omitempty"`
	ManagedNFService []ManagedNFServiceSingle           `json:"ManagedNFService,omitempty"`
	TraceJob         []TraceJobSingle                   `json:"TraceJob,omitempty"`
	EPN12            []EPN12Single                      `json:"EP_N12,omitempty"`
	EPN13            []EPN13Single                      `json:"EP_N13,omitempty"`
}

// NewAusfFunctionSingle instantiates a new AusfFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAusfFunctionSingle(id NullableString) *AusfFunctionSingle {
	this := AusfFunctionSingle{}
	this.Id = id
	return &this
}

// NewAusfFunctionSingleWithDefaults instantiates a new AusfFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAusfFunctionSingleWithDefaults() *AusfFunctionSingle {
	this := AusfFunctionSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AusfFunctionSingle) GetAttributes() AusfFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AusfFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingle) GetAttributesOk() (*AusfFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AusfFunctionSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AusfFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *AusfFunctionSingle) SetAttributes(v AusfFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *AusfFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || IsNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || IsNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *AusfFunctionSingle) HasPerfMetricJob() bool {
	if o != nil && !IsNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *AusfFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *AusfFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || IsNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || IsNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *AusfFunctionSingle) HasThresholdMonitor() bool {
	if o != nil && !IsNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *AusfFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetManagedNFService returns the ManagedNFService field value if set, zero value otherwise.
func (o *AusfFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle {
	if o == nil || IsNil(o.ManagedNFService) {
		var ret []ManagedNFServiceSingle
		return ret
	}
	return o.ManagedNFService
}

// GetManagedNFServiceOk returns a tuple with the ManagedNFService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingle) GetManagedNFServiceOk() ([]ManagedNFServiceSingle, bool) {
	if o == nil || IsNil(o.ManagedNFService) {
		return nil, false
	}
	return o.ManagedNFService, true
}

// HasManagedNFService returns a boolean if a field has been set.
func (o *AusfFunctionSingle) HasManagedNFService() bool {
	if o != nil && !IsNil(o.ManagedNFService) {
		return true
	}

	return false
}

// SetManagedNFService gets a reference to the given []ManagedNFServiceSingle and assigns it to the ManagedNFService field.
func (o *AusfFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle) {
	o.ManagedNFService = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *AusfFunctionSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || IsNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || IsNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *AusfFunctionSingle) HasTraceJob() bool {
	if o != nil && !IsNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *AusfFunctionSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetEPN12 returns the EPN12 field value if set, zero value otherwise.
func (o *AusfFunctionSingle) GetEPN12() []EPN12Single {
	if o == nil || IsNil(o.EPN12) {
		var ret []EPN12Single
		return ret
	}
	return o.EPN12
}

// GetEPN12Ok returns a tuple with the EPN12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingle) GetEPN12Ok() ([]EPN12Single, bool) {
	if o == nil || IsNil(o.EPN12) {
		return nil, false
	}
	return o.EPN12, true
}

// HasEPN12 returns a boolean if a field has been set.
func (o *AusfFunctionSingle) HasEPN12() bool {
	if o != nil && !IsNil(o.EPN12) {
		return true
	}

	return false
}

// SetEPN12 gets a reference to the given []EPN12Single and assigns it to the EPN12 field.
func (o *AusfFunctionSingle) SetEPN12(v []EPN12Single) {
	o.EPN12 = v
}

// GetEPN13 returns the EPN13 field value if set, zero value otherwise.
func (o *AusfFunctionSingle) GetEPN13() []EPN13Single {
	if o == nil || IsNil(o.EPN13) {
		var ret []EPN13Single
		return ret
	}
	return o.EPN13
}

// GetEPN13Ok returns a tuple with the EPN13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingle) GetEPN13Ok() ([]EPN13Single, bool) {
	if o == nil || IsNil(o.EPN13) {
		return nil, false
	}
	return o.EPN13, true
}

// HasEPN13 returns a boolean if a field has been set.
func (o *AusfFunctionSingle) HasEPN13() bool {
	if o != nil && !IsNil(o.EPN13) {
		return true
	}

	return false
}

// SetEPN13 gets a reference to the given []EPN13Single and assigns it to the EPN13 field.
func (o *AusfFunctionSingle) SetEPN13(v []EPN13Single) {
	o.EPN13 = v
}

func (o AusfFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AusfFunctionSingle) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EPN12) {
		toSerialize["EP_N12"] = o.EPN12
	}
	if !IsNil(o.EPN13) {
		toSerialize["EP_N13"] = o.EPN13
	}
	return toSerialize, nil
}

type NullableAusfFunctionSingle struct {
	value *AusfFunctionSingle
	isSet bool
}

func (v NullableAusfFunctionSingle) Get() *AusfFunctionSingle {
	return v.value
}

func (v *NullableAusfFunctionSingle) Set(val *AusfFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableAusfFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableAusfFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAusfFunctionSingle(val *AusfFunctionSingle) *NullableAusfFunctionSingle {
	return &NullableAusfFunctionSingle{value: val, isSet: true}
}

func (v NullableAusfFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAusfFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

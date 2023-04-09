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

// checks if the GnbCuUpFunctionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GnbCuUpFunctionSingle{}

// GnbCuUpFunctionSingle struct for GnbCuUpFunctionSingle
type GnbCuUpFunctionSingle struct {
	Top
	Attributes       *GnbCuUpFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
	PerfMetricJob    []PerfMetricJobSingle                 `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle              `json:"ThresholdMonitor,omitempty"`
	ManagedNFService []ManagedNFServiceSingle              `json:"ManagedNFService,omitempty"`
	TraceJob         []TraceJobSingle                      `json:"TraceJob,omitempty"`
	RRMPolicyRatio   []RRMPolicyRatioSingle                `json:"RRMPolicyRatio,omitempty"`
	EPE1             *EPE1Single                           `json:"EP_E1,omitempty"`
	EPXnU            []EPXnUSingle                         `json:"EP_XnU,omitempty"`
	EPF1U            []EPF1USingle                         `json:"EP_F1U,omitempty"`
	EPNgU            []EPNgUSingle                         `json:"EP_NgU,omitempty"`
	EPX2U            []EPX2USingle                         `json:"EP_X2U,omitempty"`
	EPS1U            []EPS1USingle                         `json:"EP_S1U,omitempty"`
}

// NewGnbCuUpFunctionSingle instantiates a new GnbCuUpFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGnbCuUpFunctionSingle(id NullableString) *GnbCuUpFunctionSingle {
	this := GnbCuUpFunctionSingle{}
	this.Id = id
	return &this
}

// NewGnbCuUpFunctionSingleWithDefaults instantiates a new GnbCuUpFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGnbCuUpFunctionSingleWithDefaults() *GnbCuUpFunctionSingle {
	this := GnbCuUpFunctionSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GnbCuUpFunctionSingle) GetAttributes() GnbCuUpFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GnbCuUpFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuUpFunctionSingle) GetAttributesOk() (*GnbCuUpFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GnbCuUpFunctionSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GnbCuUpFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *GnbCuUpFunctionSingle) SetAttributes(v GnbCuUpFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *GnbCuUpFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || IsNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuUpFunctionSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || IsNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *GnbCuUpFunctionSingle) HasPerfMetricJob() bool {
	if o != nil && !IsNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *GnbCuUpFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *GnbCuUpFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || IsNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuUpFunctionSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || IsNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *GnbCuUpFunctionSingle) HasThresholdMonitor() bool {
	if o != nil && !IsNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *GnbCuUpFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetManagedNFService returns the ManagedNFService field value if set, zero value otherwise.
func (o *GnbCuUpFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle {
	if o == nil || IsNil(o.ManagedNFService) {
		var ret []ManagedNFServiceSingle
		return ret
	}
	return o.ManagedNFService
}

// GetManagedNFServiceOk returns a tuple with the ManagedNFService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuUpFunctionSingle) GetManagedNFServiceOk() ([]ManagedNFServiceSingle, bool) {
	if o == nil || IsNil(o.ManagedNFService) {
		return nil, false
	}
	return o.ManagedNFService, true
}

// HasManagedNFService returns a boolean if a field has been set.
func (o *GnbCuUpFunctionSingle) HasManagedNFService() bool {
	if o != nil && !IsNil(o.ManagedNFService) {
		return true
	}

	return false
}

// SetManagedNFService gets a reference to the given []ManagedNFServiceSingle and assigns it to the ManagedNFService field.
func (o *GnbCuUpFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle) {
	o.ManagedNFService = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *GnbCuUpFunctionSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || IsNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuUpFunctionSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || IsNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *GnbCuUpFunctionSingle) HasTraceJob() bool {
	if o != nil && !IsNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *GnbCuUpFunctionSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetRRMPolicyRatio returns the RRMPolicyRatio field value if set, zero value otherwise.
func (o *GnbCuUpFunctionSingle) GetRRMPolicyRatio() []RRMPolicyRatioSingle {
	if o == nil || IsNil(o.RRMPolicyRatio) {
		var ret []RRMPolicyRatioSingle
		return ret
	}
	return o.RRMPolicyRatio
}

// GetRRMPolicyRatioOk returns a tuple with the RRMPolicyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuUpFunctionSingle) GetRRMPolicyRatioOk() ([]RRMPolicyRatioSingle, bool) {
	if o == nil || IsNil(o.RRMPolicyRatio) {
		return nil, false
	}
	return o.RRMPolicyRatio, true
}

// HasRRMPolicyRatio returns a boolean if a field has been set.
func (o *GnbCuUpFunctionSingle) HasRRMPolicyRatio() bool {
	if o != nil && !IsNil(o.RRMPolicyRatio) {
		return true
	}

	return false
}

// SetRRMPolicyRatio gets a reference to the given []RRMPolicyRatioSingle and assigns it to the RRMPolicyRatio field.
func (o *GnbCuUpFunctionSingle) SetRRMPolicyRatio(v []RRMPolicyRatioSingle) {
	o.RRMPolicyRatio = v
}

// GetEPE1 returns the EPE1 field value if set, zero value otherwise.
func (o *GnbCuUpFunctionSingle) GetEPE1() EPE1Single {
	if o == nil || IsNil(o.EPE1) {
		var ret EPE1Single
		return ret
	}
	return *o.EPE1
}

// GetEPE1Ok returns a tuple with the EPE1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuUpFunctionSingle) GetEPE1Ok() (*EPE1Single, bool) {
	if o == nil || IsNil(o.EPE1) {
		return nil, false
	}
	return o.EPE1, true
}

// HasEPE1 returns a boolean if a field has been set.
func (o *GnbCuUpFunctionSingle) HasEPE1() bool {
	if o != nil && !IsNil(o.EPE1) {
		return true
	}

	return false
}

// SetEPE1 gets a reference to the given EPE1Single and assigns it to the EPE1 field.
func (o *GnbCuUpFunctionSingle) SetEPE1(v EPE1Single) {
	o.EPE1 = &v
}

// GetEPXnU returns the EPXnU field value if set, zero value otherwise.
func (o *GnbCuUpFunctionSingle) GetEPXnU() []EPXnUSingle {
	if o == nil || IsNil(o.EPXnU) {
		var ret []EPXnUSingle
		return ret
	}
	return o.EPXnU
}

// GetEPXnUOk returns a tuple with the EPXnU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuUpFunctionSingle) GetEPXnUOk() ([]EPXnUSingle, bool) {
	if o == nil || IsNil(o.EPXnU) {
		return nil, false
	}
	return o.EPXnU, true
}

// HasEPXnU returns a boolean if a field has been set.
func (o *GnbCuUpFunctionSingle) HasEPXnU() bool {
	if o != nil && !IsNil(o.EPXnU) {
		return true
	}

	return false
}

// SetEPXnU gets a reference to the given []EPXnUSingle and assigns it to the EPXnU field.
func (o *GnbCuUpFunctionSingle) SetEPXnU(v []EPXnUSingle) {
	o.EPXnU = v
}

// GetEPF1U returns the EPF1U field value if set, zero value otherwise.
func (o *GnbCuUpFunctionSingle) GetEPF1U() []EPF1USingle {
	if o == nil || IsNil(o.EPF1U) {
		var ret []EPF1USingle
		return ret
	}
	return o.EPF1U
}

// GetEPF1UOk returns a tuple with the EPF1U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuUpFunctionSingle) GetEPF1UOk() ([]EPF1USingle, bool) {
	if o == nil || IsNil(o.EPF1U) {
		return nil, false
	}
	return o.EPF1U, true
}

// HasEPF1U returns a boolean if a field has been set.
func (o *GnbCuUpFunctionSingle) HasEPF1U() bool {
	if o != nil && !IsNil(o.EPF1U) {
		return true
	}

	return false
}

// SetEPF1U gets a reference to the given []EPF1USingle and assigns it to the EPF1U field.
func (o *GnbCuUpFunctionSingle) SetEPF1U(v []EPF1USingle) {
	o.EPF1U = v
}

// GetEPNgU returns the EPNgU field value if set, zero value otherwise.
func (o *GnbCuUpFunctionSingle) GetEPNgU() []EPNgUSingle {
	if o == nil || IsNil(o.EPNgU) {
		var ret []EPNgUSingle
		return ret
	}
	return o.EPNgU
}

// GetEPNgUOk returns a tuple with the EPNgU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuUpFunctionSingle) GetEPNgUOk() ([]EPNgUSingle, bool) {
	if o == nil || IsNil(o.EPNgU) {
		return nil, false
	}
	return o.EPNgU, true
}

// HasEPNgU returns a boolean if a field has been set.
func (o *GnbCuUpFunctionSingle) HasEPNgU() bool {
	if o != nil && !IsNil(o.EPNgU) {
		return true
	}

	return false
}

// SetEPNgU gets a reference to the given []EPNgUSingle and assigns it to the EPNgU field.
func (o *GnbCuUpFunctionSingle) SetEPNgU(v []EPNgUSingle) {
	o.EPNgU = v
}

// GetEPX2U returns the EPX2U field value if set, zero value otherwise.
func (o *GnbCuUpFunctionSingle) GetEPX2U() []EPX2USingle {
	if o == nil || IsNil(o.EPX2U) {
		var ret []EPX2USingle
		return ret
	}
	return o.EPX2U
}

// GetEPX2UOk returns a tuple with the EPX2U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuUpFunctionSingle) GetEPX2UOk() ([]EPX2USingle, bool) {
	if o == nil || IsNil(o.EPX2U) {
		return nil, false
	}
	return o.EPX2U, true
}

// HasEPX2U returns a boolean if a field has been set.
func (o *GnbCuUpFunctionSingle) HasEPX2U() bool {
	if o != nil && !IsNil(o.EPX2U) {
		return true
	}

	return false
}

// SetEPX2U gets a reference to the given []EPX2USingle and assigns it to the EPX2U field.
func (o *GnbCuUpFunctionSingle) SetEPX2U(v []EPX2USingle) {
	o.EPX2U = v
}

// GetEPS1U returns the EPS1U field value if set, zero value otherwise.
func (o *GnbCuUpFunctionSingle) GetEPS1U() []EPS1USingle {
	if o == nil || IsNil(o.EPS1U) {
		var ret []EPS1USingle
		return ret
	}
	return o.EPS1U
}

// GetEPS1UOk returns a tuple with the EPS1U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuUpFunctionSingle) GetEPS1UOk() ([]EPS1USingle, bool) {
	if o == nil || IsNil(o.EPS1U) {
		return nil, false
	}
	return o.EPS1U, true
}

// HasEPS1U returns a boolean if a field has been set.
func (o *GnbCuUpFunctionSingle) HasEPS1U() bool {
	if o != nil && !IsNil(o.EPS1U) {
		return true
	}

	return false
}

// SetEPS1U gets a reference to the given []EPS1USingle and assigns it to the EPS1U field.
func (o *GnbCuUpFunctionSingle) SetEPS1U(v []EPS1USingle) {
	o.EPS1U = v
}

func (o GnbCuUpFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GnbCuUpFunctionSingle) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RRMPolicyRatio) {
		toSerialize["RRMPolicyRatio"] = o.RRMPolicyRatio
	}
	if !IsNil(o.EPE1) {
		toSerialize["EP_E1"] = o.EPE1
	}
	if !IsNil(o.EPXnU) {
		toSerialize["EP_XnU"] = o.EPXnU
	}
	if !IsNil(o.EPF1U) {
		toSerialize["EP_F1U"] = o.EPF1U
	}
	if !IsNil(o.EPNgU) {
		toSerialize["EP_NgU"] = o.EPNgU
	}
	if !IsNil(o.EPX2U) {
		toSerialize["EP_X2U"] = o.EPX2U
	}
	if !IsNil(o.EPS1U) {
		toSerialize["EP_S1U"] = o.EPS1U
	}
	return toSerialize, nil
}

type NullableGnbCuUpFunctionSingle struct {
	value *GnbCuUpFunctionSingle
	isSet bool
}

func (v NullableGnbCuUpFunctionSingle) Get() *GnbCuUpFunctionSingle {
	return v.value
}

func (v *NullableGnbCuUpFunctionSingle) Set(val *GnbCuUpFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableGnbCuUpFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableGnbCuUpFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGnbCuUpFunctionSingle(val *GnbCuUpFunctionSingle) *NullableGnbCuUpFunctionSingle {
	return &NullableGnbCuUpFunctionSingle{value: val, isSet: true}
}

func (v NullableGnbCuUpFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGnbCuUpFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

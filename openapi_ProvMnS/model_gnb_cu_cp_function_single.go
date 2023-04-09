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

// checks if the GnbCuCpFunctionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GnbCuCpFunctionSingle{}

// GnbCuCpFunctionSingle struct for GnbCuCpFunctionSingle
type GnbCuCpFunctionSingle struct {
	Top
	Attributes             *GnbCuCpFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
	PerfMetricJob          []PerfMetricJobSingle                 `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor       []ThresholdMonitorSingle              `json:"ThresholdMonitor,omitempty"`
	ManagedNFService       []ManagedNFServiceSingle              `json:"ManagedNFService,omitempty"`
	TraceJob               []TraceJobSingle                      `json:"TraceJob,omitempty"`
	RRMPolicyRatio         []RRMPolicyRatioSingle                `json:"RRMPolicyRatio,omitempty"`
	NrCellCu               []NrCellCuSingle                      `json:"NrCellCu,omitempty"`
	EPXnC                  []EPXnCSingle                         `json:"EP_XnC,omitempty"`
	EPE1                   []EPE1Single                          `json:"EP_E1,omitempty"`
	EPF1C                  []EPF1CSingle                         `json:"EP_F1C,omitempty"`
	EPNgC                  []EPNgCSingle                         `json:"EP_NgC,omitempty"`
	EPX2C                  []EPX2CSingle                         `json:"EP_X2C,omitempty"`
	DANRManagementFunction *DANRManagementFunctionSingle         `json:"DANRManagementFunction,omitempty"`
	DESManagementFunction  *DESManagementFunctionSingle          `json:"DESManagementFunction,omitempty"`
	DMROFunction           *DMROFunctionSingle                   `json:"DMROFunction,omitempty"`
	DLBOFunction           *DLBOFunctionSingle                   `json:"DLBOFunction,omitempty"`
}

// NewGnbCuCpFunctionSingle instantiates a new GnbCuCpFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGnbCuCpFunctionSingle(id NullableString) *GnbCuCpFunctionSingle {
	this := GnbCuCpFunctionSingle{}
	this.Id = id
	return &this
}

// NewGnbCuCpFunctionSingleWithDefaults instantiates a new GnbCuCpFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGnbCuCpFunctionSingleWithDefaults() *GnbCuCpFunctionSingle {
	this := GnbCuCpFunctionSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetAttributes() GnbCuCpFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GnbCuCpFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetAttributesOk() (*GnbCuCpFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GnbCuCpFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *GnbCuCpFunctionSingle) SetAttributes(v GnbCuCpFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || IsNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || IsNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasPerfMetricJob() bool {
	if o != nil && !IsNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *GnbCuCpFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || IsNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || IsNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasThresholdMonitor() bool {
	if o != nil && !IsNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *GnbCuCpFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetManagedNFService returns the ManagedNFService field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle {
	if o == nil || IsNil(o.ManagedNFService) {
		var ret []ManagedNFServiceSingle
		return ret
	}
	return o.ManagedNFService
}

// GetManagedNFServiceOk returns a tuple with the ManagedNFService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetManagedNFServiceOk() ([]ManagedNFServiceSingle, bool) {
	if o == nil || IsNil(o.ManagedNFService) {
		return nil, false
	}
	return o.ManagedNFService, true
}

// HasManagedNFService returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasManagedNFService() bool {
	if o != nil && !IsNil(o.ManagedNFService) {
		return true
	}

	return false
}

// SetManagedNFService gets a reference to the given []ManagedNFServiceSingle and assigns it to the ManagedNFService field.
func (o *GnbCuCpFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle) {
	o.ManagedNFService = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || IsNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || IsNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasTraceJob() bool {
	if o != nil && !IsNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *GnbCuCpFunctionSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetRRMPolicyRatio returns the RRMPolicyRatio field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetRRMPolicyRatio() []RRMPolicyRatioSingle {
	if o == nil || IsNil(o.RRMPolicyRatio) {
		var ret []RRMPolicyRatioSingle
		return ret
	}
	return o.RRMPolicyRatio
}

// GetRRMPolicyRatioOk returns a tuple with the RRMPolicyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetRRMPolicyRatioOk() ([]RRMPolicyRatioSingle, bool) {
	if o == nil || IsNil(o.RRMPolicyRatio) {
		return nil, false
	}
	return o.RRMPolicyRatio, true
}

// HasRRMPolicyRatio returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasRRMPolicyRatio() bool {
	if o != nil && !IsNil(o.RRMPolicyRatio) {
		return true
	}

	return false
}

// SetRRMPolicyRatio gets a reference to the given []RRMPolicyRatioSingle and assigns it to the RRMPolicyRatio field.
func (o *GnbCuCpFunctionSingle) SetRRMPolicyRatio(v []RRMPolicyRatioSingle) {
	o.RRMPolicyRatio = v
}

// GetNrCellCu returns the NrCellCu field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetNrCellCu() []NrCellCuSingle {
	if o == nil || IsNil(o.NrCellCu) {
		var ret []NrCellCuSingle
		return ret
	}
	return o.NrCellCu
}

// GetNrCellCuOk returns a tuple with the NrCellCu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetNrCellCuOk() ([]NrCellCuSingle, bool) {
	if o == nil || IsNil(o.NrCellCu) {
		return nil, false
	}
	return o.NrCellCu, true
}

// HasNrCellCu returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasNrCellCu() bool {
	if o != nil && !IsNil(o.NrCellCu) {
		return true
	}

	return false
}

// SetNrCellCu gets a reference to the given []NrCellCuSingle and assigns it to the NrCellCu field.
func (o *GnbCuCpFunctionSingle) SetNrCellCu(v []NrCellCuSingle) {
	o.NrCellCu = v
}

// GetEPXnC returns the EPXnC field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetEPXnC() []EPXnCSingle {
	if o == nil || IsNil(o.EPXnC) {
		var ret []EPXnCSingle
		return ret
	}
	return o.EPXnC
}

// GetEPXnCOk returns a tuple with the EPXnC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetEPXnCOk() ([]EPXnCSingle, bool) {
	if o == nil || IsNil(o.EPXnC) {
		return nil, false
	}
	return o.EPXnC, true
}

// HasEPXnC returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasEPXnC() bool {
	if o != nil && !IsNil(o.EPXnC) {
		return true
	}

	return false
}

// SetEPXnC gets a reference to the given []EPXnCSingle and assigns it to the EPXnC field.
func (o *GnbCuCpFunctionSingle) SetEPXnC(v []EPXnCSingle) {
	o.EPXnC = v
}

// GetEPE1 returns the EPE1 field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetEPE1() []EPE1Single {
	if o == nil || IsNil(o.EPE1) {
		var ret []EPE1Single
		return ret
	}
	return o.EPE1
}

// GetEPE1Ok returns a tuple with the EPE1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetEPE1Ok() ([]EPE1Single, bool) {
	if o == nil || IsNil(o.EPE1) {
		return nil, false
	}
	return o.EPE1, true
}

// HasEPE1 returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasEPE1() bool {
	if o != nil && !IsNil(o.EPE1) {
		return true
	}

	return false
}

// SetEPE1 gets a reference to the given []EPE1Single and assigns it to the EPE1 field.
func (o *GnbCuCpFunctionSingle) SetEPE1(v []EPE1Single) {
	o.EPE1 = v
}

// GetEPF1C returns the EPF1C field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetEPF1C() []EPF1CSingle {
	if o == nil || IsNil(o.EPF1C) {
		var ret []EPF1CSingle
		return ret
	}
	return o.EPF1C
}

// GetEPF1COk returns a tuple with the EPF1C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetEPF1COk() ([]EPF1CSingle, bool) {
	if o == nil || IsNil(o.EPF1C) {
		return nil, false
	}
	return o.EPF1C, true
}

// HasEPF1C returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasEPF1C() bool {
	if o != nil && !IsNil(o.EPF1C) {
		return true
	}

	return false
}

// SetEPF1C gets a reference to the given []EPF1CSingle and assigns it to the EPF1C field.
func (o *GnbCuCpFunctionSingle) SetEPF1C(v []EPF1CSingle) {
	o.EPF1C = v
}

// GetEPNgC returns the EPNgC field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetEPNgC() []EPNgCSingle {
	if o == nil || IsNil(o.EPNgC) {
		var ret []EPNgCSingle
		return ret
	}
	return o.EPNgC
}

// GetEPNgCOk returns a tuple with the EPNgC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetEPNgCOk() ([]EPNgCSingle, bool) {
	if o == nil || IsNil(o.EPNgC) {
		return nil, false
	}
	return o.EPNgC, true
}

// HasEPNgC returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasEPNgC() bool {
	if o != nil && !IsNil(o.EPNgC) {
		return true
	}

	return false
}

// SetEPNgC gets a reference to the given []EPNgCSingle and assigns it to the EPNgC field.
func (o *GnbCuCpFunctionSingle) SetEPNgC(v []EPNgCSingle) {
	o.EPNgC = v
}

// GetEPX2C returns the EPX2C field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetEPX2C() []EPX2CSingle {
	if o == nil || IsNil(o.EPX2C) {
		var ret []EPX2CSingle
		return ret
	}
	return o.EPX2C
}

// GetEPX2COk returns a tuple with the EPX2C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetEPX2COk() ([]EPX2CSingle, bool) {
	if o == nil || IsNil(o.EPX2C) {
		return nil, false
	}
	return o.EPX2C, true
}

// HasEPX2C returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasEPX2C() bool {
	if o != nil && !IsNil(o.EPX2C) {
		return true
	}

	return false
}

// SetEPX2C gets a reference to the given []EPX2CSingle and assigns it to the EPX2C field.
func (o *GnbCuCpFunctionSingle) SetEPX2C(v []EPX2CSingle) {
	o.EPX2C = v
}

// GetDANRManagementFunction returns the DANRManagementFunction field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetDANRManagementFunction() DANRManagementFunctionSingle {
	if o == nil || IsNil(o.DANRManagementFunction) {
		var ret DANRManagementFunctionSingle
		return ret
	}
	return *o.DANRManagementFunction
}

// GetDANRManagementFunctionOk returns a tuple with the DANRManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetDANRManagementFunctionOk() (*DANRManagementFunctionSingle, bool) {
	if o == nil || IsNil(o.DANRManagementFunction) {
		return nil, false
	}
	return o.DANRManagementFunction, true
}

// HasDANRManagementFunction returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasDANRManagementFunction() bool {
	if o != nil && !IsNil(o.DANRManagementFunction) {
		return true
	}

	return false
}

// SetDANRManagementFunction gets a reference to the given DANRManagementFunctionSingle and assigns it to the DANRManagementFunction field.
func (o *GnbCuCpFunctionSingle) SetDANRManagementFunction(v DANRManagementFunctionSingle) {
	o.DANRManagementFunction = &v
}

// GetDESManagementFunction returns the DESManagementFunction field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetDESManagementFunction() DESManagementFunctionSingle {
	if o == nil || IsNil(o.DESManagementFunction) {
		var ret DESManagementFunctionSingle
		return ret
	}
	return *o.DESManagementFunction
}

// GetDESManagementFunctionOk returns a tuple with the DESManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetDESManagementFunctionOk() (*DESManagementFunctionSingle, bool) {
	if o == nil || IsNil(o.DESManagementFunction) {
		return nil, false
	}
	return o.DESManagementFunction, true
}

// HasDESManagementFunction returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasDESManagementFunction() bool {
	if o != nil && !IsNil(o.DESManagementFunction) {
		return true
	}

	return false
}

// SetDESManagementFunction gets a reference to the given DESManagementFunctionSingle and assigns it to the DESManagementFunction field.
func (o *GnbCuCpFunctionSingle) SetDESManagementFunction(v DESManagementFunctionSingle) {
	o.DESManagementFunction = &v
}

// GetDMROFunction returns the DMROFunction field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetDMROFunction() DMROFunctionSingle {
	if o == nil || IsNil(o.DMROFunction) {
		var ret DMROFunctionSingle
		return ret
	}
	return *o.DMROFunction
}

// GetDMROFunctionOk returns a tuple with the DMROFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetDMROFunctionOk() (*DMROFunctionSingle, bool) {
	if o == nil || IsNil(o.DMROFunction) {
		return nil, false
	}
	return o.DMROFunction, true
}

// HasDMROFunction returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasDMROFunction() bool {
	if o != nil && !IsNil(o.DMROFunction) {
		return true
	}

	return false
}

// SetDMROFunction gets a reference to the given DMROFunctionSingle and assigns it to the DMROFunction field.
func (o *GnbCuCpFunctionSingle) SetDMROFunction(v DMROFunctionSingle) {
	o.DMROFunction = &v
}

// GetDLBOFunction returns the DLBOFunction field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingle) GetDLBOFunction() DLBOFunctionSingle {
	if o == nil || IsNil(o.DLBOFunction) {
		var ret DLBOFunctionSingle
		return ret
	}
	return *o.DLBOFunction
}

// GetDLBOFunctionOk returns a tuple with the DLBOFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingle) GetDLBOFunctionOk() (*DLBOFunctionSingle, bool) {
	if o == nil || IsNil(o.DLBOFunction) {
		return nil, false
	}
	return o.DLBOFunction, true
}

// HasDLBOFunction returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingle) HasDLBOFunction() bool {
	if o != nil && !IsNil(o.DLBOFunction) {
		return true
	}

	return false
}

// SetDLBOFunction gets a reference to the given DLBOFunctionSingle and assigns it to the DLBOFunction field.
func (o *GnbCuCpFunctionSingle) SetDLBOFunction(v DLBOFunctionSingle) {
	o.DLBOFunction = &v
}

func (o GnbCuCpFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GnbCuCpFunctionSingle) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.NrCellCu) {
		toSerialize["NrCellCu"] = o.NrCellCu
	}
	if !IsNil(o.EPXnC) {
		toSerialize["EP_XnC"] = o.EPXnC
	}
	if !IsNil(o.EPE1) {
		toSerialize["EP_E1"] = o.EPE1
	}
	if !IsNil(o.EPF1C) {
		toSerialize["EP_F1C"] = o.EPF1C
	}
	if !IsNil(o.EPNgC) {
		toSerialize["EP_NgC"] = o.EPNgC
	}
	if !IsNil(o.EPX2C) {
		toSerialize["EP_X2C"] = o.EPX2C
	}
	if !IsNil(o.DANRManagementFunction) {
		toSerialize["DANRManagementFunction"] = o.DANRManagementFunction
	}
	if !IsNil(o.DESManagementFunction) {
		toSerialize["DESManagementFunction"] = o.DESManagementFunction
	}
	if !IsNil(o.DMROFunction) {
		toSerialize["DMROFunction"] = o.DMROFunction
	}
	if !IsNil(o.DLBOFunction) {
		toSerialize["DLBOFunction"] = o.DLBOFunction
	}
	return toSerialize, nil
}

type NullableGnbCuCpFunctionSingle struct {
	value *GnbCuCpFunctionSingle
	isSet bool
}

func (v NullableGnbCuCpFunctionSingle) Get() *GnbCuCpFunctionSingle {
	return v.value
}

func (v *NullableGnbCuCpFunctionSingle) Set(val *GnbCuCpFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableGnbCuCpFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableGnbCuCpFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGnbCuCpFunctionSingle(val *GnbCuCpFunctionSingle) *NullableGnbCuCpFunctionSingle {
	return &NullableGnbCuCpFunctionSingle{value: val, isSet: true}
}

func (v NullableGnbCuCpFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGnbCuCpFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the CNSliceSubnetProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CNSliceSubnetProfile{}

// CNSliceSubnetProfile struct for CNSliceSubnetProfile
type CNSliceSubnetProfile struct {
	MaxNumberofUEs         *int32                  `json:"maxNumberofUEs,omitempty"`
	DLLatency              *float32                `json:"dLLatency,omitempty"`
	ULLatency              *float32                `json:"uLLatency,omitempty"`
	DLThptPerSliceSubnet   *XLThpt                 `json:"dLThptPerSliceSubnet,omitempty"`
	DLThptPerUE            *XLThpt                 `json:"dLThptPerUE,omitempty"`
	ULThptPerSliceSubnet   *XLThpt                 `json:"uLThptPerSliceSubnet,omitempty"`
	ULThptPerUE            *XLThpt                 `json:"uLThptPerUE,omitempty"`
	MaxNumberOfPDUSessions *int32                  `json:"maxNumberOfPDUSessions,omitempty"`
	CoverageAreaTAList     []int32                 `json:"coverageAreaTAList,omitempty"`
	ResourceSharingLevel   *SharingLevel           `json:"resourceSharingLevel,omitempty"`
	DLMaxPktSize           *int32                  `json:"dLMaxPktSize,omitempty"`
	ULMaxPktSize           *int32                  `json:"uLMaxPktSize,omitempty"`
	DelayTolerance         *DelayTolerance         `json:"delayTolerance,omitempty"`
	Synchronicity          *SynchronicityRANSubnet `json:"synchronicity,omitempty"`
	SliceSimultaneousUse   *SliceSimultaneousUse   `json:"sliceSimultaneousUse,omitempty"`
	Reliability            *float32                `json:"reliability,omitempty"`
	EnergyEfficiency       *float32                `json:"energyEfficiency,omitempty"`
	DLDeterministicComm    *DeterministicComm      `json:"dLDeterministicComm,omitempty"`
	ULDeterministicComm    *DeterministicComm      `json:"uLDeterministicComm,omitempty"`
	SurvivalTime           *float32                `json:"survivalTime,omitempty"`
	NssaaSupport           *NSSAASupport           `json:"nssaaSupport,omitempty"`
	N6Protection           *N6Protection           `json:"n6Protection,omitempty"`
}

// NewCNSliceSubnetProfile instantiates a new CNSliceSubnetProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCNSliceSubnetProfile() *CNSliceSubnetProfile {
	this := CNSliceSubnetProfile{}
	return &this
}

// NewCNSliceSubnetProfileWithDefaults instantiates a new CNSliceSubnetProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCNSliceSubnetProfileWithDefaults() *CNSliceSubnetProfile {
	this := CNSliceSubnetProfile{}
	return &this
}

// GetMaxNumberofUEs returns the MaxNumberofUEs field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetMaxNumberofUEs() int32 {
	if o == nil || IsNil(o.MaxNumberofUEs) {
		var ret int32
		return ret
	}
	return *o.MaxNumberofUEs
}

// GetMaxNumberofUEsOk returns a tuple with the MaxNumberofUEs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetMaxNumberofUEsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumberofUEs) {
		return nil, false
	}
	return o.MaxNumberofUEs, true
}

// HasMaxNumberofUEs returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasMaxNumberofUEs() bool {
	if o != nil && !IsNil(o.MaxNumberofUEs) {
		return true
	}

	return false
}

// SetMaxNumberofUEs gets a reference to the given int32 and assigns it to the MaxNumberofUEs field.
func (o *CNSliceSubnetProfile) SetMaxNumberofUEs(v int32) {
	o.MaxNumberofUEs = &v
}

// GetDLLatency returns the DLLatency field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetDLLatency() float32 {
	if o == nil || IsNil(o.DLLatency) {
		var ret float32
		return ret
	}
	return *o.DLLatency
}

// GetDLLatencyOk returns a tuple with the DLLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetDLLatencyOk() (*float32, bool) {
	if o == nil || IsNil(o.DLLatency) {
		return nil, false
	}
	return o.DLLatency, true
}

// HasDLLatency returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasDLLatency() bool {
	if o != nil && !IsNil(o.DLLatency) {
		return true
	}

	return false
}

// SetDLLatency gets a reference to the given float32 and assigns it to the DLLatency field.
func (o *CNSliceSubnetProfile) SetDLLatency(v float32) {
	o.DLLatency = &v
}

// GetULLatency returns the ULLatency field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetULLatency() float32 {
	if o == nil || IsNil(o.ULLatency) {
		var ret float32
		return ret
	}
	return *o.ULLatency
}

// GetULLatencyOk returns a tuple with the ULLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetULLatencyOk() (*float32, bool) {
	if o == nil || IsNil(o.ULLatency) {
		return nil, false
	}
	return o.ULLatency, true
}

// HasULLatency returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasULLatency() bool {
	if o != nil && !IsNil(o.ULLatency) {
		return true
	}

	return false
}

// SetULLatency gets a reference to the given float32 and assigns it to the ULLatency field.
func (o *CNSliceSubnetProfile) SetULLatency(v float32) {
	o.ULLatency = &v
}

// GetDLThptPerSliceSubnet returns the DLThptPerSliceSubnet field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetDLThptPerSliceSubnet() XLThpt {
	if o == nil || IsNil(o.DLThptPerSliceSubnet) {
		var ret XLThpt
		return ret
	}
	return *o.DLThptPerSliceSubnet
}

// GetDLThptPerSliceSubnetOk returns a tuple with the DLThptPerSliceSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetDLThptPerSliceSubnetOk() (*XLThpt, bool) {
	if o == nil || IsNil(o.DLThptPerSliceSubnet) {
		return nil, false
	}
	return o.DLThptPerSliceSubnet, true
}

// HasDLThptPerSliceSubnet returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasDLThptPerSliceSubnet() bool {
	if o != nil && !IsNil(o.DLThptPerSliceSubnet) {
		return true
	}

	return false
}

// SetDLThptPerSliceSubnet gets a reference to the given XLThpt and assigns it to the DLThptPerSliceSubnet field.
func (o *CNSliceSubnetProfile) SetDLThptPerSliceSubnet(v XLThpt) {
	o.DLThptPerSliceSubnet = &v
}

// GetDLThptPerUE returns the DLThptPerUE field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetDLThptPerUE() XLThpt {
	if o == nil || IsNil(o.DLThptPerUE) {
		var ret XLThpt
		return ret
	}
	return *o.DLThptPerUE
}

// GetDLThptPerUEOk returns a tuple with the DLThptPerUE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetDLThptPerUEOk() (*XLThpt, bool) {
	if o == nil || IsNil(o.DLThptPerUE) {
		return nil, false
	}
	return o.DLThptPerUE, true
}

// HasDLThptPerUE returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasDLThptPerUE() bool {
	if o != nil && !IsNil(o.DLThptPerUE) {
		return true
	}

	return false
}

// SetDLThptPerUE gets a reference to the given XLThpt and assigns it to the DLThptPerUE field.
func (o *CNSliceSubnetProfile) SetDLThptPerUE(v XLThpt) {
	o.DLThptPerUE = &v
}

// GetULThptPerSliceSubnet returns the ULThptPerSliceSubnet field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetULThptPerSliceSubnet() XLThpt {
	if o == nil || IsNil(o.ULThptPerSliceSubnet) {
		var ret XLThpt
		return ret
	}
	return *o.ULThptPerSliceSubnet
}

// GetULThptPerSliceSubnetOk returns a tuple with the ULThptPerSliceSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetULThptPerSliceSubnetOk() (*XLThpt, bool) {
	if o == nil || IsNil(o.ULThptPerSliceSubnet) {
		return nil, false
	}
	return o.ULThptPerSliceSubnet, true
}

// HasULThptPerSliceSubnet returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasULThptPerSliceSubnet() bool {
	if o != nil && !IsNil(o.ULThptPerSliceSubnet) {
		return true
	}

	return false
}

// SetULThptPerSliceSubnet gets a reference to the given XLThpt and assigns it to the ULThptPerSliceSubnet field.
func (o *CNSliceSubnetProfile) SetULThptPerSliceSubnet(v XLThpt) {
	o.ULThptPerSliceSubnet = &v
}

// GetULThptPerUE returns the ULThptPerUE field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetULThptPerUE() XLThpt {
	if o == nil || IsNil(o.ULThptPerUE) {
		var ret XLThpt
		return ret
	}
	return *o.ULThptPerUE
}

// GetULThptPerUEOk returns a tuple with the ULThptPerUE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetULThptPerUEOk() (*XLThpt, bool) {
	if o == nil || IsNil(o.ULThptPerUE) {
		return nil, false
	}
	return o.ULThptPerUE, true
}

// HasULThptPerUE returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasULThptPerUE() bool {
	if o != nil && !IsNil(o.ULThptPerUE) {
		return true
	}

	return false
}

// SetULThptPerUE gets a reference to the given XLThpt and assigns it to the ULThptPerUE field.
func (o *CNSliceSubnetProfile) SetULThptPerUE(v XLThpt) {
	o.ULThptPerUE = &v
}

// GetMaxNumberOfPDUSessions returns the MaxNumberOfPDUSessions field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetMaxNumberOfPDUSessions() int32 {
	if o == nil || IsNil(o.MaxNumberOfPDUSessions) {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfPDUSessions
}

// GetMaxNumberOfPDUSessionsOk returns a tuple with the MaxNumberOfPDUSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetMaxNumberOfPDUSessionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumberOfPDUSessions) {
		return nil, false
	}
	return o.MaxNumberOfPDUSessions, true
}

// HasMaxNumberOfPDUSessions returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasMaxNumberOfPDUSessions() bool {
	if o != nil && !IsNil(o.MaxNumberOfPDUSessions) {
		return true
	}

	return false
}

// SetMaxNumberOfPDUSessions gets a reference to the given int32 and assigns it to the MaxNumberOfPDUSessions field.
func (o *CNSliceSubnetProfile) SetMaxNumberOfPDUSessions(v int32) {
	o.MaxNumberOfPDUSessions = &v
}

// GetCoverageAreaTAList returns the CoverageAreaTAList field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetCoverageAreaTAList() []int32 {
	if o == nil || IsNil(o.CoverageAreaTAList) {
		var ret []int32
		return ret
	}
	return o.CoverageAreaTAList
}

// GetCoverageAreaTAListOk returns a tuple with the CoverageAreaTAList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetCoverageAreaTAListOk() ([]int32, bool) {
	if o == nil || IsNil(o.CoverageAreaTAList) {
		return nil, false
	}
	return o.CoverageAreaTAList, true
}

// HasCoverageAreaTAList returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasCoverageAreaTAList() bool {
	if o != nil && !IsNil(o.CoverageAreaTAList) {
		return true
	}

	return false
}

// SetCoverageAreaTAList gets a reference to the given []int32 and assigns it to the CoverageAreaTAList field.
func (o *CNSliceSubnetProfile) SetCoverageAreaTAList(v []int32) {
	o.CoverageAreaTAList = v
}

// GetResourceSharingLevel returns the ResourceSharingLevel field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetResourceSharingLevel() SharingLevel {
	if o == nil || IsNil(o.ResourceSharingLevel) {
		var ret SharingLevel
		return ret
	}
	return *o.ResourceSharingLevel
}

// GetResourceSharingLevelOk returns a tuple with the ResourceSharingLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetResourceSharingLevelOk() (*SharingLevel, bool) {
	if o == nil || IsNil(o.ResourceSharingLevel) {
		return nil, false
	}
	return o.ResourceSharingLevel, true
}

// HasResourceSharingLevel returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasResourceSharingLevel() bool {
	if o != nil && !IsNil(o.ResourceSharingLevel) {
		return true
	}

	return false
}

// SetResourceSharingLevel gets a reference to the given SharingLevel and assigns it to the ResourceSharingLevel field.
func (o *CNSliceSubnetProfile) SetResourceSharingLevel(v SharingLevel) {
	o.ResourceSharingLevel = &v
}

// GetDLMaxPktSize returns the DLMaxPktSize field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetDLMaxPktSize() int32 {
	if o == nil || IsNil(o.DLMaxPktSize) {
		var ret int32
		return ret
	}
	return *o.DLMaxPktSize
}

// GetDLMaxPktSizeOk returns a tuple with the DLMaxPktSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetDLMaxPktSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.DLMaxPktSize) {
		return nil, false
	}
	return o.DLMaxPktSize, true
}

// HasDLMaxPktSize returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasDLMaxPktSize() bool {
	if o != nil && !IsNil(o.DLMaxPktSize) {
		return true
	}

	return false
}

// SetDLMaxPktSize gets a reference to the given int32 and assigns it to the DLMaxPktSize field.
func (o *CNSliceSubnetProfile) SetDLMaxPktSize(v int32) {
	o.DLMaxPktSize = &v
}

// GetULMaxPktSize returns the ULMaxPktSize field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetULMaxPktSize() int32 {
	if o == nil || IsNil(o.ULMaxPktSize) {
		var ret int32
		return ret
	}
	return *o.ULMaxPktSize
}

// GetULMaxPktSizeOk returns a tuple with the ULMaxPktSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetULMaxPktSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.ULMaxPktSize) {
		return nil, false
	}
	return o.ULMaxPktSize, true
}

// HasULMaxPktSize returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasULMaxPktSize() bool {
	if o != nil && !IsNil(o.ULMaxPktSize) {
		return true
	}

	return false
}

// SetULMaxPktSize gets a reference to the given int32 and assigns it to the ULMaxPktSize field.
func (o *CNSliceSubnetProfile) SetULMaxPktSize(v int32) {
	o.ULMaxPktSize = &v
}

// GetDelayTolerance returns the DelayTolerance field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetDelayTolerance() DelayTolerance {
	if o == nil || IsNil(o.DelayTolerance) {
		var ret DelayTolerance
		return ret
	}
	return *o.DelayTolerance
}

// GetDelayToleranceOk returns a tuple with the DelayTolerance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetDelayToleranceOk() (*DelayTolerance, bool) {
	if o == nil || IsNil(o.DelayTolerance) {
		return nil, false
	}
	return o.DelayTolerance, true
}

// HasDelayTolerance returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasDelayTolerance() bool {
	if o != nil && !IsNil(o.DelayTolerance) {
		return true
	}

	return false
}

// SetDelayTolerance gets a reference to the given DelayTolerance and assigns it to the DelayTolerance field.
func (o *CNSliceSubnetProfile) SetDelayTolerance(v DelayTolerance) {
	o.DelayTolerance = &v
}

// GetSynchronicity returns the Synchronicity field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetSynchronicity() SynchronicityRANSubnet {
	if o == nil || IsNil(o.Synchronicity) {
		var ret SynchronicityRANSubnet
		return ret
	}
	return *o.Synchronicity
}

// GetSynchronicityOk returns a tuple with the Synchronicity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetSynchronicityOk() (*SynchronicityRANSubnet, bool) {
	if o == nil || IsNil(o.Synchronicity) {
		return nil, false
	}
	return o.Synchronicity, true
}

// HasSynchronicity returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasSynchronicity() bool {
	if o != nil && !IsNil(o.Synchronicity) {
		return true
	}

	return false
}

// SetSynchronicity gets a reference to the given SynchronicityRANSubnet and assigns it to the Synchronicity field.
func (o *CNSliceSubnetProfile) SetSynchronicity(v SynchronicityRANSubnet) {
	o.Synchronicity = &v
}

// GetSliceSimultaneousUse returns the SliceSimultaneousUse field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetSliceSimultaneousUse() SliceSimultaneousUse {
	if o == nil || IsNil(o.SliceSimultaneousUse) {
		var ret SliceSimultaneousUse
		return ret
	}
	return *o.SliceSimultaneousUse
}

// GetSliceSimultaneousUseOk returns a tuple with the SliceSimultaneousUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetSliceSimultaneousUseOk() (*SliceSimultaneousUse, bool) {
	if o == nil || IsNil(o.SliceSimultaneousUse) {
		return nil, false
	}
	return o.SliceSimultaneousUse, true
}

// HasSliceSimultaneousUse returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasSliceSimultaneousUse() bool {
	if o != nil && !IsNil(o.SliceSimultaneousUse) {
		return true
	}

	return false
}

// SetSliceSimultaneousUse gets a reference to the given SliceSimultaneousUse and assigns it to the SliceSimultaneousUse field.
func (o *CNSliceSubnetProfile) SetSliceSimultaneousUse(v SliceSimultaneousUse) {
	o.SliceSimultaneousUse = &v
}

// GetReliability returns the Reliability field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetReliability() float32 {
	if o == nil || IsNil(o.Reliability) {
		var ret float32
		return ret
	}
	return *o.Reliability
}

// GetReliabilityOk returns a tuple with the Reliability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetReliabilityOk() (*float32, bool) {
	if o == nil || IsNil(o.Reliability) {
		return nil, false
	}
	return o.Reliability, true
}

// HasReliability returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasReliability() bool {
	if o != nil && !IsNil(o.Reliability) {
		return true
	}

	return false
}

// SetReliability gets a reference to the given float32 and assigns it to the Reliability field.
func (o *CNSliceSubnetProfile) SetReliability(v float32) {
	o.Reliability = &v
}

// GetEnergyEfficiency returns the EnergyEfficiency field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetEnergyEfficiency() float32 {
	if o == nil || IsNil(o.EnergyEfficiency) {
		var ret float32
		return ret
	}
	return *o.EnergyEfficiency
}

// GetEnergyEfficiencyOk returns a tuple with the EnergyEfficiency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetEnergyEfficiencyOk() (*float32, bool) {
	if o == nil || IsNil(o.EnergyEfficiency) {
		return nil, false
	}
	return o.EnergyEfficiency, true
}

// HasEnergyEfficiency returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasEnergyEfficiency() bool {
	if o != nil && !IsNil(o.EnergyEfficiency) {
		return true
	}

	return false
}

// SetEnergyEfficiency gets a reference to the given float32 and assigns it to the EnergyEfficiency field.
func (o *CNSliceSubnetProfile) SetEnergyEfficiency(v float32) {
	o.EnergyEfficiency = &v
}

// GetDLDeterministicComm returns the DLDeterministicComm field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetDLDeterministicComm() DeterministicComm {
	if o == nil || IsNil(o.DLDeterministicComm) {
		var ret DeterministicComm
		return ret
	}
	return *o.DLDeterministicComm
}

// GetDLDeterministicCommOk returns a tuple with the DLDeterministicComm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetDLDeterministicCommOk() (*DeterministicComm, bool) {
	if o == nil || IsNil(o.DLDeterministicComm) {
		return nil, false
	}
	return o.DLDeterministicComm, true
}

// HasDLDeterministicComm returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasDLDeterministicComm() bool {
	if o != nil && !IsNil(o.DLDeterministicComm) {
		return true
	}

	return false
}

// SetDLDeterministicComm gets a reference to the given DeterministicComm and assigns it to the DLDeterministicComm field.
func (o *CNSliceSubnetProfile) SetDLDeterministicComm(v DeterministicComm) {
	o.DLDeterministicComm = &v
}

// GetULDeterministicComm returns the ULDeterministicComm field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetULDeterministicComm() DeterministicComm {
	if o == nil || IsNil(o.ULDeterministicComm) {
		var ret DeterministicComm
		return ret
	}
	return *o.ULDeterministicComm
}

// GetULDeterministicCommOk returns a tuple with the ULDeterministicComm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetULDeterministicCommOk() (*DeterministicComm, bool) {
	if o == nil || IsNil(o.ULDeterministicComm) {
		return nil, false
	}
	return o.ULDeterministicComm, true
}

// HasULDeterministicComm returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasULDeterministicComm() bool {
	if o != nil && !IsNil(o.ULDeterministicComm) {
		return true
	}

	return false
}

// SetULDeterministicComm gets a reference to the given DeterministicComm and assigns it to the ULDeterministicComm field.
func (o *CNSliceSubnetProfile) SetULDeterministicComm(v DeterministicComm) {
	o.ULDeterministicComm = &v
}

// GetSurvivalTime returns the SurvivalTime field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetSurvivalTime() float32 {
	if o == nil || IsNil(o.SurvivalTime) {
		var ret float32
		return ret
	}
	return *o.SurvivalTime
}

// GetSurvivalTimeOk returns a tuple with the SurvivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetSurvivalTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.SurvivalTime) {
		return nil, false
	}
	return o.SurvivalTime, true
}

// HasSurvivalTime returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasSurvivalTime() bool {
	if o != nil && !IsNil(o.SurvivalTime) {
		return true
	}

	return false
}

// SetSurvivalTime gets a reference to the given float32 and assigns it to the SurvivalTime field.
func (o *CNSliceSubnetProfile) SetSurvivalTime(v float32) {
	o.SurvivalTime = &v
}

// GetNssaaSupport returns the NssaaSupport field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetNssaaSupport() NSSAASupport {
	if o == nil || IsNil(o.NssaaSupport) {
		var ret NSSAASupport
		return ret
	}
	return *o.NssaaSupport
}

// GetNssaaSupportOk returns a tuple with the NssaaSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetNssaaSupportOk() (*NSSAASupport, bool) {
	if o == nil || IsNil(o.NssaaSupport) {
		return nil, false
	}
	return o.NssaaSupport, true
}

// HasNssaaSupport returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasNssaaSupport() bool {
	if o != nil && !IsNil(o.NssaaSupport) {
		return true
	}

	return false
}

// SetNssaaSupport gets a reference to the given NSSAASupport and assigns it to the NssaaSupport field.
func (o *CNSliceSubnetProfile) SetNssaaSupport(v NSSAASupport) {
	o.NssaaSupport = &v
}

// GetN6Protection returns the N6Protection field value if set, zero value otherwise.
func (o *CNSliceSubnetProfile) GetN6Protection() N6Protection {
	if o == nil || IsNil(o.N6Protection) {
		var ret N6Protection
		return ret
	}
	return *o.N6Protection
}

// GetN6ProtectionOk returns a tuple with the N6Protection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNSliceSubnetProfile) GetN6ProtectionOk() (*N6Protection, bool) {
	if o == nil || IsNil(o.N6Protection) {
		return nil, false
	}
	return o.N6Protection, true
}

// HasN6Protection returns a boolean if a field has been set.
func (o *CNSliceSubnetProfile) HasN6Protection() bool {
	if o != nil && !IsNil(o.N6Protection) {
		return true
	}

	return false
}

// SetN6Protection gets a reference to the given N6Protection and assigns it to the N6Protection field.
func (o *CNSliceSubnetProfile) SetN6Protection(v N6Protection) {
	o.N6Protection = &v
}

func (o CNSliceSubnetProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CNSliceSubnetProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxNumberofUEs) {
		toSerialize["maxNumberofUEs"] = o.MaxNumberofUEs
	}
	if !IsNil(o.DLLatency) {
		toSerialize["dLLatency"] = o.DLLatency
	}
	if !IsNil(o.ULLatency) {
		toSerialize["uLLatency"] = o.ULLatency
	}
	if !IsNil(o.DLThptPerSliceSubnet) {
		toSerialize["dLThptPerSliceSubnet"] = o.DLThptPerSliceSubnet
	}
	if !IsNil(o.DLThptPerUE) {
		toSerialize["dLThptPerUE"] = o.DLThptPerUE
	}
	if !IsNil(o.ULThptPerSliceSubnet) {
		toSerialize["uLThptPerSliceSubnet"] = o.ULThptPerSliceSubnet
	}
	if !IsNil(o.ULThptPerUE) {
		toSerialize["uLThptPerUE"] = o.ULThptPerUE
	}
	if !IsNil(o.MaxNumberOfPDUSessions) {
		toSerialize["maxNumberOfPDUSessions"] = o.MaxNumberOfPDUSessions
	}
	if !IsNil(o.CoverageAreaTAList) {
		toSerialize["coverageAreaTAList"] = o.CoverageAreaTAList
	}
	if !IsNil(o.ResourceSharingLevel) {
		toSerialize["resourceSharingLevel"] = o.ResourceSharingLevel
	}
	if !IsNil(o.DLMaxPktSize) {
		toSerialize["dLMaxPktSize"] = o.DLMaxPktSize
	}
	if !IsNil(o.ULMaxPktSize) {
		toSerialize["uLMaxPktSize"] = o.ULMaxPktSize
	}
	if !IsNil(o.DelayTolerance) {
		toSerialize["delayTolerance"] = o.DelayTolerance
	}
	if !IsNil(o.Synchronicity) {
		toSerialize["synchronicity"] = o.Synchronicity
	}
	if !IsNil(o.SliceSimultaneousUse) {
		toSerialize["sliceSimultaneousUse"] = o.SliceSimultaneousUse
	}
	if !IsNil(o.Reliability) {
		toSerialize["reliability"] = o.Reliability
	}
	if !IsNil(o.EnergyEfficiency) {
		toSerialize["energyEfficiency"] = o.EnergyEfficiency
	}
	if !IsNil(o.DLDeterministicComm) {
		toSerialize["dLDeterministicComm"] = o.DLDeterministicComm
	}
	if !IsNil(o.ULDeterministicComm) {
		toSerialize["uLDeterministicComm"] = o.ULDeterministicComm
	}
	if !IsNil(o.SurvivalTime) {
		toSerialize["survivalTime"] = o.SurvivalTime
	}
	if !IsNil(o.NssaaSupport) {
		toSerialize["nssaaSupport"] = o.NssaaSupport
	}
	if !IsNil(o.N6Protection) {
		toSerialize["n6Protection"] = o.N6Protection
	}
	return toSerialize, nil
}

type NullableCNSliceSubnetProfile struct {
	value *CNSliceSubnetProfile
	isSet bool
}

func (v NullableCNSliceSubnetProfile) Get() *CNSliceSubnetProfile {
	return v.value
}

func (v *NullableCNSliceSubnetProfile) Set(val *CNSliceSubnetProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableCNSliceSubnetProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableCNSliceSubnetProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCNSliceSubnetProfile(val *CNSliceSubnetProfile) *NullableCNSliceSubnetProfile {
	return &NullableCNSliceSubnetProfile{value: val, isSet: true}
}

func (v NullableCNSliceSubnetProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCNSliceSubnetProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

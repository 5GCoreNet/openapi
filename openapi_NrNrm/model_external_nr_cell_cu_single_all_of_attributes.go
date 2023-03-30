/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the ExternalNrCellCuSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalNrCellCuSingleAllOfAttributes{}

// ExternalNrCellCuSingleAllOfAttributes struct for ExternalNrCellCuSingleAllOfAttributes
type ExternalNrCellCuSingleAllOfAttributes struct {
	UserLabel *string `json:"userLabel,omitempty"`
	VnfParametersList []VnfParameter `json:"vnfParametersList,omitempty"`
	PeeParametersList []PeeParameter `json:"peeParametersList,omitempty"`
	PriorityLabel *int32 `json:"priorityLabel,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
	CellLocalId *int32 `json:"cellLocalId,omitempty"`
	NrPci *int32 `json:"nrPci,omitempty"`
	PlmnIdList []PlmnId `json:"plmnIdList,omitempty"`
	NRFrequencyRef *string `json:"nRFrequencyRef,omitempty"`
}

// NewExternalNrCellCuSingleAllOfAttributes instantiates a new ExternalNrCellCuSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalNrCellCuSingleAllOfAttributes() *ExternalNrCellCuSingleAllOfAttributes {
	this := ExternalNrCellCuSingleAllOfAttributes{}
	return &this
}

// NewExternalNrCellCuSingleAllOfAttributesWithDefaults instantiates a new ExternalNrCellCuSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalNrCellCuSingleAllOfAttributesWithDefaults() *ExternalNrCellCuSingleAllOfAttributes {
	this := ExternalNrCellCuSingleAllOfAttributes{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ExternalNrCellCuSingleAllOfAttributes) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetVnfParametersList returns the VnfParametersList field value if set, zero value otherwise.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetVnfParametersList() []VnfParameter {
	if o == nil || IsNil(o.VnfParametersList) {
		var ret []VnfParameter
		return ret
	}
	return o.VnfParametersList
}

// GetVnfParametersListOk returns a tuple with the VnfParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetVnfParametersListOk() ([]VnfParameter, bool) {
	if o == nil || IsNil(o.VnfParametersList) {
		return nil, false
	}
	return o.VnfParametersList, true
}

// HasVnfParametersList returns a boolean if a field has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) HasVnfParametersList() bool {
	if o != nil && !IsNil(o.VnfParametersList) {
		return true
	}

	return false
}

// SetVnfParametersList gets a reference to the given []VnfParameter and assigns it to the VnfParametersList field.
func (o *ExternalNrCellCuSingleAllOfAttributes) SetVnfParametersList(v []VnfParameter) {
	o.VnfParametersList = v
}

// GetPeeParametersList returns the PeeParametersList field value if set, zero value otherwise.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetPeeParametersList() []PeeParameter {
	if o == nil || IsNil(o.PeeParametersList) {
		var ret []PeeParameter
		return ret
	}
	return o.PeeParametersList
}

// GetPeeParametersListOk returns a tuple with the PeeParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetPeeParametersListOk() ([]PeeParameter, bool) {
	if o == nil || IsNil(o.PeeParametersList) {
		return nil, false
	}
	return o.PeeParametersList, true
}

// HasPeeParametersList returns a boolean if a field has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) HasPeeParametersList() bool {
	if o != nil && !IsNil(o.PeeParametersList) {
		return true
	}

	return false
}

// SetPeeParametersList gets a reference to the given []PeeParameter and assigns it to the PeeParametersList field.
func (o *ExternalNrCellCuSingleAllOfAttributes) SetPeeParametersList(v []PeeParameter) {
	o.PeeParametersList = v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetPriorityLabel() int32 {
	if o == nil || IsNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLabel) {
		return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) HasPriorityLabel() bool {
	if o != nil && !IsNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *ExternalNrCellCuSingleAllOfAttributes) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) HasSupportedPerfMetricGroups() bool {
	if o != nil && !IsNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *ExternalNrCellCuSingleAllOfAttributes) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetSupportedTraceMetrics returns the SupportedTraceMetrics field value if set, zero value otherwise.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetSupportedTraceMetrics() []string {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedTraceMetrics
}

// GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetSupportedTraceMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		return nil, false
	}
	return o.SupportedTraceMetrics, true
}

// HasSupportedTraceMetrics returns a boolean if a field has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) HasSupportedTraceMetrics() bool {
	if o != nil && !IsNil(o.SupportedTraceMetrics) {
		return true
	}

	return false
}

// SetSupportedTraceMetrics gets a reference to the given []string and assigns it to the SupportedTraceMetrics field.
func (o *ExternalNrCellCuSingleAllOfAttributes) SetSupportedTraceMetrics(v []string) {
	o.SupportedTraceMetrics = v
}

// GetCellLocalId returns the CellLocalId field value if set, zero value otherwise.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetCellLocalId() int32 {
	if o == nil || IsNil(o.CellLocalId) {
		var ret int32
		return ret
	}
	return *o.CellLocalId
}

// GetCellLocalIdOk returns a tuple with the CellLocalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetCellLocalIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CellLocalId) {
		return nil, false
	}
	return o.CellLocalId, true
}

// HasCellLocalId returns a boolean if a field has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) HasCellLocalId() bool {
	if o != nil && !IsNil(o.CellLocalId) {
		return true
	}

	return false
}

// SetCellLocalId gets a reference to the given int32 and assigns it to the CellLocalId field.
func (o *ExternalNrCellCuSingleAllOfAttributes) SetCellLocalId(v int32) {
	o.CellLocalId = &v
}

// GetNrPci returns the NrPci field value if set, zero value otherwise.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetNrPci() int32 {
	if o == nil || IsNil(o.NrPci) {
		var ret int32
		return ret
	}
	return *o.NrPci
}

// GetNrPciOk returns a tuple with the NrPci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetNrPciOk() (*int32, bool) {
	if o == nil || IsNil(o.NrPci) {
		return nil, false
	}
	return o.NrPci, true
}

// HasNrPci returns a boolean if a field has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) HasNrPci() bool {
	if o != nil && !IsNil(o.NrPci) {
		return true
	}

	return false
}

// SetNrPci gets a reference to the given int32 and assigns it to the NrPci field.
func (o *ExternalNrCellCuSingleAllOfAttributes) SetNrPci(v int32) {
	o.NrPci = &v
}

// GetPlmnIdList returns the PlmnIdList field value if set, zero value otherwise.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetPlmnIdList() []PlmnId {
	if o == nil || IsNil(o.PlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.PlmnIdList
}

// GetPlmnIdListOk returns a tuple with the PlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.PlmnIdList) {
		return nil, false
	}
	return o.PlmnIdList, true
}

// HasPlmnIdList returns a boolean if a field has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) HasPlmnIdList() bool {
	if o != nil && !IsNil(o.PlmnIdList) {
		return true
	}

	return false
}

// SetPlmnIdList gets a reference to the given []PlmnId and assigns it to the PlmnIdList field.
func (o *ExternalNrCellCuSingleAllOfAttributes) SetPlmnIdList(v []PlmnId) {
	o.PlmnIdList = v
}

// GetNRFrequencyRef returns the NRFrequencyRef field value if set, zero value otherwise.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetNRFrequencyRef() string {
	if o == nil || IsNil(o.NRFrequencyRef) {
		var ret string
		return ret
	}
	return *o.NRFrequencyRef
}

// GetNRFrequencyRefOk returns a tuple with the NRFrequencyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) GetNRFrequencyRefOk() (*string, bool) {
	if o == nil || IsNil(o.NRFrequencyRef) {
		return nil, false
	}
	return o.NRFrequencyRef, true
}

// HasNRFrequencyRef returns a boolean if a field has been set.
func (o *ExternalNrCellCuSingleAllOfAttributes) HasNRFrequencyRef() bool {
	if o != nil && !IsNil(o.NRFrequencyRef) {
		return true
	}

	return false
}

// SetNRFrequencyRef gets a reference to the given string and assigns it to the NRFrequencyRef field.
func (o *ExternalNrCellCuSingleAllOfAttributes) SetNRFrequencyRef(v string) {
	o.NRFrequencyRef = &v
}

func (o ExternalNrCellCuSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalNrCellCuSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserLabel) {
		toSerialize["userLabel"] = o.UserLabel
	}
	if !IsNil(o.VnfParametersList) {
		toSerialize["vnfParametersList"] = o.VnfParametersList
	}
	if !IsNil(o.PeeParametersList) {
		toSerialize["peeParametersList"] = o.PeeParametersList
	}
	if !IsNil(o.PriorityLabel) {
		toSerialize["priorityLabel"] = o.PriorityLabel
	}
	if !IsNil(o.SupportedPerfMetricGroups) {
		toSerialize["supportedPerfMetricGroups"] = o.SupportedPerfMetricGroups
	}
	if !IsNil(o.SupportedTraceMetrics) {
		toSerialize["supportedTraceMetrics"] = o.SupportedTraceMetrics
	}
	if !IsNil(o.CellLocalId) {
		toSerialize["cellLocalId"] = o.CellLocalId
	}
	if !IsNil(o.NrPci) {
		toSerialize["nrPci"] = o.NrPci
	}
	if !IsNil(o.PlmnIdList) {
		toSerialize["plmnIdList"] = o.PlmnIdList
	}
	if !IsNil(o.NRFrequencyRef) {
		toSerialize["nRFrequencyRef"] = o.NRFrequencyRef
	}
	return toSerialize, nil
}

type NullableExternalNrCellCuSingleAllOfAttributes struct {
	value *ExternalNrCellCuSingleAllOfAttributes
	isSet bool
}

func (v NullableExternalNrCellCuSingleAllOfAttributes) Get() *ExternalNrCellCuSingleAllOfAttributes {
	return v.value
}

func (v *NullableExternalNrCellCuSingleAllOfAttributes) Set(val *ExternalNrCellCuSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalNrCellCuSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalNrCellCuSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalNrCellCuSingleAllOfAttributes(val *ExternalNrCellCuSingleAllOfAttributes) *NullableExternalNrCellCuSingleAllOfAttributes {
	return &NullableExternalNrCellCuSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableExternalNrCellCuSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalNrCellCuSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



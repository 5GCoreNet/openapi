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

// checks if the NrCellCuSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NrCellCuSingleAllOfAttributes{}

// NrCellCuSingleAllOfAttributes struct for NrCellCuSingleAllOfAttributes
type NrCellCuSingleAllOfAttributes struct {
	UserLabel *string `json:"userLabel,omitempty"`
	VnfParametersList []VnfParameter `json:"vnfParametersList,omitempty"`
	PeeParametersList []PeeParameter `json:"peeParametersList,omitempty"`
	PriorityLabel *int32 `json:"priorityLabel,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
	CellLocalId *int32 `json:"cellLocalId,omitempty"`
	PlmnInfoList []PlmnInfo `json:"plmnInfoList,omitempty"`
	NRFrequencyRef *string `json:"nRFrequencyRef,omitempty"`
}

// NewNrCellCuSingleAllOfAttributes instantiates a new NrCellCuSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrCellCuSingleAllOfAttributes() *NrCellCuSingleAllOfAttributes {
	this := NrCellCuSingleAllOfAttributes{}
	return &this
}

// NewNrCellCuSingleAllOfAttributesWithDefaults instantiates a new NrCellCuSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrCellCuSingleAllOfAttributesWithDefaults() *NrCellCuSingleAllOfAttributes {
	this := NrCellCuSingleAllOfAttributes{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOfAttributes) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOfAttributes) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOfAttributes) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *NrCellCuSingleAllOfAttributes) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetVnfParametersList returns the VnfParametersList field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOfAttributes) GetVnfParametersList() []VnfParameter {
	if o == nil || IsNil(o.VnfParametersList) {
		var ret []VnfParameter
		return ret
	}
	return o.VnfParametersList
}

// GetVnfParametersListOk returns a tuple with the VnfParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOfAttributes) GetVnfParametersListOk() ([]VnfParameter, bool) {
	if o == nil || IsNil(o.VnfParametersList) {
		return nil, false
	}
	return o.VnfParametersList, true
}

// HasVnfParametersList returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOfAttributes) HasVnfParametersList() bool {
	if o != nil && !IsNil(o.VnfParametersList) {
		return true
	}

	return false
}

// SetVnfParametersList gets a reference to the given []VnfParameter and assigns it to the VnfParametersList field.
func (o *NrCellCuSingleAllOfAttributes) SetVnfParametersList(v []VnfParameter) {
	o.VnfParametersList = v
}

// GetPeeParametersList returns the PeeParametersList field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOfAttributes) GetPeeParametersList() []PeeParameter {
	if o == nil || IsNil(o.PeeParametersList) {
		var ret []PeeParameter
		return ret
	}
	return o.PeeParametersList
}

// GetPeeParametersListOk returns a tuple with the PeeParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOfAttributes) GetPeeParametersListOk() ([]PeeParameter, bool) {
	if o == nil || IsNil(o.PeeParametersList) {
		return nil, false
	}
	return o.PeeParametersList, true
}

// HasPeeParametersList returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOfAttributes) HasPeeParametersList() bool {
	if o != nil && !IsNil(o.PeeParametersList) {
		return true
	}

	return false
}

// SetPeeParametersList gets a reference to the given []PeeParameter and assigns it to the PeeParametersList field.
func (o *NrCellCuSingleAllOfAttributes) SetPeeParametersList(v []PeeParameter) {
	o.PeeParametersList = v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOfAttributes) GetPriorityLabel() int32 {
	if o == nil || IsNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOfAttributes) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLabel) {
		return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOfAttributes) HasPriorityLabel() bool {
	if o != nil && !IsNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *NrCellCuSingleAllOfAttributes) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOfAttributes) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOfAttributes) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOfAttributes) HasSupportedPerfMetricGroups() bool {
	if o != nil && !IsNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *NrCellCuSingleAllOfAttributes) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetSupportedTraceMetrics returns the SupportedTraceMetrics field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOfAttributes) GetSupportedTraceMetrics() []string {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedTraceMetrics
}

// GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOfAttributes) GetSupportedTraceMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		return nil, false
	}
	return o.SupportedTraceMetrics, true
}

// HasSupportedTraceMetrics returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOfAttributes) HasSupportedTraceMetrics() bool {
	if o != nil && !IsNil(o.SupportedTraceMetrics) {
		return true
	}

	return false
}

// SetSupportedTraceMetrics gets a reference to the given []string and assigns it to the SupportedTraceMetrics field.
func (o *NrCellCuSingleAllOfAttributes) SetSupportedTraceMetrics(v []string) {
	o.SupportedTraceMetrics = v
}

// GetCellLocalId returns the CellLocalId field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOfAttributes) GetCellLocalId() int32 {
	if o == nil || IsNil(o.CellLocalId) {
		var ret int32
		return ret
	}
	return *o.CellLocalId
}

// GetCellLocalIdOk returns a tuple with the CellLocalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOfAttributes) GetCellLocalIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CellLocalId) {
		return nil, false
	}
	return o.CellLocalId, true
}

// HasCellLocalId returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOfAttributes) HasCellLocalId() bool {
	if o != nil && !IsNil(o.CellLocalId) {
		return true
	}

	return false
}

// SetCellLocalId gets a reference to the given int32 and assigns it to the CellLocalId field.
func (o *NrCellCuSingleAllOfAttributes) SetCellLocalId(v int32) {
	o.CellLocalId = &v
}

// GetPlmnInfoList returns the PlmnInfoList field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOfAttributes) GetPlmnInfoList() []PlmnInfo {
	if o == nil || IsNil(o.PlmnInfoList) {
		var ret []PlmnInfo
		return ret
	}
	return o.PlmnInfoList
}

// GetPlmnInfoListOk returns a tuple with the PlmnInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOfAttributes) GetPlmnInfoListOk() ([]PlmnInfo, bool) {
	if o == nil || IsNil(o.PlmnInfoList) {
		return nil, false
	}
	return o.PlmnInfoList, true
}

// HasPlmnInfoList returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOfAttributes) HasPlmnInfoList() bool {
	if o != nil && !IsNil(o.PlmnInfoList) {
		return true
	}

	return false
}

// SetPlmnInfoList gets a reference to the given []PlmnInfo and assigns it to the PlmnInfoList field.
func (o *NrCellCuSingleAllOfAttributes) SetPlmnInfoList(v []PlmnInfo) {
	o.PlmnInfoList = v
}

// GetNRFrequencyRef returns the NRFrequencyRef field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOfAttributes) GetNRFrequencyRef() string {
	if o == nil || IsNil(o.NRFrequencyRef) {
		var ret string
		return ret
	}
	return *o.NRFrequencyRef
}

// GetNRFrequencyRefOk returns a tuple with the NRFrequencyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOfAttributes) GetNRFrequencyRefOk() (*string, bool) {
	if o == nil || IsNil(o.NRFrequencyRef) {
		return nil, false
	}
	return o.NRFrequencyRef, true
}

// HasNRFrequencyRef returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOfAttributes) HasNRFrequencyRef() bool {
	if o != nil && !IsNil(o.NRFrequencyRef) {
		return true
	}

	return false
}

// SetNRFrequencyRef gets a reference to the given string and assigns it to the NRFrequencyRef field.
func (o *NrCellCuSingleAllOfAttributes) SetNRFrequencyRef(v string) {
	o.NRFrequencyRef = &v
}

func (o NrCellCuSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NrCellCuSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PlmnInfoList) {
		toSerialize["plmnInfoList"] = o.PlmnInfoList
	}
	if !IsNil(o.NRFrequencyRef) {
		toSerialize["nRFrequencyRef"] = o.NRFrequencyRef
	}
	return toSerialize, nil
}

type NullableNrCellCuSingleAllOfAttributes struct {
	value *NrCellCuSingleAllOfAttributes
	isSet bool
}

func (v NullableNrCellCuSingleAllOfAttributes) Get() *NrCellCuSingleAllOfAttributes {
	return v.value
}

func (v *NullableNrCellCuSingleAllOfAttributes) Set(val *NrCellCuSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNrCellCuSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNrCellCuSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrCellCuSingleAllOfAttributes(val *NrCellCuSingleAllOfAttributes) *NullableNrCellCuSingleAllOfAttributes {
	return &NullableNrCellCuSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableNrCellCuSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrCellCuSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



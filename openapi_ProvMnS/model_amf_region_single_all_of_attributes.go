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

// checks if the AmfRegionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfRegionSingleAllOfAttributes{}

// AmfRegionSingleAllOfAttributes struct for AmfRegionSingleAllOfAttributes
type AmfRegionSingleAllOfAttributes struct {
	UserLabel *string `json:"userLabel,omitempty"`
	VnfParametersList []VnfParameter `json:"vnfParametersList,omitempty"`
	PeeParametersList []PeeParameter `json:"peeParametersList,omitempty"`
	PriorityLabel *int32 `json:"priorityLabel,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
	PlmnIdList []PlmnId `json:"plmnIdList,omitempty"`
	NRTACList []int32 `json:"nRTACList,omitempty"`
	// AmfRegionId is defined in TS 23.003
	AmfRegionId *int32 `json:"amfRegionId,omitempty"`
	SnssaiList []Snssai `json:"snssaiList,omitempty"`
	AMFSetListRef []string `json:"aMFSetListRef,omitempty"`
}

// NewAmfRegionSingleAllOfAttributes instantiates a new AmfRegionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfRegionSingleAllOfAttributes() *AmfRegionSingleAllOfAttributes {
	this := AmfRegionSingleAllOfAttributes{}
	return &this
}

// NewAmfRegionSingleAllOfAttributesWithDefaults instantiates a new AmfRegionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfRegionSingleAllOfAttributesWithDefaults() *AmfRegionSingleAllOfAttributes {
	this := AmfRegionSingleAllOfAttributes{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributes) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributes) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributes) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *AmfRegionSingleAllOfAttributes) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetVnfParametersList returns the VnfParametersList field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributes) GetVnfParametersList() []VnfParameter {
	if o == nil || IsNil(o.VnfParametersList) {
		var ret []VnfParameter
		return ret
	}
	return o.VnfParametersList
}

// GetVnfParametersListOk returns a tuple with the VnfParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributes) GetVnfParametersListOk() ([]VnfParameter, bool) {
	if o == nil || IsNil(o.VnfParametersList) {
		return nil, false
	}
	return o.VnfParametersList, true
}

// HasVnfParametersList returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributes) HasVnfParametersList() bool {
	if o != nil && !IsNil(o.VnfParametersList) {
		return true
	}

	return false
}

// SetVnfParametersList gets a reference to the given []VnfParameter and assigns it to the VnfParametersList field.
func (o *AmfRegionSingleAllOfAttributes) SetVnfParametersList(v []VnfParameter) {
	o.VnfParametersList = v
}

// GetPeeParametersList returns the PeeParametersList field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributes) GetPeeParametersList() []PeeParameter {
	if o == nil || IsNil(o.PeeParametersList) {
		var ret []PeeParameter
		return ret
	}
	return o.PeeParametersList
}

// GetPeeParametersListOk returns a tuple with the PeeParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributes) GetPeeParametersListOk() ([]PeeParameter, bool) {
	if o == nil || IsNil(o.PeeParametersList) {
		return nil, false
	}
	return o.PeeParametersList, true
}

// HasPeeParametersList returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributes) HasPeeParametersList() bool {
	if o != nil && !IsNil(o.PeeParametersList) {
		return true
	}

	return false
}

// SetPeeParametersList gets a reference to the given []PeeParameter and assigns it to the PeeParametersList field.
func (o *AmfRegionSingleAllOfAttributes) SetPeeParametersList(v []PeeParameter) {
	o.PeeParametersList = v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributes) GetPriorityLabel() int32 {
	if o == nil || IsNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributes) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLabel) {
		return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributes) HasPriorityLabel() bool {
	if o != nil && !IsNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *AmfRegionSingleAllOfAttributes) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributes) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributes) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributes) HasSupportedPerfMetricGroups() bool {
	if o != nil && !IsNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *AmfRegionSingleAllOfAttributes) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetSupportedTraceMetrics returns the SupportedTraceMetrics field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributes) GetSupportedTraceMetrics() []string {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedTraceMetrics
}

// GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributes) GetSupportedTraceMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		return nil, false
	}
	return o.SupportedTraceMetrics, true
}

// HasSupportedTraceMetrics returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributes) HasSupportedTraceMetrics() bool {
	if o != nil && !IsNil(o.SupportedTraceMetrics) {
		return true
	}

	return false
}

// SetSupportedTraceMetrics gets a reference to the given []string and assigns it to the SupportedTraceMetrics field.
func (o *AmfRegionSingleAllOfAttributes) SetSupportedTraceMetrics(v []string) {
	o.SupportedTraceMetrics = v
}

// GetPlmnIdList returns the PlmnIdList field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributes) GetPlmnIdList() []PlmnId {
	if o == nil || IsNil(o.PlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.PlmnIdList
}

// GetPlmnIdListOk returns a tuple with the PlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributes) GetPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.PlmnIdList) {
		return nil, false
	}
	return o.PlmnIdList, true
}

// HasPlmnIdList returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributes) HasPlmnIdList() bool {
	if o != nil && !IsNil(o.PlmnIdList) {
		return true
	}

	return false
}

// SetPlmnIdList gets a reference to the given []PlmnId and assigns it to the PlmnIdList field.
func (o *AmfRegionSingleAllOfAttributes) SetPlmnIdList(v []PlmnId) {
	o.PlmnIdList = v
}

// GetNRTACList returns the NRTACList field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributes) GetNRTACList() []int32 {
	if o == nil || IsNil(o.NRTACList) {
		var ret []int32
		return ret
	}
	return o.NRTACList
}

// GetNRTACListOk returns a tuple with the NRTACList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributes) GetNRTACListOk() ([]int32, bool) {
	if o == nil || IsNil(o.NRTACList) {
		return nil, false
	}
	return o.NRTACList, true
}

// HasNRTACList returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributes) HasNRTACList() bool {
	if o != nil && !IsNil(o.NRTACList) {
		return true
	}

	return false
}

// SetNRTACList gets a reference to the given []int32 and assigns it to the NRTACList field.
func (o *AmfRegionSingleAllOfAttributes) SetNRTACList(v []int32) {
	o.NRTACList = v
}

// GetAmfRegionId returns the AmfRegionId field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributes) GetAmfRegionId() int32 {
	if o == nil || IsNil(o.AmfRegionId) {
		var ret int32
		return ret
	}
	return *o.AmfRegionId
}

// GetAmfRegionIdOk returns a tuple with the AmfRegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributes) GetAmfRegionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AmfRegionId) {
		return nil, false
	}
	return o.AmfRegionId, true
}

// HasAmfRegionId returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributes) HasAmfRegionId() bool {
	if o != nil && !IsNil(o.AmfRegionId) {
		return true
	}

	return false
}

// SetAmfRegionId gets a reference to the given int32 and assigns it to the AmfRegionId field.
func (o *AmfRegionSingleAllOfAttributes) SetAmfRegionId(v int32) {
	o.AmfRegionId = &v
}

// GetSnssaiList returns the SnssaiList field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributes) GetSnssaiList() []Snssai {
	if o == nil || IsNil(o.SnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributes) GetSnssaiListOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.SnssaiList) {
		return nil, false
	}
	return o.SnssaiList, true
}

// HasSnssaiList returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributes) HasSnssaiList() bool {
	if o != nil && !IsNil(o.SnssaiList) {
		return true
	}

	return false
}

// SetSnssaiList gets a reference to the given []Snssai and assigns it to the SnssaiList field.
func (o *AmfRegionSingleAllOfAttributes) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

// GetAMFSetListRef returns the AMFSetListRef field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributes) GetAMFSetListRef() []string {
	if o == nil || IsNil(o.AMFSetListRef) {
		var ret []string
		return ret
	}
	return o.AMFSetListRef
}

// GetAMFSetListRefOk returns a tuple with the AMFSetListRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributes) GetAMFSetListRefOk() ([]string, bool) {
	if o == nil || IsNil(o.AMFSetListRef) {
		return nil, false
	}
	return o.AMFSetListRef, true
}

// HasAMFSetListRef returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributes) HasAMFSetListRef() bool {
	if o != nil && !IsNil(o.AMFSetListRef) {
		return true
	}

	return false
}

// SetAMFSetListRef gets a reference to the given []string and assigns it to the AMFSetListRef field.
func (o *AmfRegionSingleAllOfAttributes) SetAMFSetListRef(v []string) {
	o.AMFSetListRef = v
}

func (o AmfRegionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfRegionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PlmnIdList) {
		toSerialize["plmnIdList"] = o.PlmnIdList
	}
	if !IsNil(o.NRTACList) {
		toSerialize["nRTACList"] = o.NRTACList
	}
	if !IsNil(o.AmfRegionId) {
		toSerialize["amfRegionId"] = o.AmfRegionId
	}
	if !IsNil(o.SnssaiList) {
		toSerialize["snssaiList"] = o.SnssaiList
	}
	if !IsNil(o.AMFSetListRef) {
		toSerialize["aMFSetListRef"] = o.AMFSetListRef
	}
	return toSerialize, nil
}

type NullableAmfRegionSingleAllOfAttributes struct {
	value *AmfRegionSingleAllOfAttributes
	isSet bool
}

func (v NullableAmfRegionSingleAllOfAttributes) Get() *AmfRegionSingleAllOfAttributes {
	return v.value
}

func (v *NullableAmfRegionSingleAllOfAttributes) Set(val *AmfRegionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfRegionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfRegionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfRegionSingleAllOfAttributes(val *AmfRegionSingleAllOfAttributes) *NullableAmfRegionSingleAllOfAttributes {
	return &NullableAmfRegionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableAmfRegionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfRegionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



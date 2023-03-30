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

// checks if the ScpFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScpFunctionSingleAllOfAttributes{}

// ScpFunctionSingleAllOfAttributes struct for ScpFunctionSingleAllOfAttributes
type ScpFunctionSingleAllOfAttributes struct {
	UserLabel *string `json:"userLabel,omitempty"`
	VnfParametersList []VnfParameter `json:"vnfParametersList,omitempty"`
	PeeParametersList []PeeParameter `json:"peeParametersList,omitempty"`
	PriorityLabel *int32 `json:"priorityLabel,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
	SupportedFuncList []SupportedFunc `json:"supportedFuncList,omitempty"`
	Address *HostAddr `json:"address,omitempty"`
	ScpInfo *ScpInfo `json:"scpInfo,omitempty"`
}

// NewScpFunctionSingleAllOfAttributes instantiates a new ScpFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScpFunctionSingleAllOfAttributes() *ScpFunctionSingleAllOfAttributes {
	this := ScpFunctionSingleAllOfAttributes{}
	return &this
}

// NewScpFunctionSingleAllOfAttributesWithDefaults instantiates a new ScpFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScpFunctionSingleAllOfAttributesWithDefaults() *ScpFunctionSingleAllOfAttributes {
	this := ScpFunctionSingleAllOfAttributes{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ScpFunctionSingleAllOfAttributes) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpFunctionSingleAllOfAttributes) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ScpFunctionSingleAllOfAttributes) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ScpFunctionSingleAllOfAttributes) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetVnfParametersList returns the VnfParametersList field value if set, zero value otherwise.
func (o *ScpFunctionSingleAllOfAttributes) GetVnfParametersList() []VnfParameter {
	if o == nil || IsNil(o.VnfParametersList) {
		var ret []VnfParameter
		return ret
	}
	return o.VnfParametersList
}

// GetVnfParametersListOk returns a tuple with the VnfParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpFunctionSingleAllOfAttributes) GetVnfParametersListOk() ([]VnfParameter, bool) {
	if o == nil || IsNil(o.VnfParametersList) {
		return nil, false
	}
	return o.VnfParametersList, true
}

// HasVnfParametersList returns a boolean if a field has been set.
func (o *ScpFunctionSingleAllOfAttributes) HasVnfParametersList() bool {
	if o != nil && !IsNil(o.VnfParametersList) {
		return true
	}

	return false
}

// SetVnfParametersList gets a reference to the given []VnfParameter and assigns it to the VnfParametersList field.
func (o *ScpFunctionSingleAllOfAttributes) SetVnfParametersList(v []VnfParameter) {
	o.VnfParametersList = v
}

// GetPeeParametersList returns the PeeParametersList field value if set, zero value otherwise.
func (o *ScpFunctionSingleAllOfAttributes) GetPeeParametersList() []PeeParameter {
	if o == nil || IsNil(o.PeeParametersList) {
		var ret []PeeParameter
		return ret
	}
	return o.PeeParametersList
}

// GetPeeParametersListOk returns a tuple with the PeeParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpFunctionSingleAllOfAttributes) GetPeeParametersListOk() ([]PeeParameter, bool) {
	if o == nil || IsNil(o.PeeParametersList) {
		return nil, false
	}
	return o.PeeParametersList, true
}

// HasPeeParametersList returns a boolean if a field has been set.
func (o *ScpFunctionSingleAllOfAttributes) HasPeeParametersList() bool {
	if o != nil && !IsNil(o.PeeParametersList) {
		return true
	}

	return false
}

// SetPeeParametersList gets a reference to the given []PeeParameter and assigns it to the PeeParametersList field.
func (o *ScpFunctionSingleAllOfAttributes) SetPeeParametersList(v []PeeParameter) {
	o.PeeParametersList = v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *ScpFunctionSingleAllOfAttributes) GetPriorityLabel() int32 {
	if o == nil || IsNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpFunctionSingleAllOfAttributes) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLabel) {
		return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *ScpFunctionSingleAllOfAttributes) HasPriorityLabel() bool {
	if o != nil && !IsNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *ScpFunctionSingleAllOfAttributes) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *ScpFunctionSingleAllOfAttributes) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpFunctionSingleAllOfAttributes) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *ScpFunctionSingleAllOfAttributes) HasSupportedPerfMetricGroups() bool {
	if o != nil && !IsNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *ScpFunctionSingleAllOfAttributes) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetSupportedTraceMetrics returns the SupportedTraceMetrics field value if set, zero value otherwise.
func (o *ScpFunctionSingleAllOfAttributes) GetSupportedTraceMetrics() []string {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedTraceMetrics
}

// GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpFunctionSingleAllOfAttributes) GetSupportedTraceMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		return nil, false
	}
	return o.SupportedTraceMetrics, true
}

// HasSupportedTraceMetrics returns a boolean if a field has been set.
func (o *ScpFunctionSingleAllOfAttributes) HasSupportedTraceMetrics() bool {
	if o != nil && !IsNil(o.SupportedTraceMetrics) {
		return true
	}

	return false
}

// SetSupportedTraceMetrics gets a reference to the given []string and assigns it to the SupportedTraceMetrics field.
func (o *ScpFunctionSingleAllOfAttributes) SetSupportedTraceMetrics(v []string) {
	o.SupportedTraceMetrics = v
}

// GetSupportedFuncList returns the SupportedFuncList field value if set, zero value otherwise.
func (o *ScpFunctionSingleAllOfAttributes) GetSupportedFuncList() []SupportedFunc {
	if o == nil || IsNil(o.SupportedFuncList) {
		var ret []SupportedFunc
		return ret
	}
	return o.SupportedFuncList
}

// GetSupportedFuncListOk returns a tuple with the SupportedFuncList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpFunctionSingleAllOfAttributes) GetSupportedFuncListOk() ([]SupportedFunc, bool) {
	if o == nil || IsNil(o.SupportedFuncList) {
		return nil, false
	}
	return o.SupportedFuncList, true
}

// HasSupportedFuncList returns a boolean if a field has been set.
func (o *ScpFunctionSingleAllOfAttributes) HasSupportedFuncList() bool {
	if o != nil && !IsNil(o.SupportedFuncList) {
		return true
	}

	return false
}

// SetSupportedFuncList gets a reference to the given []SupportedFunc and assigns it to the SupportedFuncList field.
func (o *ScpFunctionSingleAllOfAttributes) SetSupportedFuncList(v []SupportedFunc) {
	o.SupportedFuncList = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ScpFunctionSingleAllOfAttributes) GetAddress() HostAddr {
	if o == nil || IsNil(o.Address) {
		var ret HostAddr
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpFunctionSingleAllOfAttributes) GetAddressOk() (*HostAddr, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ScpFunctionSingleAllOfAttributes) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given HostAddr and assigns it to the Address field.
func (o *ScpFunctionSingleAllOfAttributes) SetAddress(v HostAddr) {
	o.Address = &v
}

// GetScpInfo returns the ScpInfo field value if set, zero value otherwise.
func (o *ScpFunctionSingleAllOfAttributes) GetScpInfo() ScpInfo {
	if o == nil || IsNil(o.ScpInfo) {
		var ret ScpInfo
		return ret
	}
	return *o.ScpInfo
}

// GetScpInfoOk returns a tuple with the ScpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpFunctionSingleAllOfAttributes) GetScpInfoOk() (*ScpInfo, bool) {
	if o == nil || IsNil(o.ScpInfo) {
		return nil, false
	}
	return o.ScpInfo, true
}

// HasScpInfo returns a boolean if a field has been set.
func (o *ScpFunctionSingleAllOfAttributes) HasScpInfo() bool {
	if o != nil && !IsNil(o.ScpInfo) {
		return true
	}

	return false
}

// SetScpInfo gets a reference to the given ScpInfo and assigns it to the ScpInfo field.
func (o *ScpFunctionSingleAllOfAttributes) SetScpInfo(v ScpInfo) {
	o.ScpInfo = &v
}

func (o ScpFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScpFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SupportedFuncList) {
		toSerialize["supportedFuncList"] = o.SupportedFuncList
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.ScpInfo) {
		toSerialize["scpInfo"] = o.ScpInfo
	}
	return toSerialize, nil
}

type NullableScpFunctionSingleAllOfAttributes struct {
	value *ScpFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableScpFunctionSingleAllOfAttributes) Get() *ScpFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableScpFunctionSingleAllOfAttributes) Set(val *ScpFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableScpFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableScpFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScpFunctionSingleAllOfAttributes(val *ScpFunctionSingleAllOfAttributes) *NullableScpFunctionSingleAllOfAttributes {
	return &NullableScpFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableScpFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScpFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// checks if the ExternalNrfFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalNrfFunctionSingleAllOfAttributes{}

// ExternalNrfFunctionSingleAllOfAttributes struct for ExternalNrfFunctionSingleAllOfAttributes
type ExternalNrfFunctionSingleAllOfAttributes struct {
	UserLabel *string `json:"userLabel,omitempty"`
	VnfParametersList []VnfParameter `json:"vnfParametersList,omitempty"`
	PeeParametersList []PeeParameter `json:"peeParametersList,omitempty"`
	PriorityLabel *int32 `json:"priorityLabel,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
	PlmnIdList []PlmnId `json:"plmnIdList,omitempty"`
}

// NewExternalNrfFunctionSingleAllOfAttributes instantiates a new ExternalNrfFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalNrfFunctionSingleAllOfAttributes() *ExternalNrfFunctionSingleAllOfAttributes {
	this := ExternalNrfFunctionSingleAllOfAttributes{}
	return &this
}

// NewExternalNrfFunctionSingleAllOfAttributesWithDefaults instantiates a new ExternalNrfFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalNrfFunctionSingleAllOfAttributesWithDefaults() *ExternalNrfFunctionSingleAllOfAttributes {
	this := ExternalNrfFunctionSingleAllOfAttributes{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ExternalNrfFunctionSingleAllOfAttributes) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetVnfParametersList returns the VnfParametersList field value if set, zero value otherwise.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetVnfParametersList() []VnfParameter {
	if o == nil || IsNil(o.VnfParametersList) {
		var ret []VnfParameter
		return ret
	}
	return o.VnfParametersList
}

// GetVnfParametersListOk returns a tuple with the VnfParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetVnfParametersListOk() ([]VnfParameter, bool) {
	if o == nil || IsNil(o.VnfParametersList) {
		return nil, false
	}
	return o.VnfParametersList, true
}

// HasVnfParametersList returns a boolean if a field has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) HasVnfParametersList() bool {
	if o != nil && !IsNil(o.VnfParametersList) {
		return true
	}

	return false
}

// SetVnfParametersList gets a reference to the given []VnfParameter and assigns it to the VnfParametersList field.
func (o *ExternalNrfFunctionSingleAllOfAttributes) SetVnfParametersList(v []VnfParameter) {
	o.VnfParametersList = v
}

// GetPeeParametersList returns the PeeParametersList field value if set, zero value otherwise.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetPeeParametersList() []PeeParameter {
	if o == nil || IsNil(o.PeeParametersList) {
		var ret []PeeParameter
		return ret
	}
	return o.PeeParametersList
}

// GetPeeParametersListOk returns a tuple with the PeeParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetPeeParametersListOk() ([]PeeParameter, bool) {
	if o == nil || IsNil(o.PeeParametersList) {
		return nil, false
	}
	return o.PeeParametersList, true
}

// HasPeeParametersList returns a boolean if a field has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) HasPeeParametersList() bool {
	if o != nil && !IsNil(o.PeeParametersList) {
		return true
	}

	return false
}

// SetPeeParametersList gets a reference to the given []PeeParameter and assigns it to the PeeParametersList field.
func (o *ExternalNrfFunctionSingleAllOfAttributes) SetPeeParametersList(v []PeeParameter) {
	o.PeeParametersList = v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetPriorityLabel() int32 {
	if o == nil || IsNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLabel) {
		return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) HasPriorityLabel() bool {
	if o != nil && !IsNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *ExternalNrfFunctionSingleAllOfAttributes) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) HasSupportedPerfMetricGroups() bool {
	if o != nil && !IsNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *ExternalNrfFunctionSingleAllOfAttributes) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetSupportedTraceMetrics returns the SupportedTraceMetrics field value if set, zero value otherwise.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetSupportedTraceMetrics() []string {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedTraceMetrics
}

// GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetSupportedTraceMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		return nil, false
	}
	return o.SupportedTraceMetrics, true
}

// HasSupportedTraceMetrics returns a boolean if a field has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) HasSupportedTraceMetrics() bool {
	if o != nil && !IsNil(o.SupportedTraceMetrics) {
		return true
	}

	return false
}

// SetSupportedTraceMetrics gets a reference to the given []string and assigns it to the SupportedTraceMetrics field.
func (o *ExternalNrfFunctionSingleAllOfAttributes) SetSupportedTraceMetrics(v []string) {
	o.SupportedTraceMetrics = v
}

// GetPlmnIdList returns the PlmnIdList field value if set, zero value otherwise.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetPlmnIdList() []PlmnId {
	if o == nil || IsNil(o.PlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.PlmnIdList
}

// GetPlmnIdListOk returns a tuple with the PlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.PlmnIdList) {
		return nil, false
	}
	return o.PlmnIdList, true
}

// HasPlmnIdList returns a boolean if a field has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) HasPlmnIdList() bool {
	if o != nil && !IsNil(o.PlmnIdList) {
		return true
	}

	return false
}

// SetPlmnIdList gets a reference to the given []PlmnId and assigns it to the PlmnIdList field.
func (o *ExternalNrfFunctionSingleAllOfAttributes) SetPlmnIdList(v []PlmnId) {
	o.PlmnIdList = v
}

func (o ExternalNrfFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalNrfFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableExternalNrfFunctionSingleAllOfAttributes struct {
	value *ExternalNrfFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableExternalNrfFunctionSingleAllOfAttributes) Get() *ExternalNrfFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableExternalNrfFunctionSingleAllOfAttributes) Set(val *ExternalNrfFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalNrfFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalNrfFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalNrfFunctionSingleAllOfAttributes(val *ExternalNrfFunctionSingleAllOfAttributes) *NullableExternalNrfFunctionSingleAllOfAttributes {
	return &NullableExternalNrfFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableExternalNrfFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalNrfFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



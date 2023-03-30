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

// checks if the ExternalGnbCuCpFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalGnbCuCpFunctionSingleAllOfAttributes{}

// ExternalGnbCuCpFunctionSingleAllOfAttributes struct for ExternalGnbCuCpFunctionSingleAllOfAttributes
type ExternalGnbCuCpFunctionSingleAllOfAttributes struct {
	UserLabel *string `json:"userLabel,omitempty"`
	VnfParametersList []VnfParameter `json:"vnfParametersList,omitempty"`
	PeeParametersList []PeeParameter `json:"peeParametersList,omitempty"`
	PriorityLabel *int32 `json:"priorityLabel,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
	GnbId *string `json:"gnbId,omitempty"`
	GnbIdLength *int32 `json:"gnbIdLength,omitempty"`
	PlmnId *PlmnId `json:"plmnId,omitempty"`
}

// NewExternalGnbCuCpFunctionSingleAllOfAttributes instantiates a new ExternalGnbCuCpFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGnbCuCpFunctionSingleAllOfAttributes() *ExternalGnbCuCpFunctionSingleAllOfAttributes {
	this := ExternalGnbCuCpFunctionSingleAllOfAttributes{}
	return &this
}

// NewExternalGnbCuCpFunctionSingleAllOfAttributesWithDefaults instantiates a new ExternalGnbCuCpFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGnbCuCpFunctionSingleAllOfAttributesWithDefaults() *ExternalGnbCuCpFunctionSingleAllOfAttributes {
	this := ExternalGnbCuCpFunctionSingleAllOfAttributes{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetVnfParametersList returns the VnfParametersList field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetVnfParametersList() []VnfParameter {
	if o == nil || IsNil(o.VnfParametersList) {
		var ret []VnfParameter
		return ret
	}
	return o.VnfParametersList
}

// GetVnfParametersListOk returns a tuple with the VnfParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetVnfParametersListOk() ([]VnfParameter, bool) {
	if o == nil || IsNil(o.VnfParametersList) {
		return nil, false
	}
	return o.VnfParametersList, true
}

// HasVnfParametersList returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) HasVnfParametersList() bool {
	if o != nil && !IsNil(o.VnfParametersList) {
		return true
	}

	return false
}

// SetVnfParametersList gets a reference to the given []VnfParameter and assigns it to the VnfParametersList field.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) SetVnfParametersList(v []VnfParameter) {
	o.VnfParametersList = v
}

// GetPeeParametersList returns the PeeParametersList field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetPeeParametersList() []PeeParameter {
	if o == nil || IsNil(o.PeeParametersList) {
		var ret []PeeParameter
		return ret
	}
	return o.PeeParametersList
}

// GetPeeParametersListOk returns a tuple with the PeeParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetPeeParametersListOk() ([]PeeParameter, bool) {
	if o == nil || IsNil(o.PeeParametersList) {
		return nil, false
	}
	return o.PeeParametersList, true
}

// HasPeeParametersList returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) HasPeeParametersList() bool {
	if o != nil && !IsNil(o.PeeParametersList) {
		return true
	}

	return false
}

// SetPeeParametersList gets a reference to the given []PeeParameter and assigns it to the PeeParametersList field.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) SetPeeParametersList(v []PeeParameter) {
	o.PeeParametersList = v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetPriorityLabel() int32 {
	if o == nil || IsNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLabel) {
		return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) HasPriorityLabel() bool {
	if o != nil && !IsNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) HasSupportedPerfMetricGroups() bool {
	if o != nil && !IsNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetSupportedTraceMetrics returns the SupportedTraceMetrics field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetSupportedTraceMetrics() []string {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedTraceMetrics
}

// GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetSupportedTraceMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		return nil, false
	}
	return o.SupportedTraceMetrics, true
}

// HasSupportedTraceMetrics returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) HasSupportedTraceMetrics() bool {
	if o != nil && !IsNil(o.SupportedTraceMetrics) {
		return true
	}

	return false
}

// SetSupportedTraceMetrics gets a reference to the given []string and assigns it to the SupportedTraceMetrics field.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) SetSupportedTraceMetrics(v []string) {
	o.SupportedTraceMetrics = v
}

// GetGnbId returns the GnbId field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetGnbId() string {
	if o == nil || IsNil(o.GnbId) {
		var ret string
		return ret
	}
	return *o.GnbId
}

// GetGnbIdOk returns a tuple with the GnbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetGnbIdOk() (*string, bool) {
	if o == nil || IsNil(o.GnbId) {
		return nil, false
	}
	return o.GnbId, true
}

// HasGnbId returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) HasGnbId() bool {
	if o != nil && !IsNil(o.GnbId) {
		return true
	}

	return false
}

// SetGnbId gets a reference to the given string and assigns it to the GnbId field.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) SetGnbId(v string) {
	o.GnbId = &v
}

// GetGnbIdLength returns the GnbIdLength field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetGnbIdLength() int32 {
	if o == nil || IsNil(o.GnbIdLength) {
		var ret int32
		return ret
	}
	return *o.GnbIdLength
}

// GetGnbIdLengthOk returns a tuple with the GnbIdLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetGnbIdLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.GnbIdLength) {
		return nil, false
	}
	return o.GnbIdLength, true
}

// HasGnbIdLength returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) HasGnbIdLength() bool {
	if o != nil && !IsNil(o.GnbIdLength) {
		return true
	}

	return false
}

// SetGnbIdLength gets a reference to the given int32 and assigns it to the GnbIdLength field.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) SetGnbIdLength(v int32) {
	o.GnbIdLength = &v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetPlmnId() PlmnId {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *ExternalGnbCuCpFunctionSingleAllOfAttributes) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

func (o ExternalGnbCuCpFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalGnbCuCpFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.GnbId) {
		toSerialize["gnbId"] = o.GnbId
	}
	if !IsNil(o.GnbIdLength) {
		toSerialize["gnbIdLength"] = o.GnbIdLength
	}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	return toSerialize, nil
}

type NullableExternalGnbCuCpFunctionSingleAllOfAttributes struct {
	value *ExternalGnbCuCpFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableExternalGnbCuCpFunctionSingleAllOfAttributes) Get() *ExternalGnbCuCpFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableExternalGnbCuCpFunctionSingleAllOfAttributes) Set(val *ExternalGnbCuCpFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGnbCuCpFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGnbCuCpFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGnbCuCpFunctionSingleAllOfAttributes(val *ExternalGnbCuCpFunctionSingleAllOfAttributes) *NullableExternalGnbCuCpFunctionSingleAllOfAttributes {
	return &NullableExternalGnbCuCpFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableExternalGnbCuCpFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGnbCuCpFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


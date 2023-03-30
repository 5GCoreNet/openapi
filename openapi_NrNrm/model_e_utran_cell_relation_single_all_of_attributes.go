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

// checks if the EUtranCellRelationSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EUtranCellRelationSingleAllOfAttributes{}

// EUtranCellRelationSingleAllOfAttributes struct for EUtranCellRelationSingleAllOfAttributes
type EUtranCellRelationSingleAllOfAttributes struct {
	UserLabel *string `json:"userLabel,omitempty"`
	VnfParametersList []VnfParameter `json:"vnfParametersList,omitempty"`
	PeeParametersList []PeeParameter `json:"peeParametersList,omitempty"`
	PriorityLabel *int32 `json:"priorityLabel,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
	AdjacentEUtranCellRef *string `json:"adjacentEUtranCellRef,omitempty"`
}

// NewEUtranCellRelationSingleAllOfAttributes instantiates a new EUtranCellRelationSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEUtranCellRelationSingleAllOfAttributes() *EUtranCellRelationSingleAllOfAttributes {
	this := EUtranCellRelationSingleAllOfAttributes{}
	return &this
}

// NewEUtranCellRelationSingleAllOfAttributesWithDefaults instantiates a new EUtranCellRelationSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEUtranCellRelationSingleAllOfAttributesWithDefaults() *EUtranCellRelationSingleAllOfAttributes {
	this := EUtranCellRelationSingleAllOfAttributes{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *EUtranCellRelationSingleAllOfAttributes) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *EUtranCellRelationSingleAllOfAttributes) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetVnfParametersList returns the VnfParametersList field value if set, zero value otherwise.
func (o *EUtranCellRelationSingleAllOfAttributes) GetVnfParametersList() []VnfParameter {
	if o == nil || IsNil(o.VnfParametersList) {
		var ret []VnfParameter
		return ret
	}
	return o.VnfParametersList
}

// GetVnfParametersListOk returns a tuple with the VnfParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) GetVnfParametersListOk() ([]VnfParameter, bool) {
	if o == nil || IsNil(o.VnfParametersList) {
		return nil, false
	}
	return o.VnfParametersList, true
}

// HasVnfParametersList returns a boolean if a field has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) HasVnfParametersList() bool {
	if o != nil && !IsNil(o.VnfParametersList) {
		return true
	}

	return false
}

// SetVnfParametersList gets a reference to the given []VnfParameter and assigns it to the VnfParametersList field.
func (o *EUtranCellRelationSingleAllOfAttributes) SetVnfParametersList(v []VnfParameter) {
	o.VnfParametersList = v
}

// GetPeeParametersList returns the PeeParametersList field value if set, zero value otherwise.
func (o *EUtranCellRelationSingleAllOfAttributes) GetPeeParametersList() []PeeParameter {
	if o == nil || IsNil(o.PeeParametersList) {
		var ret []PeeParameter
		return ret
	}
	return o.PeeParametersList
}

// GetPeeParametersListOk returns a tuple with the PeeParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) GetPeeParametersListOk() ([]PeeParameter, bool) {
	if o == nil || IsNil(o.PeeParametersList) {
		return nil, false
	}
	return o.PeeParametersList, true
}

// HasPeeParametersList returns a boolean if a field has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) HasPeeParametersList() bool {
	if o != nil && !IsNil(o.PeeParametersList) {
		return true
	}

	return false
}

// SetPeeParametersList gets a reference to the given []PeeParameter and assigns it to the PeeParametersList field.
func (o *EUtranCellRelationSingleAllOfAttributes) SetPeeParametersList(v []PeeParameter) {
	o.PeeParametersList = v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *EUtranCellRelationSingleAllOfAttributes) GetPriorityLabel() int32 {
	if o == nil || IsNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLabel) {
		return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) HasPriorityLabel() bool {
	if o != nil && !IsNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *EUtranCellRelationSingleAllOfAttributes) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *EUtranCellRelationSingleAllOfAttributes) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) HasSupportedPerfMetricGroups() bool {
	if o != nil && !IsNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *EUtranCellRelationSingleAllOfAttributes) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetSupportedTraceMetrics returns the SupportedTraceMetrics field value if set, zero value otherwise.
func (o *EUtranCellRelationSingleAllOfAttributes) GetSupportedTraceMetrics() []string {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedTraceMetrics
}

// GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) GetSupportedTraceMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		return nil, false
	}
	return o.SupportedTraceMetrics, true
}

// HasSupportedTraceMetrics returns a boolean if a field has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) HasSupportedTraceMetrics() bool {
	if o != nil && !IsNil(o.SupportedTraceMetrics) {
		return true
	}

	return false
}

// SetSupportedTraceMetrics gets a reference to the given []string and assigns it to the SupportedTraceMetrics field.
func (o *EUtranCellRelationSingleAllOfAttributes) SetSupportedTraceMetrics(v []string) {
	o.SupportedTraceMetrics = v
}

// GetAdjacentEUtranCellRef returns the AdjacentEUtranCellRef field value if set, zero value otherwise.
func (o *EUtranCellRelationSingleAllOfAttributes) GetAdjacentEUtranCellRef() string {
	if o == nil || IsNil(o.AdjacentEUtranCellRef) {
		var ret string
		return ret
	}
	return *o.AdjacentEUtranCellRef
}

// GetAdjacentEUtranCellRefOk returns a tuple with the AdjacentEUtranCellRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) GetAdjacentEUtranCellRefOk() (*string, bool) {
	if o == nil || IsNil(o.AdjacentEUtranCellRef) {
		return nil, false
	}
	return o.AdjacentEUtranCellRef, true
}

// HasAdjacentEUtranCellRef returns a boolean if a field has been set.
func (o *EUtranCellRelationSingleAllOfAttributes) HasAdjacentEUtranCellRef() bool {
	if o != nil && !IsNil(o.AdjacentEUtranCellRef) {
		return true
	}

	return false
}

// SetAdjacentEUtranCellRef gets a reference to the given string and assigns it to the AdjacentEUtranCellRef field.
func (o *EUtranCellRelationSingleAllOfAttributes) SetAdjacentEUtranCellRef(v string) {
	o.AdjacentEUtranCellRef = &v
}

func (o EUtranCellRelationSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EUtranCellRelationSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AdjacentEUtranCellRef) {
		toSerialize["adjacentEUtranCellRef"] = o.AdjacentEUtranCellRef
	}
	return toSerialize, nil
}

type NullableEUtranCellRelationSingleAllOfAttributes struct {
	value *EUtranCellRelationSingleAllOfAttributes
	isSet bool
}

func (v NullableEUtranCellRelationSingleAllOfAttributes) Get() *EUtranCellRelationSingleAllOfAttributes {
	return v.value
}

func (v *NullableEUtranCellRelationSingleAllOfAttributes) Set(val *EUtranCellRelationSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEUtranCellRelationSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEUtranCellRelationSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEUtranCellRelationSingleAllOfAttributes(val *EUtranCellRelationSingleAllOfAttributes) *NullableEUtranCellRelationSingleAllOfAttributes {
	return &NullableEUtranCellRelationSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableEUtranCellRelationSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEUtranCellRelationSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



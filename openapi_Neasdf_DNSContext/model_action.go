/*
Neasdf_DNSContext

EASDF DNS Context Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Neasdf_DNSContext

import (
	"encoding/json"
)

// checks if the Action type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Action{}

// Action Action to apply to DNS messages matching a message detection template
type Action struct {
	ApplyAction           ApplyAction           `json:"applyAction"`
	FwdParas              *ForwardingParameters `json:"fwdParas,omitempty"`
	ReportingOnceInd      *bool                 `json:"reportingOnceInd,omitempty"`
	ResetReportingOnceInd *bool                 `json:"resetReportingOnceInd,omitempty"`
}

// NewAction instantiates a new Action object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAction(applyAction ApplyAction) *Action {
	this := Action{}
	this.ApplyAction = applyAction
	var reportingOnceInd bool = false
	this.ReportingOnceInd = &reportingOnceInd
	var resetReportingOnceInd bool = false
	this.ResetReportingOnceInd = &resetReportingOnceInd
	return &this
}

// NewActionWithDefaults instantiates a new Action object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionWithDefaults() *Action {
	this := Action{}
	var reportingOnceInd bool = false
	this.ReportingOnceInd = &reportingOnceInd
	var resetReportingOnceInd bool = false
	this.ResetReportingOnceInd = &resetReportingOnceInd
	return &this
}

// GetApplyAction returns the ApplyAction field value
func (o *Action) GetApplyAction() ApplyAction {
	if o == nil {
		var ret ApplyAction
		return ret
	}

	return o.ApplyAction
}

// GetApplyActionOk returns a tuple with the ApplyAction field value
// and a boolean to check if the value has been set.
func (o *Action) GetApplyActionOk() (*ApplyAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplyAction, true
}

// SetApplyAction sets field value
func (o *Action) SetApplyAction(v ApplyAction) {
	o.ApplyAction = v
}

// GetFwdParas returns the FwdParas field value if set, zero value otherwise.
func (o *Action) GetFwdParas() ForwardingParameters {
	if o == nil || IsNil(o.FwdParas) {
		var ret ForwardingParameters
		return ret
	}
	return *o.FwdParas
}

// GetFwdParasOk returns a tuple with the FwdParas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetFwdParasOk() (*ForwardingParameters, bool) {
	if o == nil || IsNil(o.FwdParas) {
		return nil, false
	}
	return o.FwdParas, true
}

// HasFwdParas returns a boolean if a field has been set.
func (o *Action) HasFwdParas() bool {
	if o != nil && !IsNil(o.FwdParas) {
		return true
	}

	return false
}

// SetFwdParas gets a reference to the given ForwardingParameters and assigns it to the FwdParas field.
func (o *Action) SetFwdParas(v ForwardingParameters) {
	o.FwdParas = &v
}

// GetReportingOnceInd returns the ReportingOnceInd field value if set, zero value otherwise.
func (o *Action) GetReportingOnceInd() bool {
	if o == nil || IsNil(o.ReportingOnceInd) {
		var ret bool
		return ret
	}
	return *o.ReportingOnceInd
}

// GetReportingOnceIndOk returns a tuple with the ReportingOnceInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetReportingOnceIndOk() (*bool, bool) {
	if o == nil || IsNil(o.ReportingOnceInd) {
		return nil, false
	}
	return o.ReportingOnceInd, true
}

// HasReportingOnceInd returns a boolean if a field has been set.
func (o *Action) HasReportingOnceInd() bool {
	if o != nil && !IsNil(o.ReportingOnceInd) {
		return true
	}

	return false
}

// SetReportingOnceInd gets a reference to the given bool and assigns it to the ReportingOnceInd field.
func (o *Action) SetReportingOnceInd(v bool) {
	o.ReportingOnceInd = &v
}

// GetResetReportingOnceInd returns the ResetReportingOnceInd field value if set, zero value otherwise.
func (o *Action) GetResetReportingOnceInd() bool {
	if o == nil || IsNil(o.ResetReportingOnceInd) {
		var ret bool
		return ret
	}
	return *o.ResetReportingOnceInd
}

// GetResetReportingOnceIndOk returns a tuple with the ResetReportingOnceInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetResetReportingOnceIndOk() (*bool, bool) {
	if o == nil || IsNil(o.ResetReportingOnceInd) {
		return nil, false
	}
	return o.ResetReportingOnceInd, true
}

// HasResetReportingOnceInd returns a boolean if a field has been set.
func (o *Action) HasResetReportingOnceInd() bool {
	if o != nil && !IsNil(o.ResetReportingOnceInd) {
		return true
	}

	return false
}

// SetResetReportingOnceInd gets a reference to the given bool and assigns it to the ResetReportingOnceInd field.
func (o *Action) SetResetReportingOnceInd(v bool) {
	o.ResetReportingOnceInd = &v
}

func (o Action) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Action) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applyAction"] = o.ApplyAction
	if !IsNil(o.FwdParas) {
		toSerialize["fwdParas"] = o.FwdParas
	}
	if !IsNil(o.ReportingOnceInd) {
		toSerialize["reportingOnceInd"] = o.ReportingOnceInd
	}
	if !IsNil(o.ResetReportingOnceInd) {
		toSerialize["resetReportingOnceInd"] = o.ResetReportingOnceInd
	}
	return toSerialize, nil
}

type NullableAction struct {
	value *Action
	isSet bool
}

func (v NullableAction) Get() *Action {
	return v.value
}

func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

func (v NullableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

func (v NullableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

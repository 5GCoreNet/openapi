/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
)

// checks if the RuleReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleReport{}

// RuleReport Reports the status of PCC.
type RuleReport struct {
	// Contains the identifier of the affected PCC rule(s).
	PccRuleIds []string   `json:"pccRuleIds"`
	RuleStatus RuleStatus `json:"ruleStatus"`
	// Indicates the version of a PCC rule.
	ContVers    []int32      `json:"contVers,omitempty"`
	FailureCode *FailureCode `json:"failureCode,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RetryAfter *int32           `json:"retryAfter,omitempty"`
	FinUnitAct *FinalUnitAction `json:"finUnitAct,omitempty"`
	// indicates the RAN or NAS release cause code information.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`
	// Indicates the alternative QoS parameter set that the NG-RAN can guarantee. It is included during the report of successfull resource allocation and indicates that NG-RAN used an alternative QoS profile because the requested QoS could not be allocated..
	AltQosParamId *string `json:"altQosParamId,omitempty"`
	// When present and set to true it indicates that the Alternative QoS profiles are not  supported by NG-RAN.
	AltQosNotSuppInd *bool `json:"altQosNotSuppInd,omitempty"`
}

// NewRuleReport instantiates a new RuleReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleReport(pccRuleIds []string, ruleStatus RuleStatus) *RuleReport {
	this := RuleReport{}
	this.PccRuleIds = pccRuleIds
	this.RuleStatus = ruleStatus
	return &this
}

// NewRuleReportWithDefaults instantiates a new RuleReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleReportWithDefaults() *RuleReport {
	this := RuleReport{}
	return &this
}

// GetPccRuleIds returns the PccRuleIds field value
func (o *RuleReport) GetPccRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PccRuleIds
}

// GetPccRuleIdsOk returns a tuple with the PccRuleIds field value
// and a boolean to check if the value has been set.
func (o *RuleReport) GetPccRuleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PccRuleIds, true
}

// SetPccRuleIds sets field value
func (o *RuleReport) SetPccRuleIds(v []string) {
	o.PccRuleIds = v
}

// GetRuleStatus returns the RuleStatus field value
func (o *RuleReport) GetRuleStatus() RuleStatus {
	if o == nil {
		var ret RuleStatus
		return ret
	}

	return o.RuleStatus
}

// GetRuleStatusOk returns a tuple with the RuleStatus field value
// and a boolean to check if the value has been set.
func (o *RuleReport) GetRuleStatusOk() (*RuleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleStatus, true
}

// SetRuleStatus sets field value
func (o *RuleReport) SetRuleStatus(v RuleStatus) {
	o.RuleStatus = v
}

// GetContVers returns the ContVers field value if set, zero value otherwise.
func (o *RuleReport) GetContVers() []int32 {
	if o == nil || IsNil(o.ContVers) {
		var ret []int32
		return ret
	}
	return o.ContVers
}

// GetContVersOk returns a tuple with the ContVers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleReport) GetContVersOk() ([]int32, bool) {
	if o == nil || IsNil(o.ContVers) {
		return nil, false
	}
	return o.ContVers, true
}

// HasContVers returns a boolean if a field has been set.
func (o *RuleReport) HasContVers() bool {
	if o != nil && !IsNil(o.ContVers) {
		return true
	}

	return false
}

// SetContVers gets a reference to the given []int32 and assigns it to the ContVers field.
func (o *RuleReport) SetContVers(v []int32) {
	o.ContVers = v
}

// GetFailureCode returns the FailureCode field value if set, zero value otherwise.
func (o *RuleReport) GetFailureCode() FailureCode {
	if o == nil || IsNil(o.FailureCode) {
		var ret FailureCode
		return ret
	}
	return *o.FailureCode
}

// GetFailureCodeOk returns a tuple with the FailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleReport) GetFailureCodeOk() (*FailureCode, bool) {
	if o == nil || IsNil(o.FailureCode) {
		return nil, false
	}
	return o.FailureCode, true
}

// HasFailureCode returns a boolean if a field has been set.
func (o *RuleReport) HasFailureCode() bool {
	if o != nil && !IsNil(o.FailureCode) {
		return true
	}

	return false
}

// SetFailureCode gets a reference to the given FailureCode and assigns it to the FailureCode field.
func (o *RuleReport) SetFailureCode(v FailureCode) {
	o.FailureCode = &v
}

// GetRetryAfter returns the RetryAfter field value if set, zero value otherwise.
func (o *RuleReport) GetRetryAfter() int32 {
	if o == nil || IsNil(o.RetryAfter) {
		var ret int32
		return ret
	}
	return *o.RetryAfter
}

// GetRetryAfterOk returns a tuple with the RetryAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleReport) GetRetryAfterOk() (*int32, bool) {
	if o == nil || IsNil(o.RetryAfter) {
		return nil, false
	}
	return o.RetryAfter, true
}

// HasRetryAfter returns a boolean if a field has been set.
func (o *RuleReport) HasRetryAfter() bool {
	if o != nil && !IsNil(o.RetryAfter) {
		return true
	}

	return false
}

// SetRetryAfter gets a reference to the given int32 and assigns it to the RetryAfter field.
func (o *RuleReport) SetRetryAfter(v int32) {
	o.RetryAfter = &v
}

// GetFinUnitAct returns the FinUnitAct field value if set, zero value otherwise.
func (o *RuleReport) GetFinUnitAct() FinalUnitAction {
	if o == nil || IsNil(o.FinUnitAct) {
		var ret FinalUnitAction
		return ret
	}
	return *o.FinUnitAct
}

// GetFinUnitActOk returns a tuple with the FinUnitAct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleReport) GetFinUnitActOk() (*FinalUnitAction, bool) {
	if o == nil || IsNil(o.FinUnitAct) {
		return nil, false
	}
	return o.FinUnitAct, true
}

// HasFinUnitAct returns a boolean if a field has been set.
func (o *RuleReport) HasFinUnitAct() bool {
	if o != nil && !IsNil(o.FinUnitAct) {
		return true
	}

	return false
}

// SetFinUnitAct gets a reference to the given FinalUnitAction and assigns it to the FinUnitAct field.
func (o *RuleReport) SetFinUnitAct(v FinalUnitAction) {
	o.FinUnitAct = &v
}

// GetRanNasRelCauses returns the RanNasRelCauses field value if set, zero value otherwise.
func (o *RuleReport) GetRanNasRelCauses() []RanNasRelCause {
	if o == nil || IsNil(o.RanNasRelCauses) {
		var ret []RanNasRelCause
		return ret
	}
	return o.RanNasRelCauses
}

// GetRanNasRelCausesOk returns a tuple with the RanNasRelCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleReport) GetRanNasRelCausesOk() ([]RanNasRelCause, bool) {
	if o == nil || IsNil(o.RanNasRelCauses) {
		return nil, false
	}
	return o.RanNasRelCauses, true
}

// HasRanNasRelCauses returns a boolean if a field has been set.
func (o *RuleReport) HasRanNasRelCauses() bool {
	if o != nil && !IsNil(o.RanNasRelCauses) {
		return true
	}

	return false
}

// SetRanNasRelCauses gets a reference to the given []RanNasRelCause and assigns it to the RanNasRelCauses field.
func (o *RuleReport) SetRanNasRelCauses(v []RanNasRelCause) {
	o.RanNasRelCauses = v
}

// GetAltQosParamId returns the AltQosParamId field value if set, zero value otherwise.
func (o *RuleReport) GetAltQosParamId() string {
	if o == nil || IsNil(o.AltQosParamId) {
		var ret string
		return ret
	}
	return *o.AltQosParamId
}

// GetAltQosParamIdOk returns a tuple with the AltQosParamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleReport) GetAltQosParamIdOk() (*string, bool) {
	if o == nil || IsNil(o.AltQosParamId) {
		return nil, false
	}
	return o.AltQosParamId, true
}

// HasAltQosParamId returns a boolean if a field has been set.
func (o *RuleReport) HasAltQosParamId() bool {
	if o != nil && !IsNil(o.AltQosParamId) {
		return true
	}

	return false
}

// SetAltQosParamId gets a reference to the given string and assigns it to the AltQosParamId field.
func (o *RuleReport) SetAltQosParamId(v string) {
	o.AltQosParamId = &v
}

// GetAltQosNotSuppInd returns the AltQosNotSuppInd field value if set, zero value otherwise.
func (o *RuleReport) GetAltQosNotSuppInd() bool {
	if o == nil || IsNil(o.AltQosNotSuppInd) {
		var ret bool
		return ret
	}
	return *o.AltQosNotSuppInd
}

// GetAltQosNotSuppIndOk returns a tuple with the AltQosNotSuppInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleReport) GetAltQosNotSuppIndOk() (*bool, bool) {
	if o == nil || IsNil(o.AltQosNotSuppInd) {
		return nil, false
	}
	return o.AltQosNotSuppInd, true
}

// HasAltQosNotSuppInd returns a boolean if a field has been set.
func (o *RuleReport) HasAltQosNotSuppInd() bool {
	if o != nil && !IsNil(o.AltQosNotSuppInd) {
		return true
	}

	return false
}

// SetAltQosNotSuppInd gets a reference to the given bool and assigns it to the AltQosNotSuppInd field.
func (o *RuleReport) SetAltQosNotSuppInd(v bool) {
	o.AltQosNotSuppInd = &v
}

func (o RuleReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pccRuleIds"] = o.PccRuleIds
	toSerialize["ruleStatus"] = o.RuleStatus
	if !IsNil(o.ContVers) {
		toSerialize["contVers"] = o.ContVers
	}
	if !IsNil(o.FailureCode) {
		toSerialize["failureCode"] = o.FailureCode
	}
	if !IsNil(o.RetryAfter) {
		toSerialize["retryAfter"] = o.RetryAfter
	}
	if !IsNil(o.FinUnitAct) {
		toSerialize["finUnitAct"] = o.FinUnitAct
	}
	if !IsNil(o.RanNasRelCauses) {
		toSerialize["ranNasRelCauses"] = o.RanNasRelCauses
	}
	if !IsNil(o.AltQosParamId) {
		toSerialize["altQosParamId"] = o.AltQosParamId
	}
	if !IsNil(o.AltQosNotSuppInd) {
		toSerialize["altQosNotSuppInd"] = o.AltQosNotSuppInd
	}
	return toSerialize, nil
}

type NullableRuleReport struct {
	value *RuleReport
	isSet bool
}

func (v NullableRuleReport) Get() *RuleReport {
	return v.value
}

func (v *NullableRuleReport) Set(val *RuleReport) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleReport) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleReport(val *RuleReport) *NullableRuleReport {
	return &NullableRuleReport{value: val, isSet: true}
}

func (v NullableRuleReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

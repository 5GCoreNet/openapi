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

// checks if the ErrorReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorReport{}

// ErrorReport Contains the rule,policy decision and/or condition data error reports.
type ErrorReport struct {
	Error *ProblemDetails `json:"error,omitempty"`
	// Used to report the PCC rule failure.
	RuleReports []RuleReport `json:"ruleReports,omitempty"`
	// Used to report the session rule failure.
	SessRuleReports []SessionRuleReport `json:"sessRuleReports,omitempty"`
	// Used to report failure of the policy decision and/or condition data.
	PolDecFailureReports []PolicyDecisionFailureCode `json:"polDecFailureReports,omitempty"`
	// Indicates the invalid parameters for the reported type(s) of the failed policy decision  and/or condition data. 
	InvalidPolicyDecs []InvalidParam `json:"invalidPolicyDecs,omitempty"`
}

// NewErrorReport instantiates a new ErrorReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorReport() *ErrorReport {
	this := ErrorReport{}
	return &this
}

// NewErrorReportWithDefaults instantiates a new ErrorReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorReportWithDefaults() *ErrorReport {
	this := ErrorReport{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ErrorReport) GetError() ProblemDetails {
	if o == nil || isNil(o.Error) {
		var ret ProblemDetails
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorReport) GetErrorOk() (*ProblemDetails, bool) {
	if o == nil || isNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ErrorReport) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ProblemDetails and assigns it to the Error field.
func (o *ErrorReport) SetError(v ProblemDetails) {
	o.Error = &v
}

// GetRuleReports returns the RuleReports field value if set, zero value otherwise.
func (o *ErrorReport) GetRuleReports() []RuleReport {
	if o == nil || isNil(o.RuleReports) {
		var ret []RuleReport
		return ret
	}
	return o.RuleReports
}

// GetRuleReportsOk returns a tuple with the RuleReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorReport) GetRuleReportsOk() ([]RuleReport, bool) {
	if o == nil || isNil(o.RuleReports) {
		return nil, false
	}
	return o.RuleReports, true
}

// HasRuleReports returns a boolean if a field has been set.
func (o *ErrorReport) HasRuleReports() bool {
	if o != nil && !isNil(o.RuleReports) {
		return true
	}

	return false
}

// SetRuleReports gets a reference to the given []RuleReport and assigns it to the RuleReports field.
func (o *ErrorReport) SetRuleReports(v []RuleReport) {
	o.RuleReports = v
}

// GetSessRuleReports returns the SessRuleReports field value if set, zero value otherwise.
func (o *ErrorReport) GetSessRuleReports() []SessionRuleReport {
	if o == nil || isNil(o.SessRuleReports) {
		var ret []SessionRuleReport
		return ret
	}
	return o.SessRuleReports
}

// GetSessRuleReportsOk returns a tuple with the SessRuleReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorReport) GetSessRuleReportsOk() ([]SessionRuleReport, bool) {
	if o == nil || isNil(o.SessRuleReports) {
		return nil, false
	}
	return o.SessRuleReports, true
}

// HasSessRuleReports returns a boolean if a field has been set.
func (o *ErrorReport) HasSessRuleReports() bool {
	if o != nil && !isNil(o.SessRuleReports) {
		return true
	}

	return false
}

// SetSessRuleReports gets a reference to the given []SessionRuleReport and assigns it to the SessRuleReports field.
func (o *ErrorReport) SetSessRuleReports(v []SessionRuleReport) {
	o.SessRuleReports = v
}

// GetPolDecFailureReports returns the PolDecFailureReports field value if set, zero value otherwise.
func (o *ErrorReport) GetPolDecFailureReports() []PolicyDecisionFailureCode {
	if o == nil || isNil(o.PolDecFailureReports) {
		var ret []PolicyDecisionFailureCode
		return ret
	}
	return o.PolDecFailureReports
}

// GetPolDecFailureReportsOk returns a tuple with the PolDecFailureReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorReport) GetPolDecFailureReportsOk() ([]PolicyDecisionFailureCode, bool) {
	if o == nil || isNil(o.PolDecFailureReports) {
		return nil, false
	}
	return o.PolDecFailureReports, true
}

// HasPolDecFailureReports returns a boolean if a field has been set.
func (o *ErrorReport) HasPolDecFailureReports() bool {
	if o != nil && !isNil(o.PolDecFailureReports) {
		return true
	}

	return false
}

// SetPolDecFailureReports gets a reference to the given []PolicyDecisionFailureCode and assigns it to the PolDecFailureReports field.
func (o *ErrorReport) SetPolDecFailureReports(v []PolicyDecisionFailureCode) {
	o.PolDecFailureReports = v
}

// GetInvalidPolicyDecs returns the InvalidPolicyDecs field value if set, zero value otherwise.
func (o *ErrorReport) GetInvalidPolicyDecs() []InvalidParam {
	if o == nil || isNil(o.InvalidPolicyDecs) {
		var ret []InvalidParam
		return ret
	}
	return o.InvalidPolicyDecs
}

// GetInvalidPolicyDecsOk returns a tuple with the InvalidPolicyDecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorReport) GetInvalidPolicyDecsOk() ([]InvalidParam, bool) {
	if o == nil || isNil(o.InvalidPolicyDecs) {
		return nil, false
	}
	return o.InvalidPolicyDecs, true
}

// HasInvalidPolicyDecs returns a boolean if a field has been set.
func (o *ErrorReport) HasInvalidPolicyDecs() bool {
	if o != nil && !isNil(o.InvalidPolicyDecs) {
		return true
	}

	return false
}

// SetInvalidPolicyDecs gets a reference to the given []InvalidParam and assigns it to the InvalidPolicyDecs field.
func (o *ErrorReport) SetInvalidPolicyDecs(v []InvalidParam) {
	o.InvalidPolicyDecs = v
}

func (o ErrorReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.RuleReports) {
		toSerialize["ruleReports"] = o.RuleReports
	}
	if !isNil(o.SessRuleReports) {
		toSerialize["sessRuleReports"] = o.SessRuleReports
	}
	if !isNil(o.PolDecFailureReports) {
		toSerialize["polDecFailureReports"] = o.PolDecFailureReports
	}
	if !isNil(o.InvalidPolicyDecs) {
		toSerialize["invalidPolicyDecs"] = o.InvalidPolicyDecs
	}
	return toSerialize, nil
}

type NullableErrorReport struct {
	value *ErrorReport
	isSet bool
}

func (v NullableErrorReport) Get() *ErrorReport {
	return v.value
}

func (v *NullableErrorReport) Set(val *ErrorReport) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorReport) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorReport(val *ErrorReport) *NullableErrorReport {
	return &NullableErrorReport{value: val, isSet: true}
}

func (v NullableErrorReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



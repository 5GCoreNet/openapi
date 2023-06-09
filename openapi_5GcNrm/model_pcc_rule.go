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

// checks if the PccRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PccRule{}

// PccRule struct for PccRule
type PccRule struct {
	// Univocally identifies the PCC rule within a PDU session.
	PccRuleId     *string           `json:"pccRuleId,omitempty"`
	FlowInfoList  []FlowInformation `json:"flowInfoList,omitempty"`
	ApplicationId *string           `json:"applicationId,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	AppDescriptor *string `json:"appDescriptor,omitempty"`
	// Represents the content version of some content.
	ContentVersion *int32 `json:"contentVersion,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Precedence         *int32                      `json:"precedence,omitempty"`
	AfSigProtocol      *AfSigProtocol              `json:"afSigProtocol,omitempty"`
	IsAppRelocatable   *bool                       `json:"isAppRelocatable,omitempty"`
	IsUeAddrPreserved  *bool                       `json:"isUeAddrPreserved,omitempty"`
	QosData            [][]QosData                 `json:"qosData,omitempty"`
	AltQosParams       [][]QosData                 `json:"altQosParams,omitempty"`
	TrafficControlData [][]TrafficControlData      `json:"trafficControlData,omitempty"`
	ConditionData      NullableConditionData       `json:"conditionData,omitempty"`
	TscaiInputDl       NullableTscaiInputContainer `json:"tscaiInputDl,omitempty"`
	TscaiInputUl       NullableTscaiInputContainer `json:"tscaiInputUl,omitempty"`
}

// NewPccRule instantiates a new PccRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPccRule() *PccRule {
	this := PccRule{}
	return &this
}

// NewPccRuleWithDefaults instantiates a new PccRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPccRuleWithDefaults() *PccRule {
	this := PccRule{}
	return &this
}

// GetPccRuleId returns the PccRuleId field value if set, zero value otherwise.
func (o *PccRule) GetPccRuleId() string {
	if o == nil || IsNil(o.PccRuleId) {
		var ret string
		return ret
	}
	return *o.PccRuleId
}

// GetPccRuleIdOk returns a tuple with the PccRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetPccRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.PccRuleId) {
		return nil, false
	}
	return o.PccRuleId, true
}

// HasPccRuleId returns a boolean if a field has been set.
func (o *PccRule) HasPccRuleId() bool {
	if o != nil && !IsNil(o.PccRuleId) {
		return true
	}

	return false
}

// SetPccRuleId gets a reference to the given string and assigns it to the PccRuleId field.
func (o *PccRule) SetPccRuleId(v string) {
	o.PccRuleId = &v
}

// GetFlowInfoList returns the FlowInfoList field value if set, zero value otherwise.
func (o *PccRule) GetFlowInfoList() []FlowInformation {
	if o == nil || IsNil(o.FlowInfoList) {
		var ret []FlowInformation
		return ret
	}
	return o.FlowInfoList
}

// GetFlowInfoListOk returns a tuple with the FlowInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetFlowInfoListOk() ([]FlowInformation, bool) {
	if o == nil || IsNil(o.FlowInfoList) {
		return nil, false
	}
	return o.FlowInfoList, true
}

// HasFlowInfoList returns a boolean if a field has been set.
func (o *PccRule) HasFlowInfoList() bool {
	if o != nil && !IsNil(o.FlowInfoList) {
		return true
	}

	return false
}

// SetFlowInfoList gets a reference to the given []FlowInformation and assigns it to the FlowInfoList field.
func (o *PccRule) SetFlowInfoList(v []FlowInformation) {
	o.FlowInfoList = v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *PccRule) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *PccRule) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *PccRule) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetAppDescriptor returns the AppDescriptor field value if set, zero value otherwise.
func (o *PccRule) GetAppDescriptor() string {
	if o == nil || IsNil(o.AppDescriptor) {
		var ret string
		return ret
	}
	return *o.AppDescriptor
}

// GetAppDescriptorOk returns a tuple with the AppDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetAppDescriptorOk() (*string, bool) {
	if o == nil || IsNil(o.AppDescriptor) {
		return nil, false
	}
	return o.AppDescriptor, true
}

// HasAppDescriptor returns a boolean if a field has been set.
func (o *PccRule) HasAppDescriptor() bool {
	if o != nil && !IsNil(o.AppDescriptor) {
		return true
	}

	return false
}

// SetAppDescriptor gets a reference to the given string and assigns it to the AppDescriptor field.
func (o *PccRule) SetAppDescriptor(v string) {
	o.AppDescriptor = &v
}

// GetContentVersion returns the ContentVersion field value if set, zero value otherwise.
func (o *PccRule) GetContentVersion() int32 {
	if o == nil || IsNil(o.ContentVersion) {
		var ret int32
		return ret
	}
	return *o.ContentVersion
}

// GetContentVersionOk returns a tuple with the ContentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetContentVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.ContentVersion) {
		return nil, false
	}
	return o.ContentVersion, true
}

// HasContentVersion returns a boolean if a field has been set.
func (o *PccRule) HasContentVersion() bool {
	if o != nil && !IsNil(o.ContentVersion) {
		return true
	}

	return false
}

// SetContentVersion gets a reference to the given int32 and assigns it to the ContentVersion field.
func (o *PccRule) SetContentVersion(v int32) {
	o.ContentVersion = &v
}

// GetPrecedence returns the Precedence field value if set, zero value otherwise.
func (o *PccRule) GetPrecedence() int32 {
	if o == nil || IsNil(o.Precedence) {
		var ret int32
		return ret
	}
	return *o.Precedence
}

// GetPrecedenceOk returns a tuple with the Precedence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetPrecedenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Precedence) {
		return nil, false
	}
	return o.Precedence, true
}

// HasPrecedence returns a boolean if a field has been set.
func (o *PccRule) HasPrecedence() bool {
	if o != nil && !IsNil(o.Precedence) {
		return true
	}

	return false
}

// SetPrecedence gets a reference to the given int32 and assigns it to the Precedence field.
func (o *PccRule) SetPrecedence(v int32) {
	o.Precedence = &v
}

// GetAfSigProtocol returns the AfSigProtocol field value if set, zero value otherwise.
func (o *PccRule) GetAfSigProtocol() AfSigProtocol {
	if o == nil || IsNil(o.AfSigProtocol) {
		var ret AfSigProtocol
		return ret
	}
	return *o.AfSigProtocol
}

// GetAfSigProtocolOk returns a tuple with the AfSigProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetAfSigProtocolOk() (*AfSigProtocol, bool) {
	if o == nil || IsNil(o.AfSigProtocol) {
		return nil, false
	}
	return o.AfSigProtocol, true
}

// HasAfSigProtocol returns a boolean if a field has been set.
func (o *PccRule) HasAfSigProtocol() bool {
	if o != nil && !IsNil(o.AfSigProtocol) {
		return true
	}

	return false
}

// SetAfSigProtocol gets a reference to the given AfSigProtocol and assigns it to the AfSigProtocol field.
func (o *PccRule) SetAfSigProtocol(v AfSigProtocol) {
	o.AfSigProtocol = &v
}

// GetIsAppRelocatable returns the IsAppRelocatable field value if set, zero value otherwise.
func (o *PccRule) GetIsAppRelocatable() bool {
	if o == nil || IsNil(o.IsAppRelocatable) {
		var ret bool
		return ret
	}
	return *o.IsAppRelocatable
}

// GetIsAppRelocatableOk returns a tuple with the IsAppRelocatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetIsAppRelocatableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAppRelocatable) {
		return nil, false
	}
	return o.IsAppRelocatable, true
}

// HasIsAppRelocatable returns a boolean if a field has been set.
func (o *PccRule) HasIsAppRelocatable() bool {
	if o != nil && !IsNil(o.IsAppRelocatable) {
		return true
	}

	return false
}

// SetIsAppRelocatable gets a reference to the given bool and assigns it to the IsAppRelocatable field.
func (o *PccRule) SetIsAppRelocatable(v bool) {
	o.IsAppRelocatable = &v
}

// GetIsUeAddrPreserved returns the IsUeAddrPreserved field value if set, zero value otherwise.
func (o *PccRule) GetIsUeAddrPreserved() bool {
	if o == nil || IsNil(o.IsUeAddrPreserved) {
		var ret bool
		return ret
	}
	return *o.IsUeAddrPreserved
}

// GetIsUeAddrPreservedOk returns a tuple with the IsUeAddrPreserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetIsUeAddrPreservedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUeAddrPreserved) {
		return nil, false
	}
	return o.IsUeAddrPreserved, true
}

// HasIsUeAddrPreserved returns a boolean if a field has been set.
func (o *PccRule) HasIsUeAddrPreserved() bool {
	if o != nil && !IsNil(o.IsUeAddrPreserved) {
		return true
	}

	return false
}

// SetIsUeAddrPreserved gets a reference to the given bool and assigns it to the IsUeAddrPreserved field.
func (o *PccRule) SetIsUeAddrPreserved(v bool) {
	o.IsUeAddrPreserved = &v
}

// GetQosData returns the QosData field value if set, zero value otherwise.
func (o *PccRule) GetQosData() [][]QosData {
	if o == nil || IsNil(o.QosData) {
		var ret [][]QosData
		return ret
	}
	return o.QosData
}

// GetQosDataOk returns a tuple with the QosData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetQosDataOk() ([][]QosData, bool) {
	if o == nil || IsNil(o.QosData) {
		return nil, false
	}
	return o.QosData, true
}

// HasQosData returns a boolean if a field has been set.
func (o *PccRule) HasQosData() bool {
	if o != nil && !IsNil(o.QosData) {
		return true
	}

	return false
}

// SetQosData gets a reference to the given [][]QosData and assigns it to the QosData field.
func (o *PccRule) SetQosData(v [][]QosData) {
	o.QosData = v
}

// GetAltQosParams returns the AltQosParams field value if set, zero value otherwise.
func (o *PccRule) GetAltQosParams() [][]QosData {
	if o == nil || IsNil(o.AltQosParams) {
		var ret [][]QosData
		return ret
	}
	return o.AltQosParams
}

// GetAltQosParamsOk returns a tuple with the AltQosParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetAltQosParamsOk() ([][]QosData, bool) {
	if o == nil || IsNil(o.AltQosParams) {
		return nil, false
	}
	return o.AltQosParams, true
}

// HasAltQosParams returns a boolean if a field has been set.
func (o *PccRule) HasAltQosParams() bool {
	if o != nil && !IsNil(o.AltQosParams) {
		return true
	}

	return false
}

// SetAltQosParams gets a reference to the given [][]QosData and assigns it to the AltQosParams field.
func (o *PccRule) SetAltQosParams(v [][]QosData) {
	o.AltQosParams = v
}

// GetTrafficControlData returns the TrafficControlData field value if set, zero value otherwise.
func (o *PccRule) GetTrafficControlData() [][]TrafficControlData {
	if o == nil || IsNil(o.TrafficControlData) {
		var ret [][]TrafficControlData
		return ret
	}
	return o.TrafficControlData
}

// GetTrafficControlDataOk returns a tuple with the TrafficControlData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetTrafficControlDataOk() ([][]TrafficControlData, bool) {
	if o == nil || IsNil(o.TrafficControlData) {
		return nil, false
	}
	return o.TrafficControlData, true
}

// HasTrafficControlData returns a boolean if a field has been set.
func (o *PccRule) HasTrafficControlData() bool {
	if o != nil && !IsNil(o.TrafficControlData) {
		return true
	}

	return false
}

// SetTrafficControlData gets a reference to the given [][]TrafficControlData and assigns it to the TrafficControlData field.
func (o *PccRule) SetTrafficControlData(v [][]TrafficControlData) {
	o.TrafficControlData = v
}

// GetConditionData returns the ConditionData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetConditionData() ConditionData {
	if o == nil || IsNil(o.ConditionData.Get()) {
		var ret ConditionData
		return ret
	}
	return *o.ConditionData.Get()
}

// GetConditionDataOk returns a tuple with the ConditionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetConditionDataOk() (*ConditionData, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConditionData.Get(), o.ConditionData.IsSet()
}

// HasConditionData returns a boolean if a field has been set.
func (o *PccRule) HasConditionData() bool {
	if o != nil && o.ConditionData.IsSet() {
		return true
	}

	return false
}

// SetConditionData gets a reference to the given NullableConditionData and assigns it to the ConditionData field.
func (o *PccRule) SetConditionData(v ConditionData) {
	o.ConditionData.Set(&v)
}

// SetConditionDataNil sets the value for ConditionData to be an explicit nil
func (o *PccRule) SetConditionDataNil() {
	o.ConditionData.Set(nil)
}

// UnsetConditionData ensures that no value is present for ConditionData, not even an explicit nil
func (o *PccRule) UnsetConditionData() {
	o.ConditionData.Unset()
}

// GetTscaiInputDl returns the TscaiInputDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetTscaiInputDl() TscaiInputContainer {
	if o == nil || IsNil(o.TscaiInputDl.Get()) {
		var ret TscaiInputContainer
		return ret
	}
	return *o.TscaiInputDl.Get()
}

// GetTscaiInputDlOk returns a tuple with the TscaiInputDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetTscaiInputDlOk() (*TscaiInputContainer, bool) {
	if o == nil {
		return nil, false
	}
	return o.TscaiInputDl.Get(), o.TscaiInputDl.IsSet()
}

// HasTscaiInputDl returns a boolean if a field has been set.
func (o *PccRule) HasTscaiInputDl() bool {
	if o != nil && o.TscaiInputDl.IsSet() {
		return true
	}

	return false
}

// SetTscaiInputDl gets a reference to the given NullableTscaiInputContainer and assigns it to the TscaiInputDl field.
func (o *PccRule) SetTscaiInputDl(v TscaiInputContainer) {
	o.TscaiInputDl.Set(&v)
}

// SetTscaiInputDlNil sets the value for TscaiInputDl to be an explicit nil
func (o *PccRule) SetTscaiInputDlNil() {
	o.TscaiInputDl.Set(nil)
}

// UnsetTscaiInputDl ensures that no value is present for TscaiInputDl, not even an explicit nil
func (o *PccRule) UnsetTscaiInputDl() {
	o.TscaiInputDl.Unset()
}

// GetTscaiInputUl returns the TscaiInputUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetTscaiInputUl() TscaiInputContainer {
	if o == nil || IsNil(o.TscaiInputUl.Get()) {
		var ret TscaiInputContainer
		return ret
	}
	return *o.TscaiInputUl.Get()
}

// GetTscaiInputUlOk returns a tuple with the TscaiInputUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetTscaiInputUlOk() (*TscaiInputContainer, bool) {
	if o == nil {
		return nil, false
	}
	return o.TscaiInputUl.Get(), o.TscaiInputUl.IsSet()
}

// HasTscaiInputUl returns a boolean if a field has been set.
func (o *PccRule) HasTscaiInputUl() bool {
	if o != nil && o.TscaiInputUl.IsSet() {
		return true
	}

	return false
}

// SetTscaiInputUl gets a reference to the given NullableTscaiInputContainer and assigns it to the TscaiInputUl field.
func (o *PccRule) SetTscaiInputUl(v TscaiInputContainer) {
	o.TscaiInputUl.Set(&v)
}

// SetTscaiInputUlNil sets the value for TscaiInputUl to be an explicit nil
func (o *PccRule) SetTscaiInputUlNil() {
	o.TscaiInputUl.Set(nil)
}

// UnsetTscaiInputUl ensures that no value is present for TscaiInputUl, not even an explicit nil
func (o *PccRule) UnsetTscaiInputUl() {
	o.TscaiInputUl.Unset()
}

func (o PccRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PccRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PccRuleId) {
		toSerialize["pccRuleId"] = o.PccRuleId
	}
	if !IsNil(o.FlowInfoList) {
		toSerialize["flowInfoList"] = o.FlowInfoList
	}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.AppDescriptor) {
		toSerialize["appDescriptor"] = o.AppDescriptor
	}
	if !IsNil(o.ContentVersion) {
		toSerialize["contentVersion"] = o.ContentVersion
	}
	if !IsNil(o.Precedence) {
		toSerialize["precedence"] = o.Precedence
	}
	if !IsNil(o.AfSigProtocol) {
		toSerialize["afSigProtocol"] = o.AfSigProtocol
	}
	if !IsNil(o.IsAppRelocatable) {
		toSerialize["isAppRelocatable"] = o.IsAppRelocatable
	}
	if !IsNil(o.IsUeAddrPreserved) {
		toSerialize["isUeAddrPreserved"] = o.IsUeAddrPreserved
	}
	if !IsNil(o.QosData) {
		toSerialize["qosData"] = o.QosData
	}
	if !IsNil(o.AltQosParams) {
		toSerialize["altQosParams"] = o.AltQosParams
	}
	if !IsNil(o.TrafficControlData) {
		toSerialize["trafficControlData"] = o.TrafficControlData
	}
	if o.ConditionData.IsSet() {
		toSerialize["conditionData"] = o.ConditionData.Get()
	}
	if o.TscaiInputDl.IsSet() {
		toSerialize["tscaiInputDl"] = o.TscaiInputDl.Get()
	}
	if o.TscaiInputUl.IsSet() {
		toSerialize["tscaiInputUl"] = o.TscaiInputUl.Get()
	}
	return toSerialize, nil
}

type NullablePccRule struct {
	value *PccRule
	isSet bool
}

func (v NullablePccRule) Get() *PccRule {
	return v.value
}

func (v *NullablePccRule) Set(val *PccRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePccRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePccRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePccRule(val *PccRule) *NullablePccRule {
	return &NullablePccRule{value: val, isSet: true}
}

func (v NullablePccRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePccRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

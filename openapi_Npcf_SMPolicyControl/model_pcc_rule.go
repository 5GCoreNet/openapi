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

// checks if the PccRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PccRule{}

// PccRule Contains a PCC rule information.
type PccRule struct {
	// An array of IP flow packet filter information.
	FlowInfos []FlowInformation `json:"flowInfos,omitempty"`
	// A reference to the application detection filter configured at the UPF.
	AppId *string `json:"appId,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	AppDescriptor *string `json:"appDescriptor,omitempty"`
	// Represents the content version of some content.
	ContVer *int32 `json:"contVer,omitempty"`
	// Univocally identifies the PCC rule within a PDU session.
	PccRuleId string `json:"pccRuleId"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Precedence *int32 `json:"precedence,omitempty"`
	AfSigProtocol *AfSigProtocol `json:"afSigProtocol,omitempty"`
	// Indication of application relocation possibility.
	AppReloc *bool `json:"appReloc,omitempty"`
	// Indicates the EAS rediscovery is required.
	EasRedisInd *bool `json:"easRedisInd,omitempty"`
	// A reference to the QosData policy decision type. It is the qosId described in  clause 5.6.2.8. 
	RefQosData []string `json:"refQosData,omitempty"`
	// A Reference to the QosData policy decision type for the Alternative QoS parameter sets  of the service data flow. 
	RefAltQosParams []string `json:"refAltQosParams,omitempty"`
	// A reference to the TrafficControlData policy decision type. It is the tcId described in  clause 5.6.2.10. 
	RefTcData []string `json:"refTcData,omitempty"`
	// A reference to the ChargingData policy decision type. It is the chgId described in  clause 5.6.2.11. 
	RefChgData []string `json:"refChgData,omitempty"`
	// A reference to the ChargingData policy decision type only applicable to Non-3GPP access if \"ATSSS\" feature is supported. It is the chgId described in clause 5.6.2.11. 
	RefChgN3gData []string `json:"refChgN3gData,omitempty"`
	// A reference to UsageMonitoringData policy decision type. It is the umId described in  clause 5.6.2.12. 
	RefUmData []string `json:"refUmData,omitempty"`
	// A reference to UsageMonitoringData policy decision type only applicable to Non-3GPP access if \"ATSSS\" feature is supported. It is the umId described in clause 5.6.2.12.  
	RefUmN3gData []string `json:"refUmN3gData,omitempty"`
	// A reference to the condition data. It is the condId described in clause 5.6.2.9. 
	RefCondData NullableString `json:"refCondData,omitempty"`
	// A reference to the QosMonitoringData policy decision type. It is the qmId described in  clause 5.6.2.40.  
	RefQosMon []string `json:"refQosMon,omitempty"`
	AddrPreserInd NullableBool `json:"addrPreserInd,omitempty"`
	TscaiInputDl NullableTscaiInputContainer `json:"tscaiInputDl,omitempty"`
	TscaiInputUl NullableTscaiInputContainer `json:"tscaiInputUl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TscaiTimeDom *int32 `json:"tscaiTimeDom,omitempty"`
	DdNotifCtrl *DownlinkDataNotificationControl `json:"ddNotifCtrl,omitempty"`
	DdNotifCtrl2 NullableDownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	DisUeNotif NullableBool `json:"disUeNotif,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PackFiltAllPrec *int32 `json:"packFiltAllPrec,omitempty"`
}

// NewPccRule instantiates a new PccRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPccRule(pccRuleId string) *PccRule {
	this := PccRule{}
	this.PccRuleId = pccRuleId
	return &this
}

// NewPccRuleWithDefaults instantiates a new PccRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPccRuleWithDefaults() *PccRule {
	this := PccRule{}
	return &this
}

// GetFlowInfos returns the FlowInfos field value if set, zero value otherwise.
func (o *PccRule) GetFlowInfos() []FlowInformation {
	if o == nil || IsNil(o.FlowInfos) {
		var ret []FlowInformation
		return ret
	}
	return o.FlowInfos
}

// GetFlowInfosOk returns a tuple with the FlowInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetFlowInfosOk() ([]FlowInformation, bool) {
	if o == nil || IsNil(o.FlowInfos) {
		return nil, false
	}
	return o.FlowInfos, true
}

// HasFlowInfos returns a boolean if a field has been set.
func (o *PccRule) HasFlowInfos() bool {
	if o != nil && !IsNil(o.FlowInfos) {
		return true
	}

	return false
}

// SetFlowInfos gets a reference to the given []FlowInformation and assigns it to the FlowInfos field.
func (o *PccRule) SetFlowInfos(v []FlowInformation) {
	o.FlowInfos = v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *PccRule) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *PccRule) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *PccRule) SetAppId(v string) {
	o.AppId = &v
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

// GetContVer returns the ContVer field value if set, zero value otherwise.
func (o *PccRule) GetContVer() int32 {
	if o == nil || IsNil(o.ContVer) {
		var ret int32
		return ret
	}
	return *o.ContVer
}

// GetContVerOk returns a tuple with the ContVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetContVerOk() (*int32, bool) {
	if o == nil || IsNil(o.ContVer) {
		return nil, false
	}
	return o.ContVer, true
}

// HasContVer returns a boolean if a field has been set.
func (o *PccRule) HasContVer() bool {
	if o != nil && !IsNil(o.ContVer) {
		return true
	}

	return false
}

// SetContVer gets a reference to the given int32 and assigns it to the ContVer field.
func (o *PccRule) SetContVer(v int32) {
	o.ContVer = &v
}

// GetPccRuleId returns the PccRuleId field value
func (o *PccRule) GetPccRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PccRuleId
}

// GetPccRuleIdOk returns a tuple with the PccRuleId field value
// and a boolean to check if the value has been set.
func (o *PccRule) GetPccRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PccRuleId, true
}

// SetPccRuleId sets field value
func (o *PccRule) SetPccRuleId(v string) {
	o.PccRuleId = v
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

// GetAppReloc returns the AppReloc field value if set, zero value otherwise.
func (o *PccRule) GetAppReloc() bool {
	if o == nil || IsNil(o.AppReloc) {
		var ret bool
		return ret
	}
	return *o.AppReloc
}

// GetAppRelocOk returns a tuple with the AppReloc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetAppRelocOk() (*bool, bool) {
	if o == nil || IsNil(o.AppReloc) {
		return nil, false
	}
	return o.AppReloc, true
}

// HasAppReloc returns a boolean if a field has been set.
func (o *PccRule) HasAppReloc() bool {
	if o != nil && !IsNil(o.AppReloc) {
		return true
	}

	return false
}

// SetAppReloc gets a reference to the given bool and assigns it to the AppReloc field.
func (o *PccRule) SetAppReloc(v bool) {
	o.AppReloc = &v
}

// GetEasRedisInd returns the EasRedisInd field value if set, zero value otherwise.
func (o *PccRule) GetEasRedisInd() bool {
	if o == nil || IsNil(o.EasRedisInd) {
		var ret bool
		return ret
	}
	return *o.EasRedisInd
}

// GetEasRedisIndOk returns a tuple with the EasRedisInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetEasRedisIndOk() (*bool, bool) {
	if o == nil || IsNil(o.EasRedisInd) {
		return nil, false
	}
	return o.EasRedisInd, true
}

// HasEasRedisInd returns a boolean if a field has been set.
func (o *PccRule) HasEasRedisInd() bool {
	if o != nil && !IsNil(o.EasRedisInd) {
		return true
	}

	return false
}

// SetEasRedisInd gets a reference to the given bool and assigns it to the EasRedisInd field.
func (o *PccRule) SetEasRedisInd(v bool) {
	o.EasRedisInd = &v
}

// GetRefQosData returns the RefQosData field value if set, zero value otherwise.
func (o *PccRule) GetRefQosData() []string {
	if o == nil || IsNil(o.RefQosData) {
		var ret []string
		return ret
	}
	return o.RefQosData
}

// GetRefQosDataOk returns a tuple with the RefQosData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetRefQosDataOk() ([]string, bool) {
	if o == nil || IsNil(o.RefQosData) {
		return nil, false
	}
	return o.RefQosData, true
}

// HasRefQosData returns a boolean if a field has been set.
func (o *PccRule) HasRefQosData() bool {
	if o != nil && !IsNil(o.RefQosData) {
		return true
	}

	return false
}

// SetRefQosData gets a reference to the given []string and assigns it to the RefQosData field.
func (o *PccRule) SetRefQosData(v []string) {
	o.RefQosData = v
}

// GetRefAltQosParams returns the RefAltQosParams field value if set, zero value otherwise.
func (o *PccRule) GetRefAltQosParams() []string {
	if o == nil || IsNil(o.RefAltQosParams) {
		var ret []string
		return ret
	}
	return o.RefAltQosParams
}

// GetRefAltQosParamsOk returns a tuple with the RefAltQosParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetRefAltQosParamsOk() ([]string, bool) {
	if o == nil || IsNil(o.RefAltQosParams) {
		return nil, false
	}
	return o.RefAltQosParams, true
}

// HasRefAltQosParams returns a boolean if a field has been set.
func (o *PccRule) HasRefAltQosParams() bool {
	if o != nil && !IsNil(o.RefAltQosParams) {
		return true
	}

	return false
}

// SetRefAltQosParams gets a reference to the given []string and assigns it to the RefAltQosParams field.
func (o *PccRule) SetRefAltQosParams(v []string) {
	o.RefAltQosParams = v
}

// GetRefTcData returns the RefTcData field value if set, zero value otherwise.
func (o *PccRule) GetRefTcData() []string {
	if o == nil || IsNil(o.RefTcData) {
		var ret []string
		return ret
	}
	return o.RefTcData
}

// GetRefTcDataOk returns a tuple with the RefTcData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetRefTcDataOk() ([]string, bool) {
	if o == nil || IsNil(o.RefTcData) {
		return nil, false
	}
	return o.RefTcData, true
}

// HasRefTcData returns a boolean if a field has been set.
func (o *PccRule) HasRefTcData() bool {
	if o != nil && !IsNil(o.RefTcData) {
		return true
	}

	return false
}

// SetRefTcData gets a reference to the given []string and assigns it to the RefTcData field.
func (o *PccRule) SetRefTcData(v []string) {
	o.RefTcData = v
}

// GetRefChgData returns the RefChgData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetRefChgData() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RefChgData
}

// GetRefChgDataOk returns a tuple with the RefChgData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetRefChgDataOk() ([]string, bool) {
	if o == nil || IsNil(o.RefChgData) {
		return nil, false
	}
	return o.RefChgData, true
}

// HasRefChgData returns a boolean if a field has been set.
func (o *PccRule) HasRefChgData() bool {
	if o != nil && IsNil(o.RefChgData) {
		return true
	}

	return false
}

// SetRefChgData gets a reference to the given []string and assigns it to the RefChgData field.
func (o *PccRule) SetRefChgData(v []string) {
	o.RefChgData = v
}

// GetRefChgN3gData returns the RefChgN3gData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetRefChgN3gData() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RefChgN3gData
}

// GetRefChgN3gDataOk returns a tuple with the RefChgN3gData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetRefChgN3gDataOk() ([]string, bool) {
	if o == nil || IsNil(o.RefChgN3gData) {
		return nil, false
	}
	return o.RefChgN3gData, true
}

// HasRefChgN3gData returns a boolean if a field has been set.
func (o *PccRule) HasRefChgN3gData() bool {
	if o != nil && IsNil(o.RefChgN3gData) {
		return true
	}

	return false
}

// SetRefChgN3gData gets a reference to the given []string and assigns it to the RefChgN3gData field.
func (o *PccRule) SetRefChgN3gData(v []string) {
	o.RefChgN3gData = v
}

// GetRefUmData returns the RefUmData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetRefUmData() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RefUmData
}

// GetRefUmDataOk returns a tuple with the RefUmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetRefUmDataOk() ([]string, bool) {
	if o == nil || IsNil(o.RefUmData) {
		return nil, false
	}
	return o.RefUmData, true
}

// HasRefUmData returns a boolean if a field has been set.
func (o *PccRule) HasRefUmData() bool {
	if o != nil && IsNil(o.RefUmData) {
		return true
	}

	return false
}

// SetRefUmData gets a reference to the given []string and assigns it to the RefUmData field.
func (o *PccRule) SetRefUmData(v []string) {
	o.RefUmData = v
}

// GetRefUmN3gData returns the RefUmN3gData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetRefUmN3gData() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RefUmN3gData
}

// GetRefUmN3gDataOk returns a tuple with the RefUmN3gData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetRefUmN3gDataOk() ([]string, bool) {
	if o == nil || IsNil(o.RefUmN3gData) {
		return nil, false
	}
	return o.RefUmN3gData, true
}

// HasRefUmN3gData returns a boolean if a field has been set.
func (o *PccRule) HasRefUmN3gData() bool {
	if o != nil && IsNil(o.RefUmN3gData) {
		return true
	}

	return false
}

// SetRefUmN3gData gets a reference to the given []string and assigns it to the RefUmN3gData field.
func (o *PccRule) SetRefUmN3gData(v []string) {
	o.RefUmN3gData = v
}

// GetRefCondData returns the RefCondData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetRefCondData() string {
	if o == nil || IsNil(o.RefCondData.Get()) {
		var ret string
		return ret
	}
	return *o.RefCondData.Get()
}

// GetRefCondDataOk returns a tuple with the RefCondData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetRefCondDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefCondData.Get(), o.RefCondData.IsSet()
}

// HasRefCondData returns a boolean if a field has been set.
func (o *PccRule) HasRefCondData() bool {
	if o != nil && o.RefCondData.IsSet() {
		return true
	}

	return false
}

// SetRefCondData gets a reference to the given NullableString and assigns it to the RefCondData field.
func (o *PccRule) SetRefCondData(v string) {
	o.RefCondData.Set(&v)
}
// SetRefCondDataNil sets the value for RefCondData to be an explicit nil
func (o *PccRule) SetRefCondDataNil() {
	o.RefCondData.Set(nil)
}

// UnsetRefCondData ensures that no value is present for RefCondData, not even an explicit nil
func (o *PccRule) UnsetRefCondData() {
	o.RefCondData.Unset()
}

// GetRefQosMon returns the RefQosMon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetRefQosMon() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RefQosMon
}

// GetRefQosMonOk returns a tuple with the RefQosMon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetRefQosMonOk() ([]string, bool) {
	if o == nil || IsNil(o.RefQosMon) {
		return nil, false
	}
	return o.RefQosMon, true
}

// HasRefQosMon returns a boolean if a field has been set.
func (o *PccRule) HasRefQosMon() bool {
	if o != nil && IsNil(o.RefQosMon) {
		return true
	}

	return false
}

// SetRefQosMon gets a reference to the given []string and assigns it to the RefQosMon field.
func (o *PccRule) SetRefQosMon(v []string) {
	o.RefQosMon = v
}

// GetAddrPreserInd returns the AddrPreserInd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetAddrPreserInd() bool {
	if o == nil || IsNil(o.AddrPreserInd.Get()) {
		var ret bool
		return ret
	}
	return *o.AddrPreserInd.Get()
}

// GetAddrPreserIndOk returns a tuple with the AddrPreserInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetAddrPreserIndOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddrPreserInd.Get(), o.AddrPreserInd.IsSet()
}

// HasAddrPreserInd returns a boolean if a field has been set.
func (o *PccRule) HasAddrPreserInd() bool {
	if o != nil && o.AddrPreserInd.IsSet() {
		return true
	}

	return false
}

// SetAddrPreserInd gets a reference to the given NullableBool and assigns it to the AddrPreserInd field.
func (o *PccRule) SetAddrPreserInd(v bool) {
	o.AddrPreserInd.Set(&v)
}
// SetAddrPreserIndNil sets the value for AddrPreserInd to be an explicit nil
func (o *PccRule) SetAddrPreserIndNil() {
	o.AddrPreserInd.Set(nil)
}

// UnsetAddrPreserInd ensures that no value is present for AddrPreserInd, not even an explicit nil
func (o *PccRule) UnsetAddrPreserInd() {
	o.AddrPreserInd.Unset()
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

// GetTscaiTimeDom returns the TscaiTimeDom field value if set, zero value otherwise.
func (o *PccRule) GetTscaiTimeDom() int32 {
	if o == nil || IsNil(o.TscaiTimeDom) {
		var ret int32
		return ret
	}
	return *o.TscaiTimeDom
}

// GetTscaiTimeDomOk returns a tuple with the TscaiTimeDom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetTscaiTimeDomOk() (*int32, bool) {
	if o == nil || IsNil(o.TscaiTimeDom) {
		return nil, false
	}
	return o.TscaiTimeDom, true
}

// HasTscaiTimeDom returns a boolean if a field has been set.
func (o *PccRule) HasTscaiTimeDom() bool {
	if o != nil && !IsNil(o.TscaiTimeDom) {
		return true
	}

	return false
}

// SetTscaiTimeDom gets a reference to the given int32 and assigns it to the TscaiTimeDom field.
func (o *PccRule) SetTscaiTimeDom(v int32) {
	o.TscaiTimeDom = &v
}

// GetDdNotifCtrl returns the DdNotifCtrl field value if set, zero value otherwise.
func (o *PccRule) GetDdNotifCtrl() DownlinkDataNotificationControl {
	if o == nil || IsNil(o.DdNotifCtrl) {
		var ret DownlinkDataNotificationControl
		return ret
	}
	return *o.DdNotifCtrl
}

// GetDdNotifCtrlOk returns a tuple with the DdNotifCtrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetDdNotifCtrlOk() (*DownlinkDataNotificationControl, bool) {
	if o == nil || IsNil(o.DdNotifCtrl) {
		return nil, false
	}
	return o.DdNotifCtrl, true
}

// HasDdNotifCtrl returns a boolean if a field has been set.
func (o *PccRule) HasDdNotifCtrl() bool {
	if o != nil && !IsNil(o.DdNotifCtrl) {
		return true
	}

	return false
}

// SetDdNotifCtrl gets a reference to the given DownlinkDataNotificationControl and assigns it to the DdNotifCtrl field.
func (o *PccRule) SetDdNotifCtrl(v DownlinkDataNotificationControl) {
	o.DdNotifCtrl = &v
}

// GetDdNotifCtrl2 returns the DdNotifCtrl2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetDdNotifCtrl2() DownlinkDataNotificationControlRm {
	if o == nil || IsNil(o.DdNotifCtrl2.Get()) {
		var ret DownlinkDataNotificationControlRm
		return ret
	}
	return *o.DdNotifCtrl2.Get()
}

// GetDdNotifCtrl2Ok returns a tuple with the DdNotifCtrl2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetDdNotifCtrl2Ok() (*DownlinkDataNotificationControlRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.DdNotifCtrl2.Get(), o.DdNotifCtrl2.IsSet()
}

// HasDdNotifCtrl2 returns a boolean if a field has been set.
func (o *PccRule) HasDdNotifCtrl2() bool {
	if o != nil && o.DdNotifCtrl2.IsSet() {
		return true
	}

	return false
}

// SetDdNotifCtrl2 gets a reference to the given NullableDownlinkDataNotificationControlRm and assigns it to the DdNotifCtrl2 field.
func (o *PccRule) SetDdNotifCtrl2(v DownlinkDataNotificationControlRm) {
	o.DdNotifCtrl2.Set(&v)
}
// SetDdNotifCtrl2Nil sets the value for DdNotifCtrl2 to be an explicit nil
func (o *PccRule) SetDdNotifCtrl2Nil() {
	o.DdNotifCtrl2.Set(nil)
}

// UnsetDdNotifCtrl2 ensures that no value is present for DdNotifCtrl2, not even an explicit nil
func (o *PccRule) UnsetDdNotifCtrl2() {
	o.DdNotifCtrl2.Unset()
}

// GetDisUeNotif returns the DisUeNotif field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PccRule) GetDisUeNotif() bool {
	if o == nil || IsNil(o.DisUeNotif.Get()) {
		var ret bool
		return ret
	}
	return *o.DisUeNotif.Get()
}

// GetDisUeNotifOk returns a tuple with the DisUeNotif field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PccRule) GetDisUeNotifOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisUeNotif.Get(), o.DisUeNotif.IsSet()
}

// HasDisUeNotif returns a boolean if a field has been set.
func (o *PccRule) HasDisUeNotif() bool {
	if o != nil && o.DisUeNotif.IsSet() {
		return true
	}

	return false
}

// SetDisUeNotif gets a reference to the given NullableBool and assigns it to the DisUeNotif field.
func (o *PccRule) SetDisUeNotif(v bool) {
	o.DisUeNotif.Set(&v)
}
// SetDisUeNotifNil sets the value for DisUeNotif to be an explicit nil
func (o *PccRule) SetDisUeNotifNil() {
	o.DisUeNotif.Set(nil)
}

// UnsetDisUeNotif ensures that no value is present for DisUeNotif, not even an explicit nil
func (o *PccRule) UnsetDisUeNotif() {
	o.DisUeNotif.Unset()
}

// GetPackFiltAllPrec returns the PackFiltAllPrec field value if set, zero value otherwise.
func (o *PccRule) GetPackFiltAllPrec() int32 {
	if o == nil || IsNil(o.PackFiltAllPrec) {
		var ret int32
		return ret
	}
	return *o.PackFiltAllPrec
}

// GetPackFiltAllPrecOk returns a tuple with the PackFiltAllPrec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PccRule) GetPackFiltAllPrecOk() (*int32, bool) {
	if o == nil || IsNil(o.PackFiltAllPrec) {
		return nil, false
	}
	return o.PackFiltAllPrec, true
}

// HasPackFiltAllPrec returns a boolean if a field has been set.
func (o *PccRule) HasPackFiltAllPrec() bool {
	if o != nil && !IsNil(o.PackFiltAllPrec) {
		return true
	}

	return false
}

// SetPackFiltAllPrec gets a reference to the given int32 and assigns it to the PackFiltAllPrec field.
func (o *PccRule) SetPackFiltAllPrec(v int32) {
	o.PackFiltAllPrec = &v
}

func (o PccRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PccRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlowInfos) {
		toSerialize["flowInfos"] = o.FlowInfos
	}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.AppDescriptor) {
		toSerialize["appDescriptor"] = o.AppDescriptor
	}
	if !IsNil(o.ContVer) {
		toSerialize["contVer"] = o.ContVer
	}
	toSerialize["pccRuleId"] = o.PccRuleId
	if !IsNil(o.Precedence) {
		toSerialize["precedence"] = o.Precedence
	}
	if !IsNil(o.AfSigProtocol) {
		toSerialize["afSigProtocol"] = o.AfSigProtocol
	}
	if !IsNil(o.AppReloc) {
		toSerialize["appReloc"] = o.AppReloc
	}
	if !IsNil(o.EasRedisInd) {
		toSerialize["easRedisInd"] = o.EasRedisInd
	}
	if !IsNil(o.RefQosData) {
		toSerialize["refQosData"] = o.RefQosData
	}
	if !IsNil(o.RefAltQosParams) {
		toSerialize["refAltQosParams"] = o.RefAltQosParams
	}
	if !IsNil(o.RefTcData) {
		toSerialize["refTcData"] = o.RefTcData
	}
	if o.RefChgData != nil {
		toSerialize["refChgData"] = o.RefChgData
	}
	if o.RefChgN3gData != nil {
		toSerialize["refChgN3gData"] = o.RefChgN3gData
	}
	if o.RefUmData != nil {
		toSerialize["refUmData"] = o.RefUmData
	}
	if o.RefUmN3gData != nil {
		toSerialize["refUmN3gData"] = o.RefUmN3gData
	}
	if o.RefCondData.IsSet() {
		toSerialize["refCondData"] = o.RefCondData.Get()
	}
	if o.RefQosMon != nil {
		toSerialize["refQosMon"] = o.RefQosMon
	}
	if o.AddrPreserInd.IsSet() {
		toSerialize["addrPreserInd"] = o.AddrPreserInd.Get()
	}
	if o.TscaiInputDl.IsSet() {
		toSerialize["tscaiInputDl"] = o.TscaiInputDl.Get()
	}
	if o.TscaiInputUl.IsSet() {
		toSerialize["tscaiInputUl"] = o.TscaiInputUl.Get()
	}
	if !IsNil(o.TscaiTimeDom) {
		toSerialize["tscaiTimeDom"] = o.TscaiTimeDom
	}
	if !IsNil(o.DdNotifCtrl) {
		toSerialize["ddNotifCtrl"] = o.DdNotifCtrl
	}
	if o.DdNotifCtrl2.IsSet() {
		toSerialize["ddNotifCtrl2"] = o.DdNotifCtrl2.Get()
	}
	if o.DisUeNotif.IsSet() {
		toSerialize["disUeNotif"] = o.DisUeNotif.Get()
	}
	if !IsNil(o.PackFiltAllPrec) {
		toSerialize["packFiltAllPrec"] = o.PackFiltAllPrec
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


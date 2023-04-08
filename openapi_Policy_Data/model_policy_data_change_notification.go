/*
Unified Data Repository Service API file for policy data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Policy_Data

import (
	"encoding/json"
)

// checks if the PolicyDataChangeNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyDataChangeNotification{}

// PolicyDataChangeNotification Contains changed policy data for which notification was requested.
type PolicyDataChangeNotification struct {
	AmPolicyData *AmPolicyData `json:"amPolicyData,omitempty"`
	UePolicySet *UePolicySet `json:"uePolicySet,omitempty"`
	PlmnUePolicySet *UePolicySet `json:"plmnUePolicySet,omitempty"`
	SmPolicyData *SmPolicyData `json:"smPolicyData,omitempty"`
	UsageMonData *UsageMonData `json:"usageMonData,omitempty"`
	SponsorConnectivityData *SponsorConnectivityData `json:"SponsorConnectivityData,omitempty"`
	BdtData *BdtData `json:"bdtData,omitempty"`
	OpSpecData *OperatorSpecificDataContainer `json:"opSpecData,omitempty"`
	// Operator Specific Data resource data, if changed and notification was requested. The key of the map is operator specific data element name and the value is the operator specific data of the UE. 
	OpSpecDataMap *map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`
	// String represents the SUPI or GPSI
	UeId *string `json:"ueId,omitempty"`
	SponsorId *string `json:"sponsorId,omitempty"`
	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	BdtRefId *string `json:"bdtRefId,omitempty"`
	UsageMonId *string `json:"usageMonId,omitempty"`
	PlmnId *PlmnId `json:"plmnId,omitempty"`
	DelResources []string `json:"delResources,omitempty"`
	NotifId *string `json:"notifId,omitempty"`
	ReportedFragments []NotificationItem `json:"reportedFragments,omitempty"`
	SlicePolicyData *SlicePolicyData `json:"slicePolicyData,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
}

// NewPolicyDataChangeNotification instantiates a new PolicyDataChangeNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyDataChangeNotification() *PolicyDataChangeNotification {
	this := PolicyDataChangeNotification{}
	return &this
}

// NewPolicyDataChangeNotificationWithDefaults instantiates a new PolicyDataChangeNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyDataChangeNotificationWithDefaults() *PolicyDataChangeNotification {
	this := PolicyDataChangeNotification{}
	return &this
}

// GetAmPolicyData returns the AmPolicyData field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetAmPolicyData() AmPolicyData {
	if o == nil || isNil(o.AmPolicyData) {
		var ret AmPolicyData
		return ret
	}
	return *o.AmPolicyData
}

// GetAmPolicyDataOk returns a tuple with the AmPolicyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetAmPolicyDataOk() (*AmPolicyData, bool) {
	if o == nil || isNil(o.AmPolicyData) {
		return nil, false
	}
	return o.AmPolicyData, true
}

// HasAmPolicyData returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasAmPolicyData() bool {
	if o != nil && !isNil(o.AmPolicyData) {
		return true
	}

	return false
}

// SetAmPolicyData gets a reference to the given AmPolicyData and assigns it to the AmPolicyData field.
func (o *PolicyDataChangeNotification) SetAmPolicyData(v AmPolicyData) {
	o.AmPolicyData = &v
}

// GetUePolicySet returns the UePolicySet field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetUePolicySet() UePolicySet {
	if o == nil || isNil(o.UePolicySet) {
		var ret UePolicySet
		return ret
	}
	return *o.UePolicySet
}

// GetUePolicySetOk returns a tuple with the UePolicySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetUePolicySetOk() (*UePolicySet, bool) {
	if o == nil || isNil(o.UePolicySet) {
		return nil, false
	}
	return o.UePolicySet, true
}

// HasUePolicySet returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasUePolicySet() bool {
	if o != nil && !isNil(o.UePolicySet) {
		return true
	}

	return false
}

// SetUePolicySet gets a reference to the given UePolicySet and assigns it to the UePolicySet field.
func (o *PolicyDataChangeNotification) SetUePolicySet(v UePolicySet) {
	o.UePolicySet = &v
}

// GetPlmnUePolicySet returns the PlmnUePolicySet field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetPlmnUePolicySet() UePolicySet {
	if o == nil || isNil(o.PlmnUePolicySet) {
		var ret UePolicySet
		return ret
	}
	return *o.PlmnUePolicySet
}

// GetPlmnUePolicySetOk returns a tuple with the PlmnUePolicySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetPlmnUePolicySetOk() (*UePolicySet, bool) {
	if o == nil || isNil(o.PlmnUePolicySet) {
		return nil, false
	}
	return o.PlmnUePolicySet, true
}

// HasPlmnUePolicySet returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasPlmnUePolicySet() bool {
	if o != nil && !isNil(o.PlmnUePolicySet) {
		return true
	}

	return false
}

// SetPlmnUePolicySet gets a reference to the given UePolicySet and assigns it to the PlmnUePolicySet field.
func (o *PolicyDataChangeNotification) SetPlmnUePolicySet(v UePolicySet) {
	o.PlmnUePolicySet = &v
}

// GetSmPolicyData returns the SmPolicyData field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetSmPolicyData() SmPolicyData {
	if o == nil || isNil(o.SmPolicyData) {
		var ret SmPolicyData
		return ret
	}
	return *o.SmPolicyData
}

// GetSmPolicyDataOk returns a tuple with the SmPolicyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetSmPolicyDataOk() (*SmPolicyData, bool) {
	if o == nil || isNil(o.SmPolicyData) {
		return nil, false
	}
	return o.SmPolicyData, true
}

// HasSmPolicyData returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasSmPolicyData() bool {
	if o != nil && !isNil(o.SmPolicyData) {
		return true
	}

	return false
}

// SetSmPolicyData gets a reference to the given SmPolicyData and assigns it to the SmPolicyData field.
func (o *PolicyDataChangeNotification) SetSmPolicyData(v SmPolicyData) {
	o.SmPolicyData = &v
}

// GetUsageMonData returns the UsageMonData field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetUsageMonData() UsageMonData {
	if o == nil || isNil(o.UsageMonData) {
		var ret UsageMonData
		return ret
	}
	return *o.UsageMonData
}

// GetUsageMonDataOk returns a tuple with the UsageMonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetUsageMonDataOk() (*UsageMonData, bool) {
	if o == nil || isNil(o.UsageMonData) {
		return nil, false
	}
	return o.UsageMonData, true
}

// HasUsageMonData returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasUsageMonData() bool {
	if o != nil && !isNil(o.UsageMonData) {
		return true
	}

	return false
}

// SetUsageMonData gets a reference to the given UsageMonData and assigns it to the UsageMonData field.
func (o *PolicyDataChangeNotification) SetUsageMonData(v UsageMonData) {
	o.UsageMonData = &v
}

// GetSponsorConnectivityData returns the SponsorConnectivityData field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetSponsorConnectivityData() SponsorConnectivityData {
	if o == nil || isNil(o.SponsorConnectivityData) {
		var ret SponsorConnectivityData
		return ret
	}
	return *o.SponsorConnectivityData
}

// GetSponsorConnectivityDataOk returns a tuple with the SponsorConnectivityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetSponsorConnectivityDataOk() (*SponsorConnectivityData, bool) {
	if o == nil || isNil(o.SponsorConnectivityData) {
		return nil, false
	}
	return o.SponsorConnectivityData, true
}

// HasSponsorConnectivityData returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasSponsorConnectivityData() bool {
	if o != nil && !isNil(o.SponsorConnectivityData) {
		return true
	}

	return false
}

// SetSponsorConnectivityData gets a reference to the given SponsorConnectivityData and assigns it to the SponsorConnectivityData field.
func (o *PolicyDataChangeNotification) SetSponsorConnectivityData(v SponsorConnectivityData) {
	o.SponsorConnectivityData = &v
}

// GetBdtData returns the BdtData field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetBdtData() BdtData {
	if o == nil || isNil(o.BdtData) {
		var ret BdtData
		return ret
	}
	return *o.BdtData
}

// GetBdtDataOk returns a tuple with the BdtData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetBdtDataOk() (*BdtData, bool) {
	if o == nil || isNil(o.BdtData) {
		return nil, false
	}
	return o.BdtData, true
}

// HasBdtData returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasBdtData() bool {
	if o != nil && !isNil(o.BdtData) {
		return true
	}

	return false
}

// SetBdtData gets a reference to the given BdtData and assigns it to the BdtData field.
func (o *PolicyDataChangeNotification) SetBdtData(v BdtData) {
	o.BdtData = &v
}

// GetOpSpecData returns the OpSpecData field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetOpSpecData() OperatorSpecificDataContainer {
	if o == nil || isNil(o.OpSpecData) {
		var ret OperatorSpecificDataContainer
		return ret
	}
	return *o.OpSpecData
}

// GetOpSpecDataOk returns a tuple with the OpSpecData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetOpSpecDataOk() (*OperatorSpecificDataContainer, bool) {
	if o == nil || isNil(o.OpSpecData) {
		return nil, false
	}
	return o.OpSpecData, true
}

// HasOpSpecData returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasOpSpecData() bool {
	if o != nil && !isNil(o.OpSpecData) {
		return true
	}

	return false
}

// SetOpSpecData gets a reference to the given OperatorSpecificDataContainer and assigns it to the OpSpecData field.
func (o *PolicyDataChangeNotification) SetOpSpecData(v OperatorSpecificDataContainer) {
	o.OpSpecData = &v
}

// GetOpSpecDataMap returns the OpSpecDataMap field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetOpSpecDataMap() map[string]OperatorSpecificDataContainer {
	if o == nil || isNil(o.OpSpecDataMap) {
		var ret map[string]OperatorSpecificDataContainer
		return ret
	}
	return *o.OpSpecDataMap
}

// GetOpSpecDataMapOk returns a tuple with the OpSpecDataMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetOpSpecDataMapOk() (*map[string]OperatorSpecificDataContainer, bool) {
	if o == nil || isNil(o.OpSpecDataMap) {
		return nil, false
	}
	return o.OpSpecDataMap, true
}

// HasOpSpecDataMap returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasOpSpecDataMap() bool {
	if o != nil && !isNil(o.OpSpecDataMap) {
		return true
	}

	return false
}

// SetOpSpecDataMap gets a reference to the given map[string]OperatorSpecificDataContainer and assigns it to the OpSpecDataMap field.
func (o *PolicyDataChangeNotification) SetOpSpecDataMap(v map[string]OperatorSpecificDataContainer) {
	o.OpSpecDataMap = &v
}

// GetUeId returns the UeId field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetUeId() string {
	if o == nil || isNil(o.UeId) {
		var ret string
		return ret
	}
	return *o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetUeIdOk() (*string, bool) {
	if o == nil || isNil(o.UeId) {
		return nil, false
	}
	return o.UeId, true
}

// HasUeId returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasUeId() bool {
	if o != nil && !isNil(o.UeId) {
		return true
	}

	return false
}

// SetUeId gets a reference to the given string and assigns it to the UeId field.
func (o *PolicyDataChangeNotification) SetUeId(v string) {
	o.UeId = &v
}

// GetSponsorId returns the SponsorId field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetSponsorId() string {
	if o == nil || isNil(o.SponsorId) {
		var ret string
		return ret
	}
	return *o.SponsorId
}

// GetSponsorIdOk returns a tuple with the SponsorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetSponsorIdOk() (*string, bool) {
	if o == nil || isNil(o.SponsorId) {
		return nil, false
	}
	return o.SponsorId, true
}

// HasSponsorId returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasSponsorId() bool {
	if o != nil && !isNil(o.SponsorId) {
		return true
	}

	return false
}

// SetSponsorId gets a reference to the given string and assigns it to the SponsorId field.
func (o *PolicyDataChangeNotification) SetSponsorId(v string) {
	o.SponsorId = &v
}

// GetBdtRefId returns the BdtRefId field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetBdtRefId() string {
	if o == nil || isNil(o.BdtRefId) {
		var ret string
		return ret
	}
	return *o.BdtRefId
}

// GetBdtRefIdOk returns a tuple with the BdtRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetBdtRefIdOk() (*string, bool) {
	if o == nil || isNil(o.BdtRefId) {
		return nil, false
	}
	return o.BdtRefId, true
}

// HasBdtRefId returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasBdtRefId() bool {
	if o != nil && !isNil(o.BdtRefId) {
		return true
	}

	return false
}

// SetBdtRefId gets a reference to the given string and assigns it to the BdtRefId field.
func (o *PolicyDataChangeNotification) SetBdtRefId(v string) {
	o.BdtRefId = &v
}

// GetUsageMonId returns the UsageMonId field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetUsageMonId() string {
	if o == nil || isNil(o.UsageMonId) {
		var ret string
		return ret
	}
	return *o.UsageMonId
}

// GetUsageMonIdOk returns a tuple with the UsageMonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetUsageMonIdOk() (*string, bool) {
	if o == nil || isNil(o.UsageMonId) {
		return nil, false
	}
	return o.UsageMonId, true
}

// HasUsageMonId returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasUsageMonId() bool {
	if o != nil && !isNil(o.UsageMonId) {
		return true
	}

	return false
}

// SetUsageMonId gets a reference to the given string and assigns it to the UsageMonId field.
func (o *PolicyDataChangeNotification) SetUsageMonId(v string) {
	o.UsageMonId = &v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetPlmnId() PlmnId {
	if o == nil || isNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || isNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasPlmnId() bool {
	if o != nil && !isNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *PolicyDataChangeNotification) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetDelResources returns the DelResources field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetDelResources() []string {
	if o == nil || isNil(o.DelResources) {
		var ret []string
		return ret
	}
	return o.DelResources
}

// GetDelResourcesOk returns a tuple with the DelResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetDelResourcesOk() ([]string, bool) {
	if o == nil || isNil(o.DelResources) {
		return nil, false
	}
	return o.DelResources, true
}

// HasDelResources returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasDelResources() bool {
	if o != nil && !isNil(o.DelResources) {
		return true
	}

	return false
}

// SetDelResources gets a reference to the given []string and assigns it to the DelResources field.
func (o *PolicyDataChangeNotification) SetDelResources(v []string) {
	o.DelResources = v
}

// GetNotifId returns the NotifId field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetNotifId() string {
	if o == nil || isNil(o.NotifId) {
		var ret string
		return ret
	}
	return *o.NotifId
}

// GetNotifIdOk returns a tuple with the NotifId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetNotifIdOk() (*string, bool) {
	if o == nil || isNil(o.NotifId) {
		return nil, false
	}
	return o.NotifId, true
}

// HasNotifId returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasNotifId() bool {
	if o != nil && !isNil(o.NotifId) {
		return true
	}

	return false
}

// SetNotifId gets a reference to the given string and assigns it to the NotifId field.
func (o *PolicyDataChangeNotification) SetNotifId(v string) {
	o.NotifId = &v
}

// GetReportedFragments returns the ReportedFragments field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetReportedFragments() []NotificationItem {
	if o == nil || isNil(o.ReportedFragments) {
		var ret []NotificationItem
		return ret
	}
	return o.ReportedFragments
}

// GetReportedFragmentsOk returns a tuple with the ReportedFragments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetReportedFragmentsOk() ([]NotificationItem, bool) {
	if o == nil || isNil(o.ReportedFragments) {
		return nil, false
	}
	return o.ReportedFragments, true
}

// HasReportedFragments returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasReportedFragments() bool {
	if o != nil && !isNil(o.ReportedFragments) {
		return true
	}

	return false
}

// SetReportedFragments gets a reference to the given []NotificationItem and assigns it to the ReportedFragments field.
func (o *PolicyDataChangeNotification) SetReportedFragments(v []NotificationItem) {
	o.ReportedFragments = v
}

// GetSlicePolicyData returns the SlicePolicyData field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetSlicePolicyData() SlicePolicyData {
	if o == nil || isNil(o.SlicePolicyData) {
		var ret SlicePolicyData
		return ret
	}
	return *o.SlicePolicyData
}

// GetSlicePolicyDataOk returns a tuple with the SlicePolicyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetSlicePolicyDataOk() (*SlicePolicyData, bool) {
	if o == nil || isNil(o.SlicePolicyData) {
		return nil, false
	}
	return o.SlicePolicyData, true
}

// HasSlicePolicyData returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasSlicePolicyData() bool {
	if o != nil && !isNil(o.SlicePolicyData) {
		return true
	}

	return false
}

// SetSlicePolicyData gets a reference to the given SlicePolicyData and assigns it to the SlicePolicyData field.
func (o *PolicyDataChangeNotification) SetSlicePolicyData(v SlicePolicyData) {
	o.SlicePolicyData = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *PolicyDataChangeNotification) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataChangeNotification) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *PolicyDataChangeNotification) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *PolicyDataChangeNotification) SetSnssai(v Snssai) {
	o.Snssai = &v
}

func (o PolicyDataChangeNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyDataChangeNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AmPolicyData) {
		toSerialize["amPolicyData"] = o.AmPolicyData
	}
	if !isNil(o.UePolicySet) {
		toSerialize["uePolicySet"] = o.UePolicySet
	}
	if !isNil(o.PlmnUePolicySet) {
		toSerialize["plmnUePolicySet"] = o.PlmnUePolicySet
	}
	if !isNil(o.SmPolicyData) {
		toSerialize["smPolicyData"] = o.SmPolicyData
	}
	if !isNil(o.UsageMonData) {
		toSerialize["usageMonData"] = o.UsageMonData
	}
	if !isNil(o.SponsorConnectivityData) {
		toSerialize["SponsorConnectivityData"] = o.SponsorConnectivityData
	}
	if !isNil(o.BdtData) {
		toSerialize["bdtData"] = o.BdtData
	}
	if !isNil(o.OpSpecData) {
		toSerialize["opSpecData"] = o.OpSpecData
	}
	if !isNil(o.OpSpecDataMap) {
		toSerialize["opSpecDataMap"] = o.OpSpecDataMap
	}
	if !isNil(o.UeId) {
		toSerialize["ueId"] = o.UeId
	}
	if !isNil(o.SponsorId) {
		toSerialize["sponsorId"] = o.SponsorId
	}
	if !isNil(o.BdtRefId) {
		toSerialize["bdtRefId"] = o.BdtRefId
	}
	if !isNil(o.UsageMonId) {
		toSerialize["usageMonId"] = o.UsageMonId
	}
	if !isNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !isNil(o.DelResources) {
		toSerialize["delResources"] = o.DelResources
	}
	if !isNil(o.NotifId) {
		toSerialize["notifId"] = o.NotifId
	}
	if !isNil(o.ReportedFragments) {
		toSerialize["reportedFragments"] = o.ReportedFragments
	}
	if !isNil(o.SlicePolicyData) {
		toSerialize["slicePolicyData"] = o.SlicePolicyData
	}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	return toSerialize, nil
}

type NullablePolicyDataChangeNotification struct {
	value *PolicyDataChangeNotification
	isSet bool
}

func (v NullablePolicyDataChangeNotification) Get() *PolicyDataChangeNotification {
	return v.value
}

func (v *NullablePolicyDataChangeNotification) Set(val *PolicyDataChangeNotification) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyDataChangeNotification) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyDataChangeNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyDataChangeNotification(val *PolicyDataChangeNotification) *NullablePolicyDataChangeNotification {
	return &NullablePolicyDataChangeNotification{value: val, isSet: true}
}

func (v NullablePolicyDataChangeNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyDataChangeNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



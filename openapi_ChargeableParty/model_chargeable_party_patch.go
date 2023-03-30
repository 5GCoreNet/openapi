/*
3gpp-chargeable-party

API for Chargeable Party management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ChargeableParty

import (
	"encoding/json"
)

// checks if the ChargeablePartyPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeablePartyPatch{}

// ChargeablePartyPatch Represents a modification request of a chargeable party resource.
type ChargeablePartyPatch struct {
	// Describes the IP flows.
	FlowInfo []FlowInfo `json:"flowInfo,omitempty"`
	// Identifies the external Application Identifier.
	ExterAppId *string `json:"exterAppId,omitempty"`
	// Identifies Ethernet packet flows.
	EthFlowInfo []EthFlowDescription `json:"ethFlowInfo,omitempty"`
	// Indicates whether the sponsoring data connectivity is enabled (true) or not (false). 
	SponsoringEnabled *bool `json:"sponsoringEnabled,omitempty"`
	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	ReferenceId *string `json:"referenceId,omitempty"`
	UsageThreshold NullableUsageThresholdRm `json:"usageThreshold,omitempty"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
	// Represents the list of event(s) to which the SCS/AS requests to subscribe to.
	Events []Event `json:"events,omitempty"`
}

// NewChargeablePartyPatch instantiates a new ChargeablePartyPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeablePartyPatch() *ChargeablePartyPatch {
	this := ChargeablePartyPatch{}
	return &this
}

// NewChargeablePartyPatchWithDefaults instantiates a new ChargeablePartyPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeablePartyPatchWithDefaults() *ChargeablePartyPatch {
	this := ChargeablePartyPatch{}
	return &this
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *ChargeablePartyPatch) GetFlowInfo() []FlowInfo {
	if o == nil || IsNil(o.FlowInfo) {
		var ret []FlowInfo
		return ret
	}
	return o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeablePartyPatch) GetFlowInfoOk() ([]FlowInfo, bool) {
	if o == nil || IsNil(o.FlowInfo) {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *ChargeablePartyPatch) HasFlowInfo() bool {
	if o != nil && !IsNil(o.FlowInfo) {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given []FlowInfo and assigns it to the FlowInfo field.
func (o *ChargeablePartyPatch) SetFlowInfo(v []FlowInfo) {
	o.FlowInfo = v
}

// GetExterAppId returns the ExterAppId field value if set, zero value otherwise.
func (o *ChargeablePartyPatch) GetExterAppId() string {
	if o == nil || IsNil(o.ExterAppId) {
		var ret string
		return ret
	}
	return *o.ExterAppId
}

// GetExterAppIdOk returns a tuple with the ExterAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeablePartyPatch) GetExterAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExterAppId) {
		return nil, false
	}
	return o.ExterAppId, true
}

// HasExterAppId returns a boolean if a field has been set.
func (o *ChargeablePartyPatch) HasExterAppId() bool {
	if o != nil && !IsNil(o.ExterAppId) {
		return true
	}

	return false
}

// SetExterAppId gets a reference to the given string and assigns it to the ExterAppId field.
func (o *ChargeablePartyPatch) SetExterAppId(v string) {
	o.ExterAppId = &v
}

// GetEthFlowInfo returns the EthFlowInfo field value if set, zero value otherwise.
func (o *ChargeablePartyPatch) GetEthFlowInfo() []EthFlowDescription {
	if o == nil || IsNil(o.EthFlowInfo) {
		var ret []EthFlowDescription
		return ret
	}
	return o.EthFlowInfo
}

// GetEthFlowInfoOk returns a tuple with the EthFlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeablePartyPatch) GetEthFlowInfoOk() ([]EthFlowDescription, bool) {
	if o == nil || IsNil(o.EthFlowInfo) {
		return nil, false
	}
	return o.EthFlowInfo, true
}

// HasEthFlowInfo returns a boolean if a field has been set.
func (o *ChargeablePartyPatch) HasEthFlowInfo() bool {
	if o != nil && !IsNil(o.EthFlowInfo) {
		return true
	}

	return false
}

// SetEthFlowInfo gets a reference to the given []EthFlowDescription and assigns it to the EthFlowInfo field.
func (o *ChargeablePartyPatch) SetEthFlowInfo(v []EthFlowDescription) {
	o.EthFlowInfo = v
}

// GetSponsoringEnabled returns the SponsoringEnabled field value if set, zero value otherwise.
func (o *ChargeablePartyPatch) GetSponsoringEnabled() bool {
	if o == nil || IsNil(o.SponsoringEnabled) {
		var ret bool
		return ret
	}
	return *o.SponsoringEnabled
}

// GetSponsoringEnabledOk returns a tuple with the SponsoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeablePartyPatch) GetSponsoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SponsoringEnabled) {
		return nil, false
	}
	return o.SponsoringEnabled, true
}

// HasSponsoringEnabled returns a boolean if a field has been set.
func (o *ChargeablePartyPatch) HasSponsoringEnabled() bool {
	if o != nil && !IsNil(o.SponsoringEnabled) {
		return true
	}

	return false
}

// SetSponsoringEnabled gets a reference to the given bool and assigns it to the SponsoringEnabled field.
func (o *ChargeablePartyPatch) SetSponsoringEnabled(v bool) {
	o.SponsoringEnabled = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *ChargeablePartyPatch) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeablePartyPatch) GetReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *ChargeablePartyPatch) HasReferenceId() bool {
	if o != nil && !IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *ChargeablePartyPatch) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetUsageThreshold returns the UsageThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeablePartyPatch) GetUsageThreshold() UsageThresholdRm {
	if o == nil || IsNil(o.UsageThreshold.Get()) {
		var ret UsageThresholdRm
		return ret
	}
	return *o.UsageThreshold.Get()
}

// GetUsageThresholdOk returns a tuple with the UsageThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeablePartyPatch) GetUsageThresholdOk() (*UsageThresholdRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageThreshold.Get(), o.UsageThreshold.IsSet()
}

// HasUsageThreshold returns a boolean if a field has been set.
func (o *ChargeablePartyPatch) HasUsageThreshold() bool {
	if o != nil && o.UsageThreshold.IsSet() {
		return true
	}

	return false
}

// SetUsageThreshold gets a reference to the given NullableUsageThresholdRm and assigns it to the UsageThreshold field.
func (o *ChargeablePartyPatch) SetUsageThreshold(v UsageThresholdRm) {
	o.UsageThreshold.Set(&v)
}
// SetUsageThresholdNil sets the value for UsageThreshold to be an explicit nil
func (o *ChargeablePartyPatch) SetUsageThresholdNil() {
	o.UsageThreshold.Set(nil)
}

// UnsetUsageThreshold ensures that no value is present for UsageThreshold, not even an explicit nil
func (o *ChargeablePartyPatch) UnsetUsageThreshold() {
	o.UsageThreshold.Unset()
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *ChargeablePartyPatch) GetNotificationDestination() string {
	if o == nil || IsNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeablePartyPatch) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationDestination) {
		return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *ChargeablePartyPatch) HasNotificationDestination() bool {
	if o != nil && !IsNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *ChargeablePartyPatch) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ChargeablePartyPatch) GetEvents() []Event {
	if o == nil || IsNil(o.Events) {
		var ret []Event
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeablePartyPatch) GetEventsOk() ([]Event, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ChargeablePartyPatch) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []Event and assigns it to the Events field.
func (o *ChargeablePartyPatch) SetEvents(v []Event) {
	o.Events = v
}

func (o ChargeablePartyPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeablePartyPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlowInfo) {
		toSerialize["flowInfo"] = o.FlowInfo
	}
	if !IsNil(o.ExterAppId) {
		toSerialize["exterAppId"] = o.ExterAppId
	}
	if !IsNil(o.EthFlowInfo) {
		toSerialize["ethFlowInfo"] = o.EthFlowInfo
	}
	if !IsNil(o.SponsoringEnabled) {
		toSerialize["sponsoringEnabled"] = o.SponsoringEnabled
	}
	if !IsNil(o.ReferenceId) {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if o.UsageThreshold.IsSet() {
		toSerialize["usageThreshold"] = o.UsageThreshold.Get()
	}
	if !IsNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullableChargeablePartyPatch struct {
	value *ChargeablePartyPatch
	isSet bool
}

func (v NullableChargeablePartyPatch) Get() *ChargeablePartyPatch {
	return v.value
}

func (v *NullableChargeablePartyPatch) Set(val *ChargeablePartyPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeablePartyPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeablePartyPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeablePartyPatch(val *ChargeablePartyPatch) *NullableChargeablePartyPatch {
	return &NullableChargeablePartyPatch{value: val, isSet: true}
}

func (v NullableChargeablePartyPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeablePartyPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



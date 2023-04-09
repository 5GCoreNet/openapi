/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
)

// checks if the MonitoringNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringNotification{}

// MonitoringNotification Represents an event monitoring notification.
type MonitoringNotification struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Subscription string `json:"subscription"`
	// Each element identifies a notification of grouping configuration result.
	ConfigResults []ConfigResult `json:"configResults,omitempty"`
	// Monitoring event reports.
	MonitoringEventReports []MonitoringEventReport `json:"monitoringEventReports,omitempty"`
	// Identifies the added external Identifier(s) within the active group via the \"externalGroupId\" attribute within the MonitoringEventSubscription data type.
	AddedExternalIds []string `json:"addedExternalIds,omitempty"`
	// Identifies the added MSISDN(s) within the active group via the \"externalGroupId\" attribute within the MonitoringEventSubscription data type.
	AddedMsisdns []string `json:"addedMsisdns,omitempty"`
	// Identifies the cancelled external Identifier(s) within the active group via the \"externalGroupId\" attribute within the MonitoringEventSubscription data type.
	CancelExternalIds []string `json:"cancelExternalIds,omitempty"`
	// Identifies the cancelled MSISDN(s) within the active group via the \"externalGroupId\" attribute within the MonitoringEventSubscription data type.
	CancelMsisdns []string `json:"cancelMsisdns,omitempty"`
	// Indicates whether to request to cancel the corresponding monitoring subscription. Set to false or omitted otherwise.
	CancelInd    *bool                          `json:"cancelInd,omitempty"`
	AppliedParam *AppliedParameterConfiguration `json:"appliedParam,omitempty"`
}

// NewMonitoringNotification instantiates a new MonitoringNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringNotification(subscription string) *MonitoringNotification {
	this := MonitoringNotification{}
	this.Subscription = subscription
	return &this
}

// NewMonitoringNotificationWithDefaults instantiates a new MonitoringNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringNotificationWithDefaults() *MonitoringNotification {
	this := MonitoringNotification{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *MonitoringNotification) GetSubscription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *MonitoringNotification) GetSubscriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *MonitoringNotification) SetSubscription(v string) {
	o.Subscription = v
}

// GetConfigResults returns the ConfigResults field value if set, zero value otherwise.
func (o *MonitoringNotification) GetConfigResults() []ConfigResult {
	if o == nil || IsNil(o.ConfigResults) {
		var ret []ConfigResult
		return ret
	}
	return o.ConfigResults
}

// GetConfigResultsOk returns a tuple with the ConfigResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringNotification) GetConfigResultsOk() ([]ConfigResult, bool) {
	if o == nil || IsNil(o.ConfigResults) {
		return nil, false
	}
	return o.ConfigResults, true
}

// HasConfigResults returns a boolean if a field has been set.
func (o *MonitoringNotification) HasConfigResults() bool {
	if o != nil && !IsNil(o.ConfigResults) {
		return true
	}

	return false
}

// SetConfigResults gets a reference to the given []ConfigResult and assigns it to the ConfigResults field.
func (o *MonitoringNotification) SetConfigResults(v []ConfigResult) {
	o.ConfigResults = v
}

// GetMonitoringEventReports returns the MonitoringEventReports field value if set, zero value otherwise.
func (o *MonitoringNotification) GetMonitoringEventReports() []MonitoringEventReport {
	if o == nil || IsNil(o.MonitoringEventReports) {
		var ret []MonitoringEventReport
		return ret
	}
	return o.MonitoringEventReports
}

// GetMonitoringEventReportsOk returns a tuple with the MonitoringEventReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringNotification) GetMonitoringEventReportsOk() ([]MonitoringEventReport, bool) {
	if o == nil || IsNil(o.MonitoringEventReports) {
		return nil, false
	}
	return o.MonitoringEventReports, true
}

// HasMonitoringEventReports returns a boolean if a field has been set.
func (o *MonitoringNotification) HasMonitoringEventReports() bool {
	if o != nil && !IsNil(o.MonitoringEventReports) {
		return true
	}

	return false
}

// SetMonitoringEventReports gets a reference to the given []MonitoringEventReport and assigns it to the MonitoringEventReports field.
func (o *MonitoringNotification) SetMonitoringEventReports(v []MonitoringEventReport) {
	o.MonitoringEventReports = v
}

// GetAddedExternalIds returns the AddedExternalIds field value if set, zero value otherwise.
func (o *MonitoringNotification) GetAddedExternalIds() []string {
	if o == nil || IsNil(o.AddedExternalIds) {
		var ret []string
		return ret
	}
	return o.AddedExternalIds
}

// GetAddedExternalIdsOk returns a tuple with the AddedExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringNotification) GetAddedExternalIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AddedExternalIds) {
		return nil, false
	}
	return o.AddedExternalIds, true
}

// HasAddedExternalIds returns a boolean if a field has been set.
func (o *MonitoringNotification) HasAddedExternalIds() bool {
	if o != nil && !IsNil(o.AddedExternalIds) {
		return true
	}

	return false
}

// SetAddedExternalIds gets a reference to the given []string and assigns it to the AddedExternalIds field.
func (o *MonitoringNotification) SetAddedExternalIds(v []string) {
	o.AddedExternalIds = v
}

// GetAddedMsisdns returns the AddedMsisdns field value if set, zero value otherwise.
func (o *MonitoringNotification) GetAddedMsisdns() []string {
	if o == nil || IsNil(o.AddedMsisdns) {
		var ret []string
		return ret
	}
	return o.AddedMsisdns
}

// GetAddedMsisdnsOk returns a tuple with the AddedMsisdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringNotification) GetAddedMsisdnsOk() ([]string, bool) {
	if o == nil || IsNil(o.AddedMsisdns) {
		return nil, false
	}
	return o.AddedMsisdns, true
}

// HasAddedMsisdns returns a boolean if a field has been set.
func (o *MonitoringNotification) HasAddedMsisdns() bool {
	if o != nil && !IsNil(o.AddedMsisdns) {
		return true
	}

	return false
}

// SetAddedMsisdns gets a reference to the given []string and assigns it to the AddedMsisdns field.
func (o *MonitoringNotification) SetAddedMsisdns(v []string) {
	o.AddedMsisdns = v
}

// GetCancelExternalIds returns the CancelExternalIds field value if set, zero value otherwise.
func (o *MonitoringNotification) GetCancelExternalIds() []string {
	if o == nil || IsNil(o.CancelExternalIds) {
		var ret []string
		return ret
	}
	return o.CancelExternalIds
}

// GetCancelExternalIdsOk returns a tuple with the CancelExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringNotification) GetCancelExternalIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CancelExternalIds) {
		return nil, false
	}
	return o.CancelExternalIds, true
}

// HasCancelExternalIds returns a boolean if a field has been set.
func (o *MonitoringNotification) HasCancelExternalIds() bool {
	if o != nil && !IsNil(o.CancelExternalIds) {
		return true
	}

	return false
}

// SetCancelExternalIds gets a reference to the given []string and assigns it to the CancelExternalIds field.
func (o *MonitoringNotification) SetCancelExternalIds(v []string) {
	o.CancelExternalIds = v
}

// GetCancelMsisdns returns the CancelMsisdns field value if set, zero value otherwise.
func (o *MonitoringNotification) GetCancelMsisdns() []string {
	if o == nil || IsNil(o.CancelMsisdns) {
		var ret []string
		return ret
	}
	return o.CancelMsisdns
}

// GetCancelMsisdnsOk returns a tuple with the CancelMsisdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringNotification) GetCancelMsisdnsOk() ([]string, bool) {
	if o == nil || IsNil(o.CancelMsisdns) {
		return nil, false
	}
	return o.CancelMsisdns, true
}

// HasCancelMsisdns returns a boolean if a field has been set.
func (o *MonitoringNotification) HasCancelMsisdns() bool {
	if o != nil && !IsNil(o.CancelMsisdns) {
		return true
	}

	return false
}

// SetCancelMsisdns gets a reference to the given []string and assigns it to the CancelMsisdns field.
func (o *MonitoringNotification) SetCancelMsisdns(v []string) {
	o.CancelMsisdns = v
}

// GetCancelInd returns the CancelInd field value if set, zero value otherwise.
func (o *MonitoringNotification) GetCancelInd() bool {
	if o == nil || IsNil(o.CancelInd) {
		var ret bool
		return ret
	}
	return *o.CancelInd
}

// GetCancelIndOk returns a tuple with the CancelInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringNotification) GetCancelIndOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelInd) {
		return nil, false
	}
	return o.CancelInd, true
}

// HasCancelInd returns a boolean if a field has been set.
func (o *MonitoringNotification) HasCancelInd() bool {
	if o != nil && !IsNil(o.CancelInd) {
		return true
	}

	return false
}

// SetCancelInd gets a reference to the given bool and assigns it to the CancelInd field.
func (o *MonitoringNotification) SetCancelInd(v bool) {
	o.CancelInd = &v
}

// GetAppliedParam returns the AppliedParam field value if set, zero value otherwise.
func (o *MonitoringNotification) GetAppliedParam() AppliedParameterConfiguration {
	if o == nil || IsNil(o.AppliedParam) {
		var ret AppliedParameterConfiguration
		return ret
	}
	return *o.AppliedParam
}

// GetAppliedParamOk returns a tuple with the AppliedParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringNotification) GetAppliedParamOk() (*AppliedParameterConfiguration, bool) {
	if o == nil || IsNil(o.AppliedParam) {
		return nil, false
	}
	return o.AppliedParam, true
}

// HasAppliedParam returns a boolean if a field has been set.
func (o *MonitoringNotification) HasAppliedParam() bool {
	if o != nil && !IsNil(o.AppliedParam) {
		return true
	}

	return false
}

// SetAppliedParam gets a reference to the given AppliedParameterConfiguration and assigns it to the AppliedParam field.
func (o *MonitoringNotification) SetAppliedParam(v AppliedParameterConfiguration) {
	o.AppliedParam = &v
}

func (o MonitoringNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription"] = o.Subscription
	if !IsNil(o.ConfigResults) {
		toSerialize["configResults"] = o.ConfigResults
	}
	if !IsNil(o.MonitoringEventReports) {
		toSerialize["monitoringEventReports"] = o.MonitoringEventReports
	}
	if !IsNil(o.AddedExternalIds) {
		toSerialize["addedExternalIds"] = o.AddedExternalIds
	}
	if !IsNil(o.AddedMsisdns) {
		toSerialize["addedMsisdns"] = o.AddedMsisdns
	}
	if !IsNil(o.CancelExternalIds) {
		toSerialize["cancelExternalIds"] = o.CancelExternalIds
	}
	if !IsNil(o.CancelMsisdns) {
		toSerialize["cancelMsisdns"] = o.CancelMsisdns
	}
	if !IsNil(o.CancelInd) {
		toSerialize["cancelInd"] = o.CancelInd
	}
	if !IsNil(o.AppliedParam) {
		toSerialize["appliedParam"] = o.AppliedParam
	}
	return toSerialize, nil
}

type NullableMonitoringNotification struct {
	value *MonitoringNotification
	isSet bool
}

func (v NullableMonitoringNotification) Get() *MonitoringNotification {
	return v.value
}

func (v *NullableMonitoringNotification) Set(val *MonitoringNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringNotification(val *MonitoringNotification) *NullableMonitoringNotification {
	return &NullableMonitoringNotification{value: val, isSet: true}
}

func (v NullableMonitoringNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

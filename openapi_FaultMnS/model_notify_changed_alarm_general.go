/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
	"time"
)

// checks if the NotifyChangedAlarmGeneral type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyChangedAlarmGeneral{}

// NotifyChangedAlarmGeneral struct for NotifyChangedAlarmGeneral
type NotifyChangedAlarmGeneral struct {
	Href string `json:"href"`
	NotificationId int32 `json:"notificationId"`
	NotificationType NotificationType `json:"notificationType"`
	EventTime time.Time `json:"eventTime"`
	SystemDN string `json:"systemDN"`
	AlarmId string `json:"alarmId"`
	AlarmType AlarmType `json:"alarmType"`
	ProbableCause *ProbableCause `json:"probableCause,omitempty"`
	SpecificProblem *SpecificProblem `json:"specificProblem,omitempty"`
	PerceivedSeverity *PerceivedSeverity `json:"perceivedSeverity,omitempty"`
	CorrelatedNotifications []CorrelatedNotification `json:"correlatedNotifications,omitempty"`
	BackedUpStatus *bool `json:"backedUpStatus,omitempty"`
	BackUpObject *string `json:"backUpObject,omitempty"`
	TrendIndication *TrendIndication `json:"trendIndication,omitempty"`
	ThresholdInfo *ThresholdInfo `json:"thresholdInfo,omitempty"`
	// The first array item contains the attribute name value pairs with the new values, and the second array item the attribute name value pairs with the optional old values.
	StateChangeDefinition []map[string]interface{} `json:"stateChangeDefinition,omitempty"`
	// The key of this map is the attribute name, and the value the attribute value.
	MonitoredAttributes map[string]interface{} `json:"monitoredAttributes,omitempty"`
	ProposedRepairActions *string `json:"proposedRepairActions,omitempty"`
	AdditionalText *string `json:"additionalText,omitempty"`
	// The key of this map is the attribute name, and the value the attribute value.
	AdditionalInformation map[string]interface{} `json:"additionalInformation,omitempty"`
	RootCauseIndicator *bool `json:"rootCauseIndicator,omitempty"`
	// The key of this map is the attribute name, and the value the attribute value.
	ChangedAlarmAttributes map[string]interface{} `json:"changedAlarmAttributes,omitempty"`
}

// NewNotifyChangedAlarmGeneral instantiates a new NotifyChangedAlarmGeneral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyChangedAlarmGeneral(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, alarmId string, alarmType AlarmType) *NotifyChangedAlarmGeneral {
	this := NotifyChangedAlarmGeneral{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	this.AlarmId = alarmId
	this.AlarmType = alarmType
	return &this
}

// NewNotifyChangedAlarmGeneralWithDefaults instantiates a new NotifyChangedAlarmGeneral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyChangedAlarmGeneralWithDefaults() *NotifyChangedAlarmGeneral {
	this := NotifyChangedAlarmGeneral{}
	return &this
}

// GetHref returns the Href field value
func (o *NotifyChangedAlarmGeneral) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *NotifyChangedAlarmGeneral) SetHref(v string) {
	o.Href = v
}

// GetNotificationId returns the NotificationId field value
func (o *NotifyChangedAlarmGeneral) GetNotificationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetNotificationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationId, true
}

// SetNotificationId sets field value
func (o *NotifyChangedAlarmGeneral) SetNotificationId(v int32) {
	o.NotificationId = v
}

// GetNotificationType returns the NotificationType field value
func (o *NotifyChangedAlarmGeneral) GetNotificationType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetNotificationTypeOk() (*NotificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *NotifyChangedAlarmGeneral) SetNotificationType(v NotificationType) {
	o.NotificationType = v
}

// GetEventTime returns the EventTime field value
func (o *NotifyChangedAlarmGeneral) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *NotifyChangedAlarmGeneral) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetSystemDN returns the SystemDN field value
func (o *NotifyChangedAlarmGeneral) GetSystemDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemDN
}

// GetSystemDNOk returns a tuple with the SystemDN field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetSystemDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemDN, true
}

// SetSystemDN sets field value
func (o *NotifyChangedAlarmGeneral) SetSystemDN(v string) {
	o.SystemDN = v
}

// GetAlarmId returns the AlarmId field value
func (o *NotifyChangedAlarmGeneral) GetAlarmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlarmId
}

// GetAlarmIdOk returns a tuple with the AlarmId field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetAlarmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmId, true
}

// SetAlarmId sets field value
func (o *NotifyChangedAlarmGeneral) SetAlarmId(v string) {
	o.AlarmId = v
}

// GetAlarmType returns the AlarmType field value
func (o *NotifyChangedAlarmGeneral) GetAlarmType() AlarmType {
	if o == nil {
		var ret AlarmType
		return ret
	}

	return o.AlarmType
}

// GetAlarmTypeOk returns a tuple with the AlarmType field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetAlarmTypeOk() (*AlarmType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmType, true
}

// SetAlarmType sets field value
func (o *NotifyChangedAlarmGeneral) SetAlarmType(v AlarmType) {
	o.AlarmType = v
}

// GetProbableCause returns the ProbableCause field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetProbableCause() ProbableCause {
	if o == nil || IsNil(o.ProbableCause) {
		var ret ProbableCause
		return ret
	}
	return *o.ProbableCause
}

// GetProbableCauseOk returns a tuple with the ProbableCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetProbableCauseOk() (*ProbableCause, bool) {
	if o == nil || IsNil(o.ProbableCause) {
		return nil, false
	}
	return o.ProbableCause, true
}

// HasProbableCause returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasProbableCause() bool {
	if o != nil && !IsNil(o.ProbableCause) {
		return true
	}

	return false
}

// SetProbableCause gets a reference to the given ProbableCause and assigns it to the ProbableCause field.
func (o *NotifyChangedAlarmGeneral) SetProbableCause(v ProbableCause) {
	o.ProbableCause = &v
}

// GetSpecificProblem returns the SpecificProblem field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetSpecificProblem() SpecificProblem {
	if o == nil || IsNil(o.SpecificProblem) {
		var ret SpecificProblem
		return ret
	}
	return *o.SpecificProblem
}

// GetSpecificProblemOk returns a tuple with the SpecificProblem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetSpecificProblemOk() (*SpecificProblem, bool) {
	if o == nil || IsNil(o.SpecificProblem) {
		return nil, false
	}
	return o.SpecificProblem, true
}

// HasSpecificProblem returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasSpecificProblem() bool {
	if o != nil && !IsNil(o.SpecificProblem) {
		return true
	}

	return false
}

// SetSpecificProblem gets a reference to the given SpecificProblem and assigns it to the SpecificProblem field.
func (o *NotifyChangedAlarmGeneral) SetSpecificProblem(v SpecificProblem) {
	o.SpecificProblem = &v
}

// GetPerceivedSeverity returns the PerceivedSeverity field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetPerceivedSeverity() PerceivedSeverity {
	if o == nil || IsNil(o.PerceivedSeverity) {
		var ret PerceivedSeverity
		return ret
	}
	return *o.PerceivedSeverity
}

// GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetPerceivedSeverityOk() (*PerceivedSeverity, bool) {
	if o == nil || IsNil(o.PerceivedSeverity) {
		return nil, false
	}
	return o.PerceivedSeverity, true
}

// HasPerceivedSeverity returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasPerceivedSeverity() bool {
	if o != nil && !IsNil(o.PerceivedSeverity) {
		return true
	}

	return false
}

// SetPerceivedSeverity gets a reference to the given PerceivedSeverity and assigns it to the PerceivedSeverity field.
func (o *NotifyChangedAlarmGeneral) SetPerceivedSeverity(v PerceivedSeverity) {
	o.PerceivedSeverity = &v
}

// GetCorrelatedNotifications returns the CorrelatedNotifications field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetCorrelatedNotifications() []CorrelatedNotification {
	if o == nil || IsNil(o.CorrelatedNotifications) {
		var ret []CorrelatedNotification
		return ret
	}
	return o.CorrelatedNotifications
}

// GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetCorrelatedNotificationsOk() ([]CorrelatedNotification, bool) {
	if o == nil || IsNil(o.CorrelatedNotifications) {
		return nil, false
	}
	return o.CorrelatedNotifications, true
}

// HasCorrelatedNotifications returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasCorrelatedNotifications() bool {
	if o != nil && !IsNil(o.CorrelatedNotifications) {
		return true
	}

	return false
}

// SetCorrelatedNotifications gets a reference to the given []CorrelatedNotification and assigns it to the CorrelatedNotifications field.
func (o *NotifyChangedAlarmGeneral) SetCorrelatedNotifications(v []CorrelatedNotification) {
	o.CorrelatedNotifications = v
}

// GetBackedUpStatus returns the BackedUpStatus field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetBackedUpStatus() bool {
	if o == nil || IsNil(o.BackedUpStatus) {
		var ret bool
		return ret
	}
	return *o.BackedUpStatus
}

// GetBackedUpStatusOk returns a tuple with the BackedUpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetBackedUpStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.BackedUpStatus) {
		return nil, false
	}
	return o.BackedUpStatus, true
}

// HasBackedUpStatus returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasBackedUpStatus() bool {
	if o != nil && !IsNil(o.BackedUpStatus) {
		return true
	}

	return false
}

// SetBackedUpStatus gets a reference to the given bool and assigns it to the BackedUpStatus field.
func (o *NotifyChangedAlarmGeneral) SetBackedUpStatus(v bool) {
	o.BackedUpStatus = &v
}

// GetBackUpObject returns the BackUpObject field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetBackUpObject() string {
	if o == nil || IsNil(o.BackUpObject) {
		var ret string
		return ret
	}
	return *o.BackUpObject
}

// GetBackUpObjectOk returns a tuple with the BackUpObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetBackUpObjectOk() (*string, bool) {
	if o == nil || IsNil(o.BackUpObject) {
		return nil, false
	}
	return o.BackUpObject, true
}

// HasBackUpObject returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasBackUpObject() bool {
	if o != nil && !IsNil(o.BackUpObject) {
		return true
	}

	return false
}

// SetBackUpObject gets a reference to the given string and assigns it to the BackUpObject field.
func (o *NotifyChangedAlarmGeneral) SetBackUpObject(v string) {
	o.BackUpObject = &v
}

// GetTrendIndication returns the TrendIndication field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetTrendIndication() TrendIndication {
	if o == nil || IsNil(o.TrendIndication) {
		var ret TrendIndication
		return ret
	}
	return *o.TrendIndication
}

// GetTrendIndicationOk returns a tuple with the TrendIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetTrendIndicationOk() (*TrendIndication, bool) {
	if o == nil || IsNil(o.TrendIndication) {
		return nil, false
	}
	return o.TrendIndication, true
}

// HasTrendIndication returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasTrendIndication() bool {
	if o != nil && !IsNil(o.TrendIndication) {
		return true
	}

	return false
}

// SetTrendIndication gets a reference to the given TrendIndication and assigns it to the TrendIndication field.
func (o *NotifyChangedAlarmGeneral) SetTrendIndication(v TrendIndication) {
	o.TrendIndication = &v
}

// GetThresholdInfo returns the ThresholdInfo field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetThresholdInfo() ThresholdInfo {
	if o == nil || IsNil(o.ThresholdInfo) {
		var ret ThresholdInfo
		return ret
	}
	return *o.ThresholdInfo
}

// GetThresholdInfoOk returns a tuple with the ThresholdInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetThresholdInfoOk() (*ThresholdInfo, bool) {
	if o == nil || IsNil(o.ThresholdInfo) {
		return nil, false
	}
	return o.ThresholdInfo, true
}

// HasThresholdInfo returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasThresholdInfo() bool {
	if o != nil && !IsNil(o.ThresholdInfo) {
		return true
	}

	return false
}

// SetThresholdInfo gets a reference to the given ThresholdInfo and assigns it to the ThresholdInfo field.
func (o *NotifyChangedAlarmGeneral) SetThresholdInfo(v ThresholdInfo) {
	o.ThresholdInfo = &v
}

// GetStateChangeDefinition returns the StateChangeDefinition field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetStateChangeDefinition() []map[string]interface{} {
	if o == nil || IsNil(o.StateChangeDefinition) {
		var ret []map[string]interface{}
		return ret
	}
	return o.StateChangeDefinition
}

// GetStateChangeDefinitionOk returns a tuple with the StateChangeDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetStateChangeDefinitionOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.StateChangeDefinition) {
		return nil, false
	}
	return o.StateChangeDefinition, true
}

// HasStateChangeDefinition returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasStateChangeDefinition() bool {
	if o != nil && !IsNil(o.StateChangeDefinition) {
		return true
	}

	return false
}

// SetStateChangeDefinition gets a reference to the given []map[string]interface{} and assigns it to the StateChangeDefinition field.
func (o *NotifyChangedAlarmGeneral) SetStateChangeDefinition(v []map[string]interface{}) {
	o.StateChangeDefinition = v
}

// GetMonitoredAttributes returns the MonitoredAttributes field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetMonitoredAttributes() map[string]interface{} {
	if o == nil || IsNil(o.MonitoredAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.MonitoredAttributes
}

// GetMonitoredAttributesOk returns a tuple with the MonitoredAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetMonitoredAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MonitoredAttributes) {
		return map[string]interface{}{}, false
	}
	return o.MonitoredAttributes, true
}

// HasMonitoredAttributes returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasMonitoredAttributes() bool {
	if o != nil && !IsNil(o.MonitoredAttributes) {
		return true
	}

	return false
}

// SetMonitoredAttributes gets a reference to the given map[string]interface{} and assigns it to the MonitoredAttributes field.
func (o *NotifyChangedAlarmGeneral) SetMonitoredAttributes(v map[string]interface{}) {
	o.MonitoredAttributes = v
}

// GetProposedRepairActions returns the ProposedRepairActions field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetProposedRepairActions() string {
	if o == nil || IsNil(o.ProposedRepairActions) {
		var ret string
		return ret
	}
	return *o.ProposedRepairActions
}

// GetProposedRepairActionsOk returns a tuple with the ProposedRepairActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetProposedRepairActionsOk() (*string, bool) {
	if o == nil || IsNil(o.ProposedRepairActions) {
		return nil, false
	}
	return o.ProposedRepairActions, true
}

// HasProposedRepairActions returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasProposedRepairActions() bool {
	if o != nil && !IsNil(o.ProposedRepairActions) {
		return true
	}

	return false
}

// SetProposedRepairActions gets a reference to the given string and assigns it to the ProposedRepairActions field.
func (o *NotifyChangedAlarmGeneral) SetProposedRepairActions(v string) {
	o.ProposedRepairActions = &v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetAdditionalText() string {
	if o == nil || IsNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetAdditionalTextOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalText) {
		return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasAdditionalText() bool {
	if o != nil && !IsNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *NotifyChangedAlarmGeneral) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetAdditionalInformation() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalInformation) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetAdditionalInformationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalInformation) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasAdditionalInformation() bool {
	if o != nil && !IsNil(o.AdditionalInformation) {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given map[string]interface{} and assigns it to the AdditionalInformation field.
func (o *NotifyChangedAlarmGeneral) SetAdditionalInformation(v map[string]interface{}) {
	o.AdditionalInformation = v
}

// GetRootCauseIndicator returns the RootCauseIndicator field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetRootCauseIndicator() bool {
	if o == nil || IsNil(o.RootCauseIndicator) {
		var ret bool
		return ret
	}
	return *o.RootCauseIndicator
}

// GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetRootCauseIndicatorOk() (*bool, bool) {
	if o == nil || IsNil(o.RootCauseIndicator) {
		return nil, false
	}
	return o.RootCauseIndicator, true
}

// HasRootCauseIndicator returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasRootCauseIndicator() bool {
	if o != nil && !IsNil(o.RootCauseIndicator) {
		return true
	}

	return false
}

// SetRootCauseIndicator gets a reference to the given bool and assigns it to the RootCauseIndicator field.
func (o *NotifyChangedAlarmGeneral) SetRootCauseIndicator(v bool) {
	o.RootCauseIndicator = &v
}

// GetChangedAlarmAttributes returns the ChangedAlarmAttributes field value if set, zero value otherwise.
func (o *NotifyChangedAlarmGeneral) GetChangedAlarmAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ChangedAlarmAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ChangedAlarmAttributes
}

// GetChangedAlarmAttributesOk returns a tuple with the ChangedAlarmAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarmGeneral) GetChangedAlarmAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ChangedAlarmAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ChangedAlarmAttributes, true
}

// HasChangedAlarmAttributes returns a boolean if a field has been set.
func (o *NotifyChangedAlarmGeneral) HasChangedAlarmAttributes() bool {
	if o != nil && !IsNil(o.ChangedAlarmAttributes) {
		return true
	}

	return false
}

// SetChangedAlarmAttributes gets a reference to the given map[string]interface{} and assigns it to the ChangedAlarmAttributes field.
func (o *NotifyChangedAlarmGeneral) SetChangedAlarmAttributes(v map[string]interface{}) {
	o.ChangedAlarmAttributes = v
}

func (o NotifyChangedAlarmGeneral) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyChangedAlarmGeneral) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["notificationId"] = o.NotificationId
	toSerialize["notificationType"] = o.NotificationType
	toSerialize["eventTime"] = o.EventTime
	toSerialize["systemDN"] = o.SystemDN
	toSerialize["alarmId"] = o.AlarmId
	toSerialize["alarmType"] = o.AlarmType
	if !IsNil(o.ProbableCause) {
		toSerialize["probableCause"] = o.ProbableCause
	}
	if !IsNil(o.SpecificProblem) {
		toSerialize["specificProblem"] = o.SpecificProblem
	}
	if !IsNil(o.PerceivedSeverity) {
		toSerialize["perceivedSeverity"] = o.PerceivedSeverity
	}
	if !IsNil(o.CorrelatedNotifications) {
		toSerialize["correlatedNotifications"] = o.CorrelatedNotifications
	}
	if !IsNil(o.BackedUpStatus) {
		toSerialize["backedUpStatus"] = o.BackedUpStatus
	}
	if !IsNil(o.BackUpObject) {
		toSerialize["backUpObject"] = o.BackUpObject
	}
	if !IsNil(o.TrendIndication) {
		toSerialize["trendIndication"] = o.TrendIndication
	}
	if !IsNil(o.ThresholdInfo) {
		toSerialize["thresholdInfo"] = o.ThresholdInfo
	}
	if !IsNil(o.StateChangeDefinition) {
		toSerialize["stateChangeDefinition"] = o.StateChangeDefinition
	}
	if !IsNil(o.MonitoredAttributes) {
		toSerialize["monitoredAttributes"] = o.MonitoredAttributes
	}
	if !IsNil(o.ProposedRepairActions) {
		toSerialize["proposedRepairActions"] = o.ProposedRepairActions
	}
	if !IsNil(o.AdditionalText) {
		toSerialize["additionalText"] = o.AdditionalText
	}
	if !IsNil(o.AdditionalInformation) {
		toSerialize["additionalInformation"] = o.AdditionalInformation
	}
	if !IsNil(o.RootCauseIndicator) {
		toSerialize["rootCauseIndicator"] = o.RootCauseIndicator
	}
	if !IsNil(o.ChangedAlarmAttributes) {
		toSerialize["changedAlarmAttributes"] = o.ChangedAlarmAttributes
	}
	return toSerialize, nil
}

type NullableNotifyChangedAlarmGeneral struct {
	value *NotifyChangedAlarmGeneral
	isSet bool
}

func (v NullableNotifyChangedAlarmGeneral) Get() *NotifyChangedAlarmGeneral {
	return v.value
}

func (v *NullableNotifyChangedAlarmGeneral) Set(val *NotifyChangedAlarmGeneral) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyChangedAlarmGeneral) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyChangedAlarmGeneral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyChangedAlarmGeneral(val *NotifyChangedAlarmGeneral) *NullableNotifyChangedAlarmGeneral {
	return &NullableNotifyChangedAlarmGeneral{value: val, isSet: true}
}

func (v NullableNotifyChangedAlarmGeneral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyChangedAlarmGeneral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



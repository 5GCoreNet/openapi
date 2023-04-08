/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"time"
)

// checks if the AlarmRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlarmRecord{}

// AlarmRecord The alarmId is not a property of an alarm record. It is used as key in the map of alarm records instead.
type AlarmRecord struct {
	ObjectInstance *string `json:"objectInstance,omitempty"`
	NotificationId *int32 `json:"notificationId,omitempty"`
	AlarmRaisedTime *time.Time `json:"alarmRaisedTime,omitempty"`
	AlarmChangedTime *time.Time `json:"alarmChangedTime,omitempty"`
	AlarmClearedTime *time.Time `json:"alarmClearedTime,omitempty"`
	AlarmType *AlarmType `json:"alarmType,omitempty"`
	ProbableCause *ProbableCause `json:"probableCause,omitempty"`
	SpecificProblem *SpecificProblem `json:"specificProblem,omitempty"`
	PerceivedSeverity *PerceivedSeverity `json:"perceivedSeverity,omitempty"`
	BackedUpStatus *bool `json:"backedUpStatus,omitempty"`
	BackUpObject *string `json:"backUpObject,omitempty"`
	TrendIndication *TrendIndication `json:"trendIndication,omitempty"`
	Thresholdinfo *ThresholdInfo2 `json:"thresholdinfo,omitempty"`
	CorrelatedNotifications []CorrelatedNotification1 `json:"correlatedNotifications,omitempty"`
	// The first array item contains the attribute name value pairs with the new values, and the second array item the attribute name value pairs with the optional old values.
	StateChangeDefinition []map[string]interface{} `json:"stateChangeDefinition,omitempty"`
	// The key of this map is the attribute name, and the value the attribute value.
	MonitoredAttributes map[string]interface{} `json:"monitoredAttributes,omitempty"`
	ProposedRepairActions *string `json:"proposedRepairActions,omitempty"`
	AdditionalText *string `json:"additionalText,omitempty"`
	// The key of this map is the attribute name, and the value the attribute value.
	AdditionalInformation map[string]interface{} `json:"additionalInformation,omitempty"`
	RootCauseIndicator *bool `json:"rootCauseIndicator,omitempty"`
	AckTime *time.Time `json:"ackTime,omitempty"`
	AckUserId *string `json:"ackUserId,omitempty"`
	AckSystemId *string `json:"ackSystemId,omitempty"`
	AckState *AckState `json:"ackState,omitempty"`
	ClearUserId *string `json:"clearUserId,omitempty"`
	ClearSystemId *string `json:"clearSystemId,omitempty"`
	ServiceUser *string `json:"serviceUser,omitempty"`
	ServiceProvider *string `json:"serviceProvider,omitempty"`
	SecurityAlarmDetector *string `json:"securityAlarmDetector,omitempty"`
}

// NewAlarmRecord instantiates a new AlarmRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmRecord() *AlarmRecord {
	this := AlarmRecord{}
	return &this
}

// NewAlarmRecordWithDefaults instantiates a new AlarmRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmRecordWithDefaults() *AlarmRecord {
	this := AlarmRecord{}
	return &this
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *AlarmRecord) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *AlarmRecord) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *AlarmRecord) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetNotificationId returns the NotificationId field value if set, zero value otherwise.
func (o *AlarmRecord) GetNotificationId() int32 {
	if o == nil || isNil(o.NotificationId) {
		var ret int32
		return ret
	}
	return *o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetNotificationIdOk() (*int32, bool) {
	if o == nil || isNil(o.NotificationId) {
		return nil, false
	}
	return o.NotificationId, true
}

// HasNotificationId returns a boolean if a field has been set.
func (o *AlarmRecord) HasNotificationId() bool {
	if o != nil && !isNil(o.NotificationId) {
		return true
	}

	return false
}

// SetNotificationId gets a reference to the given int32 and assigns it to the NotificationId field.
func (o *AlarmRecord) SetNotificationId(v int32) {
	o.NotificationId = &v
}

// GetAlarmRaisedTime returns the AlarmRaisedTime field value if set, zero value otherwise.
func (o *AlarmRecord) GetAlarmRaisedTime() time.Time {
	if o == nil || isNil(o.AlarmRaisedTime) {
		var ret time.Time
		return ret
	}
	return *o.AlarmRaisedTime
}

// GetAlarmRaisedTimeOk returns a tuple with the AlarmRaisedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetAlarmRaisedTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.AlarmRaisedTime) {
		return nil, false
	}
	return o.AlarmRaisedTime, true
}

// HasAlarmRaisedTime returns a boolean if a field has been set.
func (o *AlarmRecord) HasAlarmRaisedTime() bool {
	if o != nil && !isNil(o.AlarmRaisedTime) {
		return true
	}

	return false
}

// SetAlarmRaisedTime gets a reference to the given time.Time and assigns it to the AlarmRaisedTime field.
func (o *AlarmRecord) SetAlarmRaisedTime(v time.Time) {
	o.AlarmRaisedTime = &v
}

// GetAlarmChangedTime returns the AlarmChangedTime field value if set, zero value otherwise.
func (o *AlarmRecord) GetAlarmChangedTime() time.Time {
	if o == nil || isNil(o.AlarmChangedTime) {
		var ret time.Time
		return ret
	}
	return *o.AlarmChangedTime
}

// GetAlarmChangedTimeOk returns a tuple with the AlarmChangedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetAlarmChangedTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.AlarmChangedTime) {
		return nil, false
	}
	return o.AlarmChangedTime, true
}

// HasAlarmChangedTime returns a boolean if a field has been set.
func (o *AlarmRecord) HasAlarmChangedTime() bool {
	if o != nil && !isNil(o.AlarmChangedTime) {
		return true
	}

	return false
}

// SetAlarmChangedTime gets a reference to the given time.Time and assigns it to the AlarmChangedTime field.
func (o *AlarmRecord) SetAlarmChangedTime(v time.Time) {
	o.AlarmChangedTime = &v
}

// GetAlarmClearedTime returns the AlarmClearedTime field value if set, zero value otherwise.
func (o *AlarmRecord) GetAlarmClearedTime() time.Time {
	if o == nil || isNil(o.AlarmClearedTime) {
		var ret time.Time
		return ret
	}
	return *o.AlarmClearedTime
}

// GetAlarmClearedTimeOk returns a tuple with the AlarmClearedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetAlarmClearedTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.AlarmClearedTime) {
		return nil, false
	}
	return o.AlarmClearedTime, true
}

// HasAlarmClearedTime returns a boolean if a field has been set.
func (o *AlarmRecord) HasAlarmClearedTime() bool {
	if o != nil && !isNil(o.AlarmClearedTime) {
		return true
	}

	return false
}

// SetAlarmClearedTime gets a reference to the given time.Time and assigns it to the AlarmClearedTime field.
func (o *AlarmRecord) SetAlarmClearedTime(v time.Time) {
	o.AlarmClearedTime = &v
}

// GetAlarmType returns the AlarmType field value if set, zero value otherwise.
func (o *AlarmRecord) GetAlarmType() AlarmType {
	if o == nil || isNil(o.AlarmType) {
		var ret AlarmType
		return ret
	}
	return *o.AlarmType
}

// GetAlarmTypeOk returns a tuple with the AlarmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetAlarmTypeOk() (*AlarmType, bool) {
	if o == nil || isNil(o.AlarmType) {
		return nil, false
	}
	return o.AlarmType, true
}

// HasAlarmType returns a boolean if a field has been set.
func (o *AlarmRecord) HasAlarmType() bool {
	if o != nil && !isNil(o.AlarmType) {
		return true
	}

	return false
}

// SetAlarmType gets a reference to the given AlarmType and assigns it to the AlarmType field.
func (o *AlarmRecord) SetAlarmType(v AlarmType) {
	o.AlarmType = &v
}

// GetProbableCause returns the ProbableCause field value if set, zero value otherwise.
func (o *AlarmRecord) GetProbableCause() ProbableCause {
	if o == nil || isNil(o.ProbableCause) {
		var ret ProbableCause
		return ret
	}
	return *o.ProbableCause
}

// GetProbableCauseOk returns a tuple with the ProbableCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetProbableCauseOk() (*ProbableCause, bool) {
	if o == nil || isNil(o.ProbableCause) {
		return nil, false
	}
	return o.ProbableCause, true
}

// HasProbableCause returns a boolean if a field has been set.
func (o *AlarmRecord) HasProbableCause() bool {
	if o != nil && !isNil(o.ProbableCause) {
		return true
	}

	return false
}

// SetProbableCause gets a reference to the given ProbableCause and assigns it to the ProbableCause field.
func (o *AlarmRecord) SetProbableCause(v ProbableCause) {
	o.ProbableCause = &v
}

// GetSpecificProblem returns the SpecificProblem field value if set, zero value otherwise.
func (o *AlarmRecord) GetSpecificProblem() SpecificProblem {
	if o == nil || isNil(o.SpecificProblem) {
		var ret SpecificProblem
		return ret
	}
	return *o.SpecificProblem
}

// GetSpecificProblemOk returns a tuple with the SpecificProblem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetSpecificProblemOk() (*SpecificProblem, bool) {
	if o == nil || isNil(o.SpecificProblem) {
		return nil, false
	}
	return o.SpecificProblem, true
}

// HasSpecificProblem returns a boolean if a field has been set.
func (o *AlarmRecord) HasSpecificProblem() bool {
	if o != nil && !isNil(o.SpecificProblem) {
		return true
	}

	return false
}

// SetSpecificProblem gets a reference to the given SpecificProblem and assigns it to the SpecificProblem field.
func (o *AlarmRecord) SetSpecificProblem(v SpecificProblem) {
	o.SpecificProblem = &v
}

// GetPerceivedSeverity returns the PerceivedSeverity field value if set, zero value otherwise.
func (o *AlarmRecord) GetPerceivedSeverity() PerceivedSeverity {
	if o == nil || isNil(o.PerceivedSeverity) {
		var ret PerceivedSeverity
		return ret
	}
	return *o.PerceivedSeverity
}

// GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetPerceivedSeverityOk() (*PerceivedSeverity, bool) {
	if o == nil || isNil(o.PerceivedSeverity) {
		return nil, false
	}
	return o.PerceivedSeverity, true
}

// HasPerceivedSeverity returns a boolean if a field has been set.
func (o *AlarmRecord) HasPerceivedSeverity() bool {
	if o != nil && !isNil(o.PerceivedSeverity) {
		return true
	}

	return false
}

// SetPerceivedSeverity gets a reference to the given PerceivedSeverity and assigns it to the PerceivedSeverity field.
func (o *AlarmRecord) SetPerceivedSeverity(v PerceivedSeverity) {
	o.PerceivedSeverity = &v
}

// GetBackedUpStatus returns the BackedUpStatus field value if set, zero value otherwise.
func (o *AlarmRecord) GetBackedUpStatus() bool {
	if o == nil || isNil(o.BackedUpStatus) {
		var ret bool
		return ret
	}
	return *o.BackedUpStatus
}

// GetBackedUpStatusOk returns a tuple with the BackedUpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetBackedUpStatusOk() (*bool, bool) {
	if o == nil || isNil(o.BackedUpStatus) {
		return nil, false
	}
	return o.BackedUpStatus, true
}

// HasBackedUpStatus returns a boolean if a field has been set.
func (o *AlarmRecord) HasBackedUpStatus() bool {
	if o != nil && !isNil(o.BackedUpStatus) {
		return true
	}

	return false
}

// SetBackedUpStatus gets a reference to the given bool and assigns it to the BackedUpStatus field.
func (o *AlarmRecord) SetBackedUpStatus(v bool) {
	o.BackedUpStatus = &v
}

// GetBackUpObject returns the BackUpObject field value if set, zero value otherwise.
func (o *AlarmRecord) GetBackUpObject() string {
	if o == nil || isNil(o.BackUpObject) {
		var ret string
		return ret
	}
	return *o.BackUpObject
}

// GetBackUpObjectOk returns a tuple with the BackUpObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetBackUpObjectOk() (*string, bool) {
	if o == nil || isNil(o.BackUpObject) {
		return nil, false
	}
	return o.BackUpObject, true
}

// HasBackUpObject returns a boolean if a field has been set.
func (o *AlarmRecord) HasBackUpObject() bool {
	if o != nil && !isNil(o.BackUpObject) {
		return true
	}

	return false
}

// SetBackUpObject gets a reference to the given string and assigns it to the BackUpObject field.
func (o *AlarmRecord) SetBackUpObject(v string) {
	o.BackUpObject = &v
}

// GetTrendIndication returns the TrendIndication field value if set, zero value otherwise.
func (o *AlarmRecord) GetTrendIndication() TrendIndication {
	if o == nil || isNil(o.TrendIndication) {
		var ret TrendIndication
		return ret
	}
	return *o.TrendIndication
}

// GetTrendIndicationOk returns a tuple with the TrendIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetTrendIndicationOk() (*TrendIndication, bool) {
	if o == nil || isNil(o.TrendIndication) {
		return nil, false
	}
	return o.TrendIndication, true
}

// HasTrendIndication returns a boolean if a field has been set.
func (o *AlarmRecord) HasTrendIndication() bool {
	if o != nil && !isNil(o.TrendIndication) {
		return true
	}

	return false
}

// SetTrendIndication gets a reference to the given TrendIndication and assigns it to the TrendIndication field.
func (o *AlarmRecord) SetTrendIndication(v TrendIndication) {
	o.TrendIndication = &v
}

// GetThresholdinfo returns the Thresholdinfo field value if set, zero value otherwise.
func (o *AlarmRecord) GetThresholdinfo() ThresholdInfo2 {
	if o == nil || isNil(o.Thresholdinfo) {
		var ret ThresholdInfo2
		return ret
	}
	return *o.Thresholdinfo
}

// GetThresholdinfoOk returns a tuple with the Thresholdinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetThresholdinfoOk() (*ThresholdInfo2, bool) {
	if o == nil || isNil(o.Thresholdinfo) {
		return nil, false
	}
	return o.Thresholdinfo, true
}

// HasThresholdinfo returns a boolean if a field has been set.
func (o *AlarmRecord) HasThresholdinfo() bool {
	if o != nil && !isNil(o.Thresholdinfo) {
		return true
	}

	return false
}

// SetThresholdinfo gets a reference to the given ThresholdInfo2 and assigns it to the Thresholdinfo field.
func (o *AlarmRecord) SetThresholdinfo(v ThresholdInfo2) {
	o.Thresholdinfo = &v
}

// GetCorrelatedNotifications returns the CorrelatedNotifications field value if set, zero value otherwise.
func (o *AlarmRecord) GetCorrelatedNotifications() []CorrelatedNotification1 {
	if o == nil || isNil(o.CorrelatedNotifications) {
		var ret []CorrelatedNotification1
		return ret
	}
	return o.CorrelatedNotifications
}

// GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetCorrelatedNotificationsOk() ([]CorrelatedNotification1, bool) {
	if o == nil || isNil(o.CorrelatedNotifications) {
		return nil, false
	}
	return o.CorrelatedNotifications, true
}

// HasCorrelatedNotifications returns a boolean if a field has been set.
func (o *AlarmRecord) HasCorrelatedNotifications() bool {
	if o != nil && !isNil(o.CorrelatedNotifications) {
		return true
	}

	return false
}

// SetCorrelatedNotifications gets a reference to the given []CorrelatedNotification1 and assigns it to the CorrelatedNotifications field.
func (o *AlarmRecord) SetCorrelatedNotifications(v []CorrelatedNotification1) {
	o.CorrelatedNotifications = v
}

// GetStateChangeDefinition returns the StateChangeDefinition field value if set, zero value otherwise.
func (o *AlarmRecord) GetStateChangeDefinition() []map[string]interface{} {
	if o == nil || isNil(o.StateChangeDefinition) {
		var ret []map[string]interface{}
		return ret
	}
	return o.StateChangeDefinition
}

// GetStateChangeDefinitionOk returns a tuple with the StateChangeDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetStateChangeDefinitionOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.StateChangeDefinition) {
		return nil, false
	}
	return o.StateChangeDefinition, true
}

// HasStateChangeDefinition returns a boolean if a field has been set.
func (o *AlarmRecord) HasStateChangeDefinition() bool {
	if o != nil && !isNil(o.StateChangeDefinition) {
		return true
	}

	return false
}

// SetStateChangeDefinition gets a reference to the given []map[string]interface{} and assigns it to the StateChangeDefinition field.
func (o *AlarmRecord) SetStateChangeDefinition(v []map[string]interface{}) {
	o.StateChangeDefinition = v
}

// GetMonitoredAttributes returns the MonitoredAttributes field value if set, zero value otherwise.
func (o *AlarmRecord) GetMonitoredAttributes() map[string]interface{} {
	if o == nil || isNil(o.MonitoredAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.MonitoredAttributes
}

// GetMonitoredAttributesOk returns a tuple with the MonitoredAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetMonitoredAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.MonitoredAttributes) {
		return map[string]interface{}{}, false
	}
	return o.MonitoredAttributes, true
}

// HasMonitoredAttributes returns a boolean if a field has been set.
func (o *AlarmRecord) HasMonitoredAttributes() bool {
	if o != nil && !isNil(o.MonitoredAttributes) {
		return true
	}

	return false
}

// SetMonitoredAttributes gets a reference to the given map[string]interface{} and assigns it to the MonitoredAttributes field.
func (o *AlarmRecord) SetMonitoredAttributes(v map[string]interface{}) {
	o.MonitoredAttributes = v
}

// GetProposedRepairActions returns the ProposedRepairActions field value if set, zero value otherwise.
func (o *AlarmRecord) GetProposedRepairActions() string {
	if o == nil || isNil(o.ProposedRepairActions) {
		var ret string
		return ret
	}
	return *o.ProposedRepairActions
}

// GetProposedRepairActionsOk returns a tuple with the ProposedRepairActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetProposedRepairActionsOk() (*string, bool) {
	if o == nil || isNil(o.ProposedRepairActions) {
		return nil, false
	}
	return o.ProposedRepairActions, true
}

// HasProposedRepairActions returns a boolean if a field has been set.
func (o *AlarmRecord) HasProposedRepairActions() bool {
	if o != nil && !isNil(o.ProposedRepairActions) {
		return true
	}

	return false
}

// SetProposedRepairActions gets a reference to the given string and assigns it to the ProposedRepairActions field.
func (o *AlarmRecord) SetProposedRepairActions(v string) {
	o.ProposedRepairActions = &v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *AlarmRecord) GetAdditionalText() string {
	if o == nil || isNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetAdditionalTextOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalText) {
		return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *AlarmRecord) HasAdditionalText() bool {
	if o != nil && !isNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *AlarmRecord) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *AlarmRecord) GetAdditionalInformation() map[string]interface{} {
	if o == nil || isNil(o.AdditionalInformation) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetAdditionalInformationOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.AdditionalInformation) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *AlarmRecord) HasAdditionalInformation() bool {
	if o != nil && !isNil(o.AdditionalInformation) {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given map[string]interface{} and assigns it to the AdditionalInformation field.
func (o *AlarmRecord) SetAdditionalInformation(v map[string]interface{}) {
	o.AdditionalInformation = v
}

// GetRootCauseIndicator returns the RootCauseIndicator field value if set, zero value otherwise.
func (o *AlarmRecord) GetRootCauseIndicator() bool {
	if o == nil || isNil(o.RootCauseIndicator) {
		var ret bool
		return ret
	}
	return *o.RootCauseIndicator
}

// GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetRootCauseIndicatorOk() (*bool, bool) {
	if o == nil || isNil(o.RootCauseIndicator) {
		return nil, false
	}
	return o.RootCauseIndicator, true
}

// HasRootCauseIndicator returns a boolean if a field has been set.
func (o *AlarmRecord) HasRootCauseIndicator() bool {
	if o != nil && !isNil(o.RootCauseIndicator) {
		return true
	}

	return false
}

// SetRootCauseIndicator gets a reference to the given bool and assigns it to the RootCauseIndicator field.
func (o *AlarmRecord) SetRootCauseIndicator(v bool) {
	o.RootCauseIndicator = &v
}

// GetAckTime returns the AckTime field value if set, zero value otherwise.
func (o *AlarmRecord) GetAckTime() time.Time {
	if o == nil || isNil(o.AckTime) {
		var ret time.Time
		return ret
	}
	return *o.AckTime
}

// GetAckTimeOk returns a tuple with the AckTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetAckTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.AckTime) {
		return nil, false
	}
	return o.AckTime, true
}

// HasAckTime returns a boolean if a field has been set.
func (o *AlarmRecord) HasAckTime() bool {
	if o != nil && !isNil(o.AckTime) {
		return true
	}

	return false
}

// SetAckTime gets a reference to the given time.Time and assigns it to the AckTime field.
func (o *AlarmRecord) SetAckTime(v time.Time) {
	o.AckTime = &v
}

// GetAckUserId returns the AckUserId field value if set, zero value otherwise.
func (o *AlarmRecord) GetAckUserId() string {
	if o == nil || isNil(o.AckUserId) {
		var ret string
		return ret
	}
	return *o.AckUserId
}

// GetAckUserIdOk returns a tuple with the AckUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetAckUserIdOk() (*string, bool) {
	if o == nil || isNil(o.AckUserId) {
		return nil, false
	}
	return o.AckUserId, true
}

// HasAckUserId returns a boolean if a field has been set.
func (o *AlarmRecord) HasAckUserId() bool {
	if o != nil && !isNil(o.AckUserId) {
		return true
	}

	return false
}

// SetAckUserId gets a reference to the given string and assigns it to the AckUserId field.
func (o *AlarmRecord) SetAckUserId(v string) {
	o.AckUserId = &v
}

// GetAckSystemId returns the AckSystemId field value if set, zero value otherwise.
func (o *AlarmRecord) GetAckSystemId() string {
	if o == nil || isNil(o.AckSystemId) {
		var ret string
		return ret
	}
	return *o.AckSystemId
}

// GetAckSystemIdOk returns a tuple with the AckSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetAckSystemIdOk() (*string, bool) {
	if o == nil || isNil(o.AckSystemId) {
		return nil, false
	}
	return o.AckSystemId, true
}

// HasAckSystemId returns a boolean if a field has been set.
func (o *AlarmRecord) HasAckSystemId() bool {
	if o != nil && !isNil(o.AckSystemId) {
		return true
	}

	return false
}

// SetAckSystemId gets a reference to the given string and assigns it to the AckSystemId field.
func (o *AlarmRecord) SetAckSystemId(v string) {
	o.AckSystemId = &v
}

// GetAckState returns the AckState field value if set, zero value otherwise.
func (o *AlarmRecord) GetAckState() AckState {
	if o == nil || isNil(o.AckState) {
		var ret AckState
		return ret
	}
	return *o.AckState
}

// GetAckStateOk returns a tuple with the AckState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetAckStateOk() (*AckState, bool) {
	if o == nil || isNil(o.AckState) {
		return nil, false
	}
	return o.AckState, true
}

// HasAckState returns a boolean if a field has been set.
func (o *AlarmRecord) HasAckState() bool {
	if o != nil && !isNil(o.AckState) {
		return true
	}

	return false
}

// SetAckState gets a reference to the given AckState and assigns it to the AckState field.
func (o *AlarmRecord) SetAckState(v AckState) {
	o.AckState = &v
}

// GetClearUserId returns the ClearUserId field value if set, zero value otherwise.
func (o *AlarmRecord) GetClearUserId() string {
	if o == nil || isNil(o.ClearUserId) {
		var ret string
		return ret
	}
	return *o.ClearUserId
}

// GetClearUserIdOk returns a tuple with the ClearUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetClearUserIdOk() (*string, bool) {
	if o == nil || isNil(o.ClearUserId) {
		return nil, false
	}
	return o.ClearUserId, true
}

// HasClearUserId returns a boolean if a field has been set.
func (o *AlarmRecord) HasClearUserId() bool {
	if o != nil && !isNil(o.ClearUserId) {
		return true
	}

	return false
}

// SetClearUserId gets a reference to the given string and assigns it to the ClearUserId field.
func (o *AlarmRecord) SetClearUserId(v string) {
	o.ClearUserId = &v
}

// GetClearSystemId returns the ClearSystemId field value if set, zero value otherwise.
func (o *AlarmRecord) GetClearSystemId() string {
	if o == nil || isNil(o.ClearSystemId) {
		var ret string
		return ret
	}
	return *o.ClearSystemId
}

// GetClearSystemIdOk returns a tuple with the ClearSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetClearSystemIdOk() (*string, bool) {
	if o == nil || isNil(o.ClearSystemId) {
		return nil, false
	}
	return o.ClearSystemId, true
}

// HasClearSystemId returns a boolean if a field has been set.
func (o *AlarmRecord) HasClearSystemId() bool {
	if o != nil && !isNil(o.ClearSystemId) {
		return true
	}

	return false
}

// SetClearSystemId gets a reference to the given string and assigns it to the ClearSystemId field.
func (o *AlarmRecord) SetClearSystemId(v string) {
	o.ClearSystemId = &v
}

// GetServiceUser returns the ServiceUser field value if set, zero value otherwise.
func (o *AlarmRecord) GetServiceUser() string {
	if o == nil || isNil(o.ServiceUser) {
		var ret string
		return ret
	}
	return *o.ServiceUser
}

// GetServiceUserOk returns a tuple with the ServiceUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetServiceUserOk() (*string, bool) {
	if o == nil || isNil(o.ServiceUser) {
		return nil, false
	}
	return o.ServiceUser, true
}

// HasServiceUser returns a boolean if a field has been set.
func (o *AlarmRecord) HasServiceUser() bool {
	if o != nil && !isNil(o.ServiceUser) {
		return true
	}

	return false
}

// SetServiceUser gets a reference to the given string and assigns it to the ServiceUser field.
func (o *AlarmRecord) SetServiceUser(v string) {
	o.ServiceUser = &v
}

// GetServiceProvider returns the ServiceProvider field value if set, zero value otherwise.
func (o *AlarmRecord) GetServiceProvider() string {
	if o == nil || isNil(o.ServiceProvider) {
		var ret string
		return ret
	}
	return *o.ServiceProvider
}

// GetServiceProviderOk returns a tuple with the ServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetServiceProviderOk() (*string, bool) {
	if o == nil || isNil(o.ServiceProvider) {
		return nil, false
	}
	return o.ServiceProvider, true
}

// HasServiceProvider returns a boolean if a field has been set.
func (o *AlarmRecord) HasServiceProvider() bool {
	if o != nil && !isNil(o.ServiceProvider) {
		return true
	}

	return false
}

// SetServiceProvider gets a reference to the given string and assigns it to the ServiceProvider field.
func (o *AlarmRecord) SetServiceProvider(v string) {
	o.ServiceProvider = &v
}

// GetSecurityAlarmDetector returns the SecurityAlarmDetector field value if set, zero value otherwise.
func (o *AlarmRecord) GetSecurityAlarmDetector() string {
	if o == nil || isNil(o.SecurityAlarmDetector) {
		var ret string
		return ret
	}
	return *o.SecurityAlarmDetector
}

// GetSecurityAlarmDetectorOk returns a tuple with the SecurityAlarmDetector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRecord) GetSecurityAlarmDetectorOk() (*string, bool) {
	if o == nil || isNil(o.SecurityAlarmDetector) {
		return nil, false
	}
	return o.SecurityAlarmDetector, true
}

// HasSecurityAlarmDetector returns a boolean if a field has been set.
func (o *AlarmRecord) HasSecurityAlarmDetector() bool {
	if o != nil && !isNil(o.SecurityAlarmDetector) {
		return true
	}

	return false
}

// SetSecurityAlarmDetector gets a reference to the given string and assigns it to the SecurityAlarmDetector field.
func (o *AlarmRecord) SetSecurityAlarmDetector(v string) {
	o.SecurityAlarmDetector = &v
}

func (o AlarmRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlarmRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ObjectInstance) {
		toSerialize["objectInstance"] = o.ObjectInstance
	}
	if !isNil(o.NotificationId) {
		toSerialize["notificationId"] = o.NotificationId
	}
	if !isNil(o.AlarmRaisedTime) {
		toSerialize["alarmRaisedTime"] = o.AlarmRaisedTime
	}
	if !isNil(o.AlarmChangedTime) {
		toSerialize["alarmChangedTime"] = o.AlarmChangedTime
	}
	if !isNil(o.AlarmClearedTime) {
		toSerialize["alarmClearedTime"] = o.AlarmClearedTime
	}
	if !isNil(o.AlarmType) {
		toSerialize["alarmType"] = o.AlarmType
	}
	if !isNil(o.ProbableCause) {
		toSerialize["probableCause"] = o.ProbableCause
	}
	if !isNil(o.SpecificProblem) {
		toSerialize["specificProblem"] = o.SpecificProblem
	}
	if !isNil(o.PerceivedSeverity) {
		toSerialize["perceivedSeverity"] = o.PerceivedSeverity
	}
	if !isNil(o.BackedUpStatus) {
		toSerialize["backedUpStatus"] = o.BackedUpStatus
	}
	if !isNil(o.BackUpObject) {
		toSerialize["backUpObject"] = o.BackUpObject
	}
	if !isNil(o.TrendIndication) {
		toSerialize["trendIndication"] = o.TrendIndication
	}
	if !isNil(o.Thresholdinfo) {
		toSerialize["thresholdinfo"] = o.Thresholdinfo
	}
	if !isNil(o.CorrelatedNotifications) {
		toSerialize["correlatedNotifications"] = o.CorrelatedNotifications
	}
	if !isNil(o.StateChangeDefinition) {
		toSerialize["stateChangeDefinition"] = o.StateChangeDefinition
	}
	if !isNil(o.MonitoredAttributes) {
		toSerialize["monitoredAttributes"] = o.MonitoredAttributes
	}
	if !isNil(o.ProposedRepairActions) {
		toSerialize["proposedRepairActions"] = o.ProposedRepairActions
	}
	if !isNil(o.AdditionalText) {
		toSerialize["additionalText"] = o.AdditionalText
	}
	if !isNil(o.AdditionalInformation) {
		toSerialize["additionalInformation"] = o.AdditionalInformation
	}
	if !isNil(o.RootCauseIndicator) {
		toSerialize["rootCauseIndicator"] = o.RootCauseIndicator
	}
	if !isNil(o.AckTime) {
		toSerialize["ackTime"] = o.AckTime
	}
	if !isNil(o.AckUserId) {
		toSerialize["ackUserId"] = o.AckUserId
	}
	if !isNil(o.AckSystemId) {
		toSerialize["ackSystemId"] = o.AckSystemId
	}
	if !isNil(o.AckState) {
		toSerialize["ackState"] = o.AckState
	}
	if !isNil(o.ClearUserId) {
		toSerialize["clearUserId"] = o.ClearUserId
	}
	if !isNil(o.ClearSystemId) {
		toSerialize["clearSystemId"] = o.ClearSystemId
	}
	if !isNil(o.ServiceUser) {
		toSerialize["serviceUser"] = o.ServiceUser
	}
	if !isNil(o.ServiceProvider) {
		toSerialize["serviceProvider"] = o.ServiceProvider
	}
	if !isNil(o.SecurityAlarmDetector) {
		toSerialize["securityAlarmDetector"] = o.SecurityAlarmDetector
	}
	return toSerialize, nil
}

type NullableAlarmRecord struct {
	value *AlarmRecord
	isSet bool
}

func (v NullableAlarmRecord) Get() *AlarmRecord {
	return v.value
}

func (v *NullableAlarmRecord) Set(val *AlarmRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmRecord(val *AlarmRecord) *NullableAlarmRecord {
	return &NullableAlarmRecord{value: val, isSet: true}
}

func (v NullableAlarmRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



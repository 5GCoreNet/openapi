/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
)

// checks if the NotifyNewSecAlarm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyNewSecAlarm{}

// NotifyNewSecAlarm struct for NotifyNewSecAlarm
type NotifyNewSecAlarm struct {
	Href string `json:"href"`
	NotificationId int32 `json:"notificationId"`
	NotificationType NotificationType `json:"notificationType"`
	EventTime time.Time `json:"eventTime"`
	SystemDN string `json:"systemDN"`
	AlarmId string `json:"alarmId"`
	AlarmType AlarmType `json:"alarmType"`
	ProbableCause ProbableCause `json:"probableCause"`
	PerceivedSeverity PerceivedSeverity `json:"perceivedSeverity"`
	CorrelatedNotifications []CorrelatedNotification `json:"correlatedNotifications,omitempty"`
	AdditionalText *string `json:"additionalText,omitempty"`
	// The key of this map is the attribute name, and the value the attribute value.
	AdditionalInformation map[string]interface{} `json:"additionalInformation,omitempty"`
	RootCauseIndicator *bool `json:"rootCauseIndicator,omitempty"`
	ServiceUser string `json:"serviceUser"`
	ServiceProvider string `json:"serviceProvider"`
	SecurityAlarmDetector string `json:"securityAlarmDetector"`
}

// NewNotifyNewSecAlarm instantiates a new NotifyNewSecAlarm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyNewSecAlarm(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, alarmId string, alarmType AlarmType, probableCause ProbableCause, perceivedSeverity PerceivedSeverity, serviceUser string, serviceProvider string, securityAlarmDetector string) *NotifyNewSecAlarm {
	this := NotifyNewSecAlarm{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	this.AlarmId = alarmId
	this.AlarmType = alarmType
	this.ProbableCause = probableCause
	this.PerceivedSeverity = perceivedSeverity
	this.ServiceUser = serviceUser
	this.ServiceProvider = serviceProvider
	this.SecurityAlarmDetector = securityAlarmDetector
	return &this
}

// NewNotifyNewSecAlarmWithDefaults instantiates a new NotifyNewSecAlarm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyNewSecAlarmWithDefaults() *NotifyNewSecAlarm {
	this := NotifyNewSecAlarm{}
	return &this
}

// GetHref returns the Href field value
func (o *NotifyNewSecAlarm) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *NotifyNewSecAlarm) SetHref(v string) {
	o.Href = v
}

// GetNotificationId returns the NotificationId field value
func (o *NotifyNewSecAlarm) GetNotificationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetNotificationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationId, true
}

// SetNotificationId sets field value
func (o *NotifyNewSecAlarm) SetNotificationId(v int32) {
	o.NotificationId = v
}

// GetNotificationType returns the NotificationType field value
func (o *NotifyNewSecAlarm) GetNotificationType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetNotificationTypeOk() (*NotificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *NotifyNewSecAlarm) SetNotificationType(v NotificationType) {
	o.NotificationType = v
}

// GetEventTime returns the EventTime field value
func (o *NotifyNewSecAlarm) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *NotifyNewSecAlarm) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetSystemDN returns the SystemDN field value
func (o *NotifyNewSecAlarm) GetSystemDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemDN
}

// GetSystemDNOk returns a tuple with the SystemDN field value
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetSystemDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemDN, true
}

// SetSystemDN sets field value
func (o *NotifyNewSecAlarm) SetSystemDN(v string) {
	o.SystemDN = v
}

// GetAlarmId returns the AlarmId field value
func (o *NotifyNewSecAlarm) GetAlarmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlarmId
}

// GetAlarmIdOk returns a tuple with the AlarmId field value
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetAlarmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmId, true
}

// SetAlarmId sets field value
func (o *NotifyNewSecAlarm) SetAlarmId(v string) {
	o.AlarmId = v
}

// GetAlarmType returns the AlarmType field value
func (o *NotifyNewSecAlarm) GetAlarmType() AlarmType {
	if o == nil {
		var ret AlarmType
		return ret
	}

	return o.AlarmType
}

// GetAlarmTypeOk returns a tuple with the AlarmType field value
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetAlarmTypeOk() (*AlarmType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmType, true
}

// SetAlarmType sets field value
func (o *NotifyNewSecAlarm) SetAlarmType(v AlarmType) {
	o.AlarmType = v
}

// GetProbableCause returns the ProbableCause field value
func (o *NotifyNewSecAlarm) GetProbableCause() ProbableCause {
	if o == nil {
		var ret ProbableCause
		return ret
	}

	return o.ProbableCause
}

// GetProbableCauseOk returns a tuple with the ProbableCause field value
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetProbableCauseOk() (*ProbableCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProbableCause, true
}

// SetProbableCause sets field value
func (o *NotifyNewSecAlarm) SetProbableCause(v ProbableCause) {
	o.ProbableCause = v
}

// GetPerceivedSeverity returns the PerceivedSeverity field value
func (o *NotifyNewSecAlarm) GetPerceivedSeverity() PerceivedSeverity {
	if o == nil {
		var ret PerceivedSeverity
		return ret
	}

	return o.PerceivedSeverity
}

// GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field value
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetPerceivedSeverityOk() (*PerceivedSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerceivedSeverity, true
}

// SetPerceivedSeverity sets field value
func (o *NotifyNewSecAlarm) SetPerceivedSeverity(v PerceivedSeverity) {
	o.PerceivedSeverity = v
}

// GetCorrelatedNotifications returns the CorrelatedNotifications field value if set, zero value otherwise.
func (o *NotifyNewSecAlarm) GetCorrelatedNotifications() []CorrelatedNotification {
	if o == nil || isNil(o.CorrelatedNotifications) {
		var ret []CorrelatedNotification
		return ret
	}
	return o.CorrelatedNotifications
}

// GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetCorrelatedNotificationsOk() ([]CorrelatedNotification, bool) {
	if o == nil || isNil(o.CorrelatedNotifications) {
		return nil, false
	}
	return o.CorrelatedNotifications, true
}

// HasCorrelatedNotifications returns a boolean if a field has been set.
func (o *NotifyNewSecAlarm) HasCorrelatedNotifications() bool {
	if o != nil && !isNil(o.CorrelatedNotifications) {
		return true
	}

	return false
}

// SetCorrelatedNotifications gets a reference to the given []CorrelatedNotification and assigns it to the CorrelatedNotifications field.
func (o *NotifyNewSecAlarm) SetCorrelatedNotifications(v []CorrelatedNotification) {
	o.CorrelatedNotifications = v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *NotifyNewSecAlarm) GetAdditionalText() string {
	if o == nil || isNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetAdditionalTextOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalText) {
		return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *NotifyNewSecAlarm) HasAdditionalText() bool {
	if o != nil && !isNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *NotifyNewSecAlarm) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *NotifyNewSecAlarm) GetAdditionalInformation() map[string]interface{} {
	if o == nil || isNil(o.AdditionalInformation) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetAdditionalInformationOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.AdditionalInformation) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *NotifyNewSecAlarm) HasAdditionalInformation() bool {
	if o != nil && !isNil(o.AdditionalInformation) {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given map[string]interface{} and assigns it to the AdditionalInformation field.
func (o *NotifyNewSecAlarm) SetAdditionalInformation(v map[string]interface{}) {
	o.AdditionalInformation = v
}

// GetRootCauseIndicator returns the RootCauseIndicator field value if set, zero value otherwise.
func (o *NotifyNewSecAlarm) GetRootCauseIndicator() bool {
	if o == nil || isNil(o.RootCauseIndicator) {
		var ret bool
		return ret
	}
	return *o.RootCauseIndicator
}

// GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetRootCauseIndicatorOk() (*bool, bool) {
	if o == nil || isNil(o.RootCauseIndicator) {
		return nil, false
	}
	return o.RootCauseIndicator, true
}

// HasRootCauseIndicator returns a boolean if a field has been set.
func (o *NotifyNewSecAlarm) HasRootCauseIndicator() bool {
	if o != nil && !isNil(o.RootCauseIndicator) {
		return true
	}

	return false
}

// SetRootCauseIndicator gets a reference to the given bool and assigns it to the RootCauseIndicator field.
func (o *NotifyNewSecAlarm) SetRootCauseIndicator(v bool) {
	o.RootCauseIndicator = &v
}

// GetServiceUser returns the ServiceUser field value
func (o *NotifyNewSecAlarm) GetServiceUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceUser
}

// GetServiceUserOk returns a tuple with the ServiceUser field value
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetServiceUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUser, true
}

// SetServiceUser sets field value
func (o *NotifyNewSecAlarm) SetServiceUser(v string) {
	o.ServiceUser = v
}

// GetServiceProvider returns the ServiceProvider field value
func (o *NotifyNewSecAlarm) GetServiceProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceProvider
}

// GetServiceProviderOk returns a tuple with the ServiceProvider field value
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetServiceProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceProvider, true
}

// SetServiceProvider sets field value
func (o *NotifyNewSecAlarm) SetServiceProvider(v string) {
	o.ServiceProvider = v
}

// GetSecurityAlarmDetector returns the SecurityAlarmDetector field value
func (o *NotifyNewSecAlarm) GetSecurityAlarmDetector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityAlarmDetector
}

// GetSecurityAlarmDetectorOk returns a tuple with the SecurityAlarmDetector field value
// and a boolean to check if the value has been set.
func (o *NotifyNewSecAlarm) GetSecurityAlarmDetectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityAlarmDetector, true
}

// SetSecurityAlarmDetector sets field value
func (o *NotifyNewSecAlarm) SetSecurityAlarmDetector(v string) {
	o.SecurityAlarmDetector = v
}

func (o NotifyNewSecAlarm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyNewSecAlarm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["notificationId"] = o.NotificationId
	toSerialize["notificationType"] = o.NotificationType
	toSerialize["eventTime"] = o.EventTime
	toSerialize["systemDN"] = o.SystemDN
	toSerialize["alarmId"] = o.AlarmId
	toSerialize["alarmType"] = o.AlarmType
	toSerialize["probableCause"] = o.ProbableCause
	toSerialize["perceivedSeverity"] = o.PerceivedSeverity
	if !isNil(o.CorrelatedNotifications) {
		toSerialize["correlatedNotifications"] = o.CorrelatedNotifications
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
	toSerialize["serviceUser"] = o.ServiceUser
	toSerialize["serviceProvider"] = o.ServiceProvider
	toSerialize["securityAlarmDetector"] = o.SecurityAlarmDetector
	return toSerialize, nil
}

type NullableNotifyNewSecAlarm struct {
	value *NotifyNewSecAlarm
	isSet bool
}

func (v NullableNotifyNewSecAlarm) Get() *NotifyNewSecAlarm {
	return v.value
}

func (v *NullableNotifyNewSecAlarm) Set(val *NotifyNewSecAlarm) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyNewSecAlarm) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyNewSecAlarm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyNewSecAlarm(val *NotifyNewSecAlarm) *NullableNotifyNewSecAlarm {
	return &NullableNotifyNewSecAlarm{value: val, isSet: true}
}

func (v NullableNotifyNewSecAlarm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyNewSecAlarm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



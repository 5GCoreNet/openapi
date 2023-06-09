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

// checks if the NotifyChangedSecAlarmGeneral type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyChangedSecAlarmGeneral{}

// NotifyChangedSecAlarmGeneral struct for NotifyChangedSecAlarmGeneral
type NotifyChangedSecAlarmGeneral struct {
	NotificationHeader
	AlarmId                 string                   `json:"alarmId"`
	AlarmType               AlarmType                `json:"alarmType"`
	ProbableCause           *ProbableCause           `json:"probableCause,omitempty"`
	PerceivedSeverity       *PerceivedSeverity       `json:"perceivedSeverity,omitempty"`
	CorrelatedNotifications []CorrelatedNotification `json:"correlatedNotifications,omitempty"`
	AdditionalText          *string                  `json:"additionalText,omitempty"`
	// The key of this map is the attribute name, and the value the attribute value.
	AdditionalInformation map[string]interface{} `json:"additionalInformation,omitempty"`
	RootCauseIndicator    *bool                  `json:"rootCauseIndicator,omitempty"`
	ServiceUser           string                 `json:"serviceUser"`
	ServiceProvider       string                 `json:"serviceProvider"`
	SecurityAlarmDetector string                 `json:"securityAlarmDetector"`
	// The key of this map is the attribute name, and the value the attribute value.
	ChangedAlarmAttributes map[string]interface{} `json:"changedAlarmAttributes,omitempty"`
}

// NewNotifyChangedSecAlarmGeneral instantiates a new NotifyChangedSecAlarmGeneral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyChangedSecAlarmGeneral(alarmId string, alarmType AlarmType, serviceUser string, serviceProvider string, securityAlarmDetector string, href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string) *NotifyChangedSecAlarmGeneral {
	this := NotifyChangedSecAlarmGeneral{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	this.AlarmId = alarmId
	this.AlarmType = alarmType
	this.ServiceUser = serviceUser
	this.ServiceProvider = serviceProvider
	this.SecurityAlarmDetector = securityAlarmDetector
	return &this
}

// NewNotifyChangedSecAlarmGeneralWithDefaults instantiates a new NotifyChangedSecAlarmGeneral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyChangedSecAlarmGeneralWithDefaults() *NotifyChangedSecAlarmGeneral {
	this := NotifyChangedSecAlarmGeneral{}
	return &this
}

// GetAlarmId returns the AlarmId field value
func (o *NotifyChangedSecAlarmGeneral) GetAlarmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlarmId
}

// GetAlarmIdOk returns a tuple with the AlarmId field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedSecAlarmGeneral) GetAlarmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmId, true
}

// SetAlarmId sets field value
func (o *NotifyChangedSecAlarmGeneral) SetAlarmId(v string) {
	o.AlarmId = v
}

// GetAlarmType returns the AlarmType field value
func (o *NotifyChangedSecAlarmGeneral) GetAlarmType() AlarmType {
	if o == nil {
		var ret AlarmType
		return ret
	}

	return o.AlarmType
}

// GetAlarmTypeOk returns a tuple with the AlarmType field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedSecAlarmGeneral) GetAlarmTypeOk() (*AlarmType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmType, true
}

// SetAlarmType sets field value
func (o *NotifyChangedSecAlarmGeneral) SetAlarmType(v AlarmType) {
	o.AlarmType = v
}

// GetProbableCause returns the ProbableCause field value if set, zero value otherwise.
func (o *NotifyChangedSecAlarmGeneral) GetProbableCause() ProbableCause {
	if o == nil || IsNil(o.ProbableCause) {
		var ret ProbableCause
		return ret
	}
	return *o.ProbableCause
}

// GetProbableCauseOk returns a tuple with the ProbableCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedSecAlarmGeneral) GetProbableCauseOk() (*ProbableCause, bool) {
	if o == nil || IsNil(o.ProbableCause) {
		return nil, false
	}
	return o.ProbableCause, true
}

// HasProbableCause returns a boolean if a field has been set.
func (o *NotifyChangedSecAlarmGeneral) HasProbableCause() bool {
	if o != nil && !IsNil(o.ProbableCause) {
		return true
	}

	return false
}

// SetProbableCause gets a reference to the given ProbableCause and assigns it to the ProbableCause field.
func (o *NotifyChangedSecAlarmGeneral) SetProbableCause(v ProbableCause) {
	o.ProbableCause = &v
}

// GetPerceivedSeverity returns the PerceivedSeverity field value if set, zero value otherwise.
func (o *NotifyChangedSecAlarmGeneral) GetPerceivedSeverity() PerceivedSeverity {
	if o == nil || IsNil(o.PerceivedSeverity) {
		var ret PerceivedSeverity
		return ret
	}
	return *o.PerceivedSeverity
}

// GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedSecAlarmGeneral) GetPerceivedSeverityOk() (*PerceivedSeverity, bool) {
	if o == nil || IsNil(o.PerceivedSeverity) {
		return nil, false
	}
	return o.PerceivedSeverity, true
}

// HasPerceivedSeverity returns a boolean if a field has been set.
func (o *NotifyChangedSecAlarmGeneral) HasPerceivedSeverity() bool {
	if o != nil && !IsNil(o.PerceivedSeverity) {
		return true
	}

	return false
}

// SetPerceivedSeverity gets a reference to the given PerceivedSeverity and assigns it to the PerceivedSeverity field.
func (o *NotifyChangedSecAlarmGeneral) SetPerceivedSeverity(v PerceivedSeverity) {
	o.PerceivedSeverity = &v
}

// GetCorrelatedNotifications returns the CorrelatedNotifications field value if set, zero value otherwise.
func (o *NotifyChangedSecAlarmGeneral) GetCorrelatedNotifications() []CorrelatedNotification {
	if o == nil || IsNil(o.CorrelatedNotifications) {
		var ret []CorrelatedNotification
		return ret
	}
	return o.CorrelatedNotifications
}

// GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedSecAlarmGeneral) GetCorrelatedNotificationsOk() ([]CorrelatedNotification, bool) {
	if o == nil || IsNil(o.CorrelatedNotifications) {
		return nil, false
	}
	return o.CorrelatedNotifications, true
}

// HasCorrelatedNotifications returns a boolean if a field has been set.
func (o *NotifyChangedSecAlarmGeneral) HasCorrelatedNotifications() bool {
	if o != nil && !IsNil(o.CorrelatedNotifications) {
		return true
	}

	return false
}

// SetCorrelatedNotifications gets a reference to the given []CorrelatedNotification and assigns it to the CorrelatedNotifications field.
func (o *NotifyChangedSecAlarmGeneral) SetCorrelatedNotifications(v []CorrelatedNotification) {
	o.CorrelatedNotifications = v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *NotifyChangedSecAlarmGeneral) GetAdditionalText() string {
	if o == nil || IsNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedSecAlarmGeneral) GetAdditionalTextOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalText) {
		return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *NotifyChangedSecAlarmGeneral) HasAdditionalText() bool {
	if o != nil && !IsNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *NotifyChangedSecAlarmGeneral) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *NotifyChangedSecAlarmGeneral) GetAdditionalInformation() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalInformation) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedSecAlarmGeneral) GetAdditionalInformationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalInformation) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *NotifyChangedSecAlarmGeneral) HasAdditionalInformation() bool {
	if o != nil && !IsNil(o.AdditionalInformation) {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given map[string]interface{} and assigns it to the AdditionalInformation field.
func (o *NotifyChangedSecAlarmGeneral) SetAdditionalInformation(v map[string]interface{}) {
	o.AdditionalInformation = v
}

// GetRootCauseIndicator returns the RootCauseIndicator field value if set, zero value otherwise.
func (o *NotifyChangedSecAlarmGeneral) GetRootCauseIndicator() bool {
	if o == nil || IsNil(o.RootCauseIndicator) {
		var ret bool
		return ret
	}
	return *o.RootCauseIndicator
}

// GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedSecAlarmGeneral) GetRootCauseIndicatorOk() (*bool, bool) {
	if o == nil || IsNil(o.RootCauseIndicator) {
		return nil, false
	}
	return o.RootCauseIndicator, true
}

// HasRootCauseIndicator returns a boolean if a field has been set.
func (o *NotifyChangedSecAlarmGeneral) HasRootCauseIndicator() bool {
	if o != nil && !IsNil(o.RootCauseIndicator) {
		return true
	}

	return false
}

// SetRootCauseIndicator gets a reference to the given bool and assigns it to the RootCauseIndicator field.
func (o *NotifyChangedSecAlarmGeneral) SetRootCauseIndicator(v bool) {
	o.RootCauseIndicator = &v
}

// GetServiceUser returns the ServiceUser field value
func (o *NotifyChangedSecAlarmGeneral) GetServiceUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceUser
}

// GetServiceUserOk returns a tuple with the ServiceUser field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedSecAlarmGeneral) GetServiceUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUser, true
}

// SetServiceUser sets field value
func (o *NotifyChangedSecAlarmGeneral) SetServiceUser(v string) {
	o.ServiceUser = v
}

// GetServiceProvider returns the ServiceProvider field value
func (o *NotifyChangedSecAlarmGeneral) GetServiceProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceProvider
}

// GetServiceProviderOk returns a tuple with the ServiceProvider field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedSecAlarmGeneral) GetServiceProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceProvider, true
}

// SetServiceProvider sets field value
func (o *NotifyChangedSecAlarmGeneral) SetServiceProvider(v string) {
	o.ServiceProvider = v
}

// GetSecurityAlarmDetector returns the SecurityAlarmDetector field value
func (o *NotifyChangedSecAlarmGeneral) GetSecurityAlarmDetector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityAlarmDetector
}

// GetSecurityAlarmDetectorOk returns a tuple with the SecurityAlarmDetector field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedSecAlarmGeneral) GetSecurityAlarmDetectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityAlarmDetector, true
}

// SetSecurityAlarmDetector sets field value
func (o *NotifyChangedSecAlarmGeneral) SetSecurityAlarmDetector(v string) {
	o.SecurityAlarmDetector = v
}

// GetChangedAlarmAttributes returns the ChangedAlarmAttributes field value if set, zero value otherwise.
func (o *NotifyChangedSecAlarmGeneral) GetChangedAlarmAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ChangedAlarmAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ChangedAlarmAttributes
}

// GetChangedAlarmAttributesOk returns a tuple with the ChangedAlarmAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyChangedSecAlarmGeneral) GetChangedAlarmAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ChangedAlarmAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ChangedAlarmAttributes, true
}

// HasChangedAlarmAttributes returns a boolean if a field has been set.
func (o *NotifyChangedSecAlarmGeneral) HasChangedAlarmAttributes() bool {
	if o != nil && !IsNil(o.ChangedAlarmAttributes) {
		return true
	}

	return false
}

// SetChangedAlarmAttributes gets a reference to the given map[string]interface{} and assigns it to the ChangedAlarmAttributes field.
func (o *NotifyChangedSecAlarmGeneral) SetChangedAlarmAttributes(v map[string]interface{}) {
	o.ChangedAlarmAttributes = v
}

func (o NotifyChangedSecAlarmGeneral) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyChangedSecAlarmGeneral) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNotificationHeader, errNotificationHeader := json.Marshal(o.NotificationHeader)
	if errNotificationHeader != nil {
		return map[string]interface{}{}, errNotificationHeader
	}
	errNotificationHeader = json.Unmarshal([]byte(serializedNotificationHeader), &toSerialize)
	if errNotificationHeader != nil {
		return map[string]interface{}{}, errNotificationHeader
	}
	toSerialize["alarmId"] = o.AlarmId
	toSerialize["alarmType"] = o.AlarmType
	if !IsNil(o.ProbableCause) {
		toSerialize["probableCause"] = o.ProbableCause
	}
	if !IsNil(o.PerceivedSeverity) {
		toSerialize["perceivedSeverity"] = o.PerceivedSeverity
	}
	if !IsNil(o.CorrelatedNotifications) {
		toSerialize["correlatedNotifications"] = o.CorrelatedNotifications
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
	toSerialize["serviceUser"] = o.ServiceUser
	toSerialize["serviceProvider"] = o.ServiceProvider
	toSerialize["securityAlarmDetector"] = o.SecurityAlarmDetector
	if !IsNil(o.ChangedAlarmAttributes) {
		toSerialize["changedAlarmAttributes"] = o.ChangedAlarmAttributes
	}
	return toSerialize, nil
}

type NullableNotifyChangedSecAlarmGeneral struct {
	value *NotifyChangedSecAlarmGeneral
	isSet bool
}

func (v NullableNotifyChangedSecAlarmGeneral) Get() *NotifyChangedSecAlarmGeneral {
	return v.value
}

func (v *NullableNotifyChangedSecAlarmGeneral) Set(val *NotifyChangedSecAlarmGeneral) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyChangedSecAlarmGeneral) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyChangedSecAlarmGeneral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyChangedSecAlarmGeneral(val *NotifyChangedSecAlarmGeneral) *NullableNotifyChangedSecAlarmGeneral {
	return &NullableNotifyChangedSecAlarmGeneral{value: val, isSet: true}
}

func (v NullableNotifyChangedSecAlarmGeneral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyChangedSecAlarmGeneral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"time"
)

// checks if the EventNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventNotification{}

// EventNotification Represents a notification on events that occurred.
type EventNotification struct {
	Event NwdafEvent `json:"event"`
	// string with format 'date-time' as defined in OpenAPI.
	Start *time.Time `json:"start,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStampGen *time.Time `json:"timeStampGen,omitempty"`
	FailNotifyCode *NwdafFailureCode `json:"failNotifyCode,omitempty"`
	// indicating a time in seconds.
	RvWaitTime *int32 `json:"rvWaitTime,omitempty"`
	AnaMetaInfo *AnalyticsMetadataInfo `json:"anaMetaInfo,omitempty"`
	NfLoadLevelInfos []NfLoadLevelInformation `json:"nfLoadLevelInfos,omitempty"`
	NsiLoadLevelInfos []NsiLoadLevelInfo `json:"nsiLoadLevelInfos,omitempty"`
	SliceLoadLevelInfo *SliceLoadLevelInformation `json:"sliceLoadLevelInfo,omitempty"`
	SvcExps []ServiceExperienceInfo `json:"svcExps,omitempty"`
	QosSustainInfos []QosSustainabilityInfo `json:"qosSustainInfos,omitempty"`
	UeComms []UeCommunication `json:"ueComms,omitempty"`
	UeMobs []UeMobility `json:"ueMobs,omitempty"`
	UserDataCongInfos []UserDataCongestionInfo `json:"userDataCongInfos,omitempty"`
	AbnorBehavrs []AbnormalBehaviour `json:"abnorBehavrs,omitempty"`
	NwPerfs []NetworkPerfInfo `json:"nwPerfs,omitempty"`
	DnPerfInfos []DnPerfInfo `json:"dnPerfInfos,omitempty"`
	DisperInfos []DispersionInfo `json:"disperInfos,omitempty"`
	RedTransInfos []RedundantTransmissionExpInfo `json:"redTransInfos,omitempty"`
	WlanInfos []WlanPerformanceInfo `json:"wlanInfos,omitempty"`
	SmccExps []SmcceInfo `json:"smccExps,omitempty"`
}

// NewEventNotification instantiates a new EventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventNotification(event NwdafEvent) *EventNotification {
	this := EventNotification{}
	this.Event = event
	return &this
}

// NewEventNotificationWithDefaults instantiates a new EventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventNotificationWithDefaults() *EventNotification {
	this := EventNotification{}
	return &this
}

// GetEvent returns the Event field value
func (o *EventNotification) GetEvent() NwdafEvent {
	if o == nil {
		var ret NwdafEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *EventNotification) GetEventOk() (*NwdafEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *EventNotification) SetEvent(v NwdafEvent) {
	o.Event = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *EventNotification) GetStart() time.Time {
	if o == nil || IsNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *EventNotification) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *EventNotification) SetStart(v time.Time) {
	o.Start = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *EventNotification) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *EventNotification) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *EventNotification) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetTimeStampGen returns the TimeStampGen field value if set, zero value otherwise.
func (o *EventNotification) GetTimeStampGen() time.Time {
	if o == nil || IsNil(o.TimeStampGen) {
		var ret time.Time
		return ret
	}
	return *o.TimeStampGen
}

// GetTimeStampGenOk returns a tuple with the TimeStampGen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetTimeStampGenOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeStampGen) {
		return nil, false
	}
	return o.TimeStampGen, true
}

// HasTimeStampGen returns a boolean if a field has been set.
func (o *EventNotification) HasTimeStampGen() bool {
	if o != nil && !IsNil(o.TimeStampGen) {
		return true
	}

	return false
}

// SetTimeStampGen gets a reference to the given time.Time and assigns it to the TimeStampGen field.
func (o *EventNotification) SetTimeStampGen(v time.Time) {
	o.TimeStampGen = &v
}

// GetFailNotifyCode returns the FailNotifyCode field value if set, zero value otherwise.
func (o *EventNotification) GetFailNotifyCode() NwdafFailureCode {
	if o == nil || IsNil(o.FailNotifyCode) {
		var ret NwdafFailureCode
		return ret
	}
	return *o.FailNotifyCode
}

// GetFailNotifyCodeOk returns a tuple with the FailNotifyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetFailNotifyCodeOk() (*NwdafFailureCode, bool) {
	if o == nil || IsNil(o.FailNotifyCode) {
		return nil, false
	}
	return o.FailNotifyCode, true
}

// HasFailNotifyCode returns a boolean if a field has been set.
func (o *EventNotification) HasFailNotifyCode() bool {
	if o != nil && !IsNil(o.FailNotifyCode) {
		return true
	}

	return false
}

// SetFailNotifyCode gets a reference to the given NwdafFailureCode and assigns it to the FailNotifyCode field.
func (o *EventNotification) SetFailNotifyCode(v NwdafFailureCode) {
	o.FailNotifyCode = &v
}

// GetRvWaitTime returns the RvWaitTime field value if set, zero value otherwise.
func (o *EventNotification) GetRvWaitTime() int32 {
	if o == nil || IsNil(o.RvWaitTime) {
		var ret int32
		return ret
	}
	return *o.RvWaitTime
}

// GetRvWaitTimeOk returns a tuple with the RvWaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetRvWaitTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.RvWaitTime) {
		return nil, false
	}
	return o.RvWaitTime, true
}

// HasRvWaitTime returns a boolean if a field has been set.
func (o *EventNotification) HasRvWaitTime() bool {
	if o != nil && !IsNil(o.RvWaitTime) {
		return true
	}

	return false
}

// SetRvWaitTime gets a reference to the given int32 and assigns it to the RvWaitTime field.
func (o *EventNotification) SetRvWaitTime(v int32) {
	o.RvWaitTime = &v
}

// GetAnaMetaInfo returns the AnaMetaInfo field value if set, zero value otherwise.
func (o *EventNotification) GetAnaMetaInfo() AnalyticsMetadataInfo {
	if o == nil || IsNil(o.AnaMetaInfo) {
		var ret AnalyticsMetadataInfo
		return ret
	}
	return *o.AnaMetaInfo
}

// GetAnaMetaInfoOk returns a tuple with the AnaMetaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetAnaMetaInfoOk() (*AnalyticsMetadataInfo, bool) {
	if o == nil || IsNil(o.AnaMetaInfo) {
		return nil, false
	}
	return o.AnaMetaInfo, true
}

// HasAnaMetaInfo returns a boolean if a field has been set.
func (o *EventNotification) HasAnaMetaInfo() bool {
	if o != nil && !IsNil(o.AnaMetaInfo) {
		return true
	}

	return false
}

// SetAnaMetaInfo gets a reference to the given AnalyticsMetadataInfo and assigns it to the AnaMetaInfo field.
func (o *EventNotification) SetAnaMetaInfo(v AnalyticsMetadataInfo) {
	o.AnaMetaInfo = &v
}

// GetNfLoadLevelInfos returns the NfLoadLevelInfos field value if set, zero value otherwise.
func (o *EventNotification) GetNfLoadLevelInfos() []NfLoadLevelInformation {
	if o == nil || IsNil(o.NfLoadLevelInfos) {
		var ret []NfLoadLevelInformation
		return ret
	}
	return o.NfLoadLevelInfos
}

// GetNfLoadLevelInfosOk returns a tuple with the NfLoadLevelInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetNfLoadLevelInfosOk() ([]NfLoadLevelInformation, bool) {
	if o == nil || IsNil(o.NfLoadLevelInfos) {
		return nil, false
	}
	return o.NfLoadLevelInfos, true
}

// HasNfLoadLevelInfos returns a boolean if a field has been set.
func (o *EventNotification) HasNfLoadLevelInfos() bool {
	if o != nil && !IsNil(o.NfLoadLevelInfos) {
		return true
	}

	return false
}

// SetNfLoadLevelInfos gets a reference to the given []NfLoadLevelInformation and assigns it to the NfLoadLevelInfos field.
func (o *EventNotification) SetNfLoadLevelInfos(v []NfLoadLevelInformation) {
	o.NfLoadLevelInfos = v
}

// GetNsiLoadLevelInfos returns the NsiLoadLevelInfos field value if set, zero value otherwise.
func (o *EventNotification) GetNsiLoadLevelInfos() []NsiLoadLevelInfo {
	if o == nil || IsNil(o.NsiLoadLevelInfos) {
		var ret []NsiLoadLevelInfo
		return ret
	}
	return o.NsiLoadLevelInfos
}

// GetNsiLoadLevelInfosOk returns a tuple with the NsiLoadLevelInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetNsiLoadLevelInfosOk() ([]NsiLoadLevelInfo, bool) {
	if o == nil || IsNil(o.NsiLoadLevelInfos) {
		return nil, false
	}
	return o.NsiLoadLevelInfos, true
}

// HasNsiLoadLevelInfos returns a boolean if a field has been set.
func (o *EventNotification) HasNsiLoadLevelInfos() bool {
	if o != nil && !IsNil(o.NsiLoadLevelInfos) {
		return true
	}

	return false
}

// SetNsiLoadLevelInfos gets a reference to the given []NsiLoadLevelInfo and assigns it to the NsiLoadLevelInfos field.
func (o *EventNotification) SetNsiLoadLevelInfos(v []NsiLoadLevelInfo) {
	o.NsiLoadLevelInfos = v
}

// GetSliceLoadLevelInfo returns the SliceLoadLevelInfo field value if set, zero value otherwise.
func (o *EventNotification) GetSliceLoadLevelInfo() SliceLoadLevelInformation {
	if o == nil || IsNil(o.SliceLoadLevelInfo) {
		var ret SliceLoadLevelInformation
		return ret
	}
	return *o.SliceLoadLevelInfo
}

// GetSliceLoadLevelInfoOk returns a tuple with the SliceLoadLevelInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetSliceLoadLevelInfoOk() (*SliceLoadLevelInformation, bool) {
	if o == nil || IsNil(o.SliceLoadLevelInfo) {
		return nil, false
	}
	return o.SliceLoadLevelInfo, true
}

// HasSliceLoadLevelInfo returns a boolean if a field has been set.
func (o *EventNotification) HasSliceLoadLevelInfo() bool {
	if o != nil && !IsNil(o.SliceLoadLevelInfo) {
		return true
	}

	return false
}

// SetSliceLoadLevelInfo gets a reference to the given SliceLoadLevelInformation and assigns it to the SliceLoadLevelInfo field.
func (o *EventNotification) SetSliceLoadLevelInfo(v SliceLoadLevelInformation) {
	o.SliceLoadLevelInfo = &v
}

// GetSvcExps returns the SvcExps field value if set, zero value otherwise.
func (o *EventNotification) GetSvcExps() []ServiceExperienceInfo {
	if o == nil || IsNil(o.SvcExps) {
		var ret []ServiceExperienceInfo
		return ret
	}
	return o.SvcExps
}

// GetSvcExpsOk returns a tuple with the SvcExps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetSvcExpsOk() ([]ServiceExperienceInfo, bool) {
	if o == nil || IsNil(o.SvcExps) {
		return nil, false
	}
	return o.SvcExps, true
}

// HasSvcExps returns a boolean if a field has been set.
func (o *EventNotification) HasSvcExps() bool {
	if o != nil && !IsNil(o.SvcExps) {
		return true
	}

	return false
}

// SetSvcExps gets a reference to the given []ServiceExperienceInfo and assigns it to the SvcExps field.
func (o *EventNotification) SetSvcExps(v []ServiceExperienceInfo) {
	o.SvcExps = v
}

// GetQosSustainInfos returns the QosSustainInfos field value if set, zero value otherwise.
func (o *EventNotification) GetQosSustainInfos() []QosSustainabilityInfo {
	if o == nil || IsNil(o.QosSustainInfos) {
		var ret []QosSustainabilityInfo
		return ret
	}
	return o.QosSustainInfos
}

// GetQosSustainInfosOk returns a tuple with the QosSustainInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetQosSustainInfosOk() ([]QosSustainabilityInfo, bool) {
	if o == nil || IsNil(o.QosSustainInfos) {
		return nil, false
	}
	return o.QosSustainInfos, true
}

// HasQosSustainInfos returns a boolean if a field has been set.
func (o *EventNotification) HasQosSustainInfos() bool {
	if o != nil && !IsNil(o.QosSustainInfos) {
		return true
	}

	return false
}

// SetQosSustainInfos gets a reference to the given []QosSustainabilityInfo and assigns it to the QosSustainInfos field.
func (o *EventNotification) SetQosSustainInfos(v []QosSustainabilityInfo) {
	o.QosSustainInfos = v
}

// GetUeComms returns the UeComms field value if set, zero value otherwise.
func (o *EventNotification) GetUeComms() []UeCommunication {
	if o == nil || IsNil(o.UeComms) {
		var ret []UeCommunication
		return ret
	}
	return o.UeComms
}

// GetUeCommsOk returns a tuple with the UeComms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetUeCommsOk() ([]UeCommunication, bool) {
	if o == nil || IsNil(o.UeComms) {
		return nil, false
	}
	return o.UeComms, true
}

// HasUeComms returns a boolean if a field has been set.
func (o *EventNotification) HasUeComms() bool {
	if o != nil && !IsNil(o.UeComms) {
		return true
	}

	return false
}

// SetUeComms gets a reference to the given []UeCommunication and assigns it to the UeComms field.
func (o *EventNotification) SetUeComms(v []UeCommunication) {
	o.UeComms = v
}

// GetUeMobs returns the UeMobs field value if set, zero value otherwise.
func (o *EventNotification) GetUeMobs() []UeMobility {
	if o == nil || IsNil(o.UeMobs) {
		var ret []UeMobility
		return ret
	}
	return o.UeMobs
}

// GetUeMobsOk returns a tuple with the UeMobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetUeMobsOk() ([]UeMobility, bool) {
	if o == nil || IsNil(o.UeMobs) {
		return nil, false
	}
	return o.UeMobs, true
}

// HasUeMobs returns a boolean if a field has been set.
func (o *EventNotification) HasUeMobs() bool {
	if o != nil && !IsNil(o.UeMobs) {
		return true
	}

	return false
}

// SetUeMobs gets a reference to the given []UeMobility and assigns it to the UeMobs field.
func (o *EventNotification) SetUeMobs(v []UeMobility) {
	o.UeMobs = v
}

// GetUserDataCongInfos returns the UserDataCongInfos field value if set, zero value otherwise.
func (o *EventNotification) GetUserDataCongInfos() []UserDataCongestionInfo {
	if o == nil || IsNil(o.UserDataCongInfos) {
		var ret []UserDataCongestionInfo
		return ret
	}
	return o.UserDataCongInfos
}

// GetUserDataCongInfosOk returns a tuple with the UserDataCongInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetUserDataCongInfosOk() ([]UserDataCongestionInfo, bool) {
	if o == nil || IsNil(o.UserDataCongInfos) {
		return nil, false
	}
	return o.UserDataCongInfos, true
}

// HasUserDataCongInfos returns a boolean if a field has been set.
func (o *EventNotification) HasUserDataCongInfos() bool {
	if o != nil && !IsNil(o.UserDataCongInfos) {
		return true
	}

	return false
}

// SetUserDataCongInfos gets a reference to the given []UserDataCongestionInfo and assigns it to the UserDataCongInfos field.
func (o *EventNotification) SetUserDataCongInfos(v []UserDataCongestionInfo) {
	o.UserDataCongInfos = v
}

// GetAbnorBehavrs returns the AbnorBehavrs field value if set, zero value otherwise.
func (o *EventNotification) GetAbnorBehavrs() []AbnormalBehaviour {
	if o == nil || IsNil(o.AbnorBehavrs) {
		var ret []AbnormalBehaviour
		return ret
	}
	return o.AbnorBehavrs
}

// GetAbnorBehavrsOk returns a tuple with the AbnorBehavrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetAbnorBehavrsOk() ([]AbnormalBehaviour, bool) {
	if o == nil || IsNil(o.AbnorBehavrs) {
		return nil, false
	}
	return o.AbnorBehavrs, true
}

// HasAbnorBehavrs returns a boolean if a field has been set.
func (o *EventNotification) HasAbnorBehavrs() bool {
	if o != nil && !IsNil(o.AbnorBehavrs) {
		return true
	}

	return false
}

// SetAbnorBehavrs gets a reference to the given []AbnormalBehaviour and assigns it to the AbnorBehavrs field.
func (o *EventNotification) SetAbnorBehavrs(v []AbnormalBehaviour) {
	o.AbnorBehavrs = v
}

// GetNwPerfs returns the NwPerfs field value if set, zero value otherwise.
func (o *EventNotification) GetNwPerfs() []NetworkPerfInfo {
	if o == nil || IsNil(o.NwPerfs) {
		var ret []NetworkPerfInfo
		return ret
	}
	return o.NwPerfs
}

// GetNwPerfsOk returns a tuple with the NwPerfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetNwPerfsOk() ([]NetworkPerfInfo, bool) {
	if o == nil || IsNil(o.NwPerfs) {
		return nil, false
	}
	return o.NwPerfs, true
}

// HasNwPerfs returns a boolean if a field has been set.
func (o *EventNotification) HasNwPerfs() bool {
	if o != nil && !IsNil(o.NwPerfs) {
		return true
	}

	return false
}

// SetNwPerfs gets a reference to the given []NetworkPerfInfo and assigns it to the NwPerfs field.
func (o *EventNotification) SetNwPerfs(v []NetworkPerfInfo) {
	o.NwPerfs = v
}

// GetDnPerfInfos returns the DnPerfInfos field value if set, zero value otherwise.
func (o *EventNotification) GetDnPerfInfos() []DnPerfInfo {
	if o == nil || IsNil(o.DnPerfInfos) {
		var ret []DnPerfInfo
		return ret
	}
	return o.DnPerfInfos
}

// GetDnPerfInfosOk returns a tuple with the DnPerfInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetDnPerfInfosOk() ([]DnPerfInfo, bool) {
	if o == nil || IsNil(o.DnPerfInfos) {
		return nil, false
	}
	return o.DnPerfInfos, true
}

// HasDnPerfInfos returns a boolean if a field has been set.
func (o *EventNotification) HasDnPerfInfos() bool {
	if o != nil && !IsNil(o.DnPerfInfos) {
		return true
	}

	return false
}

// SetDnPerfInfos gets a reference to the given []DnPerfInfo and assigns it to the DnPerfInfos field.
func (o *EventNotification) SetDnPerfInfos(v []DnPerfInfo) {
	o.DnPerfInfos = v
}

// GetDisperInfos returns the DisperInfos field value if set, zero value otherwise.
func (o *EventNotification) GetDisperInfos() []DispersionInfo {
	if o == nil || IsNil(o.DisperInfos) {
		var ret []DispersionInfo
		return ret
	}
	return o.DisperInfos
}

// GetDisperInfosOk returns a tuple with the DisperInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetDisperInfosOk() ([]DispersionInfo, bool) {
	if o == nil || IsNil(o.DisperInfos) {
		return nil, false
	}
	return o.DisperInfos, true
}

// HasDisperInfos returns a boolean if a field has been set.
func (o *EventNotification) HasDisperInfos() bool {
	if o != nil && !IsNil(o.DisperInfos) {
		return true
	}

	return false
}

// SetDisperInfos gets a reference to the given []DispersionInfo and assigns it to the DisperInfos field.
func (o *EventNotification) SetDisperInfos(v []DispersionInfo) {
	o.DisperInfos = v
}

// GetRedTransInfos returns the RedTransInfos field value if set, zero value otherwise.
func (o *EventNotification) GetRedTransInfos() []RedundantTransmissionExpInfo {
	if o == nil || IsNil(o.RedTransInfos) {
		var ret []RedundantTransmissionExpInfo
		return ret
	}
	return o.RedTransInfos
}

// GetRedTransInfosOk returns a tuple with the RedTransInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetRedTransInfosOk() ([]RedundantTransmissionExpInfo, bool) {
	if o == nil || IsNil(o.RedTransInfos) {
		return nil, false
	}
	return o.RedTransInfos, true
}

// HasRedTransInfos returns a boolean if a field has been set.
func (o *EventNotification) HasRedTransInfos() bool {
	if o != nil && !IsNil(o.RedTransInfos) {
		return true
	}

	return false
}

// SetRedTransInfos gets a reference to the given []RedundantTransmissionExpInfo and assigns it to the RedTransInfos field.
func (o *EventNotification) SetRedTransInfos(v []RedundantTransmissionExpInfo) {
	o.RedTransInfos = v
}

// GetWlanInfos returns the WlanInfos field value if set, zero value otherwise.
func (o *EventNotification) GetWlanInfos() []WlanPerformanceInfo {
	if o == nil || IsNil(o.WlanInfos) {
		var ret []WlanPerformanceInfo
		return ret
	}
	return o.WlanInfos
}

// GetWlanInfosOk returns a tuple with the WlanInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetWlanInfosOk() ([]WlanPerformanceInfo, bool) {
	if o == nil || IsNil(o.WlanInfos) {
		return nil, false
	}
	return o.WlanInfos, true
}

// HasWlanInfos returns a boolean if a field has been set.
func (o *EventNotification) HasWlanInfos() bool {
	if o != nil && !IsNil(o.WlanInfos) {
		return true
	}

	return false
}

// SetWlanInfos gets a reference to the given []WlanPerformanceInfo and assigns it to the WlanInfos field.
func (o *EventNotification) SetWlanInfos(v []WlanPerformanceInfo) {
	o.WlanInfos = v
}

// GetSmccExps returns the SmccExps field value if set, zero value otherwise.
func (o *EventNotification) GetSmccExps() []SmcceInfo {
	if o == nil || IsNil(o.SmccExps) {
		var ret []SmcceInfo
		return ret
	}
	return o.SmccExps
}

// GetSmccExpsOk returns a tuple with the SmccExps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetSmccExpsOk() ([]SmcceInfo, bool) {
	if o == nil || IsNil(o.SmccExps) {
		return nil, false
	}
	return o.SmccExps, true
}

// HasSmccExps returns a boolean if a field has been set.
func (o *EventNotification) HasSmccExps() bool {
	if o != nil && !IsNil(o.SmccExps) {
		return true
	}

	return false
}

// SetSmccExps gets a reference to the given []SmcceInfo and assigns it to the SmccExps field.
func (o *EventNotification) SetSmccExps(v []SmcceInfo) {
	o.SmccExps = v
}

func (o EventNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.TimeStampGen) {
		toSerialize["timeStampGen"] = o.TimeStampGen
	}
	if !IsNil(o.FailNotifyCode) {
		toSerialize["failNotifyCode"] = o.FailNotifyCode
	}
	if !IsNil(o.RvWaitTime) {
		toSerialize["rvWaitTime"] = o.RvWaitTime
	}
	if !IsNil(o.AnaMetaInfo) {
		toSerialize["anaMetaInfo"] = o.AnaMetaInfo
	}
	if !IsNil(o.NfLoadLevelInfos) {
		toSerialize["nfLoadLevelInfos"] = o.NfLoadLevelInfos
	}
	if !IsNil(o.NsiLoadLevelInfos) {
		toSerialize["nsiLoadLevelInfos"] = o.NsiLoadLevelInfos
	}
	if !IsNil(o.SliceLoadLevelInfo) {
		toSerialize["sliceLoadLevelInfo"] = o.SliceLoadLevelInfo
	}
	if !IsNil(o.SvcExps) {
		toSerialize["svcExps"] = o.SvcExps
	}
	if !IsNil(o.QosSustainInfos) {
		toSerialize["qosSustainInfos"] = o.QosSustainInfos
	}
	if !IsNil(o.UeComms) {
		toSerialize["ueComms"] = o.UeComms
	}
	if !IsNil(o.UeMobs) {
		toSerialize["ueMobs"] = o.UeMobs
	}
	if !IsNil(o.UserDataCongInfos) {
		toSerialize["userDataCongInfos"] = o.UserDataCongInfos
	}
	if !IsNil(o.AbnorBehavrs) {
		toSerialize["abnorBehavrs"] = o.AbnorBehavrs
	}
	if !IsNil(o.NwPerfs) {
		toSerialize["nwPerfs"] = o.NwPerfs
	}
	if !IsNil(o.DnPerfInfos) {
		toSerialize["dnPerfInfos"] = o.DnPerfInfos
	}
	if !IsNil(o.DisperInfos) {
		toSerialize["disperInfos"] = o.DisperInfos
	}
	if !IsNil(o.RedTransInfos) {
		toSerialize["redTransInfos"] = o.RedTransInfos
	}
	if !IsNil(o.WlanInfos) {
		toSerialize["wlanInfos"] = o.WlanInfos
	}
	if !IsNil(o.SmccExps) {
		toSerialize["smccExps"] = o.SmccExps
	}
	return toSerialize, nil
}

type NullableEventNotification struct {
	value *EventNotification
	isSet bool
}

func (v NullableEventNotification) Get() *EventNotification {
	return v.value
}

func (v *NullableEventNotification) Set(val *EventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventNotification(val *EventNotification) *NullableEventNotification {
	return &NullableEventNotification{value: val, isSet: true}
}

func (v NullableEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



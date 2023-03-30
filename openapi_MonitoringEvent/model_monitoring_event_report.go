/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
	"time"
)

// checks if the MonitoringEventReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringEventReport{}

// MonitoringEventReport Represents an event monitoring report.
type MonitoringEventReport struct {
	ImeiChange *AssociationType `json:"imeiChange,omitempty"`
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information.
	ExternalId *string `json:"externalId,omitempty"`
	IdleStatusInfo *IdleStatusInfo `json:"idleStatusInfo,omitempty"`
	LocationInfo *LocationInfo `json:"locationInfo,omitempty"`
	LocFailureCause *LocationFailureCause `json:"locFailureCause,omitempty"`
	// If \"monitoringType\" is \"LOSS_OF_CONNECTIVITY\", this parameter shall be included if available to identify the reason why loss of connectivity is reported. Refer to 3GPP TS 29.336 clause 8.4.58.
	LossOfConnectReason *int32 `json:"lossOfConnectReason,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	MaxUEAvailabilityTime *time.Time `json:"maxUEAvailabilityTime,omitempty"`
	// string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN.
	Msisdn *string `json:"msisdn,omitempty"`
	MonitoringType MonitoringType `json:"monitoringType"`
	UePerLocationReport *UePerLocationReport `json:"uePerLocationReport,omitempty"`
	PlmnId *PlmnId `json:"plmnId,omitempty"`
	ReachabilityType *ReachabilityType `json:"reachabilityType,omitempty"`
	// If \"monitoringType\" is \"ROAMING_STATUS\", this parameter shall be set to \"true\" if the UE is on roaming status. Set to false or omitted otherwise.
	RoamingStatus *bool `json:"roamingStatus,omitempty"`
	FailureCause *FailureCause `json:"failureCause,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	EventTime *time.Time `json:"eventTime,omitempty"`
	PdnConnInfoList []PdnConnectionInformation `json:"pdnConnInfoList,omitempty"`
	DddStatus *DlDataDeliveryStatus `json:"dddStatus,omitempty"`
	DddTrafDescriptor *DddTrafficDescriptor `json:"dddTrafDescriptor,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	MaxWaitTime *time.Time `json:"maxWaitTime,omitempty"`
	ApiCaps []ApiCapabilityInfo `json:"apiCaps,omitempty"`
	NSStatusInfo *SACEventStatus `json:"nSStatusInfo,omitempty"`
	AfServiceId *string `json:"afServiceId,omitempty"`
	// If \"monitoringType\" is \"AREA_OF_INTEREST\", this parameter may be included to identify the UAV.
	ServLevelDevId *string `json:"servLevelDevId,omitempty"`
	// If \"monitoringType\" is \"AREA_OF_INTEREST\", this parameter shall be set to true if the specified UAV is in the monitoring area. Set to false or omitted otherwise.
	UavPresInd *bool `json:"uavPresInd,omitempty"`
}

// NewMonitoringEventReport instantiates a new MonitoringEventReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringEventReport(monitoringType MonitoringType) *MonitoringEventReport {
	this := MonitoringEventReport{}
	this.MonitoringType = monitoringType
	return &this
}

// NewMonitoringEventReportWithDefaults instantiates a new MonitoringEventReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringEventReportWithDefaults() *MonitoringEventReport {
	this := MonitoringEventReport{}
	return &this
}

// GetImeiChange returns the ImeiChange field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetImeiChange() AssociationType {
	if o == nil || IsNil(o.ImeiChange) {
		var ret AssociationType
		return ret
	}
	return *o.ImeiChange
}

// GetImeiChangeOk returns a tuple with the ImeiChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetImeiChangeOk() (*AssociationType, bool) {
	if o == nil || IsNil(o.ImeiChange) {
		return nil, false
	}
	return o.ImeiChange, true
}

// HasImeiChange returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasImeiChange() bool {
	if o != nil && !IsNil(o.ImeiChange) {
		return true
	}

	return false
}

// SetImeiChange gets a reference to the given AssociationType and assigns it to the ImeiChange field.
func (o *MonitoringEventReport) SetImeiChange(v AssociationType) {
	o.ImeiChange = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *MonitoringEventReport) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetIdleStatusInfo returns the IdleStatusInfo field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetIdleStatusInfo() IdleStatusInfo {
	if o == nil || IsNil(o.IdleStatusInfo) {
		var ret IdleStatusInfo
		return ret
	}
	return *o.IdleStatusInfo
}

// GetIdleStatusInfoOk returns a tuple with the IdleStatusInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetIdleStatusInfoOk() (*IdleStatusInfo, bool) {
	if o == nil || IsNil(o.IdleStatusInfo) {
		return nil, false
	}
	return o.IdleStatusInfo, true
}

// HasIdleStatusInfo returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasIdleStatusInfo() bool {
	if o != nil && !IsNil(o.IdleStatusInfo) {
		return true
	}

	return false
}

// SetIdleStatusInfo gets a reference to the given IdleStatusInfo and assigns it to the IdleStatusInfo field.
func (o *MonitoringEventReport) SetIdleStatusInfo(v IdleStatusInfo) {
	o.IdleStatusInfo = &v
}

// GetLocationInfo returns the LocationInfo field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetLocationInfo() LocationInfo {
	if o == nil || IsNil(o.LocationInfo) {
		var ret LocationInfo
		return ret
	}
	return *o.LocationInfo
}

// GetLocationInfoOk returns a tuple with the LocationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetLocationInfoOk() (*LocationInfo, bool) {
	if o == nil || IsNil(o.LocationInfo) {
		return nil, false
	}
	return o.LocationInfo, true
}

// HasLocationInfo returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasLocationInfo() bool {
	if o != nil && !IsNil(o.LocationInfo) {
		return true
	}

	return false
}

// SetLocationInfo gets a reference to the given LocationInfo and assigns it to the LocationInfo field.
func (o *MonitoringEventReport) SetLocationInfo(v LocationInfo) {
	o.LocationInfo = &v
}

// GetLocFailureCause returns the LocFailureCause field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetLocFailureCause() LocationFailureCause {
	if o == nil || IsNil(o.LocFailureCause) {
		var ret LocationFailureCause
		return ret
	}
	return *o.LocFailureCause
}

// GetLocFailureCauseOk returns a tuple with the LocFailureCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetLocFailureCauseOk() (*LocationFailureCause, bool) {
	if o == nil || IsNil(o.LocFailureCause) {
		return nil, false
	}
	return o.LocFailureCause, true
}

// HasLocFailureCause returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasLocFailureCause() bool {
	if o != nil && !IsNil(o.LocFailureCause) {
		return true
	}

	return false
}

// SetLocFailureCause gets a reference to the given LocationFailureCause and assigns it to the LocFailureCause field.
func (o *MonitoringEventReport) SetLocFailureCause(v LocationFailureCause) {
	o.LocFailureCause = &v
}

// GetLossOfConnectReason returns the LossOfConnectReason field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetLossOfConnectReason() int32 {
	if o == nil || IsNil(o.LossOfConnectReason) {
		var ret int32
		return ret
	}
	return *o.LossOfConnectReason
}

// GetLossOfConnectReasonOk returns a tuple with the LossOfConnectReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetLossOfConnectReasonOk() (*int32, bool) {
	if o == nil || IsNil(o.LossOfConnectReason) {
		return nil, false
	}
	return o.LossOfConnectReason, true
}

// HasLossOfConnectReason returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasLossOfConnectReason() bool {
	if o != nil && !IsNil(o.LossOfConnectReason) {
		return true
	}

	return false
}

// SetLossOfConnectReason gets a reference to the given int32 and assigns it to the LossOfConnectReason field.
func (o *MonitoringEventReport) SetLossOfConnectReason(v int32) {
	o.LossOfConnectReason = &v
}

// GetMaxUEAvailabilityTime returns the MaxUEAvailabilityTime field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetMaxUEAvailabilityTime() time.Time {
	if o == nil || IsNil(o.MaxUEAvailabilityTime) {
		var ret time.Time
		return ret
	}
	return *o.MaxUEAvailabilityTime
}

// GetMaxUEAvailabilityTimeOk returns a tuple with the MaxUEAvailabilityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetMaxUEAvailabilityTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MaxUEAvailabilityTime) {
		return nil, false
	}
	return o.MaxUEAvailabilityTime, true
}

// HasMaxUEAvailabilityTime returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasMaxUEAvailabilityTime() bool {
	if o != nil && !IsNil(o.MaxUEAvailabilityTime) {
		return true
	}

	return false
}

// SetMaxUEAvailabilityTime gets a reference to the given time.Time and assigns it to the MaxUEAvailabilityTime field.
func (o *MonitoringEventReport) SetMaxUEAvailabilityTime(v time.Time) {
	o.MaxUEAvailabilityTime = &v
}

// GetMsisdn returns the Msisdn field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetMsisdn() string {
	if o == nil || IsNil(o.Msisdn) {
		var ret string
		return ret
	}
	return *o.Msisdn
}

// GetMsisdnOk returns a tuple with the Msisdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetMsisdnOk() (*string, bool) {
	if o == nil || IsNil(o.Msisdn) {
		return nil, false
	}
	return o.Msisdn, true
}

// HasMsisdn returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasMsisdn() bool {
	if o != nil && !IsNil(o.Msisdn) {
		return true
	}

	return false
}

// SetMsisdn gets a reference to the given string and assigns it to the Msisdn field.
func (o *MonitoringEventReport) SetMsisdn(v string) {
	o.Msisdn = &v
}

// GetMonitoringType returns the MonitoringType field value
func (o *MonitoringEventReport) GetMonitoringType() MonitoringType {
	if o == nil {
		var ret MonitoringType
		return ret
	}

	return o.MonitoringType
}

// GetMonitoringTypeOk returns a tuple with the MonitoringType field value
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetMonitoringTypeOk() (*MonitoringType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitoringType, true
}

// SetMonitoringType sets field value
func (o *MonitoringEventReport) SetMonitoringType(v MonitoringType) {
	o.MonitoringType = v
}

// GetUePerLocationReport returns the UePerLocationReport field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetUePerLocationReport() UePerLocationReport {
	if o == nil || IsNil(o.UePerLocationReport) {
		var ret UePerLocationReport
		return ret
	}
	return *o.UePerLocationReport
}

// GetUePerLocationReportOk returns a tuple with the UePerLocationReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetUePerLocationReportOk() (*UePerLocationReport, bool) {
	if o == nil || IsNil(o.UePerLocationReport) {
		return nil, false
	}
	return o.UePerLocationReport, true
}

// HasUePerLocationReport returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasUePerLocationReport() bool {
	if o != nil && !IsNil(o.UePerLocationReport) {
		return true
	}

	return false
}

// SetUePerLocationReport gets a reference to the given UePerLocationReport and assigns it to the UePerLocationReport field.
func (o *MonitoringEventReport) SetUePerLocationReport(v UePerLocationReport) {
	o.UePerLocationReport = &v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetPlmnId() PlmnId {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *MonitoringEventReport) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetReachabilityType returns the ReachabilityType field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetReachabilityType() ReachabilityType {
	if o == nil || IsNil(o.ReachabilityType) {
		var ret ReachabilityType
		return ret
	}
	return *o.ReachabilityType
}

// GetReachabilityTypeOk returns a tuple with the ReachabilityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetReachabilityTypeOk() (*ReachabilityType, bool) {
	if o == nil || IsNil(o.ReachabilityType) {
		return nil, false
	}
	return o.ReachabilityType, true
}

// HasReachabilityType returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasReachabilityType() bool {
	if o != nil && !IsNil(o.ReachabilityType) {
		return true
	}

	return false
}

// SetReachabilityType gets a reference to the given ReachabilityType and assigns it to the ReachabilityType field.
func (o *MonitoringEventReport) SetReachabilityType(v ReachabilityType) {
	o.ReachabilityType = &v
}

// GetRoamingStatus returns the RoamingStatus field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetRoamingStatus() bool {
	if o == nil || IsNil(o.RoamingStatus) {
		var ret bool
		return ret
	}
	return *o.RoamingStatus
}

// GetRoamingStatusOk returns a tuple with the RoamingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetRoamingStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.RoamingStatus) {
		return nil, false
	}
	return o.RoamingStatus, true
}

// HasRoamingStatus returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasRoamingStatus() bool {
	if o != nil && !IsNil(o.RoamingStatus) {
		return true
	}

	return false
}

// SetRoamingStatus gets a reference to the given bool and assigns it to the RoamingStatus field.
func (o *MonitoringEventReport) SetRoamingStatus(v bool) {
	o.RoamingStatus = &v
}

// GetFailureCause returns the FailureCause field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetFailureCause() FailureCause {
	if o == nil || IsNil(o.FailureCause) {
		var ret FailureCause
		return ret
	}
	return *o.FailureCause
}

// GetFailureCauseOk returns a tuple with the FailureCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetFailureCauseOk() (*FailureCause, bool) {
	if o == nil || IsNil(o.FailureCause) {
		return nil, false
	}
	return o.FailureCause, true
}

// HasFailureCause returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasFailureCause() bool {
	if o != nil && !IsNil(o.FailureCause) {
		return true
	}

	return false
}

// SetFailureCause gets a reference to the given FailureCause and assigns it to the FailureCause field.
func (o *MonitoringEventReport) SetFailureCause(v FailureCause) {
	o.FailureCause = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetEventTime() time.Time {
	if o == nil || IsNil(o.EventTime) {
		var ret time.Time
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetEventTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given time.Time and assigns it to the EventTime field.
func (o *MonitoringEventReport) SetEventTime(v time.Time) {
	o.EventTime = &v
}

// GetPdnConnInfoList returns the PdnConnInfoList field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetPdnConnInfoList() []PdnConnectionInformation {
	if o == nil || IsNil(o.PdnConnInfoList) {
		var ret []PdnConnectionInformation
		return ret
	}
	return o.PdnConnInfoList
}

// GetPdnConnInfoListOk returns a tuple with the PdnConnInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetPdnConnInfoListOk() ([]PdnConnectionInformation, bool) {
	if o == nil || IsNil(o.PdnConnInfoList) {
		return nil, false
	}
	return o.PdnConnInfoList, true
}

// HasPdnConnInfoList returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasPdnConnInfoList() bool {
	if o != nil && !IsNil(o.PdnConnInfoList) {
		return true
	}

	return false
}

// SetPdnConnInfoList gets a reference to the given []PdnConnectionInformation and assigns it to the PdnConnInfoList field.
func (o *MonitoringEventReport) SetPdnConnInfoList(v []PdnConnectionInformation) {
	o.PdnConnInfoList = v
}

// GetDddStatus returns the DddStatus field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetDddStatus() DlDataDeliveryStatus {
	if o == nil || IsNil(o.DddStatus) {
		var ret DlDataDeliveryStatus
		return ret
	}
	return *o.DddStatus
}

// GetDddStatusOk returns a tuple with the DddStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetDddStatusOk() (*DlDataDeliveryStatus, bool) {
	if o == nil || IsNil(o.DddStatus) {
		return nil, false
	}
	return o.DddStatus, true
}

// HasDddStatus returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasDddStatus() bool {
	if o != nil && !IsNil(o.DddStatus) {
		return true
	}

	return false
}

// SetDddStatus gets a reference to the given DlDataDeliveryStatus and assigns it to the DddStatus field.
func (o *MonitoringEventReport) SetDddStatus(v DlDataDeliveryStatus) {
	o.DddStatus = &v
}

// GetDddTrafDescriptor returns the DddTrafDescriptor field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetDddTrafDescriptor() DddTrafficDescriptor {
	if o == nil || IsNil(o.DddTrafDescriptor) {
		var ret DddTrafficDescriptor
		return ret
	}
	return *o.DddTrafDescriptor
}

// GetDddTrafDescriptorOk returns a tuple with the DddTrafDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetDddTrafDescriptorOk() (*DddTrafficDescriptor, bool) {
	if o == nil || IsNil(o.DddTrafDescriptor) {
		return nil, false
	}
	return o.DddTrafDescriptor, true
}

// HasDddTrafDescriptor returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasDddTrafDescriptor() bool {
	if o != nil && !IsNil(o.DddTrafDescriptor) {
		return true
	}

	return false
}

// SetDddTrafDescriptor gets a reference to the given DddTrafficDescriptor and assigns it to the DddTrafDescriptor field.
func (o *MonitoringEventReport) SetDddTrafDescriptor(v DddTrafficDescriptor) {
	o.DddTrafDescriptor = &v
}

// GetMaxWaitTime returns the MaxWaitTime field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetMaxWaitTime() time.Time {
	if o == nil || IsNil(o.MaxWaitTime) {
		var ret time.Time
		return ret
	}
	return *o.MaxWaitTime
}

// GetMaxWaitTimeOk returns a tuple with the MaxWaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetMaxWaitTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MaxWaitTime) {
		return nil, false
	}
	return o.MaxWaitTime, true
}

// HasMaxWaitTime returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasMaxWaitTime() bool {
	if o != nil && !IsNil(o.MaxWaitTime) {
		return true
	}

	return false
}

// SetMaxWaitTime gets a reference to the given time.Time and assigns it to the MaxWaitTime field.
func (o *MonitoringEventReport) SetMaxWaitTime(v time.Time) {
	o.MaxWaitTime = &v
}

// GetApiCaps returns the ApiCaps field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetApiCaps() []ApiCapabilityInfo {
	if o == nil || IsNil(o.ApiCaps) {
		var ret []ApiCapabilityInfo
		return ret
	}
	return o.ApiCaps
}

// GetApiCapsOk returns a tuple with the ApiCaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetApiCapsOk() ([]ApiCapabilityInfo, bool) {
	if o == nil || IsNil(o.ApiCaps) {
		return nil, false
	}
	return o.ApiCaps, true
}

// HasApiCaps returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasApiCaps() bool {
	if o != nil && !IsNil(o.ApiCaps) {
		return true
	}

	return false
}

// SetApiCaps gets a reference to the given []ApiCapabilityInfo and assigns it to the ApiCaps field.
func (o *MonitoringEventReport) SetApiCaps(v []ApiCapabilityInfo) {
	o.ApiCaps = v
}

// GetNSStatusInfo returns the NSStatusInfo field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetNSStatusInfo() SACEventStatus {
	if o == nil || IsNil(o.NSStatusInfo) {
		var ret SACEventStatus
		return ret
	}
	return *o.NSStatusInfo
}

// GetNSStatusInfoOk returns a tuple with the NSStatusInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetNSStatusInfoOk() (*SACEventStatus, bool) {
	if o == nil || IsNil(o.NSStatusInfo) {
		return nil, false
	}
	return o.NSStatusInfo, true
}

// HasNSStatusInfo returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasNSStatusInfo() bool {
	if o != nil && !IsNil(o.NSStatusInfo) {
		return true
	}

	return false
}

// SetNSStatusInfo gets a reference to the given SACEventStatus and assigns it to the NSStatusInfo field.
func (o *MonitoringEventReport) SetNSStatusInfo(v SACEventStatus) {
	o.NSStatusInfo = &v
}

// GetAfServiceId returns the AfServiceId field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetAfServiceId() string {
	if o == nil || IsNil(o.AfServiceId) {
		var ret string
		return ret
	}
	return *o.AfServiceId
}

// GetAfServiceIdOk returns a tuple with the AfServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetAfServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfServiceId) {
		return nil, false
	}
	return o.AfServiceId, true
}

// HasAfServiceId returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasAfServiceId() bool {
	if o != nil && !IsNil(o.AfServiceId) {
		return true
	}

	return false
}

// SetAfServiceId gets a reference to the given string and assigns it to the AfServiceId field.
func (o *MonitoringEventReport) SetAfServiceId(v string) {
	o.AfServiceId = &v
}

// GetServLevelDevId returns the ServLevelDevId field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetServLevelDevId() string {
	if o == nil || IsNil(o.ServLevelDevId) {
		var ret string
		return ret
	}
	return *o.ServLevelDevId
}

// GetServLevelDevIdOk returns a tuple with the ServLevelDevId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetServLevelDevIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServLevelDevId) {
		return nil, false
	}
	return o.ServLevelDevId, true
}

// HasServLevelDevId returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasServLevelDevId() bool {
	if o != nil && !IsNil(o.ServLevelDevId) {
		return true
	}

	return false
}

// SetServLevelDevId gets a reference to the given string and assigns it to the ServLevelDevId field.
func (o *MonitoringEventReport) SetServLevelDevId(v string) {
	o.ServLevelDevId = &v
}

// GetUavPresInd returns the UavPresInd field value if set, zero value otherwise.
func (o *MonitoringEventReport) GetUavPresInd() bool {
	if o == nil || IsNil(o.UavPresInd) {
		var ret bool
		return ret
	}
	return *o.UavPresInd
}

// GetUavPresIndOk returns a tuple with the UavPresInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventReport) GetUavPresIndOk() (*bool, bool) {
	if o == nil || IsNil(o.UavPresInd) {
		return nil, false
	}
	return o.UavPresInd, true
}

// HasUavPresInd returns a boolean if a field has been set.
func (o *MonitoringEventReport) HasUavPresInd() bool {
	if o != nil && !IsNil(o.UavPresInd) {
		return true
	}

	return false
}

// SetUavPresInd gets a reference to the given bool and assigns it to the UavPresInd field.
func (o *MonitoringEventReport) SetUavPresInd(v bool) {
	o.UavPresInd = &v
}

func (o MonitoringEventReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringEventReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImeiChange) {
		toSerialize["imeiChange"] = o.ImeiChange
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.IdleStatusInfo) {
		toSerialize["idleStatusInfo"] = o.IdleStatusInfo
	}
	if !IsNil(o.LocationInfo) {
		toSerialize["locationInfo"] = o.LocationInfo
	}
	if !IsNil(o.LocFailureCause) {
		toSerialize["locFailureCause"] = o.LocFailureCause
	}
	if !IsNil(o.LossOfConnectReason) {
		toSerialize["lossOfConnectReason"] = o.LossOfConnectReason
	}
	if !IsNil(o.MaxUEAvailabilityTime) {
		toSerialize["maxUEAvailabilityTime"] = o.MaxUEAvailabilityTime
	}
	if !IsNil(o.Msisdn) {
		toSerialize["msisdn"] = o.Msisdn
	}
	toSerialize["monitoringType"] = o.MonitoringType
	if !IsNil(o.UePerLocationReport) {
		toSerialize["uePerLocationReport"] = o.UePerLocationReport
	}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !IsNil(o.ReachabilityType) {
		toSerialize["reachabilityType"] = o.ReachabilityType
	}
	if !IsNil(o.RoamingStatus) {
		toSerialize["roamingStatus"] = o.RoamingStatus
	}
	if !IsNil(o.FailureCause) {
		toSerialize["failureCause"] = o.FailureCause
	}
	if !IsNil(o.EventTime) {
		toSerialize["eventTime"] = o.EventTime
	}
	if !IsNil(o.PdnConnInfoList) {
		toSerialize["pdnConnInfoList"] = o.PdnConnInfoList
	}
	if !IsNil(o.DddStatus) {
		toSerialize["dddStatus"] = o.DddStatus
	}
	if !IsNil(o.DddTrafDescriptor) {
		toSerialize["dddTrafDescriptor"] = o.DddTrafDescriptor
	}
	if !IsNil(o.MaxWaitTime) {
		toSerialize["maxWaitTime"] = o.MaxWaitTime
	}
	if !IsNil(o.ApiCaps) {
		toSerialize["apiCaps"] = o.ApiCaps
	}
	if !IsNil(o.NSStatusInfo) {
		toSerialize["nSStatusInfo"] = o.NSStatusInfo
	}
	if !IsNil(o.AfServiceId) {
		toSerialize["afServiceId"] = o.AfServiceId
	}
	if !IsNil(o.ServLevelDevId) {
		toSerialize["servLevelDevId"] = o.ServLevelDevId
	}
	if !IsNil(o.UavPresInd) {
		toSerialize["uavPresInd"] = o.UavPresInd
	}
	return toSerialize, nil
}

type NullableMonitoringEventReport struct {
	value *MonitoringEventReport
	isSet bool
}

func (v NullableMonitoringEventReport) Get() *MonitoringEventReport {
	return v.value
}

func (v *NullableMonitoringEventReport) Set(val *MonitoringEventReport) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringEventReport) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringEventReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringEventReport(val *MonitoringEventReport) *NullableMonitoringEventReport {
	return &NullableMonitoringEventReport{value: val, isSet: true}
}

func (v NullableMonitoringEventReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringEventReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


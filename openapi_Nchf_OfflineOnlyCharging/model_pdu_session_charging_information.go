/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_OfflineOnlyCharging

import (
	"encoding/json"
	"time"
)

// checks if the PDUSessionChargingInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PDUSessionChargingInformation{}

// PDUSessionChargingInformation struct for PDUSessionChargingInformation
type PDUSessionChargingInformation struct {
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	// Deprecated
	ChargingId *int32 `json:"chargingId,omitempty"`
	SMFChargingId *string `json:"sMFChargingId,omitempty"`
	UserInformation *UserInformation `json:"userInformation,omitempty"`
	UserLocationinfo *UserLocation `json:"userLocationinfo,omitempty"`
	MAPDUNon3GPPUserLocationInfo *UserLocation `json:"mAPDUNon3GPPUserLocationInfo,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	UserLocationTime *time.Time `json:"userLocationTime,omitempty"`
	PresenceReportingAreaInformation *map[string]PresenceInfo `json:"presenceReportingAreaInformation,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UetimeZone *string `json:"uetimeZone,omitempty"`
	PduSessionInformation PDUSessionInformation `json:"pduSessionInformation"`
	// indicating a time in seconds.
	UnitCountInactivityTimer *int32 `json:"unitCountInactivityTimer,omitempty"`
	RANSecondaryRATUsageReport *RANSecondaryRATUsageReport `json:"rANSecondaryRATUsageReport,omitempty"`
}

// NewPDUSessionChargingInformation instantiates a new PDUSessionChargingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPDUSessionChargingInformation(pduSessionInformation PDUSessionInformation) *PDUSessionChargingInformation {
	this := PDUSessionChargingInformation{}
	this.PduSessionInformation = pduSessionInformation
	return &this
}

// NewPDUSessionChargingInformationWithDefaults instantiates a new PDUSessionChargingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPDUSessionChargingInformationWithDefaults() *PDUSessionChargingInformation {
	this := PDUSessionChargingInformation{}
	return &this
}

// GetChargingId returns the ChargingId field value if set, zero value otherwise.
// Deprecated
func (o *PDUSessionChargingInformation) GetChargingId() int32 {
	if o == nil || IsNil(o.ChargingId) {
		var ret int32
		return ret
	}
	return *o.ChargingId
}

// GetChargingIdOk returns a tuple with the ChargingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PDUSessionChargingInformation) GetChargingIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ChargingId) {
		return nil, false
	}
	return o.ChargingId, true
}

// HasChargingId returns a boolean if a field has been set.
func (o *PDUSessionChargingInformation) HasChargingId() bool {
	if o != nil && !IsNil(o.ChargingId) {
		return true
	}

	return false
}

// SetChargingId gets a reference to the given int32 and assigns it to the ChargingId field.
// Deprecated
func (o *PDUSessionChargingInformation) SetChargingId(v int32) {
	o.ChargingId = &v
}

// GetSMFChargingId returns the SMFChargingId field value if set, zero value otherwise.
func (o *PDUSessionChargingInformation) GetSMFChargingId() string {
	if o == nil || IsNil(o.SMFChargingId) {
		var ret string
		return ret
	}
	return *o.SMFChargingId
}

// GetSMFChargingIdOk returns a tuple with the SMFChargingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUSessionChargingInformation) GetSMFChargingIdOk() (*string, bool) {
	if o == nil || IsNil(o.SMFChargingId) {
		return nil, false
	}
	return o.SMFChargingId, true
}

// HasSMFChargingId returns a boolean if a field has been set.
func (o *PDUSessionChargingInformation) HasSMFChargingId() bool {
	if o != nil && !IsNil(o.SMFChargingId) {
		return true
	}

	return false
}

// SetSMFChargingId gets a reference to the given string and assigns it to the SMFChargingId field.
func (o *PDUSessionChargingInformation) SetSMFChargingId(v string) {
	o.SMFChargingId = &v
}

// GetUserInformation returns the UserInformation field value if set, zero value otherwise.
func (o *PDUSessionChargingInformation) GetUserInformation() UserInformation {
	if o == nil || IsNil(o.UserInformation) {
		var ret UserInformation
		return ret
	}
	return *o.UserInformation
}

// GetUserInformationOk returns a tuple with the UserInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUSessionChargingInformation) GetUserInformationOk() (*UserInformation, bool) {
	if o == nil || IsNil(o.UserInformation) {
		return nil, false
	}
	return o.UserInformation, true
}

// HasUserInformation returns a boolean if a field has been set.
func (o *PDUSessionChargingInformation) HasUserInformation() bool {
	if o != nil && !IsNil(o.UserInformation) {
		return true
	}

	return false
}

// SetUserInformation gets a reference to the given UserInformation and assigns it to the UserInformation field.
func (o *PDUSessionChargingInformation) SetUserInformation(v UserInformation) {
	o.UserInformation = &v
}

// GetUserLocationinfo returns the UserLocationinfo field value if set, zero value otherwise.
func (o *PDUSessionChargingInformation) GetUserLocationinfo() UserLocation {
	if o == nil || IsNil(o.UserLocationinfo) {
		var ret UserLocation
		return ret
	}
	return *o.UserLocationinfo
}

// GetUserLocationinfoOk returns a tuple with the UserLocationinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUSessionChargingInformation) GetUserLocationinfoOk() (*UserLocation, bool) {
	if o == nil || IsNil(o.UserLocationinfo) {
		return nil, false
	}
	return o.UserLocationinfo, true
}

// HasUserLocationinfo returns a boolean if a field has been set.
func (o *PDUSessionChargingInformation) HasUserLocationinfo() bool {
	if o != nil && !IsNil(o.UserLocationinfo) {
		return true
	}

	return false
}

// SetUserLocationinfo gets a reference to the given UserLocation and assigns it to the UserLocationinfo field.
func (o *PDUSessionChargingInformation) SetUserLocationinfo(v UserLocation) {
	o.UserLocationinfo = &v
}

// GetMAPDUNon3GPPUserLocationInfo returns the MAPDUNon3GPPUserLocationInfo field value if set, zero value otherwise.
func (o *PDUSessionChargingInformation) GetMAPDUNon3GPPUserLocationInfo() UserLocation {
	if o == nil || IsNil(o.MAPDUNon3GPPUserLocationInfo) {
		var ret UserLocation
		return ret
	}
	return *o.MAPDUNon3GPPUserLocationInfo
}

// GetMAPDUNon3GPPUserLocationInfoOk returns a tuple with the MAPDUNon3GPPUserLocationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUSessionChargingInformation) GetMAPDUNon3GPPUserLocationInfoOk() (*UserLocation, bool) {
	if o == nil || IsNil(o.MAPDUNon3GPPUserLocationInfo) {
		return nil, false
	}
	return o.MAPDUNon3GPPUserLocationInfo, true
}

// HasMAPDUNon3GPPUserLocationInfo returns a boolean if a field has been set.
func (o *PDUSessionChargingInformation) HasMAPDUNon3GPPUserLocationInfo() bool {
	if o != nil && !IsNil(o.MAPDUNon3GPPUserLocationInfo) {
		return true
	}

	return false
}

// SetMAPDUNon3GPPUserLocationInfo gets a reference to the given UserLocation and assigns it to the MAPDUNon3GPPUserLocationInfo field.
func (o *PDUSessionChargingInformation) SetMAPDUNon3GPPUserLocationInfo(v UserLocation) {
	o.MAPDUNon3GPPUserLocationInfo = &v
}

// GetUserLocationTime returns the UserLocationTime field value if set, zero value otherwise.
func (o *PDUSessionChargingInformation) GetUserLocationTime() time.Time {
	if o == nil || IsNil(o.UserLocationTime) {
		var ret time.Time
		return ret
	}
	return *o.UserLocationTime
}

// GetUserLocationTimeOk returns a tuple with the UserLocationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUSessionChargingInformation) GetUserLocationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UserLocationTime) {
		return nil, false
	}
	return o.UserLocationTime, true
}

// HasUserLocationTime returns a boolean if a field has been set.
func (o *PDUSessionChargingInformation) HasUserLocationTime() bool {
	if o != nil && !IsNil(o.UserLocationTime) {
		return true
	}

	return false
}

// SetUserLocationTime gets a reference to the given time.Time and assigns it to the UserLocationTime field.
func (o *PDUSessionChargingInformation) SetUserLocationTime(v time.Time) {
	o.UserLocationTime = &v
}

// GetPresenceReportingAreaInformation returns the PresenceReportingAreaInformation field value if set, zero value otherwise.
func (o *PDUSessionChargingInformation) GetPresenceReportingAreaInformation() map[string]PresenceInfo {
	if o == nil || IsNil(o.PresenceReportingAreaInformation) {
		var ret map[string]PresenceInfo
		return ret
	}
	return *o.PresenceReportingAreaInformation
}

// GetPresenceReportingAreaInformationOk returns a tuple with the PresenceReportingAreaInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUSessionChargingInformation) GetPresenceReportingAreaInformationOk() (*map[string]PresenceInfo, bool) {
	if o == nil || IsNil(o.PresenceReportingAreaInformation) {
		return nil, false
	}
	return o.PresenceReportingAreaInformation, true
}

// HasPresenceReportingAreaInformation returns a boolean if a field has been set.
func (o *PDUSessionChargingInformation) HasPresenceReportingAreaInformation() bool {
	if o != nil && !IsNil(o.PresenceReportingAreaInformation) {
		return true
	}

	return false
}

// SetPresenceReportingAreaInformation gets a reference to the given map[string]PresenceInfo and assigns it to the PresenceReportingAreaInformation field.
func (o *PDUSessionChargingInformation) SetPresenceReportingAreaInformation(v map[string]PresenceInfo) {
	o.PresenceReportingAreaInformation = &v
}

// GetUetimeZone returns the UetimeZone field value if set, zero value otherwise.
func (o *PDUSessionChargingInformation) GetUetimeZone() string {
	if o == nil || IsNil(o.UetimeZone) {
		var ret string
		return ret
	}
	return *o.UetimeZone
}

// GetUetimeZoneOk returns a tuple with the UetimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUSessionChargingInformation) GetUetimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.UetimeZone) {
		return nil, false
	}
	return o.UetimeZone, true
}

// HasUetimeZone returns a boolean if a field has been set.
func (o *PDUSessionChargingInformation) HasUetimeZone() bool {
	if o != nil && !IsNil(o.UetimeZone) {
		return true
	}

	return false
}

// SetUetimeZone gets a reference to the given string and assigns it to the UetimeZone field.
func (o *PDUSessionChargingInformation) SetUetimeZone(v string) {
	o.UetimeZone = &v
}

// GetPduSessionInformation returns the PduSessionInformation field value
func (o *PDUSessionChargingInformation) GetPduSessionInformation() PDUSessionInformation {
	if o == nil {
		var ret PDUSessionInformation
		return ret
	}

	return o.PduSessionInformation
}

// GetPduSessionInformationOk returns a tuple with the PduSessionInformation field value
// and a boolean to check if the value has been set.
func (o *PDUSessionChargingInformation) GetPduSessionInformationOk() (*PDUSessionInformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PduSessionInformation, true
}

// SetPduSessionInformation sets field value
func (o *PDUSessionChargingInformation) SetPduSessionInformation(v PDUSessionInformation) {
	o.PduSessionInformation = v
}

// GetUnitCountInactivityTimer returns the UnitCountInactivityTimer field value if set, zero value otherwise.
func (o *PDUSessionChargingInformation) GetUnitCountInactivityTimer() int32 {
	if o == nil || IsNil(o.UnitCountInactivityTimer) {
		var ret int32
		return ret
	}
	return *o.UnitCountInactivityTimer
}

// GetUnitCountInactivityTimerOk returns a tuple with the UnitCountInactivityTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUSessionChargingInformation) GetUnitCountInactivityTimerOk() (*int32, bool) {
	if o == nil || IsNil(o.UnitCountInactivityTimer) {
		return nil, false
	}
	return o.UnitCountInactivityTimer, true
}

// HasUnitCountInactivityTimer returns a boolean if a field has been set.
func (o *PDUSessionChargingInformation) HasUnitCountInactivityTimer() bool {
	if o != nil && !IsNil(o.UnitCountInactivityTimer) {
		return true
	}

	return false
}

// SetUnitCountInactivityTimer gets a reference to the given int32 and assigns it to the UnitCountInactivityTimer field.
func (o *PDUSessionChargingInformation) SetUnitCountInactivityTimer(v int32) {
	o.UnitCountInactivityTimer = &v
}

// GetRANSecondaryRATUsageReport returns the RANSecondaryRATUsageReport field value if set, zero value otherwise.
func (o *PDUSessionChargingInformation) GetRANSecondaryRATUsageReport() RANSecondaryRATUsageReport {
	if o == nil || IsNil(o.RANSecondaryRATUsageReport) {
		var ret RANSecondaryRATUsageReport
		return ret
	}
	return *o.RANSecondaryRATUsageReport
}

// GetRANSecondaryRATUsageReportOk returns a tuple with the RANSecondaryRATUsageReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUSessionChargingInformation) GetRANSecondaryRATUsageReportOk() (*RANSecondaryRATUsageReport, bool) {
	if o == nil || IsNil(o.RANSecondaryRATUsageReport) {
		return nil, false
	}
	return o.RANSecondaryRATUsageReport, true
}

// HasRANSecondaryRATUsageReport returns a boolean if a field has been set.
func (o *PDUSessionChargingInformation) HasRANSecondaryRATUsageReport() bool {
	if o != nil && !IsNil(o.RANSecondaryRATUsageReport) {
		return true
	}

	return false
}

// SetRANSecondaryRATUsageReport gets a reference to the given RANSecondaryRATUsageReport and assigns it to the RANSecondaryRATUsageReport field.
func (o *PDUSessionChargingInformation) SetRANSecondaryRATUsageReport(v RANSecondaryRATUsageReport) {
	o.RANSecondaryRATUsageReport = &v
}

func (o PDUSessionChargingInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PDUSessionChargingInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChargingId) {
		toSerialize["chargingId"] = o.ChargingId
	}
	if !IsNil(o.SMFChargingId) {
		toSerialize["sMFChargingId"] = o.SMFChargingId
	}
	if !IsNil(o.UserInformation) {
		toSerialize["userInformation"] = o.UserInformation
	}
	if !IsNil(o.UserLocationinfo) {
		toSerialize["userLocationinfo"] = o.UserLocationinfo
	}
	if !IsNil(o.MAPDUNon3GPPUserLocationInfo) {
		toSerialize["mAPDUNon3GPPUserLocationInfo"] = o.MAPDUNon3GPPUserLocationInfo
	}
	if !IsNil(o.UserLocationTime) {
		toSerialize["userLocationTime"] = o.UserLocationTime
	}
	if !IsNil(o.PresenceReportingAreaInformation) {
		toSerialize["presenceReportingAreaInformation"] = o.PresenceReportingAreaInformation
	}
	if !IsNil(o.UetimeZone) {
		toSerialize["uetimeZone"] = o.UetimeZone
	}
	toSerialize["pduSessionInformation"] = o.PduSessionInformation
	if !IsNil(o.UnitCountInactivityTimer) {
		toSerialize["unitCountInactivityTimer"] = o.UnitCountInactivityTimer
	}
	if !IsNil(o.RANSecondaryRATUsageReport) {
		toSerialize["rANSecondaryRATUsageReport"] = o.RANSecondaryRATUsageReport
	}
	return toSerialize, nil
}

type NullablePDUSessionChargingInformation struct {
	value *PDUSessionChargingInformation
	isSet bool
}

func (v NullablePDUSessionChargingInformation) Get() *PDUSessionChargingInformation {
	return v.value
}

func (v *NullablePDUSessionChargingInformation) Set(val *PDUSessionChargingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePDUSessionChargingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePDUSessionChargingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePDUSessionChargingInformation(val *PDUSessionChargingInformation) *NullablePDUSessionChargingInformation {
	return &NullablePDUSessionChargingInformation{value: val, isSet: true}
}

func (v NullablePDUSessionChargingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePDUSessionChargingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



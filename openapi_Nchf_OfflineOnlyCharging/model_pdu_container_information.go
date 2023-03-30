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

// checks if the PDUContainerInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PDUContainerInformation{}

// PDUContainerInformation struct for PDUContainerInformation
type PDUContainerInformation struct {
	// string with format 'date-time' as defined in OpenAPI.
	TimeofFirstUsage *time.Time `json:"timeofFirstUsage,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeofLastUsage *time.Time `json:"timeofLastUsage,omitempty"`
	QoSInformation NullableQosData `json:"qoSInformation,omitempty"`
	QoSCharacteristics *QosCharacteristics `json:"qoSCharacteristics,omitempty"`
	AFCorrelationInformation *string `json:"aFCorrelationInformation,omitempty"`
	UserLocationInformation *UserLocation `json:"userLocationInformation,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UetimeZone *string `json:"uetimeZone,omitempty"`
	RATType *RatType `json:"rATType,omitempty"`
	ServingNodeID []ServingNetworkFunctionID `json:"servingNodeID,omitempty"`
	PresenceReportingAreaInformation *map[string]PresenceInfo `json:"presenceReportingAreaInformation,omitempty"`
	Var3gppPSDataOffStatus *Model3GPPPSDataOffStatus `json:"3gppPSDataOffStatus,omitempty"`
	SponsorIdentity *string `json:"sponsorIdentity,omitempty"`
	ApplicationserviceProviderIdentity *string `json:"applicationserviceProviderIdentity,omitempty"`
	ChargingRuleBaseName *string `json:"chargingRuleBaseName,omitempty"`
	MAPDUSteeringFunctionality *SteeringFunctionality `json:"mAPDUSteeringFunctionality,omitempty"`
	MAPDUSteeringMode *SteeringMode `json:"mAPDUSteeringMode,omitempty"`
}

// NewPDUContainerInformation instantiates a new PDUContainerInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPDUContainerInformation() *PDUContainerInformation {
	this := PDUContainerInformation{}
	return &this
}

// NewPDUContainerInformationWithDefaults instantiates a new PDUContainerInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPDUContainerInformationWithDefaults() *PDUContainerInformation {
	this := PDUContainerInformation{}
	return &this
}

// GetTimeofFirstUsage returns the TimeofFirstUsage field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetTimeofFirstUsage() time.Time {
	if o == nil || IsNil(o.TimeofFirstUsage) {
		var ret time.Time
		return ret
	}
	return *o.TimeofFirstUsage
}

// GetTimeofFirstUsageOk returns a tuple with the TimeofFirstUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetTimeofFirstUsageOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeofFirstUsage) {
		return nil, false
	}
	return o.TimeofFirstUsage, true
}

// HasTimeofFirstUsage returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasTimeofFirstUsage() bool {
	if o != nil && !IsNil(o.TimeofFirstUsage) {
		return true
	}

	return false
}

// SetTimeofFirstUsage gets a reference to the given time.Time and assigns it to the TimeofFirstUsage field.
func (o *PDUContainerInformation) SetTimeofFirstUsage(v time.Time) {
	o.TimeofFirstUsage = &v
}

// GetTimeofLastUsage returns the TimeofLastUsage field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetTimeofLastUsage() time.Time {
	if o == nil || IsNil(o.TimeofLastUsage) {
		var ret time.Time
		return ret
	}
	return *o.TimeofLastUsage
}

// GetTimeofLastUsageOk returns a tuple with the TimeofLastUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetTimeofLastUsageOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeofLastUsage) {
		return nil, false
	}
	return o.TimeofLastUsage, true
}

// HasTimeofLastUsage returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasTimeofLastUsage() bool {
	if o != nil && !IsNil(o.TimeofLastUsage) {
		return true
	}

	return false
}

// SetTimeofLastUsage gets a reference to the given time.Time and assigns it to the TimeofLastUsage field.
func (o *PDUContainerInformation) SetTimeofLastUsage(v time.Time) {
	o.TimeofLastUsage = &v
}

// GetQoSInformation returns the QoSInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PDUContainerInformation) GetQoSInformation() QosData {
	if o == nil || IsNil(o.QoSInformation.Get()) {
		var ret QosData
		return ret
	}
	return *o.QoSInformation.Get()
}

// GetQoSInformationOk returns a tuple with the QoSInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PDUContainerInformation) GetQoSInformationOk() (*QosData, bool) {
	if o == nil {
		return nil, false
	}
	return o.QoSInformation.Get(), o.QoSInformation.IsSet()
}

// HasQoSInformation returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasQoSInformation() bool {
	if o != nil && o.QoSInformation.IsSet() {
		return true
	}

	return false
}

// SetQoSInformation gets a reference to the given NullableQosData and assigns it to the QoSInformation field.
func (o *PDUContainerInformation) SetQoSInformation(v QosData) {
	o.QoSInformation.Set(&v)
}
// SetQoSInformationNil sets the value for QoSInformation to be an explicit nil
func (o *PDUContainerInformation) SetQoSInformationNil() {
	o.QoSInformation.Set(nil)
}

// UnsetQoSInformation ensures that no value is present for QoSInformation, not even an explicit nil
func (o *PDUContainerInformation) UnsetQoSInformation() {
	o.QoSInformation.Unset()
}

// GetQoSCharacteristics returns the QoSCharacteristics field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetQoSCharacteristics() QosCharacteristics {
	if o == nil || IsNil(o.QoSCharacteristics) {
		var ret QosCharacteristics
		return ret
	}
	return *o.QoSCharacteristics
}

// GetQoSCharacteristicsOk returns a tuple with the QoSCharacteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetQoSCharacteristicsOk() (*QosCharacteristics, bool) {
	if o == nil || IsNil(o.QoSCharacteristics) {
		return nil, false
	}
	return o.QoSCharacteristics, true
}

// HasQoSCharacteristics returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasQoSCharacteristics() bool {
	if o != nil && !IsNil(o.QoSCharacteristics) {
		return true
	}

	return false
}

// SetQoSCharacteristics gets a reference to the given QosCharacteristics and assigns it to the QoSCharacteristics field.
func (o *PDUContainerInformation) SetQoSCharacteristics(v QosCharacteristics) {
	o.QoSCharacteristics = &v
}

// GetAFCorrelationInformation returns the AFCorrelationInformation field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetAFCorrelationInformation() string {
	if o == nil || IsNil(o.AFCorrelationInformation) {
		var ret string
		return ret
	}
	return *o.AFCorrelationInformation
}

// GetAFCorrelationInformationOk returns a tuple with the AFCorrelationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetAFCorrelationInformationOk() (*string, bool) {
	if o == nil || IsNil(o.AFCorrelationInformation) {
		return nil, false
	}
	return o.AFCorrelationInformation, true
}

// HasAFCorrelationInformation returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasAFCorrelationInformation() bool {
	if o != nil && !IsNil(o.AFCorrelationInformation) {
		return true
	}

	return false
}

// SetAFCorrelationInformation gets a reference to the given string and assigns it to the AFCorrelationInformation field.
func (o *PDUContainerInformation) SetAFCorrelationInformation(v string) {
	o.AFCorrelationInformation = &v
}

// GetUserLocationInformation returns the UserLocationInformation field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetUserLocationInformation() UserLocation {
	if o == nil || IsNil(o.UserLocationInformation) {
		var ret UserLocation
		return ret
	}
	return *o.UserLocationInformation
}

// GetUserLocationInformationOk returns a tuple with the UserLocationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetUserLocationInformationOk() (*UserLocation, bool) {
	if o == nil || IsNil(o.UserLocationInformation) {
		return nil, false
	}
	return o.UserLocationInformation, true
}

// HasUserLocationInformation returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasUserLocationInformation() bool {
	if o != nil && !IsNil(o.UserLocationInformation) {
		return true
	}

	return false
}

// SetUserLocationInformation gets a reference to the given UserLocation and assigns it to the UserLocationInformation field.
func (o *PDUContainerInformation) SetUserLocationInformation(v UserLocation) {
	o.UserLocationInformation = &v
}

// GetUetimeZone returns the UetimeZone field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetUetimeZone() string {
	if o == nil || IsNil(o.UetimeZone) {
		var ret string
		return ret
	}
	return *o.UetimeZone
}

// GetUetimeZoneOk returns a tuple with the UetimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetUetimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.UetimeZone) {
		return nil, false
	}
	return o.UetimeZone, true
}

// HasUetimeZone returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasUetimeZone() bool {
	if o != nil && !IsNil(o.UetimeZone) {
		return true
	}

	return false
}

// SetUetimeZone gets a reference to the given string and assigns it to the UetimeZone field.
func (o *PDUContainerInformation) SetUetimeZone(v string) {
	o.UetimeZone = &v
}

// GetRATType returns the RATType field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetRATType() RatType {
	if o == nil || IsNil(o.RATType) {
		var ret RatType
		return ret
	}
	return *o.RATType
}

// GetRATTypeOk returns a tuple with the RATType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetRATTypeOk() (*RatType, bool) {
	if o == nil || IsNil(o.RATType) {
		return nil, false
	}
	return o.RATType, true
}

// HasRATType returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasRATType() bool {
	if o != nil && !IsNil(o.RATType) {
		return true
	}

	return false
}

// SetRATType gets a reference to the given RatType and assigns it to the RATType field.
func (o *PDUContainerInformation) SetRATType(v RatType) {
	o.RATType = &v
}

// GetServingNodeID returns the ServingNodeID field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetServingNodeID() []ServingNetworkFunctionID {
	if o == nil || IsNil(o.ServingNodeID) {
		var ret []ServingNetworkFunctionID
		return ret
	}
	return o.ServingNodeID
}

// GetServingNodeIDOk returns a tuple with the ServingNodeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetServingNodeIDOk() ([]ServingNetworkFunctionID, bool) {
	if o == nil || IsNil(o.ServingNodeID) {
		return nil, false
	}
	return o.ServingNodeID, true
}

// HasServingNodeID returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasServingNodeID() bool {
	if o != nil && !IsNil(o.ServingNodeID) {
		return true
	}

	return false
}

// SetServingNodeID gets a reference to the given []ServingNetworkFunctionID and assigns it to the ServingNodeID field.
func (o *PDUContainerInformation) SetServingNodeID(v []ServingNetworkFunctionID) {
	o.ServingNodeID = v
}

// GetPresenceReportingAreaInformation returns the PresenceReportingAreaInformation field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetPresenceReportingAreaInformation() map[string]PresenceInfo {
	if o == nil || IsNil(o.PresenceReportingAreaInformation) {
		var ret map[string]PresenceInfo
		return ret
	}
	return *o.PresenceReportingAreaInformation
}

// GetPresenceReportingAreaInformationOk returns a tuple with the PresenceReportingAreaInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetPresenceReportingAreaInformationOk() (*map[string]PresenceInfo, bool) {
	if o == nil || IsNil(o.PresenceReportingAreaInformation) {
		return nil, false
	}
	return o.PresenceReportingAreaInformation, true
}

// HasPresenceReportingAreaInformation returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasPresenceReportingAreaInformation() bool {
	if o != nil && !IsNil(o.PresenceReportingAreaInformation) {
		return true
	}

	return false
}

// SetPresenceReportingAreaInformation gets a reference to the given map[string]PresenceInfo and assigns it to the PresenceReportingAreaInformation field.
func (o *PDUContainerInformation) SetPresenceReportingAreaInformation(v map[string]PresenceInfo) {
	o.PresenceReportingAreaInformation = &v
}

// GetVar3gppPSDataOffStatus returns the Var3gppPSDataOffStatus field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetVar3gppPSDataOffStatus() Model3GPPPSDataOffStatus {
	if o == nil || IsNil(o.Var3gppPSDataOffStatus) {
		var ret Model3GPPPSDataOffStatus
		return ret
	}
	return *o.Var3gppPSDataOffStatus
}

// GetVar3gppPSDataOffStatusOk returns a tuple with the Var3gppPSDataOffStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetVar3gppPSDataOffStatusOk() (*Model3GPPPSDataOffStatus, bool) {
	if o == nil || IsNil(o.Var3gppPSDataOffStatus) {
		return nil, false
	}
	return o.Var3gppPSDataOffStatus, true
}

// HasVar3gppPSDataOffStatus returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasVar3gppPSDataOffStatus() bool {
	if o != nil && !IsNil(o.Var3gppPSDataOffStatus) {
		return true
	}

	return false
}

// SetVar3gppPSDataOffStatus gets a reference to the given Model3GPPPSDataOffStatus and assigns it to the Var3gppPSDataOffStatus field.
func (o *PDUContainerInformation) SetVar3gppPSDataOffStatus(v Model3GPPPSDataOffStatus) {
	o.Var3gppPSDataOffStatus = &v
}

// GetSponsorIdentity returns the SponsorIdentity field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetSponsorIdentity() string {
	if o == nil || IsNil(o.SponsorIdentity) {
		var ret string
		return ret
	}
	return *o.SponsorIdentity
}

// GetSponsorIdentityOk returns a tuple with the SponsorIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetSponsorIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.SponsorIdentity) {
		return nil, false
	}
	return o.SponsorIdentity, true
}

// HasSponsorIdentity returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasSponsorIdentity() bool {
	if o != nil && !IsNil(o.SponsorIdentity) {
		return true
	}

	return false
}

// SetSponsorIdentity gets a reference to the given string and assigns it to the SponsorIdentity field.
func (o *PDUContainerInformation) SetSponsorIdentity(v string) {
	o.SponsorIdentity = &v
}

// GetApplicationserviceProviderIdentity returns the ApplicationserviceProviderIdentity field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetApplicationserviceProviderIdentity() string {
	if o == nil || IsNil(o.ApplicationserviceProviderIdentity) {
		var ret string
		return ret
	}
	return *o.ApplicationserviceProviderIdentity
}

// GetApplicationserviceProviderIdentityOk returns a tuple with the ApplicationserviceProviderIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetApplicationserviceProviderIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationserviceProviderIdentity) {
		return nil, false
	}
	return o.ApplicationserviceProviderIdentity, true
}

// HasApplicationserviceProviderIdentity returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasApplicationserviceProviderIdentity() bool {
	if o != nil && !IsNil(o.ApplicationserviceProviderIdentity) {
		return true
	}

	return false
}

// SetApplicationserviceProviderIdentity gets a reference to the given string and assigns it to the ApplicationserviceProviderIdentity field.
func (o *PDUContainerInformation) SetApplicationserviceProviderIdentity(v string) {
	o.ApplicationserviceProviderIdentity = &v
}

// GetChargingRuleBaseName returns the ChargingRuleBaseName field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetChargingRuleBaseName() string {
	if o == nil || IsNil(o.ChargingRuleBaseName) {
		var ret string
		return ret
	}
	return *o.ChargingRuleBaseName
}

// GetChargingRuleBaseNameOk returns a tuple with the ChargingRuleBaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetChargingRuleBaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChargingRuleBaseName) {
		return nil, false
	}
	return o.ChargingRuleBaseName, true
}

// HasChargingRuleBaseName returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasChargingRuleBaseName() bool {
	if o != nil && !IsNil(o.ChargingRuleBaseName) {
		return true
	}

	return false
}

// SetChargingRuleBaseName gets a reference to the given string and assigns it to the ChargingRuleBaseName field.
func (o *PDUContainerInformation) SetChargingRuleBaseName(v string) {
	o.ChargingRuleBaseName = &v
}

// GetMAPDUSteeringFunctionality returns the MAPDUSteeringFunctionality field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetMAPDUSteeringFunctionality() SteeringFunctionality {
	if o == nil || IsNil(o.MAPDUSteeringFunctionality) {
		var ret SteeringFunctionality
		return ret
	}
	return *o.MAPDUSteeringFunctionality
}

// GetMAPDUSteeringFunctionalityOk returns a tuple with the MAPDUSteeringFunctionality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetMAPDUSteeringFunctionalityOk() (*SteeringFunctionality, bool) {
	if o == nil || IsNil(o.MAPDUSteeringFunctionality) {
		return nil, false
	}
	return o.MAPDUSteeringFunctionality, true
}

// HasMAPDUSteeringFunctionality returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasMAPDUSteeringFunctionality() bool {
	if o != nil && !IsNil(o.MAPDUSteeringFunctionality) {
		return true
	}

	return false
}

// SetMAPDUSteeringFunctionality gets a reference to the given SteeringFunctionality and assigns it to the MAPDUSteeringFunctionality field.
func (o *PDUContainerInformation) SetMAPDUSteeringFunctionality(v SteeringFunctionality) {
	o.MAPDUSteeringFunctionality = &v
}

// GetMAPDUSteeringMode returns the MAPDUSteeringMode field value if set, zero value otherwise.
func (o *PDUContainerInformation) GetMAPDUSteeringMode() SteeringMode {
	if o == nil || IsNil(o.MAPDUSteeringMode) {
		var ret SteeringMode
		return ret
	}
	return *o.MAPDUSteeringMode
}

// GetMAPDUSteeringModeOk returns a tuple with the MAPDUSteeringMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUContainerInformation) GetMAPDUSteeringModeOk() (*SteeringMode, bool) {
	if o == nil || IsNil(o.MAPDUSteeringMode) {
		return nil, false
	}
	return o.MAPDUSteeringMode, true
}

// HasMAPDUSteeringMode returns a boolean if a field has been set.
func (o *PDUContainerInformation) HasMAPDUSteeringMode() bool {
	if o != nil && !IsNil(o.MAPDUSteeringMode) {
		return true
	}

	return false
}

// SetMAPDUSteeringMode gets a reference to the given SteeringMode and assigns it to the MAPDUSteeringMode field.
func (o *PDUContainerInformation) SetMAPDUSteeringMode(v SteeringMode) {
	o.MAPDUSteeringMode = &v
}

func (o PDUContainerInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PDUContainerInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeofFirstUsage) {
		toSerialize["timeofFirstUsage"] = o.TimeofFirstUsage
	}
	if !IsNil(o.TimeofLastUsage) {
		toSerialize["timeofLastUsage"] = o.TimeofLastUsage
	}
	if o.QoSInformation.IsSet() {
		toSerialize["qoSInformation"] = o.QoSInformation.Get()
	}
	if !IsNil(o.QoSCharacteristics) {
		toSerialize["qoSCharacteristics"] = o.QoSCharacteristics
	}
	if !IsNil(o.AFCorrelationInformation) {
		toSerialize["aFCorrelationInformation"] = o.AFCorrelationInformation
	}
	if !IsNil(o.UserLocationInformation) {
		toSerialize["userLocationInformation"] = o.UserLocationInformation
	}
	if !IsNil(o.UetimeZone) {
		toSerialize["uetimeZone"] = o.UetimeZone
	}
	if !IsNil(o.RATType) {
		toSerialize["rATType"] = o.RATType
	}
	if !IsNil(o.ServingNodeID) {
		toSerialize["servingNodeID"] = o.ServingNodeID
	}
	if !IsNil(o.PresenceReportingAreaInformation) {
		toSerialize["presenceReportingAreaInformation"] = o.PresenceReportingAreaInformation
	}
	if !IsNil(o.Var3gppPSDataOffStatus) {
		toSerialize["3gppPSDataOffStatus"] = o.Var3gppPSDataOffStatus
	}
	if !IsNil(o.SponsorIdentity) {
		toSerialize["sponsorIdentity"] = o.SponsorIdentity
	}
	if !IsNil(o.ApplicationserviceProviderIdentity) {
		toSerialize["applicationserviceProviderIdentity"] = o.ApplicationserviceProviderIdentity
	}
	if !IsNil(o.ChargingRuleBaseName) {
		toSerialize["chargingRuleBaseName"] = o.ChargingRuleBaseName
	}
	if !IsNil(o.MAPDUSteeringFunctionality) {
		toSerialize["mAPDUSteeringFunctionality"] = o.MAPDUSteeringFunctionality
	}
	if !IsNil(o.MAPDUSteeringMode) {
		toSerialize["mAPDUSteeringMode"] = o.MAPDUSteeringMode
	}
	return toSerialize, nil
}

type NullablePDUContainerInformation struct {
	value *PDUContainerInformation
	isSet bool
}

func (v NullablePDUContainerInformation) Get() *PDUContainerInformation {
	return v.value
}

func (v *NullablePDUContainerInformation) Set(val *PDUContainerInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePDUContainerInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePDUContainerInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePDUContainerInformation(val *PDUContainerInformation) *NullablePDUContainerInformation {
	return &NullablePDUContainerInformation{value: val, isSet: true}
}

func (v NullablePDUContainerInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePDUContainerInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



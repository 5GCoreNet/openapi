/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Location

import (
	"encoding/json"
	"time"
)

// checks if the NotifiedPosInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifiedPosInfo{}

// NotifiedPosInfo Data within EventNotify notification
type NotifiedPosInfo struct {
	LocationEvent LocationEvent `json:"locationEvent"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi *string `json:"supi,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi *string `json:"gpsi,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	Pei                   *string         `json:"pei,omitempty"`
	LocationEstimate      *GeographicArea `json:"locationEstimate,omitempty"`
	LocalLocationEstimate *LocalArea      `json:"localLocationEstimate,omitempty"`
	// Indicates value of the age of the location estimate.
	AgeOfLocationEstimate *int32 `json:"ageOfLocationEstimate,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimestampOfLocationEstimate *time.Time                      `json:"timestampOfLocationEstimate,omitempty"`
	VelocityEstimate            *VelocityEstimate               `json:"velocityEstimate,omitempty"`
	PositioningDataList         []PositioningMethodAndUsage     `json:"positioningDataList,omitempty"`
	GnssPositioningDataList     []GnssPositioningMethodAndUsage `json:"gnssPositioningDataList,omitempty"`
	Ecgi                        *Ecgi                           `json:"ecgi,omitempty"`
	Ncgi                        *Ncgi                           `json:"ncgi,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	ServingNode *string `json:"servingNode,omitempty"`
	// Fully Qualified Domain Name
	TargetMmeName *string `json:"targetMmeName,omitempty"`
	// Fully Qualified Domain Name
	TargetMmeRealm *string       `json:"targetMmeRealm,omitempty"`
	UtranSrvccInd  *bool         `json:"utranSrvccInd,omitempty"`
	CivicAddress   *CivicAddress `json:"civicAddress,omitempty"`
	// Specifies the measured uncompensated atmospheric pressure.
	BarometricPressure *int32 `json:"barometricPressure,omitempty"`
	// Indicates value of altitude.
	Altitude *float64 `json:"altitude,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	HgmlcCallBackURI *string `json:"hgmlcCallBackURI,omitempty"`
	// LDR Reference.
	LdrReference *string `json:"ldrReference,omitempty"`
	// LMF identification.
	ServingLMFIdentification *string           `json:"servingLMFIdentification,omitempty"`
	TerminationCause         *TerminationCause `json:"terminationCause,omitempty"`
	AchievedQos              *MinorLocationQoS `json:"achievedQos,omitempty"`
	MscServerId              *string           `json:"mscServerId,omitempty"`
}

// NewNotifiedPosInfo instantiates a new NotifiedPosInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifiedPosInfo(locationEvent LocationEvent) *NotifiedPosInfo {
	this := NotifiedPosInfo{}
	this.LocationEvent = locationEvent
	return &this
}

// NewNotifiedPosInfoWithDefaults instantiates a new NotifiedPosInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifiedPosInfoWithDefaults() *NotifiedPosInfo {
	this := NotifiedPosInfo{}
	return &this
}

// GetLocationEvent returns the LocationEvent field value
func (o *NotifiedPosInfo) GetLocationEvent() LocationEvent {
	if o == nil {
		var ret LocationEvent
		return ret
	}

	return o.LocationEvent
}

// GetLocationEventOk returns a tuple with the LocationEvent field value
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetLocationEventOk() (*LocationEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationEvent, true
}

// SetLocationEvent sets field value
func (o *NotifiedPosInfo) SetLocationEvent(v LocationEvent) {
	o.LocationEvent = v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *NotifiedPosInfo) SetSupi(v string) {
	o.Supi = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *NotifiedPosInfo) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetPei returns the Pei field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetPei() string {
	if o == nil || IsNil(o.Pei) {
		var ret string
		return ret
	}
	return *o.Pei
}

// GetPeiOk returns a tuple with the Pei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetPeiOk() (*string, bool) {
	if o == nil || IsNil(o.Pei) {
		return nil, false
	}
	return o.Pei, true
}

// HasPei returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasPei() bool {
	if o != nil && !IsNil(o.Pei) {
		return true
	}

	return false
}

// SetPei gets a reference to the given string and assigns it to the Pei field.
func (o *NotifiedPosInfo) SetPei(v string) {
	o.Pei = &v
}

// GetLocationEstimate returns the LocationEstimate field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetLocationEstimate() GeographicArea {
	if o == nil || IsNil(o.LocationEstimate) {
		var ret GeographicArea
		return ret
	}
	return *o.LocationEstimate
}

// GetLocationEstimateOk returns a tuple with the LocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetLocationEstimateOk() (*GeographicArea, bool) {
	if o == nil || IsNil(o.LocationEstimate) {
		return nil, false
	}
	return o.LocationEstimate, true
}

// HasLocationEstimate returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasLocationEstimate() bool {
	if o != nil && !IsNil(o.LocationEstimate) {
		return true
	}

	return false
}

// SetLocationEstimate gets a reference to the given GeographicArea and assigns it to the LocationEstimate field.
func (o *NotifiedPosInfo) SetLocationEstimate(v GeographicArea) {
	o.LocationEstimate = &v
}

// GetLocalLocationEstimate returns the LocalLocationEstimate field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetLocalLocationEstimate() LocalArea {
	if o == nil || IsNil(o.LocalLocationEstimate) {
		var ret LocalArea
		return ret
	}
	return *o.LocalLocationEstimate
}

// GetLocalLocationEstimateOk returns a tuple with the LocalLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetLocalLocationEstimateOk() (*LocalArea, bool) {
	if o == nil || IsNil(o.LocalLocationEstimate) {
		return nil, false
	}
	return o.LocalLocationEstimate, true
}

// HasLocalLocationEstimate returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasLocalLocationEstimate() bool {
	if o != nil && !IsNil(o.LocalLocationEstimate) {
		return true
	}

	return false
}

// SetLocalLocationEstimate gets a reference to the given LocalArea and assigns it to the LocalLocationEstimate field.
func (o *NotifiedPosInfo) SetLocalLocationEstimate(v LocalArea) {
	o.LocalLocationEstimate = &v
}

// GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetAgeOfLocationEstimate() int32 {
	if o == nil || IsNil(o.AgeOfLocationEstimate) {
		var ret int32
		return ret
	}
	return *o.AgeOfLocationEstimate
}

// GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetAgeOfLocationEstimateOk() (*int32, bool) {
	if o == nil || IsNil(o.AgeOfLocationEstimate) {
		return nil, false
	}
	return o.AgeOfLocationEstimate, true
}

// HasAgeOfLocationEstimate returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasAgeOfLocationEstimate() bool {
	if o != nil && !IsNil(o.AgeOfLocationEstimate) {
		return true
	}

	return false
}

// SetAgeOfLocationEstimate gets a reference to the given int32 and assigns it to the AgeOfLocationEstimate field.
func (o *NotifiedPosInfo) SetAgeOfLocationEstimate(v int32) {
	o.AgeOfLocationEstimate = &v
}

// GetTimestampOfLocationEstimate returns the TimestampOfLocationEstimate field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetTimestampOfLocationEstimate() time.Time {
	if o == nil || IsNil(o.TimestampOfLocationEstimate) {
		var ret time.Time
		return ret
	}
	return *o.TimestampOfLocationEstimate
}

// GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetTimestampOfLocationEstimateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampOfLocationEstimate) {
		return nil, false
	}
	return o.TimestampOfLocationEstimate, true
}

// HasTimestampOfLocationEstimate returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasTimestampOfLocationEstimate() bool {
	if o != nil && !IsNil(o.TimestampOfLocationEstimate) {
		return true
	}

	return false
}

// SetTimestampOfLocationEstimate gets a reference to the given time.Time and assigns it to the TimestampOfLocationEstimate field.
func (o *NotifiedPosInfo) SetTimestampOfLocationEstimate(v time.Time) {
	o.TimestampOfLocationEstimate = &v
}

// GetVelocityEstimate returns the VelocityEstimate field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetVelocityEstimate() VelocityEstimate {
	if o == nil || IsNil(o.VelocityEstimate) {
		var ret VelocityEstimate
		return ret
	}
	return *o.VelocityEstimate
}

// GetVelocityEstimateOk returns a tuple with the VelocityEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetVelocityEstimateOk() (*VelocityEstimate, bool) {
	if o == nil || IsNil(o.VelocityEstimate) {
		return nil, false
	}
	return o.VelocityEstimate, true
}

// HasVelocityEstimate returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasVelocityEstimate() bool {
	if o != nil && !IsNil(o.VelocityEstimate) {
		return true
	}

	return false
}

// SetVelocityEstimate gets a reference to the given VelocityEstimate and assigns it to the VelocityEstimate field.
func (o *NotifiedPosInfo) SetVelocityEstimate(v VelocityEstimate) {
	o.VelocityEstimate = &v
}

// GetPositioningDataList returns the PositioningDataList field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetPositioningDataList() []PositioningMethodAndUsage {
	if o == nil || IsNil(o.PositioningDataList) {
		var ret []PositioningMethodAndUsage
		return ret
	}
	return o.PositioningDataList
}

// GetPositioningDataListOk returns a tuple with the PositioningDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetPositioningDataListOk() ([]PositioningMethodAndUsage, bool) {
	if o == nil || IsNil(o.PositioningDataList) {
		return nil, false
	}
	return o.PositioningDataList, true
}

// HasPositioningDataList returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasPositioningDataList() bool {
	if o != nil && !IsNil(o.PositioningDataList) {
		return true
	}

	return false
}

// SetPositioningDataList gets a reference to the given []PositioningMethodAndUsage and assigns it to the PositioningDataList field.
func (o *NotifiedPosInfo) SetPositioningDataList(v []PositioningMethodAndUsage) {
	o.PositioningDataList = v
}

// GetGnssPositioningDataList returns the GnssPositioningDataList field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetGnssPositioningDataList() []GnssPositioningMethodAndUsage {
	if o == nil || IsNil(o.GnssPositioningDataList) {
		var ret []GnssPositioningMethodAndUsage
		return ret
	}
	return o.GnssPositioningDataList
}

// GetGnssPositioningDataListOk returns a tuple with the GnssPositioningDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetGnssPositioningDataListOk() ([]GnssPositioningMethodAndUsage, bool) {
	if o == nil || IsNil(o.GnssPositioningDataList) {
		return nil, false
	}
	return o.GnssPositioningDataList, true
}

// HasGnssPositioningDataList returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasGnssPositioningDataList() bool {
	if o != nil && !IsNil(o.GnssPositioningDataList) {
		return true
	}

	return false
}

// SetGnssPositioningDataList gets a reference to the given []GnssPositioningMethodAndUsage and assigns it to the GnssPositioningDataList field.
func (o *NotifiedPosInfo) SetGnssPositioningDataList(v []GnssPositioningMethodAndUsage) {
	o.GnssPositioningDataList = v
}

// GetEcgi returns the Ecgi field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetEcgi() Ecgi {
	if o == nil || IsNil(o.Ecgi) {
		var ret Ecgi
		return ret
	}
	return *o.Ecgi
}

// GetEcgiOk returns a tuple with the Ecgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetEcgiOk() (*Ecgi, bool) {
	if o == nil || IsNil(o.Ecgi) {
		return nil, false
	}
	return o.Ecgi, true
}

// HasEcgi returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasEcgi() bool {
	if o != nil && !IsNil(o.Ecgi) {
		return true
	}

	return false
}

// SetEcgi gets a reference to the given Ecgi and assigns it to the Ecgi field.
func (o *NotifiedPosInfo) SetEcgi(v Ecgi) {
	o.Ecgi = &v
}

// GetNcgi returns the Ncgi field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetNcgi() Ncgi {
	if o == nil || IsNil(o.Ncgi) {
		var ret Ncgi
		return ret
	}
	return *o.Ncgi
}

// GetNcgiOk returns a tuple with the Ncgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetNcgiOk() (*Ncgi, bool) {
	if o == nil || IsNil(o.Ncgi) {
		return nil, false
	}
	return o.Ncgi, true
}

// HasNcgi returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasNcgi() bool {
	if o != nil && !IsNil(o.Ncgi) {
		return true
	}

	return false
}

// SetNcgi gets a reference to the given Ncgi and assigns it to the Ncgi field.
func (o *NotifiedPosInfo) SetNcgi(v Ncgi) {
	o.Ncgi = &v
}

// GetServingNode returns the ServingNode field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetServingNode() string {
	if o == nil || IsNil(o.ServingNode) {
		var ret string
		return ret
	}
	return *o.ServingNode
}

// GetServingNodeOk returns a tuple with the ServingNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetServingNodeOk() (*string, bool) {
	if o == nil || IsNil(o.ServingNode) {
		return nil, false
	}
	return o.ServingNode, true
}

// HasServingNode returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasServingNode() bool {
	if o != nil && !IsNil(o.ServingNode) {
		return true
	}

	return false
}

// SetServingNode gets a reference to the given string and assigns it to the ServingNode field.
func (o *NotifiedPosInfo) SetServingNode(v string) {
	o.ServingNode = &v
}

// GetTargetMmeName returns the TargetMmeName field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetTargetMmeName() string {
	if o == nil || IsNil(o.TargetMmeName) {
		var ret string
		return ret
	}
	return *o.TargetMmeName
}

// GetTargetMmeNameOk returns a tuple with the TargetMmeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetTargetMmeNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetMmeName) {
		return nil, false
	}
	return o.TargetMmeName, true
}

// HasTargetMmeName returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasTargetMmeName() bool {
	if o != nil && !IsNil(o.TargetMmeName) {
		return true
	}

	return false
}

// SetTargetMmeName gets a reference to the given string and assigns it to the TargetMmeName field.
func (o *NotifiedPosInfo) SetTargetMmeName(v string) {
	o.TargetMmeName = &v
}

// GetTargetMmeRealm returns the TargetMmeRealm field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetTargetMmeRealm() string {
	if o == nil || IsNil(o.TargetMmeRealm) {
		var ret string
		return ret
	}
	return *o.TargetMmeRealm
}

// GetTargetMmeRealmOk returns a tuple with the TargetMmeRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetTargetMmeRealmOk() (*string, bool) {
	if o == nil || IsNil(o.TargetMmeRealm) {
		return nil, false
	}
	return o.TargetMmeRealm, true
}

// HasTargetMmeRealm returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasTargetMmeRealm() bool {
	if o != nil && !IsNil(o.TargetMmeRealm) {
		return true
	}

	return false
}

// SetTargetMmeRealm gets a reference to the given string and assigns it to the TargetMmeRealm field.
func (o *NotifiedPosInfo) SetTargetMmeRealm(v string) {
	o.TargetMmeRealm = &v
}

// GetUtranSrvccInd returns the UtranSrvccInd field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetUtranSrvccInd() bool {
	if o == nil || IsNil(o.UtranSrvccInd) {
		var ret bool
		return ret
	}
	return *o.UtranSrvccInd
}

// GetUtranSrvccIndOk returns a tuple with the UtranSrvccInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetUtranSrvccIndOk() (*bool, bool) {
	if o == nil || IsNil(o.UtranSrvccInd) {
		return nil, false
	}
	return o.UtranSrvccInd, true
}

// HasUtranSrvccInd returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasUtranSrvccInd() bool {
	if o != nil && !IsNil(o.UtranSrvccInd) {
		return true
	}

	return false
}

// SetUtranSrvccInd gets a reference to the given bool and assigns it to the UtranSrvccInd field.
func (o *NotifiedPosInfo) SetUtranSrvccInd(v bool) {
	o.UtranSrvccInd = &v
}

// GetCivicAddress returns the CivicAddress field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetCivicAddress() CivicAddress {
	if o == nil || IsNil(o.CivicAddress) {
		var ret CivicAddress
		return ret
	}
	return *o.CivicAddress
}

// GetCivicAddressOk returns a tuple with the CivicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetCivicAddressOk() (*CivicAddress, bool) {
	if o == nil || IsNil(o.CivicAddress) {
		return nil, false
	}
	return o.CivicAddress, true
}

// HasCivicAddress returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasCivicAddress() bool {
	if o != nil && !IsNil(o.CivicAddress) {
		return true
	}

	return false
}

// SetCivicAddress gets a reference to the given CivicAddress and assigns it to the CivicAddress field.
func (o *NotifiedPosInfo) SetCivicAddress(v CivicAddress) {
	o.CivicAddress = &v
}

// GetBarometricPressure returns the BarometricPressure field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetBarometricPressure() int32 {
	if o == nil || IsNil(o.BarometricPressure) {
		var ret int32
		return ret
	}
	return *o.BarometricPressure
}

// GetBarometricPressureOk returns a tuple with the BarometricPressure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetBarometricPressureOk() (*int32, bool) {
	if o == nil || IsNil(o.BarometricPressure) {
		return nil, false
	}
	return o.BarometricPressure, true
}

// HasBarometricPressure returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasBarometricPressure() bool {
	if o != nil && !IsNil(o.BarometricPressure) {
		return true
	}

	return false
}

// SetBarometricPressure gets a reference to the given int32 and assigns it to the BarometricPressure field.
func (o *NotifiedPosInfo) SetBarometricPressure(v int32) {
	o.BarometricPressure = &v
}

// GetAltitude returns the Altitude field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetAltitude() float64 {
	if o == nil || IsNil(o.Altitude) {
		var ret float64
		return ret
	}
	return *o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetAltitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Altitude) {
		return nil, false
	}
	return o.Altitude, true
}

// HasAltitude returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasAltitude() bool {
	if o != nil && !IsNil(o.Altitude) {
		return true
	}

	return false
}

// SetAltitude gets a reference to the given float64 and assigns it to the Altitude field.
func (o *NotifiedPosInfo) SetAltitude(v float64) {
	o.Altitude = &v
}

// GetHgmlcCallBackURI returns the HgmlcCallBackURI field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetHgmlcCallBackURI() string {
	if o == nil || IsNil(o.HgmlcCallBackURI) {
		var ret string
		return ret
	}
	return *o.HgmlcCallBackURI
}

// GetHgmlcCallBackURIOk returns a tuple with the HgmlcCallBackURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetHgmlcCallBackURIOk() (*string, bool) {
	if o == nil || IsNil(o.HgmlcCallBackURI) {
		return nil, false
	}
	return o.HgmlcCallBackURI, true
}

// HasHgmlcCallBackURI returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasHgmlcCallBackURI() bool {
	if o != nil && !IsNil(o.HgmlcCallBackURI) {
		return true
	}

	return false
}

// SetHgmlcCallBackURI gets a reference to the given string and assigns it to the HgmlcCallBackURI field.
func (o *NotifiedPosInfo) SetHgmlcCallBackURI(v string) {
	o.HgmlcCallBackURI = &v
}

// GetLdrReference returns the LdrReference field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetLdrReference() string {
	if o == nil || IsNil(o.LdrReference) {
		var ret string
		return ret
	}
	return *o.LdrReference
}

// GetLdrReferenceOk returns a tuple with the LdrReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetLdrReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.LdrReference) {
		return nil, false
	}
	return o.LdrReference, true
}

// HasLdrReference returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasLdrReference() bool {
	if o != nil && !IsNil(o.LdrReference) {
		return true
	}

	return false
}

// SetLdrReference gets a reference to the given string and assigns it to the LdrReference field.
func (o *NotifiedPosInfo) SetLdrReference(v string) {
	o.LdrReference = &v
}

// GetServingLMFIdentification returns the ServingLMFIdentification field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetServingLMFIdentification() string {
	if o == nil || IsNil(o.ServingLMFIdentification) {
		var ret string
		return ret
	}
	return *o.ServingLMFIdentification
}

// GetServingLMFIdentificationOk returns a tuple with the ServingLMFIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetServingLMFIdentificationOk() (*string, bool) {
	if o == nil || IsNil(o.ServingLMFIdentification) {
		return nil, false
	}
	return o.ServingLMFIdentification, true
}

// HasServingLMFIdentification returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasServingLMFIdentification() bool {
	if o != nil && !IsNil(o.ServingLMFIdentification) {
		return true
	}

	return false
}

// SetServingLMFIdentification gets a reference to the given string and assigns it to the ServingLMFIdentification field.
func (o *NotifiedPosInfo) SetServingLMFIdentification(v string) {
	o.ServingLMFIdentification = &v
}

// GetTerminationCause returns the TerminationCause field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetTerminationCause() TerminationCause {
	if o == nil || IsNil(o.TerminationCause) {
		var ret TerminationCause
		return ret
	}
	return *o.TerminationCause
}

// GetTerminationCauseOk returns a tuple with the TerminationCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetTerminationCauseOk() (*TerminationCause, bool) {
	if o == nil || IsNil(o.TerminationCause) {
		return nil, false
	}
	return o.TerminationCause, true
}

// HasTerminationCause returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasTerminationCause() bool {
	if o != nil && !IsNil(o.TerminationCause) {
		return true
	}

	return false
}

// SetTerminationCause gets a reference to the given TerminationCause and assigns it to the TerminationCause field.
func (o *NotifiedPosInfo) SetTerminationCause(v TerminationCause) {
	o.TerminationCause = &v
}

// GetAchievedQos returns the AchievedQos field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetAchievedQos() MinorLocationQoS {
	if o == nil || IsNil(o.AchievedQos) {
		var ret MinorLocationQoS
		return ret
	}
	return *o.AchievedQos
}

// GetAchievedQosOk returns a tuple with the AchievedQos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetAchievedQosOk() (*MinorLocationQoS, bool) {
	if o == nil || IsNil(o.AchievedQos) {
		return nil, false
	}
	return o.AchievedQos, true
}

// HasAchievedQos returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasAchievedQos() bool {
	if o != nil && !IsNil(o.AchievedQos) {
		return true
	}

	return false
}

// SetAchievedQos gets a reference to the given MinorLocationQoS and assigns it to the AchievedQos field.
func (o *NotifiedPosInfo) SetAchievedQos(v MinorLocationQoS) {
	o.AchievedQos = &v
}

// GetMscServerId returns the MscServerId field value if set, zero value otherwise.
func (o *NotifiedPosInfo) GetMscServerId() string {
	if o == nil || IsNil(o.MscServerId) {
		var ret string
		return ret
	}
	return *o.MscServerId
}

// GetMscServerIdOk returns a tuple with the MscServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifiedPosInfo) GetMscServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.MscServerId) {
		return nil, false
	}
	return o.MscServerId, true
}

// HasMscServerId returns a boolean if a field has been set.
func (o *NotifiedPosInfo) HasMscServerId() bool {
	if o != nil && !IsNil(o.MscServerId) {
		return true
	}

	return false
}

// SetMscServerId gets a reference to the given string and assigns it to the MscServerId field.
func (o *NotifiedPosInfo) SetMscServerId(v string) {
	o.MscServerId = &v
}

func (o NotifiedPosInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifiedPosInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locationEvent"] = o.LocationEvent
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.Pei) {
		toSerialize["pei"] = o.Pei
	}
	if !IsNil(o.LocationEstimate) {
		toSerialize["locationEstimate"] = o.LocationEstimate
	}
	if !IsNil(o.LocalLocationEstimate) {
		toSerialize["localLocationEstimate"] = o.LocalLocationEstimate
	}
	if !IsNil(o.AgeOfLocationEstimate) {
		toSerialize["ageOfLocationEstimate"] = o.AgeOfLocationEstimate
	}
	if !IsNil(o.TimestampOfLocationEstimate) {
		toSerialize["timestampOfLocationEstimate"] = o.TimestampOfLocationEstimate
	}
	if !IsNil(o.VelocityEstimate) {
		toSerialize["velocityEstimate"] = o.VelocityEstimate
	}
	if !IsNil(o.PositioningDataList) {
		toSerialize["positioningDataList"] = o.PositioningDataList
	}
	if !IsNil(o.GnssPositioningDataList) {
		toSerialize["gnssPositioningDataList"] = o.GnssPositioningDataList
	}
	if !IsNil(o.Ecgi) {
		toSerialize["ecgi"] = o.Ecgi
	}
	if !IsNil(o.Ncgi) {
		toSerialize["ncgi"] = o.Ncgi
	}
	if !IsNil(o.ServingNode) {
		toSerialize["servingNode"] = o.ServingNode
	}
	if !IsNil(o.TargetMmeName) {
		toSerialize["targetMmeName"] = o.TargetMmeName
	}
	if !IsNil(o.TargetMmeRealm) {
		toSerialize["targetMmeRealm"] = o.TargetMmeRealm
	}
	if !IsNil(o.UtranSrvccInd) {
		toSerialize["utranSrvccInd"] = o.UtranSrvccInd
	}
	if !IsNil(o.CivicAddress) {
		toSerialize["civicAddress"] = o.CivicAddress
	}
	if !IsNil(o.BarometricPressure) {
		toSerialize["barometricPressure"] = o.BarometricPressure
	}
	if !IsNil(o.Altitude) {
		toSerialize["altitude"] = o.Altitude
	}
	if !IsNil(o.HgmlcCallBackURI) {
		toSerialize["hgmlcCallBackURI"] = o.HgmlcCallBackURI
	}
	if !IsNil(o.LdrReference) {
		toSerialize["ldrReference"] = o.LdrReference
	}
	if !IsNil(o.ServingLMFIdentification) {
		toSerialize["servingLMFIdentification"] = o.ServingLMFIdentification
	}
	if !IsNil(o.TerminationCause) {
		toSerialize["terminationCause"] = o.TerminationCause
	}
	if !IsNil(o.AchievedQos) {
		toSerialize["achievedQos"] = o.AchievedQos
	}
	if !IsNil(o.MscServerId) {
		toSerialize["mscServerId"] = o.MscServerId
	}
	return toSerialize, nil
}

type NullableNotifiedPosInfo struct {
	value *NotifiedPosInfo
	isSet bool
}

func (v NullableNotifiedPosInfo) Get() *NotifiedPosInfo {
	return v.value
}

func (v *NullableNotifiedPosInfo) Set(val *NotifiedPosInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifiedPosInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifiedPosInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifiedPosInfo(val *NotifiedPosInfo) *NullableNotifiedPosInfo {
	return &NullableNotifiedPosInfo{value: val, isSet: true}
}

func (v NullableNotifiedPosInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifiedPosInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

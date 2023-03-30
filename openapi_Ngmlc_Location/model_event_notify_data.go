/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
	"time"
)

// checks if the EventNotifyData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventNotifyData{}

// EventNotifyData Contains the input parameters for the target UE in EventNotify Notification service operation
type EventNotifyData struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	// LDR Reference.
	LdrReference string `json:"ldrReference"`
	EventNotifyDataType EventNotifyDataType `json:"eventNotifyDataType"`
	LocationEstimate *GeographicArea `json:"locationEstimate,omitempty"`
	CivicAddress *CivicAddress `json:"civicAddress,omitempty"`
	LocalLocationEstimate *LocalArea `json:"localLocationEstimate,omitempty"`
	// Indicates value of the age of the location estimate.
	AgeOfLocationEstimate *int32 `json:"ageOfLocationEstimate,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimestampOfLocationEstimate *time.Time `json:"timestampOfLocationEstimate,omitempty"`
	PositioningDataList []PositioningMethodAndUsage `json:"positioningDataList,omitempty"`
	GnssPositioningDataList []GnssPositioningMethodAndUsage `json:"gnssPositioningDataList,omitempty"`
	// LMF identification.
	LmfIdentification *string `json:"lmfIdentification,omitempty"`
	// String identifying the AMF ID composed of AMF Region ID (8 bits), AMF Set ID (10 bits) and AMF  Pointer (6 bits) as specified in clause 2.10.1 of 3GPP TS 23.003. It is encoded as a string of  6 hexadecimal characters (i.e., 24 bits).  
	AmfId *string `json:"amfId,omitempty"`
	TerminationCause *TerminationCause `json:"terminationCause,omitempty"`
	VelocityEstimate *VelocityEstimate `json:"velocityEstimate,omitempty"`
	// Indicates value of altitude.
	Altitude *float64 `json:"altitude,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	TargetNode *string `json:"targetNode,omitempty"`
	AccuracyFulfilmentIndicator *AccuracyFulfilmentIndicator `json:"accuracyFulfilmentIndicator,omitempty"`
	FailureCause *FailureCause `json:"failureCause,omitempty"`
	AchievedQos *MinorLocationQoS `json:"achievedQos,omitempty"`
}

// NewEventNotifyData instantiates a new EventNotifyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventNotifyData(ldrReference string, eventNotifyDataType EventNotifyDataType) *EventNotifyData {
	this := EventNotifyData{}
	this.LdrReference = ldrReference
	this.EventNotifyDataType = eventNotifyDataType
	return &this
}

// NewEventNotifyDataWithDefaults instantiates a new EventNotifyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventNotifyDataWithDefaults() *EventNotifyData {
	this := EventNotifyData{}
	return &this
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *EventNotifyData) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *EventNotifyData) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *EventNotifyData) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *EventNotifyData) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *EventNotifyData) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *EventNotifyData) SetSupi(v string) {
	o.Supi = &v
}

// GetLdrReference returns the LdrReference field value
func (o *EventNotifyData) GetLdrReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LdrReference
}

// GetLdrReferenceOk returns a tuple with the LdrReference field value
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetLdrReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LdrReference, true
}

// SetLdrReference sets field value
func (o *EventNotifyData) SetLdrReference(v string) {
	o.LdrReference = v
}

// GetEventNotifyDataType returns the EventNotifyDataType field value
func (o *EventNotifyData) GetEventNotifyDataType() EventNotifyDataType {
	if o == nil {
		var ret EventNotifyDataType
		return ret
	}

	return o.EventNotifyDataType
}

// GetEventNotifyDataTypeOk returns a tuple with the EventNotifyDataType field value
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetEventNotifyDataTypeOk() (*EventNotifyDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventNotifyDataType, true
}

// SetEventNotifyDataType sets field value
func (o *EventNotifyData) SetEventNotifyDataType(v EventNotifyDataType) {
	o.EventNotifyDataType = v
}

// GetLocationEstimate returns the LocationEstimate field value if set, zero value otherwise.
func (o *EventNotifyData) GetLocationEstimate() GeographicArea {
	if o == nil || IsNil(o.LocationEstimate) {
		var ret GeographicArea
		return ret
	}
	return *o.LocationEstimate
}

// GetLocationEstimateOk returns a tuple with the LocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetLocationEstimateOk() (*GeographicArea, bool) {
	if o == nil || IsNil(o.LocationEstimate) {
		return nil, false
	}
	return o.LocationEstimate, true
}

// HasLocationEstimate returns a boolean if a field has been set.
func (o *EventNotifyData) HasLocationEstimate() bool {
	if o != nil && !IsNil(o.LocationEstimate) {
		return true
	}

	return false
}

// SetLocationEstimate gets a reference to the given GeographicArea and assigns it to the LocationEstimate field.
func (o *EventNotifyData) SetLocationEstimate(v GeographicArea) {
	o.LocationEstimate = &v
}

// GetCivicAddress returns the CivicAddress field value if set, zero value otherwise.
func (o *EventNotifyData) GetCivicAddress() CivicAddress {
	if o == nil || IsNil(o.CivicAddress) {
		var ret CivicAddress
		return ret
	}
	return *o.CivicAddress
}

// GetCivicAddressOk returns a tuple with the CivicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetCivicAddressOk() (*CivicAddress, bool) {
	if o == nil || IsNil(o.CivicAddress) {
		return nil, false
	}
	return o.CivicAddress, true
}

// HasCivicAddress returns a boolean if a field has been set.
func (o *EventNotifyData) HasCivicAddress() bool {
	if o != nil && !IsNil(o.CivicAddress) {
		return true
	}

	return false
}

// SetCivicAddress gets a reference to the given CivicAddress and assigns it to the CivicAddress field.
func (o *EventNotifyData) SetCivicAddress(v CivicAddress) {
	o.CivicAddress = &v
}

// GetLocalLocationEstimate returns the LocalLocationEstimate field value if set, zero value otherwise.
func (o *EventNotifyData) GetLocalLocationEstimate() LocalArea {
	if o == nil || IsNil(o.LocalLocationEstimate) {
		var ret LocalArea
		return ret
	}
	return *o.LocalLocationEstimate
}

// GetLocalLocationEstimateOk returns a tuple with the LocalLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetLocalLocationEstimateOk() (*LocalArea, bool) {
	if o == nil || IsNil(o.LocalLocationEstimate) {
		return nil, false
	}
	return o.LocalLocationEstimate, true
}

// HasLocalLocationEstimate returns a boolean if a field has been set.
func (o *EventNotifyData) HasLocalLocationEstimate() bool {
	if o != nil && !IsNil(o.LocalLocationEstimate) {
		return true
	}

	return false
}

// SetLocalLocationEstimate gets a reference to the given LocalArea and assigns it to the LocalLocationEstimate field.
func (o *EventNotifyData) SetLocalLocationEstimate(v LocalArea) {
	o.LocalLocationEstimate = &v
}

// GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field value if set, zero value otherwise.
func (o *EventNotifyData) GetAgeOfLocationEstimate() int32 {
	if o == nil || IsNil(o.AgeOfLocationEstimate) {
		var ret int32
		return ret
	}
	return *o.AgeOfLocationEstimate
}

// GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetAgeOfLocationEstimateOk() (*int32, bool) {
	if o == nil || IsNil(o.AgeOfLocationEstimate) {
		return nil, false
	}
	return o.AgeOfLocationEstimate, true
}

// HasAgeOfLocationEstimate returns a boolean if a field has been set.
func (o *EventNotifyData) HasAgeOfLocationEstimate() bool {
	if o != nil && !IsNil(o.AgeOfLocationEstimate) {
		return true
	}

	return false
}

// SetAgeOfLocationEstimate gets a reference to the given int32 and assigns it to the AgeOfLocationEstimate field.
func (o *EventNotifyData) SetAgeOfLocationEstimate(v int32) {
	o.AgeOfLocationEstimate = &v
}

// GetTimestampOfLocationEstimate returns the TimestampOfLocationEstimate field value if set, zero value otherwise.
func (o *EventNotifyData) GetTimestampOfLocationEstimate() time.Time {
	if o == nil || IsNil(o.TimestampOfLocationEstimate) {
		var ret time.Time
		return ret
	}
	return *o.TimestampOfLocationEstimate
}

// GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetTimestampOfLocationEstimateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampOfLocationEstimate) {
		return nil, false
	}
	return o.TimestampOfLocationEstimate, true
}

// HasTimestampOfLocationEstimate returns a boolean if a field has been set.
func (o *EventNotifyData) HasTimestampOfLocationEstimate() bool {
	if o != nil && !IsNil(o.TimestampOfLocationEstimate) {
		return true
	}

	return false
}

// SetTimestampOfLocationEstimate gets a reference to the given time.Time and assigns it to the TimestampOfLocationEstimate field.
func (o *EventNotifyData) SetTimestampOfLocationEstimate(v time.Time) {
	o.TimestampOfLocationEstimate = &v
}

// GetPositioningDataList returns the PositioningDataList field value if set, zero value otherwise.
func (o *EventNotifyData) GetPositioningDataList() []PositioningMethodAndUsage {
	if o == nil || IsNil(o.PositioningDataList) {
		var ret []PositioningMethodAndUsage
		return ret
	}
	return o.PositioningDataList
}

// GetPositioningDataListOk returns a tuple with the PositioningDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetPositioningDataListOk() ([]PositioningMethodAndUsage, bool) {
	if o == nil || IsNil(o.PositioningDataList) {
		return nil, false
	}
	return o.PositioningDataList, true
}

// HasPositioningDataList returns a boolean if a field has been set.
func (o *EventNotifyData) HasPositioningDataList() bool {
	if o != nil && !IsNil(o.PositioningDataList) {
		return true
	}

	return false
}

// SetPositioningDataList gets a reference to the given []PositioningMethodAndUsage and assigns it to the PositioningDataList field.
func (o *EventNotifyData) SetPositioningDataList(v []PositioningMethodAndUsage) {
	o.PositioningDataList = v
}

// GetGnssPositioningDataList returns the GnssPositioningDataList field value if set, zero value otherwise.
func (o *EventNotifyData) GetGnssPositioningDataList() []GnssPositioningMethodAndUsage {
	if o == nil || IsNil(o.GnssPositioningDataList) {
		var ret []GnssPositioningMethodAndUsage
		return ret
	}
	return o.GnssPositioningDataList
}

// GetGnssPositioningDataListOk returns a tuple with the GnssPositioningDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetGnssPositioningDataListOk() ([]GnssPositioningMethodAndUsage, bool) {
	if o == nil || IsNil(o.GnssPositioningDataList) {
		return nil, false
	}
	return o.GnssPositioningDataList, true
}

// HasGnssPositioningDataList returns a boolean if a field has been set.
func (o *EventNotifyData) HasGnssPositioningDataList() bool {
	if o != nil && !IsNil(o.GnssPositioningDataList) {
		return true
	}

	return false
}

// SetGnssPositioningDataList gets a reference to the given []GnssPositioningMethodAndUsage and assigns it to the GnssPositioningDataList field.
func (o *EventNotifyData) SetGnssPositioningDataList(v []GnssPositioningMethodAndUsage) {
	o.GnssPositioningDataList = v
}

// GetLmfIdentification returns the LmfIdentification field value if set, zero value otherwise.
func (o *EventNotifyData) GetLmfIdentification() string {
	if o == nil || IsNil(o.LmfIdentification) {
		var ret string
		return ret
	}
	return *o.LmfIdentification
}

// GetLmfIdentificationOk returns a tuple with the LmfIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetLmfIdentificationOk() (*string, bool) {
	if o == nil || IsNil(o.LmfIdentification) {
		return nil, false
	}
	return o.LmfIdentification, true
}

// HasLmfIdentification returns a boolean if a field has been set.
func (o *EventNotifyData) HasLmfIdentification() bool {
	if o != nil && !IsNil(o.LmfIdentification) {
		return true
	}

	return false
}

// SetLmfIdentification gets a reference to the given string and assigns it to the LmfIdentification field.
func (o *EventNotifyData) SetLmfIdentification(v string) {
	o.LmfIdentification = &v
}

// GetAmfId returns the AmfId field value if set, zero value otherwise.
func (o *EventNotifyData) GetAmfId() string {
	if o == nil || IsNil(o.AmfId) {
		var ret string
		return ret
	}
	return *o.AmfId
}

// GetAmfIdOk returns a tuple with the AmfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetAmfIdOk() (*string, bool) {
	if o == nil || IsNil(o.AmfId) {
		return nil, false
	}
	return o.AmfId, true
}

// HasAmfId returns a boolean if a field has been set.
func (o *EventNotifyData) HasAmfId() bool {
	if o != nil && !IsNil(o.AmfId) {
		return true
	}

	return false
}

// SetAmfId gets a reference to the given string and assigns it to the AmfId field.
func (o *EventNotifyData) SetAmfId(v string) {
	o.AmfId = &v
}

// GetTerminationCause returns the TerminationCause field value if set, zero value otherwise.
func (o *EventNotifyData) GetTerminationCause() TerminationCause {
	if o == nil || IsNil(o.TerminationCause) {
		var ret TerminationCause
		return ret
	}
	return *o.TerminationCause
}

// GetTerminationCauseOk returns a tuple with the TerminationCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetTerminationCauseOk() (*TerminationCause, bool) {
	if o == nil || IsNil(o.TerminationCause) {
		return nil, false
	}
	return o.TerminationCause, true
}

// HasTerminationCause returns a boolean if a field has been set.
func (o *EventNotifyData) HasTerminationCause() bool {
	if o != nil && !IsNil(o.TerminationCause) {
		return true
	}

	return false
}

// SetTerminationCause gets a reference to the given TerminationCause and assigns it to the TerminationCause field.
func (o *EventNotifyData) SetTerminationCause(v TerminationCause) {
	o.TerminationCause = &v
}

// GetVelocityEstimate returns the VelocityEstimate field value if set, zero value otherwise.
func (o *EventNotifyData) GetVelocityEstimate() VelocityEstimate {
	if o == nil || IsNil(o.VelocityEstimate) {
		var ret VelocityEstimate
		return ret
	}
	return *o.VelocityEstimate
}

// GetVelocityEstimateOk returns a tuple with the VelocityEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetVelocityEstimateOk() (*VelocityEstimate, bool) {
	if o == nil || IsNil(o.VelocityEstimate) {
		return nil, false
	}
	return o.VelocityEstimate, true
}

// HasVelocityEstimate returns a boolean if a field has been set.
func (o *EventNotifyData) HasVelocityEstimate() bool {
	if o != nil && !IsNil(o.VelocityEstimate) {
		return true
	}

	return false
}

// SetVelocityEstimate gets a reference to the given VelocityEstimate and assigns it to the VelocityEstimate field.
func (o *EventNotifyData) SetVelocityEstimate(v VelocityEstimate) {
	o.VelocityEstimate = &v
}

// GetAltitude returns the Altitude field value if set, zero value otherwise.
func (o *EventNotifyData) GetAltitude() float64 {
	if o == nil || IsNil(o.Altitude) {
		var ret float64
		return ret
	}
	return *o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetAltitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Altitude) {
		return nil, false
	}
	return o.Altitude, true
}

// HasAltitude returns a boolean if a field has been set.
func (o *EventNotifyData) HasAltitude() bool {
	if o != nil && !IsNil(o.Altitude) {
		return true
	}

	return false
}

// SetAltitude gets a reference to the given float64 and assigns it to the Altitude field.
func (o *EventNotifyData) SetAltitude(v float64) {
	o.Altitude = &v
}

// GetTargetNode returns the TargetNode field value if set, zero value otherwise.
func (o *EventNotifyData) GetTargetNode() string {
	if o == nil || IsNil(o.TargetNode) {
		var ret string
		return ret
	}
	return *o.TargetNode
}

// GetTargetNodeOk returns a tuple with the TargetNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetTargetNodeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetNode) {
		return nil, false
	}
	return o.TargetNode, true
}

// HasTargetNode returns a boolean if a field has been set.
func (o *EventNotifyData) HasTargetNode() bool {
	if o != nil && !IsNil(o.TargetNode) {
		return true
	}

	return false
}

// SetTargetNode gets a reference to the given string and assigns it to the TargetNode field.
func (o *EventNotifyData) SetTargetNode(v string) {
	o.TargetNode = &v
}

// GetAccuracyFulfilmentIndicator returns the AccuracyFulfilmentIndicator field value if set, zero value otherwise.
func (o *EventNotifyData) GetAccuracyFulfilmentIndicator() AccuracyFulfilmentIndicator {
	if o == nil || IsNil(o.AccuracyFulfilmentIndicator) {
		var ret AccuracyFulfilmentIndicator
		return ret
	}
	return *o.AccuracyFulfilmentIndicator
}

// GetAccuracyFulfilmentIndicatorOk returns a tuple with the AccuracyFulfilmentIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetAccuracyFulfilmentIndicatorOk() (*AccuracyFulfilmentIndicator, bool) {
	if o == nil || IsNil(o.AccuracyFulfilmentIndicator) {
		return nil, false
	}
	return o.AccuracyFulfilmentIndicator, true
}

// HasAccuracyFulfilmentIndicator returns a boolean if a field has been set.
func (o *EventNotifyData) HasAccuracyFulfilmentIndicator() bool {
	if o != nil && !IsNil(o.AccuracyFulfilmentIndicator) {
		return true
	}

	return false
}

// SetAccuracyFulfilmentIndicator gets a reference to the given AccuracyFulfilmentIndicator and assigns it to the AccuracyFulfilmentIndicator field.
func (o *EventNotifyData) SetAccuracyFulfilmentIndicator(v AccuracyFulfilmentIndicator) {
	o.AccuracyFulfilmentIndicator = &v
}

// GetFailureCause returns the FailureCause field value if set, zero value otherwise.
func (o *EventNotifyData) GetFailureCause() FailureCause {
	if o == nil || IsNil(o.FailureCause) {
		var ret FailureCause
		return ret
	}
	return *o.FailureCause
}

// GetFailureCauseOk returns a tuple with the FailureCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetFailureCauseOk() (*FailureCause, bool) {
	if o == nil || IsNil(o.FailureCause) {
		return nil, false
	}
	return o.FailureCause, true
}

// HasFailureCause returns a boolean if a field has been set.
func (o *EventNotifyData) HasFailureCause() bool {
	if o != nil && !IsNil(o.FailureCause) {
		return true
	}

	return false
}

// SetFailureCause gets a reference to the given FailureCause and assigns it to the FailureCause field.
func (o *EventNotifyData) SetFailureCause(v FailureCause) {
	o.FailureCause = &v
}

// GetAchievedQos returns the AchievedQos field value if set, zero value otherwise.
func (o *EventNotifyData) GetAchievedQos() MinorLocationQoS {
	if o == nil || IsNil(o.AchievedQos) {
		var ret MinorLocationQoS
		return ret
	}
	return *o.AchievedQos
}

// GetAchievedQosOk returns a tuple with the AchievedQos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetAchievedQosOk() (*MinorLocationQoS, bool) {
	if o == nil || IsNil(o.AchievedQos) {
		return nil, false
	}
	return o.AchievedQos, true
}

// HasAchievedQos returns a boolean if a field has been set.
func (o *EventNotifyData) HasAchievedQos() bool {
	if o != nil && !IsNil(o.AchievedQos) {
		return true
	}

	return false
}

// SetAchievedQos gets a reference to the given MinorLocationQoS and assigns it to the AchievedQos field.
func (o *EventNotifyData) SetAchievedQos(v MinorLocationQoS) {
	o.AchievedQos = &v
}

func (o EventNotifyData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventNotifyData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	toSerialize["ldrReference"] = o.LdrReference
	toSerialize["eventNotifyDataType"] = o.EventNotifyDataType
	if !IsNil(o.LocationEstimate) {
		toSerialize["locationEstimate"] = o.LocationEstimate
	}
	if !IsNil(o.CivicAddress) {
		toSerialize["civicAddress"] = o.CivicAddress
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
	if !IsNil(o.PositioningDataList) {
		toSerialize["positioningDataList"] = o.PositioningDataList
	}
	if !IsNil(o.GnssPositioningDataList) {
		toSerialize["gnssPositioningDataList"] = o.GnssPositioningDataList
	}
	if !IsNil(o.LmfIdentification) {
		toSerialize["lmfIdentification"] = o.LmfIdentification
	}
	if !IsNil(o.AmfId) {
		toSerialize["amfId"] = o.AmfId
	}
	if !IsNil(o.TerminationCause) {
		toSerialize["terminationCause"] = o.TerminationCause
	}
	if !IsNil(o.VelocityEstimate) {
		toSerialize["velocityEstimate"] = o.VelocityEstimate
	}
	if !IsNil(o.Altitude) {
		toSerialize["altitude"] = o.Altitude
	}
	if !IsNil(o.TargetNode) {
		toSerialize["targetNode"] = o.TargetNode
	}
	if !IsNil(o.AccuracyFulfilmentIndicator) {
		toSerialize["accuracyFulfilmentIndicator"] = o.AccuracyFulfilmentIndicator
	}
	if !IsNil(o.FailureCause) {
		toSerialize["failureCause"] = o.FailureCause
	}
	if !IsNil(o.AchievedQos) {
		toSerialize["achievedQos"] = o.AchievedQos
	}
	return toSerialize, nil
}

type NullableEventNotifyData struct {
	value *EventNotifyData
	isSet bool
}

func (v NullableEventNotifyData) Get() *EventNotifyData {
	return v.value
}

func (v *NullableEventNotifyData) Set(val *EventNotifyData) {
	v.value = val
	v.isSet = true
}

func (v NullableEventNotifyData) IsSet() bool {
	return v.isSet
}

func (v *NullableEventNotifyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventNotifyData(val *EventNotifyData) *NullableEventNotifyData {
	return &NullableEventNotifyData{value: val, isSet: true}
}

func (v NullableEventNotifyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventNotifyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



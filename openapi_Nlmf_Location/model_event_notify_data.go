/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Location

import (
	"encoding/json"
	"time"
)

// checks if the EventNotifyData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventNotifyData{}

// EventNotifyData Information within Event Notify Request.
type EventNotifyData struct {
	ReportedEventType ReportedEventType `json:"reportedEventType"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	HgmlcCallBackURI *string `json:"hgmlcCallBackURI,omitempty"`
	// LDR Reference.
	LdrReference string `json:"ldrReference"`
	LocationEstimate *GeographicArea `json:"locationEstimate,omitempty"`
	// Indicates value of the age of the location estimate.
	AgeOfLocationEstimate *int32 `json:"ageOfLocationEstimate,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimestampOfLocationEstimate *time.Time `json:"timestampOfLocationEstimate,omitempty"`
	CivicAddress *CivicAddress `json:"civicAddress,omitempty"`
	LocalLocationEstimate *LocalArea `json:"localLocationEstimate,omitempty"`
	PositioningDataList []PositioningMethodAndUsage `json:"positioningDataList,omitempty"`
	GnssPositioningDataList []GnssPositioningMethodAndUsage `json:"gnssPositioningDataList,omitempty"`
	// LMF identification.
	ServingLMFidentification *string `json:"servingLMFidentification,omitempty"`
	TerminationCause *TerminationCause `json:"terminationCause,omitempty"`
	VelocityEstimate *VelocityEstimate `json:"velocityEstimate,omitempty"`
	// Indicates value of altitude.
	Altitude *float64 `json:"altitude,omitempty"`
	AchievedQos *MinorLocationQoS `json:"achievedQos,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewEventNotifyData instantiates a new EventNotifyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventNotifyData(reportedEventType ReportedEventType, ldrReference string) *EventNotifyData {
	this := EventNotifyData{}
	this.ReportedEventType = reportedEventType
	this.LdrReference = ldrReference
	return &this
}

// NewEventNotifyDataWithDefaults instantiates a new EventNotifyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventNotifyDataWithDefaults() *EventNotifyData {
	this := EventNotifyData{}
	return &this
}

// GetReportedEventType returns the ReportedEventType field value
func (o *EventNotifyData) GetReportedEventType() ReportedEventType {
	if o == nil {
		var ret ReportedEventType
		return ret
	}

	return o.ReportedEventType
}

// GetReportedEventTypeOk returns a tuple with the ReportedEventType field value
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetReportedEventTypeOk() (*ReportedEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportedEventType, true
}

// SetReportedEventType sets field value
func (o *EventNotifyData) SetReportedEventType(v ReportedEventType) {
	o.ReportedEventType = v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *EventNotifyData) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *EventNotifyData) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *EventNotifyData) SetSupi(v string) {
	o.Supi = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *EventNotifyData) GetGpsi() string {
	if o == nil || isNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetGpsiOk() (*string, bool) {
	if o == nil || isNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *EventNotifyData) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *EventNotifyData) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetHgmlcCallBackURI returns the HgmlcCallBackURI field value if set, zero value otherwise.
func (o *EventNotifyData) GetHgmlcCallBackURI() string {
	if o == nil || isNil(o.HgmlcCallBackURI) {
		var ret string
		return ret
	}
	return *o.HgmlcCallBackURI
}

// GetHgmlcCallBackURIOk returns a tuple with the HgmlcCallBackURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetHgmlcCallBackURIOk() (*string, bool) {
	if o == nil || isNil(o.HgmlcCallBackURI) {
		return nil, false
	}
	return o.HgmlcCallBackURI, true
}

// HasHgmlcCallBackURI returns a boolean if a field has been set.
func (o *EventNotifyData) HasHgmlcCallBackURI() bool {
	if o != nil && !isNil(o.HgmlcCallBackURI) {
		return true
	}

	return false
}

// SetHgmlcCallBackURI gets a reference to the given string and assigns it to the HgmlcCallBackURI field.
func (o *EventNotifyData) SetHgmlcCallBackURI(v string) {
	o.HgmlcCallBackURI = &v
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

// GetLocationEstimate returns the LocationEstimate field value if set, zero value otherwise.
func (o *EventNotifyData) GetLocationEstimate() GeographicArea {
	if o == nil || isNil(o.LocationEstimate) {
		var ret GeographicArea
		return ret
	}
	return *o.LocationEstimate
}

// GetLocationEstimateOk returns a tuple with the LocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetLocationEstimateOk() (*GeographicArea, bool) {
	if o == nil || isNil(o.LocationEstimate) {
		return nil, false
	}
	return o.LocationEstimate, true
}

// HasLocationEstimate returns a boolean if a field has been set.
func (o *EventNotifyData) HasLocationEstimate() bool {
	if o != nil && !isNil(o.LocationEstimate) {
		return true
	}

	return false
}

// SetLocationEstimate gets a reference to the given GeographicArea and assigns it to the LocationEstimate field.
func (o *EventNotifyData) SetLocationEstimate(v GeographicArea) {
	o.LocationEstimate = &v
}

// GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field value if set, zero value otherwise.
func (o *EventNotifyData) GetAgeOfLocationEstimate() int32 {
	if o == nil || isNil(o.AgeOfLocationEstimate) {
		var ret int32
		return ret
	}
	return *o.AgeOfLocationEstimate
}

// GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetAgeOfLocationEstimateOk() (*int32, bool) {
	if o == nil || isNil(o.AgeOfLocationEstimate) {
		return nil, false
	}
	return o.AgeOfLocationEstimate, true
}

// HasAgeOfLocationEstimate returns a boolean if a field has been set.
func (o *EventNotifyData) HasAgeOfLocationEstimate() bool {
	if o != nil && !isNil(o.AgeOfLocationEstimate) {
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
	if o == nil || isNil(o.TimestampOfLocationEstimate) {
		var ret time.Time
		return ret
	}
	return *o.TimestampOfLocationEstimate
}

// GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetTimestampOfLocationEstimateOk() (*time.Time, bool) {
	if o == nil || isNil(o.TimestampOfLocationEstimate) {
		return nil, false
	}
	return o.TimestampOfLocationEstimate, true
}

// HasTimestampOfLocationEstimate returns a boolean if a field has been set.
func (o *EventNotifyData) HasTimestampOfLocationEstimate() bool {
	if o != nil && !isNil(o.TimestampOfLocationEstimate) {
		return true
	}

	return false
}

// SetTimestampOfLocationEstimate gets a reference to the given time.Time and assigns it to the TimestampOfLocationEstimate field.
func (o *EventNotifyData) SetTimestampOfLocationEstimate(v time.Time) {
	o.TimestampOfLocationEstimate = &v
}

// GetCivicAddress returns the CivicAddress field value if set, zero value otherwise.
func (o *EventNotifyData) GetCivicAddress() CivicAddress {
	if o == nil || isNil(o.CivicAddress) {
		var ret CivicAddress
		return ret
	}
	return *o.CivicAddress
}

// GetCivicAddressOk returns a tuple with the CivicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetCivicAddressOk() (*CivicAddress, bool) {
	if o == nil || isNil(o.CivicAddress) {
		return nil, false
	}
	return o.CivicAddress, true
}

// HasCivicAddress returns a boolean if a field has been set.
func (o *EventNotifyData) HasCivicAddress() bool {
	if o != nil && !isNil(o.CivicAddress) {
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
	if o == nil || isNil(o.LocalLocationEstimate) {
		var ret LocalArea
		return ret
	}
	return *o.LocalLocationEstimate
}

// GetLocalLocationEstimateOk returns a tuple with the LocalLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetLocalLocationEstimateOk() (*LocalArea, bool) {
	if o == nil || isNil(o.LocalLocationEstimate) {
		return nil, false
	}
	return o.LocalLocationEstimate, true
}

// HasLocalLocationEstimate returns a boolean if a field has been set.
func (o *EventNotifyData) HasLocalLocationEstimate() bool {
	if o != nil && !isNil(o.LocalLocationEstimate) {
		return true
	}

	return false
}

// SetLocalLocationEstimate gets a reference to the given LocalArea and assigns it to the LocalLocationEstimate field.
func (o *EventNotifyData) SetLocalLocationEstimate(v LocalArea) {
	o.LocalLocationEstimate = &v
}

// GetPositioningDataList returns the PositioningDataList field value if set, zero value otherwise.
func (o *EventNotifyData) GetPositioningDataList() []PositioningMethodAndUsage {
	if o == nil || isNil(o.PositioningDataList) {
		var ret []PositioningMethodAndUsage
		return ret
	}
	return o.PositioningDataList
}

// GetPositioningDataListOk returns a tuple with the PositioningDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetPositioningDataListOk() ([]PositioningMethodAndUsage, bool) {
	if o == nil || isNil(o.PositioningDataList) {
		return nil, false
	}
	return o.PositioningDataList, true
}

// HasPositioningDataList returns a boolean if a field has been set.
func (o *EventNotifyData) HasPositioningDataList() bool {
	if o != nil && !isNil(o.PositioningDataList) {
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
	if o == nil || isNil(o.GnssPositioningDataList) {
		var ret []GnssPositioningMethodAndUsage
		return ret
	}
	return o.GnssPositioningDataList
}

// GetGnssPositioningDataListOk returns a tuple with the GnssPositioningDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetGnssPositioningDataListOk() ([]GnssPositioningMethodAndUsage, bool) {
	if o == nil || isNil(o.GnssPositioningDataList) {
		return nil, false
	}
	return o.GnssPositioningDataList, true
}

// HasGnssPositioningDataList returns a boolean if a field has been set.
func (o *EventNotifyData) HasGnssPositioningDataList() bool {
	if o != nil && !isNil(o.GnssPositioningDataList) {
		return true
	}

	return false
}

// SetGnssPositioningDataList gets a reference to the given []GnssPositioningMethodAndUsage and assigns it to the GnssPositioningDataList field.
func (o *EventNotifyData) SetGnssPositioningDataList(v []GnssPositioningMethodAndUsage) {
	o.GnssPositioningDataList = v
}

// GetServingLMFidentification returns the ServingLMFidentification field value if set, zero value otherwise.
func (o *EventNotifyData) GetServingLMFidentification() string {
	if o == nil || isNil(o.ServingLMFidentification) {
		var ret string
		return ret
	}
	return *o.ServingLMFidentification
}

// GetServingLMFidentificationOk returns a tuple with the ServingLMFidentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetServingLMFidentificationOk() (*string, bool) {
	if o == nil || isNil(o.ServingLMFidentification) {
		return nil, false
	}
	return o.ServingLMFidentification, true
}

// HasServingLMFidentification returns a boolean if a field has been set.
func (o *EventNotifyData) HasServingLMFidentification() bool {
	if o != nil && !isNil(o.ServingLMFidentification) {
		return true
	}

	return false
}

// SetServingLMFidentification gets a reference to the given string and assigns it to the ServingLMFidentification field.
func (o *EventNotifyData) SetServingLMFidentification(v string) {
	o.ServingLMFidentification = &v
}

// GetTerminationCause returns the TerminationCause field value if set, zero value otherwise.
func (o *EventNotifyData) GetTerminationCause() TerminationCause {
	if o == nil || isNil(o.TerminationCause) {
		var ret TerminationCause
		return ret
	}
	return *o.TerminationCause
}

// GetTerminationCauseOk returns a tuple with the TerminationCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetTerminationCauseOk() (*TerminationCause, bool) {
	if o == nil || isNil(o.TerminationCause) {
		return nil, false
	}
	return o.TerminationCause, true
}

// HasTerminationCause returns a boolean if a field has been set.
func (o *EventNotifyData) HasTerminationCause() bool {
	if o != nil && !isNil(o.TerminationCause) {
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
	if o == nil || isNil(o.VelocityEstimate) {
		var ret VelocityEstimate
		return ret
	}
	return *o.VelocityEstimate
}

// GetVelocityEstimateOk returns a tuple with the VelocityEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetVelocityEstimateOk() (*VelocityEstimate, bool) {
	if o == nil || isNil(o.VelocityEstimate) {
		return nil, false
	}
	return o.VelocityEstimate, true
}

// HasVelocityEstimate returns a boolean if a field has been set.
func (o *EventNotifyData) HasVelocityEstimate() bool {
	if o != nil && !isNil(o.VelocityEstimate) {
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
	if o == nil || isNil(o.Altitude) {
		var ret float64
		return ret
	}
	return *o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetAltitudeOk() (*float64, bool) {
	if o == nil || isNil(o.Altitude) {
		return nil, false
	}
	return o.Altitude, true
}

// HasAltitude returns a boolean if a field has been set.
func (o *EventNotifyData) HasAltitude() bool {
	if o != nil && !isNil(o.Altitude) {
		return true
	}

	return false
}

// SetAltitude gets a reference to the given float64 and assigns it to the Altitude field.
func (o *EventNotifyData) SetAltitude(v float64) {
	o.Altitude = &v
}

// GetAchievedQos returns the AchievedQos field value if set, zero value otherwise.
func (o *EventNotifyData) GetAchievedQos() MinorLocationQoS {
	if o == nil || isNil(o.AchievedQos) {
		var ret MinorLocationQoS
		return ret
	}
	return *o.AchievedQos
}

// GetAchievedQosOk returns a tuple with the AchievedQos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetAchievedQosOk() (*MinorLocationQoS, bool) {
	if o == nil || isNil(o.AchievedQos) {
		return nil, false
	}
	return o.AchievedQos, true
}

// HasAchievedQos returns a boolean if a field has been set.
func (o *EventNotifyData) HasAchievedQos() bool {
	if o != nil && !isNil(o.AchievedQos) {
		return true
	}

	return false
}

// SetAchievedQos gets a reference to the given MinorLocationQoS and assigns it to the AchievedQos field.
func (o *EventNotifyData) SetAchievedQos(v MinorLocationQoS) {
	o.AchievedQos = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *EventNotifyData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *EventNotifyData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *EventNotifyData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
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
	toSerialize["reportedEventType"] = o.ReportedEventType
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !isNil(o.HgmlcCallBackURI) {
		toSerialize["hgmlcCallBackURI"] = o.HgmlcCallBackURI
	}
	toSerialize["ldrReference"] = o.LdrReference
	if !isNil(o.LocationEstimate) {
		toSerialize["locationEstimate"] = o.LocationEstimate
	}
	if !isNil(o.AgeOfLocationEstimate) {
		toSerialize["ageOfLocationEstimate"] = o.AgeOfLocationEstimate
	}
	if !isNil(o.TimestampOfLocationEstimate) {
		toSerialize["timestampOfLocationEstimate"] = o.TimestampOfLocationEstimate
	}
	if !isNil(o.CivicAddress) {
		toSerialize["civicAddress"] = o.CivicAddress
	}
	if !isNil(o.LocalLocationEstimate) {
		toSerialize["localLocationEstimate"] = o.LocalLocationEstimate
	}
	if !isNil(o.PositioningDataList) {
		toSerialize["positioningDataList"] = o.PositioningDataList
	}
	if !isNil(o.GnssPositioningDataList) {
		toSerialize["gnssPositioningDataList"] = o.GnssPositioningDataList
	}
	if !isNil(o.ServingLMFidentification) {
		toSerialize["servingLMFidentification"] = o.ServingLMFidentification
	}
	if !isNil(o.TerminationCause) {
		toSerialize["terminationCause"] = o.TerminationCause
	}
	if !isNil(o.VelocityEstimate) {
		toSerialize["velocityEstimate"] = o.VelocityEstimate
	}
	if !isNil(o.Altitude) {
		toSerialize["altitude"] = o.Altitude
	}
	if !isNil(o.AchievedQos) {
		toSerialize["achievedQos"] = o.AchievedQos
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
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



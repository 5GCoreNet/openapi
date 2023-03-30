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

// checks if the LocUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocUpdateData{}

// LocUpdateData Contains the input parameters in LocationUpdate service operation
type LocUpdateData struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	PseudonymIndicator *PseudonymIndicator `json:"pseudonymIndicator,omitempty"`
	LocationRequestType LocationRequestType `json:"locationRequestType"`
	LocationEstimate GeographicArea `json:"locationEstimate"`
	// Indicates value of the age of the location estimate.
	AgeOfLocationEstimate int32 `json:"ageOfLocationEstimate"`
	// string with format 'date-time' as defined in OpenAPI.
	TimestampOfLocationEstimate *time.Time `json:"timestampOfLocationEstimate,omitempty"`
	AccuracyFulfilmentIndicator AccuracyFulfilmentIndicator `json:"accuracyFulfilmentIndicator"`
	CivicAddress *CivicAddress `json:"civicAddress,omitempty"`
	LcsQosClass LcsQosClass `json:"lcsQosClass"`
	// Contains the external client identification
	ExternalClientIdentification *string `json:"externalClientIdentification,omitempty"`
	AfId *string `json:"afId,omitempty"`
	GmlcNumber *string `json:"gmlcNumber,omitempty"`
	// LCS Service Type Id.
	LcsServiceType *int32 `json:"lcsServiceType,omitempty"`
}

// NewLocUpdateData instantiates a new LocUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocUpdateData(locationRequestType LocationRequestType, locationEstimate GeographicArea, ageOfLocationEstimate int32, accuracyFulfilmentIndicator AccuracyFulfilmentIndicator, lcsQosClass LcsQosClass) *LocUpdateData {
	this := LocUpdateData{}
	this.LocationRequestType = locationRequestType
	this.LocationEstimate = locationEstimate
	this.AgeOfLocationEstimate = ageOfLocationEstimate
	this.AccuracyFulfilmentIndicator = accuracyFulfilmentIndicator
	this.LcsQosClass = lcsQosClass
	return &this
}

// NewLocUpdateDataWithDefaults instantiates a new LocUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocUpdateDataWithDefaults() *LocUpdateData {
	this := LocUpdateData{}
	return &this
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *LocUpdateData) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *LocUpdateData) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *LocUpdateData) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *LocUpdateData) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *LocUpdateData) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *LocUpdateData) SetSupi(v string) {
	o.Supi = &v
}

// GetPseudonymIndicator returns the PseudonymIndicator field value if set, zero value otherwise.
func (o *LocUpdateData) GetPseudonymIndicator() PseudonymIndicator {
	if o == nil || IsNil(o.PseudonymIndicator) {
		var ret PseudonymIndicator
		return ret
	}
	return *o.PseudonymIndicator
}

// GetPseudonymIndicatorOk returns a tuple with the PseudonymIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetPseudonymIndicatorOk() (*PseudonymIndicator, bool) {
	if o == nil || IsNil(o.PseudonymIndicator) {
		return nil, false
	}
	return o.PseudonymIndicator, true
}

// HasPseudonymIndicator returns a boolean if a field has been set.
func (o *LocUpdateData) HasPseudonymIndicator() bool {
	if o != nil && !IsNil(o.PseudonymIndicator) {
		return true
	}

	return false
}

// SetPseudonymIndicator gets a reference to the given PseudonymIndicator and assigns it to the PseudonymIndicator field.
func (o *LocUpdateData) SetPseudonymIndicator(v PseudonymIndicator) {
	o.PseudonymIndicator = &v
}

// GetLocationRequestType returns the LocationRequestType field value
func (o *LocUpdateData) GetLocationRequestType() LocationRequestType {
	if o == nil {
		var ret LocationRequestType
		return ret
	}

	return o.LocationRequestType
}

// GetLocationRequestTypeOk returns a tuple with the LocationRequestType field value
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetLocationRequestTypeOk() (*LocationRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationRequestType, true
}

// SetLocationRequestType sets field value
func (o *LocUpdateData) SetLocationRequestType(v LocationRequestType) {
	o.LocationRequestType = v
}

// GetLocationEstimate returns the LocationEstimate field value
func (o *LocUpdateData) GetLocationEstimate() GeographicArea {
	if o == nil {
		var ret GeographicArea
		return ret
	}

	return o.LocationEstimate
}

// GetLocationEstimateOk returns a tuple with the LocationEstimate field value
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetLocationEstimateOk() (*GeographicArea, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationEstimate, true
}

// SetLocationEstimate sets field value
func (o *LocUpdateData) SetLocationEstimate(v GeographicArea) {
	o.LocationEstimate = v
}

// GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field value
func (o *LocUpdateData) GetAgeOfLocationEstimate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AgeOfLocationEstimate
}

// GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field value
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetAgeOfLocationEstimateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgeOfLocationEstimate, true
}

// SetAgeOfLocationEstimate sets field value
func (o *LocUpdateData) SetAgeOfLocationEstimate(v int32) {
	o.AgeOfLocationEstimate = v
}

// GetTimestampOfLocationEstimate returns the TimestampOfLocationEstimate field value if set, zero value otherwise.
func (o *LocUpdateData) GetTimestampOfLocationEstimate() time.Time {
	if o == nil || IsNil(o.TimestampOfLocationEstimate) {
		var ret time.Time
		return ret
	}
	return *o.TimestampOfLocationEstimate
}

// GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetTimestampOfLocationEstimateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampOfLocationEstimate) {
		return nil, false
	}
	return o.TimestampOfLocationEstimate, true
}

// HasTimestampOfLocationEstimate returns a boolean if a field has been set.
func (o *LocUpdateData) HasTimestampOfLocationEstimate() bool {
	if o != nil && !IsNil(o.TimestampOfLocationEstimate) {
		return true
	}

	return false
}

// SetTimestampOfLocationEstimate gets a reference to the given time.Time and assigns it to the TimestampOfLocationEstimate field.
func (o *LocUpdateData) SetTimestampOfLocationEstimate(v time.Time) {
	o.TimestampOfLocationEstimate = &v
}

// GetAccuracyFulfilmentIndicator returns the AccuracyFulfilmentIndicator field value
func (o *LocUpdateData) GetAccuracyFulfilmentIndicator() AccuracyFulfilmentIndicator {
	if o == nil {
		var ret AccuracyFulfilmentIndicator
		return ret
	}

	return o.AccuracyFulfilmentIndicator
}

// GetAccuracyFulfilmentIndicatorOk returns a tuple with the AccuracyFulfilmentIndicator field value
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetAccuracyFulfilmentIndicatorOk() (*AccuracyFulfilmentIndicator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccuracyFulfilmentIndicator, true
}

// SetAccuracyFulfilmentIndicator sets field value
func (o *LocUpdateData) SetAccuracyFulfilmentIndicator(v AccuracyFulfilmentIndicator) {
	o.AccuracyFulfilmentIndicator = v
}

// GetCivicAddress returns the CivicAddress field value if set, zero value otherwise.
func (o *LocUpdateData) GetCivicAddress() CivicAddress {
	if o == nil || IsNil(o.CivicAddress) {
		var ret CivicAddress
		return ret
	}
	return *o.CivicAddress
}

// GetCivicAddressOk returns a tuple with the CivicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetCivicAddressOk() (*CivicAddress, bool) {
	if o == nil || IsNil(o.CivicAddress) {
		return nil, false
	}
	return o.CivicAddress, true
}

// HasCivicAddress returns a boolean if a field has been set.
func (o *LocUpdateData) HasCivicAddress() bool {
	if o != nil && !IsNil(o.CivicAddress) {
		return true
	}

	return false
}

// SetCivicAddress gets a reference to the given CivicAddress and assigns it to the CivicAddress field.
func (o *LocUpdateData) SetCivicAddress(v CivicAddress) {
	o.CivicAddress = &v
}

// GetLcsQosClass returns the LcsQosClass field value
func (o *LocUpdateData) GetLcsQosClass() LcsQosClass {
	if o == nil {
		var ret LcsQosClass
		return ret
	}

	return o.LcsQosClass
}

// GetLcsQosClassOk returns a tuple with the LcsQosClass field value
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetLcsQosClassOk() (*LcsQosClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LcsQosClass, true
}

// SetLcsQosClass sets field value
func (o *LocUpdateData) SetLcsQosClass(v LcsQosClass) {
	o.LcsQosClass = v
}

// GetExternalClientIdentification returns the ExternalClientIdentification field value if set, zero value otherwise.
func (o *LocUpdateData) GetExternalClientIdentification() string {
	if o == nil || IsNil(o.ExternalClientIdentification) {
		var ret string
		return ret
	}
	return *o.ExternalClientIdentification
}

// GetExternalClientIdentificationOk returns a tuple with the ExternalClientIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetExternalClientIdentificationOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalClientIdentification) {
		return nil, false
	}
	return o.ExternalClientIdentification, true
}

// HasExternalClientIdentification returns a boolean if a field has been set.
func (o *LocUpdateData) HasExternalClientIdentification() bool {
	if o != nil && !IsNil(o.ExternalClientIdentification) {
		return true
	}

	return false
}

// SetExternalClientIdentification gets a reference to the given string and assigns it to the ExternalClientIdentification field.
func (o *LocUpdateData) SetExternalClientIdentification(v string) {
	o.ExternalClientIdentification = &v
}

// GetAfId returns the AfId field value if set, zero value otherwise.
func (o *LocUpdateData) GetAfId() string {
	if o == nil || IsNil(o.AfId) {
		var ret string
		return ret
	}
	return *o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetAfIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfId) {
		return nil, false
	}
	return o.AfId, true
}

// HasAfId returns a boolean if a field has been set.
func (o *LocUpdateData) HasAfId() bool {
	if o != nil && !IsNil(o.AfId) {
		return true
	}

	return false
}

// SetAfId gets a reference to the given string and assigns it to the AfId field.
func (o *LocUpdateData) SetAfId(v string) {
	o.AfId = &v
}

// GetGmlcNumber returns the GmlcNumber field value if set, zero value otherwise.
func (o *LocUpdateData) GetGmlcNumber() string {
	if o == nil || IsNil(o.GmlcNumber) {
		var ret string
		return ret
	}
	return *o.GmlcNumber
}

// GetGmlcNumberOk returns a tuple with the GmlcNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetGmlcNumberOk() (*string, bool) {
	if o == nil || IsNil(o.GmlcNumber) {
		return nil, false
	}
	return o.GmlcNumber, true
}

// HasGmlcNumber returns a boolean if a field has been set.
func (o *LocUpdateData) HasGmlcNumber() bool {
	if o != nil && !IsNil(o.GmlcNumber) {
		return true
	}

	return false
}

// SetGmlcNumber gets a reference to the given string and assigns it to the GmlcNumber field.
func (o *LocUpdateData) SetGmlcNumber(v string) {
	o.GmlcNumber = &v
}

// GetLcsServiceType returns the LcsServiceType field value if set, zero value otherwise.
func (o *LocUpdateData) GetLcsServiceType() int32 {
	if o == nil || IsNil(o.LcsServiceType) {
		var ret int32
		return ret
	}
	return *o.LcsServiceType
}

// GetLcsServiceTypeOk returns a tuple with the LcsServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetLcsServiceTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.LcsServiceType) {
		return nil, false
	}
	return o.LcsServiceType, true
}

// HasLcsServiceType returns a boolean if a field has been set.
func (o *LocUpdateData) HasLcsServiceType() bool {
	if o != nil && !IsNil(o.LcsServiceType) {
		return true
	}

	return false
}

// SetLcsServiceType gets a reference to the given int32 and assigns it to the LcsServiceType field.
func (o *LocUpdateData) SetLcsServiceType(v int32) {
	o.LcsServiceType = &v
}

func (o LocUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.PseudonymIndicator) {
		toSerialize["pseudonymIndicator"] = o.PseudonymIndicator
	}
	toSerialize["locationRequestType"] = o.LocationRequestType
	toSerialize["locationEstimate"] = o.LocationEstimate
	toSerialize["ageOfLocationEstimate"] = o.AgeOfLocationEstimate
	if !IsNil(o.TimestampOfLocationEstimate) {
		toSerialize["timestampOfLocationEstimate"] = o.TimestampOfLocationEstimate
	}
	toSerialize["accuracyFulfilmentIndicator"] = o.AccuracyFulfilmentIndicator
	if !IsNil(o.CivicAddress) {
		toSerialize["civicAddress"] = o.CivicAddress
	}
	toSerialize["lcsQosClass"] = o.LcsQosClass
	if !IsNil(o.ExternalClientIdentification) {
		toSerialize["externalClientIdentification"] = o.ExternalClientIdentification
	}
	if !IsNil(o.AfId) {
		toSerialize["afId"] = o.AfId
	}
	if !IsNil(o.GmlcNumber) {
		toSerialize["gmlcNumber"] = o.GmlcNumber
	}
	if !IsNil(o.LcsServiceType) {
		toSerialize["lcsServiceType"] = o.LcsServiceType
	}
	return toSerialize, nil
}

type NullableLocUpdateData struct {
	value *LocUpdateData
	isSet bool
}

func (v NullableLocUpdateData) Get() *LocUpdateData {
	return v.value
}

func (v *NullableLocUpdateData) Set(val *LocUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableLocUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableLocUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocUpdateData(val *LocUpdateData) *NullableLocUpdateData {
	return &NullableLocUpdateData{value: val, isSet: true}
}

func (v NullableLocUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


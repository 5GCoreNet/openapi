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

// checks if the EutraLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EutraLocation{}

// EutraLocation Contains the E-UTRA user location.
type EutraLocation struct {
	Tai Tai `json:"tai"`
	IgnoreTai *bool `json:"ignoreTai,omitempty"`
	Ecgi Ecgi `json:"ecgi"`
	// This flag when present shall indicate that the Ecgi shall be ignored When present, it shall be set as follows: - true: ecgi shall be ignored. - false (default): ecgi shall not be ignored. 
	IgnoreEcgi *bool `json:"ignoreEcgi,omitempty"`
	// The value represents the elapsed time in minutes since the last network contact of the mobile station.  Value \"0\" indicates that the location information was obtained after a successful paging procedure for Active Location Retrieval when the UE is in idle mode or after a successful NG-RAN location reporting procedure with the eNB when the UE is in connected mode.  Any other value than \"0\" indicates that the location information is the last known one.  See 3GPP TS 29.002 clause 17.7.8. 
	AgeOfLocationInformation *int32 `json:"ageOfLocationInformation,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	UeLocationTimestamp *time.Time `json:"ueLocationTimestamp,omitempty"`
	// Refer to geographical Information. See 3GPP TS 23.032 clause 7.3.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used. 
	GeographicalInformation *string `json:"geographicalInformation,omitempty"`
	// Refers to Calling Geodetic Location. See ITU-T Recommendation Q.763 (1999) [24] clause 3.88.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used. 
	GeodeticInformation *string `json:"geodeticInformation,omitempty"`
	GlobalNgenbId *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
	GlobalENbId *GlobalRanNodeId `json:"globalENbId,omitempty"`
}

// NewEutraLocation instantiates a new EutraLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEutraLocation(tai Tai, ecgi Ecgi) *EutraLocation {
	this := EutraLocation{}
	this.Tai = tai
	var ignoreTai bool = false
	this.IgnoreTai = &ignoreTai
	this.Ecgi = ecgi
	var ignoreEcgi bool = false
	this.IgnoreEcgi = &ignoreEcgi
	return &this
}

// NewEutraLocationWithDefaults instantiates a new EutraLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEutraLocationWithDefaults() *EutraLocation {
	this := EutraLocation{}
	var ignoreTai bool = false
	this.IgnoreTai = &ignoreTai
	var ignoreEcgi bool = false
	this.IgnoreEcgi = &ignoreEcgi
	return &this
}

// GetTai returns the Tai field value
func (o *EutraLocation) GetTai() Tai {
	if o == nil {
		var ret Tai
		return ret
	}

	return o.Tai
}

// GetTaiOk returns a tuple with the Tai field value
// and a boolean to check if the value has been set.
func (o *EutraLocation) GetTaiOk() (*Tai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tai, true
}

// SetTai sets field value
func (o *EutraLocation) SetTai(v Tai) {
	o.Tai = v
}

// GetIgnoreTai returns the IgnoreTai field value if set, zero value otherwise.
func (o *EutraLocation) GetIgnoreTai() bool {
	if o == nil || isNil(o.IgnoreTai) {
		var ret bool
		return ret
	}
	return *o.IgnoreTai
}

// GetIgnoreTaiOk returns a tuple with the IgnoreTai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EutraLocation) GetIgnoreTaiOk() (*bool, bool) {
	if o == nil || isNil(o.IgnoreTai) {
		return nil, false
	}
	return o.IgnoreTai, true
}

// HasIgnoreTai returns a boolean if a field has been set.
func (o *EutraLocation) HasIgnoreTai() bool {
	if o != nil && !isNil(o.IgnoreTai) {
		return true
	}

	return false
}

// SetIgnoreTai gets a reference to the given bool and assigns it to the IgnoreTai field.
func (o *EutraLocation) SetIgnoreTai(v bool) {
	o.IgnoreTai = &v
}

// GetEcgi returns the Ecgi field value
func (o *EutraLocation) GetEcgi() Ecgi {
	if o == nil {
		var ret Ecgi
		return ret
	}

	return o.Ecgi
}

// GetEcgiOk returns a tuple with the Ecgi field value
// and a boolean to check if the value has been set.
func (o *EutraLocation) GetEcgiOk() (*Ecgi, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ecgi, true
}

// SetEcgi sets field value
func (o *EutraLocation) SetEcgi(v Ecgi) {
	o.Ecgi = v
}

// GetIgnoreEcgi returns the IgnoreEcgi field value if set, zero value otherwise.
func (o *EutraLocation) GetIgnoreEcgi() bool {
	if o == nil || isNil(o.IgnoreEcgi) {
		var ret bool
		return ret
	}
	return *o.IgnoreEcgi
}

// GetIgnoreEcgiOk returns a tuple with the IgnoreEcgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EutraLocation) GetIgnoreEcgiOk() (*bool, bool) {
	if o == nil || isNil(o.IgnoreEcgi) {
		return nil, false
	}
	return o.IgnoreEcgi, true
}

// HasIgnoreEcgi returns a boolean if a field has been set.
func (o *EutraLocation) HasIgnoreEcgi() bool {
	if o != nil && !isNil(o.IgnoreEcgi) {
		return true
	}

	return false
}

// SetIgnoreEcgi gets a reference to the given bool and assigns it to the IgnoreEcgi field.
func (o *EutraLocation) SetIgnoreEcgi(v bool) {
	o.IgnoreEcgi = &v
}

// GetAgeOfLocationInformation returns the AgeOfLocationInformation field value if set, zero value otherwise.
func (o *EutraLocation) GetAgeOfLocationInformation() int32 {
	if o == nil || isNil(o.AgeOfLocationInformation) {
		var ret int32
		return ret
	}
	return *o.AgeOfLocationInformation
}

// GetAgeOfLocationInformationOk returns a tuple with the AgeOfLocationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EutraLocation) GetAgeOfLocationInformationOk() (*int32, bool) {
	if o == nil || isNil(o.AgeOfLocationInformation) {
		return nil, false
	}
	return o.AgeOfLocationInformation, true
}

// HasAgeOfLocationInformation returns a boolean if a field has been set.
func (o *EutraLocation) HasAgeOfLocationInformation() bool {
	if o != nil && !isNil(o.AgeOfLocationInformation) {
		return true
	}

	return false
}

// SetAgeOfLocationInformation gets a reference to the given int32 and assigns it to the AgeOfLocationInformation field.
func (o *EutraLocation) SetAgeOfLocationInformation(v int32) {
	o.AgeOfLocationInformation = &v
}

// GetUeLocationTimestamp returns the UeLocationTimestamp field value if set, zero value otherwise.
func (o *EutraLocation) GetUeLocationTimestamp() time.Time {
	if o == nil || isNil(o.UeLocationTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.UeLocationTimestamp
}

// GetUeLocationTimestampOk returns a tuple with the UeLocationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EutraLocation) GetUeLocationTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.UeLocationTimestamp) {
		return nil, false
	}
	return o.UeLocationTimestamp, true
}

// HasUeLocationTimestamp returns a boolean if a field has been set.
func (o *EutraLocation) HasUeLocationTimestamp() bool {
	if o != nil && !isNil(o.UeLocationTimestamp) {
		return true
	}

	return false
}

// SetUeLocationTimestamp gets a reference to the given time.Time and assigns it to the UeLocationTimestamp field.
func (o *EutraLocation) SetUeLocationTimestamp(v time.Time) {
	o.UeLocationTimestamp = &v
}

// GetGeographicalInformation returns the GeographicalInformation field value if set, zero value otherwise.
func (o *EutraLocation) GetGeographicalInformation() string {
	if o == nil || isNil(o.GeographicalInformation) {
		var ret string
		return ret
	}
	return *o.GeographicalInformation
}

// GetGeographicalInformationOk returns a tuple with the GeographicalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EutraLocation) GetGeographicalInformationOk() (*string, bool) {
	if o == nil || isNil(o.GeographicalInformation) {
		return nil, false
	}
	return o.GeographicalInformation, true
}

// HasGeographicalInformation returns a boolean if a field has been set.
func (o *EutraLocation) HasGeographicalInformation() bool {
	if o != nil && !isNil(o.GeographicalInformation) {
		return true
	}

	return false
}

// SetGeographicalInformation gets a reference to the given string and assigns it to the GeographicalInformation field.
func (o *EutraLocation) SetGeographicalInformation(v string) {
	o.GeographicalInformation = &v
}

// GetGeodeticInformation returns the GeodeticInformation field value if set, zero value otherwise.
func (o *EutraLocation) GetGeodeticInformation() string {
	if o == nil || isNil(o.GeodeticInformation) {
		var ret string
		return ret
	}
	return *o.GeodeticInformation
}

// GetGeodeticInformationOk returns a tuple with the GeodeticInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EutraLocation) GetGeodeticInformationOk() (*string, bool) {
	if o == nil || isNil(o.GeodeticInformation) {
		return nil, false
	}
	return o.GeodeticInformation, true
}

// HasGeodeticInformation returns a boolean if a field has been set.
func (o *EutraLocation) HasGeodeticInformation() bool {
	if o != nil && !isNil(o.GeodeticInformation) {
		return true
	}

	return false
}

// SetGeodeticInformation gets a reference to the given string and assigns it to the GeodeticInformation field.
func (o *EutraLocation) SetGeodeticInformation(v string) {
	o.GeodeticInformation = &v
}

// GetGlobalNgenbId returns the GlobalNgenbId field value if set, zero value otherwise.
func (o *EutraLocation) GetGlobalNgenbId() GlobalRanNodeId {
	if o == nil || isNil(o.GlobalNgenbId) {
		var ret GlobalRanNodeId
		return ret
	}
	return *o.GlobalNgenbId
}

// GetGlobalNgenbIdOk returns a tuple with the GlobalNgenbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EutraLocation) GetGlobalNgenbIdOk() (*GlobalRanNodeId, bool) {
	if o == nil || isNil(o.GlobalNgenbId) {
		return nil, false
	}
	return o.GlobalNgenbId, true
}

// HasGlobalNgenbId returns a boolean if a field has been set.
func (o *EutraLocation) HasGlobalNgenbId() bool {
	if o != nil && !isNil(o.GlobalNgenbId) {
		return true
	}

	return false
}

// SetGlobalNgenbId gets a reference to the given GlobalRanNodeId and assigns it to the GlobalNgenbId field.
func (o *EutraLocation) SetGlobalNgenbId(v GlobalRanNodeId) {
	o.GlobalNgenbId = &v
}

// GetGlobalENbId returns the GlobalENbId field value if set, zero value otherwise.
func (o *EutraLocation) GetGlobalENbId() GlobalRanNodeId {
	if o == nil || isNil(o.GlobalENbId) {
		var ret GlobalRanNodeId
		return ret
	}
	return *o.GlobalENbId
}

// GetGlobalENbIdOk returns a tuple with the GlobalENbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EutraLocation) GetGlobalENbIdOk() (*GlobalRanNodeId, bool) {
	if o == nil || isNil(o.GlobalENbId) {
		return nil, false
	}
	return o.GlobalENbId, true
}

// HasGlobalENbId returns a boolean if a field has been set.
func (o *EutraLocation) HasGlobalENbId() bool {
	if o != nil && !isNil(o.GlobalENbId) {
		return true
	}

	return false
}

// SetGlobalENbId gets a reference to the given GlobalRanNodeId and assigns it to the GlobalENbId field.
func (o *EutraLocation) SetGlobalENbId(v GlobalRanNodeId) {
	o.GlobalENbId = &v
}

func (o EutraLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EutraLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tai"] = o.Tai
	if !isNil(o.IgnoreTai) {
		toSerialize["ignoreTai"] = o.IgnoreTai
	}
	toSerialize["ecgi"] = o.Ecgi
	if !isNil(o.IgnoreEcgi) {
		toSerialize["ignoreEcgi"] = o.IgnoreEcgi
	}
	if !isNil(o.AgeOfLocationInformation) {
		toSerialize["ageOfLocationInformation"] = o.AgeOfLocationInformation
	}
	if !isNil(o.UeLocationTimestamp) {
		toSerialize["ueLocationTimestamp"] = o.UeLocationTimestamp
	}
	if !isNil(o.GeographicalInformation) {
		toSerialize["geographicalInformation"] = o.GeographicalInformation
	}
	if !isNil(o.GeodeticInformation) {
		toSerialize["geodeticInformation"] = o.GeodeticInformation
	}
	if !isNil(o.GlobalNgenbId) {
		toSerialize["globalNgenbId"] = o.GlobalNgenbId
	}
	if !isNil(o.GlobalENbId) {
		toSerialize["globalENbId"] = o.GlobalENbId
	}
	return toSerialize, nil
}

type NullableEutraLocation struct {
	value *EutraLocation
	isSet bool
}

func (v NullableEutraLocation) Get() *EutraLocation {
	return v.value
}

func (v *NullableEutraLocation) Set(val *EutraLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableEutraLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableEutraLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEutraLocation(val *EutraLocation) *NullableEutraLocation {
	return &NullableEutraLocation{value: val, isSet: true}
}

func (v NullableEutraLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEutraLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



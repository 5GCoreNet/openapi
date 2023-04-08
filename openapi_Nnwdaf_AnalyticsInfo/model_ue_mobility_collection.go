/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// checks if the UeMobilityCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeMobilityCollection{}

// UeMobilityCollection Contains UE mobility information associated with an application. If the allAppInd attribute  is present and set to true, then the value in the appId shall be ignored, which indicates  the collected UE mobility information is applicable to all the applications for the UE. 
type UeMobilityCollection struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	// String providing an application identifier.
	AppId string `json:"appId"`
	// Indicates applicable to all applications if set to true, otherwise set to false.  Default value is false if omitted. 
	AllAppInd *bool `json:"allAppInd,omitempty"`
	UeTrajs []UeTrajectoryCollection `json:"ueTrajs"`
}

// NewUeMobilityCollection instantiates a new UeMobilityCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeMobilityCollection(appId string, ueTrajs []UeTrajectoryCollection) *UeMobilityCollection {
	this := UeMobilityCollection{}
	this.AppId = appId
	this.UeTrajs = ueTrajs
	return &this
}

// NewUeMobilityCollectionWithDefaults instantiates a new UeMobilityCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeMobilityCollectionWithDefaults() *UeMobilityCollection {
	this := UeMobilityCollection{}
	return &this
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *UeMobilityCollection) GetGpsi() string {
	if o == nil || isNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobilityCollection) GetGpsiOk() (*string, bool) {
	if o == nil || isNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *UeMobilityCollection) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *UeMobilityCollection) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *UeMobilityCollection) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobilityCollection) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *UeMobilityCollection) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *UeMobilityCollection) SetSupi(v string) {
	o.Supi = &v
}

// GetAppId returns the AppId field value
func (o *UeMobilityCollection) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *UeMobilityCollection) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *UeMobilityCollection) SetAppId(v string) {
	o.AppId = v
}

// GetAllAppInd returns the AllAppInd field value if set, zero value otherwise.
func (o *UeMobilityCollection) GetAllAppInd() bool {
	if o == nil || isNil(o.AllAppInd) {
		var ret bool
		return ret
	}
	return *o.AllAppInd
}

// GetAllAppIndOk returns a tuple with the AllAppInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobilityCollection) GetAllAppIndOk() (*bool, bool) {
	if o == nil || isNil(o.AllAppInd) {
		return nil, false
	}
	return o.AllAppInd, true
}

// HasAllAppInd returns a boolean if a field has been set.
func (o *UeMobilityCollection) HasAllAppInd() bool {
	if o != nil && !isNil(o.AllAppInd) {
		return true
	}

	return false
}

// SetAllAppInd gets a reference to the given bool and assigns it to the AllAppInd field.
func (o *UeMobilityCollection) SetAllAppInd(v bool) {
	o.AllAppInd = &v
}

// GetUeTrajs returns the UeTrajs field value
func (o *UeMobilityCollection) GetUeTrajs() []UeTrajectoryCollection {
	if o == nil {
		var ret []UeTrajectoryCollection
		return ret
	}

	return o.UeTrajs
}

// GetUeTrajsOk returns a tuple with the UeTrajs field value
// and a boolean to check if the value has been set.
func (o *UeMobilityCollection) GetUeTrajsOk() ([]UeTrajectoryCollection, bool) {
	if o == nil {
		return nil, false
	}
	return o.UeTrajs, true
}

// SetUeTrajs sets field value
func (o *UeMobilityCollection) SetUeTrajs(v []UeTrajectoryCollection) {
	o.UeTrajs = v
}

func (o UeMobilityCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeMobilityCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	toSerialize["appId"] = o.AppId
	if !isNil(o.AllAppInd) {
		toSerialize["allAppInd"] = o.AllAppInd
	}
	toSerialize["ueTrajs"] = o.UeTrajs
	return toSerialize, nil
}

type NullableUeMobilityCollection struct {
	value *UeMobilityCollection
	isSet bool
}

func (v NullableUeMobilityCollection) Get() *UeMobilityCollection {
	return v.value
}

func (v *NullableUeMobilityCollection) Set(val *UeMobilityCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableUeMobilityCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableUeMobilityCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeMobilityCollection(val *UeMobilityCollection) *NullableUeMobilityCollection {
	return &NullableUeMobilityCollection{value: val, isSet: true}
}

func (v NullableUeMobilityCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeMobilityCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



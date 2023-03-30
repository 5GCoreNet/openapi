/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// checks if the UeCommunicationCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeCommunicationCollection{}

// UeCommunicationCollection Contains UE communication information associated with an application.
type UeCommunicationCollection struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	ExterGroupId *string `json:"exterGroupId,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	InterGroupId *string `json:"interGroupId,omitempty"`
	// String providing an application identifier.
	AppId string `json:"appId"`
	Comms []CommunicationCollection `json:"comms"`
}

// NewUeCommunicationCollection instantiates a new UeCommunicationCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeCommunicationCollection(appId string, comms []CommunicationCollection) *UeCommunicationCollection {
	this := UeCommunicationCollection{}
	this.AppId = appId
	this.Comms = comms
	return &this
}

// NewUeCommunicationCollectionWithDefaults instantiates a new UeCommunicationCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeCommunicationCollectionWithDefaults() *UeCommunicationCollection {
	this := UeCommunicationCollection{}
	return &this
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *UeCommunicationCollection) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunicationCollection) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *UeCommunicationCollection) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *UeCommunicationCollection) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *UeCommunicationCollection) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunicationCollection) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *UeCommunicationCollection) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *UeCommunicationCollection) SetSupi(v string) {
	o.Supi = &v
}

// GetExterGroupId returns the ExterGroupId field value if set, zero value otherwise.
func (o *UeCommunicationCollection) GetExterGroupId() string {
	if o == nil || IsNil(o.ExterGroupId) {
		var ret string
		return ret
	}
	return *o.ExterGroupId
}

// GetExterGroupIdOk returns a tuple with the ExterGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunicationCollection) GetExterGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExterGroupId) {
		return nil, false
	}
	return o.ExterGroupId, true
}

// HasExterGroupId returns a boolean if a field has been set.
func (o *UeCommunicationCollection) HasExterGroupId() bool {
	if o != nil && !IsNil(o.ExterGroupId) {
		return true
	}

	return false
}

// SetExterGroupId gets a reference to the given string and assigns it to the ExterGroupId field.
func (o *UeCommunicationCollection) SetExterGroupId(v string) {
	o.ExterGroupId = &v
}

// GetInterGroupId returns the InterGroupId field value if set, zero value otherwise.
func (o *UeCommunicationCollection) GetInterGroupId() string {
	if o == nil || IsNil(o.InterGroupId) {
		var ret string
		return ret
	}
	return *o.InterGroupId
}

// GetInterGroupIdOk returns a tuple with the InterGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunicationCollection) GetInterGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.InterGroupId) {
		return nil, false
	}
	return o.InterGroupId, true
}

// HasInterGroupId returns a boolean if a field has been set.
func (o *UeCommunicationCollection) HasInterGroupId() bool {
	if o != nil && !IsNil(o.InterGroupId) {
		return true
	}

	return false
}

// SetInterGroupId gets a reference to the given string and assigns it to the InterGroupId field.
func (o *UeCommunicationCollection) SetInterGroupId(v string) {
	o.InterGroupId = &v
}

// GetAppId returns the AppId field value
func (o *UeCommunicationCollection) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *UeCommunicationCollection) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *UeCommunicationCollection) SetAppId(v string) {
	o.AppId = v
}

// GetComms returns the Comms field value
func (o *UeCommunicationCollection) GetComms() []CommunicationCollection {
	if o == nil {
		var ret []CommunicationCollection
		return ret
	}

	return o.Comms
}

// GetCommsOk returns a tuple with the Comms field value
// and a boolean to check if the value has been set.
func (o *UeCommunicationCollection) GetCommsOk() ([]CommunicationCollection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comms, true
}

// SetComms sets field value
func (o *UeCommunicationCollection) SetComms(v []CommunicationCollection) {
	o.Comms = v
}

func (o UeCommunicationCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeCommunicationCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.ExterGroupId) {
		toSerialize["exterGroupId"] = o.ExterGroupId
	}
	if !IsNil(o.InterGroupId) {
		toSerialize["interGroupId"] = o.InterGroupId
	}
	toSerialize["appId"] = o.AppId
	toSerialize["comms"] = o.Comms
	return toSerialize, nil
}

type NullableUeCommunicationCollection struct {
	value *UeCommunicationCollection
	isSet bool
}

func (v NullableUeCommunicationCollection) Get() *UeCommunicationCollection {
	return v.value
}

func (v *NullableUeCommunicationCollection) Set(val *UeCommunicationCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableUeCommunicationCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableUeCommunicationCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeCommunicationCollection(val *UeCommunicationCollection) *NullableUeCommunicationCollection {
	return &NullableUeCommunicationCollection{value: val, isSet: true}
}

func (v NullableUeCommunicationCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeCommunicationCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



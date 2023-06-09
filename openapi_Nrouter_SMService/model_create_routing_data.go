/*
Nrouter_SMService Service API

SMS Router SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nrouter_SMService

import (
	"encoding/json"
)

// checks if the CreateRoutingData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoutingData{}

// CreateRoutingData Information used for creating or updating the routing information of the user.
type CreateRoutingData struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SmsfId string `json:"smsfId"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi *string `json:"supi,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewCreateRoutingData instantiates a new CreateRoutingData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoutingData(smsfId string) *CreateRoutingData {
	this := CreateRoutingData{}
	this.SmsfId = smsfId
	return &this
}

// NewCreateRoutingDataWithDefaults instantiates a new CreateRoutingData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoutingDataWithDefaults() *CreateRoutingData {
	this := CreateRoutingData{}
	return &this
}

// GetSmsfId returns the SmsfId field value
func (o *CreateRoutingData) GetSmsfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmsfId
}

// GetSmsfIdOk returns a tuple with the SmsfId field value
// and a boolean to check if the value has been set.
func (o *CreateRoutingData) GetSmsfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmsfId, true
}

// SetSmsfId sets field value
func (o *CreateRoutingData) SetSmsfId(v string) {
	o.SmsfId = v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *CreateRoutingData) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoutingData) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *CreateRoutingData) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *CreateRoutingData) SetSupi(v string) {
	o.Supi = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *CreateRoutingData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoutingData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *CreateRoutingData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *CreateRoutingData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o CreateRoutingData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoutingData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["smsfId"] = o.SmsfId
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableCreateRoutingData struct {
	value *CreateRoutingData
	isSet bool
}

func (v NullableCreateRoutingData) Get() *CreateRoutingData {
	return v.value
}

func (v *NullableCreateRoutingData) Set(val *CreateRoutingData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoutingData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoutingData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoutingData(val *CreateRoutingData) *NullableCreateRoutingData {
	return &NullableCreateRoutingData{value: val, isSet: true}
}

func (v NullableCreateRoutingData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoutingData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

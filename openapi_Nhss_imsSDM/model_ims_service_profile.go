/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the ImsServiceProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImsServiceProfile{}

// ImsServiceProfile IMS Service Profile of the UE, containing the list of Public Identifiers and optionally a list of IFCs 
type ImsServiceProfile struct {
	PublicIdentifierList []PublicIdentifier `json:"publicIdentifierList"`
	Ifcs *Ifcs `json:"ifcs,omitempty"`
	CnServiceAuthorization *CoreNetworkServiceAuthorization `json:"cnServiceAuthorization,omitempty"`
}

// NewImsServiceProfile instantiates a new ImsServiceProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImsServiceProfile(publicIdentifierList []PublicIdentifier) *ImsServiceProfile {
	this := ImsServiceProfile{}
	this.PublicIdentifierList = publicIdentifierList
	return &this
}

// NewImsServiceProfileWithDefaults instantiates a new ImsServiceProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImsServiceProfileWithDefaults() *ImsServiceProfile {
	this := ImsServiceProfile{}
	return &this
}

// GetPublicIdentifierList returns the PublicIdentifierList field value
func (o *ImsServiceProfile) GetPublicIdentifierList() []PublicIdentifier {
	if o == nil {
		var ret []PublicIdentifier
		return ret
	}

	return o.PublicIdentifierList
}

// GetPublicIdentifierListOk returns a tuple with the PublicIdentifierList field value
// and a boolean to check if the value has been set.
func (o *ImsServiceProfile) GetPublicIdentifierListOk() ([]PublicIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicIdentifierList, true
}

// SetPublicIdentifierList sets field value
func (o *ImsServiceProfile) SetPublicIdentifierList(v []PublicIdentifier) {
	o.PublicIdentifierList = v
}

// GetIfcs returns the Ifcs field value if set, zero value otherwise.
func (o *ImsServiceProfile) GetIfcs() Ifcs {
	if o == nil || IsNil(o.Ifcs) {
		var ret Ifcs
		return ret
	}
	return *o.Ifcs
}

// GetIfcsOk returns a tuple with the Ifcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImsServiceProfile) GetIfcsOk() (*Ifcs, bool) {
	if o == nil || IsNil(o.Ifcs) {
		return nil, false
	}
	return o.Ifcs, true
}

// HasIfcs returns a boolean if a field has been set.
func (o *ImsServiceProfile) HasIfcs() bool {
	if o != nil && !IsNil(o.Ifcs) {
		return true
	}

	return false
}

// SetIfcs gets a reference to the given Ifcs and assigns it to the Ifcs field.
func (o *ImsServiceProfile) SetIfcs(v Ifcs) {
	o.Ifcs = &v
}

// GetCnServiceAuthorization returns the CnServiceAuthorization field value if set, zero value otherwise.
func (o *ImsServiceProfile) GetCnServiceAuthorization() CoreNetworkServiceAuthorization {
	if o == nil || IsNil(o.CnServiceAuthorization) {
		var ret CoreNetworkServiceAuthorization
		return ret
	}
	return *o.CnServiceAuthorization
}

// GetCnServiceAuthorizationOk returns a tuple with the CnServiceAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImsServiceProfile) GetCnServiceAuthorizationOk() (*CoreNetworkServiceAuthorization, bool) {
	if o == nil || IsNil(o.CnServiceAuthorization) {
		return nil, false
	}
	return o.CnServiceAuthorization, true
}

// HasCnServiceAuthorization returns a boolean if a field has been set.
func (o *ImsServiceProfile) HasCnServiceAuthorization() bool {
	if o != nil && !IsNil(o.CnServiceAuthorization) {
		return true
	}

	return false
}

// SetCnServiceAuthorization gets a reference to the given CoreNetworkServiceAuthorization and assigns it to the CnServiceAuthorization field.
func (o *ImsServiceProfile) SetCnServiceAuthorization(v CoreNetworkServiceAuthorization) {
	o.CnServiceAuthorization = &v
}

func (o ImsServiceProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImsServiceProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["publicIdentifierList"] = o.PublicIdentifierList
	if !IsNil(o.Ifcs) {
		toSerialize["ifcs"] = o.Ifcs
	}
	if !IsNil(o.CnServiceAuthorization) {
		toSerialize["cnServiceAuthorization"] = o.CnServiceAuthorization
	}
	return toSerialize, nil
}

type NullableImsServiceProfile struct {
	value *ImsServiceProfile
	isSet bool
}

func (v NullableImsServiceProfile) Get() *ImsServiceProfile {
	return v.value
}

func (v *NullableImsServiceProfile) Set(val *ImsServiceProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableImsServiceProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableImsServiceProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImsServiceProfile(val *ImsServiceProfile) *NullableImsServiceProfile {
	return &NullableImsServiceProfile{value: val, isSet: true}
}

func (v NullableImsServiceProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImsServiceProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



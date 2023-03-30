/*
Nnssaaf_NSSAA

Network Slice-Specific Authentication and Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssaaf_NSSAA

import (
	"encoding/json"
)

// checks if the SliceAuthReauthNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SliceAuthReauthNotification{}

// SliceAuthReauthNotification struct for SliceAuthReauthNotification
type SliceAuthReauthNotification struct {
	NotifType SliceAuthNotificationType `json:"notifType"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi"`
	Snssai Snssai `json:"snssai"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
}

// NewSliceAuthReauthNotification instantiates a new SliceAuthReauthNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSliceAuthReauthNotification(notifType SliceAuthNotificationType, gpsi string, snssai Snssai) *SliceAuthReauthNotification {
	this := SliceAuthReauthNotification{}
	this.NotifType = notifType
	this.Gpsi = gpsi
	this.Snssai = snssai
	return &this
}

// NewSliceAuthReauthNotificationWithDefaults instantiates a new SliceAuthReauthNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliceAuthReauthNotificationWithDefaults() *SliceAuthReauthNotification {
	this := SliceAuthReauthNotification{}
	return &this
}

// GetNotifType returns the NotifType field value
func (o *SliceAuthReauthNotification) GetNotifType() SliceAuthNotificationType {
	if o == nil {
		var ret SliceAuthNotificationType
		return ret
	}

	return o.NotifType
}

// GetNotifTypeOk returns a tuple with the NotifType field value
// and a boolean to check if the value has been set.
func (o *SliceAuthReauthNotification) GetNotifTypeOk() (*SliceAuthNotificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifType, true
}

// SetNotifType sets field value
func (o *SliceAuthReauthNotification) SetNotifType(v SliceAuthNotificationType) {
	o.NotifType = v
}

// GetGpsi returns the Gpsi field value
func (o *SliceAuthReauthNotification) GetGpsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value
// and a boolean to check if the value has been set.
func (o *SliceAuthReauthNotification) GetGpsiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gpsi, true
}

// SetGpsi sets field value
func (o *SliceAuthReauthNotification) SetGpsi(v string) {
	o.Gpsi = v
}

// GetSnssai returns the Snssai field value
func (o *SliceAuthReauthNotification) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *SliceAuthReauthNotification) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *SliceAuthReauthNotification) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *SliceAuthReauthNotification) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceAuthReauthNotification) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *SliceAuthReauthNotification) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *SliceAuthReauthNotification) SetSupi(v string) {
	o.Supi = &v
}

func (o SliceAuthReauthNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SliceAuthReauthNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifType"] = o.NotifType
	toSerialize["gpsi"] = o.Gpsi
	toSerialize["snssai"] = o.Snssai
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	return toSerialize, nil
}

type NullableSliceAuthReauthNotification struct {
	value *SliceAuthReauthNotification
	isSet bool
}

func (v NullableSliceAuthReauthNotification) Get() *SliceAuthReauthNotification {
	return v.value
}

func (v *NullableSliceAuthReauthNotification) Set(val *SliceAuthReauthNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableSliceAuthReauthNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableSliceAuthReauthNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliceAuthReauthNotification(val *SliceAuthReauthNotification) *NullableSliceAuthReauthNotification {
	return &NullableSliceAuthReauthNotification{value: val, isSet: true}
}

func (v NullableSliceAuthReauthNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliceAuthReauthNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



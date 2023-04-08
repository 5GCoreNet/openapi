/*
Nudm_UECM

Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UECM

import (
	"encoding/json"
)

// checks if the PcscfRestorationNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcscfRestorationNotification{}

// PcscfRestorationNotification struct for PcscfRestorationNotification
type PcscfRestorationNotification struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi"`
	FailedPcscf *PcscfAddress `json:"failedPcscf,omitempty"`
}

// NewPcscfRestorationNotification instantiates a new PcscfRestorationNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcscfRestorationNotification(supi string) *PcscfRestorationNotification {
	this := PcscfRestorationNotification{}
	this.Supi = supi
	return &this
}

// NewPcscfRestorationNotificationWithDefaults instantiates a new PcscfRestorationNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcscfRestorationNotificationWithDefaults() *PcscfRestorationNotification {
	this := PcscfRestorationNotification{}
	return &this
}

// GetSupi returns the Supi field value
func (o *PcscfRestorationNotification) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *PcscfRestorationNotification) GetSupiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *PcscfRestorationNotification) SetSupi(v string) {
	o.Supi = v
}

// GetFailedPcscf returns the FailedPcscf field value if set, zero value otherwise.
func (o *PcscfRestorationNotification) GetFailedPcscf() PcscfAddress {
	if o == nil || isNil(o.FailedPcscf) {
		var ret PcscfAddress
		return ret
	}
	return *o.FailedPcscf
}

// GetFailedPcscfOk returns a tuple with the FailedPcscf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfRestorationNotification) GetFailedPcscfOk() (*PcscfAddress, bool) {
	if o == nil || isNil(o.FailedPcscf) {
		return nil, false
	}
	return o.FailedPcscf, true
}

// HasFailedPcscf returns a boolean if a field has been set.
func (o *PcscfRestorationNotification) HasFailedPcscf() bool {
	if o != nil && !isNil(o.FailedPcscf) {
		return true
	}

	return false
}

// SetFailedPcscf gets a reference to the given PcscfAddress and assigns it to the FailedPcscf field.
func (o *PcscfRestorationNotification) SetFailedPcscf(v PcscfAddress) {
	o.FailedPcscf = &v
}

func (o PcscfRestorationNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcscfRestorationNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supi"] = o.Supi
	if !isNil(o.FailedPcscf) {
		toSerialize["failedPcscf"] = o.FailedPcscf
	}
	return toSerialize, nil
}

type NullablePcscfRestorationNotification struct {
	value *PcscfRestorationNotification
	isSet bool
}

func (v NullablePcscfRestorationNotification) Get() *PcscfRestorationNotification {
	return v.value
}

func (v *NullablePcscfRestorationNotification) Set(val *PcscfRestorationNotification) {
	v.value = val
	v.isSet = true
}

func (v NullablePcscfRestorationNotification) IsSet() bool {
	return v.isSet
}

func (v *NullablePcscfRestorationNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcscfRestorationNotification(val *PcscfRestorationNotification) *NullablePcscfRestorationNotification {
	return &NullablePcscfRestorationNotification{value: val, isSet: true}
}

func (v NullablePcscfRestorationNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcscfRestorationNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



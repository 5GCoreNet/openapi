/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// checks if the PduSessionEventNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PduSessionEventNotification{}

// PduSessionEventNotification Indicates PDU session information for the concerned established/terminated PDU session. 
type PduSessionEventNotification struct {
	EvNotif AfEventNotification `json:"evNotif"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	UeIpv4 *string `json:"ueIpv4,omitempty"`
	UeIpv6 *Ipv6Addr `json:"ueIpv6,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	UeMac *string `json:"ueMac,omitempty"`
	Status *PduSessionStatus `json:"status,omitempty"`
	PcfInfo *PcfAddressingInfo `json:"pcfInfo,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
}

// NewPduSessionEventNotification instantiates a new PduSessionEventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSessionEventNotification(evNotif AfEventNotification) *PduSessionEventNotification {
	this := PduSessionEventNotification{}
	this.EvNotif = evNotif
	return &this
}

// NewPduSessionEventNotificationWithDefaults instantiates a new PduSessionEventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSessionEventNotificationWithDefaults() *PduSessionEventNotification {
	this := PduSessionEventNotification{}
	return &this
}

// GetEvNotif returns the EvNotif field value
func (o *PduSessionEventNotification) GetEvNotif() AfEventNotification {
	if o == nil {
		var ret AfEventNotification
		return ret
	}

	return o.EvNotif
}

// GetEvNotifOk returns a tuple with the EvNotif field value
// and a boolean to check if the value has been set.
func (o *PduSessionEventNotification) GetEvNotifOk() (*AfEventNotification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvNotif, true
}

// SetEvNotif sets field value
func (o *PduSessionEventNotification) SetEvNotif(v AfEventNotification) {
	o.EvNotif = v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *PduSessionEventNotification) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionEventNotification) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *PduSessionEventNotification) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *PduSessionEventNotification) SetSupi(v string) {
	o.Supi = &v
}

// GetUeIpv4 returns the UeIpv4 field value if set, zero value otherwise.
func (o *PduSessionEventNotification) GetUeIpv4() string {
	if o == nil || isNil(o.UeIpv4) {
		var ret string
		return ret
	}
	return *o.UeIpv4
}

// GetUeIpv4Ok returns a tuple with the UeIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionEventNotification) GetUeIpv4Ok() (*string, bool) {
	if o == nil || isNil(o.UeIpv4) {
		return nil, false
	}
	return o.UeIpv4, true
}

// HasUeIpv4 returns a boolean if a field has been set.
func (o *PduSessionEventNotification) HasUeIpv4() bool {
	if o != nil && !isNil(o.UeIpv4) {
		return true
	}

	return false
}

// SetUeIpv4 gets a reference to the given string and assigns it to the UeIpv4 field.
func (o *PduSessionEventNotification) SetUeIpv4(v string) {
	o.UeIpv4 = &v
}

// GetUeIpv6 returns the UeIpv6 field value if set, zero value otherwise.
func (o *PduSessionEventNotification) GetUeIpv6() Ipv6Addr {
	if o == nil || isNil(o.UeIpv6) {
		var ret Ipv6Addr
		return ret
	}
	return *o.UeIpv6
}

// GetUeIpv6Ok returns a tuple with the UeIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionEventNotification) GetUeIpv6Ok() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.UeIpv6) {
		return nil, false
	}
	return o.UeIpv6, true
}

// HasUeIpv6 returns a boolean if a field has been set.
func (o *PduSessionEventNotification) HasUeIpv6() bool {
	if o != nil && !isNil(o.UeIpv6) {
		return true
	}

	return false
}

// SetUeIpv6 gets a reference to the given Ipv6Addr and assigns it to the UeIpv6 field.
func (o *PduSessionEventNotification) SetUeIpv6(v Ipv6Addr) {
	o.UeIpv6 = &v
}

// GetUeMac returns the UeMac field value if set, zero value otherwise.
func (o *PduSessionEventNotification) GetUeMac() string {
	if o == nil || isNil(o.UeMac) {
		var ret string
		return ret
	}
	return *o.UeMac
}

// GetUeMacOk returns a tuple with the UeMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionEventNotification) GetUeMacOk() (*string, bool) {
	if o == nil || isNil(o.UeMac) {
		return nil, false
	}
	return o.UeMac, true
}

// HasUeMac returns a boolean if a field has been set.
func (o *PduSessionEventNotification) HasUeMac() bool {
	if o != nil && !isNil(o.UeMac) {
		return true
	}

	return false
}

// SetUeMac gets a reference to the given string and assigns it to the UeMac field.
func (o *PduSessionEventNotification) SetUeMac(v string) {
	o.UeMac = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PduSessionEventNotification) GetStatus() PduSessionStatus {
	if o == nil || isNil(o.Status) {
		var ret PduSessionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionEventNotification) GetStatusOk() (*PduSessionStatus, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PduSessionEventNotification) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PduSessionStatus and assigns it to the Status field.
func (o *PduSessionEventNotification) SetStatus(v PduSessionStatus) {
	o.Status = &v
}

// GetPcfInfo returns the PcfInfo field value if set, zero value otherwise.
func (o *PduSessionEventNotification) GetPcfInfo() PcfAddressingInfo {
	if o == nil || isNil(o.PcfInfo) {
		var ret PcfAddressingInfo
		return ret
	}
	return *o.PcfInfo
}

// GetPcfInfoOk returns a tuple with the PcfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionEventNotification) GetPcfInfoOk() (*PcfAddressingInfo, bool) {
	if o == nil || isNil(o.PcfInfo) {
		return nil, false
	}
	return o.PcfInfo, true
}

// HasPcfInfo returns a boolean if a field has been set.
func (o *PduSessionEventNotification) HasPcfInfo() bool {
	if o != nil && !isNil(o.PcfInfo) {
		return true
	}

	return false
}

// SetPcfInfo gets a reference to the given PcfAddressingInfo and assigns it to the PcfInfo field.
func (o *PduSessionEventNotification) SetPcfInfo(v PcfAddressingInfo) {
	o.PcfInfo = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *PduSessionEventNotification) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionEventNotification) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *PduSessionEventNotification) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *PduSessionEventNotification) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *PduSessionEventNotification) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionEventNotification) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *PduSessionEventNotification) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *PduSessionEventNotification) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *PduSessionEventNotification) GetGpsi() string {
	if o == nil || isNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionEventNotification) GetGpsiOk() (*string, bool) {
	if o == nil || isNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *PduSessionEventNotification) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *PduSessionEventNotification) SetGpsi(v string) {
	o.Gpsi = &v
}

func (o PduSessionEventNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PduSessionEventNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["evNotif"] = o.EvNotif
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.UeIpv4) {
		toSerialize["ueIpv4"] = o.UeIpv4
	}
	if !isNil(o.UeIpv6) {
		toSerialize["ueIpv6"] = o.UeIpv6
	}
	if !isNil(o.UeMac) {
		toSerialize["ueMac"] = o.UeMac
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.PcfInfo) {
		toSerialize["pcfInfo"] = o.PcfInfo
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	return toSerialize, nil
}

type NullablePduSessionEventNotification struct {
	value *PduSessionEventNotification
	isSet bool
}

func (v NullablePduSessionEventNotification) Get() *PduSessionEventNotification {
	return v.value
}

func (v *NullablePduSessionEventNotification) Set(val *PduSessionEventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionEventNotification(val *PduSessionEventNotification) *NullablePduSessionEventNotification {
	return &NullablePduSessionEventNotification{value: val, isSet: true}
}

func (v NullablePduSessionEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



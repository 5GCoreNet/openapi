/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the N2SmInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &N2SmInformation{}

// N2SmInformation Represents the session management SMF related N2 information data part
type N2SmInformation struct {
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessionId int32 `json:"pduSessionId"`
	N2InfoContent *N2InfoContent `json:"n2InfoContent,omitempty"`
	SNssai *Snssai `json:"sNssai,omitempty"`
	HomePlmnSnssai *Snssai `json:"homePlmnSnssai,omitempty"`
	IwkSnssai *Snssai `json:"iwkSnssai,omitempty"`
	SubjectToHo *bool `json:"subjectToHo,omitempty"`
}

// NewN2SmInformation instantiates a new N2SmInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN2SmInformation(pduSessionId int32) *N2SmInformation {
	this := N2SmInformation{}
	this.PduSessionId = pduSessionId
	return &this
}

// NewN2SmInformationWithDefaults instantiates a new N2SmInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN2SmInformationWithDefaults() *N2SmInformation {
	this := N2SmInformation{}
	return &this
}

// GetPduSessionId returns the PduSessionId field value
func (o *N2SmInformation) GetPduSessionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value
// and a boolean to check if the value has been set.
func (o *N2SmInformation) GetPduSessionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PduSessionId, true
}

// SetPduSessionId sets field value
func (o *N2SmInformation) SetPduSessionId(v int32) {
	o.PduSessionId = v
}

// GetN2InfoContent returns the N2InfoContent field value if set, zero value otherwise.
func (o *N2SmInformation) GetN2InfoContent() N2InfoContent {
	if o == nil || IsNil(o.N2InfoContent) {
		var ret N2InfoContent
		return ret
	}
	return *o.N2InfoContent
}

// GetN2InfoContentOk returns a tuple with the N2InfoContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2SmInformation) GetN2InfoContentOk() (*N2InfoContent, bool) {
	if o == nil || IsNil(o.N2InfoContent) {
		return nil, false
	}
	return o.N2InfoContent, true
}

// HasN2InfoContent returns a boolean if a field has been set.
func (o *N2SmInformation) HasN2InfoContent() bool {
	if o != nil && !IsNil(o.N2InfoContent) {
		return true
	}

	return false
}

// SetN2InfoContent gets a reference to the given N2InfoContent and assigns it to the N2InfoContent field.
func (o *N2SmInformation) SetN2InfoContent(v N2InfoContent) {
	o.N2InfoContent = &v
}

// GetSNssai returns the SNssai field value if set, zero value otherwise.
func (o *N2SmInformation) GetSNssai() Snssai {
	if o == nil || IsNil(o.SNssai) {
		var ret Snssai
		return ret
	}
	return *o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2SmInformation) GetSNssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.SNssai) {
		return nil, false
	}
	return o.SNssai, true
}

// HasSNssai returns a boolean if a field has been set.
func (o *N2SmInformation) HasSNssai() bool {
	if o != nil && !IsNil(o.SNssai) {
		return true
	}

	return false
}

// SetSNssai gets a reference to the given Snssai and assigns it to the SNssai field.
func (o *N2SmInformation) SetSNssai(v Snssai) {
	o.SNssai = &v
}

// GetHomePlmnSnssai returns the HomePlmnSnssai field value if set, zero value otherwise.
func (o *N2SmInformation) GetHomePlmnSnssai() Snssai {
	if o == nil || IsNil(o.HomePlmnSnssai) {
		var ret Snssai
		return ret
	}
	return *o.HomePlmnSnssai
}

// GetHomePlmnSnssaiOk returns a tuple with the HomePlmnSnssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2SmInformation) GetHomePlmnSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.HomePlmnSnssai) {
		return nil, false
	}
	return o.HomePlmnSnssai, true
}

// HasHomePlmnSnssai returns a boolean if a field has been set.
func (o *N2SmInformation) HasHomePlmnSnssai() bool {
	if o != nil && !IsNil(o.HomePlmnSnssai) {
		return true
	}

	return false
}

// SetHomePlmnSnssai gets a reference to the given Snssai and assigns it to the HomePlmnSnssai field.
func (o *N2SmInformation) SetHomePlmnSnssai(v Snssai) {
	o.HomePlmnSnssai = &v
}

// GetIwkSnssai returns the IwkSnssai field value if set, zero value otherwise.
func (o *N2SmInformation) GetIwkSnssai() Snssai {
	if o == nil || IsNil(o.IwkSnssai) {
		var ret Snssai
		return ret
	}
	return *o.IwkSnssai
}

// GetIwkSnssaiOk returns a tuple with the IwkSnssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2SmInformation) GetIwkSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.IwkSnssai) {
		return nil, false
	}
	return o.IwkSnssai, true
}

// HasIwkSnssai returns a boolean if a field has been set.
func (o *N2SmInformation) HasIwkSnssai() bool {
	if o != nil && !IsNil(o.IwkSnssai) {
		return true
	}

	return false
}

// SetIwkSnssai gets a reference to the given Snssai and assigns it to the IwkSnssai field.
func (o *N2SmInformation) SetIwkSnssai(v Snssai) {
	o.IwkSnssai = &v
}

// GetSubjectToHo returns the SubjectToHo field value if set, zero value otherwise.
func (o *N2SmInformation) GetSubjectToHo() bool {
	if o == nil || IsNil(o.SubjectToHo) {
		var ret bool
		return ret
	}
	return *o.SubjectToHo
}

// GetSubjectToHoOk returns a tuple with the SubjectToHo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2SmInformation) GetSubjectToHoOk() (*bool, bool) {
	if o == nil || IsNil(o.SubjectToHo) {
		return nil, false
	}
	return o.SubjectToHo, true
}

// HasSubjectToHo returns a boolean if a field has been set.
func (o *N2SmInformation) HasSubjectToHo() bool {
	if o != nil && !IsNil(o.SubjectToHo) {
		return true
	}

	return false
}

// SetSubjectToHo gets a reference to the given bool and assigns it to the SubjectToHo field.
func (o *N2SmInformation) SetSubjectToHo(v bool) {
	o.SubjectToHo = &v
}

func (o N2SmInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o N2SmInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pduSessionId"] = o.PduSessionId
	if !IsNil(o.N2InfoContent) {
		toSerialize["n2InfoContent"] = o.N2InfoContent
	}
	if !IsNil(o.SNssai) {
		toSerialize["sNssai"] = o.SNssai
	}
	if !IsNil(o.HomePlmnSnssai) {
		toSerialize["homePlmnSnssai"] = o.HomePlmnSnssai
	}
	if !IsNil(o.IwkSnssai) {
		toSerialize["iwkSnssai"] = o.IwkSnssai
	}
	if !IsNil(o.SubjectToHo) {
		toSerialize["subjectToHo"] = o.SubjectToHo
	}
	return toSerialize, nil
}

type NullableN2SmInformation struct {
	value *N2SmInformation
	isSet bool
}

func (v NullableN2SmInformation) Get() *N2SmInformation {
	return v.value
}

func (v *NullableN2SmInformation) Set(val *N2SmInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableN2SmInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableN2SmInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2SmInformation(val *N2SmInformation) *NullableN2SmInformation {
	return &NullableN2SmInformation{value: val, isSet: true}
}

func (v NullableN2SmInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2SmInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



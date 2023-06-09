/*
Nnsacf_NSAC

Nnsacf_NSAC Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnsacf_NSAC

import (
	"encoding/json"
)

// checks if the PduACRequestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PduACRequestInfo{}

// PduACRequestInfo struct for PduACRequestInfo
type PduACRequestInfo struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi   string     `json:"supi"`
	AnType AccessType `json:"anType"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.
	PduSessionId     int32              `json:"pduSessionId"`
	AcuOperationList []AcuOperationItem `json:"acuOperationList"`
	AdditionalAnType *AccessType        `json:"additionalAnType,omitempty"`
}

// NewPduACRequestInfo instantiates a new PduACRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduACRequestInfo(supi string, anType AccessType, pduSessionId int32, acuOperationList []AcuOperationItem) *PduACRequestInfo {
	this := PduACRequestInfo{}
	this.Supi = supi
	this.AnType = anType
	this.PduSessionId = pduSessionId
	this.AcuOperationList = acuOperationList
	return &this
}

// NewPduACRequestInfoWithDefaults instantiates a new PduACRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduACRequestInfoWithDefaults() *PduACRequestInfo {
	this := PduACRequestInfo{}
	return &this
}

// GetSupi returns the Supi field value
func (o *PduACRequestInfo) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *PduACRequestInfo) GetSupiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *PduACRequestInfo) SetSupi(v string) {
	o.Supi = v
}

// GetAnType returns the AnType field value
func (o *PduACRequestInfo) GetAnType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AnType
}

// GetAnTypeOk returns a tuple with the AnType field value
// and a boolean to check if the value has been set.
func (o *PduACRequestInfo) GetAnTypeOk() (*AccessType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnType, true
}

// SetAnType sets field value
func (o *PduACRequestInfo) SetAnType(v AccessType) {
	o.AnType = v
}

// GetPduSessionId returns the PduSessionId field value
func (o *PduACRequestInfo) GetPduSessionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value
// and a boolean to check if the value has been set.
func (o *PduACRequestInfo) GetPduSessionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PduSessionId, true
}

// SetPduSessionId sets field value
func (o *PduACRequestInfo) SetPduSessionId(v int32) {
	o.PduSessionId = v
}

// GetAcuOperationList returns the AcuOperationList field value
func (o *PduACRequestInfo) GetAcuOperationList() []AcuOperationItem {
	if o == nil {
		var ret []AcuOperationItem
		return ret
	}

	return o.AcuOperationList
}

// GetAcuOperationListOk returns a tuple with the AcuOperationList field value
// and a boolean to check if the value has been set.
func (o *PduACRequestInfo) GetAcuOperationListOk() ([]AcuOperationItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcuOperationList, true
}

// SetAcuOperationList sets field value
func (o *PduACRequestInfo) SetAcuOperationList(v []AcuOperationItem) {
	o.AcuOperationList = v
}

// GetAdditionalAnType returns the AdditionalAnType field value if set, zero value otherwise.
func (o *PduACRequestInfo) GetAdditionalAnType() AccessType {
	if o == nil || IsNil(o.AdditionalAnType) {
		var ret AccessType
		return ret
	}
	return *o.AdditionalAnType
}

// GetAdditionalAnTypeOk returns a tuple with the AdditionalAnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduACRequestInfo) GetAdditionalAnTypeOk() (*AccessType, bool) {
	if o == nil || IsNil(o.AdditionalAnType) {
		return nil, false
	}
	return o.AdditionalAnType, true
}

// HasAdditionalAnType returns a boolean if a field has been set.
func (o *PduACRequestInfo) HasAdditionalAnType() bool {
	if o != nil && !IsNil(o.AdditionalAnType) {
		return true
	}

	return false
}

// SetAdditionalAnType gets a reference to the given AccessType and assigns it to the AdditionalAnType field.
func (o *PduACRequestInfo) SetAdditionalAnType(v AccessType) {
	o.AdditionalAnType = &v
}

func (o PduACRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PduACRequestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supi"] = o.Supi
	toSerialize["anType"] = o.AnType
	toSerialize["pduSessionId"] = o.PduSessionId
	toSerialize["acuOperationList"] = o.AcuOperationList
	if !IsNil(o.AdditionalAnType) {
		toSerialize["additionalAnType"] = o.AdditionalAnType
	}
	return toSerialize, nil
}

type NullablePduACRequestInfo struct {
	value *PduACRequestInfo
	isSet bool
}

func (v NullablePduACRequestInfo) Get() *PduACRequestInfo {
	return v.value
}

func (v *NullablePduACRequestInfo) Set(val *PduACRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePduACRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePduACRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduACRequestInfo(val *PduACRequestInfo) *NullablePduACRequestInfo {
	return &NullablePduACRequestInfo{value: val, isSet: true}
}

func (v NullablePduACRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduACRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

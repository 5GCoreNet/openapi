/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
)

// checks if the TriggeringEventsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggeringEventsType{}

// TriggeringEventsType Specifies when to start a Trace Recording Session and which message shall be recorded first, when to stop a Trace Recording Session and which message shall be recorded last respectively. See 3GPP TS 32.422 clause 5.1 for additional detials.
type TriggeringEventsType struct {
	MSC_SERVER []string `json:"MSC_SERVER,omitempty"`
	SGSN       []string `json:"SGSN,omitempty"`
	MGW        []string `json:"MGW,omitempty"`
	GGSN       []string `json:"GGSN,omitempty"`
	IMS        []string `json:"IMS,omitempty"`
	BM_SC      []string `json:"BM_SC,omitempty"`
	MME        []string `json:"MME,omitempty"`
	SGW        []string `json:"SGW,omitempty"`
	PGW        []string `json:"PGW,omitempty"`
	AMF        []string `json:"AMF,omitempty"`
	SMF        []string `json:"SMF,omitempty"`
	PCF        []string `json:"PCF,omitempty"`
	UPF        []string `json:"UPF,omitempty"`
	AUSF       []string `json:"AUSF,omitempty"`
	NEF        []string `json:"NEF,omitempty"`
	NRF        []string `json:"NRF,omitempty"`
	NSSF       []string `json:"NSSF,omitempty"`
	SMSF       []string `json:"SMSF,omitempty"`
	UDM        []string `json:"UDM,omitempty"`
}

// NewTriggeringEventsType instantiates a new TriggeringEventsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggeringEventsType() *TriggeringEventsType {
	this := TriggeringEventsType{}
	return &this
}

// NewTriggeringEventsTypeWithDefaults instantiates a new TriggeringEventsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggeringEventsTypeWithDefaults() *TriggeringEventsType {
	this := TriggeringEventsType{}
	return &this
}

// GetMSC_SERVER returns the MSC_SERVER field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetMSC_SERVER() []string {
	if o == nil || IsNil(o.MSC_SERVER) {
		var ret []string
		return ret
	}
	return o.MSC_SERVER
}

// GetMSC_SERVEROk returns a tuple with the MSC_SERVER field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetMSC_SERVEROk() ([]string, bool) {
	if o == nil || IsNil(o.MSC_SERVER) {
		return nil, false
	}
	return o.MSC_SERVER, true
}

// HasMSC_SERVER returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasMSC_SERVER() bool {
	if o != nil && !IsNil(o.MSC_SERVER) {
		return true
	}

	return false
}

// SetMSC_SERVER gets a reference to the given []string and assigns it to the MSC_SERVER field.
func (o *TriggeringEventsType) SetMSC_SERVER(v []string) {
	o.MSC_SERVER = v
}

// GetSGSN returns the SGSN field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetSGSN() []string {
	if o == nil || IsNil(o.SGSN) {
		var ret []string
		return ret
	}
	return o.SGSN
}

// GetSGSNOk returns a tuple with the SGSN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetSGSNOk() ([]string, bool) {
	if o == nil || IsNil(o.SGSN) {
		return nil, false
	}
	return o.SGSN, true
}

// HasSGSN returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasSGSN() bool {
	if o != nil && !IsNil(o.SGSN) {
		return true
	}

	return false
}

// SetSGSN gets a reference to the given []string and assigns it to the SGSN field.
func (o *TriggeringEventsType) SetSGSN(v []string) {
	o.SGSN = v
}

// GetMGW returns the MGW field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetMGW() []string {
	if o == nil || IsNil(o.MGW) {
		var ret []string
		return ret
	}
	return o.MGW
}

// GetMGWOk returns a tuple with the MGW field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetMGWOk() ([]string, bool) {
	if o == nil || IsNil(o.MGW) {
		return nil, false
	}
	return o.MGW, true
}

// HasMGW returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasMGW() bool {
	if o != nil && !IsNil(o.MGW) {
		return true
	}

	return false
}

// SetMGW gets a reference to the given []string and assigns it to the MGW field.
func (o *TriggeringEventsType) SetMGW(v []string) {
	o.MGW = v
}

// GetGGSN returns the GGSN field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetGGSN() []string {
	if o == nil || IsNil(o.GGSN) {
		var ret []string
		return ret
	}
	return o.GGSN
}

// GetGGSNOk returns a tuple with the GGSN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetGGSNOk() ([]string, bool) {
	if o == nil || IsNil(o.GGSN) {
		return nil, false
	}
	return o.GGSN, true
}

// HasGGSN returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasGGSN() bool {
	if o != nil && !IsNil(o.GGSN) {
		return true
	}

	return false
}

// SetGGSN gets a reference to the given []string and assigns it to the GGSN field.
func (o *TriggeringEventsType) SetGGSN(v []string) {
	o.GGSN = v
}

// GetIMS returns the IMS field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetIMS() []string {
	if o == nil || IsNil(o.IMS) {
		var ret []string
		return ret
	}
	return o.IMS
}

// GetIMSOk returns a tuple with the IMS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetIMSOk() ([]string, bool) {
	if o == nil || IsNil(o.IMS) {
		return nil, false
	}
	return o.IMS, true
}

// HasIMS returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasIMS() bool {
	if o != nil && !IsNil(o.IMS) {
		return true
	}

	return false
}

// SetIMS gets a reference to the given []string and assigns it to the IMS field.
func (o *TriggeringEventsType) SetIMS(v []string) {
	o.IMS = v
}

// GetBM_SC returns the BM_SC field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetBM_SC() []string {
	if o == nil || IsNil(o.BM_SC) {
		var ret []string
		return ret
	}
	return o.BM_SC
}

// GetBM_SCOk returns a tuple with the BM_SC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetBM_SCOk() ([]string, bool) {
	if o == nil || IsNil(o.BM_SC) {
		return nil, false
	}
	return o.BM_SC, true
}

// HasBM_SC returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasBM_SC() bool {
	if o != nil && !IsNil(o.BM_SC) {
		return true
	}

	return false
}

// SetBM_SC gets a reference to the given []string and assigns it to the BM_SC field.
func (o *TriggeringEventsType) SetBM_SC(v []string) {
	o.BM_SC = v
}

// GetMME returns the MME field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetMME() []string {
	if o == nil || IsNil(o.MME) {
		var ret []string
		return ret
	}
	return o.MME
}

// GetMMEOk returns a tuple with the MME field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetMMEOk() ([]string, bool) {
	if o == nil || IsNil(o.MME) {
		return nil, false
	}
	return o.MME, true
}

// HasMME returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasMME() bool {
	if o != nil && !IsNil(o.MME) {
		return true
	}

	return false
}

// SetMME gets a reference to the given []string and assigns it to the MME field.
func (o *TriggeringEventsType) SetMME(v []string) {
	o.MME = v
}

// GetSGW returns the SGW field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetSGW() []string {
	if o == nil || IsNil(o.SGW) {
		var ret []string
		return ret
	}
	return o.SGW
}

// GetSGWOk returns a tuple with the SGW field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetSGWOk() ([]string, bool) {
	if o == nil || IsNil(o.SGW) {
		return nil, false
	}
	return o.SGW, true
}

// HasSGW returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasSGW() bool {
	if o != nil && !IsNil(o.SGW) {
		return true
	}

	return false
}

// SetSGW gets a reference to the given []string and assigns it to the SGW field.
func (o *TriggeringEventsType) SetSGW(v []string) {
	o.SGW = v
}

// GetPGW returns the PGW field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetPGW() []string {
	if o == nil || IsNil(o.PGW) {
		var ret []string
		return ret
	}
	return o.PGW
}

// GetPGWOk returns a tuple with the PGW field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetPGWOk() ([]string, bool) {
	if o == nil || IsNil(o.PGW) {
		return nil, false
	}
	return o.PGW, true
}

// HasPGW returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasPGW() bool {
	if o != nil && !IsNil(o.PGW) {
		return true
	}

	return false
}

// SetPGW gets a reference to the given []string and assigns it to the PGW field.
func (o *TriggeringEventsType) SetPGW(v []string) {
	o.PGW = v
}

// GetAMF returns the AMF field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetAMF() []string {
	if o == nil || IsNil(o.AMF) {
		var ret []string
		return ret
	}
	return o.AMF
}

// GetAMFOk returns a tuple with the AMF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetAMFOk() ([]string, bool) {
	if o == nil || IsNil(o.AMF) {
		return nil, false
	}
	return o.AMF, true
}

// HasAMF returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasAMF() bool {
	if o != nil && !IsNil(o.AMF) {
		return true
	}

	return false
}

// SetAMF gets a reference to the given []string and assigns it to the AMF field.
func (o *TriggeringEventsType) SetAMF(v []string) {
	o.AMF = v
}

// GetSMF returns the SMF field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetSMF() []string {
	if o == nil || IsNil(o.SMF) {
		var ret []string
		return ret
	}
	return o.SMF
}

// GetSMFOk returns a tuple with the SMF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetSMFOk() ([]string, bool) {
	if o == nil || IsNil(o.SMF) {
		return nil, false
	}
	return o.SMF, true
}

// HasSMF returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasSMF() bool {
	if o != nil && !IsNil(o.SMF) {
		return true
	}

	return false
}

// SetSMF gets a reference to the given []string and assigns it to the SMF field.
func (o *TriggeringEventsType) SetSMF(v []string) {
	o.SMF = v
}

// GetPCF returns the PCF field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetPCF() []string {
	if o == nil || IsNil(o.PCF) {
		var ret []string
		return ret
	}
	return o.PCF
}

// GetPCFOk returns a tuple with the PCF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetPCFOk() ([]string, bool) {
	if o == nil || IsNil(o.PCF) {
		return nil, false
	}
	return o.PCF, true
}

// HasPCF returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasPCF() bool {
	if o != nil && !IsNil(o.PCF) {
		return true
	}

	return false
}

// SetPCF gets a reference to the given []string and assigns it to the PCF field.
func (o *TriggeringEventsType) SetPCF(v []string) {
	o.PCF = v
}

// GetUPF returns the UPF field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetUPF() []string {
	if o == nil || IsNil(o.UPF) {
		var ret []string
		return ret
	}
	return o.UPF
}

// GetUPFOk returns a tuple with the UPF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetUPFOk() ([]string, bool) {
	if o == nil || IsNil(o.UPF) {
		return nil, false
	}
	return o.UPF, true
}

// HasUPF returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasUPF() bool {
	if o != nil && !IsNil(o.UPF) {
		return true
	}

	return false
}

// SetUPF gets a reference to the given []string and assigns it to the UPF field.
func (o *TriggeringEventsType) SetUPF(v []string) {
	o.UPF = v
}

// GetAUSF returns the AUSF field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetAUSF() []string {
	if o == nil || IsNil(o.AUSF) {
		var ret []string
		return ret
	}
	return o.AUSF
}

// GetAUSFOk returns a tuple with the AUSF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetAUSFOk() ([]string, bool) {
	if o == nil || IsNil(o.AUSF) {
		return nil, false
	}
	return o.AUSF, true
}

// HasAUSF returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasAUSF() bool {
	if o != nil && !IsNil(o.AUSF) {
		return true
	}

	return false
}

// SetAUSF gets a reference to the given []string and assigns it to the AUSF field.
func (o *TriggeringEventsType) SetAUSF(v []string) {
	o.AUSF = v
}

// GetNEF returns the NEF field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetNEF() []string {
	if o == nil || IsNil(o.NEF) {
		var ret []string
		return ret
	}
	return o.NEF
}

// GetNEFOk returns a tuple with the NEF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetNEFOk() ([]string, bool) {
	if o == nil || IsNil(o.NEF) {
		return nil, false
	}
	return o.NEF, true
}

// HasNEF returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasNEF() bool {
	if o != nil && !IsNil(o.NEF) {
		return true
	}

	return false
}

// SetNEF gets a reference to the given []string and assigns it to the NEF field.
func (o *TriggeringEventsType) SetNEF(v []string) {
	o.NEF = v
}

// GetNRF returns the NRF field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetNRF() []string {
	if o == nil || IsNil(o.NRF) {
		var ret []string
		return ret
	}
	return o.NRF
}

// GetNRFOk returns a tuple with the NRF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetNRFOk() ([]string, bool) {
	if o == nil || IsNil(o.NRF) {
		return nil, false
	}
	return o.NRF, true
}

// HasNRF returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasNRF() bool {
	if o != nil && !IsNil(o.NRF) {
		return true
	}

	return false
}

// SetNRF gets a reference to the given []string and assigns it to the NRF field.
func (o *TriggeringEventsType) SetNRF(v []string) {
	o.NRF = v
}

// GetNSSF returns the NSSF field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetNSSF() []string {
	if o == nil || IsNil(o.NSSF) {
		var ret []string
		return ret
	}
	return o.NSSF
}

// GetNSSFOk returns a tuple with the NSSF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetNSSFOk() ([]string, bool) {
	if o == nil || IsNil(o.NSSF) {
		return nil, false
	}
	return o.NSSF, true
}

// HasNSSF returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasNSSF() bool {
	if o != nil && !IsNil(o.NSSF) {
		return true
	}

	return false
}

// SetNSSF gets a reference to the given []string and assigns it to the NSSF field.
func (o *TriggeringEventsType) SetNSSF(v []string) {
	o.NSSF = v
}

// GetSMSF returns the SMSF field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetSMSF() []string {
	if o == nil || IsNil(o.SMSF) {
		var ret []string
		return ret
	}
	return o.SMSF
}

// GetSMSFOk returns a tuple with the SMSF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetSMSFOk() ([]string, bool) {
	if o == nil || IsNil(o.SMSF) {
		return nil, false
	}
	return o.SMSF, true
}

// HasSMSF returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasSMSF() bool {
	if o != nil && !IsNil(o.SMSF) {
		return true
	}

	return false
}

// SetSMSF gets a reference to the given []string and assigns it to the SMSF field.
func (o *TriggeringEventsType) SetSMSF(v []string) {
	o.SMSF = v
}

// GetUDM returns the UDM field value if set, zero value otherwise.
func (o *TriggeringEventsType) GetUDM() []string {
	if o == nil || IsNil(o.UDM) {
		var ret []string
		return ret
	}
	return o.UDM
}

// GetUDMOk returns a tuple with the UDM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggeringEventsType) GetUDMOk() ([]string, bool) {
	if o == nil || IsNil(o.UDM) {
		return nil, false
	}
	return o.UDM, true
}

// HasUDM returns a boolean if a field has been set.
func (o *TriggeringEventsType) HasUDM() bool {
	if o != nil && !IsNil(o.UDM) {
		return true
	}

	return false
}

// SetUDM gets a reference to the given []string and assigns it to the UDM field.
func (o *TriggeringEventsType) SetUDM(v []string) {
	o.UDM = v
}

func (o TriggeringEventsType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggeringEventsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MSC_SERVER) {
		toSerialize["MSC_SERVER"] = o.MSC_SERVER
	}
	if !IsNil(o.SGSN) {
		toSerialize["SGSN"] = o.SGSN
	}
	if !IsNil(o.MGW) {
		toSerialize["MGW"] = o.MGW
	}
	if !IsNil(o.GGSN) {
		toSerialize["GGSN"] = o.GGSN
	}
	if !IsNil(o.IMS) {
		toSerialize["IMS"] = o.IMS
	}
	if !IsNil(o.BM_SC) {
		toSerialize["BM_SC"] = o.BM_SC
	}
	if !IsNil(o.MME) {
		toSerialize["MME"] = o.MME
	}
	if !IsNil(o.SGW) {
		toSerialize["SGW"] = o.SGW
	}
	if !IsNil(o.PGW) {
		toSerialize["PGW"] = o.PGW
	}
	if !IsNil(o.AMF) {
		toSerialize["AMF"] = o.AMF
	}
	if !IsNil(o.SMF) {
		toSerialize["SMF"] = o.SMF
	}
	if !IsNil(o.PCF) {
		toSerialize["PCF"] = o.PCF
	}
	if !IsNil(o.UPF) {
		toSerialize["UPF"] = o.UPF
	}
	if !IsNil(o.AUSF) {
		toSerialize["AUSF"] = o.AUSF
	}
	if !IsNil(o.NEF) {
		toSerialize["NEF"] = o.NEF
	}
	if !IsNil(o.NRF) {
		toSerialize["NRF"] = o.NRF
	}
	if !IsNil(o.NSSF) {
		toSerialize["NSSF"] = o.NSSF
	}
	if !IsNil(o.SMSF) {
		toSerialize["SMSF"] = o.SMSF
	}
	if !IsNil(o.UDM) {
		toSerialize["UDM"] = o.UDM
	}
	return toSerialize, nil
}

type NullableTriggeringEventsType struct {
	value *TriggeringEventsType
	isSet bool
}

func (v NullableTriggeringEventsType) Get() *TriggeringEventsType {
	return v.value
}

func (v *NullableTriggeringEventsType) Set(val *TriggeringEventsType) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggeringEventsType) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggeringEventsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggeringEventsType(val *TriggeringEventsType) *NullableTriggeringEventsType {
	return &NullableTriggeringEventsType{value: val, isSet: true}
}

func (v NullableTriggeringEventsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggeringEventsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

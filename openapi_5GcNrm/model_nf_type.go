/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
	"fmt"
)

// NFType  NF name defined in TS 23.501 or TS 29.510
type NFType string

// List of NFType
const (
	NRF NFType = "NRF"
	UDM NFType = "UDM"
	AMF NFType = "AMF"
	SMF NFType = "SMF"
	AUSF NFType = "AUSF"
	NEF NFType = "NEF"
	PCF NFType = "PCF"
	SMSF NFType = "SMSF"
	NSSF NFType = "NSSF"
	UDR NFType = "UDR"
	LMF NFType = "LMF"
	GMLC NFType = "GMLC"
	_5_G_EIR NFType = "5G_EIR"
	SEPP NFType = "SEPP"
	UPF NFType = "UPF"
	N3_IWF NFType = "N3IWF"
	AF NFType = "AF"
	UDSF NFType = "UDSF"
	DN NFType = "DN"
	BSF NFType = "BSF"
	CHF NFType = "CHF"
	NWDAF NFType = "NWDAF"
	PCSCF NFType = "PCSCF"
	CBCF NFType = "CBCF"
	HSS NFType = "HSS"
	UCMF NFType = "UCMF"
	SOR_AF NFType = "SOR_AF"
	SPAF NFType = "SPAF"
	MME NFType = "MME"
	SCSAS NFType = "SCSAS"
	SCEF NFType = "SCEF"
	SCP NFType = "SCP"
	NSSAAF NFType = "NSSAAF"
	ICSCF NFType = "ICSCF"
	SCSCF NFType = "SCSCF"
	DRA NFType = "DRA"
	IMS_AS NFType = "IMS_AS"
	AANF NFType = "AANF"
	_5_G_DDNMF NFType = "5G_DDNMF"
	NSACF NFType = "NSACF"
	MFAF NFType = "MFAF"
	EASDF NFType = "EASDF"
	DCCF NFType = "DCCF"
	MB_SMF NFType = "MB_SMF"
	TSCTSF NFType = "TSCTSF"
	ADRF NFType = "ADRF"
	GBA_BSF NFType = "GBA_BSF"
	CEF NFType = "CEF"
	MB_UPF NFType = "MB_UPF"
	NSWOF NFType = "NSWOF"
	PKMF NFType = "PKMF"
	MNPF NFType = "MNPF"
	SMS_GMSC NFType = "SMS_GMSC"
	SMS_IWMSC NFType = "SMS_IWMSC"
	MBSF NFType = "MBSF"
	MBSTF NFType = "MBSTF"
	PANF NFType = "PANF"
	TNGF NFType = "TNGF"
	W_AGF NFType = "W_AGF"
	TWIF NFType = "TWIF"
	TSN_AF NFType = "TSN_AF"
)

// All allowed values of NFType enum
var AllowedNFTypeEnumValues = []NFType{
	"NRF",
	"UDM",
	"AMF",
	"SMF",
	"AUSF",
	"NEF",
	"PCF",
	"SMSF",
	"NSSF",
	"UDR",
	"LMF",
	"GMLC",
	"5G_EIR",
	"SEPP",
	"UPF",
	"N3IWF",
	"AF",
	"UDSF",
	"DN",
	"BSF",
	"CHF",
	"NWDAF",
	"PCSCF",
	"CBCF",
	"HSS",
	"UCMF",
	"SOR_AF",
	"SPAF",
	"MME",
	"SCSAS",
	"SCEF",
	"SCP",
	"NSSAAF",
	"ICSCF",
	"SCSCF",
	"DRA",
	"IMS_AS",
	"AANF",
	"5G_DDNMF",
	"NSACF",
	"MFAF",
	"EASDF",
	"DCCF",
	"MB_SMF",
	"TSCTSF",
	"ADRF",
	"GBA_BSF",
	"CEF",
	"MB_UPF",
	"NSWOF",
	"PKMF",
	"MNPF",
	"SMS_GMSC",
	"SMS_IWMSC",
	"MBSF",
	"MBSTF",
	"PANF",
	"TNGF",
	"W_AGF",
	"TWIF",
	"TSN_AF",
}

func (v *NFType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NFType(value)
	for _, existing := range AllowedNFTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NFType", value)
}

// NewNFTypeFromValue returns a pointer to a valid NFType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNFTypeFromValue(v string) (*NFType, error) {
	ev := NFType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NFType: valid values are %v", v, AllowedNFTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NFType) IsValid() bool {
	for _, existing := range AllowedNFTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NFType value
func (v NFType) Ptr() *NFType {
	return &v
}

type NullableNFType struct {
	value *NFType
	isSet bool
}

func (v NullableNFType) Get() *NFType {
	return v.value
}

func (v *NullableNFType) Set(val *NFType) {
	v.value = val
	v.isSet = true
}

func (v NullableNFType) IsSet() bool {
	return v.isSet
}

func (v *NullableNFType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFType(val *NFType) *NullableNFType {
	return &NullableNFType{value: val, isSet: true}
}

func (v NullableNFType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


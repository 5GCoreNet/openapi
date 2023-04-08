/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// UETransferType struct for UETransferType
type UETransferType struct {
	UETransferTypeAnyOf *UETransferTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UETransferType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UETransferTypeAnyOf
	err = json.Unmarshal(data, &dst.UETransferTypeAnyOf);
	if err == nil {
		jsonUETransferTypeAnyOf, _ := json.Marshal(dst.UETransferTypeAnyOf)
		if string(jsonUETransferTypeAnyOf) == "{}" { // empty struct
			dst.UETransferTypeAnyOf = nil
		} else {
			return nil // data stored in dst.UETransferTypeAnyOf, return on the first match
		}
	} else {
		dst.UETransferTypeAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(UETransferType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UETransferType) MarshalJSON() ([]byte, error) {
	if src.UETransferTypeAnyOf != nil {
		return json.Marshal(&src.UETransferTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUETransferType struct {
	value *UETransferType
	isSet bool
}

func (v NullableUETransferType) Get() *UETransferType {
	return v.value
}

func (v *NullableUETransferType) Set(val *UETransferType) {
	v.value = val
	v.isSet = true
}

func (v NullableUETransferType) IsSet() bool {
	return v.isSet
}

func (v *NullableUETransferType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUETransferType(val *UETransferType) *NullableUETransferType {
	return &NullableUETransferType{value: val, isSet: true}
}

func (v NullableUETransferType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUETransferType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



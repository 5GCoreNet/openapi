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

// RedundantTransmissionType struct for RedundantTransmissionType
type RedundantTransmissionType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RedundantTransmissionType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RedundantTransmissionType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RedundantTransmissionType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRedundantTransmissionType struct {
	value *RedundantTransmissionType
	isSet bool
}

func (v NullableRedundantTransmissionType) Get() *RedundantTransmissionType {
	return v.value
}

func (v *NullableRedundantTransmissionType) Set(val *RedundantTransmissionType) {
	v.value = val
	v.isSet = true
}

func (v NullableRedundantTransmissionType) IsSet() bool {
	return v.isSet
}

func (v *NullableRedundantTransmissionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedundantTransmissionType(val *RedundantTransmissionType) *NullableRedundantTransmissionType {
	return &NullableRedundantTransmissionType{value: val, isSet: true}
}

func (v NullableRedundantTransmissionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedundantTransmissionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



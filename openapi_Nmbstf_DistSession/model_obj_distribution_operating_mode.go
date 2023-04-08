/*
Nmbstf-distsession

MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbstf_DistSession

import (
	"encoding/json"
	"fmt"
)

// ObjDistributionOperatingMode Mode of data ingestion for Object distribution method
type ObjDistributionOperatingMode struct {
	ObjDistributionOperatingModeAnyOf *ObjDistributionOperatingModeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ObjDistributionOperatingMode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ObjDistributionOperatingModeAnyOf
	err = json.Unmarshal(data, &dst.ObjDistributionOperatingModeAnyOf);
	if err == nil {
		jsonObjDistributionOperatingModeAnyOf, _ := json.Marshal(dst.ObjDistributionOperatingModeAnyOf)
		if string(jsonObjDistributionOperatingModeAnyOf) == "{}" { // empty struct
			dst.ObjDistributionOperatingModeAnyOf = nil
		} else {
			return nil // data stored in dst.ObjDistributionOperatingModeAnyOf, return on the first match
		}
	} else {
		dst.ObjDistributionOperatingModeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ObjDistributionOperatingMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ObjDistributionOperatingMode) MarshalJSON() ([]byte, error) {
	if src.ObjDistributionOperatingModeAnyOf != nil {
		return json.Marshal(&src.ObjDistributionOperatingModeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableObjDistributionOperatingMode struct {
	value *ObjDistributionOperatingMode
	isSet bool
}

func (v NullableObjDistributionOperatingMode) Get() *ObjDistributionOperatingMode {
	return v.value
}

func (v *NullableObjDistributionOperatingMode) Set(val *ObjDistributionOperatingMode) {
	v.value = val
	v.isSet = true
}

func (v NullableObjDistributionOperatingMode) IsSet() bool {
	return v.isSet
}

func (v *NullableObjDistributionOperatingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjDistributionOperatingMode(val *ObjDistributionOperatingMode) *NullableObjDistributionOperatingMode {
	return &NullableObjDistributionOperatingMode{value: val, isSet: true}
}

func (v NullableObjDistributionOperatingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjDistributionOperatingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



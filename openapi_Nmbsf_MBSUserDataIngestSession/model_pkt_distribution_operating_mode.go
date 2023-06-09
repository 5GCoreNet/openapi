/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// PktDistributionOperatingMode Mode of data ingestion for Packet distribution method
type PktDistributionOperatingMode struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PktDistributionOperatingMode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
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

	return fmt.Errorf("data failed to match schemas in anyOf(PktDistributionOperatingMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PktDistributionOperatingMode) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePktDistributionOperatingMode struct {
	value *PktDistributionOperatingMode
	isSet bool
}

func (v NullablePktDistributionOperatingMode) Get() *PktDistributionOperatingMode {
	return v.value
}

func (v *NullablePktDistributionOperatingMode) Set(val *PktDistributionOperatingMode) {
	v.value = val
	v.isSet = true
}

func (v NullablePktDistributionOperatingMode) IsSet() bool {
	return v.isSet
}

func (v *NullablePktDistributionOperatingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePktDistributionOperatingMode(val *PktDistributionOperatingMode) *NullablePktDistributionOperatingMode {
	return &NullablePktDistributionOperatingMode{value: val, isSet: true}
}

func (v NullablePktDistributionOperatingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePktDistributionOperatingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

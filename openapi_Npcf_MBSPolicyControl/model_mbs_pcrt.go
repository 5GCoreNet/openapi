/*
Npcf_MBSPolicyControl API

MBS Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_MBSPolicyControl

import (
	"encoding/json"
	"fmt"
)

// MbsPcrt Possible values are: - MBS_SESSION_UPDATE: Indicates the MBS Session Update policy control request trigger. 
type MbsPcrt struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MbsPcrt) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(MbsPcrt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MbsPcrt) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMbsPcrt struct {
	value *MbsPcrt
	isSet bool
}

func (v NullableMbsPcrt) Get() *MbsPcrt {
	return v.value
}

func (v *NullableMbsPcrt) Set(val *MbsPcrt) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsPcrt) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsPcrt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsPcrt(val *MbsPcrt) *NullableMbsPcrt {
	return &NullableMbsPcrt{value: val, isSet: true}
}

func (v NullableMbsPcrt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsPcrt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



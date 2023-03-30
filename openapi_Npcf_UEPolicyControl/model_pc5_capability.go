/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_UEPolicyControl

import (
	"encoding/json"
	"fmt"
)

// Pc5Capability Possible values are: - LTE_PC5: This value is used to indicate that UE supports PC5 LTE RAT for V2X communications   over the PC5 reference point. - NR_PC5: This value is used to indicate that UE supports PC5 NR RAT for V2X communications   over the PC5 reference point. - LTE_NR_PC5: This value is used to indicate that UE supports both PC5 LTE and NR RAT for   V2X communications over the PC5 reference point.  
type Pc5Capability struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Pc5Capability) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(Pc5Capability)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Pc5Capability) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePc5Capability struct {
	value *Pc5Capability
	isSet bool
}

func (v NullablePc5Capability) Get() *Pc5Capability {
	return v.value
}

func (v *NullablePc5Capability) Set(val *Pc5Capability) {
	v.value = val
	v.isSet = true
}

func (v NullablePc5Capability) IsSet() bool {
	return v.isSet
}

func (v *NullablePc5Capability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePc5Capability(val *Pc5Capability) *NullablePc5Capability {
	return &NullablePc5Capability{value: val, isSet: true}
}

func (v NullablePc5Capability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePc5Capability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



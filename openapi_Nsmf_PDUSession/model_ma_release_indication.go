/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// MaReleaseIndication Multi-Access PDU session release Indication. Possible values are   - REL_MAPDU_OVER_3GPP   - REL_MAPDU_OVER_N3GPP
type MaReleaseIndication struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MaReleaseIndication) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(MaReleaseIndication)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MaReleaseIndication) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMaReleaseIndication struct {
	value *MaReleaseIndication
	isSet bool
}

func (v NullableMaReleaseIndication) Get() *MaReleaseIndication {
	return v.value
}

func (v *NullableMaReleaseIndication) Set(val *MaReleaseIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableMaReleaseIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableMaReleaseIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaReleaseIndication(val *MaReleaseIndication) *NullableMaReleaseIndication {
	return &NullableMaReleaseIndication{value: val, isSet: true}
}

func (v NullableMaReleaseIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaReleaseIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

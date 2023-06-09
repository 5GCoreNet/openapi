/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReporting

import (
	"encoding/json"
	"fmt"
)

// PositioningMode Indicates supported modes used for positioning method.
type PositioningMode struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PositioningMode) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(PositioningMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PositioningMode) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePositioningMode struct {
	value *PositioningMode
	isSet bool
}

func (v NullablePositioningMode) Get() *PositioningMode {
	return v.value
}

func (v *NullablePositioningMode) Set(val *PositioningMode) {
	v.value = val
	v.isSet = true
}

func (v NullablePositioningMode) IsSet() bool {
	return v.isSet
}

func (v *NullablePositioningMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositioningMode(val *PositioningMode) *NullablePositioningMode {
	return &NullablePositioningMode{value: val, isSet: true}
}

func (v NullablePositioningMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositioningMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

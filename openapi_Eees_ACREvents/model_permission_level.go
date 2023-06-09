/*
Eees_ACREvents

API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACREvents

import (
	"encoding/json"
	"fmt"
)

// PermissionLevel Possible values are: - TRIAL: Level of service permission supported is TRIAL. - GOLD: Level of service permission supported is GOLD. - SILVER: Level of service permission supported is SILVER. - OTHER: Any other level of service permissions supported.
type PermissionLevel struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PermissionLevel) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(PermissionLevel)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PermissionLevel) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePermissionLevel struct {
	value *PermissionLevel
	isSet bool
}

func (v NullablePermissionLevel) Get() *PermissionLevel {
	return v.value
}

func (v *NullablePermissionLevel) Set(val *PermissionLevel) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionLevel) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionLevel(val *PermissionLevel) *NullablePermissionLevel {
	return &NullablePermissionLevel{value: val, isSet: true}
}

func (v NullablePermissionLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

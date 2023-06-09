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

// ResourceStatus Status of SM context or PDU session resource. Possible values are - RELEASED - UNCHANGED - TRANSFERRED - UPDATED - ALT_ANCHOR_SMF
type ResourceStatus struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResourceStatus) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ResourceStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ResourceStatus) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableResourceStatus struct {
	value *ResourceStatus
	isSet bool
}

func (v NullableResourceStatus) Get() *ResourceStatus {
	return v.value
}

func (v *NullableResourceStatus) Set(val *ResourceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceStatus(val *ResourceStatus) *NullableResourceStatus {
	return &NullableResourceStatus{value: val, isSet: true}
}

func (v NullableResourceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

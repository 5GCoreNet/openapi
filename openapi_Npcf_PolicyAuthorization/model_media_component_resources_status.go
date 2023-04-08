/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// MediaComponentResourcesStatus Indicates whether the media component is active or inactive.
type MediaComponentResourcesStatus struct {
	MediaComponentResourcesStatusAnyOf *MediaComponentResourcesStatusAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MediaComponentResourcesStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MediaComponentResourcesStatusAnyOf
	err = json.Unmarshal(data, &dst.MediaComponentResourcesStatusAnyOf);
	if err == nil {
		jsonMediaComponentResourcesStatusAnyOf, _ := json.Marshal(dst.MediaComponentResourcesStatusAnyOf)
		if string(jsonMediaComponentResourcesStatusAnyOf) == "{}" { // empty struct
			dst.MediaComponentResourcesStatusAnyOf = nil
		} else {
			return nil // data stored in dst.MediaComponentResourcesStatusAnyOf, return on the first match
		}
	} else {
		dst.MediaComponentResourcesStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MediaComponentResourcesStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MediaComponentResourcesStatus) MarshalJSON() ([]byte, error) {
	if src.MediaComponentResourcesStatusAnyOf != nil {
		return json.Marshal(&src.MediaComponentResourcesStatusAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMediaComponentResourcesStatus struct {
	value *MediaComponentResourcesStatus
	isSet bool
}

func (v NullableMediaComponentResourcesStatus) Get() *MediaComponentResourcesStatus {
	return v.value
}

func (v *NullableMediaComponentResourcesStatus) Set(val *MediaComponentResourcesStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaComponentResourcesStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaComponentResourcesStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaComponentResourcesStatus(val *MediaComponentResourcesStatus) *NullableMediaComponentResourcesStatus {
	return &NullableMediaComponentResourcesStatus{value: val, isSet: true}
}

func (v NullableMediaComponentResourcesStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaComponentResourcesStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



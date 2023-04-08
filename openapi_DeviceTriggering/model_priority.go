/*
3gpp-device-triggering

API for device trigger.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_DeviceTriggering

import (
	"encoding/json"
	"fmt"
)

// Priority Possible values are - NO_PRIORITY: This value indicates that the device trigger has no priority. - PRIORITY: This value indicates that the device trigger has priority. 
type Priority struct {
	PriorityAnyOf *PriorityAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Priority) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PriorityAnyOf
	err = json.Unmarshal(data, &dst.PriorityAnyOf);
	if err == nil {
		jsonPriorityAnyOf, _ := json.Marshal(dst.PriorityAnyOf)
		if string(jsonPriorityAnyOf) == "{}" { // empty struct
			dst.PriorityAnyOf = nil
		} else {
			return nil // data stored in dst.PriorityAnyOf, return on the first match
		}
	} else {
		dst.PriorityAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(Priority)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Priority) MarshalJSON() ([]byte, error) {
	if src.PriorityAnyOf != nil {
		return json.Marshal(&src.PriorityAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePriority struct {
	value *Priority
	isSet bool
}

func (v NullablePriority) Get() *Priority {
	return v.value
}

func (v *NullablePriority) Set(val *Priority) {
	v.value = val
	v.isSet = true
}

func (v NullablePriority) IsSet() bool {
	return v.isSet
}

func (v *NullablePriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriority(val *Priority) *NullablePriority {
	return &NullablePriority{value: val, isSet: true}
}

func (v NullablePriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



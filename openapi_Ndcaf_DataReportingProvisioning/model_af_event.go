/*
Ndcaf_DataReportingProvisioning

Data Collection AF: Provisioning Sessions API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReportingProvisioning

import (
	"encoding/json"
	"fmt"
)

// AfEvent Represents Application Events.
type AfEvent struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AfEvent) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AfEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AfEvent) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAfEvent struct {
	value *AfEvent
	isSet bool
}

func (v NullableAfEvent) Get() *AfEvent {
	return v.value
}

func (v *NullableAfEvent) Set(val *AfEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAfEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAfEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfEvent(val *AfEvent) *NullableAfEvent {
	return &NullableAfEvent{value: val, isSet: true}
}

func (v NullableAfEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

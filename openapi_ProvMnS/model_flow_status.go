/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// FlowStatus Describes whether the IP flow(s) are enabled or disabled.
type FlowStatus struct {
	FlowStatusAnyOf *FlowStatusAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FlowStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FlowStatusAnyOf
	err = json.Unmarshal(data, &dst.FlowStatusAnyOf);
	if err == nil {
		jsonFlowStatusAnyOf, _ := json.Marshal(dst.FlowStatusAnyOf)
		if string(jsonFlowStatusAnyOf) == "{}" { // empty struct
			dst.FlowStatusAnyOf = nil
		} else {
			return nil // data stored in dst.FlowStatusAnyOf, return on the first match
		}
	} else {
		dst.FlowStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(FlowStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FlowStatus) MarshalJSON() ([]byte, error) {
	if src.FlowStatusAnyOf != nil {
		return json.Marshal(&src.FlowStatusAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFlowStatus struct {
	value *FlowStatus
	isSet bool
}

func (v NullableFlowStatus) Get() *FlowStatus {
	return v.value
}

func (v *NullableFlowStatus) Set(val *FlowStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowStatus(val *FlowStatus) *NullableFlowStatus {
	return &NullableFlowStatus{value: val, isSet: true}
}

func (v NullableFlowStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



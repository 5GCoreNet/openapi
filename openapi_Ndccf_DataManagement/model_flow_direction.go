/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// FlowDirection Possible values are: - DOWNLINK: The corresponding filter applies for traffic to the UE. - UPLINK: The corresponding filter applies for traffic from the UE. - BIDIRECTIONAL: The corresponding filter applies for traffic both to and from the UE. - UNSPECIFIED: The corresponding filter applies for traffic to the UE (downlink), but has no  specific direction declared. The service data flow detection shall apply the filter for  uplink traffic as if the filter was bidirectional. The PCF shall not use the value  UNSPECIFIED in filters created by the network in NW-initiated procedures. The PCF shall only  include the value UNSPECIFIED in filters in UE-initiated procedures if the same value is  received from the SMF. 
type FlowDirection struct {
	FlowDirectionAnyOf *FlowDirectionAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FlowDirection) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FlowDirectionAnyOf
	err = json.Unmarshal(data, &dst.FlowDirectionAnyOf);
	if err == nil {
		jsonFlowDirectionAnyOf, _ := json.Marshal(dst.FlowDirectionAnyOf)
		if string(jsonFlowDirectionAnyOf) == "{}" { // empty struct
			dst.FlowDirectionAnyOf = nil
		} else {
			return nil // data stored in dst.FlowDirectionAnyOf, return on the first match
		}
	} else {
		dst.FlowDirectionAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(FlowDirection)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FlowDirection) MarshalJSON() ([]byte, error) {
	if src.FlowDirectionAnyOf != nil {
		return json.Marshal(&src.FlowDirectionAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFlowDirection struct {
	value *FlowDirection
	isSet bool
}

func (v NullableFlowDirection) Get() *FlowDirection {
	return v.value
}

func (v *NullableFlowDirection) Set(val *FlowDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowDirection(val *FlowDirection) *NullableFlowDirection {
	return &NullableFlowDirection{value: val, isSet: true}
}

func (v NullableFlowDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UEAU

import (
	"encoding/json"
	"fmt"
)

// NodeType struct for NodeType
type NodeType struct {
	NodeTypeAnyOf *NodeTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NodeType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NodeTypeAnyOf
	err = json.Unmarshal(data, &dst.NodeTypeAnyOf);
	if err == nil {
		jsonNodeTypeAnyOf, _ := json.Marshal(dst.NodeTypeAnyOf)
		if string(jsonNodeTypeAnyOf) == "{}" { // empty struct
			dst.NodeTypeAnyOf = nil
		} else {
			return nil // data stored in dst.NodeTypeAnyOf, return on the first match
		}
	} else {
		dst.NodeTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NodeType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NodeType) MarshalJSON() ([]byte, error) {
	if src.NodeTypeAnyOf != nil {
		return json.Marshal(&src.NodeTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNodeType struct {
	value *NodeType
	isSet bool
}

func (v NullableNodeType) Get() *NodeType {
	return v.value
}

func (v *NullableNodeType) Set(val *NodeType) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeType) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeType(val *NodeType) *NullableNodeType {
	return &NullableNodeType{value: val, isSet: true}
}

func (v NullableNodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



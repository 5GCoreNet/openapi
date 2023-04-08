/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// RequestedNode Represents the type of serving node for which certain data is requested
type RequestedNode struct {
	RequestedNodeAnyOf *RequestedNodeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RequestedNode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RequestedNodeAnyOf
	err = json.Unmarshal(data, &dst.RequestedNodeAnyOf);
	if err == nil {
		jsonRequestedNodeAnyOf, _ := json.Marshal(dst.RequestedNodeAnyOf)
		if string(jsonRequestedNodeAnyOf) == "{}" { // empty struct
			dst.RequestedNodeAnyOf = nil
		} else {
			return nil // data stored in dst.RequestedNodeAnyOf, return on the first match
		}
	} else {
		dst.RequestedNodeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RequestedNode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RequestedNode) MarshalJSON() ([]byte, error) {
	if src.RequestedNodeAnyOf != nil {
		return json.Marshal(&src.RequestedNodeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRequestedNode struct {
	value *RequestedNode
	isSet bool
}

func (v NullableRequestedNode) Get() *RequestedNode {
	return v.value
}

func (v *NullableRequestedNode) Set(val *RequestedNode) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedNode) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedNode(val *RequestedNode) *NullableRequestedNode {
	return &NullableRequestedNode{value: val, isSet: true}
}

func (v NullableRequestedNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestedNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



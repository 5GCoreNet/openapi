/*
CAPIF_Publish_Service_API

API for publishing service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Publish_Service_API

import (
	"encoding/json"
	"fmt"
)

// Protocol Possible values are: - HTTP_1_1: HTTP version 1.1 - HTTP_2: HTTP version 2 
type Protocol struct {
	ProtocolAnyOf *ProtocolAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Protocol) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ProtocolAnyOf
	err = json.Unmarshal(data, &dst.ProtocolAnyOf);
	if err == nil {
		jsonProtocolAnyOf, _ := json.Marshal(dst.ProtocolAnyOf)
		if string(jsonProtocolAnyOf) == "{}" { // empty struct
			dst.ProtocolAnyOf = nil
		} else {
			return nil // data stored in dst.ProtocolAnyOf, return on the first match
		}
	} else {
		dst.ProtocolAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(Protocol)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Protocol) MarshalJSON() ([]byte, error) {
	if src.ProtocolAnyOf != nil {
		return json.Marshal(&src.ProtocolAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProtocol struct {
	value *Protocol
	isSet bool
}

func (v NullableProtocol) Get() *Protocol {
	return v.value
}

func (v *NullableProtocol) Set(val *Protocol) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocol(val *Protocol) *NullableProtocol {
	return &NullableProtocol{value: val, isSet: true}
}

func (v NullableProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



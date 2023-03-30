/*
Namf_EventExposure

AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// TransportProtocol Possible values are: - UDP: User Datagram Protocol. - TCP: Transmission Control Protocol.  
type TransportProtocol struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TransportProtocol) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TransportProtocol)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TransportProtocol) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTransportProtocol struct {
	value *TransportProtocol
	isSet bool
}

func (v NullableTransportProtocol) Get() *TransportProtocol {
	return v.value
}

func (v *NullableTransportProtocol) Set(val *TransportProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportProtocol(val *TransportProtocol) *NullableTransportProtocol {
	return &NullableTransportProtocol{value: val, isSet: true}
}

func (v NullableTransportProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



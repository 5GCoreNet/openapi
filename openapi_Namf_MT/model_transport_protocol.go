/*
Namf_MT

AMF Mobile Terminated Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MT

import (
	"encoding/json"
	"fmt"
)

// TransportProtocol Possible values are: - UDP: User Datagram Protocol. - TCP: Transmission Control Protocol.  
type TransportProtocol struct {
	TransportProtocolAnyOf *TransportProtocolAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TransportProtocol) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TransportProtocolAnyOf
	err = json.Unmarshal(data, &dst.TransportProtocolAnyOf);
	if err == nil {
		jsonTransportProtocolAnyOf, _ := json.Marshal(dst.TransportProtocolAnyOf)
		if string(jsonTransportProtocolAnyOf) == "{}" { // empty struct
			dst.TransportProtocolAnyOf = nil
		} else {
			return nil // data stored in dst.TransportProtocolAnyOf, return on the first match
		}
	} else {
		dst.TransportProtocolAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(TransportProtocol)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TransportProtocol) MarshalJSON() ([]byte, error) {
	if src.TransportProtocolAnyOf != nil {
		return json.Marshal(&src.TransportProtocolAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
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



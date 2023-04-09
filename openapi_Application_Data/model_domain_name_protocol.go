/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Application_Data

import (
	"encoding/json"
	"fmt"
)

// DomainNameProtocol Possible values are - DNS_QNAME: Identifies the DNS protocol and the question name in DNS query. - TLS_SNI: Identifies the Server Name Indication in TLS ClientHello message. - TLS_SAN: Identifies the Subject Alternative Name in TLS ServerCertificate message. - TLS_SCN: Identifies the Subject Common Name in TLS ServerCertificate message.
type DomainNameProtocol struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DomainNameProtocol) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(DomainNameProtocol)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DomainNameProtocol) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDomainNameProtocol struct {
	value *DomainNameProtocol
	isSet bool
}

func (v NullableDomainNameProtocol) Get() *DomainNameProtocol {
	return v.value
}

func (v *NullableDomainNameProtocol) Set(val *DomainNameProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainNameProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainNameProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainNameProtocol(val *DomainNameProtocol) *NullableDomainNameProtocol {
	return &NullableDomainNameProtocol{value: val, isSet: true}
}

func (v NullableDomainNameProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainNameProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

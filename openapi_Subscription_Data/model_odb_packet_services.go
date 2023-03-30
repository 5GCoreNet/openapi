/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// OdbPacketServices The enumeration OdbPacketServices defines the Barring of Packet Oriented Services. See 3GPP TS 23.015 for further description. It shall comply with the provisions defined in table 5.7.3.2-1 
type OdbPacketServices struct {
	NullValue *NullValue
	OdbPacketServicesAnyOf *OdbPacketServicesAnyOf
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *OdbPacketServices) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	// try to unmarshal JSON data into OdbPacketServicesAnyOf
	err = json.Unmarshal(data, &dst.OdbPacketServicesAnyOf);
	if err == nil {
		jsonOdbPacketServicesAnyOf, _ := json.Marshal(dst.OdbPacketServicesAnyOf)
		if string(jsonOdbPacketServicesAnyOf) == "{}" { // empty struct
			dst.OdbPacketServicesAnyOf = nil
		} else {
			return nil // data stored in dst.OdbPacketServicesAnyOf, return on the first match
		}
	} else {
		dst.OdbPacketServicesAnyOf = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(OdbPacketServices)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *OdbPacketServices) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.OdbPacketServicesAnyOf != nil {
		return json.Marshal(&src.OdbPacketServicesAnyOf)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableOdbPacketServices struct {
	value *OdbPacketServices
	isSet bool
}

func (v NullableOdbPacketServices) Get() *OdbPacketServices {
	return v.value
}

func (v *NullableOdbPacketServices) Set(val *OdbPacketServices) {
	v.value = val
	v.isSet = true
}

func (v NullableOdbPacketServices) IsSet() bool {
	return v.isSet
}

func (v *NullableOdbPacketServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOdbPacketServices(val *OdbPacketServices) *NullableOdbPacketServices {
	return &NullableOdbPacketServices{value: val, isSet: true}
}

func (v NullableOdbPacketServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOdbPacketServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



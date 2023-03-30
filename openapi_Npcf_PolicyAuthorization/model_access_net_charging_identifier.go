/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// AccessNetChargingIdentifier - Describes the access network charging identifier.
type AccessNetChargingIdentifier struct {
	Interface{} *interface{}
}

// interface{}AsAccessNetChargingIdentifier is a convenience function that returns interface{} wrapped in AccessNetChargingIdentifier
func Interface{}AsAccessNetChargingIdentifier(v *interface{}) AccessNetChargingIdentifier {
	return AccessNetChargingIdentifier{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AccessNetChargingIdentifier) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AccessNetChargingIdentifier)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AccessNetChargingIdentifier)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AccessNetChargingIdentifier) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AccessNetChargingIdentifier) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableAccessNetChargingIdentifier struct {
	value *AccessNetChargingIdentifier
	isSet bool
}

func (v NullableAccessNetChargingIdentifier) Get() *AccessNetChargingIdentifier {
	return v.value
}

func (v *NullableAccessNetChargingIdentifier) Set(val *AccessNetChargingIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessNetChargingIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessNetChargingIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessNetChargingIdentifier(val *AccessNetChargingIdentifier) *NullableAccessNetChargingIdentifier {
	return &NullableAccessNetChargingIdentifier{value: val, isSet: true}
}

func (v NullableAccessNetChargingIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessNetChargingIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



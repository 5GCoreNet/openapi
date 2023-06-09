/*
AUSF API

AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nausf_UEAuthentication

import (
	"encoding/json"
	"fmt"
)

// UEAuthenticationCtx5gAuthData - struct for UEAuthenticationCtx5gAuthData
type UEAuthenticationCtx5gAuthData struct {
	Av5gAka *Av5gAka
	String  *string
}

// Av5gAkaAsUEAuthenticationCtx5gAuthData is a convenience function that returns Av5gAka wrapped in UEAuthenticationCtx5gAuthData
func Av5gAkaAsUEAuthenticationCtx5gAuthData(v *Av5gAka) UEAuthenticationCtx5gAuthData {
	return UEAuthenticationCtx5gAuthData{
		Av5gAka: v,
	}
}

// stringAsUEAuthenticationCtx5gAuthData is a convenience function that returns string wrapped in UEAuthenticationCtx5gAuthData
func StringAsUEAuthenticationCtx5gAuthData(v *string) UEAuthenticationCtx5gAuthData {
	return UEAuthenticationCtx5gAuthData{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UEAuthenticationCtx5gAuthData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Av5gAka
	err = newStrictDecoder(data).Decode(&dst.Av5gAka)
	if err == nil {
		jsonAv5gAka, _ := json.Marshal(dst.Av5gAka)
		if string(jsonAv5gAka) == "{}" { // empty struct
			dst.Av5gAka = nil
		} else {
			match++
		}
	} else {
		dst.Av5gAka = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Av5gAka = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UEAuthenticationCtx5gAuthData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UEAuthenticationCtx5gAuthData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UEAuthenticationCtx5gAuthData) MarshalJSON() ([]byte, error) {
	if src.Av5gAka != nil {
		return json.Marshal(&src.Av5gAka)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UEAuthenticationCtx5gAuthData) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Av5gAka != nil {
		return obj.Av5gAka
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUEAuthenticationCtx5gAuthData struct {
	value *UEAuthenticationCtx5gAuthData
	isSet bool
}

func (v NullableUEAuthenticationCtx5gAuthData) Get() *UEAuthenticationCtx5gAuthData {
	return v.value
}

func (v *NullableUEAuthenticationCtx5gAuthData) Set(val *UEAuthenticationCtx5gAuthData) {
	v.value = val
	v.isSet = true
}

func (v NullableUEAuthenticationCtx5gAuthData) IsSet() bool {
	return v.isSet
}

func (v *NullableUEAuthenticationCtx5gAuthData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUEAuthenticationCtx5gAuthData(val *UEAuthenticationCtx5gAuthData) *NullableUEAuthenticationCtx5gAuthData {
	return &NullableUEAuthenticationCtx5gAuthData{value: val, isSet: true}
}

func (v NullableUEAuthenticationCtx5gAuthData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUEAuthenticationCtx5gAuthData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

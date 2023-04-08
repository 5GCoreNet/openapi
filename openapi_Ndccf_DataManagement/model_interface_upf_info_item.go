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

// InterfaceUpfInfoItem Information of a given IP interface of an UPF
type InterfaceUpfInfoItem struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *InterfaceUpfInfoItem) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.Interface);
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			return nil // data stored in dst.Interface, return on the first match
		}
	} else {
		dst.Interface = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(InterfaceUpfInfoItem)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *InterfaceUpfInfoItem) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableInterfaceUpfInfoItem struct {
	value *InterfaceUpfInfoItem
	isSet bool
}

func (v NullableInterfaceUpfInfoItem) Get() *InterfaceUpfInfoItem {
	return v.value
}

func (v *NullableInterfaceUpfInfoItem) Set(val *InterfaceUpfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceUpfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceUpfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceUpfInfoItem(val *InterfaceUpfInfoItem) *NullableInterfaceUpfInfoItem {
	return &NullableInterfaceUpfInfoItem{value: val, isSet: true}
}

func (v NullableInterfaceUpfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceUpfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



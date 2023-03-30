/*
M5_ConsumptionReporting

5GMS AF M5 Consumption Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M5_ConsumptionReporting

import (
	"encoding/json"
	"fmt"
)

// CellIdentifierType struct for CellIdentifierType
type CellIdentifierType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CellIdentifierType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(CellIdentifierType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CellIdentifierType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCellIdentifierType struct {
	value *CellIdentifierType
	isSet bool
}

func (v NullableCellIdentifierType) Get() *CellIdentifierType {
	return v.value
}

func (v *NullableCellIdentifierType) Set(val *CellIdentifierType) {
	v.value = val
	v.isSet = true
}

func (v NullableCellIdentifierType) IsSet() bool {
	return v.isSet
}

func (v *NullableCellIdentifierType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCellIdentifierType(val *CellIdentifierType) *NullableCellIdentifierType {
	return &NullableCellIdentifierType{value: val, isSet: true}
}

func (v NullableCellIdentifierType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCellIdentifierType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



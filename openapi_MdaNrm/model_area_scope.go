/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
	"fmt"
)

// AreaScope - struct for AreaScope
type AreaScope struct {
	ArrayOfEutraCellId *[]EutraCellId
	ArrayOfNrCellId *[]NrCellId
	ArrayOfTai *[]Tai
	ArrayOfString *[]string
}

// []EutraCellIdAsAreaScope is a convenience function that returns []EutraCellId wrapped in AreaScope
func ArrayOfEutraCellIdAsAreaScope(v *[]EutraCellId) AreaScope {
	return AreaScope{
		ArrayOfEutraCellId: v,
	}
}

// []NrCellIdAsAreaScope is a convenience function that returns []NrCellId wrapped in AreaScope
func ArrayOfNrCellIdAsAreaScope(v *[]NrCellId) AreaScope {
	return AreaScope{
		ArrayOfNrCellId: v,
	}
}

// []TaiAsAreaScope is a convenience function that returns []Tai wrapped in AreaScope
func ArrayOfTaiAsAreaScope(v *[]Tai) AreaScope {
	return AreaScope{
		ArrayOfTai: v,
	}
}

// []stringAsAreaScope is a convenience function that returns []string wrapped in AreaScope
func ArrayOfStringAsAreaScope(v *[]string) AreaScope {
	return AreaScope{
		ArrayOfString: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AreaScope) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfEutraCellId
	err = newStrictDecoder(data).Decode(&dst.ArrayOfEutraCellId)
	if err == nil {
		jsonArrayOfEutraCellId, _ := json.Marshal(dst.ArrayOfEutraCellId)
		if string(jsonArrayOfEutraCellId) == "{}" { // empty struct
			dst.ArrayOfEutraCellId = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfEutraCellId = nil
	}

	// try to unmarshal data into ArrayOfNrCellId
	err = newStrictDecoder(data).Decode(&dst.ArrayOfNrCellId)
	if err == nil {
		jsonArrayOfNrCellId, _ := json.Marshal(dst.ArrayOfNrCellId)
		if string(jsonArrayOfNrCellId) == "{}" { // empty struct
			dst.ArrayOfNrCellId = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfNrCellId = nil
	}

	// try to unmarshal data into ArrayOfTai
	err = newStrictDecoder(data).Decode(&dst.ArrayOfTai)
	if err == nil {
		jsonArrayOfTai, _ := json.Marshal(dst.ArrayOfTai)
		if string(jsonArrayOfTai) == "{}" { // empty struct
			dst.ArrayOfTai = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfTai = nil
	}

	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfString = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfEutraCellId = nil
		dst.ArrayOfNrCellId = nil
		dst.ArrayOfTai = nil
		dst.ArrayOfString = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AreaScope)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AreaScope)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AreaScope) MarshalJSON() ([]byte, error) {
	if src.ArrayOfEutraCellId != nil {
		return json.Marshal(&src.ArrayOfEutraCellId)
	}

	if src.ArrayOfNrCellId != nil {
		return json.Marshal(&src.ArrayOfNrCellId)
	}

	if src.ArrayOfTai != nil {
		return json.Marshal(&src.ArrayOfTai)
	}

	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AreaScope) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfEutraCellId != nil {
		return obj.ArrayOfEutraCellId
	}

	if obj.ArrayOfNrCellId != nil {
		return obj.ArrayOfNrCellId
	}

	if obj.ArrayOfTai != nil {
		return obj.ArrayOfTai
	}

	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	// all schemas are nil
	return nil
}

type NullableAreaScope struct {
	value *AreaScope
	isSet bool
}

func (v NullableAreaScope) Get() *AreaScope {
	return v.value
}

func (v *NullableAreaScope) Set(val *AreaScope) {
	v.value = val
	v.isSet = true
}

func (v NullableAreaScope) IsSet() bool {
	return v.isSet
}

func (v *NullableAreaScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAreaScope(val *AreaScope) *NullableAreaScope {
	return &NullableAreaScope{value: val, isSet: true}
}

func (v NullableAreaScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAreaScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



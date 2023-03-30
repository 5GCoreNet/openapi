/*
Common Type Definitions

OAS 3.0.1 specification of common type definitions in the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ComDefs

import (
	"encoding/json"
	"fmt"
)

// AreaOfInterest - struct for AreaOfInterest
type AreaOfInterest struct {
	GeoAreaToCellMapping *GeoAreaToCellMapping
	ArrayOfTai *[]Tai
	ArrayOfInt32 *[]int32
	ArrayOfString *[]string
}

// GeoAreaToCellMappingAsAreaOfInterest is a convenience function that returns GeoAreaToCellMapping wrapped in AreaOfInterest
func GeoAreaToCellMappingAsAreaOfInterest(v *GeoAreaToCellMapping) AreaOfInterest {
	return AreaOfInterest{
		GeoAreaToCellMapping: v,
	}
}

// []TaiAsAreaOfInterest is a convenience function that returns []Tai wrapped in AreaOfInterest
func ArrayOfTaiAsAreaOfInterest(v *[]Tai) AreaOfInterest {
	return AreaOfInterest{
		ArrayOfTai: v,
	}
}

// []int32AsAreaOfInterest is a convenience function that returns []int32 wrapped in AreaOfInterest
func ArrayOfInt32AsAreaOfInterest(v *[]int32) AreaOfInterest {
	return AreaOfInterest{
		ArrayOfInt32: v,
	}
}

// []stringAsAreaOfInterest is a convenience function that returns []string wrapped in AreaOfInterest
func ArrayOfStringAsAreaOfInterest(v *[]string) AreaOfInterest {
	return AreaOfInterest{
		ArrayOfString: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AreaOfInterest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GeoAreaToCellMapping
	err = newStrictDecoder(data).Decode(&dst.GeoAreaToCellMapping)
	if err == nil {
		jsonGeoAreaToCellMapping, _ := json.Marshal(dst.GeoAreaToCellMapping)
		if string(jsonGeoAreaToCellMapping) == "{}" { // empty struct
			dst.GeoAreaToCellMapping = nil
		} else {
			match++
		}
	} else {
		dst.GeoAreaToCellMapping = nil
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

	// try to unmarshal data into ArrayOfInt32
	err = newStrictDecoder(data).Decode(&dst.ArrayOfInt32)
	if err == nil {
		jsonArrayOfInt32, _ := json.Marshal(dst.ArrayOfInt32)
		if string(jsonArrayOfInt32) == "{}" { // empty struct
			dst.ArrayOfInt32 = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfInt32 = nil
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
		dst.GeoAreaToCellMapping = nil
		dst.ArrayOfTai = nil
		dst.ArrayOfInt32 = nil
		dst.ArrayOfString = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AreaOfInterest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AreaOfInterest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AreaOfInterest) MarshalJSON() ([]byte, error) {
	if src.GeoAreaToCellMapping != nil {
		return json.Marshal(&src.GeoAreaToCellMapping)
	}

	if src.ArrayOfTai != nil {
		return json.Marshal(&src.ArrayOfTai)
	}

	if src.ArrayOfInt32 != nil {
		return json.Marshal(&src.ArrayOfInt32)
	}

	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AreaOfInterest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GeoAreaToCellMapping != nil {
		return obj.GeoAreaToCellMapping
	}

	if obj.ArrayOfTai != nil {
		return obj.ArrayOfTai
	}

	if obj.ArrayOfInt32 != nil {
		return obj.ArrayOfInt32
	}

	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	// all schemas are nil
	return nil
}

type NullableAreaOfInterest struct {
	value *AreaOfInterest
	isSet bool
}

func (v NullableAreaOfInterest) Get() *AreaOfInterest {
	return v.value
}

func (v *NullableAreaOfInterest) Set(val *AreaOfInterest) {
	v.value = val
	v.isSet = true
}

func (v NullableAreaOfInterest) IsSet() bool {
	return v.isSet
}

func (v *NullableAreaOfInterest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAreaOfInterest(val *AreaOfInterest) *NullableAreaOfInterest {
	return &NullableAreaOfInterest{value: val, isSet: true}
}

func (v NullableAreaOfInterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAreaOfInterest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



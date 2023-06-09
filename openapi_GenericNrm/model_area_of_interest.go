/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GenericNrm

import (
	"encoding/json"
	"fmt"
)

// AreaOfInterest - struct for AreaOfInterest
type AreaOfInterest struct {
	GeoAreaToCellMapping *GeoAreaToCellMapping
	ArrayOfTai1          *[]Tai1
	ArrayOfUtraCellId    *[]UtraCellId
	ArrayOfString        *[]string
}

// GeoAreaToCellMappingAsAreaOfInterest is a convenience function that returns GeoAreaToCellMapping wrapped in AreaOfInterest
func GeoAreaToCellMappingAsAreaOfInterest(v *GeoAreaToCellMapping) AreaOfInterest {
	return AreaOfInterest{
		GeoAreaToCellMapping: v,
	}
}

// []Tai1AsAreaOfInterest is a convenience function that returns []Tai1 wrapped in AreaOfInterest
func ArrayOfTai1AsAreaOfInterest(v *[]Tai1) AreaOfInterest {
	return AreaOfInterest{
		ArrayOfTai1: v,
	}
}

// []UtraCellIdAsAreaOfInterest is a convenience function that returns []UtraCellId wrapped in AreaOfInterest
func ArrayOfUtraCellIdAsAreaOfInterest(v *[]UtraCellId) AreaOfInterest {
	return AreaOfInterest{
		ArrayOfUtraCellId: v,
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

	// try to unmarshal data into ArrayOfTai1
	err = newStrictDecoder(data).Decode(&dst.ArrayOfTai1)
	if err == nil {
		jsonArrayOfTai1, _ := json.Marshal(dst.ArrayOfTai1)
		if string(jsonArrayOfTai1) == "{}" { // empty struct
			dst.ArrayOfTai1 = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfTai1 = nil
	}

	// try to unmarshal data into ArrayOfUtraCellId
	err = newStrictDecoder(data).Decode(&dst.ArrayOfUtraCellId)
	if err == nil {
		jsonArrayOfUtraCellId, _ := json.Marshal(dst.ArrayOfUtraCellId)
		if string(jsonArrayOfUtraCellId) == "{}" { // empty struct
			dst.ArrayOfUtraCellId = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfUtraCellId = nil
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
		dst.ArrayOfTai1 = nil
		dst.ArrayOfUtraCellId = nil
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

	if src.ArrayOfTai1 != nil {
		return json.Marshal(&src.ArrayOfTai1)
	}

	if src.ArrayOfUtraCellId != nil {
		return json.Marshal(&src.ArrayOfUtraCellId)
	}

	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AreaOfInterest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GeoAreaToCellMapping != nil {
		return obj.GeoAreaToCellMapping
	}

	if obj.ArrayOfTai1 != nil {
		return obj.ArrayOfTai1
	}

	if obj.ArrayOfUtraCellId != nil {
		return obj.ArrayOfUtraCellId
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

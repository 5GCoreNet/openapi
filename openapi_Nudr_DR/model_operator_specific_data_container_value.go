/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// OperatorSpecificDataContainerValue - struct for OperatorSpecificDataContainerValue
type OperatorSpecificDataContainerValue struct {
	Array *Array
	Bool *bool
	Float32 *float32
	Int32 *int32
	MapOfInterface *map[string]interface{}
	String *string
}

// ArrayAsOperatorSpecificDataContainerValue is a convenience function that returns Array wrapped in OperatorSpecificDataContainerValue
func ArrayAsOperatorSpecificDataContainerValue(v *Array) OperatorSpecificDataContainerValue {
	return OperatorSpecificDataContainerValue{
		Array: v,
	}
}

// boolAsOperatorSpecificDataContainerValue is a convenience function that returns bool wrapped in OperatorSpecificDataContainerValue
func BoolAsOperatorSpecificDataContainerValue(v *bool) OperatorSpecificDataContainerValue {
	return OperatorSpecificDataContainerValue{
		Bool: v,
	}
}

// float32AsOperatorSpecificDataContainerValue is a convenience function that returns float32 wrapped in OperatorSpecificDataContainerValue
func Float32AsOperatorSpecificDataContainerValue(v *float32) OperatorSpecificDataContainerValue {
	return OperatorSpecificDataContainerValue{
		Float32: v,
	}
}

// int32AsOperatorSpecificDataContainerValue is a convenience function that returns int32 wrapped in OperatorSpecificDataContainerValue
func Int32AsOperatorSpecificDataContainerValue(v *int32) OperatorSpecificDataContainerValue {
	return OperatorSpecificDataContainerValue{
		Int32: v,
	}
}

// map[string]interface{}AsOperatorSpecificDataContainerValue is a convenience function that returns map[string]interface{} wrapped in OperatorSpecificDataContainerValue
func MapOfInterfaceAsOperatorSpecificDataContainerValue(v *map[string]interface{}) OperatorSpecificDataContainerValue {
	return OperatorSpecificDataContainerValue{
		MapOfInterface: v,
	}
}

// stringAsOperatorSpecificDataContainerValue is a convenience function that returns string wrapped in OperatorSpecificDataContainerValue
func StringAsOperatorSpecificDataContainerValue(v *string) OperatorSpecificDataContainerValue {
	return OperatorSpecificDataContainerValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *OperatorSpecificDataContainerValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Array
	err = newStrictDecoder(data).Decode(&dst.Array)
	if err == nil {
		jsonArray, _ := json.Marshal(dst.Array)
		if string(jsonArray) == "{}" { // empty struct
			dst.Array = nil
		} else {
			match++
		}
	} else {
		dst.Array = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into MapOfInterface
	err = newStrictDecoder(data).Decode(&dst.MapOfInterface)
	if err == nil {
		jsonMapOfInterface, _ := json.Marshal(dst.MapOfInterface)
		if string(jsonMapOfInterface) == "{}" { // empty struct
			dst.MapOfInterface = nil
		} else {
			match++
		}
	} else {
		dst.MapOfInterface = nil
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
		dst.Array = nil
		dst.Bool = nil
		dst.Float32 = nil
		dst.Int32 = nil
		dst.MapOfInterface = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OperatorSpecificDataContainerValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OperatorSpecificDataContainerValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OperatorSpecificDataContainerValue) MarshalJSON() ([]byte, error) {
	if src.Array != nil {
		return json.Marshal(&src.Array)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.MapOfInterface != nil {
		return json.Marshal(&src.MapOfInterface)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OperatorSpecificDataContainerValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Array != nil {
		return obj.Array
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.MapOfInterface != nil {
		return obj.MapOfInterface
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableOperatorSpecificDataContainerValue struct {
	value *OperatorSpecificDataContainerValue
	isSet bool
}

func (v NullableOperatorSpecificDataContainerValue) Get() *OperatorSpecificDataContainerValue {
	return v.value
}

func (v *NullableOperatorSpecificDataContainerValue) Set(val *OperatorSpecificDataContainerValue) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorSpecificDataContainerValue) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorSpecificDataContainerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorSpecificDataContainerValue(val *OperatorSpecificDataContainerValue) *NullableOperatorSpecificDataContainerValue {
	return &NullableOperatorSpecificDataContainerValue{value: val, isSet: true}
}

func (v NullableOperatorSpecificDataContainerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorSpecificDataContainerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



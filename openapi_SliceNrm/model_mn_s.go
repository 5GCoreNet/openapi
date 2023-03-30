/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
	"fmt"
)

// MnS - struct for MnS
type MnS struct {
	MnSOneOf *MnSOneOf
}

// MnSOneOfAsMnS is a convenience function that returns MnSOneOf wrapped in MnS
func MnSOneOfAsMnS(v *MnSOneOf) MnS {
	return MnS{
		MnSOneOf: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MnS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MnSOneOf
	err = newStrictDecoder(data).Decode(&dst.MnSOneOf)
	if err == nil {
		jsonMnSOneOf, _ := json.Marshal(dst.MnSOneOf)
		if string(jsonMnSOneOf) == "{}" { // empty struct
			dst.MnSOneOf = nil
		} else {
			match++
		}
	} else {
		dst.MnSOneOf = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MnSOneOf = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MnS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MnS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MnS) MarshalJSON() ([]byte, error) {
	if src.MnSOneOf != nil {
		return json.Marshal(&src.MnSOneOf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MnS) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MnSOneOf != nil {
		return obj.MnSOneOf
	}

	// all schemas are nil
	return nil
}

type NullableMnS struct {
	value *MnS
	isSet bool
}

func (v NullableMnS) Get() *MnS {
	return v.value
}

func (v *NullableMnS) Set(val *MnS) {
	v.value = val
	v.isSet = true
}

func (v NullableMnS) IsSet() bool {
	return v.isSet
}

func (v *NullableMnS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMnS(val *MnS) *NullableMnS {
	return &NullableMnS{value: val, isSet: true}
}

func (v NullableMnS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMnS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



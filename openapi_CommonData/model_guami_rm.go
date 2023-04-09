/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"fmt"
)

// GuamiRm This data type is defined in the same way as the 'Guami' data type, but with the OpenAPI 'nullable: true' property.
type GuamiRm struct {
	Guami     *Guami
	NullValue *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GuamiRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Guami
	err = json.Unmarshal(data, &dst.Guami)
	if err == nil {
		jsonGuami, _ := json.Marshal(dst.Guami)
		if string(jsonGuami) == "{}" { // empty struct
			dst.Guami = nil
		} else {
			return nil // data stored in dst.Guami, return on the first match
		}
	} else {
		dst.Guami = nil
	}

	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue)
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

	return fmt.Errorf("data failed to match schemas in anyOf(GuamiRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GuamiRm) MarshalJSON() ([]byte, error) {
	if src.Guami != nil {
		return json.Marshal(&src.Guami)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGuamiRm struct {
	value *GuamiRm
	isSet bool
}

func (v NullableGuamiRm) Get() *GuamiRm {
	return v.value
}

func (v *NullableGuamiRm) Set(val *GuamiRm) {
	v.value = val
	v.isSet = true
}

func (v NullableGuamiRm) IsSet() bool {
	return v.isSet
}

func (v *NullableGuamiRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuamiRm(val *GuamiRm) *NullableGuamiRm {
	return &NullableGuamiRm{value: val, isSet: true}
}

func (v NullableGuamiRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuamiRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

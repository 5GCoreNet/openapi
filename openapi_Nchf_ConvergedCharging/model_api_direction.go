/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// APIDirection struct for APIDirection
type APIDirection struct {
	APIDirectionAnyOf *APIDirectionAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *APIDirection) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into APIDirectionAnyOf
	err = json.Unmarshal(data, &dst.APIDirectionAnyOf);
	if err == nil {
		jsonAPIDirectionAnyOf, _ := json.Marshal(dst.APIDirectionAnyOf)
		if string(jsonAPIDirectionAnyOf) == "{}" { // empty struct
			dst.APIDirectionAnyOf = nil
		} else {
			return nil // data stored in dst.APIDirectionAnyOf, return on the first match
		}
	} else {
		dst.APIDirectionAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(APIDirection)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *APIDirection) MarshalJSON() ([]byte, error) {
	if src.APIDirectionAnyOf != nil {
		return json.Marshal(&src.APIDirectionAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAPIDirection struct {
	value *APIDirection
	isSet bool
}

func (v NullableAPIDirection) Get() *APIDirection {
	return v.value
}

func (v *NullableAPIDirection) Set(val *APIDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIDirection(val *APIDirection) *NullableAPIDirection {
	return &NullableAPIDirection{value: val, isSet: true}
}

func (v NullableAPIDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



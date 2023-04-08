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

// ParticipantActionType struct for ParticipantActionType
type ParticipantActionType struct {
	ParticipantActionTypeAnyOf *ParticipantActionTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ParticipantActionType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ParticipantActionTypeAnyOf
	err = json.Unmarshal(data, &dst.ParticipantActionTypeAnyOf);
	if err == nil {
		jsonParticipantActionTypeAnyOf, _ := json.Marshal(dst.ParticipantActionTypeAnyOf)
		if string(jsonParticipantActionTypeAnyOf) == "{}" { // empty struct
			dst.ParticipantActionTypeAnyOf = nil
		} else {
			return nil // data stored in dst.ParticipantActionTypeAnyOf, return on the first match
		}
	} else {
		dst.ParticipantActionTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ParticipantActionType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ParticipantActionType) MarshalJSON() ([]byte, error) {
	if src.ParticipantActionTypeAnyOf != nil {
		return json.Marshal(&src.ParticipantActionTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableParticipantActionType struct {
	value *ParticipantActionType
	isSet bool
}

func (v NullableParticipantActionType) Get() *ParticipantActionType {
	return v.value
}

func (v *NullableParticipantActionType) Set(val *ParticipantActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableParticipantActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableParticipantActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParticipantActionType(val *ParticipantActionType) *NullableParticipantActionType {
	return &NullableParticipantActionType{value: val, isSet: true}
}

func (v NullableParticipantActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParticipantActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



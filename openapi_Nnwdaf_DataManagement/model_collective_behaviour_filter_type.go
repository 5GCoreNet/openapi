/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// CollectiveBehaviourFilterType Represents collective behaviour parameter type.
type CollectiveBehaviourFilterType struct {
	CollectiveBehaviourFilterTypeAnyOf *CollectiveBehaviourFilterTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CollectiveBehaviourFilterType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CollectiveBehaviourFilterTypeAnyOf
	err = json.Unmarshal(data, &dst.CollectiveBehaviourFilterTypeAnyOf);
	if err == nil {
		jsonCollectiveBehaviourFilterTypeAnyOf, _ := json.Marshal(dst.CollectiveBehaviourFilterTypeAnyOf)
		if string(jsonCollectiveBehaviourFilterTypeAnyOf) == "{}" { // empty struct
			dst.CollectiveBehaviourFilterTypeAnyOf = nil
		} else {
			return nil // data stored in dst.CollectiveBehaviourFilterTypeAnyOf, return on the first match
		}
	} else {
		dst.CollectiveBehaviourFilterTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CollectiveBehaviourFilterType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CollectiveBehaviourFilterType) MarshalJSON() ([]byte, error) {
	if src.CollectiveBehaviourFilterTypeAnyOf != nil {
		return json.Marshal(&src.CollectiveBehaviourFilterTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCollectiveBehaviourFilterType struct {
	value *CollectiveBehaviourFilterType
	isSet bool
}

func (v NullableCollectiveBehaviourFilterType) Get() *CollectiveBehaviourFilterType {
	return v.value
}

func (v *NullableCollectiveBehaviourFilterType) Set(val *CollectiveBehaviourFilterType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectiveBehaviourFilterType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectiveBehaviourFilterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectiveBehaviourFilterType(val *CollectiveBehaviourFilterType) *NullableCollectiveBehaviourFilterType {
	return &NullableCollectiveBehaviourFilterType{value: val, isSet: true}
}

func (v NullableCollectiveBehaviourFilterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectiveBehaviourFilterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



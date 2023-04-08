/*
Naf_ProSe API

Naf_ProSe Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_ProSe

import (
	"encoding/json"
	"fmt"
)

// MetadataIndic Possible values are: - NO_METADATA: This value may be used to indicate that there is no metadata associated with the target RPAUID. This is the default value applicable if this IE is not supplied. - METADATA_UPDATE_DISALLOWED: This value shall be used to indicate that there exists metadata associated with the target RPAUID, but the metadata is not allowed to be updated. - METADATA_UPDATE_ALLOWED: This value shall be used to indicate that there exists metadata associated with the target RPAUID, and the metadata is allowed to be updated. 
type MetadataIndic struct {
	MetadataIndicAnyOf *MetadataIndicAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MetadataIndic) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MetadataIndicAnyOf
	err = json.Unmarshal(data, &dst.MetadataIndicAnyOf);
	if err == nil {
		jsonMetadataIndicAnyOf, _ := json.Marshal(dst.MetadataIndicAnyOf)
		if string(jsonMetadataIndicAnyOf) == "{}" { // empty struct
			dst.MetadataIndicAnyOf = nil
		} else {
			return nil // data stored in dst.MetadataIndicAnyOf, return on the first match
		}
	} else {
		dst.MetadataIndicAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MetadataIndic)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MetadataIndic) MarshalJSON() ([]byte, error) {
	if src.MetadataIndicAnyOf != nil {
		return json.Marshal(&src.MetadataIndicAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMetadataIndic struct {
	value *MetadataIndic
	isSet bool
}

func (v NullableMetadataIndic) Get() *MetadataIndic {
	return v.value
}

func (v *NullableMetadataIndic) Set(val *MetadataIndic) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataIndic) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataIndic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataIndic(val *MetadataIndic) *NullableMetadataIndic {
	return &NullableMetadataIndic{value: val, isSet: true}
}

func (v NullableMetadataIndic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataIndic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



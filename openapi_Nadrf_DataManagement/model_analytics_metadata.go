/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// AnalyticsMetadata Possible values are: - NUM_OF_SAMPLES: Number of data samples used for the generation of the output analytics. - DATA_WINDOW: Data time window of the data samples. - DATA_STAT_PROPS: Dataset statistical properties of the data used to generate the analytics. - STRATEGY: Output strategy used for the reporting of the analytics. - ACCURACY: Level of accuracy reached for the analytics. 
type AnalyticsMetadata struct {
	AnalyticsMetadataAnyOf *AnalyticsMetadataAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AnalyticsMetadata) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AnalyticsMetadataAnyOf
	err = json.Unmarshal(data, &dst.AnalyticsMetadataAnyOf);
	if err == nil {
		jsonAnalyticsMetadataAnyOf, _ := json.Marshal(dst.AnalyticsMetadataAnyOf)
		if string(jsonAnalyticsMetadataAnyOf) == "{}" { // empty struct
			dst.AnalyticsMetadataAnyOf = nil
		} else {
			return nil // data stored in dst.AnalyticsMetadataAnyOf, return on the first match
		}
	} else {
		dst.AnalyticsMetadataAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AnalyticsMetadata)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AnalyticsMetadata) MarshalJSON() ([]byte, error) {
	if src.AnalyticsMetadataAnyOf != nil {
		return json.Marshal(&src.AnalyticsMetadataAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAnalyticsMetadata struct {
	value *AnalyticsMetadata
	isSet bool
}

func (v NullableAnalyticsMetadata) Get() *AnalyticsMetadata {
	return v.value
}

func (v *NullableAnalyticsMetadata) Set(val *AnalyticsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsMetadata(val *AnalyticsMetadata) *NullableAnalyticsMetadata {
	return &NullableAnalyticsMetadata{value: val, isSet: true}
}

func (v NullableAnalyticsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



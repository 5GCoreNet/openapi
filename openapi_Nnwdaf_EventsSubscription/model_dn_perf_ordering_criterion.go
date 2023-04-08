/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
	"fmt"
)

// DnPerfOrderingCriterion Possible values are:   - AVERAGE_TRAFFIC_RATE: Indicates the average traffic rate.   - MAXIMUM_TRAFFIC_RATE: Indicates the maximum traffic rate.   - AVERAGE_PACKET_DELAY: Indicates the average packet delay.   - MAXIMUM_PACKET_DELAY: Indicates the maximum packet delay.   - AVERAGE_PACKET_LOSS_RATE: Indicates the average packet loss rate. 
type DnPerfOrderingCriterion struct {
	DnPerfOrderingCriterionAnyOf *DnPerfOrderingCriterionAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DnPerfOrderingCriterion) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DnPerfOrderingCriterionAnyOf
	err = json.Unmarshal(data, &dst.DnPerfOrderingCriterionAnyOf);
	if err == nil {
		jsonDnPerfOrderingCriterionAnyOf, _ := json.Marshal(dst.DnPerfOrderingCriterionAnyOf)
		if string(jsonDnPerfOrderingCriterionAnyOf) == "{}" { // empty struct
			dst.DnPerfOrderingCriterionAnyOf = nil
		} else {
			return nil // data stored in dst.DnPerfOrderingCriterionAnyOf, return on the first match
		}
	} else {
		dst.DnPerfOrderingCriterionAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DnPerfOrderingCriterion)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DnPerfOrderingCriterion) MarshalJSON() ([]byte, error) {
	if src.DnPerfOrderingCriterionAnyOf != nil {
		return json.Marshal(&src.DnPerfOrderingCriterionAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDnPerfOrderingCriterion struct {
	value *DnPerfOrderingCriterion
	isSet bool
}

func (v NullableDnPerfOrderingCriterion) Get() *DnPerfOrderingCriterion {
	return v.value
}

func (v *NullableDnPerfOrderingCriterion) Set(val *DnPerfOrderingCriterion) {
	v.value = val
	v.isSet = true
}

func (v NullableDnPerfOrderingCriterion) IsSet() bool {
	return v.isSet
}

func (v *NullableDnPerfOrderingCriterion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnPerfOrderingCriterion(val *DnPerfOrderingCriterion) *NullableDnPerfOrderingCriterion {
	return &NullableDnPerfOrderingCriterion{value: val, isSet: true}
}

func (v NullableDnPerfOrderingCriterion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnPerfOrderingCriterion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



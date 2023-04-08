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

// WlanOrderingCriterion Possible values are: - TIME_SLOT_START: Indicates the order of time slot start. - NUMBER_OF_UES: Indicates the order of number of UEs. - RSSI: Indicates the order of RSSI. - RTT: Indicates the order of RTT. - TRAFFIC_INFO: Indicates the order of Traffic information. 
type WlanOrderingCriterion struct {
	WlanOrderingCriterionAnyOf *WlanOrderingCriterionAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *WlanOrderingCriterion) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into WlanOrderingCriterionAnyOf
	err = json.Unmarshal(data, &dst.WlanOrderingCriterionAnyOf);
	if err == nil {
		jsonWlanOrderingCriterionAnyOf, _ := json.Marshal(dst.WlanOrderingCriterionAnyOf)
		if string(jsonWlanOrderingCriterionAnyOf) == "{}" { // empty struct
			dst.WlanOrderingCriterionAnyOf = nil
		} else {
			return nil // data stored in dst.WlanOrderingCriterionAnyOf, return on the first match
		}
	} else {
		dst.WlanOrderingCriterionAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(WlanOrderingCriterion)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *WlanOrderingCriterion) MarshalJSON() ([]byte, error) {
	if src.WlanOrderingCriterionAnyOf != nil {
		return json.Marshal(&src.WlanOrderingCriterionAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableWlanOrderingCriterion struct {
	value *WlanOrderingCriterion
	isSet bool
}

func (v NullableWlanOrderingCriterion) Get() *WlanOrderingCriterion {
	return v.value
}

func (v *NullableWlanOrderingCriterion) Set(val *WlanOrderingCriterion) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanOrderingCriterion) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanOrderingCriterion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanOrderingCriterion(val *WlanOrderingCriterion) *NullableWlanOrderingCriterion {
	return &NullableWlanOrderingCriterion{value: val, isSet: true}
}

func (v NullableWlanOrderingCriterion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanOrderingCriterion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



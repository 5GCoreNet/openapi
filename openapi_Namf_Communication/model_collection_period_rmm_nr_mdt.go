/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodRmmNrMdt The enumeration CollectionPeriodRmmNrMdt defines Collection period for RRM measurements NR for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.19-1
type CollectionPeriodRmmNrMdt struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CollectionPeriodRmmNrMdt) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
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

	return fmt.Errorf("data failed to match schemas in anyOf(CollectionPeriodRmmNrMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CollectionPeriodRmmNrMdt) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCollectionPeriodRmmNrMdt struct {
	value *CollectionPeriodRmmNrMdt
	isSet bool
}

func (v NullableCollectionPeriodRmmNrMdt) Get() *CollectionPeriodRmmNrMdt {
	return v.value
}

func (v *NullableCollectionPeriodRmmNrMdt) Set(val *CollectionPeriodRmmNrMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodRmmNrMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodRmmNrMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodRmmNrMdt(val *CollectionPeriodRmmNrMdt) *NullableCollectionPeriodRmmNrMdt {
	return &NullableCollectionPeriodRmmNrMdt{value: val, isSet: true}
}

func (v NullableCollectionPeriodRmmNrMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodRmmNrMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

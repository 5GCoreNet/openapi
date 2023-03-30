/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// DataNotification - Represents a Data Subscription Notification.
type DataNotification struct {
	Interface{} *interface{}
}

// interface{}AsDataNotification is a convenience function that returns interface{} wrapped in DataNotification
func Interface{}AsDataNotification(v *interface{}) DataNotification {
	return DataNotification{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DataNotification) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DataNotification)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DataNotification)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DataNotification) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DataNotification) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableDataNotification struct {
	value *DataNotification
	isSet bool
}

func (v NullableDataNotification) Get() *DataNotification {
	return v.value
}

func (v *NullableDataNotification) Set(val *DataNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableDataNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableDataNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataNotification(val *DataNotification) *NullableDataNotification {
	return &NullableDataNotification{value: val, isSet: true}
}

func (v NullableDataNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



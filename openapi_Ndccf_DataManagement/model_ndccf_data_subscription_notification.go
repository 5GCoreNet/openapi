/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// NdccfDataSubscriptionNotification - Represents a notification for a DCCF data subscription.
type NdccfDataSubscriptionNotification struct {
	Interface *interface{}
}

// interface{}AsNdccfDataSubscriptionNotification is a convenience function that returns interface{} wrapped in NdccfDataSubscriptionNotification
func InterfaceAsNdccfDataSubscriptionNotification(v *interface{}) NdccfDataSubscriptionNotification {
	return NdccfDataSubscriptionNotification{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NdccfDataSubscriptionNotification) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface
	err = newStrictDecoder(data).Decode(&dst.Interface)
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			match++
		}
	} else {
		dst.Interface = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface = nil

		return fmt.Errorf("data matches more than one schema in oneOf(NdccfDataSubscriptionNotification)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NdccfDataSubscriptionNotification)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NdccfDataSubscriptionNotification) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NdccfDataSubscriptionNotification) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableNdccfDataSubscriptionNotification struct {
	value *NdccfDataSubscriptionNotification
	isSet bool
}

func (v NullableNdccfDataSubscriptionNotification) Get() *NdccfDataSubscriptionNotification {
	return v.value
}

func (v *NullableNdccfDataSubscriptionNotification) Set(val *NdccfDataSubscriptionNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableNdccfDataSubscriptionNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNdccfDataSubscriptionNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNdccfDataSubscriptionNotification(val *NdccfDataSubscriptionNotification) *NullableNdccfDataSubscriptionNotification {
	return &NullableNdccfDataSubscriptionNotification{value: val, isSet: true}
}

func (v NullableNdccfDataSubscriptionNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNdccfDataSubscriptionNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

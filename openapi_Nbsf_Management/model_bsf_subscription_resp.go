/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.4.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsf_Management

import (
	"encoding/json"
	"fmt"
)

// BsfSubscriptionResp It represents a response to a modification or creation request of an Individual Binding Subscription resource. It may contain the notification of the already met events. 
type BsfSubscriptionResp struct {
	BsfNotification *BsfNotification
	BsfSubscription *BsfSubscription
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *BsfSubscriptionResp) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BsfNotification
	err = json.Unmarshal(data, &dst.BsfNotification);
	if err == nil {
		jsonBsfNotification, _ := json.Marshal(dst.BsfNotification)
		if string(jsonBsfNotification) == "{}" { // empty struct
			dst.BsfNotification = nil
		} else {
			return nil // data stored in dst.BsfNotification, return on the first match
		}
	} else {
		dst.BsfNotification = nil
	}

	// try to unmarshal JSON data into BsfSubscription
	err = json.Unmarshal(data, &dst.BsfSubscription);
	if err == nil {
		jsonBsfSubscription, _ := json.Marshal(dst.BsfSubscription)
		if string(jsonBsfSubscription) == "{}" { // empty struct
			dst.BsfSubscription = nil
		} else {
			return nil // data stored in dst.BsfSubscription, return on the first match
		}
	} else {
		dst.BsfSubscription = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(BsfSubscriptionResp)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *BsfSubscriptionResp) MarshalJSON() ([]byte, error) {
	if src.BsfNotification != nil {
		return json.Marshal(&src.BsfNotification)
	}

	if src.BsfSubscription != nil {
		return json.Marshal(&src.BsfSubscription)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableBsfSubscriptionResp struct {
	value *BsfSubscriptionResp
	isSet bool
}

func (v NullableBsfSubscriptionResp) Get() *BsfSubscriptionResp {
	return v.value
}

func (v *NullableBsfSubscriptionResp) Set(val *BsfSubscriptionResp) {
	v.value = val
	v.isSet = true
}

func (v NullableBsfSubscriptionResp) IsSet() bool {
	return v.isSet
}

func (v *NullableBsfSubscriptionResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBsfSubscriptionResp(val *BsfSubscriptionResp) *NullableBsfSubscriptionResp {
	return &NullableBsfSubscriptionResp{value: val, isSet: true}
}

func (v NullableBsfSubscriptionResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBsfSubscriptionResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


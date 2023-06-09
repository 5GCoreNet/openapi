/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceAdaptation

import (
	"encoding/json"
	"fmt"
)

// DeliveryMode Possible values are: - UNICAST: Unicast delivery. - MULTICAST: Multicast delivery.
type DeliveryMode struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DeliveryMode) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(DeliveryMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DeliveryMode) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDeliveryMode struct {
	value *DeliveryMode
	isSet bool
}

func (v NullableDeliveryMode) Get() *DeliveryMode {
	return v.value
}

func (v *NullableDeliveryMode) Set(val *DeliveryMode) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryMode) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryMode(val *DeliveryMode) *NullableDeliveryMode {
	return &NullableDeliveryMode{value: val, isSet: true}
}

func (v NullableDeliveryMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

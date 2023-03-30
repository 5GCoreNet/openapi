/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// IMSAddress struct for IMSAddress
type IMSAddress struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IMSAddress) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(IMSAddress)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *IMSAddress) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIMSAddress struct {
	value *IMSAddress
	isSet bool
}

func (v NullableIMSAddress) Get() *IMSAddress {
	return v.value
}

func (v *NullableIMSAddress) Set(val *IMSAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIMSAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIMSAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIMSAddress(val *IMSAddress) *NullableIMSAddress {
	return &NullableIMSAddress{value: val, isSet: true}
}

func (v NullableIMSAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIMSAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



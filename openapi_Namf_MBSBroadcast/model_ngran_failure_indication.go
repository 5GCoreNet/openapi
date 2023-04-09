/*
Namf_MBSBroadcast

AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSBroadcast

import (
	"encoding/json"
	"fmt"
)

// NgranFailureIndication Indicates a NG-RAN failure event.
type NgranFailureIndication struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NgranFailureIndication) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(NgranFailureIndication)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NgranFailureIndication) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNgranFailureIndication struct {
	value *NgranFailureIndication
	isSet bool
}

func (v NullableNgranFailureIndication) Get() *NgranFailureIndication {
	return v.value
}

func (v *NullableNgranFailureIndication) Set(val *NgranFailureIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableNgranFailureIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableNgranFailureIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgranFailureIndication(val *NgranFailureIndication) *NullableNgranFailureIndication {
	return &NullableNgranFailureIndication{value: val, isSet: true}
}

func (v NullableNgranFailureIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgranFailureIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

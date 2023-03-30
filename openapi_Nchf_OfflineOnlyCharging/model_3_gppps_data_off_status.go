/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_OfflineOnlyCharging

import (
	"encoding/json"
	"fmt"
)

// Model3GPPPSDataOffStatus struct for Model3GPPPSDataOffStatus
type Model3GPPPSDataOffStatus struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Model3GPPPSDataOffStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Model3GPPPSDataOffStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Model3GPPPSDataOffStatus) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableModel3GPPPSDataOffStatus struct {
	value *Model3GPPPSDataOffStatus
	isSet bool
}

func (v NullableModel3GPPPSDataOffStatus) Get() *Model3GPPPSDataOffStatus {
	return v.value
}

func (v *NullableModel3GPPPSDataOffStatus) Set(val *Model3GPPPSDataOffStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableModel3GPPPSDataOffStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableModel3GPPPSDataOffStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel3GPPPSDataOffStatus(val *Model3GPPPSDataOffStatus) *NullableModel3GPPPSDataOffStatus {
	return &NullableModel3GPPPSDataOffStatus{value: val, isSet: true}
}

func (v NullableModel3GPPPSDataOffStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel3GPPPSDataOffStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


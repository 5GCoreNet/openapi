/*
Ntsctsf_TimeSynchronization Service API

TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_TimeSynchronization

import (
	"encoding/json"
	"fmt"
)

// AsTimeResource Possible values are: - ATOMIC_CLOCK: Indicates atomic clock is supported. - GNSS: Indicates Global Navigation Satellite System is supported. - TERRESTRIAL_RADIO: Indicates terrestrial radio is supported. - SERIAL_TIME_CODE: Indicates serial time code is supported. - PTP: Indicates PTP is supported. - NTP: Indicates NTP is supported. - HAND_SET: Indicates hand set is supported. - INTERNAL_OSCILLATOR: Indicates internal oscillator is supported. - OTHER: Indicates other source of time is supported. 
type AsTimeResource struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AsTimeResource) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AsTimeResource)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AsTimeResource) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAsTimeResource struct {
	value *AsTimeResource
	isSet bool
}

func (v NullableAsTimeResource) Get() *AsTimeResource {
	return v.value
}

func (v *NullableAsTimeResource) Set(val *AsTimeResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAsTimeResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAsTimeResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsTimeResource(val *AsTimeResource) *NullableAsTimeResource {
	return &NullableAsTimeResource{value: val, isSet: true}
}

func (v NullableAsTimeResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsTimeResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// DlDataDeliveryStatus Possible values are: - BUFFERED: The first downlink data is buffered with extended buffering matching the   source of the downlink traffic. - TRANSMITTED: The first downlink data matching the source of the downlink traffic is   transmitted after previous buffering or discarding of corresponding packet(s) because   the UE of the PDU Session becomes ACTIVE, and buffered data can be delivered to UE. - DISCARDED: The first downlink data matching the source of the downlink traffic is   discarded because the Extended Buffering time, as determined by the SMF, expires or   the amount of downlink data to be buffered is exceeded.
type DlDataDeliveryStatus struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DlDataDeliveryStatus) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(DlDataDeliveryStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DlDataDeliveryStatus) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDlDataDeliveryStatus struct {
	value *DlDataDeliveryStatus
	isSet bool
}

func (v NullableDlDataDeliveryStatus) Get() *DlDataDeliveryStatus {
	return v.value
}

func (v *NullableDlDataDeliveryStatus) Set(val *DlDataDeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDlDataDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDlDataDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDlDataDeliveryStatus(val *DlDataDeliveryStatus) *NullableDlDataDeliveryStatus {
	return &NullableDlDataDeliveryStatus{value: val, isSet: true}
}

func (v NullableDlDataDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDlDataDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

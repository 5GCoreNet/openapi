/*
3gpp-nidd

API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NIDD

import (
	"encoding/json"
	"fmt"
)

// DeliveryStatus Possible values are - SUCCESS: Success but details not provided - SUCCESS_NEXT_HOP_ACKNOWLEDGED: Successful delivery to the next hop with acknowledgment. - SUCCESS_NEXT_HOP_UNACKNOWLEDGED: Successful delivery to the next hop without acknowledgment - SUCCESS_ACKNOWLEDGED: Reliable delivery was acknowledged by the UE - SUCCESS_UNACKNOWLEDGED: Reliable delivery was not acknowledged by the UE - TRIGGERED: The SCEF triggered the device and is buffering the data. - BUFFERING: The SCEF is buffering the data due to no PDN connection established. - BUFFERING_TEMPORARILY_NOT_REACHABLE: The SCEF has been informed that the UE is temporarily not reachable but is buffering the data - SENDING: The SCEF has forwarded the data, but they may be stored elsewhere - FAILURE: Delivery failure but details not provided - FAILURE_RDS_DISABLED: RDS was disabled - FAILURE_NEXT_HOP: Unsuccessful delivery to the next hop. - FAILURE_TIMEOUT: Unsuccessful delivery due to timeout.  - FAILURE_TEMPORARILY_NOT_REACHABLE: The SCEF has been informed that the UE is temporarily not reachable without buffering the data.
type DeliveryStatus struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DeliveryStatus) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(DeliveryStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DeliveryStatus) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDeliveryStatus struct {
	value *DeliveryStatus
	isSet bool
}

func (v NullableDeliveryStatus) Get() *DeliveryStatus {
	return v.value
}

func (v *NullableDeliveryStatus) Set(val *DeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryStatus(val *DeliveryStatus) *NullableDeliveryStatus {
	return &NullableDeliveryStatus{value: val, isSet: true}
}

func (v NullableDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// NwdafEvent Possible values are: - SLICE_LOAD_LEVEL: Indicates that the event subscribed is load level information of Network Slice   - NETWORK_PERFORMANCE: Indicates that the event subscribed is network performance information.   - NF_LOAD: Indicates that the event subscribed is load level and status of one or several Network Functions.   - SERVICE_EXPERIENCE: Indicates that the event subscribed is service experience.   - UE_MOBILITY: Indicates that the event subscribed is UE mobility information.   - UE_COMMUNICATION: Indicates that the event subscribed is UE communication information.   - QOS_SUSTAINABILITY: Indicates that the event subscribed is QoS sustainability.   - ABNORMAL_BEHAVIOUR: Indicates that the event subscribed is abnormal behaviour.   - USER_DATA_CONGESTION: Indicates that the event subscribed is user data congestion information.   - NSI_LOAD_LEVEL: Indicates that the event subscribed is load level information of Network Slice and the optionally associated Network Slice Instance   - DN_PERFORMANCE: Indicates that the event subscribed is DN performance information.   - DISPERSION: Indicates that the event subscribed is dispersion information.   - RED_TRANS_EXP: Indicates that the event subscribed is redundant transmission experience.   - WLAN_PERFORMANCE: Indicates that the event subscribed is WLAN performance.   - SM_CONGESTION: Indicates the Session Management Congestion Control Experience information for specific DNN and/or S-NSSAI. 
type NwdafEvent struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NwdafEvent) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(NwdafEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NwdafEvent) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNwdafEvent struct {
	value *NwdafEvent
	isSet bool
}

func (v NullableNwdafEvent) Get() *NwdafEvent {
	return v.value
}

func (v *NullableNwdafEvent) Set(val *NwdafEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafEvent(val *NwdafEvent) *NullableNwdafEvent {
	return &NullableNwdafEvent{value: val, isSet: true}
}

func (v NullableNwdafEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



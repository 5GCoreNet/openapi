/*
Nmfaf_3daDataManagement

MFAF 3GPP DCCF Adaptor (3DA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3daDataManagement

import (
	"encoding/json"
	"fmt"
)

// NwdafEvent Possible values are: - SLICE_LOAD_LEVEL: Indicates that the event subscribed is load level information of Network Slice   - NETWORK_PERFORMANCE: Indicates that the event subscribed is network performance information.   - NF_LOAD: Indicates that the event subscribed is load level and status of one or several Network Functions.   - SERVICE_EXPERIENCE: Indicates that the event subscribed is service experience.   - UE_MOBILITY: Indicates that the event subscribed is UE mobility information.   - UE_COMMUNICATION: Indicates that the event subscribed is UE communication information.   - QOS_SUSTAINABILITY: Indicates that the event subscribed is QoS sustainability.   - ABNORMAL_BEHAVIOUR: Indicates that the event subscribed is abnormal behaviour.   - USER_DATA_CONGESTION: Indicates that the event subscribed is user data congestion information.   - NSI_LOAD_LEVEL: Indicates that the event subscribed is load level information of Network Slice and the optionally associated Network Slice Instance   - DN_PERFORMANCE: Indicates that the event subscribed is DN performance information.   - DISPERSION: Indicates that the event subscribed is dispersion information.   - RED_TRANS_EXP: Indicates that the event subscribed is redundant transmission experience.   - WLAN_PERFORMANCE: Indicates that the event subscribed is WLAN performance.   - SM_CONGESTION: Indicates the Session Management Congestion Control Experience information for specific DNN and/or S-NSSAI. 
type NwdafEvent struct {
	NwdafEventAnyOf *NwdafEventAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NwdafEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NwdafEventAnyOf
	err = json.Unmarshal(data, &dst.NwdafEventAnyOf);
	if err == nil {
		jsonNwdafEventAnyOf, _ := json.Marshal(dst.NwdafEventAnyOf)
		if string(jsonNwdafEventAnyOf) == "{}" { // empty struct
			dst.NwdafEventAnyOf = nil
		} else {
			return nil // data stored in dst.NwdafEventAnyOf, return on the first match
		}
	} else {
		dst.NwdafEventAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
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

	return fmt.Errorf("data failed to match schemas in anyOf(NwdafEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NwdafEvent) MarshalJSON() ([]byte, error) {
	if src.NwdafEventAnyOf != nil {
		return json.Marshal(&src.NwdafEventAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
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



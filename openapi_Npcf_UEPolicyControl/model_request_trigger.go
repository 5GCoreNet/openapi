/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_UEPolicyControl

import (
	"encoding/json"
	"fmt"
)

// RequestTrigger Possible values are: - LOC_CH: Location change (tracking area). The tracking area of the UE has changed. - PRA_CH: Change of UE presence in PRA. The AMF reports the current presence status of the UE   in a Presence Reporting Area, and notifies that the UE enters/leaves the Presence Reporting   Area. - UE_POLICY: A MANAGE UE POLICY COMPLETE message or a MANAGE UE POLICY COMMAND REJECT   message, as defined in Annex D.5 of 3GPP TS 24.501 or a \"UE POLICY PROVISIONING REQUEST\"   message, as defined in clause 7.2.1.1 of 3GPP TS 24.587 , has been received by the AMF   and is being forwarded. - PLMN_CH: PLMN change. the serving PLMN of UE has changed.  - CON_STATE_CH: Connectivity state change: the connectivity state of UE has changed.  - GROUP_ID_LIST_CHG: UE Internal Group Identifier(s) has changed. This policy control request   trigger does not require a subscription - UE_CAP_CH: UE Capabilities change: the UE provided 5G ProSe capabilities have changed.   This policy control request trigger does not require subscription.  
type RequestTrigger struct {
	RequestTriggerAnyOf *RequestTriggerAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RequestTrigger) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RequestTriggerAnyOf
	err = json.Unmarshal(data, &dst.RequestTriggerAnyOf);
	if err == nil {
		jsonRequestTriggerAnyOf, _ := json.Marshal(dst.RequestTriggerAnyOf)
		if string(jsonRequestTriggerAnyOf) == "{}" { // empty struct
			dst.RequestTriggerAnyOf = nil
		} else {
			return nil // data stored in dst.RequestTriggerAnyOf, return on the first match
		}
	} else {
		dst.RequestTriggerAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RequestTrigger)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RequestTrigger) MarshalJSON() ([]byte, error) {
	if src.RequestTriggerAnyOf != nil {
		return json.Marshal(&src.RequestTriggerAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRequestTrigger struct {
	value *RequestTrigger
	isSet bool
}

func (v NullableRequestTrigger) Get() *RequestTrigger {
	return v.value
}

func (v *NullableRequestTrigger) Set(val *RequestTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTrigger(val *RequestTrigger) *NullableRequestTrigger {
	return &NullableRequestTrigger{value: val, isSet: true}
}

func (v NullableRequestTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



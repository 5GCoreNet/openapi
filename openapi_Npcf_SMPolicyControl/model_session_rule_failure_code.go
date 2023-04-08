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

// SessionRuleFailureCode Possible values are - NF_MAL: Indicates that the PCC rule could not be successfully installed (for those  provisioned from the PCF) or activated (for those pre-defined in SMF) or enforced (for those  already successfully installed) due to SMF/UPF malfunction. - RES_LIM: Indicates that the PCC rule could not be successfully installed (for those  provisioned from PCF) or activated (for those pre-defined in SMF) or enforced (for those  already successfully installed) due to a limitation of resources at the SMF/UPF. - SESSION_RESOURCE_ALLOCATION_FAILURE: Indicates the session rule could not be successfully  enforced due to failure during the allocation of resources for the PDU session in the UE,  RAN or AMF. - UNSUCC_QOS_VAL: indicates that the QoS validation has failed. - INCORRECT_UM: The usage monitoring data of the enforced session rule is not the same for  all the provisioned session rule(s). - UE_STA_SUSP: Indicates that the UE is in suspend state. - UNKNOWN_REF_ID: Indicates that the session rule could not be successfully  installed/modified because the referenced identifier to a Policy Decision Data or to a  Condition Data is unknown to the SMF. - INCORRECT_COND_DATA: Indicates that the session rule could not be successfully  installed/modified because the referenced Condition data are incorrect. - REF_ID_COLLISION: Indicates that the session rule could not be successfully  installed/modified because the same Policy Decision is referenced by a PCC rule (e.g. the  session rule and the PCC rule refer to the same Usage Monitoring decision data). - AN_GW_FAILED: Indicates that the AN-Gateway has failed and that the PCF should refrain  from sending policy decisions to the SMF until it is informed that the S-GW has been  recovered. This value shall not be used if the SM Policy association modification procedure  is initiated for session rule removal only. 
type SessionRuleFailureCode struct {
	SessionRuleFailureCodeAnyOf *SessionRuleFailureCodeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SessionRuleFailureCode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SessionRuleFailureCodeAnyOf
	err = json.Unmarshal(data, &dst.SessionRuleFailureCodeAnyOf);
	if err == nil {
		jsonSessionRuleFailureCodeAnyOf, _ := json.Marshal(dst.SessionRuleFailureCodeAnyOf)
		if string(jsonSessionRuleFailureCodeAnyOf) == "{}" { // empty struct
			dst.SessionRuleFailureCodeAnyOf = nil
		} else {
			return nil // data stored in dst.SessionRuleFailureCodeAnyOf, return on the first match
		}
	} else {
		dst.SessionRuleFailureCodeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SessionRuleFailureCode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SessionRuleFailureCode) MarshalJSON() ([]byte, error) {
	if src.SessionRuleFailureCodeAnyOf != nil {
		return json.Marshal(&src.SessionRuleFailureCodeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSessionRuleFailureCode struct {
	value *SessionRuleFailureCode
	isSet bool
}

func (v NullableSessionRuleFailureCode) Get() *SessionRuleFailureCode {
	return v.value
}

func (v *NullableSessionRuleFailureCode) Set(val *SessionRuleFailureCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionRuleFailureCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionRuleFailureCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionRuleFailureCode(val *SessionRuleFailureCode) *NullableSessionRuleFailureCode {
	return &NullableSessionRuleFailureCode{value: val, isSet: true}
}

func (v NullableSessionRuleFailureCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionRuleFailureCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



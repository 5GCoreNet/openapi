/*
ECS EES Registration_API

API for EES Registration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eecs_EESRegistration

import (
	"encoding/json"
	"fmt"
)

// ACRScenario Possible values are: - EEC_INITIATED: Represents the EEC initiated ACR scenario. - EEC_EXECUTED_VIA_SOURCE_EES: Represents the EEC ACR scenario executed via the S-EES. - EEC_EXECUTED_VIA_TARGET_EES: Represents the EEC ACR scenario executed via the T-EES. - SOURCE_EAS_DECIDED: Represents the EEC ACR scenario where the S-EAS decides to perform ACR. - SOURCE_EES_EXECUTED: Represents the EEC ACR scenario where S-EES executes the ACR. - EEL_MANAGED_ACR: Represents the EEC ACR scenario where the ACR is managed by the Edge Enabler Layer. 
type ACRScenario struct {
	ACRScenarioAnyOf *ACRScenarioAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ACRScenario) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ACRScenarioAnyOf
	err = json.Unmarshal(data, &dst.ACRScenarioAnyOf);
	if err == nil {
		jsonACRScenarioAnyOf, _ := json.Marshal(dst.ACRScenarioAnyOf)
		if string(jsonACRScenarioAnyOf) == "{}" { // empty struct
			dst.ACRScenarioAnyOf = nil
		} else {
			return nil // data stored in dst.ACRScenarioAnyOf, return on the first match
		}
	} else {
		dst.ACRScenarioAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ACRScenario)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ACRScenario) MarshalJSON() ([]byte, error) {
	if src.ACRScenarioAnyOf != nil {
		return json.Marshal(&src.ACRScenarioAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableACRScenario struct {
	value *ACRScenario
	isSet bool
}

func (v NullableACRScenario) Get() *ACRScenario {
	return v.value
}

func (v *NullableACRScenario) Set(val *ACRScenario) {
	v.value = val
	v.isSet = true
}

func (v NullableACRScenario) IsSet() bool {
	return v.isSet
}

func (v *NullableACRScenario) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACRScenario(val *ACRScenario) *NullableACRScenario {
	return &NullableACRScenario{value: val, isSet: true}
}

func (v NullableACRScenario) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACRScenario) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



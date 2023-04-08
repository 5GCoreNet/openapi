/*
M1_PolicyTemplatesProvisioning

5GMS AF M1 Policy Templates Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_PolicyTemplatesProvisioning

import (
	"encoding/json"
	"fmt"
)

// PolicyTemplateApiType struct for PolicyTemplateApiType
type PolicyTemplateApiType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PolicyTemplateApiType) UnmarshalJSON(data []byte) error {
	var err error
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

	return fmt.Errorf("data failed to match schemas in anyOf(PolicyTemplateApiType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PolicyTemplateApiType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePolicyTemplateApiType struct {
	value *PolicyTemplateApiType
	isSet bool
}

func (v NullablePolicyTemplateApiType) Get() *PolicyTemplateApiType {
	return v.value
}

func (v *NullablePolicyTemplateApiType) Set(val *PolicyTemplateApiType) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyTemplateApiType) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyTemplateApiType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyTemplateApiType(val *PolicyTemplateApiType) *NullablePolicyTemplateApiType {
	return &NullablePolicyTemplateApiType{value: val, isSet: true}
}

func (v NullablePolicyTemplateApiType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyTemplateApiType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



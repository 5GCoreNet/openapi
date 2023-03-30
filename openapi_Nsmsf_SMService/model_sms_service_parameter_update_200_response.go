/*
Nsmsf_SMService Service API

SMSF SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmsf_SMService

import (
	"encoding/json"
	"fmt"
)

// SMSServiceParameterUpdate200Response - struct for SMSServiceParameterUpdate200Response
type SMSServiceParameterUpdate200Response struct {
	PatchResult *PatchResult
	UeSmsContextData *UeSmsContextData
}

// PatchResultAsSMSServiceParameterUpdate200Response is a convenience function that returns PatchResult wrapped in SMSServiceParameterUpdate200Response
func PatchResultAsSMSServiceParameterUpdate200Response(v *PatchResult) SMSServiceParameterUpdate200Response {
	return SMSServiceParameterUpdate200Response{
		PatchResult: v,
	}
}

// UeSmsContextDataAsSMSServiceParameterUpdate200Response is a convenience function that returns UeSmsContextData wrapped in SMSServiceParameterUpdate200Response
func UeSmsContextDataAsSMSServiceParameterUpdate200Response(v *UeSmsContextData) SMSServiceParameterUpdate200Response {
	return SMSServiceParameterUpdate200Response{
		UeSmsContextData: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SMSServiceParameterUpdate200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PatchResult
	err = newStrictDecoder(data).Decode(&dst.PatchResult)
	if err == nil {
		jsonPatchResult, _ := json.Marshal(dst.PatchResult)
		if string(jsonPatchResult) == "{}" { // empty struct
			dst.PatchResult = nil
		} else {
			match++
		}
	} else {
		dst.PatchResult = nil
	}

	// try to unmarshal data into UeSmsContextData
	err = newStrictDecoder(data).Decode(&dst.UeSmsContextData)
	if err == nil {
		jsonUeSmsContextData, _ := json.Marshal(dst.UeSmsContextData)
		if string(jsonUeSmsContextData) == "{}" { // empty struct
			dst.UeSmsContextData = nil
		} else {
			match++
		}
	} else {
		dst.UeSmsContextData = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PatchResult = nil
		dst.UeSmsContextData = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SMSServiceParameterUpdate200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SMSServiceParameterUpdate200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SMSServiceParameterUpdate200Response) MarshalJSON() ([]byte, error) {
	if src.PatchResult != nil {
		return json.Marshal(&src.PatchResult)
	}

	if src.UeSmsContextData != nil {
		return json.Marshal(&src.UeSmsContextData)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SMSServiceParameterUpdate200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PatchResult != nil {
		return obj.PatchResult
	}

	if obj.UeSmsContextData != nil {
		return obj.UeSmsContextData
	}

	// all schemas are nil
	return nil
}

type NullableSMSServiceParameterUpdate200Response struct {
	value *SMSServiceParameterUpdate200Response
	isSet bool
}

func (v NullableSMSServiceParameterUpdate200Response) Get() *SMSServiceParameterUpdate200Response {
	return v.value
}

func (v *NullableSMSServiceParameterUpdate200Response) Set(val *SMSServiceParameterUpdate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSMSServiceParameterUpdate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSMSServiceParameterUpdate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMSServiceParameterUpdate200Response(val *SMSServiceParameterUpdate200Response) *NullableSMSServiceParameterUpdate200Response {
	return &NullableSMSServiceParameterUpdate200Response{value: val, isSet: true}
}

func (v NullableSMSServiceParameterUpdate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMSServiceParameterUpdate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



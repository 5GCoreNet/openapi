/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// ChargingdataChargingDataRefUpdatePost400Response - struct for ChargingdataChargingDataRefUpdatePost400Response
type ChargingdataChargingDataRefUpdatePost400Response struct {
	ChargingDataResponse *ChargingDataResponse
	ProblemDetails       *ProblemDetails
}

// ChargingDataResponseAsChargingdataChargingDataRefUpdatePost400Response is a convenience function that returns ChargingDataResponse wrapped in ChargingdataChargingDataRefUpdatePost400Response
func ChargingDataResponseAsChargingdataChargingDataRefUpdatePost400Response(v *ChargingDataResponse) ChargingdataChargingDataRefUpdatePost400Response {
	return ChargingdataChargingDataRefUpdatePost400Response{
		ChargingDataResponse: v,
	}
}

// ProblemDetailsAsChargingdataChargingDataRefUpdatePost400Response is a convenience function that returns ProblemDetails wrapped in ChargingdataChargingDataRefUpdatePost400Response
func ProblemDetailsAsChargingdataChargingDataRefUpdatePost400Response(v *ProblemDetails) ChargingdataChargingDataRefUpdatePost400Response {
	return ChargingdataChargingDataRefUpdatePost400Response{
		ProblemDetails: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChargingdataChargingDataRefUpdatePost400Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChargingDataResponse
	err = newStrictDecoder(data).Decode(&dst.ChargingDataResponse)
	if err == nil {
		jsonChargingDataResponse, _ := json.Marshal(dst.ChargingDataResponse)
		if string(jsonChargingDataResponse) == "{}" { // empty struct
			dst.ChargingDataResponse = nil
		} else {
			match++
		}
	} else {
		dst.ChargingDataResponse = nil
	}

	// try to unmarshal data into ProblemDetails
	err = newStrictDecoder(data).Decode(&dst.ProblemDetails)
	if err == nil {
		jsonProblemDetails, _ := json.Marshal(dst.ProblemDetails)
		if string(jsonProblemDetails) == "{}" { // empty struct
			dst.ProblemDetails = nil
		} else {
			match++
		}
	} else {
		dst.ProblemDetails = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ChargingDataResponse = nil
		dst.ProblemDetails = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ChargingdataChargingDataRefUpdatePost400Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ChargingdataChargingDataRefUpdatePost400Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChargingdataChargingDataRefUpdatePost400Response) MarshalJSON() ([]byte, error) {
	if src.ChargingDataResponse != nil {
		return json.Marshal(&src.ChargingDataResponse)
	}

	if src.ProblemDetails != nil {
		return json.Marshal(&src.ProblemDetails)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ChargingdataChargingDataRefUpdatePost400Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ChargingDataResponse != nil {
		return obj.ChargingDataResponse
	}

	if obj.ProblemDetails != nil {
		return obj.ProblemDetails
	}

	// all schemas are nil
	return nil
}

type NullableChargingdataChargingDataRefUpdatePost400Response struct {
	value *ChargingdataChargingDataRefUpdatePost400Response
	isSet bool
}

func (v NullableChargingdataChargingDataRefUpdatePost400Response) Get() *ChargingdataChargingDataRefUpdatePost400Response {
	return v.value
}

func (v *NullableChargingdataChargingDataRefUpdatePost400Response) Set(val *ChargingdataChargingDataRefUpdatePost400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableChargingdataChargingDataRefUpdatePost400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableChargingdataChargingDataRefUpdatePost400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargingdataChargingDataRefUpdatePost400Response(val *ChargingdataChargingDataRefUpdatePost400Response) *NullableChargingdataChargingDataRefUpdatePost400Response {
	return &NullableChargingdataChargingDataRefUpdatePost400Response{value: val, isSet: true}
}

func (v NullableChargingdataChargingDataRefUpdatePost400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargingdataChargingDataRefUpdatePost400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

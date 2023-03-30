/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
	"fmt"
)

// AlarmsAlarmIdPatchRequest - struct for AlarmsAlarmIdPatchRequest
type AlarmsAlarmIdPatchRequest struct {
	MergePatchAcknowledgeAlarm *MergePatchAcknowledgeAlarm
	MergePatchClearAlarm *MergePatchClearAlarm
}

// MergePatchAcknowledgeAlarmAsAlarmsAlarmIdPatchRequest is a convenience function that returns MergePatchAcknowledgeAlarm wrapped in AlarmsAlarmIdPatchRequest
func MergePatchAcknowledgeAlarmAsAlarmsAlarmIdPatchRequest(v *MergePatchAcknowledgeAlarm) AlarmsAlarmIdPatchRequest {
	return AlarmsAlarmIdPatchRequest{
		MergePatchAcknowledgeAlarm: v,
	}
}

// MergePatchClearAlarmAsAlarmsAlarmIdPatchRequest is a convenience function that returns MergePatchClearAlarm wrapped in AlarmsAlarmIdPatchRequest
func MergePatchClearAlarmAsAlarmsAlarmIdPatchRequest(v *MergePatchClearAlarm) AlarmsAlarmIdPatchRequest {
	return AlarmsAlarmIdPatchRequest{
		MergePatchClearAlarm: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AlarmsAlarmIdPatchRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MergePatchAcknowledgeAlarm
	err = newStrictDecoder(data).Decode(&dst.MergePatchAcknowledgeAlarm)
	if err == nil {
		jsonMergePatchAcknowledgeAlarm, _ := json.Marshal(dst.MergePatchAcknowledgeAlarm)
		if string(jsonMergePatchAcknowledgeAlarm) == "{}" { // empty struct
			dst.MergePatchAcknowledgeAlarm = nil
		} else {
			match++
		}
	} else {
		dst.MergePatchAcknowledgeAlarm = nil
	}

	// try to unmarshal data into MergePatchClearAlarm
	err = newStrictDecoder(data).Decode(&dst.MergePatchClearAlarm)
	if err == nil {
		jsonMergePatchClearAlarm, _ := json.Marshal(dst.MergePatchClearAlarm)
		if string(jsonMergePatchClearAlarm) == "{}" { // empty struct
			dst.MergePatchClearAlarm = nil
		} else {
			match++
		}
	} else {
		dst.MergePatchClearAlarm = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MergePatchAcknowledgeAlarm = nil
		dst.MergePatchClearAlarm = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AlarmsAlarmIdPatchRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AlarmsAlarmIdPatchRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AlarmsAlarmIdPatchRequest) MarshalJSON() ([]byte, error) {
	if src.MergePatchAcknowledgeAlarm != nil {
		return json.Marshal(&src.MergePatchAcknowledgeAlarm)
	}

	if src.MergePatchClearAlarm != nil {
		return json.Marshal(&src.MergePatchClearAlarm)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AlarmsAlarmIdPatchRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MergePatchAcknowledgeAlarm != nil {
		return obj.MergePatchAcknowledgeAlarm
	}

	if obj.MergePatchClearAlarm != nil {
		return obj.MergePatchClearAlarm
	}

	// all schemas are nil
	return nil
}

type NullableAlarmsAlarmIdPatchRequest struct {
	value *AlarmsAlarmIdPatchRequest
	isSet bool
}

func (v NullableAlarmsAlarmIdPatchRequest) Get() *AlarmsAlarmIdPatchRequest {
	return v.value
}

func (v *NullableAlarmsAlarmIdPatchRequest) Set(val *AlarmsAlarmIdPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmsAlarmIdPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmsAlarmIdPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmsAlarmIdPatchRequest(val *AlarmsAlarmIdPatchRequest) *NullableAlarmsAlarmIdPatchRequest {
	return &NullableAlarmsAlarmIdPatchRequest{value: val, isSet: true}
}

func (v NullableAlarmsAlarmIdPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmsAlarmIdPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Common Type Definitions

OAS 3.0.1 specification of common type definitions in the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ComDefs

import (
	"encoding/json"
	"fmt"
)

// NotificationType - struct for NotificationType
type NotificationType struct {
	AlarmNotificationTypes     *AlarmNotificationTypes
	CmNotificationTypes        *CmNotificationTypes
	FileNotificationTypes      *FileNotificationTypes
	HeartbeatNotificationTypes *HeartbeatNotificationTypes
	PerfNotificationTypes      *PerfNotificationTypes
}

// AlarmNotificationTypesAsNotificationType is a convenience function that returns AlarmNotificationTypes wrapped in NotificationType
func AlarmNotificationTypesAsNotificationType(v *AlarmNotificationTypes) NotificationType {
	return NotificationType{
		AlarmNotificationTypes: v,
	}
}

// CmNotificationTypesAsNotificationType is a convenience function that returns CmNotificationTypes wrapped in NotificationType
func CmNotificationTypesAsNotificationType(v *CmNotificationTypes) NotificationType {
	return NotificationType{
		CmNotificationTypes: v,
	}
}

// FileNotificationTypesAsNotificationType is a convenience function that returns FileNotificationTypes wrapped in NotificationType
func FileNotificationTypesAsNotificationType(v *FileNotificationTypes) NotificationType {
	return NotificationType{
		FileNotificationTypes: v,
	}
}

// HeartbeatNotificationTypesAsNotificationType is a convenience function that returns HeartbeatNotificationTypes wrapped in NotificationType
func HeartbeatNotificationTypesAsNotificationType(v *HeartbeatNotificationTypes) NotificationType {
	return NotificationType{
		HeartbeatNotificationTypes: v,
	}
}

// PerfNotificationTypesAsNotificationType is a convenience function that returns PerfNotificationTypes wrapped in NotificationType
func PerfNotificationTypesAsNotificationType(v *PerfNotificationTypes) NotificationType {
	return NotificationType{
		PerfNotificationTypes: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NotificationType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AlarmNotificationTypes
	err = newStrictDecoder(data).Decode(&dst.AlarmNotificationTypes)
	if err == nil {
		jsonAlarmNotificationTypes, _ := json.Marshal(dst.AlarmNotificationTypes)
		if string(jsonAlarmNotificationTypes) == "{}" { // empty struct
			dst.AlarmNotificationTypes = nil
		} else {
			match++
		}
	} else {
		dst.AlarmNotificationTypes = nil
	}

	// try to unmarshal data into CmNotificationTypes
	err = newStrictDecoder(data).Decode(&dst.CmNotificationTypes)
	if err == nil {
		jsonCmNotificationTypes, _ := json.Marshal(dst.CmNotificationTypes)
		if string(jsonCmNotificationTypes) == "{}" { // empty struct
			dst.CmNotificationTypes = nil
		} else {
			match++
		}
	} else {
		dst.CmNotificationTypes = nil
	}

	// try to unmarshal data into FileNotificationTypes
	err = newStrictDecoder(data).Decode(&dst.FileNotificationTypes)
	if err == nil {
		jsonFileNotificationTypes, _ := json.Marshal(dst.FileNotificationTypes)
		if string(jsonFileNotificationTypes) == "{}" { // empty struct
			dst.FileNotificationTypes = nil
		} else {
			match++
		}
	} else {
		dst.FileNotificationTypes = nil
	}

	// try to unmarshal data into HeartbeatNotificationTypes
	err = newStrictDecoder(data).Decode(&dst.HeartbeatNotificationTypes)
	if err == nil {
		jsonHeartbeatNotificationTypes, _ := json.Marshal(dst.HeartbeatNotificationTypes)
		if string(jsonHeartbeatNotificationTypes) == "{}" { // empty struct
			dst.HeartbeatNotificationTypes = nil
		} else {
			match++
		}
	} else {
		dst.HeartbeatNotificationTypes = nil
	}

	// try to unmarshal data into PerfNotificationTypes
	err = newStrictDecoder(data).Decode(&dst.PerfNotificationTypes)
	if err == nil {
		jsonPerfNotificationTypes, _ := json.Marshal(dst.PerfNotificationTypes)
		if string(jsonPerfNotificationTypes) == "{}" { // empty struct
			dst.PerfNotificationTypes = nil
		} else {
			match++
		}
	} else {
		dst.PerfNotificationTypes = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AlarmNotificationTypes = nil
		dst.CmNotificationTypes = nil
		dst.FileNotificationTypes = nil
		dst.HeartbeatNotificationTypes = nil
		dst.PerfNotificationTypes = nil

		return fmt.Errorf("data matches more than one schema in oneOf(NotificationType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NotificationType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NotificationType) MarshalJSON() ([]byte, error) {
	if src.AlarmNotificationTypes != nil {
		return json.Marshal(&src.AlarmNotificationTypes)
	}

	if src.CmNotificationTypes != nil {
		return json.Marshal(&src.CmNotificationTypes)
	}

	if src.FileNotificationTypes != nil {
		return json.Marshal(&src.FileNotificationTypes)
	}

	if src.HeartbeatNotificationTypes != nil {
		return json.Marshal(&src.HeartbeatNotificationTypes)
	}

	if src.PerfNotificationTypes != nil {
		return json.Marshal(&src.PerfNotificationTypes)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NotificationType) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AlarmNotificationTypes != nil {
		return obj.AlarmNotificationTypes
	}

	if obj.CmNotificationTypes != nil {
		return obj.CmNotificationTypes
	}

	if obj.FileNotificationTypes != nil {
		return obj.FileNotificationTypes
	}

	if obj.HeartbeatNotificationTypes != nil {
		return obj.HeartbeatNotificationTypes
	}

	if obj.PerfNotificationTypes != nil {
		return obj.PerfNotificationTypes
	}

	// all schemas are nil
	return nil
}

type NullableNotificationType struct {
	value *NotificationType
	isSet bool
}

func (v NullableNotificationType) Get() *NotificationType {
	return v.value
}

func (v *NullableNotificationType) Set(val *NotificationType) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationType) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationType(val *NotificationType) *NullableNotificationType {
	return &NullableNotificationType{value: val, isSet: true}
}

func (v NullableNotificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

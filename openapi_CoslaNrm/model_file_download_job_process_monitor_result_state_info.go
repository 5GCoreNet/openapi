/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
	"fmt"
)

// FileDownloadJobProcessMonitorResultStateInfo - struct for FileDownloadJobProcessMonitorResultStateInfo
type FileDownloadJobProcessMonitorResultStateInfo struct {
	String *string
}

// stringAsFileDownloadJobProcessMonitorResultStateInfo is a convenience function that returns string wrapped in FileDownloadJobProcessMonitorResultStateInfo
func StringAsFileDownloadJobProcessMonitorResultStateInfo(v *string) FileDownloadJobProcessMonitorResultStateInfo {
	return FileDownloadJobProcessMonitorResultStateInfo{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FileDownloadJobProcessMonitorResultStateInfo) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(FileDownloadJobProcessMonitorResultStateInfo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(FileDownloadJobProcessMonitorResultStateInfo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FileDownloadJobProcessMonitorResultStateInfo) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FileDownloadJobProcessMonitorResultStateInfo) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableFileDownloadJobProcessMonitorResultStateInfo struct {
	value *FileDownloadJobProcessMonitorResultStateInfo
	isSet bool
}

func (v NullableFileDownloadJobProcessMonitorResultStateInfo) Get() *FileDownloadJobProcessMonitorResultStateInfo {
	return v.value
}

func (v *NullableFileDownloadJobProcessMonitorResultStateInfo) Set(val *FileDownloadJobProcessMonitorResultStateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDownloadJobProcessMonitorResultStateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDownloadJobProcessMonitorResultStateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDownloadJobProcessMonitorResultStateInfo(val *FileDownloadJobProcessMonitorResultStateInfo) *NullableFileDownloadJobProcessMonitorResultStateInfo {
	return &NullableFileDownloadJobProcessMonitorResultStateInfo{value: val, isSet: true}
}

func (v NullableFileDownloadJobProcessMonitorResultStateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDownloadJobProcessMonitorResultStateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

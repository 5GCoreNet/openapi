/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// PatchOperation the model 'PatchOperation'
type PatchOperation string

// List of PatchOperation
const (
	ADD     PatchOperation = "add"
	REPLACE PatchOperation = "replace"
	REMOVE  PatchOperation = "remove"
	COPY    PatchOperation = "copy"
	MOVE    PatchOperation = "move"
	TEST    PatchOperation = "test"
)

// All allowed values of PatchOperation enum
var AllowedPatchOperationEnumValues = []PatchOperation{
	"add",
	"replace",
	"remove",
	"copy",
	"move",
	"test",
}

func (v *PatchOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchOperation(value)
	for _, existing := range AllowedPatchOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchOperation", value)
}

// NewPatchOperationFromValue returns a pointer to a valid PatchOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchOperationFromValue(v string) (*PatchOperation, error) {
	ev := PatchOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchOperation: valid values are %v", v, AllowedPatchOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchOperation) IsValid() bool {
	for _, existing := range AllowedPatchOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchOperation value
func (v PatchOperation) Ptr() *PatchOperation {
	return &v
}

type NullablePatchOperation struct {
	value *PatchOperation
	isSet bool
}

func (v NullablePatchOperation) Get() *PatchOperation {
	return v.value
}

func (v *NullablePatchOperation) Set(val *PatchOperation) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchOperation) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchOperation(val *PatchOperation) *NullablePatchOperation {
	return &NullablePatchOperation{value: val, isSet: true}
}

func (v NullablePatchOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

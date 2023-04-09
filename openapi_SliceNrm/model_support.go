/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
	"fmt"
)

// Support the model 'Support'
type Support string

// List of Support
const (
	NOT_SUPPORTED Support = "NOT SUPPORTED"
	SUPPORTED     Support = "SUPPORTED"
)

// All allowed values of Support enum
var AllowedSupportEnumValues = []Support{
	"NOT SUPPORTED",
	"SUPPORTED",
}

func (v *Support) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Support(value)
	for _, existing := range AllowedSupportEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Support", value)
}

// NewSupportFromValue returns a pointer to a valid Support
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupportFromValue(v string) (*Support, error) {
	ev := Support(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Support: valid values are %v", v, AllowedSupportEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Support) IsValid() bool {
	for _, existing := range AllowedSupportEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Support value
func (v Support) Ptr() *Support {
	return &v
}

type NullableSupport struct {
	value *Support
	isSet bool
}

func (v NullableSupport) Get() *Support {
	return v.value
}

func (v *NullableSupport) Set(val *Support) {
	v.value = val
	v.isSet = true
}

func (v NullableSupport) IsSet() bool {
	return v.isSet
}

func (v *NullableSupport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupport(val *Support) *NullableSupport {
	return &NullableSupport{value: val, isSet: true}
}

func (v NullableSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

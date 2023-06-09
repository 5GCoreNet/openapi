/*
JOSE Protected Message Forwarding API

N32-f Message Forwarding Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_JOSEProtectedMessageForwarding

import (
	"encoding/json"
	"fmt"
)

// EncodedHttpHeaderValue - HTTP header value or index to the HTTP header value
type EncodedHttpHeaderValue struct {
	IndexToEncryptedValue *IndexToEncryptedValue
	String                *string
}

// IndexToEncryptedValueAsEncodedHttpHeaderValue is a convenience function that returns IndexToEncryptedValue wrapped in EncodedHttpHeaderValue
func IndexToEncryptedValueAsEncodedHttpHeaderValue(v *IndexToEncryptedValue) EncodedHttpHeaderValue {
	return EncodedHttpHeaderValue{
		IndexToEncryptedValue: v,
	}
}

// stringAsEncodedHttpHeaderValue is a convenience function that returns string wrapped in EncodedHttpHeaderValue
func StringAsEncodedHttpHeaderValue(v *string) EncodedHttpHeaderValue {
	return EncodedHttpHeaderValue{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EncodedHttpHeaderValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IndexToEncryptedValue
	err = newStrictDecoder(data).Decode(&dst.IndexToEncryptedValue)
	if err == nil {
		jsonIndexToEncryptedValue, _ := json.Marshal(dst.IndexToEncryptedValue)
		if string(jsonIndexToEncryptedValue) == "{}" { // empty struct
			dst.IndexToEncryptedValue = nil
		} else {
			match++
		}
	} else {
		dst.IndexToEncryptedValue = nil
	}

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
		dst.IndexToEncryptedValue = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EncodedHttpHeaderValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EncodedHttpHeaderValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EncodedHttpHeaderValue) MarshalJSON() ([]byte, error) {
	if src.IndexToEncryptedValue != nil {
		return json.Marshal(&src.IndexToEncryptedValue)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EncodedHttpHeaderValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IndexToEncryptedValue != nil {
		return obj.IndexToEncryptedValue
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableEncodedHttpHeaderValue struct {
	value *EncodedHttpHeaderValue
	isSet bool
}

func (v NullableEncodedHttpHeaderValue) Get() *EncodedHttpHeaderValue {
	return v.value
}

func (v *NullableEncodedHttpHeaderValue) Set(val *EncodedHttpHeaderValue) {
	v.value = val
	v.isSet = true
}

func (v NullableEncodedHttpHeaderValue) IsSet() bool {
	return v.isSet
}

func (v *NullableEncodedHttpHeaderValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncodedHttpHeaderValue(val *EncodedHttpHeaderValue) *NullableEncodedHttpHeaderValue {
	return &NullableEncodedHttpHeaderValue{value: val, isSet: true}
}

func (v NullableEncodedHttpHeaderValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncodedHttpHeaderValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

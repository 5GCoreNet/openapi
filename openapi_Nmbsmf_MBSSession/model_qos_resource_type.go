/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
	"fmt"
)

// QosResourceType The enumeration QosResourceType indicates whether a QoS Flow is non-GBR, delay critical GBR, or non-delay critical GBR (see clauses 5.7.3.4 and 5.7.3.5 of 3GPP TS 23.501). It shall comply with the provisions defined in table 5.5.3.6-1.  
type QosResourceType struct {
	QosResourceTypeAnyOf *QosResourceTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *QosResourceType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into QosResourceTypeAnyOf
	err = json.Unmarshal(data, &dst.QosResourceTypeAnyOf);
	if err == nil {
		jsonQosResourceTypeAnyOf, _ := json.Marshal(dst.QosResourceTypeAnyOf)
		if string(jsonQosResourceTypeAnyOf) == "{}" { // empty struct
			dst.QosResourceTypeAnyOf = nil
		} else {
			return nil // data stored in dst.QosResourceTypeAnyOf, return on the first match
		}
	} else {
		dst.QosResourceTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(QosResourceType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *QosResourceType) MarshalJSON() ([]byte, error) {
	if src.QosResourceTypeAnyOf != nil {
		return json.Marshal(&src.QosResourceTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableQosResourceType struct {
	value *QosResourceType
	isSet bool
}

func (v NullableQosResourceType) Get() *QosResourceType {
	return v.value
}

func (v *NullableQosResourceType) Set(val *QosResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableQosResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableQosResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosResourceType(val *QosResourceType) *NullableQosResourceType {
	return &NullableQosResourceType{value: val, isSet: true}
}

func (v NullableQosResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



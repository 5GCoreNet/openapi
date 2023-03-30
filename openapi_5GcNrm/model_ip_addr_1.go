/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
	"fmt"
)

// IpAddr1 - struct for IpAddr1
type IpAddr1 struct {
	Ipv6Addr *Ipv6Addr
	String *string
}

// Ipv6AddrAsIpAddr1 is a convenience function that returns Ipv6Addr wrapped in IpAddr1
func Ipv6AddrAsIpAddr1(v *Ipv6Addr) IpAddr1 {
	return IpAddr1{
		Ipv6Addr: v,
	}
}

// stringAsIpAddr1 is a convenience function that returns string wrapped in IpAddr1
func StringAsIpAddr1(v *string) IpAddr1 {
	return IpAddr1{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IpAddr1) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Ipv6Addr
	err = newStrictDecoder(data).Decode(&dst.Ipv6Addr)
	if err == nil {
		jsonIpv6Addr, _ := json.Marshal(dst.Ipv6Addr)
		if string(jsonIpv6Addr) == "{}" { // empty struct
			dst.Ipv6Addr = nil
		} else {
			match++
		}
	} else {
		dst.Ipv6Addr = nil
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
		dst.Ipv6Addr = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IpAddr1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IpAddr1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IpAddr1) MarshalJSON() ([]byte, error) {
	if src.Ipv6Addr != nil {
		return json.Marshal(&src.Ipv6Addr)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IpAddr1) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Ipv6Addr != nil {
		return obj.Ipv6Addr
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableIpAddr1 struct {
	value *IpAddr1
	isSet bool
}

func (v NullableIpAddr1) Get() *IpAddr1 {
	return v.value
}

func (v *NullableIpAddr1) Set(val *IpAddr1) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAddr1) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAddr1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAddr1(val *IpAddr1) *NullableIpAddr1 {
	return &NullableIpAddr1{value: val, isSet: true}
}

func (v NullableIpAddr1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAddr1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



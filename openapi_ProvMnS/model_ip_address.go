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

// IpAddress - struct for IpAddress
type IpAddress struct {
	Ipv6Addr *Ipv6Addr
	String *string
}

// Ipv6AddrAsIpAddress is a convenience function that returns Ipv6Addr wrapped in IpAddress
func Ipv6AddrAsIpAddress(v *Ipv6Addr) IpAddress {
	return IpAddress{
		Ipv6Addr: v,
	}
}

// stringAsIpAddress is a convenience function that returns string wrapped in IpAddress
func StringAsIpAddress(v *string) IpAddress {
	return IpAddress{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IpAddress) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(IpAddress)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IpAddress)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IpAddress) MarshalJSON() ([]byte, error) {
	if src.Ipv6Addr != nil {
		return json.Marshal(&src.Ipv6Addr)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IpAddress) GetActualInstance() (interface{}) {
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

type NullableIpAddress struct {
	value *IpAddress
	isSet bool
}

func (v NullableIpAddress) Get() *IpAddress {
	return v.value
}

func (v *NullableIpAddress) Set(val *IpAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAddress(val *IpAddress) *NullableIpAddress {
	return &NullableIpAddress{value: val, isSet: true}
}

func (v NullableIpAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


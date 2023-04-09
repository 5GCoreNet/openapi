/*
M1_ContentHostingProvisioning

5GMS AF M1 Content Hosting Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_ContentHostingProvisioning

import (
	"encoding/json"
	"fmt"
)

// DistributionNetworkType Type of distribution network.
type DistributionNetworkType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DistributionNetworkType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
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

	return fmt.Errorf("data failed to match schemas in anyOf(DistributionNetworkType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DistributionNetworkType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDistributionNetworkType struct {
	value *DistributionNetworkType
	isSet bool
}

func (v NullableDistributionNetworkType) Get() *DistributionNetworkType {
	return v.value
}

func (v *NullableDistributionNetworkType) Set(val *DistributionNetworkType) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionNetworkType) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionNetworkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionNetworkType(val *DistributionNetworkType) *NullableDistributionNetworkType {
	return &NullableDistributionNetworkType{value: val, isSet: true}
}

func (v NullableDistributionNetworkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionNetworkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

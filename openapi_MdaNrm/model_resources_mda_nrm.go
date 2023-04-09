/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
	"fmt"
)

// ResourcesMdaNrm - struct for ResourcesMdaNrm
type ResourcesMdaNrm struct {
	MDAFunctionSingle    *MDAFunctionSingle
	MDAReportSingle      *MDAReportSingle
	MDARequestSingle     *MDARequestSingle
	ManagedElementSingle *ManagedElementSingle
	SubNetworkSingle     *SubNetworkSingle
}

// MDAFunctionSingleAsResourcesMdaNrm is a convenience function that returns MDAFunctionSingle wrapped in ResourcesMdaNrm
func MDAFunctionSingleAsResourcesMdaNrm(v *MDAFunctionSingle) ResourcesMdaNrm {
	return ResourcesMdaNrm{
		MDAFunctionSingle: v,
	}
}

// MDAReportSingleAsResourcesMdaNrm is a convenience function that returns MDAReportSingle wrapped in ResourcesMdaNrm
func MDAReportSingleAsResourcesMdaNrm(v *MDAReportSingle) ResourcesMdaNrm {
	return ResourcesMdaNrm{
		MDAReportSingle: v,
	}
}

// MDARequestSingleAsResourcesMdaNrm is a convenience function that returns MDARequestSingle wrapped in ResourcesMdaNrm
func MDARequestSingleAsResourcesMdaNrm(v *MDARequestSingle) ResourcesMdaNrm {
	return ResourcesMdaNrm{
		MDARequestSingle: v,
	}
}

// ManagedElementSingleAsResourcesMdaNrm is a convenience function that returns ManagedElementSingle wrapped in ResourcesMdaNrm
func ManagedElementSingleAsResourcesMdaNrm(v *ManagedElementSingle) ResourcesMdaNrm {
	return ResourcesMdaNrm{
		ManagedElementSingle: v,
	}
}

// SubNetworkSingleAsResourcesMdaNrm is a convenience function that returns SubNetworkSingle wrapped in ResourcesMdaNrm
func SubNetworkSingleAsResourcesMdaNrm(v *SubNetworkSingle) ResourcesMdaNrm {
	return ResourcesMdaNrm{
		SubNetworkSingle: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ResourcesMdaNrm) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MDAFunctionSingle
	err = newStrictDecoder(data).Decode(&dst.MDAFunctionSingle)
	if err == nil {
		jsonMDAFunctionSingle, _ := json.Marshal(dst.MDAFunctionSingle)
		if string(jsonMDAFunctionSingle) == "{}" { // empty struct
			dst.MDAFunctionSingle = nil
		} else {
			match++
		}
	} else {
		dst.MDAFunctionSingle = nil
	}

	// try to unmarshal data into MDAReportSingle
	err = newStrictDecoder(data).Decode(&dst.MDAReportSingle)
	if err == nil {
		jsonMDAReportSingle, _ := json.Marshal(dst.MDAReportSingle)
		if string(jsonMDAReportSingle) == "{}" { // empty struct
			dst.MDAReportSingle = nil
		} else {
			match++
		}
	} else {
		dst.MDAReportSingle = nil
	}

	// try to unmarshal data into MDARequestSingle
	err = newStrictDecoder(data).Decode(&dst.MDARequestSingle)
	if err == nil {
		jsonMDARequestSingle, _ := json.Marshal(dst.MDARequestSingle)
		if string(jsonMDARequestSingle) == "{}" { // empty struct
			dst.MDARequestSingle = nil
		} else {
			match++
		}
	} else {
		dst.MDARequestSingle = nil
	}

	// try to unmarshal data into ManagedElementSingle
	err = newStrictDecoder(data).Decode(&dst.ManagedElementSingle)
	if err == nil {
		jsonManagedElementSingle, _ := json.Marshal(dst.ManagedElementSingle)
		if string(jsonManagedElementSingle) == "{}" { // empty struct
			dst.ManagedElementSingle = nil
		} else {
			match++
		}
	} else {
		dst.ManagedElementSingle = nil
	}

	// try to unmarshal data into SubNetworkSingle
	err = newStrictDecoder(data).Decode(&dst.SubNetworkSingle)
	if err == nil {
		jsonSubNetworkSingle, _ := json.Marshal(dst.SubNetworkSingle)
		if string(jsonSubNetworkSingle) == "{}" { // empty struct
			dst.SubNetworkSingle = nil
		} else {
			match++
		}
	} else {
		dst.SubNetworkSingle = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MDAFunctionSingle = nil
		dst.MDAReportSingle = nil
		dst.MDARequestSingle = nil
		dst.ManagedElementSingle = nil
		dst.SubNetworkSingle = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ResourcesMdaNrm)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ResourcesMdaNrm)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ResourcesMdaNrm) MarshalJSON() ([]byte, error) {
	if src.MDAFunctionSingle != nil {
		return json.Marshal(&src.MDAFunctionSingle)
	}

	if src.MDAReportSingle != nil {
		return json.Marshal(&src.MDAReportSingle)
	}

	if src.MDARequestSingle != nil {
		return json.Marshal(&src.MDARequestSingle)
	}

	if src.ManagedElementSingle != nil {
		return json.Marshal(&src.ManagedElementSingle)
	}

	if src.SubNetworkSingle != nil {
		return json.Marshal(&src.SubNetworkSingle)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ResourcesMdaNrm) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MDAFunctionSingle != nil {
		return obj.MDAFunctionSingle
	}

	if obj.MDAReportSingle != nil {
		return obj.MDAReportSingle
	}

	if obj.MDARequestSingle != nil {
		return obj.MDARequestSingle
	}

	if obj.ManagedElementSingle != nil {
		return obj.ManagedElementSingle
	}

	if obj.SubNetworkSingle != nil {
		return obj.SubNetworkSingle
	}

	// all schemas are nil
	return nil
}

type NullableResourcesMdaNrm struct {
	value *ResourcesMdaNrm
	isSet bool
}

func (v NullableResourcesMdaNrm) Get() *ResourcesMdaNrm {
	return v.value
}

func (v *NullableResourcesMdaNrm) Set(val *ResourcesMdaNrm) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcesMdaNrm) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcesMdaNrm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcesMdaNrm(val *ResourcesMdaNrm) *NullableResourcesMdaNrm {
	return &NullableResourcesMdaNrm{value: val, isSet: true}
}

func (v NullableResourcesMdaNrm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcesMdaNrm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

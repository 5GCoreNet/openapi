/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
	"fmt"
)

// ResourcesEdgeNrm - struct for ResourcesEdgeNrm
type ResourcesEdgeNrm struct {
	EASFunctionSingle *EASFunctionSingle
	EASRequirements *EASRequirements
	ECSFunctionSingle *ECSFunctionSingle
	EESFunctionSingle *EESFunctionSingle
	EdgeDataNetworkSingle *EdgeDataNetworkSingle
	MnS *MnS
	SubNetworkSingle *SubNetworkSingle
}

// EASFunctionSingleAsResourcesEdgeNrm is a convenience function that returns EASFunctionSingle wrapped in ResourcesEdgeNrm
func EASFunctionSingleAsResourcesEdgeNrm(v *EASFunctionSingle) ResourcesEdgeNrm {
	return ResourcesEdgeNrm{
		EASFunctionSingle: v,
	}
}

// EASRequirementsAsResourcesEdgeNrm is a convenience function that returns EASRequirements wrapped in ResourcesEdgeNrm
func EASRequirementsAsResourcesEdgeNrm(v *EASRequirements) ResourcesEdgeNrm {
	return ResourcesEdgeNrm{
		EASRequirements: v,
	}
}

// ECSFunctionSingleAsResourcesEdgeNrm is a convenience function that returns ECSFunctionSingle wrapped in ResourcesEdgeNrm
func ECSFunctionSingleAsResourcesEdgeNrm(v *ECSFunctionSingle) ResourcesEdgeNrm {
	return ResourcesEdgeNrm{
		ECSFunctionSingle: v,
	}
}

// EESFunctionSingleAsResourcesEdgeNrm is a convenience function that returns EESFunctionSingle wrapped in ResourcesEdgeNrm
func EESFunctionSingleAsResourcesEdgeNrm(v *EESFunctionSingle) ResourcesEdgeNrm {
	return ResourcesEdgeNrm{
		EESFunctionSingle: v,
	}
}

// EdgeDataNetworkSingleAsResourcesEdgeNrm is a convenience function that returns EdgeDataNetworkSingle wrapped in ResourcesEdgeNrm
func EdgeDataNetworkSingleAsResourcesEdgeNrm(v *EdgeDataNetworkSingle) ResourcesEdgeNrm {
	return ResourcesEdgeNrm{
		EdgeDataNetworkSingle: v,
	}
}

// MnSAsResourcesEdgeNrm is a convenience function that returns MnS wrapped in ResourcesEdgeNrm
func MnSAsResourcesEdgeNrm(v *MnS) ResourcesEdgeNrm {
	return ResourcesEdgeNrm{
		MnS: v,
	}
}

// SubNetworkSingleAsResourcesEdgeNrm is a convenience function that returns SubNetworkSingle wrapped in ResourcesEdgeNrm
func SubNetworkSingleAsResourcesEdgeNrm(v *SubNetworkSingle) ResourcesEdgeNrm {
	return ResourcesEdgeNrm{
		SubNetworkSingle: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ResourcesEdgeNrm) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EASFunctionSingle
	err = newStrictDecoder(data).Decode(&dst.EASFunctionSingle)
	if err == nil {
		jsonEASFunctionSingle, _ := json.Marshal(dst.EASFunctionSingle)
		if string(jsonEASFunctionSingle) == "{}" { // empty struct
			dst.EASFunctionSingle = nil
		} else {
			match++
		}
	} else {
		dst.EASFunctionSingle = nil
	}

	// try to unmarshal data into EASRequirements
	err = newStrictDecoder(data).Decode(&dst.EASRequirements)
	if err == nil {
		jsonEASRequirements, _ := json.Marshal(dst.EASRequirements)
		if string(jsonEASRequirements) == "{}" { // empty struct
			dst.EASRequirements = nil
		} else {
			match++
		}
	} else {
		dst.EASRequirements = nil
	}

	// try to unmarshal data into ECSFunctionSingle
	err = newStrictDecoder(data).Decode(&dst.ECSFunctionSingle)
	if err == nil {
		jsonECSFunctionSingle, _ := json.Marshal(dst.ECSFunctionSingle)
		if string(jsonECSFunctionSingle) == "{}" { // empty struct
			dst.ECSFunctionSingle = nil
		} else {
			match++
		}
	} else {
		dst.ECSFunctionSingle = nil
	}

	// try to unmarshal data into EESFunctionSingle
	err = newStrictDecoder(data).Decode(&dst.EESFunctionSingle)
	if err == nil {
		jsonEESFunctionSingle, _ := json.Marshal(dst.EESFunctionSingle)
		if string(jsonEESFunctionSingle) == "{}" { // empty struct
			dst.EESFunctionSingle = nil
		} else {
			match++
		}
	} else {
		dst.EESFunctionSingle = nil
	}

	// try to unmarshal data into EdgeDataNetworkSingle
	err = newStrictDecoder(data).Decode(&dst.EdgeDataNetworkSingle)
	if err == nil {
		jsonEdgeDataNetworkSingle, _ := json.Marshal(dst.EdgeDataNetworkSingle)
		if string(jsonEdgeDataNetworkSingle) == "{}" { // empty struct
			dst.EdgeDataNetworkSingle = nil
		} else {
			match++
		}
	} else {
		dst.EdgeDataNetworkSingle = nil
	}

	// try to unmarshal data into MnS
	err = newStrictDecoder(data).Decode(&dst.MnS)
	if err == nil {
		jsonMnS, _ := json.Marshal(dst.MnS)
		if string(jsonMnS) == "{}" { // empty struct
			dst.MnS = nil
		} else {
			match++
		}
	} else {
		dst.MnS = nil
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
		dst.EASFunctionSingle = nil
		dst.EASRequirements = nil
		dst.ECSFunctionSingle = nil
		dst.EESFunctionSingle = nil
		dst.EdgeDataNetworkSingle = nil
		dst.MnS = nil
		dst.SubNetworkSingle = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ResourcesEdgeNrm)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ResourcesEdgeNrm)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ResourcesEdgeNrm) MarshalJSON() ([]byte, error) {
	if src.EASFunctionSingle != nil {
		return json.Marshal(&src.EASFunctionSingle)
	}

	if src.EASRequirements != nil {
		return json.Marshal(&src.EASRequirements)
	}

	if src.ECSFunctionSingle != nil {
		return json.Marshal(&src.ECSFunctionSingle)
	}

	if src.EESFunctionSingle != nil {
		return json.Marshal(&src.EESFunctionSingle)
	}

	if src.EdgeDataNetworkSingle != nil {
		return json.Marshal(&src.EdgeDataNetworkSingle)
	}

	if src.MnS != nil {
		return json.Marshal(&src.MnS)
	}

	if src.SubNetworkSingle != nil {
		return json.Marshal(&src.SubNetworkSingle)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ResourcesEdgeNrm) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EASFunctionSingle != nil {
		return obj.EASFunctionSingle
	}

	if obj.EASRequirements != nil {
		return obj.EASRequirements
	}

	if obj.ECSFunctionSingle != nil {
		return obj.ECSFunctionSingle
	}

	if obj.EESFunctionSingle != nil {
		return obj.EESFunctionSingle
	}

	if obj.EdgeDataNetworkSingle != nil {
		return obj.EdgeDataNetworkSingle
	}

	if obj.MnS != nil {
		return obj.MnS
	}

	if obj.SubNetworkSingle != nil {
		return obj.SubNetworkSingle
	}

	// all schemas are nil
	return nil
}

type NullableResourcesEdgeNrm struct {
	value *ResourcesEdgeNrm
	isSet bool
}

func (v NullableResourcesEdgeNrm) Get() *ResourcesEdgeNrm {
	return v.value
}

func (v *NullableResourcesEdgeNrm) Set(val *ResourcesEdgeNrm) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcesEdgeNrm) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcesEdgeNrm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcesEdgeNrm(val *ResourcesEdgeNrm) *NullableResourcesEdgeNrm {
	return &NullableResourcesEdgeNrm{value: val, isSet: true}
}

func (v NullableResourcesEdgeNrm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcesEdgeNrm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


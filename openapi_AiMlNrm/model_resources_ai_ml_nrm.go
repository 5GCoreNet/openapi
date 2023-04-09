/*
AI/ML NRM

OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AiMlNrm

import (
	"encoding/json"
	"fmt"
)

// ResourcesAiMlNrm - struct for ResourcesAiMlNrm
type ResourcesAiMlNrm struct {
	MLTrainingFunctionSingle *MLTrainingFunctionSingle
	MLTrainingProcessSingle  *MLTrainingProcessSingle
	MLTrainingReportSingle   *MLTrainingReportSingle
	MLTrainingRequestSingle  *MLTrainingRequestSingle
	ManagedElementSingle     *ManagedElementSingle
	SubNetworkSingle         *SubNetworkSingle
}

// MLTrainingFunctionSingleAsResourcesAiMlNrm is a convenience function that returns MLTrainingFunctionSingle wrapped in ResourcesAiMlNrm
func MLTrainingFunctionSingleAsResourcesAiMlNrm(v *MLTrainingFunctionSingle) ResourcesAiMlNrm {
	return ResourcesAiMlNrm{
		MLTrainingFunctionSingle: v,
	}
}

// MLTrainingProcessSingleAsResourcesAiMlNrm is a convenience function that returns MLTrainingProcessSingle wrapped in ResourcesAiMlNrm
func MLTrainingProcessSingleAsResourcesAiMlNrm(v *MLTrainingProcessSingle) ResourcesAiMlNrm {
	return ResourcesAiMlNrm{
		MLTrainingProcessSingle: v,
	}
}

// MLTrainingReportSingleAsResourcesAiMlNrm is a convenience function that returns MLTrainingReportSingle wrapped in ResourcesAiMlNrm
func MLTrainingReportSingleAsResourcesAiMlNrm(v *MLTrainingReportSingle) ResourcesAiMlNrm {
	return ResourcesAiMlNrm{
		MLTrainingReportSingle: v,
	}
}

// MLTrainingRequestSingleAsResourcesAiMlNrm is a convenience function that returns MLTrainingRequestSingle wrapped in ResourcesAiMlNrm
func MLTrainingRequestSingleAsResourcesAiMlNrm(v *MLTrainingRequestSingle) ResourcesAiMlNrm {
	return ResourcesAiMlNrm{
		MLTrainingRequestSingle: v,
	}
}

// ManagedElementSingleAsResourcesAiMlNrm is a convenience function that returns ManagedElementSingle wrapped in ResourcesAiMlNrm
func ManagedElementSingleAsResourcesAiMlNrm(v *ManagedElementSingle) ResourcesAiMlNrm {
	return ResourcesAiMlNrm{
		ManagedElementSingle: v,
	}
}

// SubNetworkSingleAsResourcesAiMlNrm is a convenience function that returns SubNetworkSingle wrapped in ResourcesAiMlNrm
func SubNetworkSingleAsResourcesAiMlNrm(v *SubNetworkSingle) ResourcesAiMlNrm {
	return ResourcesAiMlNrm{
		SubNetworkSingle: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ResourcesAiMlNrm) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MLTrainingFunctionSingle
	err = newStrictDecoder(data).Decode(&dst.MLTrainingFunctionSingle)
	if err == nil {
		jsonMLTrainingFunctionSingle, _ := json.Marshal(dst.MLTrainingFunctionSingle)
		if string(jsonMLTrainingFunctionSingle) == "{}" { // empty struct
			dst.MLTrainingFunctionSingle = nil
		} else {
			match++
		}
	} else {
		dst.MLTrainingFunctionSingle = nil
	}

	// try to unmarshal data into MLTrainingProcessSingle
	err = newStrictDecoder(data).Decode(&dst.MLTrainingProcessSingle)
	if err == nil {
		jsonMLTrainingProcessSingle, _ := json.Marshal(dst.MLTrainingProcessSingle)
		if string(jsonMLTrainingProcessSingle) == "{}" { // empty struct
			dst.MLTrainingProcessSingle = nil
		} else {
			match++
		}
	} else {
		dst.MLTrainingProcessSingle = nil
	}

	// try to unmarshal data into MLTrainingReportSingle
	err = newStrictDecoder(data).Decode(&dst.MLTrainingReportSingle)
	if err == nil {
		jsonMLTrainingReportSingle, _ := json.Marshal(dst.MLTrainingReportSingle)
		if string(jsonMLTrainingReportSingle) == "{}" { // empty struct
			dst.MLTrainingReportSingle = nil
		} else {
			match++
		}
	} else {
		dst.MLTrainingReportSingle = nil
	}

	// try to unmarshal data into MLTrainingRequestSingle
	err = newStrictDecoder(data).Decode(&dst.MLTrainingRequestSingle)
	if err == nil {
		jsonMLTrainingRequestSingle, _ := json.Marshal(dst.MLTrainingRequestSingle)
		if string(jsonMLTrainingRequestSingle) == "{}" { // empty struct
			dst.MLTrainingRequestSingle = nil
		} else {
			match++
		}
	} else {
		dst.MLTrainingRequestSingle = nil
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
		dst.MLTrainingFunctionSingle = nil
		dst.MLTrainingProcessSingle = nil
		dst.MLTrainingReportSingle = nil
		dst.MLTrainingRequestSingle = nil
		dst.ManagedElementSingle = nil
		dst.SubNetworkSingle = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ResourcesAiMlNrm)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ResourcesAiMlNrm)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ResourcesAiMlNrm) MarshalJSON() ([]byte, error) {
	if src.MLTrainingFunctionSingle != nil {
		return json.Marshal(&src.MLTrainingFunctionSingle)
	}

	if src.MLTrainingProcessSingle != nil {
		return json.Marshal(&src.MLTrainingProcessSingle)
	}

	if src.MLTrainingReportSingle != nil {
		return json.Marshal(&src.MLTrainingReportSingle)
	}

	if src.MLTrainingRequestSingle != nil {
		return json.Marshal(&src.MLTrainingRequestSingle)
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
func (obj *ResourcesAiMlNrm) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MLTrainingFunctionSingle != nil {
		return obj.MLTrainingFunctionSingle
	}

	if obj.MLTrainingProcessSingle != nil {
		return obj.MLTrainingProcessSingle
	}

	if obj.MLTrainingReportSingle != nil {
		return obj.MLTrainingReportSingle
	}

	if obj.MLTrainingRequestSingle != nil {
		return obj.MLTrainingRequestSingle
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

type NullableResourcesAiMlNrm struct {
	value *ResourcesAiMlNrm
	isSet bool
}

func (v NullableResourcesAiMlNrm) Get() *ResourcesAiMlNrm {
	return v.value
}

func (v *NullableResourcesAiMlNrm) Set(val *ResourcesAiMlNrm) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcesAiMlNrm) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcesAiMlNrm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcesAiMlNrm(val *ResourcesAiMlNrm) *NullableResourcesAiMlNrm {
	return &NullableResourcesAiMlNrm{value: val, isSet: true}
}

func (v NullableResourcesAiMlNrm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcesAiMlNrm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

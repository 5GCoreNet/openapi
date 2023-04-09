/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
	"fmt"
)

// ResourcesIntentNrm - struct for ResourcesIntentNrm
type ResourcesIntentNrm struct {
	IntentSingle     *IntentSingle
	SubNetworkSingle *SubNetworkSingle
}

// IntentSingleAsResourcesIntentNrm is a convenience function that returns IntentSingle wrapped in ResourcesIntentNrm
func IntentSingleAsResourcesIntentNrm(v *IntentSingle) ResourcesIntentNrm {
	return ResourcesIntentNrm{
		IntentSingle: v,
	}
}

// SubNetworkSingleAsResourcesIntentNrm is a convenience function that returns SubNetworkSingle wrapped in ResourcesIntentNrm
func SubNetworkSingleAsResourcesIntentNrm(v *SubNetworkSingle) ResourcesIntentNrm {
	return ResourcesIntentNrm{
		SubNetworkSingle: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ResourcesIntentNrm) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IntentSingle
	err = newStrictDecoder(data).Decode(&dst.IntentSingle)
	if err == nil {
		jsonIntentSingle, _ := json.Marshal(dst.IntentSingle)
		if string(jsonIntentSingle) == "{}" { // empty struct
			dst.IntentSingle = nil
		} else {
			match++
		}
	} else {
		dst.IntentSingle = nil
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
		dst.IntentSingle = nil
		dst.SubNetworkSingle = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ResourcesIntentNrm)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ResourcesIntentNrm)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ResourcesIntentNrm) MarshalJSON() ([]byte, error) {
	if src.IntentSingle != nil {
		return json.Marshal(&src.IntentSingle)
	}

	if src.SubNetworkSingle != nil {
		return json.Marshal(&src.SubNetworkSingle)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ResourcesIntentNrm) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IntentSingle != nil {
		return obj.IntentSingle
	}

	if obj.SubNetworkSingle != nil {
		return obj.SubNetworkSingle
	}

	// all schemas are nil
	return nil
}

type NullableResourcesIntentNrm struct {
	value *ResourcesIntentNrm
	isSet bool
}

func (v NullableResourcesIntentNrm) Get() *ResourcesIntentNrm {
	return v.value
}

func (v *NullableResourcesIntentNrm) Set(val *ResourcesIntentNrm) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcesIntentNrm) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcesIntentNrm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcesIntentNrm(val *ResourcesIntentNrm) *NullableResourcesIntentNrm {
	return &NullableResourcesIntentNrm{value: val, isSet: true}
}

func (v NullableResourcesIntentNrm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcesIntentNrm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

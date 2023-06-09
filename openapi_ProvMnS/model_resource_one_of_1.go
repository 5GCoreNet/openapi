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

// ResourceOneOf1 struct for ResourceOneOf1
type ResourceOneOf1 struct {
	Resources5gcNrm     *Resources5gcNrm
	ResourcesAiMlNrm    *ResourcesAiMlNrm
	ResourcesCoslaNrm   *ResourcesCoslaNrm
	ResourcesGenericNrm *ResourcesGenericNrm
	ResourcesIntentNrm  *ResourcesIntentNrm
	ResourcesMdaNrm     *ResourcesMdaNrm
	ResourcesNrNrm      *ResourcesNrNrm
	ResourcesSliceNrm   *ResourcesSliceNrm
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResourceOneOf1) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Resources5gcNrm
	err = json.Unmarshal(data, &dst.Resources5gcNrm)
	if err == nil {
		jsonResources5gcNrm, _ := json.Marshal(dst.Resources5gcNrm)
		if string(jsonResources5gcNrm) == "{}" { // empty struct
			dst.Resources5gcNrm = nil
		} else {
			return nil // data stored in dst.Resources5gcNrm, return on the first match
		}
	} else {
		dst.Resources5gcNrm = nil
	}

	// try to unmarshal JSON data into ResourcesAiMlNrm
	err = json.Unmarshal(data, &dst.ResourcesAiMlNrm)
	if err == nil {
		jsonResourcesAiMlNrm, _ := json.Marshal(dst.ResourcesAiMlNrm)
		if string(jsonResourcesAiMlNrm) == "{}" { // empty struct
			dst.ResourcesAiMlNrm = nil
		} else {
			return nil // data stored in dst.ResourcesAiMlNrm, return on the first match
		}
	} else {
		dst.ResourcesAiMlNrm = nil
	}

	// try to unmarshal JSON data into ResourcesCoslaNrm
	err = json.Unmarshal(data, &dst.ResourcesCoslaNrm)
	if err == nil {
		jsonResourcesCoslaNrm, _ := json.Marshal(dst.ResourcesCoslaNrm)
		if string(jsonResourcesCoslaNrm) == "{}" { // empty struct
			dst.ResourcesCoslaNrm = nil
		} else {
			return nil // data stored in dst.ResourcesCoslaNrm, return on the first match
		}
	} else {
		dst.ResourcesCoslaNrm = nil
	}

	// try to unmarshal JSON data into ResourcesGenericNrm
	err = json.Unmarshal(data, &dst.ResourcesGenericNrm)
	if err == nil {
		jsonResourcesGenericNrm, _ := json.Marshal(dst.ResourcesGenericNrm)
		if string(jsonResourcesGenericNrm) == "{}" { // empty struct
			dst.ResourcesGenericNrm = nil
		} else {
			return nil // data stored in dst.ResourcesGenericNrm, return on the first match
		}
	} else {
		dst.ResourcesGenericNrm = nil
	}

	// try to unmarshal JSON data into ResourcesIntentNrm
	err = json.Unmarshal(data, &dst.ResourcesIntentNrm)
	if err == nil {
		jsonResourcesIntentNrm, _ := json.Marshal(dst.ResourcesIntentNrm)
		if string(jsonResourcesIntentNrm) == "{}" { // empty struct
			dst.ResourcesIntentNrm = nil
		} else {
			return nil // data stored in dst.ResourcesIntentNrm, return on the first match
		}
	} else {
		dst.ResourcesIntentNrm = nil
	}

	// try to unmarshal JSON data into ResourcesMdaNrm
	err = json.Unmarshal(data, &dst.ResourcesMdaNrm)
	if err == nil {
		jsonResourcesMdaNrm, _ := json.Marshal(dst.ResourcesMdaNrm)
		if string(jsonResourcesMdaNrm) == "{}" { // empty struct
			dst.ResourcesMdaNrm = nil
		} else {
			return nil // data stored in dst.ResourcesMdaNrm, return on the first match
		}
	} else {
		dst.ResourcesMdaNrm = nil
	}

	// try to unmarshal JSON data into ResourcesNrNrm
	err = json.Unmarshal(data, &dst.ResourcesNrNrm)
	if err == nil {
		jsonResourcesNrNrm, _ := json.Marshal(dst.ResourcesNrNrm)
		if string(jsonResourcesNrNrm) == "{}" { // empty struct
			dst.ResourcesNrNrm = nil
		} else {
			return nil // data stored in dst.ResourcesNrNrm, return on the first match
		}
	} else {
		dst.ResourcesNrNrm = nil
	}

	// try to unmarshal JSON data into ResourcesSliceNrm
	err = json.Unmarshal(data, &dst.ResourcesSliceNrm)
	if err == nil {
		jsonResourcesSliceNrm, _ := json.Marshal(dst.ResourcesSliceNrm)
		if string(jsonResourcesSliceNrm) == "{}" { // empty struct
			dst.ResourcesSliceNrm = nil
		} else {
			return nil // data stored in dst.ResourcesSliceNrm, return on the first match
		}
	} else {
		dst.ResourcesSliceNrm = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ResourceOneOf1)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ResourceOneOf1) MarshalJSON() ([]byte, error) {
	if src.Resources5gcNrm != nil {
		return json.Marshal(&src.Resources5gcNrm)
	}

	if src.ResourcesAiMlNrm != nil {
		return json.Marshal(&src.ResourcesAiMlNrm)
	}

	if src.ResourcesCoslaNrm != nil {
		return json.Marshal(&src.ResourcesCoslaNrm)
	}

	if src.ResourcesGenericNrm != nil {
		return json.Marshal(&src.ResourcesGenericNrm)
	}

	if src.ResourcesIntentNrm != nil {
		return json.Marshal(&src.ResourcesIntentNrm)
	}

	if src.ResourcesMdaNrm != nil {
		return json.Marshal(&src.ResourcesMdaNrm)
	}

	if src.ResourcesNrNrm != nil {
		return json.Marshal(&src.ResourcesNrNrm)
	}

	if src.ResourcesSliceNrm != nil {
		return json.Marshal(&src.ResourcesSliceNrm)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableResourceOneOf1 struct {
	value *ResourceOneOf1
	isSet bool
}

func (v NullableResourceOneOf1) Get() *ResourceOneOf1 {
	return v.value
}

func (v *NullableResourceOneOf1) Set(val *ResourceOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOneOf1(val *ResourceOneOf1) *NullableResourceOneOf1 {
	return &NullableResourceOneOf1{value: val, isSet: true}
}

func (v NullableResourceOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

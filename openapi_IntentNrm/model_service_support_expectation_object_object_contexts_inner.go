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

// ServiceSupportExpectationObjectObjectContextsInner - struct for ServiceSupportExpectationObjectObjectContextsInner
type ServiceSupportExpectationObjectObjectContextsInner struct {
	CoverageAreaTAContext        *CoverageAreaTAContext
	EdgeIdenfiticationIdContext  *EdgeIdenfiticationIdContext
	EdgeIdenfiticationLocContext *EdgeIdenfiticationLocContext
	ObjectContext                *ObjectContext
}

// CoverageAreaTAContextAsServiceSupportExpectationObjectObjectContextsInner is a convenience function that returns CoverageAreaTAContext wrapped in ServiceSupportExpectationObjectObjectContextsInner
func CoverageAreaTAContextAsServiceSupportExpectationObjectObjectContextsInner(v *CoverageAreaTAContext) ServiceSupportExpectationObjectObjectContextsInner {
	return ServiceSupportExpectationObjectObjectContextsInner{
		CoverageAreaTAContext: v,
	}
}

// EdgeIdenfiticationIdContextAsServiceSupportExpectationObjectObjectContextsInner is a convenience function that returns EdgeIdenfiticationIdContext wrapped in ServiceSupportExpectationObjectObjectContextsInner
func EdgeIdenfiticationIdContextAsServiceSupportExpectationObjectObjectContextsInner(v *EdgeIdenfiticationIdContext) ServiceSupportExpectationObjectObjectContextsInner {
	return ServiceSupportExpectationObjectObjectContextsInner{
		EdgeIdenfiticationIdContext: v,
	}
}

// EdgeIdenfiticationLocContextAsServiceSupportExpectationObjectObjectContextsInner is a convenience function that returns EdgeIdenfiticationLocContext wrapped in ServiceSupportExpectationObjectObjectContextsInner
func EdgeIdenfiticationLocContextAsServiceSupportExpectationObjectObjectContextsInner(v *EdgeIdenfiticationLocContext) ServiceSupportExpectationObjectObjectContextsInner {
	return ServiceSupportExpectationObjectObjectContextsInner{
		EdgeIdenfiticationLocContext: v,
	}
}

// ObjectContextAsServiceSupportExpectationObjectObjectContextsInner is a convenience function that returns ObjectContext wrapped in ServiceSupportExpectationObjectObjectContextsInner
func ObjectContextAsServiceSupportExpectationObjectObjectContextsInner(v *ObjectContext) ServiceSupportExpectationObjectObjectContextsInner {
	return ServiceSupportExpectationObjectObjectContextsInner{
		ObjectContext: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServiceSupportExpectationObjectObjectContextsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CoverageAreaTAContext
	err = newStrictDecoder(data).Decode(&dst.CoverageAreaTAContext)
	if err == nil {
		jsonCoverageAreaTAContext, _ := json.Marshal(dst.CoverageAreaTAContext)
		if string(jsonCoverageAreaTAContext) == "{}" { // empty struct
			dst.CoverageAreaTAContext = nil
		} else {
			match++
		}
	} else {
		dst.CoverageAreaTAContext = nil
	}

	// try to unmarshal data into EdgeIdenfiticationIdContext
	err = newStrictDecoder(data).Decode(&dst.EdgeIdenfiticationIdContext)
	if err == nil {
		jsonEdgeIdenfiticationIdContext, _ := json.Marshal(dst.EdgeIdenfiticationIdContext)
		if string(jsonEdgeIdenfiticationIdContext) == "{}" { // empty struct
			dst.EdgeIdenfiticationIdContext = nil
		} else {
			match++
		}
	} else {
		dst.EdgeIdenfiticationIdContext = nil
	}

	// try to unmarshal data into EdgeIdenfiticationLocContext
	err = newStrictDecoder(data).Decode(&dst.EdgeIdenfiticationLocContext)
	if err == nil {
		jsonEdgeIdenfiticationLocContext, _ := json.Marshal(dst.EdgeIdenfiticationLocContext)
		if string(jsonEdgeIdenfiticationLocContext) == "{}" { // empty struct
			dst.EdgeIdenfiticationLocContext = nil
		} else {
			match++
		}
	} else {
		dst.EdgeIdenfiticationLocContext = nil
	}

	// try to unmarshal data into ObjectContext
	err = newStrictDecoder(data).Decode(&dst.ObjectContext)
	if err == nil {
		jsonObjectContext, _ := json.Marshal(dst.ObjectContext)
		if string(jsonObjectContext) == "{}" { // empty struct
			dst.ObjectContext = nil
		} else {
			match++
		}
	} else {
		dst.ObjectContext = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CoverageAreaTAContext = nil
		dst.EdgeIdenfiticationIdContext = nil
		dst.EdgeIdenfiticationLocContext = nil
		dst.ObjectContext = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ServiceSupportExpectationObjectObjectContextsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ServiceSupportExpectationObjectObjectContextsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServiceSupportExpectationObjectObjectContextsInner) MarshalJSON() ([]byte, error) {
	if src.CoverageAreaTAContext != nil {
		return json.Marshal(&src.CoverageAreaTAContext)
	}

	if src.EdgeIdenfiticationIdContext != nil {
		return json.Marshal(&src.EdgeIdenfiticationIdContext)
	}

	if src.EdgeIdenfiticationLocContext != nil {
		return json.Marshal(&src.EdgeIdenfiticationLocContext)
	}

	if src.ObjectContext != nil {
		return json.Marshal(&src.ObjectContext)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServiceSupportExpectationObjectObjectContextsInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CoverageAreaTAContext != nil {
		return obj.CoverageAreaTAContext
	}

	if obj.EdgeIdenfiticationIdContext != nil {
		return obj.EdgeIdenfiticationIdContext
	}

	if obj.EdgeIdenfiticationLocContext != nil {
		return obj.EdgeIdenfiticationLocContext
	}

	if obj.ObjectContext != nil {
		return obj.ObjectContext
	}

	// all schemas are nil
	return nil
}

type NullableServiceSupportExpectationObjectObjectContextsInner struct {
	value *ServiceSupportExpectationObjectObjectContextsInner
	isSet bool
}

func (v NullableServiceSupportExpectationObjectObjectContextsInner) Get() *ServiceSupportExpectationObjectObjectContextsInner {
	return v.value
}

func (v *NullableServiceSupportExpectationObjectObjectContextsInner) Set(val *ServiceSupportExpectationObjectObjectContextsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSupportExpectationObjectObjectContextsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSupportExpectationObjectObjectContextsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSupportExpectationObjectObjectContextsInner(val *ServiceSupportExpectationObjectObjectContextsInner) *NullableServiceSupportExpectationObjectObjectContextsInner {
	return &NullableServiceSupportExpectationObjectObjectContextsInner{value: val, isSet: true}
}

func (v NullableServiceSupportExpectationObjectObjectContextsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSupportExpectationObjectObjectContextsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

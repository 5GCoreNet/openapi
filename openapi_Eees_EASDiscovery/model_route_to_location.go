/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EASDiscovery

import (
	"encoding/json"
	"fmt"
)

// RouteToLocation At least one of the \"routeInfo\" attribute and the \"routeProfId\" attribute shall be included in the \"RouteToLocation\" data type. 
type RouteToLocation struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RouteToLocation) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.Interface);
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			return nil // data stored in dst.Interface, return on the first match
		}
	} else {
		dst.Interface = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RouteToLocation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RouteToLocation) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRouteToLocation struct {
	value *RouteToLocation
	isSet bool
}

func (v NullableRouteToLocation) Get() *RouteToLocation {
	return v.value
}

func (v *NullableRouteToLocation) Set(val *RouteToLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteToLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteToLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteToLocation(val *RouteToLocation) *NullableRouteToLocation {
	return &NullableRouteToLocation{value: val, isSet: true}
}

func (v NullableRouteToLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteToLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



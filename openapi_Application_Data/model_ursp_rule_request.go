/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Application_Data

import (
	"encoding/json"
)

// checks if the UrspRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UrspRuleRequest{}

// UrspRuleRequest Contains parameters that can be used to guide the URSP.
type UrspRuleRequest struct {
	TrafficDesc *TrafficDescriptorComponents `json:"trafficDesc,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RelatPrecedence *int32 `json:"relatPrecedence,omitempty"`
	// Sets of parameters that may be used to guide the Route Selection Descriptors of the  URSP. 
	RouteSelParamSets []RouteSelectionParameterSet `json:"routeSelParamSets,omitempty"`
}

// NewUrspRuleRequest instantiates a new UrspRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrspRuleRequest() *UrspRuleRequest {
	this := UrspRuleRequest{}
	return &this
}

// NewUrspRuleRequestWithDefaults instantiates a new UrspRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrspRuleRequestWithDefaults() *UrspRuleRequest {
	this := UrspRuleRequest{}
	return &this
}

// GetTrafficDesc returns the TrafficDesc field value if set, zero value otherwise.
func (o *UrspRuleRequest) GetTrafficDesc() TrafficDescriptorComponents {
	if o == nil || isNil(o.TrafficDesc) {
		var ret TrafficDescriptorComponents
		return ret
	}
	return *o.TrafficDesc
}

// GetTrafficDescOk returns a tuple with the TrafficDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrspRuleRequest) GetTrafficDescOk() (*TrafficDescriptorComponents, bool) {
	if o == nil || isNil(o.TrafficDesc) {
		return nil, false
	}
	return o.TrafficDesc, true
}

// HasTrafficDesc returns a boolean if a field has been set.
func (o *UrspRuleRequest) HasTrafficDesc() bool {
	if o != nil && !isNil(o.TrafficDesc) {
		return true
	}

	return false
}

// SetTrafficDesc gets a reference to the given TrafficDescriptorComponents and assigns it to the TrafficDesc field.
func (o *UrspRuleRequest) SetTrafficDesc(v TrafficDescriptorComponents) {
	o.TrafficDesc = &v
}

// GetRelatPrecedence returns the RelatPrecedence field value if set, zero value otherwise.
func (o *UrspRuleRequest) GetRelatPrecedence() int32 {
	if o == nil || isNil(o.RelatPrecedence) {
		var ret int32
		return ret
	}
	return *o.RelatPrecedence
}

// GetRelatPrecedenceOk returns a tuple with the RelatPrecedence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrspRuleRequest) GetRelatPrecedenceOk() (*int32, bool) {
	if o == nil || isNil(o.RelatPrecedence) {
		return nil, false
	}
	return o.RelatPrecedence, true
}

// HasRelatPrecedence returns a boolean if a field has been set.
func (o *UrspRuleRequest) HasRelatPrecedence() bool {
	if o != nil && !isNil(o.RelatPrecedence) {
		return true
	}

	return false
}

// SetRelatPrecedence gets a reference to the given int32 and assigns it to the RelatPrecedence field.
func (o *UrspRuleRequest) SetRelatPrecedence(v int32) {
	o.RelatPrecedence = &v
}

// GetRouteSelParamSets returns the RouteSelParamSets field value if set, zero value otherwise.
func (o *UrspRuleRequest) GetRouteSelParamSets() []RouteSelectionParameterSet {
	if o == nil || isNil(o.RouteSelParamSets) {
		var ret []RouteSelectionParameterSet
		return ret
	}
	return o.RouteSelParamSets
}

// GetRouteSelParamSetsOk returns a tuple with the RouteSelParamSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrspRuleRequest) GetRouteSelParamSetsOk() ([]RouteSelectionParameterSet, bool) {
	if o == nil || isNil(o.RouteSelParamSets) {
		return nil, false
	}
	return o.RouteSelParamSets, true
}

// HasRouteSelParamSets returns a boolean if a field has been set.
func (o *UrspRuleRequest) HasRouteSelParamSets() bool {
	if o != nil && !isNil(o.RouteSelParamSets) {
		return true
	}

	return false
}

// SetRouteSelParamSets gets a reference to the given []RouteSelectionParameterSet and assigns it to the RouteSelParamSets field.
func (o *UrspRuleRequest) SetRouteSelParamSets(v []RouteSelectionParameterSet) {
	o.RouteSelParamSets = v
}

func (o UrspRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UrspRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TrafficDesc) {
		toSerialize["trafficDesc"] = o.TrafficDesc
	}
	if !isNil(o.RelatPrecedence) {
		toSerialize["relatPrecedence"] = o.RelatPrecedence
	}
	if !isNil(o.RouteSelParamSets) {
		toSerialize["routeSelParamSets"] = o.RouteSelParamSets
	}
	return toSerialize, nil
}

type NullableUrspRuleRequest struct {
	value *UrspRuleRequest
	isSet bool
}

func (v NullableUrspRuleRequest) Get() *UrspRuleRequest {
	return v.value
}

func (v *NullableUrspRuleRequest) Set(val *UrspRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUrspRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUrspRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrspRuleRequest(val *UrspRuleRequest) *NullableUrspRuleRequest {
	return &NullableUrspRuleRequest{value: val, isSet: true}
}

func (v NullableUrspRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrspRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



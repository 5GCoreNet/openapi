/*
CAPIF_Events_API

API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Events_API

import (
	"encoding/json"
)

// checks if the TopologyHiding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TopologyHiding{}

// TopologyHiding Represents the routing rules information of a service API.
type TopologyHiding struct {
	ApiId        string        `json:"apiId"`
	RoutingRules []RoutingRule `json:"routingRules"`
}

// NewTopologyHiding instantiates a new TopologyHiding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologyHiding(apiId string, routingRules []RoutingRule) *TopologyHiding {
	this := TopologyHiding{}
	this.ApiId = apiId
	this.RoutingRules = routingRules
	return &this
}

// NewTopologyHidingWithDefaults instantiates a new TopologyHiding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologyHidingWithDefaults() *TopologyHiding {
	this := TopologyHiding{}
	return &this
}

// GetApiId returns the ApiId field value
func (o *TopologyHiding) GetApiId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value
// and a boolean to check if the value has been set.
func (o *TopologyHiding) GetApiIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiId, true
}

// SetApiId sets field value
func (o *TopologyHiding) SetApiId(v string) {
	o.ApiId = v
}

// GetRoutingRules returns the RoutingRules field value
func (o *TopologyHiding) GetRoutingRules() []RoutingRule {
	if o == nil {
		var ret []RoutingRule
		return ret
	}

	return o.RoutingRules
}

// GetRoutingRulesOk returns a tuple with the RoutingRules field value
// and a boolean to check if the value has been set.
func (o *TopologyHiding) GetRoutingRulesOk() ([]RoutingRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoutingRules, true
}

// SetRoutingRules sets field value
func (o *TopologyHiding) SetRoutingRules(v []RoutingRule) {
	o.RoutingRules = v
}

func (o TopologyHiding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopologyHiding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiId"] = o.ApiId
	toSerialize["routingRules"] = o.RoutingRules
	return toSerialize, nil
}

type NullableTopologyHiding struct {
	value *TopologyHiding
	isSet bool
}

func (v NullableTopologyHiding) Get() *TopologyHiding {
	return v.value
}

func (v *NullableTopologyHiding) Set(val *TopologyHiding) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologyHiding) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologyHiding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologyHiding(val *TopologyHiding) *NullableTopologyHiding {
	return &NullableTopologyHiding{value: val, isSet: true}
}

func (v NullableTopologyHiding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologyHiding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

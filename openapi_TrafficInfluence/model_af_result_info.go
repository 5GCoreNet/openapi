/*
3gpp-traffic-influence

API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_TrafficInfluence

import (
	"encoding/json"
)

// checks if the AfResultInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AfResultInfo{}

// AfResultInfo Identifies the result of application layer handling.
type AfResultInfo struct {
	AfStatus     AfResultStatus          `json:"afStatus"`
	TrafficRoute NullableRouteToLocation `json:"trafficRoute,omitempty"`
	// If present and set to \"true\" it indicates that buffering of uplink traffic to the target DNAI is needed.
	UpBuffInd *bool `json:"upBuffInd,omitempty"`
	// Contains EAS IP replacement information.
	EasIpReplaceInfos []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
}

// NewAfResultInfo instantiates a new AfResultInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfResultInfo(afStatus AfResultStatus) *AfResultInfo {
	this := AfResultInfo{}
	this.AfStatus = afStatus
	return &this
}

// NewAfResultInfoWithDefaults instantiates a new AfResultInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfResultInfoWithDefaults() *AfResultInfo {
	this := AfResultInfo{}
	return &this
}

// GetAfStatus returns the AfStatus field value
func (o *AfResultInfo) GetAfStatus() AfResultStatus {
	if o == nil {
		var ret AfResultStatus
		return ret
	}

	return o.AfStatus
}

// GetAfStatusOk returns a tuple with the AfStatus field value
// and a boolean to check if the value has been set.
func (o *AfResultInfo) GetAfStatusOk() (*AfResultStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfStatus, true
}

// SetAfStatus sets field value
func (o *AfResultInfo) SetAfStatus(v AfResultStatus) {
	o.AfStatus = v
}

// GetTrafficRoute returns the TrafficRoute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AfResultInfo) GetTrafficRoute() RouteToLocation {
	if o == nil || IsNil(o.TrafficRoute.Get()) {
		var ret RouteToLocation
		return ret
	}
	return *o.TrafficRoute.Get()
}

// GetTrafficRouteOk returns a tuple with the TrafficRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AfResultInfo) GetTrafficRouteOk() (*RouteToLocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrafficRoute.Get(), o.TrafficRoute.IsSet()
}

// HasTrafficRoute returns a boolean if a field has been set.
func (o *AfResultInfo) HasTrafficRoute() bool {
	if o != nil && o.TrafficRoute.IsSet() {
		return true
	}

	return false
}

// SetTrafficRoute gets a reference to the given NullableRouteToLocation and assigns it to the TrafficRoute field.
func (o *AfResultInfo) SetTrafficRoute(v RouteToLocation) {
	o.TrafficRoute.Set(&v)
}

// SetTrafficRouteNil sets the value for TrafficRoute to be an explicit nil
func (o *AfResultInfo) SetTrafficRouteNil() {
	o.TrafficRoute.Set(nil)
}

// UnsetTrafficRoute ensures that no value is present for TrafficRoute, not even an explicit nil
func (o *AfResultInfo) UnsetTrafficRoute() {
	o.TrafficRoute.Unset()
}

// GetUpBuffInd returns the UpBuffInd field value if set, zero value otherwise.
func (o *AfResultInfo) GetUpBuffInd() bool {
	if o == nil || IsNil(o.UpBuffInd) {
		var ret bool
		return ret
	}
	return *o.UpBuffInd
}

// GetUpBuffIndOk returns a tuple with the UpBuffInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfResultInfo) GetUpBuffIndOk() (*bool, bool) {
	if o == nil || IsNil(o.UpBuffInd) {
		return nil, false
	}
	return o.UpBuffInd, true
}

// HasUpBuffInd returns a boolean if a field has been set.
func (o *AfResultInfo) HasUpBuffInd() bool {
	if o != nil && !IsNil(o.UpBuffInd) {
		return true
	}

	return false
}

// SetUpBuffInd gets a reference to the given bool and assigns it to the UpBuffInd field.
func (o *AfResultInfo) SetUpBuffInd(v bool) {
	o.UpBuffInd = &v
}

// GetEasIpReplaceInfos returns the EasIpReplaceInfos field value if set, zero value otherwise.
func (o *AfResultInfo) GetEasIpReplaceInfos() []EasIpReplacementInfo {
	if o == nil || IsNil(o.EasIpReplaceInfos) {
		var ret []EasIpReplacementInfo
		return ret
	}
	return o.EasIpReplaceInfos
}

// GetEasIpReplaceInfosOk returns a tuple with the EasIpReplaceInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfResultInfo) GetEasIpReplaceInfosOk() ([]EasIpReplacementInfo, bool) {
	if o == nil || IsNil(o.EasIpReplaceInfos) {
		return nil, false
	}
	return o.EasIpReplaceInfos, true
}

// HasEasIpReplaceInfos returns a boolean if a field has been set.
func (o *AfResultInfo) HasEasIpReplaceInfos() bool {
	if o != nil && !IsNil(o.EasIpReplaceInfos) {
		return true
	}

	return false
}

// SetEasIpReplaceInfos gets a reference to the given []EasIpReplacementInfo and assigns it to the EasIpReplaceInfos field.
func (o *AfResultInfo) SetEasIpReplaceInfos(v []EasIpReplacementInfo) {
	o.EasIpReplaceInfos = v
}

func (o AfResultInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AfResultInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["afStatus"] = o.AfStatus
	if o.TrafficRoute.IsSet() {
		toSerialize["trafficRoute"] = o.TrafficRoute.Get()
	}
	if !IsNil(o.UpBuffInd) {
		toSerialize["upBuffInd"] = o.UpBuffInd
	}
	if !IsNil(o.EasIpReplaceInfos) {
		toSerialize["easIpReplaceInfos"] = o.EasIpReplaceInfos
	}
	return toSerialize, nil
}

type NullableAfResultInfo struct {
	value *AfResultInfo
	isSet bool
}

func (v NullableAfResultInfo) Get() *AfResultInfo {
	return v.value
}

func (v *NullableAfResultInfo) Set(val *AfResultInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAfResultInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAfResultInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfResultInfo(val *AfResultInfo) *NullableAfResultInfo {
	return &NullableAfResultInfo{value: val, isSet: true}
}

func (v NullableAfResultInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfResultInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

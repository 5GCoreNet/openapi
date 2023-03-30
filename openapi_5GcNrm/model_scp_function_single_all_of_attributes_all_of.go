/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the ScpFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScpFunctionSingleAllOfAttributesAllOf{}

// ScpFunctionSingleAllOfAttributesAllOf struct for ScpFunctionSingleAllOfAttributesAllOf
type ScpFunctionSingleAllOfAttributesAllOf struct {
	SupportedFuncList []SupportedFunc `json:"supportedFuncList,omitempty"`
	Address *HostAddr `json:"address,omitempty"`
	ScpInfo *ScpInfo `json:"scpInfo,omitempty"`
}

// NewScpFunctionSingleAllOfAttributesAllOf instantiates a new ScpFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScpFunctionSingleAllOfAttributesAllOf() *ScpFunctionSingleAllOfAttributesAllOf {
	this := ScpFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewScpFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new ScpFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScpFunctionSingleAllOfAttributesAllOfWithDefaults() *ScpFunctionSingleAllOfAttributesAllOf {
	this := ScpFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetSupportedFuncList returns the SupportedFuncList field value if set, zero value otherwise.
func (o *ScpFunctionSingleAllOfAttributesAllOf) GetSupportedFuncList() []SupportedFunc {
	if o == nil || IsNil(o.SupportedFuncList) {
		var ret []SupportedFunc
		return ret
	}
	return o.SupportedFuncList
}

// GetSupportedFuncListOk returns a tuple with the SupportedFuncList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpFunctionSingleAllOfAttributesAllOf) GetSupportedFuncListOk() ([]SupportedFunc, bool) {
	if o == nil || IsNil(o.SupportedFuncList) {
		return nil, false
	}
	return o.SupportedFuncList, true
}

// HasSupportedFuncList returns a boolean if a field has been set.
func (o *ScpFunctionSingleAllOfAttributesAllOf) HasSupportedFuncList() bool {
	if o != nil && !IsNil(o.SupportedFuncList) {
		return true
	}

	return false
}

// SetSupportedFuncList gets a reference to the given []SupportedFunc and assigns it to the SupportedFuncList field.
func (o *ScpFunctionSingleAllOfAttributesAllOf) SetSupportedFuncList(v []SupportedFunc) {
	o.SupportedFuncList = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ScpFunctionSingleAllOfAttributesAllOf) GetAddress() HostAddr {
	if o == nil || IsNil(o.Address) {
		var ret HostAddr
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpFunctionSingleAllOfAttributesAllOf) GetAddressOk() (*HostAddr, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ScpFunctionSingleAllOfAttributesAllOf) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given HostAddr and assigns it to the Address field.
func (o *ScpFunctionSingleAllOfAttributesAllOf) SetAddress(v HostAddr) {
	o.Address = &v
}

// GetScpInfo returns the ScpInfo field value if set, zero value otherwise.
func (o *ScpFunctionSingleAllOfAttributesAllOf) GetScpInfo() ScpInfo {
	if o == nil || IsNil(o.ScpInfo) {
		var ret ScpInfo
		return ret
	}
	return *o.ScpInfo
}

// GetScpInfoOk returns a tuple with the ScpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpFunctionSingleAllOfAttributesAllOf) GetScpInfoOk() (*ScpInfo, bool) {
	if o == nil || IsNil(o.ScpInfo) {
		return nil, false
	}
	return o.ScpInfo, true
}

// HasScpInfo returns a boolean if a field has been set.
func (o *ScpFunctionSingleAllOfAttributesAllOf) HasScpInfo() bool {
	if o != nil && !IsNil(o.ScpInfo) {
		return true
	}

	return false
}

// SetScpInfo gets a reference to the given ScpInfo and assigns it to the ScpInfo field.
func (o *ScpFunctionSingleAllOfAttributesAllOf) SetScpInfo(v ScpInfo) {
	o.ScpInfo = &v
}

func (o ScpFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScpFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupportedFuncList) {
		toSerialize["supportedFuncList"] = o.SupportedFuncList
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.ScpInfo) {
		toSerialize["scpInfo"] = o.ScpInfo
	}
	return toSerialize, nil
}

type NullableScpFunctionSingleAllOfAttributesAllOf struct {
	value *ScpFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableScpFunctionSingleAllOfAttributesAllOf) Get() *ScpFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableScpFunctionSingleAllOfAttributesAllOf) Set(val *ScpFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableScpFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableScpFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScpFunctionSingleAllOfAttributesAllOf(val *ScpFunctionSingleAllOfAttributesAllOf) *NullableScpFunctionSingleAllOfAttributesAllOf {
	return &NullableScpFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableScpFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScpFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


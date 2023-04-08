/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_Events

import (
	"encoding/json"
)

// checks if the ReferenceUEDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceUEDetail{}

// ReferenceUEDetail Reference UE details, where the UEs moving in and out to be monitored.
type ReferenceUEDetail struct {
	ValTgtUe ValTargetUe `json:"valTgtUe"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	ProxRange int32 `json:"proxRange"`
	// string with format 'float' as defined in OpenAPI.
	ProxRangeFrac *float32 `json:"proxRangeFrac,omitempty"`
}

// NewReferenceUEDetail instantiates a new ReferenceUEDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceUEDetail(valTgtUe ValTargetUe, proxRange int32) *ReferenceUEDetail {
	this := ReferenceUEDetail{}
	this.ValTgtUe = valTgtUe
	this.ProxRange = proxRange
	return &this
}

// NewReferenceUEDetailWithDefaults instantiates a new ReferenceUEDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceUEDetailWithDefaults() *ReferenceUEDetail {
	this := ReferenceUEDetail{}
	return &this
}

// GetValTgtUe returns the ValTgtUe field value
func (o *ReferenceUEDetail) GetValTgtUe() ValTargetUe {
	if o == nil {
		var ret ValTargetUe
		return ret
	}

	return o.ValTgtUe
}

// GetValTgtUeOk returns a tuple with the ValTgtUe field value
// and a boolean to check if the value has been set.
func (o *ReferenceUEDetail) GetValTgtUeOk() (*ValTargetUe, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValTgtUe, true
}

// SetValTgtUe sets field value
func (o *ReferenceUEDetail) SetValTgtUe(v ValTargetUe) {
	o.ValTgtUe = v
}

// GetProxRange returns the ProxRange field value
func (o *ReferenceUEDetail) GetProxRange() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProxRange
}

// GetProxRangeOk returns a tuple with the ProxRange field value
// and a boolean to check if the value has been set.
func (o *ReferenceUEDetail) GetProxRangeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProxRange, true
}

// SetProxRange sets field value
func (o *ReferenceUEDetail) SetProxRange(v int32) {
	o.ProxRange = v
}

// GetProxRangeFrac returns the ProxRangeFrac field value if set, zero value otherwise.
func (o *ReferenceUEDetail) GetProxRangeFrac() float32 {
	if o == nil || isNil(o.ProxRangeFrac) {
		var ret float32
		return ret
	}
	return *o.ProxRangeFrac
}

// GetProxRangeFracOk returns a tuple with the ProxRangeFrac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceUEDetail) GetProxRangeFracOk() (*float32, bool) {
	if o == nil || isNil(o.ProxRangeFrac) {
		return nil, false
	}
	return o.ProxRangeFrac, true
}

// HasProxRangeFrac returns a boolean if a field has been set.
func (o *ReferenceUEDetail) HasProxRangeFrac() bool {
	if o != nil && !isNil(o.ProxRangeFrac) {
		return true
	}

	return false
}

// SetProxRangeFrac gets a reference to the given float32 and assigns it to the ProxRangeFrac field.
func (o *ReferenceUEDetail) SetProxRangeFrac(v float32) {
	o.ProxRangeFrac = &v
}

func (o ReferenceUEDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceUEDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["valTgtUe"] = o.ValTgtUe
	toSerialize["proxRange"] = o.ProxRange
	if !isNil(o.ProxRangeFrac) {
		toSerialize["proxRangeFrac"] = o.ProxRangeFrac
	}
	return toSerialize, nil
}

type NullableReferenceUEDetail struct {
	value *ReferenceUEDetail
	isSet bool
}

func (v NullableReferenceUEDetail) Get() *ReferenceUEDetail {
	return v.value
}

func (v *NullableReferenceUEDetail) Set(val *ReferenceUEDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceUEDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceUEDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceUEDetail(val *ReferenceUEDetail) *NullableReferenceUEDetail {
	return &NullableReferenceUEDetail{value: val, isSet: true}
}

func (v NullableReferenceUEDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceUEDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



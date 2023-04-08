/*
N5g-ddnmf_Discovery API

N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N5g-ddnmf_Discovery

import (
	"encoding/json"
)

// checks if the DiscoveryAuthReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveryAuthReqData{}

// DiscoveryAuthReqData Represents Data used to request the authorization for a discoverer UE.
type DiscoveryAuthReqData struct {
	DiscType DiscoveryType `json:"discType"`
	RestrictedDiscData *DiscDataForRestricted `json:"restrictedDiscData,omitempty"`
}

// NewDiscoveryAuthReqData instantiates a new DiscoveryAuthReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryAuthReqData(discType DiscoveryType) *DiscoveryAuthReqData {
	this := DiscoveryAuthReqData{}
	this.DiscType = discType
	return &this
}

// NewDiscoveryAuthReqDataWithDefaults instantiates a new DiscoveryAuthReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryAuthReqDataWithDefaults() *DiscoveryAuthReqData {
	this := DiscoveryAuthReqData{}
	return &this
}

// GetDiscType returns the DiscType field value
func (o *DiscoveryAuthReqData) GetDiscType() DiscoveryType {
	if o == nil {
		var ret DiscoveryType
		return ret
	}

	return o.DiscType
}

// GetDiscTypeOk returns a tuple with the DiscType field value
// and a boolean to check if the value has been set.
func (o *DiscoveryAuthReqData) GetDiscTypeOk() (*DiscoveryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscType, true
}

// SetDiscType sets field value
func (o *DiscoveryAuthReqData) SetDiscType(v DiscoveryType) {
	o.DiscType = v
}

// GetRestrictedDiscData returns the RestrictedDiscData field value if set, zero value otherwise.
func (o *DiscoveryAuthReqData) GetRestrictedDiscData() DiscDataForRestricted {
	if o == nil || isNil(o.RestrictedDiscData) {
		var ret DiscDataForRestricted
		return ret
	}
	return *o.RestrictedDiscData
}

// GetRestrictedDiscDataOk returns a tuple with the RestrictedDiscData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryAuthReqData) GetRestrictedDiscDataOk() (*DiscDataForRestricted, bool) {
	if o == nil || isNil(o.RestrictedDiscData) {
		return nil, false
	}
	return o.RestrictedDiscData, true
}

// HasRestrictedDiscData returns a boolean if a field has been set.
func (o *DiscoveryAuthReqData) HasRestrictedDiscData() bool {
	if o != nil && !isNil(o.RestrictedDiscData) {
		return true
	}

	return false
}

// SetRestrictedDiscData gets a reference to the given DiscDataForRestricted and assigns it to the RestrictedDiscData field.
func (o *DiscoveryAuthReqData) SetRestrictedDiscData(v DiscDataForRestricted) {
	o.RestrictedDiscData = &v
}

func (o DiscoveryAuthReqData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryAuthReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["discType"] = o.DiscType
	if !isNil(o.RestrictedDiscData) {
		toSerialize["restrictedDiscData"] = o.RestrictedDiscData
	}
	return toSerialize, nil
}

type NullableDiscoveryAuthReqData struct {
	value *DiscoveryAuthReqData
	isSet bool
}

func (v NullableDiscoveryAuthReqData) Get() *DiscoveryAuthReqData {
	return v.value
}

func (v *NullableDiscoveryAuthReqData) Set(val *DiscoveryAuthReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryAuthReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryAuthReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryAuthReqData(val *DiscoveryAuthReqData) *NullableDiscoveryAuthReqData {
	return &NullableDiscoveryAuthReqData{value: val, isSet: true}
}

func (v NullableDiscoveryAuthReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryAuthReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



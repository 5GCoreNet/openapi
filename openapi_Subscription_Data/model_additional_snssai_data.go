/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// checks if the AdditionalSnssaiData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalSnssaiData{}

// AdditionalSnssaiData struct for AdditionalSnssaiData
type AdditionalSnssaiData struct {
	RequiredAuthnAuthz *bool `json:"requiredAuthnAuthz,omitempty"`
	SubscribedUeSliceMbr *SliceMbrRm `json:"subscribedUeSliceMbr,omitempty"`
	SubscribedNsSrgList []string `json:"subscribedNsSrgList,omitempty"`
}

// NewAdditionalSnssaiData instantiates a new AdditionalSnssaiData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalSnssaiData() *AdditionalSnssaiData {
	this := AdditionalSnssaiData{}
	return &this
}

// NewAdditionalSnssaiDataWithDefaults instantiates a new AdditionalSnssaiData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalSnssaiDataWithDefaults() *AdditionalSnssaiData {
	this := AdditionalSnssaiData{}
	return &this
}

// GetRequiredAuthnAuthz returns the RequiredAuthnAuthz field value if set, zero value otherwise.
func (o *AdditionalSnssaiData) GetRequiredAuthnAuthz() bool {
	if o == nil || IsNil(o.RequiredAuthnAuthz) {
		var ret bool
		return ret
	}
	return *o.RequiredAuthnAuthz
}

// GetRequiredAuthnAuthzOk returns a tuple with the RequiredAuthnAuthz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSnssaiData) GetRequiredAuthnAuthzOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiredAuthnAuthz) {
		return nil, false
	}
	return o.RequiredAuthnAuthz, true
}

// HasRequiredAuthnAuthz returns a boolean if a field has been set.
func (o *AdditionalSnssaiData) HasRequiredAuthnAuthz() bool {
	if o != nil && !IsNil(o.RequiredAuthnAuthz) {
		return true
	}

	return false
}

// SetRequiredAuthnAuthz gets a reference to the given bool and assigns it to the RequiredAuthnAuthz field.
func (o *AdditionalSnssaiData) SetRequiredAuthnAuthz(v bool) {
	o.RequiredAuthnAuthz = &v
}

// GetSubscribedUeSliceMbr returns the SubscribedUeSliceMbr field value if set, zero value otherwise.
func (o *AdditionalSnssaiData) GetSubscribedUeSliceMbr() SliceMbrRm {
	if o == nil || IsNil(o.SubscribedUeSliceMbr) {
		var ret SliceMbrRm
		return ret
	}
	return *o.SubscribedUeSliceMbr
}

// GetSubscribedUeSliceMbrOk returns a tuple with the SubscribedUeSliceMbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSnssaiData) GetSubscribedUeSliceMbrOk() (*SliceMbrRm, bool) {
	if o == nil || IsNil(o.SubscribedUeSliceMbr) {
		return nil, false
	}
	return o.SubscribedUeSliceMbr, true
}

// HasSubscribedUeSliceMbr returns a boolean if a field has been set.
func (o *AdditionalSnssaiData) HasSubscribedUeSliceMbr() bool {
	if o != nil && !IsNil(o.SubscribedUeSliceMbr) {
		return true
	}

	return false
}

// SetSubscribedUeSliceMbr gets a reference to the given SliceMbrRm and assigns it to the SubscribedUeSliceMbr field.
func (o *AdditionalSnssaiData) SetSubscribedUeSliceMbr(v SliceMbrRm) {
	o.SubscribedUeSliceMbr = &v
}

// GetSubscribedNsSrgList returns the SubscribedNsSrgList field value if set, zero value otherwise.
func (o *AdditionalSnssaiData) GetSubscribedNsSrgList() []string {
	if o == nil || IsNil(o.SubscribedNsSrgList) {
		var ret []string
		return ret
	}
	return o.SubscribedNsSrgList
}

// GetSubscribedNsSrgListOk returns a tuple with the SubscribedNsSrgList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSnssaiData) GetSubscribedNsSrgListOk() ([]string, bool) {
	if o == nil || IsNil(o.SubscribedNsSrgList) {
		return nil, false
	}
	return o.SubscribedNsSrgList, true
}

// HasSubscribedNsSrgList returns a boolean if a field has been set.
func (o *AdditionalSnssaiData) HasSubscribedNsSrgList() bool {
	if o != nil && !IsNil(o.SubscribedNsSrgList) {
		return true
	}

	return false
}

// SetSubscribedNsSrgList gets a reference to the given []string and assigns it to the SubscribedNsSrgList field.
func (o *AdditionalSnssaiData) SetSubscribedNsSrgList(v []string) {
	o.SubscribedNsSrgList = v
}

func (o AdditionalSnssaiData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalSnssaiData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequiredAuthnAuthz) {
		toSerialize["requiredAuthnAuthz"] = o.RequiredAuthnAuthz
	}
	if !IsNil(o.SubscribedUeSliceMbr) {
		toSerialize["subscribedUeSliceMbr"] = o.SubscribedUeSliceMbr
	}
	if !IsNil(o.SubscribedNsSrgList) {
		toSerialize["subscribedNsSrgList"] = o.SubscribedNsSrgList
	}
	return toSerialize, nil
}

type NullableAdditionalSnssaiData struct {
	value *AdditionalSnssaiData
	isSet bool
}

func (v NullableAdditionalSnssaiData) Get() *AdditionalSnssaiData {
	return v.value
}

func (v *NullableAdditionalSnssaiData) Set(val *AdditionalSnssaiData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalSnssaiData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalSnssaiData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalSnssaiData(val *AdditionalSnssaiData) *NullableAdditionalSnssaiData {
	return &NullableAdditionalSnssaiData{value: val, isSet: true}
}

func (v NullableAdditionalSnssaiData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalSnssaiData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


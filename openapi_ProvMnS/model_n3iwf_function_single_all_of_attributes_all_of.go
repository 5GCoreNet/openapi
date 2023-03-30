/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the N3iwfFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &N3iwfFunctionSingleAllOfAttributesAllOf{}

// N3iwfFunctionSingleAllOfAttributesAllOf struct for N3iwfFunctionSingleAllOfAttributesAllOf
type N3iwfFunctionSingleAllOfAttributesAllOf struct {
	PlmnIdList []PlmnId `json:"plmnIdList,omitempty"`
	CommModelList []CommModel `json:"commModelList,omitempty"`
}

// NewN3iwfFunctionSingleAllOfAttributesAllOf instantiates a new N3iwfFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN3iwfFunctionSingleAllOfAttributesAllOf() *N3iwfFunctionSingleAllOfAttributesAllOf {
	this := N3iwfFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewN3iwfFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new N3iwfFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN3iwfFunctionSingleAllOfAttributesAllOfWithDefaults() *N3iwfFunctionSingleAllOfAttributesAllOf {
	this := N3iwfFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetPlmnIdList returns the PlmnIdList field value if set, zero value otherwise.
func (o *N3iwfFunctionSingleAllOfAttributesAllOf) GetPlmnIdList() []PlmnId {
	if o == nil || IsNil(o.PlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.PlmnIdList
}

// GetPlmnIdListOk returns a tuple with the PlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3iwfFunctionSingleAllOfAttributesAllOf) GetPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.PlmnIdList) {
		return nil, false
	}
	return o.PlmnIdList, true
}

// HasPlmnIdList returns a boolean if a field has been set.
func (o *N3iwfFunctionSingleAllOfAttributesAllOf) HasPlmnIdList() bool {
	if o != nil && !IsNil(o.PlmnIdList) {
		return true
	}

	return false
}

// SetPlmnIdList gets a reference to the given []PlmnId and assigns it to the PlmnIdList field.
func (o *N3iwfFunctionSingleAllOfAttributesAllOf) SetPlmnIdList(v []PlmnId) {
	o.PlmnIdList = v
}

// GetCommModelList returns the CommModelList field value if set, zero value otherwise.
func (o *N3iwfFunctionSingleAllOfAttributesAllOf) GetCommModelList() []CommModel {
	if o == nil || IsNil(o.CommModelList) {
		var ret []CommModel
		return ret
	}
	return o.CommModelList
}

// GetCommModelListOk returns a tuple with the CommModelList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3iwfFunctionSingleAllOfAttributesAllOf) GetCommModelListOk() ([]CommModel, bool) {
	if o == nil || IsNil(o.CommModelList) {
		return nil, false
	}
	return o.CommModelList, true
}

// HasCommModelList returns a boolean if a field has been set.
func (o *N3iwfFunctionSingleAllOfAttributesAllOf) HasCommModelList() bool {
	if o != nil && !IsNil(o.CommModelList) {
		return true
	}

	return false
}

// SetCommModelList gets a reference to the given []CommModel and assigns it to the CommModelList field.
func (o *N3iwfFunctionSingleAllOfAttributesAllOf) SetCommModelList(v []CommModel) {
	o.CommModelList = v
}

func (o N3iwfFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o N3iwfFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlmnIdList) {
		toSerialize["plmnIdList"] = o.PlmnIdList
	}
	if !IsNil(o.CommModelList) {
		toSerialize["commModelList"] = o.CommModelList
	}
	return toSerialize, nil
}

type NullableN3iwfFunctionSingleAllOfAttributesAllOf struct {
	value *N3iwfFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableN3iwfFunctionSingleAllOfAttributesAllOf) Get() *N3iwfFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableN3iwfFunctionSingleAllOfAttributesAllOf) Set(val *N3iwfFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableN3iwfFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableN3iwfFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN3iwfFunctionSingleAllOfAttributesAllOf(val *N3iwfFunctionSingleAllOfAttributesAllOf) *NullableN3iwfFunctionSingleAllOfAttributesAllOf {
	return &NullableN3iwfFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableN3iwfFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN3iwfFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



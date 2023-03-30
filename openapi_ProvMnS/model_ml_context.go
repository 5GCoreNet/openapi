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

// checks if the MLContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MLContext{}

// MLContext struct for MLContext
type MLContext struct {
	InferenceEntityRef []string `json:"inferenceEntityRef,omitempty"`
	DataProviderRef []string `json:"dataProviderRef,omitempty"`
}

// NewMLContext instantiates a new MLContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMLContext() *MLContext {
	this := MLContext{}
	return &this
}

// NewMLContextWithDefaults instantiates a new MLContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMLContextWithDefaults() *MLContext {
	this := MLContext{}
	return &this
}

// GetInferenceEntityRef returns the InferenceEntityRef field value if set, zero value otherwise.
func (o *MLContext) GetInferenceEntityRef() []string {
	if o == nil || IsNil(o.InferenceEntityRef) {
		var ret []string
		return ret
	}
	return o.InferenceEntityRef
}

// GetInferenceEntityRefOk returns a tuple with the InferenceEntityRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLContext) GetInferenceEntityRefOk() ([]string, bool) {
	if o == nil || IsNil(o.InferenceEntityRef) {
		return nil, false
	}
	return o.InferenceEntityRef, true
}

// HasInferenceEntityRef returns a boolean if a field has been set.
func (o *MLContext) HasInferenceEntityRef() bool {
	if o != nil && !IsNil(o.InferenceEntityRef) {
		return true
	}

	return false
}

// SetInferenceEntityRef gets a reference to the given []string and assigns it to the InferenceEntityRef field.
func (o *MLContext) SetInferenceEntityRef(v []string) {
	o.InferenceEntityRef = v
}

// GetDataProviderRef returns the DataProviderRef field value if set, zero value otherwise.
func (o *MLContext) GetDataProviderRef() []string {
	if o == nil || IsNil(o.DataProviderRef) {
		var ret []string
		return ret
	}
	return o.DataProviderRef
}

// GetDataProviderRefOk returns a tuple with the DataProviderRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLContext) GetDataProviderRefOk() ([]string, bool) {
	if o == nil || IsNil(o.DataProviderRef) {
		return nil, false
	}
	return o.DataProviderRef, true
}

// HasDataProviderRef returns a boolean if a field has been set.
func (o *MLContext) HasDataProviderRef() bool {
	if o != nil && !IsNil(o.DataProviderRef) {
		return true
	}

	return false
}

// SetDataProviderRef gets a reference to the given []string and assigns it to the DataProviderRef field.
func (o *MLContext) SetDataProviderRef(v []string) {
	o.DataProviderRef = v
}

func (o MLContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MLContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InferenceEntityRef) {
		toSerialize["inferenceEntityRef"] = o.InferenceEntityRef
	}
	if !IsNil(o.DataProviderRef) {
		toSerialize["dataProviderRef"] = o.DataProviderRef
	}
	return toSerialize, nil
}

type NullableMLContext struct {
	value *MLContext
	isSet bool
}

func (v NullableMLContext) Get() *MLContext {
	return v.value
}

func (v *NullableMLContext) Set(val *MLContext) {
	v.value = val
	v.isSet = true
}

func (v NullableMLContext) IsSet() bool {
	return v.isSet
}

func (v *NullableMLContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMLContext(val *MLContext) *NullableMLContext {
	return &NullableMLContext{value: val, isSet: true}
}

func (v NullableMLContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMLContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


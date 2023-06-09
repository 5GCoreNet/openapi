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

// checks if the ExternalNrfFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalNrfFunctionSingleAllOfAttributes{}

// ExternalNrfFunctionSingleAllOfAttributes struct for ExternalNrfFunctionSingleAllOfAttributes
type ExternalNrfFunctionSingleAllOfAttributes struct {
	ManagedFunctionAttr
	PlmnIdList []PlmnId `json:"plmnIdList,omitempty"`
}

// NewExternalNrfFunctionSingleAllOfAttributes instantiates a new ExternalNrfFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalNrfFunctionSingleAllOfAttributes() *ExternalNrfFunctionSingleAllOfAttributes {
	this := ExternalNrfFunctionSingleAllOfAttributes{}
	return &this
}

// NewExternalNrfFunctionSingleAllOfAttributesWithDefaults instantiates a new ExternalNrfFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalNrfFunctionSingleAllOfAttributesWithDefaults() *ExternalNrfFunctionSingleAllOfAttributes {
	this := ExternalNrfFunctionSingleAllOfAttributes{}
	return &this
}

// GetPlmnIdList returns the PlmnIdList field value if set, zero value otherwise.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetPlmnIdList() []PlmnId {
	if o == nil || IsNil(o.PlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.PlmnIdList
}

// GetPlmnIdListOk returns a tuple with the PlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) GetPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.PlmnIdList) {
		return nil, false
	}
	return o.PlmnIdList, true
}

// HasPlmnIdList returns a boolean if a field has been set.
func (o *ExternalNrfFunctionSingleAllOfAttributes) HasPlmnIdList() bool {
	if o != nil && !IsNil(o.PlmnIdList) {
		return true
	}

	return false
}

// SetPlmnIdList gets a reference to the given []PlmnId and assigns it to the PlmnIdList field.
func (o *ExternalNrfFunctionSingleAllOfAttributes) SetPlmnIdList(v []PlmnId) {
	o.PlmnIdList = v
}

func (o ExternalNrfFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalNrfFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedFunctionAttr, errManagedFunctionAttr := json.Marshal(o.ManagedFunctionAttr)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	errManagedFunctionAttr = json.Unmarshal([]byte(serializedManagedFunctionAttr), &toSerialize)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	if !IsNil(o.PlmnIdList) {
		toSerialize["plmnIdList"] = o.PlmnIdList
	}
	return toSerialize, nil
}

type NullableExternalNrfFunctionSingleAllOfAttributes struct {
	value *ExternalNrfFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableExternalNrfFunctionSingleAllOfAttributes) Get() *ExternalNrfFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableExternalNrfFunctionSingleAllOfAttributes) Set(val *ExternalNrfFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalNrfFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalNrfFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalNrfFunctionSingleAllOfAttributes(val *ExternalNrfFunctionSingleAllOfAttributes) *NullableExternalNrfFunctionSingleAllOfAttributes {
	return &NullableExternalNrfFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableExternalNrfFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalNrfFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

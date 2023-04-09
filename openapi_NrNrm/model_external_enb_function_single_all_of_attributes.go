/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the ExternalENBFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalENBFunctionSingleAllOfAttributes{}

// ExternalENBFunctionSingleAllOfAttributes struct for ExternalENBFunctionSingleAllOfAttributes
type ExternalENBFunctionSingleAllOfAttributes struct {
	ManagedFunctionAttr
	ENBId *int32 `json:"eNBId,omitempty"`
}

// NewExternalENBFunctionSingleAllOfAttributes instantiates a new ExternalENBFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalENBFunctionSingleAllOfAttributes() *ExternalENBFunctionSingleAllOfAttributes {
	this := ExternalENBFunctionSingleAllOfAttributes{}
	return &this
}

// NewExternalENBFunctionSingleAllOfAttributesWithDefaults instantiates a new ExternalENBFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalENBFunctionSingleAllOfAttributesWithDefaults() *ExternalENBFunctionSingleAllOfAttributes {
	this := ExternalENBFunctionSingleAllOfAttributes{}
	return &this
}

// GetENBId returns the ENBId field value if set, zero value otherwise.
func (o *ExternalENBFunctionSingleAllOfAttributes) GetENBId() int32 {
	if o == nil || IsNil(o.ENBId) {
		var ret int32
		return ret
	}
	return *o.ENBId
}

// GetENBIdOk returns a tuple with the ENBId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalENBFunctionSingleAllOfAttributes) GetENBIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ENBId) {
		return nil, false
	}
	return o.ENBId, true
}

// HasENBId returns a boolean if a field has been set.
func (o *ExternalENBFunctionSingleAllOfAttributes) HasENBId() bool {
	if o != nil && !IsNil(o.ENBId) {
		return true
	}

	return false
}

// SetENBId gets a reference to the given int32 and assigns it to the ENBId field.
func (o *ExternalENBFunctionSingleAllOfAttributes) SetENBId(v int32) {
	o.ENBId = &v
}

func (o ExternalENBFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalENBFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedFunctionAttr, errManagedFunctionAttr := json.Marshal(o.ManagedFunctionAttr)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	errManagedFunctionAttr = json.Unmarshal([]byte(serializedManagedFunctionAttr), &toSerialize)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	if !IsNil(o.ENBId) {
		toSerialize["eNBId"] = o.ENBId
	}
	return toSerialize, nil
}

type NullableExternalENBFunctionSingleAllOfAttributes struct {
	value *ExternalENBFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableExternalENBFunctionSingleAllOfAttributes) Get() *ExternalENBFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableExternalENBFunctionSingleAllOfAttributes) Set(val *ExternalENBFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalENBFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalENBFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalENBFunctionSingleAllOfAttributes(val *ExternalENBFunctionSingleAllOfAttributes) *NullableExternalENBFunctionSingleAllOfAttributes {
	return &NullableExternalENBFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableExternalENBFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalENBFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

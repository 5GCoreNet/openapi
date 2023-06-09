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

// checks if the ExternalGnbDuFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalGnbDuFunctionSingleAllOfAttributes{}

// ExternalGnbDuFunctionSingleAllOfAttributes struct for ExternalGnbDuFunctionSingleAllOfAttributes
type ExternalGnbDuFunctionSingleAllOfAttributes struct {
	ManagedFunctionAttr
	GnbId       *string `json:"gnbId,omitempty"`
	GnbIdLength *int32  `json:"gnbIdLength,omitempty"`
}

// NewExternalGnbDuFunctionSingleAllOfAttributes instantiates a new ExternalGnbDuFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGnbDuFunctionSingleAllOfAttributes() *ExternalGnbDuFunctionSingleAllOfAttributes {
	this := ExternalGnbDuFunctionSingleAllOfAttributes{}
	return &this
}

// NewExternalGnbDuFunctionSingleAllOfAttributesWithDefaults instantiates a new ExternalGnbDuFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGnbDuFunctionSingleAllOfAttributesWithDefaults() *ExternalGnbDuFunctionSingleAllOfAttributes {
	this := ExternalGnbDuFunctionSingleAllOfAttributes{}
	return &this
}

// GetGnbId returns the GnbId field value if set, zero value otherwise.
func (o *ExternalGnbDuFunctionSingleAllOfAttributes) GetGnbId() string {
	if o == nil || IsNil(o.GnbId) {
		var ret string
		return ret
	}
	return *o.GnbId
}

// GetGnbIdOk returns a tuple with the GnbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbDuFunctionSingleAllOfAttributes) GetGnbIdOk() (*string, bool) {
	if o == nil || IsNil(o.GnbId) {
		return nil, false
	}
	return o.GnbId, true
}

// HasGnbId returns a boolean if a field has been set.
func (o *ExternalGnbDuFunctionSingleAllOfAttributes) HasGnbId() bool {
	if o != nil && !IsNil(o.GnbId) {
		return true
	}

	return false
}

// SetGnbId gets a reference to the given string and assigns it to the GnbId field.
func (o *ExternalGnbDuFunctionSingleAllOfAttributes) SetGnbId(v string) {
	o.GnbId = &v
}

// GetGnbIdLength returns the GnbIdLength field value if set, zero value otherwise.
func (o *ExternalGnbDuFunctionSingleAllOfAttributes) GetGnbIdLength() int32 {
	if o == nil || IsNil(o.GnbIdLength) {
		var ret int32
		return ret
	}
	return *o.GnbIdLength
}

// GetGnbIdLengthOk returns a tuple with the GnbIdLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbDuFunctionSingleAllOfAttributes) GetGnbIdLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.GnbIdLength) {
		return nil, false
	}
	return o.GnbIdLength, true
}

// HasGnbIdLength returns a boolean if a field has been set.
func (o *ExternalGnbDuFunctionSingleAllOfAttributes) HasGnbIdLength() bool {
	if o != nil && !IsNil(o.GnbIdLength) {
		return true
	}

	return false
}

// SetGnbIdLength gets a reference to the given int32 and assigns it to the GnbIdLength field.
func (o *ExternalGnbDuFunctionSingleAllOfAttributes) SetGnbIdLength(v int32) {
	o.GnbIdLength = &v
}

func (o ExternalGnbDuFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalGnbDuFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedFunctionAttr, errManagedFunctionAttr := json.Marshal(o.ManagedFunctionAttr)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	errManagedFunctionAttr = json.Unmarshal([]byte(serializedManagedFunctionAttr), &toSerialize)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	if !IsNil(o.GnbId) {
		toSerialize["gnbId"] = o.GnbId
	}
	if !IsNil(o.GnbIdLength) {
		toSerialize["gnbIdLength"] = o.GnbIdLength
	}
	return toSerialize, nil
}

type NullableExternalGnbDuFunctionSingleAllOfAttributes struct {
	value *ExternalGnbDuFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableExternalGnbDuFunctionSingleAllOfAttributes) Get() *ExternalGnbDuFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableExternalGnbDuFunctionSingleAllOfAttributes) Set(val *ExternalGnbDuFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGnbDuFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGnbDuFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGnbDuFunctionSingleAllOfAttributes(val *ExternalGnbDuFunctionSingleAllOfAttributes) *NullableExternalGnbDuFunctionSingleAllOfAttributes {
	return &NullableExternalGnbDuFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableExternalGnbDuFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGnbDuFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

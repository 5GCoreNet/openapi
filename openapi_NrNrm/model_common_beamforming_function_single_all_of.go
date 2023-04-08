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

// checks if the CommonBeamformingFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonBeamformingFunctionSingleAllOf{}

// CommonBeamformingFunctionSingleAllOf struct for CommonBeamformingFunctionSingleAllOf
type CommonBeamformingFunctionSingleAllOf struct {
	Attributes *interface{} `json:"attributes,omitempty"`
}

// NewCommonBeamformingFunctionSingleAllOf instantiates a new CommonBeamformingFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonBeamformingFunctionSingleAllOf() *CommonBeamformingFunctionSingleAllOf {
	this := CommonBeamformingFunctionSingleAllOf{}
	return &this
}

// NewCommonBeamformingFunctionSingleAllOfWithDefaults instantiates a new CommonBeamformingFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonBeamformingFunctionSingleAllOfWithDefaults() *CommonBeamformingFunctionSingleAllOf {
	this := CommonBeamformingFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingleAllOf) GetAttributes() interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingleAllOf) GetAttributesOk() (*interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *CommonBeamformingFunctionSingleAllOf) SetAttributes(v interface{}) {
	o.Attributes = &v
}

func (o CommonBeamformingFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonBeamformingFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableCommonBeamformingFunctionSingleAllOf struct {
	value *CommonBeamformingFunctionSingleAllOf
	isSet bool
}

func (v NullableCommonBeamformingFunctionSingleAllOf) Get() *CommonBeamformingFunctionSingleAllOf {
	return v.value
}

func (v *NullableCommonBeamformingFunctionSingleAllOf) Set(val *CommonBeamformingFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonBeamformingFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonBeamformingFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonBeamformingFunctionSingleAllOf(val *CommonBeamformingFunctionSingleAllOf) *NullableCommonBeamformingFunctionSingleAllOf {
	return &NullableCommonBeamformingFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableCommonBeamformingFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonBeamformingFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the V2xCapability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2xCapability{}

// V2xCapability struct for V2xCapability
type V2xCapability struct {
	LteV2x *bool `json:"lteV2x,omitempty"`
	NrV2x *bool `json:"nrV2x,omitempty"`
}

// NewV2xCapability instantiates a new V2xCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2xCapability() *V2xCapability {
	this := V2xCapability{}
	return &this
}

// NewV2xCapabilityWithDefaults instantiates a new V2xCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2xCapabilityWithDefaults() *V2xCapability {
	this := V2xCapability{}
	return &this
}

// GetLteV2x returns the LteV2x field value if set, zero value otherwise.
func (o *V2xCapability) GetLteV2x() bool {
	if o == nil || isNil(o.LteV2x) {
		var ret bool
		return ret
	}
	return *o.LteV2x
}

// GetLteV2xOk returns a tuple with the LteV2x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2xCapability) GetLteV2xOk() (*bool, bool) {
	if o == nil || isNil(o.LteV2x) {
		return nil, false
	}
	return o.LteV2x, true
}

// HasLteV2x returns a boolean if a field has been set.
func (o *V2xCapability) HasLteV2x() bool {
	if o != nil && !isNil(o.LteV2x) {
		return true
	}

	return false
}

// SetLteV2x gets a reference to the given bool and assigns it to the LteV2x field.
func (o *V2xCapability) SetLteV2x(v bool) {
	o.LteV2x = &v
}

// GetNrV2x returns the NrV2x field value if set, zero value otherwise.
func (o *V2xCapability) GetNrV2x() bool {
	if o == nil || isNil(o.NrV2x) {
		var ret bool
		return ret
	}
	return *o.NrV2x
}

// GetNrV2xOk returns a tuple with the NrV2x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2xCapability) GetNrV2xOk() (*bool, bool) {
	if o == nil || isNil(o.NrV2x) {
		return nil, false
	}
	return o.NrV2x, true
}

// HasNrV2x returns a boolean if a field has been set.
func (o *V2xCapability) HasNrV2x() bool {
	if o != nil && !isNil(o.NrV2x) {
		return true
	}

	return false
}

// SetNrV2x gets a reference to the given bool and assigns it to the NrV2x field.
func (o *V2xCapability) SetNrV2x(v bool) {
	o.NrV2x = &v
}

func (o V2xCapability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2xCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LteV2x) {
		toSerialize["lteV2x"] = o.LteV2x
	}
	if !isNil(o.NrV2x) {
		toSerialize["nrV2x"] = o.NrV2x
	}
	return toSerialize, nil
}

type NullableV2xCapability struct {
	value *V2xCapability
	isSet bool
}

func (v NullableV2xCapability) Get() *V2xCapability {
	return v.value
}

func (v *NullableV2xCapability) Set(val *V2xCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableV2xCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableV2xCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2xCapability(val *V2xCapability) *NullableV2xCapability {
	return &NullableV2xCapability{value: val, isSet: true}
}

func (v NullableV2xCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2xCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



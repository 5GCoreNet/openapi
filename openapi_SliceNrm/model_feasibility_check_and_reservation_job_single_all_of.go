/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// checks if the FeasibilityCheckAndReservationJobSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeasibilityCheckAndReservationJobSingleAllOf{}

// FeasibilityCheckAndReservationJobSingleAllOf struct for FeasibilityCheckAndReservationJobSingleAllOf
type FeasibilityCheckAndReservationJobSingleAllOf struct {
	Attributes *FeasibilityCheckAndReservationJobSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewFeasibilityCheckAndReservationJobSingleAllOf instantiates a new FeasibilityCheckAndReservationJobSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeasibilityCheckAndReservationJobSingleAllOf() *FeasibilityCheckAndReservationJobSingleAllOf {
	this := FeasibilityCheckAndReservationJobSingleAllOf{}
	return &this
}

// NewFeasibilityCheckAndReservationJobSingleAllOfWithDefaults instantiates a new FeasibilityCheckAndReservationJobSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeasibilityCheckAndReservationJobSingleAllOfWithDefaults() *FeasibilityCheckAndReservationJobSingleAllOf {
	this := FeasibilityCheckAndReservationJobSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FeasibilityCheckAndReservationJobSingleAllOf) GetAttributes() FeasibilityCheckAndReservationJobSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret FeasibilityCheckAndReservationJobSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOf) GetAttributesOk() (*FeasibilityCheckAndReservationJobSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given FeasibilityCheckAndReservationJobSingleAllOfAttributes and assigns it to the Attributes field.
func (o *FeasibilityCheckAndReservationJobSingleAllOf) SetAttributes(v FeasibilityCheckAndReservationJobSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o FeasibilityCheckAndReservationJobSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeasibilityCheckAndReservationJobSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableFeasibilityCheckAndReservationJobSingleAllOf struct {
	value *FeasibilityCheckAndReservationJobSingleAllOf
	isSet bool
}

func (v NullableFeasibilityCheckAndReservationJobSingleAllOf) Get() *FeasibilityCheckAndReservationJobSingleAllOf {
	return v.value
}

func (v *NullableFeasibilityCheckAndReservationJobSingleAllOf) Set(val *FeasibilityCheckAndReservationJobSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFeasibilityCheckAndReservationJobSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFeasibilityCheckAndReservationJobSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeasibilityCheckAndReservationJobSingleAllOf(val *FeasibilityCheckAndReservationJobSingleAllOf) *NullableFeasibilityCheckAndReservationJobSingleAllOf {
	return &NullableFeasibilityCheckAndReservationJobSingleAllOf{value: val, isSet: true}
}

func (v NullableFeasibilityCheckAndReservationJobSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeasibilityCheckAndReservationJobSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



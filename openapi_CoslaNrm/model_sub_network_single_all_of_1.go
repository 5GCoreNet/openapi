/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// checks if the SubNetworkSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubNetworkSingleAllOf1{}

// SubNetworkSingleAllOf1 struct for SubNetworkSingleAllOf1
type SubNetworkSingleAllOf1 struct {
	AssuranceClosedControlLoop []AssuranceClosedControlLoopSingle `json:"AssuranceClosedControlLoop,omitempty"`
}

// NewSubNetworkSingleAllOf1 instantiates a new SubNetworkSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingleAllOf1() *SubNetworkSingleAllOf1 {
	this := SubNetworkSingleAllOf1{}
	return &this
}

// NewSubNetworkSingleAllOf1WithDefaults instantiates a new SubNetworkSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingleAllOf1WithDefaults() *SubNetworkSingleAllOf1 {
	this := SubNetworkSingleAllOf1{}
	return &this
}

// GetAssuranceClosedControlLoop returns the AssuranceClosedControlLoop field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetAssuranceClosedControlLoop() []AssuranceClosedControlLoopSingle {
	if o == nil || IsNil(o.AssuranceClosedControlLoop) {
		var ret []AssuranceClosedControlLoopSingle
		return ret
	}
	return o.AssuranceClosedControlLoop
}

// GetAssuranceClosedControlLoopOk returns a tuple with the AssuranceClosedControlLoop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetAssuranceClosedControlLoopOk() ([]AssuranceClosedControlLoopSingle, bool) {
	if o == nil || IsNil(o.AssuranceClosedControlLoop) {
		return nil, false
	}
	return o.AssuranceClosedControlLoop, true
}

// HasAssuranceClosedControlLoop returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasAssuranceClosedControlLoop() bool {
	if o != nil && !IsNil(o.AssuranceClosedControlLoop) {
		return true
	}

	return false
}

// SetAssuranceClosedControlLoop gets a reference to the given []AssuranceClosedControlLoopSingle and assigns it to the AssuranceClosedControlLoop field.
func (o *SubNetworkSingleAllOf1) SetAssuranceClosedControlLoop(v []AssuranceClosedControlLoopSingle) {
	o.AssuranceClosedControlLoop = v
}

func (o SubNetworkSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubNetworkSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssuranceClosedControlLoop) {
		toSerialize["AssuranceClosedControlLoop"] = o.AssuranceClosedControlLoop
	}
	return toSerialize, nil
}

type NullableSubNetworkSingleAllOf1 struct {
	value *SubNetworkSingleAllOf1
	isSet bool
}

func (v NullableSubNetworkSingleAllOf1) Get() *SubNetworkSingleAllOf1 {
	return v.value
}

func (v *NullableSubNetworkSingleAllOf1) Set(val *SubNetworkSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingleAllOf1(val *SubNetworkSingleAllOf1) *NullableSubNetworkSingleAllOf1 {
	return &NullableSubNetworkSingleAllOf1{value: val, isSet: true}
}

func (v NullableSubNetworkSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the AssuranceClosedControlLoopSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssuranceClosedControlLoopSingle{}

// AssuranceClosedControlLoopSingle struct for AssuranceClosedControlLoopSingle
type AssuranceClosedControlLoopSingle struct {
	Top
	Attributes    *AssuranceClosedControlLoopSingleAllOfAttributes `json:"attributes,omitempty"`
	AssuranceGoal []AssuranceGoalSingle                            `json:"AssuranceGoal,omitempty"`
}

// NewAssuranceClosedControlLoopSingle instantiates a new AssuranceClosedControlLoopSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssuranceClosedControlLoopSingle(id NullableString) *AssuranceClosedControlLoopSingle {
	this := AssuranceClosedControlLoopSingle{}
	this.Id = id
	return &this
}

// NewAssuranceClosedControlLoopSingleWithDefaults instantiates a new AssuranceClosedControlLoopSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssuranceClosedControlLoopSingleWithDefaults() *AssuranceClosedControlLoopSingle {
	this := AssuranceClosedControlLoopSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingle) GetAttributes() AssuranceClosedControlLoopSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AssuranceClosedControlLoopSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingle) GetAttributesOk() (*AssuranceClosedControlLoopSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AssuranceClosedControlLoopSingleAllOfAttributes and assigns it to the Attributes field.
func (o *AssuranceClosedControlLoopSingle) SetAttributes(v AssuranceClosedControlLoopSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetAssuranceGoal returns the AssuranceGoal field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingle) GetAssuranceGoal() []AssuranceGoalSingle {
	if o == nil || IsNil(o.AssuranceGoal) {
		var ret []AssuranceGoalSingle
		return ret
	}
	return o.AssuranceGoal
}

// GetAssuranceGoalOk returns a tuple with the AssuranceGoal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingle) GetAssuranceGoalOk() ([]AssuranceGoalSingle, bool) {
	if o == nil || IsNil(o.AssuranceGoal) {
		return nil, false
	}
	return o.AssuranceGoal, true
}

// HasAssuranceGoal returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingle) HasAssuranceGoal() bool {
	if o != nil && !IsNil(o.AssuranceGoal) {
		return true
	}

	return false
}

// SetAssuranceGoal gets a reference to the given []AssuranceGoalSingle and assigns it to the AssuranceGoal field.
func (o *AssuranceClosedControlLoopSingle) SetAssuranceGoal(v []AssuranceGoalSingle) {
	o.AssuranceGoal = v
}

func (o AssuranceClosedControlLoopSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssuranceClosedControlLoopSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop, errTop := json.Marshal(o.Top)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	errTop = json.Unmarshal([]byte(serializedTop), &toSerialize)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.AssuranceGoal) {
		toSerialize["AssuranceGoal"] = o.AssuranceGoal
	}
	return toSerialize, nil
}

type NullableAssuranceClosedControlLoopSingle struct {
	value *AssuranceClosedControlLoopSingle
	isSet bool
}

func (v NullableAssuranceClosedControlLoopSingle) Get() *AssuranceClosedControlLoopSingle {
	return v.value
}

func (v *NullableAssuranceClosedControlLoopSingle) Set(val *AssuranceClosedControlLoopSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableAssuranceClosedControlLoopSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableAssuranceClosedControlLoopSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssuranceClosedControlLoopSingle(val *AssuranceClosedControlLoopSingle) *NullableAssuranceClosedControlLoopSingle {
	return &NullableAssuranceClosedControlLoopSingle{value: val, isSet: true}
}

func (v NullableAssuranceClosedControlLoopSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssuranceClosedControlLoopSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

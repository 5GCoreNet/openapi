/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
)

// checks if the NRFqBandContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NRFqBandContext{}

// NRFqBandContext This data type is the \"ObjectContext\" data type with specialisations for NRFqBandContext       
type NRFqBandContext struct {
	ContextAttribute *string `json:"contextAttribute,omitempty"`
	ContextCondition *string `json:"contextCondition,omitempty"`
	ContextValueRange []string `json:"contextValueRange,omitempty"`
}

// NewNRFqBandContext instantiates a new NRFqBandContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNRFqBandContext() *NRFqBandContext {
	this := NRFqBandContext{}
	return &this
}

// NewNRFqBandContextWithDefaults instantiates a new NRFqBandContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNRFqBandContextWithDefaults() *NRFqBandContext {
	this := NRFqBandContext{}
	return &this
}

// GetContextAttribute returns the ContextAttribute field value if set, zero value otherwise.
func (o *NRFqBandContext) GetContextAttribute() string {
	if o == nil || isNil(o.ContextAttribute) {
		var ret string
		return ret
	}
	return *o.ContextAttribute
}

// GetContextAttributeOk returns a tuple with the ContextAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRFqBandContext) GetContextAttributeOk() (*string, bool) {
	if o == nil || isNil(o.ContextAttribute) {
		return nil, false
	}
	return o.ContextAttribute, true
}

// HasContextAttribute returns a boolean if a field has been set.
func (o *NRFqBandContext) HasContextAttribute() bool {
	if o != nil && !isNil(o.ContextAttribute) {
		return true
	}

	return false
}

// SetContextAttribute gets a reference to the given string and assigns it to the ContextAttribute field.
func (o *NRFqBandContext) SetContextAttribute(v string) {
	o.ContextAttribute = &v
}

// GetContextCondition returns the ContextCondition field value if set, zero value otherwise.
func (o *NRFqBandContext) GetContextCondition() string {
	if o == nil || isNil(o.ContextCondition) {
		var ret string
		return ret
	}
	return *o.ContextCondition
}

// GetContextConditionOk returns a tuple with the ContextCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRFqBandContext) GetContextConditionOk() (*string, bool) {
	if o == nil || isNil(o.ContextCondition) {
		return nil, false
	}
	return o.ContextCondition, true
}

// HasContextCondition returns a boolean if a field has been set.
func (o *NRFqBandContext) HasContextCondition() bool {
	if o != nil && !isNil(o.ContextCondition) {
		return true
	}

	return false
}

// SetContextCondition gets a reference to the given string and assigns it to the ContextCondition field.
func (o *NRFqBandContext) SetContextCondition(v string) {
	o.ContextCondition = &v
}

// GetContextValueRange returns the ContextValueRange field value if set, zero value otherwise.
func (o *NRFqBandContext) GetContextValueRange() []string {
	if o == nil || isNil(o.ContextValueRange) {
		var ret []string
		return ret
	}
	return o.ContextValueRange
}

// GetContextValueRangeOk returns a tuple with the ContextValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRFqBandContext) GetContextValueRangeOk() ([]string, bool) {
	if o == nil || isNil(o.ContextValueRange) {
		return nil, false
	}
	return o.ContextValueRange, true
}

// HasContextValueRange returns a boolean if a field has been set.
func (o *NRFqBandContext) HasContextValueRange() bool {
	if o != nil && !isNil(o.ContextValueRange) {
		return true
	}

	return false
}

// SetContextValueRange gets a reference to the given []string and assigns it to the ContextValueRange field.
func (o *NRFqBandContext) SetContextValueRange(v []string) {
	o.ContextValueRange = v
}

func (o NRFqBandContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NRFqBandContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ContextAttribute) {
		toSerialize["contextAttribute"] = o.ContextAttribute
	}
	if !isNil(o.ContextCondition) {
		toSerialize["contextCondition"] = o.ContextCondition
	}
	if !isNil(o.ContextValueRange) {
		toSerialize["contextValueRange"] = o.ContextValueRange
	}
	return toSerialize, nil
}

type NullableNRFqBandContext struct {
	value *NRFqBandContext
	isSet bool
}

func (v NullableNRFqBandContext) Get() *NRFqBandContext {
	return v.value
}

func (v *NullableNRFqBandContext) Set(val *NRFqBandContext) {
	v.value = val
	v.isSet = true
}

func (v NullableNRFqBandContext) IsSet() bool {
	return v.isSet
}

func (v *NullableNRFqBandContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNRFqBandContext(val *NRFqBandContext) *NullableNRFqBandContext {
	return &NullableNRFqBandContext{value: val, isSet: true}
}

func (v NullableNRFqBandContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNRFqBandContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the LowSINRContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LowSINRContext{}

// LowSINRContext This data type is the \"TargetContext\" data type with specialisations for LowSINRContext
type LowSINRContext struct {
	ContextAttribute *string `json:"contextAttribute,omitempty"`
	ContextCondition *string `json:"contextCondition,omitempty"`
	ContextValueRange *int32 `json:"contextValueRange,omitempty"`
}

// NewLowSINRContext instantiates a new LowSINRContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLowSINRContext() *LowSINRContext {
	this := LowSINRContext{}
	return &this
}

// NewLowSINRContextWithDefaults instantiates a new LowSINRContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLowSINRContextWithDefaults() *LowSINRContext {
	this := LowSINRContext{}
	return &this
}

// GetContextAttribute returns the ContextAttribute field value if set, zero value otherwise.
func (o *LowSINRContext) GetContextAttribute() string {
	if o == nil || IsNil(o.ContextAttribute) {
		var ret string
		return ret
	}
	return *o.ContextAttribute
}

// GetContextAttributeOk returns a tuple with the ContextAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowSINRContext) GetContextAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.ContextAttribute) {
		return nil, false
	}
	return o.ContextAttribute, true
}

// HasContextAttribute returns a boolean if a field has been set.
func (o *LowSINRContext) HasContextAttribute() bool {
	if o != nil && !IsNil(o.ContextAttribute) {
		return true
	}

	return false
}

// SetContextAttribute gets a reference to the given string and assigns it to the ContextAttribute field.
func (o *LowSINRContext) SetContextAttribute(v string) {
	o.ContextAttribute = &v
}

// GetContextCondition returns the ContextCondition field value if set, zero value otherwise.
func (o *LowSINRContext) GetContextCondition() string {
	if o == nil || IsNil(o.ContextCondition) {
		var ret string
		return ret
	}
	return *o.ContextCondition
}

// GetContextConditionOk returns a tuple with the ContextCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowSINRContext) GetContextConditionOk() (*string, bool) {
	if o == nil || IsNil(o.ContextCondition) {
		return nil, false
	}
	return o.ContextCondition, true
}

// HasContextCondition returns a boolean if a field has been set.
func (o *LowSINRContext) HasContextCondition() bool {
	if o != nil && !IsNil(o.ContextCondition) {
		return true
	}

	return false
}

// SetContextCondition gets a reference to the given string and assigns it to the ContextCondition field.
func (o *LowSINRContext) SetContextCondition(v string) {
	o.ContextCondition = &v
}

// GetContextValueRange returns the ContextValueRange field value if set, zero value otherwise.
func (o *LowSINRContext) GetContextValueRange() int32 {
	if o == nil || IsNil(o.ContextValueRange) {
		var ret int32
		return ret
	}
	return *o.ContextValueRange
}

// GetContextValueRangeOk returns a tuple with the ContextValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowSINRContext) GetContextValueRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.ContextValueRange) {
		return nil, false
	}
	return o.ContextValueRange, true
}

// HasContextValueRange returns a boolean if a field has been set.
func (o *LowSINRContext) HasContextValueRange() bool {
	if o != nil && !IsNil(o.ContextValueRange) {
		return true
	}

	return false
}

// SetContextValueRange gets a reference to the given int32 and assigns it to the ContextValueRange field.
func (o *LowSINRContext) SetContextValueRange(v int32) {
	o.ContextValueRange = &v
}

func (o LowSINRContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LowSINRContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContextAttribute) {
		toSerialize["contextAttribute"] = o.ContextAttribute
	}
	if !IsNil(o.ContextCondition) {
		toSerialize["contextCondition"] = o.ContextCondition
	}
	if !IsNil(o.ContextValueRange) {
		toSerialize["contextValueRange"] = o.ContextValueRange
	}
	return toSerialize, nil
}

type NullableLowSINRContext struct {
	value *LowSINRContext
	isSet bool
}

func (v NullableLowSINRContext) Get() *LowSINRContext {
	return v.value
}

func (v *NullableLowSINRContext) Set(val *LowSINRContext) {
	v.value = val
	v.isSet = true
}

func (v NullableLowSINRContext) IsSet() bool {
	return v.isSet
}

func (v *NullableLowSINRContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLowSINRContext(val *LowSINRContext) *NullableLowSINRContext {
	return &NullableLowSINRContext{value: val, isSet: true}
}

func (v NullableLowSINRContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLowSINRContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the CoverageAreaTAContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoverageAreaTAContext{}

// CoverageAreaTAContext This data type is the \"ObjectContext\" data type with specialisations for CoverageAreaTAContext       
type CoverageAreaTAContext struct {
	ContextAttribute *string `json:"contextAttribute,omitempty"`
	ContextCondition *string `json:"contextCondition,omitempty"`
	ContextValueRange []int32 `json:"contextValueRange,omitempty"`
}

// NewCoverageAreaTAContext instantiates a new CoverageAreaTAContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoverageAreaTAContext() *CoverageAreaTAContext {
	this := CoverageAreaTAContext{}
	return &this
}

// NewCoverageAreaTAContextWithDefaults instantiates a new CoverageAreaTAContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoverageAreaTAContextWithDefaults() *CoverageAreaTAContext {
	this := CoverageAreaTAContext{}
	return &this
}

// GetContextAttribute returns the ContextAttribute field value if set, zero value otherwise.
func (o *CoverageAreaTAContext) GetContextAttribute() string {
	if o == nil || isNil(o.ContextAttribute) {
		var ret string
		return ret
	}
	return *o.ContextAttribute
}

// GetContextAttributeOk returns a tuple with the ContextAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageAreaTAContext) GetContextAttributeOk() (*string, bool) {
	if o == nil || isNil(o.ContextAttribute) {
		return nil, false
	}
	return o.ContextAttribute, true
}

// HasContextAttribute returns a boolean if a field has been set.
func (o *CoverageAreaTAContext) HasContextAttribute() bool {
	if o != nil && !isNil(o.ContextAttribute) {
		return true
	}

	return false
}

// SetContextAttribute gets a reference to the given string and assigns it to the ContextAttribute field.
func (o *CoverageAreaTAContext) SetContextAttribute(v string) {
	o.ContextAttribute = &v
}

// GetContextCondition returns the ContextCondition field value if set, zero value otherwise.
func (o *CoverageAreaTAContext) GetContextCondition() string {
	if o == nil || isNil(o.ContextCondition) {
		var ret string
		return ret
	}
	return *o.ContextCondition
}

// GetContextConditionOk returns a tuple with the ContextCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageAreaTAContext) GetContextConditionOk() (*string, bool) {
	if o == nil || isNil(o.ContextCondition) {
		return nil, false
	}
	return o.ContextCondition, true
}

// HasContextCondition returns a boolean if a field has been set.
func (o *CoverageAreaTAContext) HasContextCondition() bool {
	if o != nil && !isNil(o.ContextCondition) {
		return true
	}

	return false
}

// SetContextCondition gets a reference to the given string and assigns it to the ContextCondition field.
func (o *CoverageAreaTAContext) SetContextCondition(v string) {
	o.ContextCondition = &v
}

// GetContextValueRange returns the ContextValueRange field value if set, zero value otherwise.
func (o *CoverageAreaTAContext) GetContextValueRange() []int32 {
	if o == nil || isNil(o.ContextValueRange) {
		var ret []int32
		return ret
	}
	return o.ContextValueRange
}

// GetContextValueRangeOk returns a tuple with the ContextValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageAreaTAContext) GetContextValueRangeOk() ([]int32, bool) {
	if o == nil || isNil(o.ContextValueRange) {
		return nil, false
	}
	return o.ContextValueRange, true
}

// HasContextValueRange returns a boolean if a field has been set.
func (o *CoverageAreaTAContext) HasContextValueRange() bool {
	if o != nil && !isNil(o.ContextValueRange) {
		return true
	}

	return false
}

// SetContextValueRange gets a reference to the given []int32 and assigns it to the ContextValueRange field.
func (o *CoverageAreaTAContext) SetContextValueRange(v []int32) {
	o.ContextValueRange = v
}

func (o CoverageAreaTAContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoverageAreaTAContext) ToMap() (map[string]interface{}, error) {
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

type NullableCoverageAreaTAContext struct {
	value *CoverageAreaTAContext
	isSet bool
}

func (v NullableCoverageAreaTAContext) Get() *CoverageAreaTAContext {
	return v.value
}

func (v *NullableCoverageAreaTAContext) Set(val *CoverageAreaTAContext) {
	v.value = val
	v.isSet = true
}

func (v NullableCoverageAreaTAContext) IsSet() bool {
	return v.isSet
}

func (v *NullableCoverageAreaTAContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoverageAreaTAContext(val *CoverageAreaTAContext) *NullableCoverageAreaTAContext {
	return &NullableCoverageAreaTAContext{value: val, isSet: true}
}

func (v NullableCoverageAreaTAContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoverageAreaTAContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



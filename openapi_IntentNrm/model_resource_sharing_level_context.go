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

// checks if the ResourceSharingLevelContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSharingLevelContext{}

// ResourceSharingLevelContext This data type is the \"ExpectationContext\" data type with specialisations for ResourceSharingLevelContext
type ResourceSharingLevelContext struct {
	ContextAttribute  *string        `json:"contextAttribute,omitempty"`
	ContextCondition  *string        `json:"contextCondition,omitempty"`
	ContextValueRange []SharingLevel `json:"contextValueRange,omitempty"`
}

// NewResourceSharingLevelContext instantiates a new ResourceSharingLevelContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSharingLevelContext() *ResourceSharingLevelContext {
	this := ResourceSharingLevelContext{}
	return &this
}

// NewResourceSharingLevelContextWithDefaults instantiates a new ResourceSharingLevelContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSharingLevelContextWithDefaults() *ResourceSharingLevelContext {
	this := ResourceSharingLevelContext{}
	return &this
}

// GetContextAttribute returns the ContextAttribute field value if set, zero value otherwise.
func (o *ResourceSharingLevelContext) GetContextAttribute() string {
	if o == nil || IsNil(o.ContextAttribute) {
		var ret string
		return ret
	}
	return *o.ContextAttribute
}

// GetContextAttributeOk returns a tuple with the ContextAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSharingLevelContext) GetContextAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.ContextAttribute) {
		return nil, false
	}
	return o.ContextAttribute, true
}

// HasContextAttribute returns a boolean if a field has been set.
func (o *ResourceSharingLevelContext) HasContextAttribute() bool {
	if o != nil && !IsNil(o.ContextAttribute) {
		return true
	}

	return false
}

// SetContextAttribute gets a reference to the given string and assigns it to the ContextAttribute field.
func (o *ResourceSharingLevelContext) SetContextAttribute(v string) {
	o.ContextAttribute = &v
}

// GetContextCondition returns the ContextCondition field value if set, zero value otherwise.
func (o *ResourceSharingLevelContext) GetContextCondition() string {
	if o == nil || IsNil(o.ContextCondition) {
		var ret string
		return ret
	}
	return *o.ContextCondition
}

// GetContextConditionOk returns a tuple with the ContextCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSharingLevelContext) GetContextConditionOk() (*string, bool) {
	if o == nil || IsNil(o.ContextCondition) {
		return nil, false
	}
	return o.ContextCondition, true
}

// HasContextCondition returns a boolean if a field has been set.
func (o *ResourceSharingLevelContext) HasContextCondition() bool {
	if o != nil && !IsNil(o.ContextCondition) {
		return true
	}

	return false
}

// SetContextCondition gets a reference to the given string and assigns it to the ContextCondition field.
func (o *ResourceSharingLevelContext) SetContextCondition(v string) {
	o.ContextCondition = &v
}

// GetContextValueRange returns the ContextValueRange field value if set, zero value otherwise.
func (o *ResourceSharingLevelContext) GetContextValueRange() []SharingLevel {
	if o == nil || IsNil(o.ContextValueRange) {
		var ret []SharingLevel
		return ret
	}
	return o.ContextValueRange
}

// GetContextValueRangeOk returns a tuple with the ContextValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSharingLevelContext) GetContextValueRangeOk() ([]SharingLevel, bool) {
	if o == nil || IsNil(o.ContextValueRange) {
		return nil, false
	}
	return o.ContextValueRange, true
}

// HasContextValueRange returns a boolean if a field has been set.
func (o *ResourceSharingLevelContext) HasContextValueRange() bool {
	if o != nil && !IsNil(o.ContextValueRange) {
		return true
	}

	return false
}

// SetContextValueRange gets a reference to the given []SharingLevel and assigns it to the ContextValueRange field.
func (o *ResourceSharingLevelContext) SetContextValueRange(v []SharingLevel) {
	o.ContextValueRange = v
}

func (o ResourceSharingLevelContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSharingLevelContext) ToMap() (map[string]interface{}, error) {
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

type NullableResourceSharingLevelContext struct {
	value *ResourceSharingLevelContext
	isSet bool
}

func (v NullableResourceSharingLevelContext) Get() *ResourceSharingLevelContext {
	return v.value
}

func (v *NullableResourceSharingLevelContext) Set(val *ResourceSharingLevelContext) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSharingLevelContext) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSharingLevelContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSharingLevelContext(val *ResourceSharingLevelContext) *NullableResourceSharingLevelContext {
	return &NullableResourceSharingLevelContext{value: val, isSet: true}
}

func (v NullableResourceSharingLevelContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSharingLevelContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

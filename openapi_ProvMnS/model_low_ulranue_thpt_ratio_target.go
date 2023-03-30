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

// checks if the LowULRANUEThptRatioTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LowULRANUEThptRatioTarget{}

// LowULRANUEThptRatioTarget This data type is the \"ExpectationTarget\" data type with specialisations for LowULRANUEThptRatioTarget        
type LowULRANUEThptRatioTarget struct {
	TargetName *string `json:"targetName,omitempty"`
	TargetCondition *string `json:"targetCondition,omitempty"`
	TargetValueRange *int32 `json:"targetValueRange,omitempty"`
	TargetContexts *LowULRANUEThptContext `json:"targetContexts,omitempty"`
	TargetFulfilmentInfo *FulfilmentInfo `json:"targetFulfilmentInfo,omitempty"`
}

// NewLowULRANUEThptRatioTarget instantiates a new LowULRANUEThptRatioTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLowULRANUEThptRatioTarget() *LowULRANUEThptRatioTarget {
	this := LowULRANUEThptRatioTarget{}
	return &this
}

// NewLowULRANUEThptRatioTargetWithDefaults instantiates a new LowULRANUEThptRatioTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLowULRANUEThptRatioTargetWithDefaults() *LowULRANUEThptRatioTarget {
	this := LowULRANUEThptRatioTarget{}
	return &this
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *LowULRANUEThptRatioTarget) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowULRANUEThptRatioTarget) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *LowULRANUEThptRatioTarget) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *LowULRANUEThptRatioTarget) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetCondition returns the TargetCondition field value if set, zero value otherwise.
func (o *LowULRANUEThptRatioTarget) GetTargetCondition() string {
	if o == nil || IsNil(o.TargetCondition) {
		var ret string
		return ret
	}
	return *o.TargetCondition
}

// GetTargetConditionOk returns a tuple with the TargetCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowULRANUEThptRatioTarget) GetTargetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetCondition) {
		return nil, false
	}
	return o.TargetCondition, true
}

// HasTargetCondition returns a boolean if a field has been set.
func (o *LowULRANUEThptRatioTarget) HasTargetCondition() bool {
	if o != nil && !IsNil(o.TargetCondition) {
		return true
	}

	return false
}

// SetTargetCondition gets a reference to the given string and assigns it to the TargetCondition field.
func (o *LowULRANUEThptRatioTarget) SetTargetCondition(v string) {
	o.TargetCondition = &v
}

// GetTargetValueRange returns the TargetValueRange field value if set, zero value otherwise.
func (o *LowULRANUEThptRatioTarget) GetTargetValueRange() int32 {
	if o == nil || IsNil(o.TargetValueRange) {
		var ret int32
		return ret
	}
	return *o.TargetValueRange
}

// GetTargetValueRangeOk returns a tuple with the TargetValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowULRANUEThptRatioTarget) GetTargetValueRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetValueRange) {
		return nil, false
	}
	return o.TargetValueRange, true
}

// HasTargetValueRange returns a boolean if a field has been set.
func (o *LowULRANUEThptRatioTarget) HasTargetValueRange() bool {
	if o != nil && !IsNil(o.TargetValueRange) {
		return true
	}

	return false
}

// SetTargetValueRange gets a reference to the given int32 and assigns it to the TargetValueRange field.
func (o *LowULRANUEThptRatioTarget) SetTargetValueRange(v int32) {
	o.TargetValueRange = &v
}

// GetTargetContexts returns the TargetContexts field value if set, zero value otherwise.
func (o *LowULRANUEThptRatioTarget) GetTargetContexts() LowULRANUEThptContext {
	if o == nil || IsNil(o.TargetContexts) {
		var ret LowULRANUEThptContext
		return ret
	}
	return *o.TargetContexts
}

// GetTargetContextsOk returns a tuple with the TargetContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowULRANUEThptRatioTarget) GetTargetContextsOk() (*LowULRANUEThptContext, bool) {
	if o == nil || IsNil(o.TargetContexts) {
		return nil, false
	}
	return o.TargetContexts, true
}

// HasTargetContexts returns a boolean if a field has been set.
func (o *LowULRANUEThptRatioTarget) HasTargetContexts() bool {
	if o != nil && !IsNil(o.TargetContexts) {
		return true
	}

	return false
}

// SetTargetContexts gets a reference to the given LowULRANUEThptContext and assigns it to the TargetContexts field.
func (o *LowULRANUEThptRatioTarget) SetTargetContexts(v LowULRANUEThptContext) {
	o.TargetContexts = &v
}

// GetTargetFulfilmentInfo returns the TargetFulfilmentInfo field value if set, zero value otherwise.
func (o *LowULRANUEThptRatioTarget) GetTargetFulfilmentInfo() FulfilmentInfo {
	if o == nil || IsNil(o.TargetFulfilmentInfo) {
		var ret FulfilmentInfo
		return ret
	}
	return *o.TargetFulfilmentInfo
}

// GetTargetFulfilmentInfoOk returns a tuple with the TargetFulfilmentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowULRANUEThptRatioTarget) GetTargetFulfilmentInfoOk() (*FulfilmentInfo, bool) {
	if o == nil || IsNil(o.TargetFulfilmentInfo) {
		return nil, false
	}
	return o.TargetFulfilmentInfo, true
}

// HasTargetFulfilmentInfo returns a boolean if a field has been set.
func (o *LowULRANUEThptRatioTarget) HasTargetFulfilmentInfo() bool {
	if o != nil && !IsNil(o.TargetFulfilmentInfo) {
		return true
	}

	return false
}

// SetTargetFulfilmentInfo gets a reference to the given FulfilmentInfo and assigns it to the TargetFulfilmentInfo field.
func (o *LowULRANUEThptRatioTarget) SetTargetFulfilmentInfo(v FulfilmentInfo) {
	o.TargetFulfilmentInfo = &v
}

func (o LowULRANUEThptRatioTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LowULRANUEThptRatioTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetName) {
		toSerialize["targetName"] = o.TargetName
	}
	if !IsNil(o.TargetCondition) {
		toSerialize["targetCondition"] = o.TargetCondition
	}
	if !IsNil(o.TargetValueRange) {
		toSerialize["targetValueRange"] = o.TargetValueRange
	}
	if !IsNil(o.TargetContexts) {
		toSerialize["targetContexts"] = o.TargetContexts
	}
	if !IsNil(o.TargetFulfilmentInfo) {
		toSerialize["targetFulfilmentInfo"] = o.TargetFulfilmentInfo
	}
	return toSerialize, nil
}

type NullableLowULRANUEThptRatioTarget struct {
	value *LowULRANUEThptRatioTarget
	isSet bool
}

func (v NullableLowULRANUEThptRatioTarget) Get() *LowULRANUEThptRatioTarget {
	return v.value
}

func (v *NullableLowULRANUEThptRatioTarget) Set(val *LowULRANUEThptRatioTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableLowULRANUEThptRatioTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableLowULRANUEThptRatioTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLowULRANUEThptRatioTarget(val *LowULRANUEThptRatioTarget) *NullableLowULRANUEThptRatioTarget {
	return &NullableLowULRANUEThptRatioTarget{value: val, isSet: true}
}

func (v NullableLowULRANUEThptRatioTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLowULRANUEThptRatioTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



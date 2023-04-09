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

// checks if the LowSINRRatioTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LowSINRRatioTarget{}

// LowSINRRatioTarget This data type is the \"ExpectationTarget\" data type with specialisations for LowSINRatioTarget
type LowSINRRatioTarget struct {
	TargetName           *string         `json:"targetName,omitempty"`
	TargetCondition      *string         `json:"targetCondition,omitempty"`
	TargetValueRange     *int32          `json:"targetValueRange,omitempty"`
	TargetContexts       *LowSINRContext `json:"targetContexts,omitempty"`
	TargetFulfilmentInfo *FulfilmentInfo `json:"targetFulfilmentInfo,omitempty"`
}

// NewLowSINRRatioTarget instantiates a new LowSINRRatioTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLowSINRRatioTarget() *LowSINRRatioTarget {
	this := LowSINRRatioTarget{}
	return &this
}

// NewLowSINRRatioTargetWithDefaults instantiates a new LowSINRRatioTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLowSINRRatioTargetWithDefaults() *LowSINRRatioTarget {
	this := LowSINRRatioTarget{}
	return &this
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *LowSINRRatioTarget) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowSINRRatioTarget) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *LowSINRRatioTarget) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *LowSINRRatioTarget) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetCondition returns the TargetCondition field value if set, zero value otherwise.
func (o *LowSINRRatioTarget) GetTargetCondition() string {
	if o == nil || IsNil(o.TargetCondition) {
		var ret string
		return ret
	}
	return *o.TargetCondition
}

// GetTargetConditionOk returns a tuple with the TargetCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowSINRRatioTarget) GetTargetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetCondition) {
		return nil, false
	}
	return o.TargetCondition, true
}

// HasTargetCondition returns a boolean if a field has been set.
func (o *LowSINRRatioTarget) HasTargetCondition() bool {
	if o != nil && !IsNil(o.TargetCondition) {
		return true
	}

	return false
}

// SetTargetCondition gets a reference to the given string and assigns it to the TargetCondition field.
func (o *LowSINRRatioTarget) SetTargetCondition(v string) {
	o.TargetCondition = &v
}

// GetTargetValueRange returns the TargetValueRange field value if set, zero value otherwise.
func (o *LowSINRRatioTarget) GetTargetValueRange() int32 {
	if o == nil || IsNil(o.TargetValueRange) {
		var ret int32
		return ret
	}
	return *o.TargetValueRange
}

// GetTargetValueRangeOk returns a tuple with the TargetValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowSINRRatioTarget) GetTargetValueRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetValueRange) {
		return nil, false
	}
	return o.TargetValueRange, true
}

// HasTargetValueRange returns a boolean if a field has been set.
func (o *LowSINRRatioTarget) HasTargetValueRange() bool {
	if o != nil && !IsNil(o.TargetValueRange) {
		return true
	}

	return false
}

// SetTargetValueRange gets a reference to the given int32 and assigns it to the TargetValueRange field.
func (o *LowSINRRatioTarget) SetTargetValueRange(v int32) {
	o.TargetValueRange = &v
}

// GetTargetContexts returns the TargetContexts field value if set, zero value otherwise.
func (o *LowSINRRatioTarget) GetTargetContexts() LowSINRContext {
	if o == nil || IsNil(o.TargetContexts) {
		var ret LowSINRContext
		return ret
	}
	return *o.TargetContexts
}

// GetTargetContextsOk returns a tuple with the TargetContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowSINRRatioTarget) GetTargetContextsOk() (*LowSINRContext, bool) {
	if o == nil || IsNil(o.TargetContexts) {
		return nil, false
	}
	return o.TargetContexts, true
}

// HasTargetContexts returns a boolean if a field has been set.
func (o *LowSINRRatioTarget) HasTargetContexts() bool {
	if o != nil && !IsNil(o.TargetContexts) {
		return true
	}

	return false
}

// SetTargetContexts gets a reference to the given LowSINRContext and assigns it to the TargetContexts field.
func (o *LowSINRRatioTarget) SetTargetContexts(v LowSINRContext) {
	o.TargetContexts = &v
}

// GetTargetFulfilmentInfo returns the TargetFulfilmentInfo field value if set, zero value otherwise.
func (o *LowSINRRatioTarget) GetTargetFulfilmentInfo() FulfilmentInfo {
	if o == nil || IsNil(o.TargetFulfilmentInfo) {
		var ret FulfilmentInfo
		return ret
	}
	return *o.TargetFulfilmentInfo
}

// GetTargetFulfilmentInfoOk returns a tuple with the TargetFulfilmentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowSINRRatioTarget) GetTargetFulfilmentInfoOk() (*FulfilmentInfo, bool) {
	if o == nil || IsNil(o.TargetFulfilmentInfo) {
		return nil, false
	}
	return o.TargetFulfilmentInfo, true
}

// HasTargetFulfilmentInfo returns a boolean if a field has been set.
func (o *LowSINRRatioTarget) HasTargetFulfilmentInfo() bool {
	if o != nil && !IsNil(o.TargetFulfilmentInfo) {
		return true
	}

	return false
}

// SetTargetFulfilmentInfo gets a reference to the given FulfilmentInfo and assigns it to the TargetFulfilmentInfo field.
func (o *LowSINRRatioTarget) SetTargetFulfilmentInfo(v FulfilmentInfo) {
	o.TargetFulfilmentInfo = &v
}

func (o LowSINRRatioTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LowSINRRatioTarget) ToMap() (map[string]interface{}, error) {
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

type NullableLowSINRRatioTarget struct {
	value *LowSINRRatioTarget
	isSet bool
}

func (v NullableLowSINRRatioTarget) Get() *LowSINRRatioTarget {
	return v.value
}

func (v *NullableLowSINRRatioTarget) Set(val *LowSINRRatioTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableLowSINRRatioTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableLowSINRRatioTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLowSINRRatioTarget(val *LowSINRRatioTarget) *NullableLowSINRRatioTarget {
	return &NullableLowSINRRatioTarget{value: val, isSet: true}
}

func (v NullableLowSINRRatioTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLowSINRRatioTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

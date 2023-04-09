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

// checks if the AnalyticsScheduleOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsScheduleOneOf1{}

// AnalyticsScheduleOneOf1 struct for AnalyticsScheduleOneOf1
type AnalyticsScheduleOneOf1 struct {
	GranularityPeriod *int32 `json:"granularityPeriod,omitempty"`
}

// NewAnalyticsScheduleOneOf1 instantiates a new AnalyticsScheduleOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsScheduleOneOf1() *AnalyticsScheduleOneOf1 {
	this := AnalyticsScheduleOneOf1{}
	return &this
}

// NewAnalyticsScheduleOneOf1WithDefaults instantiates a new AnalyticsScheduleOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsScheduleOneOf1WithDefaults() *AnalyticsScheduleOneOf1 {
	this := AnalyticsScheduleOneOf1{}
	return &this
}

// GetGranularityPeriod returns the GranularityPeriod field value if set, zero value otherwise.
func (o *AnalyticsScheduleOneOf1) GetGranularityPeriod() int32 {
	if o == nil || IsNil(o.GranularityPeriod) {
		var ret int32
		return ret
	}
	return *o.GranularityPeriod
}

// GetGranularityPeriodOk returns a tuple with the GranularityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsScheduleOneOf1) GetGranularityPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.GranularityPeriod) {
		return nil, false
	}
	return o.GranularityPeriod, true
}

// HasGranularityPeriod returns a boolean if a field has been set.
func (o *AnalyticsScheduleOneOf1) HasGranularityPeriod() bool {
	if o != nil && !IsNil(o.GranularityPeriod) {
		return true
	}

	return false
}

// SetGranularityPeriod gets a reference to the given int32 and assigns it to the GranularityPeriod field.
func (o *AnalyticsScheduleOneOf1) SetGranularityPeriod(v int32) {
	o.GranularityPeriod = &v
}

func (o AnalyticsScheduleOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsScheduleOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GranularityPeriod) {
		toSerialize["granularityPeriod"] = o.GranularityPeriod
	}
	return toSerialize, nil
}

type NullableAnalyticsScheduleOneOf1 struct {
	value *AnalyticsScheduleOneOf1
	isSet bool
}

func (v NullableAnalyticsScheduleOneOf1) Get() *AnalyticsScheduleOneOf1 {
	return v.value
}

func (v *NullableAnalyticsScheduleOneOf1) Set(val *AnalyticsScheduleOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsScheduleOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsScheduleOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsScheduleOneOf1(val *AnalyticsScheduleOneOf1) *NullableAnalyticsScheduleOneOf1 {
	return &NullableAnalyticsScheduleOneOf1{value: val, isSet: true}
}

func (v NullableAnalyticsScheduleOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsScheduleOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

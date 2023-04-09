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

// checks if the AssuranceReportSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssuranceReportSingleAllOfAttributesAllOf{}

// AssuranceReportSingleAllOfAttributesAllOf struct for AssuranceReportSingleAllOfAttributesAllOf
type AssuranceReportSingleAllOfAttributesAllOf struct {
	AssuranceGoalStatusList []AssuranceGoalStatus `json:"assuranceGoalStatusList,omitempty"`
}

// NewAssuranceReportSingleAllOfAttributesAllOf instantiates a new AssuranceReportSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssuranceReportSingleAllOfAttributesAllOf() *AssuranceReportSingleAllOfAttributesAllOf {
	this := AssuranceReportSingleAllOfAttributesAllOf{}
	return &this
}

// NewAssuranceReportSingleAllOfAttributesAllOfWithDefaults instantiates a new AssuranceReportSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssuranceReportSingleAllOfAttributesAllOfWithDefaults() *AssuranceReportSingleAllOfAttributesAllOf {
	this := AssuranceReportSingleAllOfAttributesAllOf{}
	return &this
}

// GetAssuranceGoalStatusList returns the AssuranceGoalStatusList field value if set, zero value otherwise.
func (o *AssuranceReportSingleAllOfAttributesAllOf) GetAssuranceGoalStatusList() []AssuranceGoalStatus {
	if o == nil || IsNil(o.AssuranceGoalStatusList) {
		var ret []AssuranceGoalStatus
		return ret
	}
	return o.AssuranceGoalStatusList
}

// GetAssuranceGoalStatusListOk returns a tuple with the AssuranceGoalStatusList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceReportSingleAllOfAttributesAllOf) GetAssuranceGoalStatusListOk() ([]AssuranceGoalStatus, bool) {
	if o == nil || IsNil(o.AssuranceGoalStatusList) {
		return nil, false
	}
	return o.AssuranceGoalStatusList, true
}

// HasAssuranceGoalStatusList returns a boolean if a field has been set.
func (o *AssuranceReportSingleAllOfAttributesAllOf) HasAssuranceGoalStatusList() bool {
	if o != nil && !IsNil(o.AssuranceGoalStatusList) {
		return true
	}

	return false
}

// SetAssuranceGoalStatusList gets a reference to the given []AssuranceGoalStatus and assigns it to the AssuranceGoalStatusList field.
func (o *AssuranceReportSingleAllOfAttributesAllOf) SetAssuranceGoalStatusList(v []AssuranceGoalStatus) {
	o.AssuranceGoalStatusList = v
}

func (o AssuranceReportSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssuranceReportSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssuranceGoalStatusList) {
		toSerialize["assuranceGoalStatusList"] = o.AssuranceGoalStatusList
	}
	return toSerialize, nil
}

type NullableAssuranceReportSingleAllOfAttributesAllOf struct {
	value *AssuranceReportSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableAssuranceReportSingleAllOfAttributesAllOf) Get() *AssuranceReportSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableAssuranceReportSingleAllOfAttributesAllOf) Set(val *AssuranceReportSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssuranceReportSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssuranceReportSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssuranceReportSingleAllOfAttributesAllOf(val *AssuranceReportSingleAllOfAttributesAllOf) *NullableAssuranceReportSingleAllOfAttributesAllOf {
	return &NullableAssuranceReportSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableAssuranceReportSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssuranceReportSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

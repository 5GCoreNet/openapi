/*
M1_EventDataProcessingProvisioning

5GMS AF M1 Event Data Processing Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_EventDataProcessingProvisioning

import (
	"encoding/json"
)

// checks if the DataAccessProfileTimeAccessRestrictions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataAccessProfileTimeAccessRestrictions{}

// DataAccessProfileTimeAccessRestrictions struct for DataAccessProfileTimeAccessRestrictions
type DataAccessProfileTimeAccessRestrictions struct {
	// indicating a time in seconds.
	Duration int32 `json:"duration"`
	AggregationFunctions []DataAggregationFunctionType `json:"aggregationFunctions"`
}

// NewDataAccessProfileTimeAccessRestrictions instantiates a new DataAccessProfileTimeAccessRestrictions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataAccessProfileTimeAccessRestrictions(duration int32, aggregationFunctions []DataAggregationFunctionType) *DataAccessProfileTimeAccessRestrictions {
	this := DataAccessProfileTimeAccessRestrictions{}
	this.Duration = duration
	this.AggregationFunctions = aggregationFunctions
	return &this
}

// NewDataAccessProfileTimeAccessRestrictionsWithDefaults instantiates a new DataAccessProfileTimeAccessRestrictions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataAccessProfileTimeAccessRestrictionsWithDefaults() *DataAccessProfileTimeAccessRestrictions {
	this := DataAccessProfileTimeAccessRestrictions{}
	return &this
}

// GetDuration returns the Duration field value
func (o *DataAccessProfileTimeAccessRestrictions) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *DataAccessProfileTimeAccessRestrictions) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *DataAccessProfileTimeAccessRestrictions) SetDuration(v int32) {
	o.Duration = v
}

// GetAggregationFunctions returns the AggregationFunctions field value
func (o *DataAccessProfileTimeAccessRestrictions) GetAggregationFunctions() []DataAggregationFunctionType {
	if o == nil {
		var ret []DataAggregationFunctionType
		return ret
	}

	return o.AggregationFunctions
}

// GetAggregationFunctionsOk returns a tuple with the AggregationFunctions field value
// and a boolean to check if the value has been set.
func (o *DataAccessProfileTimeAccessRestrictions) GetAggregationFunctionsOk() ([]DataAggregationFunctionType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AggregationFunctions, true
}

// SetAggregationFunctions sets field value
func (o *DataAccessProfileTimeAccessRestrictions) SetAggregationFunctions(v []DataAggregationFunctionType) {
	o.AggregationFunctions = v
}

func (o DataAccessProfileTimeAccessRestrictions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataAccessProfileTimeAccessRestrictions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration
	toSerialize["aggregationFunctions"] = o.AggregationFunctions
	return toSerialize, nil
}

type NullableDataAccessProfileTimeAccessRestrictions struct {
	value *DataAccessProfileTimeAccessRestrictions
	isSet bool
}

func (v NullableDataAccessProfileTimeAccessRestrictions) Get() *DataAccessProfileTimeAccessRestrictions {
	return v.value
}

func (v *NullableDataAccessProfileTimeAccessRestrictions) Set(val *DataAccessProfileTimeAccessRestrictions) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAccessProfileTimeAccessRestrictions) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAccessProfileTimeAccessRestrictions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAccessProfileTimeAccessRestrictions(val *DataAccessProfileTimeAccessRestrictions) *NullableDataAccessProfileTimeAccessRestrictions {
	return &NullableDataAccessProfileTimeAccessRestrictions{value: val, isSet: true}
}

func (v NullableDataAccessProfileTimeAccessRestrictions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAccessProfileTimeAccessRestrictions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


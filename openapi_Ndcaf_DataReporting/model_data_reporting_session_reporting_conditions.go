/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReporting

import (
	"encoding/json"
)

// checks if the DataReportingSessionReportingConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataReportingSessionReportingConditions{}

// DataReportingSessionReportingConditions struct for DataReportingSessionReportingConditions
type DataReportingSessionReportingConditions struct {
	Default DataDomain `json:"default"`
}

// NewDataReportingSessionReportingConditions instantiates a new DataReportingSessionReportingConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataReportingSessionReportingConditions(default_ DataDomain) *DataReportingSessionReportingConditions {
	this := DataReportingSessionReportingConditions{}
	this.Default = default_
	return &this
}

// NewDataReportingSessionReportingConditionsWithDefaults instantiates a new DataReportingSessionReportingConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataReportingSessionReportingConditionsWithDefaults() *DataReportingSessionReportingConditions {
	this := DataReportingSessionReportingConditions{}
	return &this
}

// GetDefault returns the Default field value
func (o *DataReportingSessionReportingConditions) GetDefault() DataDomain {
	if o == nil {
		var ret DataDomain
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *DataReportingSessionReportingConditions) GetDefaultOk() (*DataDomain, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *DataReportingSessionReportingConditions) SetDefault(v DataDomain) {
	o.Default = v
}

func (o DataReportingSessionReportingConditions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataReportingSessionReportingConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["default"] = o.Default
	return toSerialize, nil
}

type NullableDataReportingSessionReportingConditions struct {
	value *DataReportingSessionReportingConditions
	isSet bool
}

func (v NullableDataReportingSessionReportingConditions) Get() *DataReportingSessionReportingConditions {
	return v.value
}

func (v *NullableDataReportingSessionReportingConditions) Set(val *DataReportingSessionReportingConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableDataReportingSessionReportingConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableDataReportingSessionReportingConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataReportingSessionReportingConditions(val *DataReportingSessionReportingConditions) *NullableDataReportingSessionReportingConditions {
	return &NullableDataReportingSessionReportingConditions{value: val, isSet: true}
}

func (v NullableDataReportingSessionReportingConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataReportingSessionReportingConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the LogicalInterfaceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogicalInterfaceInfo{}

// LogicalInterfaceInfo struct for LogicalInterfaceInfo
type LogicalInterfaceInfo struct {
	LogicalInterfaceType *string `json:"logicalInterfaceType,omitempty"`
	LogicalInterfaceId   *string `json:"logicalInterfaceId,omitempty"`
}

// NewLogicalInterfaceInfo instantiates a new LogicalInterfaceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogicalInterfaceInfo() *LogicalInterfaceInfo {
	this := LogicalInterfaceInfo{}
	return &this
}

// NewLogicalInterfaceInfoWithDefaults instantiates a new LogicalInterfaceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogicalInterfaceInfoWithDefaults() *LogicalInterfaceInfo {
	this := LogicalInterfaceInfo{}
	return &this
}

// GetLogicalInterfaceType returns the LogicalInterfaceType field value if set, zero value otherwise.
func (o *LogicalInterfaceInfo) GetLogicalInterfaceType() string {
	if o == nil || IsNil(o.LogicalInterfaceType) {
		var ret string
		return ret
	}
	return *o.LogicalInterfaceType
}

// GetLogicalInterfaceTypeOk returns a tuple with the LogicalInterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogicalInterfaceInfo) GetLogicalInterfaceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LogicalInterfaceType) {
		return nil, false
	}
	return o.LogicalInterfaceType, true
}

// HasLogicalInterfaceType returns a boolean if a field has been set.
func (o *LogicalInterfaceInfo) HasLogicalInterfaceType() bool {
	if o != nil && !IsNil(o.LogicalInterfaceType) {
		return true
	}

	return false
}

// SetLogicalInterfaceType gets a reference to the given string and assigns it to the LogicalInterfaceType field.
func (o *LogicalInterfaceInfo) SetLogicalInterfaceType(v string) {
	o.LogicalInterfaceType = &v
}

// GetLogicalInterfaceId returns the LogicalInterfaceId field value if set, zero value otherwise.
func (o *LogicalInterfaceInfo) GetLogicalInterfaceId() string {
	if o == nil || IsNil(o.LogicalInterfaceId) {
		var ret string
		return ret
	}
	return *o.LogicalInterfaceId
}

// GetLogicalInterfaceIdOk returns a tuple with the LogicalInterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogicalInterfaceInfo) GetLogicalInterfaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.LogicalInterfaceId) {
		return nil, false
	}
	return o.LogicalInterfaceId, true
}

// HasLogicalInterfaceId returns a boolean if a field has been set.
func (o *LogicalInterfaceInfo) HasLogicalInterfaceId() bool {
	if o != nil && !IsNil(o.LogicalInterfaceId) {
		return true
	}

	return false
}

// SetLogicalInterfaceId gets a reference to the given string and assigns it to the LogicalInterfaceId field.
func (o *LogicalInterfaceInfo) SetLogicalInterfaceId(v string) {
	o.LogicalInterfaceId = &v
}

func (o LogicalInterfaceInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogicalInterfaceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LogicalInterfaceType) {
		toSerialize["logicalInterfaceType"] = o.LogicalInterfaceType
	}
	if !IsNil(o.LogicalInterfaceId) {
		toSerialize["logicalInterfaceId"] = o.LogicalInterfaceId
	}
	return toSerialize, nil
}

type NullableLogicalInterfaceInfo struct {
	value *LogicalInterfaceInfo
	isSet bool
}

func (v NullableLogicalInterfaceInfo) Get() *LogicalInterfaceInfo {
	return v.value
}

func (v *NullableLogicalInterfaceInfo) Set(val *LogicalInterfaceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLogicalInterfaceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLogicalInterfaceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogicalInterfaceInfo(val *LogicalInterfaceInfo) *NullableLogicalInterfaceInfo {
	return &NullableLogicalInterfaceInfo{value: val, isSet: true}
}

func (v NullableLogicalInterfaceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogicalInterfaceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

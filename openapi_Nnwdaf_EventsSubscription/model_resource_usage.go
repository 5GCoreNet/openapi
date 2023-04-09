/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
)

// checks if the ResourceUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceUsage{}

// ResourceUsage The current usage of the virtual resources assigned to the NF instances belonging to a  particular network slice instance.
type ResourceUsage struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	CpuUsage *int32 `json:"cpuUsage,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MemoryUsage *int32 `json:"memoryUsage,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	StorageUsage *int32 `json:"storageUsage,omitempty"`
}

// NewResourceUsage instantiates a new ResourceUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceUsage() *ResourceUsage {
	this := ResourceUsage{}
	return &this
}

// NewResourceUsageWithDefaults instantiates a new ResourceUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceUsageWithDefaults() *ResourceUsage {
	this := ResourceUsage{}
	return &this
}

// GetCpuUsage returns the CpuUsage field value if set, zero value otherwise.
func (o *ResourceUsage) GetCpuUsage() int32 {
	if o == nil || IsNil(o.CpuUsage) {
		var ret int32
		return ret
	}
	return *o.CpuUsage
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceUsage) GetCpuUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.CpuUsage) {
		return nil, false
	}
	return o.CpuUsage, true
}

// HasCpuUsage returns a boolean if a field has been set.
func (o *ResourceUsage) HasCpuUsage() bool {
	if o != nil && !IsNil(o.CpuUsage) {
		return true
	}

	return false
}

// SetCpuUsage gets a reference to the given int32 and assigns it to the CpuUsage field.
func (o *ResourceUsage) SetCpuUsage(v int32) {
	o.CpuUsage = &v
}

// GetMemoryUsage returns the MemoryUsage field value if set, zero value otherwise.
func (o *ResourceUsage) GetMemoryUsage() int32 {
	if o == nil || IsNil(o.MemoryUsage) {
		var ret int32
		return ret
	}
	return *o.MemoryUsage
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceUsage) GetMemoryUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.MemoryUsage) {
		return nil, false
	}
	return o.MemoryUsage, true
}

// HasMemoryUsage returns a boolean if a field has been set.
func (o *ResourceUsage) HasMemoryUsage() bool {
	if o != nil && !IsNil(o.MemoryUsage) {
		return true
	}

	return false
}

// SetMemoryUsage gets a reference to the given int32 and assigns it to the MemoryUsage field.
func (o *ResourceUsage) SetMemoryUsage(v int32) {
	o.MemoryUsage = &v
}

// GetStorageUsage returns the StorageUsage field value if set, zero value otherwise.
func (o *ResourceUsage) GetStorageUsage() int32 {
	if o == nil || IsNil(o.StorageUsage) {
		var ret int32
		return ret
	}
	return *o.StorageUsage
}

// GetStorageUsageOk returns a tuple with the StorageUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceUsage) GetStorageUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.StorageUsage) {
		return nil, false
	}
	return o.StorageUsage, true
}

// HasStorageUsage returns a boolean if a field has been set.
func (o *ResourceUsage) HasStorageUsage() bool {
	if o != nil && !IsNil(o.StorageUsage) {
		return true
	}

	return false
}

// SetStorageUsage gets a reference to the given int32 and assigns it to the StorageUsage field.
func (o *ResourceUsage) SetStorageUsage(v int32) {
	o.StorageUsage = &v
}

func (o ResourceUsage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpuUsage) {
		toSerialize["cpuUsage"] = o.CpuUsage
	}
	if !IsNil(o.MemoryUsage) {
		toSerialize["memoryUsage"] = o.MemoryUsage
	}
	if !IsNil(o.StorageUsage) {
		toSerialize["storageUsage"] = o.StorageUsage
	}
	return toSerialize, nil
}

type NullableResourceUsage struct {
	value *ResourceUsage
	isSet bool
}

func (v NullableResourceUsage) Get() *ResourceUsage {
	return v.value
}

func (v *NullableResourceUsage) Set(val *ResourceUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceUsage(val *ResourceUsage) *NullableResourceUsage {
	return &NullableResourceUsage{value: val, isSet: true}
}

func (v NullableResourceUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

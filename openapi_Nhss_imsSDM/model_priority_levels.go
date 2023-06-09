/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the PriorityLevels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriorityLevels{}

// PriorityLevels Namespaces and priority levels allowed for the IMS public Identity
type PriorityLevels struct {
	ServicePriorityLevelList []string `json:"servicePriorityLevelList"`
	ServicePriorityLevel     *int32   `json:"servicePriorityLevel,omitempty"`
}

// NewPriorityLevels instantiates a new PriorityLevels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriorityLevels(servicePriorityLevelList []string) *PriorityLevels {
	this := PriorityLevels{}
	this.ServicePriorityLevelList = servicePriorityLevelList
	return &this
}

// NewPriorityLevelsWithDefaults instantiates a new PriorityLevels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriorityLevelsWithDefaults() *PriorityLevels {
	this := PriorityLevels{}
	return &this
}

// GetServicePriorityLevelList returns the ServicePriorityLevelList field value
func (o *PriorityLevels) GetServicePriorityLevelList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ServicePriorityLevelList
}

// GetServicePriorityLevelListOk returns a tuple with the ServicePriorityLevelList field value
// and a boolean to check if the value has been set.
func (o *PriorityLevels) GetServicePriorityLevelListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServicePriorityLevelList, true
}

// SetServicePriorityLevelList sets field value
func (o *PriorityLevels) SetServicePriorityLevelList(v []string) {
	o.ServicePriorityLevelList = v
}

// GetServicePriorityLevel returns the ServicePriorityLevel field value if set, zero value otherwise.
func (o *PriorityLevels) GetServicePriorityLevel() int32 {
	if o == nil || IsNil(o.ServicePriorityLevel) {
		var ret int32
		return ret
	}
	return *o.ServicePriorityLevel
}

// GetServicePriorityLevelOk returns a tuple with the ServicePriorityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriorityLevels) GetServicePriorityLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.ServicePriorityLevel) {
		return nil, false
	}
	return o.ServicePriorityLevel, true
}

// HasServicePriorityLevel returns a boolean if a field has been set.
func (o *PriorityLevels) HasServicePriorityLevel() bool {
	if o != nil && !IsNil(o.ServicePriorityLevel) {
		return true
	}

	return false
}

// SetServicePriorityLevel gets a reference to the given int32 and assigns it to the ServicePriorityLevel field.
func (o *PriorityLevels) SetServicePriorityLevel(v int32) {
	o.ServicePriorityLevel = &v
}

func (o PriorityLevels) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriorityLevels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["servicePriorityLevelList"] = o.ServicePriorityLevelList
	if !IsNil(o.ServicePriorityLevel) {
		toSerialize["servicePriorityLevel"] = o.ServicePriorityLevel
	}
	return toSerialize, nil
}

type NullablePriorityLevels struct {
	value *PriorityLevels
	isSet bool
}

func (v NullablePriorityLevels) Get() *PriorityLevels {
	return v.value
}

func (v *NullablePriorityLevels) Set(val *PriorityLevels) {
	v.value = val
	v.isSet = true
}

func (v NullablePriorityLevels) IsSet() bool {
	return v.isSet
}

func (v *NullablePriorityLevels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriorityLevels(val *PriorityLevels) *NullablePriorityLevels {
	return &NullablePriorityLevels{value: val, isSet: true}
}

func (v NullablePriorityLevels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriorityLevels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

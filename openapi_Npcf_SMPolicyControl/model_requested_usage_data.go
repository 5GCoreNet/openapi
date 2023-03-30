/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
)

// checks if the RequestedUsageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestedUsageData{}

// RequestedUsageData Contains usage data requested by the PCF requesting usage reports for the corresponding  usage monitoring data instances. 
type RequestedUsageData struct {
	// An array of usage monitoring data id references to the usage monitoring data instances for which the PCF is requesting a usage report. This attribute shall only be provided when allUmIds is not set to true. 
	RefUmIds []string `json:"refUmIds,omitempty"`
	// This boolean indicates whether requested usage data applies to all usage monitoring data instances. When it's not included, it means requested usage data shall only apply to the usage monitoring data instances referenced by the refUmIds attribute. 
	AllUmIds *bool `json:"allUmIds,omitempty"`
}

// NewRequestedUsageData instantiates a new RequestedUsageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestedUsageData() *RequestedUsageData {
	this := RequestedUsageData{}
	return &this
}

// NewRequestedUsageDataWithDefaults instantiates a new RequestedUsageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestedUsageDataWithDefaults() *RequestedUsageData {
	this := RequestedUsageData{}
	return &this
}

// GetRefUmIds returns the RefUmIds field value if set, zero value otherwise.
func (o *RequestedUsageData) GetRefUmIds() []string {
	if o == nil || IsNil(o.RefUmIds) {
		var ret []string
		return ret
	}
	return o.RefUmIds
}

// GetRefUmIdsOk returns a tuple with the RefUmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedUsageData) GetRefUmIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RefUmIds) {
		return nil, false
	}
	return o.RefUmIds, true
}

// HasRefUmIds returns a boolean if a field has been set.
func (o *RequestedUsageData) HasRefUmIds() bool {
	if o != nil && !IsNil(o.RefUmIds) {
		return true
	}

	return false
}

// SetRefUmIds gets a reference to the given []string and assigns it to the RefUmIds field.
func (o *RequestedUsageData) SetRefUmIds(v []string) {
	o.RefUmIds = v
}

// GetAllUmIds returns the AllUmIds field value if set, zero value otherwise.
func (o *RequestedUsageData) GetAllUmIds() bool {
	if o == nil || IsNil(o.AllUmIds) {
		var ret bool
		return ret
	}
	return *o.AllUmIds
}

// GetAllUmIdsOk returns a tuple with the AllUmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedUsageData) GetAllUmIdsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllUmIds) {
		return nil, false
	}
	return o.AllUmIds, true
}

// HasAllUmIds returns a boolean if a field has been set.
func (o *RequestedUsageData) HasAllUmIds() bool {
	if o != nil && !IsNil(o.AllUmIds) {
		return true
	}

	return false
}

// SetAllUmIds gets a reference to the given bool and assigns it to the AllUmIds field.
func (o *RequestedUsageData) SetAllUmIds(v bool) {
	o.AllUmIds = &v
}

func (o RequestedUsageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestedUsageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RefUmIds) {
		toSerialize["refUmIds"] = o.RefUmIds
	}
	if !IsNil(o.AllUmIds) {
		toSerialize["allUmIds"] = o.AllUmIds
	}
	return toSerialize, nil
}

type NullableRequestedUsageData struct {
	value *RequestedUsageData
	isSet bool
}

func (v NullableRequestedUsageData) Get() *RequestedUsageData {
	return v.value
}

func (v *NullableRequestedUsageData) Set(val *RequestedUsageData) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedUsageData) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedUsageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedUsageData(val *RequestedUsageData) *NullableRequestedUsageData {
	return &NullableRequestedUsageData{value: val, isSet: true}
}

func (v NullableRequestedUsageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestedUsageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



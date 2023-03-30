/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the NwdafCapability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NwdafCapability{}

// NwdafCapability Indicates the capability supported by the NWDAF
type NwdafCapability struct {
	AnalyticsAggregation *bool `json:"analyticsAggregation,omitempty"`
	AnalyticsMetadataProvisioning *bool `json:"analyticsMetadataProvisioning,omitempty"`
}

// NewNwdafCapability instantiates a new NwdafCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwdafCapability() *NwdafCapability {
	this := NwdafCapability{}
	var analyticsAggregation bool = false
	this.AnalyticsAggregation = &analyticsAggregation
	var analyticsMetadataProvisioning bool = false
	this.AnalyticsMetadataProvisioning = &analyticsMetadataProvisioning
	return &this
}

// NewNwdafCapabilityWithDefaults instantiates a new NwdafCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwdafCapabilityWithDefaults() *NwdafCapability {
	this := NwdafCapability{}
	var analyticsAggregation bool = false
	this.AnalyticsAggregation = &analyticsAggregation
	var analyticsMetadataProvisioning bool = false
	this.AnalyticsMetadataProvisioning = &analyticsMetadataProvisioning
	return &this
}

// GetAnalyticsAggregation returns the AnalyticsAggregation field value if set, zero value otherwise.
func (o *NwdafCapability) GetAnalyticsAggregation() bool {
	if o == nil || IsNil(o.AnalyticsAggregation) {
		var ret bool
		return ret
	}
	return *o.AnalyticsAggregation
}

// GetAnalyticsAggregationOk returns a tuple with the AnalyticsAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCapability) GetAnalyticsAggregationOk() (*bool, bool) {
	if o == nil || IsNil(o.AnalyticsAggregation) {
		return nil, false
	}
	return o.AnalyticsAggregation, true
}

// HasAnalyticsAggregation returns a boolean if a field has been set.
func (o *NwdafCapability) HasAnalyticsAggregation() bool {
	if o != nil && !IsNil(o.AnalyticsAggregation) {
		return true
	}

	return false
}

// SetAnalyticsAggregation gets a reference to the given bool and assigns it to the AnalyticsAggregation field.
func (o *NwdafCapability) SetAnalyticsAggregation(v bool) {
	o.AnalyticsAggregation = &v
}

// GetAnalyticsMetadataProvisioning returns the AnalyticsMetadataProvisioning field value if set, zero value otherwise.
func (o *NwdafCapability) GetAnalyticsMetadataProvisioning() bool {
	if o == nil || IsNil(o.AnalyticsMetadataProvisioning) {
		var ret bool
		return ret
	}
	return *o.AnalyticsMetadataProvisioning
}

// GetAnalyticsMetadataProvisioningOk returns a tuple with the AnalyticsMetadataProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCapability) GetAnalyticsMetadataProvisioningOk() (*bool, bool) {
	if o == nil || IsNil(o.AnalyticsMetadataProvisioning) {
		return nil, false
	}
	return o.AnalyticsMetadataProvisioning, true
}

// HasAnalyticsMetadataProvisioning returns a boolean if a field has been set.
func (o *NwdafCapability) HasAnalyticsMetadataProvisioning() bool {
	if o != nil && !IsNil(o.AnalyticsMetadataProvisioning) {
		return true
	}

	return false
}

// SetAnalyticsMetadataProvisioning gets a reference to the given bool and assigns it to the AnalyticsMetadataProvisioning field.
func (o *NwdafCapability) SetAnalyticsMetadataProvisioning(v bool) {
	o.AnalyticsMetadataProvisioning = &v
}

func (o NwdafCapability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NwdafCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnalyticsAggregation) {
		toSerialize["analyticsAggregation"] = o.AnalyticsAggregation
	}
	if !IsNil(o.AnalyticsMetadataProvisioning) {
		toSerialize["analyticsMetadataProvisioning"] = o.AnalyticsMetadataProvisioning
	}
	return toSerialize, nil
}

type NullableNwdafCapability struct {
	value *NwdafCapability
	isSet bool
}

func (v NullableNwdafCapability) Get() *NwdafCapability {
	return v.value
}

func (v *NullableNwdafCapability) Set(val *NwdafCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafCapability(val *NwdafCapability) *NullableNwdafCapability {
	return &NullableNwdafCapability{value: val, isSet: true}
}

func (v NullableNwdafCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



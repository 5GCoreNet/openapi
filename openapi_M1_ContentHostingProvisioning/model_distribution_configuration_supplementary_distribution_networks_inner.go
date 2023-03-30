/*
M1_ContentHostingProvisioning

5GMS AF M1 Content Hosting Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_ContentHostingProvisioning

import (
	"encoding/json"
)

// checks if the DistributionConfigurationSupplementaryDistributionNetworksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistributionConfigurationSupplementaryDistributionNetworksInner{}

// DistributionConfigurationSupplementaryDistributionNetworksInner A duple tying a type of supplementary distribution network to its distribution mode.
type DistributionConfigurationSupplementaryDistributionNetworksInner struct {
	DistributionNetworkType DistributionNetworkType `json:"distributionNetworkType"`
	DistributionMode DistributionMode `json:"distributionMode"`
}

// NewDistributionConfigurationSupplementaryDistributionNetworksInner instantiates a new DistributionConfigurationSupplementaryDistributionNetworksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionConfigurationSupplementaryDistributionNetworksInner(distributionNetworkType DistributionNetworkType, distributionMode DistributionMode) *DistributionConfigurationSupplementaryDistributionNetworksInner {
	this := DistributionConfigurationSupplementaryDistributionNetworksInner{}
	this.DistributionNetworkType = distributionNetworkType
	this.DistributionMode = distributionMode
	return &this
}

// NewDistributionConfigurationSupplementaryDistributionNetworksInnerWithDefaults instantiates a new DistributionConfigurationSupplementaryDistributionNetworksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionConfigurationSupplementaryDistributionNetworksInnerWithDefaults() *DistributionConfigurationSupplementaryDistributionNetworksInner {
	this := DistributionConfigurationSupplementaryDistributionNetworksInner{}
	return &this
}

// GetDistributionNetworkType returns the DistributionNetworkType field value
func (o *DistributionConfigurationSupplementaryDistributionNetworksInner) GetDistributionNetworkType() DistributionNetworkType {
	if o == nil {
		var ret DistributionNetworkType
		return ret
	}

	return o.DistributionNetworkType
}

// GetDistributionNetworkTypeOk returns a tuple with the DistributionNetworkType field value
// and a boolean to check if the value has been set.
func (o *DistributionConfigurationSupplementaryDistributionNetworksInner) GetDistributionNetworkTypeOk() (*DistributionNetworkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistributionNetworkType, true
}

// SetDistributionNetworkType sets field value
func (o *DistributionConfigurationSupplementaryDistributionNetworksInner) SetDistributionNetworkType(v DistributionNetworkType) {
	o.DistributionNetworkType = v
}

// GetDistributionMode returns the DistributionMode field value
func (o *DistributionConfigurationSupplementaryDistributionNetworksInner) GetDistributionMode() DistributionMode {
	if o == nil {
		var ret DistributionMode
		return ret
	}

	return o.DistributionMode
}

// GetDistributionModeOk returns a tuple with the DistributionMode field value
// and a boolean to check if the value has been set.
func (o *DistributionConfigurationSupplementaryDistributionNetworksInner) GetDistributionModeOk() (*DistributionMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistributionMode, true
}

// SetDistributionMode sets field value
func (o *DistributionConfigurationSupplementaryDistributionNetworksInner) SetDistributionMode(v DistributionMode) {
	o.DistributionMode = v
}

func (o DistributionConfigurationSupplementaryDistributionNetworksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistributionConfigurationSupplementaryDistributionNetworksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["distributionNetworkType"] = o.DistributionNetworkType
	toSerialize["distributionMode"] = o.DistributionMode
	return toSerialize, nil
}

type NullableDistributionConfigurationSupplementaryDistributionNetworksInner struct {
	value *DistributionConfigurationSupplementaryDistributionNetworksInner
	isSet bool
}

func (v NullableDistributionConfigurationSupplementaryDistributionNetworksInner) Get() *DistributionConfigurationSupplementaryDistributionNetworksInner {
	return v.value
}

func (v *NullableDistributionConfigurationSupplementaryDistributionNetworksInner) Set(val *DistributionConfigurationSupplementaryDistributionNetworksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionConfigurationSupplementaryDistributionNetworksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionConfigurationSupplementaryDistributionNetworksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionConfigurationSupplementaryDistributionNetworksInner(val *DistributionConfigurationSupplementaryDistributionNetworksInner) *NullableDistributionConfigurationSupplementaryDistributionNetworksInner {
	return &NullableDistributionConfigurationSupplementaryDistributionNetworksInner{value: val, isSet: true}
}

func (v NullableDistributionConfigurationSupplementaryDistributionNetworksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionConfigurationSupplementaryDistributionNetworksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



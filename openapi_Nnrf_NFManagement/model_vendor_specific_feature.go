/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
)

// checks if the VendorSpecificFeature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VendorSpecificFeature{}

// VendorSpecificFeature Information about a vendor-specific feature
type VendorSpecificFeature struct {
	FeatureName string `json:"featureName"`
	FeatureVersion string `json:"featureVersion"`
}

// NewVendorSpecificFeature instantiates a new VendorSpecificFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorSpecificFeature(featureName string, featureVersion string) *VendorSpecificFeature {
	this := VendorSpecificFeature{}
	this.FeatureName = featureName
	this.FeatureVersion = featureVersion
	return &this
}

// NewVendorSpecificFeatureWithDefaults instantiates a new VendorSpecificFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorSpecificFeatureWithDefaults() *VendorSpecificFeature {
	this := VendorSpecificFeature{}
	return &this
}

// GetFeatureName returns the FeatureName field value
func (o *VendorSpecificFeature) GetFeatureName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureName
}

// GetFeatureNameOk returns a tuple with the FeatureName field value
// and a boolean to check if the value has been set.
func (o *VendorSpecificFeature) GetFeatureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureName, true
}

// SetFeatureName sets field value
func (o *VendorSpecificFeature) SetFeatureName(v string) {
	o.FeatureName = v
}

// GetFeatureVersion returns the FeatureVersion field value
func (o *VendorSpecificFeature) GetFeatureVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureVersion
}

// GetFeatureVersionOk returns a tuple with the FeatureVersion field value
// and a boolean to check if the value has been set.
func (o *VendorSpecificFeature) GetFeatureVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureVersion, true
}

// SetFeatureVersion sets field value
func (o *VendorSpecificFeature) SetFeatureVersion(v string) {
	o.FeatureVersion = v
}

func (o VendorSpecificFeature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VendorSpecificFeature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["featureName"] = o.FeatureName
	toSerialize["featureVersion"] = o.FeatureVersion
	return toSerialize, nil
}

type NullableVendorSpecificFeature struct {
	value *VendorSpecificFeature
	isSet bool
}

func (v NullableVendorSpecificFeature) Get() *VendorSpecificFeature {
	return v.value
}

func (v *NullableVendorSpecificFeature) Set(val *VendorSpecificFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorSpecificFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorSpecificFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorSpecificFeature(val *VendorSpecificFeature) *NullableVendorSpecificFeature {
	return &NullableVendorSpecificFeature{value: val, isSet: true}
}

func (v NullableVendorSpecificFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorSpecificFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the SubNetworkSingle4AllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubNetworkSingle4AllOf{}

// SubNetworkSingle4AllOf struct for SubNetworkSingle4AllOf
type SubNetworkSingle4AllOf struct {
	SubNetwork []SubNetworkSingle `json:"SubNetwork,omitempty"`
	Intent []IntentSingle `json:"Intent,omitempty"`
}

// NewSubNetworkSingle4AllOf instantiates a new SubNetworkSingle4AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingle4AllOf() *SubNetworkSingle4AllOf {
	this := SubNetworkSingle4AllOf{}
	return &this
}

// NewSubNetworkSingle4AllOfWithDefaults instantiates a new SubNetworkSingle4AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingle4AllOfWithDefaults() *SubNetworkSingle4AllOf {
	this := SubNetworkSingle4AllOf{}
	return &this
}

// GetSubNetwork returns the SubNetwork field value if set, zero value otherwise.
func (o *SubNetworkSingle4AllOf) GetSubNetwork() []SubNetworkSingle {
	if o == nil || IsNil(o.SubNetwork) {
		var ret []SubNetworkSingle
		return ret
	}
	return o.SubNetwork
}

// GetSubNetworkOk returns a tuple with the SubNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle4AllOf) GetSubNetworkOk() ([]SubNetworkSingle, bool) {
	if o == nil || IsNil(o.SubNetwork) {
		return nil, false
	}
	return o.SubNetwork, true
}

// HasSubNetwork returns a boolean if a field has been set.
func (o *SubNetworkSingle4AllOf) HasSubNetwork() bool {
	if o != nil && !IsNil(o.SubNetwork) {
		return true
	}

	return false
}

// SetSubNetwork gets a reference to the given []SubNetworkSingle and assigns it to the SubNetwork field.
func (o *SubNetworkSingle4AllOf) SetSubNetwork(v []SubNetworkSingle) {
	o.SubNetwork = v
}

// GetIntent returns the Intent field value if set, zero value otherwise.
func (o *SubNetworkSingle4AllOf) GetIntent() []IntentSingle {
	if o == nil || IsNil(o.Intent) {
		var ret []IntentSingle
		return ret
	}
	return o.Intent
}

// GetIntentOk returns a tuple with the Intent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle4AllOf) GetIntentOk() ([]IntentSingle, bool) {
	if o == nil || IsNil(o.Intent) {
		return nil, false
	}
	return o.Intent, true
}

// HasIntent returns a boolean if a field has been set.
func (o *SubNetworkSingle4AllOf) HasIntent() bool {
	if o != nil && !IsNil(o.Intent) {
		return true
	}

	return false
}

// SetIntent gets a reference to the given []IntentSingle and assigns it to the Intent field.
func (o *SubNetworkSingle4AllOf) SetIntent(v []IntentSingle) {
	o.Intent = v
}

func (o SubNetworkSingle4AllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubNetworkSingle4AllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubNetwork) {
		toSerialize["SubNetwork"] = o.SubNetwork
	}
	if !IsNil(o.Intent) {
		toSerialize["Intent"] = o.Intent
	}
	return toSerialize, nil
}

type NullableSubNetworkSingle4AllOf struct {
	value *SubNetworkSingle4AllOf
	isSet bool
}

func (v NullableSubNetworkSingle4AllOf) Get() *SubNetworkSingle4AllOf {
	return v.value
}

func (v *NullableSubNetworkSingle4AllOf) Set(val *SubNetworkSingle4AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingle4AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingle4AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingle4AllOf(val *SubNetworkSingle4AllOf) *NullableSubNetworkSingle4AllOf {
	return &NullableSubNetworkSingle4AllOf{value: val, isSet: true}
}

func (v NullableSubNetworkSingle4AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingle4AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



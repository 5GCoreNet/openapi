/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
)

// checks if the SubNetworkSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubNetworkSingleAllOf1{}

// SubNetworkSingleAllOf1 struct for SubNetworkSingleAllOf1
type SubNetworkSingleAllOf1 struct {
	SubNetwork []SubNetworkSingle `json:"SubNetwork,omitempty"`
	Intent []IntentSingle `json:"Intent,omitempty"`
}

// NewSubNetworkSingleAllOf1 instantiates a new SubNetworkSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingleAllOf1() *SubNetworkSingleAllOf1 {
	this := SubNetworkSingleAllOf1{}
	return &this
}

// NewSubNetworkSingleAllOf1WithDefaults instantiates a new SubNetworkSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingleAllOf1WithDefaults() *SubNetworkSingleAllOf1 {
	this := SubNetworkSingleAllOf1{}
	return &this
}

// GetSubNetwork returns the SubNetwork field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetSubNetwork() []SubNetworkSingle {
	if o == nil || IsNil(o.SubNetwork) {
		var ret []SubNetworkSingle
		return ret
	}
	return o.SubNetwork
}

// GetSubNetworkOk returns a tuple with the SubNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetSubNetworkOk() ([]SubNetworkSingle, bool) {
	if o == nil || IsNil(o.SubNetwork) {
		return nil, false
	}
	return o.SubNetwork, true
}

// HasSubNetwork returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasSubNetwork() bool {
	if o != nil && !IsNil(o.SubNetwork) {
		return true
	}

	return false
}

// SetSubNetwork gets a reference to the given []SubNetworkSingle and assigns it to the SubNetwork field.
func (o *SubNetworkSingleAllOf1) SetSubNetwork(v []SubNetworkSingle) {
	o.SubNetwork = v
}

// GetIntent returns the Intent field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetIntent() []IntentSingle {
	if o == nil || IsNil(o.Intent) {
		var ret []IntentSingle
		return ret
	}
	return o.Intent
}

// GetIntentOk returns a tuple with the Intent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetIntentOk() ([]IntentSingle, bool) {
	if o == nil || IsNil(o.Intent) {
		return nil, false
	}
	return o.Intent, true
}

// HasIntent returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasIntent() bool {
	if o != nil && !IsNil(o.Intent) {
		return true
	}

	return false
}

// SetIntent gets a reference to the given []IntentSingle and assigns it to the Intent field.
func (o *SubNetworkSingleAllOf1) SetIntent(v []IntentSingle) {
	o.Intent = v
}

func (o SubNetworkSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubNetworkSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubNetwork) {
		toSerialize["SubNetwork"] = o.SubNetwork
	}
	if !IsNil(o.Intent) {
		toSerialize["Intent"] = o.Intent
	}
	return toSerialize, nil
}

type NullableSubNetworkSingleAllOf1 struct {
	value *SubNetworkSingleAllOf1
	isSet bool
}

func (v NullableSubNetworkSingleAllOf1) Get() *SubNetworkSingleAllOf1 {
	return v.value
}

func (v *NullableSubNetworkSingleAllOf1) Set(val *SubNetworkSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingleAllOf1(val *SubNetworkSingleAllOf1) *NullableSubNetworkSingleAllOf1 {
	return &NullableSubNetworkSingleAllOf1{value: val, isSet: true}
}

func (v NullableSubNetworkSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


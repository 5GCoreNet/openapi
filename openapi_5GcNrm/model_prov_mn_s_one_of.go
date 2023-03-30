/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the ProvMnSOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvMnSOneOf{}

// ProvMnSOneOf struct for ProvMnSOneOf
type ProvMnSOneOf struct {
	SubNetwork []SubNetworkSingle `json:"SubNetwork,omitempty"`
}

// NewProvMnSOneOf instantiates a new ProvMnSOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvMnSOneOf() *ProvMnSOneOf {
	this := ProvMnSOneOf{}
	return &this
}

// NewProvMnSOneOfWithDefaults instantiates a new ProvMnSOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvMnSOneOfWithDefaults() *ProvMnSOneOf {
	this := ProvMnSOneOf{}
	return &this
}

// GetSubNetwork returns the SubNetwork field value if set, zero value otherwise.
func (o *ProvMnSOneOf) GetSubNetwork() []SubNetworkSingle {
	if o == nil || IsNil(o.SubNetwork) {
		var ret []SubNetworkSingle
		return ret
	}
	return o.SubNetwork
}

// GetSubNetworkOk returns a tuple with the SubNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvMnSOneOf) GetSubNetworkOk() ([]SubNetworkSingle, bool) {
	if o == nil || IsNil(o.SubNetwork) {
		return nil, false
	}
	return o.SubNetwork, true
}

// HasSubNetwork returns a boolean if a field has been set.
func (o *ProvMnSOneOf) HasSubNetwork() bool {
	if o != nil && !IsNil(o.SubNetwork) {
		return true
	}

	return false
}

// SetSubNetwork gets a reference to the given []SubNetworkSingle and assigns it to the SubNetwork field.
func (o *ProvMnSOneOf) SetSubNetwork(v []SubNetworkSingle) {
	o.SubNetwork = v
}

func (o ProvMnSOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvMnSOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubNetwork) {
		toSerialize["SubNetwork"] = o.SubNetwork
	}
	return toSerialize, nil
}

type NullableProvMnSOneOf struct {
	value *ProvMnSOneOf
	isSet bool
}

func (v NullableProvMnSOneOf) Get() *ProvMnSOneOf {
	return v.value
}

func (v *NullableProvMnSOneOf) Set(val *ProvMnSOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProvMnSOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProvMnSOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvMnSOneOf(val *ProvMnSOneOf) *NullableProvMnSOneOf {
	return &NullableProvMnSOneOf{value: val, isSet: true}
}

func (v NullableProvMnSOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvMnSOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


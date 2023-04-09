/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// checks if the SubNetworkSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubNetworkSingleAllOfAttributes{}

// SubNetworkSingleAllOfAttributes struct for SubNetworkSingleAllOfAttributes
type SubNetworkSingleAllOfAttributes struct {
	SubNetworkAttr
}

// NewSubNetworkSingleAllOfAttributes instantiates a new SubNetworkSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingleAllOfAttributes() *SubNetworkSingleAllOfAttributes {
	this := SubNetworkSingleAllOfAttributes{}
	return &this
}

// NewSubNetworkSingleAllOfAttributesWithDefaults instantiates a new SubNetworkSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingleAllOfAttributesWithDefaults() *SubNetworkSingleAllOfAttributes {
	this := SubNetworkSingleAllOfAttributes{}
	return &this
}

func (o SubNetworkSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubNetworkSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSubNetworkAttr, errSubNetworkAttr := json.Marshal(o.SubNetworkAttr)
	if errSubNetworkAttr != nil {
		return map[string]interface{}{}, errSubNetworkAttr
	}
	errSubNetworkAttr = json.Unmarshal([]byte(serializedSubNetworkAttr), &toSerialize)
	if errSubNetworkAttr != nil {
		return map[string]interface{}{}, errSubNetworkAttr
	}
	return toSerialize, nil
}

type NullableSubNetworkSingleAllOfAttributes struct {
	value *SubNetworkSingleAllOfAttributes
	isSet bool
}

func (v NullableSubNetworkSingleAllOfAttributes) Get() *SubNetworkSingleAllOfAttributes {
	return v.value
}

func (v *NullableSubNetworkSingleAllOfAttributes) Set(val *SubNetworkSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingleAllOfAttributes(val *SubNetworkSingleAllOfAttributes) *NullableSubNetworkSingleAllOfAttributes {
	return &NullableSubNetworkSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableSubNetworkSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

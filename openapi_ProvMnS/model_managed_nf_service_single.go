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

// checks if the ManagedNFServiceSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedNFServiceSingle{}

// ManagedNFServiceSingle struct for ManagedNFServiceSingle
type ManagedNFServiceSingle struct {
	Top
	Attributes *ManagedNFServiceSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewManagedNFServiceSingle instantiates a new ManagedNFServiceSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedNFServiceSingle(id NullableString) *ManagedNFServiceSingle {
	this := ManagedNFServiceSingle{}
	this.Id = id
	return &this
}

// NewManagedNFServiceSingleWithDefaults instantiates a new ManagedNFServiceSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedNFServiceSingleWithDefaults() *ManagedNFServiceSingle {
	this := ManagedNFServiceSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagedNFServiceSingle) GetAttributes() ManagedNFServiceSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ManagedNFServiceSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNFServiceSingle) GetAttributesOk() (*ManagedNFServiceSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagedNFServiceSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedNFServiceSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ManagedNFServiceSingle) SetAttributes(v ManagedNFServiceSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o ManagedNFServiceSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedNFServiceSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop, errTop := json.Marshal(o.Top)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	errTop = json.Unmarshal([]byte(serializedTop), &toSerialize)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableManagedNFServiceSingle struct {
	value *ManagedNFServiceSingle
	isSet bool
}

func (v NullableManagedNFServiceSingle) Get() *ManagedNFServiceSingle {
	return v.value
}

func (v *NullableManagedNFServiceSingle) Set(val *ManagedNFServiceSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedNFServiceSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedNFServiceSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedNFServiceSingle(val *ManagedNFServiceSingle) *NullableManagedNFServiceSingle {
	return &NullableManagedNFServiceSingle{value: val, isSet: true}
}

func (v NullableManagedNFServiceSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedNFServiceSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

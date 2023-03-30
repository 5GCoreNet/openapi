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

// checks if the Configurable5QISetSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Configurable5QISetSingleAllOfAttributesAllOf{}

// Configurable5QISetSingleAllOfAttributesAllOf struct for Configurable5QISetSingleAllOfAttributesAllOf
type Configurable5QISetSingleAllOfAttributesAllOf struct {
	Configurable5QIs []FiveQICharacteristicsSingle `json:"configurable5QIs,omitempty"`
}

// NewConfigurable5QISetSingleAllOfAttributesAllOf instantiates a new Configurable5QISetSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurable5QISetSingleAllOfAttributesAllOf() *Configurable5QISetSingleAllOfAttributesAllOf {
	this := Configurable5QISetSingleAllOfAttributesAllOf{}
	return &this
}

// NewConfigurable5QISetSingleAllOfAttributesAllOfWithDefaults instantiates a new Configurable5QISetSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurable5QISetSingleAllOfAttributesAllOfWithDefaults() *Configurable5QISetSingleAllOfAttributesAllOf {
	this := Configurable5QISetSingleAllOfAttributesAllOf{}
	return &this
}

// GetConfigurable5QIs returns the Configurable5QIs field value if set, zero value otherwise.
func (o *Configurable5QISetSingleAllOfAttributesAllOf) GetConfigurable5QIs() []FiveQICharacteristicsSingle {
	if o == nil || IsNil(o.Configurable5QIs) {
		var ret []FiveQICharacteristicsSingle
		return ret
	}
	return o.Configurable5QIs
}

// GetConfigurable5QIsOk returns a tuple with the Configurable5QIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configurable5QISetSingleAllOfAttributesAllOf) GetConfigurable5QIsOk() ([]FiveQICharacteristicsSingle, bool) {
	if o == nil || IsNil(o.Configurable5QIs) {
		return nil, false
	}
	return o.Configurable5QIs, true
}

// HasConfigurable5QIs returns a boolean if a field has been set.
func (o *Configurable5QISetSingleAllOfAttributesAllOf) HasConfigurable5QIs() bool {
	if o != nil && !IsNil(o.Configurable5QIs) {
		return true
	}

	return false
}

// SetConfigurable5QIs gets a reference to the given []FiveQICharacteristicsSingle and assigns it to the Configurable5QIs field.
func (o *Configurable5QISetSingleAllOfAttributesAllOf) SetConfigurable5QIs(v []FiveQICharacteristicsSingle) {
	o.Configurable5QIs = v
}

func (o Configurable5QISetSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Configurable5QISetSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configurable5QIs) {
		toSerialize["configurable5QIs"] = o.Configurable5QIs
	}
	return toSerialize, nil
}

type NullableConfigurable5QISetSingleAllOfAttributesAllOf struct {
	value *Configurable5QISetSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableConfigurable5QISetSingleAllOfAttributesAllOf) Get() *Configurable5QISetSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableConfigurable5QISetSingleAllOfAttributesAllOf) Set(val *Configurable5QISetSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurable5QISetSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurable5QISetSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurable5QISetSingleAllOfAttributesAllOf(val *Configurable5QISetSingleAllOfAttributesAllOf) *NullableConfigurable5QISetSingleAllOfAttributesAllOf {
	return &NullableConfigurable5QISetSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableConfigurable5QISetSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurable5QISetSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



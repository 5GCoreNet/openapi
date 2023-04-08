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

// checks if the DANRManagementFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DANRManagementFunctionSingleAllOfAttributes{}

// DANRManagementFunctionSingleAllOfAttributes struct for DANRManagementFunctionSingleAllOfAttributes
type DANRManagementFunctionSingleAllOfAttributes struct {
	IntrasystemANRManagementSwitch *bool `json:"intrasystemANRManagementSwitch,omitempty"`
	IntersystemANRManagementSwitch *bool `json:"intersystemANRManagementSwitch,omitempty"`
}

// NewDANRManagementFunctionSingleAllOfAttributes instantiates a new DANRManagementFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDANRManagementFunctionSingleAllOfAttributes() *DANRManagementFunctionSingleAllOfAttributes {
	this := DANRManagementFunctionSingleAllOfAttributes{}
	return &this
}

// NewDANRManagementFunctionSingleAllOfAttributesWithDefaults instantiates a new DANRManagementFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDANRManagementFunctionSingleAllOfAttributesWithDefaults() *DANRManagementFunctionSingleAllOfAttributes {
	this := DANRManagementFunctionSingleAllOfAttributes{}
	return &this
}

// GetIntrasystemANRManagementSwitch returns the IntrasystemANRManagementSwitch field value if set, zero value otherwise.
func (o *DANRManagementFunctionSingleAllOfAttributes) GetIntrasystemANRManagementSwitch() bool {
	if o == nil || isNil(o.IntrasystemANRManagementSwitch) {
		var ret bool
		return ret
	}
	return *o.IntrasystemANRManagementSwitch
}

// GetIntrasystemANRManagementSwitchOk returns a tuple with the IntrasystemANRManagementSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DANRManagementFunctionSingleAllOfAttributes) GetIntrasystemANRManagementSwitchOk() (*bool, bool) {
	if o == nil || isNil(o.IntrasystemANRManagementSwitch) {
		return nil, false
	}
	return o.IntrasystemANRManagementSwitch, true
}

// HasIntrasystemANRManagementSwitch returns a boolean if a field has been set.
func (o *DANRManagementFunctionSingleAllOfAttributes) HasIntrasystemANRManagementSwitch() bool {
	if o != nil && !isNil(o.IntrasystemANRManagementSwitch) {
		return true
	}

	return false
}

// SetIntrasystemANRManagementSwitch gets a reference to the given bool and assigns it to the IntrasystemANRManagementSwitch field.
func (o *DANRManagementFunctionSingleAllOfAttributes) SetIntrasystemANRManagementSwitch(v bool) {
	o.IntrasystemANRManagementSwitch = &v
}

// GetIntersystemANRManagementSwitch returns the IntersystemANRManagementSwitch field value if set, zero value otherwise.
func (o *DANRManagementFunctionSingleAllOfAttributes) GetIntersystemANRManagementSwitch() bool {
	if o == nil || isNil(o.IntersystemANRManagementSwitch) {
		var ret bool
		return ret
	}
	return *o.IntersystemANRManagementSwitch
}

// GetIntersystemANRManagementSwitchOk returns a tuple with the IntersystemANRManagementSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DANRManagementFunctionSingleAllOfAttributes) GetIntersystemANRManagementSwitchOk() (*bool, bool) {
	if o == nil || isNil(o.IntersystemANRManagementSwitch) {
		return nil, false
	}
	return o.IntersystemANRManagementSwitch, true
}

// HasIntersystemANRManagementSwitch returns a boolean if a field has been set.
func (o *DANRManagementFunctionSingleAllOfAttributes) HasIntersystemANRManagementSwitch() bool {
	if o != nil && !isNil(o.IntersystemANRManagementSwitch) {
		return true
	}

	return false
}

// SetIntersystemANRManagementSwitch gets a reference to the given bool and assigns it to the IntersystemANRManagementSwitch field.
func (o *DANRManagementFunctionSingleAllOfAttributes) SetIntersystemANRManagementSwitch(v bool) {
	o.IntersystemANRManagementSwitch = &v
}

func (o DANRManagementFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DANRManagementFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IntrasystemANRManagementSwitch) {
		toSerialize["intrasystemANRManagementSwitch"] = o.IntrasystemANRManagementSwitch
	}
	if !isNil(o.IntersystemANRManagementSwitch) {
		toSerialize["intersystemANRManagementSwitch"] = o.IntersystemANRManagementSwitch
	}
	return toSerialize, nil
}

type NullableDANRManagementFunctionSingleAllOfAttributes struct {
	value *DANRManagementFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableDANRManagementFunctionSingleAllOfAttributes) Get() *DANRManagementFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableDANRManagementFunctionSingleAllOfAttributes) Set(val *DANRManagementFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDANRManagementFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDANRManagementFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDANRManagementFunctionSingleAllOfAttributes(val *DANRManagementFunctionSingleAllOfAttributes) *NullableDANRManagementFunctionSingleAllOfAttributes {
	return &NullableDANRManagementFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableDANRManagementFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDANRManagementFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



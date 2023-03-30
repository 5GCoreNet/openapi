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

// checks if the ExternalGnbCuUpFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalGnbCuUpFunctionSingleAllOf{}

// ExternalGnbCuUpFunctionSingleAllOf struct for ExternalGnbCuUpFunctionSingleAllOf
type ExternalGnbCuUpFunctionSingleAllOf struct {
	EPE1 []EPE1Single `json:"EP_E1,omitempty"`
	EPF1U []EPF1USingle `json:"EP_F1U,omitempty"`
	EPXnU []EPXnUSingle `json:"EP_XnU,omitempty"`
}

// NewExternalGnbCuUpFunctionSingleAllOf instantiates a new ExternalGnbCuUpFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGnbCuUpFunctionSingleAllOf() *ExternalGnbCuUpFunctionSingleAllOf {
	this := ExternalGnbCuUpFunctionSingleAllOf{}
	return &this
}

// NewExternalGnbCuUpFunctionSingleAllOfWithDefaults instantiates a new ExternalGnbCuUpFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGnbCuUpFunctionSingleAllOfWithDefaults() *ExternalGnbCuUpFunctionSingleAllOf {
	this := ExternalGnbCuUpFunctionSingleAllOf{}
	return &this
}

// GetEPE1 returns the EPE1 field value if set, zero value otherwise.
func (o *ExternalGnbCuUpFunctionSingleAllOf) GetEPE1() []EPE1Single {
	if o == nil || IsNil(o.EPE1) {
		var ret []EPE1Single
		return ret
	}
	return o.EPE1
}

// GetEPE1Ok returns a tuple with the EPE1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuUpFunctionSingleAllOf) GetEPE1Ok() ([]EPE1Single, bool) {
	if o == nil || IsNil(o.EPE1) {
		return nil, false
	}
	return o.EPE1, true
}

// HasEPE1 returns a boolean if a field has been set.
func (o *ExternalGnbCuUpFunctionSingleAllOf) HasEPE1() bool {
	if o != nil && !IsNil(o.EPE1) {
		return true
	}

	return false
}

// SetEPE1 gets a reference to the given []EPE1Single and assigns it to the EPE1 field.
func (o *ExternalGnbCuUpFunctionSingleAllOf) SetEPE1(v []EPE1Single) {
	o.EPE1 = v
}

// GetEPF1U returns the EPF1U field value if set, zero value otherwise.
func (o *ExternalGnbCuUpFunctionSingleAllOf) GetEPF1U() []EPF1USingle {
	if o == nil || IsNil(o.EPF1U) {
		var ret []EPF1USingle
		return ret
	}
	return o.EPF1U
}

// GetEPF1UOk returns a tuple with the EPF1U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuUpFunctionSingleAllOf) GetEPF1UOk() ([]EPF1USingle, bool) {
	if o == nil || IsNil(o.EPF1U) {
		return nil, false
	}
	return o.EPF1U, true
}

// HasEPF1U returns a boolean if a field has been set.
func (o *ExternalGnbCuUpFunctionSingleAllOf) HasEPF1U() bool {
	if o != nil && !IsNil(o.EPF1U) {
		return true
	}

	return false
}

// SetEPF1U gets a reference to the given []EPF1USingle and assigns it to the EPF1U field.
func (o *ExternalGnbCuUpFunctionSingleAllOf) SetEPF1U(v []EPF1USingle) {
	o.EPF1U = v
}

// GetEPXnU returns the EPXnU field value if set, zero value otherwise.
func (o *ExternalGnbCuUpFunctionSingleAllOf) GetEPXnU() []EPXnUSingle {
	if o == nil || IsNil(o.EPXnU) {
		var ret []EPXnUSingle
		return ret
	}
	return o.EPXnU
}

// GetEPXnUOk returns a tuple with the EPXnU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuUpFunctionSingleAllOf) GetEPXnUOk() ([]EPXnUSingle, bool) {
	if o == nil || IsNil(o.EPXnU) {
		return nil, false
	}
	return o.EPXnU, true
}

// HasEPXnU returns a boolean if a field has been set.
func (o *ExternalGnbCuUpFunctionSingleAllOf) HasEPXnU() bool {
	if o != nil && !IsNil(o.EPXnU) {
		return true
	}

	return false
}

// SetEPXnU gets a reference to the given []EPXnUSingle and assigns it to the EPXnU field.
func (o *ExternalGnbCuUpFunctionSingleAllOf) SetEPXnU(v []EPXnUSingle) {
	o.EPXnU = v
}

func (o ExternalGnbCuUpFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalGnbCuUpFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EPE1) {
		toSerialize["EP_E1"] = o.EPE1
	}
	if !IsNil(o.EPF1U) {
		toSerialize["EP_F1U"] = o.EPF1U
	}
	if !IsNil(o.EPXnU) {
		toSerialize["EP_XnU"] = o.EPXnU
	}
	return toSerialize, nil
}

type NullableExternalGnbCuUpFunctionSingleAllOf struct {
	value *ExternalGnbCuUpFunctionSingleAllOf
	isSet bool
}

func (v NullableExternalGnbCuUpFunctionSingleAllOf) Get() *ExternalGnbCuUpFunctionSingleAllOf {
	return v.value
}

func (v *NullableExternalGnbCuUpFunctionSingleAllOf) Set(val *ExternalGnbCuUpFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGnbCuUpFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGnbCuUpFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGnbCuUpFunctionSingleAllOf(val *ExternalGnbCuUpFunctionSingleAllOf) *NullableExternalGnbCuUpFunctionSingleAllOf {
	return &NullableExternalGnbCuUpFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableExternalGnbCuUpFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGnbCuUpFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



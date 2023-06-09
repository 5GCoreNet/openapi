/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the CPCIConfigurationFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CPCIConfigurationFunctionSingleAllOfAttributes{}

// CPCIConfigurationFunctionSingleAllOfAttributes struct for CPCIConfigurationFunctionSingleAllOfAttributes
type CPCIConfigurationFunctionSingleAllOfAttributes struct {
	CPciConfigurationControl *bool        `json:"cPciConfigurationControl,omitempty"`
	CSonPciList              *CSonPciList `json:"cSonPciList,omitempty"`
}

// NewCPCIConfigurationFunctionSingleAllOfAttributes instantiates a new CPCIConfigurationFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCPCIConfigurationFunctionSingleAllOfAttributes() *CPCIConfigurationFunctionSingleAllOfAttributes {
	this := CPCIConfigurationFunctionSingleAllOfAttributes{}
	return &this
}

// NewCPCIConfigurationFunctionSingleAllOfAttributesWithDefaults instantiates a new CPCIConfigurationFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCPCIConfigurationFunctionSingleAllOfAttributesWithDefaults() *CPCIConfigurationFunctionSingleAllOfAttributes {
	this := CPCIConfigurationFunctionSingleAllOfAttributes{}
	return &this
}

// GetCPciConfigurationControl returns the CPciConfigurationControl field value if set, zero value otherwise.
func (o *CPCIConfigurationFunctionSingleAllOfAttributes) GetCPciConfigurationControl() bool {
	if o == nil || IsNil(o.CPciConfigurationControl) {
		var ret bool
		return ret
	}
	return *o.CPciConfigurationControl
}

// GetCPciConfigurationControlOk returns a tuple with the CPciConfigurationControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPCIConfigurationFunctionSingleAllOfAttributes) GetCPciConfigurationControlOk() (*bool, bool) {
	if o == nil || IsNil(o.CPciConfigurationControl) {
		return nil, false
	}
	return o.CPciConfigurationControl, true
}

// HasCPciConfigurationControl returns a boolean if a field has been set.
func (o *CPCIConfigurationFunctionSingleAllOfAttributes) HasCPciConfigurationControl() bool {
	if o != nil && !IsNil(o.CPciConfigurationControl) {
		return true
	}

	return false
}

// SetCPciConfigurationControl gets a reference to the given bool and assigns it to the CPciConfigurationControl field.
func (o *CPCIConfigurationFunctionSingleAllOfAttributes) SetCPciConfigurationControl(v bool) {
	o.CPciConfigurationControl = &v
}

// GetCSonPciList returns the CSonPciList field value if set, zero value otherwise.
func (o *CPCIConfigurationFunctionSingleAllOfAttributes) GetCSonPciList() CSonPciList {
	if o == nil || IsNil(o.CSonPciList) {
		var ret CSonPciList
		return ret
	}
	return *o.CSonPciList
}

// GetCSonPciListOk returns a tuple with the CSonPciList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CPCIConfigurationFunctionSingleAllOfAttributes) GetCSonPciListOk() (*CSonPciList, bool) {
	if o == nil || IsNil(o.CSonPciList) {
		return nil, false
	}
	return o.CSonPciList, true
}

// HasCSonPciList returns a boolean if a field has been set.
func (o *CPCIConfigurationFunctionSingleAllOfAttributes) HasCSonPciList() bool {
	if o != nil && !IsNil(o.CSonPciList) {
		return true
	}

	return false
}

// SetCSonPciList gets a reference to the given CSonPciList and assigns it to the CSonPciList field.
func (o *CPCIConfigurationFunctionSingleAllOfAttributes) SetCSonPciList(v CSonPciList) {
	o.CSonPciList = &v
}

func (o CPCIConfigurationFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CPCIConfigurationFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CPciConfigurationControl) {
		toSerialize["cPciConfigurationControl"] = o.CPciConfigurationControl
	}
	if !IsNil(o.CSonPciList) {
		toSerialize["cSonPciList"] = o.CSonPciList
	}
	return toSerialize, nil
}

type NullableCPCIConfigurationFunctionSingleAllOfAttributes struct {
	value *CPCIConfigurationFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableCPCIConfigurationFunctionSingleAllOfAttributes) Get() *CPCIConfigurationFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableCPCIConfigurationFunctionSingleAllOfAttributes) Set(val *CPCIConfigurationFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCPCIConfigurationFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCPCIConfigurationFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCPCIConfigurationFunctionSingleAllOfAttributes(val *CPCIConfigurationFunctionSingleAllOfAttributes) *NullableCPCIConfigurationFunctionSingleAllOfAttributes {
	return &NullableCPCIConfigurationFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableCPCIConfigurationFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCPCIConfigurationFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

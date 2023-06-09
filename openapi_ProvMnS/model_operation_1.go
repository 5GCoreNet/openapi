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

// checks if the Operation1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Operation1{}

// Operation1 struct for Operation1
type Operation1 struct {
	Name               *string             `json:"name,omitempty"`
	AllowedNFTypes     *NFType             `json:"allowedNFTypes,omitempty"`
	OperationSemantics *OperationSemantics `json:"operationSemantics,omitempty"`
}

// NewOperation1 instantiates a new Operation1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperation1() *Operation1 {
	this := Operation1{}
	return &this
}

// NewOperation1WithDefaults instantiates a new Operation1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperation1WithDefaults() *Operation1 {
	this := Operation1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Operation1) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation1) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Operation1) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Operation1) SetName(v string) {
	o.Name = &v
}

// GetAllowedNFTypes returns the AllowedNFTypes field value if set, zero value otherwise.
func (o *Operation1) GetAllowedNFTypes() NFType {
	if o == nil || IsNil(o.AllowedNFTypes) {
		var ret NFType
		return ret
	}
	return *o.AllowedNFTypes
}

// GetAllowedNFTypesOk returns a tuple with the AllowedNFTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation1) GetAllowedNFTypesOk() (*NFType, bool) {
	if o == nil || IsNil(o.AllowedNFTypes) {
		return nil, false
	}
	return o.AllowedNFTypes, true
}

// HasAllowedNFTypes returns a boolean if a field has been set.
func (o *Operation1) HasAllowedNFTypes() bool {
	if o != nil && !IsNil(o.AllowedNFTypes) {
		return true
	}

	return false
}

// SetAllowedNFTypes gets a reference to the given NFType and assigns it to the AllowedNFTypes field.
func (o *Operation1) SetAllowedNFTypes(v NFType) {
	o.AllowedNFTypes = &v
}

// GetOperationSemantics returns the OperationSemantics field value if set, zero value otherwise.
func (o *Operation1) GetOperationSemantics() OperationSemantics {
	if o == nil || IsNil(o.OperationSemantics) {
		var ret OperationSemantics
		return ret
	}
	return *o.OperationSemantics
}

// GetOperationSemanticsOk returns a tuple with the OperationSemantics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation1) GetOperationSemanticsOk() (*OperationSemantics, bool) {
	if o == nil || IsNil(o.OperationSemantics) {
		return nil, false
	}
	return o.OperationSemantics, true
}

// HasOperationSemantics returns a boolean if a field has been set.
func (o *Operation1) HasOperationSemantics() bool {
	if o != nil && !IsNil(o.OperationSemantics) {
		return true
	}

	return false
}

// SetOperationSemantics gets a reference to the given OperationSemantics and assigns it to the OperationSemantics field.
func (o *Operation1) SetOperationSemantics(v OperationSemantics) {
	o.OperationSemantics = &v
}

func (o Operation1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Operation1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AllowedNFTypes) {
		toSerialize["allowedNFTypes"] = o.AllowedNFTypes
	}
	if !IsNil(o.OperationSemantics) {
		toSerialize["operationSemantics"] = o.OperationSemantics
	}
	return toSerialize, nil
}

type NullableOperation1 struct {
	value *Operation1
	isSet bool
}

func (v NullableOperation1) Get() *Operation1 {
	return v.value
}

func (v *NullableOperation1) Set(val *Operation1) {
	v.value = val
	v.isSet = true
}

func (v NullableOperation1) IsSet() bool {
	return v.isSet
}

func (v *NullableOperation1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperation1(val *Operation1) *NullableOperation1 {
	return &NullableOperation1{value: val, isSet: true}
}

func (v NullableOperation1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperation1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

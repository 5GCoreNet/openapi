/*
CAPIF_Routing_Info_API

API for Routing information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Routing_Info_API

import (
	"encoding/json"
)

// checks if the CustomOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomOperation{}

// CustomOperation Represents the description of a custom operation.
type CustomOperation struct {
	CommType CommunicationType `json:"commType"`
	// it is set as {custOpName} part of the URI structure for a custom operation without resource association as defined in clause 5.2.4 of 3GPP TS 29.122.
	CustOpName string `json:"custOpName"`
	// Supported HTTP methods for the API resource. Only applicable when the protocol in AefProfile indicates HTTP.
	Operations []Operation `json:"operations,omitempty"`
	// Text description of the custom operation
	Description *string `json:"description,omitempty"`
}

// NewCustomOperation instantiates a new CustomOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomOperation(commType CommunicationType, custOpName string) *CustomOperation {
	this := CustomOperation{}
	this.CommType = commType
	this.CustOpName = custOpName
	return &this
}

// NewCustomOperationWithDefaults instantiates a new CustomOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomOperationWithDefaults() *CustomOperation {
	this := CustomOperation{}
	return &this
}

// GetCommType returns the CommType field value
func (o *CustomOperation) GetCommType() CommunicationType {
	if o == nil {
		var ret CommunicationType
		return ret
	}

	return o.CommType
}

// GetCommTypeOk returns a tuple with the CommType field value
// and a boolean to check if the value has been set.
func (o *CustomOperation) GetCommTypeOk() (*CommunicationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommType, true
}

// SetCommType sets field value
func (o *CustomOperation) SetCommType(v CommunicationType) {
	o.CommType = v
}

// GetCustOpName returns the CustOpName field value
func (o *CustomOperation) GetCustOpName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustOpName
}

// GetCustOpNameOk returns a tuple with the CustOpName field value
// and a boolean to check if the value has been set.
func (o *CustomOperation) GetCustOpNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustOpName, true
}

// SetCustOpName sets field value
func (o *CustomOperation) SetCustOpName(v string) {
	o.CustOpName = v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *CustomOperation) GetOperations() []Operation {
	if o == nil || IsNil(o.Operations) {
		var ret []Operation
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOperation) GetOperationsOk() ([]Operation, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *CustomOperation) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []Operation and assigns it to the Operations field.
func (o *CustomOperation) SetOperations(v []Operation) {
	o.Operations = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomOperation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOperation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomOperation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomOperation) SetDescription(v string) {
	o.Description = &v
}

func (o CustomOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commType"] = o.CommType
	toSerialize["custOpName"] = o.CustOpName
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableCustomOperation struct {
	value *CustomOperation
	isSet bool
}

func (v NullableCustomOperation) Get() *CustomOperation {
	return v.value
}

func (v *NullableCustomOperation) Set(val *CustomOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomOperation(val *CustomOperation) *NullableCustomOperation {
	return &NullableCustomOperation{value: val, isSet: true}
}

func (v NullableCustomOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

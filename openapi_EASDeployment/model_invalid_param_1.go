/*
3gpp-eas-deployment

API for AF provisioned EAS Deployment.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EASDeployment

import (
	"encoding/json"
)

// checks if the InvalidParam1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidParam1{}

// InvalidParam1 It contains an invalid parameter and a related description.
type InvalidParam1 struct {
	// If the invalid parameter is an attribute in a JSON body, this IE shall contain the  attribute's name and shall be encoded as a JSON Pointer. If the invalid parameter is  an HTTP header, this IE shall be formatted as the concatenation of the string \"header \"  plus the name of such header. If the invalid parameter is a query parameter, this IE  shall be formatted as the concatenation of the string \"query \" plus the name of such  query parameter. If the invalid parameter is a variable part in the path of a resource  URI, this IE shall contain the name of the variable, including the symbols \"{\" and \"}\"  used in OpenAPI specification as the notation to represent variable path segments. 
	Param string `json:"param"`
	// A human-readable reason, e.g. \"must be a positive integer\". In cases involving failed  operations in a PATCH request, the reason string should identify the operation that  failed using the operation's array index to assist in correlation of the invalid  parameter with the failed operation, e.g.\" Replacement value invalid for attribute  (failed operation index= 4)\" 
	Reason *string `json:"reason,omitempty"`
}

// NewInvalidParam1 instantiates a new InvalidParam1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidParam1(param string) *InvalidParam1 {
	this := InvalidParam1{}
	this.Param = param
	return &this
}

// NewInvalidParam1WithDefaults instantiates a new InvalidParam1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidParam1WithDefaults() *InvalidParam1 {
	this := InvalidParam1{}
	return &this
}

// GetParam returns the Param field value
func (o *InvalidParam1) GetParam() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Param
}

// GetParamOk returns a tuple with the Param field value
// and a boolean to check if the value has been set.
func (o *InvalidParam1) GetParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Param, true
}

// SetParam sets field value
func (o *InvalidParam1) SetParam(v string) {
	o.Param = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InvalidParam1) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidParam1) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InvalidParam1) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InvalidParam1) SetReason(v string) {
	o.Reason = &v
}

func (o InvalidParam1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvalidParam1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["param"] = o.Param
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableInvalidParam1 struct {
	value *InvalidParam1
	isSet bool
}

func (v NullableInvalidParam1) Get() *InvalidParam1 {
	return v.value
}

func (v *NullableInvalidParam1) Set(val *InvalidParam1) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidParam1) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidParam1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidParam1(val *InvalidParam1) *NullableInvalidParam1 {
	return &NullableInvalidParam1{value: val, isSet: true}
}

func (v NullableInvalidParam1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidParam1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



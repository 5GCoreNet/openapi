/*
Nnssaaf_NSSAA

Network Slice-Specific Authentication and Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssaaf_NSSAA

import (
	"encoding/json"
)

// checks if the InvalidParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidParam{}

// InvalidParam It contains an invalid parameter and a related description.
type InvalidParam struct {
	// If the invalid parameter is an attribute in a JSON body, this IE shall contain the  attribute's name and shall be encoded as a JSON Pointer. If the invalid parameter is  an HTTP header, this IE shall be formatted as the concatenation of the string \"header \"  plus the name of such header. If the invalid parameter is a query parameter, this IE  shall be formatted as the concatenation of the string \"query \" plus the name of such  query parameter. If the invalid parameter is a variable part in the path of a resource  URI, this IE shall contain the name of the variable, including the symbols \"{\" and \"}\"  used in OpenAPI specification as the notation to represent variable path segments. 
	Param string `json:"param"`
	// A human-readable reason, e.g. \"must be a positive integer\". In cases involving failed  operations in a PATCH request, the reason string should identify the operation that  failed using the operation's array index to assist in correlation of the invalid  parameter with the failed operation, e.g.\" Replacement value invalid for attribute  (failed operation index= 4)\" 
	Reason *string `json:"reason,omitempty"`
}

// NewInvalidParam instantiates a new InvalidParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidParam(param string) *InvalidParam {
	this := InvalidParam{}
	this.Param = param
	return &this
}

// NewInvalidParamWithDefaults instantiates a new InvalidParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidParamWithDefaults() *InvalidParam {
	this := InvalidParam{}
	return &this
}

// GetParam returns the Param field value
func (o *InvalidParam) GetParam() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Param
}

// GetParamOk returns a tuple with the Param field value
// and a boolean to check if the value has been set.
func (o *InvalidParam) GetParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Param, true
}

// SetParam sets field value
func (o *InvalidParam) SetParam(v string) {
	o.Param = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InvalidParam) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidParam) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InvalidParam) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InvalidParam) SetReason(v string) {
	o.Reason = &v
}

func (o InvalidParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvalidParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["param"] = o.Param
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableInvalidParam struct {
	value *InvalidParam
	isSet bool
}

func (v NullableInvalidParam) Get() *InvalidParam {
	return v.value
}

func (v *NullableInvalidParam) Set(val *InvalidParam) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidParam) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidParam(val *InvalidParam) *NullableInvalidParam {
	return &NullableInvalidParam{value: val, isSet: true}
}

func (v NullableInvalidParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



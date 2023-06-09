/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_StreamingDataMnS

import (
	"encoding/json"
)

// checks if the FailedConnectionResponseType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailedConnectionResponseType{}

// FailedConnectionResponseType struct for FailedConnectionResponseType
type FailedConnectionResponseType struct {
	Error []FailedConnectionResponseTypeErrorInner `json:"error,omitempty"`
}

// NewFailedConnectionResponseType instantiates a new FailedConnectionResponseType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailedConnectionResponseType() *FailedConnectionResponseType {
	this := FailedConnectionResponseType{}
	return &this
}

// NewFailedConnectionResponseTypeWithDefaults instantiates a new FailedConnectionResponseType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailedConnectionResponseTypeWithDefaults() *FailedConnectionResponseType {
	this := FailedConnectionResponseType{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *FailedConnectionResponseType) GetError() []FailedConnectionResponseTypeErrorInner {
	if o == nil || IsNil(o.Error) {
		var ret []FailedConnectionResponseTypeErrorInner
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailedConnectionResponseType) GetErrorOk() ([]FailedConnectionResponseTypeErrorInner, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *FailedConnectionResponseType) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []FailedConnectionResponseTypeErrorInner and assigns it to the Error field.
func (o *FailedConnectionResponseType) SetError(v []FailedConnectionResponseTypeErrorInner) {
	o.Error = v
}

func (o FailedConnectionResponseType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailedConnectionResponseType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableFailedConnectionResponseType struct {
	value *FailedConnectionResponseType
	isSet bool
}

func (v NullableFailedConnectionResponseType) Get() *FailedConnectionResponseType {
	return v.value
}

func (v *NullableFailedConnectionResponseType) Set(val *FailedConnectionResponseType) {
	v.value = val
	v.isSet = true
}

func (v NullableFailedConnectionResponseType) IsSet() bool {
	return v.isSet
}

func (v *NullableFailedConnectionResponseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailedConnectionResponseType(val *FailedConnectionResponseType) *NullableFailedConnectionResponseType {
	return &NullableFailedConnectionResponseType{value: val, isSet: true}
}

func (v NullableFailedConnectionResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailedConnectionResponseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

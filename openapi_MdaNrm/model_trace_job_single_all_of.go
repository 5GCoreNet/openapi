/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
)

// checks if the TraceJobSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceJobSingleAllOf{}

// TraceJobSingleAllOf struct for TraceJobSingleAllOf
type TraceJobSingleAllOf struct {
	Attributes *TraceJobAttr `json:"attributes,omitempty"`
	Files []FilesSingle `json:"Files,omitempty"`
}

// NewTraceJobSingleAllOf instantiates a new TraceJobSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceJobSingleAllOf() *TraceJobSingleAllOf {
	this := TraceJobSingleAllOf{}
	return &this
}

// NewTraceJobSingleAllOfWithDefaults instantiates a new TraceJobSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceJobSingleAllOfWithDefaults() *TraceJobSingleAllOf {
	this := TraceJobSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TraceJobSingleAllOf) GetAttributes() TraceJobAttr {
	if o == nil || isNil(o.Attributes) {
		var ret TraceJobAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceJobSingleAllOf) GetAttributesOk() (*TraceJobAttr, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TraceJobSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TraceJobAttr and assigns it to the Attributes field.
func (o *TraceJobSingleAllOf) SetAttributes(v TraceJobAttr) {
	o.Attributes = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *TraceJobSingleAllOf) GetFiles() []FilesSingle {
	if o == nil || isNil(o.Files) {
		var ret []FilesSingle
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceJobSingleAllOf) GetFilesOk() ([]FilesSingle, bool) {
	if o == nil || isNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *TraceJobSingleAllOf) HasFiles() bool {
	if o != nil && !isNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FilesSingle and assigns it to the Files field.
func (o *TraceJobSingleAllOf) SetFiles(v []FilesSingle) {
	o.Files = v
}

func (o TraceJobSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceJobSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Files) {
		toSerialize["Files"] = o.Files
	}
	return toSerialize, nil
}

type NullableTraceJobSingleAllOf struct {
	value *TraceJobSingleAllOf
	isSet bool
}

func (v NullableTraceJobSingleAllOf) Get() *TraceJobSingleAllOf {
	return v.value
}

func (v *NullableTraceJobSingleAllOf) Set(val *TraceJobSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceJobSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceJobSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceJobSingleAllOf(val *TraceJobSingleAllOf) *NullableTraceJobSingleAllOf {
	return &NullableTraceJobSingleAllOf{value: val, isSet: true}
}

func (v NullableTraceJobSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceJobSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



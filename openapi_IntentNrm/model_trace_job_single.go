/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
)

// checks if the TraceJobSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceJobSingle{}

// TraceJobSingle struct for TraceJobSingle
type TraceJobSingle struct {
	Top
	Attributes *TraceJobAttr `json:"attributes,omitempty"`
	Files      []FilesSingle `json:"Files,omitempty"`
}

// NewTraceJobSingle instantiates a new TraceJobSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceJobSingle(id NullableString) *TraceJobSingle {
	this := TraceJobSingle{}
	this.Id = id
	return &this
}

// NewTraceJobSingleWithDefaults instantiates a new TraceJobSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceJobSingleWithDefaults() *TraceJobSingle {
	this := TraceJobSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TraceJobSingle) GetAttributes() TraceJobAttr {
	if o == nil || IsNil(o.Attributes) {
		var ret TraceJobAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceJobSingle) GetAttributesOk() (*TraceJobAttr, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TraceJobSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TraceJobAttr and assigns it to the Attributes field.
func (o *TraceJobSingle) SetAttributes(v TraceJobAttr) {
	o.Attributes = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *TraceJobSingle) GetFiles() []FilesSingle {
	if o == nil || IsNil(o.Files) {
		var ret []FilesSingle
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceJobSingle) GetFilesOk() ([]FilesSingle, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *TraceJobSingle) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FilesSingle and assigns it to the Files field.
func (o *TraceJobSingle) SetFiles(v []FilesSingle) {
	o.Files = v
}

func (o TraceJobSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceJobSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop, errTop := json.Marshal(o.Top)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	errTop = json.Unmarshal([]byte(serializedTop), &toSerialize)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Files) {
		toSerialize["Files"] = o.Files
	}
	return toSerialize, nil
}

type NullableTraceJobSingle struct {
	value *TraceJobSingle
	isSet bool
}

func (v NullableTraceJobSingle) Get() *TraceJobSingle {
	return v.value
}

func (v *NullableTraceJobSingle) Set(val *TraceJobSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceJobSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceJobSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceJobSingle(val *TraceJobSingle) *NullableTraceJobSingle {
	return &NullableTraceJobSingle{value: val, isSet: true}
}

func (v NullableTraceJobSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceJobSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the PerfMetricJobSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerfMetricJobSingle{}

// PerfMetricJobSingle struct for PerfMetricJobSingle
type PerfMetricJobSingle struct {
	Top
	Attributes *PerfMetricJobSingleAllOfAttributes `json:"attributes,omitempty"`
	Files      []FilesSingle                       `json:"Files,omitempty"`
}

// NewPerfMetricJobSingle instantiates a new PerfMetricJobSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerfMetricJobSingle(id NullableString) *PerfMetricJobSingle {
	this := PerfMetricJobSingle{}
	this.Id = id
	return &this
}

// NewPerfMetricJobSingleWithDefaults instantiates a new PerfMetricJobSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerfMetricJobSingleWithDefaults() *PerfMetricJobSingle {
	this := PerfMetricJobSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PerfMetricJobSingle) GetAttributes() PerfMetricJobSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret PerfMetricJobSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfMetricJobSingle) GetAttributesOk() (*PerfMetricJobSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PerfMetricJobSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PerfMetricJobSingleAllOfAttributes and assigns it to the Attributes field.
func (o *PerfMetricJobSingle) SetAttributes(v PerfMetricJobSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *PerfMetricJobSingle) GetFiles() []FilesSingle {
	if o == nil || IsNil(o.Files) {
		var ret []FilesSingle
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfMetricJobSingle) GetFilesOk() ([]FilesSingle, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *PerfMetricJobSingle) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FilesSingle and assigns it to the Files field.
func (o *PerfMetricJobSingle) SetFiles(v []FilesSingle) {
	o.Files = v
}

func (o PerfMetricJobSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerfMetricJobSingle) ToMap() (map[string]interface{}, error) {
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

type NullablePerfMetricJobSingle struct {
	value *PerfMetricJobSingle
	isSet bool
}

func (v NullablePerfMetricJobSingle) Get() *PerfMetricJobSingle {
	return v.value
}

func (v *NullablePerfMetricJobSingle) Set(val *PerfMetricJobSingle) {
	v.value = val
	v.isSet = true
}

func (v NullablePerfMetricJobSingle) IsSet() bool {
	return v.isSet
}

func (v *NullablePerfMetricJobSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerfMetricJobSingle(val *PerfMetricJobSingle) *NullablePerfMetricJobSingle {
	return &NullablePerfMetricJobSingle{value: val, isSet: true}
}

func (v NullablePerfMetricJobSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerfMetricJobSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

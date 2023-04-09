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

// checks if the PerfMetricJobSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerfMetricJobSingleAllOf{}

// PerfMetricJobSingleAllOf struct for PerfMetricJobSingleAllOf
type PerfMetricJobSingleAllOf struct {
	Attributes *PerfMetricJobSingleAllOfAttributes `json:"attributes,omitempty"`
	Files      []FilesSingle                       `json:"Files,omitempty"`
}

// NewPerfMetricJobSingleAllOf instantiates a new PerfMetricJobSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerfMetricJobSingleAllOf() *PerfMetricJobSingleAllOf {
	this := PerfMetricJobSingleAllOf{}
	return &this
}

// NewPerfMetricJobSingleAllOfWithDefaults instantiates a new PerfMetricJobSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerfMetricJobSingleAllOfWithDefaults() *PerfMetricJobSingleAllOf {
	this := PerfMetricJobSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PerfMetricJobSingleAllOf) GetAttributes() PerfMetricJobSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret PerfMetricJobSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfMetricJobSingleAllOf) GetAttributesOk() (*PerfMetricJobSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PerfMetricJobSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PerfMetricJobSingleAllOfAttributes and assigns it to the Attributes field.
func (o *PerfMetricJobSingleAllOf) SetAttributes(v PerfMetricJobSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *PerfMetricJobSingleAllOf) GetFiles() []FilesSingle {
	if o == nil || IsNil(o.Files) {
		var ret []FilesSingle
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfMetricJobSingleAllOf) GetFilesOk() ([]FilesSingle, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *PerfMetricJobSingleAllOf) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FilesSingle and assigns it to the Files field.
func (o *PerfMetricJobSingleAllOf) SetFiles(v []FilesSingle) {
	o.Files = v
}

func (o PerfMetricJobSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerfMetricJobSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Files) {
		toSerialize["Files"] = o.Files
	}
	return toSerialize, nil
}

type NullablePerfMetricJobSingleAllOf struct {
	value *PerfMetricJobSingleAllOf
	isSet bool
}

func (v NullablePerfMetricJobSingleAllOf) Get() *PerfMetricJobSingleAllOf {
	return v.value
}

func (v *NullablePerfMetricJobSingleAllOf) Set(val *PerfMetricJobSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePerfMetricJobSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePerfMetricJobSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerfMetricJobSingleAllOf(val *PerfMetricJobSingleAllOf) *NullablePerfMetricJobSingleAllOf {
	return &NullablePerfMetricJobSingleAllOf{value: val, isSet: true}
}

func (v NullablePerfMetricJobSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerfMetricJobSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

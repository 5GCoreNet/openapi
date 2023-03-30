/*
Nmfaf_3daDataManagement

MFAF 3GPP DCCF Adaptor (3DA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3daDataManagement

import (
	"encoding/json"
)

// checks if the ParameterProcessingInstruction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterProcessingInstruction{}

// ParameterProcessingInstruction Contains an event parameter name and the respective event parameter values and sets of attributes to be used in summarized reports. 
type ParameterProcessingInstruction struct {
	// A JSON pointer value that references an attribute within the notification object to which the processing instruction is applied. 
	Name string `json:"name"`
	// A list of values for the attribute identified by the name attribute.
	Values []interface{} `json:"values"`
	// Attributes requested to be used in the summarized reports.
	SumAttrs []SummarizationAttribute `json:"sumAttrs"`
	AggrLevel *AggregationLevel `json:"aggrLevel,omitempty"`
	// Indicates the UEs for which processed reports are requested.
	Supis []string `json:"supis,omitempty"`
	// Indicates the Areas of Interest for which processed reports are requested.
	Areas []NetworkAreaInfo `json:"areas,omitempty"`
}

// NewParameterProcessingInstruction instantiates a new ParameterProcessingInstruction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterProcessingInstruction(name string, values []interface{}, sumAttrs []SummarizationAttribute) *ParameterProcessingInstruction {
	this := ParameterProcessingInstruction{}
	this.Name = name
	this.Values = values
	this.SumAttrs = sumAttrs
	return &this
}

// NewParameterProcessingInstructionWithDefaults instantiates a new ParameterProcessingInstruction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterProcessingInstructionWithDefaults() *ParameterProcessingInstruction {
	this := ParameterProcessingInstruction{}
	return &this
}

// GetName returns the Name field value
func (o *ParameterProcessingInstruction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterProcessingInstruction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParameterProcessingInstruction) SetName(v string) {
	o.Name = v
}

// GetValues returns the Values field value
func (o *ParameterProcessingInstruction) GetValues() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ParameterProcessingInstruction) GetValuesOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ParameterProcessingInstruction) SetValues(v []interface{}) {
	o.Values = v
}

// GetSumAttrs returns the SumAttrs field value
func (o *ParameterProcessingInstruction) GetSumAttrs() []SummarizationAttribute {
	if o == nil {
		var ret []SummarizationAttribute
		return ret
	}

	return o.SumAttrs
}

// GetSumAttrsOk returns a tuple with the SumAttrs field value
// and a boolean to check if the value has been set.
func (o *ParameterProcessingInstruction) GetSumAttrsOk() ([]SummarizationAttribute, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumAttrs, true
}

// SetSumAttrs sets field value
func (o *ParameterProcessingInstruction) SetSumAttrs(v []SummarizationAttribute) {
	o.SumAttrs = v
}

// GetAggrLevel returns the AggrLevel field value if set, zero value otherwise.
func (o *ParameterProcessingInstruction) GetAggrLevel() AggregationLevel {
	if o == nil || IsNil(o.AggrLevel) {
		var ret AggregationLevel
		return ret
	}
	return *o.AggrLevel
}

// GetAggrLevelOk returns a tuple with the AggrLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProcessingInstruction) GetAggrLevelOk() (*AggregationLevel, bool) {
	if o == nil || IsNil(o.AggrLevel) {
		return nil, false
	}
	return o.AggrLevel, true
}

// HasAggrLevel returns a boolean if a field has been set.
func (o *ParameterProcessingInstruction) HasAggrLevel() bool {
	if o != nil && !IsNil(o.AggrLevel) {
		return true
	}

	return false
}

// SetAggrLevel gets a reference to the given AggregationLevel and assigns it to the AggrLevel field.
func (o *ParameterProcessingInstruction) SetAggrLevel(v AggregationLevel) {
	o.AggrLevel = &v
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *ParameterProcessingInstruction) GetSupis() []string {
	if o == nil || IsNil(o.Supis) {
		var ret []string
		return ret
	}
	return o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProcessingInstruction) GetSupisOk() ([]string, bool) {
	if o == nil || IsNil(o.Supis) {
		return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *ParameterProcessingInstruction) HasSupis() bool {
	if o != nil && !IsNil(o.Supis) {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []string and assigns it to the Supis field.
func (o *ParameterProcessingInstruction) SetSupis(v []string) {
	o.Supis = v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *ParameterProcessingInstruction) GetAreas() []NetworkAreaInfo {
	if o == nil || IsNil(o.Areas) {
		var ret []NetworkAreaInfo
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProcessingInstruction) GetAreasOk() ([]NetworkAreaInfo, bool) {
	if o == nil || IsNil(o.Areas) {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *ParameterProcessingInstruction) HasAreas() bool {
	if o != nil && !IsNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []NetworkAreaInfo and assigns it to the Areas field.
func (o *ParameterProcessingInstruction) SetAreas(v []NetworkAreaInfo) {
	o.Areas = v
}

func (o ParameterProcessingInstruction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterProcessingInstruction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["values"] = o.Values
	toSerialize["sumAttrs"] = o.SumAttrs
	if !IsNil(o.AggrLevel) {
		toSerialize["aggrLevel"] = o.AggrLevel
	}
	if !IsNil(o.Supis) {
		toSerialize["supis"] = o.Supis
	}
	if !IsNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	return toSerialize, nil
}

type NullableParameterProcessingInstruction struct {
	value *ParameterProcessingInstruction
	isSet bool
}

func (v NullableParameterProcessingInstruction) Get() *ParameterProcessingInstruction {
	return v.value
}

func (v *NullableParameterProcessingInstruction) Set(val *ParameterProcessingInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterProcessingInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterProcessingInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterProcessingInstruction(val *ParameterProcessingInstruction) *NullableParameterProcessingInstruction {
	return &NullableParameterProcessingInstruction{value: val, isSet: true}
}

func (v NullableParameterProcessingInstruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterProcessingInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



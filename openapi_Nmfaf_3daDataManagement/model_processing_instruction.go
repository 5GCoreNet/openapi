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

// checks if the ProcessingInstruction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessingInstruction{}

// ProcessingInstruction Contains instructions related to the processing of notifications.
type ProcessingInstruction struct {
	EventId DccfEvent `json:"eventId"`
	// indicating a time in seconds.
	ProcInterval int32 `json:"procInterval"`
	// List of event parameter names, and for each event parameter name, respective event parameter values and sets of the attributes to be used in the summarized reports. 
	ParamProcInstructs []ParameterProcessingInstruction `json:"paramProcInstructs,omitempty"`
}

// NewProcessingInstruction instantiates a new ProcessingInstruction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessingInstruction(eventId DccfEvent, procInterval int32) *ProcessingInstruction {
	this := ProcessingInstruction{}
	this.EventId = eventId
	this.ProcInterval = procInterval
	return &this
}

// NewProcessingInstructionWithDefaults instantiates a new ProcessingInstruction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessingInstructionWithDefaults() *ProcessingInstruction {
	this := ProcessingInstruction{}
	return &this
}

// GetEventId returns the EventId field value
func (o *ProcessingInstruction) GetEventId() DccfEvent {
	if o == nil {
		var ret DccfEvent
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *ProcessingInstruction) GetEventIdOk() (*DccfEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *ProcessingInstruction) SetEventId(v DccfEvent) {
	o.EventId = v
}

// GetProcInterval returns the ProcInterval field value
func (o *ProcessingInstruction) GetProcInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProcInterval
}

// GetProcIntervalOk returns a tuple with the ProcInterval field value
// and a boolean to check if the value has been set.
func (o *ProcessingInstruction) GetProcIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcInterval, true
}

// SetProcInterval sets field value
func (o *ProcessingInstruction) SetProcInterval(v int32) {
	o.ProcInterval = v
}

// GetParamProcInstructs returns the ParamProcInstructs field value if set, zero value otherwise.
func (o *ProcessingInstruction) GetParamProcInstructs() []ParameterProcessingInstruction {
	if o == nil || isNil(o.ParamProcInstructs) {
		var ret []ParameterProcessingInstruction
		return ret
	}
	return o.ParamProcInstructs
}

// GetParamProcInstructsOk returns a tuple with the ParamProcInstructs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessingInstruction) GetParamProcInstructsOk() ([]ParameterProcessingInstruction, bool) {
	if o == nil || isNil(o.ParamProcInstructs) {
		return nil, false
	}
	return o.ParamProcInstructs, true
}

// HasParamProcInstructs returns a boolean if a field has been set.
func (o *ProcessingInstruction) HasParamProcInstructs() bool {
	if o != nil && !isNil(o.ParamProcInstructs) {
		return true
	}

	return false
}

// SetParamProcInstructs gets a reference to the given []ParameterProcessingInstruction and assigns it to the ParamProcInstructs field.
func (o *ProcessingInstruction) SetParamProcInstructs(v []ParameterProcessingInstruction) {
	o.ParamProcInstructs = v
}

func (o ProcessingInstruction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessingInstruction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventId"] = o.EventId
	toSerialize["procInterval"] = o.ProcInterval
	if !isNil(o.ParamProcInstructs) {
		toSerialize["paramProcInstructs"] = o.ParamProcInstructs
	}
	return toSerialize, nil
}

type NullableProcessingInstruction struct {
	value *ProcessingInstruction
	isSet bool
}

func (v NullableProcessingInstruction) Get() *ProcessingInstruction {
	return v.value
}

func (v *NullableProcessingInstruction) Set(val *ProcessingInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessingInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessingInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessingInstruction(val *ProcessingInstruction) *NullableProcessingInstruction {
	return &NullableProcessingInstruction{value: val, isSet: true}
}

func (v NullableProcessingInstruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessingInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



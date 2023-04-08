/*
3gpp-time-sync-exposure

API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_TimeSyncExposure

import (
	"encoding/json"
)

// checks if the EventFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventFilter{}

// EventFilter Contains the filter conditions to match for notifying the event(s) of time synchronization capabilities. 
type EventFilter struct {
	InstanceTypes []InstanceType `json:"instanceTypes,omitempty"`
	TransProtocols []Protocol `json:"transProtocols,omitempty"`
	PtpProfiles []string `json:"ptpProfiles,omitempty"`
}

// NewEventFilter instantiates a new EventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventFilter() *EventFilter {
	this := EventFilter{}
	return &this
}

// NewEventFilterWithDefaults instantiates a new EventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventFilterWithDefaults() *EventFilter {
	this := EventFilter{}
	return &this
}

// GetInstanceTypes returns the InstanceTypes field value if set, zero value otherwise.
func (o *EventFilter) GetInstanceTypes() []InstanceType {
	if o == nil || isNil(o.InstanceTypes) {
		var ret []InstanceType
		return ret
	}
	return o.InstanceTypes
}

// GetInstanceTypesOk returns a tuple with the InstanceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetInstanceTypesOk() ([]InstanceType, bool) {
	if o == nil || isNil(o.InstanceTypes) {
		return nil, false
	}
	return o.InstanceTypes, true
}

// HasInstanceTypes returns a boolean if a field has been set.
func (o *EventFilter) HasInstanceTypes() bool {
	if o != nil && !isNil(o.InstanceTypes) {
		return true
	}

	return false
}

// SetInstanceTypes gets a reference to the given []InstanceType and assigns it to the InstanceTypes field.
func (o *EventFilter) SetInstanceTypes(v []InstanceType) {
	o.InstanceTypes = v
}

// GetTransProtocols returns the TransProtocols field value if set, zero value otherwise.
func (o *EventFilter) GetTransProtocols() []Protocol {
	if o == nil || isNil(o.TransProtocols) {
		var ret []Protocol
		return ret
	}
	return o.TransProtocols
}

// GetTransProtocolsOk returns a tuple with the TransProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetTransProtocolsOk() ([]Protocol, bool) {
	if o == nil || isNil(o.TransProtocols) {
		return nil, false
	}
	return o.TransProtocols, true
}

// HasTransProtocols returns a boolean if a field has been set.
func (o *EventFilter) HasTransProtocols() bool {
	if o != nil && !isNil(o.TransProtocols) {
		return true
	}

	return false
}

// SetTransProtocols gets a reference to the given []Protocol and assigns it to the TransProtocols field.
func (o *EventFilter) SetTransProtocols(v []Protocol) {
	o.TransProtocols = v
}

// GetPtpProfiles returns the PtpProfiles field value if set, zero value otherwise.
func (o *EventFilter) GetPtpProfiles() []string {
	if o == nil || isNil(o.PtpProfiles) {
		var ret []string
		return ret
	}
	return o.PtpProfiles
}

// GetPtpProfilesOk returns a tuple with the PtpProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetPtpProfilesOk() ([]string, bool) {
	if o == nil || isNil(o.PtpProfiles) {
		return nil, false
	}
	return o.PtpProfiles, true
}

// HasPtpProfiles returns a boolean if a field has been set.
func (o *EventFilter) HasPtpProfiles() bool {
	if o != nil && !isNil(o.PtpProfiles) {
		return true
	}

	return false
}

// SetPtpProfiles gets a reference to the given []string and assigns it to the PtpProfiles field.
func (o *EventFilter) SetPtpProfiles(v []string) {
	o.PtpProfiles = v
}

func (o EventFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InstanceTypes) {
		toSerialize["instanceTypes"] = o.InstanceTypes
	}
	if !isNil(o.TransProtocols) {
		toSerialize["transProtocols"] = o.TransProtocols
	}
	if !isNil(o.PtpProfiles) {
		toSerialize["ptpProfiles"] = o.PtpProfiles
	}
	return toSerialize, nil
}

type NullableEventFilter struct {
	value *EventFilter
	isSet bool
}

func (v NullableEventFilter) Get() *EventFilter {
	return v.value
}

func (v *NullableEventFilter) Set(val *EventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventFilter(val *EventFilter) *NullableEventFilter {
	return &NullableEventFilter{value: val, isSet: true}
}

func (v NullableEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



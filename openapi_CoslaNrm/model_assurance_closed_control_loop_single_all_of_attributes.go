/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// checks if the AssuranceClosedControlLoopSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssuranceClosedControlLoopSingleAllOfAttributes{}

// AssuranceClosedControlLoopSingleAllOfAttributes struct for AssuranceClosedControlLoopSingleAllOfAttributes
type AssuranceClosedControlLoopSingleAllOfAttributes struct {
	OperationalState *OperationalState `json:"operationalState,omitempty"`
	AdministrativeState *AdministrativeState `json:"administrativeState,omitempty"`
	ControlLoopLifeCyclePhase *ControlLoopLifeCyclePhase `json:"controlLoopLifeCyclePhase,omitempty"`
	ACCLDisallowedList *ACCLDisallowedAttributes `json:"aCCLDisallowedList,omitempty"`
}

// NewAssuranceClosedControlLoopSingleAllOfAttributes instantiates a new AssuranceClosedControlLoopSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssuranceClosedControlLoopSingleAllOfAttributes() *AssuranceClosedControlLoopSingleAllOfAttributes {
	this := AssuranceClosedControlLoopSingleAllOfAttributes{}
	return &this
}

// NewAssuranceClosedControlLoopSingleAllOfAttributesWithDefaults instantiates a new AssuranceClosedControlLoopSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssuranceClosedControlLoopSingleAllOfAttributesWithDefaults() *AssuranceClosedControlLoopSingleAllOfAttributes {
	this := AssuranceClosedControlLoopSingleAllOfAttributes{}
	return &this
}

// GetOperationalState returns the OperationalState field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) GetOperationalState() OperationalState {
	if o == nil || IsNil(o.OperationalState) {
		var ret OperationalState
		return ret
	}
	return *o.OperationalState
}

// GetOperationalStateOk returns a tuple with the OperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) GetOperationalStateOk() (*OperationalState, bool) {
	if o == nil || IsNil(o.OperationalState) {
		return nil, false
	}
	return o.OperationalState, true
}

// HasOperationalState returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) HasOperationalState() bool {
	if o != nil && !IsNil(o.OperationalState) {
		return true
	}

	return false
}

// SetOperationalState gets a reference to the given OperationalState and assigns it to the OperationalState field.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) SetOperationalState(v OperationalState) {
	o.OperationalState = &v
}

// GetAdministrativeState returns the AdministrativeState field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) GetAdministrativeState() AdministrativeState {
	if o == nil || IsNil(o.AdministrativeState) {
		var ret AdministrativeState
		return ret
	}
	return *o.AdministrativeState
}

// GetAdministrativeStateOk returns a tuple with the AdministrativeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) GetAdministrativeStateOk() (*AdministrativeState, bool) {
	if o == nil || IsNil(o.AdministrativeState) {
		return nil, false
	}
	return o.AdministrativeState, true
}

// HasAdministrativeState returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) HasAdministrativeState() bool {
	if o != nil && !IsNil(o.AdministrativeState) {
		return true
	}

	return false
}

// SetAdministrativeState gets a reference to the given AdministrativeState and assigns it to the AdministrativeState field.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) SetAdministrativeState(v AdministrativeState) {
	o.AdministrativeState = &v
}

// GetControlLoopLifeCyclePhase returns the ControlLoopLifeCyclePhase field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) GetControlLoopLifeCyclePhase() ControlLoopLifeCyclePhase {
	if o == nil || IsNil(o.ControlLoopLifeCyclePhase) {
		var ret ControlLoopLifeCyclePhase
		return ret
	}
	return *o.ControlLoopLifeCyclePhase
}

// GetControlLoopLifeCyclePhaseOk returns a tuple with the ControlLoopLifeCyclePhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) GetControlLoopLifeCyclePhaseOk() (*ControlLoopLifeCyclePhase, bool) {
	if o == nil || IsNil(o.ControlLoopLifeCyclePhase) {
		return nil, false
	}
	return o.ControlLoopLifeCyclePhase, true
}

// HasControlLoopLifeCyclePhase returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) HasControlLoopLifeCyclePhase() bool {
	if o != nil && !IsNil(o.ControlLoopLifeCyclePhase) {
		return true
	}

	return false
}

// SetControlLoopLifeCyclePhase gets a reference to the given ControlLoopLifeCyclePhase and assigns it to the ControlLoopLifeCyclePhase field.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) SetControlLoopLifeCyclePhase(v ControlLoopLifeCyclePhase) {
	o.ControlLoopLifeCyclePhase = &v
}

// GetACCLDisallowedList returns the ACCLDisallowedList field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) GetACCLDisallowedList() ACCLDisallowedAttributes {
	if o == nil || IsNil(o.ACCLDisallowedList) {
		var ret ACCLDisallowedAttributes
		return ret
	}
	return *o.ACCLDisallowedList
}

// GetACCLDisallowedListOk returns a tuple with the ACCLDisallowedList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) GetACCLDisallowedListOk() (*ACCLDisallowedAttributes, bool) {
	if o == nil || IsNil(o.ACCLDisallowedList) {
		return nil, false
	}
	return o.ACCLDisallowedList, true
}

// HasACCLDisallowedList returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) HasACCLDisallowedList() bool {
	if o != nil && !IsNil(o.ACCLDisallowedList) {
		return true
	}

	return false
}

// SetACCLDisallowedList gets a reference to the given ACCLDisallowedAttributes and assigns it to the ACCLDisallowedList field.
func (o *AssuranceClosedControlLoopSingleAllOfAttributes) SetACCLDisallowedList(v ACCLDisallowedAttributes) {
	o.ACCLDisallowedList = &v
}

func (o AssuranceClosedControlLoopSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssuranceClosedControlLoopSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OperationalState) {
		toSerialize["operationalState"] = o.OperationalState
	}
	if !IsNil(o.AdministrativeState) {
		toSerialize["administrativeState"] = o.AdministrativeState
	}
	if !IsNil(o.ControlLoopLifeCyclePhase) {
		toSerialize["controlLoopLifeCyclePhase"] = o.ControlLoopLifeCyclePhase
	}
	if !IsNil(o.ACCLDisallowedList) {
		toSerialize["aCCLDisallowedList"] = o.ACCLDisallowedList
	}
	return toSerialize, nil
}

type NullableAssuranceClosedControlLoopSingleAllOfAttributes struct {
	value *AssuranceClosedControlLoopSingleAllOfAttributes
	isSet bool
}

func (v NullableAssuranceClosedControlLoopSingleAllOfAttributes) Get() *AssuranceClosedControlLoopSingleAllOfAttributes {
	return v.value
}

func (v *NullableAssuranceClosedControlLoopSingleAllOfAttributes) Set(val *AssuranceClosedControlLoopSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAssuranceClosedControlLoopSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAssuranceClosedControlLoopSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssuranceClosedControlLoopSingleAllOfAttributes(val *AssuranceClosedControlLoopSingleAllOfAttributes) *NullableAssuranceClosedControlLoopSingleAllOfAttributes {
	return &NullableAssuranceClosedControlLoopSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableAssuranceClosedControlLoopSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssuranceClosedControlLoopSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



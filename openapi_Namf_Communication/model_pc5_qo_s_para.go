/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the Pc5QoSPara type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pc5QoSPara{}

// Pc5QoSPara Contains policy data on the PC5 QoS parameters.
type Pc5QoSPara struct {
	Pc5QosFlowList []Pc5QosFlowItem `json:"pc5QosFlowList"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	Pc5LinkAmbr *string `json:"pc5LinkAmbr,omitempty"`
}

// NewPc5QoSPara instantiates a new Pc5QoSPara object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPc5QoSPara(pc5QosFlowList []Pc5QosFlowItem) *Pc5QoSPara {
	this := Pc5QoSPara{}
	this.Pc5QosFlowList = pc5QosFlowList
	return &this
}

// NewPc5QoSParaWithDefaults instantiates a new Pc5QoSPara object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPc5QoSParaWithDefaults() *Pc5QoSPara {
	this := Pc5QoSPara{}
	return &this
}

// GetPc5QosFlowList returns the Pc5QosFlowList field value
func (o *Pc5QoSPara) GetPc5QosFlowList() []Pc5QosFlowItem {
	if o == nil {
		var ret []Pc5QosFlowItem
		return ret
	}

	return o.Pc5QosFlowList
}

// GetPc5QosFlowListOk returns a tuple with the Pc5QosFlowList field value
// and a boolean to check if the value has been set.
func (o *Pc5QoSPara) GetPc5QosFlowListOk() ([]Pc5QosFlowItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pc5QosFlowList, true
}

// SetPc5QosFlowList sets field value
func (o *Pc5QoSPara) SetPc5QosFlowList(v []Pc5QosFlowItem) {
	o.Pc5QosFlowList = v
}

// GetPc5LinkAmbr returns the Pc5LinkAmbr field value if set, zero value otherwise.
func (o *Pc5QoSPara) GetPc5LinkAmbr() string {
	if o == nil || IsNil(o.Pc5LinkAmbr) {
		var ret string
		return ret
	}
	return *o.Pc5LinkAmbr
}

// GetPc5LinkAmbrOk returns a tuple with the Pc5LinkAmbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pc5QoSPara) GetPc5LinkAmbrOk() (*string, bool) {
	if o == nil || IsNil(o.Pc5LinkAmbr) {
		return nil, false
	}
	return o.Pc5LinkAmbr, true
}

// HasPc5LinkAmbr returns a boolean if a field has been set.
func (o *Pc5QoSPara) HasPc5LinkAmbr() bool {
	if o != nil && !IsNil(o.Pc5LinkAmbr) {
		return true
	}

	return false
}

// SetPc5LinkAmbr gets a reference to the given string and assigns it to the Pc5LinkAmbr field.
func (o *Pc5QoSPara) SetPc5LinkAmbr(v string) {
	o.Pc5LinkAmbr = &v
}

func (o Pc5QoSPara) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pc5QoSPara) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pc5QosFlowList"] = o.Pc5QosFlowList
	if !IsNil(o.Pc5LinkAmbr) {
		toSerialize["pc5LinkAmbr"] = o.Pc5LinkAmbr
	}
	return toSerialize, nil
}

type NullablePc5QoSPara struct {
	value *Pc5QoSPara
	isSet bool
}

func (v NullablePc5QoSPara) Get() *Pc5QoSPara {
	return v.value
}

func (v *NullablePc5QoSPara) Set(val *Pc5QoSPara) {
	v.value = val
	v.isSet = true
}

func (v NullablePc5QoSPara) IsSet() bool {
	return v.isSet
}

func (v *NullablePc5QoSPara) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePc5QoSPara(val *Pc5QoSPara) *NullablePc5QoSPara {
	return &NullablePc5QoSPara{value: val, isSet: true}
}

func (v NullablePc5QoSPara) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePc5QoSPara) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

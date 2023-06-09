/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// checks if the KPIMonitoring type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KPIMonitoring{}

// KPIMonitoring struct for KPIMonitoring
type KPIMonitoring struct {
	ServAttrCom *ServAttrCom `json:"servAttrCom,omitempty"`
	KPIList     []string     `json:"kPIList,omitempty"`
}

// NewKPIMonitoring instantiates a new KPIMonitoring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKPIMonitoring() *KPIMonitoring {
	this := KPIMonitoring{}
	return &this
}

// NewKPIMonitoringWithDefaults instantiates a new KPIMonitoring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKPIMonitoringWithDefaults() *KPIMonitoring {
	this := KPIMonitoring{}
	return &this
}

// GetServAttrCom returns the ServAttrCom field value if set, zero value otherwise.
func (o *KPIMonitoring) GetServAttrCom() ServAttrCom {
	if o == nil || IsNil(o.ServAttrCom) {
		var ret ServAttrCom
		return ret
	}
	return *o.ServAttrCom
}

// GetServAttrComOk returns a tuple with the ServAttrCom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KPIMonitoring) GetServAttrComOk() (*ServAttrCom, bool) {
	if o == nil || IsNil(o.ServAttrCom) {
		return nil, false
	}
	return o.ServAttrCom, true
}

// HasServAttrCom returns a boolean if a field has been set.
func (o *KPIMonitoring) HasServAttrCom() bool {
	if o != nil && !IsNil(o.ServAttrCom) {
		return true
	}

	return false
}

// SetServAttrCom gets a reference to the given ServAttrCom and assigns it to the ServAttrCom field.
func (o *KPIMonitoring) SetServAttrCom(v ServAttrCom) {
	o.ServAttrCom = &v
}

// GetKPIList returns the KPIList field value if set, zero value otherwise.
func (o *KPIMonitoring) GetKPIList() []string {
	if o == nil || IsNil(o.KPIList) {
		var ret []string
		return ret
	}
	return o.KPIList
}

// GetKPIListOk returns a tuple with the KPIList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KPIMonitoring) GetKPIListOk() ([]string, bool) {
	if o == nil || IsNil(o.KPIList) {
		return nil, false
	}
	return o.KPIList, true
}

// HasKPIList returns a boolean if a field has been set.
func (o *KPIMonitoring) HasKPIList() bool {
	if o != nil && !IsNil(o.KPIList) {
		return true
	}

	return false
}

// SetKPIList gets a reference to the given []string and assigns it to the KPIList field.
func (o *KPIMonitoring) SetKPIList(v []string) {
	o.KPIList = v
}

func (o KPIMonitoring) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KPIMonitoring) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServAttrCom) {
		toSerialize["servAttrCom"] = o.ServAttrCom
	}
	if !IsNil(o.KPIList) {
		toSerialize["kPIList"] = o.KPIList
	}
	return toSerialize, nil
}

type NullableKPIMonitoring struct {
	value *KPIMonitoring
	isSet bool
}

func (v NullableKPIMonitoring) Get() *KPIMonitoring {
	return v.value
}

func (v *NullableKPIMonitoring) Set(val *KPIMonitoring) {
	v.value = val
	v.isSet = true
}

func (v NullableKPIMonitoring) IsSet() bool {
	return v.isSet
}

func (v *NullableKPIMonitoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKPIMonitoring(val *KPIMonitoring) *NullableKPIMonitoring {
	return &NullableKPIMonitoring{value: val, isSet: true}
}

func (v NullableKPIMonitoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKPIMonitoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

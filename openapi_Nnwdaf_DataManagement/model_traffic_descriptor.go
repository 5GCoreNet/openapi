/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// checks if the TrafficDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficDescriptor{}

// TrafficDescriptor Represents the Traffic Descriptor
type TrafficDescriptor struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn                      *string                `json:"dnn,omitempty"`
	SNssai                   *Snssai                `json:"sNssai,omitempty"`
	DddTrafficDescriptorList []DddTrafficDescriptor `json:"dddTrafficDescriptorList,omitempty"`
}

// NewTrafficDescriptor instantiates a new TrafficDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficDescriptor() *TrafficDescriptor {
	this := TrafficDescriptor{}
	return &this
}

// NewTrafficDescriptorWithDefaults instantiates a new TrafficDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficDescriptorWithDefaults() *TrafficDescriptor {
	this := TrafficDescriptor{}
	return &this
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *TrafficDescriptor) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficDescriptor) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *TrafficDescriptor) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *TrafficDescriptor) SetDnn(v string) {
	o.Dnn = &v
}

// GetSNssai returns the SNssai field value if set, zero value otherwise.
func (o *TrafficDescriptor) GetSNssai() Snssai {
	if o == nil || IsNil(o.SNssai) {
		var ret Snssai
		return ret
	}
	return *o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficDescriptor) GetSNssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.SNssai) {
		return nil, false
	}
	return o.SNssai, true
}

// HasSNssai returns a boolean if a field has been set.
func (o *TrafficDescriptor) HasSNssai() bool {
	if o != nil && !IsNil(o.SNssai) {
		return true
	}

	return false
}

// SetSNssai gets a reference to the given Snssai and assigns it to the SNssai field.
func (o *TrafficDescriptor) SetSNssai(v Snssai) {
	o.SNssai = &v
}

// GetDddTrafficDescriptorList returns the DddTrafficDescriptorList field value if set, zero value otherwise.
func (o *TrafficDescriptor) GetDddTrafficDescriptorList() []DddTrafficDescriptor {
	if o == nil || IsNil(o.DddTrafficDescriptorList) {
		var ret []DddTrafficDescriptor
		return ret
	}
	return o.DddTrafficDescriptorList
}

// GetDddTrafficDescriptorListOk returns a tuple with the DddTrafficDescriptorList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficDescriptor) GetDddTrafficDescriptorListOk() ([]DddTrafficDescriptor, bool) {
	if o == nil || IsNil(o.DddTrafficDescriptorList) {
		return nil, false
	}
	return o.DddTrafficDescriptorList, true
}

// HasDddTrafficDescriptorList returns a boolean if a field has been set.
func (o *TrafficDescriptor) HasDddTrafficDescriptorList() bool {
	if o != nil && !IsNil(o.DddTrafficDescriptorList) {
		return true
	}

	return false
}

// SetDddTrafficDescriptorList gets a reference to the given []DddTrafficDescriptor and assigns it to the DddTrafficDescriptorList field.
func (o *TrafficDescriptor) SetDddTrafficDescriptorList(v []DddTrafficDescriptor) {
	o.DddTrafficDescriptorList = v
}

func (o TrafficDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.SNssai) {
		toSerialize["sNssai"] = o.SNssai
	}
	if !IsNil(o.DddTrafficDescriptorList) {
		toSerialize["dddTrafficDescriptorList"] = o.DddTrafficDescriptorList
	}
	return toSerialize, nil
}

type NullableTrafficDescriptor struct {
	value *TrafficDescriptor
	isSet bool
}

func (v NullableTrafficDescriptor) Get() *TrafficDescriptor {
	return v.value
}

func (v *NullableTrafficDescriptor) Set(val *TrafficDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficDescriptor(val *TrafficDescriptor) *NullableTrafficDescriptor {
	return &NullableTrafficDescriptor{value: val, isSet: true}
}

func (v NullableTrafficDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GenericNrm

import (
	"encoding/json"
)

// checks if the SAP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAP{}

// SAP struct for SAP
type SAP struct {
	Host *HostAddr `json:"host,omitempty"`
	Port *int32    `json:"port,omitempty"`
}

// NewSAP instantiates a new SAP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAP() *SAP {
	this := SAP{}
	return &this
}

// NewSAPWithDefaults instantiates a new SAP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAPWithDefaults() *SAP {
	this := SAP{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *SAP) GetHost() HostAddr {
	if o == nil || IsNil(o.Host) {
		var ret HostAddr
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAP) GetHostOk() (*HostAddr, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *SAP) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostAddr and assigns it to the Host field.
func (o *SAP) SetHost(v HostAddr) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SAP) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAP) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SAP) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *SAP) SetPort(v int32) {
	o.Port = &v
}

func (o SAP) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableSAP struct {
	value *SAP
	isSet bool
}

func (v NullableSAP) Get() *SAP {
	return v.value
}

func (v *NullableSAP) Set(val *SAP) {
	v.value = val
	v.isSet = true
}

func (v NullableSAP) IsSet() bool {
	return v.isSet
}

func (v *NullableSAP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAP(val *SAP) *NullableSAP {
	return &NullableSAP{value: val, isSet: true}
}

func (v NullableSAP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

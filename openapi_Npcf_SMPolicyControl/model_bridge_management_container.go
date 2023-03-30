/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
)

// checks if the BridgeManagementContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BridgeManagementContainer{}

// BridgeManagementContainer Contains the UMIC.
type BridgeManagementContainer struct {
	// string with format 'bytes' as defined in OpenAPI
	BridgeManCont string `json:"bridgeManCont"`
}

// NewBridgeManagementContainer instantiates a new BridgeManagementContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBridgeManagementContainer(bridgeManCont string) *BridgeManagementContainer {
	this := BridgeManagementContainer{}
	this.BridgeManCont = bridgeManCont
	return &this
}

// NewBridgeManagementContainerWithDefaults instantiates a new BridgeManagementContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBridgeManagementContainerWithDefaults() *BridgeManagementContainer {
	this := BridgeManagementContainer{}
	return &this
}

// GetBridgeManCont returns the BridgeManCont field value
func (o *BridgeManagementContainer) GetBridgeManCont() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BridgeManCont
}

// GetBridgeManContOk returns a tuple with the BridgeManCont field value
// and a boolean to check if the value has been set.
func (o *BridgeManagementContainer) GetBridgeManContOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BridgeManCont, true
}

// SetBridgeManCont sets field value
func (o *BridgeManagementContainer) SetBridgeManCont(v string) {
	o.BridgeManCont = v
}

func (o BridgeManagementContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BridgeManagementContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bridgeManCont"] = o.BridgeManCont
	return toSerialize, nil
}

type NullableBridgeManagementContainer struct {
	value *BridgeManagementContainer
	isSet bool
}

func (v NullableBridgeManagementContainer) Get() *BridgeManagementContainer {
	return v.value
}

func (v *NullableBridgeManagementContainer) Set(val *BridgeManagementContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableBridgeManagementContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableBridgeManagementContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBridgeManagementContainer(val *BridgeManagementContainer) *NullableBridgeManagementContainer {
	return &NullableBridgeManagementContainer{value: val, isSet: true}
}

func (v NullableBridgeManagementContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBridgeManagementContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


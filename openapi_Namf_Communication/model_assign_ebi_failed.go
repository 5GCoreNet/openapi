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

// checks if the AssignEbiFailed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignEbiFailed{}

// AssignEbiFailed Represents failed assignment of EBI(s)
type AssignEbiFailed struct {
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessionId int32 `json:"pduSessionId"`
	FailedArpList []Arp `json:"failedArpList,omitempty"`
}

// NewAssignEbiFailed instantiates a new AssignEbiFailed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignEbiFailed(pduSessionId int32) *AssignEbiFailed {
	this := AssignEbiFailed{}
	this.PduSessionId = pduSessionId
	return &this
}

// NewAssignEbiFailedWithDefaults instantiates a new AssignEbiFailed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignEbiFailedWithDefaults() *AssignEbiFailed {
	this := AssignEbiFailed{}
	return &this
}

// GetPduSessionId returns the PduSessionId field value
func (o *AssignEbiFailed) GetPduSessionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value
// and a boolean to check if the value has been set.
func (o *AssignEbiFailed) GetPduSessionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PduSessionId, true
}

// SetPduSessionId sets field value
func (o *AssignEbiFailed) SetPduSessionId(v int32) {
	o.PduSessionId = v
}

// GetFailedArpList returns the FailedArpList field value if set, zero value otherwise.
func (o *AssignEbiFailed) GetFailedArpList() []Arp {
	if o == nil || isNil(o.FailedArpList) {
		var ret []Arp
		return ret
	}
	return o.FailedArpList
}

// GetFailedArpListOk returns a tuple with the FailedArpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignEbiFailed) GetFailedArpListOk() ([]Arp, bool) {
	if o == nil || isNil(o.FailedArpList) {
		return nil, false
	}
	return o.FailedArpList, true
}

// HasFailedArpList returns a boolean if a field has been set.
func (o *AssignEbiFailed) HasFailedArpList() bool {
	if o != nil && !isNil(o.FailedArpList) {
		return true
	}

	return false
}

// SetFailedArpList gets a reference to the given []Arp and assigns it to the FailedArpList field.
func (o *AssignEbiFailed) SetFailedArpList(v []Arp) {
	o.FailedArpList = v
}

func (o AssignEbiFailed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignEbiFailed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pduSessionId"] = o.PduSessionId
	if !isNil(o.FailedArpList) {
		toSerialize["failedArpList"] = o.FailedArpList
	}
	return toSerialize, nil
}

type NullableAssignEbiFailed struct {
	value *AssignEbiFailed
	isSet bool
}

func (v NullableAssignEbiFailed) Get() *AssignEbiFailed {
	return v.value
}

func (v *NullableAssignEbiFailed) Set(val *AssignEbiFailed) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignEbiFailed) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignEbiFailed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignEbiFailed(val *AssignEbiFailed) *NullableAssignEbiFailed {
	return &NullableAssignEbiFailed{value: val, isSet: true}
}

func (v NullableAssignEbiFailed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignEbiFailed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



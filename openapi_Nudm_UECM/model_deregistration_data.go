/*
Nudm_UECM

Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UECM

import (
	"encoding/json"
)

// checks if the DeregistrationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeregistrationData{}

// DeregistrationData struct for DeregistrationData
type DeregistrationData struct {
	DeregReason DeregistrationReason `json:"deregReason"`
	AccessType *AccessType `json:"accessType,omitempty"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessionId *int32 `json:"pduSessionId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NewSmfInstanceId *string `json:"newSmfInstanceId,omitempty"`
}

// NewDeregistrationData instantiates a new DeregistrationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeregistrationData(deregReason DeregistrationReason) *DeregistrationData {
	this := DeregistrationData{}
	this.DeregReason = deregReason
	return &this
}

// NewDeregistrationDataWithDefaults instantiates a new DeregistrationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeregistrationDataWithDefaults() *DeregistrationData {
	this := DeregistrationData{}
	return &this
}

// GetDeregReason returns the DeregReason field value
func (o *DeregistrationData) GetDeregReason() DeregistrationReason {
	if o == nil {
		var ret DeregistrationReason
		return ret
	}

	return o.DeregReason
}

// GetDeregReasonOk returns a tuple with the DeregReason field value
// and a boolean to check if the value has been set.
func (o *DeregistrationData) GetDeregReasonOk() (*DeregistrationReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeregReason, true
}

// SetDeregReason sets field value
func (o *DeregistrationData) SetDeregReason(v DeregistrationReason) {
	o.DeregReason = v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *DeregistrationData) GetAccessType() AccessType {
	if o == nil || IsNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeregistrationData) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *DeregistrationData) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *DeregistrationData) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetPduSessionId returns the PduSessionId field value if set, zero value otherwise.
func (o *DeregistrationData) GetPduSessionId() int32 {
	if o == nil || IsNil(o.PduSessionId) {
		var ret int32
		return ret
	}
	return *o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeregistrationData) GetPduSessionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PduSessionId) {
		return nil, false
	}
	return o.PduSessionId, true
}

// HasPduSessionId returns a boolean if a field has been set.
func (o *DeregistrationData) HasPduSessionId() bool {
	if o != nil && !IsNil(o.PduSessionId) {
		return true
	}

	return false
}

// SetPduSessionId gets a reference to the given int32 and assigns it to the PduSessionId field.
func (o *DeregistrationData) SetPduSessionId(v int32) {
	o.PduSessionId = &v
}

// GetNewSmfInstanceId returns the NewSmfInstanceId field value if set, zero value otherwise.
func (o *DeregistrationData) GetNewSmfInstanceId() string {
	if o == nil || IsNil(o.NewSmfInstanceId) {
		var ret string
		return ret
	}
	return *o.NewSmfInstanceId
}

// GetNewSmfInstanceIdOk returns a tuple with the NewSmfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeregistrationData) GetNewSmfInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.NewSmfInstanceId) {
		return nil, false
	}
	return o.NewSmfInstanceId, true
}

// HasNewSmfInstanceId returns a boolean if a field has been set.
func (o *DeregistrationData) HasNewSmfInstanceId() bool {
	if o != nil && !IsNil(o.NewSmfInstanceId) {
		return true
	}

	return false
}

// SetNewSmfInstanceId gets a reference to the given string and assigns it to the NewSmfInstanceId field.
func (o *DeregistrationData) SetNewSmfInstanceId(v string) {
	o.NewSmfInstanceId = &v
}

func (o DeregistrationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeregistrationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deregReason"] = o.DeregReason
	if !IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !IsNil(o.PduSessionId) {
		toSerialize["pduSessionId"] = o.PduSessionId
	}
	if !IsNil(o.NewSmfInstanceId) {
		toSerialize["newSmfInstanceId"] = o.NewSmfInstanceId
	}
	return toSerialize, nil
}

type NullableDeregistrationData struct {
	value *DeregistrationData
	isSet bool
}

func (v NullableDeregistrationData) Get() *DeregistrationData {
	return v.value
}

func (v *NullableDeregistrationData) Set(val *DeregistrationData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeregistrationData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeregistrationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeregistrationData(val *DeregistrationData) *NullableDeregistrationData {
	return &NullableDeregistrationData{value: val, isSet: true}
}

func (v NullableDeregistrationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeregistrationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


/*
JOSE Protected Message Forwarding API

N32-f Message Forwarding Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_JOSEProtectedMessageForwarding

import (
	"encoding/json"
)

// checks if the N32fReformattedRspMsg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &N32fReformattedRspMsg{}

// N32fReformattedRspMsg Contains the reformatted HTTP/2 response message
type N32fReformattedRspMsg struct {
	ReformattedData    FlatJweJson   `json:"reformattedData"`
	ModificationsBlock []FlatJwsJson `json:"modificationsBlock,omitempty"`
}

// NewN32fReformattedRspMsg instantiates a new N32fReformattedRspMsg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN32fReformattedRspMsg(reformattedData FlatJweJson) *N32fReformattedRspMsg {
	this := N32fReformattedRspMsg{}
	this.ReformattedData = reformattedData
	return &this
}

// NewN32fReformattedRspMsgWithDefaults instantiates a new N32fReformattedRspMsg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN32fReformattedRspMsgWithDefaults() *N32fReformattedRspMsg {
	this := N32fReformattedRspMsg{}
	return &this
}

// GetReformattedData returns the ReformattedData field value
func (o *N32fReformattedRspMsg) GetReformattedData() FlatJweJson {
	if o == nil {
		var ret FlatJweJson
		return ret
	}

	return o.ReformattedData
}

// GetReformattedDataOk returns a tuple with the ReformattedData field value
// and a boolean to check if the value has been set.
func (o *N32fReformattedRspMsg) GetReformattedDataOk() (*FlatJweJson, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReformattedData, true
}

// SetReformattedData sets field value
func (o *N32fReformattedRspMsg) SetReformattedData(v FlatJweJson) {
	o.ReformattedData = v
}

// GetModificationsBlock returns the ModificationsBlock field value if set, zero value otherwise.
func (o *N32fReformattedRspMsg) GetModificationsBlock() []FlatJwsJson {
	if o == nil || IsNil(o.ModificationsBlock) {
		var ret []FlatJwsJson
		return ret
	}
	return o.ModificationsBlock
}

// GetModificationsBlockOk returns a tuple with the ModificationsBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N32fReformattedRspMsg) GetModificationsBlockOk() ([]FlatJwsJson, bool) {
	if o == nil || IsNil(o.ModificationsBlock) {
		return nil, false
	}
	return o.ModificationsBlock, true
}

// HasModificationsBlock returns a boolean if a field has been set.
func (o *N32fReformattedRspMsg) HasModificationsBlock() bool {
	if o != nil && !IsNil(o.ModificationsBlock) {
		return true
	}

	return false
}

// SetModificationsBlock gets a reference to the given []FlatJwsJson and assigns it to the ModificationsBlock field.
func (o *N32fReformattedRspMsg) SetModificationsBlock(v []FlatJwsJson) {
	o.ModificationsBlock = v
}

func (o N32fReformattedRspMsg) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o N32fReformattedRspMsg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reformattedData"] = o.ReformattedData
	if !IsNil(o.ModificationsBlock) {
		toSerialize["modificationsBlock"] = o.ModificationsBlock
	}
	return toSerialize, nil
}

type NullableN32fReformattedRspMsg struct {
	value *N32fReformattedRspMsg
	isSet bool
}

func (v NullableN32fReformattedRspMsg) Get() *N32fReformattedRspMsg {
	return v.value
}

func (v *NullableN32fReformattedRspMsg) Set(val *N32fReformattedRspMsg) {
	v.value = val
	v.isSet = true
}

func (v NullableN32fReformattedRspMsg) IsSet() bool {
	return v.isSet
}

func (v *NullableN32fReformattedRspMsg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN32fReformattedRspMsg(val *N32fReformattedRspMsg) *NullableN32fReformattedRspMsg {
	return &NullableN32fReformattedRspMsg{value: val, isSet: true}
}

func (v NullableN32fReformattedRspMsg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN32fReformattedRspMsg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

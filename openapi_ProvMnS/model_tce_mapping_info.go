/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the TceMappingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TceMappingInfo{}

// TceMappingInfo struct for TceMappingInfo
type TceMappingInfo struct {
	TceIPAddress *TceMappingInfoTceIPAddress `json:"TceIPAddress,omitempty"`
	TceID *int32 `json:"TceID,omitempty"`
	PlmnTarget *PlmnId `json:"PlmnTarget,omitempty"`
}

// NewTceMappingInfo instantiates a new TceMappingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTceMappingInfo() *TceMappingInfo {
	this := TceMappingInfo{}
	return &this
}

// NewTceMappingInfoWithDefaults instantiates a new TceMappingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTceMappingInfoWithDefaults() *TceMappingInfo {
	this := TceMappingInfo{}
	return &this
}

// GetTceIPAddress returns the TceIPAddress field value if set, zero value otherwise.
func (o *TceMappingInfo) GetTceIPAddress() TceMappingInfoTceIPAddress {
	if o == nil || IsNil(o.TceIPAddress) {
		var ret TceMappingInfoTceIPAddress
		return ret
	}
	return *o.TceIPAddress
}

// GetTceIPAddressOk returns a tuple with the TceIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TceMappingInfo) GetTceIPAddressOk() (*TceMappingInfoTceIPAddress, bool) {
	if o == nil || IsNil(o.TceIPAddress) {
		return nil, false
	}
	return o.TceIPAddress, true
}

// HasTceIPAddress returns a boolean if a field has been set.
func (o *TceMappingInfo) HasTceIPAddress() bool {
	if o != nil && !IsNil(o.TceIPAddress) {
		return true
	}

	return false
}

// SetTceIPAddress gets a reference to the given TceMappingInfoTceIPAddress and assigns it to the TceIPAddress field.
func (o *TceMappingInfo) SetTceIPAddress(v TceMappingInfoTceIPAddress) {
	o.TceIPAddress = &v
}

// GetTceID returns the TceID field value if set, zero value otherwise.
func (o *TceMappingInfo) GetTceID() int32 {
	if o == nil || IsNil(o.TceID) {
		var ret int32
		return ret
	}
	return *o.TceID
}

// GetTceIDOk returns a tuple with the TceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TceMappingInfo) GetTceIDOk() (*int32, bool) {
	if o == nil || IsNil(o.TceID) {
		return nil, false
	}
	return o.TceID, true
}

// HasTceID returns a boolean if a field has been set.
func (o *TceMappingInfo) HasTceID() bool {
	if o != nil && !IsNil(o.TceID) {
		return true
	}

	return false
}

// SetTceID gets a reference to the given int32 and assigns it to the TceID field.
func (o *TceMappingInfo) SetTceID(v int32) {
	o.TceID = &v
}

// GetPlmnTarget returns the PlmnTarget field value if set, zero value otherwise.
func (o *TceMappingInfo) GetPlmnTarget() PlmnId {
	if o == nil || IsNil(o.PlmnTarget) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnTarget
}

// GetPlmnTargetOk returns a tuple with the PlmnTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TceMappingInfo) GetPlmnTargetOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.PlmnTarget) {
		return nil, false
	}
	return o.PlmnTarget, true
}

// HasPlmnTarget returns a boolean if a field has been set.
func (o *TceMappingInfo) HasPlmnTarget() bool {
	if o != nil && !IsNil(o.PlmnTarget) {
		return true
	}

	return false
}

// SetPlmnTarget gets a reference to the given PlmnId and assigns it to the PlmnTarget field.
func (o *TceMappingInfo) SetPlmnTarget(v PlmnId) {
	o.PlmnTarget = &v
}

func (o TceMappingInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TceMappingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TceIPAddress) {
		toSerialize["TceIPAddress"] = o.TceIPAddress
	}
	if !IsNil(o.TceID) {
		toSerialize["TceID"] = o.TceID
	}
	if !IsNil(o.PlmnTarget) {
		toSerialize["PlmnTarget"] = o.PlmnTarget
	}
	return toSerialize, nil
}

type NullableTceMappingInfo struct {
	value *TceMappingInfo
	isSet bool
}

func (v NullableTceMappingInfo) Get() *TceMappingInfo {
	return v.value
}

func (v *NullableTceMappingInfo) Set(val *TceMappingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTceMappingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTceMappingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTceMappingInfo(val *TceMappingInfo) *NullableTceMappingInfo {
	return &NullableTceMappingInfo{value: val, isSet: true}
}

func (v NullableTceMappingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTceMappingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


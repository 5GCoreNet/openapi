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

// checks if the Model5GCNfConnEcmInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model5GCNfConnEcmInfo{}

// Model5GCNfConnEcmInfo Store the 5GC NF connection information
type Model5GCNfConnEcmInfo struct {
	Var5GCNFType *string `json:"5GCNFType,omitempty"`
	Var5GCNFIpAddress *string `json:"5GCNFIpAddress,omitempty"`
	Var5GCNFRef *string `json:"5GCNFRef,omitempty"`
}

// NewModel5GCNfConnEcmInfo instantiates a new Model5GCNfConnEcmInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel5GCNfConnEcmInfo() *Model5GCNfConnEcmInfo {
	this := Model5GCNfConnEcmInfo{}
	return &this
}

// NewModel5GCNfConnEcmInfoWithDefaults instantiates a new Model5GCNfConnEcmInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel5GCNfConnEcmInfoWithDefaults() *Model5GCNfConnEcmInfo {
	this := Model5GCNfConnEcmInfo{}
	return &this
}

// GetVar5GCNFType returns the Var5GCNFType field value if set, zero value otherwise.
func (o *Model5GCNfConnEcmInfo) GetVar5GCNFType() string {
	if o == nil || IsNil(o.Var5GCNFType) {
		var ret string
		return ret
	}
	return *o.Var5GCNFType
}

// GetVar5GCNFTypeOk returns a tuple with the Var5GCNFType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GCNfConnEcmInfo) GetVar5GCNFTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Var5GCNFType) {
		return nil, false
	}
	return o.Var5GCNFType, true
}

// HasVar5GCNFType returns a boolean if a field has been set.
func (o *Model5GCNfConnEcmInfo) HasVar5GCNFType() bool {
	if o != nil && !IsNil(o.Var5GCNFType) {
		return true
	}

	return false
}

// SetVar5GCNFType gets a reference to the given string and assigns it to the Var5GCNFType field.
func (o *Model5GCNfConnEcmInfo) SetVar5GCNFType(v string) {
	o.Var5GCNFType = &v
}

// GetVar5GCNFIpAddress returns the Var5GCNFIpAddress field value if set, zero value otherwise.
func (o *Model5GCNfConnEcmInfo) GetVar5GCNFIpAddress() string {
	if o == nil || IsNil(o.Var5GCNFIpAddress) {
		var ret string
		return ret
	}
	return *o.Var5GCNFIpAddress
}

// GetVar5GCNFIpAddressOk returns a tuple with the Var5GCNFIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GCNfConnEcmInfo) GetVar5GCNFIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Var5GCNFIpAddress) {
		return nil, false
	}
	return o.Var5GCNFIpAddress, true
}

// HasVar5GCNFIpAddress returns a boolean if a field has been set.
func (o *Model5GCNfConnEcmInfo) HasVar5GCNFIpAddress() bool {
	if o != nil && !IsNil(o.Var5GCNFIpAddress) {
		return true
	}

	return false
}

// SetVar5GCNFIpAddress gets a reference to the given string and assigns it to the Var5GCNFIpAddress field.
func (o *Model5GCNfConnEcmInfo) SetVar5GCNFIpAddress(v string) {
	o.Var5GCNFIpAddress = &v
}

// GetVar5GCNFRef returns the Var5GCNFRef field value if set, zero value otherwise.
func (o *Model5GCNfConnEcmInfo) GetVar5GCNFRef() string {
	if o == nil || IsNil(o.Var5GCNFRef) {
		var ret string
		return ret
	}
	return *o.Var5GCNFRef
}

// GetVar5GCNFRefOk returns a tuple with the Var5GCNFRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GCNfConnEcmInfo) GetVar5GCNFRefOk() (*string, bool) {
	if o == nil || IsNil(o.Var5GCNFRef) {
		return nil, false
	}
	return o.Var5GCNFRef, true
}

// HasVar5GCNFRef returns a boolean if a field has been set.
func (o *Model5GCNfConnEcmInfo) HasVar5GCNFRef() bool {
	if o != nil && !IsNil(o.Var5GCNFRef) {
		return true
	}

	return false
}

// SetVar5GCNFRef gets a reference to the given string and assigns it to the Var5GCNFRef field.
func (o *Model5GCNfConnEcmInfo) SetVar5GCNFRef(v string) {
	o.Var5GCNFRef = &v
}

func (o Model5GCNfConnEcmInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model5GCNfConnEcmInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var5GCNFType) {
		toSerialize["5GCNFType"] = o.Var5GCNFType
	}
	if !IsNil(o.Var5GCNFIpAddress) {
		toSerialize["5GCNFIpAddress"] = o.Var5GCNFIpAddress
	}
	if !IsNil(o.Var5GCNFRef) {
		toSerialize["5GCNFRef"] = o.Var5GCNFRef
	}
	return toSerialize, nil
}

type NullableModel5GCNfConnEcmInfo struct {
	value *Model5GCNfConnEcmInfo
	isSet bool
}

func (v NullableModel5GCNfConnEcmInfo) Get() *Model5GCNfConnEcmInfo {
	return v.value
}

func (v *NullableModel5GCNfConnEcmInfo) Set(val *Model5GCNfConnEcmInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModel5GCNfConnEcmInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModel5GCNfConnEcmInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel5GCNfConnEcmInfo(val *Model5GCNfConnEcmInfo) *NullableModel5GCNfConnEcmInfo {
	return &NullableModel5GCNfConnEcmInfo{value: val, isSet: true}
}

func (v NullableModel5GCNfConnEcmInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel5GCNfConnEcmInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



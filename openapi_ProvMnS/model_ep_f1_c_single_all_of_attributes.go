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

// checks if the EPF1CSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPF1CSingleAllOfAttributes{}

// EPF1CSingleAllOfAttributes struct for EPF1CSingleAllOfAttributes
type EPF1CSingleAllOfAttributes struct {
	EPRPAttr
	LocalAddress  *LocalAddress  `json:"localAddress,omitempty"`
	RemoteAddress *RemoteAddress `json:"remoteAddress,omitempty"`
}

// NewEPF1CSingleAllOfAttributes instantiates a new EPF1CSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPF1CSingleAllOfAttributes() *EPF1CSingleAllOfAttributes {
	this := EPF1CSingleAllOfAttributes{}
	return &this
}

// NewEPF1CSingleAllOfAttributesWithDefaults instantiates a new EPF1CSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPF1CSingleAllOfAttributesWithDefaults() *EPF1CSingleAllOfAttributes {
	this := EPF1CSingleAllOfAttributes{}
	return &this
}

// GetLocalAddress returns the LocalAddress field value if set, zero value otherwise.
func (o *EPF1CSingleAllOfAttributes) GetLocalAddress() LocalAddress {
	if o == nil || IsNil(o.LocalAddress) {
		var ret LocalAddress
		return ret
	}
	return *o.LocalAddress
}

// GetLocalAddressOk returns a tuple with the LocalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPF1CSingleAllOfAttributes) GetLocalAddressOk() (*LocalAddress, bool) {
	if o == nil || IsNil(o.LocalAddress) {
		return nil, false
	}
	return o.LocalAddress, true
}

// HasLocalAddress returns a boolean if a field has been set.
func (o *EPF1CSingleAllOfAttributes) HasLocalAddress() bool {
	if o != nil && !IsNil(o.LocalAddress) {
		return true
	}

	return false
}

// SetLocalAddress gets a reference to the given LocalAddress and assigns it to the LocalAddress field.
func (o *EPF1CSingleAllOfAttributes) SetLocalAddress(v LocalAddress) {
	o.LocalAddress = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *EPF1CSingleAllOfAttributes) GetRemoteAddress() RemoteAddress {
	if o == nil || IsNil(o.RemoteAddress) {
		var ret RemoteAddress
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPF1CSingleAllOfAttributes) GetRemoteAddressOk() (*RemoteAddress, bool) {
	if o == nil || IsNil(o.RemoteAddress) {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *EPF1CSingleAllOfAttributes) HasRemoteAddress() bool {
	if o != nil && !IsNil(o.RemoteAddress) {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given RemoteAddress and assigns it to the RemoteAddress field.
func (o *EPF1CSingleAllOfAttributes) SetRemoteAddress(v RemoteAddress) {
	o.RemoteAddress = &v
}

func (o EPF1CSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPF1CSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEPRPAttr, errEPRPAttr := json.Marshal(o.EPRPAttr)
	if errEPRPAttr != nil {
		return map[string]interface{}{}, errEPRPAttr
	}
	errEPRPAttr = json.Unmarshal([]byte(serializedEPRPAttr), &toSerialize)
	if errEPRPAttr != nil {
		return map[string]interface{}{}, errEPRPAttr
	}
	if !IsNil(o.LocalAddress) {
		toSerialize["localAddress"] = o.LocalAddress
	}
	if !IsNil(o.RemoteAddress) {
		toSerialize["remoteAddress"] = o.RemoteAddress
	}
	return toSerialize, nil
}

type NullableEPF1CSingleAllOfAttributes struct {
	value *EPF1CSingleAllOfAttributes
	isSet bool
}

func (v NullableEPF1CSingleAllOfAttributes) Get() *EPF1CSingleAllOfAttributes {
	return v.value
}

func (v *NullableEPF1CSingleAllOfAttributes) Set(val *EPF1CSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEPF1CSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEPF1CSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPF1CSingleAllOfAttributes(val *EPF1CSingleAllOfAttributes) *NullableEPF1CSingleAllOfAttributes {
	return &NullableEPF1CSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableEPF1CSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPF1CSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

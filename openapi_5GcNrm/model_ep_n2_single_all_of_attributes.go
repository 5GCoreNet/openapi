/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the EPN2SingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPN2SingleAllOfAttributes{}

// EPN2SingleAllOfAttributes struct for EPN2SingleAllOfAttributes
type EPN2SingleAllOfAttributes struct {
	EPRPAttr
	LocalAddress  *LocalAddress  `json:"localAddress,omitempty"`
	RemoteAddress *RemoteAddress `json:"remoteAddress,omitempty"`
}

// NewEPN2SingleAllOfAttributes instantiates a new EPN2SingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPN2SingleAllOfAttributes() *EPN2SingleAllOfAttributes {
	this := EPN2SingleAllOfAttributes{}
	return &this
}

// NewEPN2SingleAllOfAttributesWithDefaults instantiates a new EPN2SingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPN2SingleAllOfAttributesWithDefaults() *EPN2SingleAllOfAttributes {
	this := EPN2SingleAllOfAttributes{}
	return &this
}

// GetLocalAddress returns the LocalAddress field value if set, zero value otherwise.
func (o *EPN2SingleAllOfAttributes) GetLocalAddress() LocalAddress {
	if o == nil || IsNil(o.LocalAddress) {
		var ret LocalAddress
		return ret
	}
	return *o.LocalAddress
}

// GetLocalAddressOk returns a tuple with the LocalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPN2SingleAllOfAttributes) GetLocalAddressOk() (*LocalAddress, bool) {
	if o == nil || IsNil(o.LocalAddress) {
		return nil, false
	}
	return o.LocalAddress, true
}

// HasLocalAddress returns a boolean if a field has been set.
func (o *EPN2SingleAllOfAttributes) HasLocalAddress() bool {
	if o != nil && !IsNil(o.LocalAddress) {
		return true
	}

	return false
}

// SetLocalAddress gets a reference to the given LocalAddress and assigns it to the LocalAddress field.
func (o *EPN2SingleAllOfAttributes) SetLocalAddress(v LocalAddress) {
	o.LocalAddress = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *EPN2SingleAllOfAttributes) GetRemoteAddress() RemoteAddress {
	if o == nil || IsNil(o.RemoteAddress) {
		var ret RemoteAddress
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPN2SingleAllOfAttributes) GetRemoteAddressOk() (*RemoteAddress, bool) {
	if o == nil || IsNil(o.RemoteAddress) {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *EPN2SingleAllOfAttributes) HasRemoteAddress() bool {
	if o != nil && !IsNil(o.RemoteAddress) {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given RemoteAddress and assigns it to the RemoteAddress field.
func (o *EPN2SingleAllOfAttributes) SetRemoteAddress(v RemoteAddress) {
	o.RemoteAddress = &v
}

func (o EPN2SingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPN2SingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
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

type NullableEPN2SingleAllOfAttributes struct {
	value *EPN2SingleAllOfAttributes
	isSet bool
}

func (v NullableEPN2SingleAllOfAttributes) Get() *EPN2SingleAllOfAttributes {
	return v.value
}

func (v *NullableEPN2SingleAllOfAttributes) Set(val *EPN2SingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEPN2SingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEPN2SingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPN2SingleAllOfAttributes(val *EPN2SingleAllOfAttributes) *NullableEPN2SingleAllOfAttributes {
	return &NullableEPN2SingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableEPN2SingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPN2SingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the PublicIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicIdentity{}

// PublicIdentity IMS Public Identity and related data (Alias Group Id, IRS default indication, Identity Type)
type PublicIdentity struct {
	// String containing an IMS Public Identity in SIP URI format or TEL URI format
	ImsPublicId  string       `json:"imsPublicId"`
	IdentityType IdentityType `json:"identityType"`
	IrsIsDefault *bool        `json:"irsIsDefault,omitempty"`
	AliasGroupId *string      `json:"aliasGroupId,omitempty"`
}

// NewPublicIdentity instantiates a new PublicIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIdentity(imsPublicId string, identityType IdentityType) *PublicIdentity {
	this := PublicIdentity{}
	this.ImsPublicId = imsPublicId
	this.IdentityType = identityType
	return &this
}

// NewPublicIdentityWithDefaults instantiates a new PublicIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIdentityWithDefaults() *PublicIdentity {
	this := PublicIdentity{}
	return &this
}

// GetImsPublicId returns the ImsPublicId field value
func (o *PublicIdentity) GetImsPublicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImsPublicId
}

// GetImsPublicIdOk returns a tuple with the ImsPublicId field value
// and a boolean to check if the value has been set.
func (o *PublicIdentity) GetImsPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImsPublicId, true
}

// SetImsPublicId sets field value
func (o *PublicIdentity) SetImsPublicId(v string) {
	o.ImsPublicId = v
}

// GetIdentityType returns the IdentityType field value
func (o *PublicIdentity) GetIdentityType() IdentityType {
	if o == nil {
		var ret IdentityType
		return ret
	}

	return o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value
// and a boolean to check if the value has been set.
func (o *PublicIdentity) GetIdentityTypeOk() (*IdentityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityType, true
}

// SetIdentityType sets field value
func (o *PublicIdentity) SetIdentityType(v IdentityType) {
	o.IdentityType = v
}

// GetIrsIsDefault returns the IrsIsDefault field value if set, zero value otherwise.
func (o *PublicIdentity) GetIrsIsDefault() bool {
	if o == nil || IsNil(o.IrsIsDefault) {
		var ret bool
		return ret
	}
	return *o.IrsIsDefault
}

// GetIrsIsDefaultOk returns a tuple with the IrsIsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentity) GetIrsIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IrsIsDefault) {
		return nil, false
	}
	return o.IrsIsDefault, true
}

// HasIrsIsDefault returns a boolean if a field has been set.
func (o *PublicIdentity) HasIrsIsDefault() bool {
	if o != nil && !IsNil(o.IrsIsDefault) {
		return true
	}

	return false
}

// SetIrsIsDefault gets a reference to the given bool and assigns it to the IrsIsDefault field.
func (o *PublicIdentity) SetIrsIsDefault(v bool) {
	o.IrsIsDefault = &v
}

// GetAliasGroupId returns the AliasGroupId field value if set, zero value otherwise.
func (o *PublicIdentity) GetAliasGroupId() string {
	if o == nil || IsNil(o.AliasGroupId) {
		var ret string
		return ret
	}
	return *o.AliasGroupId
}

// GetAliasGroupIdOk returns a tuple with the AliasGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentity) GetAliasGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AliasGroupId) {
		return nil, false
	}
	return o.AliasGroupId, true
}

// HasAliasGroupId returns a boolean if a field has been set.
func (o *PublicIdentity) HasAliasGroupId() bool {
	if o != nil && !IsNil(o.AliasGroupId) {
		return true
	}

	return false
}

// SetAliasGroupId gets a reference to the given string and assigns it to the AliasGroupId field.
func (o *PublicIdentity) SetAliasGroupId(v string) {
	o.AliasGroupId = &v
}

func (o PublicIdentity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["imsPublicId"] = o.ImsPublicId
	toSerialize["identityType"] = o.IdentityType
	if !IsNil(o.IrsIsDefault) {
		toSerialize["irsIsDefault"] = o.IrsIsDefault
	}
	if !IsNil(o.AliasGroupId) {
		toSerialize["aliasGroupId"] = o.AliasGroupId
	}
	return toSerialize, nil
}

type NullablePublicIdentity struct {
	value *PublicIdentity
	isSet bool
}

func (v NullablePublicIdentity) Get() *PublicIdentity {
	return v.value
}

func (v *NullablePublicIdentity) Set(val *PublicIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicIdentity(val *PublicIdentity) *NullablePublicIdentity {
	return &NullablePublicIdentity{value: val, isSet: true}
}

func (v NullablePublicIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

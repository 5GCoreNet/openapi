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

// checks if the UdsfFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UdsfFunctionSingleAllOfAttributesAllOf{}

// UdsfFunctionSingleAllOfAttributesAllOf struct for UdsfFunctionSingleAllOfAttributesAllOf
type UdsfFunctionSingleAllOfAttributesAllOf struct {
	PlmnInfoList     []PlmnInfo        `json:"plmnInfoList,omitempty"`
	SBIFqdn          *string           `json:"sBIFqdn,omitempty"`
	ManagedNFProfile *ManagedNFProfile `json:"managedNFProfile,omitempty"`
	UdsfInfo         *UdsfInfo         `json:"udsfInfo,omitempty"`
}

// NewUdsfFunctionSingleAllOfAttributesAllOf instantiates a new UdsfFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdsfFunctionSingleAllOfAttributesAllOf() *UdsfFunctionSingleAllOfAttributesAllOf {
	this := UdsfFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewUdsfFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new UdsfFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdsfFunctionSingleAllOfAttributesAllOfWithDefaults() *UdsfFunctionSingleAllOfAttributesAllOf {
	this := UdsfFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetPlmnInfoList returns the PlmnInfoList field value if set, zero value otherwise.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) GetPlmnInfoList() []PlmnInfo {
	if o == nil || IsNil(o.PlmnInfoList) {
		var ret []PlmnInfo
		return ret
	}
	return o.PlmnInfoList
}

// GetPlmnInfoListOk returns a tuple with the PlmnInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) GetPlmnInfoListOk() ([]PlmnInfo, bool) {
	if o == nil || IsNil(o.PlmnInfoList) {
		return nil, false
	}
	return o.PlmnInfoList, true
}

// HasPlmnInfoList returns a boolean if a field has been set.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) HasPlmnInfoList() bool {
	if o != nil && !IsNil(o.PlmnInfoList) {
		return true
	}

	return false
}

// SetPlmnInfoList gets a reference to the given []PlmnInfo and assigns it to the PlmnInfoList field.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) SetPlmnInfoList(v []PlmnInfo) {
	o.PlmnInfoList = v
}

// GetSBIFqdn returns the SBIFqdn field value if set, zero value otherwise.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) GetSBIFqdn() string {
	if o == nil || IsNil(o.SBIFqdn) {
		var ret string
		return ret
	}
	return *o.SBIFqdn
}

// GetSBIFqdnOk returns a tuple with the SBIFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) GetSBIFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.SBIFqdn) {
		return nil, false
	}
	return o.SBIFqdn, true
}

// HasSBIFqdn returns a boolean if a field has been set.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) HasSBIFqdn() bool {
	if o != nil && !IsNil(o.SBIFqdn) {
		return true
	}

	return false
}

// SetSBIFqdn gets a reference to the given string and assigns it to the SBIFqdn field.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) SetSBIFqdn(v string) {
	o.SBIFqdn = &v
}

// GetManagedNFProfile returns the ManagedNFProfile field value if set, zero value otherwise.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) GetManagedNFProfile() ManagedNFProfile {
	if o == nil || IsNil(o.ManagedNFProfile) {
		var ret ManagedNFProfile
		return ret
	}
	return *o.ManagedNFProfile
}

// GetManagedNFProfileOk returns a tuple with the ManagedNFProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) GetManagedNFProfileOk() (*ManagedNFProfile, bool) {
	if o == nil || IsNil(o.ManagedNFProfile) {
		return nil, false
	}
	return o.ManagedNFProfile, true
}

// HasManagedNFProfile returns a boolean if a field has been set.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) HasManagedNFProfile() bool {
	if o != nil && !IsNil(o.ManagedNFProfile) {
		return true
	}

	return false
}

// SetManagedNFProfile gets a reference to the given ManagedNFProfile and assigns it to the ManagedNFProfile field.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) SetManagedNFProfile(v ManagedNFProfile) {
	o.ManagedNFProfile = &v
}

// GetUdsfInfo returns the UdsfInfo field value if set, zero value otherwise.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) GetUdsfInfo() UdsfInfo {
	if o == nil || IsNil(o.UdsfInfo) {
		var ret UdsfInfo
		return ret
	}
	return *o.UdsfInfo
}

// GetUdsfInfoOk returns a tuple with the UdsfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) GetUdsfInfoOk() (*UdsfInfo, bool) {
	if o == nil || IsNil(o.UdsfInfo) {
		return nil, false
	}
	return o.UdsfInfo, true
}

// HasUdsfInfo returns a boolean if a field has been set.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) HasUdsfInfo() bool {
	if o != nil && !IsNil(o.UdsfInfo) {
		return true
	}

	return false
}

// SetUdsfInfo gets a reference to the given UdsfInfo and assigns it to the UdsfInfo field.
func (o *UdsfFunctionSingleAllOfAttributesAllOf) SetUdsfInfo(v UdsfInfo) {
	o.UdsfInfo = &v
}

func (o UdsfFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UdsfFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlmnInfoList) {
		toSerialize["plmnInfoList"] = o.PlmnInfoList
	}
	if !IsNil(o.SBIFqdn) {
		toSerialize["sBIFqdn"] = o.SBIFqdn
	}
	if !IsNil(o.ManagedNFProfile) {
		toSerialize["managedNFProfile"] = o.ManagedNFProfile
	}
	if !IsNil(o.UdsfInfo) {
		toSerialize["udsfInfo"] = o.UdsfInfo
	}
	return toSerialize, nil
}

type NullableUdsfFunctionSingleAllOfAttributesAllOf struct {
	value *UdsfFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableUdsfFunctionSingleAllOfAttributesAllOf) Get() *UdsfFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableUdsfFunctionSingleAllOfAttributesAllOf) Set(val *UdsfFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUdsfFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUdsfFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdsfFunctionSingleAllOfAttributesAllOf(val *UdsfFunctionSingleAllOfAttributesAllOf) *NullableUdsfFunctionSingleAllOfAttributesAllOf {
	return &NullableUdsfFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableUdsfFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdsfFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

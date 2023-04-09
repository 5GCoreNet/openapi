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

// checks if the NetworkSliceSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkSliceSingleAllOfAttributes{}

// NetworkSliceSingleAllOfAttributes struct for NetworkSliceSingleAllOfAttributes
type NetworkSliceSingleAllOfAttributes struct {
	NetworkSliceSubnetRef *string              `json:"networkSliceSubnetRef,omitempty"`
	OperationalState      *OperationalState    `json:"operationalState,omitempty"`
	AdministrativeState   *AdministrativeState `json:"administrativeState,omitempty"`
	ServiceProfileList    []ServiceProfile     `json:"serviceProfileList,omitempty"`
}

// NewNetworkSliceSingleAllOfAttributes instantiates a new NetworkSliceSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSliceSingleAllOfAttributes() *NetworkSliceSingleAllOfAttributes {
	this := NetworkSliceSingleAllOfAttributes{}
	return &this
}

// NewNetworkSliceSingleAllOfAttributesWithDefaults instantiates a new NetworkSliceSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSliceSingleAllOfAttributesWithDefaults() *NetworkSliceSingleAllOfAttributes {
	this := NetworkSliceSingleAllOfAttributes{}
	return &this
}

// GetNetworkSliceSubnetRef returns the NetworkSliceSubnetRef field value if set, zero value otherwise.
func (o *NetworkSliceSingleAllOfAttributes) GetNetworkSliceSubnetRef() string {
	if o == nil || IsNil(o.NetworkSliceSubnetRef) {
		var ret string
		return ret
	}
	return *o.NetworkSliceSubnetRef
}

// GetNetworkSliceSubnetRefOk returns a tuple with the NetworkSliceSubnetRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSingleAllOfAttributes) GetNetworkSliceSubnetRefOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkSliceSubnetRef) {
		return nil, false
	}
	return o.NetworkSliceSubnetRef, true
}

// HasNetworkSliceSubnetRef returns a boolean if a field has been set.
func (o *NetworkSliceSingleAllOfAttributes) HasNetworkSliceSubnetRef() bool {
	if o != nil && !IsNil(o.NetworkSliceSubnetRef) {
		return true
	}

	return false
}

// SetNetworkSliceSubnetRef gets a reference to the given string and assigns it to the NetworkSliceSubnetRef field.
func (o *NetworkSliceSingleAllOfAttributes) SetNetworkSliceSubnetRef(v string) {
	o.NetworkSliceSubnetRef = &v
}

// GetOperationalState returns the OperationalState field value if set, zero value otherwise.
func (o *NetworkSliceSingleAllOfAttributes) GetOperationalState() OperationalState {
	if o == nil || IsNil(o.OperationalState) {
		var ret OperationalState
		return ret
	}
	return *o.OperationalState
}

// GetOperationalStateOk returns a tuple with the OperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSingleAllOfAttributes) GetOperationalStateOk() (*OperationalState, bool) {
	if o == nil || IsNil(o.OperationalState) {
		return nil, false
	}
	return o.OperationalState, true
}

// HasOperationalState returns a boolean if a field has been set.
func (o *NetworkSliceSingleAllOfAttributes) HasOperationalState() bool {
	if o != nil && !IsNil(o.OperationalState) {
		return true
	}

	return false
}

// SetOperationalState gets a reference to the given OperationalState and assigns it to the OperationalState field.
func (o *NetworkSliceSingleAllOfAttributes) SetOperationalState(v OperationalState) {
	o.OperationalState = &v
}

// GetAdministrativeState returns the AdministrativeState field value if set, zero value otherwise.
func (o *NetworkSliceSingleAllOfAttributes) GetAdministrativeState() AdministrativeState {
	if o == nil || IsNil(o.AdministrativeState) {
		var ret AdministrativeState
		return ret
	}
	return *o.AdministrativeState
}

// GetAdministrativeStateOk returns a tuple with the AdministrativeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSingleAllOfAttributes) GetAdministrativeStateOk() (*AdministrativeState, bool) {
	if o == nil || IsNil(o.AdministrativeState) {
		return nil, false
	}
	return o.AdministrativeState, true
}

// HasAdministrativeState returns a boolean if a field has been set.
func (o *NetworkSliceSingleAllOfAttributes) HasAdministrativeState() bool {
	if o != nil && !IsNil(o.AdministrativeState) {
		return true
	}

	return false
}

// SetAdministrativeState gets a reference to the given AdministrativeState and assigns it to the AdministrativeState field.
func (o *NetworkSliceSingleAllOfAttributes) SetAdministrativeState(v AdministrativeState) {
	o.AdministrativeState = &v
}

// GetServiceProfileList returns the ServiceProfileList field value if set, zero value otherwise.
func (o *NetworkSliceSingleAllOfAttributes) GetServiceProfileList() []ServiceProfile {
	if o == nil || IsNil(o.ServiceProfileList) {
		var ret []ServiceProfile
		return ret
	}
	return o.ServiceProfileList
}

// GetServiceProfileListOk returns a tuple with the ServiceProfileList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSingleAllOfAttributes) GetServiceProfileListOk() ([]ServiceProfile, bool) {
	if o == nil || IsNil(o.ServiceProfileList) {
		return nil, false
	}
	return o.ServiceProfileList, true
}

// HasServiceProfileList returns a boolean if a field has been set.
func (o *NetworkSliceSingleAllOfAttributes) HasServiceProfileList() bool {
	if o != nil && !IsNil(o.ServiceProfileList) {
		return true
	}

	return false
}

// SetServiceProfileList gets a reference to the given []ServiceProfile and assigns it to the ServiceProfileList field.
func (o *NetworkSliceSingleAllOfAttributes) SetServiceProfileList(v []ServiceProfile) {
	o.ServiceProfileList = v
}

func (o NetworkSliceSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkSliceSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkSliceSubnetRef) {
		toSerialize["networkSliceSubnetRef"] = o.NetworkSliceSubnetRef
	}
	if !IsNil(o.OperationalState) {
		toSerialize["operationalState"] = o.OperationalState
	}
	if !IsNil(o.AdministrativeState) {
		toSerialize["administrativeState"] = o.AdministrativeState
	}
	if !IsNil(o.ServiceProfileList) {
		toSerialize["serviceProfileList"] = o.ServiceProfileList
	}
	return toSerialize, nil
}

type NullableNetworkSliceSingleAllOfAttributes struct {
	value *NetworkSliceSingleAllOfAttributes
	isSet bool
}

func (v NullableNetworkSliceSingleAllOfAttributes) Get() *NetworkSliceSingleAllOfAttributes {
	return v.value
}

func (v *NullableNetworkSliceSingleAllOfAttributes) Set(val *NetworkSliceSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSliceSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSliceSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSliceSingleAllOfAttributes(val *NetworkSliceSingleAllOfAttributes) *NullableNetworkSliceSingleAllOfAttributes {
	return &NullableNetworkSliceSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableNetworkSliceSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSliceSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

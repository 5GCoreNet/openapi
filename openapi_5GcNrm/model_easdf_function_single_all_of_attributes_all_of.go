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

// checks if the EASDFFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EASDFFunctionSingleAllOfAttributesAllOf{}

// EASDFFunctionSingleAllOfAttributesAllOf struct for EASDFFunctionSingleAllOfAttributesAllOf
type EASDFFunctionSingleAllOfAttributesAllOf struct {
	PlmnId           *PlmnId           `json:"plmnId,omitempty"`
	SBIFqdn          *string           `json:"sBIFqdn,omitempty"`
	ManagedNFProfile *ManagedNFProfile `json:"managedNFProfile,omitempty"`
	ServerAddr       *string           `json:"serverAddr,omitempty"`
}

// NewEASDFFunctionSingleAllOfAttributesAllOf instantiates a new EASDFFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEASDFFunctionSingleAllOfAttributesAllOf() *EASDFFunctionSingleAllOfAttributesAllOf {
	this := EASDFFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewEASDFFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new EASDFFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEASDFFunctionSingleAllOfAttributesAllOfWithDefaults() *EASDFFunctionSingleAllOfAttributesAllOf {
	this := EASDFFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) GetPlmnId() PlmnId {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetSBIFqdn returns the SBIFqdn field value if set, zero value otherwise.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) GetSBIFqdn() string {
	if o == nil || IsNil(o.SBIFqdn) {
		var ret string
		return ret
	}
	return *o.SBIFqdn
}

// GetSBIFqdnOk returns a tuple with the SBIFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) GetSBIFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.SBIFqdn) {
		return nil, false
	}
	return o.SBIFqdn, true
}

// HasSBIFqdn returns a boolean if a field has been set.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) HasSBIFqdn() bool {
	if o != nil && !IsNil(o.SBIFqdn) {
		return true
	}

	return false
}

// SetSBIFqdn gets a reference to the given string and assigns it to the SBIFqdn field.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) SetSBIFqdn(v string) {
	o.SBIFqdn = &v
}

// GetManagedNFProfile returns the ManagedNFProfile field value if set, zero value otherwise.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) GetManagedNFProfile() ManagedNFProfile {
	if o == nil || IsNil(o.ManagedNFProfile) {
		var ret ManagedNFProfile
		return ret
	}
	return *o.ManagedNFProfile
}

// GetManagedNFProfileOk returns a tuple with the ManagedNFProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) GetManagedNFProfileOk() (*ManagedNFProfile, bool) {
	if o == nil || IsNil(o.ManagedNFProfile) {
		return nil, false
	}
	return o.ManagedNFProfile, true
}

// HasManagedNFProfile returns a boolean if a field has been set.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) HasManagedNFProfile() bool {
	if o != nil && !IsNil(o.ManagedNFProfile) {
		return true
	}

	return false
}

// SetManagedNFProfile gets a reference to the given ManagedNFProfile and assigns it to the ManagedNFProfile field.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) SetManagedNFProfile(v ManagedNFProfile) {
	o.ManagedNFProfile = &v
}

// GetServerAddr returns the ServerAddr field value if set, zero value otherwise.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) GetServerAddr() string {
	if o == nil || IsNil(o.ServerAddr) {
		var ret string
		return ret
	}
	return *o.ServerAddr
}

// GetServerAddrOk returns a tuple with the ServerAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) GetServerAddrOk() (*string, bool) {
	if o == nil || IsNil(o.ServerAddr) {
		return nil, false
	}
	return o.ServerAddr, true
}

// HasServerAddr returns a boolean if a field has been set.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) HasServerAddr() bool {
	if o != nil && !IsNil(o.ServerAddr) {
		return true
	}

	return false
}

// SetServerAddr gets a reference to the given string and assigns it to the ServerAddr field.
func (o *EASDFFunctionSingleAllOfAttributesAllOf) SetServerAddr(v string) {
	o.ServerAddr = &v
}

func (o EASDFFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EASDFFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !IsNil(o.SBIFqdn) {
		toSerialize["sBIFqdn"] = o.SBIFqdn
	}
	if !IsNil(o.ManagedNFProfile) {
		toSerialize["managedNFProfile"] = o.ManagedNFProfile
	}
	if !IsNil(o.ServerAddr) {
		toSerialize["serverAddr"] = o.ServerAddr
	}
	return toSerialize, nil
}

type NullableEASDFFunctionSingleAllOfAttributesAllOf struct {
	value *EASDFFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableEASDFFunctionSingleAllOfAttributesAllOf) Get() *EASDFFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableEASDFFunctionSingleAllOfAttributesAllOf) Set(val *EASDFFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEASDFFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEASDFFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEASDFFunctionSingleAllOfAttributesAllOf(val *EASDFFunctionSingleAllOfAttributesAllOf) *NullableEASDFFunctionSingleAllOfAttributesAllOf {
	return &NullableEASDFFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableEASDFFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEASDFFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

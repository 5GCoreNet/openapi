/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
)

// checks if the EpsIwkPgw type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EpsIwkPgw{}

// EpsIwkPgw struct for EpsIwkPgw
type EpsIwkPgw struct {
	// Fully Qualified Domain Name
	PgwFqdn string `json:"pgwFqdn"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SmfInstanceId string  `json:"smfInstanceId"`
	PlmnId        *PlmnId `json:"plmnId,omitempty"`
}

// NewEpsIwkPgw instantiates a new EpsIwkPgw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpsIwkPgw(pgwFqdn string, smfInstanceId string) *EpsIwkPgw {
	this := EpsIwkPgw{}
	this.PgwFqdn = pgwFqdn
	this.SmfInstanceId = smfInstanceId
	return &this
}

// NewEpsIwkPgwWithDefaults instantiates a new EpsIwkPgw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpsIwkPgwWithDefaults() *EpsIwkPgw {
	this := EpsIwkPgw{}
	return &this
}

// GetPgwFqdn returns the PgwFqdn field value
func (o *EpsIwkPgw) GetPgwFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PgwFqdn
}

// GetPgwFqdnOk returns a tuple with the PgwFqdn field value
// and a boolean to check if the value has been set.
func (o *EpsIwkPgw) GetPgwFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PgwFqdn, true
}

// SetPgwFqdn sets field value
func (o *EpsIwkPgw) SetPgwFqdn(v string) {
	o.PgwFqdn = v
}

// GetSmfInstanceId returns the SmfInstanceId field value
func (o *EpsIwkPgw) GetSmfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmfInstanceId
}

// GetSmfInstanceIdOk returns a tuple with the SmfInstanceId field value
// and a boolean to check if the value has been set.
func (o *EpsIwkPgw) GetSmfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmfInstanceId, true
}

// SetSmfInstanceId sets field value
func (o *EpsIwkPgw) SetSmfInstanceId(v string) {
	o.SmfInstanceId = v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *EpsIwkPgw) GetPlmnId() PlmnId {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpsIwkPgw) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *EpsIwkPgw) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *EpsIwkPgw) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

func (o EpsIwkPgw) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EpsIwkPgw) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pgwFqdn"] = o.PgwFqdn
	toSerialize["smfInstanceId"] = o.SmfInstanceId
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	return toSerialize, nil
}

type NullableEpsIwkPgw struct {
	value *EpsIwkPgw
	isSet bool
}

func (v NullableEpsIwkPgw) Get() *EpsIwkPgw {
	return v.value
}

func (v *NullableEpsIwkPgw) Set(val *EpsIwkPgw) {
	v.value = val
	v.isSet = true
}

func (v NullableEpsIwkPgw) IsSet() bool {
	return v.isSet
}

func (v *NullableEpsIwkPgw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpsIwkPgw(val *EpsIwkPgw) *NullableEpsIwkPgw {
	return &NullableEpsIwkPgw{value: val, isSet: true}
}

func (v NullableEpsIwkPgw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpsIwkPgw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

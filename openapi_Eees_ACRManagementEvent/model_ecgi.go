/*
EES ACR Management Event_API

API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRManagementEvent

import (
	"encoding/json"
)

// checks if the Ecgi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ecgi{}

// Ecgi Contains the ECGI (E-UTRAN Cell Global Identity), as described in 3GPP 23.003
type Ecgi struct {
	PlmnId PlmnId `json:"plmnId"`
	// 28-bit string identifying an E-UTRA Cell Id as specified in clause 9.3.1.9 of  3GPP TS 38.413, in hexadecimal representation. Each character in the string shall take a  value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most  significant character representing the 4 most significant bits of the Cell Id shall appear  first in the string, and the character representing the 4 least significant bit of the  Cell Id shall appear last in the string.
	EutraCellId string `json:"eutraCellId"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid *string `json:"nid,omitempty"`
}

// NewEcgi instantiates a new Ecgi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcgi(plmnId PlmnId, eutraCellId string) *Ecgi {
	this := Ecgi{}
	this.PlmnId = plmnId
	this.EutraCellId = eutraCellId
	return &this
}

// NewEcgiWithDefaults instantiates a new Ecgi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcgiWithDefaults() *Ecgi {
	this := Ecgi{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *Ecgi) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *Ecgi) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *Ecgi) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetEutraCellId returns the EutraCellId field value
func (o *Ecgi) GetEutraCellId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EutraCellId
}

// GetEutraCellIdOk returns a tuple with the EutraCellId field value
// and a boolean to check if the value has been set.
func (o *Ecgi) GetEutraCellIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EutraCellId, true
}

// SetEutraCellId sets field value
func (o *Ecgi) SetEutraCellId(v string) {
	o.EutraCellId = v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *Ecgi) GetNid() string {
	if o == nil || IsNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ecgi) GetNidOk() (*string, bool) {
	if o == nil || IsNil(o.Nid) {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *Ecgi) HasNid() bool {
	if o != nil && !IsNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *Ecgi) SetNid(v string) {
	o.Nid = &v
}

func (o Ecgi) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ecgi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	toSerialize["eutraCellId"] = o.EutraCellId
	if !IsNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	return toSerialize, nil
}

type NullableEcgi struct {
	value *Ecgi
	isSet bool
}

func (v NullableEcgi) Get() *Ecgi {
	return v.value
}

func (v *NullableEcgi) Set(val *Ecgi) {
	v.value = val
	v.isSet = true
}

func (v NullableEcgi) IsSet() bool {
	return v.isSet
}

func (v *NullableEcgi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcgi(val *Ecgi) *NullableEcgi {
	return &NullableEcgi{value: val, isSet: true}
}

func (v NullableEcgi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcgi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

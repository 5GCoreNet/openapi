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

// checks if the PlmnIdNid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlmnIdNid{}

// PlmnIdNid struct for PlmnIdNid
type PlmnIdNid struct {
	Mcc *string `json:"mcc,omitempty"`
	Mnc *string `json:"mnc,omitempty"`
	Nid *string `json:"nid,omitempty"`
}

// NewPlmnIdNid instantiates a new PlmnIdNid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnIdNid() *PlmnIdNid {
	this := PlmnIdNid{}
	return &this
}

// NewPlmnIdNidWithDefaults instantiates a new PlmnIdNid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnIdNidWithDefaults() *PlmnIdNid {
	this := PlmnIdNid{}
	return &this
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *PlmnIdNid) GetMcc() string {
	if o == nil || IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnIdNid) GetMccOk() (*string, bool) {
	if o == nil || IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *PlmnIdNid) HasMcc() bool {
	if o != nil && !IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *PlmnIdNid) SetMcc(v string) {
	o.Mcc = &v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *PlmnIdNid) GetMnc() string {
	if o == nil || IsNil(o.Mnc) {
		var ret string
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnIdNid) GetMncOk() (*string, bool) {
	if o == nil || IsNil(o.Mnc) {
		return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *PlmnIdNid) HasMnc() bool {
	if o != nil && !IsNil(o.Mnc) {
		return true
	}

	return false
}

// SetMnc gets a reference to the given string and assigns it to the Mnc field.
func (o *PlmnIdNid) SetMnc(v string) {
	o.Mnc = &v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *PlmnIdNid) GetNid() string {
	if o == nil || IsNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnIdNid) GetNidOk() (*string, bool) {
	if o == nil || IsNil(o.Nid) {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *PlmnIdNid) HasNid() bool {
	if o != nil && !IsNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *PlmnIdNid) SetNid(v string) {
	o.Nid = &v
}

func (o PlmnIdNid) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlmnIdNid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !IsNil(o.Mnc) {
		toSerialize["mnc"] = o.Mnc
	}
	if !IsNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	return toSerialize, nil
}

type NullablePlmnIdNid struct {
	value *PlmnIdNid
	isSet bool
}

func (v NullablePlmnIdNid) Get() *PlmnIdNid {
	return v.value
}

func (v *NullablePlmnIdNid) Set(val *PlmnIdNid) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnIdNid) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnIdNid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnIdNid(val *PlmnIdNid) *NullablePlmnIdNid {
	return &NullablePlmnIdNid{value: val, isSet: true}
}

func (v NullablePlmnIdNid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnIdNid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

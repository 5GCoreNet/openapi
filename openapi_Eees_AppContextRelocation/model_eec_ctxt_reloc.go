/*
Eees Application Context Relocation Service

Eees Application Context Relocation Service.   © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_AppContextRelocation

import (
	"encoding/json"
)

// checks if the EecCtxtReloc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EecCtxtReloc{}

// EecCtxtReloc Represents EEC Context relocation information.
type EecCtxtReloc struct {
	EecCtxtId string `json:"eecCtxtId"`
	SEesId *string `json:"sEesId,omitempty"`
	SEecEndpoint *EndPoint `json:"sEecEndpoint,omitempty"`
	TEesId *string `json:"tEesId,omitempty"`
	TEecEndpoint *EndPoint `json:"tEecEndpoint,omitempty"`
}

// NewEecCtxtReloc instantiates a new EecCtxtReloc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEecCtxtReloc(eecCtxtId string) *EecCtxtReloc {
	this := EecCtxtReloc{}
	this.EecCtxtId = eecCtxtId
	return &this
}

// NewEecCtxtRelocWithDefaults instantiates a new EecCtxtReloc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEecCtxtRelocWithDefaults() *EecCtxtReloc {
	this := EecCtxtReloc{}
	return &this
}

// GetEecCtxtId returns the EecCtxtId field value
func (o *EecCtxtReloc) GetEecCtxtId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EecCtxtId
}

// GetEecCtxtIdOk returns a tuple with the EecCtxtId field value
// and a boolean to check if the value has been set.
func (o *EecCtxtReloc) GetEecCtxtIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EecCtxtId, true
}

// SetEecCtxtId sets field value
func (o *EecCtxtReloc) SetEecCtxtId(v string) {
	o.EecCtxtId = v
}

// GetSEesId returns the SEesId field value if set, zero value otherwise.
func (o *EecCtxtReloc) GetSEesId() string {
	if o == nil || IsNil(o.SEesId) {
		var ret string
		return ret
	}
	return *o.SEesId
}

// GetSEesIdOk returns a tuple with the SEesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EecCtxtReloc) GetSEesIdOk() (*string, bool) {
	if o == nil || IsNil(o.SEesId) {
		return nil, false
	}
	return o.SEesId, true
}

// HasSEesId returns a boolean if a field has been set.
func (o *EecCtxtReloc) HasSEesId() bool {
	if o != nil && !IsNil(o.SEesId) {
		return true
	}

	return false
}

// SetSEesId gets a reference to the given string and assigns it to the SEesId field.
func (o *EecCtxtReloc) SetSEesId(v string) {
	o.SEesId = &v
}

// GetSEecEndpoint returns the SEecEndpoint field value if set, zero value otherwise.
func (o *EecCtxtReloc) GetSEecEndpoint() EndPoint {
	if o == nil || IsNil(o.SEecEndpoint) {
		var ret EndPoint
		return ret
	}
	return *o.SEecEndpoint
}

// GetSEecEndpointOk returns a tuple with the SEecEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EecCtxtReloc) GetSEecEndpointOk() (*EndPoint, bool) {
	if o == nil || IsNil(o.SEecEndpoint) {
		return nil, false
	}
	return o.SEecEndpoint, true
}

// HasSEecEndpoint returns a boolean if a field has been set.
func (o *EecCtxtReloc) HasSEecEndpoint() bool {
	if o != nil && !IsNil(o.SEecEndpoint) {
		return true
	}

	return false
}

// SetSEecEndpoint gets a reference to the given EndPoint and assigns it to the SEecEndpoint field.
func (o *EecCtxtReloc) SetSEecEndpoint(v EndPoint) {
	o.SEecEndpoint = &v
}

// GetTEesId returns the TEesId field value if set, zero value otherwise.
func (o *EecCtxtReloc) GetTEesId() string {
	if o == nil || IsNil(o.TEesId) {
		var ret string
		return ret
	}
	return *o.TEesId
}

// GetTEesIdOk returns a tuple with the TEesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EecCtxtReloc) GetTEesIdOk() (*string, bool) {
	if o == nil || IsNil(o.TEesId) {
		return nil, false
	}
	return o.TEesId, true
}

// HasTEesId returns a boolean if a field has been set.
func (o *EecCtxtReloc) HasTEesId() bool {
	if o != nil && !IsNil(o.TEesId) {
		return true
	}

	return false
}

// SetTEesId gets a reference to the given string and assigns it to the TEesId field.
func (o *EecCtxtReloc) SetTEesId(v string) {
	o.TEesId = &v
}

// GetTEecEndpoint returns the TEecEndpoint field value if set, zero value otherwise.
func (o *EecCtxtReloc) GetTEecEndpoint() EndPoint {
	if o == nil || IsNil(o.TEecEndpoint) {
		var ret EndPoint
		return ret
	}
	return *o.TEecEndpoint
}

// GetTEecEndpointOk returns a tuple with the TEecEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EecCtxtReloc) GetTEecEndpointOk() (*EndPoint, bool) {
	if o == nil || IsNil(o.TEecEndpoint) {
		return nil, false
	}
	return o.TEecEndpoint, true
}

// HasTEecEndpoint returns a boolean if a field has been set.
func (o *EecCtxtReloc) HasTEecEndpoint() bool {
	if o != nil && !IsNil(o.TEecEndpoint) {
		return true
	}

	return false
}

// SetTEecEndpoint gets a reference to the given EndPoint and assigns it to the TEecEndpoint field.
func (o *EecCtxtReloc) SetTEecEndpoint(v EndPoint) {
	o.TEecEndpoint = &v
}

func (o EecCtxtReloc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EecCtxtReloc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eecCtxtId"] = o.EecCtxtId
	if !IsNil(o.SEesId) {
		toSerialize["sEesId"] = o.SEesId
	}
	if !IsNil(o.SEecEndpoint) {
		toSerialize["sEecEndpoint"] = o.SEecEndpoint
	}
	if !IsNil(o.TEesId) {
		toSerialize["tEesId"] = o.TEesId
	}
	if !IsNil(o.TEecEndpoint) {
		toSerialize["tEecEndpoint"] = o.TEecEndpoint
	}
	return toSerialize, nil
}

type NullableEecCtxtReloc struct {
	value *EecCtxtReloc
	isSet bool
}

func (v NullableEecCtxtReloc) Get() *EecCtxtReloc {
	return v.value
}

func (v *NullableEecCtxtReloc) Set(val *EecCtxtReloc) {
	v.value = val
	v.isSet = true
}

func (v NullableEecCtxtReloc) IsSet() bool {
	return v.isSet
}

func (v *NullableEecCtxtReloc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEecCtxtReloc(val *EecCtxtReloc) *NullableEecCtxtReloc {
	return &NullableEecCtxtReloc{value: val, isSet: true}
}

func (v NullableEecCtxtReloc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEecCtxtReloc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the V2xContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2xContext{}

// V2xContext Represents the V2X services related parameters
type V2xContext struct {
	NrV2xServicesAuth  *NrV2xAuth  `json:"nrV2xServicesAuth,omitempty"`
	LteV2xServicesAuth *LteV2xAuth `json:"lteV2xServicesAuth,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	NrUeSidelinkAmbr *string `json:"nrUeSidelinkAmbr,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	LteUeSidelinkAmbr *string     `json:"lteUeSidelinkAmbr,omitempty"`
	Pc5QoSPara        *Pc5QoSPara `json:"pc5QoSPara,omitempty"`
}

// NewV2xContext instantiates a new V2xContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2xContext() *V2xContext {
	this := V2xContext{}
	return &this
}

// NewV2xContextWithDefaults instantiates a new V2xContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2xContextWithDefaults() *V2xContext {
	this := V2xContext{}
	return &this
}

// GetNrV2xServicesAuth returns the NrV2xServicesAuth field value if set, zero value otherwise.
func (o *V2xContext) GetNrV2xServicesAuth() NrV2xAuth {
	if o == nil || IsNil(o.NrV2xServicesAuth) {
		var ret NrV2xAuth
		return ret
	}
	return *o.NrV2xServicesAuth
}

// GetNrV2xServicesAuthOk returns a tuple with the NrV2xServicesAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2xContext) GetNrV2xServicesAuthOk() (*NrV2xAuth, bool) {
	if o == nil || IsNil(o.NrV2xServicesAuth) {
		return nil, false
	}
	return o.NrV2xServicesAuth, true
}

// HasNrV2xServicesAuth returns a boolean if a field has been set.
func (o *V2xContext) HasNrV2xServicesAuth() bool {
	if o != nil && !IsNil(o.NrV2xServicesAuth) {
		return true
	}

	return false
}

// SetNrV2xServicesAuth gets a reference to the given NrV2xAuth and assigns it to the NrV2xServicesAuth field.
func (o *V2xContext) SetNrV2xServicesAuth(v NrV2xAuth) {
	o.NrV2xServicesAuth = &v
}

// GetLteV2xServicesAuth returns the LteV2xServicesAuth field value if set, zero value otherwise.
func (o *V2xContext) GetLteV2xServicesAuth() LteV2xAuth {
	if o == nil || IsNil(o.LteV2xServicesAuth) {
		var ret LteV2xAuth
		return ret
	}
	return *o.LteV2xServicesAuth
}

// GetLteV2xServicesAuthOk returns a tuple with the LteV2xServicesAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2xContext) GetLteV2xServicesAuthOk() (*LteV2xAuth, bool) {
	if o == nil || IsNil(o.LteV2xServicesAuth) {
		return nil, false
	}
	return o.LteV2xServicesAuth, true
}

// HasLteV2xServicesAuth returns a boolean if a field has been set.
func (o *V2xContext) HasLteV2xServicesAuth() bool {
	if o != nil && !IsNil(o.LteV2xServicesAuth) {
		return true
	}

	return false
}

// SetLteV2xServicesAuth gets a reference to the given LteV2xAuth and assigns it to the LteV2xServicesAuth field.
func (o *V2xContext) SetLteV2xServicesAuth(v LteV2xAuth) {
	o.LteV2xServicesAuth = &v
}

// GetNrUeSidelinkAmbr returns the NrUeSidelinkAmbr field value if set, zero value otherwise.
func (o *V2xContext) GetNrUeSidelinkAmbr() string {
	if o == nil || IsNil(o.NrUeSidelinkAmbr) {
		var ret string
		return ret
	}
	return *o.NrUeSidelinkAmbr
}

// GetNrUeSidelinkAmbrOk returns a tuple with the NrUeSidelinkAmbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2xContext) GetNrUeSidelinkAmbrOk() (*string, bool) {
	if o == nil || IsNil(o.NrUeSidelinkAmbr) {
		return nil, false
	}
	return o.NrUeSidelinkAmbr, true
}

// HasNrUeSidelinkAmbr returns a boolean if a field has been set.
func (o *V2xContext) HasNrUeSidelinkAmbr() bool {
	if o != nil && !IsNil(o.NrUeSidelinkAmbr) {
		return true
	}

	return false
}

// SetNrUeSidelinkAmbr gets a reference to the given string and assigns it to the NrUeSidelinkAmbr field.
func (o *V2xContext) SetNrUeSidelinkAmbr(v string) {
	o.NrUeSidelinkAmbr = &v
}

// GetLteUeSidelinkAmbr returns the LteUeSidelinkAmbr field value if set, zero value otherwise.
func (o *V2xContext) GetLteUeSidelinkAmbr() string {
	if o == nil || IsNil(o.LteUeSidelinkAmbr) {
		var ret string
		return ret
	}
	return *o.LteUeSidelinkAmbr
}

// GetLteUeSidelinkAmbrOk returns a tuple with the LteUeSidelinkAmbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2xContext) GetLteUeSidelinkAmbrOk() (*string, bool) {
	if o == nil || IsNil(o.LteUeSidelinkAmbr) {
		return nil, false
	}
	return o.LteUeSidelinkAmbr, true
}

// HasLteUeSidelinkAmbr returns a boolean if a field has been set.
func (o *V2xContext) HasLteUeSidelinkAmbr() bool {
	if o != nil && !IsNil(o.LteUeSidelinkAmbr) {
		return true
	}

	return false
}

// SetLteUeSidelinkAmbr gets a reference to the given string and assigns it to the LteUeSidelinkAmbr field.
func (o *V2xContext) SetLteUeSidelinkAmbr(v string) {
	o.LteUeSidelinkAmbr = &v
}

// GetPc5QoSPara returns the Pc5QoSPara field value if set, zero value otherwise.
func (o *V2xContext) GetPc5QoSPara() Pc5QoSPara {
	if o == nil || IsNil(o.Pc5QoSPara) {
		var ret Pc5QoSPara
		return ret
	}
	return *o.Pc5QoSPara
}

// GetPc5QoSParaOk returns a tuple with the Pc5QoSPara field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2xContext) GetPc5QoSParaOk() (*Pc5QoSPara, bool) {
	if o == nil || IsNil(o.Pc5QoSPara) {
		return nil, false
	}
	return o.Pc5QoSPara, true
}

// HasPc5QoSPara returns a boolean if a field has been set.
func (o *V2xContext) HasPc5QoSPara() bool {
	if o != nil && !IsNil(o.Pc5QoSPara) {
		return true
	}

	return false
}

// SetPc5QoSPara gets a reference to the given Pc5QoSPara and assigns it to the Pc5QoSPara field.
func (o *V2xContext) SetPc5QoSPara(v Pc5QoSPara) {
	o.Pc5QoSPara = &v
}

func (o V2xContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2xContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NrV2xServicesAuth) {
		toSerialize["nrV2xServicesAuth"] = o.NrV2xServicesAuth
	}
	if !IsNil(o.LteV2xServicesAuth) {
		toSerialize["lteV2xServicesAuth"] = o.LteV2xServicesAuth
	}
	if !IsNil(o.NrUeSidelinkAmbr) {
		toSerialize["nrUeSidelinkAmbr"] = o.NrUeSidelinkAmbr
	}
	if !IsNil(o.LteUeSidelinkAmbr) {
		toSerialize["lteUeSidelinkAmbr"] = o.LteUeSidelinkAmbr
	}
	if !IsNil(o.Pc5QoSPara) {
		toSerialize["pc5QoSPara"] = o.Pc5QoSPara
	}
	return toSerialize, nil
}

type NullableV2xContext struct {
	value *V2xContext
	isSet bool
}

func (v NullableV2xContext) Get() *V2xContext {
	return v.value
}

func (v *NullableV2xContext) Set(val *V2xContext) {
	v.value = val
	v.isSet = true
}

func (v NullableV2xContext) IsSet() bool {
	return v.isSet
}

func (v *NullableV2xContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2xContext(val *V2xContext) *NullableV2xContext {
	return &NullableV2xContext{value: val, isSet: true}
}

func (v NullableV2xContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2xContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

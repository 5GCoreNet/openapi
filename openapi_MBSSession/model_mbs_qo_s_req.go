/*
3gpp-mbs-session

API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSSession

import (
	"encoding/json"
)

// checks if the MbsQoSReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsQoSReq{}

// MbsQoSReq Represent MBS QoS requirements.
type MbsQoSReq struct {
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255. 
	Var5qi int32 `json:"5qi"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	GuarBitRate *string `json:"guarBitRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxBitRate *string `json:"maxBitRate,omitempty"`
	// Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AverWindow *int32 `json:"averWindow,omitempty"`
	ReqMbsArp *Arp `json:"reqMbsArp,omitempty"`
}

// NewMbsQoSReq instantiates a new MbsQoSReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsQoSReq(var5qi int32) *MbsQoSReq {
	this := MbsQoSReq{}
	this.Var5qi = var5qi
	var averWindow int32 = 2000
	this.AverWindow = &averWindow
	return &this
}

// NewMbsQoSReqWithDefaults instantiates a new MbsQoSReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsQoSReqWithDefaults() *MbsQoSReq {
	this := MbsQoSReq{}
	var averWindow int32 = 2000
	this.AverWindow = &averWindow
	return &this
}

// GetVar5qi returns the Var5qi field value
func (o *MbsQoSReq) GetVar5qi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var5qi
}

// GetVar5qiOk returns a tuple with the Var5qi field value
// and a boolean to check if the value has been set.
func (o *MbsQoSReq) GetVar5qiOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var5qi, true
}

// SetVar5qi sets field value
func (o *MbsQoSReq) SetVar5qi(v int32) {
	o.Var5qi = v
}

// GetGuarBitRate returns the GuarBitRate field value if set, zero value otherwise.
func (o *MbsQoSReq) GetGuarBitRate() string {
	if o == nil || isNil(o.GuarBitRate) {
		var ret string
		return ret
	}
	return *o.GuarBitRate
}

// GetGuarBitRateOk returns a tuple with the GuarBitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsQoSReq) GetGuarBitRateOk() (*string, bool) {
	if o == nil || isNil(o.GuarBitRate) {
		return nil, false
	}
	return o.GuarBitRate, true
}

// HasGuarBitRate returns a boolean if a field has been set.
func (o *MbsQoSReq) HasGuarBitRate() bool {
	if o != nil && !isNil(o.GuarBitRate) {
		return true
	}

	return false
}

// SetGuarBitRate gets a reference to the given string and assigns it to the GuarBitRate field.
func (o *MbsQoSReq) SetGuarBitRate(v string) {
	o.GuarBitRate = &v
}

// GetMaxBitRate returns the MaxBitRate field value if set, zero value otherwise.
func (o *MbsQoSReq) GetMaxBitRate() string {
	if o == nil || isNil(o.MaxBitRate) {
		var ret string
		return ret
	}
	return *o.MaxBitRate
}

// GetMaxBitRateOk returns a tuple with the MaxBitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsQoSReq) GetMaxBitRateOk() (*string, bool) {
	if o == nil || isNil(o.MaxBitRate) {
		return nil, false
	}
	return o.MaxBitRate, true
}

// HasMaxBitRate returns a boolean if a field has been set.
func (o *MbsQoSReq) HasMaxBitRate() bool {
	if o != nil && !isNil(o.MaxBitRate) {
		return true
	}

	return false
}

// SetMaxBitRate gets a reference to the given string and assigns it to the MaxBitRate field.
func (o *MbsQoSReq) SetMaxBitRate(v string) {
	o.MaxBitRate = &v
}

// GetAverWindow returns the AverWindow field value if set, zero value otherwise.
func (o *MbsQoSReq) GetAverWindow() int32 {
	if o == nil || isNil(o.AverWindow) {
		var ret int32
		return ret
	}
	return *o.AverWindow
}

// GetAverWindowOk returns a tuple with the AverWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsQoSReq) GetAverWindowOk() (*int32, bool) {
	if o == nil || isNil(o.AverWindow) {
		return nil, false
	}
	return o.AverWindow, true
}

// HasAverWindow returns a boolean if a field has been set.
func (o *MbsQoSReq) HasAverWindow() bool {
	if o != nil && !isNil(o.AverWindow) {
		return true
	}

	return false
}

// SetAverWindow gets a reference to the given int32 and assigns it to the AverWindow field.
func (o *MbsQoSReq) SetAverWindow(v int32) {
	o.AverWindow = &v
}

// GetReqMbsArp returns the ReqMbsArp field value if set, zero value otherwise.
func (o *MbsQoSReq) GetReqMbsArp() Arp {
	if o == nil || isNil(o.ReqMbsArp) {
		var ret Arp
		return ret
	}
	return *o.ReqMbsArp
}

// GetReqMbsArpOk returns a tuple with the ReqMbsArp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsQoSReq) GetReqMbsArpOk() (*Arp, bool) {
	if o == nil || isNil(o.ReqMbsArp) {
		return nil, false
	}
	return o.ReqMbsArp, true
}

// HasReqMbsArp returns a boolean if a field has been set.
func (o *MbsQoSReq) HasReqMbsArp() bool {
	if o != nil && !isNil(o.ReqMbsArp) {
		return true
	}

	return false
}

// SetReqMbsArp gets a reference to the given Arp and assigns it to the ReqMbsArp field.
func (o *MbsQoSReq) SetReqMbsArp(v Arp) {
	o.ReqMbsArp = &v
}

func (o MbsQoSReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsQoSReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["5qi"] = o.Var5qi
	if !isNil(o.GuarBitRate) {
		toSerialize["guarBitRate"] = o.GuarBitRate
	}
	if !isNil(o.MaxBitRate) {
		toSerialize["maxBitRate"] = o.MaxBitRate
	}
	if !isNil(o.AverWindow) {
		toSerialize["averWindow"] = o.AverWindow
	}
	if !isNil(o.ReqMbsArp) {
		toSerialize["reqMbsArp"] = o.ReqMbsArp
	}
	return toSerialize, nil
}

type NullableMbsQoSReq struct {
	value *MbsQoSReq
	isSet bool
}

func (v NullableMbsQoSReq) Get() *MbsQoSReq {
	return v.value
}

func (v *NullableMbsQoSReq) Set(val *MbsQoSReq) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsQoSReq) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsQoSReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsQoSReq(val *MbsQoSReq) *NullableMbsQoSReq {
	return &NullableMbsQoSReq{value: val, isSet: true}
}

func (v NullableMbsQoSReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsQoSReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



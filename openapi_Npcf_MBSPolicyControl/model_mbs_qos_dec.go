/*
Npcf_MBSPolicyControl API

MBS Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_MBSPolicyControl

import (
	"encoding/json"
)

// checks if the MbsQosDec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsQosDec{}

// MbsQosDec Represents the parameters constituting an MBS QoS Decision.
type MbsQosDec struct {
	MbsQosId string `json:"mbsQosId"`
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255.
	Var5qi *int32 `json:"5qi,omitempty"`
	// Unsigned integer indicating the 5QI Priority Level (see clauses 5.7.3.3 and 5.7.4 of 3GPP TS 23.501, within the range 1 to 127.Values are ordered in decreasing order of priority,  i.e. with 1 as the highest priority and 127 as the lowest priority.
	PriorityLevel *int32 `json:"priorityLevel,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MbrDl *string `json:"mbrDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GbrDl *string `json:"gbrDl,omitempty"`
	Arp   *Arp    `json:"arp,omitempty"`
	// Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AverWindow *int32 `json:"averWindow,omitempty"`
	// Represents the MBS Maximum Data Burst Volume expressed in Bytes.
	MbsMaxDataBurstVol *int32 `json:"mbsMaxDataBurstVol,omitempty"`
}

// NewMbsQosDec instantiates a new MbsQosDec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsQosDec(mbsQosId string) *MbsQosDec {
	this := MbsQosDec{}
	this.MbsQosId = mbsQosId
	var averWindow int32 = 2000
	this.AverWindow = &averWindow
	return &this
}

// NewMbsQosDecWithDefaults instantiates a new MbsQosDec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsQosDecWithDefaults() *MbsQosDec {
	this := MbsQosDec{}
	var averWindow int32 = 2000
	this.AverWindow = &averWindow
	return &this
}

// GetMbsQosId returns the MbsQosId field value
func (o *MbsQosDec) GetMbsQosId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MbsQosId
}

// GetMbsQosIdOk returns a tuple with the MbsQosId field value
// and a boolean to check if the value has been set.
func (o *MbsQosDec) GetMbsQosIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MbsQosId, true
}

// SetMbsQosId sets field value
func (o *MbsQosDec) SetMbsQosId(v string) {
	o.MbsQosId = v
}

// GetVar5qi returns the Var5qi field value if set, zero value otherwise.
func (o *MbsQosDec) GetVar5qi() int32 {
	if o == nil || IsNil(o.Var5qi) {
		var ret int32
		return ret
	}
	return *o.Var5qi
}

// GetVar5qiOk returns a tuple with the Var5qi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsQosDec) GetVar5qiOk() (*int32, bool) {
	if o == nil || IsNil(o.Var5qi) {
		return nil, false
	}
	return o.Var5qi, true
}

// HasVar5qi returns a boolean if a field has been set.
func (o *MbsQosDec) HasVar5qi() bool {
	if o != nil && !IsNil(o.Var5qi) {
		return true
	}

	return false
}

// SetVar5qi gets a reference to the given int32 and assigns it to the Var5qi field.
func (o *MbsQosDec) SetVar5qi(v int32) {
	o.Var5qi = &v
}

// GetPriorityLevel returns the PriorityLevel field value if set, zero value otherwise.
func (o *MbsQosDec) GetPriorityLevel() int32 {
	if o == nil || IsNil(o.PriorityLevel) {
		var ret int32
		return ret
	}
	return *o.PriorityLevel
}

// GetPriorityLevelOk returns a tuple with the PriorityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsQosDec) GetPriorityLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLevel) {
		return nil, false
	}
	return o.PriorityLevel, true
}

// HasPriorityLevel returns a boolean if a field has been set.
func (o *MbsQosDec) HasPriorityLevel() bool {
	if o != nil && !IsNil(o.PriorityLevel) {
		return true
	}

	return false
}

// SetPriorityLevel gets a reference to the given int32 and assigns it to the PriorityLevel field.
func (o *MbsQosDec) SetPriorityLevel(v int32) {
	o.PriorityLevel = &v
}

// GetMbrDl returns the MbrDl field value if set, zero value otherwise.
func (o *MbsQosDec) GetMbrDl() string {
	if o == nil || IsNil(o.MbrDl) {
		var ret string
		return ret
	}
	return *o.MbrDl
}

// GetMbrDlOk returns a tuple with the MbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsQosDec) GetMbrDlOk() (*string, bool) {
	if o == nil || IsNil(o.MbrDl) {
		return nil, false
	}
	return o.MbrDl, true
}

// HasMbrDl returns a boolean if a field has been set.
func (o *MbsQosDec) HasMbrDl() bool {
	if o != nil && !IsNil(o.MbrDl) {
		return true
	}

	return false
}

// SetMbrDl gets a reference to the given string and assigns it to the MbrDl field.
func (o *MbsQosDec) SetMbrDl(v string) {
	o.MbrDl = &v
}

// GetGbrDl returns the GbrDl field value if set, zero value otherwise.
func (o *MbsQosDec) GetGbrDl() string {
	if o == nil || IsNil(o.GbrDl) {
		var ret string
		return ret
	}
	return *o.GbrDl
}

// GetGbrDlOk returns a tuple with the GbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsQosDec) GetGbrDlOk() (*string, bool) {
	if o == nil || IsNil(o.GbrDl) {
		return nil, false
	}
	return o.GbrDl, true
}

// HasGbrDl returns a boolean if a field has been set.
func (o *MbsQosDec) HasGbrDl() bool {
	if o != nil && !IsNil(o.GbrDl) {
		return true
	}

	return false
}

// SetGbrDl gets a reference to the given string and assigns it to the GbrDl field.
func (o *MbsQosDec) SetGbrDl(v string) {
	o.GbrDl = &v
}

// GetArp returns the Arp field value if set, zero value otherwise.
func (o *MbsQosDec) GetArp() Arp {
	if o == nil || IsNil(o.Arp) {
		var ret Arp
		return ret
	}
	return *o.Arp
}

// GetArpOk returns a tuple with the Arp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsQosDec) GetArpOk() (*Arp, bool) {
	if o == nil || IsNil(o.Arp) {
		return nil, false
	}
	return o.Arp, true
}

// HasArp returns a boolean if a field has been set.
func (o *MbsQosDec) HasArp() bool {
	if o != nil && !IsNil(o.Arp) {
		return true
	}

	return false
}

// SetArp gets a reference to the given Arp and assigns it to the Arp field.
func (o *MbsQosDec) SetArp(v Arp) {
	o.Arp = &v
}

// GetAverWindow returns the AverWindow field value if set, zero value otherwise.
func (o *MbsQosDec) GetAverWindow() int32 {
	if o == nil || IsNil(o.AverWindow) {
		var ret int32
		return ret
	}
	return *o.AverWindow
}

// GetAverWindowOk returns a tuple with the AverWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsQosDec) GetAverWindowOk() (*int32, bool) {
	if o == nil || IsNil(o.AverWindow) {
		return nil, false
	}
	return o.AverWindow, true
}

// HasAverWindow returns a boolean if a field has been set.
func (o *MbsQosDec) HasAverWindow() bool {
	if o != nil && !IsNil(o.AverWindow) {
		return true
	}

	return false
}

// SetAverWindow gets a reference to the given int32 and assigns it to the AverWindow field.
func (o *MbsQosDec) SetAverWindow(v int32) {
	o.AverWindow = &v
}

// GetMbsMaxDataBurstVol returns the MbsMaxDataBurstVol field value if set, zero value otherwise.
func (o *MbsQosDec) GetMbsMaxDataBurstVol() int32 {
	if o == nil || IsNil(o.MbsMaxDataBurstVol) {
		var ret int32
		return ret
	}
	return *o.MbsMaxDataBurstVol
}

// GetMbsMaxDataBurstVolOk returns a tuple with the MbsMaxDataBurstVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsQosDec) GetMbsMaxDataBurstVolOk() (*int32, bool) {
	if o == nil || IsNil(o.MbsMaxDataBurstVol) {
		return nil, false
	}
	return o.MbsMaxDataBurstVol, true
}

// HasMbsMaxDataBurstVol returns a boolean if a field has been set.
func (o *MbsQosDec) HasMbsMaxDataBurstVol() bool {
	if o != nil && !IsNil(o.MbsMaxDataBurstVol) {
		return true
	}

	return false
}

// SetMbsMaxDataBurstVol gets a reference to the given int32 and assigns it to the MbsMaxDataBurstVol field.
func (o *MbsQosDec) SetMbsMaxDataBurstVol(v int32) {
	o.MbsMaxDataBurstVol = &v
}

func (o MbsQosDec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsQosDec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mbsQosId"] = o.MbsQosId
	if !IsNil(o.Var5qi) {
		toSerialize["5qi"] = o.Var5qi
	}
	if !IsNil(o.PriorityLevel) {
		toSerialize["priorityLevel"] = o.PriorityLevel
	}
	if !IsNil(o.MbrDl) {
		toSerialize["mbrDl"] = o.MbrDl
	}
	if !IsNil(o.GbrDl) {
		toSerialize["gbrDl"] = o.GbrDl
	}
	if !IsNil(o.Arp) {
		toSerialize["arp"] = o.Arp
	}
	if !IsNil(o.AverWindow) {
		toSerialize["averWindow"] = o.AverWindow
	}
	if !IsNil(o.MbsMaxDataBurstVol) {
		toSerialize["mbsMaxDataBurstVol"] = o.MbsMaxDataBurstVol
	}
	return toSerialize, nil
}

type NullableMbsQosDec struct {
	value *MbsQosDec
	isSet bool
}

func (v NullableMbsQosDec) Get() *MbsQosDec {
	return v.value
}

func (v *NullableMbsQosDec) Set(val *MbsQosDec) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsQosDec) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsQosDec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsQosDec(val *MbsQosDec) *NullableMbsQosDec {
	return &NullableMbsQosDec{value: val, isSet: true}
}

func (v NullableMbsQosDec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsQosDec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

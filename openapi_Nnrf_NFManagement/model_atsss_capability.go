/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
)

// checks if the AtsssCapability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AtsssCapability{}

// AtsssCapability Containes Capability to support procedures related to Access Traffic Steering, Switching, Splitting. 
type AtsssCapability struct {
	// Indicates the ATSSS-LL capability to support procedures related to Access Traffic Steering, Switching, Splitting (see clauses 4.2.10, 5.32 of 3GPP TS 23.501). true: Supported false (default): Not Supported 
	AtsssLL *bool `json:"atsssLL,omitempty"`
	// Indicates the MPTCP capability to support procedures related to Access Traffic Steering, Switching, Splitting (see clauses 4.2.10, 5.32 of 3GPP TS 23.501 true: Supported false (default): Not Supported 
	Mptcp *bool `json:"mptcp,omitempty"`
	// This IE is only used by the UPF to indicate whether the UPF supports RTT measurement without PMF (see clauses 5.32.2, 6.3.3.3 of 3GPP TS 23.501 true: Supported false (default): Not Supported 
	RttWithoutPmf *bool `json:"rttWithoutPmf,omitempty"`
}

// NewAtsssCapability instantiates a new AtsssCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtsssCapability() *AtsssCapability {
	this := AtsssCapability{}
	var atsssLL bool = false
	this.AtsssLL = &atsssLL
	var mptcp bool = false
	this.Mptcp = &mptcp
	var rttWithoutPmf bool = false
	this.RttWithoutPmf = &rttWithoutPmf
	return &this
}

// NewAtsssCapabilityWithDefaults instantiates a new AtsssCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtsssCapabilityWithDefaults() *AtsssCapability {
	this := AtsssCapability{}
	var atsssLL bool = false
	this.AtsssLL = &atsssLL
	var mptcp bool = false
	this.Mptcp = &mptcp
	var rttWithoutPmf bool = false
	this.RttWithoutPmf = &rttWithoutPmf
	return &this
}

// GetAtsssLL returns the AtsssLL field value if set, zero value otherwise.
func (o *AtsssCapability) GetAtsssLL() bool {
	if o == nil || isNil(o.AtsssLL) {
		var ret bool
		return ret
	}
	return *o.AtsssLL
}

// GetAtsssLLOk returns a tuple with the AtsssLL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsssCapability) GetAtsssLLOk() (*bool, bool) {
	if o == nil || isNil(o.AtsssLL) {
		return nil, false
	}
	return o.AtsssLL, true
}

// HasAtsssLL returns a boolean if a field has been set.
func (o *AtsssCapability) HasAtsssLL() bool {
	if o != nil && !isNil(o.AtsssLL) {
		return true
	}

	return false
}

// SetAtsssLL gets a reference to the given bool and assigns it to the AtsssLL field.
func (o *AtsssCapability) SetAtsssLL(v bool) {
	o.AtsssLL = &v
}

// GetMptcp returns the Mptcp field value if set, zero value otherwise.
func (o *AtsssCapability) GetMptcp() bool {
	if o == nil || isNil(o.Mptcp) {
		var ret bool
		return ret
	}
	return *o.Mptcp
}

// GetMptcpOk returns a tuple with the Mptcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsssCapability) GetMptcpOk() (*bool, bool) {
	if o == nil || isNil(o.Mptcp) {
		return nil, false
	}
	return o.Mptcp, true
}

// HasMptcp returns a boolean if a field has been set.
func (o *AtsssCapability) HasMptcp() bool {
	if o != nil && !isNil(o.Mptcp) {
		return true
	}

	return false
}

// SetMptcp gets a reference to the given bool and assigns it to the Mptcp field.
func (o *AtsssCapability) SetMptcp(v bool) {
	o.Mptcp = &v
}

// GetRttWithoutPmf returns the RttWithoutPmf field value if set, zero value otherwise.
func (o *AtsssCapability) GetRttWithoutPmf() bool {
	if o == nil || isNil(o.RttWithoutPmf) {
		var ret bool
		return ret
	}
	return *o.RttWithoutPmf
}

// GetRttWithoutPmfOk returns a tuple with the RttWithoutPmf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsssCapability) GetRttWithoutPmfOk() (*bool, bool) {
	if o == nil || isNil(o.RttWithoutPmf) {
		return nil, false
	}
	return o.RttWithoutPmf, true
}

// HasRttWithoutPmf returns a boolean if a field has been set.
func (o *AtsssCapability) HasRttWithoutPmf() bool {
	if o != nil && !isNil(o.RttWithoutPmf) {
		return true
	}

	return false
}

// SetRttWithoutPmf gets a reference to the given bool and assigns it to the RttWithoutPmf field.
func (o *AtsssCapability) SetRttWithoutPmf(v bool) {
	o.RttWithoutPmf = &v
}

func (o AtsssCapability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AtsssCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AtsssLL) {
		toSerialize["atsssLL"] = o.AtsssLL
	}
	if !isNil(o.Mptcp) {
		toSerialize["mptcp"] = o.Mptcp
	}
	if !isNil(o.RttWithoutPmf) {
		toSerialize["rttWithoutPmf"] = o.RttWithoutPmf
	}
	return toSerialize, nil
}

type NullableAtsssCapability struct {
	value *AtsssCapability
	isSet bool
}

func (v NullableAtsssCapability) Get() *AtsssCapability {
	return v.value
}

func (v *NullableAtsssCapability) Set(val *AtsssCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableAtsssCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableAtsssCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtsssCapability(val *AtsssCapability) *NullableAtsssCapability {
	return &NullableAtsssCapability{value: val, isSet: true}
}

func (v NullableAtsssCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtsssCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



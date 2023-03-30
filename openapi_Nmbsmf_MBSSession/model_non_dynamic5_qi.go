/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
)

// checks if the NonDynamic5Qi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonDynamic5Qi{}

// NonDynamic5Qi It indicates the QoS Characteristics for a standardized or pre-configured 5QI for downlink and uplink. 
type NonDynamic5Qi struct {
	// Unsigned integer indicating the 5QI Priority Level (see clauses 5.7.3.3 and 5.7.4 of 3GPP TS 23.501, within the range 1 to 127.Values are ordered in decreasing order of priority,  i.e. with 1 as the highest priority and 127 as the lowest priority.  
	PriorityLevel *int32 `json:"priorityLevel,omitempty"`
	// Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AverWindow *int32 `json:"averWindow,omitempty"`
	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.  
	MaxDataBurstVol *int32 `json:"maxDataBurstVol,omitempty"`
	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.  
	ExtMaxDataBurstVol *int32 `json:"extMaxDataBurstVol,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501 [8])), expressed in 0.01 milliseconds. 
	CnPacketDelayBudgetDl *int32 `json:"cnPacketDelayBudgetDl,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501 [8])), expressed in 0.01 milliseconds. 
	CnPacketDelayBudgetUl *int32 `json:"cnPacketDelayBudgetUl,omitempty"`
}

// NewNonDynamic5Qi instantiates a new NonDynamic5Qi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonDynamic5Qi() *NonDynamic5Qi {
	this := NonDynamic5Qi{}
	var averWindow int32 = 2000
	this.AverWindow = &averWindow
	return &this
}

// NewNonDynamic5QiWithDefaults instantiates a new NonDynamic5Qi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonDynamic5QiWithDefaults() *NonDynamic5Qi {
	this := NonDynamic5Qi{}
	var averWindow int32 = 2000
	this.AverWindow = &averWindow
	return &this
}

// GetPriorityLevel returns the PriorityLevel field value if set, zero value otherwise.
func (o *NonDynamic5Qi) GetPriorityLevel() int32 {
	if o == nil || IsNil(o.PriorityLevel) {
		var ret int32
		return ret
	}
	return *o.PriorityLevel
}

// GetPriorityLevelOk returns a tuple with the PriorityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonDynamic5Qi) GetPriorityLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLevel) {
		return nil, false
	}
	return o.PriorityLevel, true
}

// HasPriorityLevel returns a boolean if a field has been set.
func (o *NonDynamic5Qi) HasPriorityLevel() bool {
	if o != nil && !IsNil(o.PriorityLevel) {
		return true
	}

	return false
}

// SetPriorityLevel gets a reference to the given int32 and assigns it to the PriorityLevel field.
func (o *NonDynamic5Qi) SetPriorityLevel(v int32) {
	o.PriorityLevel = &v
}

// GetAverWindow returns the AverWindow field value if set, zero value otherwise.
func (o *NonDynamic5Qi) GetAverWindow() int32 {
	if o == nil || IsNil(o.AverWindow) {
		var ret int32
		return ret
	}
	return *o.AverWindow
}

// GetAverWindowOk returns a tuple with the AverWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonDynamic5Qi) GetAverWindowOk() (*int32, bool) {
	if o == nil || IsNil(o.AverWindow) {
		return nil, false
	}
	return o.AverWindow, true
}

// HasAverWindow returns a boolean if a field has been set.
func (o *NonDynamic5Qi) HasAverWindow() bool {
	if o != nil && !IsNil(o.AverWindow) {
		return true
	}

	return false
}

// SetAverWindow gets a reference to the given int32 and assigns it to the AverWindow field.
func (o *NonDynamic5Qi) SetAverWindow(v int32) {
	o.AverWindow = &v
}

// GetMaxDataBurstVol returns the MaxDataBurstVol field value if set, zero value otherwise.
func (o *NonDynamic5Qi) GetMaxDataBurstVol() int32 {
	if o == nil || IsNil(o.MaxDataBurstVol) {
		var ret int32
		return ret
	}
	return *o.MaxDataBurstVol
}

// GetMaxDataBurstVolOk returns a tuple with the MaxDataBurstVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonDynamic5Qi) GetMaxDataBurstVolOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDataBurstVol) {
		return nil, false
	}
	return o.MaxDataBurstVol, true
}

// HasMaxDataBurstVol returns a boolean if a field has been set.
func (o *NonDynamic5Qi) HasMaxDataBurstVol() bool {
	if o != nil && !IsNil(o.MaxDataBurstVol) {
		return true
	}

	return false
}

// SetMaxDataBurstVol gets a reference to the given int32 and assigns it to the MaxDataBurstVol field.
func (o *NonDynamic5Qi) SetMaxDataBurstVol(v int32) {
	o.MaxDataBurstVol = &v
}

// GetExtMaxDataBurstVol returns the ExtMaxDataBurstVol field value if set, zero value otherwise.
func (o *NonDynamic5Qi) GetExtMaxDataBurstVol() int32 {
	if o == nil || IsNil(o.ExtMaxDataBurstVol) {
		var ret int32
		return ret
	}
	return *o.ExtMaxDataBurstVol
}

// GetExtMaxDataBurstVolOk returns a tuple with the ExtMaxDataBurstVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonDynamic5Qi) GetExtMaxDataBurstVolOk() (*int32, bool) {
	if o == nil || IsNil(o.ExtMaxDataBurstVol) {
		return nil, false
	}
	return o.ExtMaxDataBurstVol, true
}

// HasExtMaxDataBurstVol returns a boolean if a field has been set.
func (o *NonDynamic5Qi) HasExtMaxDataBurstVol() bool {
	if o != nil && !IsNil(o.ExtMaxDataBurstVol) {
		return true
	}

	return false
}

// SetExtMaxDataBurstVol gets a reference to the given int32 and assigns it to the ExtMaxDataBurstVol field.
func (o *NonDynamic5Qi) SetExtMaxDataBurstVol(v int32) {
	o.ExtMaxDataBurstVol = &v
}

// GetCnPacketDelayBudgetDl returns the CnPacketDelayBudgetDl field value if set, zero value otherwise.
func (o *NonDynamic5Qi) GetCnPacketDelayBudgetDl() int32 {
	if o == nil || IsNil(o.CnPacketDelayBudgetDl) {
		var ret int32
		return ret
	}
	return *o.CnPacketDelayBudgetDl
}

// GetCnPacketDelayBudgetDlOk returns a tuple with the CnPacketDelayBudgetDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonDynamic5Qi) GetCnPacketDelayBudgetDlOk() (*int32, bool) {
	if o == nil || IsNil(o.CnPacketDelayBudgetDl) {
		return nil, false
	}
	return o.CnPacketDelayBudgetDl, true
}

// HasCnPacketDelayBudgetDl returns a boolean if a field has been set.
func (o *NonDynamic5Qi) HasCnPacketDelayBudgetDl() bool {
	if o != nil && !IsNil(o.CnPacketDelayBudgetDl) {
		return true
	}

	return false
}

// SetCnPacketDelayBudgetDl gets a reference to the given int32 and assigns it to the CnPacketDelayBudgetDl field.
func (o *NonDynamic5Qi) SetCnPacketDelayBudgetDl(v int32) {
	o.CnPacketDelayBudgetDl = &v
}

// GetCnPacketDelayBudgetUl returns the CnPacketDelayBudgetUl field value if set, zero value otherwise.
func (o *NonDynamic5Qi) GetCnPacketDelayBudgetUl() int32 {
	if o == nil || IsNil(o.CnPacketDelayBudgetUl) {
		var ret int32
		return ret
	}
	return *o.CnPacketDelayBudgetUl
}

// GetCnPacketDelayBudgetUlOk returns a tuple with the CnPacketDelayBudgetUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonDynamic5Qi) GetCnPacketDelayBudgetUlOk() (*int32, bool) {
	if o == nil || IsNil(o.CnPacketDelayBudgetUl) {
		return nil, false
	}
	return o.CnPacketDelayBudgetUl, true
}

// HasCnPacketDelayBudgetUl returns a boolean if a field has been set.
func (o *NonDynamic5Qi) HasCnPacketDelayBudgetUl() bool {
	if o != nil && !IsNil(o.CnPacketDelayBudgetUl) {
		return true
	}

	return false
}

// SetCnPacketDelayBudgetUl gets a reference to the given int32 and assigns it to the CnPacketDelayBudgetUl field.
func (o *NonDynamic5Qi) SetCnPacketDelayBudgetUl(v int32) {
	o.CnPacketDelayBudgetUl = &v
}

func (o NonDynamic5Qi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonDynamic5Qi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriorityLevel) {
		toSerialize["priorityLevel"] = o.PriorityLevel
	}
	if !IsNil(o.AverWindow) {
		toSerialize["averWindow"] = o.AverWindow
	}
	if !IsNil(o.MaxDataBurstVol) {
		toSerialize["maxDataBurstVol"] = o.MaxDataBurstVol
	}
	if !IsNil(o.ExtMaxDataBurstVol) {
		toSerialize["extMaxDataBurstVol"] = o.ExtMaxDataBurstVol
	}
	if !IsNil(o.CnPacketDelayBudgetDl) {
		toSerialize["cnPacketDelayBudgetDl"] = o.CnPacketDelayBudgetDl
	}
	if !IsNil(o.CnPacketDelayBudgetUl) {
		toSerialize["cnPacketDelayBudgetUl"] = o.CnPacketDelayBudgetUl
	}
	return toSerialize, nil
}

type NullableNonDynamic5Qi struct {
	value *NonDynamic5Qi
	isSet bool
}

func (v NullableNonDynamic5Qi) Get() *NonDynamic5Qi {
	return v.value
}

func (v *NullableNonDynamic5Qi) Set(val *NonDynamic5Qi) {
	v.value = val
	v.isSet = true
}

func (v NullableNonDynamic5Qi) IsSet() bool {
	return v.isSet
}

func (v *NullableNonDynamic5Qi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonDynamic5Qi(val *NonDynamic5Qi) *NullableNonDynamic5Qi {
	return &NullableNonDynamic5Qi{value: val, isSet: true}
}

func (v NullableNonDynamic5Qi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonDynamic5Qi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



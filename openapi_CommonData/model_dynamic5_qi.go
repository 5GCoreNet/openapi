/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
)

// checks if the Dynamic5Qi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dynamic5Qi{}

// Dynamic5Qi It indicates the QoS Characteristics for a Non-standardised or not pre-configured 5QI for downlink and uplink. 
type Dynamic5Qi struct {
	ResourceType QosResourceType `json:"resourceType"`
	// Unsigned integer indicating the 5QI Priority Level (see clauses 5.7.3.3 and 5.7.4 of 3GPP TS 23.501, within the range 1 to 127.Values are ordered in decreasing order of priority,  i.e. with 1 as the highest priority and 127 as the lowest priority.  
	PriorityLevel int32 `json:"priorityLevel"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	PacketDelayBudget int32 `json:"packetDelayBudget"`
	// String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \"scalar x 10-k\" where the scalar and the exponent k are each encoded as one decimal digit. 
	PacketErrRate string `json:"packetErrRate"`
	// Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AverWindow *int32 `json:"averWindow,omitempty"`
	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.  
	MaxDataBurstVol *int32 `json:"maxDataBurstVol,omitempty"`
	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.  
	ExtMaxDataBurstVol *int32 `json:"extMaxDataBurstVol,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501 [8])), expressed in 0.01 milliseconds. 
	ExtPacketDelBudget *int32 `json:"extPacketDelBudget,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501 [8])), expressed in 0.01 milliseconds. 
	CnPacketDelayBudgetDl *int32 `json:"cnPacketDelayBudgetDl,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501 [8])), expressed in 0.01 milliseconds. 
	CnPacketDelayBudgetUl *int32 `json:"cnPacketDelayBudgetUl,omitempty"`
}

// NewDynamic5Qi instantiates a new Dynamic5Qi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamic5Qi(resourceType QosResourceType, priorityLevel int32, packetDelayBudget int32, packetErrRate string) *Dynamic5Qi {
	this := Dynamic5Qi{}
	this.ResourceType = resourceType
	this.PriorityLevel = priorityLevel
	this.PacketDelayBudget = packetDelayBudget
	this.PacketErrRate = packetErrRate
	var averWindow int32 = 2000
	this.AverWindow = &averWindow
	return &this
}

// NewDynamic5QiWithDefaults instantiates a new Dynamic5Qi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamic5QiWithDefaults() *Dynamic5Qi {
	this := Dynamic5Qi{}
	var averWindow int32 = 2000
	this.AverWindow = &averWindow
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *Dynamic5Qi) GetResourceType() QosResourceType {
	if o == nil {
		var ret QosResourceType
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetResourceTypeOk() (*QosResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *Dynamic5Qi) SetResourceType(v QosResourceType) {
	o.ResourceType = v
}

// GetPriorityLevel returns the PriorityLevel field value
func (o *Dynamic5Qi) GetPriorityLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PriorityLevel
}

// GetPriorityLevelOk returns a tuple with the PriorityLevel field value
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetPriorityLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriorityLevel, true
}

// SetPriorityLevel sets field value
func (o *Dynamic5Qi) SetPriorityLevel(v int32) {
	o.PriorityLevel = v
}

// GetPacketDelayBudget returns the PacketDelayBudget field value
func (o *Dynamic5Qi) GetPacketDelayBudget() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PacketDelayBudget
}

// GetPacketDelayBudgetOk returns a tuple with the PacketDelayBudget field value
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetPacketDelayBudgetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PacketDelayBudget, true
}

// SetPacketDelayBudget sets field value
func (o *Dynamic5Qi) SetPacketDelayBudget(v int32) {
	o.PacketDelayBudget = v
}

// GetPacketErrRate returns the PacketErrRate field value
func (o *Dynamic5Qi) GetPacketErrRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PacketErrRate
}

// GetPacketErrRateOk returns a tuple with the PacketErrRate field value
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetPacketErrRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PacketErrRate, true
}

// SetPacketErrRate sets field value
func (o *Dynamic5Qi) SetPacketErrRate(v string) {
	o.PacketErrRate = v
}

// GetAverWindow returns the AverWindow field value if set, zero value otherwise.
func (o *Dynamic5Qi) GetAverWindow() int32 {
	if o == nil || isNil(o.AverWindow) {
		var ret int32
		return ret
	}
	return *o.AverWindow
}

// GetAverWindowOk returns a tuple with the AverWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetAverWindowOk() (*int32, bool) {
	if o == nil || isNil(o.AverWindow) {
		return nil, false
	}
	return o.AverWindow, true
}

// HasAverWindow returns a boolean if a field has been set.
func (o *Dynamic5Qi) HasAverWindow() bool {
	if o != nil && !isNil(o.AverWindow) {
		return true
	}

	return false
}

// SetAverWindow gets a reference to the given int32 and assigns it to the AverWindow field.
func (o *Dynamic5Qi) SetAverWindow(v int32) {
	o.AverWindow = &v
}

// GetMaxDataBurstVol returns the MaxDataBurstVol field value if set, zero value otherwise.
func (o *Dynamic5Qi) GetMaxDataBurstVol() int32 {
	if o == nil || isNil(o.MaxDataBurstVol) {
		var ret int32
		return ret
	}
	return *o.MaxDataBurstVol
}

// GetMaxDataBurstVolOk returns a tuple with the MaxDataBurstVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetMaxDataBurstVolOk() (*int32, bool) {
	if o == nil || isNil(o.MaxDataBurstVol) {
		return nil, false
	}
	return o.MaxDataBurstVol, true
}

// HasMaxDataBurstVol returns a boolean if a field has been set.
func (o *Dynamic5Qi) HasMaxDataBurstVol() bool {
	if o != nil && !isNil(o.MaxDataBurstVol) {
		return true
	}

	return false
}

// SetMaxDataBurstVol gets a reference to the given int32 and assigns it to the MaxDataBurstVol field.
func (o *Dynamic5Qi) SetMaxDataBurstVol(v int32) {
	o.MaxDataBurstVol = &v
}

// GetExtMaxDataBurstVol returns the ExtMaxDataBurstVol field value if set, zero value otherwise.
func (o *Dynamic5Qi) GetExtMaxDataBurstVol() int32 {
	if o == nil || isNil(o.ExtMaxDataBurstVol) {
		var ret int32
		return ret
	}
	return *o.ExtMaxDataBurstVol
}

// GetExtMaxDataBurstVolOk returns a tuple with the ExtMaxDataBurstVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetExtMaxDataBurstVolOk() (*int32, bool) {
	if o == nil || isNil(o.ExtMaxDataBurstVol) {
		return nil, false
	}
	return o.ExtMaxDataBurstVol, true
}

// HasExtMaxDataBurstVol returns a boolean if a field has been set.
func (o *Dynamic5Qi) HasExtMaxDataBurstVol() bool {
	if o != nil && !isNil(o.ExtMaxDataBurstVol) {
		return true
	}

	return false
}

// SetExtMaxDataBurstVol gets a reference to the given int32 and assigns it to the ExtMaxDataBurstVol field.
func (o *Dynamic5Qi) SetExtMaxDataBurstVol(v int32) {
	o.ExtMaxDataBurstVol = &v
}

// GetExtPacketDelBudget returns the ExtPacketDelBudget field value if set, zero value otherwise.
func (o *Dynamic5Qi) GetExtPacketDelBudget() int32 {
	if o == nil || isNil(o.ExtPacketDelBudget) {
		var ret int32
		return ret
	}
	return *o.ExtPacketDelBudget
}

// GetExtPacketDelBudgetOk returns a tuple with the ExtPacketDelBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetExtPacketDelBudgetOk() (*int32, bool) {
	if o == nil || isNil(o.ExtPacketDelBudget) {
		return nil, false
	}
	return o.ExtPacketDelBudget, true
}

// HasExtPacketDelBudget returns a boolean if a field has been set.
func (o *Dynamic5Qi) HasExtPacketDelBudget() bool {
	if o != nil && !isNil(o.ExtPacketDelBudget) {
		return true
	}

	return false
}

// SetExtPacketDelBudget gets a reference to the given int32 and assigns it to the ExtPacketDelBudget field.
func (o *Dynamic5Qi) SetExtPacketDelBudget(v int32) {
	o.ExtPacketDelBudget = &v
}

// GetCnPacketDelayBudgetDl returns the CnPacketDelayBudgetDl field value if set, zero value otherwise.
func (o *Dynamic5Qi) GetCnPacketDelayBudgetDl() int32 {
	if o == nil || isNil(o.CnPacketDelayBudgetDl) {
		var ret int32
		return ret
	}
	return *o.CnPacketDelayBudgetDl
}

// GetCnPacketDelayBudgetDlOk returns a tuple with the CnPacketDelayBudgetDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetCnPacketDelayBudgetDlOk() (*int32, bool) {
	if o == nil || isNil(o.CnPacketDelayBudgetDl) {
		return nil, false
	}
	return o.CnPacketDelayBudgetDl, true
}

// HasCnPacketDelayBudgetDl returns a boolean if a field has been set.
func (o *Dynamic5Qi) HasCnPacketDelayBudgetDl() bool {
	if o != nil && !isNil(o.CnPacketDelayBudgetDl) {
		return true
	}

	return false
}

// SetCnPacketDelayBudgetDl gets a reference to the given int32 and assigns it to the CnPacketDelayBudgetDl field.
func (o *Dynamic5Qi) SetCnPacketDelayBudgetDl(v int32) {
	o.CnPacketDelayBudgetDl = &v
}

// GetCnPacketDelayBudgetUl returns the CnPacketDelayBudgetUl field value if set, zero value otherwise.
func (o *Dynamic5Qi) GetCnPacketDelayBudgetUl() int32 {
	if o == nil || isNil(o.CnPacketDelayBudgetUl) {
		var ret int32
		return ret
	}
	return *o.CnPacketDelayBudgetUl
}

// GetCnPacketDelayBudgetUlOk returns a tuple with the CnPacketDelayBudgetUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dynamic5Qi) GetCnPacketDelayBudgetUlOk() (*int32, bool) {
	if o == nil || isNil(o.CnPacketDelayBudgetUl) {
		return nil, false
	}
	return o.CnPacketDelayBudgetUl, true
}

// HasCnPacketDelayBudgetUl returns a boolean if a field has been set.
func (o *Dynamic5Qi) HasCnPacketDelayBudgetUl() bool {
	if o != nil && !isNil(o.CnPacketDelayBudgetUl) {
		return true
	}

	return false
}

// SetCnPacketDelayBudgetUl gets a reference to the given int32 and assigns it to the CnPacketDelayBudgetUl field.
func (o *Dynamic5Qi) SetCnPacketDelayBudgetUl(v int32) {
	o.CnPacketDelayBudgetUl = &v
}

func (o Dynamic5Qi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dynamic5Qi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["priorityLevel"] = o.PriorityLevel
	toSerialize["packetDelayBudget"] = o.PacketDelayBudget
	toSerialize["packetErrRate"] = o.PacketErrRate
	if !isNil(o.AverWindow) {
		toSerialize["averWindow"] = o.AverWindow
	}
	if !isNil(o.MaxDataBurstVol) {
		toSerialize["maxDataBurstVol"] = o.MaxDataBurstVol
	}
	if !isNil(o.ExtMaxDataBurstVol) {
		toSerialize["extMaxDataBurstVol"] = o.ExtMaxDataBurstVol
	}
	if !isNil(o.ExtPacketDelBudget) {
		toSerialize["extPacketDelBudget"] = o.ExtPacketDelBudget
	}
	if !isNil(o.CnPacketDelayBudgetDl) {
		toSerialize["cnPacketDelayBudgetDl"] = o.CnPacketDelayBudgetDl
	}
	if !isNil(o.CnPacketDelayBudgetUl) {
		toSerialize["cnPacketDelayBudgetUl"] = o.CnPacketDelayBudgetUl
	}
	return toSerialize, nil
}

type NullableDynamic5Qi struct {
	value *Dynamic5Qi
	isSet bool
}

func (v NullableDynamic5Qi) Get() *Dynamic5Qi {
	return v.value
}

func (v *NullableDynamic5Qi) Set(val *Dynamic5Qi) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamic5Qi) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamic5Qi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamic5Qi(val *Dynamic5Qi) *NullableDynamic5Qi {
	return &NullableDynamic5Qi{value: val, isSet: true}
}

func (v NullableDynamic5Qi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamic5Qi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



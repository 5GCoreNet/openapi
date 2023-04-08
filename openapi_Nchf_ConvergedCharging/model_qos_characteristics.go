/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the QosCharacteristics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosCharacteristics{}

// QosCharacteristics Contains QoS characteristics for a non-standardized or a non-configured 5QI.
type QosCharacteristics struct {
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255. 
	Var5qi int32 `json:"5qi"`
	ResourceType QosResourceType `json:"resourceType"`
	// Unsigned integer indicating the 5QI Priority Level (see clauses 5.7.3.3 and 5.7.4 of 3GPP TS 23.501, within the range 1 to 127.Values are ordered in decreasing order of priority,  i.e. with 1 as the highest priority and 127 as the lowest priority.  
	PriorityLevel int32 `json:"priorityLevel"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	PacketDelayBudget int32 `json:"packetDelayBudget"`
	// String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \"scalar x 10-k\" where the scalar and the exponent k are each encoded as one decimal digit. 
	PacketErrorRate string `json:"packetErrorRate"`
	// Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AveragingWindow *int32 `json:"averagingWindow,omitempty"`
	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.  
	MaxDataBurstVol *int32 `json:"maxDataBurstVol,omitempty"`
	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.  
	ExtMaxDataBurstVol *int32 `json:"extMaxDataBurstVol,omitempty"`
}

// NewQosCharacteristics instantiates a new QosCharacteristics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosCharacteristics(var5qi int32, resourceType QosResourceType, priorityLevel int32, packetDelayBudget int32, packetErrorRate string) *QosCharacteristics {
	this := QosCharacteristics{}
	this.Var5qi = var5qi
	this.ResourceType = resourceType
	this.PriorityLevel = priorityLevel
	this.PacketDelayBudget = packetDelayBudget
	this.PacketErrorRate = packetErrorRate
	var averagingWindow int32 = 2000
	this.AveragingWindow = &averagingWindow
	return &this
}

// NewQosCharacteristicsWithDefaults instantiates a new QosCharacteristics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosCharacteristicsWithDefaults() *QosCharacteristics {
	this := QosCharacteristics{}
	var averagingWindow int32 = 2000
	this.AveragingWindow = &averagingWindow
	return &this
}

// GetVar5qi returns the Var5qi field value
func (o *QosCharacteristics) GetVar5qi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var5qi
}

// GetVar5qiOk returns a tuple with the Var5qi field value
// and a boolean to check if the value has been set.
func (o *QosCharacteristics) GetVar5qiOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var5qi, true
}

// SetVar5qi sets field value
func (o *QosCharacteristics) SetVar5qi(v int32) {
	o.Var5qi = v
}

// GetResourceType returns the ResourceType field value
func (o *QosCharacteristics) GetResourceType() QosResourceType {
	if o == nil {
		var ret QosResourceType
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *QosCharacteristics) GetResourceTypeOk() (*QosResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *QosCharacteristics) SetResourceType(v QosResourceType) {
	o.ResourceType = v
}

// GetPriorityLevel returns the PriorityLevel field value
func (o *QosCharacteristics) GetPriorityLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PriorityLevel
}

// GetPriorityLevelOk returns a tuple with the PriorityLevel field value
// and a boolean to check if the value has been set.
func (o *QosCharacteristics) GetPriorityLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriorityLevel, true
}

// SetPriorityLevel sets field value
func (o *QosCharacteristics) SetPriorityLevel(v int32) {
	o.PriorityLevel = v
}

// GetPacketDelayBudget returns the PacketDelayBudget field value
func (o *QosCharacteristics) GetPacketDelayBudget() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PacketDelayBudget
}

// GetPacketDelayBudgetOk returns a tuple with the PacketDelayBudget field value
// and a boolean to check if the value has been set.
func (o *QosCharacteristics) GetPacketDelayBudgetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PacketDelayBudget, true
}

// SetPacketDelayBudget sets field value
func (o *QosCharacteristics) SetPacketDelayBudget(v int32) {
	o.PacketDelayBudget = v
}

// GetPacketErrorRate returns the PacketErrorRate field value
func (o *QosCharacteristics) GetPacketErrorRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PacketErrorRate
}

// GetPacketErrorRateOk returns a tuple with the PacketErrorRate field value
// and a boolean to check if the value has been set.
func (o *QosCharacteristics) GetPacketErrorRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PacketErrorRate, true
}

// SetPacketErrorRate sets field value
func (o *QosCharacteristics) SetPacketErrorRate(v string) {
	o.PacketErrorRate = v
}

// GetAveragingWindow returns the AveragingWindow field value if set, zero value otherwise.
func (o *QosCharacteristics) GetAveragingWindow() int32 {
	if o == nil || isNil(o.AveragingWindow) {
		var ret int32
		return ret
	}
	return *o.AveragingWindow
}

// GetAveragingWindowOk returns a tuple with the AveragingWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosCharacteristics) GetAveragingWindowOk() (*int32, bool) {
	if o == nil || isNil(o.AveragingWindow) {
		return nil, false
	}
	return o.AveragingWindow, true
}

// HasAveragingWindow returns a boolean if a field has been set.
func (o *QosCharacteristics) HasAveragingWindow() bool {
	if o != nil && !isNil(o.AveragingWindow) {
		return true
	}

	return false
}

// SetAveragingWindow gets a reference to the given int32 and assigns it to the AveragingWindow field.
func (o *QosCharacteristics) SetAveragingWindow(v int32) {
	o.AveragingWindow = &v
}

// GetMaxDataBurstVol returns the MaxDataBurstVol field value if set, zero value otherwise.
func (o *QosCharacteristics) GetMaxDataBurstVol() int32 {
	if o == nil || isNil(o.MaxDataBurstVol) {
		var ret int32
		return ret
	}
	return *o.MaxDataBurstVol
}

// GetMaxDataBurstVolOk returns a tuple with the MaxDataBurstVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosCharacteristics) GetMaxDataBurstVolOk() (*int32, bool) {
	if o == nil || isNil(o.MaxDataBurstVol) {
		return nil, false
	}
	return o.MaxDataBurstVol, true
}

// HasMaxDataBurstVol returns a boolean if a field has been set.
func (o *QosCharacteristics) HasMaxDataBurstVol() bool {
	if o != nil && !isNil(o.MaxDataBurstVol) {
		return true
	}

	return false
}

// SetMaxDataBurstVol gets a reference to the given int32 and assigns it to the MaxDataBurstVol field.
func (o *QosCharacteristics) SetMaxDataBurstVol(v int32) {
	o.MaxDataBurstVol = &v
}

// GetExtMaxDataBurstVol returns the ExtMaxDataBurstVol field value if set, zero value otherwise.
func (o *QosCharacteristics) GetExtMaxDataBurstVol() int32 {
	if o == nil || isNil(o.ExtMaxDataBurstVol) {
		var ret int32
		return ret
	}
	return *o.ExtMaxDataBurstVol
}

// GetExtMaxDataBurstVolOk returns a tuple with the ExtMaxDataBurstVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosCharacteristics) GetExtMaxDataBurstVolOk() (*int32, bool) {
	if o == nil || isNil(o.ExtMaxDataBurstVol) {
		return nil, false
	}
	return o.ExtMaxDataBurstVol, true
}

// HasExtMaxDataBurstVol returns a boolean if a field has been set.
func (o *QosCharacteristics) HasExtMaxDataBurstVol() bool {
	if o != nil && !isNil(o.ExtMaxDataBurstVol) {
		return true
	}

	return false
}

// SetExtMaxDataBurstVol gets a reference to the given int32 and assigns it to the ExtMaxDataBurstVol field.
func (o *QosCharacteristics) SetExtMaxDataBurstVol(v int32) {
	o.ExtMaxDataBurstVol = &v
}

func (o QosCharacteristics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosCharacteristics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["5qi"] = o.Var5qi
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["priorityLevel"] = o.PriorityLevel
	toSerialize["packetDelayBudget"] = o.PacketDelayBudget
	toSerialize["packetErrorRate"] = o.PacketErrorRate
	if !isNil(o.AveragingWindow) {
		toSerialize["averagingWindow"] = o.AveragingWindow
	}
	if !isNil(o.MaxDataBurstVol) {
		toSerialize["maxDataBurstVol"] = o.MaxDataBurstVol
	}
	if !isNil(o.ExtMaxDataBurstVol) {
		toSerialize["extMaxDataBurstVol"] = o.ExtMaxDataBurstVol
	}
	return toSerialize, nil
}

type NullableQosCharacteristics struct {
	value *QosCharacteristics
	isSet bool
}

func (v NullableQosCharacteristics) Get() *QosCharacteristics {
	return v.value
}

func (v *NullableQosCharacteristics) Set(val *QosCharacteristics) {
	v.value = val
	v.isSet = true
}

func (v NullableQosCharacteristics) IsSet() bool {
	return v.isSet
}

func (v *NullableQosCharacteristics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosCharacteristics(val *QosCharacteristics) *NullableQosCharacteristics {
	return &NullableQosCharacteristics{value: val, isSet: true}
}

func (v NullableQosCharacteristics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosCharacteristics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



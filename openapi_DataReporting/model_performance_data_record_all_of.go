/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_DataReporting

import (
	"encoding/json"
)

// checks if the PerformanceDataRecordAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceDataRecordAllOf{}

// PerformanceDataRecordAllOf struct for PerformanceDataRecordAllOf
type PerformanceDataRecordAllOf struct {
	TimeInterval TimeWindow `json:"timeInterval"`
	Location *LocationArea5G `json:"location,omitempty"`
	RemoteEndpoint *AddrFqdn `json:"remoteEndpoint,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	PacketDelayBudget *int32 `json:"packetDelayBudget,omitempty"`
	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. 
	PacketLossRate *int32 `json:"packetLossRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	UplinkThroughput *string `json:"uplinkThroughput,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	DownlinkThrougput *string `json:"downlinkThrougput,omitempty"`
}

// NewPerformanceDataRecordAllOf instantiates a new PerformanceDataRecordAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceDataRecordAllOf(timeInterval TimeWindow) *PerformanceDataRecordAllOf {
	this := PerformanceDataRecordAllOf{}
	this.TimeInterval = timeInterval
	return &this
}

// NewPerformanceDataRecordAllOfWithDefaults instantiates a new PerformanceDataRecordAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceDataRecordAllOfWithDefaults() *PerformanceDataRecordAllOf {
	this := PerformanceDataRecordAllOf{}
	return &this
}

// GetTimeInterval returns the TimeInterval field value
func (o *PerformanceDataRecordAllOf) GetTimeInterval() TimeWindow {
	if o == nil {
		var ret TimeWindow
		return ret
	}

	return o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value
// and a boolean to check if the value has been set.
func (o *PerformanceDataRecordAllOf) GetTimeIntervalOk() (*TimeWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeInterval, true
}

// SetTimeInterval sets field value
func (o *PerformanceDataRecordAllOf) SetTimeInterval(v TimeWindow) {
	o.TimeInterval = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PerformanceDataRecordAllOf) GetLocation() LocationArea5G {
	if o == nil || isNil(o.Location) {
		var ret LocationArea5G
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataRecordAllOf) GetLocationOk() (*LocationArea5G, bool) {
	if o == nil || isNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PerformanceDataRecordAllOf) HasLocation() bool {
	if o != nil && !isNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given LocationArea5G and assigns it to the Location field.
func (o *PerformanceDataRecordAllOf) SetLocation(v LocationArea5G) {
	o.Location = &v
}

// GetRemoteEndpoint returns the RemoteEndpoint field value if set, zero value otherwise.
func (o *PerformanceDataRecordAllOf) GetRemoteEndpoint() AddrFqdn {
	if o == nil || isNil(o.RemoteEndpoint) {
		var ret AddrFqdn
		return ret
	}
	return *o.RemoteEndpoint
}

// GetRemoteEndpointOk returns a tuple with the RemoteEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataRecordAllOf) GetRemoteEndpointOk() (*AddrFqdn, bool) {
	if o == nil || isNil(o.RemoteEndpoint) {
		return nil, false
	}
	return o.RemoteEndpoint, true
}

// HasRemoteEndpoint returns a boolean if a field has been set.
func (o *PerformanceDataRecordAllOf) HasRemoteEndpoint() bool {
	if o != nil && !isNil(o.RemoteEndpoint) {
		return true
	}

	return false
}

// SetRemoteEndpoint gets a reference to the given AddrFqdn and assigns it to the RemoteEndpoint field.
func (o *PerformanceDataRecordAllOf) SetRemoteEndpoint(v AddrFqdn) {
	o.RemoteEndpoint = &v
}

// GetPacketDelayBudget returns the PacketDelayBudget field value if set, zero value otherwise.
func (o *PerformanceDataRecordAllOf) GetPacketDelayBudget() int32 {
	if o == nil || isNil(o.PacketDelayBudget) {
		var ret int32
		return ret
	}
	return *o.PacketDelayBudget
}

// GetPacketDelayBudgetOk returns a tuple with the PacketDelayBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataRecordAllOf) GetPacketDelayBudgetOk() (*int32, bool) {
	if o == nil || isNil(o.PacketDelayBudget) {
		return nil, false
	}
	return o.PacketDelayBudget, true
}

// HasPacketDelayBudget returns a boolean if a field has been set.
func (o *PerformanceDataRecordAllOf) HasPacketDelayBudget() bool {
	if o != nil && !isNil(o.PacketDelayBudget) {
		return true
	}

	return false
}

// SetPacketDelayBudget gets a reference to the given int32 and assigns it to the PacketDelayBudget field.
func (o *PerformanceDataRecordAllOf) SetPacketDelayBudget(v int32) {
	o.PacketDelayBudget = &v
}

// GetPacketLossRate returns the PacketLossRate field value if set, zero value otherwise.
func (o *PerformanceDataRecordAllOf) GetPacketLossRate() int32 {
	if o == nil || isNil(o.PacketLossRate) {
		var ret int32
		return ret
	}
	return *o.PacketLossRate
}

// GetPacketLossRateOk returns a tuple with the PacketLossRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataRecordAllOf) GetPacketLossRateOk() (*int32, bool) {
	if o == nil || isNil(o.PacketLossRate) {
		return nil, false
	}
	return o.PacketLossRate, true
}

// HasPacketLossRate returns a boolean if a field has been set.
func (o *PerformanceDataRecordAllOf) HasPacketLossRate() bool {
	if o != nil && !isNil(o.PacketLossRate) {
		return true
	}

	return false
}

// SetPacketLossRate gets a reference to the given int32 and assigns it to the PacketLossRate field.
func (o *PerformanceDataRecordAllOf) SetPacketLossRate(v int32) {
	o.PacketLossRate = &v
}

// GetUplinkThroughput returns the UplinkThroughput field value if set, zero value otherwise.
func (o *PerformanceDataRecordAllOf) GetUplinkThroughput() string {
	if o == nil || isNil(o.UplinkThroughput) {
		var ret string
		return ret
	}
	return *o.UplinkThroughput
}

// GetUplinkThroughputOk returns a tuple with the UplinkThroughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataRecordAllOf) GetUplinkThroughputOk() (*string, bool) {
	if o == nil || isNil(o.UplinkThroughput) {
		return nil, false
	}
	return o.UplinkThroughput, true
}

// HasUplinkThroughput returns a boolean if a field has been set.
func (o *PerformanceDataRecordAllOf) HasUplinkThroughput() bool {
	if o != nil && !isNil(o.UplinkThroughput) {
		return true
	}

	return false
}

// SetUplinkThroughput gets a reference to the given string and assigns it to the UplinkThroughput field.
func (o *PerformanceDataRecordAllOf) SetUplinkThroughput(v string) {
	o.UplinkThroughput = &v
}

// GetDownlinkThrougput returns the DownlinkThrougput field value if set, zero value otherwise.
func (o *PerformanceDataRecordAllOf) GetDownlinkThrougput() string {
	if o == nil || isNil(o.DownlinkThrougput) {
		var ret string
		return ret
	}
	return *o.DownlinkThrougput
}

// GetDownlinkThrougputOk returns a tuple with the DownlinkThrougput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataRecordAllOf) GetDownlinkThrougputOk() (*string, bool) {
	if o == nil || isNil(o.DownlinkThrougput) {
		return nil, false
	}
	return o.DownlinkThrougput, true
}

// HasDownlinkThrougput returns a boolean if a field has been set.
func (o *PerformanceDataRecordAllOf) HasDownlinkThrougput() bool {
	if o != nil && !isNil(o.DownlinkThrougput) {
		return true
	}

	return false
}

// SetDownlinkThrougput gets a reference to the given string and assigns it to the DownlinkThrougput field.
func (o *PerformanceDataRecordAllOf) SetDownlinkThrougput(v string) {
	o.DownlinkThrougput = &v
}

func (o PerformanceDataRecordAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerformanceDataRecordAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timeInterval"] = o.TimeInterval
	if !isNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !isNil(o.RemoteEndpoint) {
		toSerialize["remoteEndpoint"] = o.RemoteEndpoint
	}
	if !isNil(o.PacketDelayBudget) {
		toSerialize["packetDelayBudget"] = o.PacketDelayBudget
	}
	if !isNil(o.PacketLossRate) {
		toSerialize["packetLossRate"] = o.PacketLossRate
	}
	if !isNil(o.UplinkThroughput) {
		toSerialize["uplinkThroughput"] = o.UplinkThroughput
	}
	if !isNil(o.DownlinkThrougput) {
		toSerialize["downlinkThrougput"] = o.DownlinkThrougput
	}
	return toSerialize, nil
}

type NullablePerformanceDataRecordAllOf struct {
	value *PerformanceDataRecordAllOf
	isSet bool
}

func (v NullablePerformanceDataRecordAllOf) Get() *PerformanceDataRecordAllOf {
	return v.value
}

func (v *NullablePerformanceDataRecordAllOf) Set(val *PerformanceDataRecordAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceDataRecordAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceDataRecordAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceDataRecordAllOf(val *PerformanceDataRecordAllOf) *NullablePerformanceDataRecordAllOf {
	return &NullablePerformanceDataRecordAllOf{value: val, isSet: true}
}

func (v NullablePerformanceDataRecordAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceDataRecordAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



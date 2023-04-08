/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
)

// checks if the ThresholdLevel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThresholdLevel{}

// ThresholdLevel Represents a threshold level.
type ThresholdLevel struct {
	CongLevel *int32 `json:"congLevel,omitempty"`
	NfLoadLevel *int32 `json:"nfLoadLevel,omitempty"`
	NfCpuUsage *int32 `json:"nfCpuUsage,omitempty"`
	NfMemoryUsage *int32 `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage *int32 `json:"nfStorageUsage,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	AvgTrafficRate *string `json:"avgTrafficRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxTrafficRate *string `json:"maxTrafficRate,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	AvgPacketDelay *int32 `json:"avgPacketDelay,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	MaxPacketDelay *int32 `json:"maxPacketDelay,omitempty"`
	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. 
	AvgPacketLossRate *int32 `json:"avgPacketLossRate,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	SvcExpLevel *float32 `json:"svcExpLevel,omitempty"`
}

// NewThresholdLevel instantiates a new ThresholdLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdLevel() *ThresholdLevel {
	this := ThresholdLevel{}
	return &this
}

// NewThresholdLevelWithDefaults instantiates a new ThresholdLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdLevelWithDefaults() *ThresholdLevel {
	this := ThresholdLevel{}
	return &this
}

// GetCongLevel returns the CongLevel field value if set, zero value otherwise.
func (o *ThresholdLevel) GetCongLevel() int32 {
	if o == nil || isNil(o.CongLevel) {
		var ret int32
		return ret
	}
	return *o.CongLevel
}

// GetCongLevelOk returns a tuple with the CongLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevel) GetCongLevelOk() (*int32, bool) {
	if o == nil || isNil(o.CongLevel) {
		return nil, false
	}
	return o.CongLevel, true
}

// HasCongLevel returns a boolean if a field has been set.
func (o *ThresholdLevel) HasCongLevel() bool {
	if o != nil && !isNil(o.CongLevel) {
		return true
	}

	return false
}

// SetCongLevel gets a reference to the given int32 and assigns it to the CongLevel field.
func (o *ThresholdLevel) SetCongLevel(v int32) {
	o.CongLevel = &v
}

// GetNfLoadLevel returns the NfLoadLevel field value if set, zero value otherwise.
func (o *ThresholdLevel) GetNfLoadLevel() int32 {
	if o == nil || isNil(o.NfLoadLevel) {
		var ret int32
		return ret
	}
	return *o.NfLoadLevel
}

// GetNfLoadLevelOk returns a tuple with the NfLoadLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevel) GetNfLoadLevelOk() (*int32, bool) {
	if o == nil || isNil(o.NfLoadLevel) {
		return nil, false
	}
	return o.NfLoadLevel, true
}

// HasNfLoadLevel returns a boolean if a field has been set.
func (o *ThresholdLevel) HasNfLoadLevel() bool {
	if o != nil && !isNil(o.NfLoadLevel) {
		return true
	}

	return false
}

// SetNfLoadLevel gets a reference to the given int32 and assigns it to the NfLoadLevel field.
func (o *ThresholdLevel) SetNfLoadLevel(v int32) {
	o.NfLoadLevel = &v
}

// GetNfCpuUsage returns the NfCpuUsage field value if set, zero value otherwise.
func (o *ThresholdLevel) GetNfCpuUsage() int32 {
	if o == nil || isNil(o.NfCpuUsage) {
		var ret int32
		return ret
	}
	return *o.NfCpuUsage
}

// GetNfCpuUsageOk returns a tuple with the NfCpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevel) GetNfCpuUsageOk() (*int32, bool) {
	if o == nil || isNil(o.NfCpuUsage) {
		return nil, false
	}
	return o.NfCpuUsage, true
}

// HasNfCpuUsage returns a boolean if a field has been set.
func (o *ThresholdLevel) HasNfCpuUsage() bool {
	if o != nil && !isNil(o.NfCpuUsage) {
		return true
	}

	return false
}

// SetNfCpuUsage gets a reference to the given int32 and assigns it to the NfCpuUsage field.
func (o *ThresholdLevel) SetNfCpuUsage(v int32) {
	o.NfCpuUsage = &v
}

// GetNfMemoryUsage returns the NfMemoryUsage field value if set, zero value otherwise.
func (o *ThresholdLevel) GetNfMemoryUsage() int32 {
	if o == nil || isNil(o.NfMemoryUsage) {
		var ret int32
		return ret
	}
	return *o.NfMemoryUsage
}

// GetNfMemoryUsageOk returns a tuple with the NfMemoryUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevel) GetNfMemoryUsageOk() (*int32, bool) {
	if o == nil || isNil(o.NfMemoryUsage) {
		return nil, false
	}
	return o.NfMemoryUsage, true
}

// HasNfMemoryUsage returns a boolean if a field has been set.
func (o *ThresholdLevel) HasNfMemoryUsage() bool {
	if o != nil && !isNil(o.NfMemoryUsage) {
		return true
	}

	return false
}

// SetNfMemoryUsage gets a reference to the given int32 and assigns it to the NfMemoryUsage field.
func (o *ThresholdLevel) SetNfMemoryUsage(v int32) {
	o.NfMemoryUsage = &v
}

// GetNfStorageUsage returns the NfStorageUsage field value if set, zero value otherwise.
func (o *ThresholdLevel) GetNfStorageUsage() int32 {
	if o == nil || isNil(o.NfStorageUsage) {
		var ret int32
		return ret
	}
	return *o.NfStorageUsage
}

// GetNfStorageUsageOk returns a tuple with the NfStorageUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevel) GetNfStorageUsageOk() (*int32, bool) {
	if o == nil || isNil(o.NfStorageUsage) {
		return nil, false
	}
	return o.NfStorageUsage, true
}

// HasNfStorageUsage returns a boolean if a field has been set.
func (o *ThresholdLevel) HasNfStorageUsage() bool {
	if o != nil && !isNil(o.NfStorageUsage) {
		return true
	}

	return false
}

// SetNfStorageUsage gets a reference to the given int32 and assigns it to the NfStorageUsage field.
func (o *ThresholdLevel) SetNfStorageUsage(v int32) {
	o.NfStorageUsage = &v
}

// GetAvgTrafficRate returns the AvgTrafficRate field value if set, zero value otherwise.
func (o *ThresholdLevel) GetAvgTrafficRate() string {
	if o == nil || isNil(o.AvgTrafficRate) {
		var ret string
		return ret
	}
	return *o.AvgTrafficRate
}

// GetAvgTrafficRateOk returns a tuple with the AvgTrafficRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevel) GetAvgTrafficRateOk() (*string, bool) {
	if o == nil || isNil(o.AvgTrafficRate) {
		return nil, false
	}
	return o.AvgTrafficRate, true
}

// HasAvgTrafficRate returns a boolean if a field has been set.
func (o *ThresholdLevel) HasAvgTrafficRate() bool {
	if o != nil && !isNil(o.AvgTrafficRate) {
		return true
	}

	return false
}

// SetAvgTrafficRate gets a reference to the given string and assigns it to the AvgTrafficRate field.
func (o *ThresholdLevel) SetAvgTrafficRate(v string) {
	o.AvgTrafficRate = &v
}

// GetMaxTrafficRate returns the MaxTrafficRate field value if set, zero value otherwise.
func (o *ThresholdLevel) GetMaxTrafficRate() string {
	if o == nil || isNil(o.MaxTrafficRate) {
		var ret string
		return ret
	}
	return *o.MaxTrafficRate
}

// GetMaxTrafficRateOk returns a tuple with the MaxTrafficRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevel) GetMaxTrafficRateOk() (*string, bool) {
	if o == nil || isNil(o.MaxTrafficRate) {
		return nil, false
	}
	return o.MaxTrafficRate, true
}

// HasMaxTrafficRate returns a boolean if a field has been set.
func (o *ThresholdLevel) HasMaxTrafficRate() bool {
	if o != nil && !isNil(o.MaxTrafficRate) {
		return true
	}

	return false
}

// SetMaxTrafficRate gets a reference to the given string and assigns it to the MaxTrafficRate field.
func (o *ThresholdLevel) SetMaxTrafficRate(v string) {
	o.MaxTrafficRate = &v
}

// GetAvgPacketDelay returns the AvgPacketDelay field value if set, zero value otherwise.
func (o *ThresholdLevel) GetAvgPacketDelay() int32 {
	if o == nil || isNil(o.AvgPacketDelay) {
		var ret int32
		return ret
	}
	return *o.AvgPacketDelay
}

// GetAvgPacketDelayOk returns a tuple with the AvgPacketDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevel) GetAvgPacketDelayOk() (*int32, bool) {
	if o == nil || isNil(o.AvgPacketDelay) {
		return nil, false
	}
	return o.AvgPacketDelay, true
}

// HasAvgPacketDelay returns a boolean if a field has been set.
func (o *ThresholdLevel) HasAvgPacketDelay() bool {
	if o != nil && !isNil(o.AvgPacketDelay) {
		return true
	}

	return false
}

// SetAvgPacketDelay gets a reference to the given int32 and assigns it to the AvgPacketDelay field.
func (o *ThresholdLevel) SetAvgPacketDelay(v int32) {
	o.AvgPacketDelay = &v
}

// GetMaxPacketDelay returns the MaxPacketDelay field value if set, zero value otherwise.
func (o *ThresholdLevel) GetMaxPacketDelay() int32 {
	if o == nil || isNil(o.MaxPacketDelay) {
		var ret int32
		return ret
	}
	return *o.MaxPacketDelay
}

// GetMaxPacketDelayOk returns a tuple with the MaxPacketDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevel) GetMaxPacketDelayOk() (*int32, bool) {
	if o == nil || isNil(o.MaxPacketDelay) {
		return nil, false
	}
	return o.MaxPacketDelay, true
}

// HasMaxPacketDelay returns a boolean if a field has been set.
func (o *ThresholdLevel) HasMaxPacketDelay() bool {
	if o != nil && !isNil(o.MaxPacketDelay) {
		return true
	}

	return false
}

// SetMaxPacketDelay gets a reference to the given int32 and assigns it to the MaxPacketDelay field.
func (o *ThresholdLevel) SetMaxPacketDelay(v int32) {
	o.MaxPacketDelay = &v
}

// GetAvgPacketLossRate returns the AvgPacketLossRate field value if set, zero value otherwise.
func (o *ThresholdLevel) GetAvgPacketLossRate() int32 {
	if o == nil || isNil(o.AvgPacketLossRate) {
		var ret int32
		return ret
	}
	return *o.AvgPacketLossRate
}

// GetAvgPacketLossRateOk returns a tuple with the AvgPacketLossRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevel) GetAvgPacketLossRateOk() (*int32, bool) {
	if o == nil || isNil(o.AvgPacketLossRate) {
		return nil, false
	}
	return o.AvgPacketLossRate, true
}

// HasAvgPacketLossRate returns a boolean if a field has been set.
func (o *ThresholdLevel) HasAvgPacketLossRate() bool {
	if o != nil && !isNil(o.AvgPacketLossRate) {
		return true
	}

	return false
}

// SetAvgPacketLossRate gets a reference to the given int32 and assigns it to the AvgPacketLossRate field.
func (o *ThresholdLevel) SetAvgPacketLossRate(v int32) {
	o.AvgPacketLossRate = &v
}

// GetSvcExpLevel returns the SvcExpLevel field value if set, zero value otherwise.
func (o *ThresholdLevel) GetSvcExpLevel() float32 {
	if o == nil || isNil(o.SvcExpLevel) {
		var ret float32
		return ret
	}
	return *o.SvcExpLevel
}

// GetSvcExpLevelOk returns a tuple with the SvcExpLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdLevel) GetSvcExpLevelOk() (*float32, bool) {
	if o == nil || isNil(o.SvcExpLevel) {
		return nil, false
	}
	return o.SvcExpLevel, true
}

// HasSvcExpLevel returns a boolean if a field has been set.
func (o *ThresholdLevel) HasSvcExpLevel() bool {
	if o != nil && !isNil(o.SvcExpLevel) {
		return true
	}

	return false
}

// SetSvcExpLevel gets a reference to the given float32 and assigns it to the SvcExpLevel field.
func (o *ThresholdLevel) SetSvcExpLevel(v float32) {
	o.SvcExpLevel = &v
}

func (o ThresholdLevel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThresholdLevel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CongLevel) {
		toSerialize["congLevel"] = o.CongLevel
	}
	if !isNil(o.NfLoadLevel) {
		toSerialize["nfLoadLevel"] = o.NfLoadLevel
	}
	if !isNil(o.NfCpuUsage) {
		toSerialize["nfCpuUsage"] = o.NfCpuUsage
	}
	if !isNil(o.NfMemoryUsage) {
		toSerialize["nfMemoryUsage"] = o.NfMemoryUsage
	}
	if !isNil(o.NfStorageUsage) {
		toSerialize["nfStorageUsage"] = o.NfStorageUsage
	}
	if !isNil(o.AvgTrafficRate) {
		toSerialize["avgTrafficRate"] = o.AvgTrafficRate
	}
	if !isNil(o.MaxTrafficRate) {
		toSerialize["maxTrafficRate"] = o.MaxTrafficRate
	}
	if !isNil(o.AvgPacketDelay) {
		toSerialize["avgPacketDelay"] = o.AvgPacketDelay
	}
	if !isNil(o.MaxPacketDelay) {
		toSerialize["maxPacketDelay"] = o.MaxPacketDelay
	}
	if !isNil(o.AvgPacketLossRate) {
		toSerialize["avgPacketLossRate"] = o.AvgPacketLossRate
	}
	if !isNil(o.SvcExpLevel) {
		toSerialize["svcExpLevel"] = o.SvcExpLevel
	}
	return toSerialize, nil
}

type NullableThresholdLevel struct {
	value *ThresholdLevel
	isSet bool
}

func (v NullableThresholdLevel) Get() *ThresholdLevel {
	return v.value
}

func (v *NullableThresholdLevel) Set(val *ThresholdLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdLevel(val *ThresholdLevel) *NullableThresholdLevel {
	return &NullableThresholdLevel{value: val, isSet: true}
}

func (v NullableThresholdLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



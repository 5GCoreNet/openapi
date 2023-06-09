/*
3gpp-time-sync-exposure

API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_TimeSyncExposure

import (
	"encoding/json"
)

// checks if the ConfigForPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigForPort{}

// ConfigForPort Contains configuration for each port.
type ConfigForPort struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi             *string `json:"gpsi,omitempty"`
	N6Ind            *bool   `json:"n6Ind,omitempty"`
	PtpEnable        *bool   `json:"ptpEnable,omitempty"`
	LogSyncInter     *int32  `json:"logSyncInter,omitempty"`
	LogSyncInterInd  *bool   `json:"logSyncInterInd,omitempty"`
	LogAnnouInter    *int32  `json:"logAnnouInter,omitempty"`
	LogAnnouInterInd *bool   `json:"logAnnouInterInd,omitempty"`
}

// NewConfigForPort instantiates a new ConfigForPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigForPort() *ConfigForPort {
	this := ConfigForPort{}
	return &this
}

// NewConfigForPortWithDefaults instantiates a new ConfigForPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigForPortWithDefaults() *ConfigForPort {
	this := ConfigForPort{}
	return &this
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *ConfigForPort) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigForPort) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *ConfigForPort) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *ConfigForPort) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetN6Ind returns the N6Ind field value if set, zero value otherwise.
func (o *ConfigForPort) GetN6Ind() bool {
	if o == nil || IsNil(o.N6Ind) {
		var ret bool
		return ret
	}
	return *o.N6Ind
}

// GetN6IndOk returns a tuple with the N6Ind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigForPort) GetN6IndOk() (*bool, bool) {
	if o == nil || IsNil(o.N6Ind) {
		return nil, false
	}
	return o.N6Ind, true
}

// HasN6Ind returns a boolean if a field has been set.
func (o *ConfigForPort) HasN6Ind() bool {
	if o != nil && !IsNil(o.N6Ind) {
		return true
	}

	return false
}

// SetN6Ind gets a reference to the given bool and assigns it to the N6Ind field.
func (o *ConfigForPort) SetN6Ind(v bool) {
	o.N6Ind = &v
}

// GetPtpEnable returns the PtpEnable field value if set, zero value otherwise.
func (o *ConfigForPort) GetPtpEnable() bool {
	if o == nil || IsNil(o.PtpEnable) {
		var ret bool
		return ret
	}
	return *o.PtpEnable
}

// GetPtpEnableOk returns a tuple with the PtpEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigForPort) GetPtpEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.PtpEnable) {
		return nil, false
	}
	return o.PtpEnable, true
}

// HasPtpEnable returns a boolean if a field has been set.
func (o *ConfigForPort) HasPtpEnable() bool {
	if o != nil && !IsNil(o.PtpEnable) {
		return true
	}

	return false
}

// SetPtpEnable gets a reference to the given bool and assigns it to the PtpEnable field.
func (o *ConfigForPort) SetPtpEnable(v bool) {
	o.PtpEnable = &v
}

// GetLogSyncInter returns the LogSyncInter field value if set, zero value otherwise.
func (o *ConfigForPort) GetLogSyncInter() int32 {
	if o == nil || IsNil(o.LogSyncInter) {
		var ret int32
		return ret
	}
	return *o.LogSyncInter
}

// GetLogSyncInterOk returns a tuple with the LogSyncInter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigForPort) GetLogSyncInterOk() (*int32, bool) {
	if o == nil || IsNil(o.LogSyncInter) {
		return nil, false
	}
	return o.LogSyncInter, true
}

// HasLogSyncInter returns a boolean if a field has been set.
func (o *ConfigForPort) HasLogSyncInter() bool {
	if o != nil && !IsNil(o.LogSyncInter) {
		return true
	}

	return false
}

// SetLogSyncInter gets a reference to the given int32 and assigns it to the LogSyncInter field.
func (o *ConfigForPort) SetLogSyncInter(v int32) {
	o.LogSyncInter = &v
}

// GetLogSyncInterInd returns the LogSyncInterInd field value if set, zero value otherwise.
func (o *ConfigForPort) GetLogSyncInterInd() bool {
	if o == nil || IsNil(o.LogSyncInterInd) {
		var ret bool
		return ret
	}
	return *o.LogSyncInterInd
}

// GetLogSyncInterIndOk returns a tuple with the LogSyncInterInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigForPort) GetLogSyncInterIndOk() (*bool, bool) {
	if o == nil || IsNil(o.LogSyncInterInd) {
		return nil, false
	}
	return o.LogSyncInterInd, true
}

// HasLogSyncInterInd returns a boolean if a field has been set.
func (o *ConfigForPort) HasLogSyncInterInd() bool {
	if o != nil && !IsNil(o.LogSyncInterInd) {
		return true
	}

	return false
}

// SetLogSyncInterInd gets a reference to the given bool and assigns it to the LogSyncInterInd field.
func (o *ConfigForPort) SetLogSyncInterInd(v bool) {
	o.LogSyncInterInd = &v
}

// GetLogAnnouInter returns the LogAnnouInter field value if set, zero value otherwise.
func (o *ConfigForPort) GetLogAnnouInter() int32 {
	if o == nil || IsNil(o.LogAnnouInter) {
		var ret int32
		return ret
	}
	return *o.LogAnnouInter
}

// GetLogAnnouInterOk returns a tuple with the LogAnnouInter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigForPort) GetLogAnnouInterOk() (*int32, bool) {
	if o == nil || IsNil(o.LogAnnouInter) {
		return nil, false
	}
	return o.LogAnnouInter, true
}

// HasLogAnnouInter returns a boolean if a field has been set.
func (o *ConfigForPort) HasLogAnnouInter() bool {
	if o != nil && !IsNil(o.LogAnnouInter) {
		return true
	}

	return false
}

// SetLogAnnouInter gets a reference to the given int32 and assigns it to the LogAnnouInter field.
func (o *ConfigForPort) SetLogAnnouInter(v int32) {
	o.LogAnnouInter = &v
}

// GetLogAnnouInterInd returns the LogAnnouInterInd field value if set, zero value otherwise.
func (o *ConfigForPort) GetLogAnnouInterInd() bool {
	if o == nil || IsNil(o.LogAnnouInterInd) {
		var ret bool
		return ret
	}
	return *o.LogAnnouInterInd
}

// GetLogAnnouInterIndOk returns a tuple with the LogAnnouInterInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigForPort) GetLogAnnouInterIndOk() (*bool, bool) {
	if o == nil || IsNil(o.LogAnnouInterInd) {
		return nil, false
	}
	return o.LogAnnouInterInd, true
}

// HasLogAnnouInterInd returns a boolean if a field has been set.
func (o *ConfigForPort) HasLogAnnouInterInd() bool {
	if o != nil && !IsNil(o.LogAnnouInterInd) {
		return true
	}

	return false
}

// SetLogAnnouInterInd gets a reference to the given bool and assigns it to the LogAnnouInterInd field.
func (o *ConfigForPort) SetLogAnnouInterInd(v bool) {
	o.LogAnnouInterInd = &v
}

func (o ConfigForPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigForPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.N6Ind) {
		toSerialize["n6Ind"] = o.N6Ind
	}
	if !IsNil(o.PtpEnable) {
		toSerialize["ptpEnable"] = o.PtpEnable
	}
	if !IsNil(o.LogSyncInter) {
		toSerialize["logSyncInter"] = o.LogSyncInter
	}
	if !IsNil(o.LogSyncInterInd) {
		toSerialize["logSyncInterInd"] = o.LogSyncInterInd
	}
	if !IsNil(o.LogAnnouInter) {
		toSerialize["logAnnouInter"] = o.LogAnnouInter
	}
	if !IsNil(o.LogAnnouInterInd) {
		toSerialize["logAnnouInterInd"] = o.LogAnnouInterInd
	}
	return toSerialize, nil
}

type NullableConfigForPort struct {
	value *ConfigForPort
	isSet bool
}

func (v NullableConfigForPort) Get() *ConfigForPort {
	return v.value
}

func (v *NullableConfigForPort) Set(val *ConfigForPort) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigForPort) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigForPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigForPort(val *ConfigForPort) *NullableConfigForPort {
	return &NullableConfigForPort{value: val, isSet: true}
}

func (v NullableConfigForPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigForPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

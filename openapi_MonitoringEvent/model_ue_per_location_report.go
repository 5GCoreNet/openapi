/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
)

// checks if the UePerLocationReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UePerLocationReport{}

// UePerLocationReport Represents the number of UEs found at the indicated location.
type UePerLocationReport struct {
	// Identifies the number of UEs.
	UeCount int32 `json:"ueCount"`
	// Each element uniquely identifies a user.
	ExternalIds []string `json:"externalIds,omitempty"`
	// Each element identifies the MS internal PSTN/ISDN number allocated for a UE.
	Msisdns []string `json:"msisdns,omitempty"`
	// Each element uniquely identifies a UAV.
	ServLevelDevIds []string `json:"servLevelDevIds,omitempty"`
}

// NewUePerLocationReport instantiates a new UePerLocationReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUePerLocationReport(ueCount int32) *UePerLocationReport {
	this := UePerLocationReport{}
	this.UeCount = ueCount
	return &this
}

// NewUePerLocationReportWithDefaults instantiates a new UePerLocationReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUePerLocationReportWithDefaults() *UePerLocationReport {
	this := UePerLocationReport{}
	return &this
}

// GetUeCount returns the UeCount field value
func (o *UePerLocationReport) GetUeCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UeCount
}

// GetUeCountOk returns a tuple with the UeCount field value
// and a boolean to check if the value has been set.
func (o *UePerLocationReport) GetUeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UeCount, true
}

// SetUeCount sets field value
func (o *UePerLocationReport) SetUeCount(v int32) {
	o.UeCount = v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *UePerLocationReport) GetExternalIds() []string {
	if o == nil || isNil(o.ExternalIds) {
		var ret []string
		return ret
	}
	return o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UePerLocationReport) GetExternalIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *UePerLocationReport) HasExternalIds() bool {
	if o != nil && !isNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given []string and assigns it to the ExternalIds field.
func (o *UePerLocationReport) SetExternalIds(v []string) {
	o.ExternalIds = v
}

// GetMsisdns returns the Msisdns field value if set, zero value otherwise.
func (o *UePerLocationReport) GetMsisdns() []string {
	if o == nil || isNil(o.Msisdns) {
		var ret []string
		return ret
	}
	return o.Msisdns
}

// GetMsisdnsOk returns a tuple with the Msisdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UePerLocationReport) GetMsisdnsOk() ([]string, bool) {
	if o == nil || isNil(o.Msisdns) {
		return nil, false
	}
	return o.Msisdns, true
}

// HasMsisdns returns a boolean if a field has been set.
func (o *UePerLocationReport) HasMsisdns() bool {
	if o != nil && !isNil(o.Msisdns) {
		return true
	}

	return false
}

// SetMsisdns gets a reference to the given []string and assigns it to the Msisdns field.
func (o *UePerLocationReport) SetMsisdns(v []string) {
	o.Msisdns = v
}

// GetServLevelDevIds returns the ServLevelDevIds field value if set, zero value otherwise.
func (o *UePerLocationReport) GetServLevelDevIds() []string {
	if o == nil || isNil(o.ServLevelDevIds) {
		var ret []string
		return ret
	}
	return o.ServLevelDevIds
}

// GetServLevelDevIdsOk returns a tuple with the ServLevelDevIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UePerLocationReport) GetServLevelDevIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ServLevelDevIds) {
		return nil, false
	}
	return o.ServLevelDevIds, true
}

// HasServLevelDevIds returns a boolean if a field has been set.
func (o *UePerLocationReport) HasServLevelDevIds() bool {
	if o != nil && !isNil(o.ServLevelDevIds) {
		return true
	}

	return false
}

// SetServLevelDevIds gets a reference to the given []string and assigns it to the ServLevelDevIds field.
func (o *UePerLocationReport) SetServLevelDevIds(v []string) {
	o.ServLevelDevIds = v
}

func (o UePerLocationReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UePerLocationReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ueCount"] = o.UeCount
	if !isNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	if !isNil(o.Msisdns) {
		toSerialize["msisdns"] = o.Msisdns
	}
	if !isNil(o.ServLevelDevIds) {
		toSerialize["servLevelDevIds"] = o.ServLevelDevIds
	}
	return toSerialize, nil
}

type NullableUePerLocationReport struct {
	value *UePerLocationReport
	isSet bool
}

func (v NullableUePerLocationReport) Get() *UePerLocationReport {
	return v.value
}

func (v *NullableUePerLocationReport) Set(val *UePerLocationReport) {
	v.value = val
	v.isSet = true
}

func (v NullableUePerLocationReport) IsSet() bool {
	return v.isSet
}

func (v *NullableUePerLocationReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUePerLocationReport(val *UePerLocationReport) *NullableUePerLocationReport {
	return &NullableUePerLocationReport{value: val, isSet: true}
}

func (v NullableUePerLocationReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUePerLocationReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



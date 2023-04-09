/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_EE

import (
	"encoding/json"
	"time"
)

// checks if the DataRestorationNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataRestorationNotification{}

// DataRestorationNotification Contains identities representing those UEs potentially affected by a data-loss event at the UDR
type DataRestorationNotification struct {
	// string with format 'date-time' as defined in OpenAPI.
	LastReplicationTime *time.Time `json:"lastReplicationTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime *time.Time      `json:"recoveryTime,omitempty"`
	PlmnId       *PlmnId         `json:"plmnId,omitempty"`
	SupiRanges   []SupiRange     `json:"supiRanges,omitempty"`
	GpsiRanges   []IdentityRange `json:"gpsiRanges,omitempty"`
	ResetIds     []string        `json:"resetIds,omitempty"`
	SNssaiList   []Snssai        `json:"sNssaiList,omitempty"`
	DnnList      []string        `json:"dnnList,omitempty"`
	// Identifier of a group of NFs.
	UdmGroupId *string `json:"udmGroupId,omitempty"`
}

// NewDataRestorationNotification instantiates a new DataRestorationNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataRestorationNotification() *DataRestorationNotification {
	this := DataRestorationNotification{}
	return &this
}

// NewDataRestorationNotificationWithDefaults instantiates a new DataRestorationNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataRestorationNotificationWithDefaults() *DataRestorationNotification {
	this := DataRestorationNotification{}
	return &this
}

// GetLastReplicationTime returns the LastReplicationTime field value if set, zero value otherwise.
func (o *DataRestorationNotification) GetLastReplicationTime() time.Time {
	if o == nil || IsNil(o.LastReplicationTime) {
		var ret time.Time
		return ret
	}
	return *o.LastReplicationTime
}

// GetLastReplicationTimeOk returns a tuple with the LastReplicationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetLastReplicationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastReplicationTime) {
		return nil, false
	}
	return o.LastReplicationTime, true
}

// HasLastReplicationTime returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasLastReplicationTime() bool {
	if o != nil && !IsNil(o.LastReplicationTime) {
		return true
	}

	return false
}

// SetLastReplicationTime gets a reference to the given time.Time and assigns it to the LastReplicationTime field.
func (o *DataRestorationNotification) SetLastReplicationTime(v time.Time) {
	o.LastReplicationTime = &v
}

// GetRecoveryTime returns the RecoveryTime field value if set, zero value otherwise.
func (o *DataRestorationNotification) GetRecoveryTime() time.Time {
	if o == nil || IsNil(o.RecoveryTime) {
		var ret time.Time
		return ret
	}
	return *o.RecoveryTime
}

// GetRecoveryTimeOk returns a tuple with the RecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetRecoveryTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RecoveryTime) {
		return nil, false
	}
	return o.RecoveryTime, true
}

// HasRecoveryTime returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasRecoveryTime() bool {
	if o != nil && !IsNil(o.RecoveryTime) {
		return true
	}

	return false
}

// SetRecoveryTime gets a reference to the given time.Time and assigns it to the RecoveryTime field.
func (o *DataRestorationNotification) SetRecoveryTime(v time.Time) {
	o.RecoveryTime = &v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *DataRestorationNotification) GetPlmnId() PlmnId {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *DataRestorationNotification) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *DataRestorationNotification) GetSupiRanges() []SupiRange {
	if o == nil || IsNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || IsNil(o.SupiRanges) {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasSupiRanges() bool {
	if o != nil && !IsNil(o.SupiRanges) {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *DataRestorationNotification) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *DataRestorationNotification) GetGpsiRanges() []IdentityRange {
	if o == nil || IsNil(o.GpsiRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.GpsiRanges) {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasGpsiRanges() bool {
	if o != nil && !IsNil(o.GpsiRanges) {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *DataRestorationNotification) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *DataRestorationNotification) GetResetIds() []string {
	if o == nil || IsNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetResetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResetIds) {
		return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasResetIds() bool {
	if o != nil && !IsNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *DataRestorationNotification) SetResetIds(v []string) {
	o.ResetIds = v
}

// GetSNssaiList returns the SNssaiList field value if set, zero value otherwise.
func (o *DataRestorationNotification) GetSNssaiList() []Snssai {
	if o == nil || IsNil(o.SNssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SNssaiList
}

// GetSNssaiListOk returns a tuple with the SNssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetSNssaiListOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.SNssaiList) {
		return nil, false
	}
	return o.SNssaiList, true
}

// HasSNssaiList returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasSNssaiList() bool {
	if o != nil && !IsNil(o.SNssaiList) {
		return true
	}

	return false
}

// SetSNssaiList gets a reference to the given []Snssai and assigns it to the SNssaiList field.
func (o *DataRestorationNotification) SetSNssaiList(v []Snssai) {
	o.SNssaiList = v
}

// GetDnnList returns the DnnList field value if set, zero value otherwise.
func (o *DataRestorationNotification) GetDnnList() []string {
	if o == nil || IsNil(o.DnnList) {
		var ret []string
		return ret
	}
	return o.DnnList
}

// GetDnnListOk returns a tuple with the DnnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetDnnListOk() ([]string, bool) {
	if o == nil || IsNil(o.DnnList) {
		return nil, false
	}
	return o.DnnList, true
}

// HasDnnList returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasDnnList() bool {
	if o != nil && !IsNil(o.DnnList) {
		return true
	}

	return false
}

// SetDnnList gets a reference to the given []string and assigns it to the DnnList field.
func (o *DataRestorationNotification) SetDnnList(v []string) {
	o.DnnList = v
}

// GetUdmGroupId returns the UdmGroupId field value if set, zero value otherwise.
func (o *DataRestorationNotification) GetUdmGroupId() string {
	if o == nil || IsNil(o.UdmGroupId) {
		var ret string
		return ret
	}
	return *o.UdmGroupId
}

// GetUdmGroupIdOk returns a tuple with the UdmGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetUdmGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.UdmGroupId) {
		return nil, false
	}
	return o.UdmGroupId, true
}

// HasUdmGroupId returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasUdmGroupId() bool {
	if o != nil && !IsNil(o.UdmGroupId) {
		return true
	}

	return false
}

// SetUdmGroupId gets a reference to the given string and assigns it to the UdmGroupId field.
func (o *DataRestorationNotification) SetUdmGroupId(v string) {
	o.UdmGroupId = &v
}

func (o DataRestorationNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataRestorationNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastReplicationTime) {
		toSerialize["lastReplicationTime"] = o.LastReplicationTime
	}
	if !IsNil(o.RecoveryTime) {
		toSerialize["recoveryTime"] = o.RecoveryTime
	}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !IsNil(o.SupiRanges) {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if !IsNil(o.GpsiRanges) {
		toSerialize["gpsiRanges"] = o.GpsiRanges
	}
	if !IsNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	if !IsNil(o.SNssaiList) {
		toSerialize["sNssaiList"] = o.SNssaiList
	}
	if !IsNil(o.DnnList) {
		toSerialize["dnnList"] = o.DnnList
	}
	if !IsNil(o.UdmGroupId) {
		toSerialize["udmGroupId"] = o.UdmGroupId
	}
	return toSerialize, nil
}

type NullableDataRestorationNotification struct {
	value *DataRestorationNotification
	isSet bool
}

func (v NullableDataRestorationNotification) Get() *DataRestorationNotification {
	return v.value
}

func (v *NullableDataRestorationNotification) Set(val *DataRestorationNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableDataRestorationNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableDataRestorationNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataRestorationNotification(val *DataRestorationNotification) *NullableDataRestorationNotification {
	return &NullableDataRestorationNotification{value: val, isSet: true}
}

func (v NullableDataRestorationNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataRestorationNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

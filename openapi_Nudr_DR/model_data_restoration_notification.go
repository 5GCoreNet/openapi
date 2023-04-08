/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the DataRestorationNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataRestorationNotification{}

// DataRestorationNotification Contains identities representing those UEs potentially affected by a data-loss event at the UDR
type DataRestorationNotification struct {
	SupiRanges []SupiRange `json:"supiRanges,omitempty"`
	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
	SNssaiList []Snssai `json:"sNssaiList,omitempty"`
	DnnList []string `json:"dnnList,omitempty"`
	// Identifier of a group of NFs.
	UdrGroupId *string `json:"udrGroupId,omitempty"`
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

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *DataRestorationNotification) GetSupiRanges() []SupiRange {
	if o == nil || isNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || isNil(o.SupiRanges) {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasSupiRanges() bool {
	if o != nil && !isNil(o.SupiRanges) {
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
	if o == nil || isNil(o.GpsiRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || isNil(o.GpsiRanges) {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasGpsiRanges() bool {
	if o != nil && !isNil(o.GpsiRanges) {
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
	if o == nil || isNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetResetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ResetIds) {
		return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasResetIds() bool {
	if o != nil && !isNil(o.ResetIds) {
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
	if o == nil || isNil(o.SNssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SNssaiList
}

// GetSNssaiListOk returns a tuple with the SNssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetSNssaiListOk() ([]Snssai, bool) {
	if o == nil || isNil(o.SNssaiList) {
		return nil, false
	}
	return o.SNssaiList, true
}

// HasSNssaiList returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasSNssaiList() bool {
	if o != nil && !isNil(o.SNssaiList) {
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
	if o == nil || isNil(o.DnnList) {
		var ret []string
		return ret
	}
	return o.DnnList
}

// GetDnnListOk returns a tuple with the DnnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetDnnListOk() ([]string, bool) {
	if o == nil || isNil(o.DnnList) {
		return nil, false
	}
	return o.DnnList, true
}

// HasDnnList returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasDnnList() bool {
	if o != nil && !isNil(o.DnnList) {
		return true
	}

	return false
}

// SetDnnList gets a reference to the given []string and assigns it to the DnnList field.
func (o *DataRestorationNotification) SetDnnList(v []string) {
	o.DnnList = v
}

// GetUdrGroupId returns the UdrGroupId field value if set, zero value otherwise.
func (o *DataRestorationNotification) GetUdrGroupId() string {
	if o == nil || isNil(o.UdrGroupId) {
		var ret string
		return ret
	}
	return *o.UdrGroupId
}

// GetUdrGroupIdOk returns a tuple with the UdrGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataRestorationNotification) GetUdrGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.UdrGroupId) {
		return nil, false
	}
	return o.UdrGroupId, true
}

// HasUdrGroupId returns a boolean if a field has been set.
func (o *DataRestorationNotification) HasUdrGroupId() bool {
	if o != nil && !isNil(o.UdrGroupId) {
		return true
	}

	return false
}

// SetUdrGroupId gets a reference to the given string and assigns it to the UdrGroupId field.
func (o *DataRestorationNotification) SetUdrGroupId(v string) {
	o.UdrGroupId = &v
}

func (o DataRestorationNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataRestorationNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SupiRanges) {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if !isNil(o.GpsiRanges) {
		toSerialize["gpsiRanges"] = o.GpsiRanges
	}
	if !isNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	if !isNil(o.SNssaiList) {
		toSerialize["sNssaiList"] = o.SNssaiList
	}
	if !isNil(o.DnnList) {
		toSerialize["dnnList"] = o.DnnList
	}
	if !isNil(o.UdrGroupId) {
		toSerialize["udrGroupId"] = o.UdrGroupId
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



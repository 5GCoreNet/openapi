/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"time"
)

// checks if the MmTransactionSliceReportItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MmTransactionSliceReportItem{}

// MmTransactionSliceReportItem UE MM Transaction Report Item per Slice
type MmTransactionSliceReportItem struct {
	Snssai *Snssai `json:"snssai,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`
	Transactions int32 `json:"transactions"`
}

// NewMmTransactionSliceReportItem instantiates a new MmTransactionSliceReportItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMmTransactionSliceReportItem(timestamp time.Time, transactions int32) *MmTransactionSliceReportItem {
	this := MmTransactionSliceReportItem{}
	this.Timestamp = timestamp
	this.Transactions = transactions
	return &this
}

// NewMmTransactionSliceReportItemWithDefaults instantiates a new MmTransactionSliceReportItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMmTransactionSliceReportItemWithDefaults() *MmTransactionSliceReportItem {
	this := MmTransactionSliceReportItem{}
	return &this
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *MmTransactionSliceReportItem) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmTransactionSliceReportItem) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *MmTransactionSliceReportItem) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *MmTransactionSliceReportItem) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetTimestamp returns the Timestamp field value
func (o *MmTransactionSliceReportItem) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *MmTransactionSliceReportItem) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *MmTransactionSliceReportItem) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetTransactions returns the Transactions field value
func (o *MmTransactionSliceReportItem) GetTransactions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *MmTransactionSliceReportItem) GetTransactionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *MmTransactionSliceReportItem) SetTransactions(v int32) {
	o.Transactions = v
}

func (o MmTransactionSliceReportItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MmTransactionSliceReportItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["transactions"] = o.Transactions
	return toSerialize, nil
}

type NullableMmTransactionSliceReportItem struct {
	value *MmTransactionSliceReportItem
	isSet bool
}

func (v NullableMmTransactionSliceReportItem) Get() *MmTransactionSliceReportItem {
	return v.value
}

func (v *NullableMmTransactionSliceReportItem) Set(val *MmTransactionSliceReportItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMmTransactionSliceReportItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMmTransactionSliceReportItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMmTransactionSliceReportItem(val *MmTransactionSliceReportItem) *NullableMmTransactionSliceReportItem {
	return &NullableMmTransactionSliceReportItem{value: val, isSet: true}
}

func (v NullableMmTransactionSliceReportItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMmTransactionSliceReportItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// checks if the TransactionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionInfo{}

// TransactionInfo Represents SMF Transaction Information.
type TransactionInfo struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Transaction int32 `json:"transaction"`
	Snssai *Snssai `json:"snssai,omitempty"`
	AppIds []string `json:"appIds,omitempty"`
	TransacMetrics []TransactionMetric `json:"transacMetrics,omitempty"`
}

// NewTransactionInfo instantiates a new TransactionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionInfo(transaction int32) *TransactionInfo {
	this := TransactionInfo{}
	this.Transaction = transaction
	return &this
}

// NewTransactionInfoWithDefaults instantiates a new TransactionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionInfoWithDefaults() *TransactionInfo {
	this := TransactionInfo{}
	return &this
}

// GetTransaction returns the Transaction field value
func (o *TransactionInfo) GetTransaction() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetTransactionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *TransactionInfo) SetTransaction(v int32) {
	o.Transaction = v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *TransactionInfo) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *TransactionInfo) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *TransactionInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *TransactionInfo) GetAppIds() []string {
	if o == nil || isNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetAppIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AppIds) {
		return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *TransactionInfo) HasAppIds() bool {
	if o != nil && !isNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *TransactionInfo) SetAppIds(v []string) {
	o.AppIds = v
}

// GetTransacMetrics returns the TransacMetrics field value if set, zero value otherwise.
func (o *TransactionInfo) GetTransacMetrics() []TransactionMetric {
	if o == nil || isNil(o.TransacMetrics) {
		var ret []TransactionMetric
		return ret
	}
	return o.TransacMetrics
}

// GetTransacMetricsOk returns a tuple with the TransacMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionInfo) GetTransacMetricsOk() ([]TransactionMetric, bool) {
	if o == nil || isNil(o.TransacMetrics) {
		return nil, false
	}
	return o.TransacMetrics, true
}

// HasTransacMetrics returns a boolean if a field has been set.
func (o *TransactionInfo) HasTransacMetrics() bool {
	if o != nil && !isNil(o.TransacMetrics) {
		return true
	}

	return false
}

// SetTransacMetrics gets a reference to the given []TransactionMetric and assigns it to the TransacMetrics field.
func (o *TransactionInfo) SetTransacMetrics(v []TransactionMetric) {
	o.TransacMetrics = v
}

func (o TransactionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction"] = o.Transaction
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !isNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	if !isNil(o.TransacMetrics) {
		toSerialize["transacMetrics"] = o.TransacMetrics
	}
	return toSerialize, nil
}

type NullableTransactionInfo struct {
	value *TransactionInfo
	isSet bool
}

func (v NullableTransactionInfo) Get() *TransactionInfo {
	return v.value
}

func (v *NullableTransactionInfo) Set(val *TransactionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInfo(val *TransactionInfo) *NullableTransactionInfo {
	return &NullableTransactionInfo{value: val, isSet: true}
}

func (v NullableTransactionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



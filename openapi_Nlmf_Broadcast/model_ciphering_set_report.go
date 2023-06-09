/*
LMF Broadcast

LMF Broadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Broadcast

import (
	"encoding/json"
)

// checks if the CipheringSetReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CipheringSetReport{}

// CipheringSetReport Represents a report of Ciphering Data Set storage.
type CipheringSetReport struct {
	// Ciphering Data Set Identifier.
	CipheringSetID int32          `json:"cipheringSetID"`
	StorageOutcome StorageOutcome `json:"storageOutcome"`
}

// NewCipheringSetReport instantiates a new CipheringSetReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCipheringSetReport(cipheringSetID int32, storageOutcome StorageOutcome) *CipheringSetReport {
	this := CipheringSetReport{}
	this.CipheringSetID = cipheringSetID
	this.StorageOutcome = storageOutcome
	return &this
}

// NewCipheringSetReportWithDefaults instantiates a new CipheringSetReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCipheringSetReportWithDefaults() *CipheringSetReport {
	this := CipheringSetReport{}
	return &this
}

// GetCipheringSetID returns the CipheringSetID field value
func (o *CipheringSetReport) GetCipheringSetID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CipheringSetID
}

// GetCipheringSetIDOk returns a tuple with the CipheringSetID field value
// and a boolean to check if the value has been set.
func (o *CipheringSetReport) GetCipheringSetIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CipheringSetID, true
}

// SetCipheringSetID sets field value
func (o *CipheringSetReport) SetCipheringSetID(v int32) {
	o.CipheringSetID = v
}

// GetStorageOutcome returns the StorageOutcome field value
func (o *CipheringSetReport) GetStorageOutcome() StorageOutcome {
	if o == nil {
		var ret StorageOutcome
		return ret
	}

	return o.StorageOutcome
}

// GetStorageOutcomeOk returns a tuple with the StorageOutcome field value
// and a boolean to check if the value has been set.
func (o *CipheringSetReport) GetStorageOutcomeOk() (*StorageOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageOutcome, true
}

// SetStorageOutcome sets field value
func (o *CipheringSetReport) SetStorageOutcome(v StorageOutcome) {
	o.StorageOutcome = v
}

func (o CipheringSetReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CipheringSetReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cipheringSetID"] = o.CipheringSetID
	toSerialize["storageOutcome"] = o.StorageOutcome
	return toSerialize, nil
}

type NullableCipheringSetReport struct {
	value *CipheringSetReport
	isSet bool
}

func (v NullableCipheringSetReport) Get() *CipheringSetReport {
	return v.value
}

func (v *NullableCipheringSetReport) Set(val *CipheringSetReport) {
	v.value = val
	v.isSet = true
}

func (v NullableCipheringSetReport) IsSet() bool {
	return v.isSet
}

func (v *NullableCipheringSetReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCipheringSetReport(val *CipheringSetReport) *NullableCipheringSetReport {
	return &NullableCipheringSetReport{value: val, isSet: true}
}

func (v NullableCipheringSetReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCipheringSetReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

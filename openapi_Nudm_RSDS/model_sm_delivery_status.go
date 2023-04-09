/*
Nudm_ReportSMDeliveryStatus

UDM Report SM Delivery Status Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_RSDS

import (
	"encoding/json"
)

// checks if the SmDeliveryStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmDeliveryStatus{}

// SmDeliveryStatus Represents SM Delivery Status.
type SmDeliveryStatus struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi           string `json:"gpsi"`
	SmStatusReport string `json:"smStatusReport"`
}

// NewSmDeliveryStatus instantiates a new SmDeliveryStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmDeliveryStatus(gpsi string, smStatusReport string) *SmDeliveryStatus {
	this := SmDeliveryStatus{}
	this.Gpsi = gpsi
	this.SmStatusReport = smStatusReport
	return &this
}

// NewSmDeliveryStatusWithDefaults instantiates a new SmDeliveryStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmDeliveryStatusWithDefaults() *SmDeliveryStatus {
	this := SmDeliveryStatus{}
	return &this
}

// GetGpsi returns the Gpsi field value
func (o *SmDeliveryStatus) GetGpsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value
// and a boolean to check if the value has been set.
func (o *SmDeliveryStatus) GetGpsiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gpsi, true
}

// SetGpsi sets field value
func (o *SmDeliveryStatus) SetGpsi(v string) {
	o.Gpsi = v
}

// GetSmStatusReport returns the SmStatusReport field value
func (o *SmDeliveryStatus) GetSmStatusReport() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmStatusReport
}

// GetSmStatusReportOk returns a tuple with the SmStatusReport field value
// and a boolean to check if the value has been set.
func (o *SmDeliveryStatus) GetSmStatusReportOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmStatusReport, true
}

// SetSmStatusReport sets field value
func (o *SmDeliveryStatus) SetSmStatusReport(v string) {
	o.SmStatusReport = v
}

func (o SmDeliveryStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmDeliveryStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gpsi"] = o.Gpsi
	toSerialize["smStatusReport"] = o.SmStatusReport
	return toSerialize, nil
}

type NullableSmDeliveryStatus struct {
	value *SmDeliveryStatus
	isSet bool
}

func (v NullableSmDeliveryStatus) Get() *SmDeliveryStatus {
	return v.value
}

func (v *NullableSmDeliveryStatus) Set(val *SmDeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSmDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSmDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmDeliveryStatus(val *SmDeliveryStatus) *NullableSmDeliveryStatus {
	return &NullableSmDeliveryStatus{value: val, isSet: true}
}

func (v NullableSmDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
MSGG_N3GDelivery

API for MSGG N3G Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSGG_N3GDelivery

import (
	"encoding/json"
)

// checks if the DeliveryStatusReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryStatusReport{}

// DeliveryStatusReport Contains the delivery status report data
type DeliveryStatusReport struct {
	OriAddr      Address              `json:"oriAddr"`
	DestAddr     Address              `json:"destAddr"`
	MsgId        string               `json:"msgId"`
	SecCred      *string              `json:"secCred,omitempty"`
	FailureCause *string              `json:"failureCause,omitempty"`
	DelivSt      ReportDeliveryStatus `json:"delivSt"`
}

// NewDeliveryStatusReport instantiates a new DeliveryStatusReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryStatusReport(oriAddr Address, destAddr Address, msgId string, delivSt ReportDeliveryStatus) *DeliveryStatusReport {
	this := DeliveryStatusReport{}
	this.OriAddr = oriAddr
	this.DestAddr = destAddr
	this.MsgId = msgId
	this.DelivSt = delivSt
	return &this
}

// NewDeliveryStatusReportWithDefaults instantiates a new DeliveryStatusReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryStatusReportWithDefaults() *DeliveryStatusReport {
	this := DeliveryStatusReport{}
	return &this
}

// GetOriAddr returns the OriAddr field value
func (o *DeliveryStatusReport) GetOriAddr() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.OriAddr
}

// GetOriAddrOk returns a tuple with the OriAddr field value
// and a boolean to check if the value has been set.
func (o *DeliveryStatusReport) GetOriAddrOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriAddr, true
}

// SetOriAddr sets field value
func (o *DeliveryStatusReport) SetOriAddr(v Address) {
	o.OriAddr = v
}

// GetDestAddr returns the DestAddr field value
func (o *DeliveryStatusReport) GetDestAddr() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.DestAddr
}

// GetDestAddrOk returns a tuple with the DestAddr field value
// and a boolean to check if the value has been set.
func (o *DeliveryStatusReport) GetDestAddrOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestAddr, true
}

// SetDestAddr sets field value
func (o *DeliveryStatusReport) SetDestAddr(v Address) {
	o.DestAddr = v
}

// GetMsgId returns the MsgId field value
func (o *DeliveryStatusReport) GetMsgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MsgId
}

// GetMsgIdOk returns a tuple with the MsgId field value
// and a boolean to check if the value has been set.
func (o *DeliveryStatusReport) GetMsgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsgId, true
}

// SetMsgId sets field value
func (o *DeliveryStatusReport) SetMsgId(v string) {
	o.MsgId = v
}

// GetSecCred returns the SecCred field value if set, zero value otherwise.
func (o *DeliveryStatusReport) GetSecCred() string {
	if o == nil || IsNil(o.SecCred) {
		var ret string
		return ret
	}
	return *o.SecCred
}

// GetSecCredOk returns a tuple with the SecCred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryStatusReport) GetSecCredOk() (*string, bool) {
	if o == nil || IsNil(o.SecCred) {
		return nil, false
	}
	return o.SecCred, true
}

// HasSecCred returns a boolean if a field has been set.
func (o *DeliveryStatusReport) HasSecCred() bool {
	if o != nil && !IsNil(o.SecCred) {
		return true
	}

	return false
}

// SetSecCred gets a reference to the given string and assigns it to the SecCred field.
func (o *DeliveryStatusReport) SetSecCred(v string) {
	o.SecCred = &v
}

// GetFailureCause returns the FailureCause field value if set, zero value otherwise.
func (o *DeliveryStatusReport) GetFailureCause() string {
	if o == nil || IsNil(o.FailureCause) {
		var ret string
		return ret
	}
	return *o.FailureCause
}

// GetFailureCauseOk returns a tuple with the FailureCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryStatusReport) GetFailureCauseOk() (*string, bool) {
	if o == nil || IsNil(o.FailureCause) {
		return nil, false
	}
	return o.FailureCause, true
}

// HasFailureCause returns a boolean if a field has been set.
func (o *DeliveryStatusReport) HasFailureCause() bool {
	if o != nil && !IsNil(o.FailureCause) {
		return true
	}

	return false
}

// SetFailureCause gets a reference to the given string and assigns it to the FailureCause field.
func (o *DeliveryStatusReport) SetFailureCause(v string) {
	o.FailureCause = &v
}

// GetDelivSt returns the DelivSt field value
func (o *DeliveryStatusReport) GetDelivSt() ReportDeliveryStatus {
	if o == nil {
		var ret ReportDeliveryStatus
		return ret
	}

	return o.DelivSt
}

// GetDelivStOk returns a tuple with the DelivSt field value
// and a boolean to check if the value has been set.
func (o *DeliveryStatusReport) GetDelivStOk() (*ReportDeliveryStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DelivSt, true
}

// SetDelivSt sets field value
func (o *DeliveryStatusReport) SetDelivSt(v ReportDeliveryStatus) {
	o.DelivSt = v
}

func (o DeliveryStatusReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryStatusReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["oriAddr"] = o.OriAddr
	toSerialize["destAddr"] = o.DestAddr
	toSerialize["msgId"] = o.MsgId
	if !IsNil(o.SecCred) {
		toSerialize["secCred"] = o.SecCred
	}
	if !IsNil(o.FailureCause) {
		toSerialize["failureCause"] = o.FailureCause
	}
	toSerialize["delivSt"] = o.DelivSt
	return toSerialize, nil
}

type NullableDeliveryStatusReport struct {
	value *DeliveryStatusReport
	isSet bool
}

func (v NullableDeliveryStatusReport) Get() *DeliveryStatusReport {
	return v.value
}

func (v *NullableDeliveryStatusReport) Set(val *DeliveryStatusReport) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryStatusReport) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryStatusReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryStatusReport(val *DeliveryStatusReport) *NullableDeliveryStatusReport {
	return &NullableDeliveryStatusReport{value: val, isSet: true}
}

func (v NullableDeliveryStatusReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryStatusReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

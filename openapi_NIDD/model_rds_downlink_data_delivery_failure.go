/*
3gpp-nidd

API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NIDD

import (
	"encoding/json"
	"time"
)

// checks if the RdsDownlinkDataDeliveryFailure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RdsDownlinkDataDeliveryFailure{}

// RdsDownlinkDataDeliveryFailure Represents the failure delivery result for RDS.
type RdsDownlinkDataDeliveryFailure struct {
	ProblemDetails
	// string with format \"date-time\" as defined in OpenAPI.
	RequestedRetransmissionTime *time.Time `json:"requestedRetransmissionTime,omitempty"`
	// Indicates the serialization format(s) that are supported by the UE on the associated RDS port.
	SupportedUeFormats []SerializationFormat `json:"supportedUeFormats,omitempty"`
}

// NewRdsDownlinkDataDeliveryFailure instantiates a new RdsDownlinkDataDeliveryFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRdsDownlinkDataDeliveryFailure() *RdsDownlinkDataDeliveryFailure {
	this := RdsDownlinkDataDeliveryFailure{}
	return &this
}

// NewRdsDownlinkDataDeliveryFailureWithDefaults instantiates a new RdsDownlinkDataDeliveryFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRdsDownlinkDataDeliveryFailureWithDefaults() *RdsDownlinkDataDeliveryFailure {
	this := RdsDownlinkDataDeliveryFailure{}
	return &this
}

// GetRequestedRetransmissionTime returns the RequestedRetransmissionTime field value if set, zero value otherwise.
func (o *RdsDownlinkDataDeliveryFailure) GetRequestedRetransmissionTime() time.Time {
	if o == nil || IsNil(o.RequestedRetransmissionTime) {
		var ret time.Time
		return ret
	}
	return *o.RequestedRetransmissionTime
}

// GetRequestedRetransmissionTimeOk returns a tuple with the RequestedRetransmissionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsDownlinkDataDeliveryFailure) GetRequestedRetransmissionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestedRetransmissionTime) {
		return nil, false
	}
	return o.RequestedRetransmissionTime, true
}

// HasRequestedRetransmissionTime returns a boolean if a field has been set.
func (o *RdsDownlinkDataDeliveryFailure) HasRequestedRetransmissionTime() bool {
	if o != nil && !IsNil(o.RequestedRetransmissionTime) {
		return true
	}

	return false
}

// SetRequestedRetransmissionTime gets a reference to the given time.Time and assigns it to the RequestedRetransmissionTime field.
func (o *RdsDownlinkDataDeliveryFailure) SetRequestedRetransmissionTime(v time.Time) {
	o.RequestedRetransmissionTime = &v
}

// GetSupportedUeFormats returns the SupportedUeFormats field value if set, zero value otherwise.
func (o *RdsDownlinkDataDeliveryFailure) GetSupportedUeFormats() []SerializationFormat {
	if o == nil || IsNil(o.SupportedUeFormats) {
		var ret []SerializationFormat
		return ret
	}
	return o.SupportedUeFormats
}

// GetSupportedUeFormatsOk returns a tuple with the SupportedUeFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsDownlinkDataDeliveryFailure) GetSupportedUeFormatsOk() ([]SerializationFormat, bool) {
	if o == nil || IsNil(o.SupportedUeFormats) {
		return nil, false
	}
	return o.SupportedUeFormats, true
}

// HasSupportedUeFormats returns a boolean if a field has been set.
func (o *RdsDownlinkDataDeliveryFailure) HasSupportedUeFormats() bool {
	if o != nil && !IsNil(o.SupportedUeFormats) {
		return true
	}

	return false
}

// SetSupportedUeFormats gets a reference to the given []SerializationFormat and assigns it to the SupportedUeFormats field.
func (o *RdsDownlinkDataDeliveryFailure) SetSupportedUeFormats(v []SerializationFormat) {
	o.SupportedUeFormats = v
}

func (o RdsDownlinkDataDeliveryFailure) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RdsDownlinkDataDeliveryFailure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedProblemDetails, errProblemDetails := json.Marshal(o.ProblemDetails)
	if errProblemDetails != nil {
		return map[string]interface{}{}, errProblemDetails
	}
	errProblemDetails = json.Unmarshal([]byte(serializedProblemDetails), &toSerialize)
	if errProblemDetails != nil {
		return map[string]interface{}{}, errProblemDetails
	}
	if !IsNil(o.RequestedRetransmissionTime) {
		toSerialize["requestedRetransmissionTime"] = o.RequestedRetransmissionTime
	}
	if !IsNil(o.SupportedUeFormats) {
		toSerialize["supportedUeFormats"] = o.SupportedUeFormats
	}
	return toSerialize, nil
}

type NullableRdsDownlinkDataDeliveryFailure struct {
	value *RdsDownlinkDataDeliveryFailure
	isSet bool
}

func (v NullableRdsDownlinkDataDeliveryFailure) Get() *RdsDownlinkDataDeliveryFailure {
	return v.value
}

func (v *NullableRdsDownlinkDataDeliveryFailure) Set(val *RdsDownlinkDataDeliveryFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableRdsDownlinkDataDeliveryFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableRdsDownlinkDataDeliveryFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRdsDownlinkDataDeliveryFailure(val *RdsDownlinkDataDeliveryFailure) *NullableRdsDownlinkDataDeliveryFailure {
	return &NullableRdsDownlinkDataDeliveryFailure{value: val, isSet: true}
}

func (v NullableRdsDownlinkDataDeliveryFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRdsDownlinkDataDeliveryFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

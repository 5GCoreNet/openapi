/*
3gpp-bdt

API for BDT resouce management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ResourceManagementOfBdt

import (
	"encoding/json"
)

// checks if the TransferPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferPolicy{}

// TransferPolicy Represents an offered transfer policy sent from the SCEF to the SCS/AS, or a selected transfer policy sent from the SCS/AS to the SCEF.
type TransferPolicy struct {
	// Identifier for the transfer policy
	BdtPolicyId int32 `json:"bdtPolicyId"`
	// integer indicating a bandwidth in bits per second.
	MaxUplinkBandwidth *int32 `json:"maxUplinkBandwidth,omitempty"`
	// integer indicating a bandwidth in bits per second.
	MaxDownlinkBandwidth *int32 `json:"maxDownlinkBandwidth,omitempty"`
	// Indicates the rating group during the time window.
	RatingGroup int32      `json:"ratingGroup"`
	TimeWindow  TimeWindow `json:"timeWindow"`
}

// NewTransferPolicy instantiates a new TransferPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferPolicy(bdtPolicyId int32, ratingGroup int32, timeWindow TimeWindow) *TransferPolicy {
	this := TransferPolicy{}
	this.BdtPolicyId = bdtPolicyId
	this.RatingGroup = ratingGroup
	this.TimeWindow = timeWindow
	return &this
}

// NewTransferPolicyWithDefaults instantiates a new TransferPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferPolicyWithDefaults() *TransferPolicy {
	this := TransferPolicy{}
	return &this
}

// GetBdtPolicyId returns the BdtPolicyId field value
func (o *TransferPolicy) GetBdtPolicyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BdtPolicyId
}

// GetBdtPolicyIdOk returns a tuple with the BdtPolicyId field value
// and a boolean to check if the value has been set.
func (o *TransferPolicy) GetBdtPolicyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BdtPolicyId, true
}

// SetBdtPolicyId sets field value
func (o *TransferPolicy) SetBdtPolicyId(v int32) {
	o.BdtPolicyId = v
}

// GetMaxUplinkBandwidth returns the MaxUplinkBandwidth field value if set, zero value otherwise.
func (o *TransferPolicy) GetMaxUplinkBandwidth() int32 {
	if o == nil || IsNil(o.MaxUplinkBandwidth) {
		var ret int32
		return ret
	}
	return *o.MaxUplinkBandwidth
}

// GetMaxUplinkBandwidthOk returns a tuple with the MaxUplinkBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferPolicy) GetMaxUplinkBandwidthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUplinkBandwidth) {
		return nil, false
	}
	return o.MaxUplinkBandwidth, true
}

// HasMaxUplinkBandwidth returns a boolean if a field has been set.
func (o *TransferPolicy) HasMaxUplinkBandwidth() bool {
	if o != nil && !IsNil(o.MaxUplinkBandwidth) {
		return true
	}

	return false
}

// SetMaxUplinkBandwidth gets a reference to the given int32 and assigns it to the MaxUplinkBandwidth field.
func (o *TransferPolicy) SetMaxUplinkBandwidth(v int32) {
	o.MaxUplinkBandwidth = &v
}

// GetMaxDownlinkBandwidth returns the MaxDownlinkBandwidth field value if set, zero value otherwise.
func (o *TransferPolicy) GetMaxDownlinkBandwidth() int32 {
	if o == nil || IsNil(o.MaxDownlinkBandwidth) {
		var ret int32
		return ret
	}
	return *o.MaxDownlinkBandwidth
}

// GetMaxDownlinkBandwidthOk returns a tuple with the MaxDownlinkBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferPolicy) GetMaxDownlinkBandwidthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDownlinkBandwidth) {
		return nil, false
	}
	return o.MaxDownlinkBandwidth, true
}

// HasMaxDownlinkBandwidth returns a boolean if a field has been set.
func (o *TransferPolicy) HasMaxDownlinkBandwidth() bool {
	if o != nil && !IsNil(o.MaxDownlinkBandwidth) {
		return true
	}

	return false
}

// SetMaxDownlinkBandwidth gets a reference to the given int32 and assigns it to the MaxDownlinkBandwidth field.
func (o *TransferPolicy) SetMaxDownlinkBandwidth(v int32) {
	o.MaxDownlinkBandwidth = &v
}

// GetRatingGroup returns the RatingGroup field value
func (o *TransferPolicy) GetRatingGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RatingGroup
}

// GetRatingGroupOk returns a tuple with the RatingGroup field value
// and a boolean to check if the value has been set.
func (o *TransferPolicy) GetRatingGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RatingGroup, true
}

// SetRatingGroup sets field value
func (o *TransferPolicy) SetRatingGroup(v int32) {
	o.RatingGroup = v
}

// GetTimeWindow returns the TimeWindow field value
func (o *TransferPolicy) GetTimeWindow() TimeWindow {
	if o == nil {
		var ret TimeWindow
		return ret
	}

	return o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value
// and a boolean to check if the value has been set.
func (o *TransferPolicy) GetTimeWindowOk() (*TimeWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeWindow, true
}

// SetTimeWindow sets field value
func (o *TransferPolicy) SetTimeWindow(v TimeWindow) {
	o.TimeWindow = v
}

func (o TransferPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bdtPolicyId"] = o.BdtPolicyId
	if !IsNil(o.MaxUplinkBandwidth) {
		toSerialize["maxUplinkBandwidth"] = o.MaxUplinkBandwidth
	}
	if !IsNil(o.MaxDownlinkBandwidth) {
		toSerialize["maxDownlinkBandwidth"] = o.MaxDownlinkBandwidth
	}
	toSerialize["ratingGroup"] = o.RatingGroup
	toSerialize["timeWindow"] = o.TimeWindow
	return toSerialize, nil
}

type NullableTransferPolicy struct {
	value *TransferPolicy
	isSet bool
}

func (v NullableTransferPolicy) Get() *TransferPolicy {
	return v.value
}

func (v *NullableTransferPolicy) Set(val *TransferPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferPolicy(val *TransferPolicy) *NullableTransferPolicy {
	return &NullableTransferPolicy{value: val, isSet: true}
}

func (v NullableTransferPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

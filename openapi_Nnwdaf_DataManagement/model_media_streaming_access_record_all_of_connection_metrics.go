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

// checks if the MediaStreamingAccessRecordAllOfConnectionMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaStreamingAccessRecordAllOfConnectionMetrics{}

// MediaStreamingAccessRecordAllOfConnectionMetrics struct for MediaStreamingAccessRecordAllOfConnectionMetrics
type MediaStreamingAccessRecordAllOfConnectionMetrics struct {
	// string with format 'float' as defined in OpenAPI.
	MeanNetworkRoundTripTime float32 `json:"meanNetworkRoundTripTime"`
	// string with format 'float' as defined in OpenAPI.
	NetworkRoundTripTimeVariation float32 `json:"networkRoundTripTimeVariation"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	CongestionWindowSize int32 `json:"congestionWindowSize"`
}

// NewMediaStreamingAccessRecordAllOfConnectionMetrics instantiates a new MediaStreamingAccessRecordAllOfConnectionMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaStreamingAccessRecordAllOfConnectionMetrics(meanNetworkRoundTripTime float32, networkRoundTripTimeVariation float32, congestionWindowSize int32) *MediaStreamingAccessRecordAllOfConnectionMetrics {
	this := MediaStreamingAccessRecordAllOfConnectionMetrics{}
	this.MeanNetworkRoundTripTime = meanNetworkRoundTripTime
	this.NetworkRoundTripTimeVariation = networkRoundTripTimeVariation
	this.CongestionWindowSize = congestionWindowSize
	return &this
}

// NewMediaStreamingAccessRecordAllOfConnectionMetricsWithDefaults instantiates a new MediaStreamingAccessRecordAllOfConnectionMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaStreamingAccessRecordAllOfConnectionMetricsWithDefaults() *MediaStreamingAccessRecordAllOfConnectionMetrics {
	this := MediaStreamingAccessRecordAllOfConnectionMetrics{}
	return &this
}

// GetMeanNetworkRoundTripTime returns the MeanNetworkRoundTripTime field value
func (o *MediaStreamingAccessRecordAllOfConnectionMetrics) GetMeanNetworkRoundTripTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MeanNetworkRoundTripTime
}

// GetMeanNetworkRoundTripTimeOk returns a tuple with the MeanNetworkRoundTripTime field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfConnectionMetrics) GetMeanNetworkRoundTripTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeanNetworkRoundTripTime, true
}

// SetMeanNetworkRoundTripTime sets field value
func (o *MediaStreamingAccessRecordAllOfConnectionMetrics) SetMeanNetworkRoundTripTime(v float32) {
	o.MeanNetworkRoundTripTime = v
}

// GetNetworkRoundTripTimeVariation returns the NetworkRoundTripTimeVariation field value
func (o *MediaStreamingAccessRecordAllOfConnectionMetrics) GetNetworkRoundTripTimeVariation() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NetworkRoundTripTimeVariation
}

// GetNetworkRoundTripTimeVariationOk returns a tuple with the NetworkRoundTripTimeVariation field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfConnectionMetrics) GetNetworkRoundTripTimeVariationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkRoundTripTimeVariation, true
}

// SetNetworkRoundTripTimeVariation sets field value
func (o *MediaStreamingAccessRecordAllOfConnectionMetrics) SetNetworkRoundTripTimeVariation(v float32) {
	o.NetworkRoundTripTimeVariation = v
}

// GetCongestionWindowSize returns the CongestionWindowSize field value
func (o *MediaStreamingAccessRecordAllOfConnectionMetrics) GetCongestionWindowSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CongestionWindowSize
}

// GetCongestionWindowSizeOk returns a tuple with the CongestionWindowSize field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfConnectionMetrics) GetCongestionWindowSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CongestionWindowSize, true
}

// SetCongestionWindowSize sets field value
func (o *MediaStreamingAccessRecordAllOfConnectionMetrics) SetCongestionWindowSize(v int32) {
	o.CongestionWindowSize = v
}

func (o MediaStreamingAccessRecordAllOfConnectionMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaStreamingAccessRecordAllOfConnectionMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meanNetworkRoundTripTime"] = o.MeanNetworkRoundTripTime
	toSerialize["networkRoundTripTimeVariation"] = o.NetworkRoundTripTimeVariation
	toSerialize["congestionWindowSize"] = o.CongestionWindowSize
	return toSerialize, nil
}

type NullableMediaStreamingAccessRecordAllOfConnectionMetrics struct {
	value *MediaStreamingAccessRecordAllOfConnectionMetrics
	isSet bool
}

func (v NullableMediaStreamingAccessRecordAllOfConnectionMetrics) Get() *MediaStreamingAccessRecordAllOfConnectionMetrics {
	return v.value
}

func (v *NullableMediaStreamingAccessRecordAllOfConnectionMetrics) Set(val *MediaStreamingAccessRecordAllOfConnectionMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaStreamingAccessRecordAllOfConnectionMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaStreamingAccessRecordAllOfConnectionMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaStreamingAccessRecordAllOfConnectionMetrics(val *MediaStreamingAccessRecordAllOfConnectionMetrics) *NullableMediaStreamingAccessRecordAllOfConnectionMetrics {
	return &NullableMediaStreamingAccessRecordAllOfConnectionMetrics{value: val, isSet: true}
}

func (v NullableMediaStreamingAccessRecordAllOfConnectionMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaStreamingAccessRecordAllOfConnectionMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

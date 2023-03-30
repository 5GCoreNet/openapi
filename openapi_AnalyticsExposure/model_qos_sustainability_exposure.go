/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
	"time"
)

// checks if the QosSustainabilityExposure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosSustainabilityExposure{}

// QosSustainabilityExposure Represents a QoS sustainability information.
type QosSustainabilityExposure struct {
	LocArea LocationArea5G `json:"locArea"`
	// string with format \"date-time\" as defined in OpenAPI.
	StartTs time.Time `json:"startTs"`
	// string with format \"date-time\" as defined in OpenAPI.
	EndTs time.Time `json:"endTs"`
	QosFlowRetThd *RetainabilityThreshold `json:"qosFlowRetThd,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	RanUeThrouThd *string `json:"ranUeThrouThd,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
}

// NewQosSustainabilityExposure instantiates a new QosSustainabilityExposure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosSustainabilityExposure(locArea LocationArea5G, startTs time.Time, endTs time.Time) *QosSustainabilityExposure {
	this := QosSustainabilityExposure{}
	this.LocArea = locArea
	this.StartTs = startTs
	this.EndTs = endTs
	return &this
}

// NewQosSustainabilityExposureWithDefaults instantiates a new QosSustainabilityExposure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosSustainabilityExposureWithDefaults() *QosSustainabilityExposure {
	this := QosSustainabilityExposure{}
	return &this
}

// GetLocArea returns the LocArea field value
func (o *QosSustainabilityExposure) GetLocArea() LocationArea5G {
	if o == nil {
		var ret LocationArea5G
		return ret
	}

	return o.LocArea
}

// GetLocAreaOk returns a tuple with the LocArea field value
// and a boolean to check if the value has been set.
func (o *QosSustainabilityExposure) GetLocAreaOk() (*LocationArea5G, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocArea, true
}

// SetLocArea sets field value
func (o *QosSustainabilityExposure) SetLocArea(v LocationArea5G) {
	o.LocArea = v
}

// GetStartTs returns the StartTs field value
func (o *QosSustainabilityExposure) GetStartTs() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value
// and a boolean to check if the value has been set.
func (o *QosSustainabilityExposure) GetStartTsOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTs, true
}

// SetStartTs sets field value
func (o *QosSustainabilityExposure) SetStartTs(v time.Time) {
	o.StartTs = v
}

// GetEndTs returns the EndTs field value
func (o *QosSustainabilityExposure) GetEndTs() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value
// and a boolean to check if the value has been set.
func (o *QosSustainabilityExposure) GetEndTsOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTs, true
}

// SetEndTs sets field value
func (o *QosSustainabilityExposure) SetEndTs(v time.Time) {
	o.EndTs = v
}

// GetQosFlowRetThd returns the QosFlowRetThd field value if set, zero value otherwise.
func (o *QosSustainabilityExposure) GetQosFlowRetThd() RetainabilityThreshold {
	if o == nil || IsNil(o.QosFlowRetThd) {
		var ret RetainabilityThreshold
		return ret
	}
	return *o.QosFlowRetThd
}

// GetQosFlowRetThdOk returns a tuple with the QosFlowRetThd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosSustainabilityExposure) GetQosFlowRetThdOk() (*RetainabilityThreshold, bool) {
	if o == nil || IsNil(o.QosFlowRetThd) {
		return nil, false
	}
	return o.QosFlowRetThd, true
}

// HasQosFlowRetThd returns a boolean if a field has been set.
func (o *QosSustainabilityExposure) HasQosFlowRetThd() bool {
	if o != nil && !IsNil(o.QosFlowRetThd) {
		return true
	}

	return false
}

// SetQosFlowRetThd gets a reference to the given RetainabilityThreshold and assigns it to the QosFlowRetThd field.
func (o *QosSustainabilityExposure) SetQosFlowRetThd(v RetainabilityThreshold) {
	o.QosFlowRetThd = &v
}

// GetRanUeThrouThd returns the RanUeThrouThd field value if set, zero value otherwise.
func (o *QosSustainabilityExposure) GetRanUeThrouThd() string {
	if o == nil || IsNil(o.RanUeThrouThd) {
		var ret string
		return ret
	}
	return *o.RanUeThrouThd
}

// GetRanUeThrouThdOk returns a tuple with the RanUeThrouThd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosSustainabilityExposure) GetRanUeThrouThdOk() (*string, bool) {
	if o == nil || IsNil(o.RanUeThrouThd) {
		return nil, false
	}
	return o.RanUeThrouThd, true
}

// HasRanUeThrouThd returns a boolean if a field has been set.
func (o *QosSustainabilityExposure) HasRanUeThrouThd() bool {
	if o != nil && !IsNil(o.RanUeThrouThd) {
		return true
	}

	return false
}

// SetRanUeThrouThd gets a reference to the given string and assigns it to the RanUeThrouThd field.
func (o *QosSustainabilityExposure) SetRanUeThrouThd(v string) {
	o.RanUeThrouThd = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *QosSustainabilityExposure) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosSustainabilityExposure) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *QosSustainabilityExposure) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *QosSustainabilityExposure) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *QosSustainabilityExposure) GetConfidence() int32 {
	if o == nil || IsNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosSustainabilityExposure) GetConfidenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *QosSustainabilityExposure) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *QosSustainabilityExposure) SetConfidence(v int32) {
	o.Confidence = &v
}

func (o QosSustainabilityExposure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosSustainabilityExposure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locArea"] = o.LocArea
	toSerialize["startTs"] = o.StartTs
	toSerialize["endTs"] = o.EndTs
	if !IsNil(o.QosFlowRetThd) {
		toSerialize["qosFlowRetThd"] = o.QosFlowRetThd
	}
	if !IsNil(o.RanUeThrouThd) {
		toSerialize["ranUeThrouThd"] = o.RanUeThrouThd
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	return toSerialize, nil
}

type NullableQosSustainabilityExposure struct {
	value *QosSustainabilityExposure
	isSet bool
}

func (v NullableQosSustainabilityExposure) Get() *QosSustainabilityExposure {
	return v.value
}

func (v *NullableQosSustainabilityExposure) Set(val *QosSustainabilityExposure) {
	v.value = val
	v.isSet = true
}

func (v NullableQosSustainabilityExposure) IsSet() bool {
	return v.isSet
}

func (v *NullableQosSustainabilityExposure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosSustainabilityExposure(val *QosSustainabilityExposure) *NullableQosSustainabilityExposure {
	return &NullableQosSustainabilityExposure{value: val, isSet: true}
}

func (v NullableQosSustainabilityExposure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosSustainabilityExposure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the AnalyticsMetadataInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsMetadataInfo{}

// AnalyticsMetadataInfo Contains analytics metadata information required for analytics aggregation.
type AnalyticsMetadataInfo struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NumSamples    *int32                       `json:"numSamples,omitempty"`
	DataWindow    *TimeWindow                  `json:"dataWindow,omitempty"`
	DataStatProps []DatasetStatisticalProperty `json:"dataStatProps,omitempty"`
	Strategy      *OutputStrategy              `json:"strategy,omitempty"`
	Accuracy      *Accuracy                    `json:"accuracy,omitempty"`
}

// NewAnalyticsMetadataInfo instantiates a new AnalyticsMetadataInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsMetadataInfo() *AnalyticsMetadataInfo {
	this := AnalyticsMetadataInfo{}
	return &this
}

// NewAnalyticsMetadataInfoWithDefaults instantiates a new AnalyticsMetadataInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsMetadataInfoWithDefaults() *AnalyticsMetadataInfo {
	this := AnalyticsMetadataInfo{}
	return &this
}

// GetNumSamples returns the NumSamples field value if set, zero value otherwise.
func (o *AnalyticsMetadataInfo) GetNumSamples() int32 {
	if o == nil || IsNil(o.NumSamples) {
		var ret int32
		return ret
	}
	return *o.NumSamples
}

// GetNumSamplesOk returns a tuple with the NumSamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsMetadataInfo) GetNumSamplesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumSamples) {
		return nil, false
	}
	return o.NumSamples, true
}

// HasNumSamples returns a boolean if a field has been set.
func (o *AnalyticsMetadataInfo) HasNumSamples() bool {
	if o != nil && !IsNil(o.NumSamples) {
		return true
	}

	return false
}

// SetNumSamples gets a reference to the given int32 and assigns it to the NumSamples field.
func (o *AnalyticsMetadataInfo) SetNumSamples(v int32) {
	o.NumSamples = &v
}

// GetDataWindow returns the DataWindow field value if set, zero value otherwise.
func (o *AnalyticsMetadataInfo) GetDataWindow() TimeWindow {
	if o == nil || IsNil(o.DataWindow) {
		var ret TimeWindow
		return ret
	}
	return *o.DataWindow
}

// GetDataWindowOk returns a tuple with the DataWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsMetadataInfo) GetDataWindowOk() (*TimeWindow, bool) {
	if o == nil || IsNil(o.DataWindow) {
		return nil, false
	}
	return o.DataWindow, true
}

// HasDataWindow returns a boolean if a field has been set.
func (o *AnalyticsMetadataInfo) HasDataWindow() bool {
	if o != nil && !IsNil(o.DataWindow) {
		return true
	}

	return false
}

// SetDataWindow gets a reference to the given TimeWindow and assigns it to the DataWindow field.
func (o *AnalyticsMetadataInfo) SetDataWindow(v TimeWindow) {
	o.DataWindow = &v
}

// GetDataStatProps returns the DataStatProps field value if set, zero value otherwise.
func (o *AnalyticsMetadataInfo) GetDataStatProps() []DatasetStatisticalProperty {
	if o == nil || IsNil(o.DataStatProps) {
		var ret []DatasetStatisticalProperty
		return ret
	}
	return o.DataStatProps
}

// GetDataStatPropsOk returns a tuple with the DataStatProps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsMetadataInfo) GetDataStatPropsOk() ([]DatasetStatisticalProperty, bool) {
	if o == nil || IsNil(o.DataStatProps) {
		return nil, false
	}
	return o.DataStatProps, true
}

// HasDataStatProps returns a boolean if a field has been set.
func (o *AnalyticsMetadataInfo) HasDataStatProps() bool {
	if o != nil && !IsNil(o.DataStatProps) {
		return true
	}

	return false
}

// SetDataStatProps gets a reference to the given []DatasetStatisticalProperty and assigns it to the DataStatProps field.
func (o *AnalyticsMetadataInfo) SetDataStatProps(v []DatasetStatisticalProperty) {
	o.DataStatProps = v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *AnalyticsMetadataInfo) GetStrategy() OutputStrategy {
	if o == nil || IsNil(o.Strategy) {
		var ret OutputStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsMetadataInfo) GetStrategyOk() (*OutputStrategy, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *AnalyticsMetadataInfo) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given OutputStrategy and assigns it to the Strategy field.
func (o *AnalyticsMetadataInfo) SetStrategy(v OutputStrategy) {
	o.Strategy = &v
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *AnalyticsMetadataInfo) GetAccuracy() Accuracy {
	if o == nil || IsNil(o.Accuracy) {
		var ret Accuracy
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsMetadataInfo) GetAccuracyOk() (*Accuracy, bool) {
	if o == nil || IsNil(o.Accuracy) {
		return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *AnalyticsMetadataInfo) HasAccuracy() bool {
	if o != nil && !IsNil(o.Accuracy) {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given Accuracy and assigns it to the Accuracy field.
func (o *AnalyticsMetadataInfo) SetAccuracy(v Accuracy) {
	o.Accuracy = &v
}

func (o AnalyticsMetadataInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsMetadataInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumSamples) {
		toSerialize["numSamples"] = o.NumSamples
	}
	if !IsNil(o.DataWindow) {
		toSerialize["dataWindow"] = o.DataWindow
	}
	if !IsNil(o.DataStatProps) {
		toSerialize["dataStatProps"] = o.DataStatProps
	}
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !IsNil(o.Accuracy) {
		toSerialize["accuracy"] = o.Accuracy
	}
	return toSerialize, nil
}

type NullableAnalyticsMetadataInfo struct {
	value *AnalyticsMetadataInfo
	isSet bool
}

func (v NullableAnalyticsMetadataInfo) Get() *AnalyticsMetadataInfo {
	return v.value
}

func (v *NullableAnalyticsMetadataInfo) Set(val *AnalyticsMetadataInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsMetadataInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsMetadataInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsMetadataInfo(val *AnalyticsMetadataInfo) *NullableAnalyticsMetadataInfo {
	return &NullableAnalyticsMetadataInfo{value: val, isSet: true}
}

func (v NullableAnalyticsMetadataInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsMetadataInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

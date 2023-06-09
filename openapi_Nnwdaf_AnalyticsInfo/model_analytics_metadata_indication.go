/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// checks if the AnalyticsMetadataIndication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsMetadataIndication{}

// AnalyticsMetadataIndication Contains analytics metadata information requested to be used during analytics generation.
type AnalyticsMetadataIndication struct {
	DataWindow    *TimeWindow                  `json:"dataWindow,omitempty"`
	DataStatProps []DatasetStatisticalProperty `json:"dataStatProps,omitempty"`
	Strategy      *OutputStrategy              `json:"strategy,omitempty"`
	AggrNwdafIds  []string                     `json:"aggrNwdafIds,omitempty"`
}

// NewAnalyticsMetadataIndication instantiates a new AnalyticsMetadataIndication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsMetadataIndication() *AnalyticsMetadataIndication {
	this := AnalyticsMetadataIndication{}
	return &this
}

// NewAnalyticsMetadataIndicationWithDefaults instantiates a new AnalyticsMetadataIndication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsMetadataIndicationWithDefaults() *AnalyticsMetadataIndication {
	this := AnalyticsMetadataIndication{}
	return &this
}

// GetDataWindow returns the DataWindow field value if set, zero value otherwise.
func (o *AnalyticsMetadataIndication) GetDataWindow() TimeWindow {
	if o == nil || IsNil(o.DataWindow) {
		var ret TimeWindow
		return ret
	}
	return *o.DataWindow
}

// GetDataWindowOk returns a tuple with the DataWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsMetadataIndication) GetDataWindowOk() (*TimeWindow, bool) {
	if o == nil || IsNil(o.DataWindow) {
		return nil, false
	}
	return o.DataWindow, true
}

// HasDataWindow returns a boolean if a field has been set.
func (o *AnalyticsMetadataIndication) HasDataWindow() bool {
	if o != nil && !IsNil(o.DataWindow) {
		return true
	}

	return false
}

// SetDataWindow gets a reference to the given TimeWindow and assigns it to the DataWindow field.
func (o *AnalyticsMetadataIndication) SetDataWindow(v TimeWindow) {
	o.DataWindow = &v
}

// GetDataStatProps returns the DataStatProps field value if set, zero value otherwise.
func (o *AnalyticsMetadataIndication) GetDataStatProps() []DatasetStatisticalProperty {
	if o == nil || IsNil(o.DataStatProps) {
		var ret []DatasetStatisticalProperty
		return ret
	}
	return o.DataStatProps
}

// GetDataStatPropsOk returns a tuple with the DataStatProps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsMetadataIndication) GetDataStatPropsOk() ([]DatasetStatisticalProperty, bool) {
	if o == nil || IsNil(o.DataStatProps) {
		return nil, false
	}
	return o.DataStatProps, true
}

// HasDataStatProps returns a boolean if a field has been set.
func (o *AnalyticsMetadataIndication) HasDataStatProps() bool {
	if o != nil && !IsNil(o.DataStatProps) {
		return true
	}

	return false
}

// SetDataStatProps gets a reference to the given []DatasetStatisticalProperty and assigns it to the DataStatProps field.
func (o *AnalyticsMetadataIndication) SetDataStatProps(v []DatasetStatisticalProperty) {
	o.DataStatProps = v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *AnalyticsMetadataIndication) GetStrategy() OutputStrategy {
	if o == nil || IsNil(o.Strategy) {
		var ret OutputStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsMetadataIndication) GetStrategyOk() (*OutputStrategy, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *AnalyticsMetadataIndication) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given OutputStrategy and assigns it to the Strategy field.
func (o *AnalyticsMetadataIndication) SetStrategy(v OutputStrategy) {
	o.Strategy = &v
}

// GetAggrNwdafIds returns the AggrNwdafIds field value if set, zero value otherwise.
func (o *AnalyticsMetadataIndication) GetAggrNwdafIds() []string {
	if o == nil || IsNil(o.AggrNwdafIds) {
		var ret []string
		return ret
	}
	return o.AggrNwdafIds
}

// GetAggrNwdafIdsOk returns a tuple with the AggrNwdafIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsMetadataIndication) GetAggrNwdafIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AggrNwdafIds) {
		return nil, false
	}
	return o.AggrNwdafIds, true
}

// HasAggrNwdafIds returns a boolean if a field has been set.
func (o *AnalyticsMetadataIndication) HasAggrNwdafIds() bool {
	if o != nil && !IsNil(o.AggrNwdafIds) {
		return true
	}

	return false
}

// SetAggrNwdafIds gets a reference to the given []string and assigns it to the AggrNwdafIds field.
func (o *AnalyticsMetadataIndication) SetAggrNwdafIds(v []string) {
	o.AggrNwdafIds = v
}

func (o AnalyticsMetadataIndication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsMetadataIndication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataWindow) {
		toSerialize["dataWindow"] = o.DataWindow
	}
	if !IsNil(o.DataStatProps) {
		toSerialize["dataStatProps"] = o.DataStatProps
	}
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !IsNil(o.AggrNwdafIds) {
		toSerialize["aggrNwdafIds"] = o.AggrNwdafIds
	}
	return toSerialize, nil
}

type NullableAnalyticsMetadataIndication struct {
	value *AnalyticsMetadataIndication
	isSet bool
}

func (v NullableAnalyticsMetadataIndication) Get() *AnalyticsMetadataIndication {
	return v.value
}

func (v *NullableAnalyticsMetadataIndication) Set(val *AnalyticsMetadataIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsMetadataIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsMetadataIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsMetadataIndication(val *AnalyticsMetadataIndication) *NullableAnalyticsMetadataIndication {
	return &NullableAnalyticsMetadataIndication{value: val, isSet: true}
}

func (v NullableAnalyticsMetadataIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsMetadataIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

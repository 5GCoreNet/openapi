/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the SubNetworkSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubNetworkSingleAllOfAttributes{}

// SubNetworkSingleAllOfAttributes struct for SubNetworkSingleAllOfAttributes
type SubNetworkSingleAllOfAttributes struct {
	DnPrefix *string `json:"dnPrefix,omitempty"`
	UserLabel *string `json:"userLabel,omitempty"`
	UserDefinedNetworkType *string `json:"userDefinedNetworkType,omitempty"`
	SetOfMcc []string `json:"setOfMcc,omitempty"`
	PriorityLabel *int32 `json:"priorityLabel,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
}

// NewSubNetworkSingleAllOfAttributes instantiates a new SubNetworkSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingleAllOfAttributes() *SubNetworkSingleAllOfAttributes {
	this := SubNetworkSingleAllOfAttributes{}
	return &this
}

// NewSubNetworkSingleAllOfAttributesWithDefaults instantiates a new SubNetworkSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingleAllOfAttributesWithDefaults() *SubNetworkSingleAllOfAttributes {
	this := SubNetworkSingleAllOfAttributes{}
	return &this
}

// GetDnPrefix returns the DnPrefix field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOfAttributes) GetDnPrefix() string {
	if o == nil || IsNil(o.DnPrefix) {
		var ret string
		return ret
	}
	return *o.DnPrefix
}

// GetDnPrefixOk returns a tuple with the DnPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOfAttributes) GetDnPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.DnPrefix) {
		return nil, false
	}
	return o.DnPrefix, true
}

// HasDnPrefix returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOfAttributes) HasDnPrefix() bool {
	if o != nil && !IsNil(o.DnPrefix) {
		return true
	}

	return false
}

// SetDnPrefix gets a reference to the given string and assigns it to the DnPrefix field.
func (o *SubNetworkSingleAllOfAttributes) SetDnPrefix(v string) {
	o.DnPrefix = &v
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOfAttributes) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOfAttributes) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOfAttributes) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *SubNetworkSingleAllOfAttributes) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetUserDefinedNetworkType returns the UserDefinedNetworkType field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOfAttributes) GetUserDefinedNetworkType() string {
	if o == nil || IsNil(o.UserDefinedNetworkType) {
		var ret string
		return ret
	}
	return *o.UserDefinedNetworkType
}

// GetUserDefinedNetworkTypeOk returns a tuple with the UserDefinedNetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOfAttributes) GetUserDefinedNetworkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserDefinedNetworkType) {
		return nil, false
	}
	return o.UserDefinedNetworkType, true
}

// HasUserDefinedNetworkType returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOfAttributes) HasUserDefinedNetworkType() bool {
	if o != nil && !IsNil(o.UserDefinedNetworkType) {
		return true
	}

	return false
}

// SetUserDefinedNetworkType gets a reference to the given string and assigns it to the UserDefinedNetworkType field.
func (o *SubNetworkSingleAllOfAttributes) SetUserDefinedNetworkType(v string) {
	o.UserDefinedNetworkType = &v
}

// GetSetOfMcc returns the SetOfMcc field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOfAttributes) GetSetOfMcc() []string {
	if o == nil || IsNil(o.SetOfMcc) {
		var ret []string
		return ret
	}
	return o.SetOfMcc
}

// GetSetOfMccOk returns a tuple with the SetOfMcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOfAttributes) GetSetOfMccOk() ([]string, bool) {
	if o == nil || IsNil(o.SetOfMcc) {
		return nil, false
	}
	return o.SetOfMcc, true
}

// HasSetOfMcc returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOfAttributes) HasSetOfMcc() bool {
	if o != nil && !IsNil(o.SetOfMcc) {
		return true
	}

	return false
}

// SetSetOfMcc gets a reference to the given []string and assigns it to the SetOfMcc field.
func (o *SubNetworkSingleAllOfAttributes) SetSetOfMcc(v []string) {
	o.SetOfMcc = v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOfAttributes) GetPriorityLabel() int32 {
	if o == nil || IsNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOfAttributes) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLabel) {
		return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOfAttributes) HasPriorityLabel() bool {
	if o != nil && !IsNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *SubNetworkSingleAllOfAttributes) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOfAttributes) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOfAttributes) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOfAttributes) HasSupportedPerfMetricGroups() bool {
	if o != nil && !IsNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *SubNetworkSingleAllOfAttributes) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetSupportedTraceMetrics returns the SupportedTraceMetrics field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOfAttributes) GetSupportedTraceMetrics() []string {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedTraceMetrics
}

// GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOfAttributes) GetSupportedTraceMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		return nil, false
	}
	return o.SupportedTraceMetrics, true
}

// HasSupportedTraceMetrics returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOfAttributes) HasSupportedTraceMetrics() bool {
	if o != nil && !IsNil(o.SupportedTraceMetrics) {
		return true
	}

	return false
}

// SetSupportedTraceMetrics gets a reference to the given []string and assigns it to the SupportedTraceMetrics field.
func (o *SubNetworkSingleAllOfAttributes) SetSupportedTraceMetrics(v []string) {
	o.SupportedTraceMetrics = v
}

func (o SubNetworkSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubNetworkSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DnPrefix) {
		toSerialize["dnPrefix"] = o.DnPrefix
	}
	if !IsNil(o.UserLabel) {
		toSerialize["userLabel"] = o.UserLabel
	}
	if !IsNil(o.UserDefinedNetworkType) {
		toSerialize["userDefinedNetworkType"] = o.UserDefinedNetworkType
	}
	if !IsNil(o.SetOfMcc) {
		toSerialize["setOfMcc"] = o.SetOfMcc
	}
	if !IsNil(o.PriorityLabel) {
		toSerialize["priorityLabel"] = o.PriorityLabel
	}
	if !IsNil(o.SupportedPerfMetricGroups) {
		toSerialize["supportedPerfMetricGroups"] = o.SupportedPerfMetricGroups
	}
	if !IsNil(o.SupportedTraceMetrics) {
		toSerialize["supportedTraceMetrics"] = o.SupportedTraceMetrics
	}
	return toSerialize, nil
}

type NullableSubNetworkSingleAllOfAttributes struct {
	value *SubNetworkSingleAllOfAttributes
	isSet bool
}

func (v NullableSubNetworkSingleAllOfAttributes) Get() *SubNetworkSingleAllOfAttributes {
	return v.value
}

func (v *NullableSubNetworkSingleAllOfAttributes) Set(val *SubNetworkSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingleAllOfAttributes(val *SubNetworkSingleAllOfAttributes) *NullableSubNetworkSingleAllOfAttributes {
	return &NullableSubNetworkSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableSubNetworkSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



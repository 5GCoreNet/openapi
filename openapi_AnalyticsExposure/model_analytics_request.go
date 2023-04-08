/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
)

// checks if the AnalyticsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsRequest{}

// AnalyticsRequest Represents the parameters to request to retrieve analytics information.
type AnalyticsRequest struct {
	AnalyEvent AnalyticsEvent `json:"analyEvent"`
	AnalyEventFilter *AnalyticsEventFilter `json:"analyEventFilter,omitempty"`
	AnalyRep *EventReportingRequirement `json:"analyRep,omitempty"`
	TgtUe *TargetUeId `json:"tgtUe,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat"`
}

// NewAnalyticsRequest instantiates a new AnalyticsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsRequest(analyEvent AnalyticsEvent, suppFeat string) *AnalyticsRequest {
	this := AnalyticsRequest{}
	this.AnalyEvent = analyEvent
	this.SuppFeat = suppFeat
	return &this
}

// NewAnalyticsRequestWithDefaults instantiates a new AnalyticsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsRequestWithDefaults() *AnalyticsRequest {
	this := AnalyticsRequest{}
	return &this
}

// GetAnalyEvent returns the AnalyEvent field value
func (o *AnalyticsRequest) GetAnalyEvent() AnalyticsEvent {
	if o == nil {
		var ret AnalyticsEvent
		return ret
	}

	return o.AnalyEvent
}

// GetAnalyEventOk returns a tuple with the AnalyEvent field value
// and a boolean to check if the value has been set.
func (o *AnalyticsRequest) GetAnalyEventOk() (*AnalyticsEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalyEvent, true
}

// SetAnalyEvent sets field value
func (o *AnalyticsRequest) SetAnalyEvent(v AnalyticsEvent) {
	o.AnalyEvent = v
}

// GetAnalyEventFilter returns the AnalyEventFilter field value if set, zero value otherwise.
func (o *AnalyticsRequest) GetAnalyEventFilter() AnalyticsEventFilter {
	if o == nil || isNil(o.AnalyEventFilter) {
		var ret AnalyticsEventFilter
		return ret
	}
	return *o.AnalyEventFilter
}

// GetAnalyEventFilterOk returns a tuple with the AnalyEventFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsRequest) GetAnalyEventFilterOk() (*AnalyticsEventFilter, bool) {
	if o == nil || isNil(o.AnalyEventFilter) {
		return nil, false
	}
	return o.AnalyEventFilter, true
}

// HasAnalyEventFilter returns a boolean if a field has been set.
func (o *AnalyticsRequest) HasAnalyEventFilter() bool {
	if o != nil && !isNil(o.AnalyEventFilter) {
		return true
	}

	return false
}

// SetAnalyEventFilter gets a reference to the given AnalyticsEventFilter and assigns it to the AnalyEventFilter field.
func (o *AnalyticsRequest) SetAnalyEventFilter(v AnalyticsEventFilter) {
	o.AnalyEventFilter = &v
}

// GetAnalyRep returns the AnalyRep field value if set, zero value otherwise.
func (o *AnalyticsRequest) GetAnalyRep() EventReportingRequirement {
	if o == nil || isNil(o.AnalyRep) {
		var ret EventReportingRequirement
		return ret
	}
	return *o.AnalyRep
}

// GetAnalyRepOk returns a tuple with the AnalyRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsRequest) GetAnalyRepOk() (*EventReportingRequirement, bool) {
	if o == nil || isNil(o.AnalyRep) {
		return nil, false
	}
	return o.AnalyRep, true
}

// HasAnalyRep returns a boolean if a field has been set.
func (o *AnalyticsRequest) HasAnalyRep() bool {
	if o != nil && !isNil(o.AnalyRep) {
		return true
	}

	return false
}

// SetAnalyRep gets a reference to the given EventReportingRequirement and assigns it to the AnalyRep field.
func (o *AnalyticsRequest) SetAnalyRep(v EventReportingRequirement) {
	o.AnalyRep = &v
}

// GetTgtUe returns the TgtUe field value if set, zero value otherwise.
func (o *AnalyticsRequest) GetTgtUe() TargetUeId {
	if o == nil || isNil(o.TgtUe) {
		var ret TargetUeId
		return ret
	}
	return *o.TgtUe
}

// GetTgtUeOk returns a tuple with the TgtUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsRequest) GetTgtUeOk() (*TargetUeId, bool) {
	if o == nil || isNil(o.TgtUe) {
		return nil, false
	}
	return o.TgtUe, true
}

// HasTgtUe returns a boolean if a field has been set.
func (o *AnalyticsRequest) HasTgtUe() bool {
	if o != nil && !isNil(o.TgtUe) {
		return true
	}

	return false
}

// SetTgtUe gets a reference to the given TargetUeId and assigns it to the TgtUe field.
func (o *AnalyticsRequest) SetTgtUe(v TargetUeId) {
	o.TgtUe = &v
}

// GetSuppFeat returns the SuppFeat field value
func (o *AnalyticsRequest) GetSuppFeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value
// and a boolean to check if the value has been set.
func (o *AnalyticsRequest) GetSuppFeatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuppFeat, true
}

// SetSuppFeat sets field value
func (o *AnalyticsRequest) SetSuppFeat(v string) {
	o.SuppFeat = v
}

func (o AnalyticsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["analyEvent"] = o.AnalyEvent
	if !isNil(o.AnalyEventFilter) {
		toSerialize["analyEventFilter"] = o.AnalyEventFilter
	}
	if !isNil(o.AnalyRep) {
		toSerialize["analyRep"] = o.AnalyRep
	}
	if !isNil(o.TgtUe) {
		toSerialize["tgtUe"] = o.TgtUe
	}
	toSerialize["suppFeat"] = o.SuppFeat
	return toSerialize, nil
}

type NullableAnalyticsRequest struct {
	value *AnalyticsRequest
	isSet bool
}

func (v NullableAnalyticsRequest) Get() *AnalyticsRequest {
	return v.value
}

func (v *NullableAnalyticsRequest) Set(val *AnalyticsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsRequest(val *AnalyticsRequest) *NullableAnalyticsRequest {
	return &NullableAnalyticsRequest{value: val, isSet: true}
}

func (v NullableAnalyticsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



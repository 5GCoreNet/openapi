/*
Namf_EventExposure

AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_EventExposure

import (
	"encoding/json"
)

// checks if the AmfCreatedEventSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfCreatedEventSubscription{}

// AmfCreatedEventSubscription Data within a create AMF event subscription response
type AmfCreatedEventSubscription struct {
	Subscription AmfEventSubscription `json:"subscription"`
	// String providing an URI formatted according to RFC 3986.
	SubscriptionId string `json:"subscriptionId"`
	ReportList []AmfEventReport `json:"reportList,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewAmfCreatedEventSubscription instantiates a new AmfCreatedEventSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfCreatedEventSubscription(subscription AmfEventSubscription, subscriptionId string) *AmfCreatedEventSubscription {
	this := AmfCreatedEventSubscription{}
	this.Subscription = subscription
	this.SubscriptionId = subscriptionId
	return &this
}

// NewAmfCreatedEventSubscriptionWithDefaults instantiates a new AmfCreatedEventSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfCreatedEventSubscriptionWithDefaults() *AmfCreatedEventSubscription {
	this := AmfCreatedEventSubscription{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *AmfCreatedEventSubscription) GetSubscription() AmfEventSubscription {
	if o == nil {
		var ret AmfEventSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *AmfCreatedEventSubscription) GetSubscriptionOk() (*AmfEventSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *AmfCreatedEventSubscription) SetSubscription(v AmfEventSubscription) {
	o.Subscription = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *AmfCreatedEventSubscription) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *AmfCreatedEventSubscription) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *AmfCreatedEventSubscription) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetReportList returns the ReportList field value if set, zero value otherwise.
func (o *AmfCreatedEventSubscription) GetReportList() []AmfEventReport {
	if o == nil || isNil(o.ReportList) {
		var ret []AmfEventReport
		return ret
	}
	return o.ReportList
}

// GetReportListOk returns a tuple with the ReportList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfCreatedEventSubscription) GetReportListOk() ([]AmfEventReport, bool) {
	if o == nil || isNil(o.ReportList) {
		return nil, false
	}
	return o.ReportList, true
}

// HasReportList returns a boolean if a field has been set.
func (o *AmfCreatedEventSubscription) HasReportList() bool {
	if o != nil && !isNil(o.ReportList) {
		return true
	}

	return false
}

// SetReportList gets a reference to the given []AmfEventReport and assigns it to the ReportList field.
func (o *AmfCreatedEventSubscription) SetReportList(v []AmfEventReport) {
	o.ReportList = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *AmfCreatedEventSubscription) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfCreatedEventSubscription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *AmfCreatedEventSubscription) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *AmfCreatedEventSubscription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o AmfCreatedEventSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfCreatedEventSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription"] = o.Subscription
	toSerialize["subscriptionId"] = o.SubscriptionId
	if !isNil(o.ReportList) {
		toSerialize["reportList"] = o.ReportList
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableAmfCreatedEventSubscription struct {
	value *AmfCreatedEventSubscription
	isSet bool
}

func (v NullableAmfCreatedEventSubscription) Get() *AmfCreatedEventSubscription {
	return v.value
}

func (v *NullableAmfCreatedEventSubscription) Set(val *AmfCreatedEventSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfCreatedEventSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfCreatedEventSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfCreatedEventSubscription(val *AmfCreatedEventSubscription) *NullableAmfCreatedEventSubscription {
	return &NullableAmfCreatedEventSubscription{value: val, isSet: true}
}

func (v NullableAmfCreatedEventSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfCreatedEventSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



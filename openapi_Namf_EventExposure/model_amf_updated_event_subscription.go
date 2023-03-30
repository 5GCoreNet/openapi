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

// checks if the AmfUpdatedEventSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfUpdatedEventSubscription{}

// AmfUpdatedEventSubscription Represents a successful update on an AMF Event Subscription
type AmfUpdatedEventSubscription struct {
	Subscription AmfEventSubscription `json:"subscription"`
	ReportList []AmfEventReport `json:"reportList,omitempty"`
}

// NewAmfUpdatedEventSubscription instantiates a new AmfUpdatedEventSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfUpdatedEventSubscription(subscription AmfEventSubscription) *AmfUpdatedEventSubscription {
	this := AmfUpdatedEventSubscription{}
	this.Subscription = subscription
	return &this
}

// NewAmfUpdatedEventSubscriptionWithDefaults instantiates a new AmfUpdatedEventSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfUpdatedEventSubscriptionWithDefaults() *AmfUpdatedEventSubscription {
	this := AmfUpdatedEventSubscription{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *AmfUpdatedEventSubscription) GetSubscription() AmfEventSubscription {
	if o == nil {
		var ret AmfEventSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *AmfUpdatedEventSubscription) GetSubscriptionOk() (*AmfEventSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *AmfUpdatedEventSubscription) SetSubscription(v AmfEventSubscription) {
	o.Subscription = v
}

// GetReportList returns the ReportList field value if set, zero value otherwise.
func (o *AmfUpdatedEventSubscription) GetReportList() []AmfEventReport {
	if o == nil || IsNil(o.ReportList) {
		var ret []AmfEventReport
		return ret
	}
	return o.ReportList
}

// GetReportListOk returns a tuple with the ReportList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfUpdatedEventSubscription) GetReportListOk() ([]AmfEventReport, bool) {
	if o == nil || IsNil(o.ReportList) {
		return nil, false
	}
	return o.ReportList, true
}

// HasReportList returns a boolean if a field has been set.
func (o *AmfUpdatedEventSubscription) HasReportList() bool {
	if o != nil && !IsNil(o.ReportList) {
		return true
	}

	return false
}

// SetReportList gets a reference to the given []AmfEventReport and assigns it to the ReportList field.
func (o *AmfUpdatedEventSubscription) SetReportList(v []AmfEventReport) {
	o.ReportList = v
}

func (o AmfUpdatedEventSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfUpdatedEventSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription"] = o.Subscription
	if !IsNil(o.ReportList) {
		toSerialize["reportList"] = o.ReportList
	}
	return toSerialize, nil
}

type NullableAmfUpdatedEventSubscription struct {
	value *AmfUpdatedEventSubscription
	isSet bool
}

func (v NullableAmfUpdatedEventSubscription) Get() *AmfUpdatedEventSubscription {
	return v.value
}

func (v *NullableAmfUpdatedEventSubscription) Set(val *AmfUpdatedEventSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfUpdatedEventSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfUpdatedEventSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfUpdatedEventSubscription(val *AmfUpdatedEventSubscription) *NullableAmfUpdatedEventSubscription {
	return &NullableAmfUpdatedEventSubscription{value: val, isSet: true}
}

func (v NullableAmfUpdatedEventSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfUpdatedEventSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



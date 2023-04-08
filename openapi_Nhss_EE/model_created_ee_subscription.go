/*
Nhss_EE

HSS Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_EE

import (
	"encoding/json"
)

// checks if the CreatedEeSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatedEeSubscription{}

// CreatedEeSubscription It represents the response body of the subscription request, containing data of the created subscription in the HSS
type CreatedEeSubscription struct {
	EeSubscription EeSubscription `json:"eeSubscription"`
	EventReports []MonitoringReport `json:"eventReports,omitempty"`
	// A map (list of key-value pairs where referenceId converted from integer to string serves as key; see clause 6.4.6.3.2) of FailedMonitoringConfiguration
	FailedMonitoringConfigs *map[string]FailedMonitoringConfiguration `json:"failedMonitoringConfigs,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewCreatedEeSubscription instantiates a new CreatedEeSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedEeSubscription(eeSubscription EeSubscription) *CreatedEeSubscription {
	this := CreatedEeSubscription{}
	this.EeSubscription = eeSubscription
	return &this
}

// NewCreatedEeSubscriptionWithDefaults instantiates a new CreatedEeSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedEeSubscriptionWithDefaults() *CreatedEeSubscription {
	this := CreatedEeSubscription{}
	return &this
}

// GetEeSubscription returns the EeSubscription field value
func (o *CreatedEeSubscription) GetEeSubscription() EeSubscription {
	if o == nil {
		var ret EeSubscription
		return ret
	}

	return o.EeSubscription
}

// GetEeSubscriptionOk returns a tuple with the EeSubscription field value
// and a boolean to check if the value has been set.
func (o *CreatedEeSubscription) GetEeSubscriptionOk() (*EeSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EeSubscription, true
}

// SetEeSubscription sets field value
func (o *CreatedEeSubscription) SetEeSubscription(v EeSubscription) {
	o.EeSubscription = v
}

// GetEventReports returns the EventReports field value if set, zero value otherwise.
func (o *CreatedEeSubscription) GetEventReports() []MonitoringReport {
	if o == nil || isNil(o.EventReports) {
		var ret []MonitoringReport
		return ret
	}
	return o.EventReports
}

// GetEventReportsOk returns a tuple with the EventReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedEeSubscription) GetEventReportsOk() ([]MonitoringReport, bool) {
	if o == nil || isNil(o.EventReports) {
		return nil, false
	}
	return o.EventReports, true
}

// HasEventReports returns a boolean if a field has been set.
func (o *CreatedEeSubscription) HasEventReports() bool {
	if o != nil && !isNil(o.EventReports) {
		return true
	}

	return false
}

// SetEventReports gets a reference to the given []MonitoringReport and assigns it to the EventReports field.
func (o *CreatedEeSubscription) SetEventReports(v []MonitoringReport) {
	o.EventReports = v
}

// GetFailedMonitoringConfigs returns the FailedMonitoringConfigs field value if set, zero value otherwise.
func (o *CreatedEeSubscription) GetFailedMonitoringConfigs() map[string]FailedMonitoringConfiguration {
	if o == nil || isNil(o.FailedMonitoringConfigs) {
		var ret map[string]FailedMonitoringConfiguration
		return ret
	}
	return *o.FailedMonitoringConfigs
}

// GetFailedMonitoringConfigsOk returns a tuple with the FailedMonitoringConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedEeSubscription) GetFailedMonitoringConfigsOk() (*map[string]FailedMonitoringConfiguration, bool) {
	if o == nil || isNil(o.FailedMonitoringConfigs) {
		return nil, false
	}
	return o.FailedMonitoringConfigs, true
}

// HasFailedMonitoringConfigs returns a boolean if a field has been set.
func (o *CreatedEeSubscription) HasFailedMonitoringConfigs() bool {
	if o != nil && !isNil(o.FailedMonitoringConfigs) {
		return true
	}

	return false
}

// SetFailedMonitoringConfigs gets a reference to the given map[string]FailedMonitoringConfiguration and assigns it to the FailedMonitoringConfigs field.
func (o *CreatedEeSubscription) SetFailedMonitoringConfigs(v map[string]FailedMonitoringConfiguration) {
	o.FailedMonitoringConfigs = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *CreatedEeSubscription) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedEeSubscription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *CreatedEeSubscription) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *CreatedEeSubscription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o CreatedEeSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatedEeSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eeSubscription"] = o.EeSubscription
	if !isNil(o.EventReports) {
		toSerialize["eventReports"] = o.EventReports
	}
	if !isNil(o.FailedMonitoringConfigs) {
		toSerialize["failedMonitoringConfigs"] = o.FailedMonitoringConfigs
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableCreatedEeSubscription struct {
	value *CreatedEeSubscription
	isSet bool
}

func (v NullableCreatedEeSubscription) Get() *CreatedEeSubscription {
	return v.value
}

func (v *NullableCreatedEeSubscription) Set(val *CreatedEeSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedEeSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedEeSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedEeSubscription(val *CreatedEeSubscription) *NullableCreatedEeSubscription {
	return &NullableCreatedEeSubscription{value: val, isSet: true}
}

func (v NullableCreatedEeSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedEeSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



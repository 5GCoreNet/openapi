/*
Nsoraf_SOR

Nsoraf Steering Of Roaming Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsoraf_SOR

import (
	"encoding/json"
	"time"
)

// checks if the SorInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SorInformation{}

// SorInformation Represents the SoR information to be conveyed to a UE.
type SorInformation struct {
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	SteeringContainer *SteeringContainer `json:"steeringContainer,omitempty"`
	SorAckIndication bool `json:"sorAckIndication"`
	// string with format 'bytes' as defined in OpenAPI
	SorCmci *string `json:"sorCmci,omitempty"`
	StoreSorCmciInMe *bool `json:"storeSorCmciInMe,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	SorSendingTime time.Time `json:"sorSendingTime"`
}

// NewSorInformation instantiates a new SorInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSorInformation(sorAckIndication bool, sorSendingTime time.Time) *SorInformation {
	this := SorInformation{}
	this.SorAckIndication = sorAckIndication
	this.SorSendingTime = sorSendingTime
	return &this
}

// NewSorInformationWithDefaults instantiates a new SorInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSorInformationWithDefaults() *SorInformation {
	this := SorInformation{}
	return &this
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SorInformation) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorInformation) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SorInformation) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SorInformation) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetSteeringContainer returns the SteeringContainer field value if set, zero value otherwise.
func (o *SorInformation) GetSteeringContainer() SteeringContainer {
	if o == nil || isNil(o.SteeringContainer) {
		var ret SteeringContainer
		return ret
	}
	return *o.SteeringContainer
}

// GetSteeringContainerOk returns a tuple with the SteeringContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorInformation) GetSteeringContainerOk() (*SteeringContainer, bool) {
	if o == nil || isNil(o.SteeringContainer) {
		return nil, false
	}
	return o.SteeringContainer, true
}

// HasSteeringContainer returns a boolean if a field has been set.
func (o *SorInformation) HasSteeringContainer() bool {
	if o != nil && !isNil(o.SteeringContainer) {
		return true
	}

	return false
}

// SetSteeringContainer gets a reference to the given SteeringContainer and assigns it to the SteeringContainer field.
func (o *SorInformation) SetSteeringContainer(v SteeringContainer) {
	o.SteeringContainer = &v
}

// GetSorAckIndication returns the SorAckIndication field value
func (o *SorInformation) GetSorAckIndication() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SorAckIndication
}

// GetSorAckIndicationOk returns a tuple with the SorAckIndication field value
// and a boolean to check if the value has been set.
func (o *SorInformation) GetSorAckIndicationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SorAckIndication, true
}

// SetSorAckIndication sets field value
func (o *SorInformation) SetSorAckIndication(v bool) {
	o.SorAckIndication = v
}

// GetSorCmci returns the SorCmci field value if set, zero value otherwise.
func (o *SorInformation) GetSorCmci() string {
	if o == nil || isNil(o.SorCmci) {
		var ret string
		return ret
	}
	return *o.SorCmci
}

// GetSorCmciOk returns a tuple with the SorCmci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorInformation) GetSorCmciOk() (*string, bool) {
	if o == nil || isNil(o.SorCmci) {
		return nil, false
	}
	return o.SorCmci, true
}

// HasSorCmci returns a boolean if a field has been set.
func (o *SorInformation) HasSorCmci() bool {
	if o != nil && !isNil(o.SorCmci) {
		return true
	}

	return false
}

// SetSorCmci gets a reference to the given string and assigns it to the SorCmci field.
func (o *SorInformation) SetSorCmci(v string) {
	o.SorCmci = &v
}

// GetStoreSorCmciInMe returns the StoreSorCmciInMe field value if set, zero value otherwise.
func (o *SorInformation) GetStoreSorCmciInMe() bool {
	if o == nil || isNil(o.StoreSorCmciInMe) {
		var ret bool
		return ret
	}
	return *o.StoreSorCmciInMe
}

// GetStoreSorCmciInMeOk returns a tuple with the StoreSorCmciInMe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorInformation) GetStoreSorCmciInMeOk() (*bool, bool) {
	if o == nil || isNil(o.StoreSorCmciInMe) {
		return nil, false
	}
	return o.StoreSorCmciInMe, true
}

// HasStoreSorCmciInMe returns a boolean if a field has been set.
func (o *SorInformation) HasStoreSorCmciInMe() bool {
	if o != nil && !isNil(o.StoreSorCmciInMe) {
		return true
	}

	return false
}

// SetStoreSorCmciInMe gets a reference to the given bool and assigns it to the StoreSorCmciInMe field.
func (o *SorInformation) SetStoreSorCmciInMe(v bool) {
	o.StoreSorCmciInMe = &v
}

// GetSorSendingTime returns the SorSendingTime field value
func (o *SorInformation) GetSorSendingTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SorSendingTime
}

// GetSorSendingTimeOk returns a tuple with the SorSendingTime field value
// and a boolean to check if the value has been set.
func (o *SorInformation) GetSorSendingTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SorSendingTime, true
}

// SetSorSendingTime sets field value
func (o *SorInformation) SetSorSendingTime(v time.Time) {
	o.SorSendingTime = v
}

func (o SorInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SorInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.SteeringContainer) {
		toSerialize["steeringContainer"] = o.SteeringContainer
	}
	toSerialize["sorAckIndication"] = o.SorAckIndication
	if !isNil(o.SorCmci) {
		toSerialize["sorCmci"] = o.SorCmci
	}
	if !isNil(o.StoreSorCmciInMe) {
		toSerialize["storeSorCmciInMe"] = o.StoreSorCmciInMe
	}
	toSerialize["sorSendingTime"] = o.SorSendingTime
	return toSerialize, nil
}

type NullableSorInformation struct {
	value *SorInformation
	isSet bool
}

func (v NullableSorInformation) Get() *SorInformation {
	return v.value
}

func (v *NullableSorInformation) Set(val *SorInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableSorInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableSorInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSorInformation(val *SorInformation) *NullableSorInformation {
	return &NullableSorInformation{value: val, isSet: true}
}

func (v NullableSorInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSorInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
3gpp-msisdn-less-mo-sms

API for MSISDN-less Mobile Originated SMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MsisdnLessMoSms

import (
	"encoding/json"
)

// checks if the MsisdnLessMoSmsNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MsisdnLessMoSmsNotification{}

// MsisdnLessMoSmsNotification Represents a MSISDN-less MO SMS notification.
type MsisdnLessMoSmsNotification struct {
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures"`
	// String with format \"byte\" as defined in OpenAPI Specification, i.e, base64-encoded characters.
	Sms string `json:"sms"`
	// External identifier has the form username@realm.
	ExternalId string `json:"externalId"`
	// Unsigned integer with valid values between 0 and 65535.
	ApplicationPort int32 `json:"applicationPort"`
}

// NewMsisdnLessMoSmsNotification instantiates a new MsisdnLessMoSmsNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsisdnLessMoSmsNotification(supportedFeatures string, sms string, externalId string, applicationPort int32) *MsisdnLessMoSmsNotification {
	this := MsisdnLessMoSmsNotification{}
	this.SupportedFeatures = supportedFeatures
	this.Sms = sms
	this.ExternalId = externalId
	this.ApplicationPort = applicationPort
	return &this
}

// NewMsisdnLessMoSmsNotificationWithDefaults instantiates a new MsisdnLessMoSmsNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsisdnLessMoSmsNotificationWithDefaults() *MsisdnLessMoSmsNotification {
	this := MsisdnLessMoSmsNotification{}
	return &this
}

// GetSupportedFeatures returns the SupportedFeatures field value
func (o *MsisdnLessMoSmsNotification) GetSupportedFeatures() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value
// and a boolean to check if the value has been set.
func (o *MsisdnLessMoSmsNotification) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedFeatures, true
}

// SetSupportedFeatures sets field value
func (o *MsisdnLessMoSmsNotification) SetSupportedFeatures(v string) {
	o.SupportedFeatures = v
}

// GetSms returns the Sms field value
func (o *MsisdnLessMoSmsNotification) GetSms() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sms
}

// GetSmsOk returns a tuple with the Sms field value
// and a boolean to check if the value has been set.
func (o *MsisdnLessMoSmsNotification) GetSmsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sms, true
}

// SetSms sets field value
func (o *MsisdnLessMoSmsNotification) SetSms(v string) {
	o.Sms = v
}

// GetExternalId returns the ExternalId field value
func (o *MsisdnLessMoSmsNotification) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *MsisdnLessMoSmsNotification) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *MsisdnLessMoSmsNotification) SetExternalId(v string) {
	o.ExternalId = v
}

// GetApplicationPort returns the ApplicationPort field value
func (o *MsisdnLessMoSmsNotification) GetApplicationPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationPort
}

// GetApplicationPortOk returns a tuple with the ApplicationPort field value
// and a boolean to check if the value has been set.
func (o *MsisdnLessMoSmsNotification) GetApplicationPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationPort, true
}

// SetApplicationPort sets field value
func (o *MsisdnLessMoSmsNotification) SetApplicationPort(v int32) {
	o.ApplicationPort = v
}

func (o MsisdnLessMoSmsNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MsisdnLessMoSmsNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supportedFeatures"] = o.SupportedFeatures
	toSerialize["sms"] = o.Sms
	toSerialize["externalId"] = o.ExternalId
	toSerialize["applicationPort"] = o.ApplicationPort
	return toSerialize, nil
}

type NullableMsisdnLessMoSmsNotification struct {
	value *MsisdnLessMoSmsNotification
	isSet bool
}

func (v NullableMsisdnLessMoSmsNotification) Get() *MsisdnLessMoSmsNotification {
	return v.value
}

func (v *NullableMsisdnLessMoSmsNotification) Set(val *MsisdnLessMoSmsNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableMsisdnLessMoSmsNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableMsisdnLessMoSmsNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsisdnLessMoSmsNotification(val *MsisdnLessMoSmsNotification) *NullableMsisdnLessMoSmsNotification {
	return &NullableMsisdnLessMoSmsNotification{value: val, isSet: true}
}

func (v NullableMsisdnLessMoSmsNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsisdnLessMoSmsNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

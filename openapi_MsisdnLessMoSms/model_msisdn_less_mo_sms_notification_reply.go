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

// checks if the MsisdnLessMoSmsNotificationReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MsisdnLessMoSmsNotificationReply{}

// MsisdnLessMoSmsNotificationReply Represents a reply to a MSISDN-less MO SMS notification.
type MsisdnLessMoSmsNotificationReply struct {
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures"`
}

// NewMsisdnLessMoSmsNotificationReply instantiates a new MsisdnLessMoSmsNotificationReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsisdnLessMoSmsNotificationReply(supportedFeatures string) *MsisdnLessMoSmsNotificationReply {
	this := MsisdnLessMoSmsNotificationReply{}
	this.SupportedFeatures = supportedFeatures
	return &this
}

// NewMsisdnLessMoSmsNotificationReplyWithDefaults instantiates a new MsisdnLessMoSmsNotificationReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsisdnLessMoSmsNotificationReplyWithDefaults() *MsisdnLessMoSmsNotificationReply {
	this := MsisdnLessMoSmsNotificationReply{}
	return &this
}

// GetSupportedFeatures returns the SupportedFeatures field value
func (o *MsisdnLessMoSmsNotificationReply) GetSupportedFeatures() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value
// and a boolean to check if the value has been set.
func (o *MsisdnLessMoSmsNotificationReply) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedFeatures, true
}

// SetSupportedFeatures sets field value
func (o *MsisdnLessMoSmsNotificationReply) SetSupportedFeatures(v string) {
	o.SupportedFeatures = v
}

func (o MsisdnLessMoSmsNotificationReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MsisdnLessMoSmsNotificationReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supportedFeatures"] = o.SupportedFeatures
	return toSerialize, nil
}

type NullableMsisdnLessMoSmsNotificationReply struct {
	value *MsisdnLessMoSmsNotificationReply
	isSet bool
}

func (v NullableMsisdnLessMoSmsNotificationReply) Get() *MsisdnLessMoSmsNotificationReply {
	return v.value
}

func (v *NullableMsisdnLessMoSmsNotificationReply) Set(val *MsisdnLessMoSmsNotificationReply) {
	v.value = val
	v.isSet = true
}

func (v NullableMsisdnLessMoSmsNotificationReply) IsSet() bool {
	return v.isSet
}

func (v *NullableMsisdnLessMoSmsNotificationReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsisdnLessMoSmsNotificationReply(val *MsisdnLessMoSmsNotificationReply) *NullableMsisdnLessMoSmsNotificationReply {
	return &NullableMsisdnLessMoSmsNotificationReply{value: val, isSet: true}
}

func (v NullableMsisdnLessMoSmsNotificationReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsisdnLessMoSmsNotificationReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



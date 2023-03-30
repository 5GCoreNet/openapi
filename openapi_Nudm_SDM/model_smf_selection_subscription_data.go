/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
)

// checks if the SmfSelectionSubscriptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmfSelectionSubscriptionData{}

// SmfSelectionSubscriptionData struct for SmfSelectionSubscriptionData
type SmfSelectionSubscriptionData struct {
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// A map(list of key-value pairs) where singleNssai serves as key of SnssaiInfo
	SubscribedSnssaiInfos *map[string]SnssaiInfo `json:"subscribedSnssaiInfos,omitempty"`
	SharedSnssaiInfosId *string `json:"sharedSnssaiInfosId,omitempty"`
	// Identifier of a group of NFs.
	HssGroupId *string `json:"hssGroupId,omitempty"`
}

// NewSmfSelectionSubscriptionData instantiates a new SmfSelectionSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmfSelectionSubscriptionData() *SmfSelectionSubscriptionData {
	this := SmfSelectionSubscriptionData{}
	return &this
}

// NewSmfSelectionSubscriptionDataWithDefaults instantiates a new SmfSelectionSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmfSelectionSubscriptionDataWithDefaults() *SmfSelectionSubscriptionData {
	this := SmfSelectionSubscriptionData{}
	return &this
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SmfSelectionSubscriptionData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfSelectionSubscriptionData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SmfSelectionSubscriptionData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SmfSelectionSubscriptionData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetSubscribedSnssaiInfos returns the SubscribedSnssaiInfos field value if set, zero value otherwise.
func (o *SmfSelectionSubscriptionData) GetSubscribedSnssaiInfos() map[string]SnssaiInfo {
	if o == nil || IsNil(o.SubscribedSnssaiInfos) {
		var ret map[string]SnssaiInfo
		return ret
	}
	return *o.SubscribedSnssaiInfos
}

// GetSubscribedSnssaiInfosOk returns a tuple with the SubscribedSnssaiInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfSelectionSubscriptionData) GetSubscribedSnssaiInfosOk() (*map[string]SnssaiInfo, bool) {
	if o == nil || IsNil(o.SubscribedSnssaiInfos) {
		return nil, false
	}
	return o.SubscribedSnssaiInfos, true
}

// HasSubscribedSnssaiInfos returns a boolean if a field has been set.
func (o *SmfSelectionSubscriptionData) HasSubscribedSnssaiInfos() bool {
	if o != nil && !IsNil(o.SubscribedSnssaiInfos) {
		return true
	}

	return false
}

// SetSubscribedSnssaiInfos gets a reference to the given map[string]SnssaiInfo and assigns it to the SubscribedSnssaiInfos field.
func (o *SmfSelectionSubscriptionData) SetSubscribedSnssaiInfos(v map[string]SnssaiInfo) {
	o.SubscribedSnssaiInfos = &v
}

// GetSharedSnssaiInfosId returns the SharedSnssaiInfosId field value if set, zero value otherwise.
func (o *SmfSelectionSubscriptionData) GetSharedSnssaiInfosId() string {
	if o == nil || IsNil(o.SharedSnssaiInfosId) {
		var ret string
		return ret
	}
	return *o.SharedSnssaiInfosId
}

// GetSharedSnssaiInfosIdOk returns a tuple with the SharedSnssaiInfosId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfSelectionSubscriptionData) GetSharedSnssaiInfosIdOk() (*string, bool) {
	if o == nil || IsNil(o.SharedSnssaiInfosId) {
		return nil, false
	}
	return o.SharedSnssaiInfosId, true
}

// HasSharedSnssaiInfosId returns a boolean if a field has been set.
func (o *SmfSelectionSubscriptionData) HasSharedSnssaiInfosId() bool {
	if o != nil && !IsNil(o.SharedSnssaiInfosId) {
		return true
	}

	return false
}

// SetSharedSnssaiInfosId gets a reference to the given string and assigns it to the SharedSnssaiInfosId field.
func (o *SmfSelectionSubscriptionData) SetSharedSnssaiInfosId(v string) {
	o.SharedSnssaiInfosId = &v
}

// GetHssGroupId returns the HssGroupId field value if set, zero value otherwise.
func (o *SmfSelectionSubscriptionData) GetHssGroupId() string {
	if o == nil || IsNil(o.HssGroupId) {
		var ret string
		return ret
	}
	return *o.HssGroupId
}

// GetHssGroupIdOk returns a tuple with the HssGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfSelectionSubscriptionData) GetHssGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.HssGroupId) {
		return nil, false
	}
	return o.HssGroupId, true
}

// HasHssGroupId returns a boolean if a field has been set.
func (o *SmfSelectionSubscriptionData) HasHssGroupId() bool {
	if o != nil && !IsNil(o.HssGroupId) {
		return true
	}

	return false
}

// SetHssGroupId gets a reference to the given string and assigns it to the HssGroupId field.
func (o *SmfSelectionSubscriptionData) SetHssGroupId(v string) {
	o.HssGroupId = &v
}

func (o SmfSelectionSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmfSelectionSubscriptionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.SubscribedSnssaiInfos) {
		toSerialize["subscribedSnssaiInfos"] = o.SubscribedSnssaiInfos
	}
	if !IsNil(o.SharedSnssaiInfosId) {
		toSerialize["sharedSnssaiInfosId"] = o.SharedSnssaiInfosId
	}
	if !IsNil(o.HssGroupId) {
		toSerialize["hssGroupId"] = o.HssGroupId
	}
	return toSerialize, nil
}

type NullableSmfSelectionSubscriptionData struct {
	value *SmfSelectionSubscriptionData
	isSet bool
}

func (v NullableSmfSelectionSubscriptionData) Get() *SmfSelectionSubscriptionData {
	return v.value
}

func (v *NullableSmfSelectionSubscriptionData) Set(val *SmfSelectionSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmfSelectionSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmfSelectionSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmfSelectionSubscriptionData(val *SmfSelectionSubscriptionData) *NullableSmfSelectionSubscriptionData {
	return &NullableSmfSelectionSubscriptionData{value: val, isSet: true}
}

func (v NullableSmfSelectionSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmfSelectionSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// checks if the EeProfileData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EeProfileData{}

// EeProfileData Event Exposure Profile Data.
type EeProfileData struct {
	RestrictedEventTypes []EventType `json:"restrictedEventTypes,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// A map (list of key-value pairs where EventType serves as key) of MTC provider lists. In addition to defined EventTypes, the key value \"ALL\" may be used to identify a map entry which contains a list of MtcProviders that are allowed monitoring all Event Types.
	AllowedMtcProvider *map[string][]MtcProvider `json:"allowedMtcProvider,omitempty"`
	IwkEpcRestricted   *bool                     `json:"iwkEpcRestricted,omitempty"`
	Imsi               *string                   `json:"imsi,omitempty"`
	// Identifier of a group of NFs.
	HssGroupId *string `json:"hssGroupId,omitempty"`
}

// NewEeProfileData instantiates a new EeProfileData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEeProfileData() *EeProfileData {
	this := EeProfileData{}
	var iwkEpcRestricted bool = false
	this.IwkEpcRestricted = &iwkEpcRestricted
	return &this
}

// NewEeProfileDataWithDefaults instantiates a new EeProfileData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEeProfileDataWithDefaults() *EeProfileData {
	this := EeProfileData{}
	var iwkEpcRestricted bool = false
	this.IwkEpcRestricted = &iwkEpcRestricted
	return &this
}

// GetRestrictedEventTypes returns the RestrictedEventTypes field value if set, zero value otherwise.
func (o *EeProfileData) GetRestrictedEventTypes() []EventType {
	if o == nil || IsNil(o.RestrictedEventTypes) {
		var ret []EventType
		return ret
	}
	return o.RestrictedEventTypes
}

// GetRestrictedEventTypesOk returns a tuple with the RestrictedEventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeProfileData) GetRestrictedEventTypesOk() ([]EventType, bool) {
	if o == nil || IsNil(o.RestrictedEventTypes) {
		return nil, false
	}
	return o.RestrictedEventTypes, true
}

// HasRestrictedEventTypes returns a boolean if a field has been set.
func (o *EeProfileData) HasRestrictedEventTypes() bool {
	if o != nil && !IsNil(o.RestrictedEventTypes) {
		return true
	}

	return false
}

// SetRestrictedEventTypes gets a reference to the given []EventType and assigns it to the RestrictedEventTypes field.
func (o *EeProfileData) SetRestrictedEventTypes(v []EventType) {
	o.RestrictedEventTypes = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *EeProfileData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeProfileData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *EeProfileData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *EeProfileData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetAllowedMtcProvider returns the AllowedMtcProvider field value if set, zero value otherwise.
func (o *EeProfileData) GetAllowedMtcProvider() map[string][]MtcProvider {
	if o == nil || IsNil(o.AllowedMtcProvider) {
		var ret map[string][]MtcProvider
		return ret
	}
	return *o.AllowedMtcProvider
}

// GetAllowedMtcProviderOk returns a tuple with the AllowedMtcProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeProfileData) GetAllowedMtcProviderOk() (*map[string][]MtcProvider, bool) {
	if o == nil || IsNil(o.AllowedMtcProvider) {
		return nil, false
	}
	return o.AllowedMtcProvider, true
}

// HasAllowedMtcProvider returns a boolean if a field has been set.
func (o *EeProfileData) HasAllowedMtcProvider() bool {
	if o != nil && !IsNil(o.AllowedMtcProvider) {
		return true
	}

	return false
}

// SetAllowedMtcProvider gets a reference to the given map[string][]MtcProvider and assigns it to the AllowedMtcProvider field.
func (o *EeProfileData) SetAllowedMtcProvider(v map[string][]MtcProvider) {
	o.AllowedMtcProvider = &v
}

// GetIwkEpcRestricted returns the IwkEpcRestricted field value if set, zero value otherwise.
func (o *EeProfileData) GetIwkEpcRestricted() bool {
	if o == nil || IsNil(o.IwkEpcRestricted) {
		var ret bool
		return ret
	}
	return *o.IwkEpcRestricted
}

// GetIwkEpcRestrictedOk returns a tuple with the IwkEpcRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeProfileData) GetIwkEpcRestrictedOk() (*bool, bool) {
	if o == nil || IsNil(o.IwkEpcRestricted) {
		return nil, false
	}
	return o.IwkEpcRestricted, true
}

// HasIwkEpcRestricted returns a boolean if a field has been set.
func (o *EeProfileData) HasIwkEpcRestricted() bool {
	if o != nil && !IsNil(o.IwkEpcRestricted) {
		return true
	}

	return false
}

// SetIwkEpcRestricted gets a reference to the given bool and assigns it to the IwkEpcRestricted field.
func (o *EeProfileData) SetIwkEpcRestricted(v bool) {
	o.IwkEpcRestricted = &v
}

// GetImsi returns the Imsi field value if set, zero value otherwise.
func (o *EeProfileData) GetImsi() string {
	if o == nil || IsNil(o.Imsi) {
		var ret string
		return ret
	}
	return *o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeProfileData) GetImsiOk() (*string, bool) {
	if o == nil || IsNil(o.Imsi) {
		return nil, false
	}
	return o.Imsi, true
}

// HasImsi returns a boolean if a field has been set.
func (o *EeProfileData) HasImsi() bool {
	if o != nil && !IsNil(o.Imsi) {
		return true
	}

	return false
}

// SetImsi gets a reference to the given string and assigns it to the Imsi field.
func (o *EeProfileData) SetImsi(v string) {
	o.Imsi = &v
}

// GetHssGroupId returns the HssGroupId field value if set, zero value otherwise.
func (o *EeProfileData) GetHssGroupId() string {
	if o == nil || IsNil(o.HssGroupId) {
		var ret string
		return ret
	}
	return *o.HssGroupId
}

// GetHssGroupIdOk returns a tuple with the HssGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeProfileData) GetHssGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.HssGroupId) {
		return nil, false
	}
	return o.HssGroupId, true
}

// HasHssGroupId returns a boolean if a field has been set.
func (o *EeProfileData) HasHssGroupId() bool {
	if o != nil && !IsNil(o.HssGroupId) {
		return true
	}

	return false
}

// SetHssGroupId gets a reference to the given string and assigns it to the HssGroupId field.
func (o *EeProfileData) SetHssGroupId(v string) {
	o.HssGroupId = &v
}

func (o EeProfileData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EeProfileData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RestrictedEventTypes) {
		toSerialize["restrictedEventTypes"] = o.RestrictedEventTypes
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.AllowedMtcProvider) {
		toSerialize["allowedMtcProvider"] = o.AllowedMtcProvider
	}
	if !IsNil(o.IwkEpcRestricted) {
		toSerialize["iwkEpcRestricted"] = o.IwkEpcRestricted
	}
	if !IsNil(o.Imsi) {
		toSerialize["imsi"] = o.Imsi
	}
	if !IsNil(o.HssGroupId) {
		toSerialize["hssGroupId"] = o.HssGroupId
	}
	return toSerialize, nil
}

type NullableEeProfileData struct {
	value *EeProfileData
	isSet bool
}

func (v NullableEeProfileData) Get() *EeProfileData {
	return v.value
}

func (v *NullableEeProfileData) Set(val *EeProfileData) {
	v.value = val
	v.isSet = true
}

func (v NullableEeProfileData) IsSet() bool {
	return v.isSet
}

func (v *NullableEeProfileData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEeProfileData(val *EeProfileData) *NullableEeProfileData {
	return &NullableEeProfileData{value: val, isSet: true}
}

func (v NullableEeProfileData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEeProfileData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
UAE Server C2 Operation Mode Management Service

UAE Server C2 Operation Mode Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_UAE_C2OperationModeManagement

import (
	"encoding/json"
)

// checks if the ConfigureData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigureData{}

// ConfigureData Represents the parameters to request to provision C2 Operation Mode configuration information for a UAS (i.e. pair of UAV and UAV-C).
type ConfigureData struct {
	// string providing an URI formatted according to IETF RFC 3986.
	UassId                string                `json:"uassId"`
	UasId                 UasId                 `json:"uasId"`
	AllowedC2CommModes    []C2CommMode          `json:"allowedC2CommModes"`
	C2CommModeSwitchTypes []C2CommModeSwitching `json:"c2CommModeSwitchTypes"`
	// string providing an URI formatted according to IETF RFC 3986.
	NotificationUri     string           `json:"notificationUri"`
	PrimaryC2CommMode   C2CommMode       `json:"primaryC2CommMode"`
	SecondaryC2CommMode *C2CommMode      `json:"secondaryC2CommMode,omitempty"`
	C2SwitchPolicies    C2SwitchPolicies `json:"c2SwitchPolicies"`
	C2ServiceArea       *C2ServiceArea   `json:"c2ServiceArea,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewConfigureData instantiates a new ConfigureData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigureData(uassId string, uasId UasId, allowedC2CommModes []C2CommMode, c2CommModeSwitchTypes []C2CommModeSwitching, notificationUri string, primaryC2CommMode C2CommMode, c2SwitchPolicies C2SwitchPolicies) *ConfigureData {
	this := ConfigureData{}
	this.UassId = uassId
	this.UasId = uasId
	this.AllowedC2CommModes = allowedC2CommModes
	this.C2CommModeSwitchTypes = c2CommModeSwitchTypes
	this.NotificationUri = notificationUri
	this.PrimaryC2CommMode = primaryC2CommMode
	this.C2SwitchPolicies = c2SwitchPolicies
	return &this
}

// NewConfigureDataWithDefaults instantiates a new ConfigureData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigureDataWithDefaults() *ConfigureData {
	this := ConfigureData{}
	return &this
}

// GetUassId returns the UassId field value
func (o *ConfigureData) GetUassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UassId
}

// GetUassIdOk returns a tuple with the UassId field value
// and a boolean to check if the value has been set.
func (o *ConfigureData) GetUassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UassId, true
}

// SetUassId sets field value
func (o *ConfigureData) SetUassId(v string) {
	o.UassId = v
}

// GetUasId returns the UasId field value
func (o *ConfigureData) GetUasId() UasId {
	if o == nil {
		var ret UasId
		return ret
	}

	return o.UasId
}

// GetUasIdOk returns a tuple with the UasId field value
// and a boolean to check if the value has been set.
func (o *ConfigureData) GetUasIdOk() (*UasId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UasId, true
}

// SetUasId sets field value
func (o *ConfigureData) SetUasId(v UasId) {
	o.UasId = v
}

// GetAllowedC2CommModes returns the AllowedC2CommModes field value
func (o *ConfigureData) GetAllowedC2CommModes() []C2CommMode {
	if o == nil {
		var ret []C2CommMode
		return ret
	}

	return o.AllowedC2CommModes
}

// GetAllowedC2CommModesOk returns a tuple with the AllowedC2CommModes field value
// and a boolean to check if the value has been set.
func (o *ConfigureData) GetAllowedC2CommModesOk() ([]C2CommMode, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedC2CommModes, true
}

// SetAllowedC2CommModes sets field value
func (o *ConfigureData) SetAllowedC2CommModes(v []C2CommMode) {
	o.AllowedC2CommModes = v
}

// GetC2CommModeSwitchTypes returns the C2CommModeSwitchTypes field value
func (o *ConfigureData) GetC2CommModeSwitchTypes() []C2CommModeSwitching {
	if o == nil {
		var ret []C2CommModeSwitching
		return ret
	}

	return o.C2CommModeSwitchTypes
}

// GetC2CommModeSwitchTypesOk returns a tuple with the C2CommModeSwitchTypes field value
// and a boolean to check if the value has been set.
func (o *ConfigureData) GetC2CommModeSwitchTypesOk() ([]C2CommModeSwitching, bool) {
	if o == nil {
		return nil, false
	}
	return o.C2CommModeSwitchTypes, true
}

// SetC2CommModeSwitchTypes sets field value
func (o *ConfigureData) SetC2CommModeSwitchTypes(v []C2CommModeSwitching) {
	o.C2CommModeSwitchTypes = v
}

// GetNotificationUri returns the NotificationUri field value
func (o *ConfigureData) GetNotificationUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value
// and a boolean to check if the value has been set.
func (o *ConfigureData) GetNotificationUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationUri, true
}

// SetNotificationUri sets field value
func (o *ConfigureData) SetNotificationUri(v string) {
	o.NotificationUri = v
}

// GetPrimaryC2CommMode returns the PrimaryC2CommMode field value
func (o *ConfigureData) GetPrimaryC2CommMode() C2CommMode {
	if o == nil {
		var ret C2CommMode
		return ret
	}

	return o.PrimaryC2CommMode
}

// GetPrimaryC2CommModeOk returns a tuple with the PrimaryC2CommMode field value
// and a boolean to check if the value has been set.
func (o *ConfigureData) GetPrimaryC2CommModeOk() (*C2CommMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryC2CommMode, true
}

// SetPrimaryC2CommMode sets field value
func (o *ConfigureData) SetPrimaryC2CommMode(v C2CommMode) {
	o.PrimaryC2CommMode = v
}

// GetSecondaryC2CommMode returns the SecondaryC2CommMode field value if set, zero value otherwise.
func (o *ConfigureData) GetSecondaryC2CommMode() C2CommMode {
	if o == nil || IsNil(o.SecondaryC2CommMode) {
		var ret C2CommMode
		return ret
	}
	return *o.SecondaryC2CommMode
}

// GetSecondaryC2CommModeOk returns a tuple with the SecondaryC2CommMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigureData) GetSecondaryC2CommModeOk() (*C2CommMode, bool) {
	if o == nil || IsNil(o.SecondaryC2CommMode) {
		return nil, false
	}
	return o.SecondaryC2CommMode, true
}

// HasSecondaryC2CommMode returns a boolean if a field has been set.
func (o *ConfigureData) HasSecondaryC2CommMode() bool {
	if o != nil && !IsNil(o.SecondaryC2CommMode) {
		return true
	}

	return false
}

// SetSecondaryC2CommMode gets a reference to the given C2CommMode and assigns it to the SecondaryC2CommMode field.
func (o *ConfigureData) SetSecondaryC2CommMode(v C2CommMode) {
	o.SecondaryC2CommMode = &v
}

// GetC2SwitchPolicies returns the C2SwitchPolicies field value
func (o *ConfigureData) GetC2SwitchPolicies() C2SwitchPolicies {
	if o == nil {
		var ret C2SwitchPolicies
		return ret
	}

	return o.C2SwitchPolicies
}

// GetC2SwitchPoliciesOk returns a tuple with the C2SwitchPolicies field value
// and a boolean to check if the value has been set.
func (o *ConfigureData) GetC2SwitchPoliciesOk() (*C2SwitchPolicies, bool) {
	if o == nil {
		return nil, false
	}
	return &o.C2SwitchPolicies, true
}

// SetC2SwitchPolicies sets field value
func (o *ConfigureData) SetC2SwitchPolicies(v C2SwitchPolicies) {
	o.C2SwitchPolicies = v
}

// GetC2ServiceArea returns the C2ServiceArea field value if set, zero value otherwise.
func (o *ConfigureData) GetC2ServiceArea() C2ServiceArea {
	if o == nil || IsNil(o.C2ServiceArea) {
		var ret C2ServiceArea
		return ret
	}
	return *o.C2ServiceArea
}

// GetC2ServiceAreaOk returns a tuple with the C2ServiceArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigureData) GetC2ServiceAreaOk() (*C2ServiceArea, bool) {
	if o == nil || IsNil(o.C2ServiceArea) {
		return nil, false
	}
	return o.C2ServiceArea, true
}

// HasC2ServiceArea returns a boolean if a field has been set.
func (o *ConfigureData) HasC2ServiceArea() bool {
	if o != nil && !IsNil(o.C2ServiceArea) {
		return true
	}

	return false
}

// SetC2ServiceArea gets a reference to the given C2ServiceArea and assigns it to the C2ServiceArea field.
func (o *ConfigureData) SetC2ServiceArea(v C2ServiceArea) {
	o.C2ServiceArea = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *ConfigureData) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigureData) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *ConfigureData) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *ConfigureData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o ConfigureData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigureData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uassId"] = o.UassId
	toSerialize["uasId"] = o.UasId
	toSerialize["allowedC2CommModes"] = o.AllowedC2CommModes
	toSerialize["c2CommModeSwitchTypes"] = o.C2CommModeSwitchTypes
	toSerialize["notificationUri"] = o.NotificationUri
	toSerialize["primaryC2CommMode"] = o.PrimaryC2CommMode
	if !IsNil(o.SecondaryC2CommMode) {
		toSerialize["secondaryC2CommMode"] = o.SecondaryC2CommMode
	}
	toSerialize["c2SwitchPolicies"] = o.C2SwitchPolicies
	if !IsNil(o.C2ServiceArea) {
		toSerialize["c2ServiceArea"] = o.C2ServiceArea
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableConfigureData struct {
	value *ConfigureData
	isSet bool
}

func (v NullableConfigureData) Get() *ConfigureData {
	return v.value
}

func (v *NullableConfigureData) Set(val *ConfigureData) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigureData) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigureData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigureData(val *ConfigureData) *NullableConfigureData {
	return &NullableConfigureData{value: val, isSet: true}
}

func (v NullableConfigureData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigureData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

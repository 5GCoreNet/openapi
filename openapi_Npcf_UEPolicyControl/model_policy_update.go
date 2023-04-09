/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_UEPolicyControl

import (
	"encoding/json"
)

// checks if the PolicyUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyUpdate{}

// PolicyUpdate Represents updated policies that the PCF provides in a notification or in the reply to an Update Request.
type PolicyUpdate struct {
	// String providing an URI formatted according to RFC 3986.
	ResourceUri string `json:"resourceUri"`
	// string with format 'bytes' as defined in OpenAPI
	UePolicy      *string        `json:"uePolicy,omitempty"`
	N2Pc5Pol      *N2InfoContent `json:"n2Pc5Pol,omitempty"`
	N2Pc5ProSePol *N2InfoContent `json:"n2Pc5ProSePol,omitempty"`
	// Request Triggers that the PCF subscribes. Only values \"LOC_CH\" and \"PRA_CH\" are permitted.
	Triggers []RequestTrigger `json:"triggers,omitempty"`
	// Contains the presence reporting area(s) for which reporting was requested. The praId attribute within the PresenceInfo data type is the key of the map.
	Pras map[string]PresenceInfo `json:"pras,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewPolicyUpdate instantiates a new PolicyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyUpdate(resourceUri string) *PolicyUpdate {
	this := PolicyUpdate{}
	this.ResourceUri = resourceUri
	return &this
}

// NewPolicyUpdateWithDefaults instantiates a new PolicyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyUpdateWithDefaults() *PolicyUpdate {
	this := PolicyUpdate{}
	return &this
}

// GetResourceUri returns the ResourceUri field value
func (o *PolicyUpdate) GetResourceUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceUri
}

// GetResourceUriOk returns a tuple with the ResourceUri field value
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetResourceUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceUri, true
}

// SetResourceUri sets field value
func (o *PolicyUpdate) SetResourceUri(v string) {
	o.ResourceUri = v
}

// GetUePolicy returns the UePolicy field value if set, zero value otherwise.
func (o *PolicyUpdate) GetUePolicy() string {
	if o == nil || IsNil(o.UePolicy) {
		var ret string
		return ret
	}
	return *o.UePolicy
}

// GetUePolicyOk returns a tuple with the UePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetUePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.UePolicy) {
		return nil, false
	}
	return o.UePolicy, true
}

// HasUePolicy returns a boolean if a field has been set.
func (o *PolicyUpdate) HasUePolicy() bool {
	if o != nil && !IsNil(o.UePolicy) {
		return true
	}

	return false
}

// SetUePolicy gets a reference to the given string and assigns it to the UePolicy field.
func (o *PolicyUpdate) SetUePolicy(v string) {
	o.UePolicy = &v
}

// GetN2Pc5Pol returns the N2Pc5Pol field value if set, zero value otherwise.
func (o *PolicyUpdate) GetN2Pc5Pol() N2InfoContent {
	if o == nil || IsNil(o.N2Pc5Pol) {
		var ret N2InfoContent
		return ret
	}
	return *o.N2Pc5Pol
}

// GetN2Pc5PolOk returns a tuple with the N2Pc5Pol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetN2Pc5PolOk() (*N2InfoContent, bool) {
	if o == nil || IsNil(o.N2Pc5Pol) {
		return nil, false
	}
	return o.N2Pc5Pol, true
}

// HasN2Pc5Pol returns a boolean if a field has been set.
func (o *PolicyUpdate) HasN2Pc5Pol() bool {
	if o != nil && !IsNil(o.N2Pc5Pol) {
		return true
	}

	return false
}

// SetN2Pc5Pol gets a reference to the given N2InfoContent and assigns it to the N2Pc5Pol field.
func (o *PolicyUpdate) SetN2Pc5Pol(v N2InfoContent) {
	o.N2Pc5Pol = &v
}

// GetN2Pc5ProSePol returns the N2Pc5ProSePol field value if set, zero value otherwise.
func (o *PolicyUpdate) GetN2Pc5ProSePol() N2InfoContent {
	if o == nil || IsNil(o.N2Pc5ProSePol) {
		var ret N2InfoContent
		return ret
	}
	return *o.N2Pc5ProSePol
}

// GetN2Pc5ProSePolOk returns a tuple with the N2Pc5ProSePol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetN2Pc5ProSePolOk() (*N2InfoContent, bool) {
	if o == nil || IsNil(o.N2Pc5ProSePol) {
		return nil, false
	}
	return o.N2Pc5ProSePol, true
}

// HasN2Pc5ProSePol returns a boolean if a field has been set.
func (o *PolicyUpdate) HasN2Pc5ProSePol() bool {
	if o != nil && !IsNil(o.N2Pc5ProSePol) {
		return true
	}

	return false
}

// SetN2Pc5ProSePol gets a reference to the given N2InfoContent and assigns it to the N2Pc5ProSePol field.
func (o *PolicyUpdate) SetN2Pc5ProSePol(v N2InfoContent) {
	o.N2Pc5ProSePol = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyUpdate) GetTriggers() []RequestTrigger {
	if o == nil {
		var ret []RequestTrigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyUpdate) GetTriggersOk() ([]RequestTrigger, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *PolicyUpdate) HasTriggers() bool {
	if o != nil && IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []RequestTrigger and assigns it to the Triggers field.
func (o *PolicyUpdate) SetTriggers(v []RequestTrigger) {
	o.Triggers = v
}

// GetPras returns the Pras field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyUpdate) GetPras() map[string]PresenceInfo {
	if o == nil {
		var ret map[string]PresenceInfo
		return ret
	}
	return o.Pras
}

// GetPrasOk returns a tuple with the Pras field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyUpdate) GetPrasOk() (*map[string]PresenceInfo, bool) {
	if o == nil || IsNil(o.Pras) {
		return nil, false
	}
	return &o.Pras, true
}

// HasPras returns a boolean if a field has been set.
func (o *PolicyUpdate) HasPras() bool {
	if o != nil && IsNil(o.Pras) {
		return true
	}

	return false
}

// SetPras gets a reference to the given map[string]PresenceInfo and assigns it to the Pras field.
func (o *PolicyUpdate) SetPras(v map[string]PresenceInfo) {
	o.Pras = v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *PolicyUpdate) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *PolicyUpdate) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *PolicyUpdate) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o PolicyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceUri"] = o.ResourceUri
	if !IsNil(o.UePolicy) {
		toSerialize["uePolicy"] = o.UePolicy
	}
	if !IsNil(o.N2Pc5Pol) {
		toSerialize["n2Pc5Pol"] = o.N2Pc5Pol
	}
	if !IsNil(o.N2Pc5ProSePol) {
		toSerialize["n2Pc5ProSePol"] = o.N2Pc5ProSePol
	}
	if o.Triggers != nil {
		toSerialize["triggers"] = o.Triggers
	}
	if o.Pras != nil {
		toSerialize["pras"] = o.Pras
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullablePolicyUpdate struct {
	value *PolicyUpdate
	isSet bool
}

func (v NullablePolicyUpdate) Get() *PolicyUpdate {
	return v.value
}

func (v *NullablePolicyUpdate) Set(val *PolicyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyUpdate(val *PolicyUpdate) *NullablePolicyUpdate {
	return &NullablePolicyUpdate{value: val, isSet: true}
}

func (v NullablePolicyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

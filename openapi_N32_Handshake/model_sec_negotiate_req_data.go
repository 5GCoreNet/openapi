/*
N32 Handshake API

N32-c Handshake Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N32_Handshake

import (
	"encoding/json"
)

// checks if the SecNegotiateReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecNegotiateReqData{}

// SecNegotiateReqData Defines the security capabilities of a SEPP sent to a receiving SEPP
type SecNegotiateReqData struct {
	// Fully Qualified Domain Name
	Sender string `json:"sender"`
	SupportedSecCapabilityList []SecurityCapability `json:"supportedSecCapabilityList"`
	Var3GppSbiTargetApiRootSupported *bool `json:"3GppSbiTargetApiRootSupported,omitempty"`
	PlmnIdList []PlmnId `json:"plmnIdList,omitempty"`
	SnpnIdList []PlmnIdNid `json:"snpnIdList,omitempty"`
	TargetPlmnId *PlmnId `json:"targetPlmnId,omitempty"`
	TargetSnpnId *PlmnIdNid `json:"targetSnpnId,omitempty"`
	IntendedUsagePurpose []IntendedN32Purpose `json:"intendedUsagePurpose,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// Fully Qualified Domain Name
	SenderN32fFqdn *string `json:"senderN32fFqdn,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	SenderN32fPort *int32 `json:"senderN32fPort,omitempty"`
}

// NewSecNegotiateReqData instantiates a new SecNegotiateReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecNegotiateReqData(sender string, supportedSecCapabilityList []SecurityCapability) *SecNegotiateReqData {
	this := SecNegotiateReqData{}
	this.Sender = sender
	this.SupportedSecCapabilityList = supportedSecCapabilityList
	var var3GppSbiTargetApiRootSupported bool = false
	this.Var3GppSbiTargetApiRootSupported = &var3GppSbiTargetApiRootSupported
	return &this
}

// NewSecNegotiateReqDataWithDefaults instantiates a new SecNegotiateReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecNegotiateReqDataWithDefaults() *SecNegotiateReqData {
	this := SecNegotiateReqData{}
	var var3GppSbiTargetApiRootSupported bool = false
	this.Var3GppSbiTargetApiRootSupported = &var3GppSbiTargetApiRootSupported
	return &this
}

// GetSender returns the Sender field value
func (o *SecNegotiateReqData) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *SecNegotiateReqData) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *SecNegotiateReqData) SetSender(v string) {
	o.Sender = v
}

// GetSupportedSecCapabilityList returns the SupportedSecCapabilityList field value
func (o *SecNegotiateReqData) GetSupportedSecCapabilityList() []SecurityCapability {
	if o == nil {
		var ret []SecurityCapability
		return ret
	}

	return o.SupportedSecCapabilityList
}

// GetSupportedSecCapabilityListOk returns a tuple with the SupportedSecCapabilityList field value
// and a boolean to check if the value has been set.
func (o *SecNegotiateReqData) GetSupportedSecCapabilityListOk() ([]SecurityCapability, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedSecCapabilityList, true
}

// SetSupportedSecCapabilityList sets field value
func (o *SecNegotiateReqData) SetSupportedSecCapabilityList(v []SecurityCapability) {
	o.SupportedSecCapabilityList = v
}

// GetVar3GppSbiTargetApiRootSupported returns the Var3GppSbiTargetApiRootSupported field value if set, zero value otherwise.
func (o *SecNegotiateReqData) GetVar3GppSbiTargetApiRootSupported() bool {
	if o == nil || IsNil(o.Var3GppSbiTargetApiRootSupported) {
		var ret bool
		return ret
	}
	return *o.Var3GppSbiTargetApiRootSupported
}

// GetVar3GppSbiTargetApiRootSupportedOk returns a tuple with the Var3GppSbiTargetApiRootSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecNegotiateReqData) GetVar3GppSbiTargetApiRootSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.Var3GppSbiTargetApiRootSupported) {
		return nil, false
	}
	return o.Var3GppSbiTargetApiRootSupported, true
}

// HasVar3GppSbiTargetApiRootSupported returns a boolean if a field has been set.
func (o *SecNegotiateReqData) HasVar3GppSbiTargetApiRootSupported() bool {
	if o != nil && !IsNil(o.Var3GppSbiTargetApiRootSupported) {
		return true
	}

	return false
}

// SetVar3GppSbiTargetApiRootSupported gets a reference to the given bool and assigns it to the Var3GppSbiTargetApiRootSupported field.
func (o *SecNegotiateReqData) SetVar3GppSbiTargetApiRootSupported(v bool) {
	o.Var3GppSbiTargetApiRootSupported = &v
}

// GetPlmnIdList returns the PlmnIdList field value if set, zero value otherwise.
func (o *SecNegotiateReqData) GetPlmnIdList() []PlmnId {
	if o == nil || IsNil(o.PlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.PlmnIdList
}

// GetPlmnIdListOk returns a tuple with the PlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecNegotiateReqData) GetPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.PlmnIdList) {
		return nil, false
	}
	return o.PlmnIdList, true
}

// HasPlmnIdList returns a boolean if a field has been set.
func (o *SecNegotiateReqData) HasPlmnIdList() bool {
	if o != nil && !IsNil(o.PlmnIdList) {
		return true
	}

	return false
}

// SetPlmnIdList gets a reference to the given []PlmnId and assigns it to the PlmnIdList field.
func (o *SecNegotiateReqData) SetPlmnIdList(v []PlmnId) {
	o.PlmnIdList = v
}

// GetSnpnIdList returns the SnpnIdList field value if set, zero value otherwise.
func (o *SecNegotiateReqData) GetSnpnIdList() []PlmnIdNid {
	if o == nil || IsNil(o.SnpnIdList) {
		var ret []PlmnIdNid
		return ret
	}
	return o.SnpnIdList
}

// GetSnpnIdListOk returns a tuple with the SnpnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecNegotiateReqData) GetSnpnIdListOk() ([]PlmnIdNid, bool) {
	if o == nil || IsNil(o.SnpnIdList) {
		return nil, false
	}
	return o.SnpnIdList, true
}

// HasSnpnIdList returns a boolean if a field has been set.
func (o *SecNegotiateReqData) HasSnpnIdList() bool {
	if o != nil && !IsNil(o.SnpnIdList) {
		return true
	}

	return false
}

// SetSnpnIdList gets a reference to the given []PlmnIdNid and assigns it to the SnpnIdList field.
func (o *SecNegotiateReqData) SetSnpnIdList(v []PlmnIdNid) {
	o.SnpnIdList = v
}

// GetTargetPlmnId returns the TargetPlmnId field value if set, zero value otherwise.
func (o *SecNegotiateReqData) GetTargetPlmnId() PlmnId {
	if o == nil || IsNil(o.TargetPlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.TargetPlmnId
}

// GetTargetPlmnIdOk returns a tuple with the TargetPlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecNegotiateReqData) GetTargetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.TargetPlmnId) {
		return nil, false
	}
	return o.TargetPlmnId, true
}

// HasTargetPlmnId returns a boolean if a field has been set.
func (o *SecNegotiateReqData) HasTargetPlmnId() bool {
	if o != nil && !IsNil(o.TargetPlmnId) {
		return true
	}

	return false
}

// SetTargetPlmnId gets a reference to the given PlmnId and assigns it to the TargetPlmnId field.
func (o *SecNegotiateReqData) SetTargetPlmnId(v PlmnId) {
	o.TargetPlmnId = &v
}

// GetTargetSnpnId returns the TargetSnpnId field value if set, zero value otherwise.
func (o *SecNegotiateReqData) GetTargetSnpnId() PlmnIdNid {
	if o == nil || IsNil(o.TargetSnpnId) {
		var ret PlmnIdNid
		return ret
	}
	return *o.TargetSnpnId
}

// GetTargetSnpnIdOk returns a tuple with the TargetSnpnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecNegotiateReqData) GetTargetSnpnIdOk() (*PlmnIdNid, bool) {
	if o == nil || IsNil(o.TargetSnpnId) {
		return nil, false
	}
	return o.TargetSnpnId, true
}

// HasTargetSnpnId returns a boolean if a field has been set.
func (o *SecNegotiateReqData) HasTargetSnpnId() bool {
	if o != nil && !IsNil(o.TargetSnpnId) {
		return true
	}

	return false
}

// SetTargetSnpnId gets a reference to the given PlmnIdNid and assigns it to the TargetSnpnId field.
func (o *SecNegotiateReqData) SetTargetSnpnId(v PlmnIdNid) {
	o.TargetSnpnId = &v
}

// GetIntendedUsagePurpose returns the IntendedUsagePurpose field value if set, zero value otherwise.
func (o *SecNegotiateReqData) GetIntendedUsagePurpose() []IntendedN32Purpose {
	if o == nil || IsNil(o.IntendedUsagePurpose) {
		var ret []IntendedN32Purpose
		return ret
	}
	return o.IntendedUsagePurpose
}

// GetIntendedUsagePurposeOk returns a tuple with the IntendedUsagePurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecNegotiateReqData) GetIntendedUsagePurposeOk() ([]IntendedN32Purpose, bool) {
	if o == nil || IsNil(o.IntendedUsagePurpose) {
		return nil, false
	}
	return o.IntendedUsagePurpose, true
}

// HasIntendedUsagePurpose returns a boolean if a field has been set.
func (o *SecNegotiateReqData) HasIntendedUsagePurpose() bool {
	if o != nil && !IsNil(o.IntendedUsagePurpose) {
		return true
	}

	return false
}

// SetIntendedUsagePurpose gets a reference to the given []IntendedN32Purpose and assigns it to the IntendedUsagePurpose field.
func (o *SecNegotiateReqData) SetIntendedUsagePurpose(v []IntendedN32Purpose) {
	o.IntendedUsagePurpose = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SecNegotiateReqData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecNegotiateReqData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SecNegotiateReqData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SecNegotiateReqData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetSenderN32fFqdn returns the SenderN32fFqdn field value if set, zero value otherwise.
func (o *SecNegotiateReqData) GetSenderN32fFqdn() string {
	if o == nil || IsNil(o.SenderN32fFqdn) {
		var ret string
		return ret
	}
	return *o.SenderN32fFqdn
}

// GetSenderN32fFqdnOk returns a tuple with the SenderN32fFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecNegotiateReqData) GetSenderN32fFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.SenderN32fFqdn) {
		return nil, false
	}
	return o.SenderN32fFqdn, true
}

// HasSenderN32fFqdn returns a boolean if a field has been set.
func (o *SecNegotiateReqData) HasSenderN32fFqdn() bool {
	if o != nil && !IsNil(o.SenderN32fFqdn) {
		return true
	}

	return false
}

// SetSenderN32fFqdn gets a reference to the given string and assigns it to the SenderN32fFqdn field.
func (o *SecNegotiateReqData) SetSenderN32fFqdn(v string) {
	o.SenderN32fFqdn = &v
}

// GetSenderN32fPort returns the SenderN32fPort field value if set, zero value otherwise.
func (o *SecNegotiateReqData) GetSenderN32fPort() int32 {
	if o == nil || IsNil(o.SenderN32fPort) {
		var ret int32
		return ret
	}
	return *o.SenderN32fPort
}

// GetSenderN32fPortOk returns a tuple with the SenderN32fPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecNegotiateReqData) GetSenderN32fPortOk() (*int32, bool) {
	if o == nil || IsNil(o.SenderN32fPort) {
		return nil, false
	}
	return o.SenderN32fPort, true
}

// HasSenderN32fPort returns a boolean if a field has been set.
func (o *SecNegotiateReqData) HasSenderN32fPort() bool {
	if o != nil && !IsNil(o.SenderN32fPort) {
		return true
	}

	return false
}

// SetSenderN32fPort gets a reference to the given int32 and assigns it to the SenderN32fPort field.
func (o *SecNegotiateReqData) SetSenderN32fPort(v int32) {
	o.SenderN32fPort = &v
}

func (o SecNegotiateReqData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecNegotiateReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sender"] = o.Sender
	toSerialize["supportedSecCapabilityList"] = o.SupportedSecCapabilityList
	if !IsNil(o.Var3GppSbiTargetApiRootSupported) {
		toSerialize["3GppSbiTargetApiRootSupported"] = o.Var3GppSbiTargetApiRootSupported
	}
	if !IsNil(o.PlmnIdList) {
		toSerialize["plmnIdList"] = o.PlmnIdList
	}
	if !IsNil(o.SnpnIdList) {
		toSerialize["snpnIdList"] = o.SnpnIdList
	}
	if !IsNil(o.TargetPlmnId) {
		toSerialize["targetPlmnId"] = o.TargetPlmnId
	}
	if !IsNil(o.TargetSnpnId) {
		toSerialize["targetSnpnId"] = o.TargetSnpnId
	}
	if !IsNil(o.IntendedUsagePurpose) {
		toSerialize["intendedUsagePurpose"] = o.IntendedUsagePurpose
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.SenderN32fFqdn) {
		toSerialize["senderN32fFqdn"] = o.SenderN32fFqdn
	}
	if !IsNil(o.SenderN32fPort) {
		toSerialize["senderN32fPort"] = o.SenderN32fPort
	}
	return toSerialize, nil
}

type NullableSecNegotiateReqData struct {
	value *SecNegotiateReqData
	isSet bool
}

func (v NullableSecNegotiateReqData) Get() *SecNegotiateReqData {
	return v.value
}

func (v *NullableSecNegotiateReqData) Set(val *SecNegotiateReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableSecNegotiateReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableSecNegotiateReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecNegotiateReqData(val *SecNegotiateReqData) *NullableSecNegotiateReqData {
	return &NullableSecNegotiateReqData{value: val, isSet: true}
}

func (v NullableSecNegotiateReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecNegotiateReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


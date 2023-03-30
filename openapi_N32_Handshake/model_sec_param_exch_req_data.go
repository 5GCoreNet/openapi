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

// checks if the SecParamExchReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecParamExchReqData{}

// SecParamExchReqData Request data structure for parameter exchange
type SecParamExchReqData struct {
	N32fContextId string `json:"n32fContextId"`
	JweCipherSuiteList []string `json:"jweCipherSuiteList,omitempty"`
	JwsCipherSuiteList []string `json:"jwsCipherSuiteList,omitempty"`
	ProtectionPolicyInfo *ProtectionPolicy `json:"protectionPolicyInfo,omitempty"`
	IpxProviderSecInfoList []IpxProviderSecInfo `json:"ipxProviderSecInfoList,omitempty"`
	// Fully Qualified Domain Name
	Sender *string `json:"sender,omitempty"`
}

// NewSecParamExchReqData instantiates a new SecParamExchReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecParamExchReqData(n32fContextId string) *SecParamExchReqData {
	this := SecParamExchReqData{}
	this.N32fContextId = n32fContextId
	return &this
}

// NewSecParamExchReqDataWithDefaults instantiates a new SecParamExchReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecParamExchReqDataWithDefaults() *SecParamExchReqData {
	this := SecParamExchReqData{}
	return &this
}

// GetN32fContextId returns the N32fContextId field value
func (o *SecParamExchReqData) GetN32fContextId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.N32fContextId
}

// GetN32fContextIdOk returns a tuple with the N32fContextId field value
// and a boolean to check if the value has been set.
func (o *SecParamExchReqData) GetN32fContextIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N32fContextId, true
}

// SetN32fContextId sets field value
func (o *SecParamExchReqData) SetN32fContextId(v string) {
	o.N32fContextId = v
}

// GetJweCipherSuiteList returns the JweCipherSuiteList field value if set, zero value otherwise.
func (o *SecParamExchReqData) GetJweCipherSuiteList() []string {
	if o == nil || IsNil(o.JweCipherSuiteList) {
		var ret []string
		return ret
	}
	return o.JweCipherSuiteList
}

// GetJweCipherSuiteListOk returns a tuple with the JweCipherSuiteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecParamExchReqData) GetJweCipherSuiteListOk() ([]string, bool) {
	if o == nil || IsNil(o.JweCipherSuiteList) {
		return nil, false
	}
	return o.JweCipherSuiteList, true
}

// HasJweCipherSuiteList returns a boolean if a field has been set.
func (o *SecParamExchReqData) HasJweCipherSuiteList() bool {
	if o != nil && !IsNil(o.JweCipherSuiteList) {
		return true
	}

	return false
}

// SetJweCipherSuiteList gets a reference to the given []string and assigns it to the JweCipherSuiteList field.
func (o *SecParamExchReqData) SetJweCipherSuiteList(v []string) {
	o.JweCipherSuiteList = v
}

// GetJwsCipherSuiteList returns the JwsCipherSuiteList field value if set, zero value otherwise.
func (o *SecParamExchReqData) GetJwsCipherSuiteList() []string {
	if o == nil || IsNil(o.JwsCipherSuiteList) {
		var ret []string
		return ret
	}
	return o.JwsCipherSuiteList
}

// GetJwsCipherSuiteListOk returns a tuple with the JwsCipherSuiteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecParamExchReqData) GetJwsCipherSuiteListOk() ([]string, bool) {
	if o == nil || IsNil(o.JwsCipherSuiteList) {
		return nil, false
	}
	return o.JwsCipherSuiteList, true
}

// HasJwsCipherSuiteList returns a boolean if a field has been set.
func (o *SecParamExchReqData) HasJwsCipherSuiteList() bool {
	if o != nil && !IsNil(o.JwsCipherSuiteList) {
		return true
	}

	return false
}

// SetJwsCipherSuiteList gets a reference to the given []string and assigns it to the JwsCipherSuiteList field.
func (o *SecParamExchReqData) SetJwsCipherSuiteList(v []string) {
	o.JwsCipherSuiteList = v
}

// GetProtectionPolicyInfo returns the ProtectionPolicyInfo field value if set, zero value otherwise.
func (o *SecParamExchReqData) GetProtectionPolicyInfo() ProtectionPolicy {
	if o == nil || IsNil(o.ProtectionPolicyInfo) {
		var ret ProtectionPolicy
		return ret
	}
	return *o.ProtectionPolicyInfo
}

// GetProtectionPolicyInfoOk returns a tuple with the ProtectionPolicyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecParamExchReqData) GetProtectionPolicyInfoOk() (*ProtectionPolicy, bool) {
	if o == nil || IsNil(o.ProtectionPolicyInfo) {
		return nil, false
	}
	return o.ProtectionPolicyInfo, true
}

// HasProtectionPolicyInfo returns a boolean if a field has been set.
func (o *SecParamExchReqData) HasProtectionPolicyInfo() bool {
	if o != nil && !IsNil(o.ProtectionPolicyInfo) {
		return true
	}

	return false
}

// SetProtectionPolicyInfo gets a reference to the given ProtectionPolicy and assigns it to the ProtectionPolicyInfo field.
func (o *SecParamExchReqData) SetProtectionPolicyInfo(v ProtectionPolicy) {
	o.ProtectionPolicyInfo = &v
}

// GetIpxProviderSecInfoList returns the IpxProviderSecInfoList field value if set, zero value otherwise.
func (o *SecParamExchReqData) GetIpxProviderSecInfoList() []IpxProviderSecInfo {
	if o == nil || IsNil(o.IpxProviderSecInfoList) {
		var ret []IpxProviderSecInfo
		return ret
	}
	return o.IpxProviderSecInfoList
}

// GetIpxProviderSecInfoListOk returns a tuple with the IpxProviderSecInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecParamExchReqData) GetIpxProviderSecInfoListOk() ([]IpxProviderSecInfo, bool) {
	if o == nil || IsNil(o.IpxProviderSecInfoList) {
		return nil, false
	}
	return o.IpxProviderSecInfoList, true
}

// HasIpxProviderSecInfoList returns a boolean if a field has been set.
func (o *SecParamExchReqData) HasIpxProviderSecInfoList() bool {
	if o != nil && !IsNil(o.IpxProviderSecInfoList) {
		return true
	}

	return false
}

// SetIpxProviderSecInfoList gets a reference to the given []IpxProviderSecInfo and assigns it to the IpxProviderSecInfoList field.
func (o *SecParamExchReqData) SetIpxProviderSecInfoList(v []IpxProviderSecInfo) {
	o.IpxProviderSecInfoList = v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *SecParamExchReqData) GetSender() string {
	if o == nil || IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecParamExchReqData) GetSenderOk() (*string, bool) {
	if o == nil || IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *SecParamExchReqData) HasSender() bool {
	if o != nil && !IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *SecParamExchReqData) SetSender(v string) {
	o.Sender = &v
}

func (o SecParamExchReqData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecParamExchReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["n32fContextId"] = o.N32fContextId
	if !IsNil(o.JweCipherSuiteList) {
		toSerialize["jweCipherSuiteList"] = o.JweCipherSuiteList
	}
	if !IsNil(o.JwsCipherSuiteList) {
		toSerialize["jwsCipherSuiteList"] = o.JwsCipherSuiteList
	}
	if !IsNil(o.ProtectionPolicyInfo) {
		toSerialize["protectionPolicyInfo"] = o.ProtectionPolicyInfo
	}
	if !IsNil(o.IpxProviderSecInfoList) {
		toSerialize["ipxProviderSecInfoList"] = o.IpxProviderSecInfoList
	}
	if !IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	return toSerialize, nil
}

type NullableSecParamExchReqData struct {
	value *SecParamExchReqData
	isSet bool
}

func (v NullableSecParamExchReqData) Get() *SecParamExchReqData {
	return v.value
}

func (v *NullableSecParamExchReqData) Set(val *SecParamExchReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableSecParamExchReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableSecParamExchReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecParamExchReqData(val *SecParamExchReqData) *NullableSecParamExchReqData {
	return &NullableSecParamExchReqData{value: val, isSet: true}
}

func (v NullableSecParamExchReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecParamExchReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


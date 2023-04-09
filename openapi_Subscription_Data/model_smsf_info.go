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

// checks if the SmsfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsfInfo{}

// SmsfInfo struct for SmsfInfo
type SmsfInfo struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SmsfInstanceId string `json:"smsfInstanceId"`
	PlmnId         PlmnId `json:"plmnId"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
	SmsfSetId *string `json:"smsfSetId,omitempty"`
}

// NewSmsfInfo instantiates a new SmsfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsfInfo(smsfInstanceId string, plmnId PlmnId) *SmsfInfo {
	this := SmsfInfo{}
	this.SmsfInstanceId = smsfInstanceId
	this.PlmnId = plmnId
	return &this
}

// NewSmsfInfoWithDefaults instantiates a new SmsfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsfInfoWithDefaults() *SmsfInfo {
	this := SmsfInfo{}
	return &this
}

// GetSmsfInstanceId returns the SmsfInstanceId field value
func (o *SmsfInfo) GetSmsfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmsfInstanceId
}

// GetSmsfInstanceIdOk returns a tuple with the SmsfInstanceId field value
// and a boolean to check if the value has been set.
func (o *SmsfInfo) GetSmsfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmsfInstanceId, true
}

// SetSmsfInstanceId sets field value
func (o *SmsfInfo) SetSmsfInstanceId(v string) {
	o.SmsfInstanceId = v
}

// GetPlmnId returns the PlmnId field value
func (o *SmsfInfo) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *SmsfInfo) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *SmsfInfo) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetSmsfSetId returns the SmsfSetId field value if set, zero value otherwise.
func (o *SmsfInfo) GetSmsfSetId() string {
	if o == nil || IsNil(o.SmsfSetId) {
		var ret string
		return ret
	}
	return *o.SmsfSetId
}

// GetSmsfSetIdOk returns a tuple with the SmsfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfInfo) GetSmsfSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.SmsfSetId) {
		return nil, false
	}
	return o.SmsfSetId, true
}

// HasSmsfSetId returns a boolean if a field has been set.
func (o *SmsfInfo) HasSmsfSetId() bool {
	if o != nil && !IsNil(o.SmsfSetId) {
		return true
	}

	return false
}

// SetSmsfSetId gets a reference to the given string and assigns it to the SmsfSetId field.
func (o *SmsfInfo) SetSmsfSetId(v string) {
	o.SmsfSetId = &v
}

func (o SmsfInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["smsfInstanceId"] = o.SmsfInstanceId
	toSerialize["plmnId"] = o.PlmnId
	if !IsNil(o.SmsfSetId) {
		toSerialize["smsfSetId"] = o.SmsfSetId
	}
	return toSerialize, nil
}

type NullableSmsfInfo struct {
	value *SmsfInfo
	isSet bool
}

func (v NullableSmsfInfo) Get() *SmsfInfo {
	return v.value
}

func (v *NullableSmsfInfo) Set(val *SmsfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsfInfo(val *SmsfInfo) *NullableSmsfInfo {
	return &NullableSmsfInfo{value: val, isSet: true}
}

func (v NullableSmsfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

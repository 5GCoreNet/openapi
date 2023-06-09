/*
Neasdf_BaselineDNSPattern

EASDF Baseline DNS Pattern Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Neasdf_BaselineDNSPattern

import (
	"encoding/json"
)

// checks if the VarNfId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VarNfId{}

// VarNfId SMF or SMF Set Id or Set Id part in NF Set Id
type VarNfId struct {
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
	SmfSetId *string `json:"smfSetId,omitempty"`
	SetId    *string `json:"setId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SmfInstanceId *string `json:"smfInstanceId,omitempty"`
}

// NewVarNfId instantiates a new VarNfId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVarNfId() *VarNfId {
	this := VarNfId{}
	return &this
}

// NewVarNfIdWithDefaults instantiates a new VarNfId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVarNfIdWithDefaults() *VarNfId {
	this := VarNfId{}
	return &this
}

// GetSmfSetId returns the SmfSetId field value if set, zero value otherwise.
func (o *VarNfId) GetSmfSetId() string {
	if o == nil || IsNil(o.SmfSetId) {
		var ret string
		return ret
	}
	return *o.SmfSetId
}

// GetSmfSetIdOk returns a tuple with the SmfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VarNfId) GetSmfSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.SmfSetId) {
		return nil, false
	}
	return o.SmfSetId, true
}

// HasSmfSetId returns a boolean if a field has been set.
func (o *VarNfId) HasSmfSetId() bool {
	if o != nil && !IsNil(o.SmfSetId) {
		return true
	}

	return false
}

// SetSmfSetId gets a reference to the given string and assigns it to the SmfSetId field.
func (o *VarNfId) SetSmfSetId(v string) {
	o.SmfSetId = &v
}

// GetSetId returns the SetId field value if set, zero value otherwise.
func (o *VarNfId) GetSetId() string {
	if o == nil || IsNil(o.SetId) {
		var ret string
		return ret
	}
	return *o.SetId
}

// GetSetIdOk returns a tuple with the SetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VarNfId) GetSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.SetId) {
		return nil, false
	}
	return o.SetId, true
}

// HasSetId returns a boolean if a field has been set.
func (o *VarNfId) HasSetId() bool {
	if o != nil && !IsNil(o.SetId) {
		return true
	}

	return false
}

// SetSetId gets a reference to the given string and assigns it to the SetId field.
func (o *VarNfId) SetSetId(v string) {
	o.SetId = &v
}

// GetSmfInstanceId returns the SmfInstanceId field value if set, zero value otherwise.
func (o *VarNfId) GetSmfInstanceId() string {
	if o == nil || IsNil(o.SmfInstanceId) {
		var ret string
		return ret
	}
	return *o.SmfInstanceId
}

// GetSmfInstanceIdOk returns a tuple with the SmfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VarNfId) GetSmfInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SmfInstanceId) {
		return nil, false
	}
	return o.SmfInstanceId, true
}

// HasSmfInstanceId returns a boolean if a field has been set.
func (o *VarNfId) HasSmfInstanceId() bool {
	if o != nil && !IsNil(o.SmfInstanceId) {
		return true
	}

	return false
}

// SetSmfInstanceId gets a reference to the given string and assigns it to the SmfInstanceId field.
func (o *VarNfId) SetSmfInstanceId(v string) {
	o.SmfInstanceId = &v
}

func (o VarNfId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VarNfId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SmfSetId) {
		toSerialize["smfSetId"] = o.SmfSetId
	}
	if !IsNil(o.SetId) {
		toSerialize["setId"] = o.SetId
	}
	if !IsNil(o.SmfInstanceId) {
		toSerialize["smfInstanceId"] = o.SmfInstanceId
	}
	return toSerialize, nil
}

type NullableVarNfId struct {
	value *VarNfId
	isSet bool
}

func (v NullableVarNfId) Get() *VarNfId {
	return v.value
}

func (v *NullableVarNfId) Set(val *VarNfId) {
	v.value = val
	v.isSet = true
}

func (v NullableVarNfId) IsSet() bool {
	return v.isSet
}

func (v *NullableVarNfId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVarNfId(val *VarNfId) *NullableVarNfId {
	return &NullableVarNfId{value: val, isSet: true}
}

func (v NullableVarNfId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVarNfId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

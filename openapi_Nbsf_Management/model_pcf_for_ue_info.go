/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.4.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsf_Management

import (
	"encoding/json"
)

// checks if the PcfForUeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcfForUeInfo{}

// PcfForUeInfo Contains the information of the PCF for a UE.
type PcfForUeInfo struct {
	// Fully Qualified Domain Name
	PcfFqdn *string `json:"pcfFqdn,omitempty"`
	// IP end points of the PCF hosting the Npcf_AmPolicyAuthorization service.
	PcfIpEndPoints []IpEndPoint `json:"pcfIpEndPoints,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	PcfId *string `json:"pcfId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	PcfSetId *string `json:"pcfSetId,omitempty"`
	BindLevel *BindingLevel `json:"bindLevel,omitempty"`
}

// NewPcfForUeInfo instantiates a new PcfForUeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfForUeInfo() *PcfForUeInfo {
	this := PcfForUeInfo{}
	return &this
}

// NewPcfForUeInfoWithDefaults instantiates a new PcfForUeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfForUeInfoWithDefaults() *PcfForUeInfo {
	this := PcfForUeInfo{}
	return &this
}

// GetPcfFqdn returns the PcfFqdn field value if set, zero value otherwise.
func (o *PcfForUeInfo) GetPcfFqdn() string {
	if o == nil || isNil(o.PcfFqdn) {
		var ret string
		return ret
	}
	return *o.PcfFqdn
}

// GetPcfFqdnOk returns a tuple with the PcfFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfForUeInfo) GetPcfFqdnOk() (*string, bool) {
	if o == nil || isNil(o.PcfFqdn) {
		return nil, false
	}
	return o.PcfFqdn, true
}

// HasPcfFqdn returns a boolean if a field has been set.
func (o *PcfForUeInfo) HasPcfFqdn() bool {
	if o != nil && !isNil(o.PcfFqdn) {
		return true
	}

	return false
}

// SetPcfFqdn gets a reference to the given string and assigns it to the PcfFqdn field.
func (o *PcfForUeInfo) SetPcfFqdn(v string) {
	o.PcfFqdn = &v
}

// GetPcfIpEndPoints returns the PcfIpEndPoints field value if set, zero value otherwise.
func (o *PcfForUeInfo) GetPcfIpEndPoints() []IpEndPoint {
	if o == nil || isNil(o.PcfIpEndPoints) {
		var ret []IpEndPoint
		return ret
	}
	return o.PcfIpEndPoints
}

// GetPcfIpEndPointsOk returns a tuple with the PcfIpEndPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfForUeInfo) GetPcfIpEndPointsOk() ([]IpEndPoint, bool) {
	if o == nil || isNil(o.PcfIpEndPoints) {
		return nil, false
	}
	return o.PcfIpEndPoints, true
}

// HasPcfIpEndPoints returns a boolean if a field has been set.
func (o *PcfForUeInfo) HasPcfIpEndPoints() bool {
	if o != nil && !isNil(o.PcfIpEndPoints) {
		return true
	}

	return false
}

// SetPcfIpEndPoints gets a reference to the given []IpEndPoint and assigns it to the PcfIpEndPoints field.
func (o *PcfForUeInfo) SetPcfIpEndPoints(v []IpEndPoint) {
	o.PcfIpEndPoints = v
}

// GetPcfId returns the PcfId field value if set, zero value otherwise.
func (o *PcfForUeInfo) GetPcfId() string {
	if o == nil || isNil(o.PcfId) {
		var ret string
		return ret
	}
	return *o.PcfId
}

// GetPcfIdOk returns a tuple with the PcfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfForUeInfo) GetPcfIdOk() (*string, bool) {
	if o == nil || isNil(o.PcfId) {
		return nil, false
	}
	return o.PcfId, true
}

// HasPcfId returns a boolean if a field has been set.
func (o *PcfForUeInfo) HasPcfId() bool {
	if o != nil && !isNil(o.PcfId) {
		return true
	}

	return false
}

// SetPcfId gets a reference to the given string and assigns it to the PcfId field.
func (o *PcfForUeInfo) SetPcfId(v string) {
	o.PcfId = &v
}

// GetPcfSetId returns the PcfSetId field value if set, zero value otherwise.
func (o *PcfForUeInfo) GetPcfSetId() string {
	if o == nil || isNil(o.PcfSetId) {
		var ret string
		return ret
	}
	return *o.PcfSetId
}

// GetPcfSetIdOk returns a tuple with the PcfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfForUeInfo) GetPcfSetIdOk() (*string, bool) {
	if o == nil || isNil(o.PcfSetId) {
		return nil, false
	}
	return o.PcfSetId, true
}

// HasPcfSetId returns a boolean if a field has been set.
func (o *PcfForUeInfo) HasPcfSetId() bool {
	if o != nil && !isNil(o.PcfSetId) {
		return true
	}

	return false
}

// SetPcfSetId gets a reference to the given string and assigns it to the PcfSetId field.
func (o *PcfForUeInfo) SetPcfSetId(v string) {
	o.PcfSetId = &v
}

// GetBindLevel returns the BindLevel field value if set, zero value otherwise.
func (o *PcfForUeInfo) GetBindLevel() BindingLevel {
	if o == nil || isNil(o.BindLevel) {
		var ret BindingLevel
		return ret
	}
	return *o.BindLevel
}

// GetBindLevelOk returns a tuple with the BindLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfForUeInfo) GetBindLevelOk() (*BindingLevel, bool) {
	if o == nil || isNil(o.BindLevel) {
		return nil, false
	}
	return o.BindLevel, true
}

// HasBindLevel returns a boolean if a field has been set.
func (o *PcfForUeInfo) HasBindLevel() bool {
	if o != nil && !isNil(o.BindLevel) {
		return true
	}

	return false
}

// SetBindLevel gets a reference to the given BindingLevel and assigns it to the BindLevel field.
func (o *PcfForUeInfo) SetBindLevel(v BindingLevel) {
	o.BindLevel = &v
}

func (o PcfForUeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcfForUeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PcfFqdn) {
		toSerialize["pcfFqdn"] = o.PcfFqdn
	}
	if !isNil(o.PcfIpEndPoints) {
		toSerialize["pcfIpEndPoints"] = o.PcfIpEndPoints
	}
	if !isNil(o.PcfId) {
		toSerialize["pcfId"] = o.PcfId
	}
	if !isNil(o.PcfSetId) {
		toSerialize["pcfSetId"] = o.PcfSetId
	}
	if !isNil(o.BindLevel) {
		toSerialize["bindLevel"] = o.BindLevel
	}
	return toSerialize, nil
}

type NullablePcfForUeInfo struct {
	value *PcfForUeInfo
	isSet bool
}

func (v NullablePcfForUeInfo) Get() *PcfForUeInfo {
	return v.value
}

func (v *NullablePcfForUeInfo) Set(val *PcfForUeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfForUeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfForUeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfForUeInfo(val *PcfForUeInfo) *NullablePcfForUeInfo {
	return &NullablePcfForUeInfo{value: val, isSet: true}
}

func (v NullablePcfForUeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfForUeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



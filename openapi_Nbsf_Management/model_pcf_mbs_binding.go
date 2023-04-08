/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.4.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsf_Management

import (
	"encoding/json"
	"time"
)

// checks if the PcfMbsBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcfMbsBinding{}

// PcfMbsBinding Represents an Individual PCF for an MBS Session binding.
type PcfMbsBinding struct {
	MbsSessionId MbsSessionId `json:"mbsSessionId"`
	// Fully Qualified Domain Name
	PcfFqdn *string `json:"pcfFqdn,omitempty"`
	PcfIpEndPoints []IpEndPoint `json:"pcfIpEndPoints,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	PcfId *string `json:"pcfId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	PcfSetId *string `json:"pcfSetId,omitempty"`
	BindLevel *BindingLevel `json:"bindLevel,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime *time.Time `json:"recoveryTime,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewPcfMbsBinding instantiates a new PcfMbsBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfMbsBinding(mbsSessionId MbsSessionId) *PcfMbsBinding {
	this := PcfMbsBinding{}
	this.MbsSessionId = mbsSessionId
	return &this
}

// NewPcfMbsBindingWithDefaults instantiates a new PcfMbsBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfMbsBindingWithDefaults() *PcfMbsBinding {
	this := PcfMbsBinding{}
	return &this
}

// GetMbsSessionId returns the MbsSessionId field value
func (o *PcfMbsBinding) GetMbsSessionId() MbsSessionId {
	if o == nil {
		var ret MbsSessionId
		return ret
	}

	return o.MbsSessionId
}

// GetMbsSessionIdOk returns a tuple with the MbsSessionId field value
// and a boolean to check if the value has been set.
func (o *PcfMbsBinding) GetMbsSessionIdOk() (*MbsSessionId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MbsSessionId, true
}

// SetMbsSessionId sets field value
func (o *PcfMbsBinding) SetMbsSessionId(v MbsSessionId) {
	o.MbsSessionId = v
}

// GetPcfFqdn returns the PcfFqdn field value if set, zero value otherwise.
func (o *PcfMbsBinding) GetPcfFqdn() string {
	if o == nil || isNil(o.PcfFqdn) {
		var ret string
		return ret
	}
	return *o.PcfFqdn
}

// GetPcfFqdnOk returns a tuple with the PcfFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfMbsBinding) GetPcfFqdnOk() (*string, bool) {
	if o == nil || isNil(o.PcfFqdn) {
		return nil, false
	}
	return o.PcfFqdn, true
}

// HasPcfFqdn returns a boolean if a field has been set.
func (o *PcfMbsBinding) HasPcfFqdn() bool {
	if o != nil && !isNil(o.PcfFqdn) {
		return true
	}

	return false
}

// SetPcfFqdn gets a reference to the given string and assigns it to the PcfFqdn field.
func (o *PcfMbsBinding) SetPcfFqdn(v string) {
	o.PcfFqdn = &v
}

// GetPcfIpEndPoints returns the PcfIpEndPoints field value if set, zero value otherwise.
func (o *PcfMbsBinding) GetPcfIpEndPoints() []IpEndPoint {
	if o == nil || isNil(o.PcfIpEndPoints) {
		var ret []IpEndPoint
		return ret
	}
	return o.PcfIpEndPoints
}

// GetPcfIpEndPointsOk returns a tuple with the PcfIpEndPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfMbsBinding) GetPcfIpEndPointsOk() ([]IpEndPoint, bool) {
	if o == nil || isNil(o.PcfIpEndPoints) {
		return nil, false
	}
	return o.PcfIpEndPoints, true
}

// HasPcfIpEndPoints returns a boolean if a field has been set.
func (o *PcfMbsBinding) HasPcfIpEndPoints() bool {
	if o != nil && !isNil(o.PcfIpEndPoints) {
		return true
	}

	return false
}

// SetPcfIpEndPoints gets a reference to the given []IpEndPoint and assigns it to the PcfIpEndPoints field.
func (o *PcfMbsBinding) SetPcfIpEndPoints(v []IpEndPoint) {
	o.PcfIpEndPoints = v
}

// GetPcfId returns the PcfId field value if set, zero value otherwise.
func (o *PcfMbsBinding) GetPcfId() string {
	if o == nil || isNil(o.PcfId) {
		var ret string
		return ret
	}
	return *o.PcfId
}

// GetPcfIdOk returns a tuple with the PcfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfMbsBinding) GetPcfIdOk() (*string, bool) {
	if o == nil || isNil(o.PcfId) {
		return nil, false
	}
	return o.PcfId, true
}

// HasPcfId returns a boolean if a field has been set.
func (o *PcfMbsBinding) HasPcfId() bool {
	if o != nil && !isNil(o.PcfId) {
		return true
	}

	return false
}

// SetPcfId gets a reference to the given string and assigns it to the PcfId field.
func (o *PcfMbsBinding) SetPcfId(v string) {
	o.PcfId = &v
}

// GetPcfSetId returns the PcfSetId field value if set, zero value otherwise.
func (o *PcfMbsBinding) GetPcfSetId() string {
	if o == nil || isNil(o.PcfSetId) {
		var ret string
		return ret
	}
	return *o.PcfSetId
}

// GetPcfSetIdOk returns a tuple with the PcfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfMbsBinding) GetPcfSetIdOk() (*string, bool) {
	if o == nil || isNil(o.PcfSetId) {
		return nil, false
	}
	return o.PcfSetId, true
}

// HasPcfSetId returns a boolean if a field has been set.
func (o *PcfMbsBinding) HasPcfSetId() bool {
	if o != nil && !isNil(o.PcfSetId) {
		return true
	}

	return false
}

// SetPcfSetId gets a reference to the given string and assigns it to the PcfSetId field.
func (o *PcfMbsBinding) SetPcfSetId(v string) {
	o.PcfSetId = &v
}

// GetBindLevel returns the BindLevel field value if set, zero value otherwise.
func (o *PcfMbsBinding) GetBindLevel() BindingLevel {
	if o == nil || isNil(o.BindLevel) {
		var ret BindingLevel
		return ret
	}
	return *o.BindLevel
}

// GetBindLevelOk returns a tuple with the BindLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfMbsBinding) GetBindLevelOk() (*BindingLevel, bool) {
	if o == nil || isNil(o.BindLevel) {
		return nil, false
	}
	return o.BindLevel, true
}

// HasBindLevel returns a boolean if a field has been set.
func (o *PcfMbsBinding) HasBindLevel() bool {
	if o != nil && !isNil(o.BindLevel) {
		return true
	}

	return false
}

// SetBindLevel gets a reference to the given BindingLevel and assigns it to the BindLevel field.
func (o *PcfMbsBinding) SetBindLevel(v BindingLevel) {
	o.BindLevel = &v
}

// GetRecoveryTime returns the RecoveryTime field value if set, zero value otherwise.
func (o *PcfMbsBinding) GetRecoveryTime() time.Time {
	if o == nil || isNil(o.RecoveryTime) {
		var ret time.Time
		return ret
	}
	return *o.RecoveryTime
}

// GetRecoveryTimeOk returns a tuple with the RecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfMbsBinding) GetRecoveryTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.RecoveryTime) {
		return nil, false
	}
	return o.RecoveryTime, true
}

// HasRecoveryTime returns a boolean if a field has been set.
func (o *PcfMbsBinding) HasRecoveryTime() bool {
	if o != nil && !isNil(o.RecoveryTime) {
		return true
	}

	return false
}

// SetRecoveryTime gets a reference to the given time.Time and assigns it to the RecoveryTime field.
func (o *PcfMbsBinding) SetRecoveryTime(v time.Time) {
	o.RecoveryTime = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *PcfMbsBinding) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfMbsBinding) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *PcfMbsBinding) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *PcfMbsBinding) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o PcfMbsBinding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcfMbsBinding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mbsSessionId"] = o.MbsSessionId
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
	if !isNil(o.RecoveryTime) {
		toSerialize["recoveryTime"] = o.RecoveryTime
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullablePcfMbsBinding struct {
	value *PcfMbsBinding
	isSet bool
}

func (v NullablePcfMbsBinding) Get() *PcfMbsBinding {
	return v.value
}

func (v *NullablePcfMbsBinding) Set(val *PcfMbsBinding) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfMbsBinding) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfMbsBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfMbsBinding(val *PcfMbsBinding) *NullablePcfMbsBinding {
	return &NullablePcfMbsBinding{value: val, isSet: true}
}

func (v NullablePcfMbsBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfMbsBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



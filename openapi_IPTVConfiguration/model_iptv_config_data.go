/*
3gpp-iptvconfiguration

API for IPTV configuration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IPTVConfiguration

import (
	"encoding/json"
)

// checks if the IptvConfigData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IptvConfigData{}

// IptvConfigData Represents an individual IPTV Configuration resource.
type IptvConfigData struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self *string `json:"self,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExterGroupId *string `json:"exterGroupId,omitempty"`
	AfAppId string `json:"afAppId"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// Identifies a list of multicast address access control information. Any string value can be used as a key of the map. 
	MultiAccCtrls map[string]MulticastAccessControl `json:"multiAccCtrls"`
	// String uniquely identifying MTC provider information.
	MtcProviderId *string `json:"mtcProviderId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat"`
}

// NewIptvConfigData instantiates a new IptvConfigData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIptvConfigData(afAppId string, multiAccCtrls map[string]MulticastAccessControl, suppFeat string) *IptvConfigData {
	this := IptvConfigData{}
	this.AfAppId = afAppId
	this.MultiAccCtrls = multiAccCtrls
	this.SuppFeat = suppFeat
	return &this
}

// NewIptvConfigDataWithDefaults instantiates a new IptvConfigData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIptvConfigDataWithDefaults() *IptvConfigData {
	this := IptvConfigData{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *IptvConfigData) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IptvConfigData) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *IptvConfigData) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *IptvConfigData) SetSelf(v string) {
	o.Self = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *IptvConfigData) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IptvConfigData) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *IptvConfigData) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *IptvConfigData) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetExterGroupId returns the ExterGroupId field value if set, zero value otherwise.
func (o *IptvConfigData) GetExterGroupId() string {
	if o == nil || IsNil(o.ExterGroupId) {
		var ret string
		return ret
	}
	return *o.ExterGroupId
}

// GetExterGroupIdOk returns a tuple with the ExterGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IptvConfigData) GetExterGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExterGroupId) {
		return nil, false
	}
	return o.ExterGroupId, true
}

// HasExterGroupId returns a boolean if a field has been set.
func (o *IptvConfigData) HasExterGroupId() bool {
	if o != nil && !IsNil(o.ExterGroupId) {
		return true
	}

	return false
}

// SetExterGroupId gets a reference to the given string and assigns it to the ExterGroupId field.
func (o *IptvConfigData) SetExterGroupId(v string) {
	o.ExterGroupId = &v
}

// GetAfAppId returns the AfAppId field value
func (o *IptvConfigData) GetAfAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfAppId
}

// GetAfAppIdOk returns a tuple with the AfAppId field value
// and a boolean to check if the value has been set.
func (o *IptvConfigData) GetAfAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfAppId, true
}

// SetAfAppId sets field value
func (o *IptvConfigData) SetAfAppId(v string) {
	o.AfAppId = v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *IptvConfigData) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IptvConfigData) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *IptvConfigData) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *IptvConfigData) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *IptvConfigData) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IptvConfigData) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *IptvConfigData) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *IptvConfigData) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetMultiAccCtrls returns the MultiAccCtrls field value
func (o *IptvConfigData) GetMultiAccCtrls() map[string]MulticastAccessControl {
	if o == nil {
		var ret map[string]MulticastAccessControl
		return ret
	}

	return o.MultiAccCtrls
}

// GetMultiAccCtrlsOk returns a tuple with the MultiAccCtrls field value
// and a boolean to check if the value has been set.
func (o *IptvConfigData) GetMultiAccCtrlsOk() (*map[string]MulticastAccessControl, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultiAccCtrls, true
}

// SetMultiAccCtrls sets field value
func (o *IptvConfigData) SetMultiAccCtrls(v map[string]MulticastAccessControl) {
	o.MultiAccCtrls = v
}

// GetMtcProviderId returns the MtcProviderId field value if set, zero value otherwise.
func (o *IptvConfigData) GetMtcProviderId() string {
	if o == nil || IsNil(o.MtcProviderId) {
		var ret string
		return ret
	}
	return *o.MtcProviderId
}

// GetMtcProviderIdOk returns a tuple with the MtcProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IptvConfigData) GetMtcProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.MtcProviderId) {
		return nil, false
	}
	return o.MtcProviderId, true
}

// HasMtcProviderId returns a boolean if a field has been set.
func (o *IptvConfigData) HasMtcProviderId() bool {
	if o != nil && !IsNil(o.MtcProviderId) {
		return true
	}

	return false
}

// SetMtcProviderId gets a reference to the given string and assigns it to the MtcProviderId field.
func (o *IptvConfigData) SetMtcProviderId(v string) {
	o.MtcProviderId = &v
}

// GetSuppFeat returns the SuppFeat field value
func (o *IptvConfigData) GetSuppFeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value
// and a boolean to check if the value has been set.
func (o *IptvConfigData) GetSuppFeatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuppFeat, true
}

// SetSuppFeat sets field value
func (o *IptvConfigData) SetSuppFeat(v string) {
	o.SuppFeat = v
}

func (o IptvConfigData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IptvConfigData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.ExterGroupId) {
		toSerialize["exterGroupId"] = o.ExterGroupId
	}
	toSerialize["afAppId"] = o.AfAppId
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	toSerialize["multiAccCtrls"] = o.MultiAccCtrls
	if !IsNil(o.MtcProviderId) {
		toSerialize["mtcProviderId"] = o.MtcProviderId
	}
	toSerialize["suppFeat"] = o.SuppFeat
	return toSerialize, nil
}

type NullableIptvConfigData struct {
	value *IptvConfigData
	isSet bool
}

func (v NullableIptvConfigData) Get() *IptvConfigData {
	return v.value
}

func (v *NullableIptvConfigData) Set(val *IptvConfigData) {
	v.value = val
	v.isSet = true
}

func (v NullableIptvConfigData) IsSet() bool {
	return v.isSet
}

func (v *NullableIptvConfigData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIptvConfigData(val *IptvConfigData) *NullableIptvConfigData {
	return &NullableIptvConfigData{value: val, isSet: true}
}

func (v NullableIptvConfigData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIptvConfigData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



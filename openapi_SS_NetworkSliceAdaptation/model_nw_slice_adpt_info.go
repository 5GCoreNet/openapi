/*
SS_NetworkSliceAdaptation

API for SEAL Network Slice Adaptation.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkSliceAdaptation

import (
	"encoding/json"
)

// checks if the NwSliceAdptInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NwSliceAdptInfo{}

// NwSliceAdptInfo Represents the information associated with requested network slice adaptation with the underlying network.
type NwSliceAdptInfo struct {
	ValServiceId string   `json:"valServiceId"`
	ValTgtUeIds  []string `json:"valTgtUeIds"`
	Snssai       *Snssai  `json:"snssai,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn *string `json:"dnn,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewNwSliceAdptInfo instantiates a new NwSliceAdptInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwSliceAdptInfo(valServiceId string, valTgtUeIds []string) *NwSliceAdptInfo {
	this := NwSliceAdptInfo{}
	this.ValServiceId = valServiceId
	this.ValTgtUeIds = valTgtUeIds
	return &this
}

// NewNwSliceAdptInfoWithDefaults instantiates a new NwSliceAdptInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwSliceAdptInfoWithDefaults() *NwSliceAdptInfo {
	this := NwSliceAdptInfo{}
	return &this
}

// GetValServiceId returns the ValServiceId field value
func (o *NwSliceAdptInfo) GetValServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValServiceId
}

// GetValServiceIdOk returns a tuple with the ValServiceId field value
// and a boolean to check if the value has been set.
func (o *NwSliceAdptInfo) GetValServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValServiceId, true
}

// SetValServiceId sets field value
func (o *NwSliceAdptInfo) SetValServiceId(v string) {
	o.ValServiceId = v
}

// GetValTgtUeIds returns the ValTgtUeIds field value
func (o *NwSliceAdptInfo) GetValTgtUeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ValTgtUeIds
}

// GetValTgtUeIdsOk returns a tuple with the ValTgtUeIds field value
// and a boolean to check if the value has been set.
func (o *NwSliceAdptInfo) GetValTgtUeIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValTgtUeIds, true
}

// SetValTgtUeIds sets field value
func (o *NwSliceAdptInfo) SetValTgtUeIds(v []string) {
	o.ValTgtUeIds = v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *NwSliceAdptInfo) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwSliceAdptInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *NwSliceAdptInfo) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *NwSliceAdptInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *NwSliceAdptInfo) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwSliceAdptInfo) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *NwSliceAdptInfo) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *NwSliceAdptInfo) SetDnn(v string) {
	o.Dnn = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *NwSliceAdptInfo) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwSliceAdptInfo) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *NwSliceAdptInfo) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *NwSliceAdptInfo) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o NwSliceAdptInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NwSliceAdptInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["valServiceId"] = o.ValServiceId
	toSerialize["valTgtUeIds"] = o.ValTgtUeIds
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableNwSliceAdptInfo struct {
	value *NwSliceAdptInfo
	isSet bool
}

func (v NullableNwSliceAdptInfo) Get() *NwSliceAdptInfo {
	return v.value
}

func (v *NullableNwSliceAdptInfo) Set(val *NwSliceAdptInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNwSliceAdptInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNwSliceAdptInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwSliceAdptInfo(val *NwSliceAdptInfo) *NullableNwSliceAdptInfo {
	return &NullableNwSliceAdptInfo{value: val, isSet: true}
}

func (v NullableNwSliceAdptInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwSliceAdptInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

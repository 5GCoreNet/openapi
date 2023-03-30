/*
Nnef_EASDeployment

NEF EAS Deployment service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_EASDeployment

import (
	"encoding/json"
)

// checks if the EasDeployInfoData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EasDeployInfoData{}

// EasDeployInfoData Represents the EAS Deployment Information to be reported.
type EasDeployInfoData struct {
	AppId *string `json:"appId,omitempty"`
	// list of DNS server identifier (consisting of IP address and port) and/or IP address(s)  of the EAS in the local DN for each DNAI. The key of map is the DNAI. 
	DnaiInfos *map[string]DnaiInformation `json:"dnaiInfos,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	FqdnPatternList []FqdnPatternMatchingRule `json:"fqdnPatternList"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	InternalGroupId *string `json:"internalGroupId,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
}

// NewEasDeployInfoData instantiates a new EasDeployInfoData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasDeployInfoData(fqdnPatternList []FqdnPatternMatchingRule) *EasDeployInfoData {
	this := EasDeployInfoData{}
	this.FqdnPatternList = fqdnPatternList
	return &this
}

// NewEasDeployInfoDataWithDefaults instantiates a new EasDeployInfoData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasDeployInfoDataWithDefaults() *EasDeployInfoData {
	this := EasDeployInfoData{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *EasDeployInfoData) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDeployInfoData) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *EasDeployInfoData) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *EasDeployInfoData) SetAppId(v string) {
	o.AppId = &v
}

// GetDnaiInfos returns the DnaiInfos field value if set, zero value otherwise.
func (o *EasDeployInfoData) GetDnaiInfos() map[string]DnaiInformation {
	if o == nil || IsNil(o.DnaiInfos) {
		var ret map[string]DnaiInformation
		return ret
	}
	return *o.DnaiInfos
}

// GetDnaiInfosOk returns a tuple with the DnaiInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDeployInfoData) GetDnaiInfosOk() (*map[string]DnaiInformation, bool) {
	if o == nil || IsNil(o.DnaiInfos) {
		return nil, false
	}
	return o.DnaiInfos, true
}

// HasDnaiInfos returns a boolean if a field has been set.
func (o *EasDeployInfoData) HasDnaiInfos() bool {
	if o != nil && !IsNil(o.DnaiInfos) {
		return true
	}

	return false
}

// SetDnaiInfos gets a reference to the given map[string]DnaiInformation and assigns it to the DnaiInfos field.
func (o *EasDeployInfoData) SetDnaiInfos(v map[string]DnaiInformation) {
	o.DnaiInfos = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *EasDeployInfoData) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDeployInfoData) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *EasDeployInfoData) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *EasDeployInfoData) SetDnn(v string) {
	o.Dnn = &v
}

// GetFqdnPatternList returns the FqdnPatternList field value
func (o *EasDeployInfoData) GetFqdnPatternList() []FqdnPatternMatchingRule {
	if o == nil {
		var ret []FqdnPatternMatchingRule
		return ret
	}

	return o.FqdnPatternList
}

// GetFqdnPatternListOk returns a tuple with the FqdnPatternList field value
// and a boolean to check if the value has been set.
func (o *EasDeployInfoData) GetFqdnPatternListOk() ([]FqdnPatternMatchingRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.FqdnPatternList, true
}

// SetFqdnPatternList sets field value
func (o *EasDeployInfoData) SetFqdnPatternList(v []FqdnPatternMatchingRule) {
	o.FqdnPatternList = v
}

// GetInternalGroupId returns the InternalGroupId field value if set, zero value otherwise.
func (o *EasDeployInfoData) GetInternalGroupId() string {
	if o == nil || IsNil(o.InternalGroupId) {
		var ret string
		return ret
	}
	return *o.InternalGroupId
}

// GetInternalGroupIdOk returns a tuple with the InternalGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDeployInfoData) GetInternalGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.InternalGroupId) {
		return nil, false
	}
	return o.InternalGroupId, true
}

// HasInternalGroupId returns a boolean if a field has been set.
func (o *EasDeployInfoData) HasInternalGroupId() bool {
	if o != nil && !IsNil(o.InternalGroupId) {
		return true
	}

	return false
}

// SetInternalGroupId gets a reference to the given string and assigns it to the InternalGroupId field.
func (o *EasDeployInfoData) SetInternalGroupId(v string) {
	o.InternalGroupId = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *EasDeployInfoData) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDeployInfoData) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *EasDeployInfoData) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *EasDeployInfoData) SetSnssai(v Snssai) {
	o.Snssai = &v
}

func (o EasDeployInfoData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EasDeployInfoData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.DnaiInfos) {
		toSerialize["dnaiInfos"] = o.DnaiInfos
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	toSerialize["fqdnPatternList"] = o.FqdnPatternList
	if !IsNil(o.InternalGroupId) {
		toSerialize["internalGroupId"] = o.InternalGroupId
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	return toSerialize, nil
}

type NullableEasDeployInfoData struct {
	value *EasDeployInfoData
	isSet bool
}

func (v NullableEasDeployInfoData) Get() *EasDeployInfoData {
	return v.value
}

func (v *NullableEasDeployInfoData) Set(val *EasDeployInfoData) {
	v.value = val
	v.isSet = true
}

func (v NullableEasDeployInfoData) IsSet() bool {
	return v.isSet
}

func (v *NullableEasDeployInfoData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasDeployInfoData(val *EasDeployInfoData) *NullableEasDeployInfoData {
	return &NullableEasDeployInfoData{value: val, isSet: true}
}

func (v NullableEasDeployInfoData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasDeployInfoData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



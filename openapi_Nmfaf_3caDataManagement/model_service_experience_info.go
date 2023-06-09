/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// checks if the ServiceExperienceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceExperienceInfo{}

// ServiceExperienceInfo Represents service experience information.
type ServiceExperienceInfo struct {
	SvcExprc SvcExperience `json:"svcExprc"`
	// string with format 'float' as defined in OpenAPI.
	SvcExprcVariance *float32 `json:"svcExprcVariance,omitempty"`
	Supis            []string `json:"supis,omitempty"`
	Snssai           *Snssai  `json:"snssai,omitempty"`
	// String providing an application identifier.
	AppId       *string                `json:"appId,omitempty"`
	SrvExpcType *ServiceExperienceType `json:"srvExpcType,omitempty"`
	UeLocs      []LocationInfo         `json:"ueLocs,omitempty"`
	UpfInfo     *UpfInformation        `json:"upfInfo,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai          *string   `json:"dnai,omitempty"`
	AppServerInst *AddrFqdn `json:"appServerInst,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn         *string          `json:"dnn,omitempty"`
	NetworkArea *NetworkAreaInfo `json:"networkArea,omitempty"`
	// Contains the Identifier of the selected Network Slice instance
	NsiId *string `json:"nsiId,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio   *int32              `json:"ratio,omitempty"`
	RatFreq *RatFreqInformation `json:"ratFreq,omitempty"`
}

// NewServiceExperienceInfo instantiates a new ServiceExperienceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceExperienceInfo(svcExprc SvcExperience) *ServiceExperienceInfo {
	this := ServiceExperienceInfo{}
	this.SvcExprc = svcExprc
	return &this
}

// NewServiceExperienceInfoWithDefaults instantiates a new ServiceExperienceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceExperienceInfoWithDefaults() *ServiceExperienceInfo {
	this := ServiceExperienceInfo{}
	return &this
}

// GetSvcExprc returns the SvcExprc field value
func (o *ServiceExperienceInfo) GetSvcExprc() SvcExperience {
	if o == nil {
		var ret SvcExperience
		return ret
	}

	return o.SvcExprc
}

// GetSvcExprcOk returns a tuple with the SvcExprc field value
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetSvcExprcOk() (*SvcExperience, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SvcExprc, true
}

// SetSvcExprc sets field value
func (o *ServiceExperienceInfo) SetSvcExprc(v SvcExperience) {
	o.SvcExprc = v
}

// GetSvcExprcVariance returns the SvcExprcVariance field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetSvcExprcVariance() float32 {
	if o == nil || IsNil(o.SvcExprcVariance) {
		var ret float32
		return ret
	}
	return *o.SvcExprcVariance
}

// GetSvcExprcVarianceOk returns a tuple with the SvcExprcVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetSvcExprcVarianceOk() (*float32, bool) {
	if o == nil || IsNil(o.SvcExprcVariance) {
		return nil, false
	}
	return o.SvcExprcVariance, true
}

// HasSvcExprcVariance returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasSvcExprcVariance() bool {
	if o != nil && !IsNil(o.SvcExprcVariance) {
		return true
	}

	return false
}

// SetSvcExprcVariance gets a reference to the given float32 and assigns it to the SvcExprcVariance field.
func (o *ServiceExperienceInfo) SetSvcExprcVariance(v float32) {
	o.SvcExprcVariance = &v
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetSupis() []string {
	if o == nil || IsNil(o.Supis) {
		var ret []string
		return ret
	}
	return o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetSupisOk() ([]string, bool) {
	if o == nil || IsNil(o.Supis) {
		return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasSupis() bool {
	if o != nil && !IsNil(o.Supis) {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []string and assigns it to the Supis field.
func (o *ServiceExperienceInfo) SetSupis(v []string) {
	o.Supis = v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *ServiceExperienceInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *ServiceExperienceInfo) SetAppId(v string) {
	o.AppId = &v
}

// GetSrvExpcType returns the SrvExpcType field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetSrvExpcType() ServiceExperienceType {
	if o == nil || IsNil(o.SrvExpcType) {
		var ret ServiceExperienceType
		return ret
	}
	return *o.SrvExpcType
}

// GetSrvExpcTypeOk returns a tuple with the SrvExpcType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetSrvExpcTypeOk() (*ServiceExperienceType, bool) {
	if o == nil || IsNil(o.SrvExpcType) {
		return nil, false
	}
	return o.SrvExpcType, true
}

// HasSrvExpcType returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasSrvExpcType() bool {
	if o != nil && !IsNil(o.SrvExpcType) {
		return true
	}

	return false
}

// SetSrvExpcType gets a reference to the given ServiceExperienceType and assigns it to the SrvExpcType field.
func (o *ServiceExperienceInfo) SetSrvExpcType(v ServiceExperienceType) {
	o.SrvExpcType = &v
}

// GetUeLocs returns the UeLocs field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetUeLocs() []LocationInfo {
	if o == nil || IsNil(o.UeLocs) {
		var ret []LocationInfo
		return ret
	}
	return o.UeLocs
}

// GetUeLocsOk returns a tuple with the UeLocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetUeLocsOk() ([]LocationInfo, bool) {
	if o == nil || IsNil(o.UeLocs) {
		return nil, false
	}
	return o.UeLocs, true
}

// HasUeLocs returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasUeLocs() bool {
	if o != nil && !IsNil(o.UeLocs) {
		return true
	}

	return false
}

// SetUeLocs gets a reference to the given []LocationInfo and assigns it to the UeLocs field.
func (o *ServiceExperienceInfo) SetUeLocs(v []LocationInfo) {
	o.UeLocs = v
}

// GetUpfInfo returns the UpfInfo field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetUpfInfo() UpfInformation {
	if o == nil || IsNil(o.UpfInfo) {
		var ret UpfInformation
		return ret
	}
	return *o.UpfInfo
}

// GetUpfInfoOk returns a tuple with the UpfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetUpfInfoOk() (*UpfInformation, bool) {
	if o == nil || IsNil(o.UpfInfo) {
		return nil, false
	}
	return o.UpfInfo, true
}

// HasUpfInfo returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasUpfInfo() bool {
	if o != nil && !IsNil(o.UpfInfo) {
		return true
	}

	return false
}

// SetUpfInfo gets a reference to the given UpfInformation and assigns it to the UpfInfo field.
func (o *ServiceExperienceInfo) SetUpfInfo(v UpfInformation) {
	o.UpfInfo = &v
}

// GetDnai returns the Dnai field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetDnai() string {
	if o == nil || IsNil(o.Dnai) {
		var ret string
		return ret
	}
	return *o.Dnai
}

// GetDnaiOk returns a tuple with the Dnai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetDnaiOk() (*string, bool) {
	if o == nil || IsNil(o.Dnai) {
		return nil, false
	}
	return o.Dnai, true
}

// HasDnai returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasDnai() bool {
	if o != nil && !IsNil(o.Dnai) {
		return true
	}

	return false
}

// SetDnai gets a reference to the given string and assigns it to the Dnai field.
func (o *ServiceExperienceInfo) SetDnai(v string) {
	o.Dnai = &v
}

// GetAppServerInst returns the AppServerInst field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetAppServerInst() AddrFqdn {
	if o == nil || IsNil(o.AppServerInst) {
		var ret AddrFqdn
		return ret
	}
	return *o.AppServerInst
}

// GetAppServerInstOk returns a tuple with the AppServerInst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetAppServerInstOk() (*AddrFqdn, bool) {
	if o == nil || IsNil(o.AppServerInst) {
		return nil, false
	}
	return o.AppServerInst, true
}

// HasAppServerInst returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasAppServerInst() bool {
	if o != nil && !IsNil(o.AppServerInst) {
		return true
	}

	return false
}

// SetAppServerInst gets a reference to the given AddrFqdn and assigns it to the AppServerInst field.
func (o *ServiceExperienceInfo) SetAppServerInst(v AddrFqdn) {
	o.AppServerInst = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetConfidence() int32 {
	if o == nil || IsNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetConfidenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *ServiceExperienceInfo) SetConfidence(v int32) {
	o.Confidence = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *ServiceExperienceInfo) SetDnn(v string) {
	o.Dnn = &v
}

// GetNetworkArea returns the NetworkArea field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetNetworkArea() NetworkAreaInfo {
	if o == nil || IsNil(o.NetworkArea) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.NetworkArea
}

// GetNetworkAreaOk returns a tuple with the NetworkArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetNetworkAreaOk() (*NetworkAreaInfo, bool) {
	if o == nil || IsNil(o.NetworkArea) {
		return nil, false
	}
	return o.NetworkArea, true
}

// HasNetworkArea returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasNetworkArea() bool {
	if o != nil && !IsNil(o.NetworkArea) {
		return true
	}

	return false
}

// SetNetworkArea gets a reference to the given NetworkAreaInfo and assigns it to the NetworkArea field.
func (o *ServiceExperienceInfo) SetNetworkArea(v NetworkAreaInfo) {
	o.NetworkArea = &v
}

// GetNsiId returns the NsiId field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetNsiId() string {
	if o == nil || IsNil(o.NsiId) {
		var ret string
		return ret
	}
	return *o.NsiId
}

// GetNsiIdOk returns a tuple with the NsiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetNsiIdOk() (*string, bool) {
	if o == nil || IsNil(o.NsiId) {
		return nil, false
	}
	return o.NsiId, true
}

// HasNsiId returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasNsiId() bool {
	if o != nil && !IsNil(o.NsiId) {
		return true
	}

	return false
}

// SetNsiId gets a reference to the given string and assigns it to the NsiId field.
func (o *ServiceExperienceInfo) SetNsiId(v string) {
	o.NsiId = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetRatio() int32 {
	if o == nil || IsNil(o.Ratio) {
		var ret int32
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetRatioOk() (*int32, bool) {
	if o == nil || IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasRatio() bool {
	if o != nil && !IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given int32 and assigns it to the Ratio field.
func (o *ServiceExperienceInfo) SetRatio(v int32) {
	o.Ratio = &v
}

// GetRatFreq returns the RatFreq field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetRatFreq() RatFreqInformation {
	if o == nil || IsNil(o.RatFreq) {
		var ret RatFreqInformation
		return ret
	}
	return *o.RatFreq
}

// GetRatFreqOk returns a tuple with the RatFreq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetRatFreqOk() (*RatFreqInformation, bool) {
	if o == nil || IsNil(o.RatFreq) {
		return nil, false
	}
	return o.RatFreq, true
}

// HasRatFreq returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasRatFreq() bool {
	if o != nil && !IsNil(o.RatFreq) {
		return true
	}

	return false
}

// SetRatFreq gets a reference to the given RatFreqInformation and assigns it to the RatFreq field.
func (o *ServiceExperienceInfo) SetRatFreq(v RatFreqInformation) {
	o.RatFreq = &v
}

func (o ServiceExperienceInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceExperienceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["svcExprc"] = o.SvcExprc
	if !IsNil(o.SvcExprcVariance) {
		toSerialize["svcExprcVariance"] = o.SvcExprcVariance
	}
	if !IsNil(o.Supis) {
		toSerialize["supis"] = o.Supis
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.SrvExpcType) {
		toSerialize["srvExpcType"] = o.SrvExpcType
	}
	if !IsNil(o.UeLocs) {
		toSerialize["ueLocs"] = o.UeLocs
	}
	if !IsNil(o.UpfInfo) {
		toSerialize["upfInfo"] = o.UpfInfo
	}
	if !IsNil(o.Dnai) {
		toSerialize["dnai"] = o.Dnai
	}
	if !IsNil(o.AppServerInst) {
		toSerialize["appServerInst"] = o.AppServerInst
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.NetworkArea) {
		toSerialize["networkArea"] = o.NetworkArea
	}
	if !IsNil(o.NsiId) {
		toSerialize["nsiId"] = o.NsiId
	}
	if !IsNil(o.Ratio) {
		toSerialize["ratio"] = o.Ratio
	}
	if !IsNil(o.RatFreq) {
		toSerialize["ratFreq"] = o.RatFreq
	}
	return toSerialize, nil
}

type NullableServiceExperienceInfo struct {
	value *ServiceExperienceInfo
	isSet bool
}

func (v NullableServiceExperienceInfo) Get() *ServiceExperienceInfo {
	return v.value
}

func (v *NullableServiceExperienceInfo) Set(val *ServiceExperienceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceExperienceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceExperienceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceExperienceInfo(val *ServiceExperienceInfo) *NullableServiceExperienceInfo {
	return &NullableServiceExperienceInfo{value: val, isSet: true}
}

func (v NullableServiceExperienceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceExperienceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

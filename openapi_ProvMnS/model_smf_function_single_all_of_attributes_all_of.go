/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the SmfFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmfFunctionSingleAllOfAttributesAllOf{}

// SmfFunctionSingleAllOfAttributesAllOf struct for SmfFunctionSingleAllOfAttributesAllOf
type SmfFunctionSingleAllOfAttributesAllOf struct {
	PLMNInfoList []PlmnInfo `json:"pLMNInfoList,omitempty"`
	NRTACList []int32 `json:"nRTACList,omitempty"`
	SBIFqdn *string `json:"sBIFqdn,omitempty"`
	SNssaiSmfInfoList []SNssaiSmfInfoItem `json:"sNssaiSmfInfoList,omitempty"`
	TaiList []Tai `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	PwgFqdn *string `json:"pwgFqdn,omitempty"`
	PgwAddrList []IpAddr1 `json:"pgwAddrList,omitempty"`
	AccessType *AccessType `json:"accessType,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	CNSIIdList []string `json:"cNSIIdList,omitempty"`
	VsmfSupportInd *bool `json:"vsmfSupportInd,omitempty"`
	PwgFqdnList []string `json:"pwgFqdnList,omitempty"`
	ManagedNFProfile *ManagedNFProfile `json:"managedNFProfile,omitempty"`
	CommModelList []CommModel `json:"commModelList,omitempty"`
	Configurable5QISetRef *string `json:"configurable5QISetRef,omitempty"`
	Dynamic5QISetRef *string `json:"dynamic5QISetRef,omitempty"`
}

// NewSmfFunctionSingleAllOfAttributesAllOf instantiates a new SmfFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmfFunctionSingleAllOfAttributesAllOf() *SmfFunctionSingleAllOfAttributesAllOf {
	this := SmfFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewSmfFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new SmfFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmfFunctionSingleAllOfAttributesAllOfWithDefaults() *SmfFunctionSingleAllOfAttributesAllOf {
	this := SmfFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetPLMNInfoList returns the PLMNInfoList field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetPLMNInfoList() []PlmnInfo {
	if o == nil || IsNil(o.PLMNInfoList) {
		var ret []PlmnInfo
		return ret
	}
	return o.PLMNInfoList
}

// GetPLMNInfoListOk returns a tuple with the PLMNInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetPLMNInfoListOk() ([]PlmnInfo, bool) {
	if o == nil || IsNil(o.PLMNInfoList) {
		return nil, false
	}
	return o.PLMNInfoList, true
}

// HasPLMNInfoList returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasPLMNInfoList() bool {
	if o != nil && !IsNil(o.PLMNInfoList) {
		return true
	}

	return false
}

// SetPLMNInfoList gets a reference to the given []PlmnInfo and assigns it to the PLMNInfoList field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetPLMNInfoList(v []PlmnInfo) {
	o.PLMNInfoList = v
}

// GetNRTACList returns the NRTACList field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetNRTACList() []int32 {
	if o == nil || IsNil(o.NRTACList) {
		var ret []int32
		return ret
	}
	return o.NRTACList
}

// GetNRTACListOk returns a tuple with the NRTACList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetNRTACListOk() ([]int32, bool) {
	if o == nil || IsNil(o.NRTACList) {
		return nil, false
	}
	return o.NRTACList, true
}

// HasNRTACList returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasNRTACList() bool {
	if o != nil && !IsNil(o.NRTACList) {
		return true
	}

	return false
}

// SetNRTACList gets a reference to the given []int32 and assigns it to the NRTACList field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetNRTACList(v []int32) {
	o.NRTACList = v
}

// GetSBIFqdn returns the SBIFqdn field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetSBIFqdn() string {
	if o == nil || IsNil(o.SBIFqdn) {
		var ret string
		return ret
	}
	return *o.SBIFqdn
}

// GetSBIFqdnOk returns a tuple with the SBIFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetSBIFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.SBIFqdn) {
		return nil, false
	}
	return o.SBIFqdn, true
}

// HasSBIFqdn returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasSBIFqdn() bool {
	if o != nil && !IsNil(o.SBIFqdn) {
		return true
	}

	return false
}

// SetSBIFqdn gets a reference to the given string and assigns it to the SBIFqdn field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetSBIFqdn(v string) {
	o.SBIFqdn = &v
}

// GetSNssaiSmfInfoList returns the SNssaiSmfInfoList field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetSNssaiSmfInfoList() []SNssaiSmfInfoItem {
	if o == nil || IsNil(o.SNssaiSmfInfoList) {
		var ret []SNssaiSmfInfoItem
		return ret
	}
	return o.SNssaiSmfInfoList
}

// GetSNssaiSmfInfoListOk returns a tuple with the SNssaiSmfInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetSNssaiSmfInfoListOk() ([]SNssaiSmfInfoItem, bool) {
	if o == nil || IsNil(o.SNssaiSmfInfoList) {
		return nil, false
	}
	return o.SNssaiSmfInfoList, true
}

// HasSNssaiSmfInfoList returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasSNssaiSmfInfoList() bool {
	if o != nil && !IsNil(o.SNssaiSmfInfoList) {
		return true
	}

	return false
}

// SetSNssaiSmfInfoList gets a reference to the given []SNssaiSmfInfoItem and assigns it to the SNssaiSmfInfoList field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetSNssaiSmfInfoList(v []SNssaiSmfInfoItem) {
	o.SNssaiSmfInfoList = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetTaiList() []Tai {
	if o == nil || IsNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetTaiListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TaiList) {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasTaiList() bool {
	if o != nil && !IsNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetTaiRangeList() []TaiRange {
	if o == nil || IsNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || IsNil(o.TaiRangeList) {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasTaiRangeList() bool {
	if o != nil && !IsNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetPwgFqdn returns the PwgFqdn field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetPwgFqdn() string {
	if o == nil || IsNil(o.PwgFqdn) {
		var ret string
		return ret
	}
	return *o.PwgFqdn
}

// GetPwgFqdnOk returns a tuple with the PwgFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetPwgFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.PwgFqdn) {
		return nil, false
	}
	return o.PwgFqdn, true
}

// HasPwgFqdn returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasPwgFqdn() bool {
	if o != nil && !IsNil(o.PwgFqdn) {
		return true
	}

	return false
}

// SetPwgFqdn gets a reference to the given string and assigns it to the PwgFqdn field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetPwgFqdn(v string) {
	o.PwgFqdn = &v
}

// GetPgwAddrList returns the PgwAddrList field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetPgwAddrList() []IpAddr1 {
	if o == nil || IsNil(o.PgwAddrList) {
		var ret []IpAddr1
		return ret
	}
	return o.PgwAddrList
}

// GetPgwAddrListOk returns a tuple with the PgwAddrList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetPgwAddrListOk() ([]IpAddr1, bool) {
	if o == nil || IsNil(o.PgwAddrList) {
		return nil, false
	}
	return o.PgwAddrList, true
}

// HasPgwAddrList returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasPgwAddrList() bool {
	if o != nil && !IsNil(o.PgwAddrList) {
		return true
	}

	return false
}

// SetPgwAddrList gets a reference to the given []IpAddr1 and assigns it to the PgwAddrList field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetPgwAddrList(v []IpAddr1) {
	o.PgwAddrList = v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetAccessType() AccessType {
	if o == nil || IsNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetPriority(v int32) {
	o.Priority = &v
}

// GetCNSIIdList returns the CNSIIdList field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetCNSIIdList() []string {
	if o == nil || IsNil(o.CNSIIdList) {
		var ret []string
		return ret
	}
	return o.CNSIIdList
}

// GetCNSIIdListOk returns a tuple with the CNSIIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetCNSIIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.CNSIIdList) {
		return nil, false
	}
	return o.CNSIIdList, true
}

// HasCNSIIdList returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasCNSIIdList() bool {
	if o != nil && !IsNil(o.CNSIIdList) {
		return true
	}

	return false
}

// SetCNSIIdList gets a reference to the given []string and assigns it to the CNSIIdList field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetCNSIIdList(v []string) {
	o.CNSIIdList = v
}

// GetVsmfSupportInd returns the VsmfSupportInd field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetVsmfSupportInd() bool {
	if o == nil || IsNil(o.VsmfSupportInd) {
		var ret bool
		return ret
	}
	return *o.VsmfSupportInd
}

// GetVsmfSupportIndOk returns a tuple with the VsmfSupportInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetVsmfSupportIndOk() (*bool, bool) {
	if o == nil || IsNil(o.VsmfSupportInd) {
		return nil, false
	}
	return o.VsmfSupportInd, true
}

// HasVsmfSupportInd returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasVsmfSupportInd() bool {
	if o != nil && !IsNil(o.VsmfSupportInd) {
		return true
	}

	return false
}

// SetVsmfSupportInd gets a reference to the given bool and assigns it to the VsmfSupportInd field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetVsmfSupportInd(v bool) {
	o.VsmfSupportInd = &v
}

// GetPwgFqdnList returns the PwgFqdnList field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetPwgFqdnList() []string {
	if o == nil || IsNil(o.PwgFqdnList) {
		var ret []string
		return ret
	}
	return o.PwgFqdnList
}

// GetPwgFqdnListOk returns a tuple with the PwgFqdnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetPwgFqdnListOk() ([]string, bool) {
	if o == nil || IsNil(o.PwgFqdnList) {
		return nil, false
	}
	return o.PwgFqdnList, true
}

// HasPwgFqdnList returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasPwgFqdnList() bool {
	if o != nil && !IsNil(o.PwgFqdnList) {
		return true
	}

	return false
}

// SetPwgFqdnList gets a reference to the given []string and assigns it to the PwgFqdnList field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetPwgFqdnList(v []string) {
	o.PwgFqdnList = v
}

// GetManagedNFProfile returns the ManagedNFProfile field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetManagedNFProfile() ManagedNFProfile {
	if o == nil || IsNil(o.ManagedNFProfile) {
		var ret ManagedNFProfile
		return ret
	}
	return *o.ManagedNFProfile
}

// GetManagedNFProfileOk returns a tuple with the ManagedNFProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetManagedNFProfileOk() (*ManagedNFProfile, bool) {
	if o == nil || IsNil(o.ManagedNFProfile) {
		return nil, false
	}
	return o.ManagedNFProfile, true
}

// HasManagedNFProfile returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasManagedNFProfile() bool {
	if o != nil && !IsNil(o.ManagedNFProfile) {
		return true
	}

	return false
}

// SetManagedNFProfile gets a reference to the given ManagedNFProfile and assigns it to the ManagedNFProfile field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetManagedNFProfile(v ManagedNFProfile) {
	o.ManagedNFProfile = &v
}

// GetCommModelList returns the CommModelList field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetCommModelList() []CommModel {
	if o == nil || IsNil(o.CommModelList) {
		var ret []CommModel
		return ret
	}
	return o.CommModelList
}

// GetCommModelListOk returns a tuple with the CommModelList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetCommModelListOk() ([]CommModel, bool) {
	if o == nil || IsNil(o.CommModelList) {
		return nil, false
	}
	return o.CommModelList, true
}

// HasCommModelList returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasCommModelList() bool {
	if o != nil && !IsNil(o.CommModelList) {
		return true
	}

	return false
}

// SetCommModelList gets a reference to the given []CommModel and assigns it to the CommModelList field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetCommModelList(v []CommModel) {
	o.CommModelList = v
}

// GetConfigurable5QISetRef returns the Configurable5QISetRef field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetConfigurable5QISetRef() string {
	if o == nil || IsNil(o.Configurable5QISetRef) {
		var ret string
		return ret
	}
	return *o.Configurable5QISetRef
}

// GetConfigurable5QISetRefOk returns a tuple with the Configurable5QISetRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetConfigurable5QISetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Configurable5QISetRef) {
		return nil, false
	}
	return o.Configurable5QISetRef, true
}

// HasConfigurable5QISetRef returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasConfigurable5QISetRef() bool {
	if o != nil && !IsNil(o.Configurable5QISetRef) {
		return true
	}

	return false
}

// SetConfigurable5QISetRef gets a reference to the given string and assigns it to the Configurable5QISetRef field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetConfigurable5QISetRef(v string) {
	o.Configurable5QISetRef = &v
}

// GetDynamic5QISetRef returns the Dynamic5QISetRef field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetDynamic5QISetRef() string {
	if o == nil || IsNil(o.Dynamic5QISetRef) {
		var ret string
		return ret
	}
	return *o.Dynamic5QISetRef
}

// GetDynamic5QISetRefOk returns a tuple with the Dynamic5QISetRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) GetDynamic5QISetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Dynamic5QISetRef) {
		return nil, false
	}
	return o.Dynamic5QISetRef, true
}

// HasDynamic5QISetRef returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOfAttributesAllOf) HasDynamic5QISetRef() bool {
	if o != nil && !IsNil(o.Dynamic5QISetRef) {
		return true
	}

	return false
}

// SetDynamic5QISetRef gets a reference to the given string and assigns it to the Dynamic5QISetRef field.
func (o *SmfFunctionSingleAllOfAttributesAllOf) SetDynamic5QISetRef(v string) {
	o.Dynamic5QISetRef = &v
}

func (o SmfFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmfFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PLMNInfoList) {
		toSerialize["pLMNInfoList"] = o.PLMNInfoList
	}
	if !IsNil(o.NRTACList) {
		toSerialize["nRTACList"] = o.NRTACList
	}
	if !IsNil(o.SBIFqdn) {
		toSerialize["sBIFqdn"] = o.SBIFqdn
	}
	if !IsNil(o.SNssaiSmfInfoList) {
		toSerialize["sNssaiSmfInfoList"] = o.SNssaiSmfInfoList
	}
	if !IsNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !IsNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !IsNil(o.PwgFqdn) {
		toSerialize["pwgFqdn"] = o.PwgFqdn
	}
	if !IsNil(o.PgwAddrList) {
		toSerialize["pgwAddrList"] = o.PgwAddrList
	}
	if !IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.CNSIIdList) {
		toSerialize["cNSIIdList"] = o.CNSIIdList
	}
	if !IsNil(o.VsmfSupportInd) {
		toSerialize["vsmfSupportInd"] = o.VsmfSupportInd
	}
	if !IsNil(o.PwgFqdnList) {
		toSerialize["pwgFqdnList"] = o.PwgFqdnList
	}
	if !IsNil(o.ManagedNFProfile) {
		toSerialize["managedNFProfile"] = o.ManagedNFProfile
	}
	if !IsNil(o.CommModelList) {
		toSerialize["commModelList"] = o.CommModelList
	}
	if !IsNil(o.Configurable5QISetRef) {
		toSerialize["configurable5QISetRef"] = o.Configurable5QISetRef
	}
	if !IsNil(o.Dynamic5QISetRef) {
		toSerialize["dynamic5QISetRef"] = o.Dynamic5QISetRef
	}
	return toSerialize, nil
}

type NullableSmfFunctionSingleAllOfAttributesAllOf struct {
	value *SmfFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableSmfFunctionSingleAllOfAttributesAllOf) Get() *SmfFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableSmfFunctionSingleAllOfAttributesAllOf) Set(val *SmfFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSmfFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSmfFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmfFunctionSingleAllOfAttributesAllOf(val *SmfFunctionSingleAllOfAttributesAllOf) *NullableSmfFunctionSingleAllOfAttributesAllOf {
	return &NullableSmfFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableSmfFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmfFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


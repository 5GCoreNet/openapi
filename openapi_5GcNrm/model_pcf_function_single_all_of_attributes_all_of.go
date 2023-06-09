/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the PcfFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcfFunctionSingleAllOfAttributesAllOf{}

// PcfFunctionSingleAllOfAttributesAllOf struct for PcfFunctionSingleAllOfAttributesAllOf
type PcfFunctionSingleAllOfAttributesAllOf struct {
	PLMNInfoList          []PlmnInfo        `json:"pLMNInfoList,omitempty"`
	SBIFqdn               *string           `json:"sBIFqdn,omitempty"`
	ManagedNFProfile      *ManagedNFProfile `json:"managedNFProfile,omitempty"`
	CommModelList         []CommModel       `json:"commModelList,omitempty"`
	GroupId               *string           `json:"groupId,omitempty"`
	DnnList               []string          `json:"dnnList,omitempty"`
	SupiRanges            []SupiRange       `json:"supiRanges,omitempty"`
	GpsiRanges            []IdentityRange   `json:"gpsiRanges,omitempty"`
	RxDiamHost            *string           `json:"rxDiamHost,omitempty"`
	RxDiamRealm           *string           `json:"rxDiamRealm,omitempty"`
	V2xSupportInd         *bool             `json:"v2xSupportInd,omitempty"`
	ProseSupportInd       *bool             `json:"proseSupportInd,omitempty"`
	ProseCapability       *ProseCapability  `json:"proseCapability,omitempty"`
	V2xCapability         *V2xCapability    `json:"v2xCapability,omitempty"`
	Configurable5QISetRef *string           `json:"configurable5QISetRef,omitempty"`
	Dynamic5QISetRef      *string           `json:"dynamic5QISetRef,omitempty"`
	SupportedBMOList      []string          `json:"supportedBMOList,omitempty"`
}

// NewPcfFunctionSingleAllOfAttributesAllOf instantiates a new PcfFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfFunctionSingleAllOfAttributesAllOf() *PcfFunctionSingleAllOfAttributesAllOf {
	this := PcfFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewPcfFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new PcfFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfFunctionSingleAllOfAttributesAllOfWithDefaults() *PcfFunctionSingleAllOfAttributesAllOf {
	this := PcfFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetPLMNInfoList returns the PLMNInfoList field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetPLMNInfoList() []PlmnInfo {
	if o == nil || IsNil(o.PLMNInfoList) {
		var ret []PlmnInfo
		return ret
	}
	return o.PLMNInfoList
}

// GetPLMNInfoListOk returns a tuple with the PLMNInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetPLMNInfoListOk() ([]PlmnInfo, bool) {
	if o == nil || IsNil(o.PLMNInfoList) {
		return nil, false
	}
	return o.PLMNInfoList, true
}

// HasPLMNInfoList returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasPLMNInfoList() bool {
	if o != nil && !IsNil(o.PLMNInfoList) {
		return true
	}

	return false
}

// SetPLMNInfoList gets a reference to the given []PlmnInfo and assigns it to the PLMNInfoList field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetPLMNInfoList(v []PlmnInfo) {
	o.PLMNInfoList = v
}

// GetSBIFqdn returns the SBIFqdn field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetSBIFqdn() string {
	if o == nil || IsNil(o.SBIFqdn) {
		var ret string
		return ret
	}
	return *o.SBIFqdn
}

// GetSBIFqdnOk returns a tuple with the SBIFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetSBIFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.SBIFqdn) {
		return nil, false
	}
	return o.SBIFqdn, true
}

// HasSBIFqdn returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasSBIFqdn() bool {
	if o != nil && !IsNil(o.SBIFqdn) {
		return true
	}

	return false
}

// SetSBIFqdn gets a reference to the given string and assigns it to the SBIFqdn field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetSBIFqdn(v string) {
	o.SBIFqdn = &v
}

// GetManagedNFProfile returns the ManagedNFProfile field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetManagedNFProfile() ManagedNFProfile {
	if o == nil || IsNil(o.ManagedNFProfile) {
		var ret ManagedNFProfile
		return ret
	}
	return *o.ManagedNFProfile
}

// GetManagedNFProfileOk returns a tuple with the ManagedNFProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetManagedNFProfileOk() (*ManagedNFProfile, bool) {
	if o == nil || IsNil(o.ManagedNFProfile) {
		return nil, false
	}
	return o.ManagedNFProfile, true
}

// HasManagedNFProfile returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasManagedNFProfile() bool {
	if o != nil && !IsNil(o.ManagedNFProfile) {
		return true
	}

	return false
}

// SetManagedNFProfile gets a reference to the given ManagedNFProfile and assigns it to the ManagedNFProfile field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetManagedNFProfile(v ManagedNFProfile) {
	o.ManagedNFProfile = &v
}

// GetCommModelList returns the CommModelList field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetCommModelList() []CommModel {
	if o == nil || IsNil(o.CommModelList) {
		var ret []CommModel
		return ret
	}
	return o.CommModelList
}

// GetCommModelListOk returns a tuple with the CommModelList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetCommModelListOk() ([]CommModel, bool) {
	if o == nil || IsNil(o.CommModelList) {
		return nil, false
	}
	return o.CommModelList, true
}

// HasCommModelList returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasCommModelList() bool {
	if o != nil && !IsNil(o.CommModelList) {
		return true
	}

	return false
}

// SetCommModelList gets a reference to the given []CommModel and assigns it to the CommModelList field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetCommModelList(v []CommModel) {
	o.CommModelList = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetGroupId(v string) {
	o.GroupId = &v
}

// GetDnnList returns the DnnList field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetDnnList() []string {
	if o == nil || IsNil(o.DnnList) {
		var ret []string
		return ret
	}
	return o.DnnList
}

// GetDnnListOk returns a tuple with the DnnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetDnnListOk() ([]string, bool) {
	if o == nil || IsNil(o.DnnList) {
		return nil, false
	}
	return o.DnnList, true
}

// HasDnnList returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasDnnList() bool {
	if o != nil && !IsNil(o.DnnList) {
		return true
	}

	return false
}

// SetDnnList gets a reference to the given []string and assigns it to the DnnList field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetDnnList(v []string) {
	o.DnnList = v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetSupiRanges() []SupiRange {
	if o == nil || IsNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || IsNil(o.SupiRanges) {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasSupiRanges() bool {
	if o != nil && !IsNil(o.SupiRanges) {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetGpsiRanges() []IdentityRange {
	if o == nil || IsNil(o.GpsiRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.GpsiRanges) {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasGpsiRanges() bool {
	if o != nil && !IsNil(o.GpsiRanges) {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

// GetRxDiamHost returns the RxDiamHost field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetRxDiamHost() string {
	if o == nil || IsNil(o.RxDiamHost) {
		var ret string
		return ret
	}
	return *o.RxDiamHost
}

// GetRxDiamHostOk returns a tuple with the RxDiamHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetRxDiamHostOk() (*string, bool) {
	if o == nil || IsNil(o.RxDiamHost) {
		return nil, false
	}
	return o.RxDiamHost, true
}

// HasRxDiamHost returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasRxDiamHost() bool {
	if o != nil && !IsNil(o.RxDiamHost) {
		return true
	}

	return false
}

// SetRxDiamHost gets a reference to the given string and assigns it to the RxDiamHost field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetRxDiamHost(v string) {
	o.RxDiamHost = &v
}

// GetRxDiamRealm returns the RxDiamRealm field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetRxDiamRealm() string {
	if o == nil || IsNil(o.RxDiamRealm) {
		var ret string
		return ret
	}
	return *o.RxDiamRealm
}

// GetRxDiamRealmOk returns a tuple with the RxDiamRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetRxDiamRealmOk() (*string, bool) {
	if o == nil || IsNil(o.RxDiamRealm) {
		return nil, false
	}
	return o.RxDiamRealm, true
}

// HasRxDiamRealm returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasRxDiamRealm() bool {
	if o != nil && !IsNil(o.RxDiamRealm) {
		return true
	}

	return false
}

// SetRxDiamRealm gets a reference to the given string and assigns it to the RxDiamRealm field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetRxDiamRealm(v string) {
	o.RxDiamRealm = &v
}

// GetV2xSupportInd returns the V2xSupportInd field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetV2xSupportInd() bool {
	if o == nil || IsNil(o.V2xSupportInd) {
		var ret bool
		return ret
	}
	return *o.V2xSupportInd
}

// GetV2xSupportIndOk returns a tuple with the V2xSupportInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetV2xSupportIndOk() (*bool, bool) {
	if o == nil || IsNil(o.V2xSupportInd) {
		return nil, false
	}
	return o.V2xSupportInd, true
}

// HasV2xSupportInd returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasV2xSupportInd() bool {
	if o != nil && !IsNil(o.V2xSupportInd) {
		return true
	}

	return false
}

// SetV2xSupportInd gets a reference to the given bool and assigns it to the V2xSupportInd field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetV2xSupportInd(v bool) {
	o.V2xSupportInd = &v
}

// GetProseSupportInd returns the ProseSupportInd field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetProseSupportInd() bool {
	if o == nil || IsNil(o.ProseSupportInd) {
		var ret bool
		return ret
	}
	return *o.ProseSupportInd
}

// GetProseSupportIndOk returns a tuple with the ProseSupportInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetProseSupportIndOk() (*bool, bool) {
	if o == nil || IsNil(o.ProseSupportInd) {
		return nil, false
	}
	return o.ProseSupportInd, true
}

// HasProseSupportInd returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasProseSupportInd() bool {
	if o != nil && !IsNil(o.ProseSupportInd) {
		return true
	}

	return false
}

// SetProseSupportInd gets a reference to the given bool and assigns it to the ProseSupportInd field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetProseSupportInd(v bool) {
	o.ProseSupportInd = &v
}

// GetProseCapability returns the ProseCapability field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetProseCapability() ProseCapability {
	if o == nil || IsNil(o.ProseCapability) {
		var ret ProseCapability
		return ret
	}
	return *o.ProseCapability
}

// GetProseCapabilityOk returns a tuple with the ProseCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetProseCapabilityOk() (*ProseCapability, bool) {
	if o == nil || IsNil(o.ProseCapability) {
		return nil, false
	}
	return o.ProseCapability, true
}

// HasProseCapability returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasProseCapability() bool {
	if o != nil && !IsNil(o.ProseCapability) {
		return true
	}

	return false
}

// SetProseCapability gets a reference to the given ProseCapability and assigns it to the ProseCapability field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetProseCapability(v ProseCapability) {
	o.ProseCapability = &v
}

// GetV2xCapability returns the V2xCapability field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetV2xCapability() V2xCapability {
	if o == nil || IsNil(o.V2xCapability) {
		var ret V2xCapability
		return ret
	}
	return *o.V2xCapability
}

// GetV2xCapabilityOk returns a tuple with the V2xCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetV2xCapabilityOk() (*V2xCapability, bool) {
	if o == nil || IsNil(o.V2xCapability) {
		return nil, false
	}
	return o.V2xCapability, true
}

// HasV2xCapability returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasV2xCapability() bool {
	if o != nil && !IsNil(o.V2xCapability) {
		return true
	}

	return false
}

// SetV2xCapability gets a reference to the given V2xCapability and assigns it to the V2xCapability field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetV2xCapability(v V2xCapability) {
	o.V2xCapability = &v
}

// GetConfigurable5QISetRef returns the Configurable5QISetRef field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetConfigurable5QISetRef() string {
	if o == nil || IsNil(o.Configurable5QISetRef) {
		var ret string
		return ret
	}
	return *o.Configurable5QISetRef
}

// GetConfigurable5QISetRefOk returns a tuple with the Configurable5QISetRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetConfigurable5QISetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Configurable5QISetRef) {
		return nil, false
	}
	return o.Configurable5QISetRef, true
}

// HasConfigurable5QISetRef returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasConfigurable5QISetRef() bool {
	if o != nil && !IsNil(o.Configurable5QISetRef) {
		return true
	}

	return false
}

// SetConfigurable5QISetRef gets a reference to the given string and assigns it to the Configurable5QISetRef field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetConfigurable5QISetRef(v string) {
	o.Configurable5QISetRef = &v
}

// GetDynamic5QISetRef returns the Dynamic5QISetRef field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetDynamic5QISetRef() string {
	if o == nil || IsNil(o.Dynamic5QISetRef) {
		var ret string
		return ret
	}
	return *o.Dynamic5QISetRef
}

// GetDynamic5QISetRefOk returns a tuple with the Dynamic5QISetRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetDynamic5QISetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Dynamic5QISetRef) {
		return nil, false
	}
	return o.Dynamic5QISetRef, true
}

// HasDynamic5QISetRef returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasDynamic5QISetRef() bool {
	if o != nil && !IsNil(o.Dynamic5QISetRef) {
		return true
	}

	return false
}

// SetDynamic5QISetRef gets a reference to the given string and assigns it to the Dynamic5QISetRef field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetDynamic5QISetRef(v string) {
	o.Dynamic5QISetRef = &v
}

// GetSupportedBMOList returns the SupportedBMOList field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetSupportedBMOList() []string {
	if o == nil || IsNil(o.SupportedBMOList) {
		var ret []string
		return ret
	}
	return o.SupportedBMOList
}

// GetSupportedBMOListOk returns a tuple with the SupportedBMOList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) GetSupportedBMOListOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedBMOList) {
		return nil, false
	}
	return o.SupportedBMOList, true
}

// HasSupportedBMOList returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOfAttributesAllOf) HasSupportedBMOList() bool {
	if o != nil && !IsNil(o.SupportedBMOList) {
		return true
	}

	return false
}

// SetSupportedBMOList gets a reference to the given []string and assigns it to the SupportedBMOList field.
func (o *PcfFunctionSingleAllOfAttributesAllOf) SetSupportedBMOList(v []string) {
	o.SupportedBMOList = v
}

func (o PcfFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcfFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PLMNInfoList) {
		toSerialize["pLMNInfoList"] = o.PLMNInfoList
	}
	if !IsNil(o.SBIFqdn) {
		toSerialize["sBIFqdn"] = o.SBIFqdn
	}
	if !IsNil(o.ManagedNFProfile) {
		toSerialize["managedNFProfile"] = o.ManagedNFProfile
	}
	if !IsNil(o.CommModelList) {
		toSerialize["commModelList"] = o.CommModelList
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.DnnList) {
		toSerialize["dnnList"] = o.DnnList
	}
	if !IsNil(o.SupiRanges) {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if !IsNil(o.GpsiRanges) {
		toSerialize["gpsiRanges"] = o.GpsiRanges
	}
	if !IsNil(o.RxDiamHost) {
		toSerialize["rxDiamHost"] = o.RxDiamHost
	}
	if !IsNil(o.RxDiamRealm) {
		toSerialize["rxDiamRealm"] = o.RxDiamRealm
	}
	if !IsNil(o.V2xSupportInd) {
		toSerialize["v2xSupportInd"] = o.V2xSupportInd
	}
	if !IsNil(o.ProseSupportInd) {
		toSerialize["proseSupportInd"] = o.ProseSupportInd
	}
	if !IsNil(o.ProseCapability) {
		toSerialize["proseCapability"] = o.ProseCapability
	}
	if !IsNil(o.V2xCapability) {
		toSerialize["v2xCapability"] = o.V2xCapability
	}
	if !IsNil(o.Configurable5QISetRef) {
		toSerialize["configurable5QISetRef"] = o.Configurable5QISetRef
	}
	if !IsNil(o.Dynamic5QISetRef) {
		toSerialize["dynamic5QISetRef"] = o.Dynamic5QISetRef
	}
	if !IsNil(o.SupportedBMOList) {
		toSerialize["supportedBMOList"] = o.SupportedBMOList
	}
	return toSerialize, nil
}

type NullablePcfFunctionSingleAllOfAttributesAllOf struct {
	value *PcfFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullablePcfFunctionSingleAllOfAttributesAllOf) Get() *PcfFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullablePcfFunctionSingleAllOfAttributesAllOf) Set(val *PcfFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfFunctionSingleAllOfAttributesAllOf(val *PcfFunctionSingleAllOfAttributesAllOf) *NullablePcfFunctionSingleAllOfAttributesAllOf {
	return &NullablePcfFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullablePcfFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

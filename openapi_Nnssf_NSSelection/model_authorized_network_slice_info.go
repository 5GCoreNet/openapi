/*
NSSF NS Selection

NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssf_NSSelection

import (
	"encoding/json"
)

// checks if the AuthorizedNetworkSliceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizedNetworkSliceInfo{}

// AuthorizedNetworkSliceInfo Contains the authorized network slice information
type AuthorizedNetworkSliceInfo struct {
	AllowedNssaiList []AllowedNssai `json:"allowedNssaiList,omitempty"`
	ConfiguredNssai []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	TargetAmfSet *string `json:"targetAmfSet,omitempty"`
	CandidateAmfList []string `json:"candidateAmfList,omitempty"`
	RejectedNssaiInPlmn []Snssai `json:"rejectedNssaiInPlmn,omitempty"`
	RejectedNssaiInTa []Snssai `json:"rejectedNssaiInTa,omitempty"`
	NsiInformation *NsiInformation `json:"nsiInformation,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NrfAmfSet *string `json:"nrfAmfSet,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NrfAmfSetNfMgtUri *string `json:"nrfAmfSetNfMgtUri,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NrfAmfSetAccessTokenUri *string `json:"nrfAmfSetAccessTokenUri,omitempty"`
	// Map indicating whether the NRF requires Oauth2-based authorization for accessing its services. The key of the map shall be the name of an NRF service, e.g. \"nnrf-nfm\" or \"nnrf-disc\" 
	NrfOauth2Required *map[string]bool `json:"nrfOauth2Required,omitempty"`
	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\"Nid\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit. 
	TargetAmfServiceSet *string `json:"targetAmfServiceSet,omitempty"`
	TargetNssai []Snssai `json:"targetNssai,omitempty"`
	NsagInfos []NsagInfo `json:"nsagInfos,omitempty"`
}

// NewAuthorizedNetworkSliceInfo instantiates a new AuthorizedNetworkSliceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizedNetworkSliceInfo() *AuthorizedNetworkSliceInfo {
	this := AuthorizedNetworkSliceInfo{}
	return &this
}

// NewAuthorizedNetworkSliceInfoWithDefaults instantiates a new AuthorizedNetworkSliceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizedNetworkSliceInfoWithDefaults() *AuthorizedNetworkSliceInfo {
	this := AuthorizedNetworkSliceInfo{}
	return &this
}

// GetAllowedNssaiList returns the AllowedNssaiList field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetAllowedNssaiList() []AllowedNssai {
	if o == nil || isNil(o.AllowedNssaiList) {
		var ret []AllowedNssai
		return ret
	}
	return o.AllowedNssaiList
}

// GetAllowedNssaiListOk returns a tuple with the AllowedNssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetAllowedNssaiListOk() ([]AllowedNssai, bool) {
	if o == nil || isNil(o.AllowedNssaiList) {
		return nil, false
	}
	return o.AllowedNssaiList, true
}

// HasAllowedNssaiList returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasAllowedNssaiList() bool {
	if o != nil && !isNil(o.AllowedNssaiList) {
		return true
	}

	return false
}

// SetAllowedNssaiList gets a reference to the given []AllowedNssai and assigns it to the AllowedNssaiList field.
func (o *AuthorizedNetworkSliceInfo) SetAllowedNssaiList(v []AllowedNssai) {
	o.AllowedNssaiList = v
}

// GetConfiguredNssai returns the ConfiguredNssai field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetConfiguredNssai() []ConfiguredSnssai {
	if o == nil || isNil(o.ConfiguredNssai) {
		var ret []ConfiguredSnssai
		return ret
	}
	return o.ConfiguredNssai
}

// GetConfiguredNssaiOk returns a tuple with the ConfiguredNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetConfiguredNssaiOk() ([]ConfiguredSnssai, bool) {
	if o == nil || isNil(o.ConfiguredNssai) {
		return nil, false
	}
	return o.ConfiguredNssai, true
}

// HasConfiguredNssai returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasConfiguredNssai() bool {
	if o != nil && !isNil(o.ConfiguredNssai) {
		return true
	}

	return false
}

// SetConfiguredNssai gets a reference to the given []ConfiguredSnssai and assigns it to the ConfiguredNssai field.
func (o *AuthorizedNetworkSliceInfo) SetConfiguredNssai(v []ConfiguredSnssai) {
	o.ConfiguredNssai = v
}

// GetTargetAmfSet returns the TargetAmfSet field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetTargetAmfSet() string {
	if o == nil || isNil(o.TargetAmfSet) {
		var ret string
		return ret
	}
	return *o.TargetAmfSet
}

// GetTargetAmfSetOk returns a tuple with the TargetAmfSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetTargetAmfSetOk() (*string, bool) {
	if o == nil || isNil(o.TargetAmfSet) {
		return nil, false
	}
	return o.TargetAmfSet, true
}

// HasTargetAmfSet returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasTargetAmfSet() bool {
	if o != nil && !isNil(o.TargetAmfSet) {
		return true
	}

	return false
}

// SetTargetAmfSet gets a reference to the given string and assigns it to the TargetAmfSet field.
func (o *AuthorizedNetworkSliceInfo) SetTargetAmfSet(v string) {
	o.TargetAmfSet = &v
}

// GetCandidateAmfList returns the CandidateAmfList field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetCandidateAmfList() []string {
	if o == nil || isNil(o.CandidateAmfList) {
		var ret []string
		return ret
	}
	return o.CandidateAmfList
}

// GetCandidateAmfListOk returns a tuple with the CandidateAmfList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetCandidateAmfListOk() ([]string, bool) {
	if o == nil || isNil(o.CandidateAmfList) {
		return nil, false
	}
	return o.CandidateAmfList, true
}

// HasCandidateAmfList returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasCandidateAmfList() bool {
	if o != nil && !isNil(o.CandidateAmfList) {
		return true
	}

	return false
}

// SetCandidateAmfList gets a reference to the given []string and assigns it to the CandidateAmfList field.
func (o *AuthorizedNetworkSliceInfo) SetCandidateAmfList(v []string) {
	o.CandidateAmfList = v
}

// GetRejectedNssaiInPlmn returns the RejectedNssaiInPlmn field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetRejectedNssaiInPlmn() []Snssai {
	if o == nil || isNil(o.RejectedNssaiInPlmn) {
		var ret []Snssai
		return ret
	}
	return o.RejectedNssaiInPlmn
}

// GetRejectedNssaiInPlmnOk returns a tuple with the RejectedNssaiInPlmn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetRejectedNssaiInPlmnOk() ([]Snssai, bool) {
	if o == nil || isNil(o.RejectedNssaiInPlmn) {
		return nil, false
	}
	return o.RejectedNssaiInPlmn, true
}

// HasRejectedNssaiInPlmn returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasRejectedNssaiInPlmn() bool {
	if o != nil && !isNil(o.RejectedNssaiInPlmn) {
		return true
	}

	return false
}

// SetRejectedNssaiInPlmn gets a reference to the given []Snssai and assigns it to the RejectedNssaiInPlmn field.
func (o *AuthorizedNetworkSliceInfo) SetRejectedNssaiInPlmn(v []Snssai) {
	o.RejectedNssaiInPlmn = v
}

// GetRejectedNssaiInTa returns the RejectedNssaiInTa field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetRejectedNssaiInTa() []Snssai {
	if o == nil || isNil(o.RejectedNssaiInTa) {
		var ret []Snssai
		return ret
	}
	return o.RejectedNssaiInTa
}

// GetRejectedNssaiInTaOk returns a tuple with the RejectedNssaiInTa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetRejectedNssaiInTaOk() ([]Snssai, bool) {
	if o == nil || isNil(o.RejectedNssaiInTa) {
		return nil, false
	}
	return o.RejectedNssaiInTa, true
}

// HasRejectedNssaiInTa returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasRejectedNssaiInTa() bool {
	if o != nil && !isNil(o.RejectedNssaiInTa) {
		return true
	}

	return false
}

// SetRejectedNssaiInTa gets a reference to the given []Snssai and assigns it to the RejectedNssaiInTa field.
func (o *AuthorizedNetworkSliceInfo) SetRejectedNssaiInTa(v []Snssai) {
	o.RejectedNssaiInTa = v
}

// GetNsiInformation returns the NsiInformation field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetNsiInformation() NsiInformation {
	if o == nil || isNil(o.NsiInformation) {
		var ret NsiInformation
		return ret
	}
	return *o.NsiInformation
}

// GetNsiInformationOk returns a tuple with the NsiInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetNsiInformationOk() (*NsiInformation, bool) {
	if o == nil || isNil(o.NsiInformation) {
		return nil, false
	}
	return o.NsiInformation, true
}

// HasNsiInformation returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasNsiInformation() bool {
	if o != nil && !isNil(o.NsiInformation) {
		return true
	}

	return false
}

// SetNsiInformation gets a reference to the given NsiInformation and assigns it to the NsiInformation field.
func (o *AuthorizedNetworkSliceInfo) SetNsiInformation(v NsiInformation) {
	o.NsiInformation = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *AuthorizedNetworkSliceInfo) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetNrfAmfSet returns the NrfAmfSet field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetNrfAmfSet() string {
	if o == nil || isNil(o.NrfAmfSet) {
		var ret string
		return ret
	}
	return *o.NrfAmfSet
}

// GetNrfAmfSetOk returns a tuple with the NrfAmfSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetNrfAmfSetOk() (*string, bool) {
	if o == nil || isNil(o.NrfAmfSet) {
		return nil, false
	}
	return o.NrfAmfSet, true
}

// HasNrfAmfSet returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasNrfAmfSet() bool {
	if o != nil && !isNil(o.NrfAmfSet) {
		return true
	}

	return false
}

// SetNrfAmfSet gets a reference to the given string and assigns it to the NrfAmfSet field.
func (o *AuthorizedNetworkSliceInfo) SetNrfAmfSet(v string) {
	o.NrfAmfSet = &v
}

// GetNrfAmfSetNfMgtUri returns the NrfAmfSetNfMgtUri field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetNrfAmfSetNfMgtUri() string {
	if o == nil || isNil(o.NrfAmfSetNfMgtUri) {
		var ret string
		return ret
	}
	return *o.NrfAmfSetNfMgtUri
}

// GetNrfAmfSetNfMgtUriOk returns a tuple with the NrfAmfSetNfMgtUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetNrfAmfSetNfMgtUriOk() (*string, bool) {
	if o == nil || isNil(o.NrfAmfSetNfMgtUri) {
		return nil, false
	}
	return o.NrfAmfSetNfMgtUri, true
}

// HasNrfAmfSetNfMgtUri returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasNrfAmfSetNfMgtUri() bool {
	if o != nil && !isNil(o.NrfAmfSetNfMgtUri) {
		return true
	}

	return false
}

// SetNrfAmfSetNfMgtUri gets a reference to the given string and assigns it to the NrfAmfSetNfMgtUri field.
func (o *AuthorizedNetworkSliceInfo) SetNrfAmfSetNfMgtUri(v string) {
	o.NrfAmfSetNfMgtUri = &v
}

// GetNrfAmfSetAccessTokenUri returns the NrfAmfSetAccessTokenUri field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetNrfAmfSetAccessTokenUri() string {
	if o == nil || isNil(o.NrfAmfSetAccessTokenUri) {
		var ret string
		return ret
	}
	return *o.NrfAmfSetAccessTokenUri
}

// GetNrfAmfSetAccessTokenUriOk returns a tuple with the NrfAmfSetAccessTokenUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetNrfAmfSetAccessTokenUriOk() (*string, bool) {
	if o == nil || isNil(o.NrfAmfSetAccessTokenUri) {
		return nil, false
	}
	return o.NrfAmfSetAccessTokenUri, true
}

// HasNrfAmfSetAccessTokenUri returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasNrfAmfSetAccessTokenUri() bool {
	if o != nil && !isNil(o.NrfAmfSetAccessTokenUri) {
		return true
	}

	return false
}

// SetNrfAmfSetAccessTokenUri gets a reference to the given string and assigns it to the NrfAmfSetAccessTokenUri field.
func (o *AuthorizedNetworkSliceInfo) SetNrfAmfSetAccessTokenUri(v string) {
	o.NrfAmfSetAccessTokenUri = &v
}

// GetNrfOauth2Required returns the NrfOauth2Required field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetNrfOauth2Required() map[string]bool {
	if o == nil || isNil(o.NrfOauth2Required) {
		var ret map[string]bool
		return ret
	}
	return *o.NrfOauth2Required
}

// GetNrfOauth2RequiredOk returns a tuple with the NrfOauth2Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetNrfOauth2RequiredOk() (*map[string]bool, bool) {
	if o == nil || isNil(o.NrfOauth2Required) {
		return nil, false
	}
	return o.NrfOauth2Required, true
}

// HasNrfOauth2Required returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasNrfOauth2Required() bool {
	if o != nil && !isNil(o.NrfOauth2Required) {
		return true
	}

	return false
}

// SetNrfOauth2Required gets a reference to the given map[string]bool and assigns it to the NrfOauth2Required field.
func (o *AuthorizedNetworkSliceInfo) SetNrfOauth2Required(v map[string]bool) {
	o.NrfOauth2Required = &v
}

// GetTargetAmfServiceSet returns the TargetAmfServiceSet field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetTargetAmfServiceSet() string {
	if o == nil || isNil(o.TargetAmfServiceSet) {
		var ret string
		return ret
	}
	return *o.TargetAmfServiceSet
}

// GetTargetAmfServiceSetOk returns a tuple with the TargetAmfServiceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetTargetAmfServiceSetOk() (*string, bool) {
	if o == nil || isNil(o.TargetAmfServiceSet) {
		return nil, false
	}
	return o.TargetAmfServiceSet, true
}

// HasTargetAmfServiceSet returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasTargetAmfServiceSet() bool {
	if o != nil && !isNil(o.TargetAmfServiceSet) {
		return true
	}

	return false
}

// SetTargetAmfServiceSet gets a reference to the given string and assigns it to the TargetAmfServiceSet field.
func (o *AuthorizedNetworkSliceInfo) SetTargetAmfServiceSet(v string) {
	o.TargetAmfServiceSet = &v
}

// GetTargetNssai returns the TargetNssai field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetTargetNssai() []Snssai {
	if o == nil || isNil(o.TargetNssai) {
		var ret []Snssai
		return ret
	}
	return o.TargetNssai
}

// GetTargetNssaiOk returns a tuple with the TargetNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetTargetNssaiOk() ([]Snssai, bool) {
	if o == nil || isNil(o.TargetNssai) {
		return nil, false
	}
	return o.TargetNssai, true
}

// HasTargetNssai returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasTargetNssai() bool {
	if o != nil && !isNil(o.TargetNssai) {
		return true
	}

	return false
}

// SetTargetNssai gets a reference to the given []Snssai and assigns it to the TargetNssai field.
func (o *AuthorizedNetworkSliceInfo) SetTargetNssai(v []Snssai) {
	o.TargetNssai = v
}

// GetNsagInfos returns the NsagInfos field value if set, zero value otherwise.
func (o *AuthorizedNetworkSliceInfo) GetNsagInfos() []NsagInfo {
	if o == nil || isNil(o.NsagInfos) {
		var ret []NsagInfo
		return ret
	}
	return o.NsagInfos
}

// GetNsagInfosOk returns a tuple with the NsagInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNetworkSliceInfo) GetNsagInfosOk() ([]NsagInfo, bool) {
	if o == nil || isNil(o.NsagInfos) {
		return nil, false
	}
	return o.NsagInfos, true
}

// HasNsagInfos returns a boolean if a field has been set.
func (o *AuthorizedNetworkSliceInfo) HasNsagInfos() bool {
	if o != nil && !isNil(o.NsagInfos) {
		return true
	}

	return false
}

// SetNsagInfos gets a reference to the given []NsagInfo and assigns it to the NsagInfos field.
func (o *AuthorizedNetworkSliceInfo) SetNsagInfos(v []NsagInfo) {
	o.NsagInfos = v
}

func (o AuthorizedNetworkSliceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizedNetworkSliceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowedNssaiList) {
		toSerialize["allowedNssaiList"] = o.AllowedNssaiList
	}
	if !isNil(o.ConfiguredNssai) {
		toSerialize["configuredNssai"] = o.ConfiguredNssai
	}
	if !isNil(o.TargetAmfSet) {
		toSerialize["targetAmfSet"] = o.TargetAmfSet
	}
	if !isNil(o.CandidateAmfList) {
		toSerialize["candidateAmfList"] = o.CandidateAmfList
	}
	if !isNil(o.RejectedNssaiInPlmn) {
		toSerialize["rejectedNssaiInPlmn"] = o.RejectedNssaiInPlmn
	}
	if !isNil(o.RejectedNssaiInTa) {
		toSerialize["rejectedNssaiInTa"] = o.RejectedNssaiInTa
	}
	if !isNil(o.NsiInformation) {
		toSerialize["nsiInformation"] = o.NsiInformation
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.NrfAmfSet) {
		toSerialize["nrfAmfSet"] = o.NrfAmfSet
	}
	if !isNil(o.NrfAmfSetNfMgtUri) {
		toSerialize["nrfAmfSetNfMgtUri"] = o.NrfAmfSetNfMgtUri
	}
	if !isNil(o.NrfAmfSetAccessTokenUri) {
		toSerialize["nrfAmfSetAccessTokenUri"] = o.NrfAmfSetAccessTokenUri
	}
	if !isNil(o.NrfOauth2Required) {
		toSerialize["nrfOauth2Required"] = o.NrfOauth2Required
	}
	if !isNil(o.TargetAmfServiceSet) {
		toSerialize["targetAmfServiceSet"] = o.TargetAmfServiceSet
	}
	if !isNil(o.TargetNssai) {
		toSerialize["targetNssai"] = o.TargetNssai
	}
	if !isNil(o.NsagInfos) {
		toSerialize["nsagInfos"] = o.NsagInfos
	}
	return toSerialize, nil
}

type NullableAuthorizedNetworkSliceInfo struct {
	value *AuthorizedNetworkSliceInfo
	isSet bool
}

func (v NullableAuthorizedNetworkSliceInfo) Get() *AuthorizedNetworkSliceInfo {
	return v.value
}

func (v *NullableAuthorizedNetworkSliceInfo) Set(val *AuthorizedNetworkSliceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizedNetworkSliceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizedNetworkSliceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizedNetworkSliceInfo(val *AuthorizedNetworkSliceInfo) *NullableAuthorizedNetworkSliceInfo {
	return &NullableAuthorizedNetworkSliceInfo{value: val, isSet: true}
}

func (v NullableAuthorizedNetworkSliceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizedNetworkSliceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
)

// checks if the AccessTokenReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenReq{}

// AccessTokenReq Contains information related to the access token request
type AccessTokenReq struct {
	GrantType string `json:"grant_type"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId string `json:"nfInstanceId"`
	NfType *NFType `json:"nfType,omitempty"`
	TargetNfType *NFType `json:"targetNfType,omitempty"`
	Scope string `json:"scope"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	TargetNfInstanceId *string `json:"targetNfInstanceId,omitempty"`
	RequesterPlmn *PlmnId `json:"requesterPlmn,omitempty"`
	RequesterPlmnList []PlmnId `json:"requesterPlmnList,omitempty"`
	RequesterSnssaiList []Snssai `json:"requesterSnssaiList,omitempty"`
	// Fully Qualified Domain Name
	RequesterFqdn *string `json:"requesterFqdn,omitempty"`
	RequesterSnpnList []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetPlmn *PlmnId `json:"targetPlmn,omitempty"`
	TargetSnpn *PlmnIdNid `json:"targetSnpn,omitempty"`
	TargetSnssaiList []Snssai `json:"targetSnssaiList,omitempty"`
	TargetNsiList []string `json:"targetNsiList,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	TargetNfSetId *string `json:"targetNfSetId,omitempty"`
	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\"Nid\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit. 
	TargetNfServiceSetId *string `json:"targetNfServiceSetId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	HnrfAccessTokenUri *string `json:"hnrfAccessTokenUri,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SourceNfInstanceId *string `json:"sourceNfInstanceId,omitempty"`
}

// NewAccessTokenReq instantiates a new AccessTokenReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenReq(grantType string, nfInstanceId string, scope string) *AccessTokenReq {
	this := AccessTokenReq{}
	this.GrantType = grantType
	this.NfInstanceId = nfInstanceId
	this.Scope = scope
	return &this
}

// NewAccessTokenReqWithDefaults instantiates a new AccessTokenReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenReqWithDefaults() *AccessTokenReq {
	this := AccessTokenReq{}
	return &this
}

// GetGrantType returns the GrantType field value
func (o *AccessTokenReq) GetGrantType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetGrantTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *AccessTokenReq) SetGrantType(v string) {
	o.GrantType = v
}

// GetNfInstanceId returns the NfInstanceId field value
func (o *AccessTokenReq) GetNfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfInstanceId
}

// GetNfInstanceIdOk returns a tuple with the NfInstanceId field value
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetNfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfInstanceId, true
}

// SetNfInstanceId sets field value
func (o *AccessTokenReq) SetNfInstanceId(v string) {
	o.NfInstanceId = v
}

// GetNfType returns the NfType field value if set, zero value otherwise.
func (o *AccessTokenReq) GetNfType() NFType {
	if o == nil || isNil(o.NfType) {
		var ret NFType
		return ret
	}
	return *o.NfType
}

// GetNfTypeOk returns a tuple with the NfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetNfTypeOk() (*NFType, bool) {
	if o == nil || isNil(o.NfType) {
		return nil, false
	}
	return o.NfType, true
}

// HasNfType returns a boolean if a field has been set.
func (o *AccessTokenReq) HasNfType() bool {
	if o != nil && !isNil(o.NfType) {
		return true
	}

	return false
}

// SetNfType gets a reference to the given NFType and assigns it to the NfType field.
func (o *AccessTokenReq) SetNfType(v NFType) {
	o.NfType = &v
}

// GetTargetNfType returns the TargetNfType field value if set, zero value otherwise.
func (o *AccessTokenReq) GetTargetNfType() NFType {
	if o == nil || isNil(o.TargetNfType) {
		var ret NFType
		return ret
	}
	return *o.TargetNfType
}

// GetTargetNfTypeOk returns a tuple with the TargetNfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetTargetNfTypeOk() (*NFType, bool) {
	if o == nil || isNil(o.TargetNfType) {
		return nil, false
	}
	return o.TargetNfType, true
}

// HasTargetNfType returns a boolean if a field has been set.
func (o *AccessTokenReq) HasTargetNfType() bool {
	if o != nil && !isNil(o.TargetNfType) {
		return true
	}

	return false
}

// SetTargetNfType gets a reference to the given NFType and assigns it to the TargetNfType field.
func (o *AccessTokenReq) SetTargetNfType(v NFType) {
	o.TargetNfType = &v
}

// GetScope returns the Scope field value
func (o *AccessTokenReq) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *AccessTokenReq) SetScope(v string) {
	o.Scope = v
}

// GetTargetNfInstanceId returns the TargetNfInstanceId field value if set, zero value otherwise.
func (o *AccessTokenReq) GetTargetNfInstanceId() string {
	if o == nil || isNil(o.TargetNfInstanceId) {
		var ret string
		return ret
	}
	return *o.TargetNfInstanceId
}

// GetTargetNfInstanceIdOk returns a tuple with the TargetNfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetTargetNfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.TargetNfInstanceId) {
		return nil, false
	}
	return o.TargetNfInstanceId, true
}

// HasTargetNfInstanceId returns a boolean if a field has been set.
func (o *AccessTokenReq) HasTargetNfInstanceId() bool {
	if o != nil && !isNil(o.TargetNfInstanceId) {
		return true
	}

	return false
}

// SetTargetNfInstanceId gets a reference to the given string and assigns it to the TargetNfInstanceId field.
func (o *AccessTokenReq) SetTargetNfInstanceId(v string) {
	o.TargetNfInstanceId = &v
}

// GetRequesterPlmn returns the RequesterPlmn field value if set, zero value otherwise.
func (o *AccessTokenReq) GetRequesterPlmn() PlmnId {
	if o == nil || isNil(o.RequesterPlmn) {
		var ret PlmnId
		return ret
	}
	return *o.RequesterPlmn
}

// GetRequesterPlmnOk returns a tuple with the RequesterPlmn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetRequesterPlmnOk() (*PlmnId, bool) {
	if o == nil || isNil(o.RequesterPlmn) {
		return nil, false
	}
	return o.RequesterPlmn, true
}

// HasRequesterPlmn returns a boolean if a field has been set.
func (o *AccessTokenReq) HasRequesterPlmn() bool {
	if o != nil && !isNil(o.RequesterPlmn) {
		return true
	}

	return false
}

// SetRequesterPlmn gets a reference to the given PlmnId and assigns it to the RequesterPlmn field.
func (o *AccessTokenReq) SetRequesterPlmn(v PlmnId) {
	o.RequesterPlmn = &v
}

// GetRequesterPlmnList returns the RequesterPlmnList field value if set, zero value otherwise.
func (o *AccessTokenReq) GetRequesterPlmnList() []PlmnId {
	if o == nil || isNil(o.RequesterPlmnList) {
		var ret []PlmnId
		return ret
	}
	return o.RequesterPlmnList
}

// GetRequesterPlmnListOk returns a tuple with the RequesterPlmnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetRequesterPlmnListOk() ([]PlmnId, bool) {
	if o == nil || isNil(o.RequesterPlmnList) {
		return nil, false
	}
	return o.RequesterPlmnList, true
}

// HasRequesterPlmnList returns a boolean if a field has been set.
func (o *AccessTokenReq) HasRequesterPlmnList() bool {
	if o != nil && !isNil(o.RequesterPlmnList) {
		return true
	}

	return false
}

// SetRequesterPlmnList gets a reference to the given []PlmnId and assigns it to the RequesterPlmnList field.
func (o *AccessTokenReq) SetRequesterPlmnList(v []PlmnId) {
	o.RequesterPlmnList = v
}

// GetRequesterSnssaiList returns the RequesterSnssaiList field value if set, zero value otherwise.
func (o *AccessTokenReq) GetRequesterSnssaiList() []Snssai {
	if o == nil || isNil(o.RequesterSnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.RequesterSnssaiList
}

// GetRequesterSnssaiListOk returns a tuple with the RequesterSnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetRequesterSnssaiListOk() ([]Snssai, bool) {
	if o == nil || isNil(o.RequesterSnssaiList) {
		return nil, false
	}
	return o.RequesterSnssaiList, true
}

// HasRequesterSnssaiList returns a boolean if a field has been set.
func (o *AccessTokenReq) HasRequesterSnssaiList() bool {
	if o != nil && !isNil(o.RequesterSnssaiList) {
		return true
	}

	return false
}

// SetRequesterSnssaiList gets a reference to the given []Snssai and assigns it to the RequesterSnssaiList field.
func (o *AccessTokenReq) SetRequesterSnssaiList(v []Snssai) {
	o.RequesterSnssaiList = v
}

// GetRequesterFqdn returns the RequesterFqdn field value if set, zero value otherwise.
func (o *AccessTokenReq) GetRequesterFqdn() string {
	if o == nil || isNil(o.RequesterFqdn) {
		var ret string
		return ret
	}
	return *o.RequesterFqdn
}

// GetRequesterFqdnOk returns a tuple with the RequesterFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetRequesterFqdnOk() (*string, bool) {
	if o == nil || isNil(o.RequesterFqdn) {
		return nil, false
	}
	return o.RequesterFqdn, true
}

// HasRequesterFqdn returns a boolean if a field has been set.
func (o *AccessTokenReq) HasRequesterFqdn() bool {
	if o != nil && !isNil(o.RequesterFqdn) {
		return true
	}

	return false
}

// SetRequesterFqdn gets a reference to the given string and assigns it to the RequesterFqdn field.
func (o *AccessTokenReq) SetRequesterFqdn(v string) {
	o.RequesterFqdn = &v
}

// GetRequesterSnpnList returns the RequesterSnpnList field value if set, zero value otherwise.
func (o *AccessTokenReq) GetRequesterSnpnList() []PlmnIdNid {
	if o == nil || isNil(o.RequesterSnpnList) {
		var ret []PlmnIdNid
		return ret
	}
	return o.RequesterSnpnList
}

// GetRequesterSnpnListOk returns a tuple with the RequesterSnpnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetRequesterSnpnListOk() ([]PlmnIdNid, bool) {
	if o == nil || isNil(o.RequesterSnpnList) {
		return nil, false
	}
	return o.RequesterSnpnList, true
}

// HasRequesterSnpnList returns a boolean if a field has been set.
func (o *AccessTokenReq) HasRequesterSnpnList() bool {
	if o != nil && !isNil(o.RequesterSnpnList) {
		return true
	}

	return false
}

// SetRequesterSnpnList gets a reference to the given []PlmnIdNid and assigns it to the RequesterSnpnList field.
func (o *AccessTokenReq) SetRequesterSnpnList(v []PlmnIdNid) {
	o.RequesterSnpnList = v
}

// GetTargetPlmn returns the TargetPlmn field value if set, zero value otherwise.
func (o *AccessTokenReq) GetTargetPlmn() PlmnId {
	if o == nil || isNil(o.TargetPlmn) {
		var ret PlmnId
		return ret
	}
	return *o.TargetPlmn
}

// GetTargetPlmnOk returns a tuple with the TargetPlmn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetTargetPlmnOk() (*PlmnId, bool) {
	if o == nil || isNil(o.TargetPlmn) {
		return nil, false
	}
	return o.TargetPlmn, true
}

// HasTargetPlmn returns a boolean if a field has been set.
func (o *AccessTokenReq) HasTargetPlmn() bool {
	if o != nil && !isNil(o.TargetPlmn) {
		return true
	}

	return false
}

// SetTargetPlmn gets a reference to the given PlmnId and assigns it to the TargetPlmn field.
func (o *AccessTokenReq) SetTargetPlmn(v PlmnId) {
	o.TargetPlmn = &v
}

// GetTargetSnpn returns the TargetSnpn field value if set, zero value otherwise.
func (o *AccessTokenReq) GetTargetSnpn() PlmnIdNid {
	if o == nil || isNil(o.TargetSnpn) {
		var ret PlmnIdNid
		return ret
	}
	return *o.TargetSnpn
}

// GetTargetSnpnOk returns a tuple with the TargetSnpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetTargetSnpnOk() (*PlmnIdNid, bool) {
	if o == nil || isNil(o.TargetSnpn) {
		return nil, false
	}
	return o.TargetSnpn, true
}

// HasTargetSnpn returns a boolean if a field has been set.
func (o *AccessTokenReq) HasTargetSnpn() bool {
	if o != nil && !isNil(o.TargetSnpn) {
		return true
	}

	return false
}

// SetTargetSnpn gets a reference to the given PlmnIdNid and assigns it to the TargetSnpn field.
func (o *AccessTokenReq) SetTargetSnpn(v PlmnIdNid) {
	o.TargetSnpn = &v
}

// GetTargetSnssaiList returns the TargetSnssaiList field value if set, zero value otherwise.
func (o *AccessTokenReq) GetTargetSnssaiList() []Snssai {
	if o == nil || isNil(o.TargetSnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.TargetSnssaiList
}

// GetTargetSnssaiListOk returns a tuple with the TargetSnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetTargetSnssaiListOk() ([]Snssai, bool) {
	if o == nil || isNil(o.TargetSnssaiList) {
		return nil, false
	}
	return o.TargetSnssaiList, true
}

// HasTargetSnssaiList returns a boolean if a field has been set.
func (o *AccessTokenReq) HasTargetSnssaiList() bool {
	if o != nil && !isNil(o.TargetSnssaiList) {
		return true
	}

	return false
}

// SetTargetSnssaiList gets a reference to the given []Snssai and assigns it to the TargetSnssaiList field.
func (o *AccessTokenReq) SetTargetSnssaiList(v []Snssai) {
	o.TargetSnssaiList = v
}

// GetTargetNsiList returns the TargetNsiList field value if set, zero value otherwise.
func (o *AccessTokenReq) GetTargetNsiList() []string {
	if o == nil || isNil(o.TargetNsiList) {
		var ret []string
		return ret
	}
	return o.TargetNsiList
}

// GetTargetNsiListOk returns a tuple with the TargetNsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetTargetNsiListOk() ([]string, bool) {
	if o == nil || isNil(o.TargetNsiList) {
		return nil, false
	}
	return o.TargetNsiList, true
}

// HasTargetNsiList returns a boolean if a field has been set.
func (o *AccessTokenReq) HasTargetNsiList() bool {
	if o != nil && !isNil(o.TargetNsiList) {
		return true
	}

	return false
}

// SetTargetNsiList gets a reference to the given []string and assigns it to the TargetNsiList field.
func (o *AccessTokenReq) SetTargetNsiList(v []string) {
	o.TargetNsiList = v
}

// GetTargetNfSetId returns the TargetNfSetId field value if set, zero value otherwise.
func (o *AccessTokenReq) GetTargetNfSetId() string {
	if o == nil || isNil(o.TargetNfSetId) {
		var ret string
		return ret
	}
	return *o.TargetNfSetId
}

// GetTargetNfSetIdOk returns a tuple with the TargetNfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetTargetNfSetIdOk() (*string, bool) {
	if o == nil || isNil(o.TargetNfSetId) {
		return nil, false
	}
	return o.TargetNfSetId, true
}

// HasTargetNfSetId returns a boolean if a field has been set.
func (o *AccessTokenReq) HasTargetNfSetId() bool {
	if o != nil && !isNil(o.TargetNfSetId) {
		return true
	}

	return false
}

// SetTargetNfSetId gets a reference to the given string and assigns it to the TargetNfSetId field.
func (o *AccessTokenReq) SetTargetNfSetId(v string) {
	o.TargetNfSetId = &v
}

// GetTargetNfServiceSetId returns the TargetNfServiceSetId field value if set, zero value otherwise.
func (o *AccessTokenReq) GetTargetNfServiceSetId() string {
	if o == nil || isNil(o.TargetNfServiceSetId) {
		var ret string
		return ret
	}
	return *o.TargetNfServiceSetId
}

// GetTargetNfServiceSetIdOk returns a tuple with the TargetNfServiceSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetTargetNfServiceSetIdOk() (*string, bool) {
	if o == nil || isNil(o.TargetNfServiceSetId) {
		return nil, false
	}
	return o.TargetNfServiceSetId, true
}

// HasTargetNfServiceSetId returns a boolean if a field has been set.
func (o *AccessTokenReq) HasTargetNfServiceSetId() bool {
	if o != nil && !isNil(o.TargetNfServiceSetId) {
		return true
	}

	return false
}

// SetTargetNfServiceSetId gets a reference to the given string and assigns it to the TargetNfServiceSetId field.
func (o *AccessTokenReq) SetTargetNfServiceSetId(v string) {
	o.TargetNfServiceSetId = &v
}

// GetHnrfAccessTokenUri returns the HnrfAccessTokenUri field value if set, zero value otherwise.
func (o *AccessTokenReq) GetHnrfAccessTokenUri() string {
	if o == nil || isNil(o.HnrfAccessTokenUri) {
		var ret string
		return ret
	}
	return *o.HnrfAccessTokenUri
}

// GetHnrfAccessTokenUriOk returns a tuple with the HnrfAccessTokenUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetHnrfAccessTokenUriOk() (*string, bool) {
	if o == nil || isNil(o.HnrfAccessTokenUri) {
		return nil, false
	}
	return o.HnrfAccessTokenUri, true
}

// HasHnrfAccessTokenUri returns a boolean if a field has been set.
func (o *AccessTokenReq) HasHnrfAccessTokenUri() bool {
	if o != nil && !isNil(o.HnrfAccessTokenUri) {
		return true
	}

	return false
}

// SetHnrfAccessTokenUri gets a reference to the given string and assigns it to the HnrfAccessTokenUri field.
func (o *AccessTokenReq) SetHnrfAccessTokenUri(v string) {
	o.HnrfAccessTokenUri = &v
}

// GetSourceNfInstanceId returns the SourceNfInstanceId field value if set, zero value otherwise.
func (o *AccessTokenReq) GetSourceNfInstanceId() string {
	if o == nil || isNil(o.SourceNfInstanceId) {
		var ret string
		return ret
	}
	return *o.SourceNfInstanceId
}

// GetSourceNfInstanceIdOk returns a tuple with the SourceNfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenReq) GetSourceNfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceNfInstanceId) {
		return nil, false
	}
	return o.SourceNfInstanceId, true
}

// HasSourceNfInstanceId returns a boolean if a field has been set.
func (o *AccessTokenReq) HasSourceNfInstanceId() bool {
	if o != nil && !isNil(o.SourceNfInstanceId) {
		return true
	}

	return false
}

// SetSourceNfInstanceId gets a reference to the given string and assigns it to the SourceNfInstanceId field.
func (o *AccessTokenReq) SetSourceNfInstanceId(v string) {
	o.SourceNfInstanceId = &v
}

func (o AccessTokenReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grant_type"] = o.GrantType
	toSerialize["nfInstanceId"] = o.NfInstanceId
	if !isNil(o.NfType) {
		toSerialize["nfType"] = o.NfType
	}
	if !isNil(o.TargetNfType) {
		toSerialize["targetNfType"] = o.TargetNfType
	}
	toSerialize["scope"] = o.Scope
	if !isNil(o.TargetNfInstanceId) {
		toSerialize["targetNfInstanceId"] = o.TargetNfInstanceId
	}
	if !isNil(o.RequesterPlmn) {
		toSerialize["requesterPlmn"] = o.RequesterPlmn
	}
	if !isNil(o.RequesterPlmnList) {
		toSerialize["requesterPlmnList"] = o.RequesterPlmnList
	}
	if !isNil(o.RequesterSnssaiList) {
		toSerialize["requesterSnssaiList"] = o.RequesterSnssaiList
	}
	if !isNil(o.RequesterFqdn) {
		toSerialize["requesterFqdn"] = o.RequesterFqdn
	}
	if !isNil(o.RequesterSnpnList) {
		toSerialize["requesterSnpnList"] = o.RequesterSnpnList
	}
	if !isNil(o.TargetPlmn) {
		toSerialize["targetPlmn"] = o.TargetPlmn
	}
	if !isNil(o.TargetSnpn) {
		toSerialize["targetSnpn"] = o.TargetSnpn
	}
	if !isNil(o.TargetSnssaiList) {
		toSerialize["targetSnssaiList"] = o.TargetSnssaiList
	}
	if !isNil(o.TargetNsiList) {
		toSerialize["targetNsiList"] = o.TargetNsiList
	}
	if !isNil(o.TargetNfSetId) {
		toSerialize["targetNfSetId"] = o.TargetNfSetId
	}
	if !isNil(o.TargetNfServiceSetId) {
		toSerialize["targetNfServiceSetId"] = o.TargetNfServiceSetId
	}
	if !isNil(o.HnrfAccessTokenUri) {
		toSerialize["hnrfAccessTokenUri"] = o.HnrfAccessTokenUri
	}
	if !isNil(o.SourceNfInstanceId) {
		toSerialize["sourceNfInstanceId"] = o.SourceNfInstanceId
	}
	return toSerialize, nil
}

type NullableAccessTokenReq struct {
	value *AccessTokenReq
	isSet bool
}

func (v NullableAccessTokenReq) Get() *AccessTokenReq {
	return v.value
}

func (v *NullableAccessTokenReq) Set(val *AccessTokenReq) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenReq) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenReq(val *AccessTokenReq) *NullableAccessTokenReq {
	return &NullableAccessTokenReq{value: val, isSet: true}
}

func (v NullableAccessTokenReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



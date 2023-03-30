/*
NRF OAuth2

NRF OAuth2 Authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_AccessToken

import (
	"encoding/json"
)

// checks if the AccessTokenClaims type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenClaims{}

// AccessTokenClaims The claims data structure for the access token
type AccessTokenClaims struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	Iss string `json:"iss"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	Sub string `json:"sub"`
	Aud AccessTokenClaimsAud `json:"aud"`
	Scope string `json:"scope"`
	Exp int32 `json:"exp"`
	ConsumerPlmnId *PlmnId `json:"consumerPlmnId,omitempty"`
	ConsumerSnpnId *PlmnIdNid `json:"consumerSnpnId,omitempty"`
	ProducerPlmnId *PlmnId `json:"producerPlmnId,omitempty"`
	ProducerSnpnId *PlmnIdNid `json:"producerSnpnId,omitempty"`
	ProducerSnssaiList []Snssai `json:"producerSnssaiList,omitempty"`
	ProducerNsiList []string `json:"producerNsiList,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	ProducerNfSetId *string `json:"producerNfSetId,omitempty"`
	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\"Nid\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit. 
	ProducerNfServiceSetId *string `json:"producerNfServiceSetId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SourceNfInstanceId *string `json:"sourceNfInstanceId,omitempty"`
}

// NewAccessTokenClaims instantiates a new AccessTokenClaims object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenClaims(iss string, sub string, aud AccessTokenClaimsAud, scope string, exp int32) *AccessTokenClaims {
	this := AccessTokenClaims{}
	this.Iss = iss
	this.Sub = sub
	this.Aud = aud
	this.Scope = scope
	this.Exp = exp
	return &this
}

// NewAccessTokenClaimsWithDefaults instantiates a new AccessTokenClaims object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenClaimsWithDefaults() *AccessTokenClaims {
	this := AccessTokenClaims{}
	return &this
}

// GetIss returns the Iss field value
func (o *AccessTokenClaims) GetIss() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iss
}

// GetIssOk returns a tuple with the Iss field value
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetIssOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iss, true
}

// SetIss sets field value
func (o *AccessTokenClaims) SetIss(v string) {
	o.Iss = v
}

// GetSub returns the Sub field value
func (o *AccessTokenClaims) GetSub() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sub
}

// GetSubOk returns a tuple with the Sub field value
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetSubOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sub, true
}

// SetSub sets field value
func (o *AccessTokenClaims) SetSub(v string) {
	o.Sub = v
}

// GetAud returns the Aud field value
func (o *AccessTokenClaims) GetAud() AccessTokenClaimsAud {
	if o == nil {
		var ret AccessTokenClaimsAud
		return ret
	}

	return o.Aud
}

// GetAudOk returns a tuple with the Aud field value
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetAudOk() (*AccessTokenClaimsAud, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aud, true
}

// SetAud sets field value
func (o *AccessTokenClaims) SetAud(v AccessTokenClaimsAud) {
	o.Aud = v
}

// GetScope returns the Scope field value
func (o *AccessTokenClaims) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *AccessTokenClaims) SetScope(v string) {
	o.Scope = v
}

// GetExp returns the Exp field value
func (o *AccessTokenClaims) GetExp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Exp
}

// GetExpOk returns a tuple with the Exp field value
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetExpOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exp, true
}

// SetExp sets field value
func (o *AccessTokenClaims) SetExp(v int32) {
	o.Exp = v
}

// GetConsumerPlmnId returns the ConsumerPlmnId field value if set, zero value otherwise.
func (o *AccessTokenClaims) GetConsumerPlmnId() PlmnId {
	if o == nil || IsNil(o.ConsumerPlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.ConsumerPlmnId
}

// GetConsumerPlmnIdOk returns a tuple with the ConsumerPlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetConsumerPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.ConsumerPlmnId) {
		return nil, false
	}
	return o.ConsumerPlmnId, true
}

// HasConsumerPlmnId returns a boolean if a field has been set.
func (o *AccessTokenClaims) HasConsumerPlmnId() bool {
	if o != nil && !IsNil(o.ConsumerPlmnId) {
		return true
	}

	return false
}

// SetConsumerPlmnId gets a reference to the given PlmnId and assigns it to the ConsumerPlmnId field.
func (o *AccessTokenClaims) SetConsumerPlmnId(v PlmnId) {
	o.ConsumerPlmnId = &v
}

// GetConsumerSnpnId returns the ConsumerSnpnId field value if set, zero value otherwise.
func (o *AccessTokenClaims) GetConsumerSnpnId() PlmnIdNid {
	if o == nil || IsNil(o.ConsumerSnpnId) {
		var ret PlmnIdNid
		return ret
	}
	return *o.ConsumerSnpnId
}

// GetConsumerSnpnIdOk returns a tuple with the ConsumerSnpnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetConsumerSnpnIdOk() (*PlmnIdNid, bool) {
	if o == nil || IsNil(o.ConsumerSnpnId) {
		return nil, false
	}
	return o.ConsumerSnpnId, true
}

// HasConsumerSnpnId returns a boolean if a field has been set.
func (o *AccessTokenClaims) HasConsumerSnpnId() bool {
	if o != nil && !IsNil(o.ConsumerSnpnId) {
		return true
	}

	return false
}

// SetConsumerSnpnId gets a reference to the given PlmnIdNid and assigns it to the ConsumerSnpnId field.
func (o *AccessTokenClaims) SetConsumerSnpnId(v PlmnIdNid) {
	o.ConsumerSnpnId = &v
}

// GetProducerPlmnId returns the ProducerPlmnId field value if set, zero value otherwise.
func (o *AccessTokenClaims) GetProducerPlmnId() PlmnId {
	if o == nil || IsNil(o.ProducerPlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.ProducerPlmnId
}

// GetProducerPlmnIdOk returns a tuple with the ProducerPlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetProducerPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.ProducerPlmnId) {
		return nil, false
	}
	return o.ProducerPlmnId, true
}

// HasProducerPlmnId returns a boolean if a field has been set.
func (o *AccessTokenClaims) HasProducerPlmnId() bool {
	if o != nil && !IsNil(o.ProducerPlmnId) {
		return true
	}

	return false
}

// SetProducerPlmnId gets a reference to the given PlmnId and assigns it to the ProducerPlmnId field.
func (o *AccessTokenClaims) SetProducerPlmnId(v PlmnId) {
	o.ProducerPlmnId = &v
}

// GetProducerSnpnId returns the ProducerSnpnId field value if set, zero value otherwise.
func (o *AccessTokenClaims) GetProducerSnpnId() PlmnIdNid {
	if o == nil || IsNil(o.ProducerSnpnId) {
		var ret PlmnIdNid
		return ret
	}
	return *o.ProducerSnpnId
}

// GetProducerSnpnIdOk returns a tuple with the ProducerSnpnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetProducerSnpnIdOk() (*PlmnIdNid, bool) {
	if o == nil || IsNil(o.ProducerSnpnId) {
		return nil, false
	}
	return o.ProducerSnpnId, true
}

// HasProducerSnpnId returns a boolean if a field has been set.
func (o *AccessTokenClaims) HasProducerSnpnId() bool {
	if o != nil && !IsNil(o.ProducerSnpnId) {
		return true
	}

	return false
}

// SetProducerSnpnId gets a reference to the given PlmnIdNid and assigns it to the ProducerSnpnId field.
func (o *AccessTokenClaims) SetProducerSnpnId(v PlmnIdNid) {
	o.ProducerSnpnId = &v
}

// GetProducerSnssaiList returns the ProducerSnssaiList field value if set, zero value otherwise.
func (o *AccessTokenClaims) GetProducerSnssaiList() []Snssai {
	if o == nil || IsNil(o.ProducerSnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.ProducerSnssaiList
}

// GetProducerSnssaiListOk returns a tuple with the ProducerSnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetProducerSnssaiListOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.ProducerSnssaiList) {
		return nil, false
	}
	return o.ProducerSnssaiList, true
}

// HasProducerSnssaiList returns a boolean if a field has been set.
func (o *AccessTokenClaims) HasProducerSnssaiList() bool {
	if o != nil && !IsNil(o.ProducerSnssaiList) {
		return true
	}

	return false
}

// SetProducerSnssaiList gets a reference to the given []Snssai and assigns it to the ProducerSnssaiList field.
func (o *AccessTokenClaims) SetProducerSnssaiList(v []Snssai) {
	o.ProducerSnssaiList = v
}

// GetProducerNsiList returns the ProducerNsiList field value if set, zero value otherwise.
func (o *AccessTokenClaims) GetProducerNsiList() []string {
	if o == nil || IsNil(o.ProducerNsiList) {
		var ret []string
		return ret
	}
	return o.ProducerNsiList
}

// GetProducerNsiListOk returns a tuple with the ProducerNsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetProducerNsiListOk() ([]string, bool) {
	if o == nil || IsNil(o.ProducerNsiList) {
		return nil, false
	}
	return o.ProducerNsiList, true
}

// HasProducerNsiList returns a boolean if a field has been set.
func (o *AccessTokenClaims) HasProducerNsiList() bool {
	if o != nil && !IsNil(o.ProducerNsiList) {
		return true
	}

	return false
}

// SetProducerNsiList gets a reference to the given []string and assigns it to the ProducerNsiList field.
func (o *AccessTokenClaims) SetProducerNsiList(v []string) {
	o.ProducerNsiList = v
}

// GetProducerNfSetId returns the ProducerNfSetId field value if set, zero value otherwise.
func (o *AccessTokenClaims) GetProducerNfSetId() string {
	if o == nil || IsNil(o.ProducerNfSetId) {
		var ret string
		return ret
	}
	return *o.ProducerNfSetId
}

// GetProducerNfSetIdOk returns a tuple with the ProducerNfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetProducerNfSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProducerNfSetId) {
		return nil, false
	}
	return o.ProducerNfSetId, true
}

// HasProducerNfSetId returns a boolean if a field has been set.
func (o *AccessTokenClaims) HasProducerNfSetId() bool {
	if o != nil && !IsNil(o.ProducerNfSetId) {
		return true
	}

	return false
}

// SetProducerNfSetId gets a reference to the given string and assigns it to the ProducerNfSetId field.
func (o *AccessTokenClaims) SetProducerNfSetId(v string) {
	o.ProducerNfSetId = &v
}

// GetProducerNfServiceSetId returns the ProducerNfServiceSetId field value if set, zero value otherwise.
func (o *AccessTokenClaims) GetProducerNfServiceSetId() string {
	if o == nil || IsNil(o.ProducerNfServiceSetId) {
		var ret string
		return ret
	}
	return *o.ProducerNfServiceSetId
}

// GetProducerNfServiceSetIdOk returns a tuple with the ProducerNfServiceSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetProducerNfServiceSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProducerNfServiceSetId) {
		return nil, false
	}
	return o.ProducerNfServiceSetId, true
}

// HasProducerNfServiceSetId returns a boolean if a field has been set.
func (o *AccessTokenClaims) HasProducerNfServiceSetId() bool {
	if o != nil && !IsNil(o.ProducerNfServiceSetId) {
		return true
	}

	return false
}

// SetProducerNfServiceSetId gets a reference to the given string and assigns it to the ProducerNfServiceSetId field.
func (o *AccessTokenClaims) SetProducerNfServiceSetId(v string) {
	o.ProducerNfServiceSetId = &v
}

// GetSourceNfInstanceId returns the SourceNfInstanceId field value if set, zero value otherwise.
func (o *AccessTokenClaims) GetSourceNfInstanceId() string {
	if o == nil || IsNil(o.SourceNfInstanceId) {
		var ret string
		return ret
	}
	return *o.SourceNfInstanceId
}

// GetSourceNfInstanceIdOk returns a tuple with the SourceNfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetSourceNfInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceNfInstanceId) {
		return nil, false
	}
	return o.SourceNfInstanceId, true
}

// HasSourceNfInstanceId returns a boolean if a field has been set.
func (o *AccessTokenClaims) HasSourceNfInstanceId() bool {
	if o != nil && !IsNil(o.SourceNfInstanceId) {
		return true
	}

	return false
}

// SetSourceNfInstanceId gets a reference to the given string and assigns it to the SourceNfInstanceId field.
func (o *AccessTokenClaims) SetSourceNfInstanceId(v string) {
	o.SourceNfInstanceId = &v
}

func (o AccessTokenClaims) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenClaims) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iss"] = o.Iss
	toSerialize["sub"] = o.Sub
	toSerialize["aud"] = o.Aud
	toSerialize["scope"] = o.Scope
	toSerialize["exp"] = o.Exp
	if !IsNil(o.ConsumerPlmnId) {
		toSerialize["consumerPlmnId"] = o.ConsumerPlmnId
	}
	if !IsNil(o.ConsumerSnpnId) {
		toSerialize["consumerSnpnId"] = o.ConsumerSnpnId
	}
	if !IsNil(o.ProducerPlmnId) {
		toSerialize["producerPlmnId"] = o.ProducerPlmnId
	}
	if !IsNil(o.ProducerSnpnId) {
		toSerialize["producerSnpnId"] = o.ProducerSnpnId
	}
	if !IsNil(o.ProducerSnssaiList) {
		toSerialize["producerSnssaiList"] = o.ProducerSnssaiList
	}
	if !IsNil(o.ProducerNsiList) {
		toSerialize["producerNsiList"] = o.ProducerNsiList
	}
	if !IsNil(o.ProducerNfSetId) {
		toSerialize["producerNfSetId"] = o.ProducerNfSetId
	}
	if !IsNil(o.ProducerNfServiceSetId) {
		toSerialize["producerNfServiceSetId"] = o.ProducerNfServiceSetId
	}
	if !IsNil(o.SourceNfInstanceId) {
		toSerialize["sourceNfInstanceId"] = o.SourceNfInstanceId
	}
	return toSerialize, nil
}

type NullableAccessTokenClaims struct {
	value *AccessTokenClaims
	isSet bool
}

func (v NullableAccessTokenClaims) Get() *AccessTokenClaims {
	return v.value
}

func (v *NullableAccessTokenClaims) Set(val *AccessTokenClaims) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenClaims) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenClaims) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenClaims(val *AccessTokenClaims) *NullableAccessTokenClaims {
	return &NullableAccessTokenClaims{value: val, isSet: true}
}

func (v NullableAccessTokenClaims) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenClaims) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



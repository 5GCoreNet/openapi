/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the UeContextCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeContextCreateData{}

// UeContextCreateData Data within a request to create an individual ueContext resource
type UeContextCreateData struct {
	UeContext          UeContext         `json:"ueContext"`
	TargetId           NgRanTargetId     `json:"targetId"`
	SourceToTargetData N2InfoContent     `json:"sourceToTargetData"`
	PduSessionList     []N2SmInformation `json:"pduSessionList"`
	// String providing an URI formatted according to RFC 3986.
	N2NotifyUri                *string        `json:"n2NotifyUri,omitempty"`
	UeRadioCapability          *N2InfoContent `json:"ueRadioCapability,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent `json:"ueRadioCapabilityForPaging,omitempty"`
	NgapCause                  *NgApCause     `json:"ngapCause,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string    `json:"supportedFeatures,omitempty"`
	ServingNetwork    *PlmnIdNid `json:"servingNetwork,omitempty"`
}

// NewUeContextCreateData instantiates a new UeContextCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeContextCreateData(ueContext UeContext, targetId NgRanTargetId, sourceToTargetData N2InfoContent, pduSessionList []N2SmInformation) *UeContextCreateData {
	this := UeContextCreateData{}
	this.UeContext = ueContext
	this.TargetId = targetId
	this.SourceToTargetData = sourceToTargetData
	this.PduSessionList = pduSessionList
	return &this
}

// NewUeContextCreateDataWithDefaults instantiates a new UeContextCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeContextCreateDataWithDefaults() *UeContextCreateData {
	this := UeContextCreateData{}
	return &this
}

// GetUeContext returns the UeContext field value
func (o *UeContextCreateData) GetUeContext() UeContext {
	if o == nil {
		var ret UeContext
		return ret
	}

	return o.UeContext
}

// GetUeContextOk returns a tuple with the UeContext field value
// and a boolean to check if the value has been set.
func (o *UeContextCreateData) GetUeContextOk() (*UeContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UeContext, true
}

// SetUeContext sets field value
func (o *UeContextCreateData) SetUeContext(v UeContext) {
	o.UeContext = v
}

// GetTargetId returns the TargetId field value
func (o *UeContextCreateData) GetTargetId() NgRanTargetId {
	if o == nil {
		var ret NgRanTargetId
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *UeContextCreateData) GetTargetIdOk() (*NgRanTargetId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *UeContextCreateData) SetTargetId(v NgRanTargetId) {
	o.TargetId = v
}

// GetSourceToTargetData returns the SourceToTargetData field value
func (o *UeContextCreateData) GetSourceToTargetData() N2InfoContent {
	if o == nil {
		var ret N2InfoContent
		return ret
	}

	return o.SourceToTargetData
}

// GetSourceToTargetDataOk returns a tuple with the SourceToTargetData field value
// and a boolean to check if the value has been set.
func (o *UeContextCreateData) GetSourceToTargetDataOk() (*N2InfoContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceToTargetData, true
}

// SetSourceToTargetData sets field value
func (o *UeContextCreateData) SetSourceToTargetData(v N2InfoContent) {
	o.SourceToTargetData = v
}

// GetPduSessionList returns the PduSessionList field value
func (o *UeContextCreateData) GetPduSessionList() []N2SmInformation {
	if o == nil {
		var ret []N2SmInformation
		return ret
	}

	return o.PduSessionList
}

// GetPduSessionListOk returns a tuple with the PduSessionList field value
// and a boolean to check if the value has been set.
func (o *UeContextCreateData) GetPduSessionListOk() ([]N2SmInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.PduSessionList, true
}

// SetPduSessionList sets field value
func (o *UeContextCreateData) SetPduSessionList(v []N2SmInformation) {
	o.PduSessionList = v
}

// GetN2NotifyUri returns the N2NotifyUri field value if set, zero value otherwise.
func (o *UeContextCreateData) GetN2NotifyUri() string {
	if o == nil || IsNil(o.N2NotifyUri) {
		var ret string
		return ret
	}
	return *o.N2NotifyUri
}

// GetN2NotifyUriOk returns a tuple with the N2NotifyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextCreateData) GetN2NotifyUriOk() (*string, bool) {
	if o == nil || IsNil(o.N2NotifyUri) {
		return nil, false
	}
	return o.N2NotifyUri, true
}

// HasN2NotifyUri returns a boolean if a field has been set.
func (o *UeContextCreateData) HasN2NotifyUri() bool {
	if o != nil && !IsNil(o.N2NotifyUri) {
		return true
	}

	return false
}

// SetN2NotifyUri gets a reference to the given string and assigns it to the N2NotifyUri field.
func (o *UeContextCreateData) SetN2NotifyUri(v string) {
	o.N2NotifyUri = &v
}

// GetUeRadioCapability returns the UeRadioCapability field value if set, zero value otherwise.
func (o *UeContextCreateData) GetUeRadioCapability() N2InfoContent {
	if o == nil || IsNil(o.UeRadioCapability) {
		var ret N2InfoContent
		return ret
	}
	return *o.UeRadioCapability
}

// GetUeRadioCapabilityOk returns a tuple with the UeRadioCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextCreateData) GetUeRadioCapabilityOk() (*N2InfoContent, bool) {
	if o == nil || IsNil(o.UeRadioCapability) {
		return nil, false
	}
	return o.UeRadioCapability, true
}

// HasUeRadioCapability returns a boolean if a field has been set.
func (o *UeContextCreateData) HasUeRadioCapability() bool {
	if o != nil && !IsNil(o.UeRadioCapability) {
		return true
	}

	return false
}

// SetUeRadioCapability gets a reference to the given N2InfoContent and assigns it to the UeRadioCapability field.
func (o *UeContextCreateData) SetUeRadioCapability(v N2InfoContent) {
	o.UeRadioCapability = &v
}

// GetUeRadioCapabilityForPaging returns the UeRadioCapabilityForPaging field value if set, zero value otherwise.
func (o *UeContextCreateData) GetUeRadioCapabilityForPaging() N2InfoContent {
	if o == nil || IsNil(o.UeRadioCapabilityForPaging) {
		var ret N2InfoContent
		return ret
	}
	return *o.UeRadioCapabilityForPaging
}

// GetUeRadioCapabilityForPagingOk returns a tuple with the UeRadioCapabilityForPaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextCreateData) GetUeRadioCapabilityForPagingOk() (*N2InfoContent, bool) {
	if o == nil || IsNil(o.UeRadioCapabilityForPaging) {
		return nil, false
	}
	return o.UeRadioCapabilityForPaging, true
}

// HasUeRadioCapabilityForPaging returns a boolean if a field has been set.
func (o *UeContextCreateData) HasUeRadioCapabilityForPaging() bool {
	if o != nil && !IsNil(o.UeRadioCapabilityForPaging) {
		return true
	}

	return false
}

// SetUeRadioCapabilityForPaging gets a reference to the given N2InfoContent and assigns it to the UeRadioCapabilityForPaging field.
func (o *UeContextCreateData) SetUeRadioCapabilityForPaging(v N2InfoContent) {
	o.UeRadioCapabilityForPaging = &v
}

// GetNgapCause returns the NgapCause field value if set, zero value otherwise.
func (o *UeContextCreateData) GetNgapCause() NgApCause {
	if o == nil || IsNil(o.NgapCause) {
		var ret NgApCause
		return ret
	}
	return *o.NgapCause
}

// GetNgapCauseOk returns a tuple with the NgapCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextCreateData) GetNgapCauseOk() (*NgApCause, bool) {
	if o == nil || IsNil(o.NgapCause) {
		return nil, false
	}
	return o.NgapCause, true
}

// HasNgapCause returns a boolean if a field has been set.
func (o *UeContextCreateData) HasNgapCause() bool {
	if o != nil && !IsNil(o.NgapCause) {
		return true
	}

	return false
}

// SetNgapCause gets a reference to the given NgApCause and assigns it to the NgapCause field.
func (o *UeContextCreateData) SetNgapCause(v NgApCause) {
	o.NgapCause = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *UeContextCreateData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextCreateData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *UeContextCreateData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *UeContextCreateData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetServingNetwork returns the ServingNetwork field value if set, zero value otherwise.
func (o *UeContextCreateData) GetServingNetwork() PlmnIdNid {
	if o == nil || IsNil(o.ServingNetwork) {
		var ret PlmnIdNid
		return ret
	}
	return *o.ServingNetwork
}

// GetServingNetworkOk returns a tuple with the ServingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextCreateData) GetServingNetworkOk() (*PlmnIdNid, bool) {
	if o == nil || IsNil(o.ServingNetwork) {
		return nil, false
	}
	return o.ServingNetwork, true
}

// HasServingNetwork returns a boolean if a field has been set.
func (o *UeContextCreateData) HasServingNetwork() bool {
	if o != nil && !IsNil(o.ServingNetwork) {
		return true
	}

	return false
}

// SetServingNetwork gets a reference to the given PlmnIdNid and assigns it to the ServingNetwork field.
func (o *UeContextCreateData) SetServingNetwork(v PlmnIdNid) {
	o.ServingNetwork = &v
}

func (o UeContextCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeContextCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ueContext"] = o.UeContext
	toSerialize["targetId"] = o.TargetId
	toSerialize["sourceToTargetData"] = o.SourceToTargetData
	toSerialize["pduSessionList"] = o.PduSessionList
	if !IsNil(o.N2NotifyUri) {
		toSerialize["n2NotifyUri"] = o.N2NotifyUri
	}
	if !IsNil(o.UeRadioCapability) {
		toSerialize["ueRadioCapability"] = o.UeRadioCapability
	}
	if !IsNil(o.UeRadioCapabilityForPaging) {
		toSerialize["ueRadioCapabilityForPaging"] = o.UeRadioCapabilityForPaging
	}
	if !IsNil(o.NgapCause) {
		toSerialize["ngapCause"] = o.NgapCause
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.ServingNetwork) {
		toSerialize["servingNetwork"] = o.ServingNetwork
	}
	return toSerialize, nil
}

type NullableUeContextCreateData struct {
	value *UeContextCreateData
	isSet bool
}

func (v NullableUeContextCreateData) Get() *UeContextCreateData {
	return v.value
}

func (v *NullableUeContextCreateData) Set(val *UeContextCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableUeContextCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableUeContextCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeContextCreateData(val *UeContextCreateData) *NullableUeContextCreateData {
	return &NullableUeContextCreateData{value: val, isSet: true}
}

func (v NullableUeContextCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeContextCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

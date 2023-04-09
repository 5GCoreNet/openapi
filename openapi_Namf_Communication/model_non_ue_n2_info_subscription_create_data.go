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

// checks if the NonUeN2InfoSubscriptionCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonUeN2InfoSubscriptionCreateData{}

// NonUeN2InfoSubscriptionCreateData Data within a create subscription request for non-UE specific N2 information notification
type NonUeN2InfoSubscriptionCreateData struct {
	GlobalRanNodeList  []GlobalRanNodeId  `json:"globalRanNodeList,omitempty"`
	AnTypeList         []AccessType       `json:"anTypeList,omitempty"`
	N2InformationClass N2InformationClass `json:"n2InformationClass"`
	// String providing an URI formatted according to RFC 3986.
	N2NotifyCallbackUri string `json:"n2NotifyCallbackUri"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfId *string `json:"nfId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewNonUeN2InfoSubscriptionCreateData instantiates a new NonUeN2InfoSubscriptionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonUeN2InfoSubscriptionCreateData(n2InformationClass N2InformationClass, n2NotifyCallbackUri string) *NonUeN2InfoSubscriptionCreateData {
	this := NonUeN2InfoSubscriptionCreateData{}
	this.N2InformationClass = n2InformationClass
	this.N2NotifyCallbackUri = n2NotifyCallbackUri
	return &this
}

// NewNonUeN2InfoSubscriptionCreateDataWithDefaults instantiates a new NonUeN2InfoSubscriptionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonUeN2InfoSubscriptionCreateDataWithDefaults() *NonUeN2InfoSubscriptionCreateData {
	this := NonUeN2InfoSubscriptionCreateData{}
	return &this
}

// GetGlobalRanNodeList returns the GlobalRanNodeList field value if set, zero value otherwise.
func (o *NonUeN2InfoSubscriptionCreateData) GetGlobalRanNodeList() []GlobalRanNodeId {
	if o == nil || IsNil(o.GlobalRanNodeList) {
		var ret []GlobalRanNodeId
		return ret
	}
	return o.GlobalRanNodeList
}

// GetGlobalRanNodeListOk returns a tuple with the GlobalRanNodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonUeN2InfoSubscriptionCreateData) GetGlobalRanNodeListOk() ([]GlobalRanNodeId, bool) {
	if o == nil || IsNil(o.GlobalRanNodeList) {
		return nil, false
	}
	return o.GlobalRanNodeList, true
}

// HasGlobalRanNodeList returns a boolean if a field has been set.
func (o *NonUeN2InfoSubscriptionCreateData) HasGlobalRanNodeList() bool {
	if o != nil && !IsNil(o.GlobalRanNodeList) {
		return true
	}

	return false
}

// SetGlobalRanNodeList gets a reference to the given []GlobalRanNodeId and assigns it to the GlobalRanNodeList field.
func (o *NonUeN2InfoSubscriptionCreateData) SetGlobalRanNodeList(v []GlobalRanNodeId) {
	o.GlobalRanNodeList = v
}

// GetAnTypeList returns the AnTypeList field value if set, zero value otherwise.
func (o *NonUeN2InfoSubscriptionCreateData) GetAnTypeList() []AccessType {
	if o == nil || IsNil(o.AnTypeList) {
		var ret []AccessType
		return ret
	}
	return o.AnTypeList
}

// GetAnTypeListOk returns a tuple with the AnTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonUeN2InfoSubscriptionCreateData) GetAnTypeListOk() ([]AccessType, bool) {
	if o == nil || IsNil(o.AnTypeList) {
		return nil, false
	}
	return o.AnTypeList, true
}

// HasAnTypeList returns a boolean if a field has been set.
func (o *NonUeN2InfoSubscriptionCreateData) HasAnTypeList() bool {
	if o != nil && !IsNil(o.AnTypeList) {
		return true
	}

	return false
}

// SetAnTypeList gets a reference to the given []AccessType and assigns it to the AnTypeList field.
func (o *NonUeN2InfoSubscriptionCreateData) SetAnTypeList(v []AccessType) {
	o.AnTypeList = v
}

// GetN2InformationClass returns the N2InformationClass field value
func (o *NonUeN2InfoSubscriptionCreateData) GetN2InformationClass() N2InformationClass {
	if o == nil {
		var ret N2InformationClass
		return ret
	}

	return o.N2InformationClass
}

// GetN2InformationClassOk returns a tuple with the N2InformationClass field value
// and a boolean to check if the value has been set.
func (o *NonUeN2InfoSubscriptionCreateData) GetN2InformationClassOk() (*N2InformationClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N2InformationClass, true
}

// SetN2InformationClass sets field value
func (o *NonUeN2InfoSubscriptionCreateData) SetN2InformationClass(v N2InformationClass) {
	o.N2InformationClass = v
}

// GetN2NotifyCallbackUri returns the N2NotifyCallbackUri field value
func (o *NonUeN2InfoSubscriptionCreateData) GetN2NotifyCallbackUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.N2NotifyCallbackUri
}

// GetN2NotifyCallbackUriOk returns a tuple with the N2NotifyCallbackUri field value
// and a boolean to check if the value has been set.
func (o *NonUeN2InfoSubscriptionCreateData) GetN2NotifyCallbackUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N2NotifyCallbackUri, true
}

// SetN2NotifyCallbackUri sets field value
func (o *NonUeN2InfoSubscriptionCreateData) SetN2NotifyCallbackUri(v string) {
	o.N2NotifyCallbackUri = v
}

// GetNfId returns the NfId field value if set, zero value otherwise.
func (o *NonUeN2InfoSubscriptionCreateData) GetNfId() string {
	if o == nil || IsNil(o.NfId) {
		var ret string
		return ret
	}
	return *o.NfId
}

// GetNfIdOk returns a tuple with the NfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonUeN2InfoSubscriptionCreateData) GetNfIdOk() (*string, bool) {
	if o == nil || IsNil(o.NfId) {
		return nil, false
	}
	return o.NfId, true
}

// HasNfId returns a boolean if a field has been set.
func (o *NonUeN2InfoSubscriptionCreateData) HasNfId() bool {
	if o != nil && !IsNil(o.NfId) {
		return true
	}

	return false
}

// SetNfId gets a reference to the given string and assigns it to the NfId field.
func (o *NonUeN2InfoSubscriptionCreateData) SetNfId(v string) {
	o.NfId = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *NonUeN2InfoSubscriptionCreateData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonUeN2InfoSubscriptionCreateData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *NonUeN2InfoSubscriptionCreateData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *NonUeN2InfoSubscriptionCreateData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o NonUeN2InfoSubscriptionCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonUeN2InfoSubscriptionCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GlobalRanNodeList) {
		toSerialize["globalRanNodeList"] = o.GlobalRanNodeList
	}
	if !IsNil(o.AnTypeList) {
		toSerialize["anTypeList"] = o.AnTypeList
	}
	toSerialize["n2InformationClass"] = o.N2InformationClass
	toSerialize["n2NotifyCallbackUri"] = o.N2NotifyCallbackUri
	if !IsNil(o.NfId) {
		toSerialize["nfId"] = o.NfId
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableNonUeN2InfoSubscriptionCreateData struct {
	value *NonUeN2InfoSubscriptionCreateData
	isSet bool
}

func (v NullableNonUeN2InfoSubscriptionCreateData) Get() *NonUeN2InfoSubscriptionCreateData {
	return v.value
}

func (v *NullableNonUeN2InfoSubscriptionCreateData) Set(val *NonUeN2InfoSubscriptionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableNonUeN2InfoSubscriptionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableNonUeN2InfoSubscriptionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonUeN2InfoSubscriptionCreateData(val *NonUeN2InfoSubscriptionCreateData) *NullableNonUeN2InfoSubscriptionCreateData {
	return &NullableNonUeN2InfoSubscriptionCreateData{value: val, isSet: true}
}

func (v NullableNonUeN2InfoSubscriptionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonUeN2InfoSubscriptionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

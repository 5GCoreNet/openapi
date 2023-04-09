/*
Naf_ProSe API

Naf_ProSe Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_ProSe

import (
	"encoding/json"
)

// checks if the AuthDisResData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthDisResData{}

// AuthDisResData Represents the obtained authorization Data for a UE of a 5G ProSe Direct Discovery  request.
type AuthDisResData struct {
	AuthResponseType         AuthResponseType                `json:"authResponseType"`
	ProseAppCodeSuffixPool   *ProseApplicationCodeSuffixPool `json:"proseAppCodeSuffixPool,omitempty"`
	Pduids                   []string                        `json:"pduids,omitempty"`
	RestrictedCodeSuffixPool []RestrictedCodeSuffixPool      `json:"restrictedCodeSuffixPool,omitempty"`
	ProseAppMasks            []string                        `json:"proseAppMasks,omitempty"`
	ProSeRestrictedMasks     []string                        `json:"proSeRestrictedMasks,omitempty"`
	// Contains the Application Level Container.
	ResAppLevelContainer *string      `json:"resAppLevelContainer,omitempty"`
	TargetDataSet        []TargetData `json:"targetDataSet,omitempty"`
	// Contains the PDUID.
	TargetPduid *string `json:"targetPduid,omitempty"`
	// Contains the metadata.
	MetaData *string `json:"metaData,omitempty"`
}

// NewAuthDisResData instantiates a new AuthDisResData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthDisResData(authResponseType AuthResponseType) *AuthDisResData {
	this := AuthDisResData{}
	this.AuthResponseType = authResponseType
	return &this
}

// NewAuthDisResDataWithDefaults instantiates a new AuthDisResData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthDisResDataWithDefaults() *AuthDisResData {
	this := AuthDisResData{}
	return &this
}

// GetAuthResponseType returns the AuthResponseType field value
func (o *AuthDisResData) GetAuthResponseType() AuthResponseType {
	if o == nil {
		var ret AuthResponseType
		return ret
	}

	return o.AuthResponseType
}

// GetAuthResponseTypeOk returns a tuple with the AuthResponseType field value
// and a boolean to check if the value has been set.
func (o *AuthDisResData) GetAuthResponseTypeOk() (*AuthResponseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthResponseType, true
}

// SetAuthResponseType sets field value
func (o *AuthDisResData) SetAuthResponseType(v AuthResponseType) {
	o.AuthResponseType = v
}

// GetProseAppCodeSuffixPool returns the ProseAppCodeSuffixPool field value if set, zero value otherwise.
func (o *AuthDisResData) GetProseAppCodeSuffixPool() ProseApplicationCodeSuffixPool {
	if o == nil || IsNil(o.ProseAppCodeSuffixPool) {
		var ret ProseApplicationCodeSuffixPool
		return ret
	}
	return *o.ProseAppCodeSuffixPool
}

// GetProseAppCodeSuffixPoolOk returns a tuple with the ProseAppCodeSuffixPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisResData) GetProseAppCodeSuffixPoolOk() (*ProseApplicationCodeSuffixPool, bool) {
	if o == nil || IsNil(o.ProseAppCodeSuffixPool) {
		return nil, false
	}
	return o.ProseAppCodeSuffixPool, true
}

// HasProseAppCodeSuffixPool returns a boolean if a field has been set.
func (o *AuthDisResData) HasProseAppCodeSuffixPool() bool {
	if o != nil && !IsNil(o.ProseAppCodeSuffixPool) {
		return true
	}

	return false
}

// SetProseAppCodeSuffixPool gets a reference to the given ProseApplicationCodeSuffixPool and assigns it to the ProseAppCodeSuffixPool field.
func (o *AuthDisResData) SetProseAppCodeSuffixPool(v ProseApplicationCodeSuffixPool) {
	o.ProseAppCodeSuffixPool = &v
}

// GetPduids returns the Pduids field value if set, zero value otherwise.
func (o *AuthDisResData) GetPduids() []string {
	if o == nil || IsNil(o.Pduids) {
		var ret []string
		return ret
	}
	return o.Pduids
}

// GetPduidsOk returns a tuple with the Pduids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisResData) GetPduidsOk() ([]string, bool) {
	if o == nil || IsNil(o.Pduids) {
		return nil, false
	}
	return o.Pduids, true
}

// HasPduids returns a boolean if a field has been set.
func (o *AuthDisResData) HasPduids() bool {
	if o != nil && !IsNil(o.Pduids) {
		return true
	}

	return false
}

// SetPduids gets a reference to the given []string and assigns it to the Pduids field.
func (o *AuthDisResData) SetPduids(v []string) {
	o.Pduids = v
}

// GetRestrictedCodeSuffixPool returns the RestrictedCodeSuffixPool field value if set, zero value otherwise.
func (o *AuthDisResData) GetRestrictedCodeSuffixPool() []RestrictedCodeSuffixPool {
	if o == nil || IsNil(o.RestrictedCodeSuffixPool) {
		var ret []RestrictedCodeSuffixPool
		return ret
	}
	return o.RestrictedCodeSuffixPool
}

// GetRestrictedCodeSuffixPoolOk returns a tuple with the RestrictedCodeSuffixPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisResData) GetRestrictedCodeSuffixPoolOk() ([]RestrictedCodeSuffixPool, bool) {
	if o == nil || IsNil(o.RestrictedCodeSuffixPool) {
		return nil, false
	}
	return o.RestrictedCodeSuffixPool, true
}

// HasRestrictedCodeSuffixPool returns a boolean if a field has been set.
func (o *AuthDisResData) HasRestrictedCodeSuffixPool() bool {
	if o != nil && !IsNil(o.RestrictedCodeSuffixPool) {
		return true
	}

	return false
}

// SetRestrictedCodeSuffixPool gets a reference to the given []RestrictedCodeSuffixPool and assigns it to the RestrictedCodeSuffixPool field.
func (o *AuthDisResData) SetRestrictedCodeSuffixPool(v []RestrictedCodeSuffixPool) {
	o.RestrictedCodeSuffixPool = v
}

// GetProseAppMasks returns the ProseAppMasks field value if set, zero value otherwise.
func (o *AuthDisResData) GetProseAppMasks() []string {
	if o == nil || IsNil(o.ProseAppMasks) {
		var ret []string
		return ret
	}
	return o.ProseAppMasks
}

// GetProseAppMasksOk returns a tuple with the ProseAppMasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisResData) GetProseAppMasksOk() ([]string, bool) {
	if o == nil || IsNil(o.ProseAppMasks) {
		return nil, false
	}
	return o.ProseAppMasks, true
}

// HasProseAppMasks returns a boolean if a field has been set.
func (o *AuthDisResData) HasProseAppMasks() bool {
	if o != nil && !IsNil(o.ProseAppMasks) {
		return true
	}

	return false
}

// SetProseAppMasks gets a reference to the given []string and assigns it to the ProseAppMasks field.
func (o *AuthDisResData) SetProseAppMasks(v []string) {
	o.ProseAppMasks = v
}

// GetProSeRestrictedMasks returns the ProSeRestrictedMasks field value if set, zero value otherwise.
func (o *AuthDisResData) GetProSeRestrictedMasks() []string {
	if o == nil || IsNil(o.ProSeRestrictedMasks) {
		var ret []string
		return ret
	}
	return o.ProSeRestrictedMasks
}

// GetProSeRestrictedMasksOk returns a tuple with the ProSeRestrictedMasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisResData) GetProSeRestrictedMasksOk() ([]string, bool) {
	if o == nil || IsNil(o.ProSeRestrictedMasks) {
		return nil, false
	}
	return o.ProSeRestrictedMasks, true
}

// HasProSeRestrictedMasks returns a boolean if a field has been set.
func (o *AuthDisResData) HasProSeRestrictedMasks() bool {
	if o != nil && !IsNil(o.ProSeRestrictedMasks) {
		return true
	}

	return false
}

// SetProSeRestrictedMasks gets a reference to the given []string and assigns it to the ProSeRestrictedMasks field.
func (o *AuthDisResData) SetProSeRestrictedMasks(v []string) {
	o.ProSeRestrictedMasks = v
}

// GetResAppLevelContainer returns the ResAppLevelContainer field value if set, zero value otherwise.
func (o *AuthDisResData) GetResAppLevelContainer() string {
	if o == nil || IsNil(o.ResAppLevelContainer) {
		var ret string
		return ret
	}
	return *o.ResAppLevelContainer
}

// GetResAppLevelContainerOk returns a tuple with the ResAppLevelContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisResData) GetResAppLevelContainerOk() (*string, bool) {
	if o == nil || IsNil(o.ResAppLevelContainer) {
		return nil, false
	}
	return o.ResAppLevelContainer, true
}

// HasResAppLevelContainer returns a boolean if a field has been set.
func (o *AuthDisResData) HasResAppLevelContainer() bool {
	if o != nil && !IsNil(o.ResAppLevelContainer) {
		return true
	}

	return false
}

// SetResAppLevelContainer gets a reference to the given string and assigns it to the ResAppLevelContainer field.
func (o *AuthDisResData) SetResAppLevelContainer(v string) {
	o.ResAppLevelContainer = &v
}

// GetTargetDataSet returns the TargetDataSet field value if set, zero value otherwise.
func (o *AuthDisResData) GetTargetDataSet() []TargetData {
	if o == nil || IsNil(o.TargetDataSet) {
		var ret []TargetData
		return ret
	}
	return o.TargetDataSet
}

// GetTargetDataSetOk returns a tuple with the TargetDataSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisResData) GetTargetDataSetOk() ([]TargetData, bool) {
	if o == nil || IsNil(o.TargetDataSet) {
		return nil, false
	}
	return o.TargetDataSet, true
}

// HasTargetDataSet returns a boolean if a field has been set.
func (o *AuthDisResData) HasTargetDataSet() bool {
	if o != nil && !IsNil(o.TargetDataSet) {
		return true
	}

	return false
}

// SetTargetDataSet gets a reference to the given []TargetData and assigns it to the TargetDataSet field.
func (o *AuthDisResData) SetTargetDataSet(v []TargetData) {
	o.TargetDataSet = v
}

// GetTargetPduid returns the TargetPduid field value if set, zero value otherwise.
func (o *AuthDisResData) GetTargetPduid() string {
	if o == nil || IsNil(o.TargetPduid) {
		var ret string
		return ret
	}
	return *o.TargetPduid
}

// GetTargetPduidOk returns a tuple with the TargetPduid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisResData) GetTargetPduidOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPduid) {
		return nil, false
	}
	return o.TargetPduid, true
}

// HasTargetPduid returns a boolean if a field has been set.
func (o *AuthDisResData) HasTargetPduid() bool {
	if o != nil && !IsNil(o.TargetPduid) {
		return true
	}

	return false
}

// SetTargetPduid gets a reference to the given string and assigns it to the TargetPduid field.
func (o *AuthDisResData) SetTargetPduid(v string) {
	o.TargetPduid = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *AuthDisResData) GetMetaData() string {
	if o == nil || IsNil(o.MetaData) {
		var ret string
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisResData) GetMetaDataOk() (*string, bool) {
	if o == nil || IsNil(o.MetaData) {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *AuthDisResData) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given string and assigns it to the MetaData field.
func (o *AuthDisResData) SetMetaData(v string) {
	o.MetaData = &v
}

func (o AuthDisResData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthDisResData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authResponseType"] = o.AuthResponseType
	if !IsNil(o.ProseAppCodeSuffixPool) {
		toSerialize["proseAppCodeSuffixPool"] = o.ProseAppCodeSuffixPool
	}
	if !IsNil(o.Pduids) {
		toSerialize["pduids"] = o.Pduids
	}
	if !IsNil(o.RestrictedCodeSuffixPool) {
		toSerialize["restrictedCodeSuffixPool"] = o.RestrictedCodeSuffixPool
	}
	if !IsNil(o.ProseAppMasks) {
		toSerialize["proseAppMasks"] = o.ProseAppMasks
	}
	if !IsNil(o.ProSeRestrictedMasks) {
		toSerialize["proSeRestrictedMasks"] = o.ProSeRestrictedMasks
	}
	if !IsNil(o.ResAppLevelContainer) {
		toSerialize["resAppLevelContainer"] = o.ResAppLevelContainer
	}
	if !IsNil(o.TargetDataSet) {
		toSerialize["targetDataSet"] = o.TargetDataSet
	}
	if !IsNil(o.TargetPduid) {
		toSerialize["targetPduid"] = o.TargetPduid
	}
	if !IsNil(o.MetaData) {
		toSerialize["metaData"] = o.MetaData
	}
	return toSerialize, nil
}

type NullableAuthDisResData struct {
	value *AuthDisResData
	isSet bool
}

func (v NullableAuthDisResData) Get() *AuthDisResData {
	return v.value
}

func (v *NullableAuthDisResData) Set(val *AuthDisResData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthDisResData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthDisResData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthDisResData(val *AuthDisResData) *NullableAuthDisResData {
	return &NullableAuthDisResData{value: val, isSet: true}
}

func (v NullableAuthDisResData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthDisResData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

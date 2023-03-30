/*
CAPIF_API_Invoker_Management_API

API for API invoker management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_API_Invoker_Management_API

import (
	"encoding/json"
)

// checks if the ServiceAPIDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAPIDescription{}

// ServiceAPIDescription Represents the description of a service API as published by the APF.
type ServiceAPIDescription struct {
	// API name, it is set as {apiName} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122.
	ApiName string `json:"apiName"`
	// API identifier assigned by the CAPIF core function to the published service API. Shall not be present in the HTTP POST request from the API publishing function to the CAPIF core function. Shall be present in the HTTP POST response from the CAPIF core function to the API publishing function and in the HTTP GET response from the CAPIF core function to the API invoker (discovery API). 
	ApiId *string `json:"apiId,omitempty"`
	// AEF profile information, which includes the exposed API details (e.g. protocol). 
	AefProfiles []AefProfile `json:"aefProfiles,omitempty"`
	// Text description of the API
	Description *string `json:"description,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	ShareableInfo *ShareableInformation `json:"shareableInfo,omitempty"`
	ServiceAPICategory *string `json:"serviceAPICategory,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	ApiSuppFeats *string `json:"apiSuppFeats,omitempty"`
	PubApiPath *PublishedApiPath `json:"pubApiPath,omitempty"`
	// CAPIF core function identifier.
	CcfId *string `json:"ccfId,omitempty"`
}

// NewServiceAPIDescription instantiates a new ServiceAPIDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAPIDescription(apiName string) *ServiceAPIDescription {
	this := ServiceAPIDescription{}
	this.ApiName = apiName
	return &this
}

// NewServiceAPIDescriptionWithDefaults instantiates a new ServiceAPIDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAPIDescriptionWithDefaults() *ServiceAPIDescription {
	this := ServiceAPIDescription{}
	return &this
}

// GetApiName returns the ApiName field value
func (o *ServiceAPIDescription) GetApiName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value
// and a boolean to check if the value has been set.
func (o *ServiceAPIDescription) GetApiNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiName, true
}

// SetApiName sets field value
func (o *ServiceAPIDescription) SetApiName(v string) {
	o.ApiName = v
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *ServiceAPIDescription) GetApiId() string {
	if o == nil || IsNil(o.ApiId) {
		var ret string
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAPIDescription) GetApiIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiId) {
		return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *ServiceAPIDescription) HasApiId() bool {
	if o != nil && !IsNil(o.ApiId) {
		return true
	}

	return false
}

// SetApiId gets a reference to the given string and assigns it to the ApiId field.
func (o *ServiceAPIDescription) SetApiId(v string) {
	o.ApiId = &v
}

// GetAefProfiles returns the AefProfiles field value if set, zero value otherwise.
func (o *ServiceAPIDescription) GetAefProfiles() []AefProfile {
	if o == nil || IsNil(o.AefProfiles) {
		var ret []AefProfile
		return ret
	}
	return o.AefProfiles
}

// GetAefProfilesOk returns a tuple with the AefProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAPIDescription) GetAefProfilesOk() ([]AefProfile, bool) {
	if o == nil || IsNil(o.AefProfiles) {
		return nil, false
	}
	return o.AefProfiles, true
}

// HasAefProfiles returns a boolean if a field has been set.
func (o *ServiceAPIDescription) HasAefProfiles() bool {
	if o != nil && !IsNil(o.AefProfiles) {
		return true
	}

	return false
}

// SetAefProfiles gets a reference to the given []AefProfile and assigns it to the AefProfiles field.
func (o *ServiceAPIDescription) SetAefProfiles(v []AefProfile) {
	o.AefProfiles = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAPIDescription) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAPIDescription) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAPIDescription) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAPIDescription) SetDescription(v string) {
	o.Description = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ServiceAPIDescription) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAPIDescription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ServiceAPIDescription) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ServiceAPIDescription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetShareableInfo returns the ShareableInfo field value if set, zero value otherwise.
func (o *ServiceAPIDescription) GetShareableInfo() ShareableInformation {
	if o == nil || IsNil(o.ShareableInfo) {
		var ret ShareableInformation
		return ret
	}
	return *o.ShareableInfo
}

// GetShareableInfoOk returns a tuple with the ShareableInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAPIDescription) GetShareableInfoOk() (*ShareableInformation, bool) {
	if o == nil || IsNil(o.ShareableInfo) {
		return nil, false
	}
	return o.ShareableInfo, true
}

// HasShareableInfo returns a boolean if a field has been set.
func (o *ServiceAPIDescription) HasShareableInfo() bool {
	if o != nil && !IsNil(o.ShareableInfo) {
		return true
	}

	return false
}

// SetShareableInfo gets a reference to the given ShareableInformation and assigns it to the ShareableInfo field.
func (o *ServiceAPIDescription) SetShareableInfo(v ShareableInformation) {
	o.ShareableInfo = &v
}

// GetServiceAPICategory returns the ServiceAPICategory field value if set, zero value otherwise.
func (o *ServiceAPIDescription) GetServiceAPICategory() string {
	if o == nil || IsNil(o.ServiceAPICategory) {
		var ret string
		return ret
	}
	return *o.ServiceAPICategory
}

// GetServiceAPICategoryOk returns a tuple with the ServiceAPICategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAPIDescription) GetServiceAPICategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAPICategory) {
		return nil, false
	}
	return o.ServiceAPICategory, true
}

// HasServiceAPICategory returns a boolean if a field has been set.
func (o *ServiceAPIDescription) HasServiceAPICategory() bool {
	if o != nil && !IsNil(o.ServiceAPICategory) {
		return true
	}

	return false
}

// SetServiceAPICategory gets a reference to the given string and assigns it to the ServiceAPICategory field.
func (o *ServiceAPIDescription) SetServiceAPICategory(v string) {
	o.ServiceAPICategory = &v
}

// GetApiSuppFeats returns the ApiSuppFeats field value if set, zero value otherwise.
func (o *ServiceAPIDescription) GetApiSuppFeats() string {
	if o == nil || IsNil(o.ApiSuppFeats) {
		var ret string
		return ret
	}
	return *o.ApiSuppFeats
}

// GetApiSuppFeatsOk returns a tuple with the ApiSuppFeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAPIDescription) GetApiSuppFeatsOk() (*string, bool) {
	if o == nil || IsNil(o.ApiSuppFeats) {
		return nil, false
	}
	return o.ApiSuppFeats, true
}

// HasApiSuppFeats returns a boolean if a field has been set.
func (o *ServiceAPIDescription) HasApiSuppFeats() bool {
	if o != nil && !IsNil(o.ApiSuppFeats) {
		return true
	}

	return false
}

// SetApiSuppFeats gets a reference to the given string and assigns it to the ApiSuppFeats field.
func (o *ServiceAPIDescription) SetApiSuppFeats(v string) {
	o.ApiSuppFeats = &v
}

// GetPubApiPath returns the PubApiPath field value if set, zero value otherwise.
func (o *ServiceAPIDescription) GetPubApiPath() PublishedApiPath {
	if o == nil || IsNil(o.PubApiPath) {
		var ret PublishedApiPath
		return ret
	}
	return *o.PubApiPath
}

// GetPubApiPathOk returns a tuple with the PubApiPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAPIDescription) GetPubApiPathOk() (*PublishedApiPath, bool) {
	if o == nil || IsNil(o.PubApiPath) {
		return nil, false
	}
	return o.PubApiPath, true
}

// HasPubApiPath returns a boolean if a field has been set.
func (o *ServiceAPIDescription) HasPubApiPath() bool {
	if o != nil && !IsNil(o.PubApiPath) {
		return true
	}

	return false
}

// SetPubApiPath gets a reference to the given PublishedApiPath and assigns it to the PubApiPath field.
func (o *ServiceAPIDescription) SetPubApiPath(v PublishedApiPath) {
	o.PubApiPath = &v
}

// GetCcfId returns the CcfId field value if set, zero value otherwise.
func (o *ServiceAPIDescription) GetCcfId() string {
	if o == nil || IsNil(o.CcfId) {
		var ret string
		return ret
	}
	return *o.CcfId
}

// GetCcfIdOk returns a tuple with the CcfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAPIDescription) GetCcfIdOk() (*string, bool) {
	if o == nil || IsNil(o.CcfId) {
		return nil, false
	}
	return o.CcfId, true
}

// HasCcfId returns a boolean if a field has been set.
func (o *ServiceAPIDescription) HasCcfId() bool {
	if o != nil && !IsNil(o.CcfId) {
		return true
	}

	return false
}

// SetCcfId gets a reference to the given string and assigns it to the CcfId field.
func (o *ServiceAPIDescription) SetCcfId(v string) {
	o.CcfId = &v
}

func (o ServiceAPIDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAPIDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiName"] = o.ApiName
	if !IsNil(o.ApiId) {
		toSerialize["apiId"] = o.ApiId
	}
	if !IsNil(o.AefProfiles) {
		toSerialize["aefProfiles"] = o.AefProfiles
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.ShareableInfo) {
		toSerialize["shareableInfo"] = o.ShareableInfo
	}
	if !IsNil(o.ServiceAPICategory) {
		toSerialize["serviceAPICategory"] = o.ServiceAPICategory
	}
	if !IsNil(o.ApiSuppFeats) {
		toSerialize["apiSuppFeats"] = o.ApiSuppFeats
	}
	if !IsNil(o.PubApiPath) {
		toSerialize["pubApiPath"] = o.PubApiPath
	}
	if !IsNil(o.CcfId) {
		toSerialize["ccfId"] = o.CcfId
	}
	return toSerialize, nil
}

type NullableServiceAPIDescription struct {
	value *ServiceAPIDescription
	isSet bool
}

func (v NullableServiceAPIDescription) Get() *ServiceAPIDescription {
	return v.value
}

func (v *NullableServiceAPIDescription) Set(val *ServiceAPIDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAPIDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAPIDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAPIDescription(val *ServiceAPIDescription) *NullableServiceAPIDescription {
	return &NullableServiceAPIDescription{value: val, isSet: true}
}

func (v NullableServiceAPIDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAPIDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
CAPIF_API_Provider_Management_API

API for API provider domain functions management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_API_Provider_Management_API

import (
	"encoding/json"
)

// checks if the APIProviderEnrolmentDetailsPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIProviderEnrolmentDetailsPatch{}

// APIProviderEnrolmentDetailsPatch Represents a list of modifications for the API provider domain's enrolment details.
type APIProviderEnrolmentDetailsPatch struct {
	// A list of individual API provider domain functions details. When included by the API management function in the HTTP request message, it lists the API provider domain functions that the API management function intends to register/update in registration or update registration procedure.
	ApiProvFuncs []APIProviderFunctionDetails `json:"apiProvFuncs,omitempty"`
	// Generic information related to the API provider domain such as details of the API provider applications.
	ApiProvDomInfo *string `json:"apiProvDomInfo,omitempty"`
}

// NewAPIProviderEnrolmentDetailsPatch instantiates a new APIProviderEnrolmentDetailsPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIProviderEnrolmentDetailsPatch() *APIProviderEnrolmentDetailsPatch {
	this := APIProviderEnrolmentDetailsPatch{}
	return &this
}

// NewAPIProviderEnrolmentDetailsPatchWithDefaults instantiates a new APIProviderEnrolmentDetailsPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIProviderEnrolmentDetailsPatchWithDefaults() *APIProviderEnrolmentDetailsPatch {
	this := APIProviderEnrolmentDetailsPatch{}
	return &this
}

// GetApiProvFuncs returns the ApiProvFuncs field value if set, zero value otherwise.
func (o *APIProviderEnrolmentDetailsPatch) GetApiProvFuncs() []APIProviderFunctionDetails {
	if o == nil || IsNil(o.ApiProvFuncs) {
		var ret []APIProviderFunctionDetails
		return ret
	}
	return o.ApiProvFuncs
}

// GetApiProvFuncsOk returns a tuple with the ApiProvFuncs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIProviderEnrolmentDetailsPatch) GetApiProvFuncsOk() ([]APIProviderFunctionDetails, bool) {
	if o == nil || IsNil(o.ApiProvFuncs) {
		return nil, false
	}
	return o.ApiProvFuncs, true
}

// HasApiProvFuncs returns a boolean if a field has been set.
func (o *APIProviderEnrolmentDetailsPatch) HasApiProvFuncs() bool {
	if o != nil && !IsNil(o.ApiProvFuncs) {
		return true
	}

	return false
}

// SetApiProvFuncs gets a reference to the given []APIProviderFunctionDetails and assigns it to the ApiProvFuncs field.
func (o *APIProviderEnrolmentDetailsPatch) SetApiProvFuncs(v []APIProviderFunctionDetails) {
	o.ApiProvFuncs = v
}

// GetApiProvDomInfo returns the ApiProvDomInfo field value if set, zero value otherwise.
func (o *APIProviderEnrolmentDetailsPatch) GetApiProvDomInfo() string {
	if o == nil || IsNil(o.ApiProvDomInfo) {
		var ret string
		return ret
	}
	return *o.ApiProvDomInfo
}

// GetApiProvDomInfoOk returns a tuple with the ApiProvDomInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIProviderEnrolmentDetailsPatch) GetApiProvDomInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ApiProvDomInfo) {
		return nil, false
	}
	return o.ApiProvDomInfo, true
}

// HasApiProvDomInfo returns a boolean if a field has been set.
func (o *APIProviderEnrolmentDetailsPatch) HasApiProvDomInfo() bool {
	if o != nil && !IsNil(o.ApiProvDomInfo) {
		return true
	}

	return false
}

// SetApiProvDomInfo gets a reference to the given string and assigns it to the ApiProvDomInfo field.
func (o *APIProviderEnrolmentDetailsPatch) SetApiProvDomInfo(v string) {
	o.ApiProvDomInfo = &v
}

func (o APIProviderEnrolmentDetailsPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIProviderEnrolmentDetailsPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiProvFuncs) {
		toSerialize["apiProvFuncs"] = o.ApiProvFuncs
	}
	if !IsNil(o.ApiProvDomInfo) {
		toSerialize["apiProvDomInfo"] = o.ApiProvDomInfo
	}
	return toSerialize, nil
}

type NullableAPIProviderEnrolmentDetailsPatch struct {
	value *APIProviderEnrolmentDetailsPatch
	isSet bool
}

func (v NullableAPIProviderEnrolmentDetailsPatch) Get() *APIProviderEnrolmentDetailsPatch {
	return v.value
}

func (v *NullableAPIProviderEnrolmentDetailsPatch) Set(val *APIProviderEnrolmentDetailsPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIProviderEnrolmentDetailsPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIProviderEnrolmentDetailsPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIProviderEnrolmentDetailsPatch(val *APIProviderEnrolmentDetailsPatch) *NullableAPIProviderEnrolmentDetailsPatch {
	return &NullableAPIProviderEnrolmentDetailsPatch{value: val, isSet: true}
}

func (v NullableAPIProviderEnrolmentDetailsPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIProviderEnrolmentDetailsPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

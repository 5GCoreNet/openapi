/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
)

// checks if the ProblemDetails1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProblemDetails1{}

// ProblemDetails1 Provides additional information in an error response.
type ProblemDetails1 struct {
	// String providing an URI formatted according to RFC 3986.
	Type   *string `json:"type,omitempty"`
	Title  *string `json:"title,omitempty"`
	Status *int32  `json:"status,omitempty"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	Instance *string `json:"instance,omitempty"`
	// A machine-readable application error cause specific to this occurrence of the problem.  This IE should be present and provide application-related error information, if available.
	Cause         *string         `json:"cause,omitempty"`
	InvalidParams []InvalidParam1 `json:"invalidParams,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures  *string         `json:"supportedFeatures,omitempty"`
	AccessTokenError   *AccessTokenErr `json:"accessTokenError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	// Fully Qualified Domain Name
	NrfId                *string  `json:"nrfId,omitempty"`
	SupportedApiVersions []string `json:"supportedApiVersions,omitempty"`
}

// NewProblemDetails1 instantiates a new ProblemDetails1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemDetails1() *ProblemDetails1 {
	this := ProblemDetails1{}
	return &this
}

// NewProblemDetails1WithDefaults instantiates a new ProblemDetails1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemDetails1WithDefaults() *ProblemDetails1 {
	this := ProblemDetails1{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProblemDetails1) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails1) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProblemDetails1) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProblemDetails1) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ProblemDetails1) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails1) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ProblemDetails1) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ProblemDetails1) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProblemDetails1) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails1) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProblemDetails1) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ProblemDetails1) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ProblemDetails1) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails1) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ProblemDetails1) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ProblemDetails1) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ProblemDetails1) GetInstance() string {
	if o == nil || IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails1) GetInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ProblemDetails1) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ProblemDetails1) SetInstance(v string) {
	o.Instance = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *ProblemDetails1) GetCause() string {
	if o == nil || IsNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails1) GetCauseOk() (*string, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *ProblemDetails1) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *ProblemDetails1) SetCause(v string) {
	o.Cause = &v
}

// GetInvalidParams returns the InvalidParams field value if set, zero value otherwise.
func (o *ProblemDetails1) GetInvalidParams() []InvalidParam1 {
	if o == nil || IsNil(o.InvalidParams) {
		var ret []InvalidParam1
		return ret
	}
	return o.InvalidParams
}

// GetInvalidParamsOk returns a tuple with the InvalidParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails1) GetInvalidParamsOk() ([]InvalidParam1, bool) {
	if o == nil || IsNil(o.InvalidParams) {
		return nil, false
	}
	return o.InvalidParams, true
}

// HasInvalidParams returns a boolean if a field has been set.
func (o *ProblemDetails1) HasInvalidParams() bool {
	if o != nil && !IsNil(o.InvalidParams) {
		return true
	}

	return false
}

// SetInvalidParams gets a reference to the given []InvalidParam1 and assigns it to the InvalidParams field.
func (o *ProblemDetails1) SetInvalidParams(v []InvalidParam1) {
	o.InvalidParams = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ProblemDetails1) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails1) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ProblemDetails1) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ProblemDetails1) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetAccessTokenError returns the AccessTokenError field value if set, zero value otherwise.
func (o *ProblemDetails1) GetAccessTokenError() AccessTokenErr {
	if o == nil || IsNil(o.AccessTokenError) {
		var ret AccessTokenErr
		return ret
	}
	return *o.AccessTokenError
}

// GetAccessTokenErrorOk returns a tuple with the AccessTokenError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails1) GetAccessTokenErrorOk() (*AccessTokenErr, bool) {
	if o == nil || IsNil(o.AccessTokenError) {
		return nil, false
	}
	return o.AccessTokenError, true
}

// HasAccessTokenError returns a boolean if a field has been set.
func (o *ProblemDetails1) HasAccessTokenError() bool {
	if o != nil && !IsNil(o.AccessTokenError) {
		return true
	}

	return false
}

// SetAccessTokenError gets a reference to the given AccessTokenErr and assigns it to the AccessTokenError field.
func (o *ProblemDetails1) SetAccessTokenError(v AccessTokenErr) {
	o.AccessTokenError = &v
}

// GetAccessTokenRequest returns the AccessTokenRequest field value if set, zero value otherwise.
func (o *ProblemDetails1) GetAccessTokenRequest() AccessTokenReq {
	if o == nil || IsNil(o.AccessTokenRequest) {
		var ret AccessTokenReq
		return ret
	}
	return *o.AccessTokenRequest
}

// GetAccessTokenRequestOk returns a tuple with the AccessTokenRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails1) GetAccessTokenRequestOk() (*AccessTokenReq, bool) {
	if o == nil || IsNil(o.AccessTokenRequest) {
		return nil, false
	}
	return o.AccessTokenRequest, true
}

// HasAccessTokenRequest returns a boolean if a field has been set.
func (o *ProblemDetails1) HasAccessTokenRequest() bool {
	if o != nil && !IsNil(o.AccessTokenRequest) {
		return true
	}

	return false
}

// SetAccessTokenRequest gets a reference to the given AccessTokenReq and assigns it to the AccessTokenRequest field.
func (o *ProblemDetails1) SetAccessTokenRequest(v AccessTokenReq) {
	o.AccessTokenRequest = &v
}

// GetNrfId returns the NrfId field value if set, zero value otherwise.
func (o *ProblemDetails1) GetNrfId() string {
	if o == nil || IsNil(o.NrfId) {
		var ret string
		return ret
	}
	return *o.NrfId
}

// GetNrfIdOk returns a tuple with the NrfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails1) GetNrfIdOk() (*string, bool) {
	if o == nil || IsNil(o.NrfId) {
		return nil, false
	}
	return o.NrfId, true
}

// HasNrfId returns a boolean if a field has been set.
func (o *ProblemDetails1) HasNrfId() bool {
	if o != nil && !IsNil(o.NrfId) {
		return true
	}

	return false
}

// SetNrfId gets a reference to the given string and assigns it to the NrfId field.
func (o *ProblemDetails1) SetNrfId(v string) {
	o.NrfId = &v
}

// GetSupportedApiVersions returns the SupportedApiVersions field value if set, zero value otherwise.
func (o *ProblemDetails1) GetSupportedApiVersions() []string {
	if o == nil || IsNil(o.SupportedApiVersions) {
		var ret []string
		return ret
	}
	return o.SupportedApiVersions
}

// GetSupportedApiVersionsOk returns a tuple with the SupportedApiVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails1) GetSupportedApiVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedApiVersions) {
		return nil, false
	}
	return o.SupportedApiVersions, true
}

// HasSupportedApiVersions returns a boolean if a field has been set.
func (o *ProblemDetails1) HasSupportedApiVersions() bool {
	if o != nil && !IsNil(o.SupportedApiVersions) {
		return true
	}

	return false
}

// SetSupportedApiVersions gets a reference to the given []string and assigns it to the SupportedApiVersions field.
func (o *ProblemDetails1) SetSupportedApiVersions(v []string) {
	o.SupportedApiVersions = v
}

func (o ProblemDetails1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProblemDetails1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !IsNil(o.InvalidParams) {
		toSerialize["invalidParams"] = o.InvalidParams
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.AccessTokenError) {
		toSerialize["accessTokenError"] = o.AccessTokenError
	}
	if !IsNil(o.AccessTokenRequest) {
		toSerialize["accessTokenRequest"] = o.AccessTokenRequest
	}
	if !IsNil(o.NrfId) {
		toSerialize["nrfId"] = o.NrfId
	}
	if !IsNil(o.SupportedApiVersions) {
		toSerialize["supportedApiVersions"] = o.SupportedApiVersions
	}
	return toSerialize, nil
}

type NullableProblemDetails1 struct {
	value *ProblemDetails1
	isSet bool
}

func (v NullableProblemDetails1) Get() *ProblemDetails1 {
	return v.value
}

func (v *NullableProblemDetails1) Set(val *ProblemDetails1) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemDetails1) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemDetails1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemDetails1(val *ProblemDetails1) *NullableProblemDetails1 {
	return &NullableProblemDetails1{value: val, isSet: true}
}

func (v NullableProblemDetails1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemDetails1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

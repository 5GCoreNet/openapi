/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.4.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsf_Management

import (
	"encoding/json"
)

// checks if the MbsExtProblemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsExtProblemDetails{}

// MbsExtProblemDetails Contains the FQDN or IP endpoints of the existing PCF and the cause value if there is an  existing PCF binding information for the MBS session. 
type MbsExtProblemDetails struct {
	// String providing an URI formatted according to RFC 3986.
	Type *string `json:"type,omitempty"`
	Title *string `json:"title,omitempty"`
	Status *int32 `json:"status,omitempty"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	Instance *string `json:"instance,omitempty"`
	// A machine-readable application error cause specific to this occurrence of the problem.  This IE should be present and provide application-related error information, if available. 
	Cause *string `json:"cause,omitempty"`
	InvalidParams []InvalidParam `json:"invalidParams,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	AccessTokenError *AccessTokenErr `json:"accessTokenError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	// Fully Qualified Domain Name
	NrfId *string `json:"nrfId,omitempty"`
	SupportedApiVersions []string `json:"supportedApiVersions,omitempty"`
	// Fully Qualified Domain Name
	PcfFqdn *string `json:"pcfFqdn,omitempty"`
	// IP end points of the PCF handling the MBS Session.
	PcfIpEndPoints []IpEndPoint `json:"pcfIpEndPoints,omitempty"`
}

// NewMbsExtProblemDetails instantiates a new MbsExtProblemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsExtProblemDetails() *MbsExtProblemDetails {
	this := MbsExtProblemDetails{}
	return &this
}

// NewMbsExtProblemDetailsWithDefaults instantiates a new MbsExtProblemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsExtProblemDetailsWithDefaults() *MbsExtProblemDetails {
	this := MbsExtProblemDetails{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MbsExtProblemDetails) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MbsExtProblemDetails) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *MbsExtProblemDetails) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetDetail() string {
	if o == nil || isNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetDetailOk() (*string, bool) {
	if o == nil || isNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasDetail() bool {
	if o != nil && !isNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *MbsExtProblemDetails) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetInstance() string {
	if o == nil || isNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetInstanceOk() (*string, bool) {
	if o == nil || isNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasInstance() bool {
	if o != nil && !isNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *MbsExtProblemDetails) SetInstance(v string) {
	o.Instance = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetCause() string {
	if o == nil || isNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetCauseOk() (*string, bool) {
	if o == nil || isNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasCause() bool {
	if o != nil && !isNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *MbsExtProblemDetails) SetCause(v string) {
	o.Cause = &v
}

// GetInvalidParams returns the InvalidParams field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetInvalidParams() []InvalidParam {
	if o == nil || isNil(o.InvalidParams) {
		var ret []InvalidParam
		return ret
	}
	return o.InvalidParams
}

// GetInvalidParamsOk returns a tuple with the InvalidParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetInvalidParamsOk() ([]InvalidParam, bool) {
	if o == nil || isNil(o.InvalidParams) {
		return nil, false
	}
	return o.InvalidParams, true
}

// HasInvalidParams returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasInvalidParams() bool {
	if o != nil && !isNil(o.InvalidParams) {
		return true
	}

	return false
}

// SetInvalidParams gets a reference to the given []InvalidParam and assigns it to the InvalidParams field.
func (o *MbsExtProblemDetails) SetInvalidParams(v []InvalidParam) {
	o.InvalidParams = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *MbsExtProblemDetails) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetAccessTokenError returns the AccessTokenError field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetAccessTokenError() AccessTokenErr {
	if o == nil || isNil(o.AccessTokenError) {
		var ret AccessTokenErr
		return ret
	}
	return *o.AccessTokenError
}

// GetAccessTokenErrorOk returns a tuple with the AccessTokenError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetAccessTokenErrorOk() (*AccessTokenErr, bool) {
	if o == nil || isNil(o.AccessTokenError) {
		return nil, false
	}
	return o.AccessTokenError, true
}

// HasAccessTokenError returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasAccessTokenError() bool {
	if o != nil && !isNil(o.AccessTokenError) {
		return true
	}

	return false
}

// SetAccessTokenError gets a reference to the given AccessTokenErr and assigns it to the AccessTokenError field.
func (o *MbsExtProblemDetails) SetAccessTokenError(v AccessTokenErr) {
	o.AccessTokenError = &v
}

// GetAccessTokenRequest returns the AccessTokenRequest field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetAccessTokenRequest() AccessTokenReq {
	if o == nil || isNil(o.AccessTokenRequest) {
		var ret AccessTokenReq
		return ret
	}
	return *o.AccessTokenRequest
}

// GetAccessTokenRequestOk returns a tuple with the AccessTokenRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetAccessTokenRequestOk() (*AccessTokenReq, bool) {
	if o == nil || isNil(o.AccessTokenRequest) {
		return nil, false
	}
	return o.AccessTokenRequest, true
}

// HasAccessTokenRequest returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasAccessTokenRequest() bool {
	if o != nil && !isNil(o.AccessTokenRequest) {
		return true
	}

	return false
}

// SetAccessTokenRequest gets a reference to the given AccessTokenReq and assigns it to the AccessTokenRequest field.
func (o *MbsExtProblemDetails) SetAccessTokenRequest(v AccessTokenReq) {
	o.AccessTokenRequest = &v
}

// GetNrfId returns the NrfId field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetNrfId() string {
	if o == nil || isNil(o.NrfId) {
		var ret string
		return ret
	}
	return *o.NrfId
}

// GetNrfIdOk returns a tuple with the NrfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetNrfIdOk() (*string, bool) {
	if o == nil || isNil(o.NrfId) {
		return nil, false
	}
	return o.NrfId, true
}

// HasNrfId returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasNrfId() bool {
	if o != nil && !isNil(o.NrfId) {
		return true
	}

	return false
}

// SetNrfId gets a reference to the given string and assigns it to the NrfId field.
func (o *MbsExtProblemDetails) SetNrfId(v string) {
	o.NrfId = &v
}

// GetSupportedApiVersions returns the SupportedApiVersions field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetSupportedApiVersions() []string {
	if o == nil || isNil(o.SupportedApiVersions) {
		var ret []string
		return ret
	}
	return o.SupportedApiVersions
}

// GetSupportedApiVersionsOk returns a tuple with the SupportedApiVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetSupportedApiVersionsOk() ([]string, bool) {
	if o == nil || isNil(o.SupportedApiVersions) {
		return nil, false
	}
	return o.SupportedApiVersions, true
}

// HasSupportedApiVersions returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasSupportedApiVersions() bool {
	if o != nil && !isNil(o.SupportedApiVersions) {
		return true
	}

	return false
}

// SetSupportedApiVersions gets a reference to the given []string and assigns it to the SupportedApiVersions field.
func (o *MbsExtProblemDetails) SetSupportedApiVersions(v []string) {
	o.SupportedApiVersions = v
}

// GetPcfFqdn returns the PcfFqdn field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetPcfFqdn() string {
	if o == nil || isNil(o.PcfFqdn) {
		var ret string
		return ret
	}
	return *o.PcfFqdn
}

// GetPcfFqdnOk returns a tuple with the PcfFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetPcfFqdnOk() (*string, bool) {
	if o == nil || isNil(o.PcfFqdn) {
		return nil, false
	}
	return o.PcfFqdn, true
}

// HasPcfFqdn returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasPcfFqdn() bool {
	if o != nil && !isNil(o.PcfFqdn) {
		return true
	}

	return false
}

// SetPcfFqdn gets a reference to the given string and assigns it to the PcfFqdn field.
func (o *MbsExtProblemDetails) SetPcfFqdn(v string) {
	o.PcfFqdn = &v
}

// GetPcfIpEndPoints returns the PcfIpEndPoints field value if set, zero value otherwise.
func (o *MbsExtProblemDetails) GetPcfIpEndPoints() []IpEndPoint {
	if o == nil || isNil(o.PcfIpEndPoints) {
		var ret []IpEndPoint
		return ret
	}
	return o.PcfIpEndPoints
}

// GetPcfIpEndPointsOk returns a tuple with the PcfIpEndPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsExtProblemDetails) GetPcfIpEndPointsOk() ([]IpEndPoint, bool) {
	if o == nil || isNil(o.PcfIpEndPoints) {
		return nil, false
	}
	return o.PcfIpEndPoints, true
}

// HasPcfIpEndPoints returns a boolean if a field has been set.
func (o *MbsExtProblemDetails) HasPcfIpEndPoints() bool {
	if o != nil && !isNil(o.PcfIpEndPoints) {
		return true
	}

	return false
}

// SetPcfIpEndPoints gets a reference to the given []IpEndPoint and assigns it to the PcfIpEndPoints field.
func (o *MbsExtProblemDetails) SetPcfIpEndPoints(v []IpEndPoint) {
	o.PcfIpEndPoints = v
}

func (o MbsExtProblemDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsExtProblemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !isNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !isNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !isNil(o.InvalidParams) {
		toSerialize["invalidParams"] = o.InvalidParams
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.AccessTokenError) {
		toSerialize["accessTokenError"] = o.AccessTokenError
	}
	if !isNil(o.AccessTokenRequest) {
		toSerialize["accessTokenRequest"] = o.AccessTokenRequest
	}
	if !isNil(o.NrfId) {
		toSerialize["nrfId"] = o.NrfId
	}
	if !isNil(o.SupportedApiVersions) {
		toSerialize["supportedApiVersions"] = o.SupportedApiVersions
	}
	if !isNil(o.PcfFqdn) {
		toSerialize["pcfFqdn"] = o.PcfFqdn
	}
	if !isNil(o.PcfIpEndPoints) {
		toSerialize["pcfIpEndPoints"] = o.PcfIpEndPoints
	}
	return toSerialize, nil
}

type NullableMbsExtProblemDetails struct {
	value *MbsExtProblemDetails
	isSet bool
}

func (v NullableMbsExtProblemDetails) Get() *MbsExtProblemDetails {
	return v.value
}

func (v *NullableMbsExtProblemDetails) Set(val *MbsExtProblemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsExtProblemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsExtProblemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsExtProblemDetails(val *MbsExtProblemDetails) *NullableMbsExtProblemDetails {
	return &NullableMbsExtProblemDetails{value: val, isSet: true}
}

func (v NullableMbsExtProblemDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsExtProblemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



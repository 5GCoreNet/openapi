/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EASDiscovery

import (
	"encoding/json"
)

// checks if the EasDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EasDetail{}

// EasDetail EAS details.
type EasDetail struct {
	// Application identifier of the EAS.
	EasId string `json:"easId"`
	ExpectedSvcKPIs *ACServiceKPIs `json:"expectedSvcKPIs,omitempty"`
	MinimumReqSvcKPIs *ACServiceKPIs `json:"minimumReqSvcKPIs,omitempty"`
}

// NewEasDetail instantiates a new EasDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasDetail(easId string) *EasDetail {
	this := EasDetail{}
	this.EasId = easId
	return &this
}

// NewEasDetailWithDefaults instantiates a new EasDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasDetailWithDefaults() *EasDetail {
	this := EasDetail{}
	return &this
}

// GetEasId returns the EasId field value
func (o *EasDetail) GetEasId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EasId
}

// GetEasIdOk returns a tuple with the EasId field value
// and a boolean to check if the value has been set.
func (o *EasDetail) GetEasIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EasId, true
}

// SetEasId sets field value
func (o *EasDetail) SetEasId(v string) {
	o.EasId = v
}

// GetExpectedSvcKPIs returns the ExpectedSvcKPIs field value if set, zero value otherwise.
func (o *EasDetail) GetExpectedSvcKPIs() ACServiceKPIs {
	if o == nil || IsNil(o.ExpectedSvcKPIs) {
		var ret ACServiceKPIs
		return ret
	}
	return *o.ExpectedSvcKPIs
}

// GetExpectedSvcKPIsOk returns a tuple with the ExpectedSvcKPIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDetail) GetExpectedSvcKPIsOk() (*ACServiceKPIs, bool) {
	if o == nil || IsNil(o.ExpectedSvcKPIs) {
		return nil, false
	}
	return o.ExpectedSvcKPIs, true
}

// HasExpectedSvcKPIs returns a boolean if a field has been set.
func (o *EasDetail) HasExpectedSvcKPIs() bool {
	if o != nil && !IsNil(o.ExpectedSvcKPIs) {
		return true
	}

	return false
}

// SetExpectedSvcKPIs gets a reference to the given ACServiceKPIs and assigns it to the ExpectedSvcKPIs field.
func (o *EasDetail) SetExpectedSvcKPIs(v ACServiceKPIs) {
	o.ExpectedSvcKPIs = &v
}

// GetMinimumReqSvcKPIs returns the MinimumReqSvcKPIs field value if set, zero value otherwise.
func (o *EasDetail) GetMinimumReqSvcKPIs() ACServiceKPIs {
	if o == nil || IsNil(o.MinimumReqSvcKPIs) {
		var ret ACServiceKPIs
		return ret
	}
	return *o.MinimumReqSvcKPIs
}

// GetMinimumReqSvcKPIsOk returns a tuple with the MinimumReqSvcKPIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDetail) GetMinimumReqSvcKPIsOk() (*ACServiceKPIs, bool) {
	if o == nil || IsNil(o.MinimumReqSvcKPIs) {
		return nil, false
	}
	return o.MinimumReqSvcKPIs, true
}

// HasMinimumReqSvcKPIs returns a boolean if a field has been set.
func (o *EasDetail) HasMinimumReqSvcKPIs() bool {
	if o != nil && !IsNil(o.MinimumReqSvcKPIs) {
		return true
	}

	return false
}

// SetMinimumReqSvcKPIs gets a reference to the given ACServiceKPIs and assigns it to the MinimumReqSvcKPIs field.
func (o *EasDetail) SetMinimumReqSvcKPIs(v ACServiceKPIs) {
	o.MinimumReqSvcKPIs = &v
}

func (o EasDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EasDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["easId"] = o.EasId
	if !IsNil(o.ExpectedSvcKPIs) {
		toSerialize["expectedSvcKPIs"] = o.ExpectedSvcKPIs
	}
	if !IsNil(o.MinimumReqSvcKPIs) {
		toSerialize["minimumReqSvcKPIs"] = o.MinimumReqSvcKPIs
	}
	return toSerialize, nil
}

type NullableEasDetail struct {
	value *EasDetail
	isSet bool
}

func (v NullableEasDetail) Get() *EasDetail {
	return v.value
}

func (v *NullableEasDetail) Set(val *EasDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableEasDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableEasDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasDetail(val *EasDetail) *NullableEasDetail {
	return &NullableEasDetail{value: val, isSet: true}
}

func (v NullableEasDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



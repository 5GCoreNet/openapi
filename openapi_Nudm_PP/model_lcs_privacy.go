/*
Nudm_PP

Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_PP

import (
	"encoding/json"
)

// checks if the LcsPrivacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LcsPrivacy{}

// LcsPrivacy struct for LcsPrivacy
type LcsPrivacy struct {
	AfInstanceId *string `json:"afInstanceId,omitempty"`
	ReferenceId *int32 `json:"referenceId,omitempty"`
	Lpi *Lpi `json:"lpi,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation *string `json:"mtcProviderInformation,omitempty"`
}

// NewLcsPrivacy instantiates a new LcsPrivacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLcsPrivacy() *LcsPrivacy {
	this := LcsPrivacy{}
	return &this
}

// NewLcsPrivacyWithDefaults instantiates a new LcsPrivacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLcsPrivacyWithDefaults() *LcsPrivacy {
	this := LcsPrivacy{}
	return &this
}

// GetAfInstanceId returns the AfInstanceId field value if set, zero value otherwise.
func (o *LcsPrivacy) GetAfInstanceId() string {
	if o == nil || isNil(o.AfInstanceId) {
		var ret string
		return ret
	}
	return *o.AfInstanceId
}

// GetAfInstanceIdOk returns a tuple with the AfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsPrivacy) GetAfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.AfInstanceId) {
		return nil, false
	}
	return o.AfInstanceId, true
}

// HasAfInstanceId returns a boolean if a field has been set.
func (o *LcsPrivacy) HasAfInstanceId() bool {
	if o != nil && !isNil(o.AfInstanceId) {
		return true
	}

	return false
}

// SetAfInstanceId gets a reference to the given string and assigns it to the AfInstanceId field.
func (o *LcsPrivacy) SetAfInstanceId(v string) {
	o.AfInstanceId = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *LcsPrivacy) GetReferenceId() int32 {
	if o == nil || isNil(o.ReferenceId) {
		var ret int32
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsPrivacy) GetReferenceIdOk() (*int32, bool) {
	if o == nil || isNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *LcsPrivacy) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given int32 and assigns it to the ReferenceId field.
func (o *LcsPrivacy) SetReferenceId(v int32) {
	o.ReferenceId = &v
}

// GetLpi returns the Lpi field value if set, zero value otherwise.
func (o *LcsPrivacy) GetLpi() Lpi {
	if o == nil || isNil(o.Lpi) {
		var ret Lpi
		return ret
	}
	return *o.Lpi
}

// GetLpiOk returns a tuple with the Lpi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsPrivacy) GetLpiOk() (*Lpi, bool) {
	if o == nil || isNil(o.Lpi) {
		return nil, false
	}
	return o.Lpi, true
}

// HasLpi returns a boolean if a field has been set.
func (o *LcsPrivacy) HasLpi() bool {
	if o != nil && !isNil(o.Lpi) {
		return true
	}

	return false
}

// SetLpi gets a reference to the given Lpi and assigns it to the Lpi field.
func (o *LcsPrivacy) SetLpi(v Lpi) {
	o.Lpi = &v
}

// GetMtcProviderInformation returns the MtcProviderInformation field value if set, zero value otherwise.
func (o *LcsPrivacy) GetMtcProviderInformation() string {
	if o == nil || isNil(o.MtcProviderInformation) {
		var ret string
		return ret
	}
	return *o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsPrivacy) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil || isNil(o.MtcProviderInformation) {
		return nil, false
	}
	return o.MtcProviderInformation, true
}

// HasMtcProviderInformation returns a boolean if a field has been set.
func (o *LcsPrivacy) HasMtcProviderInformation() bool {
	if o != nil && !isNil(o.MtcProviderInformation) {
		return true
	}

	return false
}

// SetMtcProviderInformation gets a reference to the given string and assigns it to the MtcProviderInformation field.
func (o *LcsPrivacy) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = &v
}

func (o LcsPrivacy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LcsPrivacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AfInstanceId) {
		toSerialize["afInstanceId"] = o.AfInstanceId
	}
	if !isNil(o.ReferenceId) {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if !isNil(o.Lpi) {
		toSerialize["lpi"] = o.Lpi
	}
	if !isNil(o.MtcProviderInformation) {
		toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	}
	return toSerialize, nil
}

type NullableLcsPrivacy struct {
	value *LcsPrivacy
	isSet bool
}

func (v NullableLcsPrivacy) Get() *LcsPrivacy {
	return v.value
}

func (v *NullableLcsPrivacy) Set(val *LcsPrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsPrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsPrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsPrivacy(val *LcsPrivacy) *NullableLcsPrivacy {
	return &NullableLcsPrivacy{value: val, isSet: true}
}

func (v NullableLcsPrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsPrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



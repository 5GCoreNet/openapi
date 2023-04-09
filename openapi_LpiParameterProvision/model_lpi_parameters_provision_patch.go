/*
3gpp-lpi-pp

API for Location Privacy Indication Parameters Provisioning.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_LpiParameterProvision

import (
	"encoding/json"
)

// checks if the LpiParametersProvisionPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LpiParametersProvisionPatch{}

// LpiParametersProvisionPatch Represents the parameters to modify an existing Individual LPI Parameters Provisionings resource.
type LpiParametersProvisionPatch struct {
	Lpi *Lpi `json:"lpi,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderId *string `json:"mtcProviderId,omitempty"`
}

// NewLpiParametersProvisionPatch instantiates a new LpiParametersProvisionPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLpiParametersProvisionPatch() *LpiParametersProvisionPatch {
	this := LpiParametersProvisionPatch{}
	return &this
}

// NewLpiParametersProvisionPatchWithDefaults instantiates a new LpiParametersProvisionPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLpiParametersProvisionPatchWithDefaults() *LpiParametersProvisionPatch {
	this := LpiParametersProvisionPatch{}
	return &this
}

// GetLpi returns the Lpi field value if set, zero value otherwise.
func (o *LpiParametersProvisionPatch) GetLpi() Lpi {
	if o == nil || IsNil(o.Lpi) {
		var ret Lpi
		return ret
	}
	return *o.Lpi
}

// GetLpiOk returns a tuple with the Lpi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LpiParametersProvisionPatch) GetLpiOk() (*Lpi, bool) {
	if o == nil || IsNil(o.Lpi) {
		return nil, false
	}
	return o.Lpi, true
}

// HasLpi returns a boolean if a field has been set.
func (o *LpiParametersProvisionPatch) HasLpi() bool {
	if o != nil && !IsNil(o.Lpi) {
		return true
	}

	return false
}

// SetLpi gets a reference to the given Lpi and assigns it to the Lpi field.
func (o *LpiParametersProvisionPatch) SetLpi(v Lpi) {
	o.Lpi = &v
}

// GetMtcProviderId returns the MtcProviderId field value if set, zero value otherwise.
func (o *LpiParametersProvisionPatch) GetMtcProviderId() string {
	if o == nil || IsNil(o.MtcProviderId) {
		var ret string
		return ret
	}
	return *o.MtcProviderId
}

// GetMtcProviderIdOk returns a tuple with the MtcProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LpiParametersProvisionPatch) GetMtcProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.MtcProviderId) {
		return nil, false
	}
	return o.MtcProviderId, true
}

// HasMtcProviderId returns a boolean if a field has been set.
func (o *LpiParametersProvisionPatch) HasMtcProviderId() bool {
	if o != nil && !IsNil(o.MtcProviderId) {
		return true
	}

	return false
}

// SetMtcProviderId gets a reference to the given string and assigns it to the MtcProviderId field.
func (o *LpiParametersProvisionPatch) SetMtcProviderId(v string) {
	o.MtcProviderId = &v
}

func (o LpiParametersProvisionPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LpiParametersProvisionPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lpi) {
		toSerialize["lpi"] = o.Lpi
	}
	if !IsNil(o.MtcProviderId) {
		toSerialize["mtcProviderId"] = o.MtcProviderId
	}
	return toSerialize, nil
}

type NullableLpiParametersProvisionPatch struct {
	value *LpiParametersProvisionPatch
	isSet bool
}

func (v NullableLpiParametersProvisionPatch) Get() *LpiParametersProvisionPatch {
	return v.value
}

func (v *NullableLpiParametersProvisionPatch) Set(val *LpiParametersProvisionPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableLpiParametersProvisionPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableLpiParametersProvisionPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLpiParametersProvisionPatch(val *LpiParametersProvisionPatch) *NullableLpiParametersProvisionPatch {
	return &NullableLpiParametersProvisionPatch{value: val, isSet: true}
}

func (v NullableLpiParametersProvisionPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLpiParametersProvisionPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

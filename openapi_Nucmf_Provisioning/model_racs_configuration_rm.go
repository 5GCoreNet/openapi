/*
Nucmf_Provisioning

UCMF_Provisioning Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nucmf_Provisioning

import (
	"encoding/json"
)

// checks if the RacsConfigurationRm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RacsConfigurationRm{}

// RacsConfigurationRm Represents the same as the RacsConfiguration data type but with the nullable:true property.
type RacsConfigurationRm struct {
	// The UE radio capability data in EPS.
	RacsParamEps NullableString `json:"racsParamEps,omitempty"`
	// The UE radio capability data in 5GS.
	RacsParam5Gs NullableString `json:"racsParam5Gs,omitempty"`
	// Related UE model's IMEI-TAC values.
	ImeiTacs []string `json:"imeiTacs,omitempty"`
}

// NewRacsConfigurationRm instantiates a new RacsConfigurationRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRacsConfigurationRm() *RacsConfigurationRm {
	this := RacsConfigurationRm{}
	return &this
}

// NewRacsConfigurationRmWithDefaults instantiates a new RacsConfigurationRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRacsConfigurationRmWithDefaults() *RacsConfigurationRm {
	this := RacsConfigurationRm{}
	return &this
}

// GetRacsParamEps returns the RacsParamEps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RacsConfigurationRm) GetRacsParamEps() string {
	if o == nil || IsNil(o.RacsParamEps.Get()) {
		var ret string
		return ret
	}
	return *o.RacsParamEps.Get()
}

// GetRacsParamEpsOk returns a tuple with the RacsParamEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RacsConfigurationRm) GetRacsParamEpsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RacsParamEps.Get(), o.RacsParamEps.IsSet()
}

// HasRacsParamEps returns a boolean if a field has been set.
func (o *RacsConfigurationRm) HasRacsParamEps() bool {
	if o != nil && o.RacsParamEps.IsSet() {
		return true
	}

	return false
}

// SetRacsParamEps gets a reference to the given NullableString and assigns it to the RacsParamEps field.
func (o *RacsConfigurationRm) SetRacsParamEps(v string) {
	o.RacsParamEps.Set(&v)
}

// SetRacsParamEpsNil sets the value for RacsParamEps to be an explicit nil
func (o *RacsConfigurationRm) SetRacsParamEpsNil() {
	o.RacsParamEps.Set(nil)
}

// UnsetRacsParamEps ensures that no value is present for RacsParamEps, not even an explicit nil
func (o *RacsConfigurationRm) UnsetRacsParamEps() {
	o.RacsParamEps.Unset()
}

// GetRacsParam5Gs returns the RacsParam5Gs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RacsConfigurationRm) GetRacsParam5Gs() string {
	if o == nil || IsNil(o.RacsParam5Gs.Get()) {
		var ret string
		return ret
	}
	return *o.RacsParam5Gs.Get()
}

// GetRacsParam5GsOk returns a tuple with the RacsParam5Gs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RacsConfigurationRm) GetRacsParam5GsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RacsParam5Gs.Get(), o.RacsParam5Gs.IsSet()
}

// HasRacsParam5Gs returns a boolean if a field has been set.
func (o *RacsConfigurationRm) HasRacsParam5Gs() bool {
	if o != nil && o.RacsParam5Gs.IsSet() {
		return true
	}

	return false
}

// SetRacsParam5Gs gets a reference to the given NullableString and assigns it to the RacsParam5Gs field.
func (o *RacsConfigurationRm) SetRacsParam5Gs(v string) {
	o.RacsParam5Gs.Set(&v)
}

// SetRacsParam5GsNil sets the value for RacsParam5Gs to be an explicit nil
func (o *RacsConfigurationRm) SetRacsParam5GsNil() {
	o.RacsParam5Gs.Set(nil)
}

// UnsetRacsParam5Gs ensures that no value is present for RacsParam5Gs, not even an explicit nil
func (o *RacsConfigurationRm) UnsetRacsParam5Gs() {
	o.RacsParam5Gs.Unset()
}

// GetImeiTacs returns the ImeiTacs field value if set, zero value otherwise.
func (o *RacsConfigurationRm) GetImeiTacs() []string {
	if o == nil || IsNil(o.ImeiTacs) {
		var ret []string
		return ret
	}
	return o.ImeiTacs
}

// GetImeiTacsOk returns a tuple with the ImeiTacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RacsConfigurationRm) GetImeiTacsOk() ([]string, bool) {
	if o == nil || IsNil(o.ImeiTacs) {
		return nil, false
	}
	return o.ImeiTacs, true
}

// HasImeiTacs returns a boolean if a field has been set.
func (o *RacsConfigurationRm) HasImeiTacs() bool {
	if o != nil && !IsNil(o.ImeiTacs) {
		return true
	}

	return false
}

// SetImeiTacs gets a reference to the given []string and assigns it to the ImeiTacs field.
func (o *RacsConfigurationRm) SetImeiTacs(v []string) {
	o.ImeiTacs = v
}

func (o RacsConfigurationRm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RacsConfigurationRm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RacsParamEps.IsSet() {
		toSerialize["racsParamEps"] = o.RacsParamEps.Get()
	}
	if o.RacsParam5Gs.IsSet() {
		toSerialize["racsParam5Gs"] = o.RacsParam5Gs.Get()
	}
	if !IsNil(o.ImeiTacs) {
		toSerialize["imeiTacs"] = o.ImeiTacs
	}
	return toSerialize, nil
}

type NullableRacsConfigurationRm struct {
	value *RacsConfigurationRm
	isSet bool
}

func (v NullableRacsConfigurationRm) Get() *RacsConfigurationRm {
	return v.value
}

func (v *NullableRacsConfigurationRm) Set(val *RacsConfigurationRm) {
	v.value = val
	v.isSet = true
}

func (v NullableRacsConfigurationRm) IsSet() bool {
	return v.isSet
}

func (v *NullableRacsConfigurationRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRacsConfigurationRm(val *RacsConfigurationRm) *NullableRacsConfigurationRm {
	return &NullableRacsConfigurationRm{value: val, isSet: true}
}

func (v NullableRacsConfigurationRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRacsConfigurationRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

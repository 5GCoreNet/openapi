/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the PSCellInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PSCellInformation{}

// PSCellInformation struct for PSCellInformation
type PSCellInformation struct {
	Nrcgi *Ncgi `json:"nrcgi,omitempty"`
	Ecgi  *Ecgi `json:"ecgi,omitempty"`
}

// NewPSCellInformation instantiates a new PSCellInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPSCellInformation() *PSCellInformation {
	this := PSCellInformation{}
	return &this
}

// NewPSCellInformationWithDefaults instantiates a new PSCellInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPSCellInformationWithDefaults() *PSCellInformation {
	this := PSCellInformation{}
	return &this
}

// GetNrcgi returns the Nrcgi field value if set, zero value otherwise.
func (o *PSCellInformation) GetNrcgi() Ncgi {
	if o == nil || IsNil(o.Nrcgi) {
		var ret Ncgi
		return ret
	}
	return *o.Nrcgi
}

// GetNrcgiOk returns a tuple with the Nrcgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PSCellInformation) GetNrcgiOk() (*Ncgi, bool) {
	if o == nil || IsNil(o.Nrcgi) {
		return nil, false
	}
	return o.Nrcgi, true
}

// HasNrcgi returns a boolean if a field has been set.
func (o *PSCellInformation) HasNrcgi() bool {
	if o != nil && !IsNil(o.Nrcgi) {
		return true
	}

	return false
}

// SetNrcgi gets a reference to the given Ncgi and assigns it to the Nrcgi field.
func (o *PSCellInformation) SetNrcgi(v Ncgi) {
	o.Nrcgi = &v
}

// GetEcgi returns the Ecgi field value if set, zero value otherwise.
func (o *PSCellInformation) GetEcgi() Ecgi {
	if o == nil || IsNil(o.Ecgi) {
		var ret Ecgi
		return ret
	}
	return *o.Ecgi
}

// GetEcgiOk returns a tuple with the Ecgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PSCellInformation) GetEcgiOk() (*Ecgi, bool) {
	if o == nil || IsNil(o.Ecgi) {
		return nil, false
	}
	return o.Ecgi, true
}

// HasEcgi returns a boolean if a field has been set.
func (o *PSCellInformation) HasEcgi() bool {
	if o != nil && !IsNil(o.Ecgi) {
		return true
	}

	return false
}

// SetEcgi gets a reference to the given Ecgi and assigns it to the Ecgi field.
func (o *PSCellInformation) SetEcgi(v Ecgi) {
	o.Ecgi = &v
}

func (o PSCellInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PSCellInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nrcgi) {
		toSerialize["nrcgi"] = o.Nrcgi
	}
	if !IsNil(o.Ecgi) {
		toSerialize["ecgi"] = o.Ecgi
	}
	return toSerialize, nil
}

type NullablePSCellInformation struct {
	value *PSCellInformation
	isSet bool
}

func (v NullablePSCellInformation) Get() *PSCellInformation {
	return v.value
}

func (v *NullablePSCellInformation) Set(val *PSCellInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePSCellInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePSCellInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePSCellInformation(val *PSCellInformation) *NullablePSCellInformation {
	return &NullablePSCellInformation{value: val, isSet: true}
}

func (v NullablePSCellInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePSCellInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

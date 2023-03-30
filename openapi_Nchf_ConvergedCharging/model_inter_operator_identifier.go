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

// checks if the InterOperatorIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterOperatorIdentifier{}

// InterOperatorIdentifier struct for InterOperatorIdentifier
type InterOperatorIdentifier struct {
	OriginatingIOI *string `json:"originatingIOI,omitempty"`
	TerminatingIOI *string `json:"terminatingIOI,omitempty"`
}

// NewInterOperatorIdentifier instantiates a new InterOperatorIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterOperatorIdentifier() *InterOperatorIdentifier {
	this := InterOperatorIdentifier{}
	return &this
}

// NewInterOperatorIdentifierWithDefaults instantiates a new InterOperatorIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterOperatorIdentifierWithDefaults() *InterOperatorIdentifier {
	this := InterOperatorIdentifier{}
	return &this
}

// GetOriginatingIOI returns the OriginatingIOI field value if set, zero value otherwise.
func (o *InterOperatorIdentifier) GetOriginatingIOI() string {
	if o == nil || IsNil(o.OriginatingIOI) {
		var ret string
		return ret
	}
	return *o.OriginatingIOI
}

// GetOriginatingIOIOk returns a tuple with the OriginatingIOI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterOperatorIdentifier) GetOriginatingIOIOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatingIOI) {
		return nil, false
	}
	return o.OriginatingIOI, true
}

// HasOriginatingIOI returns a boolean if a field has been set.
func (o *InterOperatorIdentifier) HasOriginatingIOI() bool {
	if o != nil && !IsNil(o.OriginatingIOI) {
		return true
	}

	return false
}

// SetOriginatingIOI gets a reference to the given string and assigns it to the OriginatingIOI field.
func (o *InterOperatorIdentifier) SetOriginatingIOI(v string) {
	o.OriginatingIOI = &v
}

// GetTerminatingIOI returns the TerminatingIOI field value if set, zero value otherwise.
func (o *InterOperatorIdentifier) GetTerminatingIOI() string {
	if o == nil || IsNil(o.TerminatingIOI) {
		var ret string
		return ret
	}
	return *o.TerminatingIOI
}

// GetTerminatingIOIOk returns a tuple with the TerminatingIOI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterOperatorIdentifier) GetTerminatingIOIOk() (*string, bool) {
	if o == nil || IsNil(o.TerminatingIOI) {
		return nil, false
	}
	return o.TerminatingIOI, true
}

// HasTerminatingIOI returns a boolean if a field has been set.
func (o *InterOperatorIdentifier) HasTerminatingIOI() bool {
	if o != nil && !IsNil(o.TerminatingIOI) {
		return true
	}

	return false
}

// SetTerminatingIOI gets a reference to the given string and assigns it to the TerminatingIOI field.
func (o *InterOperatorIdentifier) SetTerminatingIOI(v string) {
	o.TerminatingIOI = &v
}

func (o InterOperatorIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterOperatorIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginatingIOI) {
		toSerialize["originatingIOI"] = o.OriginatingIOI
	}
	if !IsNil(o.TerminatingIOI) {
		toSerialize["terminatingIOI"] = o.TerminatingIOI
	}
	return toSerialize, nil
}

type NullableInterOperatorIdentifier struct {
	value *InterOperatorIdentifier
	isSet bool
}

func (v NullableInterOperatorIdentifier) Get() *InterOperatorIdentifier {
	return v.value
}

func (v *NullableInterOperatorIdentifier) Set(val *InterOperatorIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableInterOperatorIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableInterOperatorIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterOperatorIdentifier(val *InterOperatorIdentifier) *NullableInterOperatorIdentifier {
	return &NullableInterOperatorIdentifier{value: val, isSet: true}
}

func (v NullableInterOperatorIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterOperatorIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



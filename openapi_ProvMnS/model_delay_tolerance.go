/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the DelayTolerance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelayTolerance{}

// DelayTolerance struct for DelayTolerance
type DelayTolerance struct {
	ServAttrCom *ServAttrCom `json:"servAttrCom,omitempty"`
	Support *Support `json:"support,omitempty"`
}

// NewDelayTolerance instantiates a new DelayTolerance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelayTolerance() *DelayTolerance {
	this := DelayTolerance{}
	return &this
}

// NewDelayToleranceWithDefaults instantiates a new DelayTolerance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelayToleranceWithDefaults() *DelayTolerance {
	this := DelayTolerance{}
	return &this
}

// GetServAttrCom returns the ServAttrCom field value if set, zero value otherwise.
func (o *DelayTolerance) GetServAttrCom() ServAttrCom {
	if o == nil || IsNil(o.ServAttrCom) {
		var ret ServAttrCom
		return ret
	}
	return *o.ServAttrCom
}

// GetServAttrComOk returns a tuple with the ServAttrCom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayTolerance) GetServAttrComOk() (*ServAttrCom, bool) {
	if o == nil || IsNil(o.ServAttrCom) {
		return nil, false
	}
	return o.ServAttrCom, true
}

// HasServAttrCom returns a boolean if a field has been set.
func (o *DelayTolerance) HasServAttrCom() bool {
	if o != nil && !IsNil(o.ServAttrCom) {
		return true
	}

	return false
}

// SetServAttrCom gets a reference to the given ServAttrCom and assigns it to the ServAttrCom field.
func (o *DelayTolerance) SetServAttrCom(v ServAttrCom) {
	o.ServAttrCom = &v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *DelayTolerance) GetSupport() Support {
	if o == nil || IsNil(o.Support) {
		var ret Support
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayTolerance) GetSupportOk() (*Support, bool) {
	if o == nil || IsNil(o.Support) {
		return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *DelayTolerance) HasSupport() bool {
	if o != nil && !IsNil(o.Support) {
		return true
	}

	return false
}

// SetSupport gets a reference to the given Support and assigns it to the Support field.
func (o *DelayTolerance) SetSupport(v Support) {
	o.Support = &v
}

func (o DelayTolerance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelayTolerance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServAttrCom) {
		toSerialize["servAttrCom"] = o.ServAttrCom
	}
	if !IsNil(o.Support) {
		toSerialize["support"] = o.Support
	}
	return toSerialize, nil
}

type NullableDelayTolerance struct {
	value *DelayTolerance
	isSet bool
}

func (v NullableDelayTolerance) Get() *DelayTolerance {
	return v.value
}

func (v *NullableDelayTolerance) Set(val *DelayTolerance) {
	v.value = val
	v.isSet = true
}

func (v NullableDelayTolerance) IsSet() bool {
	return v.isSet
}

func (v *NullableDelayTolerance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelayTolerance(val *DelayTolerance) *NullableDelayTolerance {
	return &NullableDelayTolerance{value: val, isSet: true}
}

func (v NullableDelayTolerance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelayTolerance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



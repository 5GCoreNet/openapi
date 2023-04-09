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

// checks if the Synchronicity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Synchronicity{}

// Synchronicity struct for Synchronicity
type Synchronicity struct {
	ServAttrCom  *ServAttrCom     `json:"servAttrCom,omitempty"`
	Availability *SynAvailability `json:"availability,omitempty"`
	Accuracy     *float32         `json:"accuracy,omitempty"`
}

// NewSynchronicity instantiates a new Synchronicity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSynchronicity() *Synchronicity {
	this := Synchronicity{}
	return &this
}

// NewSynchronicityWithDefaults instantiates a new Synchronicity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSynchronicityWithDefaults() *Synchronicity {
	this := Synchronicity{}
	return &this
}

// GetServAttrCom returns the ServAttrCom field value if set, zero value otherwise.
func (o *Synchronicity) GetServAttrCom() ServAttrCom {
	if o == nil || IsNil(o.ServAttrCom) {
		var ret ServAttrCom
		return ret
	}
	return *o.ServAttrCom
}

// GetServAttrComOk returns a tuple with the ServAttrCom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Synchronicity) GetServAttrComOk() (*ServAttrCom, bool) {
	if o == nil || IsNil(o.ServAttrCom) {
		return nil, false
	}
	return o.ServAttrCom, true
}

// HasServAttrCom returns a boolean if a field has been set.
func (o *Synchronicity) HasServAttrCom() bool {
	if o != nil && !IsNil(o.ServAttrCom) {
		return true
	}

	return false
}

// SetServAttrCom gets a reference to the given ServAttrCom and assigns it to the ServAttrCom field.
func (o *Synchronicity) SetServAttrCom(v ServAttrCom) {
	o.ServAttrCom = &v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *Synchronicity) GetAvailability() SynAvailability {
	if o == nil || IsNil(o.Availability) {
		var ret SynAvailability
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Synchronicity) GetAvailabilityOk() (*SynAvailability, bool) {
	if o == nil || IsNil(o.Availability) {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *Synchronicity) HasAvailability() bool {
	if o != nil && !IsNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given SynAvailability and assigns it to the Availability field.
func (o *Synchronicity) SetAvailability(v SynAvailability) {
	o.Availability = &v
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *Synchronicity) GetAccuracy() float32 {
	if o == nil || IsNil(o.Accuracy) {
		var ret float32
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Synchronicity) GetAccuracyOk() (*float32, bool) {
	if o == nil || IsNil(o.Accuracy) {
		return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *Synchronicity) HasAccuracy() bool {
	if o != nil && !IsNil(o.Accuracy) {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given float32 and assigns it to the Accuracy field.
func (o *Synchronicity) SetAccuracy(v float32) {
	o.Accuracy = &v
}

func (o Synchronicity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Synchronicity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServAttrCom) {
		toSerialize["servAttrCom"] = o.ServAttrCom
	}
	if !IsNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !IsNil(o.Accuracy) {
		toSerialize["accuracy"] = o.Accuracy
	}
	return toSerialize, nil
}

type NullableSynchronicity struct {
	value *Synchronicity
	isSet bool
}

func (v NullableSynchronicity) Get() *Synchronicity {
	return v.value
}

func (v *NullableSynchronicity) Set(val *Synchronicity) {
	v.value = val
	v.isSet = true
}

func (v NullableSynchronicity) IsSet() bool {
	return v.isSet
}

func (v *NullableSynchronicity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSynchronicity(val *Synchronicity) *NullableSynchronicity {
	return &NullableSynchronicity{value: val, isSet: true}
}

func (v NullableSynchronicity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSynchronicity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EASDiscovery

import (
	"encoding/json"
	"time"
)

// checks if the DiscoveredEas type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveredEas{}

// DiscoveredEas Represents an EAS discovery information.
type DiscoveredEas struct {
	Eas EASProfile `json:"eas"`
	// string with format \"date-time\" as defined in OpenAPI.
	LifeTime *time.Time `json:"lifeTime,omitempty"`
}

// NewDiscoveredEas instantiates a new DiscoveredEas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveredEas(eas EASProfile) *DiscoveredEas {
	this := DiscoveredEas{}
	this.Eas = eas
	return &this
}

// NewDiscoveredEasWithDefaults instantiates a new DiscoveredEas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveredEasWithDefaults() *DiscoveredEas {
	this := DiscoveredEas{}
	return &this
}

// GetEas returns the Eas field value
func (o *DiscoveredEas) GetEas() EASProfile {
	if o == nil {
		var ret EASProfile
		return ret
	}

	return o.Eas
}

// GetEasOk returns a tuple with the Eas field value
// and a boolean to check if the value has been set.
func (o *DiscoveredEas) GetEasOk() (*EASProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Eas, true
}

// SetEas sets field value
func (o *DiscoveredEas) SetEas(v EASProfile) {
	o.Eas = v
}

// GetLifeTime returns the LifeTime field value if set, zero value otherwise.
func (o *DiscoveredEas) GetLifeTime() time.Time {
	if o == nil || isNil(o.LifeTime) {
		var ret time.Time
		return ret
	}
	return *o.LifeTime
}

// GetLifeTimeOk returns a tuple with the LifeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredEas) GetLifeTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LifeTime) {
		return nil, false
	}
	return o.LifeTime, true
}

// HasLifeTime returns a boolean if a field has been set.
func (o *DiscoveredEas) HasLifeTime() bool {
	if o != nil && !isNil(o.LifeTime) {
		return true
	}

	return false
}

// SetLifeTime gets a reference to the given time.Time and assigns it to the LifeTime field.
func (o *DiscoveredEas) SetLifeTime(v time.Time) {
	o.LifeTime = &v
}

func (o DiscoveredEas) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveredEas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eas"] = o.Eas
	if !isNil(o.LifeTime) {
		toSerialize["lifeTime"] = o.LifeTime
	}
	return toSerialize, nil
}

type NullableDiscoveredEas struct {
	value *DiscoveredEas
	isSet bool
}

func (v NullableDiscoveredEas) Get() *DiscoveredEas {
	return v.value
}

func (v *NullableDiscoveredEas) Set(val *DiscoveredEas) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveredEas) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveredEas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveredEas(val *DiscoveredEas) *NullableDiscoveredEas {
	return &NullableDiscoveredEas{value: val, isSet: true}
}

func (v NullableDiscoveredEas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveredEas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



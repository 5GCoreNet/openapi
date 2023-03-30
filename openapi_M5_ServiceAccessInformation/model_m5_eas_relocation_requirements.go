/*
M5_ServiceAccessInformation

5GMS AF M5 Service Access Information API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M5_ServiceAccessInformation

import (
	"encoding/json"
)

// checks if the M5EASRelocationRequirements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &M5EASRelocationRequirements{}

// M5EASRelocationRequirements Relocation requirements of an EAS.
type M5EASRelocationRequirements struct {
	Tolerance EASRelocationTolerance `json:"tolerance"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property. 
	MaxInterruptionDuration NullableInt32 `json:"maxInterruptionDuration,omitempty"`
}

// NewM5EASRelocationRequirements instantiates a new M5EASRelocationRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewM5EASRelocationRequirements(tolerance EASRelocationTolerance) *M5EASRelocationRequirements {
	this := M5EASRelocationRequirements{}
	this.Tolerance = tolerance
	return &this
}

// NewM5EASRelocationRequirementsWithDefaults instantiates a new M5EASRelocationRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewM5EASRelocationRequirementsWithDefaults() *M5EASRelocationRequirements {
	this := M5EASRelocationRequirements{}
	return &this
}

// GetTolerance returns the Tolerance field value
func (o *M5EASRelocationRequirements) GetTolerance() EASRelocationTolerance {
	if o == nil {
		var ret EASRelocationTolerance
		return ret
	}

	return o.Tolerance
}

// GetToleranceOk returns a tuple with the Tolerance field value
// and a boolean to check if the value has been set.
func (o *M5EASRelocationRequirements) GetToleranceOk() (*EASRelocationTolerance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tolerance, true
}

// SetTolerance sets field value
func (o *M5EASRelocationRequirements) SetTolerance(v EASRelocationTolerance) {
	o.Tolerance = v
}

// GetMaxInterruptionDuration returns the MaxInterruptionDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *M5EASRelocationRequirements) GetMaxInterruptionDuration() int32 {
	if o == nil || IsNil(o.MaxInterruptionDuration.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxInterruptionDuration.Get()
}

// GetMaxInterruptionDurationOk returns a tuple with the MaxInterruptionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *M5EASRelocationRequirements) GetMaxInterruptionDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxInterruptionDuration.Get(), o.MaxInterruptionDuration.IsSet()
}

// HasMaxInterruptionDuration returns a boolean if a field has been set.
func (o *M5EASRelocationRequirements) HasMaxInterruptionDuration() bool {
	if o != nil && o.MaxInterruptionDuration.IsSet() {
		return true
	}

	return false
}

// SetMaxInterruptionDuration gets a reference to the given NullableInt32 and assigns it to the MaxInterruptionDuration field.
func (o *M5EASRelocationRequirements) SetMaxInterruptionDuration(v int32) {
	o.MaxInterruptionDuration.Set(&v)
}
// SetMaxInterruptionDurationNil sets the value for MaxInterruptionDuration to be an explicit nil
func (o *M5EASRelocationRequirements) SetMaxInterruptionDurationNil() {
	o.MaxInterruptionDuration.Set(nil)
}

// UnsetMaxInterruptionDuration ensures that no value is present for MaxInterruptionDuration, not even an explicit nil
func (o *M5EASRelocationRequirements) UnsetMaxInterruptionDuration() {
	o.MaxInterruptionDuration.Unset()
}

func (o M5EASRelocationRequirements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o M5EASRelocationRequirements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tolerance"] = o.Tolerance
	if o.MaxInterruptionDuration.IsSet() {
		toSerialize["maxInterruptionDuration"] = o.MaxInterruptionDuration.Get()
	}
	return toSerialize, nil
}

type NullableM5EASRelocationRequirements struct {
	value *M5EASRelocationRequirements
	isSet bool
}

func (v NullableM5EASRelocationRequirements) Get() *M5EASRelocationRequirements {
	return v.value
}

func (v *NullableM5EASRelocationRequirements) Set(val *M5EASRelocationRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableM5EASRelocationRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableM5EASRelocationRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableM5EASRelocationRequirements(val *M5EASRelocationRequirements) *NullableM5EASRelocationRequirements {
	return &NullableM5EASRelocationRequirements{value: val, isSet: true}
}

func (v NullableM5EASRelocationRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableM5EASRelocationRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



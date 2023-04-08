/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the InternalGroupIdRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalGroupIdRange{}

// InternalGroupIdRange struct for InternalGroupIdRange
type InternalGroupIdRange struct {
	Start *string `json:"start,omitempty"`
	End *string `json:"end,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
}

// NewInternalGroupIdRange instantiates a new InternalGroupIdRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalGroupIdRange() *InternalGroupIdRange {
	this := InternalGroupIdRange{}
	return &this
}

// NewInternalGroupIdRangeWithDefaults instantiates a new InternalGroupIdRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalGroupIdRangeWithDefaults() *InternalGroupIdRange {
	this := InternalGroupIdRange{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *InternalGroupIdRange) GetStart() string {
	if o == nil || isNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalGroupIdRange) GetStartOk() (*string, bool) {
	if o == nil || isNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *InternalGroupIdRange) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *InternalGroupIdRange) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *InternalGroupIdRange) GetEnd() string {
	if o == nil || isNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalGroupIdRange) GetEndOk() (*string, bool) {
	if o == nil || isNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *InternalGroupIdRange) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *InternalGroupIdRange) SetEnd(v string) {
	o.End = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *InternalGroupIdRange) GetPattern() string {
	if o == nil || isNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalGroupIdRange) GetPatternOk() (*string, bool) {
	if o == nil || isNil(o.Pattern) {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *InternalGroupIdRange) HasPattern() bool {
	if o != nil && !isNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *InternalGroupIdRange) SetPattern(v string) {
	o.Pattern = &v
}

func (o InternalGroupIdRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalGroupIdRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !isNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	return toSerialize, nil
}

type NullableInternalGroupIdRange struct {
	value *InternalGroupIdRange
	isSet bool
}

func (v NullableInternalGroupIdRange) Get() *InternalGroupIdRange {
	return v.value
}

func (v *NullableInternalGroupIdRange) Set(val *InternalGroupIdRange) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalGroupIdRange) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalGroupIdRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalGroupIdRange(val *InternalGroupIdRange) *NullableInternalGroupIdRange {
	return &NullableInternalGroupIdRange{value: val, isSet: true}
}

func (v NullableInternalGroupIdRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalGroupIdRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



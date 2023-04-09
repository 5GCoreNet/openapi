/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_Events

import (
	"encoding/json"
)

// checks if the VALGroupFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VALGroupFilter{}

// VALGroupFilter Represents a filter of VAL group identifiers belonging to a VAL service.
type VALGroupFilter struct {
	// Identity of the VAL service
	ValSvcId *string `json:"valSvcId,omitempty"`
	// VAL group identifiers that event subscriber wants to know in the interested event.
	ValGrpIds []string `json:"valGrpIds"`
}

// NewVALGroupFilter instantiates a new VALGroupFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVALGroupFilter(valGrpIds []string) *VALGroupFilter {
	this := VALGroupFilter{}
	this.ValGrpIds = valGrpIds
	return &this
}

// NewVALGroupFilterWithDefaults instantiates a new VALGroupFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVALGroupFilterWithDefaults() *VALGroupFilter {
	this := VALGroupFilter{}
	return &this
}

// GetValSvcId returns the ValSvcId field value if set, zero value otherwise.
func (o *VALGroupFilter) GetValSvcId() string {
	if o == nil || IsNil(o.ValSvcId) {
		var ret string
		return ret
	}
	return *o.ValSvcId
}

// GetValSvcIdOk returns a tuple with the ValSvcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VALGroupFilter) GetValSvcIdOk() (*string, bool) {
	if o == nil || IsNil(o.ValSvcId) {
		return nil, false
	}
	return o.ValSvcId, true
}

// HasValSvcId returns a boolean if a field has been set.
func (o *VALGroupFilter) HasValSvcId() bool {
	if o != nil && !IsNil(o.ValSvcId) {
		return true
	}

	return false
}

// SetValSvcId gets a reference to the given string and assigns it to the ValSvcId field.
func (o *VALGroupFilter) SetValSvcId(v string) {
	o.ValSvcId = &v
}

// GetValGrpIds returns the ValGrpIds field value
func (o *VALGroupFilter) GetValGrpIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ValGrpIds
}

// GetValGrpIdsOk returns a tuple with the ValGrpIds field value
// and a boolean to check if the value has been set.
func (o *VALGroupFilter) GetValGrpIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValGrpIds, true
}

// SetValGrpIds sets field value
func (o *VALGroupFilter) SetValGrpIds(v []string) {
	o.ValGrpIds = v
}

func (o VALGroupFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VALGroupFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ValSvcId) {
		toSerialize["valSvcId"] = o.ValSvcId
	}
	toSerialize["valGrpIds"] = o.ValGrpIds
	return toSerialize, nil
}

type NullableVALGroupFilter struct {
	value *VALGroupFilter
	isSet bool
}

func (v NullableVALGroupFilter) Get() *VALGroupFilter {
	return v.value
}

func (v *NullableVALGroupFilter) Set(val *VALGroupFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableVALGroupFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableVALGroupFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVALGroupFilter(val *VALGroupFilter) *NullableVALGroupFilter {
	return &NullableVALGroupFilter{value: val, isSet: true}
}

func (v NullableVALGroupFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVALGroupFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

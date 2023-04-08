/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// checks if the UrllcEEPerfReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UrllcEEPerfReq{}

// UrllcEEPerfReq struct for UrllcEEPerfReq
type UrllcEEPerfReq struct {
	KpiType *string `json:"kpiType,omitempty"`
	Req *float32 `json:"req,omitempty"`
}

// NewUrllcEEPerfReq instantiates a new UrllcEEPerfReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrllcEEPerfReq() *UrllcEEPerfReq {
	this := UrllcEEPerfReq{}
	return &this
}

// NewUrllcEEPerfReqWithDefaults instantiates a new UrllcEEPerfReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrllcEEPerfReqWithDefaults() *UrllcEEPerfReq {
	this := UrllcEEPerfReq{}
	return &this
}

// GetKpiType returns the KpiType field value if set, zero value otherwise.
func (o *UrllcEEPerfReq) GetKpiType() string {
	if o == nil || isNil(o.KpiType) {
		var ret string
		return ret
	}
	return *o.KpiType
}

// GetKpiTypeOk returns a tuple with the KpiType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrllcEEPerfReq) GetKpiTypeOk() (*string, bool) {
	if o == nil || isNil(o.KpiType) {
		return nil, false
	}
	return o.KpiType, true
}

// HasKpiType returns a boolean if a field has been set.
func (o *UrllcEEPerfReq) HasKpiType() bool {
	if o != nil && !isNil(o.KpiType) {
		return true
	}

	return false
}

// SetKpiType gets a reference to the given string and assigns it to the KpiType field.
func (o *UrllcEEPerfReq) SetKpiType(v string) {
	o.KpiType = &v
}

// GetReq returns the Req field value if set, zero value otherwise.
func (o *UrllcEEPerfReq) GetReq() float32 {
	if o == nil || isNil(o.Req) {
		var ret float32
		return ret
	}
	return *o.Req
}

// GetReqOk returns a tuple with the Req field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrllcEEPerfReq) GetReqOk() (*float32, bool) {
	if o == nil || isNil(o.Req) {
		return nil, false
	}
	return o.Req, true
}

// HasReq returns a boolean if a field has been set.
func (o *UrllcEEPerfReq) HasReq() bool {
	if o != nil && !isNil(o.Req) {
		return true
	}

	return false
}

// SetReq gets a reference to the given float32 and assigns it to the Req field.
func (o *UrllcEEPerfReq) SetReq(v float32) {
	o.Req = &v
}

func (o UrllcEEPerfReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UrllcEEPerfReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.KpiType) {
		toSerialize["kpiType"] = o.KpiType
	}
	if !isNil(o.Req) {
		toSerialize["req"] = o.Req
	}
	return toSerialize, nil
}

type NullableUrllcEEPerfReq struct {
	value *UrllcEEPerfReq
	isSet bool
}

func (v NullableUrllcEEPerfReq) Get() *UrllcEEPerfReq {
	return v.value
}

func (v *NullableUrllcEEPerfReq) Set(val *UrllcEEPerfReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUrllcEEPerfReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUrllcEEPerfReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrllcEEPerfReq(val *UrllcEEPerfReq) *NullableUrllcEEPerfReq {
	return &NullableUrllcEEPerfReq{value: val, isSet: true}
}

func (v NullableUrllcEEPerfReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrllcEEPerfReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



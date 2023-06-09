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

// checks if the EmbbEEPerfReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbbEEPerfReq{}

// EmbbEEPerfReq struct for EmbbEEPerfReq
type EmbbEEPerfReq struct {
	KpiType *string  `json:"kpiType,omitempty"`
	Req     *float32 `json:"req,omitempty"`
}

// NewEmbbEEPerfReq instantiates a new EmbbEEPerfReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbbEEPerfReq() *EmbbEEPerfReq {
	this := EmbbEEPerfReq{}
	return &this
}

// NewEmbbEEPerfReqWithDefaults instantiates a new EmbbEEPerfReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbbEEPerfReqWithDefaults() *EmbbEEPerfReq {
	this := EmbbEEPerfReq{}
	return &this
}

// GetKpiType returns the KpiType field value if set, zero value otherwise.
func (o *EmbbEEPerfReq) GetKpiType() string {
	if o == nil || IsNil(o.KpiType) {
		var ret string
		return ret
	}
	return *o.KpiType
}

// GetKpiTypeOk returns a tuple with the KpiType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbbEEPerfReq) GetKpiTypeOk() (*string, bool) {
	if o == nil || IsNil(o.KpiType) {
		return nil, false
	}
	return o.KpiType, true
}

// HasKpiType returns a boolean if a field has been set.
func (o *EmbbEEPerfReq) HasKpiType() bool {
	if o != nil && !IsNil(o.KpiType) {
		return true
	}

	return false
}

// SetKpiType gets a reference to the given string and assigns it to the KpiType field.
func (o *EmbbEEPerfReq) SetKpiType(v string) {
	o.KpiType = &v
}

// GetReq returns the Req field value if set, zero value otherwise.
func (o *EmbbEEPerfReq) GetReq() float32 {
	if o == nil || IsNil(o.Req) {
		var ret float32
		return ret
	}
	return *o.Req
}

// GetReqOk returns a tuple with the Req field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbbEEPerfReq) GetReqOk() (*float32, bool) {
	if o == nil || IsNil(o.Req) {
		return nil, false
	}
	return o.Req, true
}

// HasReq returns a boolean if a field has been set.
func (o *EmbbEEPerfReq) HasReq() bool {
	if o != nil && !IsNil(o.Req) {
		return true
	}

	return false
}

// SetReq gets a reference to the given float32 and assigns it to the Req field.
func (o *EmbbEEPerfReq) SetReq(v float32) {
	o.Req = &v
}

func (o EmbbEEPerfReq) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbbEEPerfReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KpiType) {
		toSerialize["kpiType"] = o.KpiType
	}
	if !IsNil(o.Req) {
		toSerialize["req"] = o.Req
	}
	return toSerialize, nil
}

type NullableEmbbEEPerfReq struct {
	value *EmbbEEPerfReq
	isSet bool
}

func (v NullableEmbbEEPerfReq) Get() *EmbbEEPerfReq {
	return v.value
}

func (v *NullableEmbbEEPerfReq) Set(val *EmbbEEPerfReq) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbbEEPerfReq) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbbEEPerfReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbbEEPerfReq(val *EmbbEEPerfReq) *NullableEmbbEEPerfReq {
	return &NullableEmbbEEPerfReq{value: val, isSet: true}
}

func (v NullableEmbbEEPerfReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbbEEPerfReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the GnbDuFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GnbDuFunctionSingleAllOfAttributesAllOf{}

// GnbDuFunctionSingleAllOfAttributesAllOf struct for GnbDuFunctionSingleAllOfAttributesAllOf
type GnbDuFunctionSingleAllOfAttributesAllOf struct {
	GnbDuId         *float32         `json:"gnbDuId,omitempty"`
	GnbDuName       *string          `json:"gnbDuName,omitempty"`
	GnbId           *string          `json:"gnbId,omitempty"`
	GnbIdLength     *int32           `json:"gnbIdLength,omitempty"`
	RimRSReportConf *RimRSReportConf `json:"rimRSReportConf,omitempty"`
}

// NewGnbDuFunctionSingleAllOfAttributesAllOf instantiates a new GnbDuFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGnbDuFunctionSingleAllOfAttributesAllOf() *GnbDuFunctionSingleAllOfAttributesAllOf {
	this := GnbDuFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewGnbDuFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new GnbDuFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGnbDuFunctionSingleAllOfAttributesAllOfWithDefaults() *GnbDuFunctionSingleAllOfAttributesAllOf {
	this := GnbDuFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetGnbDuId returns the GnbDuId field value if set, zero value otherwise.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) GetGnbDuId() float32 {
	if o == nil || IsNil(o.GnbDuId) {
		var ret float32
		return ret
	}
	return *o.GnbDuId
}

// GetGnbDuIdOk returns a tuple with the GnbDuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) GetGnbDuIdOk() (*float32, bool) {
	if o == nil || IsNil(o.GnbDuId) {
		return nil, false
	}
	return o.GnbDuId, true
}

// HasGnbDuId returns a boolean if a field has been set.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) HasGnbDuId() bool {
	if o != nil && !IsNil(o.GnbDuId) {
		return true
	}

	return false
}

// SetGnbDuId gets a reference to the given float32 and assigns it to the GnbDuId field.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) SetGnbDuId(v float32) {
	o.GnbDuId = &v
}

// GetGnbDuName returns the GnbDuName field value if set, zero value otherwise.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) GetGnbDuName() string {
	if o == nil || IsNil(o.GnbDuName) {
		var ret string
		return ret
	}
	return *o.GnbDuName
}

// GetGnbDuNameOk returns a tuple with the GnbDuName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) GetGnbDuNameOk() (*string, bool) {
	if o == nil || IsNil(o.GnbDuName) {
		return nil, false
	}
	return o.GnbDuName, true
}

// HasGnbDuName returns a boolean if a field has been set.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) HasGnbDuName() bool {
	if o != nil && !IsNil(o.GnbDuName) {
		return true
	}

	return false
}

// SetGnbDuName gets a reference to the given string and assigns it to the GnbDuName field.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) SetGnbDuName(v string) {
	o.GnbDuName = &v
}

// GetGnbId returns the GnbId field value if set, zero value otherwise.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) GetGnbId() string {
	if o == nil || IsNil(o.GnbId) {
		var ret string
		return ret
	}
	return *o.GnbId
}

// GetGnbIdOk returns a tuple with the GnbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) GetGnbIdOk() (*string, bool) {
	if o == nil || IsNil(o.GnbId) {
		return nil, false
	}
	return o.GnbId, true
}

// HasGnbId returns a boolean if a field has been set.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) HasGnbId() bool {
	if o != nil && !IsNil(o.GnbId) {
		return true
	}

	return false
}

// SetGnbId gets a reference to the given string and assigns it to the GnbId field.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) SetGnbId(v string) {
	o.GnbId = &v
}

// GetGnbIdLength returns the GnbIdLength field value if set, zero value otherwise.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) GetGnbIdLength() int32 {
	if o == nil || IsNil(o.GnbIdLength) {
		var ret int32
		return ret
	}
	return *o.GnbIdLength
}

// GetGnbIdLengthOk returns a tuple with the GnbIdLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) GetGnbIdLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.GnbIdLength) {
		return nil, false
	}
	return o.GnbIdLength, true
}

// HasGnbIdLength returns a boolean if a field has been set.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) HasGnbIdLength() bool {
	if o != nil && !IsNil(o.GnbIdLength) {
		return true
	}

	return false
}

// SetGnbIdLength gets a reference to the given int32 and assigns it to the GnbIdLength field.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) SetGnbIdLength(v int32) {
	o.GnbIdLength = &v
}

// GetRimRSReportConf returns the RimRSReportConf field value if set, zero value otherwise.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) GetRimRSReportConf() RimRSReportConf {
	if o == nil || IsNil(o.RimRSReportConf) {
		var ret RimRSReportConf
		return ret
	}
	return *o.RimRSReportConf
}

// GetRimRSReportConfOk returns a tuple with the RimRSReportConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) GetRimRSReportConfOk() (*RimRSReportConf, bool) {
	if o == nil || IsNil(o.RimRSReportConf) {
		return nil, false
	}
	return o.RimRSReportConf, true
}

// HasRimRSReportConf returns a boolean if a field has been set.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) HasRimRSReportConf() bool {
	if o != nil && !IsNil(o.RimRSReportConf) {
		return true
	}

	return false
}

// SetRimRSReportConf gets a reference to the given RimRSReportConf and assigns it to the RimRSReportConf field.
func (o *GnbDuFunctionSingleAllOfAttributesAllOf) SetRimRSReportConf(v RimRSReportConf) {
	o.RimRSReportConf = &v
}

func (o GnbDuFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GnbDuFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GnbDuId) {
		toSerialize["gnbDuId"] = o.GnbDuId
	}
	if !IsNil(o.GnbDuName) {
		toSerialize["gnbDuName"] = o.GnbDuName
	}
	if !IsNil(o.GnbId) {
		toSerialize["gnbId"] = o.GnbId
	}
	if !IsNil(o.GnbIdLength) {
		toSerialize["gnbIdLength"] = o.GnbIdLength
	}
	if !IsNil(o.RimRSReportConf) {
		toSerialize["rimRSReportConf"] = o.RimRSReportConf
	}
	return toSerialize, nil
}

type NullableGnbDuFunctionSingleAllOfAttributesAllOf struct {
	value *GnbDuFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableGnbDuFunctionSingleAllOfAttributesAllOf) Get() *GnbDuFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableGnbDuFunctionSingleAllOfAttributesAllOf) Set(val *GnbDuFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGnbDuFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGnbDuFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGnbDuFunctionSingleAllOfAttributesAllOf(val *GnbDuFunctionSingleAllOfAttributesAllOf) *NullableGnbDuFunctionSingleAllOfAttributesAllOf {
	return &NullableGnbDuFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableGnbDuFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGnbDuFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

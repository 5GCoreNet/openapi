/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// checks if the RedundantTransmissionExpInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedundantTransmissionExpInfo{}

// RedundantTransmissionExpInfo The redundant transmission experience related information. When subscribed event is  \"RED_TRANS_EXP\", the \"redTransInfos\" attribute shall be included.
type RedundantTransmissionExpInfo struct {
	SpatialValidCon *NetworkAreaInfo `json:"spatialValidCon,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn          *string                         `json:"dnn,omitempty"`
	RedTransExps []RedundantTransmissionExpPerTS `json:"redTransExps"`
}

// NewRedundantTransmissionExpInfo instantiates a new RedundantTransmissionExpInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedundantTransmissionExpInfo(redTransExps []RedundantTransmissionExpPerTS) *RedundantTransmissionExpInfo {
	this := RedundantTransmissionExpInfo{}
	this.RedTransExps = redTransExps
	return &this
}

// NewRedundantTransmissionExpInfoWithDefaults instantiates a new RedundantTransmissionExpInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedundantTransmissionExpInfoWithDefaults() *RedundantTransmissionExpInfo {
	this := RedundantTransmissionExpInfo{}
	return &this
}

// GetSpatialValidCon returns the SpatialValidCon field value if set, zero value otherwise.
func (o *RedundantTransmissionExpInfo) GetSpatialValidCon() NetworkAreaInfo {
	if o == nil || IsNil(o.SpatialValidCon) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.SpatialValidCon
}

// GetSpatialValidConOk returns a tuple with the SpatialValidCon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundantTransmissionExpInfo) GetSpatialValidConOk() (*NetworkAreaInfo, bool) {
	if o == nil || IsNil(o.SpatialValidCon) {
		return nil, false
	}
	return o.SpatialValidCon, true
}

// HasSpatialValidCon returns a boolean if a field has been set.
func (o *RedundantTransmissionExpInfo) HasSpatialValidCon() bool {
	if o != nil && !IsNil(o.SpatialValidCon) {
		return true
	}

	return false
}

// SetSpatialValidCon gets a reference to the given NetworkAreaInfo and assigns it to the SpatialValidCon field.
func (o *RedundantTransmissionExpInfo) SetSpatialValidCon(v NetworkAreaInfo) {
	o.SpatialValidCon = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *RedundantTransmissionExpInfo) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundantTransmissionExpInfo) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *RedundantTransmissionExpInfo) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *RedundantTransmissionExpInfo) SetDnn(v string) {
	o.Dnn = &v
}

// GetRedTransExps returns the RedTransExps field value
func (o *RedundantTransmissionExpInfo) GetRedTransExps() []RedundantTransmissionExpPerTS {
	if o == nil {
		var ret []RedundantTransmissionExpPerTS
		return ret
	}

	return o.RedTransExps
}

// GetRedTransExpsOk returns a tuple with the RedTransExps field value
// and a boolean to check if the value has been set.
func (o *RedundantTransmissionExpInfo) GetRedTransExpsOk() ([]RedundantTransmissionExpPerTS, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedTransExps, true
}

// SetRedTransExps sets field value
func (o *RedundantTransmissionExpInfo) SetRedTransExps(v []RedundantTransmissionExpPerTS) {
	o.RedTransExps = v
}

func (o RedundantTransmissionExpInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedundantTransmissionExpInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpatialValidCon) {
		toSerialize["spatialValidCon"] = o.SpatialValidCon
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	toSerialize["redTransExps"] = o.RedTransExps
	return toSerialize, nil
}

type NullableRedundantTransmissionExpInfo struct {
	value *RedundantTransmissionExpInfo
	isSet bool
}

func (v NullableRedundantTransmissionExpInfo) Get() *RedundantTransmissionExpInfo {
	return v.value
}

func (v *NullableRedundantTransmissionExpInfo) Set(val *RedundantTransmissionExpInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRedundantTransmissionExpInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRedundantTransmissionExpInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedundantTransmissionExpInfo(val *RedundantTransmissionExpInfo) *NullableRedundantTransmissionExpInfo {
	return &NullableRedundantTransmissionExpInfo{value: val, isSet: true}
}

func (v NullableRedundantTransmissionExpInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedundantTransmissionExpInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

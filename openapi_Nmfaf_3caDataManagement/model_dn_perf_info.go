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

// checks if the DnPerfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnPerfInfo{}

// DnPerfInfo Represents DN performance information.
type DnPerfInfo struct {
	// String providing an application identifier.
	AppId *string `json:"appId,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	DnPerf []DnPerf `json:"dnPerf"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
}

// NewDnPerfInfo instantiates a new DnPerfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnPerfInfo(dnPerf []DnPerf) *DnPerfInfo {
	this := DnPerfInfo{}
	this.DnPerf = dnPerf
	return &this
}

// NewDnPerfInfoWithDefaults instantiates a new DnPerfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnPerfInfoWithDefaults() *DnPerfInfo {
	this := DnPerfInfo{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *DnPerfInfo) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnPerfInfo) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *DnPerfInfo) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *DnPerfInfo) SetAppId(v string) {
	o.AppId = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *DnPerfInfo) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnPerfInfo) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *DnPerfInfo) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *DnPerfInfo) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *DnPerfInfo) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnPerfInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *DnPerfInfo) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *DnPerfInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetDnPerf returns the DnPerf field value
func (o *DnPerfInfo) GetDnPerf() []DnPerf {
	if o == nil {
		var ret []DnPerf
		return ret
	}

	return o.DnPerf
}

// GetDnPerfOk returns a tuple with the DnPerf field value
// and a boolean to check if the value has been set.
func (o *DnPerfInfo) GetDnPerfOk() ([]DnPerf, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnPerf, true
}

// SetDnPerf sets field value
func (o *DnPerfInfo) SetDnPerf(v []DnPerf) {
	o.DnPerf = v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *DnPerfInfo) GetConfidence() int32 {
	if o == nil || isNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnPerfInfo) GetConfidenceOk() (*int32, bool) {
	if o == nil || isNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *DnPerfInfo) HasConfidence() bool {
	if o != nil && !isNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *DnPerfInfo) SetConfidence(v int32) {
	o.Confidence = &v
}

func (o DnPerfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnPerfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	toSerialize["dnPerf"] = o.DnPerf
	if !isNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	return toSerialize, nil
}

type NullableDnPerfInfo struct {
	value *DnPerfInfo
	isSet bool
}

func (v NullableDnPerfInfo) Get() *DnPerfInfo {
	return v.value
}

func (v *NullableDnPerfInfo) Set(val *DnPerfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDnPerfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDnPerfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnPerfInfo(val *DnPerfInfo) *NullableDnPerfInfo {
	return &NullableDnPerfInfo{value: val, isSet: true}
}

func (v NullableDnPerfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnPerfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



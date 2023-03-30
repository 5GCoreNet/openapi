/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// checks if the DatalinkReportingConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatalinkReportingConfiguration{}

// DatalinkReportingConfiguration struct for DatalinkReportingConfiguration
type DatalinkReportingConfiguration struct {
	DddTrafficDes []DddTrafficDescriptor `json:"dddTrafficDes,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Slice *Snssai `json:"slice,omitempty"`
	DddStatusList []DlDataDeliveryStatus `json:"dddStatusList,omitempty"`
}

// NewDatalinkReportingConfiguration instantiates a new DatalinkReportingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatalinkReportingConfiguration() *DatalinkReportingConfiguration {
	this := DatalinkReportingConfiguration{}
	return &this
}

// NewDatalinkReportingConfigurationWithDefaults instantiates a new DatalinkReportingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatalinkReportingConfigurationWithDefaults() *DatalinkReportingConfiguration {
	this := DatalinkReportingConfiguration{}
	return &this
}

// GetDddTrafficDes returns the DddTrafficDes field value if set, zero value otherwise.
func (o *DatalinkReportingConfiguration) GetDddTrafficDes() []DddTrafficDescriptor {
	if o == nil || IsNil(o.DddTrafficDes) {
		var ret []DddTrafficDescriptor
		return ret
	}
	return o.DddTrafficDes
}

// GetDddTrafficDesOk returns a tuple with the DddTrafficDes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatalinkReportingConfiguration) GetDddTrafficDesOk() ([]DddTrafficDescriptor, bool) {
	if o == nil || IsNil(o.DddTrafficDes) {
		return nil, false
	}
	return o.DddTrafficDes, true
}

// HasDddTrafficDes returns a boolean if a field has been set.
func (o *DatalinkReportingConfiguration) HasDddTrafficDes() bool {
	if o != nil && !IsNil(o.DddTrafficDes) {
		return true
	}

	return false
}

// SetDddTrafficDes gets a reference to the given []DddTrafficDescriptor and assigns it to the DddTrafficDes field.
func (o *DatalinkReportingConfiguration) SetDddTrafficDes(v []DddTrafficDescriptor) {
	o.DddTrafficDes = v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *DatalinkReportingConfiguration) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatalinkReportingConfiguration) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *DatalinkReportingConfiguration) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *DatalinkReportingConfiguration) SetDnn(v string) {
	o.Dnn = &v
}

// GetSlice returns the Slice field value if set, zero value otherwise.
func (o *DatalinkReportingConfiguration) GetSlice() Snssai {
	if o == nil || IsNil(o.Slice) {
		var ret Snssai
		return ret
	}
	return *o.Slice
}

// GetSliceOk returns a tuple with the Slice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatalinkReportingConfiguration) GetSliceOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Slice) {
		return nil, false
	}
	return o.Slice, true
}

// HasSlice returns a boolean if a field has been set.
func (o *DatalinkReportingConfiguration) HasSlice() bool {
	if o != nil && !IsNil(o.Slice) {
		return true
	}

	return false
}

// SetSlice gets a reference to the given Snssai and assigns it to the Slice field.
func (o *DatalinkReportingConfiguration) SetSlice(v Snssai) {
	o.Slice = &v
}

// GetDddStatusList returns the DddStatusList field value if set, zero value otherwise.
func (o *DatalinkReportingConfiguration) GetDddStatusList() []DlDataDeliveryStatus {
	if o == nil || IsNil(o.DddStatusList) {
		var ret []DlDataDeliveryStatus
		return ret
	}
	return o.DddStatusList
}

// GetDddStatusListOk returns a tuple with the DddStatusList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatalinkReportingConfiguration) GetDddStatusListOk() ([]DlDataDeliveryStatus, bool) {
	if o == nil || IsNil(o.DddStatusList) {
		return nil, false
	}
	return o.DddStatusList, true
}

// HasDddStatusList returns a boolean if a field has been set.
func (o *DatalinkReportingConfiguration) HasDddStatusList() bool {
	if o != nil && !IsNil(o.DddStatusList) {
		return true
	}

	return false
}

// SetDddStatusList gets a reference to the given []DlDataDeliveryStatus and assigns it to the DddStatusList field.
func (o *DatalinkReportingConfiguration) SetDddStatusList(v []DlDataDeliveryStatus) {
	o.DddStatusList = v
}

func (o DatalinkReportingConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatalinkReportingConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DddTrafficDes) {
		toSerialize["dddTrafficDes"] = o.DddTrafficDes
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.Slice) {
		toSerialize["slice"] = o.Slice
	}
	if !IsNil(o.DddStatusList) {
		toSerialize["dddStatusList"] = o.DddStatusList
	}
	return toSerialize, nil
}

type NullableDatalinkReportingConfiguration struct {
	value *DatalinkReportingConfiguration
	isSet bool
}

func (v NullableDatalinkReportingConfiguration) Get() *DatalinkReportingConfiguration {
	return v.value
}

func (v *NullableDatalinkReportingConfiguration) Set(val *DatalinkReportingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDatalinkReportingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDatalinkReportingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatalinkReportingConfiguration(val *DatalinkReportingConfiguration) *NullableDatalinkReportingConfiguration {
	return &NullableDatalinkReportingConfiguration{value: val, isSet: true}
}

func (v NullableDatalinkReportingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatalinkReportingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



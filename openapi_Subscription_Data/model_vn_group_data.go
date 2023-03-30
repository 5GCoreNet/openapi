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

// checks if the VnGroupData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnGroupData{}

// VnGroupData struct for VnGroupData
type VnGroupData struct {
	PduSessionTypes *PduSessionTypes `json:"pduSessionTypes,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	SingleNssai *Snssai `json:"singleNssai,omitempty"`
	AppDescriptors []AppDescriptor `json:"appDescriptors,omitempty"`
}

// NewVnGroupData instantiates a new VnGroupData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnGroupData() *VnGroupData {
	this := VnGroupData{}
	return &this
}

// NewVnGroupDataWithDefaults instantiates a new VnGroupData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnGroupDataWithDefaults() *VnGroupData {
	this := VnGroupData{}
	return &this
}

// GetPduSessionTypes returns the PduSessionTypes field value if set, zero value otherwise.
func (o *VnGroupData) GetPduSessionTypes() PduSessionTypes {
	if o == nil || IsNil(o.PduSessionTypes) {
		var ret PduSessionTypes
		return ret
	}
	return *o.PduSessionTypes
}

// GetPduSessionTypesOk returns a tuple with the PduSessionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnGroupData) GetPduSessionTypesOk() (*PduSessionTypes, bool) {
	if o == nil || IsNil(o.PduSessionTypes) {
		return nil, false
	}
	return o.PduSessionTypes, true
}

// HasPduSessionTypes returns a boolean if a field has been set.
func (o *VnGroupData) HasPduSessionTypes() bool {
	if o != nil && !IsNil(o.PduSessionTypes) {
		return true
	}

	return false
}

// SetPduSessionTypes gets a reference to the given PduSessionTypes and assigns it to the PduSessionTypes field.
func (o *VnGroupData) SetPduSessionTypes(v PduSessionTypes) {
	o.PduSessionTypes = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *VnGroupData) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnGroupData) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *VnGroupData) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *VnGroupData) SetDnn(v string) {
	o.Dnn = &v
}

// GetSingleNssai returns the SingleNssai field value if set, zero value otherwise.
func (o *VnGroupData) GetSingleNssai() Snssai {
	if o == nil || IsNil(o.SingleNssai) {
		var ret Snssai
		return ret
	}
	return *o.SingleNssai
}

// GetSingleNssaiOk returns a tuple with the SingleNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnGroupData) GetSingleNssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.SingleNssai) {
		return nil, false
	}
	return o.SingleNssai, true
}

// HasSingleNssai returns a boolean if a field has been set.
func (o *VnGroupData) HasSingleNssai() bool {
	if o != nil && !IsNil(o.SingleNssai) {
		return true
	}

	return false
}

// SetSingleNssai gets a reference to the given Snssai and assigns it to the SingleNssai field.
func (o *VnGroupData) SetSingleNssai(v Snssai) {
	o.SingleNssai = &v
}

// GetAppDescriptors returns the AppDescriptors field value if set, zero value otherwise.
func (o *VnGroupData) GetAppDescriptors() []AppDescriptor {
	if o == nil || IsNil(o.AppDescriptors) {
		var ret []AppDescriptor
		return ret
	}
	return o.AppDescriptors
}

// GetAppDescriptorsOk returns a tuple with the AppDescriptors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnGroupData) GetAppDescriptorsOk() ([]AppDescriptor, bool) {
	if o == nil || IsNil(o.AppDescriptors) {
		return nil, false
	}
	return o.AppDescriptors, true
}

// HasAppDescriptors returns a boolean if a field has been set.
func (o *VnGroupData) HasAppDescriptors() bool {
	if o != nil && !IsNil(o.AppDescriptors) {
		return true
	}

	return false
}

// SetAppDescriptors gets a reference to the given []AppDescriptor and assigns it to the AppDescriptors field.
func (o *VnGroupData) SetAppDescriptors(v []AppDescriptor) {
	o.AppDescriptors = v
}

func (o VnGroupData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnGroupData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PduSessionTypes) {
		toSerialize["pduSessionTypes"] = o.PduSessionTypes
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.SingleNssai) {
		toSerialize["singleNssai"] = o.SingleNssai
	}
	if !IsNil(o.AppDescriptors) {
		toSerialize["appDescriptors"] = o.AppDescriptors
	}
	return toSerialize, nil
}

type NullableVnGroupData struct {
	value *VnGroupData
	isSet bool
}

func (v NullableVnGroupData) Get() *VnGroupData {
	return v.value
}

func (v *NullableVnGroupData) Set(val *VnGroupData) {
	v.value = val
	v.isSet = true
}

func (v NullableVnGroupData) IsSet() bool {
	return v.isSet
}

func (v *NullableVnGroupData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnGroupData(val *VnGroupData) *NullableVnGroupData {
	return &NullableVnGroupData{value: val, isSet: true}
}

func (v NullableVnGroupData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnGroupData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



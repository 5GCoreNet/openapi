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

// checks if the NgeirFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NgeirFunctionSingleAllOfAttributesAllOf{}

// NgeirFunctionSingleAllOfAttributesAllOf struct for NgeirFunctionSingleAllOfAttributesAllOf
type NgeirFunctionSingleAllOfAttributesAllOf struct {
	PlmnIdList []PlmnId `json:"plmnIdList,omitempty"`
	SBIFqdn *string `json:"sBIFqdn,omitempty"`
	SnssaiList []Snssai `json:"snssaiList,omitempty"`
	ManagedNFProfile *ManagedNFProfile `json:"managedNFProfile,omitempty"`
	CommModelList []CommModel `json:"commModelList,omitempty"`
}

// NewNgeirFunctionSingleAllOfAttributesAllOf instantiates a new NgeirFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgeirFunctionSingleAllOfAttributesAllOf() *NgeirFunctionSingleAllOfAttributesAllOf {
	this := NgeirFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewNgeirFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new NgeirFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgeirFunctionSingleAllOfAttributesAllOfWithDefaults() *NgeirFunctionSingleAllOfAttributesAllOf {
	this := NgeirFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetPlmnIdList returns the PlmnIdList field value if set, zero value otherwise.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) GetPlmnIdList() []PlmnId {
	if o == nil || IsNil(o.PlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.PlmnIdList
}

// GetPlmnIdListOk returns a tuple with the PlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) GetPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.PlmnIdList) {
		return nil, false
	}
	return o.PlmnIdList, true
}

// HasPlmnIdList returns a boolean if a field has been set.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) HasPlmnIdList() bool {
	if o != nil && !IsNil(o.PlmnIdList) {
		return true
	}

	return false
}

// SetPlmnIdList gets a reference to the given []PlmnId and assigns it to the PlmnIdList field.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) SetPlmnIdList(v []PlmnId) {
	o.PlmnIdList = v
}

// GetSBIFqdn returns the SBIFqdn field value if set, zero value otherwise.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) GetSBIFqdn() string {
	if o == nil || IsNil(o.SBIFqdn) {
		var ret string
		return ret
	}
	return *o.SBIFqdn
}

// GetSBIFqdnOk returns a tuple with the SBIFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) GetSBIFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.SBIFqdn) {
		return nil, false
	}
	return o.SBIFqdn, true
}

// HasSBIFqdn returns a boolean if a field has been set.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) HasSBIFqdn() bool {
	if o != nil && !IsNil(o.SBIFqdn) {
		return true
	}

	return false
}

// SetSBIFqdn gets a reference to the given string and assigns it to the SBIFqdn field.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) SetSBIFqdn(v string) {
	o.SBIFqdn = &v
}

// GetSnssaiList returns the SnssaiList field value if set, zero value otherwise.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) GetSnssaiList() []Snssai {
	if o == nil || IsNil(o.SnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) GetSnssaiListOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.SnssaiList) {
		return nil, false
	}
	return o.SnssaiList, true
}

// HasSnssaiList returns a boolean if a field has been set.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) HasSnssaiList() bool {
	if o != nil && !IsNil(o.SnssaiList) {
		return true
	}

	return false
}

// SetSnssaiList gets a reference to the given []Snssai and assigns it to the SnssaiList field.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

// GetManagedNFProfile returns the ManagedNFProfile field value if set, zero value otherwise.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) GetManagedNFProfile() ManagedNFProfile {
	if o == nil || IsNil(o.ManagedNFProfile) {
		var ret ManagedNFProfile
		return ret
	}
	return *o.ManagedNFProfile
}

// GetManagedNFProfileOk returns a tuple with the ManagedNFProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) GetManagedNFProfileOk() (*ManagedNFProfile, bool) {
	if o == nil || IsNil(o.ManagedNFProfile) {
		return nil, false
	}
	return o.ManagedNFProfile, true
}

// HasManagedNFProfile returns a boolean if a field has been set.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) HasManagedNFProfile() bool {
	if o != nil && !IsNil(o.ManagedNFProfile) {
		return true
	}

	return false
}

// SetManagedNFProfile gets a reference to the given ManagedNFProfile and assigns it to the ManagedNFProfile field.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) SetManagedNFProfile(v ManagedNFProfile) {
	o.ManagedNFProfile = &v
}

// GetCommModelList returns the CommModelList field value if set, zero value otherwise.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) GetCommModelList() []CommModel {
	if o == nil || IsNil(o.CommModelList) {
		var ret []CommModel
		return ret
	}
	return o.CommModelList
}

// GetCommModelListOk returns a tuple with the CommModelList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) GetCommModelListOk() ([]CommModel, bool) {
	if o == nil || IsNil(o.CommModelList) {
		return nil, false
	}
	return o.CommModelList, true
}

// HasCommModelList returns a boolean if a field has been set.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) HasCommModelList() bool {
	if o != nil && !IsNil(o.CommModelList) {
		return true
	}

	return false
}

// SetCommModelList gets a reference to the given []CommModel and assigns it to the CommModelList field.
func (o *NgeirFunctionSingleAllOfAttributesAllOf) SetCommModelList(v []CommModel) {
	o.CommModelList = v
}

func (o NgeirFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NgeirFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlmnIdList) {
		toSerialize["plmnIdList"] = o.PlmnIdList
	}
	if !IsNil(o.SBIFqdn) {
		toSerialize["sBIFqdn"] = o.SBIFqdn
	}
	if !IsNil(o.SnssaiList) {
		toSerialize["snssaiList"] = o.SnssaiList
	}
	if !IsNil(o.ManagedNFProfile) {
		toSerialize["managedNFProfile"] = o.ManagedNFProfile
	}
	if !IsNil(o.CommModelList) {
		toSerialize["commModelList"] = o.CommModelList
	}
	return toSerialize, nil
}

type NullableNgeirFunctionSingleAllOfAttributesAllOf struct {
	value *NgeirFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableNgeirFunctionSingleAllOfAttributesAllOf) Get() *NgeirFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableNgeirFunctionSingleAllOfAttributesAllOf) Set(val *NgeirFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNgeirFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNgeirFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgeirFunctionSingleAllOfAttributesAllOf(val *NgeirFunctionSingleAllOfAttributesAllOf) *NullableNgeirFunctionSingleAllOfAttributesAllOf {
	return &NullableNgeirFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableNgeirFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgeirFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


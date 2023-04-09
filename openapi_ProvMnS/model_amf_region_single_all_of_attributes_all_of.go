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

// checks if the AmfRegionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfRegionSingleAllOfAttributesAllOf{}

// AmfRegionSingleAllOfAttributesAllOf struct for AmfRegionSingleAllOfAttributesAllOf
type AmfRegionSingleAllOfAttributesAllOf struct {
	PlmnIdList []PlmnId `json:"plmnIdList,omitempty"`
	NRTACList  []int32  `json:"nRTACList,omitempty"`
	// AmfRegionId is defined in TS 23.003
	AmfRegionId   *int32   `json:"amfRegionId,omitempty"`
	SnssaiList    []Snssai `json:"snssaiList,omitempty"`
	AMFSetListRef []string `json:"aMFSetListRef,omitempty"`
}

// NewAmfRegionSingleAllOfAttributesAllOf instantiates a new AmfRegionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfRegionSingleAllOfAttributesAllOf() *AmfRegionSingleAllOfAttributesAllOf {
	this := AmfRegionSingleAllOfAttributesAllOf{}
	return &this
}

// NewAmfRegionSingleAllOfAttributesAllOfWithDefaults instantiates a new AmfRegionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfRegionSingleAllOfAttributesAllOfWithDefaults() *AmfRegionSingleAllOfAttributesAllOf {
	this := AmfRegionSingleAllOfAttributesAllOf{}
	return &this
}

// GetPlmnIdList returns the PlmnIdList field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributesAllOf) GetPlmnIdList() []PlmnId {
	if o == nil || IsNil(o.PlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.PlmnIdList
}

// GetPlmnIdListOk returns a tuple with the PlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributesAllOf) GetPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.PlmnIdList) {
		return nil, false
	}
	return o.PlmnIdList, true
}

// HasPlmnIdList returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributesAllOf) HasPlmnIdList() bool {
	if o != nil && !IsNil(o.PlmnIdList) {
		return true
	}

	return false
}

// SetPlmnIdList gets a reference to the given []PlmnId and assigns it to the PlmnIdList field.
func (o *AmfRegionSingleAllOfAttributesAllOf) SetPlmnIdList(v []PlmnId) {
	o.PlmnIdList = v
}

// GetNRTACList returns the NRTACList field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributesAllOf) GetNRTACList() []int32 {
	if o == nil || IsNil(o.NRTACList) {
		var ret []int32
		return ret
	}
	return o.NRTACList
}

// GetNRTACListOk returns a tuple with the NRTACList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributesAllOf) GetNRTACListOk() ([]int32, bool) {
	if o == nil || IsNil(o.NRTACList) {
		return nil, false
	}
	return o.NRTACList, true
}

// HasNRTACList returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributesAllOf) HasNRTACList() bool {
	if o != nil && !IsNil(o.NRTACList) {
		return true
	}

	return false
}

// SetNRTACList gets a reference to the given []int32 and assigns it to the NRTACList field.
func (o *AmfRegionSingleAllOfAttributesAllOf) SetNRTACList(v []int32) {
	o.NRTACList = v
}

// GetAmfRegionId returns the AmfRegionId field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributesAllOf) GetAmfRegionId() int32 {
	if o == nil || IsNil(o.AmfRegionId) {
		var ret int32
		return ret
	}
	return *o.AmfRegionId
}

// GetAmfRegionIdOk returns a tuple with the AmfRegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributesAllOf) GetAmfRegionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AmfRegionId) {
		return nil, false
	}
	return o.AmfRegionId, true
}

// HasAmfRegionId returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributesAllOf) HasAmfRegionId() bool {
	if o != nil && !IsNil(o.AmfRegionId) {
		return true
	}

	return false
}

// SetAmfRegionId gets a reference to the given int32 and assigns it to the AmfRegionId field.
func (o *AmfRegionSingleAllOfAttributesAllOf) SetAmfRegionId(v int32) {
	o.AmfRegionId = &v
}

// GetSnssaiList returns the SnssaiList field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributesAllOf) GetSnssaiList() []Snssai {
	if o == nil || IsNil(o.SnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributesAllOf) GetSnssaiListOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.SnssaiList) {
		return nil, false
	}
	return o.SnssaiList, true
}

// HasSnssaiList returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributesAllOf) HasSnssaiList() bool {
	if o != nil && !IsNil(o.SnssaiList) {
		return true
	}

	return false
}

// SetSnssaiList gets a reference to the given []Snssai and assigns it to the SnssaiList field.
func (o *AmfRegionSingleAllOfAttributesAllOf) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

// GetAMFSetListRef returns the AMFSetListRef field value if set, zero value otherwise.
func (o *AmfRegionSingleAllOfAttributesAllOf) GetAMFSetListRef() []string {
	if o == nil || IsNil(o.AMFSetListRef) {
		var ret []string
		return ret
	}
	return o.AMFSetListRef
}

// GetAMFSetListRefOk returns a tuple with the AMFSetListRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfRegionSingleAllOfAttributesAllOf) GetAMFSetListRefOk() ([]string, bool) {
	if o == nil || IsNil(o.AMFSetListRef) {
		return nil, false
	}
	return o.AMFSetListRef, true
}

// HasAMFSetListRef returns a boolean if a field has been set.
func (o *AmfRegionSingleAllOfAttributesAllOf) HasAMFSetListRef() bool {
	if o != nil && !IsNil(o.AMFSetListRef) {
		return true
	}

	return false
}

// SetAMFSetListRef gets a reference to the given []string and assigns it to the AMFSetListRef field.
func (o *AmfRegionSingleAllOfAttributesAllOf) SetAMFSetListRef(v []string) {
	o.AMFSetListRef = v
}

func (o AmfRegionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfRegionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlmnIdList) {
		toSerialize["plmnIdList"] = o.PlmnIdList
	}
	if !IsNil(o.NRTACList) {
		toSerialize["nRTACList"] = o.NRTACList
	}
	if !IsNil(o.AmfRegionId) {
		toSerialize["amfRegionId"] = o.AmfRegionId
	}
	if !IsNil(o.SnssaiList) {
		toSerialize["snssaiList"] = o.SnssaiList
	}
	if !IsNil(o.AMFSetListRef) {
		toSerialize["aMFSetListRef"] = o.AMFSetListRef
	}
	return toSerialize, nil
}

type NullableAmfRegionSingleAllOfAttributesAllOf struct {
	value *AmfRegionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableAmfRegionSingleAllOfAttributesAllOf) Get() *AmfRegionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableAmfRegionSingleAllOfAttributesAllOf) Set(val *AmfRegionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfRegionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfRegionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfRegionSingleAllOfAttributesAllOf(val *AmfRegionSingleAllOfAttributesAllOf) *NullableAmfRegionSingleAllOfAttributesAllOf {
	return &NullableAmfRegionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableAmfRegionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfRegionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

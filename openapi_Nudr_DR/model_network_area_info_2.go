/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the NetworkAreaInfo2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkAreaInfo2{}

// NetworkAreaInfo2 Describes a network area information in which the NF service consumer requests the number of UEs. 
type NetworkAreaInfo2 struct {
	// Contains a list of E-UTRA cell identities.
	Ecgis []Ecgi1 `json:"ecgis,omitempty"`
	// Contains a list of NR cell identities.
	Ncgis []Ncgi1 `json:"ncgis,omitempty"`
	// Contains a list of NG RAN nodes.
	GRanNodeIds []GlobalRanNodeId1 `json:"gRanNodeIds,omitempty"`
	// Contains a list of tracking area identities.
	Tais []Tai1 `json:"tais,omitempty"`
}

// NewNetworkAreaInfo2 instantiates a new NetworkAreaInfo2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAreaInfo2() *NetworkAreaInfo2 {
	this := NetworkAreaInfo2{}
	return &this
}

// NewNetworkAreaInfo2WithDefaults instantiates a new NetworkAreaInfo2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAreaInfo2WithDefaults() *NetworkAreaInfo2 {
	this := NetworkAreaInfo2{}
	return &this
}

// GetEcgis returns the Ecgis field value if set, zero value otherwise.
func (o *NetworkAreaInfo2) GetEcgis() []Ecgi1 {
	if o == nil || isNil(o.Ecgis) {
		var ret []Ecgi1
		return ret
	}
	return o.Ecgis
}

// GetEcgisOk returns a tuple with the Ecgis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAreaInfo2) GetEcgisOk() ([]Ecgi1, bool) {
	if o == nil || isNil(o.Ecgis) {
		return nil, false
	}
	return o.Ecgis, true
}

// HasEcgis returns a boolean if a field has been set.
func (o *NetworkAreaInfo2) HasEcgis() bool {
	if o != nil && !isNil(o.Ecgis) {
		return true
	}

	return false
}

// SetEcgis gets a reference to the given []Ecgi1 and assigns it to the Ecgis field.
func (o *NetworkAreaInfo2) SetEcgis(v []Ecgi1) {
	o.Ecgis = v
}

// GetNcgis returns the Ncgis field value if set, zero value otherwise.
func (o *NetworkAreaInfo2) GetNcgis() []Ncgi1 {
	if o == nil || isNil(o.Ncgis) {
		var ret []Ncgi1
		return ret
	}
	return o.Ncgis
}

// GetNcgisOk returns a tuple with the Ncgis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAreaInfo2) GetNcgisOk() ([]Ncgi1, bool) {
	if o == nil || isNil(o.Ncgis) {
		return nil, false
	}
	return o.Ncgis, true
}

// HasNcgis returns a boolean if a field has been set.
func (o *NetworkAreaInfo2) HasNcgis() bool {
	if o != nil && !isNil(o.Ncgis) {
		return true
	}

	return false
}

// SetNcgis gets a reference to the given []Ncgi1 and assigns it to the Ncgis field.
func (o *NetworkAreaInfo2) SetNcgis(v []Ncgi1) {
	o.Ncgis = v
}

// GetGRanNodeIds returns the GRanNodeIds field value if set, zero value otherwise.
func (o *NetworkAreaInfo2) GetGRanNodeIds() []GlobalRanNodeId1 {
	if o == nil || isNil(o.GRanNodeIds) {
		var ret []GlobalRanNodeId1
		return ret
	}
	return o.GRanNodeIds
}

// GetGRanNodeIdsOk returns a tuple with the GRanNodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAreaInfo2) GetGRanNodeIdsOk() ([]GlobalRanNodeId1, bool) {
	if o == nil || isNil(o.GRanNodeIds) {
		return nil, false
	}
	return o.GRanNodeIds, true
}

// HasGRanNodeIds returns a boolean if a field has been set.
func (o *NetworkAreaInfo2) HasGRanNodeIds() bool {
	if o != nil && !isNil(o.GRanNodeIds) {
		return true
	}

	return false
}

// SetGRanNodeIds gets a reference to the given []GlobalRanNodeId1 and assigns it to the GRanNodeIds field.
func (o *NetworkAreaInfo2) SetGRanNodeIds(v []GlobalRanNodeId1) {
	o.GRanNodeIds = v
}

// GetTais returns the Tais field value if set, zero value otherwise.
func (o *NetworkAreaInfo2) GetTais() []Tai1 {
	if o == nil || isNil(o.Tais) {
		var ret []Tai1
		return ret
	}
	return o.Tais
}

// GetTaisOk returns a tuple with the Tais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAreaInfo2) GetTaisOk() ([]Tai1, bool) {
	if o == nil || isNil(o.Tais) {
		return nil, false
	}
	return o.Tais, true
}

// HasTais returns a boolean if a field has been set.
func (o *NetworkAreaInfo2) HasTais() bool {
	if o != nil && !isNil(o.Tais) {
		return true
	}

	return false
}

// SetTais gets a reference to the given []Tai1 and assigns it to the Tais field.
func (o *NetworkAreaInfo2) SetTais(v []Tai1) {
	o.Tais = v
}

func (o NetworkAreaInfo2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkAreaInfo2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ecgis) {
		toSerialize["ecgis"] = o.Ecgis
	}
	if !isNil(o.Ncgis) {
		toSerialize["ncgis"] = o.Ncgis
	}
	if !isNil(o.GRanNodeIds) {
		toSerialize["gRanNodeIds"] = o.GRanNodeIds
	}
	if !isNil(o.Tais) {
		toSerialize["tais"] = o.Tais
	}
	return toSerialize, nil
}

type NullableNetworkAreaInfo2 struct {
	value *NetworkAreaInfo2
	isSet bool
}

func (v NullableNetworkAreaInfo2) Get() *NetworkAreaInfo2 {
	return v.value
}

func (v *NullableNetworkAreaInfo2) Set(val *NetworkAreaInfo2) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAreaInfo2) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAreaInfo2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAreaInfo2(val *NetworkAreaInfo2) *NullableNetworkAreaInfo2 {
	return &NullableNetworkAreaInfo2{value: val, isSet: true}
}

func (v NullableNetworkAreaInfo2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAreaInfo2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



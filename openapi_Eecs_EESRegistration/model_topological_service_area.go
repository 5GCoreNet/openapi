/*
ECS EES Registration_API

API for EES Registration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eecs_EESRegistration

import (
	"encoding/json"
)

// checks if the TopologicalServiceArea type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TopologicalServiceArea{}

// TopologicalServiceArea Represents topological service area information.
type TopologicalServiceArea struct {
	// A list of E-UTRA cell identities.
	Ecgis []Ecgi `json:"ecgis,omitempty"`
	// A list of NR cell identities.
	Ncgis []Ncgi `json:"ncgis,omitempty"`
	// A list of tracking area identities.
	Tais []Tai `json:"tais,omitempty"`
	// A list of PLMN identities.
	PlmnIds []PlmnId1 `json:"plmnIds,omitempty"`
}

// NewTopologicalServiceArea instantiates a new TopologicalServiceArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologicalServiceArea() *TopologicalServiceArea {
	this := TopologicalServiceArea{}
	return &this
}

// NewTopologicalServiceAreaWithDefaults instantiates a new TopologicalServiceArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologicalServiceAreaWithDefaults() *TopologicalServiceArea {
	this := TopologicalServiceArea{}
	return &this
}

// GetEcgis returns the Ecgis field value if set, zero value otherwise.
func (o *TopologicalServiceArea) GetEcgis() []Ecgi {
	if o == nil || IsNil(o.Ecgis) {
		var ret []Ecgi
		return ret
	}
	return o.Ecgis
}

// GetEcgisOk returns a tuple with the Ecgis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologicalServiceArea) GetEcgisOk() ([]Ecgi, bool) {
	if o == nil || IsNil(o.Ecgis) {
		return nil, false
	}
	return o.Ecgis, true
}

// HasEcgis returns a boolean if a field has been set.
func (o *TopologicalServiceArea) HasEcgis() bool {
	if o != nil && !IsNil(o.Ecgis) {
		return true
	}

	return false
}

// SetEcgis gets a reference to the given []Ecgi and assigns it to the Ecgis field.
func (o *TopologicalServiceArea) SetEcgis(v []Ecgi) {
	o.Ecgis = v
}

// GetNcgis returns the Ncgis field value if set, zero value otherwise.
func (o *TopologicalServiceArea) GetNcgis() []Ncgi {
	if o == nil || IsNil(o.Ncgis) {
		var ret []Ncgi
		return ret
	}
	return o.Ncgis
}

// GetNcgisOk returns a tuple with the Ncgis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologicalServiceArea) GetNcgisOk() ([]Ncgi, bool) {
	if o == nil || IsNil(o.Ncgis) {
		return nil, false
	}
	return o.Ncgis, true
}

// HasNcgis returns a boolean if a field has been set.
func (o *TopologicalServiceArea) HasNcgis() bool {
	if o != nil && !IsNil(o.Ncgis) {
		return true
	}

	return false
}

// SetNcgis gets a reference to the given []Ncgi and assigns it to the Ncgis field.
func (o *TopologicalServiceArea) SetNcgis(v []Ncgi) {
	o.Ncgis = v
}

// GetTais returns the Tais field value if set, zero value otherwise.
func (o *TopologicalServiceArea) GetTais() []Tai {
	if o == nil || IsNil(o.Tais) {
		var ret []Tai
		return ret
	}
	return o.Tais
}

// GetTaisOk returns a tuple with the Tais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologicalServiceArea) GetTaisOk() ([]Tai, bool) {
	if o == nil || IsNil(o.Tais) {
		return nil, false
	}
	return o.Tais, true
}

// HasTais returns a boolean if a field has been set.
func (o *TopologicalServiceArea) HasTais() bool {
	if o != nil && !IsNil(o.Tais) {
		return true
	}

	return false
}

// SetTais gets a reference to the given []Tai and assigns it to the Tais field.
func (o *TopologicalServiceArea) SetTais(v []Tai) {
	o.Tais = v
}

// GetPlmnIds returns the PlmnIds field value if set, zero value otherwise.
func (o *TopologicalServiceArea) GetPlmnIds() []PlmnId1 {
	if o == nil || IsNil(o.PlmnIds) {
		var ret []PlmnId1
		return ret
	}
	return o.PlmnIds
}

// GetPlmnIdsOk returns a tuple with the PlmnIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologicalServiceArea) GetPlmnIdsOk() ([]PlmnId1, bool) {
	if o == nil || IsNil(o.PlmnIds) {
		return nil, false
	}
	return o.PlmnIds, true
}

// HasPlmnIds returns a boolean if a field has been set.
func (o *TopologicalServiceArea) HasPlmnIds() bool {
	if o != nil && !IsNil(o.PlmnIds) {
		return true
	}

	return false
}

// SetPlmnIds gets a reference to the given []PlmnId1 and assigns it to the PlmnIds field.
func (o *TopologicalServiceArea) SetPlmnIds(v []PlmnId1) {
	o.PlmnIds = v
}

func (o TopologicalServiceArea) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopologicalServiceArea) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ecgis) {
		toSerialize["ecgis"] = o.Ecgis
	}
	if !IsNil(o.Ncgis) {
		toSerialize["ncgis"] = o.Ncgis
	}
	if !IsNil(o.Tais) {
		toSerialize["tais"] = o.Tais
	}
	if !IsNil(o.PlmnIds) {
		toSerialize["plmnIds"] = o.PlmnIds
	}
	return toSerialize, nil
}

type NullableTopologicalServiceArea struct {
	value *TopologicalServiceArea
	isSet bool
}

func (v NullableTopologicalServiceArea) Get() *TopologicalServiceArea {
	return v.value
}

func (v *NullableTopologicalServiceArea) Set(val *TopologicalServiceArea) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologicalServiceArea) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologicalServiceArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologicalServiceArea(val *TopologicalServiceArea) *NullableTopologicalServiceArea {
	return &NullableTopologicalServiceArea{value: val, isSet: true}
}

func (v NullableTopologicalServiceArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologicalServiceArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



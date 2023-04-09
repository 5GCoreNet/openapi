/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Application_Data

import (
	"encoding/json"
)

// checks if the RouteSelectionParameterSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteSelectionParameterSet{}

// RouteSelectionParameterSet Contains parameters that can be used to guide the Route Selection Descriptors of the URSP.
type RouteSelectionParameterSet struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn    *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Precedence *int32 `json:"precedence,omitempty"`
	// Indicates where the route selection parameters apply. It may correspond to a geographical area, for example using a geographic shape that is known to the AF and is configured by the operator to correspond to a list of or TAIs.
	SpatialValidityAreas []GeographicalArea `json:"spatialValidityAreas,omitempty"`
	// Indicates the TAIs in which the route selection parameters apply. This attribute is  applicable only within the 5GC and it shall not be included in the request messages of  untrusted AFs for URSP guidance.
	SpatialValidityTais []Tai `json:"spatialValidityTais,omitempty"`
}

// NewRouteSelectionParameterSet instantiates a new RouteSelectionParameterSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteSelectionParameterSet() *RouteSelectionParameterSet {
	this := RouteSelectionParameterSet{}
	return &this
}

// NewRouteSelectionParameterSetWithDefaults instantiates a new RouteSelectionParameterSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteSelectionParameterSetWithDefaults() *RouteSelectionParameterSet {
	this := RouteSelectionParameterSet{}
	return &this
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *RouteSelectionParameterSet) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteSelectionParameterSet) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *RouteSelectionParameterSet) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *RouteSelectionParameterSet) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *RouteSelectionParameterSet) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteSelectionParameterSet) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *RouteSelectionParameterSet) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *RouteSelectionParameterSet) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetPrecedence returns the Precedence field value if set, zero value otherwise.
func (o *RouteSelectionParameterSet) GetPrecedence() int32 {
	if o == nil || IsNil(o.Precedence) {
		var ret int32
		return ret
	}
	return *o.Precedence
}

// GetPrecedenceOk returns a tuple with the Precedence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteSelectionParameterSet) GetPrecedenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Precedence) {
		return nil, false
	}
	return o.Precedence, true
}

// HasPrecedence returns a boolean if a field has been set.
func (o *RouteSelectionParameterSet) HasPrecedence() bool {
	if o != nil && !IsNil(o.Precedence) {
		return true
	}

	return false
}

// SetPrecedence gets a reference to the given int32 and assigns it to the Precedence field.
func (o *RouteSelectionParameterSet) SetPrecedence(v int32) {
	o.Precedence = &v
}

// GetSpatialValidityAreas returns the SpatialValidityAreas field value if set, zero value otherwise.
func (o *RouteSelectionParameterSet) GetSpatialValidityAreas() []GeographicalArea {
	if o == nil || IsNil(o.SpatialValidityAreas) {
		var ret []GeographicalArea
		return ret
	}
	return o.SpatialValidityAreas
}

// GetSpatialValidityAreasOk returns a tuple with the SpatialValidityAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteSelectionParameterSet) GetSpatialValidityAreasOk() ([]GeographicalArea, bool) {
	if o == nil || IsNil(o.SpatialValidityAreas) {
		return nil, false
	}
	return o.SpatialValidityAreas, true
}

// HasSpatialValidityAreas returns a boolean if a field has been set.
func (o *RouteSelectionParameterSet) HasSpatialValidityAreas() bool {
	if o != nil && !IsNil(o.SpatialValidityAreas) {
		return true
	}

	return false
}

// SetSpatialValidityAreas gets a reference to the given []GeographicalArea and assigns it to the SpatialValidityAreas field.
func (o *RouteSelectionParameterSet) SetSpatialValidityAreas(v []GeographicalArea) {
	o.SpatialValidityAreas = v
}

// GetSpatialValidityTais returns the SpatialValidityTais field value if set, zero value otherwise.
func (o *RouteSelectionParameterSet) GetSpatialValidityTais() []Tai {
	if o == nil || IsNil(o.SpatialValidityTais) {
		var ret []Tai
		return ret
	}
	return o.SpatialValidityTais
}

// GetSpatialValidityTaisOk returns a tuple with the SpatialValidityTais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteSelectionParameterSet) GetSpatialValidityTaisOk() ([]Tai, bool) {
	if o == nil || IsNil(o.SpatialValidityTais) {
		return nil, false
	}
	return o.SpatialValidityTais, true
}

// HasSpatialValidityTais returns a boolean if a field has been set.
func (o *RouteSelectionParameterSet) HasSpatialValidityTais() bool {
	if o != nil && !IsNil(o.SpatialValidityTais) {
		return true
	}

	return false
}

// SetSpatialValidityTais gets a reference to the given []Tai and assigns it to the SpatialValidityTais field.
func (o *RouteSelectionParameterSet) SetSpatialValidityTais(v []Tai) {
	o.SpatialValidityTais = v
}

func (o RouteSelectionParameterSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteSelectionParameterSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !IsNil(o.Precedence) {
		toSerialize["precedence"] = o.Precedence
	}
	if !IsNil(o.SpatialValidityAreas) {
		toSerialize["spatialValidityAreas"] = o.SpatialValidityAreas
	}
	if !IsNil(o.SpatialValidityTais) {
		toSerialize["spatialValidityTais"] = o.SpatialValidityTais
	}
	return toSerialize, nil
}

type NullableRouteSelectionParameterSet struct {
	value *RouteSelectionParameterSet
	isSet bool
}

func (v NullableRouteSelectionParameterSet) Get() *RouteSelectionParameterSet {
	return v.value
}

func (v *NullableRouteSelectionParameterSet) Set(val *RouteSelectionParameterSet) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteSelectionParameterSet) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteSelectionParameterSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteSelectionParameterSet(val *RouteSelectionParameterSet) *NullableRouteSelectionParameterSet {
	return &NullableRouteSelectionParameterSet{value: val, isSet: true}
}

func (v NullableRouteSelectionParameterSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteSelectionParameterSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

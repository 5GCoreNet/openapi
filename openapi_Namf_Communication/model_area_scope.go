/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the AreaScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AreaScope{}

// AreaScope Contain the area based on Cells or Tracking Areas.
type AreaScope struct {
	EutraCellIdList []string `json:"eutraCellIdList,omitempty"`
	NrCellIdList []string `json:"nrCellIdList,omitempty"`
	TacList []string `json:"tacList,omitempty"`
	// A map (list of key-value pairs) where PlmnId converted to a string serves as key 
	TacInfoPerPlmn *map[string]TacInfo `json:"tacInfoPerPlmn,omitempty"`
}

// NewAreaScope instantiates a new AreaScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAreaScope() *AreaScope {
	this := AreaScope{}
	return &this
}

// NewAreaScopeWithDefaults instantiates a new AreaScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAreaScopeWithDefaults() *AreaScope {
	this := AreaScope{}
	return &this
}

// GetEutraCellIdList returns the EutraCellIdList field value if set, zero value otherwise.
func (o *AreaScope) GetEutraCellIdList() []string {
	if o == nil || IsNil(o.EutraCellIdList) {
		var ret []string
		return ret
	}
	return o.EutraCellIdList
}

// GetEutraCellIdListOk returns a tuple with the EutraCellIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaScope) GetEutraCellIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.EutraCellIdList) {
		return nil, false
	}
	return o.EutraCellIdList, true
}

// HasEutraCellIdList returns a boolean if a field has been set.
func (o *AreaScope) HasEutraCellIdList() bool {
	if o != nil && !IsNil(o.EutraCellIdList) {
		return true
	}

	return false
}

// SetEutraCellIdList gets a reference to the given []string and assigns it to the EutraCellIdList field.
func (o *AreaScope) SetEutraCellIdList(v []string) {
	o.EutraCellIdList = v
}

// GetNrCellIdList returns the NrCellIdList field value if set, zero value otherwise.
func (o *AreaScope) GetNrCellIdList() []string {
	if o == nil || IsNil(o.NrCellIdList) {
		var ret []string
		return ret
	}
	return o.NrCellIdList
}

// GetNrCellIdListOk returns a tuple with the NrCellIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaScope) GetNrCellIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.NrCellIdList) {
		return nil, false
	}
	return o.NrCellIdList, true
}

// HasNrCellIdList returns a boolean if a field has been set.
func (o *AreaScope) HasNrCellIdList() bool {
	if o != nil && !IsNil(o.NrCellIdList) {
		return true
	}

	return false
}

// SetNrCellIdList gets a reference to the given []string and assigns it to the NrCellIdList field.
func (o *AreaScope) SetNrCellIdList(v []string) {
	o.NrCellIdList = v
}

// GetTacList returns the TacList field value if set, zero value otherwise.
func (o *AreaScope) GetTacList() []string {
	if o == nil || IsNil(o.TacList) {
		var ret []string
		return ret
	}
	return o.TacList
}

// GetTacListOk returns a tuple with the TacList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaScope) GetTacListOk() ([]string, bool) {
	if o == nil || IsNil(o.TacList) {
		return nil, false
	}
	return o.TacList, true
}

// HasTacList returns a boolean if a field has been set.
func (o *AreaScope) HasTacList() bool {
	if o != nil && !IsNil(o.TacList) {
		return true
	}

	return false
}

// SetTacList gets a reference to the given []string and assigns it to the TacList field.
func (o *AreaScope) SetTacList(v []string) {
	o.TacList = v
}

// GetTacInfoPerPlmn returns the TacInfoPerPlmn field value if set, zero value otherwise.
func (o *AreaScope) GetTacInfoPerPlmn() map[string]TacInfo {
	if o == nil || IsNil(o.TacInfoPerPlmn) {
		var ret map[string]TacInfo
		return ret
	}
	return *o.TacInfoPerPlmn
}

// GetTacInfoPerPlmnOk returns a tuple with the TacInfoPerPlmn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaScope) GetTacInfoPerPlmnOk() (*map[string]TacInfo, bool) {
	if o == nil || IsNil(o.TacInfoPerPlmn) {
		return nil, false
	}
	return o.TacInfoPerPlmn, true
}

// HasTacInfoPerPlmn returns a boolean if a field has been set.
func (o *AreaScope) HasTacInfoPerPlmn() bool {
	if o != nil && !IsNil(o.TacInfoPerPlmn) {
		return true
	}

	return false
}

// SetTacInfoPerPlmn gets a reference to the given map[string]TacInfo and assigns it to the TacInfoPerPlmn field.
func (o *AreaScope) SetTacInfoPerPlmn(v map[string]TacInfo) {
	o.TacInfoPerPlmn = &v
}

func (o AreaScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AreaScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EutraCellIdList) {
		toSerialize["eutraCellIdList"] = o.EutraCellIdList
	}
	if !IsNil(o.NrCellIdList) {
		toSerialize["nrCellIdList"] = o.NrCellIdList
	}
	if !IsNil(o.TacList) {
		toSerialize["tacList"] = o.TacList
	}
	if !IsNil(o.TacInfoPerPlmn) {
		toSerialize["tacInfoPerPlmn"] = o.TacInfoPerPlmn
	}
	return toSerialize, nil
}

type NullableAreaScope struct {
	value *AreaScope
	isSet bool
}

func (v NullableAreaScope) Get() *AreaScope {
	return v.value
}

func (v *NullableAreaScope) Set(val *AreaScope) {
	v.value = val
	v.isSet = true
}

func (v NullableAreaScope) IsSet() bool {
	return v.isSet
}

func (v *NullableAreaScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAreaScope(val *AreaScope) *NullableAreaScope {
	return &NullableAreaScope{value: val, isSet: true}
}

func (v NullableAreaScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAreaScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



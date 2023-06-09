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

// checks if the LmfFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LmfFunctionSingleAllOfAttributes{}

// LmfFunctionSingleAllOfAttributes struct for LmfFunctionSingleAllOfAttributes
type LmfFunctionSingleAllOfAttributes struct {
	ManagedFunctionAttr
	PlmnIdList       []PlmnId          `json:"plmnIdList,omitempty"`
	ManagedNFProfile *ManagedNFProfile `json:"managedNFProfile,omitempty"`
	CommModelList    []CommModel       `json:"commModelList,omitempty"`
}

// NewLmfFunctionSingleAllOfAttributes instantiates a new LmfFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLmfFunctionSingleAllOfAttributes() *LmfFunctionSingleAllOfAttributes {
	this := LmfFunctionSingleAllOfAttributes{}
	return &this
}

// NewLmfFunctionSingleAllOfAttributesWithDefaults instantiates a new LmfFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLmfFunctionSingleAllOfAttributesWithDefaults() *LmfFunctionSingleAllOfAttributes {
	this := LmfFunctionSingleAllOfAttributes{}
	return &this
}

// GetPlmnIdList returns the PlmnIdList field value if set, zero value otherwise.
func (o *LmfFunctionSingleAllOfAttributes) GetPlmnIdList() []PlmnId {
	if o == nil || IsNil(o.PlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.PlmnIdList
}

// GetPlmnIdListOk returns a tuple with the PlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfFunctionSingleAllOfAttributes) GetPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.PlmnIdList) {
		return nil, false
	}
	return o.PlmnIdList, true
}

// HasPlmnIdList returns a boolean if a field has been set.
func (o *LmfFunctionSingleAllOfAttributes) HasPlmnIdList() bool {
	if o != nil && !IsNil(o.PlmnIdList) {
		return true
	}

	return false
}

// SetPlmnIdList gets a reference to the given []PlmnId and assigns it to the PlmnIdList field.
func (o *LmfFunctionSingleAllOfAttributes) SetPlmnIdList(v []PlmnId) {
	o.PlmnIdList = v
}

// GetManagedNFProfile returns the ManagedNFProfile field value if set, zero value otherwise.
func (o *LmfFunctionSingleAllOfAttributes) GetManagedNFProfile() ManagedNFProfile {
	if o == nil || IsNil(o.ManagedNFProfile) {
		var ret ManagedNFProfile
		return ret
	}
	return *o.ManagedNFProfile
}

// GetManagedNFProfileOk returns a tuple with the ManagedNFProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfFunctionSingleAllOfAttributes) GetManagedNFProfileOk() (*ManagedNFProfile, bool) {
	if o == nil || IsNil(o.ManagedNFProfile) {
		return nil, false
	}
	return o.ManagedNFProfile, true
}

// HasManagedNFProfile returns a boolean if a field has been set.
func (o *LmfFunctionSingleAllOfAttributes) HasManagedNFProfile() bool {
	if o != nil && !IsNil(o.ManagedNFProfile) {
		return true
	}

	return false
}

// SetManagedNFProfile gets a reference to the given ManagedNFProfile and assigns it to the ManagedNFProfile field.
func (o *LmfFunctionSingleAllOfAttributes) SetManagedNFProfile(v ManagedNFProfile) {
	o.ManagedNFProfile = &v
}

// GetCommModelList returns the CommModelList field value if set, zero value otherwise.
func (o *LmfFunctionSingleAllOfAttributes) GetCommModelList() []CommModel {
	if o == nil || IsNil(o.CommModelList) {
		var ret []CommModel
		return ret
	}
	return o.CommModelList
}

// GetCommModelListOk returns a tuple with the CommModelList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfFunctionSingleAllOfAttributes) GetCommModelListOk() ([]CommModel, bool) {
	if o == nil || IsNil(o.CommModelList) {
		return nil, false
	}
	return o.CommModelList, true
}

// HasCommModelList returns a boolean if a field has been set.
func (o *LmfFunctionSingleAllOfAttributes) HasCommModelList() bool {
	if o != nil && !IsNil(o.CommModelList) {
		return true
	}

	return false
}

// SetCommModelList gets a reference to the given []CommModel and assigns it to the CommModelList field.
func (o *LmfFunctionSingleAllOfAttributes) SetCommModelList(v []CommModel) {
	o.CommModelList = v
}

func (o LmfFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LmfFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedFunctionAttr, errManagedFunctionAttr := json.Marshal(o.ManagedFunctionAttr)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	errManagedFunctionAttr = json.Unmarshal([]byte(serializedManagedFunctionAttr), &toSerialize)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	if !IsNil(o.PlmnIdList) {
		toSerialize["plmnIdList"] = o.PlmnIdList
	}
	if !IsNil(o.ManagedNFProfile) {
		toSerialize["managedNFProfile"] = o.ManagedNFProfile
	}
	if !IsNil(o.CommModelList) {
		toSerialize["commModelList"] = o.CommModelList
	}
	return toSerialize, nil
}

type NullableLmfFunctionSingleAllOfAttributes struct {
	value *LmfFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableLmfFunctionSingleAllOfAttributes) Get() *LmfFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableLmfFunctionSingleAllOfAttributes) Set(val *LmfFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableLmfFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableLmfFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLmfFunctionSingleAllOfAttributes(val *LmfFunctionSingleAllOfAttributes) *NullableLmfFunctionSingleAllOfAttributes {
	return &NullableLmfFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableLmfFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLmfFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

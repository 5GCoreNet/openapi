/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// checks if the LmfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LmfInfo{}

// LmfInfo Information of an LMF NF Instance
type LmfInfo struct {
	ServingClientTypes []ExternalClientType `json:"servingClientTypes,omitempty"`
	// LMF identification.
	LmfId              *string              `json:"lmfId,omitempty"`
	ServingAccessTypes []AccessType         `json:"servingAccessTypes,omitempty"`
	ServingAnNodeTypes []AnNodeType         `json:"servingAnNodeTypes,omitempty"`
	ServingRatTypes    []RatType            `json:"servingRatTypes,omitempty"`
	TaiList            []Tai                `json:"taiList,omitempty"`
	TaiRangeList       []TaiRange           `json:"taiRangeList,omitempty"`
	SupportedGADShapes []SupportedGADShapes `json:"supportedGADShapes,omitempty"`
}

// NewLmfInfo instantiates a new LmfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLmfInfo() *LmfInfo {
	this := LmfInfo{}
	return &this
}

// NewLmfInfoWithDefaults instantiates a new LmfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLmfInfoWithDefaults() *LmfInfo {
	this := LmfInfo{}
	return &this
}

// GetServingClientTypes returns the ServingClientTypes field value if set, zero value otherwise.
func (o *LmfInfo) GetServingClientTypes() []ExternalClientType {
	if o == nil || IsNil(o.ServingClientTypes) {
		var ret []ExternalClientType
		return ret
	}
	return o.ServingClientTypes
}

// GetServingClientTypesOk returns a tuple with the ServingClientTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfInfo) GetServingClientTypesOk() ([]ExternalClientType, bool) {
	if o == nil || IsNil(o.ServingClientTypes) {
		return nil, false
	}
	return o.ServingClientTypes, true
}

// HasServingClientTypes returns a boolean if a field has been set.
func (o *LmfInfo) HasServingClientTypes() bool {
	if o != nil && !IsNil(o.ServingClientTypes) {
		return true
	}

	return false
}

// SetServingClientTypes gets a reference to the given []ExternalClientType and assigns it to the ServingClientTypes field.
func (o *LmfInfo) SetServingClientTypes(v []ExternalClientType) {
	o.ServingClientTypes = v
}

// GetLmfId returns the LmfId field value if set, zero value otherwise.
func (o *LmfInfo) GetLmfId() string {
	if o == nil || IsNil(o.LmfId) {
		var ret string
		return ret
	}
	return *o.LmfId
}

// GetLmfIdOk returns a tuple with the LmfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfInfo) GetLmfIdOk() (*string, bool) {
	if o == nil || IsNil(o.LmfId) {
		return nil, false
	}
	return o.LmfId, true
}

// HasLmfId returns a boolean if a field has been set.
func (o *LmfInfo) HasLmfId() bool {
	if o != nil && !IsNil(o.LmfId) {
		return true
	}

	return false
}

// SetLmfId gets a reference to the given string and assigns it to the LmfId field.
func (o *LmfInfo) SetLmfId(v string) {
	o.LmfId = &v
}

// GetServingAccessTypes returns the ServingAccessTypes field value if set, zero value otherwise.
func (o *LmfInfo) GetServingAccessTypes() []AccessType {
	if o == nil || IsNil(o.ServingAccessTypes) {
		var ret []AccessType
		return ret
	}
	return o.ServingAccessTypes
}

// GetServingAccessTypesOk returns a tuple with the ServingAccessTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfInfo) GetServingAccessTypesOk() ([]AccessType, bool) {
	if o == nil || IsNil(o.ServingAccessTypes) {
		return nil, false
	}
	return o.ServingAccessTypes, true
}

// HasServingAccessTypes returns a boolean if a field has been set.
func (o *LmfInfo) HasServingAccessTypes() bool {
	if o != nil && !IsNil(o.ServingAccessTypes) {
		return true
	}

	return false
}

// SetServingAccessTypes gets a reference to the given []AccessType and assigns it to the ServingAccessTypes field.
func (o *LmfInfo) SetServingAccessTypes(v []AccessType) {
	o.ServingAccessTypes = v
}

// GetServingAnNodeTypes returns the ServingAnNodeTypes field value if set, zero value otherwise.
func (o *LmfInfo) GetServingAnNodeTypes() []AnNodeType {
	if o == nil || IsNil(o.ServingAnNodeTypes) {
		var ret []AnNodeType
		return ret
	}
	return o.ServingAnNodeTypes
}

// GetServingAnNodeTypesOk returns a tuple with the ServingAnNodeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfInfo) GetServingAnNodeTypesOk() ([]AnNodeType, bool) {
	if o == nil || IsNil(o.ServingAnNodeTypes) {
		return nil, false
	}
	return o.ServingAnNodeTypes, true
}

// HasServingAnNodeTypes returns a boolean if a field has been set.
func (o *LmfInfo) HasServingAnNodeTypes() bool {
	if o != nil && !IsNil(o.ServingAnNodeTypes) {
		return true
	}

	return false
}

// SetServingAnNodeTypes gets a reference to the given []AnNodeType and assigns it to the ServingAnNodeTypes field.
func (o *LmfInfo) SetServingAnNodeTypes(v []AnNodeType) {
	o.ServingAnNodeTypes = v
}

// GetServingRatTypes returns the ServingRatTypes field value if set, zero value otherwise.
func (o *LmfInfo) GetServingRatTypes() []RatType {
	if o == nil || IsNil(o.ServingRatTypes) {
		var ret []RatType
		return ret
	}
	return o.ServingRatTypes
}

// GetServingRatTypesOk returns a tuple with the ServingRatTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfInfo) GetServingRatTypesOk() ([]RatType, bool) {
	if o == nil || IsNil(o.ServingRatTypes) {
		return nil, false
	}
	return o.ServingRatTypes, true
}

// HasServingRatTypes returns a boolean if a field has been set.
func (o *LmfInfo) HasServingRatTypes() bool {
	if o != nil && !IsNil(o.ServingRatTypes) {
		return true
	}

	return false
}

// SetServingRatTypes gets a reference to the given []RatType and assigns it to the ServingRatTypes field.
func (o *LmfInfo) SetServingRatTypes(v []RatType) {
	o.ServingRatTypes = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *LmfInfo) GetTaiList() []Tai {
	if o == nil || IsNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TaiList) {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *LmfInfo) HasTaiList() bool {
	if o != nil && !IsNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *LmfInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *LmfInfo) GetTaiRangeList() []TaiRange {
	if o == nil || IsNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || IsNil(o.TaiRangeList) {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *LmfInfo) HasTaiRangeList() bool {
	if o != nil && !IsNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *LmfInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetSupportedGADShapes returns the SupportedGADShapes field value if set, zero value otherwise.
func (o *LmfInfo) GetSupportedGADShapes() []SupportedGADShapes {
	if o == nil || IsNil(o.SupportedGADShapes) {
		var ret []SupportedGADShapes
		return ret
	}
	return o.SupportedGADShapes
}

// GetSupportedGADShapesOk returns a tuple with the SupportedGADShapes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfInfo) GetSupportedGADShapesOk() ([]SupportedGADShapes, bool) {
	if o == nil || IsNil(o.SupportedGADShapes) {
		return nil, false
	}
	return o.SupportedGADShapes, true
}

// HasSupportedGADShapes returns a boolean if a field has been set.
func (o *LmfInfo) HasSupportedGADShapes() bool {
	if o != nil && !IsNil(o.SupportedGADShapes) {
		return true
	}

	return false
}

// SetSupportedGADShapes gets a reference to the given []SupportedGADShapes and assigns it to the SupportedGADShapes field.
func (o *LmfInfo) SetSupportedGADShapes(v []SupportedGADShapes) {
	o.SupportedGADShapes = v
}

func (o LmfInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LmfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServingClientTypes) {
		toSerialize["servingClientTypes"] = o.ServingClientTypes
	}
	if !IsNil(o.LmfId) {
		toSerialize["lmfId"] = o.LmfId
	}
	if !IsNil(o.ServingAccessTypes) {
		toSerialize["servingAccessTypes"] = o.ServingAccessTypes
	}
	if !IsNil(o.ServingAnNodeTypes) {
		toSerialize["servingAnNodeTypes"] = o.ServingAnNodeTypes
	}
	if !IsNil(o.ServingRatTypes) {
		toSerialize["servingRatTypes"] = o.ServingRatTypes
	}
	if !IsNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !IsNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !IsNil(o.SupportedGADShapes) {
		toSerialize["supportedGADShapes"] = o.SupportedGADShapes
	}
	return toSerialize, nil
}

type NullableLmfInfo struct {
	value *LmfInfo
	isSet bool
}

func (v NullableLmfInfo) Get() *LmfInfo {
	return v.value
}

func (v *NullableLmfInfo) Set(val *LmfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLmfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLmfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLmfInfo(val *LmfInfo) *NullableLmfInfo {
	return &NullableLmfInfo{value: val, isSet: true}
}

func (v NullableLmfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLmfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

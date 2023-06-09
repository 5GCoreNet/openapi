/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// checks if the ServAttrCom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServAttrCom{}

// ServAttrCom struct for ServAttrCom
type ServAttrCom struct {
	Category *Category `json:"category,omitempty"`
	Tagging  []string  `json:"tagging,omitempty"`
	Exposure *Exposure `json:"exposure,omitempty"`
}

// NewServAttrCom instantiates a new ServAttrCom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServAttrCom() *ServAttrCom {
	this := ServAttrCom{}
	return &this
}

// NewServAttrComWithDefaults instantiates a new ServAttrCom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServAttrComWithDefaults() *ServAttrCom {
	this := ServAttrCom{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ServAttrCom) GetCategory() Category {
	if o == nil || IsNil(o.Category) {
		var ret Category
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServAttrCom) GetCategoryOk() (*Category, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ServAttrCom) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given Category and assigns it to the Category field.
func (o *ServAttrCom) SetCategory(v Category) {
	o.Category = &v
}

// GetTagging returns the Tagging field value if set, zero value otherwise.
func (o *ServAttrCom) GetTagging() []string {
	if o == nil || IsNil(o.Tagging) {
		var ret []string
		return ret
	}
	return o.Tagging
}

// GetTaggingOk returns a tuple with the Tagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServAttrCom) GetTaggingOk() ([]string, bool) {
	if o == nil || IsNil(o.Tagging) {
		return nil, false
	}
	return o.Tagging, true
}

// HasTagging returns a boolean if a field has been set.
func (o *ServAttrCom) HasTagging() bool {
	if o != nil && !IsNil(o.Tagging) {
		return true
	}

	return false
}

// SetTagging gets a reference to the given []string and assigns it to the Tagging field.
func (o *ServAttrCom) SetTagging(v []string) {
	o.Tagging = v
}

// GetExposure returns the Exposure field value if set, zero value otherwise.
func (o *ServAttrCom) GetExposure() Exposure {
	if o == nil || IsNil(o.Exposure) {
		var ret Exposure
		return ret
	}
	return *o.Exposure
}

// GetExposureOk returns a tuple with the Exposure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServAttrCom) GetExposureOk() (*Exposure, bool) {
	if o == nil || IsNil(o.Exposure) {
		return nil, false
	}
	return o.Exposure, true
}

// HasExposure returns a boolean if a field has been set.
func (o *ServAttrCom) HasExposure() bool {
	if o != nil && !IsNil(o.Exposure) {
		return true
	}

	return false
}

// SetExposure gets a reference to the given Exposure and assigns it to the Exposure field.
func (o *ServAttrCom) SetExposure(v Exposure) {
	o.Exposure = &v
}

func (o ServAttrCom) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServAttrCom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Tagging) {
		toSerialize["tagging"] = o.Tagging
	}
	if !IsNil(o.Exposure) {
		toSerialize["exposure"] = o.Exposure
	}
	return toSerialize, nil
}

type NullableServAttrCom struct {
	value *ServAttrCom
	isSet bool
}

func (v NullableServAttrCom) Get() *ServAttrCom {
	return v.value
}

func (v *NullableServAttrCom) Set(val *ServAttrCom) {
	v.value = val
	v.isSet = true
}

func (v NullableServAttrCom) IsSet() bool {
	return v.isSet
}

func (v *NullableServAttrCom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServAttrCom(val *ServAttrCom) *NullableServAttrCom {
	return &NullableServAttrCom{value: val, isSet: true}
}

func (v NullableServAttrCom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServAttrCom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

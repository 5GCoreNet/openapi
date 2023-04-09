/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the CollectiveBehaviourFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectiveBehaviourFilter{}

// CollectiveBehaviourFilter Contains the collective behaviour filter information to be collected from UE.
type CollectiveBehaviourFilter struct {
	Type CollectiveBehaviourFilterType `json:"type"`
	// Value of the parameter type as in the type attribute.
	Value string `json:"value"`
	// Indicates whether request list of UE IDs that fulfill a collective behaviour within the area of interest. This attribute shall set to \"true\" if request the list of UE IDs, otherwise, set to \"false\". May only be present and sets to \"true\" if \"AfEvent\" sets to \"COLLECTIVE_BEHAVIOUR\".
	ListOfUeInd *bool `json:"listOfUeInd,omitempty"`
}

// NewCollectiveBehaviourFilter instantiates a new CollectiveBehaviourFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectiveBehaviourFilter(type_ CollectiveBehaviourFilterType, value string) *CollectiveBehaviourFilter {
	this := CollectiveBehaviourFilter{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewCollectiveBehaviourFilterWithDefaults instantiates a new CollectiveBehaviourFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectiveBehaviourFilterWithDefaults() *CollectiveBehaviourFilter {
	this := CollectiveBehaviourFilter{}
	return &this
}

// GetType returns the Type field value
func (o *CollectiveBehaviourFilter) GetType() CollectiveBehaviourFilterType {
	if o == nil {
		var ret CollectiveBehaviourFilterType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CollectiveBehaviourFilter) GetTypeOk() (*CollectiveBehaviourFilterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CollectiveBehaviourFilter) SetType(v CollectiveBehaviourFilterType) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *CollectiveBehaviourFilter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CollectiveBehaviourFilter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CollectiveBehaviourFilter) SetValue(v string) {
	o.Value = v
}

// GetListOfUeInd returns the ListOfUeInd field value if set, zero value otherwise.
func (o *CollectiveBehaviourFilter) GetListOfUeInd() bool {
	if o == nil || IsNil(o.ListOfUeInd) {
		var ret bool
		return ret
	}
	return *o.ListOfUeInd
}

// GetListOfUeIndOk returns a tuple with the ListOfUeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectiveBehaviourFilter) GetListOfUeIndOk() (*bool, bool) {
	if o == nil || IsNil(o.ListOfUeInd) {
		return nil, false
	}
	return o.ListOfUeInd, true
}

// HasListOfUeInd returns a boolean if a field has been set.
func (o *CollectiveBehaviourFilter) HasListOfUeInd() bool {
	if o != nil && !IsNil(o.ListOfUeInd) {
		return true
	}

	return false
}

// SetListOfUeInd gets a reference to the given bool and assigns it to the ListOfUeInd field.
func (o *CollectiveBehaviourFilter) SetListOfUeInd(v bool) {
	o.ListOfUeInd = &v
}

func (o CollectiveBehaviourFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectiveBehaviourFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	if !IsNil(o.ListOfUeInd) {
		toSerialize["listOfUeInd"] = o.ListOfUeInd
	}
	return toSerialize, nil
}

type NullableCollectiveBehaviourFilter struct {
	value *CollectiveBehaviourFilter
	isSet bool
}

func (v NullableCollectiveBehaviourFilter) Get() *CollectiveBehaviourFilter {
	return v.value
}

func (v *NullableCollectiveBehaviourFilter) Set(val *CollectiveBehaviourFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectiveBehaviourFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectiveBehaviourFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectiveBehaviourFilter(val *CollectiveBehaviourFilter) *NullableCollectiveBehaviourFilter {
	return &NullableCollectiveBehaviourFilter{value: val, isSet: true}
}

func (v NullableCollectiveBehaviourFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectiveBehaviourFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

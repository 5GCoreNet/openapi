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

// checks if the PredefinedPccRuleSetSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredefinedPccRuleSetSingle{}

// PredefinedPccRuleSetSingle struct for PredefinedPccRuleSetSingle
type PredefinedPccRuleSetSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *PredefinedPccRuleSetSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewPredefinedPccRuleSetSingle instantiates a new PredefinedPccRuleSetSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredefinedPccRuleSetSingle(id NullableString) *PredefinedPccRuleSetSingle {
	this := PredefinedPccRuleSetSingle{}
	this.Id = id
	return &this
}

// NewPredefinedPccRuleSetSingleWithDefaults instantiates a new PredefinedPccRuleSetSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredefinedPccRuleSetSingleWithDefaults() *PredefinedPccRuleSetSingle {
	this := PredefinedPccRuleSetSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PredefinedPccRuleSetSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PredefinedPccRuleSetSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *PredefinedPccRuleSetSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *PredefinedPccRuleSetSingle) GetObjectClass() string {
	if o == nil || IsNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredefinedPccRuleSetSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *PredefinedPccRuleSetSingle) HasObjectClass() bool {
	if o != nil && !IsNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *PredefinedPccRuleSetSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *PredefinedPccRuleSetSingle) GetObjectInstance() string {
	if o == nil || IsNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredefinedPccRuleSetSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *PredefinedPccRuleSetSingle) HasObjectInstance() bool {
	if o != nil && !IsNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *PredefinedPccRuleSetSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *PredefinedPccRuleSetSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || IsNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredefinedPccRuleSetSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || IsNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *PredefinedPccRuleSetSingle) HasVsDataContainer() bool {
	if o != nil && !IsNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *PredefinedPccRuleSetSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PredefinedPccRuleSetSingle) GetAttributes() PredefinedPccRuleSetSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret PredefinedPccRuleSetSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredefinedPccRuleSetSingle) GetAttributesOk() (*PredefinedPccRuleSetSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PredefinedPccRuleSetSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PredefinedPccRuleSetSingleAllOfAttributes and assigns it to the Attributes field.
func (o *PredefinedPccRuleSetSingle) SetAttributes(v PredefinedPccRuleSetSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o PredefinedPccRuleSetSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredefinedPccRuleSetSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if !IsNil(o.ObjectClass) {
		toSerialize["objectClass"] = o.ObjectClass
	}
	if !IsNil(o.ObjectInstance) {
		toSerialize["objectInstance"] = o.ObjectInstance
	}
	if !IsNil(o.VsDataContainer) {
		toSerialize["VsDataContainer"] = o.VsDataContainer
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullablePredefinedPccRuleSetSingle struct {
	value *PredefinedPccRuleSetSingle
	isSet bool
}

func (v NullablePredefinedPccRuleSetSingle) Get() *PredefinedPccRuleSetSingle {
	return v.value
}

func (v *NullablePredefinedPccRuleSetSingle) Set(val *PredefinedPccRuleSetSingle) {
	v.value = val
	v.isSet = true
}

func (v NullablePredefinedPccRuleSetSingle) IsSet() bool {
	return v.isSet
}

func (v *NullablePredefinedPccRuleSetSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredefinedPccRuleSetSingle(val *PredefinedPccRuleSetSingle) *NullablePredefinedPccRuleSetSingle {
	return &NullablePredefinedPccRuleSetSingle{value: val, isSet: true}
}

func (v NullablePredefinedPccRuleSetSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredefinedPccRuleSetSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


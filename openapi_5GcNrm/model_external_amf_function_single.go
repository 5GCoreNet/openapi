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

// checks if the ExternalAmfFunctionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalAmfFunctionSingle{}

// ExternalAmfFunctionSingle struct for ExternalAmfFunctionSingle
type ExternalAmfFunctionSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *ExternalAmfFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewExternalAmfFunctionSingle instantiates a new ExternalAmfFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAmfFunctionSingle(id NullableString) *ExternalAmfFunctionSingle {
	this := ExternalAmfFunctionSingle{}
	this.Id = id
	return &this
}

// NewExternalAmfFunctionSingleWithDefaults instantiates a new ExternalAmfFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAmfFunctionSingleWithDefaults() *ExternalAmfFunctionSingle {
	this := ExternalAmfFunctionSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExternalAmfFunctionSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalAmfFunctionSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *ExternalAmfFunctionSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *ExternalAmfFunctionSingle) GetObjectClass() string {
	if o == nil || IsNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAmfFunctionSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *ExternalAmfFunctionSingle) HasObjectClass() bool {
	if o != nil && !IsNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *ExternalAmfFunctionSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *ExternalAmfFunctionSingle) GetObjectInstance() string {
	if o == nil || IsNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAmfFunctionSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *ExternalAmfFunctionSingle) HasObjectInstance() bool {
	if o != nil && !IsNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *ExternalAmfFunctionSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *ExternalAmfFunctionSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || IsNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAmfFunctionSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || IsNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *ExternalAmfFunctionSingle) HasVsDataContainer() bool {
	if o != nil && !IsNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *ExternalAmfFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalAmfFunctionSingle) GetAttributes() ExternalAmfFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ExternalAmfFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAmfFunctionSingle) GetAttributesOk() (*ExternalAmfFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalAmfFunctionSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ExternalAmfFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ExternalAmfFunctionSingle) SetAttributes(v ExternalAmfFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o ExternalAmfFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalAmfFunctionSingle) ToMap() (map[string]interface{}, error) {
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

type NullableExternalAmfFunctionSingle struct {
	value *ExternalAmfFunctionSingle
	isSet bool
}

func (v NullableExternalAmfFunctionSingle) Get() *ExternalAmfFunctionSingle {
	return v.value
}

func (v *NullableExternalAmfFunctionSingle) Set(val *ExternalAmfFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAmfFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAmfFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAmfFunctionSingle(val *ExternalAmfFunctionSingle) *NullableExternalAmfFunctionSingle {
	return &NullableExternalAmfFunctionSingle{value: val, isSet: true}
}

func (v NullableExternalAmfFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAmfFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the EPX2CSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPX2CSingle{}

// EPX2CSingle struct for EPX2CSingle
type EPX2CSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *EPF1CSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEPX2CSingle instantiates a new EPX2CSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPX2CSingle(id NullableString) *EPX2CSingle {
	this := EPX2CSingle{}
	this.Id = id
	return &this
}

// NewEPX2CSingleWithDefaults instantiates a new EPX2CSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPX2CSingleWithDefaults() *EPX2CSingle {
	this := EPX2CSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EPX2CSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EPX2CSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *EPX2CSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *EPX2CSingle) GetObjectClass() string {
	if o == nil || IsNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPX2CSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *EPX2CSingle) HasObjectClass() bool {
	if o != nil && !IsNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *EPX2CSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *EPX2CSingle) GetObjectInstance() string {
	if o == nil || IsNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPX2CSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *EPX2CSingle) HasObjectInstance() bool {
	if o != nil && !IsNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *EPX2CSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *EPX2CSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || IsNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPX2CSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || IsNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *EPX2CSingle) HasVsDataContainer() bool {
	if o != nil && !IsNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *EPX2CSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPX2CSingle) GetAttributes() EPF1CSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret EPF1CSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPX2CSingle) GetAttributesOk() (*EPF1CSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPX2CSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPF1CSingleAllOfAttributes and assigns it to the Attributes field.
func (o *EPX2CSingle) SetAttributes(v EPF1CSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EPX2CSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPX2CSingle) ToMap() (map[string]interface{}, error) {
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

type NullableEPX2CSingle struct {
	value *EPX2CSingle
	isSet bool
}

func (v NullableEPX2CSingle) Get() *EPX2CSingle {
	return v.value
}

func (v *NullableEPX2CSingle) Set(val *EPX2CSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableEPX2CSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableEPX2CSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPX2CSingle(val *EPX2CSingle) *NullableEPX2CSingle {
	return &NullableEPX2CSingle{value: val, isSet: true}
}

func (v NullableEPX2CSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPX2CSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



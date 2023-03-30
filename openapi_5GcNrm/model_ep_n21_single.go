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

// checks if the EPN21Single type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPN21Single{}

// EPN21Single struct for EPN21Single
type EPN21Single struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *EPN2SingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEPN21Single instantiates a new EPN21Single object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPN21Single(id NullableString) *EPN21Single {
	this := EPN21Single{}
	this.Id = id
	return &this
}

// NewEPN21SingleWithDefaults instantiates a new EPN21Single object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPN21SingleWithDefaults() *EPN21Single {
	this := EPN21Single{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EPN21Single) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EPN21Single) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *EPN21Single) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *EPN21Single) GetObjectClass() string {
	if o == nil || IsNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPN21Single) GetObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *EPN21Single) HasObjectClass() bool {
	if o != nil && !IsNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *EPN21Single) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *EPN21Single) GetObjectInstance() string {
	if o == nil || IsNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPN21Single) GetObjectInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *EPN21Single) HasObjectInstance() bool {
	if o != nil && !IsNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *EPN21Single) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *EPN21Single) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || IsNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPN21Single) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || IsNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *EPN21Single) HasVsDataContainer() bool {
	if o != nil && !IsNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *EPN21Single) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPN21Single) GetAttributes() EPN2SingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret EPN2SingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPN21Single) GetAttributesOk() (*EPN2SingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPN21Single) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPN2SingleAllOfAttributes and assigns it to the Attributes field.
func (o *EPN21Single) SetAttributes(v EPN2SingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EPN21Single) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPN21Single) ToMap() (map[string]interface{}, error) {
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

type NullableEPN21Single struct {
	value *EPN21Single
	isSet bool
}

func (v NullableEPN21Single) Get() *EPN21Single {
	return v.value
}

func (v *NullableEPN21Single) Set(val *EPN21Single) {
	v.value = val
	v.isSet = true
}

func (v NullableEPN21Single) IsSet() bool {
	return v.isSet
}

func (v *NullableEPN21Single) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPN21Single(val *EPN21Single) *NullableEPN21Single {
	return &NullableEPN21Single{value: val, isSet: true}
}

func (v NullableEPN21Single) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPN21Single) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



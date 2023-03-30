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

// checks if the CCOOvershootCoverageParametersSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CCOOvershootCoverageParametersSingle{}

// CCOOvershootCoverageParametersSingle struct for CCOOvershootCoverageParametersSingle
type CCOOvershootCoverageParametersSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *CCOParametersAttrAllOfAttributes `json:"attributes,omitempty"`
}

// NewCCOOvershootCoverageParametersSingle instantiates a new CCOOvershootCoverageParametersSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCCOOvershootCoverageParametersSingle(id NullableString) *CCOOvershootCoverageParametersSingle {
	this := CCOOvershootCoverageParametersSingle{}
	this.Id = id
	return &this
}

// NewCCOOvershootCoverageParametersSingleWithDefaults instantiates a new CCOOvershootCoverageParametersSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCCOOvershootCoverageParametersSingleWithDefaults() *CCOOvershootCoverageParametersSingle {
	this := CCOOvershootCoverageParametersSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CCOOvershootCoverageParametersSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CCOOvershootCoverageParametersSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *CCOOvershootCoverageParametersSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *CCOOvershootCoverageParametersSingle) GetObjectClass() string {
	if o == nil || IsNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOOvershootCoverageParametersSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *CCOOvershootCoverageParametersSingle) HasObjectClass() bool {
	if o != nil && !IsNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *CCOOvershootCoverageParametersSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *CCOOvershootCoverageParametersSingle) GetObjectInstance() string {
	if o == nil || IsNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOOvershootCoverageParametersSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *CCOOvershootCoverageParametersSingle) HasObjectInstance() bool {
	if o != nil && !IsNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *CCOOvershootCoverageParametersSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *CCOOvershootCoverageParametersSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || IsNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOOvershootCoverageParametersSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || IsNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *CCOOvershootCoverageParametersSingle) HasVsDataContainer() bool {
	if o != nil && !IsNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *CCOOvershootCoverageParametersSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CCOOvershootCoverageParametersSingle) GetAttributes() CCOParametersAttrAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret CCOParametersAttrAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOOvershootCoverageParametersSingle) GetAttributesOk() (*CCOParametersAttrAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CCOOvershootCoverageParametersSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CCOParametersAttrAllOfAttributes and assigns it to the Attributes field.
func (o *CCOOvershootCoverageParametersSingle) SetAttributes(v CCOParametersAttrAllOfAttributes) {
	o.Attributes = &v
}

func (o CCOOvershootCoverageParametersSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CCOOvershootCoverageParametersSingle) ToMap() (map[string]interface{}, error) {
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

type NullableCCOOvershootCoverageParametersSingle struct {
	value *CCOOvershootCoverageParametersSingle
	isSet bool
}

func (v NullableCCOOvershootCoverageParametersSingle) Get() *CCOOvershootCoverageParametersSingle {
	return v.value
}

func (v *NullableCCOOvershootCoverageParametersSingle) Set(val *CCOOvershootCoverageParametersSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableCCOOvershootCoverageParametersSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableCCOOvershootCoverageParametersSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCCOOvershootCoverageParametersSingle(val *CCOOvershootCoverageParametersSingle) *NullableCCOOvershootCoverageParametersSingle {
	return &NullableCCOOvershootCoverageParametersSingle{value: val, isSet: true}
}

func (v NullableCCOOvershootCoverageParametersSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCCOOvershootCoverageParametersSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



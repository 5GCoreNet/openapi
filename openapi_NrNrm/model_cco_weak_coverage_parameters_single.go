/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the CCOWeakCoverageParametersSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CCOWeakCoverageParametersSingle{}

// CCOWeakCoverageParametersSingle struct for CCOWeakCoverageParametersSingle
type CCOWeakCoverageParametersSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *CCOParametersAttrAllOfAttributes `json:"attributes,omitempty"`
}

// NewCCOWeakCoverageParametersSingle instantiates a new CCOWeakCoverageParametersSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCCOWeakCoverageParametersSingle(id NullableString) *CCOWeakCoverageParametersSingle {
	this := CCOWeakCoverageParametersSingle{}
	this.Id = id
	return &this
}

// NewCCOWeakCoverageParametersSingleWithDefaults instantiates a new CCOWeakCoverageParametersSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCCOWeakCoverageParametersSingleWithDefaults() *CCOWeakCoverageParametersSingle {
	this := CCOWeakCoverageParametersSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CCOWeakCoverageParametersSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CCOWeakCoverageParametersSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *CCOWeakCoverageParametersSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *CCOWeakCoverageParametersSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOWeakCoverageParametersSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *CCOWeakCoverageParametersSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *CCOWeakCoverageParametersSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *CCOWeakCoverageParametersSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOWeakCoverageParametersSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *CCOWeakCoverageParametersSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *CCOWeakCoverageParametersSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *CCOWeakCoverageParametersSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOWeakCoverageParametersSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *CCOWeakCoverageParametersSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *CCOWeakCoverageParametersSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CCOWeakCoverageParametersSingle) GetAttributes() CCOParametersAttrAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret CCOParametersAttrAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOWeakCoverageParametersSingle) GetAttributesOk() (*CCOParametersAttrAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CCOWeakCoverageParametersSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CCOParametersAttrAllOfAttributes and assigns it to the Attributes field.
func (o *CCOWeakCoverageParametersSingle) SetAttributes(v CCOParametersAttrAllOfAttributes) {
	o.Attributes = &v
}

func (o CCOWeakCoverageParametersSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CCOWeakCoverageParametersSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if !isNil(o.ObjectClass) {
		toSerialize["objectClass"] = o.ObjectClass
	}
	if !isNil(o.ObjectInstance) {
		toSerialize["objectInstance"] = o.ObjectInstance
	}
	if !isNil(o.VsDataContainer) {
		toSerialize["VsDataContainer"] = o.VsDataContainer
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableCCOWeakCoverageParametersSingle struct {
	value *CCOWeakCoverageParametersSingle
	isSet bool
}

func (v NullableCCOWeakCoverageParametersSingle) Get() *CCOWeakCoverageParametersSingle {
	return v.value
}

func (v *NullableCCOWeakCoverageParametersSingle) Set(val *CCOWeakCoverageParametersSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableCCOWeakCoverageParametersSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableCCOWeakCoverageParametersSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCCOWeakCoverageParametersSingle(val *CCOWeakCoverageParametersSingle) *NullableCCOWeakCoverageParametersSingle {
	return &NullableCCOWeakCoverageParametersSingle{value: val, isSet: true}
}

func (v NullableCCOWeakCoverageParametersSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCCOWeakCoverageParametersSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



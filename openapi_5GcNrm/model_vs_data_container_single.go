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

// checks if the VsDataContainerSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VsDataContainerSingle{}

// VsDataContainerSingle struct for VsDataContainerSingle
type VsDataContainerSingle struct {
	Id *string `json:"id,omitempty"`
	Attributes *VsDataContainerSingleAttributes `json:"attributes,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
}

// NewVsDataContainerSingle instantiates a new VsDataContainerSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVsDataContainerSingle() *VsDataContainerSingle {
	this := VsDataContainerSingle{}
	return &this
}

// NewVsDataContainerSingleWithDefaults instantiates a new VsDataContainerSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVsDataContainerSingleWithDefaults() *VsDataContainerSingle {
	this := VsDataContainerSingle{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VsDataContainerSingle) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsDataContainerSingle) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VsDataContainerSingle) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VsDataContainerSingle) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *VsDataContainerSingle) GetAttributes() VsDataContainerSingleAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret VsDataContainerSingleAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsDataContainerSingle) GetAttributesOk() (*VsDataContainerSingleAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *VsDataContainerSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given VsDataContainerSingleAttributes and assigns it to the Attributes field.
func (o *VsDataContainerSingle) SetAttributes(v VsDataContainerSingleAttributes) {
	o.Attributes = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *VsDataContainerSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || IsNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsDataContainerSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || IsNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *VsDataContainerSingle) HasVsDataContainer() bool {
	if o != nil && !IsNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *VsDataContainerSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

func (o VsDataContainerSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VsDataContainerSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.VsDataContainer) {
		toSerialize["VsDataContainer"] = o.VsDataContainer
	}
	return toSerialize, nil
}

type NullableVsDataContainerSingle struct {
	value *VsDataContainerSingle
	isSet bool
}

func (v NullableVsDataContainerSingle) Get() *VsDataContainerSingle {
	return v.value
}

func (v *NullableVsDataContainerSingle) Set(val *VsDataContainerSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableVsDataContainerSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableVsDataContainerSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVsDataContainerSingle(val *VsDataContainerSingle) *NullableVsDataContainerSingle {
	return &NullableVsDataContainerSingle{value: val, isSet: true}
}

func (v NullableVsDataContainerSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVsDataContainerSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



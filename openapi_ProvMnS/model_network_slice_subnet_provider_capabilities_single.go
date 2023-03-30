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

// checks if the NetworkSliceSubnetProviderCapabilitiesSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkSliceSubnetProviderCapabilitiesSingle{}

// NetworkSliceSubnetProviderCapabilitiesSingle struct for NetworkSliceSubnetProviderCapabilitiesSingle
type NetworkSliceSubnetProviderCapabilitiesSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewNetworkSliceSubnetProviderCapabilitiesSingle instantiates a new NetworkSliceSubnetProviderCapabilitiesSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSliceSubnetProviderCapabilitiesSingle(id NullableString) *NetworkSliceSubnetProviderCapabilitiesSingle {
	this := NetworkSliceSubnetProviderCapabilitiesSingle{}
	this.Id = id
	return &this
}

// NewNetworkSliceSubnetProviderCapabilitiesSingleWithDefaults instantiates a new NetworkSliceSubnetProviderCapabilitiesSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSliceSubnetProviderCapabilitiesSingleWithDefaults() *NetworkSliceSubnetProviderCapabilitiesSingle {
	this := NetworkSliceSubnetProviderCapabilitiesSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetObjectClass() string {
	if o == nil || IsNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) HasObjectClass() bool {
	if o != nil && !IsNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetObjectInstance() string {
	if o == nil || IsNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) HasObjectInstance() bool {
	if o != nil && !IsNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || IsNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || IsNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) HasVsDataContainer() bool {
	if o != nil && !IsNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetAttributes() NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetAttributesOk() (*NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes and assigns it to the Attributes field.
func (o *NetworkSliceSubnetProviderCapabilitiesSingle) SetAttributes(v NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o NetworkSliceSubnetProviderCapabilitiesSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkSliceSubnetProviderCapabilitiesSingle) ToMap() (map[string]interface{}, error) {
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

type NullableNetworkSliceSubnetProviderCapabilitiesSingle struct {
	value *NetworkSliceSubnetProviderCapabilitiesSingle
	isSet bool
}

func (v NullableNetworkSliceSubnetProviderCapabilitiesSingle) Get() *NetworkSliceSubnetProviderCapabilitiesSingle {
	return v.value
}

func (v *NullableNetworkSliceSubnetProviderCapabilitiesSingle) Set(val *NetworkSliceSubnetProviderCapabilitiesSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSliceSubnetProviderCapabilitiesSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSliceSubnetProviderCapabilitiesSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSliceSubnetProviderCapabilitiesSingle(val *NetworkSliceSubnetProviderCapabilitiesSingle) *NullableNetworkSliceSubnetProviderCapabilitiesSingle {
	return &NullableNetworkSliceSubnetProviderCapabilitiesSingle{value: val, isSet: true}
}

func (v NullableNetworkSliceSubnetProviderCapabilitiesSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSliceSubnetProviderCapabilitiesSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



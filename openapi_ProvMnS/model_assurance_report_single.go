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

// checks if the AssuranceReportSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssuranceReportSingle{}

// AssuranceReportSingle struct for AssuranceReportSingle
type AssuranceReportSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *AssuranceReportSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewAssuranceReportSingle instantiates a new AssuranceReportSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssuranceReportSingle(id NullableString) *AssuranceReportSingle {
	this := AssuranceReportSingle{}
	this.Id = id
	return &this
}

// NewAssuranceReportSingleWithDefaults instantiates a new AssuranceReportSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssuranceReportSingleWithDefaults() *AssuranceReportSingle {
	this := AssuranceReportSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AssuranceReportSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssuranceReportSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *AssuranceReportSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *AssuranceReportSingle) GetObjectClass() string {
	if o == nil || IsNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceReportSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *AssuranceReportSingle) HasObjectClass() bool {
	if o != nil && !IsNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *AssuranceReportSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *AssuranceReportSingle) GetObjectInstance() string {
	if o == nil || IsNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceReportSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *AssuranceReportSingle) HasObjectInstance() bool {
	if o != nil && !IsNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *AssuranceReportSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *AssuranceReportSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || IsNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceReportSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || IsNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *AssuranceReportSingle) HasVsDataContainer() bool {
	if o != nil && !IsNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *AssuranceReportSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AssuranceReportSingle) GetAttributes() AssuranceReportSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AssuranceReportSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceReportSingle) GetAttributesOk() (*AssuranceReportSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AssuranceReportSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AssuranceReportSingleAllOfAttributes and assigns it to the Attributes field.
func (o *AssuranceReportSingle) SetAttributes(v AssuranceReportSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o AssuranceReportSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssuranceReportSingle) ToMap() (map[string]interface{}, error) {
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

type NullableAssuranceReportSingle struct {
	value *AssuranceReportSingle
	isSet bool
}

func (v NullableAssuranceReportSingle) Get() *AssuranceReportSingle {
	return v.value
}

func (v *NullableAssuranceReportSingle) Set(val *AssuranceReportSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableAssuranceReportSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableAssuranceReportSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssuranceReportSingle(val *AssuranceReportSingle) *NullableAssuranceReportSingle {
	return &NullableAssuranceReportSingle{value: val, isSet: true}
}

func (v NullableAssuranceReportSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssuranceReportSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



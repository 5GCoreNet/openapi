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

// checks if the FileDownloadJobSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileDownloadJobSingle{}

// FileDownloadJobSingle struct for FileDownloadJobSingle
type FileDownloadJobSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *FileDownloadJobSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewFileDownloadJobSingle instantiates a new FileDownloadJobSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileDownloadJobSingle(id NullableString) *FileDownloadJobSingle {
	this := FileDownloadJobSingle{}
	this.Id = id
	return &this
}

// NewFileDownloadJobSingleWithDefaults instantiates a new FileDownloadJobSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDownloadJobSingleWithDefaults() *FileDownloadJobSingle {
	this := FileDownloadJobSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FileDownloadJobSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDownloadJobSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *FileDownloadJobSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *FileDownloadJobSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *FileDownloadJobSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *FileDownloadJobSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *FileDownloadJobSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *FileDownloadJobSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *FileDownloadJobSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *FileDownloadJobSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *FileDownloadJobSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *FileDownloadJobSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FileDownloadJobSingle) GetAttributes() FileDownloadJobSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret FileDownloadJobSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobSingle) GetAttributesOk() (*FileDownloadJobSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FileDownloadJobSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given FileDownloadJobSingleAllOfAttributes and assigns it to the Attributes field.
func (o *FileDownloadJobSingle) SetAttributes(v FileDownloadJobSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o FileDownloadJobSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileDownloadJobSingle) ToMap() (map[string]interface{}, error) {
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

type NullableFileDownloadJobSingle struct {
	value *FileDownloadJobSingle
	isSet bool
}

func (v NullableFileDownloadJobSingle) Get() *FileDownloadJobSingle {
	return v.value
}

func (v *NullableFileDownloadJobSingle) Set(val *FileDownloadJobSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDownloadJobSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDownloadJobSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDownloadJobSingle(val *FileDownloadJobSingle) *NullableFileDownloadJobSingle {
	return &NullableFileDownloadJobSingle{value: val, isSet: true}
}

func (v NullableFileDownloadJobSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDownloadJobSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



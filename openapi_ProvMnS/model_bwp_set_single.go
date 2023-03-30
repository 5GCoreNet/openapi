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

// checks if the BWPSetSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BWPSetSingle{}

// BWPSetSingle struct for BWPSetSingle
type BWPSetSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	BWPlist []string `json:"bWPlist,omitempty"`
}

// NewBWPSetSingle instantiates a new BWPSetSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBWPSetSingle(id NullableString) *BWPSetSingle {
	this := BWPSetSingle{}
	this.Id = id
	return &this
}

// NewBWPSetSingleWithDefaults instantiates a new BWPSetSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBWPSetSingleWithDefaults() *BWPSetSingle {
	this := BWPSetSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BWPSetSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BWPSetSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *BWPSetSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *BWPSetSingle) GetObjectClass() string {
	if o == nil || IsNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BWPSetSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *BWPSetSingle) HasObjectClass() bool {
	if o != nil && !IsNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *BWPSetSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *BWPSetSingle) GetObjectInstance() string {
	if o == nil || IsNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BWPSetSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *BWPSetSingle) HasObjectInstance() bool {
	if o != nil && !IsNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *BWPSetSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *BWPSetSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || IsNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BWPSetSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || IsNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *BWPSetSingle) HasVsDataContainer() bool {
	if o != nil && !IsNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *BWPSetSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetBWPlist returns the BWPlist field value if set, zero value otherwise.
func (o *BWPSetSingle) GetBWPlist() []string {
	if o == nil || IsNil(o.BWPlist) {
		var ret []string
		return ret
	}
	return o.BWPlist
}

// GetBWPlistOk returns a tuple with the BWPlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BWPSetSingle) GetBWPlistOk() ([]string, bool) {
	if o == nil || IsNil(o.BWPlist) {
		return nil, false
	}
	return o.BWPlist, true
}

// HasBWPlist returns a boolean if a field has been set.
func (o *BWPSetSingle) HasBWPlist() bool {
	if o != nil && !IsNil(o.BWPlist) {
		return true
	}

	return false
}

// SetBWPlist gets a reference to the given []string and assigns it to the BWPlist field.
func (o *BWPSetSingle) SetBWPlist(v []string) {
	o.BWPlist = v
}

func (o BWPSetSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BWPSetSingle) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BWPlist) {
		toSerialize["bWPlist"] = o.BWPlist
	}
	return toSerialize, nil
}

type NullableBWPSetSingle struct {
	value *BWPSetSingle
	isSet bool
}

func (v NullableBWPSetSingle) Get() *BWPSetSingle {
	return v.value
}

func (v *NullableBWPSetSingle) Set(val *BWPSetSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableBWPSetSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableBWPSetSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBWPSetSingle(val *BWPSetSingle) *NullableBWPSetSingle {
	return &NullableBWPSetSingle{value: val, isSet: true}
}

func (v NullableBWPSetSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBWPSetSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



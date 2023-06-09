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

// checks if the ResourceOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceOneOf{}

// ResourceOneOf struct for ResourceOneOf
type ResourceOneOf struct {
	Id             string                 `json:"id"`
	ObjectClass    *string                `json:"objectClass,omitempty"`
	ObjectInstance *string                `json:"objectInstance,omitempty"`
	Attributes     map[string]interface{} `json:"attributes,omitempty"`
}

// NewResourceOneOf instantiates a new ResourceOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOneOf(id string) *ResourceOneOf {
	this := ResourceOneOf{}
	this.Id = id
	return &this
}

// NewResourceOneOfWithDefaults instantiates a new ResourceOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOneOfWithDefaults() *ResourceOneOf {
	this := ResourceOneOf{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceOneOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceOneOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceOneOf) SetId(v string) {
	o.Id = v
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *ResourceOneOf) GetObjectClass() string {
	if o == nil || IsNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOneOf) GetObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *ResourceOneOf) HasObjectClass() bool {
	if o != nil && !IsNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *ResourceOneOf) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *ResourceOneOf) GetObjectInstance() string {
	if o == nil || IsNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOneOf) GetObjectInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *ResourceOneOf) HasObjectInstance() bool {
	if o != nil && !IsNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *ResourceOneOf) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ResourceOneOf) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOneOf) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ResourceOneOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *ResourceOneOf) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o ResourceOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.ObjectClass) {
		toSerialize["objectClass"] = o.ObjectClass
	}
	if !IsNil(o.ObjectInstance) {
		toSerialize["objectInstance"] = o.ObjectInstance
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableResourceOneOf struct {
	value *ResourceOneOf
	isSet bool
}

func (v NullableResourceOneOf) Get() *ResourceOneOf {
	return v.value
}

func (v *NullableResourceOneOf) Set(val *ResourceOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOneOf(val *ResourceOneOf) *NullableResourceOneOf {
	return &NullableResourceOneOf{value: val, isSet: true}
}

func (v NullableResourceOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

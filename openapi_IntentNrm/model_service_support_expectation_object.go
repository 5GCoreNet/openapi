/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
)

// checks if the ServiceSupportExpectationObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceSupportExpectationObject{}

// ServiceSupportExpectationObject This data type is the \"ExpectationObject\" data type with specialisations for ServiceSupportExpectation
type ServiceSupportExpectationObject struct {
	ObjectType *string `json:"objectType,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	ObjectContexts []ServiceSupportExpectationObjectObjectContextsInner `json:"objectContexts,omitempty"`
}

// NewServiceSupportExpectationObject instantiates a new ServiceSupportExpectationObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSupportExpectationObject() *ServiceSupportExpectationObject {
	this := ServiceSupportExpectationObject{}
	return &this
}

// NewServiceSupportExpectationObjectWithDefaults instantiates a new ServiceSupportExpectationObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSupportExpectationObjectWithDefaults() *ServiceSupportExpectationObject {
	this := ServiceSupportExpectationObject{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *ServiceSupportExpectationObject) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSupportExpectationObject) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *ServiceSupportExpectationObject) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *ServiceSupportExpectationObject) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *ServiceSupportExpectationObject) GetObjectInstance() string {
	if o == nil || IsNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSupportExpectationObject) GetObjectInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *ServiceSupportExpectationObject) HasObjectInstance() bool {
	if o != nil && !IsNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *ServiceSupportExpectationObject) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetObjectContexts returns the ObjectContexts field value if set, zero value otherwise.
func (o *ServiceSupportExpectationObject) GetObjectContexts() []ServiceSupportExpectationObjectObjectContextsInner {
	if o == nil || IsNil(o.ObjectContexts) {
		var ret []ServiceSupportExpectationObjectObjectContextsInner
		return ret
	}
	return o.ObjectContexts
}

// GetObjectContextsOk returns a tuple with the ObjectContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSupportExpectationObject) GetObjectContextsOk() ([]ServiceSupportExpectationObjectObjectContextsInner, bool) {
	if o == nil || IsNil(o.ObjectContexts) {
		return nil, false
	}
	return o.ObjectContexts, true
}

// HasObjectContexts returns a boolean if a field has been set.
func (o *ServiceSupportExpectationObject) HasObjectContexts() bool {
	if o != nil && !IsNil(o.ObjectContexts) {
		return true
	}

	return false
}

// SetObjectContexts gets a reference to the given []ServiceSupportExpectationObjectObjectContextsInner and assigns it to the ObjectContexts field.
func (o *ServiceSupportExpectationObject) SetObjectContexts(v []ServiceSupportExpectationObjectObjectContextsInner) {
	o.ObjectContexts = v
}

func (o ServiceSupportExpectationObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceSupportExpectationObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !IsNil(o.ObjectInstance) {
		toSerialize["objectInstance"] = o.ObjectInstance
	}
	if !IsNil(o.ObjectContexts) {
		toSerialize["objectContexts"] = o.ObjectContexts
	}
	return toSerialize, nil
}

type NullableServiceSupportExpectationObject struct {
	value *ServiceSupportExpectationObject
	isSet bool
}

func (v NullableServiceSupportExpectationObject) Get() *ServiceSupportExpectationObject {
	return v.value
}

func (v *NullableServiceSupportExpectationObject) Set(val *ServiceSupportExpectationObject) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSupportExpectationObject) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSupportExpectationObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSupportExpectationObject(val *ServiceSupportExpectationObject) *NullableServiceSupportExpectationObject {
	return &NullableServiceSupportExpectationObject{value: val, isSet: true}
}

func (v NullableServiceSupportExpectationObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSupportExpectationObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


/*
AI/ML NRM

OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AiMlNrm

import (
	"encoding/json"
)

// checks if the NtfSubscriptionControlSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NtfSubscriptionControlSingle{}

// NtfSubscriptionControlSingle struct for NtfSubscriptionControlSingle
type NtfSubscriptionControlSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *NtfSubscriptionControlSingleAllOfAttributes `json:"attributes,omitempty"`
	HeartbeatControl *HeartbeatControlSingle `json:"HeartbeatControl,omitempty"`
}

// NewNtfSubscriptionControlSingle instantiates a new NtfSubscriptionControlSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNtfSubscriptionControlSingle(id NullableString) *NtfSubscriptionControlSingle {
	this := NtfSubscriptionControlSingle{}
	this.Id = id
	return &this
}

// NewNtfSubscriptionControlSingleWithDefaults instantiates a new NtfSubscriptionControlSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNtfSubscriptionControlSingleWithDefaults() *NtfSubscriptionControlSingle {
	this := NtfSubscriptionControlSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NtfSubscriptionControlSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NtfSubscriptionControlSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *NtfSubscriptionControlSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *NtfSubscriptionControlSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *NtfSubscriptionControlSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *NtfSubscriptionControlSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingle) GetAttributes() NtfSubscriptionControlSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret NtfSubscriptionControlSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingle) GetAttributesOk() (*NtfSubscriptionControlSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given NtfSubscriptionControlSingleAllOfAttributes and assigns it to the Attributes field.
func (o *NtfSubscriptionControlSingle) SetAttributes(v NtfSubscriptionControlSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetHeartbeatControl returns the HeartbeatControl field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingle) GetHeartbeatControl() HeartbeatControlSingle {
	if o == nil || isNil(o.HeartbeatControl) {
		var ret HeartbeatControlSingle
		return ret
	}
	return *o.HeartbeatControl
}

// GetHeartbeatControlOk returns a tuple with the HeartbeatControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingle) GetHeartbeatControlOk() (*HeartbeatControlSingle, bool) {
	if o == nil || isNil(o.HeartbeatControl) {
		return nil, false
	}
	return o.HeartbeatControl, true
}

// HasHeartbeatControl returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingle) HasHeartbeatControl() bool {
	if o != nil && !isNil(o.HeartbeatControl) {
		return true
	}

	return false
}

// SetHeartbeatControl gets a reference to the given HeartbeatControlSingle and assigns it to the HeartbeatControl field.
func (o *NtfSubscriptionControlSingle) SetHeartbeatControl(v HeartbeatControlSingle) {
	o.HeartbeatControl = &v
}

func (o NtfSubscriptionControlSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NtfSubscriptionControlSingle) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.HeartbeatControl) {
		toSerialize["HeartbeatControl"] = o.HeartbeatControl
	}
	return toSerialize, nil
}

type NullableNtfSubscriptionControlSingle struct {
	value *NtfSubscriptionControlSingle
	isSet bool
}

func (v NullableNtfSubscriptionControlSingle) Get() *NtfSubscriptionControlSingle {
	return v.value
}

func (v *NullableNtfSubscriptionControlSingle) Set(val *NtfSubscriptionControlSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableNtfSubscriptionControlSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableNtfSubscriptionControlSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNtfSubscriptionControlSingle(val *NtfSubscriptionControlSingle) *NullableNtfSubscriptionControlSingle {
	return &NullableNtfSubscriptionControlSingle{value: val, isSet: true}
}

func (v NullableNtfSubscriptionControlSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNtfSubscriptionControlSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



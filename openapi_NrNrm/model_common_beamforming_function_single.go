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

// checks if the CommonBeamformingFunctionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonBeamformingFunctionSingle{}

// CommonBeamformingFunctionSingle struct for CommonBeamformingFunctionSingle
type CommonBeamformingFunctionSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *interface{} `json:"attributes,omitempty"`
	Beam []BeamSingle `json:"Beam,omitempty"`
}

// NewCommonBeamformingFunctionSingle instantiates a new CommonBeamformingFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonBeamformingFunctionSingle(id NullableString) *CommonBeamformingFunctionSingle {
	this := CommonBeamformingFunctionSingle{}
	this.Id = id
	return &this
}

// NewCommonBeamformingFunctionSingleWithDefaults instantiates a new CommonBeamformingFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonBeamformingFunctionSingleWithDefaults() *CommonBeamformingFunctionSingle {
	this := CommonBeamformingFunctionSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonBeamformingFunctionSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonBeamformingFunctionSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *CommonBeamformingFunctionSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *CommonBeamformingFunctionSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *CommonBeamformingFunctionSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *CommonBeamformingFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingle) GetAttributes() interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingle) GetAttributesOk() (*interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *CommonBeamformingFunctionSingle) SetAttributes(v interface{}) {
	o.Attributes = &v
}

// GetBeam returns the Beam field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingle) GetBeam() []BeamSingle {
	if o == nil || isNil(o.Beam) {
		var ret []BeamSingle
		return ret
	}
	return o.Beam
}

// GetBeamOk returns a tuple with the Beam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingle) GetBeamOk() ([]BeamSingle, bool) {
	if o == nil || isNil(o.Beam) {
		return nil, false
	}
	return o.Beam, true
}

// HasBeam returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingle) HasBeam() bool {
	if o != nil && !isNil(o.Beam) {
		return true
	}

	return false
}

// SetBeam gets a reference to the given []BeamSingle and assigns it to the Beam field.
func (o *CommonBeamformingFunctionSingle) SetBeam(v []BeamSingle) {
	o.Beam = v
}

func (o CommonBeamformingFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonBeamformingFunctionSingle) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.Beam) {
		toSerialize["Beam"] = o.Beam
	}
	return toSerialize, nil
}

type NullableCommonBeamformingFunctionSingle struct {
	value *CommonBeamformingFunctionSingle
	isSet bool
}

func (v NullableCommonBeamformingFunctionSingle) Get() *CommonBeamformingFunctionSingle {
	return v.value
}

func (v *NullableCommonBeamformingFunctionSingle) Set(val *CommonBeamformingFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonBeamformingFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonBeamformingFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonBeamformingFunctionSingle(val *CommonBeamformingFunctionSingle) *NullableCommonBeamformingFunctionSingle {
	return &NullableCommonBeamformingFunctionSingle{value: val, isSet: true}
}

func (v NullableCommonBeamformingFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonBeamformingFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



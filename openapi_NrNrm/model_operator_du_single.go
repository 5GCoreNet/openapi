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

// checks if the OperatorDuSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatorDuSingle{}

// OperatorDuSingle struct for OperatorDuSingle
type OperatorDuSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	GnbId *string `json:"gnbId,omitempty"`
	GnbIdLength *int32 `json:"gnbIdLength,omitempty"`
	EPF1C *EPF1CSingle `json:"EP_F1C,omitempty"`
	EPF1U []EPF1USingle `json:"EP_F1U,omitempty"`
}

// NewOperatorDuSingle instantiates a new OperatorDuSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorDuSingle(id NullableString) *OperatorDuSingle {
	this := OperatorDuSingle{}
	this.Id = id
	return &this
}

// NewOperatorDuSingleWithDefaults instantiates a new OperatorDuSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorDuSingleWithDefaults() *OperatorDuSingle {
	this := OperatorDuSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OperatorDuSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperatorDuSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *OperatorDuSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *OperatorDuSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDuSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *OperatorDuSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *OperatorDuSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *OperatorDuSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDuSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *OperatorDuSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *OperatorDuSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *OperatorDuSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDuSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *OperatorDuSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *OperatorDuSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetGnbId returns the GnbId field value if set, zero value otherwise.
func (o *OperatorDuSingle) GetGnbId() string {
	if o == nil || isNil(o.GnbId) {
		var ret string
		return ret
	}
	return *o.GnbId
}

// GetGnbIdOk returns a tuple with the GnbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDuSingle) GetGnbIdOk() (*string, bool) {
	if o == nil || isNil(o.GnbId) {
		return nil, false
	}
	return o.GnbId, true
}

// HasGnbId returns a boolean if a field has been set.
func (o *OperatorDuSingle) HasGnbId() bool {
	if o != nil && !isNil(o.GnbId) {
		return true
	}

	return false
}

// SetGnbId gets a reference to the given string and assigns it to the GnbId field.
func (o *OperatorDuSingle) SetGnbId(v string) {
	o.GnbId = &v
}

// GetGnbIdLength returns the GnbIdLength field value if set, zero value otherwise.
func (o *OperatorDuSingle) GetGnbIdLength() int32 {
	if o == nil || isNil(o.GnbIdLength) {
		var ret int32
		return ret
	}
	return *o.GnbIdLength
}

// GetGnbIdLengthOk returns a tuple with the GnbIdLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDuSingle) GetGnbIdLengthOk() (*int32, bool) {
	if o == nil || isNil(o.GnbIdLength) {
		return nil, false
	}
	return o.GnbIdLength, true
}

// HasGnbIdLength returns a boolean if a field has been set.
func (o *OperatorDuSingle) HasGnbIdLength() bool {
	if o != nil && !isNil(o.GnbIdLength) {
		return true
	}

	return false
}

// SetGnbIdLength gets a reference to the given int32 and assigns it to the GnbIdLength field.
func (o *OperatorDuSingle) SetGnbIdLength(v int32) {
	o.GnbIdLength = &v
}

// GetEPF1C returns the EPF1C field value if set, zero value otherwise.
func (o *OperatorDuSingle) GetEPF1C() EPF1CSingle {
	if o == nil || isNil(o.EPF1C) {
		var ret EPF1CSingle
		return ret
	}
	return *o.EPF1C
}

// GetEPF1COk returns a tuple with the EPF1C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDuSingle) GetEPF1COk() (*EPF1CSingle, bool) {
	if o == nil || isNil(o.EPF1C) {
		return nil, false
	}
	return o.EPF1C, true
}

// HasEPF1C returns a boolean if a field has been set.
func (o *OperatorDuSingle) HasEPF1C() bool {
	if o != nil && !isNil(o.EPF1C) {
		return true
	}

	return false
}

// SetEPF1C gets a reference to the given EPF1CSingle and assigns it to the EPF1C field.
func (o *OperatorDuSingle) SetEPF1C(v EPF1CSingle) {
	o.EPF1C = &v
}

// GetEPF1U returns the EPF1U field value if set, zero value otherwise.
func (o *OperatorDuSingle) GetEPF1U() []EPF1USingle {
	if o == nil || isNil(o.EPF1U) {
		var ret []EPF1USingle
		return ret
	}
	return o.EPF1U
}

// GetEPF1UOk returns a tuple with the EPF1U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDuSingle) GetEPF1UOk() ([]EPF1USingle, bool) {
	if o == nil || isNil(o.EPF1U) {
		return nil, false
	}
	return o.EPF1U, true
}

// HasEPF1U returns a boolean if a field has been set.
func (o *OperatorDuSingle) HasEPF1U() bool {
	if o != nil && !isNil(o.EPF1U) {
		return true
	}

	return false
}

// SetEPF1U gets a reference to the given []EPF1USingle and assigns it to the EPF1U field.
func (o *OperatorDuSingle) SetEPF1U(v []EPF1USingle) {
	o.EPF1U = v
}

func (o OperatorDuSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperatorDuSingle) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.GnbId) {
		toSerialize["gnbId"] = o.GnbId
	}
	if !isNil(o.GnbIdLength) {
		toSerialize["gnbIdLength"] = o.GnbIdLength
	}
	if !isNil(o.EPF1C) {
		toSerialize["EP_F1C"] = o.EPF1C
	}
	if !isNil(o.EPF1U) {
		toSerialize["EP_F1U"] = o.EPF1U
	}
	return toSerialize, nil
}

type NullableOperatorDuSingle struct {
	value *OperatorDuSingle
	isSet bool
}

func (v NullableOperatorDuSingle) Get() *OperatorDuSingle {
	return v.value
}

func (v *NullableOperatorDuSingle) Set(val *OperatorDuSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorDuSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorDuSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorDuSingle(val *OperatorDuSingle) *NullableOperatorDuSingle {
	return &NullableOperatorDuSingle{value: val, isSet: true}
}

func (v NullableOperatorDuSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorDuSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



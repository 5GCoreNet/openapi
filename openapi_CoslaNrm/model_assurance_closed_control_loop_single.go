/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// checks if the AssuranceClosedControlLoopSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssuranceClosedControlLoopSingle{}

// AssuranceClosedControlLoopSingle struct for AssuranceClosedControlLoopSingle
type AssuranceClosedControlLoopSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *AssuranceClosedControlLoopSingleAllOfAttributes `json:"attributes,omitempty"`
	AssuranceGoal []AssuranceGoalSingle `json:"AssuranceGoal,omitempty"`
}

// NewAssuranceClosedControlLoopSingle instantiates a new AssuranceClosedControlLoopSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssuranceClosedControlLoopSingle(id NullableString) *AssuranceClosedControlLoopSingle {
	this := AssuranceClosedControlLoopSingle{}
	this.Id = id
	return &this
}

// NewAssuranceClosedControlLoopSingleWithDefaults instantiates a new AssuranceClosedControlLoopSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssuranceClosedControlLoopSingleWithDefaults() *AssuranceClosedControlLoopSingle {
	this := AssuranceClosedControlLoopSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AssuranceClosedControlLoopSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssuranceClosedControlLoopSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *AssuranceClosedControlLoopSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *AssuranceClosedControlLoopSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *AssuranceClosedControlLoopSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *AssuranceClosedControlLoopSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingle) GetAttributes() AssuranceClosedControlLoopSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret AssuranceClosedControlLoopSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingle) GetAttributesOk() (*AssuranceClosedControlLoopSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AssuranceClosedControlLoopSingleAllOfAttributes and assigns it to the Attributes field.
func (o *AssuranceClosedControlLoopSingle) SetAttributes(v AssuranceClosedControlLoopSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetAssuranceGoal returns the AssuranceGoal field value if set, zero value otherwise.
func (o *AssuranceClosedControlLoopSingle) GetAssuranceGoal() []AssuranceGoalSingle {
	if o == nil || isNil(o.AssuranceGoal) {
		var ret []AssuranceGoalSingle
		return ret
	}
	return o.AssuranceGoal
}

// GetAssuranceGoalOk returns a tuple with the AssuranceGoal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceClosedControlLoopSingle) GetAssuranceGoalOk() ([]AssuranceGoalSingle, bool) {
	if o == nil || isNil(o.AssuranceGoal) {
		return nil, false
	}
	return o.AssuranceGoal, true
}

// HasAssuranceGoal returns a boolean if a field has been set.
func (o *AssuranceClosedControlLoopSingle) HasAssuranceGoal() bool {
	if o != nil && !isNil(o.AssuranceGoal) {
		return true
	}

	return false
}

// SetAssuranceGoal gets a reference to the given []AssuranceGoalSingle and assigns it to the AssuranceGoal field.
func (o *AssuranceClosedControlLoopSingle) SetAssuranceGoal(v []AssuranceGoalSingle) {
	o.AssuranceGoal = v
}

func (o AssuranceClosedControlLoopSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssuranceClosedControlLoopSingle) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.AssuranceGoal) {
		toSerialize["AssuranceGoal"] = o.AssuranceGoal
	}
	return toSerialize, nil
}

type NullableAssuranceClosedControlLoopSingle struct {
	value *AssuranceClosedControlLoopSingle
	isSet bool
}

func (v NullableAssuranceClosedControlLoopSingle) Get() *AssuranceClosedControlLoopSingle {
	return v.value
}

func (v *NullableAssuranceClosedControlLoopSingle) Set(val *AssuranceClosedControlLoopSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableAssuranceClosedControlLoopSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableAssuranceClosedControlLoopSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssuranceClosedControlLoopSingle(val *AssuranceClosedControlLoopSingle) *NullableAssuranceClosedControlLoopSingle {
	return &NullableAssuranceClosedControlLoopSingle{value: val, isSet: true}
}

func (v NullableAssuranceClosedControlLoopSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssuranceClosedControlLoopSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



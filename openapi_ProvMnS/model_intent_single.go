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

// checks if the IntentSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntentSingle{}

// IntentSingle struct for IntentSingle
type IntentSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	UserLabel *string `json:"userLabel,omitempty"`
	IntentExpectations []OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation `json:"intentExpectations,omitempty"`
	IntentContexts []IntentContext `json:"intentContexts,omitempty"`
	IntentFulfilmentInfo *FulfilmentInfo `json:"intentFulfilmentInfo,omitempty"`
}

// NewIntentSingle instantiates a new IntentSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntentSingle(id NullableString) *IntentSingle {
	this := IntentSingle{}
	this.Id = id
	return &this
}

// NewIntentSingleWithDefaults instantiates a new IntentSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntentSingleWithDefaults() *IntentSingle {
	this := IntentSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IntentSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntentSingle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *IntentSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *IntentSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *IntentSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *IntentSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *IntentSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
		return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *IntentSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *IntentSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *IntentSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
		return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *IntentSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *IntentSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *IntentSingle) GetUserLabel() string {
	if o == nil || isNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingle) GetUserLabelOk() (*string, bool) {
	if o == nil || isNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *IntentSingle) HasUserLabel() bool {
	if o != nil && !isNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *IntentSingle) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetIntentExpectations returns the IntentExpectations field value if set, zero value otherwise.
func (o *IntentSingle) GetIntentExpectations() []OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation {
	if o == nil || isNil(o.IntentExpectations) {
		var ret []OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation
		return ret
	}
	return o.IntentExpectations
}

// GetIntentExpectationsOk returns a tuple with the IntentExpectations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingle) GetIntentExpectationsOk() ([]OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation, bool) {
	if o == nil || isNil(o.IntentExpectations) {
		return nil, false
	}
	return o.IntentExpectations, true
}

// HasIntentExpectations returns a boolean if a field has been set.
func (o *IntentSingle) HasIntentExpectations() bool {
	if o != nil && !isNil(o.IntentExpectations) {
		return true
	}

	return false
}

// SetIntentExpectations gets a reference to the given []OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation and assigns it to the IntentExpectations field.
func (o *IntentSingle) SetIntentExpectations(v []OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation) {
	o.IntentExpectations = v
}

// GetIntentContexts returns the IntentContexts field value if set, zero value otherwise.
func (o *IntentSingle) GetIntentContexts() []IntentContext {
	if o == nil || isNil(o.IntentContexts) {
		var ret []IntentContext
		return ret
	}
	return o.IntentContexts
}

// GetIntentContextsOk returns a tuple with the IntentContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingle) GetIntentContextsOk() ([]IntentContext, bool) {
	if o == nil || isNil(o.IntentContexts) {
		return nil, false
	}
	return o.IntentContexts, true
}

// HasIntentContexts returns a boolean if a field has been set.
func (o *IntentSingle) HasIntentContexts() bool {
	if o != nil && !isNil(o.IntentContexts) {
		return true
	}

	return false
}

// SetIntentContexts gets a reference to the given []IntentContext and assigns it to the IntentContexts field.
func (o *IntentSingle) SetIntentContexts(v []IntentContext) {
	o.IntentContexts = v
}

// GetIntentFulfilmentInfo returns the IntentFulfilmentInfo field value if set, zero value otherwise.
func (o *IntentSingle) GetIntentFulfilmentInfo() FulfilmentInfo {
	if o == nil || isNil(o.IntentFulfilmentInfo) {
		var ret FulfilmentInfo
		return ret
	}
	return *o.IntentFulfilmentInfo
}

// GetIntentFulfilmentInfoOk returns a tuple with the IntentFulfilmentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingle) GetIntentFulfilmentInfoOk() (*FulfilmentInfo, bool) {
	if o == nil || isNil(o.IntentFulfilmentInfo) {
		return nil, false
	}
	return o.IntentFulfilmentInfo, true
}

// HasIntentFulfilmentInfo returns a boolean if a field has been set.
func (o *IntentSingle) HasIntentFulfilmentInfo() bool {
	if o != nil && !isNil(o.IntentFulfilmentInfo) {
		return true
	}

	return false
}

// SetIntentFulfilmentInfo gets a reference to the given FulfilmentInfo and assigns it to the IntentFulfilmentInfo field.
func (o *IntentSingle) SetIntentFulfilmentInfo(v FulfilmentInfo) {
	o.IntentFulfilmentInfo = &v
}

func (o IntentSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntentSingle) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.UserLabel) {
		toSerialize["userLabel"] = o.UserLabel
	}
	if !isNil(o.IntentExpectations) {
		toSerialize["intentExpectations"] = o.IntentExpectations
	}
	if !isNil(o.IntentContexts) {
		toSerialize["intentContexts"] = o.IntentContexts
	}
	if !isNil(o.IntentFulfilmentInfo) {
		toSerialize["intentFulfilmentInfo"] = o.IntentFulfilmentInfo
	}
	return toSerialize, nil
}

type NullableIntentSingle struct {
	value *IntentSingle
	isSet bool
}

func (v NullableIntentSingle) Get() *IntentSingle {
	return v.value
}

func (v *NullableIntentSingle) Set(val *IntentSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableIntentSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableIntentSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntentSingle(val *IntentSingle) *NullableIntentSingle {
	return &NullableIntentSingle{value: val, isSet: true}
}

func (v NullableIntentSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntentSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// checks if the ServiceTypeUnrelatedClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTypeUnrelatedClass{}

// ServiceTypeUnrelatedClass struct for ServiceTypeUnrelatedClass
type ServiceTypeUnrelatedClass struct {
	// LCS service type.
	ServiceType int32 `json:"serviceType"`
	AllowedGeographicArea []GeographicArea `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction *PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	CodeWordInd *CodeWordInd `json:"codeWordInd,omitempty"`
	ValidTimePeriod *ValidTimePeriod `json:"validTimePeriod,omitempty"`
	CodeWordList []string `json:"codeWordList,omitempty"`
}

// NewServiceTypeUnrelatedClass instantiates a new ServiceTypeUnrelatedClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTypeUnrelatedClass(serviceType int32) *ServiceTypeUnrelatedClass {
	this := ServiceTypeUnrelatedClass{}
	this.ServiceType = serviceType
	return &this
}

// NewServiceTypeUnrelatedClassWithDefaults instantiates a new ServiceTypeUnrelatedClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTypeUnrelatedClassWithDefaults() *ServiceTypeUnrelatedClass {
	this := ServiceTypeUnrelatedClass{}
	return &this
}

// GetServiceType returns the ServiceType field value
func (o *ServiceTypeUnrelatedClass) GetServiceType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *ServiceTypeUnrelatedClass) GetServiceTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value
func (o *ServiceTypeUnrelatedClass) SetServiceType(v int32) {
	o.ServiceType = v
}

// GetAllowedGeographicArea returns the AllowedGeographicArea field value if set, zero value otherwise.
func (o *ServiceTypeUnrelatedClass) GetAllowedGeographicArea() []GeographicArea {
	if o == nil || IsNil(o.AllowedGeographicArea) {
		var ret []GeographicArea
		return ret
	}
	return o.AllowedGeographicArea
}

// GetAllowedGeographicAreaOk returns a tuple with the AllowedGeographicArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTypeUnrelatedClass) GetAllowedGeographicAreaOk() ([]GeographicArea, bool) {
	if o == nil || IsNil(o.AllowedGeographicArea) {
		return nil, false
	}
	return o.AllowedGeographicArea, true
}

// HasAllowedGeographicArea returns a boolean if a field has been set.
func (o *ServiceTypeUnrelatedClass) HasAllowedGeographicArea() bool {
	if o != nil && !IsNil(o.AllowedGeographicArea) {
		return true
	}

	return false
}

// SetAllowedGeographicArea gets a reference to the given []GeographicArea and assigns it to the AllowedGeographicArea field.
func (o *ServiceTypeUnrelatedClass) SetAllowedGeographicArea(v []GeographicArea) {
	o.AllowedGeographicArea = v
}

// GetPrivacyCheckRelatedAction returns the PrivacyCheckRelatedAction field value if set, zero value otherwise.
func (o *ServiceTypeUnrelatedClass) GetPrivacyCheckRelatedAction() PrivacyCheckRelatedAction {
	if o == nil || IsNil(o.PrivacyCheckRelatedAction) {
		var ret PrivacyCheckRelatedAction
		return ret
	}
	return *o.PrivacyCheckRelatedAction
}

// GetPrivacyCheckRelatedActionOk returns a tuple with the PrivacyCheckRelatedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTypeUnrelatedClass) GetPrivacyCheckRelatedActionOk() (*PrivacyCheckRelatedAction, bool) {
	if o == nil || IsNil(o.PrivacyCheckRelatedAction) {
		return nil, false
	}
	return o.PrivacyCheckRelatedAction, true
}

// HasPrivacyCheckRelatedAction returns a boolean if a field has been set.
func (o *ServiceTypeUnrelatedClass) HasPrivacyCheckRelatedAction() bool {
	if o != nil && !IsNil(o.PrivacyCheckRelatedAction) {
		return true
	}

	return false
}

// SetPrivacyCheckRelatedAction gets a reference to the given PrivacyCheckRelatedAction and assigns it to the PrivacyCheckRelatedAction field.
func (o *ServiceTypeUnrelatedClass) SetPrivacyCheckRelatedAction(v PrivacyCheckRelatedAction) {
	o.PrivacyCheckRelatedAction = &v
}

// GetCodeWordInd returns the CodeWordInd field value if set, zero value otherwise.
func (o *ServiceTypeUnrelatedClass) GetCodeWordInd() CodeWordInd {
	if o == nil || IsNil(o.CodeWordInd) {
		var ret CodeWordInd
		return ret
	}
	return *o.CodeWordInd
}

// GetCodeWordIndOk returns a tuple with the CodeWordInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTypeUnrelatedClass) GetCodeWordIndOk() (*CodeWordInd, bool) {
	if o == nil || IsNil(o.CodeWordInd) {
		return nil, false
	}
	return o.CodeWordInd, true
}

// HasCodeWordInd returns a boolean if a field has been set.
func (o *ServiceTypeUnrelatedClass) HasCodeWordInd() bool {
	if o != nil && !IsNil(o.CodeWordInd) {
		return true
	}

	return false
}

// SetCodeWordInd gets a reference to the given CodeWordInd and assigns it to the CodeWordInd field.
func (o *ServiceTypeUnrelatedClass) SetCodeWordInd(v CodeWordInd) {
	o.CodeWordInd = &v
}

// GetValidTimePeriod returns the ValidTimePeriod field value if set, zero value otherwise.
func (o *ServiceTypeUnrelatedClass) GetValidTimePeriod() ValidTimePeriod {
	if o == nil || IsNil(o.ValidTimePeriod) {
		var ret ValidTimePeriod
		return ret
	}
	return *o.ValidTimePeriod
}

// GetValidTimePeriodOk returns a tuple with the ValidTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTypeUnrelatedClass) GetValidTimePeriodOk() (*ValidTimePeriod, bool) {
	if o == nil || IsNil(o.ValidTimePeriod) {
		return nil, false
	}
	return o.ValidTimePeriod, true
}

// HasValidTimePeriod returns a boolean if a field has been set.
func (o *ServiceTypeUnrelatedClass) HasValidTimePeriod() bool {
	if o != nil && !IsNil(o.ValidTimePeriod) {
		return true
	}

	return false
}

// SetValidTimePeriod gets a reference to the given ValidTimePeriod and assigns it to the ValidTimePeriod field.
func (o *ServiceTypeUnrelatedClass) SetValidTimePeriod(v ValidTimePeriod) {
	o.ValidTimePeriod = &v
}

// GetCodeWordList returns the CodeWordList field value if set, zero value otherwise.
func (o *ServiceTypeUnrelatedClass) GetCodeWordList() []string {
	if o == nil || IsNil(o.CodeWordList) {
		var ret []string
		return ret
	}
	return o.CodeWordList
}

// GetCodeWordListOk returns a tuple with the CodeWordList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTypeUnrelatedClass) GetCodeWordListOk() ([]string, bool) {
	if o == nil || IsNil(o.CodeWordList) {
		return nil, false
	}
	return o.CodeWordList, true
}

// HasCodeWordList returns a boolean if a field has been set.
func (o *ServiceTypeUnrelatedClass) HasCodeWordList() bool {
	if o != nil && !IsNil(o.CodeWordList) {
		return true
	}

	return false
}

// SetCodeWordList gets a reference to the given []string and assigns it to the CodeWordList field.
func (o *ServiceTypeUnrelatedClass) SetCodeWordList(v []string) {
	o.CodeWordList = v
}

func (o ServiceTypeUnrelatedClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTypeUnrelatedClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serviceType"] = o.ServiceType
	if !IsNil(o.AllowedGeographicArea) {
		toSerialize["allowedGeographicArea"] = o.AllowedGeographicArea
	}
	if !IsNil(o.PrivacyCheckRelatedAction) {
		toSerialize["privacyCheckRelatedAction"] = o.PrivacyCheckRelatedAction
	}
	if !IsNil(o.CodeWordInd) {
		toSerialize["codeWordInd"] = o.CodeWordInd
	}
	if !IsNil(o.ValidTimePeriod) {
		toSerialize["validTimePeriod"] = o.ValidTimePeriod
	}
	if !IsNil(o.CodeWordList) {
		toSerialize["codeWordList"] = o.CodeWordList
	}
	return toSerialize, nil
}

type NullableServiceTypeUnrelatedClass struct {
	value *ServiceTypeUnrelatedClass
	isSet bool
}

func (v NullableServiceTypeUnrelatedClass) Get() *ServiceTypeUnrelatedClass {
	return v.value
}

func (v *NullableServiceTypeUnrelatedClass) Set(val *ServiceTypeUnrelatedClass) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTypeUnrelatedClass) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTypeUnrelatedClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTypeUnrelatedClass(val *ServiceTypeUnrelatedClass) *NullableServiceTypeUnrelatedClass {
	return &NullableServiceTypeUnrelatedClass{value: val, isSet: true}
}

func (v NullableServiceTypeUnrelatedClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTypeUnrelatedClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


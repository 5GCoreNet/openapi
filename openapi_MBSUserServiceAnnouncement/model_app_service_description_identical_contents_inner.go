/*
MBS User Service Announcement Element units’ definition

MBS User Service Announcement Element units. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSUserServiceAnnouncement

import (
	"encoding/json"
)

// checks if the AppServiceDescriptionIdenticalContentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppServiceDescriptionIdenticalContentsInner{}

// AppServiceDescriptionIdenticalContentsInner struct for AppServiceDescriptionIdenticalContentsInner
type AppServiceDescriptionIdenticalContentsInner struct {
	UnicastAppService []ApplicationService `json:"unicastAppService,omitempty"`
}

// NewAppServiceDescriptionIdenticalContentsInner instantiates a new AppServiceDescriptionIdenticalContentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppServiceDescriptionIdenticalContentsInner() *AppServiceDescriptionIdenticalContentsInner {
	this := AppServiceDescriptionIdenticalContentsInner{}
	return &this
}

// NewAppServiceDescriptionIdenticalContentsInnerWithDefaults instantiates a new AppServiceDescriptionIdenticalContentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppServiceDescriptionIdenticalContentsInnerWithDefaults() *AppServiceDescriptionIdenticalContentsInner {
	this := AppServiceDescriptionIdenticalContentsInner{}
	return &this
}

// GetUnicastAppService returns the UnicastAppService field value if set, zero value otherwise.
func (o *AppServiceDescriptionIdenticalContentsInner) GetUnicastAppService() []ApplicationService {
	if o == nil || isNil(o.UnicastAppService) {
		var ret []ApplicationService
		return ret
	}
	return o.UnicastAppService
}

// GetUnicastAppServiceOk returns a tuple with the UnicastAppService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceDescriptionIdenticalContentsInner) GetUnicastAppServiceOk() ([]ApplicationService, bool) {
	if o == nil || isNil(o.UnicastAppService) {
		return nil, false
	}
	return o.UnicastAppService, true
}

// HasUnicastAppService returns a boolean if a field has been set.
func (o *AppServiceDescriptionIdenticalContentsInner) HasUnicastAppService() bool {
	if o != nil && !isNil(o.UnicastAppService) {
		return true
	}

	return false
}

// SetUnicastAppService gets a reference to the given []ApplicationService and assigns it to the UnicastAppService field.
func (o *AppServiceDescriptionIdenticalContentsInner) SetUnicastAppService(v []ApplicationService) {
	o.UnicastAppService = v
}

func (o AppServiceDescriptionIdenticalContentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppServiceDescriptionIdenticalContentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UnicastAppService) {
		toSerialize["unicastAppService"] = o.UnicastAppService
	}
	return toSerialize, nil
}

type NullableAppServiceDescriptionIdenticalContentsInner struct {
	value *AppServiceDescriptionIdenticalContentsInner
	isSet bool
}

func (v NullableAppServiceDescriptionIdenticalContentsInner) Get() *AppServiceDescriptionIdenticalContentsInner {
	return v.value
}

func (v *NullableAppServiceDescriptionIdenticalContentsInner) Set(val *AppServiceDescriptionIdenticalContentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppServiceDescriptionIdenticalContentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppServiceDescriptionIdenticalContentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppServiceDescriptionIdenticalContentsInner(val *AppServiceDescriptionIdenticalContentsInner) *NullableAppServiceDescriptionIdenticalContentsInner {
	return &NullableAppServiceDescriptionIdenticalContentsInner{value: val, isSet: true}
}

func (v NullableAppServiceDescriptionIdenticalContentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppServiceDescriptionIdenticalContentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



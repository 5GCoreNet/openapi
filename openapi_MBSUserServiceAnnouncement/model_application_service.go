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

// checks if the ApplicationService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationService{}

// ApplicationService struct for ApplicationService
type ApplicationService struct {
	BasePattern string `json:"basePattern"`
}

// NewApplicationService instantiates a new ApplicationService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationService(basePattern string) *ApplicationService {
	this := ApplicationService{}
	this.BasePattern = basePattern
	return &this
}

// NewApplicationServiceWithDefaults instantiates a new ApplicationService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationServiceWithDefaults() *ApplicationService {
	this := ApplicationService{}
	return &this
}

// GetBasePattern returns the BasePattern field value
func (o *ApplicationService) GetBasePattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasePattern
}

// GetBasePatternOk returns a tuple with the BasePattern field value
// and a boolean to check if the value has been set.
func (o *ApplicationService) GetBasePatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasePattern, true
}

// SetBasePattern sets field value
func (o *ApplicationService) SetBasePattern(v string) {
	o.BasePattern = v
}

func (o ApplicationService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["basePattern"] = o.BasePattern
	return toSerialize, nil
}

type NullableApplicationService struct {
	value *ApplicationService
	isSet bool
}

func (v NullableApplicationService) Get() *ApplicationService {
	return v.value
}

func (v *NullableApplicationService) Set(val *ApplicationService) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationService) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationService(val *ApplicationService) *NullableApplicationService {
	return &NullableApplicationService{value: val, isSet: true}
}

func (v NullableApplicationService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



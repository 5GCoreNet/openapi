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

// checks if the UserServiceDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserServiceDescription{}

// UserServiceDescription struct for UserServiceDescription
type UserServiceDescription struct {
	Name []string `json:"name,omitempty"`
	ServiceLanguage []string `json:"serviceLanguage,omitempty"`
	ServiceId string `json:"serviceId"`
	DistributionSessionDescription *DistributionSessionDescription `json:"distributionSessionDescription,omitempty"`
	AppServiceDescription *AppServiceDescription `json:"appServiceDescription,omitempty"`
	ScheduleDescription []ServiceSchedule `json:"scheduleDescription,omitempty"`
	AvailabilityInfo []AvailabilityInformationBinding `json:"availabilityInfo,omitempty"`
}

// NewUserServiceDescription instantiates a new UserServiceDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserServiceDescription(serviceId string) *UserServiceDescription {
	this := UserServiceDescription{}
	this.ServiceId = serviceId
	return &this
}

// NewUserServiceDescriptionWithDefaults instantiates a new UserServiceDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserServiceDescriptionWithDefaults() *UserServiceDescription {
	this := UserServiceDescription{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserServiceDescription) GetName() []string {
	if o == nil || IsNil(o.Name) {
		var ret []string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceDescription) GetNameOk() ([]string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserServiceDescription) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *UserServiceDescription) SetName(v []string) {
	o.Name = v
}

// GetServiceLanguage returns the ServiceLanguage field value if set, zero value otherwise.
func (o *UserServiceDescription) GetServiceLanguage() []string {
	if o == nil || IsNil(o.ServiceLanguage) {
		var ret []string
		return ret
	}
	return o.ServiceLanguage
}

// GetServiceLanguageOk returns a tuple with the ServiceLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceDescription) GetServiceLanguageOk() ([]string, bool) {
	if o == nil || IsNil(o.ServiceLanguage) {
		return nil, false
	}
	return o.ServiceLanguage, true
}

// HasServiceLanguage returns a boolean if a field has been set.
func (o *UserServiceDescription) HasServiceLanguage() bool {
	if o != nil && !IsNil(o.ServiceLanguage) {
		return true
	}

	return false
}

// SetServiceLanguage gets a reference to the given []string and assigns it to the ServiceLanguage field.
func (o *UserServiceDescription) SetServiceLanguage(v []string) {
	o.ServiceLanguage = v
}

// GetServiceId returns the ServiceId field value
func (o *UserServiceDescription) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UserServiceDescription) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UserServiceDescription) SetServiceId(v string) {
	o.ServiceId = v
}

// GetDistributionSessionDescription returns the DistributionSessionDescription field value if set, zero value otherwise.
func (o *UserServiceDescription) GetDistributionSessionDescription() DistributionSessionDescription {
	if o == nil || IsNil(o.DistributionSessionDescription) {
		var ret DistributionSessionDescription
		return ret
	}
	return *o.DistributionSessionDescription
}

// GetDistributionSessionDescriptionOk returns a tuple with the DistributionSessionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceDescription) GetDistributionSessionDescriptionOk() (*DistributionSessionDescription, bool) {
	if o == nil || IsNil(o.DistributionSessionDescription) {
		return nil, false
	}
	return o.DistributionSessionDescription, true
}

// HasDistributionSessionDescription returns a boolean if a field has been set.
func (o *UserServiceDescription) HasDistributionSessionDescription() bool {
	if o != nil && !IsNil(o.DistributionSessionDescription) {
		return true
	}

	return false
}

// SetDistributionSessionDescription gets a reference to the given DistributionSessionDescription and assigns it to the DistributionSessionDescription field.
func (o *UserServiceDescription) SetDistributionSessionDescription(v DistributionSessionDescription) {
	o.DistributionSessionDescription = &v
}

// GetAppServiceDescription returns the AppServiceDescription field value if set, zero value otherwise.
func (o *UserServiceDescription) GetAppServiceDescription() AppServiceDescription {
	if o == nil || IsNil(o.AppServiceDescription) {
		var ret AppServiceDescription
		return ret
	}
	return *o.AppServiceDescription
}

// GetAppServiceDescriptionOk returns a tuple with the AppServiceDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceDescription) GetAppServiceDescriptionOk() (*AppServiceDescription, bool) {
	if o == nil || IsNil(o.AppServiceDescription) {
		return nil, false
	}
	return o.AppServiceDescription, true
}

// HasAppServiceDescription returns a boolean if a field has been set.
func (o *UserServiceDescription) HasAppServiceDescription() bool {
	if o != nil && !IsNil(o.AppServiceDescription) {
		return true
	}

	return false
}

// SetAppServiceDescription gets a reference to the given AppServiceDescription and assigns it to the AppServiceDescription field.
func (o *UserServiceDescription) SetAppServiceDescription(v AppServiceDescription) {
	o.AppServiceDescription = &v
}

// GetScheduleDescription returns the ScheduleDescription field value if set, zero value otherwise.
func (o *UserServiceDescription) GetScheduleDescription() []ServiceSchedule {
	if o == nil || IsNil(o.ScheduleDescription) {
		var ret []ServiceSchedule
		return ret
	}
	return o.ScheduleDescription
}

// GetScheduleDescriptionOk returns a tuple with the ScheduleDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceDescription) GetScheduleDescriptionOk() ([]ServiceSchedule, bool) {
	if o == nil || IsNil(o.ScheduleDescription) {
		return nil, false
	}
	return o.ScheduleDescription, true
}

// HasScheduleDescription returns a boolean if a field has been set.
func (o *UserServiceDescription) HasScheduleDescription() bool {
	if o != nil && !IsNil(o.ScheduleDescription) {
		return true
	}

	return false
}

// SetScheduleDescription gets a reference to the given []ServiceSchedule and assigns it to the ScheduleDescription field.
func (o *UserServiceDescription) SetScheduleDescription(v []ServiceSchedule) {
	o.ScheduleDescription = v
}

// GetAvailabilityInfo returns the AvailabilityInfo field value if set, zero value otherwise.
func (o *UserServiceDescription) GetAvailabilityInfo() []AvailabilityInformationBinding {
	if o == nil || IsNil(o.AvailabilityInfo) {
		var ret []AvailabilityInformationBinding
		return ret
	}
	return o.AvailabilityInfo
}

// GetAvailabilityInfoOk returns a tuple with the AvailabilityInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceDescription) GetAvailabilityInfoOk() ([]AvailabilityInformationBinding, bool) {
	if o == nil || IsNil(o.AvailabilityInfo) {
		return nil, false
	}
	return o.AvailabilityInfo, true
}

// HasAvailabilityInfo returns a boolean if a field has been set.
func (o *UserServiceDescription) HasAvailabilityInfo() bool {
	if o != nil && !IsNil(o.AvailabilityInfo) {
		return true
	}

	return false
}

// SetAvailabilityInfo gets a reference to the given []AvailabilityInformationBinding and assigns it to the AvailabilityInfo field.
func (o *UserServiceDescription) SetAvailabilityInfo(v []AvailabilityInformationBinding) {
	o.AvailabilityInfo = v
}

func (o UserServiceDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserServiceDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ServiceLanguage) {
		toSerialize["serviceLanguage"] = o.ServiceLanguage
	}
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.DistributionSessionDescription) {
		toSerialize["distributionSessionDescription"] = o.DistributionSessionDescription
	}
	if !IsNil(o.AppServiceDescription) {
		toSerialize["appServiceDescription"] = o.AppServiceDescription
	}
	if !IsNil(o.ScheduleDescription) {
		toSerialize["scheduleDescription"] = o.ScheduleDescription
	}
	if !IsNil(o.AvailabilityInfo) {
		toSerialize["availabilityInfo"] = o.AvailabilityInfo
	}
	return toSerialize, nil
}

type NullableUserServiceDescription struct {
	value *UserServiceDescription
	isSet bool
}

func (v NullableUserServiceDescription) Get() *UserServiceDescription {
	return v.value
}

func (v *NullableUserServiceDescription) Set(val *UserServiceDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableUserServiceDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableUserServiceDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserServiceDescription(val *UserServiceDescription) *NullableUserServiceDescription {
	return &NullableUserServiceDescription{value: val, isSet: true}
}

func (v NullableUserServiceDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserServiceDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



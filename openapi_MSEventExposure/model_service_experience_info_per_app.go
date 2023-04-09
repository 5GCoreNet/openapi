/*
3gpp-ms-event-exposure

API for Media Streaming Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSEventExposure

import (
	"encoding/json"
)

// checks if the ServiceExperienceInfoPerApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceExperienceInfoPerApp{}

// ServiceExperienceInfoPerApp Contains service experience information associated with an application.
type ServiceExperienceInfoPerApp struct {
	// String providing an application identifier.
	AppId          *string                        `json:"appId,omitempty"`
	AppServerIns   *AddrFqdn                      `json:"appServerIns,omitempty"`
	SvcExpPerFlows []ServiceExperienceInfoPerFlow `json:"svcExpPerFlows"`
	Gpsis          []string                       `json:"gpsis,omitempty"`
	Supis          []string                       `json:"supis,omitempty"`
}

// NewServiceExperienceInfoPerApp instantiates a new ServiceExperienceInfoPerApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceExperienceInfoPerApp(svcExpPerFlows []ServiceExperienceInfoPerFlow) *ServiceExperienceInfoPerApp {
	this := ServiceExperienceInfoPerApp{}
	this.SvcExpPerFlows = svcExpPerFlows
	return &this
}

// NewServiceExperienceInfoPerAppWithDefaults instantiates a new ServiceExperienceInfoPerApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceExperienceInfoPerAppWithDefaults() *ServiceExperienceInfoPerApp {
	this := ServiceExperienceInfoPerApp{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ServiceExperienceInfoPerApp) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfoPerApp) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ServiceExperienceInfoPerApp) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *ServiceExperienceInfoPerApp) SetAppId(v string) {
	o.AppId = &v
}

// GetAppServerIns returns the AppServerIns field value if set, zero value otherwise.
func (o *ServiceExperienceInfoPerApp) GetAppServerIns() AddrFqdn {
	if o == nil || IsNil(o.AppServerIns) {
		var ret AddrFqdn
		return ret
	}
	return *o.AppServerIns
}

// GetAppServerInsOk returns a tuple with the AppServerIns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfoPerApp) GetAppServerInsOk() (*AddrFqdn, bool) {
	if o == nil || IsNil(o.AppServerIns) {
		return nil, false
	}
	return o.AppServerIns, true
}

// HasAppServerIns returns a boolean if a field has been set.
func (o *ServiceExperienceInfoPerApp) HasAppServerIns() bool {
	if o != nil && !IsNil(o.AppServerIns) {
		return true
	}

	return false
}

// SetAppServerIns gets a reference to the given AddrFqdn and assigns it to the AppServerIns field.
func (o *ServiceExperienceInfoPerApp) SetAppServerIns(v AddrFqdn) {
	o.AppServerIns = &v
}

// GetSvcExpPerFlows returns the SvcExpPerFlows field value
func (o *ServiceExperienceInfoPerApp) GetSvcExpPerFlows() []ServiceExperienceInfoPerFlow {
	if o == nil {
		var ret []ServiceExperienceInfoPerFlow
		return ret
	}

	return o.SvcExpPerFlows
}

// GetSvcExpPerFlowsOk returns a tuple with the SvcExpPerFlows field value
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfoPerApp) GetSvcExpPerFlowsOk() ([]ServiceExperienceInfoPerFlow, bool) {
	if o == nil {
		return nil, false
	}
	return o.SvcExpPerFlows, true
}

// SetSvcExpPerFlows sets field value
func (o *ServiceExperienceInfoPerApp) SetSvcExpPerFlows(v []ServiceExperienceInfoPerFlow) {
	o.SvcExpPerFlows = v
}

// GetGpsis returns the Gpsis field value if set, zero value otherwise.
func (o *ServiceExperienceInfoPerApp) GetGpsis() []string {
	if o == nil || IsNil(o.Gpsis) {
		var ret []string
		return ret
	}
	return o.Gpsis
}

// GetGpsisOk returns a tuple with the Gpsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfoPerApp) GetGpsisOk() ([]string, bool) {
	if o == nil || IsNil(o.Gpsis) {
		return nil, false
	}
	return o.Gpsis, true
}

// HasGpsis returns a boolean if a field has been set.
func (o *ServiceExperienceInfoPerApp) HasGpsis() bool {
	if o != nil && !IsNil(o.Gpsis) {
		return true
	}

	return false
}

// SetGpsis gets a reference to the given []string and assigns it to the Gpsis field.
func (o *ServiceExperienceInfoPerApp) SetGpsis(v []string) {
	o.Gpsis = v
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *ServiceExperienceInfoPerApp) GetSupis() []string {
	if o == nil || IsNil(o.Supis) {
		var ret []string
		return ret
	}
	return o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfoPerApp) GetSupisOk() ([]string, bool) {
	if o == nil || IsNil(o.Supis) {
		return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *ServiceExperienceInfoPerApp) HasSupis() bool {
	if o != nil && !IsNil(o.Supis) {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []string and assigns it to the Supis field.
func (o *ServiceExperienceInfoPerApp) SetSupis(v []string) {
	o.Supis = v
}

func (o ServiceExperienceInfoPerApp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceExperienceInfoPerApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.AppServerIns) {
		toSerialize["appServerIns"] = o.AppServerIns
	}
	toSerialize["svcExpPerFlows"] = o.SvcExpPerFlows
	if !IsNil(o.Gpsis) {
		toSerialize["gpsis"] = o.Gpsis
	}
	if !IsNil(o.Supis) {
		toSerialize["supis"] = o.Supis
	}
	return toSerialize, nil
}

type NullableServiceExperienceInfoPerApp struct {
	value *ServiceExperienceInfoPerApp
	isSet bool
}

func (v NullableServiceExperienceInfoPerApp) Get() *ServiceExperienceInfoPerApp {
	return v.value
}

func (v *NullableServiceExperienceInfoPerApp) Set(val *ServiceExperienceInfoPerApp) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceExperienceInfoPerApp) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceExperienceInfoPerApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceExperienceInfoPerApp(val *ServiceExperienceInfoPerApp) *NullableServiceExperienceInfoPerApp {
	return &NullableServiceExperienceInfoPerApp{value: val, isSet: true}
}

func (v NullableServiceExperienceInfoPerApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceExperienceInfoPerApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

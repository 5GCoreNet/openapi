/*
VAE_ApplicationRequirement

API for VAE Application Requirement Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_VAE_ApplicationRequirement

import (
	"encoding/json"
)

// checks if the AppReqNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppReqNotification{}

// AppReqNotification Represents a notificaton of the result of the network resource adaptation corresponding to the V2X application requirement. 
type AppReqNotification struct {
	// String providing an URI formatted according to RFC 3986.
	ResourceUri string `json:"resourceUri"`
	Result ReservationResult `json:"result"`
}

// NewAppReqNotification instantiates a new AppReqNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppReqNotification(resourceUri string, result ReservationResult) *AppReqNotification {
	this := AppReqNotification{}
	this.ResourceUri = resourceUri
	this.Result = result
	return &this
}

// NewAppReqNotificationWithDefaults instantiates a new AppReqNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppReqNotificationWithDefaults() *AppReqNotification {
	this := AppReqNotification{}
	return &this
}

// GetResourceUri returns the ResourceUri field value
func (o *AppReqNotification) GetResourceUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceUri
}

// GetResourceUriOk returns a tuple with the ResourceUri field value
// and a boolean to check if the value has been set.
func (o *AppReqNotification) GetResourceUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceUri, true
}

// SetResourceUri sets field value
func (o *AppReqNotification) SetResourceUri(v string) {
	o.ResourceUri = v
}

// GetResult returns the Result field value
func (o *AppReqNotification) GetResult() ReservationResult {
	if o == nil {
		var ret ReservationResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *AppReqNotification) GetResultOk() (*ReservationResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *AppReqNotification) SetResult(v ReservationResult) {
	o.Result = v
}

func (o AppReqNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppReqNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceUri"] = o.ResourceUri
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableAppReqNotification struct {
	value *AppReqNotification
	isSet bool
}

func (v NullableAppReqNotification) Get() *AppReqNotification {
	return v.value
}

func (v *NullableAppReqNotification) Set(val *AppReqNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAppReqNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAppReqNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppReqNotification(val *AppReqNotification) *NullableAppReqNotification {
	return &NullableAppReqNotification{value: val, isSet: true}
}

func (v NullableAppReqNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppReqNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



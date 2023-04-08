/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the CoreNetworkServiceAuthorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreNetworkServiceAuthorization{}

// CoreNetworkServiceAuthorization Core Network Service Authorization 
type CoreNetworkServiceAuthorization struct {
	SubscribedMediaProfileId *int32 `json:"subscribedMediaProfileId,omitempty"`
}

// NewCoreNetworkServiceAuthorization instantiates a new CoreNetworkServiceAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreNetworkServiceAuthorization() *CoreNetworkServiceAuthorization {
	this := CoreNetworkServiceAuthorization{}
	return &this
}

// NewCoreNetworkServiceAuthorizationWithDefaults instantiates a new CoreNetworkServiceAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreNetworkServiceAuthorizationWithDefaults() *CoreNetworkServiceAuthorization {
	this := CoreNetworkServiceAuthorization{}
	return &this
}

// GetSubscribedMediaProfileId returns the SubscribedMediaProfileId field value if set, zero value otherwise.
func (o *CoreNetworkServiceAuthorization) GetSubscribedMediaProfileId() int32 {
	if o == nil || isNil(o.SubscribedMediaProfileId) {
		var ret int32
		return ret
	}
	return *o.SubscribedMediaProfileId
}

// GetSubscribedMediaProfileIdOk returns a tuple with the SubscribedMediaProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreNetworkServiceAuthorization) GetSubscribedMediaProfileIdOk() (*int32, bool) {
	if o == nil || isNil(o.SubscribedMediaProfileId) {
		return nil, false
	}
	return o.SubscribedMediaProfileId, true
}

// HasSubscribedMediaProfileId returns a boolean if a field has been set.
func (o *CoreNetworkServiceAuthorization) HasSubscribedMediaProfileId() bool {
	if o != nil && !isNil(o.SubscribedMediaProfileId) {
		return true
	}

	return false
}

// SetSubscribedMediaProfileId gets a reference to the given int32 and assigns it to the SubscribedMediaProfileId field.
func (o *CoreNetworkServiceAuthorization) SetSubscribedMediaProfileId(v int32) {
	o.SubscribedMediaProfileId = &v
}

func (o CoreNetworkServiceAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreNetworkServiceAuthorization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SubscribedMediaProfileId) {
		toSerialize["subscribedMediaProfileId"] = o.SubscribedMediaProfileId
	}
	return toSerialize, nil
}

type NullableCoreNetworkServiceAuthorization struct {
	value *CoreNetworkServiceAuthorization
	isSet bool
}

func (v NullableCoreNetworkServiceAuthorization) Get() *CoreNetworkServiceAuthorization {
	return v.value
}

func (v *NullableCoreNetworkServiceAuthorization) Set(val *CoreNetworkServiceAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreNetworkServiceAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreNetworkServiceAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreNetworkServiceAuthorization(val *CoreNetworkServiceAuthorization) *NullableCoreNetworkServiceAuthorization {
	return &NullableCoreNetworkServiceAuthorization{value: val, isSet: true}
}

func (v NullableCoreNetworkServiceAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreNetworkServiceAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



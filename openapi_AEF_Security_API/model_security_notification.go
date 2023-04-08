/*
AEF_Security_API

API for AEF security management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AEF_Security_API

import (
	"encoding/json"
)

// checks if the SecurityNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityNotification{}

// SecurityNotification Represents the revoked authorization notification details.
type SecurityNotification struct {
	// String identifying the API invoker assigned by the CAPIF core function.
	ApiInvokerId string `json:"apiInvokerId"`
	// String identifying the AEF.
	AefId *string `json:"aefId,omitempty"`
	// Identifier of the service API
	ApiIds []string `json:"apiIds"`
	Cause Cause `json:"cause"`
}

// NewSecurityNotification instantiates a new SecurityNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityNotification(apiInvokerId string, apiIds []string, cause Cause) *SecurityNotification {
	this := SecurityNotification{}
	this.ApiInvokerId = apiInvokerId
	this.ApiIds = apiIds
	this.Cause = cause
	return &this
}

// NewSecurityNotificationWithDefaults instantiates a new SecurityNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityNotificationWithDefaults() *SecurityNotification {
	this := SecurityNotification{}
	return &this
}

// GetApiInvokerId returns the ApiInvokerId field value
func (o *SecurityNotification) GetApiInvokerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiInvokerId
}

// GetApiInvokerIdOk returns a tuple with the ApiInvokerId field value
// and a boolean to check if the value has been set.
func (o *SecurityNotification) GetApiInvokerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiInvokerId, true
}

// SetApiInvokerId sets field value
func (o *SecurityNotification) SetApiInvokerId(v string) {
	o.ApiInvokerId = v
}

// GetAefId returns the AefId field value if set, zero value otherwise.
func (o *SecurityNotification) GetAefId() string {
	if o == nil || isNil(o.AefId) {
		var ret string
		return ret
	}
	return *o.AefId
}

// GetAefIdOk returns a tuple with the AefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityNotification) GetAefIdOk() (*string, bool) {
	if o == nil || isNil(o.AefId) {
		return nil, false
	}
	return o.AefId, true
}

// HasAefId returns a boolean if a field has been set.
func (o *SecurityNotification) HasAefId() bool {
	if o != nil && !isNil(o.AefId) {
		return true
	}

	return false
}

// SetAefId gets a reference to the given string and assigns it to the AefId field.
func (o *SecurityNotification) SetAefId(v string) {
	o.AefId = &v
}

// GetApiIds returns the ApiIds field value
func (o *SecurityNotification) GetApiIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ApiIds
}

// GetApiIdsOk returns a tuple with the ApiIds field value
// and a boolean to check if the value has been set.
func (o *SecurityNotification) GetApiIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiIds, true
}

// SetApiIds sets field value
func (o *SecurityNotification) SetApiIds(v []string) {
	o.ApiIds = v
}

// GetCause returns the Cause field value
func (o *SecurityNotification) GetCause() Cause {
	if o == nil {
		var ret Cause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *SecurityNotification) GetCauseOk() (*Cause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *SecurityNotification) SetCause(v Cause) {
	o.Cause = v
}

func (o SecurityNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiInvokerId"] = o.ApiInvokerId
	if !isNil(o.AefId) {
		toSerialize["aefId"] = o.AefId
	}
	toSerialize["apiIds"] = o.ApiIds
	toSerialize["cause"] = o.Cause
	return toSerialize, nil
}

type NullableSecurityNotification struct {
	value *SecurityNotification
	isSet bool
}

func (v NullableSecurityNotification) Get() *SecurityNotification {
	return v.value
}

func (v *NullableSecurityNotification) Set(val *SecurityNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityNotification(val *SecurityNotification) *NullableSecurityNotification {
	return &NullableSecurityNotification{value: val, isSet: true}
}

func (v NullableSecurityNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



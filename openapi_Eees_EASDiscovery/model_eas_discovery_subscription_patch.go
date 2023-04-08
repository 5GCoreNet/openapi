/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EASDiscovery

import (
	"encoding/json"
	"time"
)

// checks if the EasDiscoverySubscriptionPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EasDiscoverySubscriptionPatch{}

// EasDiscoverySubscriptionPatch Represents an Individual EAS Discovery Subscription resource.
type EasDiscoverySubscriptionPatch struct {
	EasDiscoveryFilter *EasDiscoveryFilter `json:"easDiscoveryFilter,omitempty"`
	EasDynInfoFilter *EasDynamicInfoFilter `json:"easDynInfoFilter,omitempty"`
	// Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC.
	EasSvcContinuity []ACRScenario `json:"easSvcContinuity,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	ExpTime *time.Time `json:"expTime,omitempty"`
}

// NewEasDiscoverySubscriptionPatch instantiates a new EasDiscoverySubscriptionPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasDiscoverySubscriptionPatch() *EasDiscoverySubscriptionPatch {
	this := EasDiscoverySubscriptionPatch{}
	return &this
}

// NewEasDiscoverySubscriptionPatchWithDefaults instantiates a new EasDiscoverySubscriptionPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasDiscoverySubscriptionPatchWithDefaults() *EasDiscoverySubscriptionPatch {
	this := EasDiscoverySubscriptionPatch{}
	return &this
}

// GetEasDiscoveryFilter returns the EasDiscoveryFilter field value if set, zero value otherwise.
func (o *EasDiscoverySubscriptionPatch) GetEasDiscoveryFilter() EasDiscoveryFilter {
	if o == nil || isNil(o.EasDiscoveryFilter) {
		var ret EasDiscoveryFilter
		return ret
	}
	return *o.EasDiscoveryFilter
}

// GetEasDiscoveryFilterOk returns a tuple with the EasDiscoveryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscriptionPatch) GetEasDiscoveryFilterOk() (*EasDiscoveryFilter, bool) {
	if o == nil || isNil(o.EasDiscoveryFilter) {
		return nil, false
	}
	return o.EasDiscoveryFilter, true
}

// HasEasDiscoveryFilter returns a boolean if a field has been set.
func (o *EasDiscoverySubscriptionPatch) HasEasDiscoveryFilter() bool {
	if o != nil && !isNil(o.EasDiscoveryFilter) {
		return true
	}

	return false
}

// SetEasDiscoveryFilter gets a reference to the given EasDiscoveryFilter and assigns it to the EasDiscoveryFilter field.
func (o *EasDiscoverySubscriptionPatch) SetEasDiscoveryFilter(v EasDiscoveryFilter) {
	o.EasDiscoveryFilter = &v
}

// GetEasDynInfoFilter returns the EasDynInfoFilter field value if set, zero value otherwise.
func (o *EasDiscoverySubscriptionPatch) GetEasDynInfoFilter() EasDynamicInfoFilter {
	if o == nil || isNil(o.EasDynInfoFilter) {
		var ret EasDynamicInfoFilter
		return ret
	}
	return *o.EasDynInfoFilter
}

// GetEasDynInfoFilterOk returns a tuple with the EasDynInfoFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscriptionPatch) GetEasDynInfoFilterOk() (*EasDynamicInfoFilter, bool) {
	if o == nil || isNil(o.EasDynInfoFilter) {
		return nil, false
	}
	return o.EasDynInfoFilter, true
}

// HasEasDynInfoFilter returns a boolean if a field has been set.
func (o *EasDiscoverySubscriptionPatch) HasEasDynInfoFilter() bool {
	if o != nil && !isNil(o.EasDynInfoFilter) {
		return true
	}

	return false
}

// SetEasDynInfoFilter gets a reference to the given EasDynamicInfoFilter and assigns it to the EasDynInfoFilter field.
func (o *EasDiscoverySubscriptionPatch) SetEasDynInfoFilter(v EasDynamicInfoFilter) {
	o.EasDynInfoFilter = &v
}

// GetEasSvcContinuity returns the EasSvcContinuity field value if set, zero value otherwise.
func (o *EasDiscoverySubscriptionPatch) GetEasSvcContinuity() []ACRScenario {
	if o == nil || isNil(o.EasSvcContinuity) {
		var ret []ACRScenario
		return ret
	}
	return o.EasSvcContinuity
}

// GetEasSvcContinuityOk returns a tuple with the EasSvcContinuity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscriptionPatch) GetEasSvcContinuityOk() ([]ACRScenario, bool) {
	if o == nil || isNil(o.EasSvcContinuity) {
		return nil, false
	}
	return o.EasSvcContinuity, true
}

// HasEasSvcContinuity returns a boolean if a field has been set.
func (o *EasDiscoverySubscriptionPatch) HasEasSvcContinuity() bool {
	if o != nil && !isNil(o.EasSvcContinuity) {
		return true
	}

	return false
}

// SetEasSvcContinuity gets a reference to the given []ACRScenario and assigns it to the EasSvcContinuity field.
func (o *EasDiscoverySubscriptionPatch) SetEasSvcContinuity(v []ACRScenario) {
	o.EasSvcContinuity = v
}

// GetExpTime returns the ExpTime field value if set, zero value otherwise.
func (o *EasDiscoverySubscriptionPatch) GetExpTime() time.Time {
	if o == nil || isNil(o.ExpTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpTime
}

// GetExpTimeOk returns a tuple with the ExpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscriptionPatch) GetExpTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpTime) {
		return nil, false
	}
	return o.ExpTime, true
}

// HasExpTime returns a boolean if a field has been set.
func (o *EasDiscoverySubscriptionPatch) HasExpTime() bool {
	if o != nil && !isNil(o.ExpTime) {
		return true
	}

	return false
}

// SetExpTime gets a reference to the given time.Time and assigns it to the ExpTime field.
func (o *EasDiscoverySubscriptionPatch) SetExpTime(v time.Time) {
	o.ExpTime = &v
}

func (o EasDiscoverySubscriptionPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EasDiscoverySubscriptionPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EasDiscoveryFilter) {
		toSerialize["easDiscoveryFilter"] = o.EasDiscoveryFilter
	}
	if !isNil(o.EasDynInfoFilter) {
		toSerialize["easDynInfoFilter"] = o.EasDynInfoFilter
	}
	if !isNil(o.EasSvcContinuity) {
		toSerialize["easSvcContinuity"] = o.EasSvcContinuity
	}
	if !isNil(o.ExpTime) {
		toSerialize["expTime"] = o.ExpTime
	}
	return toSerialize, nil
}

type NullableEasDiscoverySubscriptionPatch struct {
	value *EasDiscoverySubscriptionPatch
	isSet bool
}

func (v NullableEasDiscoverySubscriptionPatch) Get() *EasDiscoverySubscriptionPatch {
	return v.value
}

func (v *NullableEasDiscoverySubscriptionPatch) Set(val *EasDiscoverySubscriptionPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableEasDiscoverySubscriptionPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableEasDiscoverySubscriptionPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasDiscoverySubscriptionPatch(val *EasDiscoverySubscriptionPatch) *NullableEasDiscoverySubscriptionPatch {
	return &NullableEasDiscoverySubscriptionPatch{value: val, isSet: true}
}

func (v NullableEasDiscoverySubscriptionPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasDiscoverySubscriptionPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



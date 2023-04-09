/*
CAPIF_API_Invoker_Management_API

API for API invoker management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_API_Invoker_Management_API

import (
	"encoding/json"
)

// checks if the OnboardingNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnboardingNotification{}

// OnboardingNotification Represents a notification of on-boarding or update result.
type OnboardingNotification struct {
	// Set to \"true\" indicate successful on-boarding. Otherwise set to \"false\"
	Result bool `json:"result"`
	// string providing an URI formatted according to IETF RFC 3986.
	ResourceLocation           *string                     `json:"resourceLocation,omitempty"`
	ApiInvokerEnrolmentDetails *APIInvokerEnrolmentDetails `json:"apiInvokerEnrolmentDetails,omitempty"`
	ApiList                    *APIList                    `json:"apiList,omitempty"`
}

// NewOnboardingNotification instantiates a new OnboardingNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnboardingNotification(result bool) *OnboardingNotification {
	this := OnboardingNotification{}
	this.Result = result
	return &this
}

// NewOnboardingNotificationWithDefaults instantiates a new OnboardingNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnboardingNotificationWithDefaults() *OnboardingNotification {
	this := OnboardingNotification{}
	return &this
}

// GetResult returns the Result field value
func (o *OnboardingNotification) GetResult() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *OnboardingNotification) GetResultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *OnboardingNotification) SetResult(v bool) {
	o.Result = v
}

// GetResourceLocation returns the ResourceLocation field value if set, zero value otherwise.
func (o *OnboardingNotification) GetResourceLocation() string {
	if o == nil || IsNil(o.ResourceLocation) {
		var ret string
		return ret
	}
	return *o.ResourceLocation
}

// GetResourceLocationOk returns a tuple with the ResourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingNotification) GetResourceLocationOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceLocation) {
		return nil, false
	}
	return o.ResourceLocation, true
}

// HasResourceLocation returns a boolean if a field has been set.
func (o *OnboardingNotification) HasResourceLocation() bool {
	if o != nil && !IsNil(o.ResourceLocation) {
		return true
	}

	return false
}

// SetResourceLocation gets a reference to the given string and assigns it to the ResourceLocation field.
func (o *OnboardingNotification) SetResourceLocation(v string) {
	o.ResourceLocation = &v
}

// GetApiInvokerEnrolmentDetails returns the ApiInvokerEnrolmentDetails field value if set, zero value otherwise.
func (o *OnboardingNotification) GetApiInvokerEnrolmentDetails() APIInvokerEnrolmentDetails {
	if o == nil || IsNil(o.ApiInvokerEnrolmentDetails) {
		var ret APIInvokerEnrolmentDetails
		return ret
	}
	return *o.ApiInvokerEnrolmentDetails
}

// GetApiInvokerEnrolmentDetailsOk returns a tuple with the ApiInvokerEnrolmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingNotification) GetApiInvokerEnrolmentDetailsOk() (*APIInvokerEnrolmentDetails, bool) {
	if o == nil || IsNil(o.ApiInvokerEnrolmentDetails) {
		return nil, false
	}
	return o.ApiInvokerEnrolmentDetails, true
}

// HasApiInvokerEnrolmentDetails returns a boolean if a field has been set.
func (o *OnboardingNotification) HasApiInvokerEnrolmentDetails() bool {
	if o != nil && !IsNil(o.ApiInvokerEnrolmentDetails) {
		return true
	}

	return false
}

// SetApiInvokerEnrolmentDetails gets a reference to the given APIInvokerEnrolmentDetails and assigns it to the ApiInvokerEnrolmentDetails field.
func (o *OnboardingNotification) SetApiInvokerEnrolmentDetails(v APIInvokerEnrolmentDetails) {
	o.ApiInvokerEnrolmentDetails = &v
}

// GetApiList returns the ApiList field value if set, zero value otherwise.
func (o *OnboardingNotification) GetApiList() APIList {
	if o == nil || IsNil(o.ApiList) {
		var ret APIList
		return ret
	}
	return *o.ApiList
}

// GetApiListOk returns a tuple with the ApiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingNotification) GetApiListOk() (*APIList, bool) {
	if o == nil || IsNil(o.ApiList) {
		return nil, false
	}
	return o.ApiList, true
}

// HasApiList returns a boolean if a field has been set.
func (o *OnboardingNotification) HasApiList() bool {
	if o != nil && !IsNil(o.ApiList) {
		return true
	}

	return false
}

// SetApiList gets a reference to the given APIList and assigns it to the ApiList field.
func (o *OnboardingNotification) SetApiList(v APIList) {
	o.ApiList = &v
}

func (o OnboardingNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnboardingNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	if !IsNil(o.ResourceLocation) {
		toSerialize["resourceLocation"] = o.ResourceLocation
	}
	if !IsNil(o.ApiInvokerEnrolmentDetails) {
		toSerialize["apiInvokerEnrolmentDetails"] = o.ApiInvokerEnrolmentDetails
	}
	if !IsNil(o.ApiList) {
		toSerialize["apiList"] = o.ApiList
	}
	return toSerialize, nil
}

type NullableOnboardingNotification struct {
	value *OnboardingNotification
	isSet bool
}

func (v NullableOnboardingNotification) Get() *OnboardingNotification {
	return v.value
}

func (v *NullableOnboardingNotification) Set(val *OnboardingNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableOnboardingNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableOnboardingNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnboardingNotification(val *OnboardingNotification) *NullableOnboardingNotification {
	return &NullableOnboardingNotification{value: val, isSet: true}
}

func (v NullableOnboardingNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnboardingNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

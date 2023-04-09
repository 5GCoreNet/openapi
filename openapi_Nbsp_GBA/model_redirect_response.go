/*
GBA BSF Nbsp_GBA Service

GBA BSF Nbsp_GBA Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsp_GBA

import (
	"encoding/json"
)

// checks if the RedirectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedirectResponse{}

// RedirectResponse The response shall include a Location header field containing a different URI  (pointing to a different URI of an other service instance), or the same URI if a request  is redirected to the same target resource via a different SCP.
type RedirectResponse struct {
	Cause *string `json:"cause,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	TargetScp *string `json:"targetScp,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	TargetSepp *string `json:"targetSepp,omitempty"`
}

// NewRedirectResponse instantiates a new RedirectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectResponse() *RedirectResponse {
	this := RedirectResponse{}
	return &this
}

// NewRedirectResponseWithDefaults instantiates a new RedirectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectResponseWithDefaults() *RedirectResponse {
	this := RedirectResponse{}
	return &this
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *RedirectResponse) GetCause() string {
	if o == nil || IsNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectResponse) GetCauseOk() (*string, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *RedirectResponse) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *RedirectResponse) SetCause(v string) {
	o.Cause = &v
}

// GetTargetScp returns the TargetScp field value if set, zero value otherwise.
func (o *RedirectResponse) GetTargetScp() string {
	if o == nil || IsNil(o.TargetScp) {
		var ret string
		return ret
	}
	return *o.TargetScp
}

// GetTargetScpOk returns a tuple with the TargetScp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectResponse) GetTargetScpOk() (*string, bool) {
	if o == nil || IsNil(o.TargetScp) {
		return nil, false
	}
	return o.TargetScp, true
}

// HasTargetScp returns a boolean if a field has been set.
func (o *RedirectResponse) HasTargetScp() bool {
	if o != nil && !IsNil(o.TargetScp) {
		return true
	}

	return false
}

// SetTargetScp gets a reference to the given string and assigns it to the TargetScp field.
func (o *RedirectResponse) SetTargetScp(v string) {
	o.TargetScp = &v
}

// GetTargetSepp returns the TargetSepp field value if set, zero value otherwise.
func (o *RedirectResponse) GetTargetSepp() string {
	if o == nil || IsNil(o.TargetSepp) {
		var ret string
		return ret
	}
	return *o.TargetSepp
}

// GetTargetSeppOk returns a tuple with the TargetSepp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectResponse) GetTargetSeppOk() (*string, bool) {
	if o == nil || IsNil(o.TargetSepp) {
		return nil, false
	}
	return o.TargetSepp, true
}

// HasTargetSepp returns a boolean if a field has been set.
func (o *RedirectResponse) HasTargetSepp() bool {
	if o != nil && !IsNil(o.TargetSepp) {
		return true
	}

	return false
}

// SetTargetSepp gets a reference to the given string and assigns it to the TargetSepp field.
func (o *RedirectResponse) SetTargetSepp(v string) {
	o.TargetSepp = &v
}

func (o RedirectResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedirectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !IsNil(o.TargetScp) {
		toSerialize["targetScp"] = o.TargetScp
	}
	if !IsNil(o.TargetSepp) {
		toSerialize["targetSepp"] = o.TargetSepp
	}
	return toSerialize, nil
}

type NullableRedirectResponse struct {
	value *RedirectResponse
	isSet bool
}

func (v NullableRedirectResponse) Get() *RedirectResponse {
	return v.value
}

func (v *NullableRedirectResponse) Set(val *RedirectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectResponse(val *RedirectResponse) *NullableRedirectResponse {
	return &NullableRedirectResponse{value: val, isSet: true}
}

func (v NullableRedirectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

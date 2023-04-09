/*
Nhss_imsUECM

Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsUECM

import (
	"encoding/json"
)

// checks if the ScscfRestorationInfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScscfRestorationInfoRequest{}

// ScscfRestorationInfoRequest S-CSCF restoration information request
type ScscfRestorationInfoRequest struct {
	ScscfRestorationInfoRequest *ScscfRestorationInfo `json:"scscfRestorationInfoRequest,omitempty"`
}

// NewScscfRestorationInfoRequest instantiates a new ScscfRestorationInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScscfRestorationInfoRequest() *ScscfRestorationInfoRequest {
	this := ScscfRestorationInfoRequest{}
	return &this
}

// NewScscfRestorationInfoRequestWithDefaults instantiates a new ScscfRestorationInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScscfRestorationInfoRequestWithDefaults() *ScscfRestorationInfoRequest {
	this := ScscfRestorationInfoRequest{}
	return &this
}

// GetScscfRestorationInfoRequest returns the ScscfRestorationInfoRequest field value if set, zero value otherwise.
func (o *ScscfRestorationInfoRequest) GetScscfRestorationInfoRequest() ScscfRestorationInfo {
	if o == nil || IsNil(o.ScscfRestorationInfoRequest) {
		var ret ScscfRestorationInfo
		return ret
	}
	return *o.ScscfRestorationInfoRequest
}

// GetScscfRestorationInfoRequestOk returns a tuple with the ScscfRestorationInfoRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRestorationInfoRequest) GetScscfRestorationInfoRequestOk() (*ScscfRestorationInfo, bool) {
	if o == nil || IsNil(o.ScscfRestorationInfoRequest) {
		return nil, false
	}
	return o.ScscfRestorationInfoRequest, true
}

// HasScscfRestorationInfoRequest returns a boolean if a field has been set.
func (o *ScscfRestorationInfoRequest) HasScscfRestorationInfoRequest() bool {
	if o != nil && !IsNil(o.ScscfRestorationInfoRequest) {
		return true
	}

	return false
}

// SetScscfRestorationInfoRequest gets a reference to the given ScscfRestorationInfo and assigns it to the ScscfRestorationInfoRequest field.
func (o *ScscfRestorationInfoRequest) SetScscfRestorationInfoRequest(v ScscfRestorationInfo) {
	o.ScscfRestorationInfoRequest = &v
}

func (o ScscfRestorationInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScscfRestorationInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScscfRestorationInfoRequest) {
		toSerialize["scscfRestorationInfoRequest"] = o.ScscfRestorationInfoRequest
	}
	return toSerialize, nil
}

type NullableScscfRestorationInfoRequest struct {
	value *ScscfRestorationInfoRequest
	isSet bool
}

func (v NullableScscfRestorationInfoRequest) Get() *ScscfRestorationInfoRequest {
	return v.value
}

func (v *NullableScscfRestorationInfoRequest) Set(val *ScscfRestorationInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScscfRestorationInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScscfRestorationInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScscfRestorationInfoRequest(val *ScscfRestorationInfoRequest) *NullableScscfRestorationInfoRequest {
	return &NullableScscfRestorationInfoRequest{value: val, isSet: true}
}

func (v NullableScscfRestorationInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScscfRestorationInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

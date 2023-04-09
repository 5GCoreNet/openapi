/*
M5_NetworkAssistance

5GMS AF M5 Network Assistance API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M5_NetworkAssistance

import (
	"encoding/json"
)

// checks if the OperationSuccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationSuccessResponse{}

// OperationSuccessResponse struct for OperationSuccessResponse
type OperationSuccessResponse struct {
	Success bool    `json:"success"`
	Reason  *string `json:"reason,omitempty"`
}

// NewOperationSuccessResponse instantiates a new OperationSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationSuccessResponse(success bool) *OperationSuccessResponse {
	this := OperationSuccessResponse{}
	this.Success = success
	return &this
}

// NewOperationSuccessResponseWithDefaults instantiates a new OperationSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationSuccessResponseWithDefaults() *OperationSuccessResponse {
	this := OperationSuccessResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *OperationSuccessResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *OperationSuccessResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *OperationSuccessResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *OperationSuccessResponse) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationSuccessResponse) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *OperationSuccessResponse) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *OperationSuccessResponse) SetReason(v string) {
	o.Reason = &v
}

func (o OperationSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationSuccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableOperationSuccessResponse struct {
	value *OperationSuccessResponse
	isSet bool
}

func (v NullableOperationSuccessResponse) Get() *OperationSuccessResponse {
	return v.value
}

func (v *NullableOperationSuccessResponse) Set(val *OperationSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationSuccessResponse(val *OperationSuccessResponse) *NullableOperationSuccessResponse {
	return &NullableOperationSuccessResponse{value: val, isSet: true}
}

func (v NullableOperationSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

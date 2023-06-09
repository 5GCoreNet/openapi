/*
Nhss_UECM

HSS UE Context Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_UECM

import (
	"encoding/json"
)

// checks if the DeregistrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeregistrationRequest{}

// DeregistrationRequest It represents the request body of the deregistration request sent by UDM to HSS and contains the IMSI of the UE, the deregistration reason, etc.
type DeregistrationRequest struct {
	Imsi        string               `json:"imsi"`
	DeregReason DeregistrationReason `json:"deregReason"`
	Guami       *Guami               `json:"guami,omitempty"`
}

// NewDeregistrationRequest instantiates a new DeregistrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeregistrationRequest(imsi string, deregReason DeregistrationReason) *DeregistrationRequest {
	this := DeregistrationRequest{}
	this.Imsi = imsi
	this.DeregReason = deregReason
	return &this
}

// NewDeregistrationRequestWithDefaults instantiates a new DeregistrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeregistrationRequestWithDefaults() *DeregistrationRequest {
	this := DeregistrationRequest{}
	return &this
}

// GetImsi returns the Imsi field value
func (o *DeregistrationRequest) GetImsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value
// and a boolean to check if the value has been set.
func (o *DeregistrationRequest) GetImsiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Imsi, true
}

// SetImsi sets field value
func (o *DeregistrationRequest) SetImsi(v string) {
	o.Imsi = v
}

// GetDeregReason returns the DeregReason field value
func (o *DeregistrationRequest) GetDeregReason() DeregistrationReason {
	if o == nil {
		var ret DeregistrationReason
		return ret
	}

	return o.DeregReason
}

// GetDeregReasonOk returns a tuple with the DeregReason field value
// and a boolean to check if the value has been set.
func (o *DeregistrationRequest) GetDeregReasonOk() (*DeregistrationReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeregReason, true
}

// SetDeregReason sets field value
func (o *DeregistrationRequest) SetDeregReason(v DeregistrationReason) {
	o.DeregReason = v
}

// GetGuami returns the Guami field value if set, zero value otherwise.
func (o *DeregistrationRequest) GetGuami() Guami {
	if o == nil || IsNil(o.Guami) {
		var ret Guami
		return ret
	}
	return *o.Guami
}

// GetGuamiOk returns a tuple with the Guami field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeregistrationRequest) GetGuamiOk() (*Guami, bool) {
	if o == nil || IsNil(o.Guami) {
		return nil, false
	}
	return o.Guami, true
}

// HasGuami returns a boolean if a field has been set.
func (o *DeregistrationRequest) HasGuami() bool {
	if o != nil && !IsNil(o.Guami) {
		return true
	}

	return false
}

// SetGuami gets a reference to the given Guami and assigns it to the Guami field.
func (o *DeregistrationRequest) SetGuami(v Guami) {
	o.Guami = &v
}

func (o DeregistrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeregistrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["imsi"] = o.Imsi
	toSerialize["deregReason"] = o.DeregReason
	if !IsNil(o.Guami) {
		toSerialize["guami"] = o.Guami
	}
	return toSerialize, nil
}

type NullableDeregistrationRequest struct {
	value *DeregistrationRequest
	isSet bool
}

func (v NullableDeregistrationRequest) Get() *DeregistrationRequest {
	return v.value
}

func (v *NullableDeregistrationRequest) Set(val *DeregistrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeregistrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeregistrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeregistrationRequest(val *DeregistrationRequest) *NullableDeregistrationRequest {
	return &NullableDeregistrationRequest{value: val, isSet: true}
}

func (v NullableDeregistrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeregistrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

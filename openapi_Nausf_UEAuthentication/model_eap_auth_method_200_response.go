/*
AUSF API

AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nausf_UEAuthentication

import (
	"encoding/json"
)

// checks if the EapAuthMethod200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EapAuthMethod200Response{}

// EapAuthMethod200Response struct for EapAuthMethod200Response
type EapAuthMethod200Response struct {
	// contains an EAP packet
	EapPayload NullableString `json:"eapPayload"`
	// URI : /{eapSessionUri}, a map(list of key-value pairs) where member serves as key
	Links map[string]LinksValueSchema `json:"_links"`
}

// NewEapAuthMethod200Response instantiates a new EapAuthMethod200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEapAuthMethod200Response(eapPayload NullableString, links map[string]LinksValueSchema) *EapAuthMethod200Response {
	this := EapAuthMethod200Response{}
	this.EapPayload = eapPayload
	this.Links = links
	return &this
}

// NewEapAuthMethod200ResponseWithDefaults instantiates a new EapAuthMethod200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEapAuthMethod200ResponseWithDefaults() *EapAuthMethod200Response {
	this := EapAuthMethod200Response{}
	return &this
}

// GetEapPayload returns the EapPayload field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EapAuthMethod200Response) GetEapPayload() string {
	if o == nil || o.EapPayload.Get() == nil {
		var ret string
		return ret
	}

	return *o.EapPayload.Get()
}

// GetEapPayloadOk returns a tuple with the EapPayload field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EapAuthMethod200Response) GetEapPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EapPayload.Get(), o.EapPayload.IsSet()
}

// SetEapPayload sets field value
func (o *EapAuthMethod200Response) SetEapPayload(v string) {
	o.EapPayload.Set(&v)
}

// GetLinks returns the Links field value
func (o *EapAuthMethod200Response) GetLinks() map[string]LinksValueSchema {
	if o == nil {
		var ret map[string]LinksValueSchema
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *EapAuthMethod200Response) GetLinksOk() (*map[string]LinksValueSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *EapAuthMethod200Response) SetLinks(v map[string]LinksValueSchema) {
	o.Links = v
}

func (o EapAuthMethod200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EapAuthMethod200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eapPayload"] = o.EapPayload.Get()
	toSerialize["_links"] = o.Links
	return toSerialize, nil
}

type NullableEapAuthMethod200Response struct {
	value *EapAuthMethod200Response
	isSet bool
}

func (v NullableEapAuthMethod200Response) Get() *EapAuthMethod200Response {
	return v.value
}

func (v *NullableEapAuthMethod200Response) Set(val *EapAuthMethod200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEapAuthMethod200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEapAuthMethod200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEapAuthMethod200Response(val *EapAuthMethod200Response) *NullableEapAuthMethod200Response {
	return &NullableEapAuthMethod200Response{value: val, isSet: true}
}

func (v NullableEapAuthMethod200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEapAuthMethod200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

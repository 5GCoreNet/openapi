/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the AerialUeSubscriptionInfo1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AerialUeSubscriptionInfo1{}

// AerialUeSubscriptionInfo1 Contains the Aerial UE Subscription Information, it at least contains the Aerial UE Indication.
type AerialUeSubscriptionInfo1 struct {
	AerialUeInd AerialUeIndication `json:"aerialUeInd"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Var3gppUavId *string `json:"3gppUavId,omitempty"`
}

// NewAerialUeSubscriptionInfo1 instantiates a new AerialUeSubscriptionInfo1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAerialUeSubscriptionInfo1(aerialUeInd AerialUeIndication) *AerialUeSubscriptionInfo1 {
	this := AerialUeSubscriptionInfo1{}
	this.AerialUeInd = aerialUeInd
	return &this
}

// NewAerialUeSubscriptionInfo1WithDefaults instantiates a new AerialUeSubscriptionInfo1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAerialUeSubscriptionInfo1WithDefaults() *AerialUeSubscriptionInfo1 {
	this := AerialUeSubscriptionInfo1{}
	return &this
}

// GetAerialUeInd returns the AerialUeInd field value
func (o *AerialUeSubscriptionInfo1) GetAerialUeInd() AerialUeIndication {
	if o == nil {
		var ret AerialUeIndication
		return ret
	}

	return o.AerialUeInd
}

// GetAerialUeIndOk returns a tuple with the AerialUeInd field value
// and a boolean to check if the value has been set.
func (o *AerialUeSubscriptionInfo1) GetAerialUeIndOk() (*AerialUeIndication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AerialUeInd, true
}

// SetAerialUeInd sets field value
func (o *AerialUeSubscriptionInfo1) SetAerialUeInd(v AerialUeIndication) {
	o.AerialUeInd = v
}

// GetVar3gppUavId returns the Var3gppUavId field value if set, zero value otherwise.
func (o *AerialUeSubscriptionInfo1) GetVar3gppUavId() string {
	if o == nil || isNil(o.Var3gppUavId) {
		var ret string
		return ret
	}
	return *o.Var3gppUavId
}

// GetVar3gppUavIdOk returns a tuple with the Var3gppUavId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AerialUeSubscriptionInfo1) GetVar3gppUavIdOk() (*string, bool) {
	if o == nil || isNil(o.Var3gppUavId) {
		return nil, false
	}
	return o.Var3gppUavId, true
}

// HasVar3gppUavId returns a boolean if a field has been set.
func (o *AerialUeSubscriptionInfo1) HasVar3gppUavId() bool {
	if o != nil && !isNil(o.Var3gppUavId) {
		return true
	}

	return false
}

// SetVar3gppUavId gets a reference to the given string and assigns it to the Var3gppUavId field.
func (o *AerialUeSubscriptionInfo1) SetVar3gppUavId(v string) {
	o.Var3gppUavId = &v
}

func (o AerialUeSubscriptionInfo1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AerialUeSubscriptionInfo1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aerialUeInd"] = o.AerialUeInd
	if !isNil(o.Var3gppUavId) {
		toSerialize["3gppUavId"] = o.Var3gppUavId
	}
	return toSerialize, nil
}

type NullableAerialUeSubscriptionInfo1 struct {
	value *AerialUeSubscriptionInfo1
	isSet bool
}

func (v NullableAerialUeSubscriptionInfo1) Get() *AerialUeSubscriptionInfo1 {
	return v.value
}

func (v *NullableAerialUeSubscriptionInfo1) Set(val *AerialUeSubscriptionInfo1) {
	v.value = val
	v.isSet = true
}

func (v NullableAerialUeSubscriptionInfo1) IsSet() bool {
	return v.isSet
}

func (v *NullableAerialUeSubscriptionInfo1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAerialUeSubscriptionInfo1(val *AerialUeSubscriptionInfo1) *NullableAerialUeSubscriptionInfo1 {
	return &NullableAerialUeSubscriptionInfo1{value: val, isSet: true}
}

func (v NullableAerialUeSubscriptionInfo1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAerialUeSubscriptionInfo1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



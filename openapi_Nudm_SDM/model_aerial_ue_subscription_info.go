/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
)

// checks if the AerialUeSubscriptionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AerialUeSubscriptionInfo{}

// AerialUeSubscriptionInfo Contains the Aerial UE Subscription Information, it at least contains the Aerial UE Indication.
type AerialUeSubscriptionInfo struct {
	AerialUeInd AerialUeIndication `json:"aerialUeInd"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Var3gppUavId *string `json:"3gppUavId,omitempty"`
}

// NewAerialUeSubscriptionInfo instantiates a new AerialUeSubscriptionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAerialUeSubscriptionInfo(aerialUeInd AerialUeIndication) *AerialUeSubscriptionInfo {
	this := AerialUeSubscriptionInfo{}
	this.AerialUeInd = aerialUeInd
	return &this
}

// NewAerialUeSubscriptionInfoWithDefaults instantiates a new AerialUeSubscriptionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAerialUeSubscriptionInfoWithDefaults() *AerialUeSubscriptionInfo {
	this := AerialUeSubscriptionInfo{}
	return &this
}

// GetAerialUeInd returns the AerialUeInd field value
func (o *AerialUeSubscriptionInfo) GetAerialUeInd() AerialUeIndication {
	if o == nil {
		var ret AerialUeIndication
		return ret
	}

	return o.AerialUeInd
}

// GetAerialUeIndOk returns a tuple with the AerialUeInd field value
// and a boolean to check if the value has been set.
func (o *AerialUeSubscriptionInfo) GetAerialUeIndOk() (*AerialUeIndication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AerialUeInd, true
}

// SetAerialUeInd sets field value
func (o *AerialUeSubscriptionInfo) SetAerialUeInd(v AerialUeIndication) {
	o.AerialUeInd = v
}

// GetVar3gppUavId returns the Var3gppUavId field value if set, zero value otherwise.
func (o *AerialUeSubscriptionInfo) GetVar3gppUavId() string {
	if o == nil || isNil(o.Var3gppUavId) {
		var ret string
		return ret
	}
	return *o.Var3gppUavId
}

// GetVar3gppUavIdOk returns a tuple with the Var3gppUavId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AerialUeSubscriptionInfo) GetVar3gppUavIdOk() (*string, bool) {
	if o == nil || isNil(o.Var3gppUavId) {
		return nil, false
	}
	return o.Var3gppUavId, true
}

// HasVar3gppUavId returns a boolean if a field has been set.
func (o *AerialUeSubscriptionInfo) HasVar3gppUavId() bool {
	if o != nil && !isNil(o.Var3gppUavId) {
		return true
	}

	return false
}

// SetVar3gppUavId gets a reference to the given string and assigns it to the Var3gppUavId field.
func (o *AerialUeSubscriptionInfo) SetVar3gppUavId(v string) {
	o.Var3gppUavId = &v
}

func (o AerialUeSubscriptionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AerialUeSubscriptionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aerialUeInd"] = o.AerialUeInd
	if !isNil(o.Var3gppUavId) {
		toSerialize["3gppUavId"] = o.Var3gppUavId
	}
	return toSerialize, nil
}

type NullableAerialUeSubscriptionInfo struct {
	value *AerialUeSubscriptionInfo
	isSet bool
}

func (v NullableAerialUeSubscriptionInfo) Get() *AerialUeSubscriptionInfo {
	return v.value
}

func (v *NullableAerialUeSubscriptionInfo) Set(val *AerialUeSubscriptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAerialUeSubscriptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAerialUeSubscriptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAerialUeSubscriptionInfo(val *AerialUeSubscriptionInfo) *NullableAerialUeSubscriptionInfo {
	return &NullableAerialUeSubscriptionInfo{value: val, isSet: true}
}

func (v NullableAerialUeSubscriptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAerialUeSubscriptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



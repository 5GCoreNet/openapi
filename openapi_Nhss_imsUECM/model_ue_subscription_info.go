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

// checks if the UeSubscriptionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeSubscriptionInfo{}

// UeSubscriptionInfo Subscription information of the UE for the SIP Registration State event
type UeSubscriptionInfo struct {
	CallIdSipHeader string `json:"callIdSipHeader"`
	FromSipHeader   string `json:"fromSipHeader"`
	ToSipHeader     string `json:"toSipHeader"`
	RecordRoute     string `json:"recordRoute"`
	Contact         string `json:"contact"`
}

// NewUeSubscriptionInfo instantiates a new UeSubscriptionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeSubscriptionInfo(callIdSipHeader string, fromSipHeader string, toSipHeader string, recordRoute string, contact string) *UeSubscriptionInfo {
	this := UeSubscriptionInfo{}
	this.CallIdSipHeader = callIdSipHeader
	this.FromSipHeader = fromSipHeader
	this.ToSipHeader = toSipHeader
	this.RecordRoute = recordRoute
	this.Contact = contact
	return &this
}

// NewUeSubscriptionInfoWithDefaults instantiates a new UeSubscriptionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeSubscriptionInfoWithDefaults() *UeSubscriptionInfo {
	this := UeSubscriptionInfo{}
	return &this
}

// GetCallIdSipHeader returns the CallIdSipHeader field value
func (o *UeSubscriptionInfo) GetCallIdSipHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallIdSipHeader
}

// GetCallIdSipHeaderOk returns a tuple with the CallIdSipHeader field value
// and a boolean to check if the value has been set.
func (o *UeSubscriptionInfo) GetCallIdSipHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallIdSipHeader, true
}

// SetCallIdSipHeader sets field value
func (o *UeSubscriptionInfo) SetCallIdSipHeader(v string) {
	o.CallIdSipHeader = v
}

// GetFromSipHeader returns the FromSipHeader field value
func (o *UeSubscriptionInfo) GetFromSipHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromSipHeader
}

// GetFromSipHeaderOk returns a tuple with the FromSipHeader field value
// and a boolean to check if the value has been set.
func (o *UeSubscriptionInfo) GetFromSipHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromSipHeader, true
}

// SetFromSipHeader sets field value
func (o *UeSubscriptionInfo) SetFromSipHeader(v string) {
	o.FromSipHeader = v
}

// GetToSipHeader returns the ToSipHeader field value
func (o *UeSubscriptionInfo) GetToSipHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToSipHeader
}

// GetToSipHeaderOk returns a tuple with the ToSipHeader field value
// and a boolean to check if the value has been set.
func (o *UeSubscriptionInfo) GetToSipHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToSipHeader, true
}

// SetToSipHeader sets field value
func (o *UeSubscriptionInfo) SetToSipHeader(v string) {
	o.ToSipHeader = v
}

// GetRecordRoute returns the RecordRoute field value
func (o *UeSubscriptionInfo) GetRecordRoute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordRoute
}

// GetRecordRouteOk returns a tuple with the RecordRoute field value
// and a boolean to check if the value has been set.
func (o *UeSubscriptionInfo) GetRecordRouteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordRoute, true
}

// SetRecordRoute sets field value
func (o *UeSubscriptionInfo) SetRecordRoute(v string) {
	o.RecordRoute = v
}

// GetContact returns the Contact field value
func (o *UeSubscriptionInfo) GetContact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *UeSubscriptionInfo) GetContactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value
func (o *UeSubscriptionInfo) SetContact(v string) {
	o.Contact = v
}

func (o UeSubscriptionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeSubscriptionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["callIdSipHeader"] = o.CallIdSipHeader
	toSerialize["fromSipHeader"] = o.FromSipHeader
	toSerialize["toSipHeader"] = o.ToSipHeader
	toSerialize["recordRoute"] = o.RecordRoute
	toSerialize["contact"] = o.Contact
	return toSerialize, nil
}

type NullableUeSubscriptionInfo struct {
	value *UeSubscriptionInfo
	isSet bool
}

func (v NullableUeSubscriptionInfo) Get() *UeSubscriptionInfo {
	return v.value
}

func (v *NullableUeSubscriptionInfo) Set(val *UeSubscriptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUeSubscriptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUeSubscriptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeSubscriptionInfo(val *UeSubscriptionInfo) *NullableUeSubscriptionInfo {
	return &NullableUeSubscriptionInfo{value: val, isSet: true}
}

func (v NullableUeSubscriptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeSubscriptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the NwdafSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NwdafSubscription{}

// NwdafSubscription Individual NWDAF subscription identified by the subscription Id.
type NwdafSubscription struct {
	// String providing an URI formatted according to RFC 3986.
	NwdafEvtSubsServiceUri  string                   `json:"nwdafEvtSubsServiceUri"`
	NwdafEventsSubscription NnwdafEventsSubscription `json:"nwdafEventsSubscription"`
}

// NewNwdafSubscription instantiates a new NwdafSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwdafSubscription(nwdafEvtSubsServiceUri string, nwdafEventsSubscription NnwdafEventsSubscription) *NwdafSubscription {
	this := NwdafSubscription{}
	this.NwdafEvtSubsServiceUri = nwdafEvtSubsServiceUri
	this.NwdafEventsSubscription = nwdafEventsSubscription
	return &this
}

// NewNwdafSubscriptionWithDefaults instantiates a new NwdafSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwdafSubscriptionWithDefaults() *NwdafSubscription {
	this := NwdafSubscription{}
	return &this
}

// GetNwdafEvtSubsServiceUri returns the NwdafEvtSubsServiceUri field value
func (o *NwdafSubscription) GetNwdafEvtSubsServiceUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NwdafEvtSubsServiceUri
}

// GetNwdafEvtSubsServiceUriOk returns a tuple with the NwdafEvtSubsServiceUri field value
// and a boolean to check if the value has been set.
func (o *NwdafSubscription) GetNwdafEvtSubsServiceUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NwdafEvtSubsServiceUri, true
}

// SetNwdafEvtSubsServiceUri sets field value
func (o *NwdafSubscription) SetNwdafEvtSubsServiceUri(v string) {
	o.NwdafEvtSubsServiceUri = v
}

// GetNwdafEventsSubscription returns the NwdafEventsSubscription field value
func (o *NwdafSubscription) GetNwdafEventsSubscription() NnwdafEventsSubscription {
	if o == nil {
		var ret NnwdafEventsSubscription
		return ret
	}

	return o.NwdafEventsSubscription
}

// GetNwdafEventsSubscriptionOk returns a tuple with the NwdafEventsSubscription field value
// and a boolean to check if the value has been set.
func (o *NwdafSubscription) GetNwdafEventsSubscriptionOk() (*NnwdafEventsSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NwdafEventsSubscription, true
}

// SetNwdafEventsSubscription sets field value
func (o *NwdafSubscription) SetNwdafEventsSubscription(v NnwdafEventsSubscription) {
	o.NwdafEventsSubscription = v
}

func (o NwdafSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NwdafSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nwdafEvtSubsServiceUri"] = o.NwdafEvtSubsServiceUri
	toSerialize["nwdafEventsSubscription"] = o.NwdafEventsSubscription
	return toSerialize, nil
}

type NullableNwdafSubscription struct {
	value *NwdafSubscription
	isSet bool
}

func (v NullableNwdafSubscription) Get() *NwdafSubscription {
	return v.value
}

func (v *NullableNwdafSubscription) Set(val *NwdafSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafSubscription(val *NwdafSubscription) *NullableNwdafSubscription {
	return &NullableNwdafSubscription{value: val, isSet: true}
}

func (v NullableNwdafSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

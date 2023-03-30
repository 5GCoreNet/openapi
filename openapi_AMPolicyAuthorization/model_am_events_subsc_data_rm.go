/*
3gpp-am-policyauthorization

API for AM policy authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AMPolicyAuthorization

import (
	"encoding/json"
)

// checks if the AmEventsSubscDataRm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmEventsSubscDataRm{}

// AmEventsSubscDataRm This data type is defined in the same way as the AmEventsSubscData but with the OpenAPI nullable property set to true.
type AmEventsSubscDataRm struct {
	// String providing an URI formatted according to RFC 3986.
	EventNotifUri *string `json:"eventNotifUri,omitempty"`
	Events []AmEventData `json:"events,omitempty"`
}

// NewAmEventsSubscDataRm instantiates a new AmEventsSubscDataRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmEventsSubscDataRm() *AmEventsSubscDataRm {
	this := AmEventsSubscDataRm{}
	return &this
}

// NewAmEventsSubscDataRmWithDefaults instantiates a new AmEventsSubscDataRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmEventsSubscDataRmWithDefaults() *AmEventsSubscDataRm {
	this := AmEventsSubscDataRm{}
	return &this
}

// GetEventNotifUri returns the EventNotifUri field value if set, zero value otherwise.
func (o *AmEventsSubscDataRm) GetEventNotifUri() string {
	if o == nil || IsNil(o.EventNotifUri) {
		var ret string
		return ret
	}
	return *o.EventNotifUri
}

// GetEventNotifUriOk returns a tuple with the EventNotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmEventsSubscDataRm) GetEventNotifUriOk() (*string, bool) {
	if o == nil || IsNil(o.EventNotifUri) {
		return nil, false
	}
	return o.EventNotifUri, true
}

// HasEventNotifUri returns a boolean if a field has been set.
func (o *AmEventsSubscDataRm) HasEventNotifUri() bool {
	if o != nil && !IsNil(o.EventNotifUri) {
		return true
	}

	return false
}

// SetEventNotifUri gets a reference to the given string and assigns it to the EventNotifUri field.
func (o *AmEventsSubscDataRm) SetEventNotifUri(v string) {
	o.EventNotifUri = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *AmEventsSubscDataRm) GetEvents() []AmEventData {
	if o == nil || IsNil(o.Events) {
		var ret []AmEventData
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmEventsSubscDataRm) GetEventsOk() ([]AmEventData, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *AmEventsSubscDataRm) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []AmEventData and assigns it to the Events field.
func (o *AmEventsSubscDataRm) SetEvents(v []AmEventData) {
	o.Events = v
}

func (o AmEventsSubscDataRm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmEventsSubscDataRm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventNotifUri) {
		toSerialize["eventNotifUri"] = o.EventNotifUri
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullableAmEventsSubscDataRm struct {
	value *AmEventsSubscDataRm
	isSet bool
}

func (v NullableAmEventsSubscDataRm) Get() *AmEventsSubscDataRm {
	return v.value
}

func (v *NullableAmEventsSubscDataRm) Set(val *AmEventsSubscDataRm) {
	v.value = val
	v.isSet = true
}

func (v NullableAmEventsSubscDataRm) IsSet() bool {
	return v.isSet
}

func (v *NullableAmEventsSubscDataRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmEventsSubscDataRm(val *AmEventsSubscDataRm) *NullableAmEventsSubscDataRm {
	return &NullableAmEventsSubscDataRm{value: val, isSet: true}
}

func (v NullableAmEventsSubscDataRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmEventsSubscDataRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



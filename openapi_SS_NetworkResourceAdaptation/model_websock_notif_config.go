/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceAdaptation

import (
	"encoding/json"
)

// checks if the WebsockNotifConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebsockNotifConfig{}

// WebsockNotifConfig Represents the configuration information for the delivery of notifications over Websockets.
type WebsockNotifConfig struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	WebsocketUri *string `json:"websocketUri,omitempty"`
	// Set by the SCS/AS to indicate that the Websocket delivery is requested.
	RequestWebsocketUri *bool `json:"requestWebsocketUri,omitempty"`
}

// NewWebsockNotifConfig instantiates a new WebsockNotifConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsockNotifConfig() *WebsockNotifConfig {
	this := WebsockNotifConfig{}
	return &this
}

// NewWebsockNotifConfigWithDefaults instantiates a new WebsockNotifConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsockNotifConfigWithDefaults() *WebsockNotifConfig {
	this := WebsockNotifConfig{}
	return &this
}

// GetWebsocketUri returns the WebsocketUri field value if set, zero value otherwise.
func (o *WebsockNotifConfig) GetWebsocketUri() string {
	if o == nil || isNil(o.WebsocketUri) {
		var ret string
		return ret
	}
	return *o.WebsocketUri
}

// GetWebsocketUriOk returns a tuple with the WebsocketUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsockNotifConfig) GetWebsocketUriOk() (*string, bool) {
	if o == nil || isNil(o.WebsocketUri) {
		return nil, false
	}
	return o.WebsocketUri, true
}

// HasWebsocketUri returns a boolean if a field has been set.
func (o *WebsockNotifConfig) HasWebsocketUri() bool {
	if o != nil && !isNil(o.WebsocketUri) {
		return true
	}

	return false
}

// SetWebsocketUri gets a reference to the given string and assigns it to the WebsocketUri field.
func (o *WebsockNotifConfig) SetWebsocketUri(v string) {
	o.WebsocketUri = &v
}

// GetRequestWebsocketUri returns the RequestWebsocketUri field value if set, zero value otherwise.
func (o *WebsockNotifConfig) GetRequestWebsocketUri() bool {
	if o == nil || isNil(o.RequestWebsocketUri) {
		var ret bool
		return ret
	}
	return *o.RequestWebsocketUri
}

// GetRequestWebsocketUriOk returns a tuple with the RequestWebsocketUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsockNotifConfig) GetRequestWebsocketUriOk() (*bool, bool) {
	if o == nil || isNil(o.RequestWebsocketUri) {
		return nil, false
	}
	return o.RequestWebsocketUri, true
}

// HasRequestWebsocketUri returns a boolean if a field has been set.
func (o *WebsockNotifConfig) HasRequestWebsocketUri() bool {
	if o != nil && !isNil(o.RequestWebsocketUri) {
		return true
	}

	return false
}

// SetRequestWebsocketUri gets a reference to the given bool and assigns it to the RequestWebsocketUri field.
func (o *WebsockNotifConfig) SetRequestWebsocketUri(v bool) {
	o.RequestWebsocketUri = &v
}

func (o WebsockNotifConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebsockNotifConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WebsocketUri) {
		toSerialize["websocketUri"] = o.WebsocketUri
	}
	if !isNil(o.RequestWebsocketUri) {
		toSerialize["requestWebsocketUri"] = o.RequestWebsocketUri
	}
	return toSerialize, nil
}

type NullableWebsockNotifConfig struct {
	value *WebsockNotifConfig
	isSet bool
}

func (v NullableWebsockNotifConfig) Get() *WebsockNotifConfig {
	return v.value
}

func (v *NullableWebsockNotifConfig) Set(val *WebsockNotifConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsockNotifConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsockNotifConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsockNotifConfig(val *WebsockNotifConfig) *NullableWebsockNotifConfig {
	return &NullableWebsockNotifConfig{value: val, isSet: true}
}

func (v NullableWebsockNotifConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsockNotifConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



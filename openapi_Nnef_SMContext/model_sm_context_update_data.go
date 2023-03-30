/*
Nnef_SMContext

Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_SMContext

import (
	"encoding/json"
)

// checks if the SmContextUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmContextUpdateData{}

// SmContextUpdateData Representation of the updates to apply to the Individual SM context.
type SmContextUpdateData struct {
	// String providing an URI formatted according to RFC 3986.
	DlNiddEndPoint *string `json:"dlNiddEndPoint,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NotificationUri *string `json:"notificationUri,omitempty"`
	SmContextConfig *SmContextConfiguration `json:"smContextConfig,omitempty"`
}

// NewSmContextUpdateData instantiates a new SmContextUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmContextUpdateData() *SmContextUpdateData {
	this := SmContextUpdateData{}
	return &this
}

// NewSmContextUpdateDataWithDefaults instantiates a new SmContextUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmContextUpdateDataWithDefaults() *SmContextUpdateData {
	this := SmContextUpdateData{}
	return &this
}

// GetDlNiddEndPoint returns the DlNiddEndPoint field value if set, zero value otherwise.
func (o *SmContextUpdateData) GetDlNiddEndPoint() string {
	if o == nil || IsNil(o.DlNiddEndPoint) {
		var ret string
		return ret
	}
	return *o.DlNiddEndPoint
}

// GetDlNiddEndPointOk returns a tuple with the DlNiddEndPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextUpdateData) GetDlNiddEndPointOk() (*string, bool) {
	if o == nil || IsNil(o.DlNiddEndPoint) {
		return nil, false
	}
	return o.DlNiddEndPoint, true
}

// HasDlNiddEndPoint returns a boolean if a field has been set.
func (o *SmContextUpdateData) HasDlNiddEndPoint() bool {
	if o != nil && !IsNil(o.DlNiddEndPoint) {
		return true
	}

	return false
}

// SetDlNiddEndPoint gets a reference to the given string and assigns it to the DlNiddEndPoint field.
func (o *SmContextUpdateData) SetDlNiddEndPoint(v string) {
	o.DlNiddEndPoint = &v
}

// GetNotificationUri returns the NotificationUri field value if set, zero value otherwise.
func (o *SmContextUpdateData) GetNotificationUri() string {
	if o == nil || IsNil(o.NotificationUri) {
		var ret string
		return ret
	}
	return *o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextUpdateData) GetNotificationUriOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationUri) {
		return nil, false
	}
	return o.NotificationUri, true
}

// HasNotificationUri returns a boolean if a field has been set.
func (o *SmContextUpdateData) HasNotificationUri() bool {
	if o != nil && !IsNil(o.NotificationUri) {
		return true
	}

	return false
}

// SetNotificationUri gets a reference to the given string and assigns it to the NotificationUri field.
func (o *SmContextUpdateData) SetNotificationUri(v string) {
	o.NotificationUri = &v
}

// GetSmContextConfig returns the SmContextConfig field value if set, zero value otherwise.
func (o *SmContextUpdateData) GetSmContextConfig() SmContextConfiguration {
	if o == nil || IsNil(o.SmContextConfig) {
		var ret SmContextConfiguration
		return ret
	}
	return *o.SmContextConfig
}

// GetSmContextConfigOk returns a tuple with the SmContextConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextUpdateData) GetSmContextConfigOk() (*SmContextConfiguration, bool) {
	if o == nil || IsNil(o.SmContextConfig) {
		return nil, false
	}
	return o.SmContextConfig, true
}

// HasSmContextConfig returns a boolean if a field has been set.
func (o *SmContextUpdateData) HasSmContextConfig() bool {
	if o != nil && !IsNil(o.SmContextConfig) {
		return true
	}

	return false
}

// SetSmContextConfig gets a reference to the given SmContextConfiguration and assigns it to the SmContextConfig field.
func (o *SmContextUpdateData) SetSmContextConfig(v SmContextConfiguration) {
	o.SmContextConfig = &v
}

func (o SmContextUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmContextUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DlNiddEndPoint) {
		toSerialize["dlNiddEndPoint"] = o.DlNiddEndPoint
	}
	if !IsNil(o.NotificationUri) {
		toSerialize["notificationUri"] = o.NotificationUri
	}
	if !IsNil(o.SmContextConfig) {
		toSerialize["smContextConfig"] = o.SmContextConfig
	}
	return toSerialize, nil
}

type NullableSmContextUpdateData struct {
	value *SmContextUpdateData
	isSet bool
}

func (v NullableSmContextUpdateData) Get() *SmContextUpdateData {
	return v.value
}

func (v *NullableSmContextUpdateData) Set(val *SmContextUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmContextUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmContextUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmContextUpdateData(val *SmContextUpdateData) *NullableSmContextUpdateData {
	return &NullableSmContextUpdateData{value: val, isSet: true}
}

func (v NullableSmContextUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmContextUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



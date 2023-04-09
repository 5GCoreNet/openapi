/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the SdpDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SdpDescription{}

// SdpDescription Contains a SDP line (and optionally the value in the line) within the body (if any) of a SIP request
type SdpDescription struct {
	Line    string  `json:"line"`
	Content *string `json:"content,omitempty"`
}

// NewSdpDescription instantiates a new SdpDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdpDescription(line string) *SdpDescription {
	this := SdpDescription{}
	this.Line = line
	return &this
}

// NewSdpDescriptionWithDefaults instantiates a new SdpDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdpDescriptionWithDefaults() *SdpDescription {
	this := SdpDescription{}
	return &this
}

// GetLine returns the Line field value
func (o *SdpDescription) GetLine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Line
}

// GetLineOk returns a tuple with the Line field value
// and a boolean to check if the value has been set.
func (o *SdpDescription) GetLineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line, true
}

// SetLine sets field value
func (o *SdpDescription) SetLine(v string) {
	o.Line = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *SdpDescription) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdpDescription) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *SdpDescription) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *SdpDescription) SetContent(v string) {
	o.Content = &v
}

func (o SdpDescription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SdpDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["line"] = o.Line
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableSdpDescription struct {
	value *SdpDescription
	isSet bool
}

func (v NullableSdpDescription) Get() *SdpDescription {
	return v.value
}

func (v *NullableSdpDescription) Set(val *SdpDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableSdpDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableSdpDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdpDescription(val *SdpDescription) *NullableSdpDescription {
	return &NullableSdpDescription{value: val, isSet: true}
}

func (v NullableSdpDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdpDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

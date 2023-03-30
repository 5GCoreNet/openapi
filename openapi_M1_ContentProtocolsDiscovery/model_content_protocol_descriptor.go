/*
M1_ContentProtocolsDiscovery

5GMS AF M1 Content Protocols Discovery API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_ContentProtocolsDiscovery

import (
	"encoding/json"
)

// checks if the ContentProtocolDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentProtocolDescriptor{}

// ContentProtocolDescriptor A descriptor describing a content protocol.
type ContentProtocolDescriptor struct {
	// String providing an URI formatted according to RFC 3986.
	TermIdentifier string `json:"termIdentifier"`
	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	DescriptionLocator *string `json:"descriptionLocator,omitempty"`
}

// NewContentProtocolDescriptor instantiates a new ContentProtocolDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentProtocolDescriptor(termIdentifier string) *ContentProtocolDescriptor {
	this := ContentProtocolDescriptor{}
	this.TermIdentifier = termIdentifier
	return &this
}

// NewContentProtocolDescriptorWithDefaults instantiates a new ContentProtocolDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentProtocolDescriptorWithDefaults() *ContentProtocolDescriptor {
	this := ContentProtocolDescriptor{}
	return &this
}

// GetTermIdentifier returns the TermIdentifier field value
func (o *ContentProtocolDescriptor) GetTermIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TermIdentifier
}

// GetTermIdentifierOk returns a tuple with the TermIdentifier field value
// and a boolean to check if the value has been set.
func (o *ContentProtocolDescriptor) GetTermIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TermIdentifier, true
}

// SetTermIdentifier sets field value
func (o *ContentProtocolDescriptor) SetTermIdentifier(v string) {
	o.TermIdentifier = v
}

// GetDescriptionLocator returns the DescriptionLocator field value if set, zero value otherwise.
func (o *ContentProtocolDescriptor) GetDescriptionLocator() string {
	if o == nil || IsNil(o.DescriptionLocator) {
		var ret string
		return ret
	}
	return *o.DescriptionLocator
}

// GetDescriptionLocatorOk returns a tuple with the DescriptionLocator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentProtocolDescriptor) GetDescriptionLocatorOk() (*string, bool) {
	if o == nil || IsNil(o.DescriptionLocator) {
		return nil, false
	}
	return o.DescriptionLocator, true
}

// HasDescriptionLocator returns a boolean if a field has been set.
func (o *ContentProtocolDescriptor) HasDescriptionLocator() bool {
	if o != nil && !IsNil(o.DescriptionLocator) {
		return true
	}

	return false
}

// SetDescriptionLocator gets a reference to the given string and assigns it to the DescriptionLocator field.
func (o *ContentProtocolDescriptor) SetDescriptionLocator(v string) {
	o.DescriptionLocator = &v
}

func (o ContentProtocolDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentProtocolDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["termIdentifier"] = o.TermIdentifier
	if !IsNil(o.DescriptionLocator) {
		toSerialize["descriptionLocator"] = o.DescriptionLocator
	}
	return toSerialize, nil
}

type NullableContentProtocolDescriptor struct {
	value *ContentProtocolDescriptor
	isSet bool
}

func (v NullableContentProtocolDescriptor) Get() *ContentProtocolDescriptor {
	return v.value
}

func (v *NullableContentProtocolDescriptor) Set(val *ContentProtocolDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableContentProtocolDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableContentProtocolDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentProtocolDescriptor(val *ContentProtocolDescriptor) *NullableContentProtocolDescriptor {
	return &NullableContentProtocolDescriptor{value: val, isSet: true}
}

func (v NullableContentProtocolDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentProtocolDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


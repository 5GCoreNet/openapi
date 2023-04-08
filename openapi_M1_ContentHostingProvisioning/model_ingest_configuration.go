/*
M1_ContentHostingProvisioning

5GMS AF M1 Content Hosting Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_ContentHostingProvisioning

import (
	"encoding/json"
)

// checks if the IngestConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestConfiguration{}

// IngestConfiguration A configuration for content ingest.
type IngestConfiguration struct {
	Pull *bool `json:"pull,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	Protocol *string `json:"protocol,omitempty"`
	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	BaseURL *string `json:"baseURL,omitempty"`
}

// NewIngestConfiguration instantiates a new IngestConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestConfiguration() *IngestConfiguration {
	this := IngestConfiguration{}
	return &this
}

// NewIngestConfigurationWithDefaults instantiates a new IngestConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestConfigurationWithDefaults() *IngestConfiguration {
	this := IngestConfiguration{}
	return &this
}

// GetPull returns the Pull field value if set, zero value otherwise.
func (o *IngestConfiguration) GetPull() bool {
	if o == nil || isNil(o.Pull) {
		var ret bool
		return ret
	}
	return *o.Pull
}

// GetPullOk returns a tuple with the Pull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestConfiguration) GetPullOk() (*bool, bool) {
	if o == nil || isNil(o.Pull) {
		return nil, false
	}
	return o.Pull, true
}

// HasPull returns a boolean if a field has been set.
func (o *IngestConfiguration) HasPull() bool {
	if o != nil && !isNil(o.Pull) {
		return true
	}

	return false
}

// SetPull gets a reference to the given bool and assigns it to the Pull field.
func (o *IngestConfiguration) SetPull(v bool) {
	o.Pull = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *IngestConfiguration) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestConfiguration) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *IngestConfiguration) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *IngestConfiguration) SetProtocol(v string) {
	o.Protocol = &v
}

// GetBaseURL returns the BaseURL field value if set, zero value otherwise.
func (o *IngestConfiguration) GetBaseURL() string {
	if o == nil || isNil(o.BaseURL) {
		var ret string
		return ret
	}
	return *o.BaseURL
}

// GetBaseURLOk returns a tuple with the BaseURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestConfiguration) GetBaseURLOk() (*string, bool) {
	if o == nil || isNil(o.BaseURL) {
		return nil, false
	}
	return o.BaseURL, true
}

// HasBaseURL returns a boolean if a field has been set.
func (o *IngestConfiguration) HasBaseURL() bool {
	if o != nil && !isNil(o.BaseURL) {
		return true
	}

	return false
}

// SetBaseURL gets a reference to the given string and assigns it to the BaseURL field.
func (o *IngestConfiguration) SetBaseURL(v string) {
	o.BaseURL = &v
}

func (o IngestConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Pull) {
		toSerialize["pull"] = o.Pull
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.BaseURL) {
		toSerialize["baseURL"] = o.BaseURL
	}
	return toSerialize, nil
}

type NullableIngestConfiguration struct {
	value *IngestConfiguration
	isSet bool
}

func (v NullableIngestConfiguration) Get() *IngestConfiguration {
	return v.value
}

func (v *NullableIngestConfiguration) Set(val *IngestConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestConfiguration(val *IngestConfiguration) *NullableIngestConfiguration {
	return &NullableIngestConfiguration{value: val, isSet: true}
}

func (v NullableIngestConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



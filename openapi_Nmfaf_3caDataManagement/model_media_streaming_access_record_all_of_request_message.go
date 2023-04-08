/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// checks if the MediaStreamingAccessRecordAllOfRequestMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaStreamingAccessRecordAllOfRequestMessage{}

// MediaStreamingAccessRecordAllOfRequestMessage struct for MediaStreamingAccessRecordAllOfRequestMessage
type MediaStreamingAccessRecordAllOfRequestMessage struct {
	Method string `json:"method"`
	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	Url string `json:"url"`
	ProtocolVersion string `json:"protocolVersion"`
	Range *string `json:"range,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Size int32 `json:"size"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	BodySize int32 `json:"bodySize"`
	ContentType *string `json:"contentType,omitempty"`
	UserAgent *string `json:"userAgent,omitempty"`
	UserIdentity *string `json:"userIdentity,omitempty"`
	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	Referer *string `json:"referer,omitempty"`
}

// NewMediaStreamingAccessRecordAllOfRequestMessage instantiates a new MediaStreamingAccessRecordAllOfRequestMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaStreamingAccessRecordAllOfRequestMessage(method string, url string, protocolVersion string, size int32, bodySize int32) *MediaStreamingAccessRecordAllOfRequestMessage {
	this := MediaStreamingAccessRecordAllOfRequestMessage{}
	this.Method = method
	this.Url = url
	this.ProtocolVersion = protocolVersion
	this.Size = size
	this.BodySize = bodySize
	return &this
}

// NewMediaStreamingAccessRecordAllOfRequestMessageWithDefaults instantiates a new MediaStreamingAccessRecordAllOfRequestMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaStreamingAccessRecordAllOfRequestMessageWithDefaults() *MediaStreamingAccessRecordAllOfRequestMessage {
	this := MediaStreamingAccessRecordAllOfRequestMessage{}
	return &this
}

// GetMethod returns the Method field value
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetMethod(v string) {
	o.Method = v
}

// GetUrl returns the Url field value
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetUrl(v string) {
	o.Url = v
}

// GetProtocolVersion returns the ProtocolVersion field value
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetProtocolVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProtocolVersion
}

// GetProtocolVersionOk returns a tuple with the ProtocolVersion field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetProtocolVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtocolVersion, true
}

// SetProtocolVersion sets field value
func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetProtocolVersion(v string) {
	o.ProtocolVersion = v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetRange() string {
	if o == nil || isNil(o.Range) {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetRangeOk() (*string, bool) {
	if o == nil || isNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) HasRange() bool {
	if o != nil && !isNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetRange(v string) {
	o.Range = &v
}

// GetSize returns the Size field value
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetSize(v int32) {
	o.Size = v
}

// GetBodySize returns the BodySize field value
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetBodySize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BodySize
}

// GetBodySizeOk returns a tuple with the BodySize field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetBodySizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BodySize, true
}

// SetBodySize sets field value
func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetBodySize(v int32) {
	o.BodySize = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetContentType() string {
	if o == nil || isNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetContentTypeOk() (*string, bool) {
	if o == nil || isNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) HasContentType() bool {
	if o != nil && !isNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetContentType(v string) {
	o.ContentType = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetUserAgent() string {
	if o == nil || isNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetUserAgentOk() (*string, bool) {
	if o == nil || isNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) HasUserAgent() bool {
	if o != nil && !isNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetUserIdentity returns the UserIdentity field value if set, zero value otherwise.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetUserIdentity() string {
	if o == nil || isNil(o.UserIdentity) {
		var ret string
		return ret
	}
	return *o.UserIdentity
}

// GetUserIdentityOk returns a tuple with the UserIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetUserIdentityOk() (*string, bool) {
	if o == nil || isNil(o.UserIdentity) {
		return nil, false
	}
	return o.UserIdentity, true
}

// HasUserIdentity returns a boolean if a field has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) HasUserIdentity() bool {
	if o != nil && !isNil(o.UserIdentity) {
		return true
	}

	return false
}

// SetUserIdentity gets a reference to the given string and assigns it to the UserIdentity field.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetUserIdentity(v string) {
	o.UserIdentity = &v
}

// GetReferer returns the Referer field value if set, zero value otherwise.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetReferer() string {
	if o == nil || isNil(o.Referer) {
		var ret string
		return ret
	}
	return *o.Referer
}

// GetRefererOk returns a tuple with the Referer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetRefererOk() (*string, bool) {
	if o == nil || isNil(o.Referer) {
		return nil, false
	}
	return o.Referer, true
}

// HasReferer returns a boolean if a field has been set.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) HasReferer() bool {
	if o != nil && !isNil(o.Referer) {
		return true
	}

	return false
}

// SetReferer gets a reference to the given string and assigns it to the Referer field.
func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetReferer(v string) {
	o.Referer = &v
}

func (o MediaStreamingAccessRecordAllOfRequestMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaStreamingAccessRecordAllOfRequestMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["method"] = o.Method
	toSerialize["url"] = o.Url
	toSerialize["protocolVersion"] = o.ProtocolVersion
	if !isNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	toSerialize["size"] = o.Size
	toSerialize["bodySize"] = o.BodySize
	if !isNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !isNil(o.UserAgent) {
		toSerialize["userAgent"] = o.UserAgent
	}
	if !isNil(o.UserIdentity) {
		toSerialize["userIdentity"] = o.UserIdentity
	}
	if !isNil(o.Referer) {
		toSerialize["referer"] = o.Referer
	}
	return toSerialize, nil
}

type NullableMediaStreamingAccessRecordAllOfRequestMessage struct {
	value *MediaStreamingAccessRecordAllOfRequestMessage
	isSet bool
}

func (v NullableMediaStreamingAccessRecordAllOfRequestMessage) Get() *MediaStreamingAccessRecordAllOfRequestMessage {
	return v.value
}

func (v *NullableMediaStreamingAccessRecordAllOfRequestMessage) Set(val *MediaStreamingAccessRecordAllOfRequestMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaStreamingAccessRecordAllOfRequestMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaStreamingAccessRecordAllOfRequestMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaStreamingAccessRecordAllOfRequestMessage(val *MediaStreamingAccessRecordAllOfRequestMessage) *NullableMediaStreamingAccessRecordAllOfRequestMessage {
	return &NullableMediaStreamingAccessRecordAllOfRequestMessage{value: val, isSet: true}
}

func (v NullableMediaStreamingAccessRecordAllOfRequestMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaStreamingAccessRecordAllOfRequestMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



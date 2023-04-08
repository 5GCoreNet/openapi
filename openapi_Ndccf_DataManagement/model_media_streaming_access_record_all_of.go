/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the MediaStreamingAccessRecordAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaStreamingAccessRecordAllOf{}

// MediaStreamingAccessRecordAllOf struct for MediaStreamingAccessRecordAllOf
type MediaStreamingAccessRecordAllOf struct {
	MediaStreamHandlerEndpointAddress EndpointAddress `json:"mediaStreamHandlerEndpointAddress"`
	ApplicationServerEndpointAddress EndpointAddress `json:"applicationServerEndpointAddress"`
	SessionIdentifier *string `json:"sessionIdentifier,omitempty"`
	RequestMessage MediaStreamingAccessRecordAllOfRequestMessage `json:"requestMessage"`
	CacheStatus *CacheStatus `json:"cacheStatus,omitempty"`
	ResponseMessage MediaStreamingAccessRecordAllOfResponseMessage `json:"responseMessage"`
	// string with format 'float' as defined in OpenAPI.
	ProcessingLatency float32 `json:"processingLatency"`
	ConnectionMetrics *MediaStreamingAccessRecordAllOfConnectionMetrics `json:"connectionMetrics,omitempty"`
}

// NewMediaStreamingAccessRecordAllOf instantiates a new MediaStreamingAccessRecordAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaStreamingAccessRecordAllOf(mediaStreamHandlerEndpointAddress EndpointAddress, applicationServerEndpointAddress EndpointAddress, requestMessage MediaStreamingAccessRecordAllOfRequestMessage, responseMessage MediaStreamingAccessRecordAllOfResponseMessage, processingLatency float32) *MediaStreamingAccessRecordAllOf {
	this := MediaStreamingAccessRecordAllOf{}
	this.MediaStreamHandlerEndpointAddress = mediaStreamHandlerEndpointAddress
	this.ApplicationServerEndpointAddress = applicationServerEndpointAddress
	this.RequestMessage = requestMessage
	this.ResponseMessage = responseMessage
	this.ProcessingLatency = processingLatency
	return &this
}

// NewMediaStreamingAccessRecordAllOfWithDefaults instantiates a new MediaStreamingAccessRecordAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaStreamingAccessRecordAllOfWithDefaults() *MediaStreamingAccessRecordAllOf {
	this := MediaStreamingAccessRecordAllOf{}
	return &this
}

// GetMediaStreamHandlerEndpointAddress returns the MediaStreamHandlerEndpointAddress field value
func (o *MediaStreamingAccessRecordAllOf) GetMediaStreamHandlerEndpointAddress() EndpointAddress {
	if o == nil {
		var ret EndpointAddress
		return ret
	}

	return o.MediaStreamHandlerEndpointAddress
}

// GetMediaStreamHandlerEndpointAddressOk returns a tuple with the MediaStreamHandlerEndpointAddress field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOf) GetMediaStreamHandlerEndpointAddressOk() (*EndpointAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaStreamHandlerEndpointAddress, true
}

// SetMediaStreamHandlerEndpointAddress sets field value
func (o *MediaStreamingAccessRecordAllOf) SetMediaStreamHandlerEndpointAddress(v EndpointAddress) {
	o.MediaStreamHandlerEndpointAddress = v
}

// GetApplicationServerEndpointAddress returns the ApplicationServerEndpointAddress field value
func (o *MediaStreamingAccessRecordAllOf) GetApplicationServerEndpointAddress() EndpointAddress {
	if o == nil {
		var ret EndpointAddress
		return ret
	}

	return o.ApplicationServerEndpointAddress
}

// GetApplicationServerEndpointAddressOk returns a tuple with the ApplicationServerEndpointAddress field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOf) GetApplicationServerEndpointAddressOk() (*EndpointAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationServerEndpointAddress, true
}

// SetApplicationServerEndpointAddress sets field value
func (o *MediaStreamingAccessRecordAllOf) SetApplicationServerEndpointAddress(v EndpointAddress) {
	o.ApplicationServerEndpointAddress = v
}

// GetSessionIdentifier returns the SessionIdentifier field value if set, zero value otherwise.
func (o *MediaStreamingAccessRecordAllOf) GetSessionIdentifier() string {
	if o == nil || isNil(o.SessionIdentifier) {
		var ret string
		return ret
	}
	return *o.SessionIdentifier
}

// GetSessionIdentifierOk returns a tuple with the SessionIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOf) GetSessionIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.SessionIdentifier) {
		return nil, false
	}
	return o.SessionIdentifier, true
}

// HasSessionIdentifier returns a boolean if a field has been set.
func (o *MediaStreamingAccessRecordAllOf) HasSessionIdentifier() bool {
	if o != nil && !isNil(o.SessionIdentifier) {
		return true
	}

	return false
}

// SetSessionIdentifier gets a reference to the given string and assigns it to the SessionIdentifier field.
func (o *MediaStreamingAccessRecordAllOf) SetSessionIdentifier(v string) {
	o.SessionIdentifier = &v
}

// GetRequestMessage returns the RequestMessage field value
func (o *MediaStreamingAccessRecordAllOf) GetRequestMessage() MediaStreamingAccessRecordAllOfRequestMessage {
	if o == nil {
		var ret MediaStreamingAccessRecordAllOfRequestMessage
		return ret
	}

	return o.RequestMessage
}

// GetRequestMessageOk returns a tuple with the RequestMessage field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOf) GetRequestMessageOk() (*MediaStreamingAccessRecordAllOfRequestMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestMessage, true
}

// SetRequestMessage sets field value
func (o *MediaStreamingAccessRecordAllOf) SetRequestMessage(v MediaStreamingAccessRecordAllOfRequestMessage) {
	o.RequestMessage = v
}

// GetCacheStatus returns the CacheStatus field value if set, zero value otherwise.
func (o *MediaStreamingAccessRecordAllOf) GetCacheStatus() CacheStatus {
	if o == nil || isNil(o.CacheStatus) {
		var ret CacheStatus
		return ret
	}
	return *o.CacheStatus
}

// GetCacheStatusOk returns a tuple with the CacheStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOf) GetCacheStatusOk() (*CacheStatus, bool) {
	if o == nil || isNil(o.CacheStatus) {
		return nil, false
	}
	return o.CacheStatus, true
}

// HasCacheStatus returns a boolean if a field has been set.
func (o *MediaStreamingAccessRecordAllOf) HasCacheStatus() bool {
	if o != nil && !isNil(o.CacheStatus) {
		return true
	}

	return false
}

// SetCacheStatus gets a reference to the given CacheStatus and assigns it to the CacheStatus field.
func (o *MediaStreamingAccessRecordAllOf) SetCacheStatus(v CacheStatus) {
	o.CacheStatus = &v
}

// GetResponseMessage returns the ResponseMessage field value
func (o *MediaStreamingAccessRecordAllOf) GetResponseMessage() MediaStreamingAccessRecordAllOfResponseMessage {
	if o == nil {
		var ret MediaStreamingAccessRecordAllOfResponseMessage
		return ret
	}

	return o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOf) GetResponseMessageOk() (*MediaStreamingAccessRecordAllOfResponseMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseMessage, true
}

// SetResponseMessage sets field value
func (o *MediaStreamingAccessRecordAllOf) SetResponseMessage(v MediaStreamingAccessRecordAllOfResponseMessage) {
	o.ResponseMessage = v
}

// GetProcessingLatency returns the ProcessingLatency field value
func (o *MediaStreamingAccessRecordAllOf) GetProcessingLatency() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ProcessingLatency
}

// GetProcessingLatencyOk returns a tuple with the ProcessingLatency field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOf) GetProcessingLatencyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessingLatency, true
}

// SetProcessingLatency sets field value
func (o *MediaStreamingAccessRecordAllOf) SetProcessingLatency(v float32) {
	o.ProcessingLatency = v
}

// GetConnectionMetrics returns the ConnectionMetrics field value if set, zero value otherwise.
func (o *MediaStreamingAccessRecordAllOf) GetConnectionMetrics() MediaStreamingAccessRecordAllOfConnectionMetrics {
	if o == nil || isNil(o.ConnectionMetrics) {
		var ret MediaStreamingAccessRecordAllOfConnectionMetrics
		return ret
	}
	return *o.ConnectionMetrics
}

// GetConnectionMetricsOk returns a tuple with the ConnectionMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecordAllOf) GetConnectionMetricsOk() (*MediaStreamingAccessRecordAllOfConnectionMetrics, bool) {
	if o == nil || isNil(o.ConnectionMetrics) {
		return nil, false
	}
	return o.ConnectionMetrics, true
}

// HasConnectionMetrics returns a boolean if a field has been set.
func (o *MediaStreamingAccessRecordAllOf) HasConnectionMetrics() bool {
	if o != nil && !isNil(o.ConnectionMetrics) {
		return true
	}

	return false
}

// SetConnectionMetrics gets a reference to the given MediaStreamingAccessRecordAllOfConnectionMetrics and assigns it to the ConnectionMetrics field.
func (o *MediaStreamingAccessRecordAllOf) SetConnectionMetrics(v MediaStreamingAccessRecordAllOfConnectionMetrics) {
	o.ConnectionMetrics = &v
}

func (o MediaStreamingAccessRecordAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaStreamingAccessRecordAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mediaStreamHandlerEndpointAddress"] = o.MediaStreamHandlerEndpointAddress
	toSerialize["applicationServerEndpointAddress"] = o.ApplicationServerEndpointAddress
	if !isNil(o.SessionIdentifier) {
		toSerialize["sessionIdentifier"] = o.SessionIdentifier
	}
	toSerialize["requestMessage"] = o.RequestMessage
	if !isNil(o.CacheStatus) {
		toSerialize["cacheStatus"] = o.CacheStatus
	}
	toSerialize["responseMessage"] = o.ResponseMessage
	toSerialize["processingLatency"] = o.ProcessingLatency
	if !isNil(o.ConnectionMetrics) {
		toSerialize["connectionMetrics"] = o.ConnectionMetrics
	}
	return toSerialize, nil
}

type NullableMediaStreamingAccessRecordAllOf struct {
	value *MediaStreamingAccessRecordAllOf
	isSet bool
}

func (v NullableMediaStreamingAccessRecordAllOf) Get() *MediaStreamingAccessRecordAllOf {
	return v.value
}

func (v *NullableMediaStreamingAccessRecordAllOf) Set(val *MediaStreamingAccessRecordAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaStreamingAccessRecordAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaStreamingAccessRecordAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaStreamingAccessRecordAllOf(val *MediaStreamingAccessRecordAllOf) *NullableMediaStreamingAccessRecordAllOf {
	return &NullableMediaStreamingAccessRecordAllOf{value: val, isSet: true}
}

func (v NullableMediaStreamingAccessRecordAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaStreamingAccessRecordAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



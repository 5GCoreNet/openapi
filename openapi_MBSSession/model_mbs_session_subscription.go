/*
3gpp-mbs-session

API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSSession

import (
	"encoding/json"
	"time"
)

// checks if the MbsSessionSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsSessionSubscription{}

// MbsSessionSubscription MBS session subscription
type MbsSessionSubscription struct {
	MbsSessionId *MbsSessionId `json:"mbsSessionId,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	AreaSessionId *int32 `json:"areaSessionId,omitempty"`
	EventList []MbsSessionEvent `json:"eventList"`
	// String providing an URI formatted according to RFC 3986.
	NotifyUri string `json:"notifyUri"`
	NotifyCorrelationId *string `json:"notifyCorrelationId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ExpiryTime *time.Time `json:"expiryTime,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfcInstanceId *string `json:"nfcInstanceId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	MbsSessionSubscUri *string `json:"mbsSessionSubscUri,omitempty"`
}

// NewMbsSessionSubscription instantiates a new MbsSessionSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSessionSubscription(eventList []MbsSessionEvent, notifyUri string) *MbsSessionSubscription {
	this := MbsSessionSubscription{}
	this.EventList = eventList
	this.NotifyUri = notifyUri
	return &this
}

// NewMbsSessionSubscriptionWithDefaults instantiates a new MbsSessionSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSessionSubscriptionWithDefaults() *MbsSessionSubscription {
	this := MbsSessionSubscription{}
	return &this
}

// GetMbsSessionId returns the MbsSessionId field value if set, zero value otherwise.
func (o *MbsSessionSubscription) GetMbsSessionId() MbsSessionId {
	if o == nil || isNil(o.MbsSessionId) {
		var ret MbsSessionId
		return ret
	}
	return *o.MbsSessionId
}

// GetMbsSessionIdOk returns a tuple with the MbsSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionSubscription) GetMbsSessionIdOk() (*MbsSessionId, bool) {
	if o == nil || isNil(o.MbsSessionId) {
		return nil, false
	}
	return o.MbsSessionId, true
}

// HasMbsSessionId returns a boolean if a field has been set.
func (o *MbsSessionSubscription) HasMbsSessionId() bool {
	if o != nil && !isNil(o.MbsSessionId) {
		return true
	}

	return false
}

// SetMbsSessionId gets a reference to the given MbsSessionId and assigns it to the MbsSessionId field.
func (o *MbsSessionSubscription) SetMbsSessionId(v MbsSessionId) {
	o.MbsSessionId = &v
}

// GetAreaSessionId returns the AreaSessionId field value if set, zero value otherwise.
func (o *MbsSessionSubscription) GetAreaSessionId() int32 {
	if o == nil || isNil(o.AreaSessionId) {
		var ret int32
		return ret
	}
	return *o.AreaSessionId
}

// GetAreaSessionIdOk returns a tuple with the AreaSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionSubscription) GetAreaSessionIdOk() (*int32, bool) {
	if o == nil || isNil(o.AreaSessionId) {
		return nil, false
	}
	return o.AreaSessionId, true
}

// HasAreaSessionId returns a boolean if a field has been set.
func (o *MbsSessionSubscription) HasAreaSessionId() bool {
	if o != nil && !isNil(o.AreaSessionId) {
		return true
	}

	return false
}

// SetAreaSessionId gets a reference to the given int32 and assigns it to the AreaSessionId field.
func (o *MbsSessionSubscription) SetAreaSessionId(v int32) {
	o.AreaSessionId = &v
}

// GetEventList returns the EventList field value
func (o *MbsSessionSubscription) GetEventList() []MbsSessionEvent {
	if o == nil {
		var ret []MbsSessionEvent
		return ret
	}

	return o.EventList
}

// GetEventListOk returns a tuple with the EventList field value
// and a boolean to check if the value has been set.
func (o *MbsSessionSubscription) GetEventListOk() ([]MbsSessionEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventList, true
}

// SetEventList sets field value
func (o *MbsSessionSubscription) SetEventList(v []MbsSessionEvent) {
	o.EventList = v
}

// GetNotifyUri returns the NotifyUri field value
func (o *MbsSessionSubscription) GetNotifyUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifyUri
}

// GetNotifyUriOk returns a tuple with the NotifyUri field value
// and a boolean to check if the value has been set.
func (o *MbsSessionSubscription) GetNotifyUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyUri, true
}

// SetNotifyUri sets field value
func (o *MbsSessionSubscription) SetNotifyUri(v string) {
	o.NotifyUri = v
}

// GetNotifyCorrelationId returns the NotifyCorrelationId field value if set, zero value otherwise.
func (o *MbsSessionSubscription) GetNotifyCorrelationId() string {
	if o == nil || isNil(o.NotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrelationId
}

// GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionSubscription) GetNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || isNil(o.NotifyCorrelationId) {
		return nil, false
	}
	return o.NotifyCorrelationId, true
}

// HasNotifyCorrelationId returns a boolean if a field has been set.
func (o *MbsSessionSubscription) HasNotifyCorrelationId() bool {
	if o != nil && !isNil(o.NotifyCorrelationId) {
		return true
	}

	return false
}

// SetNotifyCorrelationId gets a reference to the given string and assigns it to the NotifyCorrelationId field.
func (o *MbsSessionSubscription) SetNotifyCorrelationId(v string) {
	o.NotifyCorrelationId = &v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *MbsSessionSubscription) GetExpiryTime() time.Time {
	if o == nil || isNil(o.ExpiryTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionSubscription) GetExpiryTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiryTime) {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *MbsSessionSubscription) HasExpiryTime() bool {
	if o != nil && !isNil(o.ExpiryTime) {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given time.Time and assigns it to the ExpiryTime field.
func (o *MbsSessionSubscription) SetExpiryTime(v time.Time) {
	o.ExpiryTime = &v
}

// GetNfcInstanceId returns the NfcInstanceId field value if set, zero value otherwise.
func (o *MbsSessionSubscription) GetNfcInstanceId() string {
	if o == nil || isNil(o.NfcInstanceId) {
		var ret string
		return ret
	}
	return *o.NfcInstanceId
}

// GetNfcInstanceIdOk returns a tuple with the NfcInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionSubscription) GetNfcInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.NfcInstanceId) {
		return nil, false
	}
	return o.NfcInstanceId, true
}

// HasNfcInstanceId returns a boolean if a field has been set.
func (o *MbsSessionSubscription) HasNfcInstanceId() bool {
	if o != nil && !isNil(o.NfcInstanceId) {
		return true
	}

	return false
}

// SetNfcInstanceId gets a reference to the given string and assigns it to the NfcInstanceId field.
func (o *MbsSessionSubscription) SetNfcInstanceId(v string) {
	o.NfcInstanceId = &v
}

// GetMbsSessionSubscUri returns the MbsSessionSubscUri field value if set, zero value otherwise.
func (o *MbsSessionSubscription) GetMbsSessionSubscUri() string {
	if o == nil || isNil(o.MbsSessionSubscUri) {
		var ret string
		return ret
	}
	return *o.MbsSessionSubscUri
}

// GetMbsSessionSubscUriOk returns a tuple with the MbsSessionSubscUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionSubscription) GetMbsSessionSubscUriOk() (*string, bool) {
	if o == nil || isNil(o.MbsSessionSubscUri) {
		return nil, false
	}
	return o.MbsSessionSubscUri, true
}

// HasMbsSessionSubscUri returns a boolean if a field has been set.
func (o *MbsSessionSubscription) HasMbsSessionSubscUri() bool {
	if o != nil && !isNil(o.MbsSessionSubscUri) {
		return true
	}

	return false
}

// SetMbsSessionSubscUri gets a reference to the given string and assigns it to the MbsSessionSubscUri field.
func (o *MbsSessionSubscription) SetMbsSessionSubscUri(v string) {
	o.MbsSessionSubscUri = &v
}

func (o MbsSessionSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsSessionSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MbsSessionId) {
		toSerialize["mbsSessionId"] = o.MbsSessionId
	}
	if !isNil(o.AreaSessionId) {
		toSerialize["areaSessionId"] = o.AreaSessionId
	}
	toSerialize["eventList"] = o.EventList
	toSerialize["notifyUri"] = o.NotifyUri
	if !isNil(o.NotifyCorrelationId) {
		toSerialize["notifyCorrelationId"] = o.NotifyCorrelationId
	}
	if !isNil(o.ExpiryTime) {
		toSerialize["expiryTime"] = o.ExpiryTime
	}
	if !isNil(o.NfcInstanceId) {
		toSerialize["nfcInstanceId"] = o.NfcInstanceId
	}
	if !isNil(o.MbsSessionSubscUri) {
		toSerialize["mbsSessionSubscUri"] = o.MbsSessionSubscUri
	}
	return toSerialize, nil
}

type NullableMbsSessionSubscription struct {
	value *MbsSessionSubscription
	isSet bool
}

func (v NullableMbsSessionSubscription) Get() *MbsSessionSubscription {
	return v.value
}

func (v *NullableMbsSessionSubscription) Set(val *MbsSessionSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionSubscription(val *MbsSessionSubscription) *NullableMbsSessionSubscription {
	return &NullableMbsSessionSubscription{value: val, isSet: true}
}

func (v NullableMbsSessionSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



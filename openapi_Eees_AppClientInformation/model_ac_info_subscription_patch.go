/*
EES Application Client Information_API

API for EES Application Client Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_AppClientInformation

import (
	"encoding/json"
	"time"
)

// checks if the ACInfoSubscriptionPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACInfoSubscriptionPatch{}

// ACInfoSubscriptionPatch Represents the partial update of Individual AC Information Subscription.
type ACInfoSubscriptionPatch struct {
	// Filters to retrieve the information about specific ACs.
	AcFltrs []ACFilters `json:"acFltrs,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	ExpTime  *time.Time            `json:"expTime,omitempty"`
	EventReq *ReportingInformation `json:"eventReq,omitempty"`
	// string providing an URI formatted according to IETF RFC 3986.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
}

// NewACInfoSubscriptionPatch instantiates a new ACInfoSubscriptionPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACInfoSubscriptionPatch() *ACInfoSubscriptionPatch {
	this := ACInfoSubscriptionPatch{}
	return &this
}

// NewACInfoSubscriptionPatchWithDefaults instantiates a new ACInfoSubscriptionPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACInfoSubscriptionPatchWithDefaults() *ACInfoSubscriptionPatch {
	this := ACInfoSubscriptionPatch{}
	return &this
}

// GetAcFltrs returns the AcFltrs field value if set, zero value otherwise.
func (o *ACInfoSubscriptionPatch) GetAcFltrs() []ACFilters {
	if o == nil || IsNil(o.AcFltrs) {
		var ret []ACFilters
		return ret
	}
	return o.AcFltrs
}

// GetAcFltrsOk returns a tuple with the AcFltrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACInfoSubscriptionPatch) GetAcFltrsOk() ([]ACFilters, bool) {
	if o == nil || IsNil(o.AcFltrs) {
		return nil, false
	}
	return o.AcFltrs, true
}

// HasAcFltrs returns a boolean if a field has been set.
func (o *ACInfoSubscriptionPatch) HasAcFltrs() bool {
	if o != nil && !IsNil(o.AcFltrs) {
		return true
	}

	return false
}

// SetAcFltrs gets a reference to the given []ACFilters and assigns it to the AcFltrs field.
func (o *ACInfoSubscriptionPatch) SetAcFltrs(v []ACFilters) {
	o.AcFltrs = v
}

// GetExpTime returns the ExpTime field value if set, zero value otherwise.
func (o *ACInfoSubscriptionPatch) GetExpTime() time.Time {
	if o == nil || IsNil(o.ExpTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpTime
}

// GetExpTimeOk returns a tuple with the ExpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACInfoSubscriptionPatch) GetExpTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpTime) {
		return nil, false
	}
	return o.ExpTime, true
}

// HasExpTime returns a boolean if a field has been set.
func (o *ACInfoSubscriptionPatch) HasExpTime() bool {
	if o != nil && !IsNil(o.ExpTime) {
		return true
	}

	return false
}

// SetExpTime gets a reference to the given time.Time and assigns it to the ExpTime field.
func (o *ACInfoSubscriptionPatch) SetExpTime(v time.Time) {
	o.ExpTime = &v
}

// GetEventReq returns the EventReq field value if set, zero value otherwise.
func (o *ACInfoSubscriptionPatch) GetEventReq() ReportingInformation {
	if o == nil || IsNil(o.EventReq) {
		var ret ReportingInformation
		return ret
	}
	return *o.EventReq
}

// GetEventReqOk returns a tuple with the EventReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACInfoSubscriptionPatch) GetEventReqOk() (*ReportingInformation, bool) {
	if o == nil || IsNil(o.EventReq) {
		return nil, false
	}
	return o.EventReq, true
}

// HasEventReq returns a boolean if a field has been set.
func (o *ACInfoSubscriptionPatch) HasEventReq() bool {
	if o != nil && !IsNil(o.EventReq) {
		return true
	}

	return false
}

// SetEventReq gets a reference to the given ReportingInformation and assigns it to the EventReq field.
func (o *ACInfoSubscriptionPatch) SetEventReq(v ReportingInformation) {
	o.EventReq = &v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *ACInfoSubscriptionPatch) GetNotificationDestination() string {
	if o == nil || IsNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACInfoSubscriptionPatch) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationDestination) {
		return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *ACInfoSubscriptionPatch) HasNotificationDestination() bool {
	if o != nil && !IsNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *ACInfoSubscriptionPatch) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

func (o ACInfoSubscriptionPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ACInfoSubscriptionPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcFltrs) {
		toSerialize["acFltrs"] = o.AcFltrs
	}
	if !IsNil(o.ExpTime) {
		toSerialize["expTime"] = o.ExpTime
	}
	if !IsNil(o.EventReq) {
		toSerialize["eventReq"] = o.EventReq
	}
	if !IsNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	return toSerialize, nil
}

type NullableACInfoSubscriptionPatch struct {
	value *ACInfoSubscriptionPatch
	isSet bool
}

func (v NullableACInfoSubscriptionPatch) Get() *ACInfoSubscriptionPatch {
	return v.value
}

func (v *NullableACInfoSubscriptionPatch) Set(val *ACInfoSubscriptionPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableACInfoSubscriptionPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableACInfoSubscriptionPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACInfoSubscriptionPatch(val *ACInfoSubscriptionPatch) *NullableACInfoSubscriptionPatch {
	return &NullableACInfoSubscriptionPatch{value: val, isSet: true}
}

func (v NullableACInfoSubscriptionPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACInfoSubscriptionPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"time"
)

// checks if the ReachabilityReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReachabilityReport{}

// ReachabilityReport struct for ReachabilityReport
type ReachabilityReport struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AmfInstanceId *string `json:"amfInstanceId,omitempty"`
	AccessTypeList []AccessType `json:"accessTypeList,omitempty"`
	Reachability *UeReachability `json:"reachability,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	MaxAvailabilityTime *time.Time `json:"maxAvailabilityTime,omitempty"`
	IdleStatusIndication *IdleStatusIndication `json:"idleStatusIndication,omitempty"`
}

// NewReachabilityReport instantiates a new ReachabilityReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReachabilityReport() *ReachabilityReport {
	this := ReachabilityReport{}
	return &this
}

// NewReachabilityReportWithDefaults instantiates a new ReachabilityReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReachabilityReportWithDefaults() *ReachabilityReport {
	this := ReachabilityReport{}
	return &this
}

// GetAmfInstanceId returns the AmfInstanceId field value if set, zero value otherwise.
func (o *ReachabilityReport) GetAmfInstanceId() string {
	if o == nil || isNil(o.AmfInstanceId) {
		var ret string
		return ret
	}
	return *o.AmfInstanceId
}

// GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReachabilityReport) GetAmfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.AmfInstanceId) {
		return nil, false
	}
	return o.AmfInstanceId, true
}

// HasAmfInstanceId returns a boolean if a field has been set.
func (o *ReachabilityReport) HasAmfInstanceId() bool {
	if o != nil && !isNil(o.AmfInstanceId) {
		return true
	}

	return false
}

// SetAmfInstanceId gets a reference to the given string and assigns it to the AmfInstanceId field.
func (o *ReachabilityReport) SetAmfInstanceId(v string) {
	o.AmfInstanceId = &v
}

// GetAccessTypeList returns the AccessTypeList field value if set, zero value otherwise.
func (o *ReachabilityReport) GetAccessTypeList() []AccessType {
	if o == nil || isNil(o.AccessTypeList) {
		var ret []AccessType
		return ret
	}
	return o.AccessTypeList
}

// GetAccessTypeListOk returns a tuple with the AccessTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReachabilityReport) GetAccessTypeListOk() ([]AccessType, bool) {
	if o == nil || isNil(o.AccessTypeList) {
		return nil, false
	}
	return o.AccessTypeList, true
}

// HasAccessTypeList returns a boolean if a field has been set.
func (o *ReachabilityReport) HasAccessTypeList() bool {
	if o != nil && !isNil(o.AccessTypeList) {
		return true
	}

	return false
}

// SetAccessTypeList gets a reference to the given []AccessType and assigns it to the AccessTypeList field.
func (o *ReachabilityReport) SetAccessTypeList(v []AccessType) {
	o.AccessTypeList = v
}

// GetReachability returns the Reachability field value if set, zero value otherwise.
func (o *ReachabilityReport) GetReachability() UeReachability {
	if o == nil || isNil(o.Reachability) {
		var ret UeReachability
		return ret
	}
	return *o.Reachability
}

// GetReachabilityOk returns a tuple with the Reachability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReachabilityReport) GetReachabilityOk() (*UeReachability, bool) {
	if o == nil || isNil(o.Reachability) {
		return nil, false
	}
	return o.Reachability, true
}

// HasReachability returns a boolean if a field has been set.
func (o *ReachabilityReport) HasReachability() bool {
	if o != nil && !isNil(o.Reachability) {
		return true
	}

	return false
}

// SetReachability gets a reference to the given UeReachability and assigns it to the Reachability field.
func (o *ReachabilityReport) SetReachability(v UeReachability) {
	o.Reachability = &v
}

// GetMaxAvailabilityTime returns the MaxAvailabilityTime field value if set, zero value otherwise.
func (o *ReachabilityReport) GetMaxAvailabilityTime() time.Time {
	if o == nil || isNil(o.MaxAvailabilityTime) {
		var ret time.Time
		return ret
	}
	return *o.MaxAvailabilityTime
}

// GetMaxAvailabilityTimeOk returns a tuple with the MaxAvailabilityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReachabilityReport) GetMaxAvailabilityTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.MaxAvailabilityTime) {
		return nil, false
	}
	return o.MaxAvailabilityTime, true
}

// HasMaxAvailabilityTime returns a boolean if a field has been set.
func (o *ReachabilityReport) HasMaxAvailabilityTime() bool {
	if o != nil && !isNil(o.MaxAvailabilityTime) {
		return true
	}

	return false
}

// SetMaxAvailabilityTime gets a reference to the given time.Time and assigns it to the MaxAvailabilityTime field.
func (o *ReachabilityReport) SetMaxAvailabilityTime(v time.Time) {
	o.MaxAvailabilityTime = &v
}

// GetIdleStatusIndication returns the IdleStatusIndication field value if set, zero value otherwise.
func (o *ReachabilityReport) GetIdleStatusIndication() IdleStatusIndication {
	if o == nil || isNil(o.IdleStatusIndication) {
		var ret IdleStatusIndication
		return ret
	}
	return *o.IdleStatusIndication
}

// GetIdleStatusIndicationOk returns a tuple with the IdleStatusIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReachabilityReport) GetIdleStatusIndicationOk() (*IdleStatusIndication, bool) {
	if o == nil || isNil(o.IdleStatusIndication) {
		return nil, false
	}
	return o.IdleStatusIndication, true
}

// HasIdleStatusIndication returns a boolean if a field has been set.
func (o *ReachabilityReport) HasIdleStatusIndication() bool {
	if o != nil && !isNil(o.IdleStatusIndication) {
		return true
	}

	return false
}

// SetIdleStatusIndication gets a reference to the given IdleStatusIndication and assigns it to the IdleStatusIndication field.
func (o *ReachabilityReport) SetIdleStatusIndication(v IdleStatusIndication) {
	o.IdleStatusIndication = &v
}

func (o ReachabilityReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReachabilityReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AmfInstanceId) {
		toSerialize["amfInstanceId"] = o.AmfInstanceId
	}
	if !isNil(o.AccessTypeList) {
		toSerialize["accessTypeList"] = o.AccessTypeList
	}
	if !isNil(o.Reachability) {
		toSerialize["reachability"] = o.Reachability
	}
	if !isNil(o.MaxAvailabilityTime) {
		toSerialize["maxAvailabilityTime"] = o.MaxAvailabilityTime
	}
	if !isNil(o.IdleStatusIndication) {
		toSerialize["idleStatusIndication"] = o.IdleStatusIndication
	}
	return toSerialize, nil
}

type NullableReachabilityReport struct {
	value *ReachabilityReport
	isSet bool
}

func (v NullableReachabilityReport) Get() *ReachabilityReport {
	return v.value
}

func (v *NullableReachabilityReport) Set(val *ReachabilityReport) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityReport) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityReport(val *ReachabilityReport) *NullableReachabilityReport {
	return &NullableReachabilityReport{value: val, isSet: true}
}

func (v NullableReachabilityReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



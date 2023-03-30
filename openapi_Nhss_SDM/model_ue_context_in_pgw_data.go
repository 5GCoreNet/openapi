/*
Nhss_SDM

HSS Subscriber Data Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_SDM

import (
	"encoding/json"
	"time"
)

// checks if the UeContextInPgwData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeContextInPgwData{}

// UeContextInPgwData Contains data about APNs and PGW-C+SMF FQDNs used in interworking with UDM, and the PGW-C+SMF FQDN to be used for emergency session
type UeContextInPgwData struct {
	PgwInfo []PgwInfo `json:"pgwInfo,omitempty"`
	// Fully Qualified Domain Name
	EmergencyFqdn *string `json:"emergencyFqdn,omitempty"`
	EmergencyPlmnId *PlmnId `json:"emergencyPlmnId,omitempty"`
	EmergencyIpAddr *IpAddress `json:"emergencyIpAddr,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	EmergencyRegistrationTime *time.Time `json:"emergencyRegistrationTime,omitempty"`
}

// NewUeContextInPgwData instantiates a new UeContextInPgwData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeContextInPgwData() *UeContextInPgwData {
	this := UeContextInPgwData{}
	return &this
}

// NewUeContextInPgwDataWithDefaults instantiates a new UeContextInPgwData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeContextInPgwDataWithDefaults() *UeContextInPgwData {
	this := UeContextInPgwData{}
	return &this
}

// GetPgwInfo returns the PgwInfo field value if set, zero value otherwise.
func (o *UeContextInPgwData) GetPgwInfo() []PgwInfo {
	if o == nil || IsNil(o.PgwInfo) {
		var ret []PgwInfo
		return ret
	}
	return o.PgwInfo
}

// GetPgwInfoOk returns a tuple with the PgwInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInPgwData) GetPgwInfoOk() ([]PgwInfo, bool) {
	if o == nil || IsNil(o.PgwInfo) {
		return nil, false
	}
	return o.PgwInfo, true
}

// HasPgwInfo returns a boolean if a field has been set.
func (o *UeContextInPgwData) HasPgwInfo() bool {
	if o != nil && !IsNil(o.PgwInfo) {
		return true
	}

	return false
}

// SetPgwInfo gets a reference to the given []PgwInfo and assigns it to the PgwInfo field.
func (o *UeContextInPgwData) SetPgwInfo(v []PgwInfo) {
	o.PgwInfo = v
}

// GetEmergencyFqdn returns the EmergencyFqdn field value if set, zero value otherwise.
func (o *UeContextInPgwData) GetEmergencyFqdn() string {
	if o == nil || IsNil(o.EmergencyFqdn) {
		var ret string
		return ret
	}
	return *o.EmergencyFqdn
}

// GetEmergencyFqdnOk returns a tuple with the EmergencyFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInPgwData) GetEmergencyFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.EmergencyFqdn) {
		return nil, false
	}
	return o.EmergencyFqdn, true
}

// HasEmergencyFqdn returns a boolean if a field has been set.
func (o *UeContextInPgwData) HasEmergencyFqdn() bool {
	if o != nil && !IsNil(o.EmergencyFqdn) {
		return true
	}

	return false
}

// SetEmergencyFqdn gets a reference to the given string and assigns it to the EmergencyFqdn field.
func (o *UeContextInPgwData) SetEmergencyFqdn(v string) {
	o.EmergencyFqdn = &v
}

// GetEmergencyPlmnId returns the EmergencyPlmnId field value if set, zero value otherwise.
func (o *UeContextInPgwData) GetEmergencyPlmnId() PlmnId {
	if o == nil || IsNil(o.EmergencyPlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.EmergencyPlmnId
}

// GetEmergencyPlmnIdOk returns a tuple with the EmergencyPlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInPgwData) GetEmergencyPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.EmergencyPlmnId) {
		return nil, false
	}
	return o.EmergencyPlmnId, true
}

// HasEmergencyPlmnId returns a boolean if a field has been set.
func (o *UeContextInPgwData) HasEmergencyPlmnId() bool {
	if o != nil && !IsNil(o.EmergencyPlmnId) {
		return true
	}

	return false
}

// SetEmergencyPlmnId gets a reference to the given PlmnId and assigns it to the EmergencyPlmnId field.
func (o *UeContextInPgwData) SetEmergencyPlmnId(v PlmnId) {
	o.EmergencyPlmnId = &v
}

// GetEmergencyIpAddr returns the EmergencyIpAddr field value if set, zero value otherwise.
func (o *UeContextInPgwData) GetEmergencyIpAddr() IpAddress {
	if o == nil || IsNil(o.EmergencyIpAddr) {
		var ret IpAddress
		return ret
	}
	return *o.EmergencyIpAddr
}

// GetEmergencyIpAddrOk returns a tuple with the EmergencyIpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInPgwData) GetEmergencyIpAddrOk() (*IpAddress, bool) {
	if o == nil || IsNil(o.EmergencyIpAddr) {
		return nil, false
	}
	return o.EmergencyIpAddr, true
}

// HasEmergencyIpAddr returns a boolean if a field has been set.
func (o *UeContextInPgwData) HasEmergencyIpAddr() bool {
	if o != nil && !IsNil(o.EmergencyIpAddr) {
		return true
	}

	return false
}

// SetEmergencyIpAddr gets a reference to the given IpAddress and assigns it to the EmergencyIpAddr field.
func (o *UeContextInPgwData) SetEmergencyIpAddr(v IpAddress) {
	o.EmergencyIpAddr = &v
}

// GetEmergencyRegistrationTime returns the EmergencyRegistrationTime field value if set, zero value otherwise.
func (o *UeContextInPgwData) GetEmergencyRegistrationTime() time.Time {
	if o == nil || IsNil(o.EmergencyRegistrationTime) {
		var ret time.Time
		return ret
	}
	return *o.EmergencyRegistrationTime
}

// GetEmergencyRegistrationTimeOk returns a tuple with the EmergencyRegistrationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInPgwData) GetEmergencyRegistrationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EmergencyRegistrationTime) {
		return nil, false
	}
	return o.EmergencyRegistrationTime, true
}

// HasEmergencyRegistrationTime returns a boolean if a field has been set.
func (o *UeContextInPgwData) HasEmergencyRegistrationTime() bool {
	if o != nil && !IsNil(o.EmergencyRegistrationTime) {
		return true
	}

	return false
}

// SetEmergencyRegistrationTime gets a reference to the given time.Time and assigns it to the EmergencyRegistrationTime field.
func (o *UeContextInPgwData) SetEmergencyRegistrationTime(v time.Time) {
	o.EmergencyRegistrationTime = &v
}

func (o UeContextInPgwData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeContextInPgwData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PgwInfo) {
		toSerialize["pgwInfo"] = o.PgwInfo
	}
	if !IsNil(o.EmergencyFqdn) {
		toSerialize["emergencyFqdn"] = o.EmergencyFqdn
	}
	if !IsNil(o.EmergencyPlmnId) {
		toSerialize["emergencyPlmnId"] = o.EmergencyPlmnId
	}
	if !IsNil(o.EmergencyIpAddr) {
		toSerialize["emergencyIpAddr"] = o.EmergencyIpAddr
	}
	if !IsNil(o.EmergencyRegistrationTime) {
		toSerialize["emergencyRegistrationTime"] = o.EmergencyRegistrationTime
	}
	return toSerialize, nil
}

type NullableUeContextInPgwData struct {
	value *UeContextInPgwData
	isSet bool
}

func (v NullableUeContextInPgwData) Get() *UeContextInPgwData {
	return v.value
}

func (v *NullableUeContextInPgwData) Set(val *UeContextInPgwData) {
	v.value = val
	v.isSet = true
}

func (v NullableUeContextInPgwData) IsSet() bool {
	return v.isSet
}

func (v *NullableUeContextInPgwData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeContextInPgwData(val *UeContextInPgwData) *NullableUeContextInPgwData {
	return &NullableUeContextInPgwData{value: val, isSet: true}
}

func (v NullableUeContextInPgwData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeContextInPgwData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


/*
Neasdf_DNSContext

EASDF DNS Context Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Neasdf_DNSContext

import (
	"encoding/json"
	"time"
)

// checks if the DnsContextEventReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsContextEventReport{}

// DnsContextEventReport DNS context event report
type DnsContextEventReport struct {
	// string with format 'date-time' as defined in OpenAPI.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	DnsRuleId      *int32          `json:"dnsRuleId,omitempty"`
	DnsQueryReport *DnsQueryReport `json:"dnsQueryReport,omitempty"`
	DnsRspReport   *DnsRspReport   `json:"dnsRspReport,omitempty"`
	DnsMsgId       *string         `json:"dnsMsgId,omitempty"`
}

// NewDnsContextEventReport instantiates a new DnsContextEventReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsContextEventReport() *DnsContextEventReport {
	this := DnsContextEventReport{}
	return &this
}

// NewDnsContextEventReportWithDefaults instantiates a new DnsContextEventReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsContextEventReportWithDefaults() *DnsContextEventReport {
	this := DnsContextEventReport{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *DnsContextEventReport) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsContextEventReport) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DnsContextEventReport) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *DnsContextEventReport) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetDnsRuleId returns the DnsRuleId field value if set, zero value otherwise.
func (o *DnsContextEventReport) GetDnsRuleId() int32 {
	if o == nil || IsNil(o.DnsRuleId) {
		var ret int32
		return ret
	}
	return *o.DnsRuleId
}

// GetDnsRuleIdOk returns a tuple with the DnsRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsContextEventReport) GetDnsRuleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DnsRuleId) {
		return nil, false
	}
	return o.DnsRuleId, true
}

// HasDnsRuleId returns a boolean if a field has been set.
func (o *DnsContextEventReport) HasDnsRuleId() bool {
	if o != nil && !IsNil(o.DnsRuleId) {
		return true
	}

	return false
}

// SetDnsRuleId gets a reference to the given int32 and assigns it to the DnsRuleId field.
func (o *DnsContextEventReport) SetDnsRuleId(v int32) {
	o.DnsRuleId = &v
}

// GetDnsQueryReport returns the DnsQueryReport field value if set, zero value otherwise.
func (o *DnsContextEventReport) GetDnsQueryReport() DnsQueryReport {
	if o == nil || IsNil(o.DnsQueryReport) {
		var ret DnsQueryReport
		return ret
	}
	return *o.DnsQueryReport
}

// GetDnsQueryReportOk returns a tuple with the DnsQueryReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsContextEventReport) GetDnsQueryReportOk() (*DnsQueryReport, bool) {
	if o == nil || IsNil(o.DnsQueryReport) {
		return nil, false
	}
	return o.DnsQueryReport, true
}

// HasDnsQueryReport returns a boolean if a field has been set.
func (o *DnsContextEventReport) HasDnsQueryReport() bool {
	if o != nil && !IsNil(o.DnsQueryReport) {
		return true
	}

	return false
}

// SetDnsQueryReport gets a reference to the given DnsQueryReport and assigns it to the DnsQueryReport field.
func (o *DnsContextEventReport) SetDnsQueryReport(v DnsQueryReport) {
	o.DnsQueryReport = &v
}

// GetDnsRspReport returns the DnsRspReport field value if set, zero value otherwise.
func (o *DnsContextEventReport) GetDnsRspReport() DnsRspReport {
	if o == nil || IsNil(o.DnsRspReport) {
		var ret DnsRspReport
		return ret
	}
	return *o.DnsRspReport
}

// GetDnsRspReportOk returns a tuple with the DnsRspReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsContextEventReport) GetDnsRspReportOk() (*DnsRspReport, bool) {
	if o == nil || IsNil(o.DnsRspReport) {
		return nil, false
	}
	return o.DnsRspReport, true
}

// HasDnsRspReport returns a boolean if a field has been set.
func (o *DnsContextEventReport) HasDnsRspReport() bool {
	if o != nil && !IsNil(o.DnsRspReport) {
		return true
	}

	return false
}

// SetDnsRspReport gets a reference to the given DnsRspReport and assigns it to the DnsRspReport field.
func (o *DnsContextEventReport) SetDnsRspReport(v DnsRspReport) {
	o.DnsRspReport = &v
}

// GetDnsMsgId returns the DnsMsgId field value if set, zero value otherwise.
func (o *DnsContextEventReport) GetDnsMsgId() string {
	if o == nil || IsNil(o.DnsMsgId) {
		var ret string
		return ret
	}
	return *o.DnsMsgId
}

// GetDnsMsgIdOk returns a tuple with the DnsMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsContextEventReport) GetDnsMsgIdOk() (*string, bool) {
	if o == nil || IsNil(o.DnsMsgId) {
		return nil, false
	}
	return o.DnsMsgId, true
}

// HasDnsMsgId returns a boolean if a field has been set.
func (o *DnsContextEventReport) HasDnsMsgId() bool {
	if o != nil && !IsNil(o.DnsMsgId) {
		return true
	}

	return false
}

// SetDnsMsgId gets a reference to the given string and assigns it to the DnsMsgId field.
func (o *DnsContextEventReport) SetDnsMsgId(v string) {
	o.DnsMsgId = &v
}

func (o DnsContextEventReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsContextEventReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.DnsRuleId) {
		toSerialize["dnsRuleId"] = o.DnsRuleId
	}
	if !IsNil(o.DnsQueryReport) {
		toSerialize["dnsQueryReport"] = o.DnsQueryReport
	}
	if !IsNil(o.DnsRspReport) {
		toSerialize["dnsRspReport"] = o.DnsRspReport
	}
	if !IsNil(o.DnsMsgId) {
		toSerialize["dnsMsgId"] = o.DnsMsgId
	}
	return toSerialize, nil
}

type NullableDnsContextEventReport struct {
	value *DnsContextEventReport
	isSet bool
}

func (v NullableDnsContextEventReport) Get() *DnsContextEventReport {
	return v.value
}

func (v *NullableDnsContextEventReport) Set(val *DnsContextEventReport) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsContextEventReport) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsContextEventReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsContextEventReport(val *DnsContextEventReport) *NullableDnsContextEventReport {
	return &NullableDnsContextEventReport{value: val, isSet: true}
}

func (v NullableDnsContextEventReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsContextEventReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

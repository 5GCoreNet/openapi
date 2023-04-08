/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReporting

import (
	"encoding/json"
)

// checks if the CommunicationRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunicationRecord{}

// CommunicationRecord struct for CommunicationRecord
type CommunicationRecord struct {
	// string with format 'date-time' as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`
	TimeInterval TimeWindow `json:"timeInterval"`
	// Unsigned integer identifying a volume in units of bytes.
	UplinkVolume *int64 `json:"uplinkVolume,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`
}

// NewCommunicationRecord instantiates a new CommunicationRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationRecord(timestamp time.Time, timeInterval TimeWindow) *CommunicationRecord {
	this := CommunicationRecord{}
	this.Timestamp = timestamp
	this.TimeInterval = timeInterval
	return &this
}

// NewCommunicationRecordWithDefaults instantiates a new CommunicationRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationRecordWithDefaults() *CommunicationRecord {
	this := CommunicationRecord{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *CommunicationRecord) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *CommunicationRecord) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *CommunicationRecord) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetTimeInterval returns the TimeInterval field value
func (o *CommunicationRecord) GetTimeInterval() TimeWindow {
	if o == nil {
		var ret TimeWindow
		return ret
	}

	return o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value
// and a boolean to check if the value has been set.
func (o *CommunicationRecord) GetTimeIntervalOk() (*TimeWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeInterval, true
}

// SetTimeInterval sets field value
func (o *CommunicationRecord) SetTimeInterval(v TimeWindow) {
	o.TimeInterval = v
}

// GetUplinkVolume returns the UplinkVolume field value if set, zero value otherwise.
func (o *CommunicationRecord) GetUplinkVolume() int64 {
	if o == nil || isNil(o.UplinkVolume) {
		var ret int64
		return ret
	}
	return *o.UplinkVolume
}

// GetUplinkVolumeOk returns a tuple with the UplinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRecord) GetUplinkVolumeOk() (*int64, bool) {
	if o == nil || isNil(o.UplinkVolume) {
		return nil, false
	}
	return o.UplinkVolume, true
}

// HasUplinkVolume returns a boolean if a field has been set.
func (o *CommunicationRecord) HasUplinkVolume() bool {
	if o != nil && !isNil(o.UplinkVolume) {
		return true
	}

	return false
}

// SetUplinkVolume gets a reference to the given int64 and assigns it to the UplinkVolume field.
func (o *CommunicationRecord) SetUplinkVolume(v int64) {
	o.UplinkVolume = &v
}

// GetDownlinkVolume returns the DownlinkVolume field value if set, zero value otherwise.
func (o *CommunicationRecord) GetDownlinkVolume() int64 {
	if o == nil || isNil(o.DownlinkVolume) {
		var ret int64
		return ret
	}
	return *o.DownlinkVolume
}

// GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRecord) GetDownlinkVolumeOk() (*int64, bool) {
	if o == nil || isNil(o.DownlinkVolume) {
		return nil, false
	}
	return o.DownlinkVolume, true
}

// HasDownlinkVolume returns a boolean if a field has been set.
func (o *CommunicationRecord) HasDownlinkVolume() bool {
	if o != nil && !isNil(o.DownlinkVolume) {
		return true
	}

	return false
}

// SetDownlinkVolume gets a reference to the given int64 and assigns it to the DownlinkVolume field.
func (o *CommunicationRecord) SetDownlinkVolume(v int64) {
	o.DownlinkVolume = &v
}

func (o CommunicationRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunicationRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["timeInterval"] = o.TimeInterval
	if !isNil(o.UplinkVolume) {
		toSerialize["uplinkVolume"] = o.UplinkVolume
	}
	if !isNil(o.DownlinkVolume) {
		toSerialize["downlinkVolume"] = o.DownlinkVolume
	}
	return toSerialize, nil
}

type NullableCommunicationRecord struct {
	value *CommunicationRecord
	isSet bool
}

func (v NullableCommunicationRecord) Get() *CommunicationRecord {
	return v.value
}

func (v *NullableCommunicationRecord) Set(val *CommunicationRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationRecord(val *CommunicationRecord) *NullableCommunicationRecord {
	return &NullableCommunicationRecord{value: val, isSet: true}
}

func (v NullableCommunicationRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"time"
)

// checks if the QosFlowsUsageReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosFlowsUsageReport{}

// QosFlowsUsageReport struct for QosFlowsUsageReport
type QosFlowsUsageReport struct {
	// Unsigned integer identifying a QoS flow, within the range 0 to 63.
	QFI *int32 `json:"qFI,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	StartTimestamp *time.Time `json:"startTimestamp,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	EndTimestamp *time.Time `json:"endTimestamp,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	UplinkVolume *int32 `json:"uplinkVolume,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	DownlinkVolume *int32 `json:"downlinkVolume,omitempty"`
}

// NewQosFlowsUsageReport instantiates a new QosFlowsUsageReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosFlowsUsageReport() *QosFlowsUsageReport {
	this := QosFlowsUsageReport{}
	return &this
}

// NewQosFlowsUsageReportWithDefaults instantiates a new QosFlowsUsageReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosFlowsUsageReportWithDefaults() *QosFlowsUsageReport {
	this := QosFlowsUsageReport{}
	return &this
}

// GetQFI returns the QFI field value if set, zero value otherwise.
func (o *QosFlowsUsageReport) GetQFI() int32 {
	if o == nil || IsNil(o.QFI) {
		var ret int32
		return ret
	}
	return *o.QFI
}

// GetQFIOk returns a tuple with the QFI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowsUsageReport) GetQFIOk() (*int32, bool) {
	if o == nil || IsNil(o.QFI) {
		return nil, false
	}
	return o.QFI, true
}

// HasQFI returns a boolean if a field has been set.
func (o *QosFlowsUsageReport) HasQFI() bool {
	if o != nil && !IsNil(o.QFI) {
		return true
	}

	return false
}

// SetQFI gets a reference to the given int32 and assigns it to the QFI field.
func (o *QosFlowsUsageReport) SetQFI(v int32) {
	o.QFI = &v
}

// GetStartTimestamp returns the StartTimestamp field value if set, zero value otherwise.
func (o *QosFlowsUsageReport) GetStartTimestamp() time.Time {
	if o == nil || IsNil(o.StartTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowsUsageReport) GetStartTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTimestamp) {
		return nil, false
	}
	return o.StartTimestamp, true
}

// HasStartTimestamp returns a boolean if a field has been set.
func (o *QosFlowsUsageReport) HasStartTimestamp() bool {
	if o != nil && !IsNil(o.StartTimestamp) {
		return true
	}

	return false
}

// SetStartTimestamp gets a reference to the given time.Time and assigns it to the StartTimestamp field.
func (o *QosFlowsUsageReport) SetStartTimestamp(v time.Time) {
	o.StartTimestamp = &v
}

// GetEndTimestamp returns the EndTimestamp field value if set, zero value otherwise.
func (o *QosFlowsUsageReport) GetEndTimestamp() time.Time {
	if o == nil || IsNil(o.EndTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.EndTimestamp
}

// GetEndTimestampOk returns a tuple with the EndTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowsUsageReport) GetEndTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTimestamp) {
		return nil, false
	}
	return o.EndTimestamp, true
}

// HasEndTimestamp returns a boolean if a field has been set.
func (o *QosFlowsUsageReport) HasEndTimestamp() bool {
	if o != nil && !IsNil(o.EndTimestamp) {
		return true
	}

	return false
}

// SetEndTimestamp gets a reference to the given time.Time and assigns it to the EndTimestamp field.
func (o *QosFlowsUsageReport) SetEndTimestamp(v time.Time) {
	o.EndTimestamp = &v
}

// GetUplinkVolume returns the UplinkVolume field value if set, zero value otherwise.
func (o *QosFlowsUsageReport) GetUplinkVolume() int32 {
	if o == nil || IsNil(o.UplinkVolume) {
		var ret int32
		return ret
	}
	return *o.UplinkVolume
}

// GetUplinkVolumeOk returns a tuple with the UplinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowsUsageReport) GetUplinkVolumeOk() (*int32, bool) {
	if o == nil || IsNil(o.UplinkVolume) {
		return nil, false
	}
	return o.UplinkVolume, true
}

// HasUplinkVolume returns a boolean if a field has been set.
func (o *QosFlowsUsageReport) HasUplinkVolume() bool {
	if o != nil && !IsNil(o.UplinkVolume) {
		return true
	}

	return false
}

// SetUplinkVolume gets a reference to the given int32 and assigns it to the UplinkVolume field.
func (o *QosFlowsUsageReport) SetUplinkVolume(v int32) {
	o.UplinkVolume = &v
}

// GetDownlinkVolume returns the DownlinkVolume field value if set, zero value otherwise.
func (o *QosFlowsUsageReport) GetDownlinkVolume() int32 {
	if o == nil || IsNil(o.DownlinkVolume) {
		var ret int32
		return ret
	}
	return *o.DownlinkVolume
}

// GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowsUsageReport) GetDownlinkVolumeOk() (*int32, bool) {
	if o == nil || IsNil(o.DownlinkVolume) {
		return nil, false
	}
	return o.DownlinkVolume, true
}

// HasDownlinkVolume returns a boolean if a field has been set.
func (o *QosFlowsUsageReport) HasDownlinkVolume() bool {
	if o != nil && !IsNil(o.DownlinkVolume) {
		return true
	}

	return false
}

// SetDownlinkVolume gets a reference to the given int32 and assigns it to the DownlinkVolume field.
func (o *QosFlowsUsageReport) SetDownlinkVolume(v int32) {
	o.DownlinkVolume = &v
}

func (o QosFlowsUsageReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosFlowsUsageReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QFI) {
		toSerialize["qFI"] = o.QFI
	}
	if !IsNil(o.StartTimestamp) {
		toSerialize["startTimestamp"] = o.StartTimestamp
	}
	if !IsNil(o.EndTimestamp) {
		toSerialize["endTimestamp"] = o.EndTimestamp
	}
	if !IsNil(o.UplinkVolume) {
		toSerialize["uplinkVolume"] = o.UplinkVolume
	}
	if !IsNil(o.DownlinkVolume) {
		toSerialize["downlinkVolume"] = o.DownlinkVolume
	}
	return toSerialize, nil
}

type NullableQosFlowsUsageReport struct {
	value *QosFlowsUsageReport
	isSet bool
}

func (v NullableQosFlowsUsageReport) Get() *QosFlowsUsageReport {
	return v.value
}

func (v *NullableQosFlowsUsageReport) Set(val *QosFlowsUsageReport) {
	v.value = val
	v.isSet = true
}

func (v NullableQosFlowsUsageReport) IsSet() bool {
	return v.isSet
}

func (v *NullableQosFlowsUsageReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosFlowsUsageReport(val *QosFlowsUsageReport) *NullableQosFlowsUsageReport {
	return &NullableQosFlowsUsageReport{value: val, isSet: true}
}

func (v NullableQosFlowsUsageReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosFlowsUsageReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
)

// checks if the AccuUsageReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccuUsageReport{}

// AccuUsageReport Contains the accumulated usage report information.
type AccuUsageReport struct {
	// An id referencing UsageMonitoringData objects associated with this usage report.
	RefUmIds string `json:"refUmIds"`
	// Unsigned integer identifying a volume in units of bytes.
	VolUsage *int64 `json:"volUsage,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	VolUsageUplink *int64 `json:"volUsageUplink,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	VolUsageDownlink *int64 `json:"volUsageDownlink,omitempty"`
	// indicating a time in seconds.
	TimeUsage *int32 `json:"timeUsage,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	NextVolUsage *int64 `json:"nextVolUsage,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	NextVolUsageUplink *int64 `json:"nextVolUsageUplink,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	NextVolUsageDownlink *int64 `json:"nextVolUsageDownlink,omitempty"`
	// indicating a time in seconds.
	NextTimeUsage *int32 `json:"nextTimeUsage,omitempty"`
}

// NewAccuUsageReport instantiates a new AccuUsageReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccuUsageReport(refUmIds string) *AccuUsageReport {
	this := AccuUsageReport{}
	this.RefUmIds = refUmIds
	return &this
}

// NewAccuUsageReportWithDefaults instantiates a new AccuUsageReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccuUsageReportWithDefaults() *AccuUsageReport {
	this := AccuUsageReport{}
	return &this
}

// GetRefUmIds returns the RefUmIds field value
func (o *AccuUsageReport) GetRefUmIds() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefUmIds
}

// GetRefUmIdsOk returns a tuple with the RefUmIds field value
// and a boolean to check if the value has been set.
func (o *AccuUsageReport) GetRefUmIdsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefUmIds, true
}

// SetRefUmIds sets field value
func (o *AccuUsageReport) SetRefUmIds(v string) {
	o.RefUmIds = v
}

// GetVolUsage returns the VolUsage field value if set, zero value otherwise.
func (o *AccuUsageReport) GetVolUsage() int64 {
	if o == nil || IsNil(o.VolUsage) {
		var ret int64
		return ret
	}
	return *o.VolUsage
}

// GetVolUsageOk returns a tuple with the VolUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccuUsageReport) GetVolUsageOk() (*int64, bool) {
	if o == nil || IsNil(o.VolUsage) {
		return nil, false
	}
	return o.VolUsage, true
}

// HasVolUsage returns a boolean if a field has been set.
func (o *AccuUsageReport) HasVolUsage() bool {
	if o != nil && !IsNil(o.VolUsage) {
		return true
	}

	return false
}

// SetVolUsage gets a reference to the given int64 and assigns it to the VolUsage field.
func (o *AccuUsageReport) SetVolUsage(v int64) {
	o.VolUsage = &v
}

// GetVolUsageUplink returns the VolUsageUplink field value if set, zero value otherwise.
func (o *AccuUsageReport) GetVolUsageUplink() int64 {
	if o == nil || IsNil(o.VolUsageUplink) {
		var ret int64
		return ret
	}
	return *o.VolUsageUplink
}

// GetVolUsageUplinkOk returns a tuple with the VolUsageUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccuUsageReport) GetVolUsageUplinkOk() (*int64, bool) {
	if o == nil || IsNil(o.VolUsageUplink) {
		return nil, false
	}
	return o.VolUsageUplink, true
}

// HasVolUsageUplink returns a boolean if a field has been set.
func (o *AccuUsageReport) HasVolUsageUplink() bool {
	if o != nil && !IsNil(o.VolUsageUplink) {
		return true
	}

	return false
}

// SetVolUsageUplink gets a reference to the given int64 and assigns it to the VolUsageUplink field.
func (o *AccuUsageReport) SetVolUsageUplink(v int64) {
	o.VolUsageUplink = &v
}

// GetVolUsageDownlink returns the VolUsageDownlink field value if set, zero value otherwise.
func (o *AccuUsageReport) GetVolUsageDownlink() int64 {
	if o == nil || IsNil(o.VolUsageDownlink) {
		var ret int64
		return ret
	}
	return *o.VolUsageDownlink
}

// GetVolUsageDownlinkOk returns a tuple with the VolUsageDownlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccuUsageReport) GetVolUsageDownlinkOk() (*int64, bool) {
	if o == nil || IsNil(o.VolUsageDownlink) {
		return nil, false
	}
	return o.VolUsageDownlink, true
}

// HasVolUsageDownlink returns a boolean if a field has been set.
func (o *AccuUsageReport) HasVolUsageDownlink() bool {
	if o != nil && !IsNil(o.VolUsageDownlink) {
		return true
	}

	return false
}

// SetVolUsageDownlink gets a reference to the given int64 and assigns it to the VolUsageDownlink field.
func (o *AccuUsageReport) SetVolUsageDownlink(v int64) {
	o.VolUsageDownlink = &v
}

// GetTimeUsage returns the TimeUsage field value if set, zero value otherwise.
func (o *AccuUsageReport) GetTimeUsage() int32 {
	if o == nil || IsNil(o.TimeUsage) {
		var ret int32
		return ret
	}
	return *o.TimeUsage
}

// GetTimeUsageOk returns a tuple with the TimeUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccuUsageReport) GetTimeUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeUsage) {
		return nil, false
	}
	return o.TimeUsage, true
}

// HasTimeUsage returns a boolean if a field has been set.
func (o *AccuUsageReport) HasTimeUsage() bool {
	if o != nil && !IsNil(o.TimeUsage) {
		return true
	}

	return false
}

// SetTimeUsage gets a reference to the given int32 and assigns it to the TimeUsage field.
func (o *AccuUsageReport) SetTimeUsage(v int32) {
	o.TimeUsage = &v
}

// GetNextVolUsage returns the NextVolUsage field value if set, zero value otherwise.
func (o *AccuUsageReport) GetNextVolUsage() int64 {
	if o == nil || IsNil(o.NextVolUsage) {
		var ret int64
		return ret
	}
	return *o.NextVolUsage
}

// GetNextVolUsageOk returns a tuple with the NextVolUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccuUsageReport) GetNextVolUsageOk() (*int64, bool) {
	if o == nil || IsNil(o.NextVolUsage) {
		return nil, false
	}
	return o.NextVolUsage, true
}

// HasNextVolUsage returns a boolean if a field has been set.
func (o *AccuUsageReport) HasNextVolUsage() bool {
	if o != nil && !IsNil(o.NextVolUsage) {
		return true
	}

	return false
}

// SetNextVolUsage gets a reference to the given int64 and assigns it to the NextVolUsage field.
func (o *AccuUsageReport) SetNextVolUsage(v int64) {
	o.NextVolUsage = &v
}

// GetNextVolUsageUplink returns the NextVolUsageUplink field value if set, zero value otherwise.
func (o *AccuUsageReport) GetNextVolUsageUplink() int64 {
	if o == nil || IsNil(o.NextVolUsageUplink) {
		var ret int64
		return ret
	}
	return *o.NextVolUsageUplink
}

// GetNextVolUsageUplinkOk returns a tuple with the NextVolUsageUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccuUsageReport) GetNextVolUsageUplinkOk() (*int64, bool) {
	if o == nil || IsNil(o.NextVolUsageUplink) {
		return nil, false
	}
	return o.NextVolUsageUplink, true
}

// HasNextVolUsageUplink returns a boolean if a field has been set.
func (o *AccuUsageReport) HasNextVolUsageUplink() bool {
	if o != nil && !IsNil(o.NextVolUsageUplink) {
		return true
	}

	return false
}

// SetNextVolUsageUplink gets a reference to the given int64 and assigns it to the NextVolUsageUplink field.
func (o *AccuUsageReport) SetNextVolUsageUplink(v int64) {
	o.NextVolUsageUplink = &v
}

// GetNextVolUsageDownlink returns the NextVolUsageDownlink field value if set, zero value otherwise.
func (o *AccuUsageReport) GetNextVolUsageDownlink() int64 {
	if o == nil || IsNil(o.NextVolUsageDownlink) {
		var ret int64
		return ret
	}
	return *o.NextVolUsageDownlink
}

// GetNextVolUsageDownlinkOk returns a tuple with the NextVolUsageDownlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccuUsageReport) GetNextVolUsageDownlinkOk() (*int64, bool) {
	if o == nil || IsNil(o.NextVolUsageDownlink) {
		return nil, false
	}
	return o.NextVolUsageDownlink, true
}

// HasNextVolUsageDownlink returns a boolean if a field has been set.
func (o *AccuUsageReport) HasNextVolUsageDownlink() bool {
	if o != nil && !IsNil(o.NextVolUsageDownlink) {
		return true
	}

	return false
}

// SetNextVolUsageDownlink gets a reference to the given int64 and assigns it to the NextVolUsageDownlink field.
func (o *AccuUsageReport) SetNextVolUsageDownlink(v int64) {
	o.NextVolUsageDownlink = &v
}

// GetNextTimeUsage returns the NextTimeUsage field value if set, zero value otherwise.
func (o *AccuUsageReport) GetNextTimeUsage() int32 {
	if o == nil || IsNil(o.NextTimeUsage) {
		var ret int32
		return ret
	}
	return *o.NextTimeUsage
}

// GetNextTimeUsageOk returns a tuple with the NextTimeUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccuUsageReport) GetNextTimeUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.NextTimeUsage) {
		return nil, false
	}
	return o.NextTimeUsage, true
}

// HasNextTimeUsage returns a boolean if a field has been set.
func (o *AccuUsageReport) HasNextTimeUsage() bool {
	if o != nil && !IsNil(o.NextTimeUsage) {
		return true
	}

	return false
}

// SetNextTimeUsage gets a reference to the given int32 and assigns it to the NextTimeUsage field.
func (o *AccuUsageReport) SetNextTimeUsage(v int32) {
	o.NextTimeUsage = &v
}

func (o AccuUsageReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccuUsageReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["refUmIds"] = o.RefUmIds
	if !IsNil(o.VolUsage) {
		toSerialize["volUsage"] = o.VolUsage
	}
	if !IsNil(o.VolUsageUplink) {
		toSerialize["volUsageUplink"] = o.VolUsageUplink
	}
	if !IsNil(o.VolUsageDownlink) {
		toSerialize["volUsageDownlink"] = o.VolUsageDownlink
	}
	if !IsNil(o.TimeUsage) {
		toSerialize["timeUsage"] = o.TimeUsage
	}
	if !IsNil(o.NextVolUsage) {
		toSerialize["nextVolUsage"] = o.NextVolUsage
	}
	if !IsNil(o.NextVolUsageUplink) {
		toSerialize["nextVolUsageUplink"] = o.NextVolUsageUplink
	}
	if !IsNil(o.NextVolUsageDownlink) {
		toSerialize["nextVolUsageDownlink"] = o.NextVolUsageDownlink
	}
	if !IsNil(o.NextTimeUsage) {
		toSerialize["nextTimeUsage"] = o.NextTimeUsage
	}
	return toSerialize, nil
}

type NullableAccuUsageReport struct {
	value *AccuUsageReport
	isSet bool
}

func (v NullableAccuUsageReport) Get() *AccuUsageReport {
	return v.value
}

func (v *NullableAccuUsageReport) Set(val *AccuUsageReport) {
	v.value = val
	v.isSet = true
}

func (v NullableAccuUsageReport) IsSet() bool {
	return v.isSet
}

func (v *NullableAccuUsageReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccuUsageReport(val *AccuUsageReport) *NullableAccuUsageReport {
	return &NullableAccuUsageReport{value: val, isSet: true}
}

func (v NullableAccuUsageReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccuUsageReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



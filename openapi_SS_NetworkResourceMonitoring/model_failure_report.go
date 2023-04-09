/*
SS_NetworkResourceMonitoring

API for SEAL Network Resource Monitoring.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceMonitoring

import (
	"encoding/json"
)

// checks if the FailureReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailureReport{}

// FailureReport Represents the failure report indicating the VAL UE(s) or VAL Stream ID(s) for which the NRM server failed to obtain the requested data.
type FailureReport struct {
	// List of VAL UE(s) whose measurement data is not obtained successfully.
	ValUeIds []ValTargetUe `json:"valUeIds,omitempty"`
	// List of VAL stream ID(s) whose measurement data is not obtained successfully.
	ValStreamIds  []string            `json:"valStreamIds,omitempty"`
	FailureReason *FailureReason      `json:"failureReason,omitempty"`
	MeasDataType  MeasurementDataType `json:"measDataType"`
}

// NewFailureReport instantiates a new FailureReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailureReport(measDataType MeasurementDataType) *FailureReport {
	this := FailureReport{}
	this.MeasDataType = measDataType
	return &this
}

// NewFailureReportWithDefaults instantiates a new FailureReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailureReportWithDefaults() *FailureReport {
	this := FailureReport{}
	return &this
}

// GetValUeIds returns the ValUeIds field value if set, zero value otherwise.
func (o *FailureReport) GetValUeIds() []ValTargetUe {
	if o == nil || IsNil(o.ValUeIds) {
		var ret []ValTargetUe
		return ret
	}
	return o.ValUeIds
}

// GetValUeIdsOk returns a tuple with the ValUeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureReport) GetValUeIdsOk() ([]ValTargetUe, bool) {
	if o == nil || IsNil(o.ValUeIds) {
		return nil, false
	}
	return o.ValUeIds, true
}

// HasValUeIds returns a boolean if a field has been set.
func (o *FailureReport) HasValUeIds() bool {
	if o != nil && !IsNil(o.ValUeIds) {
		return true
	}

	return false
}

// SetValUeIds gets a reference to the given []ValTargetUe and assigns it to the ValUeIds field.
func (o *FailureReport) SetValUeIds(v []ValTargetUe) {
	o.ValUeIds = v
}

// GetValStreamIds returns the ValStreamIds field value if set, zero value otherwise.
func (o *FailureReport) GetValStreamIds() []string {
	if o == nil || IsNil(o.ValStreamIds) {
		var ret []string
		return ret
	}
	return o.ValStreamIds
}

// GetValStreamIdsOk returns a tuple with the ValStreamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureReport) GetValStreamIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ValStreamIds) {
		return nil, false
	}
	return o.ValStreamIds, true
}

// HasValStreamIds returns a boolean if a field has been set.
func (o *FailureReport) HasValStreamIds() bool {
	if o != nil && !IsNil(o.ValStreamIds) {
		return true
	}

	return false
}

// SetValStreamIds gets a reference to the given []string and assigns it to the ValStreamIds field.
func (o *FailureReport) SetValStreamIds(v []string) {
	o.ValStreamIds = v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *FailureReport) GetFailureReason() FailureReason {
	if o == nil || IsNil(o.FailureReason) {
		var ret FailureReason
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureReport) GetFailureReasonOk() (*FailureReason, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *FailureReport) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given FailureReason and assigns it to the FailureReason field.
func (o *FailureReport) SetFailureReason(v FailureReason) {
	o.FailureReason = &v
}

// GetMeasDataType returns the MeasDataType field value
func (o *FailureReport) GetMeasDataType() MeasurementDataType {
	if o == nil {
		var ret MeasurementDataType
		return ret
	}

	return o.MeasDataType
}

// GetMeasDataTypeOk returns a tuple with the MeasDataType field value
// and a boolean to check if the value has been set.
func (o *FailureReport) GetMeasDataTypeOk() (*MeasurementDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeasDataType, true
}

// SetMeasDataType sets field value
func (o *FailureReport) SetMeasDataType(v MeasurementDataType) {
	o.MeasDataType = v
}

func (o FailureReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailureReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ValUeIds) {
		toSerialize["valUeIds"] = o.ValUeIds
	}
	if !IsNil(o.ValStreamIds) {
		toSerialize["valStreamIds"] = o.ValStreamIds
	}
	if !IsNil(o.FailureReason) {
		toSerialize["failureReason"] = o.FailureReason
	}
	toSerialize["measDataType"] = o.MeasDataType
	return toSerialize, nil
}

type NullableFailureReport struct {
	value *FailureReport
	isSet bool
}

func (v NullableFailureReport) Get() *FailureReport {
	return v.value
}

func (v *NullableFailureReport) Set(val *FailureReport) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureReport) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureReport(val *FailureReport) *NullableFailureReport {
	return &NullableFailureReport{value: val, isSet: true}
}

func (v NullableFailureReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the RimRSReportInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RimRSReportInfo{}

// RimRSReportInfo struct for RimRSReportInfo
type RimRSReportInfo struct {
	DetectedSetID *int32 `json:"detectedSetID,omitempty"`
	PropagationDelay *int32 `json:"propagationDelay,omitempty"`
	FunctionalityOfRIMRS *string `json:"functionalityOfRIMRS,omitempty"`
}

// NewRimRSReportInfo instantiates a new RimRSReportInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRimRSReportInfo() *RimRSReportInfo {
	this := RimRSReportInfo{}
	return &this
}

// NewRimRSReportInfoWithDefaults instantiates a new RimRSReportInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRimRSReportInfoWithDefaults() *RimRSReportInfo {
	this := RimRSReportInfo{}
	return &this
}

// GetDetectedSetID returns the DetectedSetID field value if set, zero value otherwise.
func (o *RimRSReportInfo) GetDetectedSetID() int32 {
	if o == nil || IsNil(o.DetectedSetID) {
		var ret int32
		return ret
	}
	return *o.DetectedSetID
}

// GetDetectedSetIDOk returns a tuple with the DetectedSetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSReportInfo) GetDetectedSetIDOk() (*int32, bool) {
	if o == nil || IsNil(o.DetectedSetID) {
		return nil, false
	}
	return o.DetectedSetID, true
}

// HasDetectedSetID returns a boolean if a field has been set.
func (o *RimRSReportInfo) HasDetectedSetID() bool {
	if o != nil && !IsNil(o.DetectedSetID) {
		return true
	}

	return false
}

// SetDetectedSetID gets a reference to the given int32 and assigns it to the DetectedSetID field.
func (o *RimRSReportInfo) SetDetectedSetID(v int32) {
	o.DetectedSetID = &v
}

// GetPropagationDelay returns the PropagationDelay field value if set, zero value otherwise.
func (o *RimRSReportInfo) GetPropagationDelay() int32 {
	if o == nil || IsNil(o.PropagationDelay) {
		var ret int32
		return ret
	}
	return *o.PropagationDelay
}

// GetPropagationDelayOk returns a tuple with the PropagationDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSReportInfo) GetPropagationDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.PropagationDelay) {
		return nil, false
	}
	return o.PropagationDelay, true
}

// HasPropagationDelay returns a boolean if a field has been set.
func (o *RimRSReportInfo) HasPropagationDelay() bool {
	if o != nil && !IsNil(o.PropagationDelay) {
		return true
	}

	return false
}

// SetPropagationDelay gets a reference to the given int32 and assigns it to the PropagationDelay field.
func (o *RimRSReportInfo) SetPropagationDelay(v int32) {
	o.PropagationDelay = &v
}

// GetFunctionalityOfRIMRS returns the FunctionalityOfRIMRS field value if set, zero value otherwise.
func (o *RimRSReportInfo) GetFunctionalityOfRIMRS() string {
	if o == nil || IsNil(o.FunctionalityOfRIMRS) {
		var ret string
		return ret
	}
	return *o.FunctionalityOfRIMRS
}

// GetFunctionalityOfRIMRSOk returns a tuple with the FunctionalityOfRIMRS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSReportInfo) GetFunctionalityOfRIMRSOk() (*string, bool) {
	if o == nil || IsNil(o.FunctionalityOfRIMRS) {
		return nil, false
	}
	return o.FunctionalityOfRIMRS, true
}

// HasFunctionalityOfRIMRS returns a boolean if a field has been set.
func (o *RimRSReportInfo) HasFunctionalityOfRIMRS() bool {
	if o != nil && !IsNil(o.FunctionalityOfRIMRS) {
		return true
	}

	return false
}

// SetFunctionalityOfRIMRS gets a reference to the given string and assigns it to the FunctionalityOfRIMRS field.
func (o *RimRSReportInfo) SetFunctionalityOfRIMRS(v string) {
	o.FunctionalityOfRIMRS = &v
}

func (o RimRSReportInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RimRSReportInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DetectedSetID) {
		toSerialize["detectedSetID"] = o.DetectedSetID
	}
	if !IsNil(o.PropagationDelay) {
		toSerialize["propagationDelay"] = o.PropagationDelay
	}
	if !IsNil(o.FunctionalityOfRIMRS) {
		toSerialize["functionalityOfRIMRS"] = o.FunctionalityOfRIMRS
	}
	return toSerialize, nil
}

type NullableRimRSReportInfo struct {
	value *RimRSReportInfo
	isSet bool
}

func (v NullableRimRSReportInfo) Get() *RimRSReportInfo {
	return v.value
}

func (v *NullableRimRSReportInfo) Set(val *RimRSReportInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRimRSReportInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRimRSReportInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRimRSReportInfo(val *RimRSReportInfo) *NullableRimRSReportInfo {
	return &NullableRimRSReportInfo{value: val, isSet: true}
}

func (v NullableRimRSReportInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRimRSReportInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



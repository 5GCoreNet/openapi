/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
)

// checks if the FailureCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailureCause{}

// FailureCause Represents the reason of communication failure.
type FailureCause struct {
	// Identifies a non-transparent copy of the BSSGP cause code. Refer to 3GPP TS 29.128.
	BssgpCause *int32 `json:"bssgpCause,omitempty"`
	// Identify the type of the S1AP-Cause. Refer to 3GPP TS 29.128.
	CauseType *int32 `json:"causeType,omitempty"`
	// Identifies a non-transparent copy of the GMM cause code. Refer to 3GPP TS 29.128.
	GmmCause *int32 `json:"gmmCause,omitempty"`
	// Identifies a non-transparent copy of the RANAP cause code. Refer to 3GPP TS 29.128.
	RanapCause *int32 `json:"ranapCause,omitempty"`
	// Indicates RAN and/or NAS release cause code information, TWAN release cause code information or untrusted WLAN release cause code information. Refer to 3GPP TS 29.214.
	RanNasCause *string `json:"ranNasCause,omitempty"`
	// Identifies a non-transparent copy of the S1AP cause code. Refer to 3GPP TS 29.128.
	S1ApCause *int32 `json:"s1ApCause,omitempty"`
	// Identifies a non-transparent copy of the SM cause code. Refer to 3GPP TS 29.128.
	SmCause *int32 `json:"smCause,omitempty"`
}

// NewFailureCause instantiates a new FailureCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailureCause() *FailureCause {
	this := FailureCause{}
	return &this
}

// NewFailureCauseWithDefaults instantiates a new FailureCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailureCauseWithDefaults() *FailureCause {
	this := FailureCause{}
	return &this
}

// GetBssgpCause returns the BssgpCause field value if set, zero value otherwise.
func (o *FailureCause) GetBssgpCause() int32 {
	if o == nil || IsNil(o.BssgpCause) {
		var ret int32
		return ret
	}
	return *o.BssgpCause
}

// GetBssgpCauseOk returns a tuple with the BssgpCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureCause) GetBssgpCauseOk() (*int32, bool) {
	if o == nil || IsNil(o.BssgpCause) {
		return nil, false
	}
	return o.BssgpCause, true
}

// HasBssgpCause returns a boolean if a field has been set.
func (o *FailureCause) HasBssgpCause() bool {
	if o != nil && !IsNil(o.BssgpCause) {
		return true
	}

	return false
}

// SetBssgpCause gets a reference to the given int32 and assigns it to the BssgpCause field.
func (o *FailureCause) SetBssgpCause(v int32) {
	o.BssgpCause = &v
}

// GetCauseType returns the CauseType field value if set, zero value otherwise.
func (o *FailureCause) GetCauseType() int32 {
	if o == nil || IsNil(o.CauseType) {
		var ret int32
		return ret
	}
	return *o.CauseType
}

// GetCauseTypeOk returns a tuple with the CauseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureCause) GetCauseTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.CauseType) {
		return nil, false
	}
	return o.CauseType, true
}

// HasCauseType returns a boolean if a field has been set.
func (o *FailureCause) HasCauseType() bool {
	if o != nil && !IsNil(o.CauseType) {
		return true
	}

	return false
}

// SetCauseType gets a reference to the given int32 and assigns it to the CauseType field.
func (o *FailureCause) SetCauseType(v int32) {
	o.CauseType = &v
}

// GetGmmCause returns the GmmCause field value if set, zero value otherwise.
func (o *FailureCause) GetGmmCause() int32 {
	if o == nil || IsNil(o.GmmCause) {
		var ret int32
		return ret
	}
	return *o.GmmCause
}

// GetGmmCauseOk returns a tuple with the GmmCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureCause) GetGmmCauseOk() (*int32, bool) {
	if o == nil || IsNil(o.GmmCause) {
		return nil, false
	}
	return o.GmmCause, true
}

// HasGmmCause returns a boolean if a field has been set.
func (o *FailureCause) HasGmmCause() bool {
	if o != nil && !IsNil(o.GmmCause) {
		return true
	}

	return false
}

// SetGmmCause gets a reference to the given int32 and assigns it to the GmmCause field.
func (o *FailureCause) SetGmmCause(v int32) {
	o.GmmCause = &v
}

// GetRanapCause returns the RanapCause field value if set, zero value otherwise.
func (o *FailureCause) GetRanapCause() int32 {
	if o == nil || IsNil(o.RanapCause) {
		var ret int32
		return ret
	}
	return *o.RanapCause
}

// GetRanapCauseOk returns a tuple with the RanapCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureCause) GetRanapCauseOk() (*int32, bool) {
	if o == nil || IsNil(o.RanapCause) {
		return nil, false
	}
	return o.RanapCause, true
}

// HasRanapCause returns a boolean if a field has been set.
func (o *FailureCause) HasRanapCause() bool {
	if o != nil && !IsNil(o.RanapCause) {
		return true
	}

	return false
}

// SetRanapCause gets a reference to the given int32 and assigns it to the RanapCause field.
func (o *FailureCause) SetRanapCause(v int32) {
	o.RanapCause = &v
}

// GetRanNasCause returns the RanNasCause field value if set, zero value otherwise.
func (o *FailureCause) GetRanNasCause() string {
	if o == nil || IsNil(o.RanNasCause) {
		var ret string
		return ret
	}
	return *o.RanNasCause
}

// GetRanNasCauseOk returns a tuple with the RanNasCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureCause) GetRanNasCauseOk() (*string, bool) {
	if o == nil || IsNil(o.RanNasCause) {
		return nil, false
	}
	return o.RanNasCause, true
}

// HasRanNasCause returns a boolean if a field has been set.
func (o *FailureCause) HasRanNasCause() bool {
	if o != nil && !IsNil(o.RanNasCause) {
		return true
	}

	return false
}

// SetRanNasCause gets a reference to the given string and assigns it to the RanNasCause field.
func (o *FailureCause) SetRanNasCause(v string) {
	o.RanNasCause = &v
}

// GetS1ApCause returns the S1ApCause field value if set, zero value otherwise.
func (o *FailureCause) GetS1ApCause() int32 {
	if o == nil || IsNil(o.S1ApCause) {
		var ret int32
		return ret
	}
	return *o.S1ApCause
}

// GetS1ApCauseOk returns a tuple with the S1ApCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureCause) GetS1ApCauseOk() (*int32, bool) {
	if o == nil || IsNil(o.S1ApCause) {
		return nil, false
	}
	return o.S1ApCause, true
}

// HasS1ApCause returns a boolean if a field has been set.
func (o *FailureCause) HasS1ApCause() bool {
	if o != nil && !IsNil(o.S1ApCause) {
		return true
	}

	return false
}

// SetS1ApCause gets a reference to the given int32 and assigns it to the S1ApCause field.
func (o *FailureCause) SetS1ApCause(v int32) {
	o.S1ApCause = &v
}

// GetSmCause returns the SmCause field value if set, zero value otherwise.
func (o *FailureCause) GetSmCause() int32 {
	if o == nil || IsNil(o.SmCause) {
		var ret int32
		return ret
	}
	return *o.SmCause
}

// GetSmCauseOk returns a tuple with the SmCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailureCause) GetSmCauseOk() (*int32, bool) {
	if o == nil || IsNil(o.SmCause) {
		return nil, false
	}
	return o.SmCause, true
}

// HasSmCause returns a boolean if a field has been set.
func (o *FailureCause) HasSmCause() bool {
	if o != nil && !IsNil(o.SmCause) {
		return true
	}

	return false
}

// SetSmCause gets a reference to the given int32 and assigns it to the SmCause field.
func (o *FailureCause) SetSmCause(v int32) {
	o.SmCause = &v
}

func (o FailureCause) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailureCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BssgpCause) {
		toSerialize["bssgpCause"] = o.BssgpCause
	}
	if !IsNil(o.CauseType) {
		toSerialize["causeType"] = o.CauseType
	}
	if !IsNil(o.GmmCause) {
		toSerialize["gmmCause"] = o.GmmCause
	}
	if !IsNil(o.RanapCause) {
		toSerialize["ranapCause"] = o.RanapCause
	}
	if !IsNil(o.RanNasCause) {
		toSerialize["ranNasCause"] = o.RanNasCause
	}
	if !IsNil(o.S1ApCause) {
		toSerialize["s1ApCause"] = o.S1ApCause
	}
	if !IsNil(o.SmCause) {
		toSerialize["smCause"] = o.SmCause
	}
	return toSerialize, nil
}

type NullableFailureCause struct {
	value *FailureCause
	isSet bool
}

func (v NullableFailureCause) Get() *FailureCause {
	return v.value
}

func (v *NullableFailureCause) Set(val *FailureCause) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureCause) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureCause(val *FailureCause) *NullableFailureCause {
	return &NullableFailureCause{value: val, isSet: true}
}

func (v NullableFailureCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

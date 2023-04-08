/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"time"
)

// checks if the HsmfUpdateError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HsmfUpdateError{}

// HsmfUpdateError Error within Update Response from H-SMF, or from SMF to I-SMF
type HsmfUpdateError struct {
	Error ProblemDetails `json:"error"`
	// Procedure Transaction Identifier
	Pti *int32 `json:"pti,omitempty"`
	N1smCause *string `json:"n1smCause,omitempty"`
	N1SmInfoToUe *RefToBinaryData `json:"n1SmInfoToUe,omitempty"`
	// indicating a time in seconds.
	BackOffTimer *int32 `json:"backOffTimer,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime *time.Time `json:"recoveryTime,omitempty"`
}

// NewHsmfUpdateError instantiates a new HsmfUpdateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHsmfUpdateError(error_ ProblemDetails) *HsmfUpdateError {
	this := HsmfUpdateError{}
	this.Error = error_
	return &this
}

// NewHsmfUpdateErrorWithDefaults instantiates a new HsmfUpdateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHsmfUpdateErrorWithDefaults() *HsmfUpdateError {
	this := HsmfUpdateError{}
	return &this
}

// GetError returns the Error field value
func (o *HsmfUpdateError) GetError() ProblemDetails {
	if o == nil {
		var ret ProblemDetails
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *HsmfUpdateError) GetErrorOk() (*ProblemDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *HsmfUpdateError) SetError(v ProblemDetails) {
	o.Error = v
}

// GetPti returns the Pti field value if set, zero value otherwise.
func (o *HsmfUpdateError) GetPti() int32 {
	if o == nil || isNil(o.Pti) {
		var ret int32
		return ret
	}
	return *o.Pti
}

// GetPtiOk returns a tuple with the Pti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdateError) GetPtiOk() (*int32, bool) {
	if o == nil || isNil(o.Pti) {
		return nil, false
	}
	return o.Pti, true
}

// HasPti returns a boolean if a field has been set.
func (o *HsmfUpdateError) HasPti() bool {
	if o != nil && !isNil(o.Pti) {
		return true
	}

	return false
}

// SetPti gets a reference to the given int32 and assigns it to the Pti field.
func (o *HsmfUpdateError) SetPti(v int32) {
	o.Pti = &v
}

// GetN1smCause returns the N1smCause field value if set, zero value otherwise.
func (o *HsmfUpdateError) GetN1smCause() string {
	if o == nil || isNil(o.N1smCause) {
		var ret string
		return ret
	}
	return *o.N1smCause
}

// GetN1smCauseOk returns a tuple with the N1smCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdateError) GetN1smCauseOk() (*string, bool) {
	if o == nil || isNil(o.N1smCause) {
		return nil, false
	}
	return o.N1smCause, true
}

// HasN1smCause returns a boolean if a field has been set.
func (o *HsmfUpdateError) HasN1smCause() bool {
	if o != nil && !isNil(o.N1smCause) {
		return true
	}

	return false
}

// SetN1smCause gets a reference to the given string and assigns it to the N1smCause field.
func (o *HsmfUpdateError) SetN1smCause(v string) {
	o.N1smCause = &v
}

// GetN1SmInfoToUe returns the N1SmInfoToUe field value if set, zero value otherwise.
func (o *HsmfUpdateError) GetN1SmInfoToUe() RefToBinaryData {
	if o == nil || isNil(o.N1SmInfoToUe) {
		var ret RefToBinaryData
		return ret
	}
	return *o.N1SmInfoToUe
}

// GetN1SmInfoToUeOk returns a tuple with the N1SmInfoToUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdateError) GetN1SmInfoToUeOk() (*RefToBinaryData, bool) {
	if o == nil || isNil(o.N1SmInfoToUe) {
		return nil, false
	}
	return o.N1SmInfoToUe, true
}

// HasN1SmInfoToUe returns a boolean if a field has been set.
func (o *HsmfUpdateError) HasN1SmInfoToUe() bool {
	if o != nil && !isNil(o.N1SmInfoToUe) {
		return true
	}

	return false
}

// SetN1SmInfoToUe gets a reference to the given RefToBinaryData and assigns it to the N1SmInfoToUe field.
func (o *HsmfUpdateError) SetN1SmInfoToUe(v RefToBinaryData) {
	o.N1SmInfoToUe = &v
}

// GetBackOffTimer returns the BackOffTimer field value if set, zero value otherwise.
func (o *HsmfUpdateError) GetBackOffTimer() int32 {
	if o == nil || isNil(o.BackOffTimer) {
		var ret int32
		return ret
	}
	return *o.BackOffTimer
}

// GetBackOffTimerOk returns a tuple with the BackOffTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdateError) GetBackOffTimerOk() (*int32, bool) {
	if o == nil || isNil(o.BackOffTimer) {
		return nil, false
	}
	return o.BackOffTimer, true
}

// HasBackOffTimer returns a boolean if a field has been set.
func (o *HsmfUpdateError) HasBackOffTimer() bool {
	if o != nil && !isNil(o.BackOffTimer) {
		return true
	}

	return false
}

// SetBackOffTimer gets a reference to the given int32 and assigns it to the BackOffTimer field.
func (o *HsmfUpdateError) SetBackOffTimer(v int32) {
	o.BackOffTimer = &v
}

// GetRecoveryTime returns the RecoveryTime field value if set, zero value otherwise.
func (o *HsmfUpdateError) GetRecoveryTime() time.Time {
	if o == nil || isNil(o.RecoveryTime) {
		var ret time.Time
		return ret
	}
	return *o.RecoveryTime
}

// GetRecoveryTimeOk returns a tuple with the RecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdateError) GetRecoveryTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.RecoveryTime) {
		return nil, false
	}
	return o.RecoveryTime, true
}

// HasRecoveryTime returns a boolean if a field has been set.
func (o *HsmfUpdateError) HasRecoveryTime() bool {
	if o != nil && !isNil(o.RecoveryTime) {
		return true
	}

	return false
}

// SetRecoveryTime gets a reference to the given time.Time and assigns it to the RecoveryTime field.
func (o *HsmfUpdateError) SetRecoveryTime(v time.Time) {
	o.RecoveryTime = &v
}

func (o HsmfUpdateError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HsmfUpdateError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	if !isNil(o.Pti) {
		toSerialize["pti"] = o.Pti
	}
	if !isNil(o.N1smCause) {
		toSerialize["n1smCause"] = o.N1smCause
	}
	if !isNil(o.N1SmInfoToUe) {
		toSerialize["n1SmInfoToUe"] = o.N1SmInfoToUe
	}
	if !isNil(o.BackOffTimer) {
		toSerialize["backOffTimer"] = o.BackOffTimer
	}
	if !isNil(o.RecoveryTime) {
		toSerialize["recoveryTime"] = o.RecoveryTime
	}
	return toSerialize, nil
}

type NullableHsmfUpdateError struct {
	value *HsmfUpdateError
	isSet bool
}

func (v NullableHsmfUpdateError) Get() *HsmfUpdateError {
	return v.value
}

func (v *NullableHsmfUpdateError) Set(val *HsmfUpdateError) {
	v.value = val
	v.isSet = true
}

func (v NullableHsmfUpdateError) IsSet() bool {
	return v.isSet
}

func (v *NullableHsmfUpdateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHsmfUpdateError(val *HsmfUpdateError) *NullableHsmfUpdateError {
	return &NullableHsmfUpdateError{value: val, isSet: true}
}

func (v NullableHsmfUpdateError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHsmfUpdateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



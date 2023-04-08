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

// checks if the Trigger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Trigger{}

// Trigger struct for Trigger
type Trigger struct {
	TriggerType TriggerType `json:"triggerType"`
	TriggerCategory TriggerCategory `json:"triggerCategory"`
	// indicating a time in seconds.
	TimeLimit *int32 `json:"timeLimit,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	VolumeLimit *int32 `json:"volumeLimit,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	VolumeLimit64 *int32 `json:"volumeLimit64,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	EventLimit *int32 `json:"eventLimit,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	MaxNumberOfccc *int32 `json:"maxNumberOfccc,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TariffTimeChange *time.Time `json:"tariffTimeChange,omitempty"`
}

// NewTrigger instantiates a new Trigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrigger(triggerType TriggerType, triggerCategory TriggerCategory) *Trigger {
	this := Trigger{}
	this.TriggerType = triggerType
	this.TriggerCategory = triggerCategory
	return &this
}

// NewTriggerWithDefaults instantiates a new Trigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerWithDefaults() *Trigger {
	this := Trigger{}
	return &this
}

// GetTriggerType returns the TriggerType field value
func (o *Trigger) GetTriggerType() TriggerType {
	if o == nil {
		var ret TriggerType
		return ret
	}

	return o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value
// and a boolean to check if the value has been set.
func (o *Trigger) GetTriggerTypeOk() (*TriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerType, true
}

// SetTriggerType sets field value
func (o *Trigger) SetTriggerType(v TriggerType) {
	o.TriggerType = v
}

// GetTriggerCategory returns the TriggerCategory field value
func (o *Trigger) GetTriggerCategory() TriggerCategory {
	if o == nil {
		var ret TriggerCategory
		return ret
	}

	return o.TriggerCategory
}

// GetTriggerCategoryOk returns a tuple with the TriggerCategory field value
// and a boolean to check if the value has been set.
func (o *Trigger) GetTriggerCategoryOk() (*TriggerCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerCategory, true
}

// SetTriggerCategory sets field value
func (o *Trigger) SetTriggerCategory(v TriggerCategory) {
	o.TriggerCategory = v
}

// GetTimeLimit returns the TimeLimit field value if set, zero value otherwise.
func (o *Trigger) GetTimeLimit() int32 {
	if o == nil || isNil(o.TimeLimit) {
		var ret int32
		return ret
	}
	return *o.TimeLimit
}

// GetTimeLimitOk returns a tuple with the TimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trigger) GetTimeLimitOk() (*int32, bool) {
	if o == nil || isNil(o.TimeLimit) {
		return nil, false
	}
	return o.TimeLimit, true
}

// HasTimeLimit returns a boolean if a field has been set.
func (o *Trigger) HasTimeLimit() bool {
	if o != nil && !isNil(o.TimeLimit) {
		return true
	}

	return false
}

// SetTimeLimit gets a reference to the given int32 and assigns it to the TimeLimit field.
func (o *Trigger) SetTimeLimit(v int32) {
	o.TimeLimit = &v
}

// GetVolumeLimit returns the VolumeLimit field value if set, zero value otherwise.
func (o *Trigger) GetVolumeLimit() int32 {
	if o == nil || isNil(o.VolumeLimit) {
		var ret int32
		return ret
	}
	return *o.VolumeLimit
}

// GetVolumeLimitOk returns a tuple with the VolumeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trigger) GetVolumeLimitOk() (*int32, bool) {
	if o == nil || isNil(o.VolumeLimit) {
		return nil, false
	}
	return o.VolumeLimit, true
}

// HasVolumeLimit returns a boolean if a field has been set.
func (o *Trigger) HasVolumeLimit() bool {
	if o != nil && !isNil(o.VolumeLimit) {
		return true
	}

	return false
}

// SetVolumeLimit gets a reference to the given int32 and assigns it to the VolumeLimit field.
func (o *Trigger) SetVolumeLimit(v int32) {
	o.VolumeLimit = &v
}

// GetVolumeLimit64 returns the VolumeLimit64 field value if set, zero value otherwise.
func (o *Trigger) GetVolumeLimit64() int32 {
	if o == nil || isNil(o.VolumeLimit64) {
		var ret int32
		return ret
	}
	return *o.VolumeLimit64
}

// GetVolumeLimit64Ok returns a tuple with the VolumeLimit64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trigger) GetVolumeLimit64Ok() (*int32, bool) {
	if o == nil || isNil(o.VolumeLimit64) {
		return nil, false
	}
	return o.VolumeLimit64, true
}

// HasVolumeLimit64 returns a boolean if a field has been set.
func (o *Trigger) HasVolumeLimit64() bool {
	if o != nil && !isNil(o.VolumeLimit64) {
		return true
	}

	return false
}

// SetVolumeLimit64 gets a reference to the given int32 and assigns it to the VolumeLimit64 field.
func (o *Trigger) SetVolumeLimit64(v int32) {
	o.VolumeLimit64 = &v
}

// GetEventLimit returns the EventLimit field value if set, zero value otherwise.
func (o *Trigger) GetEventLimit() int32 {
	if o == nil || isNil(o.EventLimit) {
		var ret int32
		return ret
	}
	return *o.EventLimit
}

// GetEventLimitOk returns a tuple with the EventLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trigger) GetEventLimitOk() (*int32, bool) {
	if o == nil || isNil(o.EventLimit) {
		return nil, false
	}
	return o.EventLimit, true
}

// HasEventLimit returns a boolean if a field has been set.
func (o *Trigger) HasEventLimit() bool {
	if o != nil && !isNil(o.EventLimit) {
		return true
	}

	return false
}

// SetEventLimit gets a reference to the given int32 and assigns it to the EventLimit field.
func (o *Trigger) SetEventLimit(v int32) {
	o.EventLimit = &v
}

// GetMaxNumberOfccc returns the MaxNumberOfccc field value if set, zero value otherwise.
func (o *Trigger) GetMaxNumberOfccc() int32 {
	if o == nil || isNil(o.MaxNumberOfccc) {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfccc
}

// GetMaxNumberOfcccOk returns a tuple with the MaxNumberOfccc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trigger) GetMaxNumberOfcccOk() (*int32, bool) {
	if o == nil || isNil(o.MaxNumberOfccc) {
		return nil, false
	}
	return o.MaxNumberOfccc, true
}

// HasMaxNumberOfccc returns a boolean if a field has been set.
func (o *Trigger) HasMaxNumberOfccc() bool {
	if o != nil && !isNil(o.MaxNumberOfccc) {
		return true
	}

	return false
}

// SetMaxNumberOfccc gets a reference to the given int32 and assigns it to the MaxNumberOfccc field.
func (o *Trigger) SetMaxNumberOfccc(v int32) {
	o.MaxNumberOfccc = &v
}

// GetTariffTimeChange returns the TariffTimeChange field value if set, zero value otherwise.
func (o *Trigger) GetTariffTimeChange() time.Time {
	if o == nil || isNil(o.TariffTimeChange) {
		var ret time.Time
		return ret
	}
	return *o.TariffTimeChange
}

// GetTariffTimeChangeOk returns a tuple with the TariffTimeChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trigger) GetTariffTimeChangeOk() (*time.Time, bool) {
	if o == nil || isNil(o.TariffTimeChange) {
		return nil, false
	}
	return o.TariffTimeChange, true
}

// HasTariffTimeChange returns a boolean if a field has been set.
func (o *Trigger) HasTariffTimeChange() bool {
	if o != nil && !isNil(o.TariffTimeChange) {
		return true
	}

	return false
}

// SetTariffTimeChange gets a reference to the given time.Time and assigns it to the TariffTimeChange field.
func (o *Trigger) SetTariffTimeChange(v time.Time) {
	o.TariffTimeChange = &v
}

func (o Trigger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Trigger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["triggerType"] = o.TriggerType
	toSerialize["triggerCategory"] = o.TriggerCategory
	if !isNil(o.TimeLimit) {
		toSerialize["timeLimit"] = o.TimeLimit
	}
	if !isNil(o.VolumeLimit) {
		toSerialize["volumeLimit"] = o.VolumeLimit
	}
	if !isNil(o.VolumeLimit64) {
		toSerialize["volumeLimit64"] = o.VolumeLimit64
	}
	if !isNil(o.EventLimit) {
		toSerialize["eventLimit"] = o.EventLimit
	}
	if !isNil(o.MaxNumberOfccc) {
		toSerialize["maxNumberOfccc"] = o.MaxNumberOfccc
	}
	if !isNil(o.TariffTimeChange) {
		toSerialize["tariffTimeChange"] = o.TariffTimeChange
	}
	return toSerialize, nil
}

type NullableTrigger struct {
	value *Trigger
	isSet bool
}

func (v NullableTrigger) Get() *Trigger {
	return v.value
}

func (v *NullableTrigger) Set(val *Trigger) {
	v.value = val
	v.isSet = true
}

func (v NullableTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrigger(val *Trigger) *NullableTrigger {
	return &NullableTrigger{value: val, isSet: true}
}

func (v NullableTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



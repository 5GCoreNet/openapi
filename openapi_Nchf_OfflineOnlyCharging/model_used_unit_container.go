/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_OfflineOnlyCharging

import (
	"encoding/json"
	"time"
)

// checks if the UsedUnitContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsedUnitContainer{}

// UsedUnitContainer struct for UsedUnitContainer
type UsedUnitContainer struct {
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	ServiceId *int32    `json:"serviceId,omitempty"`
	Triggers  []Trigger `json:"triggers,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TriggerTimestamp *time.Time `json:"triggerTimestamp,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	Time *int32 `json:"time,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	TotalVolume *int32 `json:"totalVolume,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	UplinkVolume *int32 `json:"uplinkVolume,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	DownlinkVolume *int32 `json:"downlinkVolume,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	ServiceSpecificUnits    *int32                   `json:"serviceSpecificUnits,omitempty"`
	EventTimeStamps         []time.Time              `json:"eventTimeStamps,omitempty"`
	LocalSequenceNumber     int32                    `json:"localSequenceNumber"`
	PDUContainerInformation *PDUContainerInformation `json:"pDUContainerInformation,omitempty"`
}

// NewUsedUnitContainer instantiates a new UsedUnitContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsedUnitContainer(localSequenceNumber int32) *UsedUnitContainer {
	this := UsedUnitContainer{}
	this.LocalSequenceNumber = localSequenceNumber
	return &this
}

// NewUsedUnitContainerWithDefaults instantiates a new UsedUnitContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsedUnitContainerWithDefaults() *UsedUnitContainer {
	this := UsedUnitContainer{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *UsedUnitContainer) GetServiceId() int32 {
	if o == nil || IsNil(o.ServiceId) {
		var ret int32
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsedUnitContainer) GetServiceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *UsedUnitContainer) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given int32 and assigns it to the ServiceId field.
func (o *UsedUnitContainer) SetServiceId(v int32) {
	o.ServiceId = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *UsedUnitContainer) GetTriggers() []Trigger {
	if o == nil || IsNil(o.Triggers) {
		var ret []Trigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsedUnitContainer) GetTriggersOk() ([]Trigger, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *UsedUnitContainer) HasTriggers() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []Trigger and assigns it to the Triggers field.
func (o *UsedUnitContainer) SetTriggers(v []Trigger) {
	o.Triggers = v
}

// GetTriggerTimestamp returns the TriggerTimestamp field value if set, zero value otherwise.
func (o *UsedUnitContainer) GetTriggerTimestamp() time.Time {
	if o == nil || IsNil(o.TriggerTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.TriggerTimestamp
}

// GetTriggerTimestampOk returns a tuple with the TriggerTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsedUnitContainer) GetTriggerTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TriggerTimestamp) {
		return nil, false
	}
	return o.TriggerTimestamp, true
}

// HasTriggerTimestamp returns a boolean if a field has been set.
func (o *UsedUnitContainer) HasTriggerTimestamp() bool {
	if o != nil && !IsNil(o.TriggerTimestamp) {
		return true
	}

	return false
}

// SetTriggerTimestamp gets a reference to the given time.Time and assigns it to the TriggerTimestamp field.
func (o *UsedUnitContainer) SetTriggerTimestamp(v time.Time) {
	o.TriggerTimestamp = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *UsedUnitContainer) GetTime() int32 {
	if o == nil || IsNil(o.Time) {
		var ret int32
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsedUnitContainer) GetTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *UsedUnitContainer) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int32 and assigns it to the Time field.
func (o *UsedUnitContainer) SetTime(v int32) {
	o.Time = &v
}

// GetTotalVolume returns the TotalVolume field value if set, zero value otherwise.
func (o *UsedUnitContainer) GetTotalVolume() int32 {
	if o == nil || IsNil(o.TotalVolume) {
		var ret int32
		return ret
	}
	return *o.TotalVolume
}

// GetTotalVolumeOk returns a tuple with the TotalVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsedUnitContainer) GetTotalVolumeOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalVolume) {
		return nil, false
	}
	return o.TotalVolume, true
}

// HasTotalVolume returns a boolean if a field has been set.
func (o *UsedUnitContainer) HasTotalVolume() bool {
	if o != nil && !IsNil(o.TotalVolume) {
		return true
	}

	return false
}

// SetTotalVolume gets a reference to the given int32 and assigns it to the TotalVolume field.
func (o *UsedUnitContainer) SetTotalVolume(v int32) {
	o.TotalVolume = &v
}

// GetUplinkVolume returns the UplinkVolume field value if set, zero value otherwise.
func (o *UsedUnitContainer) GetUplinkVolume() int32 {
	if o == nil || IsNil(o.UplinkVolume) {
		var ret int32
		return ret
	}
	return *o.UplinkVolume
}

// GetUplinkVolumeOk returns a tuple with the UplinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsedUnitContainer) GetUplinkVolumeOk() (*int32, bool) {
	if o == nil || IsNil(o.UplinkVolume) {
		return nil, false
	}
	return o.UplinkVolume, true
}

// HasUplinkVolume returns a boolean if a field has been set.
func (o *UsedUnitContainer) HasUplinkVolume() bool {
	if o != nil && !IsNil(o.UplinkVolume) {
		return true
	}

	return false
}

// SetUplinkVolume gets a reference to the given int32 and assigns it to the UplinkVolume field.
func (o *UsedUnitContainer) SetUplinkVolume(v int32) {
	o.UplinkVolume = &v
}

// GetDownlinkVolume returns the DownlinkVolume field value if set, zero value otherwise.
func (o *UsedUnitContainer) GetDownlinkVolume() int32 {
	if o == nil || IsNil(o.DownlinkVolume) {
		var ret int32
		return ret
	}
	return *o.DownlinkVolume
}

// GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsedUnitContainer) GetDownlinkVolumeOk() (*int32, bool) {
	if o == nil || IsNil(o.DownlinkVolume) {
		return nil, false
	}
	return o.DownlinkVolume, true
}

// HasDownlinkVolume returns a boolean if a field has been set.
func (o *UsedUnitContainer) HasDownlinkVolume() bool {
	if o != nil && !IsNil(o.DownlinkVolume) {
		return true
	}

	return false
}

// SetDownlinkVolume gets a reference to the given int32 and assigns it to the DownlinkVolume field.
func (o *UsedUnitContainer) SetDownlinkVolume(v int32) {
	o.DownlinkVolume = &v
}

// GetServiceSpecificUnits returns the ServiceSpecificUnits field value if set, zero value otherwise.
func (o *UsedUnitContainer) GetServiceSpecificUnits() int32 {
	if o == nil || IsNil(o.ServiceSpecificUnits) {
		var ret int32
		return ret
	}
	return *o.ServiceSpecificUnits
}

// GetServiceSpecificUnitsOk returns a tuple with the ServiceSpecificUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsedUnitContainer) GetServiceSpecificUnitsOk() (*int32, bool) {
	if o == nil || IsNil(o.ServiceSpecificUnits) {
		return nil, false
	}
	return o.ServiceSpecificUnits, true
}

// HasServiceSpecificUnits returns a boolean if a field has been set.
func (o *UsedUnitContainer) HasServiceSpecificUnits() bool {
	if o != nil && !IsNil(o.ServiceSpecificUnits) {
		return true
	}

	return false
}

// SetServiceSpecificUnits gets a reference to the given int32 and assigns it to the ServiceSpecificUnits field.
func (o *UsedUnitContainer) SetServiceSpecificUnits(v int32) {
	o.ServiceSpecificUnits = &v
}

// GetEventTimeStamps returns the EventTimeStamps field value if set, zero value otherwise.
func (o *UsedUnitContainer) GetEventTimeStamps() []time.Time {
	if o == nil || IsNil(o.EventTimeStamps) {
		var ret []time.Time
		return ret
	}
	return o.EventTimeStamps
}

// GetEventTimeStampsOk returns a tuple with the EventTimeStamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsedUnitContainer) GetEventTimeStampsOk() ([]time.Time, bool) {
	if o == nil || IsNil(o.EventTimeStamps) {
		return nil, false
	}
	return o.EventTimeStamps, true
}

// HasEventTimeStamps returns a boolean if a field has been set.
func (o *UsedUnitContainer) HasEventTimeStamps() bool {
	if o != nil && !IsNil(o.EventTimeStamps) {
		return true
	}

	return false
}

// SetEventTimeStamps gets a reference to the given []time.Time and assigns it to the EventTimeStamps field.
func (o *UsedUnitContainer) SetEventTimeStamps(v []time.Time) {
	o.EventTimeStamps = v
}

// GetLocalSequenceNumber returns the LocalSequenceNumber field value
func (o *UsedUnitContainer) GetLocalSequenceNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LocalSequenceNumber
}

// GetLocalSequenceNumberOk returns a tuple with the LocalSequenceNumber field value
// and a boolean to check if the value has been set.
func (o *UsedUnitContainer) GetLocalSequenceNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalSequenceNumber, true
}

// SetLocalSequenceNumber sets field value
func (o *UsedUnitContainer) SetLocalSequenceNumber(v int32) {
	o.LocalSequenceNumber = v
}

// GetPDUContainerInformation returns the PDUContainerInformation field value if set, zero value otherwise.
func (o *UsedUnitContainer) GetPDUContainerInformation() PDUContainerInformation {
	if o == nil || IsNil(o.PDUContainerInformation) {
		var ret PDUContainerInformation
		return ret
	}
	return *o.PDUContainerInformation
}

// GetPDUContainerInformationOk returns a tuple with the PDUContainerInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsedUnitContainer) GetPDUContainerInformationOk() (*PDUContainerInformation, bool) {
	if o == nil || IsNil(o.PDUContainerInformation) {
		return nil, false
	}
	return o.PDUContainerInformation, true
}

// HasPDUContainerInformation returns a boolean if a field has been set.
func (o *UsedUnitContainer) HasPDUContainerInformation() bool {
	if o != nil && !IsNil(o.PDUContainerInformation) {
		return true
	}

	return false
}

// SetPDUContainerInformation gets a reference to the given PDUContainerInformation and assigns it to the PDUContainerInformation field.
func (o *UsedUnitContainer) SetPDUContainerInformation(v PDUContainerInformation) {
	o.PDUContainerInformation = &v
}

func (o UsedUnitContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsedUnitContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !IsNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	if !IsNil(o.TriggerTimestamp) {
		toSerialize["triggerTimestamp"] = o.TriggerTimestamp
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.TotalVolume) {
		toSerialize["totalVolume"] = o.TotalVolume
	}
	if !IsNil(o.UplinkVolume) {
		toSerialize["uplinkVolume"] = o.UplinkVolume
	}
	if !IsNil(o.DownlinkVolume) {
		toSerialize["downlinkVolume"] = o.DownlinkVolume
	}
	if !IsNil(o.ServiceSpecificUnits) {
		toSerialize["serviceSpecificUnits"] = o.ServiceSpecificUnits
	}
	if !IsNil(o.EventTimeStamps) {
		toSerialize["eventTimeStamps"] = o.EventTimeStamps
	}
	toSerialize["localSequenceNumber"] = o.LocalSequenceNumber
	if !IsNil(o.PDUContainerInformation) {
		toSerialize["pDUContainerInformation"] = o.PDUContainerInformation
	}
	return toSerialize, nil
}

type NullableUsedUnitContainer struct {
	value *UsedUnitContainer
	isSet bool
}

func (v NullableUsedUnitContainer) Get() *UsedUnitContainer {
	return v.value
}

func (v *NullableUsedUnitContainer) Set(val *UsedUnitContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableUsedUnitContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableUsedUnitContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsedUnitContainer(val *UsedUnitContainer) *NullableUsedUnitContainer {
	return &NullableUsedUnitContainer{value: val, isSet: true}
}

func (v NullableUsedUnitContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsedUnitContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

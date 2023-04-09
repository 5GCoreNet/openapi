/*
3gpp-nidd

API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NIDD

import (
	"encoding/json"
)

// checks if the NiddDownlinkDataTransferPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiddDownlinkDataTransferPatch{}

// NiddDownlinkDataTransferPatch Represents the parameters to request the modification of an Individual NIDD Downlink Data Delivery resource.
type NiddDownlinkDataTransferPatch struct {
	// String with format \"byte\" as defined in OpenAPI Specification, i.e, base64-encoded characters.
	Data *string `json:"data,omitempty"`
	// Indicates whether the reliable data service (as defined in clause 4.5.14.3 of 3GPP TS  23.682) acknowledgement is requested (true) or not (false).
	ReliableDataService *bool    `json:"reliableDataService,omitempty"`
	RdsPort             *RdsPort `json:"rdsPort,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	MaximumLatency *int32 `json:"maximumLatency,omitempty"`
	// It is used to indicate the priority of the non-IP data packet relative to other non-IP data packets.
	Priority               *int32                   `json:"priority,omitempty"`
	PdnEstablishmentOption *PdnEstablishmentOptions `json:"pdnEstablishmentOption,omitempty"`
}

// NewNiddDownlinkDataTransferPatch instantiates a new NiddDownlinkDataTransferPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiddDownlinkDataTransferPatch() *NiddDownlinkDataTransferPatch {
	this := NiddDownlinkDataTransferPatch{}
	return &this
}

// NewNiddDownlinkDataTransferPatchWithDefaults instantiates a new NiddDownlinkDataTransferPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiddDownlinkDataTransferPatchWithDefaults() *NiddDownlinkDataTransferPatch {
	this := NiddDownlinkDataTransferPatch{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *NiddDownlinkDataTransferPatch) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiddDownlinkDataTransferPatch) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *NiddDownlinkDataTransferPatch) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *NiddDownlinkDataTransferPatch) SetData(v string) {
	o.Data = &v
}

// GetReliableDataService returns the ReliableDataService field value if set, zero value otherwise.
func (o *NiddDownlinkDataTransferPatch) GetReliableDataService() bool {
	if o == nil || IsNil(o.ReliableDataService) {
		var ret bool
		return ret
	}
	return *o.ReliableDataService
}

// GetReliableDataServiceOk returns a tuple with the ReliableDataService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiddDownlinkDataTransferPatch) GetReliableDataServiceOk() (*bool, bool) {
	if o == nil || IsNil(o.ReliableDataService) {
		return nil, false
	}
	return o.ReliableDataService, true
}

// HasReliableDataService returns a boolean if a field has been set.
func (o *NiddDownlinkDataTransferPatch) HasReliableDataService() bool {
	if o != nil && !IsNil(o.ReliableDataService) {
		return true
	}

	return false
}

// SetReliableDataService gets a reference to the given bool and assigns it to the ReliableDataService field.
func (o *NiddDownlinkDataTransferPatch) SetReliableDataService(v bool) {
	o.ReliableDataService = &v
}

// GetRdsPort returns the RdsPort field value if set, zero value otherwise.
func (o *NiddDownlinkDataTransferPatch) GetRdsPort() RdsPort {
	if o == nil || IsNil(o.RdsPort) {
		var ret RdsPort
		return ret
	}
	return *o.RdsPort
}

// GetRdsPortOk returns a tuple with the RdsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiddDownlinkDataTransferPatch) GetRdsPortOk() (*RdsPort, bool) {
	if o == nil || IsNil(o.RdsPort) {
		return nil, false
	}
	return o.RdsPort, true
}

// HasRdsPort returns a boolean if a field has been set.
func (o *NiddDownlinkDataTransferPatch) HasRdsPort() bool {
	if o != nil && !IsNil(o.RdsPort) {
		return true
	}

	return false
}

// SetRdsPort gets a reference to the given RdsPort and assigns it to the RdsPort field.
func (o *NiddDownlinkDataTransferPatch) SetRdsPort(v RdsPort) {
	o.RdsPort = &v
}

// GetMaximumLatency returns the MaximumLatency field value if set, zero value otherwise.
func (o *NiddDownlinkDataTransferPatch) GetMaximumLatency() int32 {
	if o == nil || IsNil(o.MaximumLatency) {
		var ret int32
		return ret
	}
	return *o.MaximumLatency
}

// GetMaximumLatencyOk returns a tuple with the MaximumLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiddDownlinkDataTransferPatch) GetMaximumLatencyOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumLatency) {
		return nil, false
	}
	return o.MaximumLatency, true
}

// HasMaximumLatency returns a boolean if a field has been set.
func (o *NiddDownlinkDataTransferPatch) HasMaximumLatency() bool {
	if o != nil && !IsNil(o.MaximumLatency) {
		return true
	}

	return false
}

// SetMaximumLatency gets a reference to the given int32 and assigns it to the MaximumLatency field.
func (o *NiddDownlinkDataTransferPatch) SetMaximumLatency(v int32) {
	o.MaximumLatency = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *NiddDownlinkDataTransferPatch) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiddDownlinkDataTransferPatch) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *NiddDownlinkDataTransferPatch) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *NiddDownlinkDataTransferPatch) SetPriority(v int32) {
	o.Priority = &v
}

// GetPdnEstablishmentOption returns the PdnEstablishmentOption field value if set, zero value otherwise.
func (o *NiddDownlinkDataTransferPatch) GetPdnEstablishmentOption() PdnEstablishmentOptions {
	if o == nil || IsNil(o.PdnEstablishmentOption) {
		var ret PdnEstablishmentOptions
		return ret
	}
	return *o.PdnEstablishmentOption
}

// GetPdnEstablishmentOptionOk returns a tuple with the PdnEstablishmentOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiddDownlinkDataTransferPatch) GetPdnEstablishmentOptionOk() (*PdnEstablishmentOptions, bool) {
	if o == nil || IsNil(o.PdnEstablishmentOption) {
		return nil, false
	}
	return o.PdnEstablishmentOption, true
}

// HasPdnEstablishmentOption returns a boolean if a field has been set.
func (o *NiddDownlinkDataTransferPatch) HasPdnEstablishmentOption() bool {
	if o != nil && !IsNil(o.PdnEstablishmentOption) {
		return true
	}

	return false
}

// SetPdnEstablishmentOption gets a reference to the given PdnEstablishmentOptions and assigns it to the PdnEstablishmentOption field.
func (o *NiddDownlinkDataTransferPatch) SetPdnEstablishmentOption(v PdnEstablishmentOptions) {
	o.PdnEstablishmentOption = &v
}

func (o NiddDownlinkDataTransferPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiddDownlinkDataTransferPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.ReliableDataService) {
		toSerialize["reliableDataService"] = o.ReliableDataService
	}
	if !IsNil(o.RdsPort) {
		toSerialize["rdsPort"] = o.RdsPort
	}
	if !IsNil(o.MaximumLatency) {
		toSerialize["maximumLatency"] = o.MaximumLatency
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.PdnEstablishmentOption) {
		toSerialize["pdnEstablishmentOption"] = o.PdnEstablishmentOption
	}
	return toSerialize, nil
}

type NullableNiddDownlinkDataTransferPatch struct {
	value *NiddDownlinkDataTransferPatch
	isSet bool
}

func (v NullableNiddDownlinkDataTransferPatch) Get() *NiddDownlinkDataTransferPatch {
	return v.value
}

func (v *NullableNiddDownlinkDataTransferPatch) Set(val *NiddDownlinkDataTransferPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableNiddDownlinkDataTransferPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableNiddDownlinkDataTransferPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiddDownlinkDataTransferPatch(val *NiddDownlinkDataTransferPatch) *NullableNiddDownlinkDataTransferPatch {
	return &NullableNiddDownlinkDataTransferPatch{value: val, isSet: true}
}

func (v NullableNiddDownlinkDataTransferPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiddDownlinkDataTransferPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

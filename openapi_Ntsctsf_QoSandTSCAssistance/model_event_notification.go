/*
Ntsctsf_QoSandTSCAssistance Service API

TSCTSF QoS and TSC Assistance Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_QoSandTSCAssistance

import (
	"encoding/json"
)

// checks if the EventNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventNotification{}

// EventNotification Describes a notification of an matched event.
type EventNotification struct {
	Event TscEvent `json:"event"`
	// Identifies the IP flows that were sent during event subscription.
	FlowIds []int32 `json:"flowIds,omitempty"`
	QosMonReports []QosMonitoringReport `json:"qosMonReports,omitempty"`
	UsgRep *AccumulatedUsage `json:"usgRep,omitempty"`
	// The currently applied alternative QoS requirement referring to an alternative QoS reference or a requested alternative QoS parameter set. Applicable for event QOS_NOT_GUARANTEED or SUCCESSFUL_RESOURCES_ALLOCATION. 
	AppliedQosRef *string `json:"appliedQosRef,omitempty"`
}

// NewEventNotification instantiates a new EventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventNotification(event TscEvent) *EventNotification {
	this := EventNotification{}
	this.Event = event
	return &this
}

// NewEventNotificationWithDefaults instantiates a new EventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventNotificationWithDefaults() *EventNotification {
	this := EventNotification{}
	return &this
}

// GetEvent returns the Event field value
func (o *EventNotification) GetEvent() TscEvent {
	if o == nil {
		var ret TscEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *EventNotification) GetEventOk() (*TscEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *EventNotification) SetEvent(v TscEvent) {
	o.Event = v
}

// GetFlowIds returns the FlowIds field value if set, zero value otherwise.
func (o *EventNotification) GetFlowIds() []int32 {
	if o == nil || isNil(o.FlowIds) {
		var ret []int32
		return ret
	}
	return o.FlowIds
}

// GetFlowIdsOk returns a tuple with the FlowIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetFlowIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.FlowIds) {
		return nil, false
	}
	return o.FlowIds, true
}

// HasFlowIds returns a boolean if a field has been set.
func (o *EventNotification) HasFlowIds() bool {
	if o != nil && !isNil(o.FlowIds) {
		return true
	}

	return false
}

// SetFlowIds gets a reference to the given []int32 and assigns it to the FlowIds field.
func (o *EventNotification) SetFlowIds(v []int32) {
	o.FlowIds = v
}

// GetQosMonReports returns the QosMonReports field value if set, zero value otherwise.
func (o *EventNotification) GetQosMonReports() []QosMonitoringReport {
	if o == nil || isNil(o.QosMonReports) {
		var ret []QosMonitoringReport
		return ret
	}
	return o.QosMonReports
}

// GetQosMonReportsOk returns a tuple with the QosMonReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetQosMonReportsOk() ([]QosMonitoringReport, bool) {
	if o == nil || isNil(o.QosMonReports) {
		return nil, false
	}
	return o.QosMonReports, true
}

// HasQosMonReports returns a boolean if a field has been set.
func (o *EventNotification) HasQosMonReports() bool {
	if o != nil && !isNil(o.QosMonReports) {
		return true
	}

	return false
}

// SetQosMonReports gets a reference to the given []QosMonitoringReport and assigns it to the QosMonReports field.
func (o *EventNotification) SetQosMonReports(v []QosMonitoringReport) {
	o.QosMonReports = v
}

// GetUsgRep returns the UsgRep field value if set, zero value otherwise.
func (o *EventNotification) GetUsgRep() AccumulatedUsage {
	if o == nil || isNil(o.UsgRep) {
		var ret AccumulatedUsage
		return ret
	}
	return *o.UsgRep
}

// GetUsgRepOk returns a tuple with the UsgRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetUsgRepOk() (*AccumulatedUsage, bool) {
	if o == nil || isNil(o.UsgRep) {
		return nil, false
	}
	return o.UsgRep, true
}

// HasUsgRep returns a boolean if a field has been set.
func (o *EventNotification) HasUsgRep() bool {
	if o != nil && !isNil(o.UsgRep) {
		return true
	}

	return false
}

// SetUsgRep gets a reference to the given AccumulatedUsage and assigns it to the UsgRep field.
func (o *EventNotification) SetUsgRep(v AccumulatedUsage) {
	o.UsgRep = &v
}

// GetAppliedQosRef returns the AppliedQosRef field value if set, zero value otherwise.
func (o *EventNotification) GetAppliedQosRef() string {
	if o == nil || isNil(o.AppliedQosRef) {
		var ret string
		return ret
	}
	return *o.AppliedQosRef
}

// GetAppliedQosRefOk returns a tuple with the AppliedQosRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetAppliedQosRefOk() (*string, bool) {
	if o == nil || isNil(o.AppliedQosRef) {
		return nil, false
	}
	return o.AppliedQosRef, true
}

// HasAppliedQosRef returns a boolean if a field has been set.
func (o *EventNotification) HasAppliedQosRef() bool {
	if o != nil && !isNil(o.AppliedQosRef) {
		return true
	}

	return false
}

// SetAppliedQosRef gets a reference to the given string and assigns it to the AppliedQosRef field.
func (o *EventNotification) SetAppliedQosRef(v string) {
	o.AppliedQosRef = &v
}

func (o EventNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if !isNil(o.FlowIds) {
		toSerialize["flowIds"] = o.FlowIds
	}
	if !isNil(o.QosMonReports) {
		toSerialize["qosMonReports"] = o.QosMonReports
	}
	if !isNil(o.UsgRep) {
		toSerialize["usgRep"] = o.UsgRep
	}
	if !isNil(o.AppliedQosRef) {
		toSerialize["appliedQosRef"] = o.AppliedQosRef
	}
	return toSerialize, nil
}

type NullableEventNotification struct {
	value *EventNotification
	isSet bool
}

func (v NullableEventNotification) Get() *EventNotification {
	return v.value
}

func (v *NullableEventNotification) Set(val *EventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventNotification(val *EventNotification) *NullableEventNotification {
	return &NullableEventNotification{value: val, isSet: true}
}

func (v NullableEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


